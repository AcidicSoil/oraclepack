package pack

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var (
	writeOutputPathRegex = regexp.MustCompile(`(?m)--write-output\s+["']?([^"'\s]+)["']?`)
	answerFormatRegex    = regexp.MustCompile(`(?i)answer\s+format`)
	directOnlyRegex      = regexp.MustCompile(`(?i)return\s+only[:\s]*direct\s+answer`)
)

// DetectOutputContract determines the expected response contract for a step.
func DetectOutputContract(step types.Step) types.OutputContract {
	if step.Code == "" {
		return types.OutputContractUnknown
	}

	hasAnswerFormat := answerFormatRegex.MatchString(step.Code)
	if !hasAnswerFormat {
		return types.OutputContractUnknown
	}

	if directOnlyRegex.MatchString(step.Code) {
		return types.OutputContractDirectAnswerOnly
	}

	return types.OutputContractAllSections
}

// StepOutputExpectations returns a map of output paths to required tokens.
// If no validation is needed, it returns nil.
func StepOutputExpectations(step *types.Step) map[string][]string {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}

	if len(paths) > 1 {
		out := map[string][]string{}
		for _, path := range paths {
			switch {
			case strings.Contains(path, "-direct-answer"):
				out[path] = []string{"### Direct answer"}
			case strings.Contains(path, "-risks-unknowns"):
				out[path] = []string{"### Risks and unknowns"}
			case strings.Contains(path, "-next-experiment"):
				out[path] = []string{"### Next experiment"}
			case strings.Contains(path, "-missing-evidence"):
				out[path] = []string{"### Missing evidence"}
			}
		}
		if len(out) == 0 {
			return nil
		}
		return out
	}

	switch DetectOutputContract(*step) {
	case types.OutputContractDirectAnswerOnly:
		return map[string][]string{paths[0]: []string{"### Direct answer"}}
	case types.OutputContractAllSections:
		return map[string][]string{paths[0]: {
			"### Direct answer",
			"### Risks and unknowns",
			"### Next experiment",
			"### Missing evidence",
		}}
	default:
		return nil
	}
}

// StepOutputExpectationsWithMode returns expectations honoring chunk mode.
// chunkMode: auto (default), single (treat as single output), multi (force suffix mapping when multiple outputs).
func StepOutputExpectationsWithMode(step *types.Step, chunkMode string) map[string][]string {
	if step == nil {
		return nil
	}
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}
	mode := strings.ToLower(strings.TrimSpace(chunkMode))
	if mode == "" {
		mode = "auto"
	}

	switch mode {
	case "single":
		// Always treat as a single output (use first path).
		return expectationsForSingle(step, paths[0])
	case "multi":
		if len(paths) > 1 {
			return expectationsForSuffixes(paths)
		}
		return expectationsForSingle(step, paths[0])
	default: // auto
		if len(paths) > 1 {
			return expectationsForSuffixes(paths)
		}
		return expectationsForSingle(step, paths[0])
	}
}

func expectationsForSuffixes(paths []string) map[string][]string {
	out := map[string][]string{}
	for _, path := range paths {
		switch {
		case strings.Contains(path, "-direct-answer"):
			out[path] = []string{"### Direct answer"}
		case strings.Contains(path, "-risks-unknowns"):
			out[path] = []string{"### Risks and unknowns"}
		case strings.Contains(path, "-next-experiment"):
			out[path] = []string{"### Next experiment"}
		case strings.Contains(path, "-missing-evidence"):
			out[path] = []string{"### Missing evidence"}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func expectationsForSingle(step *types.Step, path string) map[string][]string {
	switch DetectOutputContract(*step) {
	case types.OutputContractDirectAnswerOnly:
		return map[string][]string{path: []string{"### Direct answer"}}
	case types.OutputContractAllSections:
		return map[string][]string{path: {
			"### Direct answer",
			"### Risks and unknowns",
			"### Next experiment",
			"### Missing evidence",
		}}
	default:
		return nil
	}
}

// ExtractWriteOutputPaths returns all --write-output paths found in the step code.
func ExtractWriteOutputPaths(code string) []string {
	matches := writeOutputPathRegex.FindAllStringSubmatch(code, -1)
	if len(matches) == 0 {
		return nil
	}
	paths := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) >= 2 {
			paths = append(paths, m[1])
		}
	}
	return paths
}
