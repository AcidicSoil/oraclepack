package pack

import (
	"fmt"
	"os"
	"regexp"
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
		if !containsToken(content, tok) {
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

func containsToken(content, token string) bool {
	if strings.Contains(content, token) {
		return true
	}

	heading := strings.TrimSpace(strings.TrimPrefix(token, "###"))
	if heading == token {
		return false
	}

	alts := []string{heading}
	switch strings.ToLower(heading) {
	case "direct answer":
		alts = append(alts, "answer")
	case "risks and unknowns":
		alts = append(alts, "risks/unknowns", "risks & unknowns")
	case "next experiment":
		alts = append(alts, "next smallest concrete experiment")
	case "missing evidence":
		alts = append(alts, "if evidence is insufficient")
	}

	for _, alt := range alts {
		if alt == "" {
			continue
		}
		pat := `(?im)^\s*#{0,3}\s*` + regexp.QuoteMeta(alt) + `\b`
		if regexp.MustCompile(pat).MatchString(content) {
			return true
		}
	}

	return false
}

// VerifyStepOutputs validates output files for a step. If requireHeadings is false,
// it only checks that output files exist and are non-empty.
func VerifyStepOutputs(step *types.Step, requireHeadings bool, chunkMode string) []types.OutputFailure {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}

	if !requireHeadings {
		paths = selectChunkPaths(paths, chunkMode)
		var failures []types.OutputFailure
		for _, path := range paths {
			info, err := os.Stat(path)
			if err != nil {
				failures = append(failures, types.OutputFailure{
					Path:  path,
					Error: err.Error(),
				})
				continue
			}
			if info.Size() == 0 {
				failures = append(failures, types.OutputFailure{
					Path:  path,
					Error: fmt.Sprintf("output file is empty: %s", path),
				})
			}
		}
		return failures
	}

	expectations := StepOutputExpectationsWithMode(step, chunkMode)
	if len(expectations) == 0 {
		return nil
	}
	var failures []types.OutputFailure
	for path, required := range expectations {
		ok, failure := ValidateOutputFile(path, required)
		if !ok {
			failures = append(failures, failure)
		}
	}
	return failures
}

func selectChunkPaths(paths []string, chunkMode string) []string {
	if len(paths) == 0 {
		return paths
	}
	mode := strings.ToLower(strings.TrimSpace(chunkMode))
	if mode == "" {
		mode = "auto"
	}
	switch mode {
	case "single":
		return []string{paths[0]}
	default:
		return paths
	}
}
