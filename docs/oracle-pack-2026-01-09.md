# Oracle Pack — Unknown (Gold Stage 1)

## Parsed args
- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown
- out_dir: docs/oracle-questions-2026-01-09
- oracle_cmd: oracle
- oracle_flags: --files-report
- engine: None
- model: None
- extra_files: (none)

Notes:
- Template is the contract. Keep the pack runner-ingestible.
- Exactly one fenced bash block in this whole document.
- Exactly 20 steps, numbered 01..20.
- Each step includes: ROI= impact= confidence= effort= horizon= category= reference=
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add --dry-run summary to preview what will be sent, and keep --files-report enabled when available.

# 01) ROI=4.2 impact=High confidence=High effort=Low horizon=Immediate category=contracts/interfaces reference=internal/cli/root.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/01-contracts-interfaces-surface.md" \
  -f "README.md" \
  -f "internal/cli/root.go" \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: internal/cli/root.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.2 (impact=High, confidence=High, effort=Low)

Question:
Identify the primary public interface(s) of this system (CLI commands, flags, and pack format surface). For each, list the key inputs/outputs and where they are defined in the code.

Rationale (one sentence):
We need a trustworthy map of the system's outside-facing contract before deeper planning.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=3.6 impact=High confidence=Medium effort=Medium horizon=Immediate category=contracts/interfaces reference=internal/tools/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/02-contracts-interfaces-integration.md" \
  -f "internal/tools/types.go" \
  -f "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: internal/tools/types.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 3.6 (impact=High, confidence=Medium, effort=Medium)

Question:
What are the top integration points with external tools or systems (oracle CLI, shell environment, filesystem, optional task-master/codex/gemini prefixes)? For each, point to the contract or config that declares it.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=3.4 impact=High confidence=Medium effort=Medium horizon=NearTerm category=invariants reference=internal/pack/parser.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/03-invariants-domain.md" \
  -f "internal/pack/parser.go" \
  -f "internal/pack/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: internal/pack/parser.go
Category: invariants
Horizon: NearTerm
ROI: 3.4 (impact=High, confidence=Medium, effort=Medium)

Question:
List the system's most important invariants about pack structure and step metadata (e.g., numbering, ROI parsing, header format). For each, show where it is enforced or where enforcement is missing.

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=3.1 impact=Medium confidence=Medium effort=Medium horizon=NearTerm category=invariants reference=internal/pack/output_check.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/04-invariants-validation.md" \
  -f "internal/pack/output_check.go" \
  -f "internal/pack/parser.go" \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: internal/pack/output_check.go
Category: invariants
Horizon: NearTerm
ROI: 3.1 (impact=Medium, confidence=Medium, effort=Medium)

Question:
Where does validation happen (pack parsing, output section validation, tool detection)? Identify the validation boundaries and the most likely gaps that could cause inconsistent state or false positives.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces correctness risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=3.0 impact=Medium confidence=Medium effort=Low horizon=NearTerm category=caching/state reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/05-caching-state-layers.md" \
  -f "internal/state/types.go" \
  -f "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: internal/state/types.go
Category: caching/state
Horizon: NearTerm
ROI: 3.0 (impact=Medium, confidence=Medium, effort=Low)

Question:
What stateful components exist (run state, report artifacts, in-memory tracking within a run)? For each, describe lifecycle, persistence, and where it is implemented.

Rationale (one sentence):
State and caching are common sources of subtle bugs and performance issues.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=2.6 impact=Medium confidence=Low effort=Medium horizon=NearTerm category=caching/state reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/06-caching-state-consistency.md" \
  -f "internal/app/run.go" \
  -f "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: internal/app/run.go
Category: caching/state
Horizon: NearTerm
ROI: 2.6 (impact=Medium, confidence=Low, effort=Medium)

Question:
Identify the top consistency risks between in-memory run state and persisted state/report files (e.g., partial writes, skipped updates, resume behavior). Where are the knobs/configs that influence state persistence?

