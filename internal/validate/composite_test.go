package validate

import (
	"testing"

	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

func TestCompositeValidator(t *testing.T) {
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", Code: "codex exec \"hi\""},
		},
	}
	cv := CompositeValidator{
		ToolPresence: ToolPresenceValidator{Checker: fakeChecker{found: map[string]bool{"codex": false}}},
	}
	results := cv.ValidatePack(p)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Status != state.StatusSkipped {
		t.Fatalf("expected skipped, got %s", results[0].Status)
	}
}
