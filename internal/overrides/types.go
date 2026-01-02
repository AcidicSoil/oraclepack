package overrides

// RuntimeOverrides holds configuration for runtime flag modifications.
type RuntimeOverrides struct {
	AddedFlags   []string        // Flags to append (e.g., "--model=gpt-4")
	RemovedFlags []string        // Flags to remove (e.g., "--json")
	ChatGPTURL   string          // Optional URL to inject via --chatgpt-url
	ApplyToSteps map[string]bool // Set of step IDs to apply overrides to. If empty, applies to none.
}

// OracleInvocation represents a detected oracle command in a script.
type OracleInvocation struct {
	StartLine   int    // 0-based start line index
	EndLine     int    // 0-based end line index (inclusive)
	Command     string // The full command string (joined if multi-line)
	Indentation string // The leading whitespace
}