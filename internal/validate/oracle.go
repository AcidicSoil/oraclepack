package validate

import (
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
	"github.com/user/oraclepack/internal/types"
)

// OracleDryRunValidator is a placeholder for oracle-specific validation.
type OracleDryRunValidator struct{}

// Validate returns success for non-oracle tools and no-op for oracle until wired.
func (OracleDryRunValidator) Validate(step *types.Step, kind tools.ToolKind) (state.Status, string) {
	if kind != tools.ToolOracle {
		return state.StatusSuccess, ""
	}
	return state.StatusSuccess, ""
}
