# Attachment minimization rules (Grouped Tickets Stage 1 — Direct Attach)

Objective: keep each group pack focused and portable.

## Ticket attachments

- Ticket files are attached directly in each step via `${ticket_args[@]}`.
- Use `group_max_files` (default 25) to bound per-pack ticket count.
- If a group is larger than the cap, split into multiple packs (part 1..N).

## Non-ticket attachments (repo evidence)

- Keep explicit non-ticket attachments to **0–1 per step**.
- Prefer a single high-signal file that clarifies contracts or a key code path.

## extra_files (literal append)

- If `extra_files` is provided, append it literally to every oracle command.
- It may include additional `-f/--file` flags.
- Place `extra_files` on its own line with a comment:
  - `# extra_files appended literally`

