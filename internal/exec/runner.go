package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

// Runner handles the execution of shell scripts.
type Runner struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
	warnings    []SanitizeWarning
}

// RunnerOptions configures a Runner.
type RunnerOptions struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
}

// NewRunner creates a new Runner with options.
func NewRunner(opts RunnerOptions) *Runner {
	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}

	return &Runner{
		Shell:       shell,
		WorkDir:     opts.WorkDir,
		Env:         append(os.Environ(), opts.Env...),
		OracleFlags: opts.OracleFlags,
		Overrides:   opts.Overrides,
		ChatGPTURL:  opts.ChatGPTURL,
	}
}

// RunPrelude executes the prelude code.
func (r *Runner) RunPrelude(ctx context.Context, p *pack.Prelude, logWriter io.Writer) error {
	script, warnings := SanitizeScript(p.Code, "prelude", "")
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

// RunStep executes a single step's code.
func (r *Runner) RunStep(ctx context.Context, s *pack.Step, logWriter io.Writer) error {
	flags := ApplyChatGPTURL(r.OracleFlags, r.ChatGPTURL)
	if r.Overrides != nil {
		flags = r.Overrides.EffectiveFlags(s.ID, r.OracleFlags)
		flags = ApplyChatGPTURL(flags, r.ChatGPTURL)
	}
	code := InjectFlags(s.Code, flags)
	script, warnings := SanitizeScript(code, "step", s.ID)
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

func (r *Runner) recordWarnings(warnings []SanitizeWarning, logWriter io.Writer) {
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		r.warnings = append(r.warnings, w)
		if logWriter != nil {
			scope := w.Scope
			if scope == "" {
				scope = "script"
			}
			step := ""
			if w.StepID != "" {
				step = " step " + w.StepID
			}
			_, _ = fmt.Fprintf(logWriter, "⚠ oraclepack: sanitized label in %s%s line %d: %s\n", scope, step, w.Line, w.Token)
		}
	}
}

// DrainWarnings returns any sanitizer warnings collected since the last call.
func (r *Runner) DrainWarnings() []SanitizeWarning {
	if len(r.warnings) == 0 {
		return nil
	}
	out := make([]SanitizeWarning, len(r.warnings))
	copy(out, r.warnings)
	r.warnings = nil
	return out
}

func (r *Runner) run(ctx context.Context, script string, logWriter io.Writer) error {
	// We use bash -lc to ensure login shell (paths, aliases, etc)
	cmd := exec.CommandContext(ctx, r.Shell, "-lc", script)
	cmd.Dir = r.WorkDir
	cmd.Env = r.Env

	// Standardize stdout and stderr to the logWriter
	cmd.Stdout = logWriter
	cmd.Stderr = logWriter

	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err)
	}

	return nil
}
