1) Direct answer (1–8 bullets, evidence-cited)

No dedicated background job system is evident in the Go code shown (no queue/worker process, scheduler, or job runner). The primary execution path appears to be a synchronous CLI (Cobra) that runs work inline rather than enqueuing background tasks. Evidence: CLI entry is Cobra-based with flags (internal/cli/root.go), and execution uses context + os/exec to run an external command (internal/exec/runner.go).

Because there’s no background job substrate, classic reliability controls (DLQ/dead-lettering, queue backoff policies, concurrency limits, reprocessing UI) are not implemented as first-class “background work” features—failures would surface as immediate CLI errors rather than being quarantined/retried by a worker. Evidence: inline execution model suggested by os/exec runner (internal/exec/runner.go) and lack of worker constructs in the CLI surface (internal/cli/root.go).

What does exist is “run reliability” via status/state tracking, which is the closest analogue to reprocessing controls: the system models step/run outcomes with explicit statuses (pending/running/success/failed/skipped). Evidence: Status* constants in internal/state/types.go.

State/report paths are configurable at the app level, which suggests persistence of outcomes to disk (useful for restart/manual re-run), but this is not the same as automated retries/backoff. Evidence: app-wide config includes StatePath and Report... fields (internal/app/app.go), and run orchestration imports internal/state + internal/report (internal/app/run.go snippet).

Concurrency limiting is implicitly “1 run = 1 process” (single CLI invocation) unless internal/app/run.go spawns goroutines (not shown). No explicit worker pool/limit controls are visible in the provided excerpts. Evidence: only synchronous CLI + runner excerpts provided (internal/cli/root.go, internal/exec/runner.go, partial internal/app/run.go).

2) Risks/unknowns (bullets)

internal/app/run.go is truncated; it could still spawn goroutines, use tickers, or implement internal retry/backoff that isn’t visible in the snippet.

The persistence semantics (when state is flushed, atomicity, resume behavior) can’t be confirmed without the full implementations of internal/state/* and internal/report/*.

There may be background work outside the shown files (e.g., separate cmd/* binaries, build tags, or a different package) that wasn’t included in the excerpts.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background-work primitives and libraries, e.g. rg -n "(time\\.NewTicker|time\\.Tick|cron|scheduler|queue|worker|goroutine|go\\s+func|asynq|machinery|temporal|celery|bull)" .

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/app/run.go (full file)

internal/state/*.go

internal/report/*.go

cmd/** and internal/cli/*.go (full set, not snippets)

Any directories named or resembling jobs/, worker/, background/, scheduler/, queue/, cron/ (if present)
