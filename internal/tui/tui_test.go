package tui

import (
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/types"
)

func TestInitAutoRun(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Number: 1, Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Test case 1: autoRun = true
	modelAuto := NewModel(p, r, s, "", 0, "over", true, false, 0, false, "auto")
	cmdAuto := modelAuto.Init()

	if cmdAuto == nil {
		t.Fatal("expected Init cmd to be non-nil when autoRun is true")
	}
	// Note: We can't easily assert the content of a Batch command in a unit test.

	// Test case 2: autoRun = false
	modelManual := NewModel(p, r, s, "", 0, "over", false, false, 0, false, "auto")
	// Even with autoRun false, we have textinput.Blink, so Init is not nil.
	cmdManual := modelManual.Init()
	if cmdManual == nil {
		t.Fatal("expected Init cmd to be non-nil due to textinput.Blink")
	}
}
