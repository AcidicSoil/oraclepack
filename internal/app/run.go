package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/report"
	"github.com/user/oraclepack/internal/state"
)

func (a *App) RunPlain(ctx context.Context, out io.Writer) error {
	// Assumes a.Prepare() and a.LoadState() have been called by the CLI entrypoint.
	if a.Pack == nil {
		return fmt.Errorf("pack not loaded")
	}
	if a.State == nil {
		return fmt.Errorf("state not loaded")
	}

	if a.State.StartTime.IsZero() {
		a.State.StartTime = time.Now()
	}

	fmt.Fprintf(out, "Running pack: %s\n", a.Config.PackPath)
	fmt.Fprintf(out, "Output directory: %s\n", a.Runner.WorkDir)

	// Prelude
	if a.Pack.Prelude.Code != "" {
		fmt.Fprintln(out, "Executing prelude...")
		err := a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out)
		a.recordWarnings()
		if err != nil {
			return fmt.Errorf("prelude failed: %w", err)
		}
	}

	for _, step := range a.Pack.Steps {
		// Filter by ROI
		if a.Config.ROIThreshold > 0 {
			if a.Config.ROIMode == "under" {
				// "under" is strictly less than
				if step.ROI >= a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f >= %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			} else {
				// "over" is greater than or equal to (3.3 or higher)
				if step.ROI < a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f < %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			}
		}

		// Check resume
		if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess {
			fmt.Fprintf(out, "Skipping step %s (already succeeded)\n", step.ID)
			continue
		}

		fmt.Fprintf(out, "\n>>> Step %s: %s\n", step.ID, step.OriginalLine)

		status := state.StepStatus{
			Status:    state.StatusRunning,
			StartedAt: time.Now(),
		}
		a.State.StepStatuses[step.ID] = status
		a.saveState()

		// Execute
		err := a.runStepWithOutputVerification(ctx, &step, out)
		a.recordWarnings()

		status.EndedAt = time.Now()
		if err != nil {
			status.Status = state.StatusFailed
			status.Error = err.Error()
			a.State.StepStatuses[step.ID] = status
			a.saveState()

			if a.Config.StopOnFail {
				a.finalize(out)
				return err
			}
			continue
		}

		status.Status = state.StatusSuccess
		status.ExitCode = 0
		a.State.StepStatuses[step.ID] = status
		a.saveState()
	}

	a.finalize(out)
	return nil
}

func (a *App) runStepWithOutputVerification(ctx context.Context, step *pack.Step, out io.Writer) error {
	retries := a.Config.OutputRetries
	if retries < 0 {
		retries = 0
	}
	for attempt := 0; attempt <= retries; attempt++ {
		err := a.Runner.RunStep(ctx, step, out)
		if err != nil {
			return err
		}
		if !a.Config.OutputVerify {
			return nil
		}
		expectations := pack.StepOutputExpectations(step)
		if len(expectations) == 0 {
			return nil
		}
		var failures []string
		for path, required := range expectations {
			ok, missing, err := pack.ValidateOutputFile(path, required)
			if err != nil {
				return fmt.Errorf("output verification failed for step %s: %w", step.ID, err)
			}
			if !ok {
				failures = append(failures, fmt.Sprintf("%s missing: %s", path, strings.Join(missing, ", ")))
			}
		}
		if len(failures) == 0 {
			return nil
		}
		if attempt == retries {
			return fmt.Errorf(
				"output verification failed for step %s: %s",
				step.ID,
				strings.Join(failures, "; "),
			)
		}
		fmt.Fprintf(out, "⚠ output verification failed for step %s (%s); re-running (%d/%d)...\n",
			step.ID, strings.Join(failures, "; "), attempt+1, retries)
	}
	return nil
}

func (a *App) recordWarnings() {
	if a.State == nil || a.Runner == nil {
		return
	}
	warnings := a.Runner.DrainWarnings()
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		a.State.Warnings = append(a.State.Warnings, state.Warning{
			Scope:   w.Scope,
			StepID:  w.StepID,
			Line:    w.Line,
			Token:   w.Token,
			Message: w.Message,
		})
	}
	a.saveState()
}

func (a *App) saveState() {
	if a.Config.StatePath != "" {
		_ = state.SaveStateAtomic(a.Config.StatePath, a.State)
	}
}

func (a *App) finalize(out io.Writer) {
	if a.Config.ReportPath != "" {
		rep := report.GenerateReport(a.State, filepath.Base(a.Config.PackPath))
		data, _ := json.MarshalIndent(rep, "", "  ")
		_ = os.WriteFile(a.Config.ReportPath, data, 0644)
		fmt.Fprintf(out, "\nReport written to %s\n", a.Config.ReportPath)
	}
}