Rationale (one sentence):
Consistency failure modes often surface as "random bugs" and are expensive to debug.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=2.4 impact=Medium confidence=Low effort=Medium horizon=NearTerm category=background jobs reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/07-background-jobs-discovery.md" \
  -f "internal/app/run.go" \
  -f "internal/tui/tui.go" \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: internal/app/run.go
Category: background jobs
Horizon: NearTerm
ROI: 2.4 (impact=Medium, confidence=Low, effort=Medium)

Question:
What background jobs/workers/scheduled tasks exist (if any)? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined.

Rationale (one sentence):
Background work affects reliability, cost, and operational complexity.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=2.2 impact=Medium confidence=Low effort=High horizon=NearTerm category=background jobs reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/08-background-jobs-reliability.md" \
  -f "internal/exec/runner.go" \
  -f "internal/app/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: internal/exec/runner.go
Category: background jobs
Horizon: NearTerm
ROI: 2.2 (impact=Medium, confidence=Low, effort=High)

Question:
Where are the main reliability controls for any background or long-running work (timeouts, retries, cancellation, concurrency limits), and what is missing or inconsistent?

Rationale (one sentence):
Reliability controls prevent incident loops and data corruption.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=3.3 impact=High confidence=Medium effort=Low horizon=Immediate category=observability reference=internal/report/generate.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/09-observability-signals.md" \
  -f "internal/report/generate.go" \
  -f "internal/report/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: internal/report/generate.go
Category: observability
Horizon: Immediate
ROI: 3.3 (impact=High, confidence=Medium, effort=Low)

Question:
What observability signals exist (run reports, warnings, stdout/stderr logs), and what are the primary identifiers for correlating a step/run across components? Point to the code/config that defines them.

Rationale (one sentence):
You cannot operate or improve what you cannot measure or debug quickly.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=3.0 impact=High confidence=Medium effort=Medium horizon=Immediate category=observability reference=internal/report/io.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/10-observability-gaps.md" \
  -f "internal/report/io.go" \
  -f "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: internal/report/io.go
Category: observability
Horizon: Immediate
ROI: 3.0 (impact=High, confidence=Medium, effort=Medium)

Question:
Where are the biggest observability gaps (missing structured logs, missing metrics, missing trace spans)? Recommend the smallest additions that would most improve debugging.

Rationale (one sentence):
Targeted observability improvements compound across all future changes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=3.2 impact=High confidence=Medium effort=Low horizon=Immediate category=permissions reference=internal/cli/root.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/11-permissions-model.md" \
  -f "internal/cli/root.go" \
  -f "internal/app/app.go" \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: internal/cli/root.go
Category: permissions
Horizon: Immediate
ROI: 3.2 (impact=High, confidence=Medium, effort=Low)

Question:
What is the permission model (roles/scopes/claims/ACLs), and where is it defined? If none, explain what access assumptions the CLI makes (filesystem, shell, external tools).

Rationale (one sentence):
Permission rules are a high-risk area with security and product impact.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=2.8 impact=High confidence=Low effort=Medium horizon=Immediate category=permissions reference=internal/cli/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/12-permissions-enforcement.md" \
  -f "internal/cli/run.go" \
  -f "internal/app/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: internal/cli/run.go
Category: permissions
Horizon: Immediate
ROI: 2.8 (impact=High, confidence=Low, effort=Medium)

Question:
Where are permissions or access checks enforced (if at all)? Identify enforcement chokepoints and any bypass risks for filesystem or shell execution.

Rationale (one sentence):
Enforcement consistency prevents privilege escalation and policy drift.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=2.5 impact=Medium confidence=Low effort=Medium horizon=NearTerm category=migrations reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/13-migrations-schema.md" \
  -f "internal/state/types.go" \
  -f "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: internal/state/types.go
Category: migrations
Horizon: NearTerm
ROI: 2.5 (impact=Medium, confidence=Low, effort=Medium)

Question:
How are schema/config migrations handled (state file schema versions, report schema changes, pack format evolution)? Identify tooling, version fields, and how migrations are applied during upgrades.

