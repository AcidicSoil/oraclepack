package shell

import (
	"context"
	"strings"
	"testing"
)

func TestRunCommandLoginShell(t *testing.T) {
	res, err := RunCommand(context.Background(), "echo $PATH")
	if err != nil {
		t.Fatalf("RunCommand: %v", err)
	}
	if strings.TrimSpace(res.Stdout) == "" {
		t.Fatalf("expected PATH output, got empty stdout")
	}
	if res.ExitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", res.ExitCode)
	}
}
