package overrides

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name   string
		script string
		want   []OracleInvocation
	}{
		{
			name:   "Simple command",
			script: "oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Command: "oracle --json", Indentation: ""},
			},
		},
		{
			name:   "Indented command",
			script: "  oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Command: "  oracle --json", Indentation: "  "},
			},
		},
		{
			name: "Multiline command",
			script: `oracle \
  --json \
  --files`,
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 2, Command: "oracle \
  --json \
  --files", Indentation: ""},
			},
		},
		{
			name: "Commented command",
			script: `# oracle --json
oracle --real`,
			want: []OracleInvocation{
				{StartLine: 1, EndLine: 1, Command: "oracle --real", Indentation: ""},
			},
		},
		{
			name: "Multiple commands",
			script: "
echo start
oracle --one
echo mid
oracle --two
echo end
",
			want: []OracleInvocation{
				{StartLine: 2, EndLine: 2, Command: "oracle --one", Indentation: ""},
				{StartLine: 4, EndLine: 4, Command: "oracle --two", Indentation: ""},
			},
		},
		{
			name: "Oraclepack prefix (should not match)",
			script: "oraclepack run",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Scan(tt.script)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
