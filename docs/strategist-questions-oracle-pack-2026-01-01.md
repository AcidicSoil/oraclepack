# strategist questions — oracle pack

---

## run (exactly 20)

```bash
# shellcheck shell=bash
set -euo pipefail

out_dir="artifacts/oracle/strategist-questions/2026-01-01"
mkdir -p "$out_dir"

# 01) ROI=2.45 | Immediate | contracts/interfaces | internal/pack/types.go:Pack.WriteOutput
oracle \
  --files-report \
  --write-output "$out_dir/01-immediate-contracts-writeoutput-usage.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: contracts/interfaces.\nReference: internal/pack/types.go:Pack.WriteOutput\nQuestion: Is Pack.WriteOutput ever consumed to enforce --write-output behavior, and if not should it be wired into execution or removed to avoid dead config.\nRationale: Dead config fields create ambiguous pack semantics for users.\nSmallest experiment today: Run "ck --regex WriteOutput internal" to confirm all usages.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/pack/types.go" \
  -f "internal/pack/parser.go"

# 02) ROI=1.80 | Immediate | invariants | internal/cli/run.go:runCmd
oracle \
  --files-report \
  --write-output "$out_dir/02-immediate-invariants-roi-mode-validation.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: invariants.\nReference: internal/cli/run.go:runCmd\nQuestion: Should runCmd validate roi-mode values (over or under) and return an error on invalid input instead of silently defaulting.\nRationale: Invalid values can change step filtering without a clear error.\nSmallest experiment today: Review roiMode comparisons in internal/app/run.go.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/cli/run.go" \
  -f "internal/app/run.go"

# 03) ROI=1.63 | Immediate | observability | internal/report/generate.go:GenerateReport
oracle \
  --files-report \
  --write-output "$out_dir/03-immediate-observability-report-ordering.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: observability.\nReference: internal/report/generate.go:GenerateReport\nQuestion: Should GenerateReport sort StepReport entries by step ID to keep report JSON deterministic.\nRationale: Map iteration order is nondeterministic and complicates diffs.\nSmallest experiment today: Inspect state.RunState.StepStatuses type and where it is populated.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/report/generate.go" \
  -f "internal/state/types.go"

# 04) ROI=1.60 | Immediate | caching/state | internal/state/persist.go:SaveStateAtomic
oracle \
  --files-report \
  --write-output "$out_dir/04-immediate-state-save-errors.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: caching/state.\nReference: internal/state/persist.go:SaveStateAtomic\nQuestion: Should app.saveState surface SaveStateAtomic errors to the user instead of ignoring them.\nRationale: Silent write failures can break resume correctness.\nSmallest experiment today: Check internal/app/run.go saveState implementation for error handling.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/state/persist.go" \
  -f "internal/app/run.go"

# 05) ROI=1.40 | Immediate | invariants | internal/pack/parser.go:Validate
oracle \
  --files-report \
  --write-output "$out_dir/05-immediate-invariants-step-sequencing.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: invariants.\nReference: internal/pack/parser.go:Validate\nQuestion: Is strict sequential numbering required for execution and resume, or can packs allow gaps without breaking behavior.\nRationale: Relaxed numbering increases compatibility with existing packs.\nSmallest experiment today: Search for step.Number or step.ID assumptions in internal/ code.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/pack/parser.go"

# 06) ROI=1.20 | Immediate | failure modes | internal/exec/inject.go:InjectFlags
oracle \
  --files-report \
  --write-output "$out_dir/06-immediate-failure-injectflags-regex.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: failure modes.\nReference: internal/exec/inject.go:InjectFlags\nQuestion: Does the oracle command regex risk matching lines like oraclepack or commented oracle, and should it enforce a word boundary.\nRationale: Incorrect flag injection can mutate scripts.\nSmallest experiment today: Review inject_test.go for edge case coverage.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/exec/inject.go" \
  -f "internal/exec/inject_test.go"

# 07) ROI=1.05 | Immediate | failure modes | internal/exec/runner.go:run
oracle \
  --files-report \
  --write-output "$out_dir/07-immediate-failure-runner-exit.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: failure modes.\nReference: internal/exec/runner.go:run\nQuestion: Should Runner.run capture exit code or stderr separately so reports can include richer failure details.\nRationale: Better error metadata speeds debugging.\nSmallest experiment today: Inspect runner_test.go and report generation to see what is recorded today.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/exec/runner.go" \
  -f "internal/errors/errors.go"

# 08) ROI=1.00 | Immediate | UX flows | internal/tui/tui.go:Update
oracle \
  --files-report \
  --write-output "$out_dir/08-immediate-ux-roi-filter.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: UX flows.\nReference: internal/tui/tui.go:Update\nQuestion: Should the ROI filter UI expose the mode (over or under) and an explicit clear action to avoid confusion.\nRationale: Filtering affects which steps execute and can surprise users.\nSmallest experiment today: Trace key handling for f, enter, and esc in Update.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/tui/tui.go"

# 09) ROI=0.83 | Immediate | failure modes | internal/errors/errors.go:ExitCode
oracle \
  --files-report \
  --write-output "$out_dir/09-immediate-failure-exitcode-mapping.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: failure modes.\nReference: internal/errors/errors.go:ExitCode\nQuestion: Are ExitCode mappings sufficient for CI and scripting, or should context cancel and config errors map to distinct codes.\nRationale: Stable exit codes improve automation reliability.\nSmallest experiment today: Trace Execute error handling in internal/cli/root.go.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/errors/errors.go" \
  -f "internal/cli/root.go"

# 10) ROI=0.80 | Immediate | UX flows | internal/render/render.go:RenderMarkdown
oracle \
  --files-report \
  --write-output "$out_dir/10-immediate-ux-markdown-wrap.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: UX flows.\nReference: internal/render/render.go:RenderMarkdown\nQuestion: Should RenderMarkdown use dynamic width from the TUI instead of fixed 80 to avoid truncation.\nRationale: Fixed wrapping can reduce readability on narrow or wide terminals.\nSmallest experiment today: Search for RenderMarkdown usage and any width hints in TUI code.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/render/render.go"

# 11) ROI=0.75 | Immediate | caching/state | internal/state/types.go:RunState.SchemaVersion
oracle \
  --files-report \
  --write-output "$out_dir/11-immediate-state-schema-version.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: caching/state.\nReference: internal/state/types.go:RunState.SchemaVersion\nQuestion: Should LoadState reject or migrate mismatched SchemaVersion values.\nRationale: Schema drift can corrupt resume behavior and reports.\nSmallest experiment today: Inspect LoadState in internal/state/persist.go for version checks.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/state/types.go" \
  -f "internal/state/persist.go"

# 12) ROI=0.72 | Strategic | invariants | internal/app/app.go:Prepare
oracle \
  --files-report \
  --write-output "$out_dir/12-strategic-invariants-outdir-precedence.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: invariants.\nReference: internal/app/app.go:Prepare\nQuestion: Is the out_dir precedence (CLI over pack over default) the intended long term contract, especially with state and report paths derived from pack name.\nRationale: OutDir semantics affect reproducibility and user expectations.\nSmallest experiment today: Trace out_dir derivation in cli/run.go and app.Prepare.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/app/app.go" \
  -f "internal/cli/run.go" \
  -f "internal/pack/parser.go"

# 13) ROI=0.63 | Strategic | observability | internal/app/run.go:finalize
oracle \
  --files-report \
  --write-output "$out_dir/13-strategic-observability-report-write.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: observability.\nReference: internal/app/run.go:finalize\nQuestion: Should finalize surface report write failures or add logging so users know reports failed.\nRationale: Silent report write failures hide execution outcomes.\nSmallest experiment today: Review finalize and report.GenerateReport for error handling.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/app/run.go" \
  -f "internal/report/generate.go"

# 14) ROI=0.60 | Immediate | background jobs | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/14-immediate-background-jobs-unknown.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Immediate. Category: background jobs.\nReference: Unknown.\nMissing artifact pattern: /jobs/, /workers/, /queue/, /cron/, *scheduler*.* (no matches in repo scan).\nQuestion: Are there any background job or async execution mechanisms, or is all execution synchronous in the runner loop.\nRationale: Async execution changes failure handling and state persistence guarantees.\nSmallest experiment today: Run "fd -i \"job|worker|queue|cron|scheduler\" -d 6" to confirm absence.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "README.md" \
  -f "internal/app/run.go"

# 15) ROI=0.58 | Strategic | contracts/interfaces | internal/report/types.go:ReportV1
oracle \
  --files-report \
  --write-output "$out_dir/15-strategic-contracts-report-versioning.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: contracts/interfaces.\nReference: internal/report/types.go:ReportV1\nQuestion: Should ReportV1 include a schema version field to allow forward compatible parsing.\nRationale: Versioned reports enable safe tooling updates.\nSmallest experiment today: Search for report consumers or tests that assume a fixed schema.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/report/types.go" \
  -f "internal/report/generate.go"

# 16) ROI=0.50 | Strategic | failure modes | internal/exec/runner.go:NewRunner
oracle \
  --files-report \
  --write-output "$out_dir/16-strategic-failure-shell-default.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: failure modes.\nReference: internal/exec/runner.go:NewRunner\nQuestion: Is defaulting to /bin/bash correct for cross platform use, or should it be configurable per OS or via flag.\nRationale: Defaulting to /bin/bash can fail on Windows and minimal containers.\nSmallest experiment today: Compare README install notes with runner.NewRunner defaults.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/exec/runner.go" \
  -f "README.md"

# 17) ROI=0.45 | Strategic | permissions | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/17-strategic-permissions-unknown.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: permissions.\nReference: Unknown.\nMissing artifact pattern: /auth/, *permission*.*, *policy*.*, *rbac*.* (no matches in repo scan).\nQuestion: Does the tool need any permission or allowlist layer beyond user consent, or is none intended.\nRationale: Execution tools may need guardrails in shared environments.\nSmallest experiment today: Run "fd -i \"auth|rbac|permission|policy|acl\" -d 6" to confirm absence.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "README.md" \
  -f "internal/exec/runner.go"

# 18) ROI=0.40 | Strategic | feature flags | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/18-strategic-feature-flags-unknown.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: feature flags.\nReference: Unknown.\nMissing artifact pattern: *flag*.*, *toggle*.*, /features/ (no matches in repo scan).\nQuestion: Is there a feature flag or toggle system beyond CLI flags, or should behavior changes remain explicit flags only.\nRationale: Feature flags affect rollout strategy and compatibility.\nSmallest experiment today: Run "fd -i \"flag|toggle|feature\" -d 6" and review CLI flag definitions.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "README.md" \
  -f "internal/cli/root.go"

# 19) ROI=0.40 | Strategic | UX flows | internal/exec/stream.go:LineWriter
oracle \
  --files-report \
  --write-output "$out_dir/19-strategic-ux-linewriter-buffer.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: UX flows.\nReference: internal/exec/stream.go:LineWriter\nQuestion: Should LineWriter add a max buffer or backpressure to avoid memory growth on very large outputs.\nRationale: Long running commands can emit large output and exhaust memory.\nSmallest experiment today: Trace LineWriter usage and log channel buffering in tui.Model.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "internal/exec/stream.go" \
  -f "internal/tui/tui.go"

# 20) ROI=0.30 | Strategic | migrations | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/20-strategic-migrations-unknown.md" \
  -p $'Repo: oraclepack - Go CLI/TUI (Cobra + Bubble Tea). Entrypoint cmd/oraclepack/main.go; core logic in internal/.\nConstraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown.\nHorizon: Strategic. Category: migrations.\nReference: Unknown.\nMissing artifact pattern: /migrations/, schema.*, /prisma/, /alembic/ (no matches in repo scan).\nQuestion: Is there any persistent storage that needs migrations, or is all state file based only.\nRationale: Migration strategy influences deployment and compatibility.\nSmallest experiment today: Run "fd -i \"migrat|schema|prisma|alembic|knex|gorm\" -d 6" to confirm absence.\nAnswer format:\n- Reference\n- Direct answer\n- Evidence (cite attached paths and symbols)\n- Smallest experiment today\n- Risks or assumptions (optional, max 3 bullets)' \
  -f "README.md" \
  -f "internal/state/persist.go"
```
