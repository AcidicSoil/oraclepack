package templates

import (
	"os"
	"testing"

	"github.com/user/oraclepack/internal/pack"
)

func TestRenderTicketActionPack(t *testing.T) {
	got := RenderTicketActionPack()
	if got == "" {
		t.Fatal("expected non-empty template")
	}

	// Golden comparison
	data, err := os.ReadFile("ticket-action-pack.md")
	if err != nil {
		t.Fatalf("read template: %v", err)
	}
	if string(data) != got {
		t.Fatalf("template mismatch with golden file")
	}

	// Ensure pack is parseable and validates 20-step contract.
	p, err := pack.Parse([]byte(got))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if err := p.Validate(); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}
	if len(p.Steps) != 20 {
		t.Fatalf("expected 20 steps, got %d", len(p.Steps))
	}
}
