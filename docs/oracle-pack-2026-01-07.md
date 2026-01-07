# Oracle Pack — Unknown (Gold Stage 1)

## Parsed args
- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown
- out_dir: docs/oracle-questions-2026-01-07
- oracle_cmd: oracle
- oracle_flags: --files-report
- engine: None
- model: None
- extra_files: None

Notes:
- Template is the contract. Keep the pack runner-ingestible.
- Exactly one fenced bash block in this whole document.
- Exactly 20 steps, numbered 01..20.
- Each step includes: ROI= impact= confidence= effort= horizon= category= reference=
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=8.2 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/cli/root.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md" \
  --file "internal/cli/root.go" \
  --file "internal/cli/cmds.go" \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: internal/cli/root.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 8.2 (impact=9, confidence=0.9, effort=1)

Question:
Identify the primary public interface(s) of this system (CLI commands and the oracle-pack Markdown contract). For each, list the key inputs/outputs and where they are defined in code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.5 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md" \
  --file "internal/exec/runner.go" \
  --file "internal/exec/oracle_validate.go" \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: internal/exec/runner.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 7.5 (impact=8, confidence=0.9, effort=1)

Question:
What are the top integration points with external systems (oracle CLI invocation, filesystem outputs, environment variables), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

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

# 03) ROI=6.8 impact=8 confidence=0.85 effort=1 horizon=NearTerm category=invariants reference=internal/pack/parser.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/03-invariants-domain.md" \
  --file "internal/pack/parser.go" \
  --file "internal/pack/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: internal/pack/parser.go
Category: invariants
Horizon: NearTerm
ROI: 6.8 (impact=8, confidence=0.85, effort=1)

Question:
List the system’s most important invariants about pack structure and step ordering. For each, show where it is enforced (or where it should be enforced but currently is not).

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

# 04) ROI=6.2 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=invariants reference=internal/app/app.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/04-invariants-validation.md" \
  --file "internal/app/app.go" \
  --file "internal/pack/parser.go" \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: internal/app/app.go
Category: invariants
Horizon: NearTerm
ROI: 6.2 (impact=7, confidence=0.8, effort=1)

Question:
Where does validation happen (pack parse/validate, config resolution, runtime checks)? Identify validation boundaries and the most likely gaps that could cause inconsistent state.

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

# 05) ROI=6.0 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=caching/state reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/05-caching-state-layers.md" \
  --file "internal/state/types.go" \
  --file "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: internal/state/types.go
Category: caching/state
Horizon: NearTerm
ROI: 6.0 (impact=7, confidence=0.8, effort=1)

Question:
What stateful components exist (run state, status tracking, persisted state files)? For each, describe lifecycle, read/write points, and where it is implemented.

Rationale (one sentence):
State handling is a common source of subtle bugs and operational issues.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=5.8 impact=7 confidence=0.75 effort=1 horizon=NearTerm category=caching/state reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/06-caching-state-consistency.md" \
  --file "internal/app/run.go" \
  --file "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: internal/app/run.go
Category: caching/state
Horizon: NearTerm
ROI: 5.8 (impact=7, confidence=0.75, effort=1)

Question:
Identify the top consistency risks between in-memory run state and persisted state/report outputs (stale writes, partial updates, missing flushes). Where are the knobs/configs for state behavior?

Rationale (one sentence):
Consistency failure modes often surface as “random bugs” and are expensive to debug.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=5.0 impact=6 confidence=0.7 effort=1 horizon=NearTerm category=background jobs reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/07-background-jobs-discovery.md" \
  --file "internal/app/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: internal/app/run.go
Category: background jobs
Horizon: NearTerm
ROI: 5.0 (impact=6, confidence=0.7, effort=1)

Question:
What background jobs/workers/scheduled tasks exist, if any? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined. If none, confirm based on evidence.

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

# 08) ROI=4.2 impact=5 confidence=0.7 effort=1 horizon=NearTerm category=background jobs reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/08-background-jobs-reliability.md" \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=5, confidence=0.7, effort=1)

Question:
Where are the main reliability controls for background work (dead-lettering, backoff, concurrency limits, reprocessing), and what is missing or inconsistent? If there are no background jobs, explicitly state that and point to confirming evidence.

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

# 09) ROI=6.4 impact=8 confidence=0.8 effort=1 horizon=Immediate category=observability reference=internal/report/generate.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/09-observability-signals.md" \
  --file "internal/report/generate.go" \
  --file "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: internal/report/generate.go
Category: observability
Horizon: Immediate
ROI: 6.4 (impact=8, confidence=0.8, effort=1)

Question:
What observability signals exist (reports, warnings, log lines), and what are the primary identifiers for correlating a run/step across components? Point to the code/config that defines them.

Rationale (one sentence):
You can’t operate or improve what you can’t measure or debug quickly.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=5.6 impact=7 confidence=0.8 effort=1 horizon=Immediate category=observability reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/10-observability-gaps.md" \
  --file "internal/exec/runner.go" \
  --file "internal/errors/errors.go" \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: internal/exec/runner.go
Category: observability
Horizon: Immediate
ROI: 5.6 (impact=7, confidence=0.8, effort=1)

