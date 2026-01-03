# oracle strategist question pack

---

## parsed args

- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown
- out_dir: oracle-out
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: empty

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=1.75 impact=0.5 confidence=0.7 effort=0.2 horizon=Immediate category=observability reference=internal/report/generate.go:GenerateReport
oracle \
  --files-report \
  --write-output "oracle-out/01-observability-generate-report.md" \
  -p "Strategist question #01
Reference: internal/report/generate.go:GenerateReport
Category: observability
Horizon: Immediate
ROI: 1.75 (impact=0.5, confidence=0.7, effort=0.2)
Question: Should `GenerateReport` produce a deterministic step order (e.g., sorted by step ID) so downstream tooling can diff reports reliably?
Rationale: Reports are the primary audit artifact and nondeterministic ordering makes automation brittle.
Smallest experiment today: Review the map iteration in `GenerateReport` and prototype sorting by step ID.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/report/generate.go"

# 02 — ROI=1.50 impact=0.5 confidence=0.6 effort=0.2 horizon=Immediate category=UX flows reference=internal/tui/tui.go:refreshList
oracle \
  --files-report \
  --write-output "oracle-out/02-ux-flows-roi-filter.md" \
  -p "Strategist question #02
Reference: internal/tui/tui.go:refreshList
Category: UX flows
Horizon: Immediate
ROI: 1.50 (impact=0.5, confidence=0.6, effort=0.2)
Question: Does the ROI filter behave as intended for `under` vs `over` modes and for steps with `ROI==0`?
Rationale: Filtering is a core interaction in large packs and off-by-one logic will confuse users.
Smallest experiment today: Inspect `refreshList` and trace the ROI comparisons for zero-valued ROI.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/tui/tui.go"

# 03 — ROI=1.40 impact=0.7 confidence=0.6 effort=0.3 horizon=Immediate category=invariants reference=internal/pack/parser.go:Validate
oracle \
  --files-report \
  --write-output "oracle-out/03-invariants-validate-steps.md" \
  -p "Strategist question #03
Reference: internal/pack/parser.go:Validate
Category: invariants
Horizon: Immediate
ROI: 1.40 (impact=0.7, confidence=0.6, effort=0.3)
Question: Does `Validate` correctly allow all supported header separators while still enforcing strict sequential numbering?
Rationale: Validation is the gatekeeper for running packs and overly strict rules can reject valid files.
Smallest experiment today: Add a unit test in `internal/pack/parser_test.go` for mixed header separators and sequential numbering.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/pack/parser.go"

# 04 — ROI=1.25 impact=0.5 confidence=0.5 effort=0.2 horizon=Immediate category=caching/state reference=internal/state/types.go:RunState
oracle \
  --files-report \
  --write-output "oracle-out/04-caching-state-schema-version.md" \
  -p "Strategist question #04
Reference: internal/state/types.go:RunState
Category: caching/state
Horizon: Immediate
ROI: 1.25 (impact=0.5, confidence=0.5, effort=0.2)
Question: Is `SchemaVersion` enforced anywhere on resume, and should mismatches block or migrate state loading?
Rationale: State schema drift can silently corrupt resume behavior.
Smallest experiment today: Search for `SchemaVersion` usage and note whether load-time validation exists.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/state/types.go" \
  -f "internal/state/persist.go"

# 05 — ROI=1.20 impact=0.4 confidence=0.6 effort=0.2 horizon=Immediate category=observability reference=internal/exec/stream.go:LineWriter.Write
oracle \
  --files-report \
  --write-output "oracle-out/05-observability-linewriter-flush.md" \
  -p "Strategist question #05
Reference: internal/exec/stream.go:LineWriter.Write
Category: observability
Horizon: Immediate
ROI: 1.20 (impact=0.4, confidence=0.6, effort=0.2)
Question: Do log lines flush correctly when output lacks a trailing newline?
Rationale: Missing final lines makes debugging long runs harder.
Smallest experiment today: Review `LineWriter.Write`/`Close` to confirm buffered output is always emitted.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/exec/stream.go"

