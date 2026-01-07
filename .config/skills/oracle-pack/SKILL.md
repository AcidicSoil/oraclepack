---
name: oracle-pack
description: Generate exactly 20 evidence-cited, ROI-ranked strategist questions for the current repository (12 immediate + 8 strategic), then emit them as runnable @steipete/oracle CLI commands with minimal targeted file attachments and deterministic per-question output paths (a copy/paste-ready “oracle question pack” in the same scratch-style format as oracle-scratch.md).
---

## Quick start

Use this skill when the user wants strategist questions **and** wants to run each question through Oracle as a separate command (so a second model can answer with repo context).

Invocation:

- `$oracle-pack <free-text args>`

Primary output: a Markdown doc whose main content is a **single** fenced `bash` block containing **exactly 20** `oracle ...` commands (copy/paste-ready), in the same “scratch-style” format shown in `references/oracle-scratch-format.md`.

## Inputs

- Repository contents (current working directory).
- Free-text args (may include): `codebase_name`, `constraints`, `non_goals`, `team_size`, `deadline`, plus Oracle pack controls:
  - `out_dir` (default: `oracle-out`)
  - `oracle_cmd` (default: `oracle`)
  - `oracle_flags` (default: `--files-report`)
  - `extra_files` (optional: comma-separated extra `-f/--file` entries to include in *every* command; default: none)

## Deterministic search/tooling rules (must follow)

- Prefer deterministic, non-interactive commands so runs are reproducible.
- Before using `ck` in a session, **always run**: `ck --help` and only rely on flags that `ck --help` confirms.
- Default search tool is `ck`:
  - Conceptual: `ck --sem "<query>"`
  - Literal: `ck --regex "<pattern>"`
  - Best of both: `ck --hybrid "<query>"`
  - For post-processing: prefer `ck --jsonl | jq ... | sort | head`
- Structural search: use `ast-grep` when syntax/shape matters (see `references/inference-first-discovery.md`).
- File finding: prefer `fd` (filename/path/extension filters).

## Args parsing (no follow-ups)

Extract values if present; otherwise set defaults:

- `codebase_name`: `Unknown`
- `constraints`: `None`
- `non_goals`: `None`
- `team_size`: `Unknown`
- `deadline`: `Unknown`
- `out_dir`: `oracle-out`
- `oracle_cmd`: `oracle`
- `oracle_flags`: `--files-report`
- `extra_files`: (empty)

Honor `constraints`. Explicitly exclude `non_goals`.

## Workflow (inference-first discovery → evidence → questions → oracle pack)

### 1) Inference-first discovery (adaptive, codebase-search compliant)

Goal: infer *what this repo is* and *where the truth likely lives* before broad sweeps.

Do this in order:

1. **Discover the search interface (required)**
   - Run: `ck --help`
   - If `ck` indicates it needs setup/indexing, follow the instructions printed by `ck`.

2. **Map the repo surface (cheap, high-signal)**
   - Use `fd` to list likely roots deterministically:
     - `fd . -d 2`
     - `fd -p README.md`
     - `fd -p package.json`
     - `fd -p pyproject.toml`
     - `fd -p go.mod`
     - `fd -p Cargo.toml`
   - Read the smallest set of “index” files that describe structure:
     - README / docs index if present
     - primary manifests/config
     - obvious entrypoint referenced by scripts/manifests

3. **Infer stack + conventions from evidence**
   - From manifests/config, infer runtime + framework signals (dependencies, filenames, directory conventions).
   - From entrypoints, infer wiring patterns (router, DI/container, job scheduler, migration tool).

4. **Derive a “search plan” from inferred signals**
   - Prefer following imports/references from entrypoints into:
     - routing/handlers
     - auth middleware/policies
     - job/queue registration
     - data layer/migrations
     - feature flags config
     - observability bootstrap
   - Use `ck` in the smallest inferred app roots first:
     - Conceptual: `ck --sem "<subsystem intent>"`
     - Hybrid when unsure: `ck --hybrid "<subsystem intent>"`
   - If results are broad/noisy:
     - Narrow to directories/files from the top hits, then re-run `ck`.
     - Use `ast-grep` when you need structure, not text matches.

This improves efficiency vs. hard-targeting a long list of patterns up front: you start with the repo’s own “self-description” and expand only as needed.

### 2) Evidence harvesting (collect anchors before writing questions)

Collect **at least 20 candidate anchors** as one of:

- `{path}:{symbol}` (preferred: nearest function/class/type/handler)
- `{endpoint}` (literal endpoint string)
- `{event}` (job/queue/event name)

For each required category, attempt to identify at least one anchor:

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

Use deterministic selection of candidates:
- If using `ck --jsonl`, post-process to stable top-N:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- If using plain output, cap deterministically with `head`.

If evidence for a category cannot be found:

- still generate a question for it with reference `Unknown`
- explicitly state which missing artifact pattern would likely provide it (e.g., `**/migrations/**` not found)

### 3) Generate exactly 20 strategist questions (evidence-cited + minimal experiments)

