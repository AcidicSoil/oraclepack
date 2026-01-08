# Attachment minimization rules (Codebase Stage 1 — Direct Attach)

Objective: keep each group pack focused and portable.

## Code attachments

- Code files are attached directly in each step via `${code_args[@]}`.
- Use `group_max_files` (default 200) to bound per-pack file count.
- If a group is larger than the cap, split into multiple packs (part 1..N).
- Prefer code_glob + include_exts to avoid irrelevant files.

## Non-code attachments (extra_files)

- Keep explicit non-code attachments to **0–1 per step**.
- Prefer a single high-signal file (e.g., README, architecture doc).

## extra_files (literal append)

- If `extra_files` is provided, append it literally to every oracle command.
- It may include additional `-f/--file` flags.
- Place `extra_files` on its own line with a comment:
  - `# extra_files appended literally`
