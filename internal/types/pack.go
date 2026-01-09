package types

// Pack represents a parsed oracle pack.
type Pack struct {
	Prelude     Prelude `json:"prelude" yaml:"prelude"`
	Steps       []Step  `json:"steps" yaml:"steps"`
	Source      string  `json:"source,omitempty" yaml:"source,omitempty"`
	OutDir      string  `json:"out_dir,omitempty" yaml:"out_dir,omitempty"`
	WriteOutput bool    `json:"write_output" yaml:"write_output"`
}

// Prelude contains the shell code that runs before any steps.
type Prelude struct {
	Code string `json:"code" yaml:"code"`
}

// Step represents an individual executable step within the pack.
type Step struct {
	ID           string  `json:"id" yaml:"id"`                             // e.g., "01"
	Number       int     `json:"number" yaml:"number"`                     // e.g., 1
	Code         string  `json:"code" yaml:"code"`                         // The bash code
	OriginalLine string  `json:"original_line" yaml:"original_line"`       // The header line, e.g., "# 01)"
	ROI          float64 `json:"roi,omitempty" yaml:"roi,omitempty"`       // Return on Investment value extracted from header
	Impact       string  `json:"impact,omitempty" yaml:"impact,omitempty"` // Optional impact metadata extracted from step comments
}
