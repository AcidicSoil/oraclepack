1) Direct answer (failure-mode taxonomy; where classified/handled; what’s surfaced)

Precondition / lifecycle failures (internal “programmer error” / wrong call order): RunPlain hard-fails if a.Pack == nil or a.State == nil with plain fmt.Errorf("pack not loaded") / "state not loaded". These are not classified into internal/errors sentinel types, so callers will likely treat them as generic failures. (Evidence: internal/app/run.go:RunPlain)

Prelude execution failures (treated as fatal): If Pack.Prelude.Code != "", it runs a.Runner.RunPrelude(...); on error it records warnings, then returns fmt.Errorf("prelude failed: %w", err) without writing a report (no finalize() call on this path). (Evidence: internal/app/run.go:RunPlain)

Step execution failures (classified at the step-status layer, optionally fatal at the run layer):

Before executing a step, state is set to StatusRunning with StartedAt, then persisted via a.saveState().

On RunStep error: step status becomes StatusFailed, status.Error = err.Error(), state is persisted, and:

if StopOnFail is true: finalize() runs and RunPlain returns the underlying err;

otherwise it continues to the next step.
This is the core “execution failed” mode in practice. (Evidence: internal/app/run.go:RunPlain)

“Stop-on-fail” semantics create two user-visible failure shapes: with StopOnFail=true, the run returns immediately on the first failed step (after writing the report); with StopOnFail=false, the run returns nil even if one or more steps failed, relying on state/report for operators to discover failures. (Evidence: internal/app/run.go:RunPlain)

Skips are treated as non-fail outcomes, but are not consistently represented in state:

ROI-based skips print Skipping step... and continue without writing any step status (no StatusSkipped assignment here).

Resume-based skips require an existing StatusSuccess entry; they print and continue leaving the prior success intact.
Result: reports/operators may not be able to distinguish “intentionally skipped by ROI filter” vs “never attempted / missing status” unless report generation infers it elsewhere. (Evidence: internal/app/run.go:RunPlain)

Warnings are a parallel channel, not failures: recordWarnings() drains runner warnings and appends them to a.State.Warnings, then persists state. They are not printed directly by App (beyond whatever the runner prints), so visibility is primarily via persisted state/report. (Evidence: internal/app/run.go:recordWarnings, saveState)

State persistence failures are explicitly ignored (silent data-loss risk): saveState() discards errors from state.SaveStateAtomic(...) via _ = .... If state writes fail, the run proceeds and operators may later see stale/missing progress. (Evidence: internal/app/run.go:saveState)

Report generation / write failures are explicitly ignored, but success is always announced: finalize() discards errors from json.MarshalIndent(...) and os.WriteFile(...), yet always prints Report written to ... when ReportPath != "". This can falsely signal success to operators even when no report was produced. (Evidence: internal/app/run.go:finalize)

System-wide error classification exists, but RunPlain largely relies on upstream to use it: internal/errors/errors.go defines sentinel errors (ErrInvalidPack, ErrExecutionFailed, ErrConfigInvalid) and an ExitCode(err) mapping using `errors.Is(
