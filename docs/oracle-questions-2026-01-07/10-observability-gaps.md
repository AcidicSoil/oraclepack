1) Direct answer (1–8 bullets, evidence-cited)

No correlation IDs in execution logs, so step output cannot be tied back to a step/run reliably. run() wires cmd.Stdout/cmd.Stderr directly to logWriter with no prefix/header/trailer, and it receives no stepID/runID to stamp lines (run(ctx, script, logWriter) only gets the script + writer). Evidence: RunPrelude/RunStep call r.run(ctx, script, logWriter); run() sets cmd.Stdout = logWriter and cmd.Stderr = logWriter and emits no logs of its own. (internal/exec/runner.go: RunPrelude, RunStep, run)

Key decision points are silent (missing “what config did we actually use?” logs). Shell defaulting (/bin/bash), override application, ChatGPTURL flag application, and flag injection happen without any structured record, making misconfig/debugging difficult. Evidence: NewRunner defaults Shell; RunStep conditionally calls Overrides.EffectiveFlags(...) and ApplyChatGPTURL(...), then InjectFlags(...), with no logging. (internal/exec/runner.go: NewRunner, RunStep)

Errors lose high-value context (step, command, exit code, duration). On failure, run() returns fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err) (or ctx.Err()), which preserves a category but omits step ID, elapsed time, and process exit code/signal—often the first things needed to diagnose. Evidence: run() error path and wrapping; error taxonomy only maps categories to exit codes. (internal/exec/runner.go: run; internal/errors/errors.go: ExitCode)

Sanitizer warnings are emitted as human text only and are not emitted at “decision boundaries.” Warnings are printed as a single formatted line and accumulated in-memory, but there’s no summary count per step/run, no machine-parseable form, and no “this step was sanitized” header at step start/end. Evidence: recordWarnings prints fmt.Fprintf(... "⚠ oraclepack: sanitized label ...") and appends to r.warnings; DrainWarnings() returns them but nothing here enforces they’re logged/serialized elsewhere. (internal/exec/runner.go: recordWarnings, DrainWarnings)

No run/step outcome metrics exist at the point where the outcome is known. The runner is the natural place to emit: step_success/step_failure counts, durations, and warning counts; currently it emits none. Evidence: run() measures no time and publishes no counters; it only runs the command and returns error/nil. (internal/exec/runner.go: run)

Smallest, highest-leverage additions (minimal surface area):

Add a tiny correlation context passed into run() (at least step_id, optionally run_id) and emit two boundary lines: step.start and step.end (include duration_ms, status, exit_code if available). This alone makes raw stdout/stderr interpretable. (Hook point: RunPrelude/RunStep just before/after r.run(...); data already present: s.ID in RunStep.) (internal/exec/runner.go: RunStep, run)

On error, detect *exec.ExitError and include exit_code/`s_*
