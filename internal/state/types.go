package state

import (
	"time"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

// RunState tracks the execution progress of an oracle pack.
type RunState struct {
	SchemaVersion int                   `json:"schema_version"`
	PackHash      string                `json:"pack_hash"`
	StartTime     time.Time             `json:"start_time"`
	StepStatuses  map[string]StepStatus `json:"step_statuses"`
	ROIThreshold  float64               `json:"roi_threshold,omitempty"`
	ROIMode       string                `json:"roi_mode,omitempty"`
	Warnings      []Warning             `json:"warnings,omitempty"`
}

// StepStatus holds the outcome of an individual step.
type StepStatus struct {
	Status    Status    `json:"status"`
	ExitCode  int       `json:"exit_code"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Error     string    `json:"error,omitempty"`
}

// Warning captures a non-fatal execution note (e.g., sanitized labels).
type Warning struct {
	Scope   string `json:"scope"`
	StepID  string `json:"step_id,omitempty"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
