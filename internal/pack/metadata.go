package pack

import (
	"fmt"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/types"
)

// DeriveMetadata extracts configuration from the prelude.
func DeriveMetadata(p *types.Pack) {
	if p == nil {
		return
	}
	outDirMatch := outDirRegex.FindStringSubmatch(p.Prelude.Code)
	if len(outDirMatch) > 1 {
		p.OutDir = outDirMatch[1]
	}

	if writeOutputRegex.MatchString(p.Prelude.Code) {
		p.WriteOutput = true
	}
}

// Validate checks if the pack follows all rules.
func Validate(p *types.Pack) error {
	if p == nil {
		return fmt.Errorf("%w: pack is nil", errors.ErrInvalidPack)
	}
	if len(p.Steps) == 0 {
		return fmt.Errorf("%w: at least one step is required", errors.ErrInvalidPack)
	}
	if len(p.Steps) != 20 {
		return fmt.Errorf("%w: expected exactly 20 steps, got %d", errors.ErrInvalidPack, len(p.Steps))
	}

	seen := make(map[int]bool)
	for i, step := range p.Steps {
		if step.Number <= 0 {
			return fmt.Errorf("%w: invalid step number %d", errors.ErrInvalidPack, step.Number)
		}
		if seen[step.Number] {
			return fmt.Errorf("%w: duplicate step number %d", errors.ErrInvalidPack, step.Number)
		}
		seen[step.Number] = true

		// Ensure sequential starting from 1
		if step.Number != i+1 {
			return fmt.Errorf("%w: steps must be sequential starting from 1 (expected %d, got %d)", errors.ErrInvalidPack, i+1, step.Number)
		}
	}

	return nil
}
