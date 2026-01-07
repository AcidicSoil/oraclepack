1) Direct answer (1–8 bullets, evidence-cited)

No explicit background jobs/workers/scheduled tasks are defined in internal/app/run.go. The execution model is a single synchronous control flow in (*App).RunPlain: optional prelude, then a for-loop over steps, each executed via a.Runner.RunStep(...), followed by state updates and writes. Evidence: RunPlain(...) calls RunPrelude(...) and RunStep(...) inline (no goroutines, tickers, or async constructs shown).

Triggers: Only direct, immediate calls from RunPlain (prelude runs if a.Pack.Prelude.Code != ""; steps run when they pass ROI/resume filters). Evidence: prelude conditional and for _, step := range a.Pack.Steps { ... err := a.Runner.RunStep(...) ... }.

Payloads: The “work” payload is the current step (&step) and out writer passed into Runner.RunStep, plus the prelude (&a.Pack.Prelude). Evidence: a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out) and a.Runner.RunStep(ctx, &step, out).

Retries: No retry loop/backoff exists here. The only “try again” behavior is manual restart/resume: successful steps are skipped on subsequent runs. Evidence: resume check if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess { ... continue }.

Idempotency: “At-least-once with resume-skip on success” at the step level. Steps are marked running, persisted, then marked success/failed and persisted; rerun skips StatusSuccess steps. Evidence: status set to StatusRunning + a.saveState(), then on success sets StatusSuccess + a.saveState(), and skip logic above.

Scheduled behavior: None (no timers/cron). The only time usage is stamping StartTime, StartedAt, EndedAt via time.Now(). Evidence: a.State.StartTime = time.Now(), StartedAt: time.Now(), status.EndedAt = time.Now().

2) Risks/unknowns

a.Runner.RunPrelude, a.Runner.RunStep, and a.Runner.DrainWarnings() could hide concurrency (goroutines, worker pools, background process supervision) inside internal/exec or other packages; this file alone cannot confirm that.

state.SaveStateAtomic(...) could include locking/FS semantics or side effects relevant to “background-like” behavior (e.g., fsync patterns), though it is still invoked synchronously here.

CLI/TUI entrypoints may spawn background refreshers (status polling, UI redraw loops) outside run.go.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background/scheduling primitives:
rg -n --hidden "time\.NewTicker|time\.AfterFunc|cron|scheduler|worker|goroutine|go\s+func|errgroup|WaitGroup|queue|async" internal/

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/exec/**/*.go (especially internal/exec/runner.go and any process-management helpers)

internal/cli/**/*.go (entrypoints/TUI loops can introduce background refresh/polling)

internal/state/**/*.go (e.g., SaveStateAtomic implementation)
