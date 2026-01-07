1) Direct answer (observability + correlation identifiers)

Primary structured observability output is the run “report”: internal/report/GenerateReport(s *state.RunState, packName string) *ReportV1 builds a ReportV1 containing PackInfo{Name, Hash}, GeneratedAt, per-step results (Steps: []StepReport), and a Summary with counts + total duration. (Evidence: internal/report/generate.go, GenerateReport populating PackInfo, Steps, Summary.)

Per-step signals captured (and thus observable via the report) include: StepReport.ID (step identifier), Status, ExitCode, Duration/DurationMs, and Error. These are derived from RunState.StepStatuses map[string]StepStatus and each StepStatus’s timestamps/exit code/error. (Evidence: internal/report/generate.go builds StepReport from the for id, status := range s.StepStatuses loop; internal/state/types.go defines RunState.StepStatuses and StepStatus fields.)

Warnings are first-class, structured observability signals: RunState.Warnings []Warning are copied into report.Warnings with fields {Scope, StepID, Line, Token, Message}. (Evidence: internal/state/types.go defines Warning; internal/report/generate.go copies s.Warnings into report.Warnings.)

Primary correlation identifiers across components (based on what’s present in state/report):

pack_hash: RunState.PackHash and ReportV1.PackInfo.Hash (ties a run to a specific pack content). (Evidence: internal/state/types.go PackHash; internal/report/generate.go assigns Hash: s.PackHash.)

step_id: the map key in RunState.StepStatuses and StepReport.ID; also optionally present on warnings (Warning.StepID). This is the main join key between step execution, step status, and warnings. (Evidence: internal/state/types.go StepStatuses map[string]StepStatus, Warning.StepID; internal/report/generate.go uses loop variable id → StepReport.ID.)

start_time / generated_at: RunState.StartTime and ReportV1.GeneratedAt provide coarse run-level temporal correlation (e.g., match logs/events by time window). (Evidence: internal/state/types.go StartTime; internal/report/generate.go sets GeneratedAt: time.Now().)

There is no explicit “run_id”/trace id in the provided state/report schema; correlation appears to rely on (pack_hash, start_time) + step IDs, rather than a stable UUID. (Evidence: internal/state/types.go RunState fields shown do not include a run identifier other than PackHash/StartTime.)

Log lines: no logging primitives (log.*, fmt.Printf, etc.) are shown in the provided generate.go or state/types.go, so log-based observability and log correlation keys cannot be confirmed from these files alone. (Evidence: the provided snippets contain only struct composition / time math.)

2) Risks / unknowns

Run correlation fragility: without a dedicated run_id, correlating across components (state file writes, executor output, CLI output, report artifacts) will be brittle when multiple runs share the same pack hash or occur close in time.

Non-deterministic step ordering in reports: iterating for id, status := range s.StepStatuses over a Go map yields randomized order, which can create noisy diffs and make “compare reports across runs” harder.

Missing visibility into logging: cannot verify
