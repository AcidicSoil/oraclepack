# Oracle Pack — Unknown (Tickets Stage 1)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-08
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files:
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths:
- ticket_bundle_path: docs/oracle-questions-2026-01-08/_tickets_bundle.md
- mode: tickets

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash


# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/01-contracts-interfaces-ticket-surface.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven)

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the ticket bundle as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts). For each implied change:
- describe the new/changed interface shape
- identify the most likely code areas involved
- call out any backwards-compatibility constraints that must be preserved.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/02-contracts-interfaces-integration-points.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven)

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the ticket bundle as the primary context, identify any external integrations implied by the tickets (e.g., calling new agents/tools, new CLIs, new services). For each:
- what contract/config must be added or changed?
- what failure/timeout behavior should be defined?
- what should be the minimal “compat-safe” rollout approach?

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.2 impact=7 confidence=0.75 effort=2 horizon=NearTerm category=invariants reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/03-invariants-correctness-guardrails.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven)

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: 5.2 (impact=7, confidence=0.75, effort=2)

Question:
From the tickets, derive the key correctness invariants that must hold while implementing them (e.g., “runner-ingestible pack constraints”, “no schema drift”, “no unsafe paths”). For each invariant:
- define it precisely
- state how to enforce it (validation, linting, runtime checks)
- identify where it should live in the codebase (by file/path patterns if evidence is missing).

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/04-invariants-validation-boundaries.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven)

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the tickets, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation). Where could invalid inputs slip through (missing tickets, malformed headers, extra fences)? Propose a minimal validation plan that preserves backward compatibility.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/05-caching-state-ticket-artifacts.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven)

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Based on the tickets, what state/artifacts must be produced and preserved (ticket bundle, generated pack, validator outputs, runner outputs)? Identify any schema/format expectations that must remain backward compatible and how to keep them stable.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.1 impact=5 confidence=0.80 effort=2 horizon=NearTerm category=caching/state reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/06-caching-state-determinism-consistency.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #06  (ticket-driven)

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: 4.1 (impact=5, confidence=0.80, effort=2)

Question:
Using the tickets, identify determinism risks (non-deterministic ticket selection, unstable ordering, environment-dependent paths). Propose a minimal deterministic selection and bundling approach and how to test it.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=3.6 impact=5 confidence=0.70 effort=3 horizon=NearTerm category=background jobs reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/07-background-jobs-ticket-implications.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #07  (ticket-driven)

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: 3.6 (impact=5, confidence=0.70, effort=3)

Question:
Do any tickets imply background processing, worker modes, scheduled validation, or CI pipelines (e.g., generating packs from tickets in CI)? If yes, define:
- trigger mechanism
- inputs/outputs
- retries/idempotency constraints
If no, explicitly confirm based on the ticket bundle.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=3.9 impact=5 confidence=0.72 effort=3 horizon=NearTerm category=background jobs reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/08-background-jobs-reliability-controls.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #08  (ticket-driven)

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: 3.9 (impact=5, confidence=0.72, effort=3)

Question:
If tickets imply background/CI execution, what reliability controls are required (concurrency limits, backoff, reprocessing, artifact retention)? Tie each control to a specific ticket requirement and suggest minimal implementation points.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=4.7 impact=6 confidence=0.82 effort=2 horizon=Immediate category=observability reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/09-observability-required-signals.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #09  (ticket-driven)

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: 4.7 (impact=6, confidence=0.82, effort=2)

Question:
From the tickets, define the minimum observability required for implementing and operating ticketed changes (logs, warnings, structured outputs, correlation/run IDs). What signals must be emitted on:
- missing tickets
- validation failures
- unsafe path detection
- runner ingestion failures?

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=4.5 impact=6 confidence=0.80 effort=2 horizon=Immediate category=observability reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/10-observability-gaps-and-metrics.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #10  (ticket-driven)

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: 4.5 (impact=6, confidence=0.80, effort=2)

Question:
Using the ticket bundle, identify observability gaps that would block shipping the ticketed work safely (missing structured errors, missing per-step diagnostics, missing coverage/mismatch reporting). Recommend the smallest additions with high debugging value.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=5.0 impact=7 confidence=0.76 effort=2 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/11-permissions-security-constraints.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #11  (ticket-driven)

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 5.0 (impact=7, confidence=0.76, effort=2)

