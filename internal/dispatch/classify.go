package dispatch

import (
	"regexp"
	"strings"

	"github.com/user/oraclepack/internal/tools"
)

var classifier = regexp.MustCompile(`^(\s*)(oracle|tm|task-master|codex|gemini)\b`)

// Classification describes a parsed command prefix.
type Classification struct {
	Kind    tools.ToolKind
	Prefix  string
	Command string
}

// Classify detects a supported tool prefix and returns the remaining command.
func Classify(line string) (Classification, bool) {
	m := classifier.FindStringSubmatch(line)
	if len(m) < 3 {
		return Classification{}, false
	}
	prefix := m[2]
	kind := toolKindFromPrefix(prefix)
	if kind == nil {
		return Classification{}, false
	}
	trimmed := strings.TrimSpace(line[len(m[1])+len(prefix):])
	return Classification{Kind: *kind, Prefix: prefix, Command: strings.TrimSpace(trimmed)}, true
}

func toolKindFromPrefix(prefix string) *tools.ToolKind {
	var kind tools.ToolKind
	switch prefix {
	case "oracle":
		kind = tools.ToolOracle
	case "tm":
		kind = tools.ToolTM
	case "task-master":
		kind = tools.ToolTaskMaster
	case "codex":
		kind = tools.ToolCodex
	case "gemini":
		kind = tools.ToolGemini
	default:
		return nil
	}
	return &kind
}
