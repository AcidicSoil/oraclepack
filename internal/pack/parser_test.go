package pack

import (
	"testing"
)

func TestParse(t *testing.T) {
	content := []byte(`
# My Pack
Some description.

` + "```" + `bash
out_dir="dist"
--write-output

# 01)
echo "hello"

# 02)
echo "world"
` + "```" + `
`)

	p, err := Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if p.OutDir != "dist" {
		t.Errorf("expected OutDir dist, got %s", p.OutDir)
	}

	if !p.WriteOutput {
		t.Errorf("expected WriteOutput true, got false")
	}

	if len(p.Steps) != 2 {
		t.Errorf("expected 2 steps, got %d", len(p.Steps))
	}

	if p.Steps[0].ID != "01" || p.Steps[0].Number != 1 {
		t.Errorf("step 1 mismatch: %+v", p.Steps[0])
	}

	if err := p.Validate(); err != nil {
		t.Errorf("Validate failed: %v", err)
	}
}

func TestValidateErrors(t *testing.T) {
	tests := []struct {
		name    string
		pack    *Pack
		wantErr string
	}{
		{
			"no steps",
			&Pack{},
			"at least one step is required",
		},
		{
			"duplicate steps",
			&Pack{
				Steps: []Step{
					{Number: 1, ID: "01"},
					{Number: 1, ID: "01"},
				},
			},
			"duplicate step number 1",
		},
		{
			"non-sequential",
			&Pack{
				Steps: []Step{
					{Number: 1, ID: "01"},
					{Number: 3, ID: "03"},
				},
			},
			"steps must be sequential starting from 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pack.Validate()
			if err == nil {
				t.Error("expected error, got nil")
			} else if !contains(err.Error(), tt.wantErr) {
				t.Errorf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}
