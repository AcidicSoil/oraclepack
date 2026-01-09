package validate

import (
	"strings"

	"github.com/user/oraclepack/internal/dispatch"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// StepResult captures validation output for a step.
type StepResult struct {
	StepID   string
	ToolKind tools.ToolKind
	Status   state.Status
	Error    string
}

// DetectToolKind scans a step for a known tool prefix.
func DetectToolKind(step *types.Step) tools.ToolKind {
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
