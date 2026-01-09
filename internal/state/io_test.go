package state

import (
	"path/filepath"
	"testing"
	"time"
)

func TestWriteStateAndLoadState(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	input := &RunState{
		SchemaVersion: 1,
		PackHash:      "hash",
		StartTime:     time.Now(),
		CurrentStep:   2,
		StepStatuses: map[string]StepStatus{
			"01": {Status: StatusSuccess},
		},
	}

	if err := WriteState(path, input); err != nil {
		t.Fatalf("WriteState: %v", err)
	}

	out, err := LoadState(path)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}

	if out.CurrentStep != 2 {
		t.Fatalf("expected CurrentStep 2, got %d", out.CurrentStep)
	}
	if out.StepStatuses["01"].Status != StatusSuccess {
		t.Fatalf("expected status success, got %s", out.StepStatuses["01"].Status)
	}
}
