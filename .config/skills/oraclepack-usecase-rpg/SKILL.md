---
name: oraclepack-usecase-rpg
description: Generate a strict, oraclepack-ingestible 20-step Oracle strategist pack (single ```bash fence) with ROI-ranked step headers and RPG (Repository Planning Graph) metadata embedded in each -p prompt body.
metadata:
  short-description: Strict oraclepack pack generator (RPG-infused)
---

## Quick start

Use this skill when the user asks for any of the following:
- an “oraclepack pack”, “oracle strategist question pack”, or “20-step ROI-ranked pack”
- “oraclepack use cases” that must keep the strict runner-ingestible output
- the same, but with RPG/RPG graph/Repository Planning Graph infusion

What you produce:
- Exactly one Markdown deliverable (intended path: `docs/oracle-pack-YYYY-MM-DD.md`)
- The deliverable contains **exactly one** fenced code block that starts with **exactly** ` ```bash` then a newline (no extra tokens after `bash`)
- The fenced `bash` block contains **exactly 20** sequential steps, `01..20`

Reference template: `references/oracle-pack-template.md`

Optional validator:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

## Workflow

### 1) Parse inputs (free-text args)

Interpret the user’s trailing text as conceptual `{{args}}` and extract:
- `codebase_name` (default: `Unknown`)
- `constraints` (default: `None`)
- `non_goals` (default: `None`)
- `team_size` (default: `Unknown`)
- `deadline` (default: `Unknown`)
- `out_dir` (default: `oracle-out`)
- `oracle_cmd` (default: `oracle`)
- `oracle_flags` (default: `--files-report`)
- `extra_files` (default: empty; append **literally** to every command if provided)

If any value is missing or ambiguous, use the defaults above (do not ask questions).

### 2) Repo discovery (inference-first, deterministic)

Read small, high-signal “index” artifacts first (if they exist):
- README / overview docs
- manifests / dependency lockfiles
- entrypoints (main binaries, CLI, server startup)
- config examples
- CI configuration

Then do targeted discovery (only as needed):
- Prefer deterministic searches (e.g., `ck`, `ast-grep`, `fd`) if available.
- Cap results deterministically (fixed max hits).
- If a file/path does not exist: record `Unknown` and continue.

### 3) Plan the 20 probes (what to cover)

You must keep categories **exactly** to this set (do not add/remove/rename):
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

Across the 20 steps, ensure at least one step explicitly targets each sub-layer below (embed it in the *Question* text; map to the best required category above):
- supply chain/dependencies
- secrets/credentials
- CI/CD or release process
- runtime configuration/IaC (if present)
- privacy/PII/retention (if applicable)
- SLO/alerting/runbooks
- cost/performance hotspots
- testing strategy/coverage

Horizon mix requirement:
- Exactly **12** steps with `horizon=Immediate`
- Exactly **8** steps with `horizon=Strategic`

### 4) RPG infusion (non-breaking)

Inside each step’s `-p` prompt body, include an `RPG:` block with:
- `FunctionalNode: <Capability>::<Feature> (WHAT)`
- `StructuralNode: <module-or-dir> :: <public surface> (HOW)`
- `DependsOn: [<prior step ids, optional>]`
- `Phase: <P0|P1|P2|P3>`

Dependency rule (to preserve ROI-sorted executability):
- `DependsOn` may reference **only earlier step IDs**.

Phase mapping guidance (do not create new categories):
- `P0`: contracts/interfaces, invariants, permissions, observability
- `P1`: caching/state, background jobs, failure modes
- `P2`: migrations, feature flags
- `P3`: UX flows

### 5) ROI scoring and ordering

For each step choose:
- `impact`, `confidence`, `effort` as **one decimal** in `0.1..1.0`

Compute:
- `ROI = (impact * confidence) / effort`

Ordering requirement:
- Sort all 20 steps by ROI descending
- Tie-break by lower effort
- After sorting, assign IDs `01..20` in that sorted order

If ordering conflicts with RPG dependencies:
- Adjust impact/confidence/effort (and thus ROI) so prerequisite/foundation steps float earlier.
- Keep `DependsOn` pointing backwards only.

### 6) Render the pack (strict output)

You must output exactly one Markdown deliverable matching the structure in `references/oracle-pack-template.md`:
- Title section
- Parsed args list
- Commands section with **exactly one** fenced code block:
  - opening fence must be exactly: <code>```bash</code>
  - closing fence must be exactly: <code>```</code>
- Coverage check section

Inside the fenced `bash` block:
- Include a small prelude setting `out_dir="..."` (so pack metadata can be derived).
- Each of the 20 steps must start with a header line matching one of:
  - `# 01) ...`
  - `# 01 — ...`
  - `# 01 - ...`

Each step header line must include all tokens:
- `ROI=<number>`
- `impact=<i>`
- `confidence=<c>`
- `effort=<e>`
- `horizon=<Immediate|Strategic>`
- `category=<one of required categories>`
- `reference=<...>` (use `Unknown` if needed)

Each step must invoke the Oracle CLI once and must include:
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- A `-p` prompt body that includes (in this exact section order):
  - `Strategist question #NN`
  - `RPG:` block (fields listed above)
  - `Reference: ...`
  - `Category: ...`
  - `Horizon: ...`
  - `ROI: ... (impact=..., confidence=..., effort=...)`
  - `Question: ...`
  - `Rationale: ...` (exactly one sentence)
  - `Smallest experiment today: ...` (exactly one action)
  - `Constraints: ...`
  - `Non-goals: ...`
  - `Answer format:` with the 4 required numbered items

Quoting guidance for `-p` (recommended, robust):
- Use a heredoc in command substitution so the prompt stays readable while remaining a single oracle invocation:
  - `-p "$(cat <<'PROMPT' ... PROMPT)"`

### 7) Validate before finalizing

Run the included validator if possible:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

Fix any failures by editing the pack until it passes.

## Output contract

When invoked, produce:

1) A single line:
- `Output file: docs/oracle-pack-YYYY-MM-DD.md`

2) Immediately after, the **exact** Markdown content for that file (no extra commentary).

If the current date is not available in the execution context:
- Use `docs/oracle-pack.md`.

## Failure modes (do not guess)

- Missing repo evidence → use `Reference: Unknown` and in the prompt’s item (4) name the exact missing file/path patterns to attach next.
- Missing args → apply defaults; do not ask follow-ups.
- Conflicting constraints (ordering vs dependencies vs horizon mix) → preserve strict ingestion first:
  1) single `bash` fence
  2) 20 sequential steps
  3) required header tokens
  4) horizon mix 12/8
  Then adjust ROI inputs and dependency edges.

## Invocation examples

1) “Generate an RPG-infused oraclepack strategist pack for this repo; focus on permissions boundaries and CI/CD; keep strict oraclepack output.”

2) “Create a 20-step ROI-ranked oraclepack pack for `promptbench` with constraints=None and non_goals=None; default out_dir; strict format.”

3) “Produce an oraclepack pack emphasizing caching/state correctness and observability; include at least one step about cost/performance hotspots; strict template; RPG nodes.”

4) “Generate an oraclepack use-case pack for a Go CLI; include secrets/credentials handling and release process; keep categories fixed; strict single ```bash fence.”

5) “Make an oraclepack pack with explicit dependency edges (DependsOn) that only point backwards; preserve ROI-sorted order; strict output.”
