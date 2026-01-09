package report

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteReport writes a ReportV1 to disk.
func WriteReport(path string, rep *ReportV1) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	return nil
}
