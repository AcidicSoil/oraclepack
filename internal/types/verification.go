package types

// OutputContract describes the expected response shape for a step.
type OutputContract string

const (
	OutputContractUnknown          OutputContract = ""
	OutputContractAllSections      OutputContract = "all_sections"
	OutputContractDirectAnswerOnly OutputContract = "direct_answer_only"
	OutputContractChunkedBySuffix  OutputContract = "chunked_by_suffix"
)

// OutputFailure captures a missing or invalid output artifact.
type OutputFailure struct {
	StepID        string   `json:"step_id,omitempty" yaml:"step_id,omitempty"`
	Path          string   `json:"path,omitempty" yaml:"path,omitempty"`
	MissingTokens []string `json:"missing_tokens,omitempty" yaml:"missing_tokens,omitempty"`
	Error         string   `json:"error,omitempty" yaml:"error,omitempty"`
}

// SyntaxFinding captures a structural or syntax issue in generated bash.
type SyntaxFinding struct {
	StepID  string `json:"step_id,omitempty" yaml:"step_id,omitempty"`
	Line    int    `json:"line" yaml:"line"`
	Token   string `json:"token,omitempty" yaml:"token,omitempty"`
	Message string `json:"message" yaml:"message"`
}
