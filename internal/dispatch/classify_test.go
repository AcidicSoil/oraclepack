package dispatch

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		line    string
		wantOK  bool
		wantCmd string
	}{
		{"oracle query \"hi\"", true, "query \"hi\""},
		{"  tm list", true, "list"},
		{"task-master next", true, "next"},
		{"codex exec \"x\"", true, "exec \"x\""},
		{"gemini run", true, "run"},
		{"echo hello", false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			got, ok := Classify(tt.line)
			if ok != tt.wantOK {
				t.Fatalf("expected ok=%v got %v", tt.wantOK, ok)
			}
			if ok && got.Command != tt.wantCmd {
				t.Fatalf("expected cmd %q got %q", tt.wantCmd, got.Command)
			}
		})
	}
}
