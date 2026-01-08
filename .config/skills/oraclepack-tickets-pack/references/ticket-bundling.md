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