Produce exactly:

- **12** Immediate Exploration
- **8** Strategic Planning

Each question must include:

- **Reference**: `{path}:{symbol}` OR `{endpoint}` OR `{event}` OR `Unknown`
- **Question**: incisive, feasibility-focused, no scope creep
- **Rationale**: exactly one sentence
- **Smallest experiment today**: exactly one concrete action runnable today (targeted `ck` query, targeted `ast-grep` pattern, trace a codepath, add a log line, minimal unit/integration test, run migration dry-run, execute a single endpoint, etc.)

No duplicates (by intent or reference).

### 4) Score and order by near-term ROI (required math + sorting)

For each question compute:

- `ROI = (impact * confidence) / effort`
- `impact`, `confidence`, `effort` ∈ `{0.1..1.0}` (one decimal)

Sort all 20 descending by ROI; break ties by lower effort.

### 5) Convert the 20 questions into an Oracle “question pack” (final deliverable)

For each question (in final sorted order), emit a runnable command using:

- command: `{{oracle_cmd}}` (default `oracle`)
- include `{{oracle_flags}}` (default `--files-report`)
- deterministic output file: `--write-output "<out_dir>/<nn>-<slug>.md"`
  - `<nn>` = zero-padded 2-digit index (01..20)
  - `<slug>` = short, filesystem-safe slug derived from `{category}` + a hint of `{reference}` (fallback to category only)
- prompt: a structured prompt that includes the strategist question fields and constraints:
  - reference, question, rationale, smallest experiment today
  - constraints + non_goals (explicit)
  - requested Oracle answer shape (concise, evidence-cited, and actionable)

#### Attachment selection (must be minimal, evidence-driven)

For each command, attach only the smallest set of files that lets Oracle answer correctly:

- If reference is `{path}:{symbol}`:
  - attach that `path`
  - optionally attach **one** upstream config/router/entrypoint file that shows how it’s invoked (only if needed)
- If reference is `{endpoint}` or `{event}`:
  - attach the file(s) where you found the literal endpoint/event definition
- If reference is `Unknown`:
  - attach only “index” files (README + manifest + 1 best-guess entrypoint), and in the prompt state what artifact pattern is missing

Follow `references/attachment-minimization.md`.

Also:
- If `extra_files` was provided in args, include those `-f` entries in every command (after the evidence-minimal attachments).

### 6) Render final output using the pack template

Use `assets/oracle-pack-template.md` as the strict shape for the final Markdown deliverable.

## Output contract (strict)

Produce exactly one Markdown deliverable that:

1. Matches the structure in `assets/oracle-pack-template.md`
2. Contains **exactly 20** `oracle` invocations total
3. Commands are sorted by ROI (desc; ties by lower effort)
4. Each command:
   - includes `--write-output "<out_dir>/<nn>-<slug>.md"`
   - includes minimal targeted `-f/--file` attachments
   - embeds the strategist question (reference, question, rationale, experiment) in the `-p/--prompt` text

## Failure modes

- Missing/insufficient evidence:
  - use reference `Unknown`
  - state the missing artifact pattern that would provide evidence
  - keep the question actionable and experiment minimal
  - keep attachments limited to index files (don’t attach huge globs)
- Search tool limitations (`ck` missing/setup required):
  - run `ck --help` and follow printed setup; if unavailable, fall back to reading index files + minimal `fd`-based locating
  - proceed with partial evidence; mark affected references `Unknown`
- Ambiguous args:
  - apply defaults without follow-ups

## References

- `assets/oracle-pack-template.md` — exact final output shape
- `references/inference-first-discovery.md` — inference-driven search plan aligned to `ck`/`ast-grep`/`fd`
- `references/attachment-minimization.md` — deterministic rules for minimal `-f` selections
- `references/oracle-scratch-format.md` — the target “scratch-style” formatting to mimic

## Invocation examples

- “Generate an oracle question pack for this repo. codebase_name=Payments API; constraints=no new infra; non_goals=mobile app; team_size=3; deadline=2 weeks; out_dir=artifacts/oracle”
- “Create oracle strategist questions; constraints=ship bugfixes only; deadline=Friday; out_dir=oracle-out”
- “Produce an oracle pack: focus on auth + jobs + migrations; non_goals=refactor; out_dir=oracle-q”
- “Generate strategist oracle commands; constraints=keep DB schema stable; team_size=2; deadline=1 week”
- “Make an Oracle question pack; oracle_cmd='npx -y @steipete/oracle'; oracle_flags='--engine browser  --files-report'; out_dir=oracle-out”

### Write output to `docs/` (required)

- Ensure a `docs/` directory exists at the repo root (create it if missing).
- Write the complete final deliverable to:
  - `docs/oracle-pack-YYYY-MM-DD.md` (use today’s date; ISO 8601)
  - Fallback if date is unavailable: `docs/oracle-pack.md`
- The file must contain only the deliverable Markdown (no extra preamble/epilogue).
- In the assistant response, print the chosen path first on a single line: `Output file: <path>`, then print the same Markdown content.
