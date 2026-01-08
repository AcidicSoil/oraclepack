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
