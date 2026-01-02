package tui

import (
	"fmt"
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

func (m OverridesConfirmModel) View(over overrides.RuntimeOverrides) string {
	added := strings.Join(over.AddedFlags, ", ")
	if added == "" {
		added = "(none)"
	}
	removed := strings.Join(over.RemovedFlags, ", ")
	if removed == "" {
		removed = "(none)"
	}
	targeted := len(over.ApplyToSteps)
	url := over.ChatGPTURL
	if url == "" {
		url = "(none)"
	}

	lines := []string{
		"Summary:",
		fmt.Sprintf("Added flags: %s", added),
		fmt.Sprintf("Removed flags: %s", removed),
		fmt.Sprintf("ChatGPT URL: %s", url),
		fmt.Sprintf("Targeted steps: %d", targeted),
		"",
		"[Enter] Validate  [Esc] Cancel",
	}

	if m.validating {
		lines = append(lines, "", "Validating overrides...")
	}
	if m.errMsg != "" {
		lines = append(lines, "", "Validation failed:", m.errMsg)
	}

	return strings.Join(lines, "\n")
}
