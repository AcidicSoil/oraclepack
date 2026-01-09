package pack

import (
	"os"
	"regexp"
	"strings"
)

var writeOutputPathRegex = regexp.MustCompile(`(?m)--write-output\s+"([^"]+)"`)

// StepOutputExpectations returns a map of output paths to required tokens.
// If no validation is needed, it returns nil.
func StepOutputExpectations(step *Step) map[string][]string {
	paths := ExtractWriteOutputPaths(step.Code)
	if len(paths) == 0 {
		return nil
	}
	if len(paths) == 1 {
		tokens := expectedAnswerTokens(step.Code)
		if len(tokens) == 0 {
			return nil
		}
		return map[string][]string{paths[0]: tokens}
	}

	out := map[string][]string{}
	for _, path := range paths {
		switch {
		case strings.Contains(path, "-direct-answer"):
			out[path] = []string{"direct answer"}
		case strings.Contains(path, "-risks-unknowns"):
			out[path] = []string{"risks unknowns"}
		case strings.Contains(path, "-next-experiment"):
			out[path] = []string{"next smallest concrete experiment"}
		case strings.Contains(path, "-missing-evidence"):
			out[path] = []string{"missing file path pattern"}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
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

// ValidateOutputFile checks whether the output file contains the required answer sections.
// It returns ok=false with missing tokens when validation fails.
func ValidateOutputFile(path string, requiredTokens []string) (bool, []string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, nil, err
	}
	normalized := normalizeText(string(data))
	var missing []string
	for _, tok := range requiredTokens {
		if !strings.Contains(normalized, tok) {
			missing = append(missing, tok)
		}
	}
	if len(missing) > 0 {
		return false, missing, nil
	}
	return true, nil, nil
}

func expectedAnswerTokens(code string) []string {
	lower := strings.ToLower(code)
	if !strings.Contains(lower, "answer format") {
		return nil
	}
	return []string{
		"direct answer",
		"risks unknowns",
		"next smallest concrete experiment",
		"if evidence is insufficient",
	}
}

func normalizeText(s string) string {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	return s
}
