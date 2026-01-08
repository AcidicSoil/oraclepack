<filetree>
Project Structure:
└── .config
    └── skills
        └── oraclepack-tickets-pack
            ├── references
            │   ├── attachment-minimization.md
            │   ├── ticket-bundling.md
            │   └── tickets-pack-template.md
            ├── scripts
            │   ├── lint_attachments.py
            │   └── validate_pack.py
            └── SKILL.md

</filetree>

<source_code>
.config/skills/oraclepack-tickets-pack/SKILL.md
```
---
name: oraclepack-tickets-pack
description: Generate a runner-ingestible oraclepack Stage-1 question pack (single Markdown doc) driven by `.tickets/` content. Exactly one ```bash fence, exactly 20 steps (01..20), strict ROI header tokens, deterministic ticket bundling, minimal per-step attachments, coverage check.
metadata:
  short-description: Ticket-driven Stage-1 oraclepack pack generator + validators
---

# oraclepack-tickets-pack (Stage 1)

## Purpose

Produce a **ticket-driven** oraclepack Stage-1 pack (Markdown) that is **runner-ingestible** and **schema-compatible** with existing oraclepack pack format.

The generated pack’s questions and minimal attachments are guided primarily by a deterministic **ticket bundle** built from `.tickets/` (or explicit ticket paths).

## Use when

- You have tickets stored under `.tickets/` (or you can provide explicit ticket file paths), and you want a strict 20-step oraclepack Stage-1 question pack that:
  - references tickets as the primary context in every step
  - uses minimal attachments per step
  - preserves existing oraclepack Markdown pack schema (backward compatible)

## Hard requirements (output contract)

The produced pack (single Markdown file) MUST satisfy:

1) **Schema safety / compatibility**
- Do not break the existing oraclepack Markdown pack schema.
- Exactly **one** fenced code block labeled `bash` in the entire document.
- No other fenced code blocks anywhere (no additional ``` fences).

1) **Runner-ingestible strictness**
- Exactly **20** steps inside the single `bash` fence.
- Steps are numbered **01..20** in order.
- Each step header starts with `# NN)` and includes the strict header tokens in the header line:
  - `ROI= ... impact= ... confidence= ... effort= ... horizon= ... category= ... reference= ...`
- Every step includes:
  - `--write-output "<out_dir>/<nn>-<slug>.md"`
- Steps must be **self-contained** and must **not rely on shell variables created in previous steps**.

1) **Attachment minimization**
- Default **0–2 native attachments per step**.
- Each step should normally attach:
  - `-f "<ticket_bundle_path>"` (primary context)
  - plus at most **one** additional repo file when needed
- If `extra_files` are provided, append them **literally**, but keep the step’s native attachments ≤2.

1) **Path safety**
- `--write-output` destinations must be deterministic and must not escape `out_dir` (no `..` traversal).
- No absolute write paths.

1) **Determinism**
- Ticket discovery ordering must be stable:
  - lexicographic ordering only
  - no timestamps / mtimes used for selection

The pack MUST end with a **Coverage check** section listing all 10 categories as `OK` or `Missing(<step ids>)`.

## Inputs (do not ask follow-ups)

Parse the user’s trailing text as whitespace-separated `KEY=value` tokens (last-one-wins). Unknown keys ignored.

Supported keys (defaults in parentheses):

