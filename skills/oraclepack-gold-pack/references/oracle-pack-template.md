# Oracle Pack — {{codebase_name}} (Gold Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- engine: {{engine}}
- model: {{model}}
- extra_files: {{extra_files}}

Notes:
- Template is the **contract**. Keep the pack runner-ingestible.
- Exactly one fenced `bash` block in this whole document.
- Exactly 20 steps, numbered `01..20`.
- Each step includes: `ROI= impact= confidence= effort= horizon= category= reference=`
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the primary public interface(s) of this system (API endpoints, CLI commands, public SDK surface, event contracts). For each, list the key request/response (or input/output) shapes and where they are defined in the code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-integration.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the top integration points with external systems (databases, queues, third-party APIs, SSO, storage), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-domain.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
List the system’s most important invariants (business rules, correctness properties, “must always be true” conditions). For each, show where it is enforced (or where it should be enforced but currently is not).

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-validation.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where does validation happen (input validation, schema validation, domain validation)? Identify the validation boundaries and the most likely gaps that could cause inconsistent state.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces security/correctness risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-layers.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What stateful components exist (in-memory state, caches, sessions, client-side state, persisted state)? For each, describe lifecycle, invalidation/expiry, and where it is implemented.

Rationale (one sentence):
State and caching are common sources of subtle bugs and performance issues.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-consistency.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the top consistency risks between caches/state layers and the source of truth (stale reads, write skew, missing invalidation). Where are the knobs/configs for cache behavior?

Rationale (one sentence):
Consistency failure modes often surface as “random bugs” and are expensive to debug.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-discovery.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What background jobs/workers/scheduled tasks exist? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined.

Rationale (one sentence):
Background work affects reliability, cost, and operational complexity.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-reliability.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the main reliability controls for background work (dead-lettering, backoff, concurrency limits, reprocessing), and what is missing or inconsistent?

Rationale (one sentence):
Reliability controls prevent incident loops and data corruption.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-signals.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What observability signals exist (logs/metrics/traces/events), and what are the primary identifiers for correlating a request/job across components? Point to the code/config that defines them.

Rationale (one sentence):
You can’t operate or improve what you can’t measure or debug quickly.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-gaps.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the biggest observability gaps (missing logs around key decisions, missing metrics for SLOs, missing trace spans)? Recommend the smallest additions that would most improve debugging.

Rationale (one sentence):
Targeted observability improvements compound across all future changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-model.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the permission model (roles/scopes/claims/ACLs), and where is it defined? Provide the minimal set of files that encode “who can do what.”

Rationale (one sentence):
Permission rules are a high-risk area with security and product impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-enforcement.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are permissions enforced (middleware/guards/policies/service-layer checks), and where are likely bypass risks? Identify the enforcement chokepoints and any inconsistent patterns.

Rationale (one sentence):
Enforcement consistency prevents privilege escalation and policy drift.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-schema.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
How are schema/config migrations handled (DB migrations, data backfills, versioned configs)? Identify the tooling, directories, and how migrations are applied in CI/deploy.

Rationale (one sentence):
Migration mechanics are critical for safe releases and rollbacks.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compat.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the backward/forward compatibility expectations during migrations (rolling deploys, dual-read/dual-write, feature-flagged schema use)? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during deployments.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-primary.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the primary user flows (or primary operator workflows) and their steps? Map each to the main components/modules involved, and note the key state transitions.

Rationale (one sentence):
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edgecases.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, partial completion, retries, timeouts)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-taxonomy.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the failure-mode taxonomy of this system (timeouts, retries, partial failures, inconsistent state, dependency failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-resilience.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What resilience mechanisms exist (circuit breakers, bulkheads, retries with jitter, rate limiting, graceful degradation)? Where are they configured, and which critical path lacks them?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-inventory.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What feature-flag system exists (or how are conditional rollouts handled)? Inventory the flags (or equivalents) and identify where flags are defined, evaluated, and documented.

Rationale (one sentence):
Flags enable safe rollout and experimentation and reduce release risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-rollout.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Describe the flag/rollout lifecycle (how flags are created, tested, ramped, monitored, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

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

Mark each as `OK` or `Missing(<which step ids>)`:

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
