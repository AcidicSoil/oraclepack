<filetree>
Project Structure:
└── docs
    ├── oracle-questions-2026-01-09-030000
    │   ├── actions
    │   │   ├── 01-contracts-interfaces-ticket-surface-direct-answer.md
    │   │   ├── 01-contracts-interfaces-ticket-surface-missing-evidence.md
    │   │   ├── 01-contracts-interfaces-ticket-surface-next-experiment.md
    │   │   └── 01-contracts-interfaces-ticket-surface-risks-unknowns.md
    │   ├── packs
    │   │   ├── actions.md
    │   │   ├── mcp.md
    │   │   ├── misc.md
    │   │   ├── other.md
    │   │   ├── prd-tui.md
    │   │   └── raw-exports.md
    │   ├── _groups.json
    │   └── manifest.json
    └── oracle-pack-2026-01-09.md

</filetree>

<source_code>
docs/oracle-pack-2026-01-09.md
```
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
```

docs/oracle-questions-2026-01-09-030000/_groups.json
```
{
  "PRD-TUI": [
    ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
    ".tickets/PRD-TUI/PRD-generator URL routing.md"
  ],
  "actions": [
    ".tickets/actions/Enable Action Packs Dispatch.md",
    ".tickets/actions/Improving Oraclepack Workflow.md",
    ".tickets/actions/Oraclepack Action Pack Integration.md",
    ".tickets/actions/Oraclepack Action Pack Issue.md",
    ".tickets/actions/Oraclepack Action Packs.md",
    ".tickets/actions/Oraclepack Compatibility Issues.md"
  ],
  "mcp": [
    ".tickets/mcp/Expose Oraclepack as MCP.md",
    ".tickets/mcp/MCP Server for Oraclepack.md",
    ".tickets/mcp/gaps-still-not-covered.md",
    ".tickets/mcp/gaps_part2-mcp-builder.md",
    ".tickets/mcp/oraclepack-MCP.md",
    ".tickets/mcp/oraclepack_mcp_server.md"
  ],
  "misc": [
    ".tickets/Formalize LLM Decision Points.md",
    ".tickets/Oraclepack CLI MCP Parity.md",
    ".tickets/Oraclepack File Storage.md",
    ".tickets/Oraclepack Parity Automation.md",
    ".tickets/Oraclepack Schema Approach.md",
    ".tickets/Oraclepack bash fix.md",
    ".tickets/Oraclepack output verification issues.md",
    ".tickets/Oraclepack-CLI-agents.md",
    ".tickets/Publish OraclePack MCP.md"
  ],
  "other": [
    ".tickets/other/Oraclepack Pipeline Improvements.md",
    ".tickets/other/Oraclepack Prompt Generator.md",
    ".tickets/other/Oraclepack Workflow Enhancement.md",
    ".tickets/other/Verbose Payload Rendering TUI.md"
  ],
  "raw-exports": [
    ".tickets/raw-exports/Output verification failure.md"
  ]
}
```

docs/oracle-questions-2026-01-09-030000/manifest.json
```
{
  "groups": [
    {
      "attached_paths": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "group": "PRD-TUI",
      "original_tickets": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/prd-tui",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/prd-tui.md",
      "part": 1,
      "slug": "prd-tui"
    },
    {
      "attached_paths": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "group": "actions",
      "original_tickets": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/actions",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/actions.md",
      "part": 1,
      "slug": "actions"
    },
    {
      "attached_paths": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "group": "mcp",
      "original_tickets": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/mcp",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/mcp.md",
      "part": 1,
      "slug": "mcp"
    },
    {
      "attached_paths": [
        ".tickets/Formalize LLM Decision Points.md",
        ".tickets/Oraclepack CLI MCP Parity.md",
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Parity Automation.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Oraclepack output verification issues.md",
        ".tickets/Oraclepack-CLI-agents.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "group": "misc",
      "original_tickets": [
        ".tickets/Formalize LLM Decision Points.md",
        ".tickets/Oraclepack CLI MCP Parity.md",
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Parity Automation.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Oraclepack output verification issues.md",
        ".tickets/Oraclepack-CLI-agents.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/misc",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/misc.md",
      "part": 1,
      "slug": "misc"
    },
    {
      "attached_paths": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "group": "other",
      "original_tickets": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/other",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/other.md",
      "part": 1,
      "slug": "other"
    },
    {
      "attached_paths": [
        ".tickets/raw-exports/Output verification failure.md"
      ],
      "group": "raw-exports",
      "original_tickets": [
        ".tickets/raw-exports/Output verification failure.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-09-030000/raw-exports",
      "pack_path": "docs/oracle-questions-2026-01-09-030000/packs/raw-exports.md",
      "part": 1,
      "slug": "raw-exports"
    }
  ]
}
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-direct-answer.md
```
Direct answer

