package pack

import (
	"fmt"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

// VerifyReport captures output verification results across a pack.
type VerifyReport struct {
	TotalSteps   int
	CheckedSteps int
	Failures     []types.OutputFailure
}

// FormatVerifyReport renders a human-readable report.
func FormatVerifyReport(report VerifyReport) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Verified outputs for %d/%d steps\n", report.CheckedSteps, report.TotalSteps)
	if len(report.Failures) == 0 {
		b.WriteString("All required output tokens were found.\n")
		return b.String()
	}

	b.WriteString("Missing or invalid outputs:\n")
	for _, failure := range report.Failures {
		stepLabel := failure.StepID
		if stepLabel == "" {
			stepLabel = "unknown step"
		}
		fmt.Fprintf(&b, "- Step %s: %s", stepLabel, failure.Path)
		if failure.Error != "" {
			fmt.Fprintf(&b, " (error: %s)", failure.Error)
		} else if len(failure.MissingTokens) > 0 {
			fmt.Fprintf(&b, " missing %s", strings.Join(failure.MissingTokens, ", "))
		}
		b.WriteString("\n")
	}
	return b.String()
}
