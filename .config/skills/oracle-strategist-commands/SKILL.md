---

name: oracle-strategist-commands
description: Generate exactly 20 evidence-anchored strategist questions as runnable @steipete/oracle bash commands (12 immediate + 8 strategic), using inference-first repo discovery to avoid wasted searching.
metadata:
  short-description: 20 ROI-ranked Oracle CLI question commands for a repo (inference-first discovery)
  output: bash
  count: 20
---

## Use when

Use this skill when the user wants incisive, feasibility-focused strategist questions, but **as an executable pack** of **exactly 20** `oracle` CLI commands.

## Deliverable (strict)

Output MUST be:

- A single fenced code block: ```bash
- Optional header lines (env vars, `mkdir -p`, shared preamble)
- Then **exactly 20** Oracle invocations (lines beginning with `oracle` or `npx -y @steipete/oracle`)
- Counts:
  - **12** commands that are “Immediate” (01–12)
  - **8** commands that are “Strategic” (13–20)
- No prose outside the bash code fence.

Each command MUST include:

- `-p/--prompt`
- at least one `-f/--file`
- `--files-report`
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- `--wait` (unless user explicitly requests omitting it)

Optional but recommended:

- a commented preflight line using `--dry-run summary` before each real command.

## Inputs (do not ask follow-ups)

Interpret free-text args as optionally including:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `engine` (`api|browser`; if unspecified, omit)
- `model` (if unspecified, omit)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD`)

If missing, apply defaults and proceed.

## Core principle: inference-first discovery

Avoid “hard-targeting” specific artifact types up front (e.g., hunting for migrations/auth/jobs) because many repos won’t have them. Instead:

1) Gather cheap high-signal evidence (anchors).
2) Infer the stack and architecture hypotheses.
3) Run small, evidence-gated probes to confirm.
4) Only search/read deeper for categories that now have supporting evidence.

This reduces wasted probing and keeps attachments minimal.

## Workflow

### 1) Build a repo fingerprint (anchors only)

Read the smallest set of files that often exist and immediately reveal stack/shape:

- `README*` (first ~200 lines)
- root manifests that exist (one or more): `package.json`, `pyproject.toml`, `go.mod`, `Cargo.toml`, `pom.xml`, `build.gradle`, `composer.json`, `Gemfile`
- `Dockerfile`, `docker-compose.*` (if present)
- `.github/workflows/*` (names + one representative workflow file)
- `Makefile` / `Taskfile.*` (if present)

Also capture a lightweight file index summary (without reading everything):

- top-level directories
- dominant file extensions
- any “obvious entrypoint” indicators from scripts/config (e.g., package.json scripts, go `cmd/`, rust `src/main.rs`)

### 2) Form hypotheses (stack + entrypoints)

From the fingerprint, infer:

- primary language/runtime
- framework candidates (e.g., Express/Nest/Next, FastAPI/Django, chi/gin, Rails, Spring)
- likely entrypoints and routing surfaces
- where “truth” lives (schema, API contracts, config)

Write down 3–6 candidate entrypoints (paths) you will validate.

### 3) Evidence-gated probing loop (bounded)

Run a small loop of probes; each probe must have a clear purpose and a stop condition.

**Budgets (defaults):**

- max probes: 12
- max file reads: 25
- max search passes: 6

**Allowed probes (choose the cheapest next step):**

- Open an inferred entrypoint and follow imports to one routing/handler layer.
- Read a config file referenced by an entrypoint.
- Perform a targeted search ONLY for confirmed-framework idioms (e.g., “router”, “app.get”, “FastAPI()”, “chi.NewRouter”) after framework is evidenced.
- If you see a strong signal (e.g., `prisma` in deps), then and only then probe its likely files (`prisma/schema.prisma`, migrations).

**Do NOT** scan for categories (auth/jobs/migrations/feature flags/observability) unless you have at least one of:

- a dependency/config signal (library name, config key)
- an import path signal
- a directory naming signal found naturally via the fingerprint or followed imports

### 4) Build the evidence index (mandatory)

Create an internal evidence index with **>= 20** references that actually exist, as one of:

- `path:symbol` (function/class/type/handler name)
- explicit endpoint strings (e.g., `/api/users`)
- job/event names (queue/topic/schedule identifiers)
- config keys/flags (feature gates, env var keys)

For each reference, also record a “best attachment set” (1–6 files) that makes it intelligible.

If the repo is too small to reach 20, pad with `Unknown` entries; those questions must explicitly call out missing evidence.

### 5) Draft exactly 20 strategist questions (12 immediate + 8 strategic)

Each question must:

- be feasibility-focused and grounded in repo reality
- respect `constraints`
- explicitly exclude `non_goals`
- be actionable for `team_size` and `deadline`
- cite at least one evidence reference (or `Unknown` with a clear gap)
- fall into one of these areas (only when evidenced; otherwise treat as a “risk/unknowns” question):
  - contracts/interfaces
  - invariants/data consistency
  - auth/permissions
  - migrations/schema
  - request/UX flows (if applicable)
  - failure modes & reliability
  - feature flags/config gating
  - observability (logs/metrics/tracing)

Guidance:

- Immediate (01–12): unblockers, correctness, runtime risk, deployment friction, quick wins.
- Strategic (13–20): architectural constraints, roadmap leverage, risk retirement, scalability, maintainability.

### 6) Emit the question pack as bash commands (final output)

#### Common command shape

Each command MUST:

- call Oracle with `-p/--prompt` and `-f/--file` (repeat `-f` as needed)
- include `--files-report`
- include `--write-output "<out_dir>/<nn>-<slug>.md"`
- include `--wait` unless user opted out
- attach minimal files:
  - small base context: 1–3 anchor files (entrypoint + config)
  - focus context: files containing the cited evidence reference(s)

Avoid broad `**/*` globs. Prefer explicit files and small globs with negation excludes.

#### Prompt micro-format (required in every `-p`)

Each `-p` must require Oracle to respond using exactly:

- **Answer**
- **Evidence** (bullets; cite concrete refs)
- **Risks / unknowns**
- **Smallest experiment** (runnable today)

Keep prompts short and specific.

#### Optional preflight (recommended)

For each command, optionally emit a commented “dry-run” line immediately above it:

- same command + `--dry-run summary`
This lets the user validate attachments/token weights without sending.

## Failure handling rules

- Missing evidence for a category:
  - do NOT hunt for it
  - convert into a “risk/unknowns” question
  - attach best available anchors + nearest relevant files
- Very large repos:
  - tighten budgets
  - reduce base context to 1–2 files
  - prefer following import chains over directory scans
- Unclear entrypoint:
  - use build scripts/config to infer; otherwise choose the most central “composition root” file you can find and proceed

## Write output to docs (required)

When running this skill, also write the complete deliverable (the single `bash` fence) to:

- `docs/oracle-strategist-commands-YYYY-MM-DD.md` (ISO date)
Fallback:
- `docs/oracle-strategist-commands.md`

In the assistant response:

1) Print: `Output file: <path>`
2) Then print the exact same `bash` fenced deliverable.
