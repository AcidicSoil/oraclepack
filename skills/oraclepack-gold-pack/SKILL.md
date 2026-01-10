---
name: oraclepack-gold-pack
description: Generate a single canonical Stage-1 oraclepack question pack as Markdown:exactly one ```bash fence containing exactly 20 steps (01..20) with strict ROI header tokens, per-step --write-output, fixed categories + coverage check. Use when you need a gold, runner-ingestible pack template (Stage 1 only; not Stage 3 taskify).
metadata:
  short-description: Gold Stage-1 oraclepack pack generator + validators
---

# Oraclepack Gold Pack (Stage 1)

This skill produces the **canonical Stage-1** oraclepack question pack (20 Oracle CLI calls). It is intentionally strict to prevent schema drift.

**Non-negotiable contract (pack output):**
- Exactly **one** fenced code block: starts with exactly ` ```bash` on its own line, and ends with exactly ` ````
- No other fenced code blocks anywhere in the pack.
- Exactly **20** steps, numbered **01..20** in order.
- Every step has a header line matching:
  - `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes `--write-output "<out_dir>/<nn>-<slug>.md"`.
- Categories are fixed to this exact set (no additions/renames):
  - `contracts/interfaces`
  - `invariants`
  - `caching/state`
  - `background jobs`
  - `observability`
  - `permissions`
  - `migrations`
  - `UX flows`
  - `failure modes`
  - `feature flags`
- Pack ends with a **Coverage check** section listing all 10 categories as `OK` or `Missing(<step ids>)`.

The pack template is the contract. The scratch doc is **not** a pack format.

References:
- Contract template: `references/oracle-pack-template.md`
- Repo discovery: `references/inference-first-discovery.md`
- Attachment rules: `references/attachment-minimization.md`
- Scratch playbook (not pack): `references/oracle-scratch-format.md`

## Quick start

1) Generate a pack file (intended path):
- `docs/oracle-pack-YYYY-MM-DD-HHMMSS.md`

1) Validate it (recommended before running oraclepack):
- `python3 scripts/validate_pack.py docs/oracle-pack-YYYY-MM-DD-HHMMSS.md`

1) Optional attachment lint:
- `python3 scripts/lint_attachments.py docs/oracle-pack-YYYY-MM-DD-HHMMSS.md`


## Setup wizard (always run)

For each available argument listed in this skill, ask the user:

```
1) Use default
2) Enter custom value
```

Collect a choice for every argument (no skips). Then output a summary table of chosen values before running.

For each available argument listed in this skill, ask the user:

```
1) Use default
2) Enter custom value
```

Collect a choice for every argument (no skips). Then run the workflow using the resolved values.

## Inputs (do not ask follow-ups)

Interpret the user’s trailing text as conceptual `{{args}}`. Extract:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD-HHMMSS`)
- `oracle_cmd` (default `oracle`)
- `oracle_flags` (default `--files-report`)
- `engine` (`api|browser`; optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(engine flag unknown)`)
- `model` (optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(model flag unknown)`)
- `extra_files` (default empty; if provided, append **literally** to every command)

If any value is missing: use defaults and proceed.

## Workflow (deterministic)

### 1) Read the contract template first
Open `references/oracle-pack-template.md` and treat it as the **single source of truth** for formatting.

If `2`, collect KEY=value args and apply them in the pack’s Parsed args section; otherwise use defaults.

### 3) Repo discovery (inference-first)
Follow `references/inference-first-discovery.md`:
- Read a small set of “anchors” first.
- Infer what’s present in the repo.
- Only then choose the best 1–2 attachments per step.

### 4) Plan the 20 probes (2 per category)
Use the fixed categories and produce **exactly 2 steps per category** (20 total). Keep each step’s prompt focused and non-overlapping.

For each step:
- Pick a **reference anchor** (`reference=` token): `{path}:{symbol}` OR `{path}` OR `Unknown`.
- Pick ≤2 attachments (or fewer) using `references/attachment-minimization.md`.
- Ensure the prompt asks for:
  - Direct answer (bullets)
  - Risks/unknowns
  - Next smallest concrete experiment
  - Missing artifact patterns to request if evidence is insufficient

### 5) Emit the pack (single file)
Produce exactly one Markdown document with:
- Title + parsed args section (plain markdown; no code fences)
- Exactly one ` ```bash` fence containing the 20 steps
- A Coverage check section after the fence

### 6) Validate and correct drift
Run:
- `python3 scripts/validate_pack.py <pack.md>`
If it fails, fix the pack until it passes.

Optionally run:
- `python3 scripts/lint_attachments.py <pack.md>`
If it fails, reduce attachments to ≤2 per step (before any literal `extra_files`).

Then write a minimal manifest:
- `python3 scripts/write_manifest.py --pack-path <pack.md>`

## Output contract

When invoked, you produce:
- One runner-ingestible Markdown pack (intended filename: `docs/oracle-pack-YYYY-MM-DD-HHMMSS.md`)
- A minimal manifest (e.g., `docs/oracle-pack-YYYY-MM-DD-HHMMSS.manifest.json`) with pack_path and date.

You do **not**:
- Run oraclepack
- Generate Stage-3 “action packs” (that is `oraclepack-taskify`)

## Failure modes (do not guess)

- Missing repo evidence → set `reference=Unknown`, attach fewer files, and explicitly request missing file/path patterns in the prompt.
- Uncertain CLI flag support (`engine`, `model`) → omit flags and write `TODO(engine/model flag unknown)` in the pack’s parsed args notes (do not invent flags).
- Any schema drift → fix until `scripts/validate_pack.py` passes.

## Invocation examples

1) “Generate a gold oraclepack Stage-1 pack for this repo. out_dir=docs/oracle-questions-2026-01-06”
2) “Make the strict 20-step pack for codebase_name=AcmeAPI constraints=‘no DB changes’”
3) “Create the canonical pack; engine=browser model=gpt-5.2-pro (if supported)”
4) “Produce the gold pack; add extra_files='-f docs/ARCHITECTURE.md -f docs/API.md'”
5) “Regenerate this pack but fix headers and coverage check so it validates.”