Question:
From the tickets, determine what security/permissions constraints must exist (e.g., exec gating, tool invocation restrictions, file access restrictions, safe write paths). Define “who can do what” minimally and where enforcement should live.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=4.8 impact=7 confidence=0.74 effort=2 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/12-permissions-enforcement-chokepoints.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #12  (ticket-driven)

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 4.8 (impact=7, confidence=0.74, effort=2)

Question:
Using the tickets, identify where permissions must be enforced (CLI command gating, TUI actions, runner execution, filesystem writes). Call out bypass risks and propose the minimal enforcement chokepoints and tests.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=4.2 impact=6 confidence=0.72 effort=3 horizon=NearTerm category=migrations reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/13-migrations-schema-changes.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #13  (ticket-driven)

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.72, effort=3)

Question:
Do tickets imply schema/version changes (pack schema, state/report schema, actions artifacts)? Identify what can change vs must remain backward-compatible, and propose a minimal migration strategy (if any).

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=4.0 impact=6 confidence=0.70 effort=3 horizon=NearTerm category=migrations reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/14-migrations-compat-guardrails.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #14  (ticket-driven)

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: 4.0 (impact=6, confidence=0.70, effort=3)

Question:
Using the ticket bundle, define the compatibility expectations (backward/forward, rolling upgrades, mixed versions). Where are the risky edges, and what guardrails/tests should be required before shipping ticketed changes?

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=4.6 impact=6 confidence=0.80 effort=2 horizon=NearTerm category=UX flows reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/15-ux-flows-ticketed-user-journeys.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #15  (ticket-driven)

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: 4.6 (impact=6, confidence=0.80, effort=2)

Question:
From the ticket bundle, identify which user/operator flows are affected (TUI flows, CLI flows, non-interactive mode). For each flow:
- outline steps and state transitions
- identify key UX requirements implied by tickets
- call out compatibility constraints with existing workflows.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=4.2 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=UX flows reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/16-ux-flows-edge-cases-and-gotchas.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #16  (ticket-driven)

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: 4.2 (impact=6, confidence=0.78, effort=2)

Question:
Using the ticket bundle, enumerate top UX edge cases (“gotchas”) that must be handled (missing tickets, partial bundles, validation failures, ambiguous outputs, cancel/back navigation). Identify where handling should be implemented and what tests are required.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=5.4 impact=7 confidence=0.76 effort=2 horizon=Immediate category=failure modes reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/17-failure-modes-taxonomy-from-tickets.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #17  (ticket-driven)

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: 5.4 (impact=7, confidence=0.76, effort=2)

Question:
Derive a failure-mode taxonomy implied by the tickets (ticket discovery failures, bundle generation failures, schema violations, runner ingestion errors, tool execution failures). For each failure:
- expected user-visible behavior
- diagnostics to emit
- where to classify/handle it.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=5.1 impact=7 confidence=0.74 effort=2 horizon=Immediate category=failure modes reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/18-failure-modes-resilience-and-recovery.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #18  (ticket-driven)

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: 5.1 (impact=7, confidence=0.74, effort=2)

Question:
Using the tickets, propose resilience mechanisms appropriate for the ticketed changes (sanitize unsafe output paths, fail-fast vs partial completion, error wrapping, “stop-on-fail” behavior). Identify critical paths lacking safeguards and the smallest fixes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=4.7 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=feature flags reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/19-feature-flags-rollout-controls.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #19  (ticket-driven)

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.78, effort=2)

Question:
Do tickets imply the need for feature-flag-like controls (rollout gating, experimental flags, safety toggles, per-step targeting)? Inventory what controls should exist and where they should be defined/evaluated.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=4.3 impact=6 confidence=0.76 effort=2 horizon=NearTerm category=feature flags reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-08/20-feature-flags-lifecycle-and-flag-debt.md" \
  -f "docs/oracle-questions-2026-01-08/_tickets_bundle.md" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
   \
  -p "$(cat <<'PROMPT'
Strategist question #20  (ticket-driven)

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: 4.3 (impact=6, confidence=0.76, effort=2)

Question:
Define the rollout lifecycle required to ship the ticketed changes safely (create flag, test, ramp, monitor, retire). Identify minimum guardrails needed to prevent flag debt and to keep backward compatibility.

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

## Coverage check

Mark each as `OK` or `Missing(<step ids>)`:

- contracts/interfaces: OK
- invariants: OK
- caching/state: OK
- background jobs: OK
- observability: OK
- permissions: OK
- migrations: OK
- UX flows: OK
- failure modes: OK
- feature flags: OK
