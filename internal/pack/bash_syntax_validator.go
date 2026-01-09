package pack

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var orphanFlagRegex = regexp.MustCompile(`^\s*(?:-p|--prompt)(?:\s+.+)?$`)

// FindOrphanedFlags detects lines that contain only flags like -p/--prompt
// without being part of a continued command.
func FindOrphanedFlags(script string) []types.SyntaxFinding {
	if script == "" {
		return nil
	}

	lines := strings.Split(script, "\n")
	var findings []types.SyntaxFinding
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if !orphanFlagRegex.MatchString(line) {
			continue
		}

		if i > 0 && lineContinues(lines[i-1]) {
			continue
		}

		findings = append(findings, types.SyntaxFinding{
			Line:    i + 1,
			Token:   strings.Fields(trimmed)[0],
			Message: "Orphaned flag without a preceding command or line continuation",
		})
	}
	return findings
}

func lineContinues(line string) bool {
	trimmed := strings.TrimRight(line, " \t")
	return strings.HasSuffix(trimmed, "\\")
}

// CheckPackScripts validates prelude and step scripts for orphaned flags and bash syntax.
func CheckPackScripts(p *types.Pack) ([]types.SyntaxFinding, string, error) {
	if p == nil {
		return nil, "", nil
	}

	var findings []types.SyntaxFinding
	var warnings []string

	// Prelude orphaned flags
	for _, finding := range FindOrphanedFlags(p.Prelude.Code) {
		findings = append(findings, finding)
	}

	// Step orphaned flags and bash -n checks
	for _, step := range p.Steps {
		for _, finding := range FindOrphanedFlags(step.Code) {
			finding.StepID = step.ID
			findings = append(findings, finding)
		}

		syntaxFindings, warning, err := CheckBashSyntax(step.Code)
		if err != nil {
			return findings, warning, err
		}
		if warning != "" {
			warnings = append(warnings, warning)
		}
		for _, finding := range syntaxFindings {
			finding.StepID = step.ID
			findings = append(findings, finding)
		}
	}

	// Prelude bash -n check
	preludeFindings, warning, err := CheckBashSyntax(p.Prelude.Code)
	if err != nil {
		return findings, warning, err
	}
	if warning != "" {
		warnings = append(warnings, warning)
	}
	findings = append(findings, preludeFindings...)

	return findings, strings.Join(uniqueStrings(warnings), "; "), nil
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(values))
	var out []string
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}
