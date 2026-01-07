1) Direct answer (top consistency risks + knobs)

Silent state-write failures → persisted state can lag in-memory indefinitely. saveState() discards errors (_ = state.SaveStateAtomic(...)), so the run can proceed while the on-disk state remains stale. Evidence: internal/app/run.go: saveState().

Non-transactional step transitions (RUNNING → SUCCESS/FAILED) can leave “stuck RUNNING” on disk. Each step does: set StatusRunning + saveState(), then later set EndedAt + final status + saveState(); any crash or write failure between these creates divergence (disk shows RUNNING; memory/report show final). Evidence: internal/app/run.go: RunPlain() loop setting StatusRunning then later StatusFailed/StatusSuccess.

Warnings can be persisted out of sync with step status. recordWarnings() drains and appends warnings and immediately calls saveState(); it runs after RunStep but before the final status save, so the state file can contain warnings from a step while still showing that step as RUNNING if the later status save fails. Evidence: internal/app/run.go: recordWarnings() + call sites around RunStep.

ROI-based skipping is not reflected in persisted state/report as “skipped”. ROI filter continues before writing any StepStatus; yet the state model has StatusSkipped available. This makes outputs inconsistent with console logs (“Skipping step…”) and loses auditability. Evidence: internal/app/run.go: ROIThreshold/ROIMode filter continues + internal/state/types.go: StatusSkipped exists.

StartTime is not guaranteed to flush. StartTime is set in-memory (if IsZero() { StartTime = time.Now() }) but there’s no immediate saveState(); if there are zero steps (and no warnings) you can generate a report with the updated StartTime while the persisted state file never updates. Evidence: internal/app/run.go: StartTime set; finalize() can run without any prior save.

Report write success is claimed even if the write fails. finalize() ignores json.MarshalIndent and os.WriteFile errors and prints “Report written to …” unconditionally, so report output can be missing/partial while stdout says it exists. Evidence: internal/app/run.go: finalize().

Resume behavior depends on persisted success markers and can misbehave with stale state. Resume skips only when StepStatuses[step.ID].Status == success; if state saves fail, already-completed steps can rerun (or previously-successful steps might not be recognized). Evidence: internal/app/run.go: resume check.

State behavior knobs/configs (observable here):

Config.StatePath: enables/disables persisted run state (saveState() is a no-op if empty). Evidence: internal/app/run.go: saveState().

Config.ReportPath: enables/disables report writing. Evidence: internal/app/run.go: finalize().

Config.StopOnFail: controls early termination path (still calls finalize() on failure). Evidence: internal/app/run.go: StopOnFail branch.

Config.ROIThreshold + Config.ROIMode: controls skip filter semantics (“under” is strict <; otherwise >=). Evidence: internal/app/run.go: ROI filter.

2) Risks/unknowns

Atomicity guarantees are unknown. state.SaveStateAtomic is assumed to be atomic, but wi
