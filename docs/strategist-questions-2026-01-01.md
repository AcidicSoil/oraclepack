# Strategist Questions (Evidence-Cited, ROI-Ranked)

## Parsed args

- codebase_name: Unknown
- constraints: Unknown
- non_goals: Unknown
- team_size: Unknown
- deadline: Unknown

## Questions (exactly 20; sorted by ROI desc; ties by lower effort)

Counts must be: 12 Immediate Exploration, 8 Strategic Planning.

| # | ROI | impact | confidence | effort | Horizon | Category | Reference | Question | Rationale (1 sentence) | Smallest experiment today (1 action) |
|---:|---:|---:|---:|---:|---|---|---|---|---|---|
| 1 | 2.45 | 0.7 | 0.7 | 0.2 | Immediate | contracts/interfaces | internal/pack/types.go:Pack.WriteOutput | Is Pack.WriteOutput ever consumed to enforce --write-output behavior, and if not should it be wired into execution or removed. | Dead config fields create ambiguous pack semantics for users. | Run `ck --regex "WriteOutput" internal` to confirm all usages. |
| 2 | 1.80 | 0.6 | 0.6 | 0.2 | Immediate | invariants | internal/cli/run.go:runCmd | Should runCmd validate roi-mode values (over or under) and return an error on invalid input instead of silently defaulting. | Invalid values can change step filtering without a clear error. | Trace roiMode comparisons in `internal/app/run.go`. |
| 3 | 1.63 | 0.7 | 0.7 | 0.3 | Immediate | observability | internal/report/generate.go:GenerateReport | Should GenerateReport sort StepReport entries by step ID to keep report JSON deterministic. | Map iteration order is nondeterministic and complicates diffs. | Inspect StepStatuses map usage in `internal/state/types.go`. |
| 4 | 1.60 | 0.8 | 0.6 | 0.3 | Immediate | caching/state | internal/state/persist.go:SaveStateAtomic | Should app.saveState surface SaveStateAtomic errors to the user instead of ignoring them. | Silent write failures can break resume correctness. | Check `saveState` in `internal/app/run.go` for error handling. |
| 5 | 1.40 | 0.7 | 0.6 | 0.3 | Immediate | invariants | internal/pack/parser.go:Validate | Is strict sequential numbering required for execution and resume, or can packs allow gaps without breaking behavior. | Relaxed numbering increases compatibility with existing packs. | Search for step number assumptions in `internal/`. |
| 6 | 1.20 | 0.6 | 0.6 | 0.3 | Immediate | failure modes | internal/exec/inject.go:InjectFlags | Does the oracle command regex risk matching lines like oraclepack or commented oracle, and should it enforce a word boundary. | Incorrect flag injection can mutate scripts. | Review `internal/exec/inject_test.go` for edge case coverage. |
| 7 | 1.05 | 0.7 | 0.6 | 0.4 | Immediate | failure modes | internal/exec/runner.go:run | Should Runner.run capture exit code or stderr separately so reports can include richer failure details. | Better error metadata speeds debugging. | Inspect `internal/report/generate.go` for error fields. |
| 8 | 1.00 | 0.5 | 0.6 | 0.3 | Immediate | UX flows | internal/tui/tui.go:Update | Should the ROI filter UI expose the mode (over or under) and an explicit clear action to avoid confusion. | Filtering affects which steps execute and can surprise users. | Trace key handling for `f`, `enter`, and `esc` in `Update`. |
| 9 | 0.83 | 0.5 | 0.5 | 0.3 | Immediate | failure modes | internal/errors/errors.go:ExitCode | Are ExitCode mappings sufficient for CI and scripting, or should context cancel and config errors map to distinct codes. | Stable exit codes improve automation reliability. | Follow error handling from `internal/cli/root.go` into `ExitCode`. |
| 10 | 0.80 | 0.4 | 0.6 | 0.3 | Immediate | UX flows | internal/render/render.go:RenderMarkdown | Should RenderMarkdown use dynamic width from the TUI instead of fixed 80 to avoid truncation. | Fixed wrapping can reduce readability on narrow or wide terminals. | Search for RenderMarkdown callers in `internal/`. |
| 11 | 0.75 | 0.6 | 0.5 | 0.4 | Immediate | caching/state | internal/state/types.go:RunState.SchemaVersion | Should LoadState reject or migrate mismatched SchemaVersion values. | Schema drift can corrupt resume behavior and reports. | Inspect `internal/state/persist.go` for version checks. |
| 12 | 0.72 | 0.6 | 0.6 | 0.5 | Strategic | invariants | internal/app/app.go:Prepare | Is the out_dir precedence (CLI over pack over default) the intended long term contract, especially with state and report paths derived from pack name. | OutDir semantics affect reproducibility and user expectations. | Trace out_dir derivation in `internal/cli/run.go` and `internal/app/app.go`. |
| 13 | 0.63 | 0.5 | 0.5 | 0.4 | Strategic | observability | internal/app/run.go:finalize | Should finalize surface report write failures or add logging so users know reports failed. | Silent report write failures hide execution outcomes. | Review `finalize` and report generation in `internal/app/run.go`. |
| 14 | 0.60 | 0.3 | 0.4 | 0.2 | Immediate | background jobs | Unknown | Are there any background job or async execution mechanisms, or is all execution synchronous in the runner loop. | No /jobs/, /workers/, /queue/, /cron/, or *scheduler*.* paths were found in the repo scan. | Run `fd -i "job|worker|queue|cron|scheduler" -d 6` to confirm absence. |
| 15 | 0.58 | 0.7 | 0.5 | 0.6 | Strategic | contracts/interfaces | internal/report/types.go:ReportV1 | Should ReportV1 include a schema version field to allow forward compatible parsing. | Versioned reports enable safe tooling updates. | Search for report consumers or schema assumptions in `internal/report`. |
| 16 | 0.50 | 0.6 | 0.5 | 0.6 | Strategic | failure modes | internal/exec/runner.go:NewRunner | Is defaulting to /bin/bash correct for cross platform use, or should it be configurable per OS or via flag. | Defaulting to /bin/bash can fail on Windows and minimal containers. | Compare README install notes with `NewRunner` defaults. |
| 17 | 0.45 | 0.3 | 0.3 | 0.2 | Strategic | permissions | Unknown | Does the tool need any permission or allowlist layer beyond user consent, or is none intended. | No /auth/, *permission*.*, *policy*.*, or *rbac*.* paths were found in the repo scan. | Run `fd -i "auth|rbac|permission|policy|acl" -d 6` to confirm absence. |
| 18 | 0.40 | 0.4 | 0.3 | 0.3 | Strategic | feature flags | Unknown | Is there a feature flag or toggle system beyond CLI flags, or should behavior changes remain explicit flags only. | No *flag*.*, *toggle*.*, or /features/ paths were found in the repo scan. | Run `fd -i "flag|toggle|feature" -d 6` and review CLI flags. |
| 19 | 0.40 | 0.4 | 0.5 | 0.5 | Strategic | UX flows | internal/exec/stream.go:LineWriter | Should LineWriter add a max buffer or backpressure to avoid memory growth on very large outputs. | Long running commands can emit large output and exhaust memory. | Trace LineWriter usage and log channel buffering in `internal/tui/tui.go`. |
| 20 | 0.30 | 0.3 | 0.2 | 0.2 | Strategic | migrations | Unknown | Is there any persistent storage that needs migrations, or is all state file based only. | No /migrations/, schema.*, /prisma/, or /alembic/ paths were found in the repo scan. | Run `fd -i "migrat|schema|prisma|alembic|knex|gorm" -d 6` to confirm absence. |

## Coverage check (must be satisfied)

- contracts/interfaces: OK
- invariants: OK
- caching/state: OK
- background jobs: Missing (/jobs/, /workers/, /queue/, /cron/, *scheduler*.* not found)
- observability: OK
- permissions: Missing (/auth/, *permission*.*, *policy*.*, *rbac*.* not found)
- migrations: Missing (/migrations/, schema.*, /prisma/, /alembic/ not found)
- UX flows: OK
- failure modes: OK
- feature flags: Missing (*flag*.*, *toggle*.*, /features/ not found)
