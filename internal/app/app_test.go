package app

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
)

func TestApp_RunPlain(t *testing.T) {
	steps := buildSteps(20, "echo")
	packContent := `
# Test Pack
` + "```" + `bash
` + steps + `
` + "```" + `
`
	packFile := "test.md"
	stateFile := "test_state.json"
	reportFile := "test_report.json"
	defer os.Remove(packFile)
	defer os.Remove(stateFile)
	defer os.Remove(reportFile)

	os.WriteFile(packFile, []byte(packContent), 0644)

	cfg := Config{
		PackPath:   packFile,
		StatePath:  stateFile,
		ReportPath: reportFile,
	}

	a := New(cfg)
	if err := a.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}
	if err := a.LoadState(); err != nil {
		t.Fatalf("LoadState failed: %v", err)
	}
	
	var out bytes.Buffer
	err := a.RunPlain(context.Background(), &out)
	if err != nil {
		t.Fatalf("RunPlain failed: %v", err)
	}

	output := out.String()
	if !contains(output, "step 1") || !contains(output, "step 2") {
		t.Errorf("output missing steps: %s", output)
	}

	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		t.Error("state file was not created")
	}

	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("report file was not created")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}

func buildSteps(count int, cmd string) string {
	var b bytes.Buffer
	for i := 1; i <= count; i++ {
		if i < 10 {
			b.WriteString("# 0")
		} else {
			b.WriteString("# ")
		}
		b.WriteString(fmt.Sprintf("%d)\n", i))
		b.WriteString(cmd)
		b.WriteString(fmt.Sprintf(" \"step %d\"\n", i))
	}
	return b.String()
}
