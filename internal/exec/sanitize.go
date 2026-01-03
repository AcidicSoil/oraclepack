package exec

import (
	osexec "os/exec"
	"regexp"
	"strings"
)

// SanitizeWarning captures a label line that was converted to a safe echo.
type SanitizeWarning struct {
	Scope   string
	StepID  string
	Line    int
	Token   string
	Message string
}

var (
	labelTokenRegex   = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*$`)
	heredocStartRegex = regexp.MustCompile(`<<-?\s*['"]?([A-Za-z0-9_]+)['"]?`)
)

var shellBuiltins = map[string]bool{
	"alias":    true,
	"bg":       true,
	"break":    true,
	"cd":       true,
	"command":  true,
	"continue": true,
	"declare":  true,
	"dirs":     true,
	"echo":     true,
	"eval":     true,
	"exec":     true,
	"exit":     true,
	"export":   true,
	"fg":       true,
	"getopts":  true,
	"hash":     true,
	"help":     true,
	"jobs":     true,
	"local":    true,
	"popd":     true,
	"printf":   true,
	"pushd":    true,
	"pwd":      true,
	"readonly": true,
	"return":   true,
	"set":      true,
	"shift":    true,
	"source":   true,
	"test":     true,
	"trap":     true,
	"true":     true,
	"type":     true,
	"ulimit":   true,
	"umask":    true,
	"unalias":  true,
	"unset":    true,
	"wait":     true,
	"false":    true,
}

var shellKeywords = map[string]bool{
	"case":     true,
	"do":       true,
	"done":     true,
	"elif":     true,
	"else":     true,
	"esac":     true,
	"fi":       true,
	"for":      true,
	"function": true,
	"if":       true,
	"in":       true,
	"select":   true,
	"then":     true,
	"time":     true,
	"until":    true,
	"while":    true,
}

// SanitizeScript converts bare label-like lines into safe echo statements.
func SanitizeScript(script, scope, stepID string) (string, []SanitizeWarning) {
	if script == "" {
		return script, nil
	}

	lines := strings.Split(script, "\n")
	var warnings []SanitizeWarning
	var heredocEnd string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if heredocEnd != "" {
			if trimmed == heredocEnd {
				heredocEnd = ""
			}
			continue
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if end := heredocStartToken(trimmed); end != "" {
			heredocEnd = end
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) != 1 {
			continue
		}
		token := fields[0]
		if !labelTokenRegex.MatchString(token) {
			continue
		}
		lower := strings.ToLower(token)
		if shellBuiltins[lower] || shellKeywords[lower] {
			continue
		}
		if _, err := osexec.LookPath(token); err == nil {
			continue
		}

		indent := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
		lines[i] = indent + "echo \"" + token + "\""
		warnings = append(warnings, SanitizeWarning{
			Scope:   scope,
			StepID:  stepID,
			Line:    i + 1,
			Token:   token,
			Message: "Converted bare label to echo",
		})
	}

	return strings.Join(lines, "\n"), warnings
}

func heredocStartToken(line string) string {
	match := heredocStartRegex.FindStringSubmatch(line)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
