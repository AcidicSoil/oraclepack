package tui

import (
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestStepPreviewContentUnwrapped(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", OriginalLine: "Step 1", Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}
	m := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")
	m.width = 80
	m.previewID = "01"
	m.previewWrap = false

	content := m.stepPreviewContent()
	if !strings.Contains(content, "Step 01") {
		t.Fatalf("expected header to include step id, got %q", content)
	}
	if !strings.Contains(content, "echo hello") {
		t.Fatalf("expected content to include code, got %q", content)
	}
}
