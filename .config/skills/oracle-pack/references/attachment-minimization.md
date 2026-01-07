<!-- path: ~/.codex/skills/oracle-pack/references/attachment-minimization.md -->
# Attachment minimization rules (evidence-driven)

Objective: keep Oracle inputs small and high-signal. Attach only what Oracle needs to answer the single question.

## Always exclude

- Secrets (`.env`, key files, tokens)
- Large generated dirs unless explicitly needed (`node_modules`, `dist`, `build`, `coverage`, etc.)
- Broad globs like `src/**` unless the question has `Unknown` reference and no better anchor exists

## Core mapping: reference → files to attach

### Reference = `{path}:{symbol}`

Attach:
1) `path` (the file containing the symbol)

Optionally attach **one** more file (only if it materially affects correctness):
- nearest router/entrypoint that calls into `path`
- config file that wires the subsystem (`config/*`, `settings.*`, `env` loader)
- interface/schema referenced by the symbol (e.g., `openapi.*`, `schema.*`)

Avoid attaching more than 2 files unless the repo uses very small modules and it’s clearly required.

### Reference = `{endpoint}`

Attach:
1) the route map / router file where the endpoint is declared
2) the handler file (if different)

Prefer literal files over globs.

### Reference = `{event}`

Attach:
1) the job/worker registration file where the event is scheduled/declared
2) the worker implementation file (if different)

### Reference = `Unknown`

Attach only “index” evidence:
- README (or closest equivalent)
- primary manifest (`package.json`, `pyproject.toml`, etc.)
- one best-guess entrypoint file (if you found one)

In the prompt, explicitly request the missing artifact pattern to attach next (e.g., “attach `**/migrations/**`”).

## Optional global extras (`extra_files`)

If the user provided `extra_files`, append those attachments to every command *after* the minimal evidence attachments. Use this sparingly.
