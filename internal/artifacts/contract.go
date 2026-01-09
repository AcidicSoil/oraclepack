package artifacts

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/user/oraclepack/internal/foundation"
)

// Contract maps step IDs to required artifact paths.
type Contract map[string][]string

// DefaultContract returns the standard artifact contract.
func DefaultContract() Contract {
	base := ".oraclepack/ticketify"
	return Contract{
		"09": {filepath.Join(base, "next.json")},
		"10": {filepath.Join(base, "codex-implement.md")},
		"11": {filepath.Join(base, "codex-verify.md")},
		"12": {filepath.Join(base, "PR.md")},
	}
}

// EvaluateGates checks required artifacts for a given step.
func EvaluateGates(stepID string, contract Contract) error {
	paths, ok := contract[stepID]
	if !ok || len(paths) == 0 {
		return nil
	}
	var missing []string
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil || info.IsDir() || info.Size() == 0 {
			missing = append(missing, p)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("%w: %v", foundation.ErrArtifactMissing, missing)
	}
	return nil
}
