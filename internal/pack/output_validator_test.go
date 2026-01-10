package pack

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestValidateOutputFile_RelaxedHeadings(t *testing.T) {
	content := `Direct answer

Risks/unknowns

Next smallest concrete experiment

If evidence is insufficient`

	dir := t.TempDir()
	path := filepath.Join(dir, "out.md")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write output: %v", err)
	}

	ok, failure := ValidateOutputFile(path, []string{
		"### Direct answer",
		"### Risks and unknowns",
		"### Next experiment",
		"### Missing evidence",
	})
	if !ok {
		t.Fatalf("expected relaxed headings to pass, failure: %#v", failure)
	}

	step := &types.Step{Code: `oracle --write-output "` + path + `"`}
	failures := VerifyStepOutputs(step, true, "single")
	if len(failures) != 0 {
		t.Fatalf("expected no failures in single chunk mode, got %#v", failures)
	}
}
