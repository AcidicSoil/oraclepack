package artifacts

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEvaluateGates(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, ".oraclepack", "ticketify")
	if err := os.MkdirAll(base, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	contract := Contract{
		"09": {filepath.Join(base, "next.json")},
	}

	// Missing file should error.
	if err := EvaluateGates("09", contract); err == nil {
		t.Fatal("expected missing artifact error")
	}

	// Create file and verify pass.
	path := filepath.Join(base, "next.json")
	if err := os.WriteFile(path, []byte("ok"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := EvaluateGates("09", contract); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
