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
}

// RunnerOptions configures a Runner.
type RunnerOptions struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
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
	}
}

// RunPrelude executes the prelude code.
func (r *Runner) RunPrelude(ctx context.Context, p *pack.Prelude, logWriter io.Writer) error {
	return r.run(ctx, p.Code, logWriter)
}

// RunStep executes a single step's code.
func (r *Runner) RunStep(ctx context.Context, s *pack.Step, logWriter io.Writer) error {
	flags := r.OracleFlags
	if r.Overrides != nil {
		flags = r.Overrides.EffectiveFlags(s.ID, r.OracleFlags)
	}
	code := InjectFlags(s.Code, flags)
	return r.run(ctx, code, logWriter)
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
