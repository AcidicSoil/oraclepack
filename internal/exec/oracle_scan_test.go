package exec

import (
	"reflect"
	"testing"
)

func TestExtractOracleInvocations(t *testing.T) {
	tests := []struct {
		name   string
		script string
		want   []OracleInvocation
	}{
		{
			name:   "Simple command",
			script: "oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "oracle --json", Display: "oracle --json", Indentation: ""},
			},
		},
		{
			name:   "Indented command",
			script: "  oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "  oracle --json", Display: "oracle --json", Indentation: "  "},
			},
		},
		{
			name: "Multiline command",
			script: `oracle \
  --json \
  --files`,
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 2, Raw: `oracle \
  --json \
  --files`, Display: `oracle \
  --json \
  --files`, Indentation: ""},
			},
		},
		{
			name: "Commented command",
			script: `# oracle --json
oracle --real`,
			want: []OracleInvocation{
				{StartLine: 1, EndLine: 1, Raw: "oracle --real", Display: "oracle --real", Indentation: ""},
			},
		},
		{
			name: "Multiple commands",
			script: `
echo start
oracle --one
echo mid
oracle --two
echo end
`,
			want: []OracleInvocation{
				{StartLine: 2, EndLine: 2, Raw: "oracle --one", Display: "oracle --one", Indentation: ""},
				{StartLine: 4, EndLine: 4, Raw: "oracle --two", Display: "oracle --two", Indentation: ""},
			},
		},
		{
			name:   "Oraclepack prefix (should not match)",
			script: "oraclepack run",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractOracleInvocations(tt.script)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractOracleInvocations() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
