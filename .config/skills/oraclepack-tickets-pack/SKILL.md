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