# 06 — ROI=1.20 impact=0.6 confidence=0.6 effort=0.3 horizon=Immediate category=caching/state reference=internal/state/persist.go:SaveStateAtomic
oracle \
  --files-report \
  --write-output "oracle-out/06-caching-state-atomic-save.md" \
  -p "Strategist question #06
Reference: internal/state/persist.go:SaveStateAtomic
Category: caching/state
Horizon: Immediate
ROI: 1.20 (impact=0.6, confidence=0.6, effort=0.3)
Question: Is `SaveStateAtomic` safe across different filesystems when `statePath` is set outside the pack directory?
Rationale: Atomic renames are only guaranteed on the same filesystem.
Smallest experiment today: Verify that temp files are created alongside `statePath` and document any cross-filesystem risks.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/state/persist.go"

# 07 — ROI=1.20 impact=0.6 confidence=0.6 effort=0.3 horizon=Immediate category=failure modes reference=internal/exec/runner.go:run
oracle \
  --files-report \
  --write-output "oracle-out/07-failure-modes-runner-run.md" \
  -p "Strategist question #07
Reference: internal/exec/runner.go:run
Category: failure modes
Horizon: Immediate
ROI: 1.20 (impact=0.6, confidence=0.6, effort=0.3)
Question: Does `Runner.run` surface context cancellations distinctly from command failures and preserve the original exit status?
Rationale: Clear failure classification speeds up diagnosis and retry logic.
Smallest experiment today: Trace the error wrapping paths in `Runner.run` for context vs exec errors.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/exec/runner.go"

# 08 — ROI=1.00 impact=0.5 confidence=0.6 effort=0.3 horizon=Immediate category=failure modes reference=internal/errors/errors.go:ExitCode
oracle \
  --files-report \
  --write-output "oracle-out/08-failure-modes-exit-code.md" \
  -p "Strategist question #08
Reference: internal/errors/errors.go:ExitCode
Category: failure modes
Horizon: Immediate
ROI: 1.00 (impact=0.5, confidence=0.6, effort=0.3)
Question: Are all user-visible failures mapped to stable exit codes that CI can rely on?
Rationale: Exit codes are the primary contract for automation.
Smallest experiment today: Trace error returns from CLI/app layers to see which errors reach `ExitCode`.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/errors/errors.go"

# 09 — ROI=1.00 impact=0.6 confidence=0.5 effort=0.3 horizon=Immediate category=invariants reference=internal/pack/parser.go:Parse
oracle \
  --files-report \
  --write-output "oracle-out/09-invariants-stepheader-regex.md" \
  -p "Strategist question #09
Reference: internal/pack/parser.go:Parse
Category: invariants
Horizon: Immediate
ROI: 1.00 (impact=0.6, confidence=0.5, effort=0.3)
Question: Can `stepHeaderRegex` accidentally interpret commented lines inside bash code as new steps?
Rationale: Mis-parsing steps can execute unintended commands.
Smallest experiment today: Add a pack example with a bash comment containing `# 01)` and observe parse behavior.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/pack/parser.go"

# 10 — ROI=0.83 impact=0.5 confidence=0.5 effort=0.3 horizon=Immediate category=UX flows reference=internal/cli/run.go:runCmd
oracle \
  --files-report \
  --write-output "oracle-out/10-ux-flows-run-paths.md" \
  -p "Strategist question #10
Reference: internal/cli/run.go:runCmd
Category: UX flows
Horizon: Immediate
ROI: 0.83 (impact=0.5, confidence=0.5, effort=0.3)
Question: Should `statePath` and `reportPath` be resolved relative to the pack file rather than the current working directory?
Rationale: Users expect outputs to live alongside the pack they ran.
Smallest experiment today: Follow `runCmd` path construction and note the directory used for outputs.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/cli/run.go" \
  -f "internal/app/run.go"

# 11 — ROI=0.75 impact=0.6 confidence=0.5 effort=0.4 horizon=Immediate category=contracts/interfaces reference=internal/app/app.go:Config
oracle \
  --files-report \
  --write-output "oracle-out/11-contracts-config-propagation.md" \
  -p "Strategist question #11
