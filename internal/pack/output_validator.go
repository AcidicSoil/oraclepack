package pack

import (
	"os"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

// ValidateOutputFile checks whether the output file contains the required answer sections.
// It returns ok=false with a populated OutputFailure when validation fails.
func ValidateOutputFile(path string, requiredTokens []string) (bool, types.OutputFailure) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, types.OutputFailure{
			Path:  path,
			Error: err.Error(),
		}
	}

	content := string(data)
	var missing []string
	for _, tok := range requiredTokens {
		if !strings.Contains(content, tok) {
			missing = append(missing, tok)
		}
	}

	if len(missing) > 0 {
		return false, types.OutputFailure{
			Path:          path,
			MissingTokens: missing,
		}
	}

	return true, types.OutputFailure{}
}
