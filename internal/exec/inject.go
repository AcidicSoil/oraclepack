package exec

import "strings"

// InjectFlags scans a script and appends flags to any 'oracle' command invocation.
func InjectFlags(script string, flags []string) string {
	if len(flags) == 0 {
		return script
	}

	flagStr := strings.Join(flags, " ")

	lines := strings.Split(script, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		insertIdx := oracleInsertIndex(line)
		if insertIdx == -1 {
			continue
		}

		lines[i] = insertFlagsInLine(line, insertIdx, flagStr)
	}

	return strings.Join(lines, "\n")
}

func oracleInsertIndex(line string) int {
	i := 0
	for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
		i++
	}

	if !strings.HasPrefix(line[i:], "oracle") {
		return -1
	}

	end := i + len("oracle")
	if end < len(line) {
		next := line[end]
		if next != ' ' && next != '\t' {
			return -1
		}
	}

	return end
}

func insertFlagsInLine(line string, insertIdx int, flags string) string {
	prefix := line[:insertIdx]
	rest := line[insertIdx:]
	if rest == "" {
		return prefix + " " + flags
	}
	if rest[0] == ' ' || rest[0] == '\t' {
		return prefix + " " + flags + rest
	}
	return prefix + " " + flags + " " + rest
}