Reference: internal/app/app.go:Config
Category: contracts/interfaces
Horizon: Immediate
ROI: 0.75 (impact=0.6, confidence=0.5, effort=0.4)
Question: Is the `Config` contract missing any CLI/runtime inputs that should be propagated to `exec.Runner` (e.g., `oracle-bin`, per-step flags)?
Rationale: Gaps between CLI flags and runtime config lead to surprising behavior.
Smallest experiment today: Trace which CLI flags are copied into `app.Config` and `exec.Runner`.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/app/app.go" \
  -f "internal/cli/run.go"

# 12 — ROI=0.75 impact=0.6 confidence=0.5 effort=0.4 horizon=Immediate category=UX flows reference=internal/tui/tui.go:Model.Update
oracle \
  --files-report \
  --write-output "oracle-out/12-ux-flows-viewdone-keys.md" \
  -p "Strategist question #12
Reference: internal/tui/tui.go:Model.Update
Category: UX flows
Horizon: Immediate
ROI: 0.75 (impact=0.6, confidence=0.5, effort=0.4)
Question: Do `r`, `b`, and `n` key paths leave the model in a consistent state with cleared logs and correct selection?
Rationale: Inconsistent state transitions cause confusing reruns.
Smallest experiment today: Map the `ViewDone` key handling branch in `Model.Update`.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/tui/tui.go"

# 13 — ROI=0.56 impact=0.7 confidence=0.4 effort=0.5 horizon=Strategic category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "oracle-out/13-permissions-policy-missing.md" \
  -p "Strategist question #13
Reference: Unknown
Category: permissions
Horizon: Strategic
ROI: 0.56 (impact=0.7, confidence=0.4, effort=0.5)
Question: Should running pack steps require an explicit permission/allowlist policy beyond the existing `--yes` flag?
Rationale: No permission artifacts were found (missing `**/policy/**` or `**/permissions/**`), so safety expectations are unclear.
Smallest experiment today: Run `ck --sem \"permission allowlist confirm\" .` to check for policy files.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "README.md" \
  -f "go.mod" \
  -f "cmd/oraclepack/main.go"

# 14 — ROI=0.48 impact=0.6 confidence=0.4 effort=0.5 horizon=Strategic category=caching/state reference=internal/state/persist.go:LoadState
oracle \
  --files-report \
  --write-output "oracle-out/14-caching-state-migrations.md" \
  -p "Strategist question #14
Reference: internal/state/persist.go:LoadState
Category: caching/state
Horizon: Strategic
ROI: 0.48 (impact=0.6, confidence=0.4, effort=0.5)
Question: Do we need a formal state migration path for future schema changes to keep `--resume` reliable?
Rationale: As features evolve, resume compatibility becomes a long-term risk.
Smallest experiment today: Sketch a `SchemaVersion` switch in `LoadState` for forward/backward compatibility.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/state/persist.go" \
  -f "internal/state/types.go"

# 15 — ROI=0.48 impact=0.6 confidence=0.4 effort=0.5 horizon=Strategic category=UX flows reference=internal/tui/tui.go:Model.Init
oracle \
  --files-report \
  --write-output "oracle-out/15-ux-flows-batch-mode.md" \
  -p "Strategist question #15
Reference: internal/tui/tui.go:Model.Init
Category: UX flows
Horizon: Strategic
ROI: 0.48 (impact=0.6, confidence=0.4, effort=0.5)
Question: Should the TUI support a non-interactive batch mode with step selection to scale large packs?
Rationale: As packs grow, batch workflows could become primary usage.
Smallest experiment today: Review how `run-all` is wired between CLI and `Model.Init` to assess feasibility of persisted selections.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/tui/tui.go" \
  -f "internal/cli/run.go"

# 16 — ROI=0.47 impact=0.7 confidence=0.4 effort=0.6 horizon=Strategic category=contracts/interfaces reference=internal/pack/types.go:Pack
oracle \
  --files-report \
  --write-output "oracle-out/16-contracts-pack-version.md" \
  -p "Strategist question #16
