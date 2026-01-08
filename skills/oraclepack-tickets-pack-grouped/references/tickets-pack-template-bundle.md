# Oracle Pack — {{codebase_name}} (Tickets Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}
- ticket_root: {{ticket_root}}
- ticket_glob: {{ticket_glob}}
- ticket_paths: {{ticket_paths}}
- ticket_bundle_path: {{ticket_bundle_path}}
- mode: {{mode}}

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "{{out_dir}}/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- `## Coverage check` MUST be outside the bash fence (after the closing ```).

```bash
# Prelude (allowed inside the single bash fence)
# - Creates out_dir deterministically
# - Builds ticket_bundle_path deterministically from ticket_root/ticket_glob OR ticket_paths
# - Uses lexicographic ordering only (no mtime/timestamps)

set -euo pipefail

mkdir -p "{{out_dir}}"

python3 - <<'PY'
from __future__ import annotations

import sys
from pathlib import Path

CODEBASE_NAME = "{{codebase_name}}"
OUT_DIR = Path("{{out_dir}}")
TICKET_ROOT = Path("{{ticket_root}}")
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS_RAW = "{{ticket_paths}}".strip()
BUNDLE_PATH = Path("{{ticket_bundle_path}}")

def _read_text(p: Path) -> str:
    return p.read_text(encoding="utf-8", errors="replace")

def _title_from_md(text: str) -> str:
    for ln in text.splitlines():
        s = ln.strip()
        if s.startswith("# "):
            return s[2:].strip() or "Untitled"
    for ln in text.splitlines():
        s = ln.strip()
        if s:
            return s[:80]
    return "Untitled"

def _select_paths() -> list[Path]:
    if TICKET_PATHS_RAW:
        items = [Path(x.strip()) for x in TICKET_PATHS_RAW.split(",") if x.strip()]
        items = sorted(items, key=lambda p: str(p))
        return items

    if not TICKET_ROOT.exists():
        return []

    items = sorted(TICKET_ROOT.glob(TICKET_GLOB), key=lambda p: str(p))
    return items

paths = _select_paths()

BUNDLE_PATH.parent.mkdir(parents=True, exist_ok=True)

lines: list[str] = []
lines.append(f"# Tickets Bundle — {CODEBASE_NAME if CODEBASE_NAME else 'Unknown'}")
lines.append("")
lines.append("## Selection")
lines.append(f"- ticket_root: {TICKET_ROOT}")
lines.append(f"- ticket_glob: {TICKET_GLOB}")
lines.append(f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}")
lines.append("- ordering: lexicographic by path")
lines.append("")

if not paths:
    warn = (
        "## WARNING: No tickets found\n\n"
        "No ticket files were selected.\n\n"
        "What was attempted:\n"
        f"- ticket_root: {TICKET_ROOT}\n"
        f"- ticket_glob: {TICKET_GLOB}\n"
        f"- ticket_paths: {TICKET_PATHS_RAW if TICKET_PATHS_RAW else '(none)'}\n\n"
        "Next: provide explicit ticket_paths or create tickets under ticket_root.\n"
    )
    lines.append(warn)
    print(f"[WARN] No tickets selected; bundle will contain only WARNING.", file=sys.stderr)
else:
    lines.append("## Tickets")
    lines.append("")
    for p in paths:
        lines.append("---")
        lines.append(f"### {_title_from_md(_read_text(p))}")
        lines.append(f"- path: {p}")
        lines.append("")
        try:
            txt = _read_text(p)
        except Exception as e:
            lines.append(f"[ERROR reading file: {e}]")
            lines.append("")
            continue

        # Simple truncation policy: keep first 4000 chars if large.
        if len(txt) > 4000:
            lines.append(txt[:4000])
            lines.append("\n[... truncated ...]\n")
        else:
            lines.append(txt)

        lines.append("")

BUNDLE_PATH.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
print(f"[OK] Wrote ticket bundle: {BUNDLE_PATH}")
PY

# 01) ROI=8.0 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
Category: contracts/interfaces

Using the attached tickets bundle as the primary evidence, identify the primary public interface(s) implied by the tickets (CLI commands, APIs, file contracts, or user workflows).
For each interface:
- list key inputs/outputs
- list the exact files/modules likely defining it (if unknown, say Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.8 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-dependencies.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02
Category: contracts/interfaces

From the attached tickets bundle, infer which external dependencies/services the system must integrate with (CLIs, APIs, SaaS, databases).
For each dependency:
- what contract is required (auth, endpoints, file formats)
- what configuration surface is implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=7.6 impact=8 confidence=0.85 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-must-always-hold.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03
Category: invariants

Based on the attached tickets bundle, list the invariants that must always hold (data constraints, ordering constraints, security invariants, idempotency).
For each invariant:
- what breaks if violated
- where it should be enforced (layer/module; if unknown, Unknown)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=7.2 impact=8 confidence=0.8 effort=2 horizon=Immediate category=invariants reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-input-validation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04
Category: invariants

Using the attached tickets bundle, identify what inputs must be validated (CLI args, config fields, payloads, file paths).
For each input:
- validation rules implied
- failure message/behavior implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=7.0 impact=7 confidence=0.85 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-state-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05
Category: caching/state

From the attached tickets bundle, infer what state must be persisted or cached (files, DB, in-memory, remote).
For each state item:
- read/write lifecycle
- consistency model implied
- failure recovery requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=caching/state reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-cache-invalidation.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06
Category: caching/state

Using the attached tickets bundle, identify caching risks: staleness, invalidation, keying, or race conditions implied by the tickets.
Propose a minimal caching strategy consistent with the tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=6.9 impact=8 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-what-runs-async.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07
Category: background jobs

From the attached tickets bundle, determine what work should run asynchronously/background (schedulers, queues, cron, long-running tasks).
For each job:
- trigger
- inputs/outputs
- retry/backoff requirements

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=6.6 impact=7 confidence=0.75 effort=3 horizon=Near category=background jobs reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-idempotency.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08
Category: background jobs

Using the attached tickets bundle, list the idempotency and concurrency constraints implied for background jobs.
Recommend minimal safeguards (dedupe keys, locks, at-least-once handling) aligned with tickets.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=7.4 impact=8 confidence=0.8 effort=2 horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-logs-metrics-traces.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09
Category: observability

From the attached tickets bundle, infer required observability: logs, metrics, traces, and user-visible diagnostics.
List:
- what to log/measure
- cardinality risks
- minimal dashboards/alerts implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=7.0 impact=7 confidence=0.8 effort=2 horizon=Near category=observability reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-error-taxonomy.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10
Category: observability

Using the attached tickets bundle, define an error taxonomy consistent with ticket failure modes:
- user errors vs system errors
- retryable vs non-retryable
- how errors should surface (CLI exit codes, UI states, logs)

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=7.6 impact=9 confidence=0.75 effort=3 horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-authz-model.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11
Category: permissions

From the attached tickets bundle, infer the permissions model (roles, capabilities, scopes).
List:
- what operations require permissions
- how permissions are granted/revoked
- audit requirements implied

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=7.0 impact=8 confidence=0.75 effort=3 horizon=Near category=permissions reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-secret-handling.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12
Category: permissions

Using the attached tickets bundle, identify sensitive data/secret handling needs.
Recommend:
- where secrets come from (env, files, vault)
- redaction rules
- least-privilege defaults

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=7.2 impact=8 confidence=0.8 effort=2 horizon=Near category=migrations reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-data-changes.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #13
Category: migrations

From the attached tickets bundle, infer any data/schema/config migrations needed.
For each migration:
- trigger/versioning
- rollout plan
- rollback strategy

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=migrations reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compatibility.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #14
Category: migrations

Using the attached tickets bundle, identify backwards/forwards compatibility requirements during migration windows.
Recommend minimal compatibility shims or staged rollout steps.

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=7.4 impact=8 confidence=0.8 effort=2 horizon=Immediate category=UX flows reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-primary-journeys.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #15
Category: UX flows

From the attached tickets bundle, map the primary user journeys implied by tickets.
For each journey:
- entry points
- steps/screens/commands
- success criteria and user feedback

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=6.9 impact=7 confidence=0.8 effort=2 horizon=Near category=UX flows reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edge-cases.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16
Category: UX flows

Using the attached tickets bundle, list UX edge cases and failure UX:
- partial completion
- retries
- cancellation
- timeouts
- conflict resolution

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=7.8 impact=9 confidence=0.8 effort=2 horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-top-risks.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17
Category: failure modes

From the attached tickets bundle, enumerate the most likely failure modes.
For each failure mode:
- detection signal
- mitigation
- user-visible behavior

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=7.0 impact=8 confidence=0.75 effort=3 horizon=Near category=failure modes reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-test-plan.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18
Category: failure modes

Using the attached tickets bundle, propose a minimal test plan that covers the highest-risk failure modes.
Include:
- unit vs integration coverage split
- fixtures/mocks needed
- one smallest test to write first

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=7.3 impact=8 confidence=0.8 effort=2 horizon=Near category=feature flags reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-needed.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19
Category: feature flags

From the attached tickets bundle, infer where feature flags or staged rollouts are needed.
For each flag:
- what it gates
- default value
- sunset plan

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=6.8 impact=7 confidence=0.8 effort=2 horizon=Near category=feature flags reference=Unknown
{{oracle_cmd}} {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-observability.md" \
  -f "{{ticket_bundle_path}}" {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20
Category: feature flags

Using the attached tickets bundle, propose how to observe/validate a flagged rollout:
- success metrics
- rollback triggers
- logging/alert changes while enabled

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

## Coverage check

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