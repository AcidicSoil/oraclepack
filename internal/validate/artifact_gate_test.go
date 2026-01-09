package validate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/artifacts"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
)

func TestArtifactGateValidator(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "next.json")
	contract := artifacts.Contract{"09": {path}}
	step := &pack.Step{ID: "09"}
	v := ArtifactGateValidator{Contract: contract}

	status, _ := v.Validate(step, tools.ToolCodex, true)
	if status != state.StatusFailed {
		t.Fatalf("expected failed for missing artifact, got %s", status)
	}

	if err := os.WriteFile(path, []byte("ok"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	status, _ = v.Validate(step, tools.ToolCodex, true)
	if status != state.StatusSuccess {
		t.Fatalf("expected success, got %s", status)
	}

	status, _ = v.Validate(step, tools.ToolCodex, false)
	if status != state.StatusSkipped {
		t.Fatalf("expected skipped when tool missing, got %s", status)
	}
}