Reference: internal/pack/types.go:Pack
Category: contracts/interfaces
Horizon: Strategic
ROI: 0.47 (impact=0.7, confidence=0.4, effort=0.6)
Question: Should the pack format introduce an explicit version/schema field to enable backward-compatible feature evolution?
Rationale: Versioned contracts prevent breaking older packs when new fields are added.
Smallest experiment today: Draft a minimal `Pack.Version` field and list compatibility rules.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/pack/types.go"

# 17 — ROI=0.40 impact=0.6 confidence=0.4 effort=0.6 horizon=Strategic category=observability reference=internal/report/types.go:ReportV1
oracle \
  --files-report \
  --write-output "oracle-out/17-observability-report-telemetry.md" \
  -p "Strategist question #17
Reference: internal/report/types.go:ReportV1
Category: observability
Horizon: Strategic
ROI: 0.40 (impact=0.6, confidence=0.4, effort=0.6)
Question: What additional telemetry fields (e.g., oracle command count, stderr size) would materially improve report usefulness?
Rationale: Strategic decisions need richer, consistent metrics than today’s summary.
Smallest experiment today: List candidate fields and identify which module would populate each.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "internal/report/types.go"

# 18 — ROI=0.30 impact=0.5 confidence=0.3 effort=0.5 horizon=Strategic category=feature flags reference=Unknown
oracle \
  --files-report \
  --write-output "oracle-out/18-feature-flags-missing.md" \
  -p "Strategist question #18
Reference: Unknown
Category: feature flags
Horizon: Strategic
ROI: 0.30 (impact=0.5, confidence=0.3, effort=0.5)
Question: Should experimental features (ROI filter, run-all, overrides) be gated behind a feature-flag mechanism?
Rationale: No feature-flag system was found (missing `**/flags/**`), which limits safe rollout.
Smallest experiment today: Run `ck --sem \"feature flag toggle\" .` and note any flag infrastructure.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "README.md" \
  -f "go.mod" \
  -f "cmd/oraclepack/main.go"

# 19 — ROI=0.24 impact=0.4 confidence=0.3 effort=0.5 horizon=Strategic category=migrations reference=Unknown
oracle \
  --files-report \
  --write-output "oracle-out/19-migrations-missing.md" \
  -p "Strategist question #19
Reference: Unknown
Category: migrations
Horizon: Strategic
ROI: 0.24 (impact=0.4, confidence=0.3, effort=0.5)
Question: Is there a migration strategy for pack/report/state schema changes, or are breaking changes acceptable?
Rationale: No migration artifacts were found (missing `**/migrations/**`), so upgrade expectations are unclear.
Smallest experiment today: Search for `schema_version` or `migration` mentions across the repo to confirm absence.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "README.md" \
  -f "go.mod" \
  -f "cmd/oraclepack/main.go"

# 20 — ROI=0.20 impact=0.4 confidence=0.3 effort=0.6 horizon=Strategic category=background jobs reference=Unknown
oracle \
  --files-report \
  --write-output "oracle-out/20-background-jobs-missing.md" \
  -p "Strategist question #20
Reference: Unknown
Category: background jobs
Horizon: Strategic
ROI: 0.20 (impact=0.4, confidence=0.3, effort=0.6)
Question: Would a background job runner be valuable for long-running validation or preflight tasks?
Rationale: No job/worker artifacts were found (missing `**/jobs/**`), so long-running tasks may block the TUI.
Smallest experiment today: Run `ck --sem \"queue worker cron\" .` to confirm whether any job system exists.
Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "README.md" \
  -f "go.mod" \
  -f "cmd/oraclepack/main.go"
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: OK

*   invariants: OK

*   caching/state: OK

*   background jobs: Missing (no job artifacts found; attach `**/jobs/**` if present)

*   observability: OK

*   permissions: Missing (no policy artifacts found; attach `**/policy/**` or `**/permissions/**` if present)

*   migrations: Missing (no migration artifacts found; attach `**/migrations/**` if present)

*   UX flows: OK

*   failure modes: OK

*   feature flags: Missing (no flag artifacts found; attach `**/flags/**` if present)
