package shell

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/user/oraclepack/internal/dispatch"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
)

// Engine executes pack steps in headless mode.
type Engine struct {
	Pack       *pack.Pack
	State      *state.RunState
	StatePath  string
	StopOnFail bool
	Timeout    time.Duration
	Checker    tools.PresenceChecker
}

// Run executes all steps sequentially, updating state on disk.
func (e *Engine) Run(ctx context.Context) error {
	if e.Pack == nil {
		return nil
	}
	if e.State == nil {
		e.State = &state.RunState{
			SchemaVersion: 1,
			StepStatuses:  make(map[string]state.StepStatus),
		}
	}
	if e.State.StepStatuses == nil {
		e.State.StepStatuses = make(map[string]state.StepStatus)
	}

	for i := range e.Pack.Steps {
		step := e.Pack.Steps[i]
		e.State.CurrentStep = step.Number
		status := state.StepStatus{Status: state.StatusRunning, StartedAt: time.Now()}
		e.State.StepStatuses[step.ID] = status
		_ = state.WriteState(e.StatePath, e.State)

		kind := detectToolKind(&step)
		if shouldSkipForMissingTool(kind, e.Checker) {
			status.Status = state.StatusSkipped
			status.Error = "tool missing"
			status.EndedAt = time.Now()
			e.State.StepStatuses[step.ID] = status
			_ = state.WriteState(e.StatePath, e.State)
			continue
		}

		stepCtx := ctx
		if e.Timeout > 0 {
			var cancel context.CancelFunc
			stepCtx, cancel = context.WithTimeout(ctx, e.Timeout)
			defer cancel()
		}

		res, err := RunCommand(stepCtx, step.Code)
		if err == nil && res.ExitCode != 0 {
			err = fmt.Errorf("command failed with exit code %d", res.ExitCode)
		}
		status.EndedAt = time.Now()
		if err != nil {
			status.Status = state.StatusFailed
			status.Error = err.Error()
			e.State.StepStatuses[step.ID] = status
			_ = state.WriteState(e.StatePath, e.State)
			if e.StopOnFail {
				return err
			}
			continue
		}
		status.Status = state.StatusSuccess
		e.State.StepStatuses[step.ID] = status
		_ = state.WriteState(e.StatePath, e.State)
	}
	return nil
}

func detectToolKind(step *pack.Step) tools.ToolKind {
	if step == nil {
		return tools.ToolUnknown
	}
	lines := strings.Split(step.Code, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if cls, ok := dispatch.Classify(trimmed); ok {
			return cls.Kind
		}
		break
	}
	return tools.ToolUnknown
}

func shouldSkipForMissingTool(kind tools.ToolKind, checker tools.PresenceChecker) bool {
	if kind == tools.ToolUnknown {
		return false
	}
	meta, ok := tools.Metadata(kind)
	if !ok {
		return false
	}
	if checker == nil {
		_, found := DetectBinary(meta.Name)
		return !found
	}
	_, found := checker.DetectBinary(meta.Name)
	return !found
}
