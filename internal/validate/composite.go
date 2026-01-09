package validate

import (
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

// CompositeValidator coordinates multiple validators for all steps.
type CompositeValidator struct {
	ToolPresence ToolPresenceValidator
	OracleDryRun OracleDryRunValidator
	ArtifactGate ArtifactGateValidator
}

// ValidatePack validates all steps in a pack and returns per-step results.
func (v CompositeValidator) ValidatePack(p *pack.Pack) []StepResult {
	if p == nil {
		return nil
	}
	results := make([]StepResult, 0, len(p.Steps))
	for i := range p.Steps {
		step := &p.Steps[i]
		kind := DetectToolKind(step)

		status, reason, toolPresent := v.ToolPresence.Validate(step, kind)
		if status == state.StatusSkipped {
			results = append(results, StepResult{
				StepID:   step.ID,
				ToolKind: kind,
				Status:   status,
				Error:    reason,
			})
			continue
		}

		if oracleStatus, oracleErr := v.OracleDryRun.Validate(step, kind); oracleStatus == state.StatusFailed {
			results = append(results, StepResult{
				StepID:   step.ID,
				ToolKind: kind,
				Status:   oracleStatus,
				Error:    oracleErr,
			})
			continue
		}

		gateStatus, gateErr := v.ArtifactGate.Validate(step, kind, toolPresent)
		finalStatus := gateStatus
		errMsg := gateErr
		if finalStatus == "" {
			finalStatus = state.StatusSuccess
		}
		results = append(results, StepResult{
			StepID:   step.ID,
			ToolKind: kind,
			Status:   finalStatus,
			Error:    errMsg,
		})
	}
	return results
}
