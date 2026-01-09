package report

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteReport(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "report.json")

	rep := &ReportV1{
		PackInfo: PackInfo{Name: "pack"},
		Summary:  Summary{TotalSteps: 1},
	}
	if err := WriteReport(path, rep); err != nil {
		t.Fatalf("WriteReport: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected report file to exist: %v", err)
	}
}
