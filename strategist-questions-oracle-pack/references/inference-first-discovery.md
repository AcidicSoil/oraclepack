<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/inference-first-discovery.md -->
# Inference-first discovery (adaptive search ladder, ck-first)

Goal: avoid wasting time on hard-targeted globs/greps that may not exist by deriving the search plan from repo evidence,
using the codebase-search rules (ck → ast-grep when structure matters → deterministic output shaping).

## Step 0 — Discover the search interface (required)

- Always run: `ck --help`
- Use only flags confirmed by `ck --help`.
- If `ck` indicates indexing/setup is needed, follow its printed instructions.

## Step 1 — Read “index” artifacts first (cheap, high-signal)

Use `fd` to locate these quickly:

- Repo docs/index:
  - `fd -p README.md`
  - `fd -e md . docs | head`
- Manifests/config (choose the ones that exist):
  - `fd -p package.json`
  - `fd -p pyproject.toml`
  - `fd -p go.mod`
  - `fd -p Cargo.toml`
  - `fd -p pom.xml`
  - `fd -p build.gradle`

Read the smallest set of files that explain structure and entrypoints.

## Step 2 — Infer subsystem locations from what you actually saw

Build a short “signals” list from evidence:
- Router/framework name (from dependencies and scripts)
- Job system
- Migration tool
- Feature flag system
- Observability stack
- Cache layer

Then: follow imports/registration code rather than searching the whole repo.

## Step 3 — Derive targeted `ck` queries (conceptual-first)

Start conceptual (meaning-based), scoped to inferred roots:

- Conceptual: `ck --sem "<intent phrase>"`
  - Examples of intent phrases (adapt to the repo’s vocabulary):
    - “route registration” / “router setup”
    - “auth middleware” / “permission check”
    - “queue worker registration”
    - “migration runner” / “schema change”
    - “feature flag evaluation”
    - “telemetry initialization” / “logger setup”

When unsure, use hybrid:
- `ck --hybrid "<intent phrase>"`

For literal hits (exact tokens you already saw), use regex:
- `ck --regex "<literal-or-regex>"`

### Deterministic narrowing

If results are broad/noisy, narrow deterministically:

- Use `ck --jsonl` (if available) and cap results:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- Otherwise cap with `head` and re-run `ck` constrained to the top-hit directories/files.

## Step 4 — Structural search with `ast-grep` (when structure matters)

Use `ast-grep` when you need syntax-aware matching (AST-level), such as:
- finding a specific call pattern (e.g., permission checks around handlers)
- finding registration shapes (e.g., route definitions)
- reducing false positives vs text search

Examples:

- Find code structure:
  - `ast-grep --lang <language> -p '<pattern>'`
- List matching files (cap output):
  - `ast-grep -l --lang <language> -p '<pattern>' | head -n 10`

Use `ck` first to locate candidate files/modules, then `ast-grep` to enforce structure inside those.

## Step 5 — Progressive widening (fallback ladder)

If inference cannot locate a subsystem:

1) Run a small set of conceptual `ck --sem` / `ck --hybrid` queries using generic intent phrases.
2) If still nothing, broaden to repo-wide `ck --hybrid` for that intent.
3) If still nothing, mark evidence for that category as missing and record the most likely missing artifact pattern.

## Step 6 — Early stopping (don’t over-search)

Stop harvesting once:
- you have >= 20 candidate anchors total, and
- every required category has at least one candidate (or a clearly documented “missing artifact pattern”)

Then generate questions; do not keep scanning “just in case”.
