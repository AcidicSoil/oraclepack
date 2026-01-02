package exec

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

func TestValidateOverrides_Success(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --ok"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	_, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		[]string{"--base"},
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
}

func TestValidateOverrides_Error(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --bad"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	errs, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		nil,
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 validation error, got %d", len(errs))
	}
	msg := errs[0].ErrorMessage
	if !strings.Contains(msg, "invalid flag") && !strings.Contains(msg, "unknown option") {
		t.Fatalf("unexpected error message: %q", msg)
	}
	if !strings.Contains(errs[0].Command, "--dry-run summary") {
		t.Fatalf("expected command to include --dry-run summary, got %q", errs[0].Command)
	}
}

func writeOracleStub(t *testing.T, dir string) {
	t.Helper()
	stub := `#!/bin/sh
has_dry=0
has_summary=0
for arg in "$@"; do
  if [ "$arg" = "--dry-run" ]; then has_dry=1; fi
  if [ "$arg" = "summary" ]; then has_summary=1; fi
  if [ "$arg" = "--bad" ]; then echo "invalid flag" 1>&2; exit 1; fi
done
if [ $has_dry -eq 0 ] || [ $has_summary -eq 0 ]; then
  echo "missing dry run" 1>&2
  exit 1
fi
exit 0
`
	path := filepath.Join(dir, "oracle")
	if err := os.WriteFile(path, []byte(stub), 0o755); err != nil {
		t.Fatalf("write oracle stub: %v", err)
	}
}
