package pack

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/user/oraclepack/internal/types"
)

var bashLineRegex = regexp.MustCompile(`line\s+(\d+)`)

// CheckBashSyntax runs "bash -n" on the script and returns syntax findings.
// If bash is not found, it returns a warning and no findings.
func CheckBashSyntax(script string) ([]types.SyntaxFinding, string, error) {
	if strings.TrimSpace(script) == "" {
		return nil, "", nil
	}

	if _, err := exec.LookPath("bash"); err != nil {
		return nil, "bash not found on PATH; skipping bash -n syntax check", nil
	}

	cmd := exec.Command("bash", "-n")
	cmd.Stdin = strings.NewReader(script)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		line := 0
		if match := bashLineRegex.FindStringSubmatch(msg); len(match) > 1 {
			if parsed, parseErr := strconv.Atoi(match[1]); parseErr == nil {
				line = parsed
			}
		}
		return []types.SyntaxFinding{
			{
				Line:    line,
				Message: msg,
			},
		}, "", nil
	}

	return nil, "", nil
}
