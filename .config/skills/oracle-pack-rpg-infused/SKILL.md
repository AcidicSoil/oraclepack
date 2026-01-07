---
name: oracle-pack-rpg-infused
description: Generate a strictly oraclepack-ingestible “oracle strategist question pack” Markdown file under docs/oracle-pack-YYYY-MM-DD.md containing exactly one fenced bash block with exactly 20 numbered oracle commands (01..20), ROI-scored/sorted and category-covered, with an RPG block (WHAT/HOW + dependency edges) embedded in each -p prompt for graph extraction.
---

# Oracle Pack (RPG-infused)

## Quick start

1) Read `assets/oracle-pack-template.md` and keep its structure exactly.
2) Produce exactly one Markdown deliverable and output it verbatim in your final response:
   - First line: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
   - Then the exact Markdown content (no extra commentary, no additional code fences).
3) Ensure the bash block contains exactly 20 steps, numbered `01..20`, each with:
   - One header comment line in the required format
   - Exactly one runnable oracle command line

For the full spec, follow `references/oracle-pack-spec.md`.

## Inputs and defaults

Interpret any user-provided details as “args” and apply these defaults when missing:

- `codebase_name=Unknown`
- `constraints=None`
- `non_goals=None`
- `team_size=Unknown`
- `deadline=Unknown`
- `out_dir=oracle-out`
- `oracle_cmd=oracle`
- `oracle_flags=--files-report`
- `extra_files=empty` (only use if the user provides an explicit mechanism/flag; otherwise omit and keep evidence handling inside the prompt body as text)

If the user supplies their own values, use them verbatim unless they would break the strict output contract.

## Workflow

1) Establish minimal repo understanding (inference-first)
   - Prefer reading index artifacts (e.g., README, top-level config/manifests) before proposing file sweeps.
   - If you cannot access the repo contents, keep `reference=Unknown` and make the “missing artifact pattern(s)” explicit in the prompt body.

2) Plan the 20 questions (before writing any command lines)
   - Use only the required categories (no additions, no removals).
   - Ensure horizon mix: exactly 12 `Immediate` and 8 `Strategic`.
   - Ensure RPG infusion in every step’s `-p` prompt (WHAT/HOW + DependsOn + Phase).
   - Ensure dependencies only point backward (DependsOn may reference only earlier step IDs).
   - Ensure coverage: across all 20 steps, each required category appears at least once.

3) Assign scoring and compute ROI
   - Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
   - Compute `ROI = (impact * confidence) / effort` (one decimal).
   - Sort all 20 steps by ROI descending; tie-break by lower effort.

4) Emit the pack Markdown
   - Copy `assets/oracle-pack-template.md` structure exactly.
   - Fill in metadata values and the 20 steps.
   - Keep exactly ONE fenced bash block.
   - Keep exactly 20 steps (01..20) inside that block.

5) Emit the coverage check section
   - List the 10 required categories exactly once each with `OK` or `Missing(...)`.

6) Validate (optional but recommended when execution is available)
   - Run `scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md` and fix any reported issues.

## Output contract (must hold)

- Exactly one Markdown deliverable.
- Matches `assets/oracle-pack-template.md` structure exactly.
- Exactly one fenced bash block.
- Bash block contains exactly 20 steps, numbered 01..20.
- Each step has:
  - Header line with: `ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
  - Exactly one runnable oracle command including:
    - `${oracle_cmd}` (default `oracle`)
    - `${oracle_flags}` (default `--files-report`)
    - `--write-output "<out_dir>/<nn>-<slug>.md"`
- Steps sorted by ROI desc; tie-break by lower effort.
- Coverage check section present and correctly formatted.

## Failure modes (do not guess)

- Unknown subsystem/file references: set `reference=Unknown` and include explicit missing file/path patterns in section (4) of the prompt body.
- Missing or unclear user args: apply defaults; do not invent tool flags or repo-specific details.
- Potential contract violations: prioritize strict compliance over richer prose.

## Invocation examples

1) “Generate an RPG-infused oracle strategist pack for this repo. out_dir=oracle-out.”
2) “Create the oracle-pack for a monorepo; prioritize caching/state, observability, and permissions. Use oracle_flags=--files-report.”
3) “Make a pack for codebase_name=AcmeWeb, constraints=‘No DB migrations this sprint’, non_goals=‘No UI redesign’.”
4) “Produce a pack that emphasizes feature flags and invariants, but still covers all required categories.”
5) “Generate the pack with conservative confidence values (0.2–0.6) and document missing artifact patterns explicitly.”
