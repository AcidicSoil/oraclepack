package exec

import (
	"bufio"
	"regexp"
	"strings"
)

var (
	oracleCmdRegex = regexp.MustCompile(`^(\s*)oracle(\s+|$)`)
)

// InjectFlags scans a script and appends flags to any 'oracle' command invocation.
func InjectFlags(script string, flags []string) string {
	if len(flags) == 0 {
		return script
	}

	flagStr := " " + strings.Join(flags, " ")
	
	var result []string
	scanner := bufio.NewScanner(strings.NewReader(script))
	for scanner.Scan() {
		line := scanner.Text()
		if oracleCmdRegex.MatchString(line) {
			// Append flags to the line
			line += flagStr
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}
