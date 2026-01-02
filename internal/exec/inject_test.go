package exec

import (
	"testing"
)

func TestInjectFlags(t *testing.T) {
	tests := []struct {
		name     string
		script   string
		flags    []string
		expected string
	}{
		{
			"simple injection",
			"oracle query 'hello'",
			[]string{"--verbose"},
			"oracle --verbose query 'hello'",
		},
		{
			"indented injection",
			"  oracle query 'hello'",
			[]string{"--verbose"},
			"  oracle --verbose query 'hello'",
		},
		{
			"no injection needed",
			"echo 'hello'",
			[]string{"--verbose"},
			"echo 'hello'",
		},
		{
			"multiple lines",
			"echo 'start'\noracle query\necho 'end'",
			[]string{"--debug"},
			"echo 'start'\noracle --debug query\necho 'end'",
		},
		{
			"multiline with continuation",
			"oracle \\\n  --json \\\n  --files",
			[]string{"--flag"},
			"oracle --flag \\\n  --json \\\n  --files",
		},
		{
			"multiline with args and continuation",
			"  oracle arg \\\n  --json",
			[]string{"--flag"},
			"  oracle --flag arg \\\n  --json",
		},
		{
			"commented command",
			"# oracle --json",
			[]string{"--verbose"},
			"# oracle --json",
		},
		{
			"oracle as part of word",
			"coracle query",
			[]string{"--verbose"},
			"coracle query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InjectFlags(tt.script, tt.flags)
			if got != tt.expected {
				t.Errorf("InjectFlags() = %q, want %q", got, tt.expected)
			}
		})
	}
}
