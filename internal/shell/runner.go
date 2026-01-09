package shell

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

// Result captures command output and exit code.
type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

// RunCommand executes a command with login shell semantics using /bin/bash -lc.
func RunCommand(ctx context.Context, cmd string) (Result, error) {
	c := exec.CommandContext(ctx, "/bin/bash", "-lc", cmd)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()
	exitCode := 0
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
		} else {
			return Result{}, fmt.Errorf("run command: %w", err)
		}
	}

	return Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: exitCode,
	}, nil
}