* Clarify/document Action Pack execution semantics (CLI/TUI docs surface): steps run as `bash -lc ...` in the project root, and oraclepack’s “special handling” (flag injection/override validation) applies only to commands beginning with `oracle`; non-`oracle` tools (`tm`/`task-master`, `codex`, `gemini`) run directly as shell commands.  
* Dispatcher/command-detection contract expansion: broaden detection beyond the oracle-anchored regex (`^(\\s*)(oracle)\\b`) so steps containing `tm`/`task-master`, `codex`, and `gemini` can be treated as first-class command targets for dispatch/override/validation inclusion. Back-compat: preserve existing behavior for `oracle ...` commands. 
* Override-validation behavior change (TUI/validate surface): today validation targets only oracle invocations (runs `oracle --dry-run summary`) and skips steps without oracle lines; tickets imply extending/restructuring validation so steps containing `tm`/`task-master`, `codex`, `gemini` are not silently excluded purely due to prefix mismatch. Back-compat: do not regress the current oracle-only validation flow.  
* `ticket-action-pack.md` content contract change: replace placeholder steps (explicitly 09–13 and 16) with headless `gemini` + non-interactive `codex exec` automation. Back-compat: keep the pack “oraclepack-ingestible” (single bash fence, `# NN)` steps) and keep Steps 01–07 semantics unchanged.  
* New standardized output artifact interface under `.oraclepack/ticketify/` for the agent loop: Step 09 `next.json`, Step 10 `codex-implement.md`, Step 11 `codex-verify.md` and/or `gemini-review.json`, Step 16 `PR.md`. Back-compat: paths are treated as required/stable outputs (don’t rename/move without versioning).  
* Tool-availability/skip semantics become part of the operational contract for these steps: add `command -v ...` guards and documented “skip” behavior to avoid hard failures when `codex`/`gemini` are missing or interactive (blocking risk). Back-compat: default runs should not newly fail just because optional tools aren’t installed.  
* Taskify Action Pack generator surface change: add an opt-in “agent-mode” (suggested `mode=codex` / `mode=gemini`) that swaps the existing autopilot entrypoint step slot after Task Master expansion, while keeping the 20-step contract intact. Back-compat: default mode remains current behavior; agent-mode must not add steps. 
* Output/reporting contract stability: tickets explicitly name run outputs/artifacts (e.g., `.oraclepack/ticketify/_tickets_index.json`, `_actions.json`, `.taskmaster/docs/tickets_prd.md`, `.oraclepack/ticketify/tm-complexity.json`, plus `ticket-action-pack.state.json` / `ticket-action-pack.report.json`) and restate project-root execution (no chdir to `out_dir`). Back-compat: keep these names/locations stable to avoid breaking consumers and user expectations. 
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-missing-evidence.md
```
Missing evidence

* `.tickets/actions/*.md` (full raw tickets, not summaries)
* `ticket-action-pack.md` (and its generator/template source): `**/*ticket-action-pack*.md`, `**/*ticketify*pack*.md`, `**/templates/**`, `**/assets/**`
* “taskify” Action Pack generator implementation (for proposed `mode=codex|gemini`): `**/*taskify*.*`, `**/taskify/**`, `**/oraclepack-taskify/**`
* Oraclepack CLI/TUI public surface (commands/flags/help) and wiring: `**/cmd/**`, `**/internal/cli/**`, `**/*root*.go`, `**/*commands*.go`
* Oraclepack dispatcher / command-detection / override-injection / validation pipeline (regex currently `^(\\s*)(oracle)\\b` per tickets): `**/*dispatch*.*`, `**/*detect*.*`, `**/*override*.*`, `**/*validate*.*`, `**/internal/**`
* Oraclepack TUI override-validation flow (to assess user-visible behavior/back-compat): `**/*tui*.*`, `**/internal/**tui**/**`
* Docs explicitly referenced by the tickets: `**/oraclepack-tui.md`, `**/oracle_pack_and_taskify-skills.md`
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-next-experiment.md
```
Next smallest concrete experiment

* Run `rg -n --hidden --no-ignore-vcs 'ticket-action-pack\.md|Action Pack|override validation|dry-run summary|\^\(\\\\s\*\)\(oracle\)\\\\b|\btm\b|\btask-master\b|\bcodex\b|\bgemini\b' .`
```

docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-risks-unknowns.md
```
Risks/unknowns