- `codebase_name` (`Unknown`)
- `out_dir` (`docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (`oracle`)
- `oracle_flags` (`--files-report`)
- `extra_files` (empty; appended literally)
- `ticket_root` (`.tickets`)
- `ticket_glob` (`**/*.md` relative to `ticket_root`)
- `ticket_paths` (optional; comma-separated explicit ticket files; if present, ignore `ticket_glob`)
- `ticket_bundle_path` (`<out_dir>/_tickets_bundle.md`)
- `mode` (`tickets`; reserved)

Notes:
- `YYYY-MM-DD` is computed at pack generation time for default `out_dir`.
- If oracle flag support is uncertain (engine/model/etc), **omit unsupported flags**; never invent flags.

## Workflow (deterministic)

1) Read:
- `references/ticket-bundling.md` (how to build the bundle deterministically)
- `references/attachment-minimization.md` (attachment limits + extra_files handling)

1) Render a pack by starting from:
- `references/tickets-pack-template.md`

1) Resolve args deterministically:
- Fill placeholders for `out_dir`, `ticket_root`, `ticket_glob`, `ticket_paths`, `ticket_bundle_path`, `oracle_cmd`, `oracle_flags`, `extra_files`.
- Ensure ticket selection and concatenation are lexicographically stable.

1) Ensure the 20 steps are **ticket-scoped**:
- Use the fixed 10 categories (2 steps per category):
  - contracts/interfaces
  - invariants
  - caching/state
  - background jobs
  - observability
  - permissions
  - migrations
  - UX flows
  - failure modes
  - feature flags
- Each step prompt must explicitly reference the **ticket bundle** as primary context.
- Each step prompt must include the mandatory Answer format (4 parts).

1) Validate output:
- `python3 skills/oraclepack-tickets-pack/scripts/validate_pack.py <pack.md>`
- Optional attachment lint:
- `python3 skills/oraclepack-tickets-pack/scripts/lint_attachments.py <pack.md>`

## Failure behavior (must be encoded into the generated pack)

- `ticket_root` missing OR no tickets matched:
  - Still generate the pack.
  - Prelude must write a clear warning into the ticket bundle and emit a clear stderr message.
  - Step 01 prompt must request: “which ticket paths to attach next” (exact missing file/path pattern(s)).

- Oracle flag uncertainty:
  - Omit unsupported flags.
  - Never invent flags.

- Output path ambiguity:
  - Validator must catch missing `--write-output`, invalid numbering, invalid headers, missing coverage check, or unsafe write paths.

## Deliverables

This skill produces:
- One runner-ingestible Stage-1 oraclepack pack (single Markdown doc) that passes `scripts/validate_pack.py`.

## Reference assets

- Template pack: `references/tickets-pack-template.md`
- Ticket bundling: `references/ticket-bundling.md`
- Attachment rules: `references/attachment-minimization.md`
- Validator: `scripts/validate_pack.py`
- Optional linter: `scripts/lint_attachments.py`
```

.config/skills/oraclepack-tickets-pack/references/attachment-minimization.md
```
# Attachment minimization rules (Tickets Stage 1 packs)

Objective: keep oracle calls fast, portable, and deterministic by attaching the minimum evidence per step.

## Hard limits

- Default **native attachments**: **0–2 per step** (`-f/--file`).
- In tickets packs, the ticket bundle (`ticket_bundle_path`) is typically **the first native attachment**.
- If you need more than 2 native attachments, the step is not scoped tightly enough: split or reduce.

## extra_files (literal append)

- If `extra_files` is provided, it must be appended **literally** to every oracle command.
- It may contain additional `-f/--file` flags.
- To keep linting reliable and preserve the “native attachments ≤2” rule:
  - place `extra_files` on its own line in each command,
  - preceded by a comment line containing: `extra_files appended literally`.

This lets `scripts/lint_attachments.py` treat that line as “extra” and not part of the native attachment count.

## What to attach (rule of thumb)

For each step, prefer:
1) Ticket bundle: `-f "<ticket_bundle_path>"`
2) One repo file that best supports the question:
   - a definition/contract file (types, schemas, CLI/TUI surface), OR
   - a use-site/enforcement file

If you can’t pick confidently:
- attach only the ticket bundle
- set `reference=Unknown`
- ensure the prompt requests the exact missing file/path pattern(s) to attach next

## Avoid these attachment anti-patterns

- Attaching entire directories when one file is enough.
- Attaching duplicates.
- Attaching generated/lock files unless the ticket explicitly requires it.
- Attaching secrets.
```

.config/skills/oraclepack-tickets-pack/references/ticket-bundling.md
```
# Ticket bundling (deterministic)

Objective: create a single Markdown file (`ticket_bundle_path`) that provides stable, minimal, high-signal context for all 20 oracle steps.

## Inputs

- `ticket_root` (default `.tickets`)
- `ticket_glob` (default `**/*.md`, relative to `ticket_root`)
- `ticket_paths` (optional; comma-separated explicit files; if present, ignore `ticket_glob`)
- `ticket_bundle_path` (default `<out_dir>/_tickets_bundle.md`)

## Deterministic selection rules

1) If `ticket_paths` is non-empty:
- Split on commas, trim whitespace.
- Use that list exactly.
- Sort lexicographically by path string.

2) Else:
- If `ticket_root` does not exist: select nothing.
- Glob `ticket_root/ticket_glob`.
- Sort lexicographically by path string.

Hard rule: do not use timestamps, mtimes, file sizes, or “newest” semantics.

## Bundle format

The bundle should include:

1) Header:
- codebase name (if available)
- the selection rules and resolved values:
  - ticket_root, ticket_glob, ticket_paths (or “(none)”)
  - ordering: lexicographic by path

2) Per-ticket sections (in lex order):
- ticket title (best-effort):
  - first Markdown heading (`# ...`) if present, else first non-empty line, else “Untitled”
- ticket path
- key sections or excerpt:
  - if ticket content is small, include full content
  - otherwise include common sections when present (examples):
    - Description / Context / Problem
    - Proposal / Solution
    - Acceptance Criteria
    - Repro steps / Expected / Actual
    - Notes / Links
  - and include a “[... truncated ...]” marker when partial

## Failure handling requirements

If `ticket_root` is missing OR no tickets matched:
- Still create `ticket_bundle_path`.
- Include a **WARNING** section explaining:
  - what was attempted (root + glob or explicit paths)
  - that the bundle is empty
  - that Step 01 should request which ticket paths to attach next (exact missing file/path pattern(s))

## Why bundling exists

- Ensures every step uses the same primary evidence.
- Reduces per-step attachments (bundle is 1 attachment).
- Improves determinism and portability of oracle calls.
```

.config/skills/oraclepack-tickets-pack/references/tickets-pack-template.md
```
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
```

.config/skills/oraclepack-tickets-pack/scripts/lint_attachments.py
```
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    header: str
    lines: List[str]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_bash_fence(lines: List[str]) -> List[str]:
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]
    if len(fence_idxs) != 2:
        raise ValueError(f"Expected exactly one fenced block (2 fence lines). Found {len(fence_idxs)}.")
    open_i, close_i = fence_idxs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash.")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ```.")
    return [ln.rstrip("\n") for ln in lines[open_i + 1 : close_i]]


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if not header_idxs:
        raise ValueError("No step headers found inside bash fence.")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(Step(n=n, header=block[0], lines=block))
    return steps


