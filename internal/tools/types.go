package tools

// ToolKind identifies a supported tool prefix.
type ToolKind int

const (
	ToolUnknown ToolKind = iota
	ToolOracle
	ToolTM
	ToolTaskMaster
	ToolCodex
	ToolGemini
)

// ToolMetadata captures tool invocation details.
type ToolMetadata struct {
	Name string
	Args []string
}

var registry = map[ToolKind]ToolMetadata{
	ToolUnknown:    {Name: "unknown"},
	ToolOracle:     {Name: "oracle"},
	ToolTM:         {Name: "tm"},
	ToolTaskMaster: {Name: "task-master"},
	ToolCodex:      {Name: "codex", Args: []string{"exec"}},
	ToolGemini:     {Name: "gemini"},
}

// Metadata returns tool metadata if present.
func Metadata(kind ToolKind) (ToolMetadata, bool) {
	meta, ok := registry[kind]
	return meta, ok
}

// Name returns the canonical tool name.
func (k ToolKind) Name() string {
	if meta, ok := registry[k]; ok {
		return meta.Name
	}
	return "unknown"
}

// PresenceChecker abstracts binary detection.
type PresenceChecker interface {
	DetectBinary(name string) (string, bool)
}