Question:
Where are the biggest observability gaps (missing structured logs around key decisions, missing metrics for run outcomes, missing correlation IDs)? Recommend the smallest additions that would most improve debugging.

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

# 11) ROI=5.2 impact=7 confidence=0.75 effort=1 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/11-permissions-model.md" \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 5.2 (impact=7, confidence=0.75, effort=1)

Question:
What is the permission model (roles/scopes/ACLs), and where is it defined? Provide the minimal set of files that encode “who can do what.” If no permission model exists, confirm that with evidence.

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

# 12) ROI=4.8 impact=6 confidence=0.8 effort=1 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/12-permissions-enforcement.md" \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.8, effort=1)

Question:
Where are permissions enforced (middleware/guards/policies/service-layer checks), and where are likely bypass risks? Identify enforcement chokepoints and any inconsistent patterns.

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

# 13) ROI=4.5 impact=6 confidence=0.75 effort=1 horizon=NearTerm category=migrations reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/13-migrations-schema.md" \
  --file "internal/state/types.go" \
  --file "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: internal/state/types.go
Category: migrations
Horizon: NearTerm
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
How are schema/config migrations handled (state schema versioning, report format evolution)? Identify the tooling, files, and how migrations are applied or would be applied.

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

# 14) ROI=4.1 impact=5 confidence=0.8 effort=1 horizon=NearTerm category=migrations reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/14-migrations-compat.md" \
  --file "internal/state/types.go" \
  --file "internal/report/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: internal/state/types.go
Category: migrations
Horizon: NearTerm
ROI: 4.1 (impact=5, confidence=0.8, effort=1)

Question:
What are the backward/forward compatibility expectations for state/report schemas? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during upgrades and rollbacks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=5.7 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=UX flows reference=internal/tui/tui.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/15-ux-flows-primary.md" \
  --file "internal/tui/tui.go" \
  --file "internal/cli/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: internal/tui/tui.go
Category: UX flows
Horizon: NearTerm
ROI: 5.7 (impact=7, confidence=0.8, effort=1)

Question:
What are the primary user/operator flows in the TUI (run steps, run all, overrides wizard, URL selection)? Map each to main components/modules and key state transitions.

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

# 16) ROI=5.1 impact=6 confidence=0.85 effort=1 horizon=NearTerm category=UX flows reference=internal/tui/overrides_flow.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/16-ux-flows-edgecases.md" \
  --file "internal/tui/overrides_flow.go" \
  --file "internal/tui/overrides_flags.go" \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: internal/tui/overrides_flow.go
Category: UX flows
Horizon: NearTerm
ROI: 5.1 (impact=6, confidence=0.85, effort=1)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, cancel/back navigation, partial completion, timeouts)? Identify where these cases are handled and where they are missing.

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

# 17) ROI=6.1 impact=8 confidence=0.76 effort=1 horizon=Immediate category=failure modes reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/17-failure-modes-taxonomy.md" \
  --file "internal/app/run.go" \
  --file "internal/errors/errors.go" \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: internal/app/run.go
Category: failure modes
Horizon: Immediate
ROI: 6.1 (impact=8, confidence=0.76, effort=1)

Question:
What is the failure-mode taxonomy of this system (pack validation errors, execution failures, state/report write failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

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

# 18) ROI=5.4 impact=7 confidence=0.77 effort=1 horizon=Immediate category=failure modes reference=internal/exec/sanitize.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/18-failure-modes-resilience.md" \
  --file "internal/exec/sanitize.go" \
  --file "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: internal/exec/sanitize.go
Category: failure modes
Horizon: Immediate
ROI: 5.4 (impact=7, confidence=0.77, effort=1)

Question:
What resilience mechanisms exist (sanitization, error wrapping, stop-on-fail behavior), and which critical paths lack them? Where are they configured?

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

# 19) ROI=4.7 impact=6 confidence=0.78 effort=1 horizon=NearTerm category=feature flags reference=internal/overrides/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/19-feature-flags-inventory.md" \
  --file "internal/overrides/types.go" \
  --file "internal/tui/overrides_flags.go" \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: internal/overrides/types.go
Category: feature flags
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.78, effort=1)

Question:
What feature-flag-like controls exist (oracle flag overrides, per-step targeting, ROI filtering)? Inventory the flags/controls and identify where they are defined, evaluated, and documented.

Rationale (one sentence):
Flags and runtime controls enable safe rollout and experimentation and reduce release risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=4.3 impact=5 confidence=0.86 effort=1 horizon=NearTerm category=feature flags reference=internal/overrides/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/20-feature-flags-rollout.md" \
  --file "internal/overrides/types.go" \
  --file "internal/overrides/merge.go" \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: internal/overrides/types.go
Category: feature flags
Horizon: NearTerm
ROI: 4.3 (impact=5, confidence=0.86, effort=1)

Question:
Describe the flag/rollout lifecycle for runtime overrides (how flags are selected, applied per-step, validated, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

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

* contracts/interfaces: OK

* invariants: OK

* caching/state: OK

* background jobs: OK

* observability: OK

* permissions: OK

* migrations: OK

* UX flows: OK

* failure modes: OK

* feature flags: OK
