package validate

import (
	"fmt"

	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/shell"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
)

// ToolPresenceValidator checks that tool binaries exist on PATH.
type ToolPresenceValidator struct {
	Checker tools.PresenceChecker
}

// Validate returns skipped if the tool is missing.
func (v ToolPresenceValidator) Validate(step *pack.Step, kind tools.ToolKind) (state.Status, string, bool) {
	if kind == tools.ToolUnknown {
		return state.StatusSuccess, "", true
	}
	meta, ok := tools.Metadata(kind)
	if !ok {
		return state.StatusSuccess, "", true
	}
	checker := v.Checker
	if checker == nil {
		checker = shellChecker{}
	}
	if _, found := checker.DetectBinary(meta.Name); !found {
		return state.StatusSkipped, fmt.Sprintf("tool %s not found on PATH", meta.Name), false
	}
	return state.StatusSuccess, "", true
}

type shellChecker struct{}

func (shellChecker) DetectBinary(name string) (string, bool) {
	return shell.DetectBinary(name)
}
