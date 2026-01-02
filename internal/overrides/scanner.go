package overrides

import (
	"regexp"
	"strings"
)

var oracleCmdRegex = regexp.MustCompile(`^(\s*)(oracle)\b`)

// Scan extracts oracle invocations from a script.
func Scan(script string) []OracleInvocation {
	var invocations []OracleInvocation
	lines := strings.Split(script, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)
		
		// Skip comments
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Check for oracle command
		loc := oracleCmdRegex.FindStringSubmatchIndex(line)
		if loc != nil {
			startLine := i
			// Group 1 is the indentation
			indentation := line[loc[2]:loc[3]]

			var cmdBuilder strings.Builder
			cmdBuilder.WriteString(line)

			endLine := i
			// Handle line continuations
			// Check if line ends with backslash (ignoring trailing whitespace)
			for {
				if endLine+1 >= len(lines) {
					break
				}
				
				// Check current line for continuation
				currTrimmed := strings.TrimRight(lines[endLine], " \t")
				if !strings.HasSuffix(currTrimmed, "\") {
					break
				}

				endLine++
				cmdBuilder.WriteString("\n")
				cmdBuilder.WriteString(lines[endLine])
			}

			invocations = append(invocations, OracleInvocation{
				StartLine:   startLine,
				EndLine:     endLine,
				Command:     cmdBuilder.String(),
				Indentation: indentation,
			})

			i = endLine // Advance loop
		}
	}
	return invocations
}