def _count_native_attachments(step: Step) -> int:
    """
    Counts -f/--file occurrences excluding:
      - comment lines
      - the literal extra_files line (immediately following the marker comment)
    """
    count = 0
    ignore_next_nonempty = False

    for ln in step.lines[1:]:
        s = ln.strip()
        if not s:
            continue

        # Detect extra_files marker comment; ignore next non-empty line.
        if s.startswith("#") and "extra_files appended literally" in s.lower():
            ignore_next_nonempty = True
            continue

        if ignore_next_nonempty:
            # Skip counting attachments on the extra_files line itself.
            ignore_next_nonempty = False
            continue

        if s.startswith("#"):
            continue

        count += len(re.findall(r"(?<!\S)(-f|--file)(?!\S)", ln))
    return count

def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        native = _count_native_attachments(step)
        if native > 2:
            errors.append(
                f"Step {step.n}: has {native} native attachments; must be <= 2 (ticket bundle + at most one repo file)."
            )

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Attachment lint passed (native attachments <= 2 per step; extra_files line excluded).")

def main() -> None:
    p = argparse.ArgumentParser(
        description="Lint ticket-driven oraclepack Stage-1 packs for native attachments (<=2 per step, excluding literal extra_files line)."
    )
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
```

.config/skills/oraclepack-tickets-pack/scripts/validate_pack.py
```
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path, PurePosixPath
from typing import Dict, List, Tuple


ALLOWED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

# Required header tokens, in strict order.
HEADER_TOKEN_ORDER = [
    "ROI=",
    "impact=",
    "confidence=",
    "effort=",
    "horizon=",
    "category=",
    "reference=",
]


