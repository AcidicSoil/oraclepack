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
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
# Prelude (allowed inside the single bash fence)
# - Creates out_dir deterministically
# - Builds ticket_bundle_path deterministically from ticket_root/ticket_glob OR ticket_paths
# - Uses lexicographic ordering only (no mtime/timestamps)

set -euo pipefail

mkdir -p "{{out_dir}}"

python3 - <<'PY'
from __future__ import annotations

import os
from pathlib import Path

CODEBASE_NAME = "{{codebase_name}}"
OUT_DIR = "{{out_dir}}"
TICKET_ROOT = "{{ticket_root}}"
TICKET_GLOB = "{{ticket_glob}}"
TICKET_PATHS = "{{ticket_paths}}".strip()
BUNDLE_PATH = "{{ticket_bundle_path}}"

root = Path(TICKET_ROOT)

def read_text(p: Path) -> str:
    try:
        return p.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return p.read_text(encoding="utf-8", errors="replace")

def title_from_md(text: str) -> str:
    for ln in text.splitlines():
        s = ln.strip()
        if s.startswith("#"):
            return s.lstrip("#").strip()[:160] or "Untitled"
    for ln in text.splitlines():
        s = ln.strip()
        if s:
            return s[:160]
    return "Untitled"

def select_key_sections(text: str) -> str:
    # Heuristic: if small, include all; else include common ticket sections + top excerpt.
    lines = text.splitlines()
    if len(text) <= 8000 and len(lines) <= 250:
        return text

    keep = []
    wanted = {"description", "context", "problem", "proposal", "solution", "acceptance criteria", "ac", "steps", "repro", "expected", "actual", "notes", "links"}
    i = 0
    while i < len(lines):
        ln = lines[i]
        s = ln.strip()
        if s.startswith("##"):
            hdr = s.lstrip("#").strip().lower()
            if hdr in wanted:
                keep.append(ln)
                i += 1
                # capture until next heading
                while i < len(lines) and not lines[i].lstrip().startswith("#"):
                    keep.append(lines[i])
                    i += 1
                continue
        i += 1

    # Fallback excerpt if no sections matched
    if not any(l.strip() for l in keep):
        excerpt = "\n".join(lines[:200])
        return excerpt + "\n\n[... truncated ...]\n"
    return "\n".join(keep) + "\n\n[... truncated ...]\n"

def resolve_ticket_files() -> list[Path]:
    if TICKET_PATHS:
        items = [p.strip() for p in TICKET_PATHS.split(",") if p.strip()]
        return sorted([Path(p) for p in items], key=lambda p: str(p))
    if not root.exists():
        return []
    return sorted(list(root.glob(TICKET_GLOB)), key=lambda p: str(p))

tickets = resolve_ticket_files()

bundle_lines = []
bundle_lines.append(f"# Tickets bundle — {CODEBASE_NAME}")
bundle_lines.append("")
bundle_lines.append("## Selection rules (deterministic)")
bundle_lines.append(f"- ticket_root: {TICKET_ROOT}")
bundle_lines.append(f"- ticket_glob: {TICKET_GLOB}")
bundle_lines.append(f"- ticket_paths: {TICKET_PATHS or '(none)'}")
bundle_lines.append("- ordering: lexicographic by path")
bundle_lines.append("")

if not tickets:
    bundle_lines.append("## WARNING")
    if TICKET_PATHS:
        bundle_lines.append("- No tickets were found from ticket_paths (check paths).")
    else:
        bundle_lines.append("- ticket_root missing or no tickets matched the glob.")
    bundle_lines.append("- This bundle is empty; Step 01 should request which ticket paths to attach next.")
    bundle_lines.append("")

for p in tickets:
    try:
        txt = read_text(p)
    except Exception as e:
        txt = f"[ERROR reading file: {e}]"
    title = title_from_md(txt)
    body = select_key_sections(txt)
    bundle_lines.append(f"## {title}")
    bundle_lines.append(f"- path: {p}")
    bundle_lines.append("")
    bundle_lines.append(body)
    bundle_lines.append("")
    bundle_lines.append("---")
    bundle_lines.append("")

out_path = Path(BUNDLE_PATH)
out_path.parent.mkdir(parents=True, exist_ok=True)
out_path.write_text("\n".join(bundle_lines).rstrip() + "\n", encoding="utf-8")

print(f"OK: wrote ticket bundle: {out_path}")
PY

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-ticket-surface.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-integration-points.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-correctness-guardrails.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-validation-boundaries.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-ticket-artifacts.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-determinism-consistency.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-ticket-implications.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-reliability-controls.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-required-signals.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-gaps-and-metrics.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-security-constraints.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-enforcement-chokepoints.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-schema-changes.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compat-guardrails.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-ticketed-user-journeys.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edge-cases-and-gotchas.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-taxonomy-from-tickets.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-resilience-and-recovery.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-rollout-controls.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-lifecycle-and-flag-debt.md" \
  -f "{{ticket_bundle_path}}" \
  # optional (at most one): -f "<best_repo_file_path>" \
  # extra_files appended literally (may be empty; may include -f/--file):
  {{extra_files}} \
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

```