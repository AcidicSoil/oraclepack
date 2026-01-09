package validate

import (
	"errors"

	"github.com/user/oraclepack/internal/artifacts"
	"github.com/user/oraclepack/internal/foundation"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
	"github.com/user/oraclepack/internal/tools"
)

// ArtifactGateValidator checks expected artifacts after a step.
type ArtifactGateValidator struct {
	Contract artifacts.Contract
}

func (v ArtifactGateValidator) Validate(step *pack.Step, kind tools.ToolKind, toolPresent bool) (state.Status, string) {
	if step == nil {
		return state.StatusSuccess, ""
	}
	contract := v.Contract
	if contract == nil {
		contract = artifacts.DefaultContract()
	}
	if !toolPresent {
		if _, ok := contract[step.ID]; ok {
			return state.StatusSkipped, "tool missing; skipping artifact gate"
		}
		return state.StatusSuccess, ""
	}
	err := artifacts.EvaluateGates(step.ID, contract)
	if err == nil {
		return state.StatusSuccess, ""
	}
	if errors.Is(err, foundation.ErrArtifactMissing) {
		return state.StatusFailed, err.Error()
	}
	return state.StatusFailed, err.Error()
}