@dataclass
class Step:
    n: str
    header_line_no: int  # 1-based within fence
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly one fenced code block labeled bash
      - no other fences anywhere
      - opening fence line must be exactly ```bash
      - closing fence line must be exactly ```
    """
    fence_locs = [i for i, ln in enumerate(lines) if re.match(r"^```", ln)]
    if len(fence_locs) != 2:
        # Show all fence-like lines to help debugging.
        details = []
        for i, ln in enumerate(lines):
            if re.match(r"^```", ln):
                details.append(f"line {i+1}: {ln.rstrip()}")
        raise ValueError(
            f"Expected exactly one fenced code block (2 fence lines), found {len(fence_locs)} fence line(s). "
            + ("Fences: " + "; ".join(details) if details else "")
        )

    open_i, close_i = fence_locs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash on its own line (no spaces).")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` on its own line (no spaces).")
    if close_i <= open_i:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_i + 1 : close_i]
    outside_lines = lines[:open_i] + lines[close_i + 1 :]
    return open_i, close_i, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if len(header_idxs) != 20:
        raise ValueError(f"Expected exactly 20 step headers inside bash fence, found {len(header_idxs)}.")

    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [n for _, n in header_idxs]
    if got != expected:
        raise ValueError(f"Step numbering must be sequential 01..20. Got: {got}")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(
            Step(
                n=n,
                header_line_no=start_i + 1,
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _header_token_positions(header: str) -> Dict[str, int]:
    pos: Dict[str, int] = {}
    for t in HEADER_TOKEN_ORDER:
        pos[t] = header.find(t)
    return pos


def _parse_category_value(header: str) -> str:
    if "category=" not in header:
        return ""
    after = header.split("category=", 1)[1]
    # Category ends at the start of " reference=" (strict contract).
    end = after.find(" reference=")
    if end == -1:
        # As a fallback, try other token starts, though contract expects reference= last.
        for token in [" ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            p = after.find(token)
            if p != -1:
                end = p if end == -1 else min(end, p)
    if end == -1:
        cat = after.strip()
    else:
        cat = after[:end].strip()
    return cat


def _has_nonempty_scalar(header: str, key: str) -> bool:
    # scalar value ends at next whitespace
    m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
    return bool(m and m.group(1).strip())


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    # Strict start.
    if not re.match(rf"^#\s*{re.escape(step.n)}\)\s+", header):
        errors.append(f"Step {step.n}: header must start with '# {step.n})'. Got: {header}")

    # Tokens must appear in strict order.
    pos = _header_token_positions(header)
    for t, p in pos.items():
        if p == -1:
            errors.append(f"Step {step.n}: missing required token '{t}' in header: {header}")

    # Order check (only if all present).
    if all(p != -1 for p in pos.values()):
        last = -1
        for t in HEADER_TOKEN_ORDER:
            if pos[t] <= last:
                errors.append(
                    f"Step {step.n}: token '{t}' is out of order in header. "
                    f"Expected order: {' '.join(HEADER_TOKEN_ORDER)}. Got: {header}"
                )
                break
            last = pos[t]

    # Non-empty values.
    if not _has_nonempty_scalar(header, "ROI"):
        errors.append(f"Step {step.n}: missing/empty ROI= value in header: {header}")
    for k in ["impact", "confidence", "effort", "horizon", "reference"]:
        if not _has_nonempty_scalar(header, k):
            errors.append(f"Step {step.n}: missing/empty {k}= value in header: {header}")

    cat_val = _parse_category_value(header)
    if not cat_val:
        errors.append(f"Step {step.n}: missing/empty category= value in header: {header}")
    elif cat_val not in ALLOWED_CATEGORIES:
        errors.append(
            f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}. Header: {header}"
        )


def _validate_write_output(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)
    # Strict: must use double quotes exactly: --write-output "<path>"
    m = re.search(r'--write-output\s+"([^"]+)"', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output \"...\" (double-quoted) in step block.")
        return

    out_path = m.group(1)

    # Disallow variable expansions in write paths.
    if "$" in out_path or "`" in out_path:
        errors.append(f"Step {step.n}: --write-output path must not contain shell expansions. Got: {out_path}")

    # Disallow absolute writes (and home shortcuts).
    if out_path.startswith("/") or out_path.startswith("~"):
        errors.append(f"Step {step.n}: --write-output path must be relative (no absolute/home paths). Got: {out_path}")

    # Disallow traversal.
    if re.search(r"(^|/)\.\.(/|$)", out_path):
        errors.append(f"Step {step.n}: --write-output path must not contain '..' traversal. Got: {out_path}")

    # Basic shape: <out_dir>/<nn>-<slug>.md
    if "/" not in out_path:
        errors.append(f"Step {step.n}: --write-output path must contain a directory component. Got: {out_path}")
        return

    filename = out_path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(f"Step {step.n}: --write-output filename must start with '{step.n}-'. Got: {filename}")
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with '.md'. Got: {filename}")

    # Extra guard: ensure PurePosixPath doesn't include '..' (covers odd strings like 'a/../b').
    try:
        parts = PurePosixPath(out_path).parts
        if ".." in parts:
            errors.append(f"Step {step.n}: --write-output path contains '..' segment (unsafe). Got: {out_path}")
    except Exception:
        # Non-fatal; already handled by regex.
        pass


def _validate_ticket_bundle_reference(step: Step, errors: List[str]) -> None:
    joined = "\n".join(step.block_lines)

    # Require the bundle to be mentioned/attached.
    if "_tickets_bundle" not in joined:
        errors.append(
            f"Step {step.n}: must reference the ticket bundle (expected '_tickets_bundle' in step block)."
        )

    # Require a file attachment pointing to the bundle, double-quoted for stability.
    if re.search(r'(?<!\S)(-f|--file)(?!\S)\s+"[^"\n]*_tickets_bundle[^"\n]*"', joined) is None:
        errors.append(
            f"Step {step.n}: must attach the ticket bundle via -f/--file \"..._tickets_bundle...\"."
        )


def _validate_answer_format(step: Step, errors: List[str]) -> None:
    hay = "\n".join(step.block_lines).lower()
    required = [
        "answer format:",
        "direct answer",
        "risks/unknowns",
        "next smallest concrete experiment",
        "if evidence is insufficient",
        "missing file/path pattern",
    ]
    missing = [s for s in required if s not in hay]
    if missing:
        errors.append(f"Step {step.n}: prompt missing required Answer format components: {missing}")


def _validate_category_counts(steps: List[Step], errors: List[str]) -> None:
    counts: Dict[str, List[str]] = {c: [] for c in ALLOWED_CATEGORIES}
    for st in steps:
        cat = _parse_category_value(st.header_line)
        if cat in counts:
            counts[cat].append(st.n)

    bad = []
    for cat, ids in counts.items():
        if len(ids) != 2:
            bad.append(f"{cat}={len(ids)} (steps={ids})")
    if bad:
        errors.append(
            "Category distribution must be exactly 2 steps per category (20 total). Problems: " + ", ".join(bad)
        )


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    if m is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    after = text[m.end() :]
    for cat in ALLOWED_CATEGORIES:
        # Require a line like: "- <cat>: OK" OR "- <cat>: Missing(01,02)"
        pat = rf"^\s*[-*]\s+{re.escape(cat)}\s*:\s*(OK|Missing\([^)]*\))\s*$"
        if re.search(pat, after, flags=re.MULTILINE) is None:
            errors.append(f'Coverage check missing/invalid line for category: "{cat}"')


def validate_pack(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)

    try:
        _, _, fence_lines, outside_lines = _extract_single_bash_fence(lines)
    except ValueError as e:
        _fail([str(e)])

    try:
        steps = _parse_steps(fence_lines)
    except ValueError as e:
        _fail([str(e)])

    errors: List[str] = []
    for st in steps:
        _validate_header(st, errors)
        _validate_write_output(st, errors)
        _validate_ticket_bundle_reference(st, errors)
        _validate_answer_format(st, errors)

    _validate_category_counts(steps, errors)
    _validate_coverage_check(outside_lines, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against tickets Stage-1 contract.")

def main() -> None:
    p = argparse.ArgumentParser(
        description="Validate a ticket-driven oraclepack Stage-1 pack (single bash fence, 20 steps, strict headers/tokens, safe write paths, ticket bundle references, coverage check)."
    )
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path)


if __name__ == "__main__":
    main()
```

</source_code>