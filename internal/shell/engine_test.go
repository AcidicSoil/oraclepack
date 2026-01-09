package shell

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

type fakeChecker struct {
	found map[string]bool
}

func (f fakeChecker) DetectBinary(name string) (string, bool) {
	return name, f.found[name]
}

func TestEngineSkipsMissingTool(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "codex exec \"hi\""},
		},
	}
	dir := t.TempDir()
	engine := &Engine{
		Pack:      p,
		State:     &state.RunState{SchemaVersion: 1, StepStatuses: map[string]state.StepStatus{}},
		StatePath: filepath.Join(dir, "state.json"),
		Checker:   fakeChecker{found: map[string]bool{"codex": false}},
	}
	if err := engine.Run(context.Background()); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if engine.State.StepStatuses["01"].Status != state.StatusSkipped {
		t.Fatalf("expected skipped, got %s", engine.State.StepStatuses["01"].Status)
	}
}

func TestEngineFailsOnError(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "exit 1"},
		},
	}
	engine := &Engine{
		Pack:       p,
		State:      &state.RunState{SchemaVersion: 1, StepStatuses: map[string]state.StepStatus{}},
		StopOnFail: true,
		Checker:    fakeChecker{found: map[string]bool{}},
	}
	if err := engine.Run(context.Background()); err == nil {
		t.Fatal("expected error, got nil")
	}
	if engine.State.StepStatuses["01"].Status != state.StatusFailed {
		t.Fatalf("expected failed, got %s", engine.State.StepStatuses["01"].Status)
	}
}
