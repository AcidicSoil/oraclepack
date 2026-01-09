package validate

import (
	"testing"

	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

type fakeChecker struct {
	found map[string]bool
}

func (f fakeChecker) DetectBinary(name string) (string, bool) {
	return name, f.found[name]
}

func TestToolPresenceValidator(t *testing.T) {
	step := &types.Step{ID: "01", Code: "codex exec \"hi\""}
	v := ToolPresenceValidator{Checker: fakeChecker{found: map[string]bool{"codex": false}}}
	status, reason, present := v.Validate(step, tools.ToolCodex)
	if status != state.StatusSkipped || present {
		t.Fatalf("expected skipped/missing, got status=%s present=%v", status, present)
	}
	if reason == "" {
		t.Fatalf("expected reason for missing tool")
	}
}