* Override injection/validation expansion beyond `oracle` is underspecified: which overrides/flags apply to `tm`/`task-master`, `codex`, and `gemini`, and what “validation” means for each tool are not defined; this can easily produce partial/incorrect behavior.
* Dispatcher intent is ambiguous: tickets do not specify whether the dispatcher should “interpret actions” (semantic dispatch) vs only broaden prefix-based detection to include non-`oracle` commands.
* User-facing docs/TUI discoverability is unclear: where the “Action Pack execution semantics + failure modes” documentation should live (README vs `oraclepack-tui.md` vs TUI help text) is not specified, risking continued user confusion about what is/ isn’t routed/validated.
* Action Pack “implement” execution path has key unknowns: where `top_n` is defined/how ranking works, and the exact `<out_dir>/_actions.json` location/structure needed for deterministic dispatch are not provided. 
* Taskify “agent-mode” surface is incomplete: the tickets don’t specify how `mode=codex|gemini` is selected (CLI flag vs TUI option vs config) or which exact step slot replaces the autopilot entrypoint while keeping the 20-step contract intact. 
* `ticket-action-pack.md` placeholder replacement is underspecified in places: whether Step 11 defaults to Codex verification, Gemini diff review, or both, and whether Steps 12–13 must change are not defined, increasing the risk of a half-automated pack that still looks “successful” but doesn’t produce expected artifacts. 
* Operational “skip” contract for missing tools is not defined: tickets require `command -v` guards and “skip” behavior when `codex`/`gemini` are absent, but don’t specify exit codes/status/report semantics, creating inconsistency across CLI/TUI reporting and resume behavior.
* Runner/TUI behavior gaps remain undefined for reliability: signal handling (SIGINT propagation, atomic state/log flush), the full resume contract (what reruns, how interrupted prelude is treated), and how to avoid UX implying non-existent guarantees for non-`oracle` override/validation are not specified. 
```

docs/oracle-questions-2026-01-09-030000/packs/actions.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/actions
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md
- ticket_max_files: 6
- group_name: actions
- group_slug: actions
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/actions/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/actions"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/actions/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/mcp.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/mcp
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md
- ticket_max_files: 6
- group_name: mcp
- group_slug: mcp
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/mcp/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/mcp"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/mcp/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/misc.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/misc
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md
- ticket_max_files: 9
- group_name: misc
- group_slug: misc
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/misc/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/misc"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/misc/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Formalize LLM Decision Points.md,.tickets/Oraclepack CLI MCP Parity.md,.tickets/Oraclepack File Storage.md,.tickets/Oraclepack Parity Automation.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Oraclepack output verification issues.md,.tickets/Oraclepack-CLI-agents.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("9")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/other.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/other
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md
- ticket_max_files: 4
- group_name: other
- group_slug: other
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/other/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/other"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/other/04-invariants-validation-boundaries-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/prd-tui.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/prd-tui
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md
- ticket_max_files: 2
- group_name: PRD-TUI
- group_slug: prd-tui
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/prd-tui"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/prd-tui/04-invariants-validation-boundaries-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
[TRUNCATED]
```

docs/oracle-questions-2026-01-09-030000/packs/raw-exports.md
```
# Oracle Pack — Unknown (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: Unknown
- out_dir: docs/oracle-questions-2026-01-09-030000/raw-exports
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/raw-exports/Output verification failure.md
- ticket_max_files: 1
- group_name: raw-exports
- group_slug: raw-exports
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-09-030000/raw-exports"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/01-contracts-interfaces-ticket-surface-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/02-contracts-interfaces-integration-points-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/02-contracts-interfaces-integration-points-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/02-contracts-interfaces-integration-points-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/02-contracts-interfaces-integration-points-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/03-invariants-invariant-map-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
Return only: Direct answer (1–10 bullets, evidence-cited)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/03-invariants-invariant-map-risks-unknowns.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Risks/unknowns
Start with heading: Risks/unknowns
Return only: Risks/unknowns (bullets)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/03-invariants-invariant-map-next-experiment.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Next smallest concrete experiment
Start with heading: Next smallest concrete experiment
Return only: Next smallest concrete experiment (exactly one action)
PROMPT
)"

oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/03-invariants-invariant-map-missing-evidence.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Missing evidence
Start with heading: Missing evidence
Return only: If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"


# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=raw-exports

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path
import fnmatch

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/raw-exports/Output verification failure.md".strip()
MAX = int("1")


root = Path(TICKET_ROOT)

def read_gitignore(start: Path):
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


def gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def is_gitignored(path: Path, git_root: Path, patterns: list[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    name = rel_posix.split("/")[-1]
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        if gitignore_match(rel_posix, name, pat):
            ignored = not neg
    return ignored


git_root, git_patterns = read_gitignore(root)
def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
tickets = [p for p in tickets if not is_gitignored(p, git_root, git_patterns)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'raw-exports'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-09-030000/raw-exports/04-invariants-validation-boundaries-direct-answer.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: raw-exports)

Reference: raw-exports
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

Output section: Direct answer
Start with heading: Direct answer
[TRUNCATED]
```

</source_code>