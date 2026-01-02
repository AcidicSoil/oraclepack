package tui

import (
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

func TestStepPreviewContentUnwrapped(t *testing.T) {
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", OriginalLine: "Step 1", Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}
	m := NewModel(p, r, s, "", 0, "over", false)
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
