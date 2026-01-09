package pack

import (
	"os/exec"
	"testing"

	"github.com/user/oraclepack/internal/types"
)

func TestFindOrphanedFlags(t *testing.T) {
	script := "-p \"hello\"\n"
	findings := FindOrphanedFlags(script)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].Line != 1 {
		t.Fatalf("expected line 1, got %d", findings[0].Line)
	}

	script = "oracle \\\n  -p \"ok\"\n"
	findings = FindOrphanedFlags(script)
	if len(findings) != 0 {
		t.Fatalf("expected no findings for continued line, got %d", len(findings))
	}
}

func TestCheckBashSyntax(t *testing.T) {
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("bash not available")
	}

	script := "if true\n  echo hello\n"
	findings, warning, err := CheckBashSyntax(script)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if warning != "" {
		t.Fatalf("expected no warning, got %s", warning)
	}
	if len(findings) == 0 {
		t.Fatal("expected syntax findings")
	}
	if findings[0].Line == 0 {
		t.Fatal("expected line number in syntax finding")
	}
}

func TestCheckPackScriptsReportsStepID(t *testing.T) {
	p := &types.Pack{
		Steps: []types.Step{
			{ID: "01", Code: "-p \"oops\""},
		},
	}
	findings, _, err := CheckPackScripts(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].StepID != "01" {
		t.Fatalf("expected step ID 01, got %q", findings[0].StepID)
	}
}
