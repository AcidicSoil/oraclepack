package app

import (
	"bytes"
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestApp_RunPlain_ROI(t *testing.T) {
	steps := buildROISteps()
	packContent := `
# ROI Test Pack
` + "```" + `bash
` + steps + `
` + "```" + `
`
	packFile := "roi_test.md"
	defer os.Remove(packFile)
	os.WriteFile(packFile, []byte(packContent), 0644)

	// Test Case 1: Filter OVER 3.3 (Should run 5.0 and 3.3)
	t.Run("Filter Over 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "over",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if !strings.Contains(output, "Step 01") {
			t.Error("expected Step 01 (5.0) to run")
		}
		if !strings.Contains(output, "Step 02") {
			t.Error("expected Step 02 (3.3) to run (inclusive)")
		}
		if strings.Contains(output, "Step 03") && !strings.Contains(output, "Skipping step 03") {
			t.Error("expected Step 03 (1.0) to be skipped")
		}
	})

	// Test Case 2: Filter UNDER 3.3 (Should run 1.0 only)
	t.Run("Filter Under 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "under",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if strings.Contains(output, "Step 01") && !strings.Contains(output, "Skipping step 01") {
			t.Error("expected Step 01 (5.0) to be skipped")
		}
		if strings.Contains(output, "Step 02") && !strings.Contains(output, "Skipping step 02") {
			t.Error("expected Step 02 (3.3) to be skipped (exclusive)")
		}
		if !strings.Contains(output, "Step 03") {
			t.Error("expected Step 03 (1.0) to run")
		}
	})
}

func buildROISteps() string {
	var b strings.Builder
	for i := 1; i <= 20; i++ {
		id := i
		if id < 10 {
			b.WriteString("# 0")
		} else {
			b.WriteString("# ")
		}
		b.WriteString(strconv.Itoa(id))
		if i == 1 {
			b.WriteString(") ROI=5.0\n")
			b.WriteString("echo \"high\"\n\n")
			continue
		}
		if i == 2 {
			b.WriteString(") ROI=3.3\n")
			b.WriteString("echo \"threshold\"\n\n")
			continue
		}
		if i == 3 {
			b.WriteString(") ROI=1.0\n")
			b.WriteString("echo \"low\"\n\n")
			continue
		}
		b.WriteString(")\n")
		b.WriteString("echo \"step ")
		b.WriteString(strconv.Itoa(id))
		b.WriteString("\"\n\n")
	}
	return b.String()
}