Rationale (one sentence):
Migration mechanics are critical for safe releases and rollbacks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=2.2 impact=Medium confidence=Low effort=High horizon=NearTerm category=migrations reference=internal/state/persist.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/14-migrations-compat.md" \
  -f "internal/state/persist.go" \
  -f "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: internal/state/persist.go
Category: migrations
Horizon: NearTerm
ROI: 2.2 (impact=Medium, confidence=Low, effort=High)

Question:
What are the backward/forward compatibility expectations for state and report files (e.g., loading older state, changing schema_version)? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during upgrades.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=3.1 impact=High confidence=Medium effort=Medium horizon=NearTerm category=UX flows reference=internal/tui/tui.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/15-ux-flows-primary.md" \
  -f "internal/tui/tui.go" \
  -f "internal/cli/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: internal/tui/tui.go
Category: UX flows
Horizon: NearTerm
ROI: 3.1 (impact=High, confidence=Medium, effort=Medium)

Question:
What are the primary user/operator flows (TUI navigation, step selection, overrides, run/resume)? Map each to the main components/modules involved and note key state transitions.

Rationale (one sentence):
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=2.7 impact=Medium confidence=Medium effort=Medium horizon=NearTerm category=UX flows reference=internal/tui/overrides_flow.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/16-ux-flows-edgecases.md" \
  -f "internal/tui/overrides_flow.go" \
  -f "internal/tui/tui.go" \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: internal/tui/overrides_flow.go
Category: UX flows
Horizon: NearTerm
ROI: 2.7 (impact=Medium, confidence=Medium, effort=Medium)

Question:
For the primary flows, what are the top edge cases and "gotchas" (validation failures, canceled overrides, partial runs, retries)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=3.4 impact=High confidence=Medium effort=Low horizon=Immediate category=failure modes reference=internal/errors/errors.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/17-failure-modes-taxonomy.md" \
  -f "internal/errors/errors.go" \
  -f "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: internal/errors/errors.go
Category: failure modes
Horizon: Immediate
ROI: 3.4 (impact=High, confidence=Medium, effort=Low)

Question:
What is the failure-mode taxonomy of this system (invalid pack, execution failure, config errors, subprocess failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=3.0 impact=High confidence=Medium effort=Medium horizon=Immediate category=failure modes reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/18-failure-modes-resilience.md" \
  -f "internal/app/run.go" \
  -f "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: internal/app/run.go
Category: failure modes
Horizon: Immediate
ROI: 3.0 (impact=High, confidence=Medium, effort=Medium)

Question:
What resilience mechanisms exist (retry loops, output verification retries, stop-on-fail, resume), and which critical paths lack them? Where are these configured?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=2.6 impact=Medium confidence=Low effort=Medium horizon=NearTerm category=feature flags reference=internal/overrides/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/19-feature-flags-inventory.md" \
  -f "internal/overrides/types.go" \
  -f "internal/cli/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: internal/overrides/types.go
Category: feature flags
Horizon: NearTerm
ROI: 2.6 (impact=Medium, confidence=Low, effort=Medium)

Question:
What feature-flag or override system exists (flag injection, step targeting, ROI filtering)? Inventory the flags/overrides and identify where they are defined, evaluated, and documented.

Rationale (one sentence):
Flags enable safe rollout and experimentation and reduce release risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=2.4 impact=Medium confidence=Low effort=High horizon=NearTerm category=feature flags reference=internal/overrides/merge.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-09/20-feature-flags-rollout.md" \
  -f "internal/overrides/merge.go" \
  -f "internal/tui/overrides_flow.go" \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: internal/overrides/merge.go
Category: feature flags
Horizon: NearTerm
ROI: 2.4 (impact=Medium, confidence=Low, effort=High)

Question:
Describe the flag/override lifecycle (how overrides are created, validated, applied to steps, and retired). Identify minimum guardrails needed to prevent override drift.

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check
--------------

Mark each as OK or Missing(<which step ids>):

*   contracts/interfaces: OK

*   invariants: OK

*   caching/state: OK

*   background jobs: OK

*   observability: OK

*   permissions: OK

*   migrations: OK

*   UX flows: OK

*   failure modes: OK

*   feature flags: OK
