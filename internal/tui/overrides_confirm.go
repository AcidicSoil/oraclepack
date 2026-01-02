package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
)

type ValidationResultMsg struct {
	Errors []exec.ValidationError
	Err    error
}

type OverridesConfirmModel struct {
	validating bool
	errMsg     string
	errors     []exec.ValidationError
}

func (m OverridesConfirmModel) View(over overrides.RuntimeOverrides, baseline []string) string {
	added := strings.Join(over.AddedFlags, ", ")
	if added == "" {
		added = "(none)"
	}
	removed := strings.Join(over.RemovedFlags, ", ")
	if removed == "" {
		removed = "(none)"
	}
	targeted := len(over.ApplyToSteps)
	targetList := formatTargetList(over.ApplyToSteps, 5)
	effective := effectiveFlagsSummary(over, baseline)
	lines := []string{
		"Summary:",
		fmt.Sprintf("Added flags: %s", added),
		fmt.Sprintf("Removed flags: %s", removed),
		fmt.Sprintf("Targeted steps: %d%s", targeted, targetList),
		fmt.Sprintf("Effective flags: %s", effective),
		"",
		"[Enter] Validate  [Esc] Cancel",
	}

	if m.validating {
		lines = append(lines, "", "Validating overrides...")
	}
	if m.errMsg != "" {
		lines = append(lines, "", "Validation failed:", m.errMsg)
	}
	if len(m.errors) > 0 {
		lines = append(lines, "", fmt.Sprintf("Validation errors (%d):", len(m.errors)))
		lines = append(lines, formatValidationErrors(m.errors, 6)...)
	}

	return strings.Join(lines, "\n")
}

func formatTargetList(targets map[string]bool, limit int) string {
	if len(targets) == 0 || limit <= 0 {
		return ""
	}
	ids := make([]string, 0, len(targets))
	for id := range targets {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	if len(ids) <= limit {
		return fmt.Sprintf(" (%s)", strings.Join(ids, ", "))
	}
	return fmt.Sprintf(" (%s, +%d more)", strings.Join(ids[:limit], ", "), len(ids)-limit)
}

func effectiveFlagsSummary(over overrides.RuntimeOverrides, baseline []string) string {
	if len(over.ApplyToSteps) == 0 {
		return "(no steps targeted)"
	}
	var first string
	for id := range over.ApplyToSteps {
		first = id
		break
	}
	flags := over.EffectiveFlags(first, baseline)
	if len(flags) == 0 {
		return "(none)"
	}
	return strings.Join(flags, " ")
}

func formatValidationErrors(errors []exec.ValidationError, limit int) []string {
	if limit <= 0 {
		return nil
	}
	lines := []string{}
	for i, err := range errors {
		if i >= limit {
			lines = append(lines, fmt.Sprintf("- (+%d more)", len(errors)-limit))
			break
		}
		msg := strings.TrimSpace(err.ErrorMessage)
		if msg == "" {
			msg = "(no error message)"
		}
		lines = append(lines, fmt.Sprintf("- Step %s: %s", err.StepID, firstLine(msg)))
	}
	return lines
}

func firstLine(msg string) string {
	if idx := strings.IndexByte(msg, '\n'); idx != -1 {
		return msg[:idx]
	}
	return msg
}
