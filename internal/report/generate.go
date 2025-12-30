package report

import (
	"time"

	"github.com/user/oraclepack/internal/state"
)

// GenerateReport creates a ReportV1 from a RunState.
func GenerateReport(s *state.RunState, packName string) *ReportV1 {
	report := &ReportV1{
		PackInfo: PackInfo{
			Name: packName,
			Hash: s.PackHash,
		},
		GeneratedAt: time.Now(),
		Steps:       []StepReport{},
	}

	var totalDuration time.Duration
	success, failure, skipped := 0, 0, 0

	for id, status := range s.StepStatuses {
		duration := status.EndedAt.Sub(status.StartedAt)
		if status.EndedAt.IsZero() || status.StartedAt.IsZero() {
			duration = 0
		}
		
		totalDuration += duration

		sr := StepReport{
			ID:         id,
			Status:     string(status.Status),
			ExitCode:   status.ExitCode,
			Duration:   duration,
			DurationMs: duration.Milliseconds(),
			Error:      status.Error,
		}
		report.Steps = append(report.Steps, sr)

		switch status.Status {
		case state.StatusSuccess:
			success++
		case state.StatusFailed:
			failure++
		case state.StatusSkipped:
			skipped++
		}
	}

	report.Summary = Summary{
		TotalSteps:      len(s.StepStatuses),
		SuccessCount:    success,
		FailureCount:    failure,
		SkippedCount:    skipped,
		TotalDuration:   totalDuration,
		TotalDurationMs: totalDuration.Milliseconds(),
	}

	return report
}
