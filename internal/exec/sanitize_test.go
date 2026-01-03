package exec

import "testing"

func TestSanitizeScript_LabelLine(t *testing.T) {
	input := "GenerateReport\noracle --help\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(warnings))
	}
	if warnings[0].Token != "GenerateReport" {
		t.Fatalf("expected token GenerateReport, got %s", warnings[0].Token)
	}
	wantPrefix := "echo \"GenerateReport\""
	if got[:len(wantPrefix)] != wantPrefix {
		t.Fatalf("expected sanitized line to start with %q, got %q", wantPrefix, got)
	}
}

func TestSanitizeScript_BuiltinUnchanged(t *testing.T) {
	input := "echo\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected script unchanged, got %q", got)
	}
}

func TestSanitizeScript_HeredocUnchanged(t *testing.T) {
	input := "cat <<'EOF'\nGenerateReport\nEOF\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected heredoc unchanged, got %q", got)
	}
}
