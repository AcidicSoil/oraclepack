Title:

* Implement deterministic oraclepack pipeline improvements (strict validation, run manifests, resume/caching, Stage 3 “Actionizer”)

Summary:

* The current two-stage oraclepack workflow (Stage 1 pack generation → Stage 2 execution) is “weakly connected” and lacks deterministic handoff metadata, robust resume/retry, and an automated Stage 3 that converts 20 outputs into actionable engineering work.

    Oracle Pack Workflow Analysis

* This ticket proposes additive, backward-compatible enhancements to oraclepack and the Stage 1 generator prompts so runs are reproducible, CI-friendly, and produce machine-readable artifacts suitable for automation.

    Oracle Pack Workflow Analysis

Background / Context:

* Workflow context:

  * Stage 1: Codex skill or Gemini CLI slash command generates a single Markdown oracle question pack under `docs/*oracle-pack*.md`, following a strict oraclepack schema and containing exactly 20 `oracle ...` commands.

        Oracle Pack Workflow Analysis

  * Stage 2: oraclepack (Go wrapper around `@steipete/oracle`) executes the 20 commands and writes per-question outputs (via `--write-output`).

        Oracle Pack Workflow Analysis

  * Stage 3 is currently missing: outputs are not automatically turned into actionable implementation work.

        Oracle Pack Workflow Analysis

* Non-negotiable constraints:

  * No schema-breaking changes to the oraclepack Markdown pack schema without a backward-compatible migration path and validator-safe proof.

  * Automation must be deterministic and reproducible (no interactive steps in the critical path).

  * Stage 1 output must remain a single-pack deliverable that oraclepack can validate/run (no extra blocks/headers; no schema drift).

  * Prefer minimal file attachments per question; avoid broad globs unless unavoidable.

  * Optimize for longer runtimes with minimal human interaction (batching, resume/retry, caching, stable outputs, CI-friendly).

        Oracle Pack Workflow Analysis

Current Behavior (Actual):

* Stage 1 (generation) failure modes / friction points:

  * Packs can drift from the strict schema (extra fenced blocks, step-like headers, missing fields, wrong ordering, wrong count ≠ 20), causing ingestion/validation issues.

        Oracle Pack Workflow Analysis

  * Attachments may be bloated (broad globs, “just in case” files), increasing token cost and reducing signal-to-noise.

        Oracle Pack Workflow Analysis

  * ROI scoring can be inconsistent (unstable prioritization vs stated rationale).

        Oracle Pack Workflow Analysis

  * Coverage duplication across 20 questions (overlapping targets) wastes runs/budget.

        Oracle Pack Workflow Analysis

* Stage 2 (execution) failure modes / friction points:

  * Resume/retry semantics are weak (reruns may re-execute completed steps; partial failures require manual selection).

        Oracle Pack Workflow Analysis

  * Output determinism gaps: inconsistent output paths/slugs/out\_dir naming undermine CI diffs and Stage 3 discovery.

        Oracle Pack Workflow Analysis

  * Concurrency/rate limiting is not first-class (provider throttling/timeouts lead to nondeterministic failures).

        Oracle Pack Workflow Analysis

* Cross-stage handoff issues:

  * Missing traceability between pack ↔ outputs (no explicit manifest tying outputs to pack hash, git SHA, tool versions, provider/model settings).

        Oracle Pack Workflow Analysis

  * Stage 2 may be bypassed (pack executed “by hand”), losing wrapper state/report and consistent run directory.

        Oracle Pack Workflow Analysis

Expected Behavior:

* Stage 1 packs are always validator-safe and deterministic (single pack, exactly 20 oracle invocations, no schema drift).

    Oracle Pack Workflow Analysis

* Stage 2 produces stable, discoverable, machine-readable run artifacts that bind pack ↔ outputs and enable idempotent resume/rerun.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) exists and deterministically converts the 20 outputs into actionable engineering work artifacts (backlog + change plan + optional issue export), without duplicating work on reruns.

    Oracle Pack Workflow Analysis

* CI can run validate → run → actionize non-interactively with structured outputs and policy-driven exit codes.

    Oracle Pack Workflow Analysis

Requirements:

* Validation / linting (additive, backward-compatible):

  * Add `oraclepack validate --strict --json` that reports counts (steps=20, bash\_blocks=1, oracle\_invocations=20), ordering checks (ROI desc; ties effort asc), and required fields presence.

        Oracle Pack Workflow Analysis

* Deterministic run directory + manifest:

  * On `run`, create `.oraclepack/runs/<pack_id>/` and emit `run.json` + `steps.json`.

        Oracle Pack Workflow Analysis

  * `pack_id = YYYY-MM-DD__<gitshort>__<packhash8>`.

        Oracle Pack Workflow Analysis

  * `run.json` must include: `pack_path`, `pack_hash`, `git_sha`, `oraclepack_version`, `oracle_version`, `created_at`.

        Oracle Pack Workflow Analysis

  * `steps.json` must include: `step_id` (01..20), `reference`, `category`, `roi`, `command_hash`, `output_path`, `output_hash`, `status` (pending|ok|failed|skipped).

        Oracle Pack Workflow Analysis

* Resume/rerun semantics:

  * Make resume default: if `run.json` exists, skip steps whose output exists and matches recorded hash.

  * Support explicit overrides: `--rerun all|failed|01,03,07`.

        Oracle Pack Workflow Analysis

* Concurrency and rate limiting:

  * Add global `--max-parallel N` and optionally per-provider caps via config.

  * Implement exponential backoff + jitter on transient errors (e.g., 429/503) with a retry budget.

        Oracle Pack Workflow Analysis

* Deterministic caching (optional initially):

  * Implement caching keyed by `sha256(prompt + attached_file_hashes + oracle_flags + model)`, stored under `.oraclepack/cache/<invocation_key>.md`; rerun reuses cached outputs when key matches.

        Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) design and artifacts:

  * Implement `oraclepack actionize --run-dir .oraclepack/runs/<pack_id>`.

  * Inputs: `run.json` + 20 outputs under `.oraclepack/runs/<pack_id>/outputs/`.

        Oracle Pack Workflow Analysis

  * Deterministic processing: normalize outputs → dedupe → categorize via fixed taxonomy → generate action tasks, including blocked/conflict handling.

        Oracle Pack Workflow Analysis

  * Outputs under `.oraclepack/runs/<pack_id>/actionizer/`:

    * `normalized.jsonl`, `backlog.md`, `change-plan.md`

    * Optional: `github-issues.json`, `taskmaster.json`.

            Oracle Pack Workflow Analysis

  * Idempotency: stable IDs derived from `pack_hash` (e.g., `OP3-<packhash8>-<issue_index>-<task_index>`), stable paths, byte-identical regeneration when inputs unchanged.

        Oracle Pack Workflow Analysis

* Stage 1 prompt/skill improvements (without breaking schema):

  * Embed structured mini-metadata inside each `-p` prompt text (not new pack headers), e.g., `QuestionId`, `Category`, `Reference`, `ExpectedArtifacts`.

        Oracle Pack Workflow Analysis

  * Enforce deterministic attachment minimization heuristics (reference file + 0–2 neighbors; avoid broad globs unless evidence demands).

        Oracle Pack Workflow Analysis

  * Standardize generator prompt across Codex skills and Gemini CLI commands using a single canonical prompt file in repo.

        Oracle Pack Workflow Analysis

* CI-native mode:

  * Provide `oraclepack run --ci --non-interactive --json-log` and `oraclepack actionize --ci`.

  * CI policy can fail build if validation fails, completion rate below threshold, or action yield below threshold (threshold values: Not provided).

        Oracle Pack Workflow Analysis

* Security/safety:

  * Path safety: prevent `--write-output` from escaping run dir (reject `..` traversal).

        Oracle Pack Workflow Analysis

Out of Scope:

* Breaking changes to the existing oraclepack Markdown pack schema (unless a backward-compatible migration path and validator-safe proof are provided).

    Oracle Pack Workflow Analysis

Reproduction Steps:

1. Generate a pack via Stage 1 and save to `docs/oracle-pack-YYYY-MM-DD.md`.

    Oracle Pack Workflow Analysis

2. Run `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md` and observe schema drift / strictness gaps (exact current validator behavior: Unknown).

    Oracle Pack Workflow Analysis

3. Execute the pack, interrupt mid-run, rerun, and observe whether completed steps are skipped (current behavior: weak/unclear).

    Oracle Pack Workflow Analysis

4. Compare two runs on the same commit and observe output path/slug stability and traceability artifacts (manifest missing today).

    Oracle Pack Workflow Analysis

Environment:

* Tooling:

  * oraclepack (Go wrapper around `@steipete/oracle`).

        Oracle Pack Workflow Analysis

  * Stage 1 generators: Codex skills or Gemini CLI slash commands.

        Oracle Pack Workflow Analysis

* Repository/OS/versions: Unknown (git SHA, oraclepack version, oracle version, provider/model settings not provided in the conversation; also identified as missing traceability today).

    Oracle Pack Workflow Analysis

Evidence:

* Proposed stable artifact layout and handoff contract:

    Oracle Pack Workflow Analysis

  * `docs/oracle-pack-YYYY-MM-DD.md`

  * `.oraclepack/runs/<pack_id>/run.json`

  * `.oraclepack/runs/<pack_id>/steps.json`

  * `.oraclepack/runs/<pack_id>/outputs/01-<slug>.md … 20-<slug>.md`

  * `.oraclepack/runs/<pack_id>/actionizer/{normalized.jsonl, backlog.md, change-plan.md}`

* Proposed commands (some flag names explicitly “proposed where not already present”):

    Oracle Pack Workflow Analysis

  * `oraclepack validate --strict docs/oracle-pack-YYYY-MM-DD.md --json > .oraclepack/validate.json`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --max-parallel 4 --resume --ci`

  * `oraclepack actionize --run-dir .oraclepack/runs/<pack_id> --ci`

* Example Stage 3 output record shape (JSONL line):

    Oracle Pack Workflow Analysis

Decisions / Agreements:

* Do not break the oraclepack Markdown pack schema; any change must be backward-compatible with a validator-safe proof.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) is required and should be implemented as a first-class oraclepack subcommand (`actionize`) producing deterministic artifacts with idempotent reruns.

    Oracle Pack Workflow Analysis

* Traceability and determinism should be achieved via additive sidecar files (e.g., `run.json`, `steps.json`) rather than pack schema changes.

    Oracle Pack Workflow Analysis

Open Items / Unknowns:

* Current oraclepack CLI surface area:

  * Whether `validate --strict`, `--json`, `run --ci`, `--resume`, `--json-log`, and `actionize` exist today vs need implementation (conversation notes some flags are “proposed”).

        Oracle Pack Workflow Analysis

* Current on-disk state/report formats and locations (“state lives today (intended): oraclepack state/report + per-step outputs”; exact paths not provided).

    Oracle Pack Workflow Analysis

* Threshold definitions for CI policy (“completion rate < threshold”, “action yield < threshold”): Not provided.

    Oracle Pack Workflow Analysis

* Exact strict pack schema invariants enforced today (beyond “strict output contract” and “exactly 20” requirement): Not provided in this conversation (referenced as external inputs).

    Oracle Pack Workflow Analysis

Risks / Dependencies:

* Risk: filesystem layout changes may affect existing users; mitigation is additive behavior that preserves current out\_dir behavior.

    Oracle Pack Workflow Analysis

* Risk: caching correctness depends on hashing all attached file contents; incomplete hashing risks “cache poisoning.”

    Oracle Pack Workflow Analysis

* Risk: provider throttling/timeouts require robust transient-error classification for backoff/retry behavior.

    Oracle Pack Workflow Analysis

* Dependency: Stage 3 quality depends on stable, parseable structure in per-question outputs; mitigated by deterministic normalization heuristics and improved Stage 1 prompt shaping.

Acceptance Criteria:

* Validation:

  * `oraclepack validate --strict --json` deterministically reports schema invariants (20 steps, 20 oracle invocations, schema drift detection) and can gate CI.

        Oracle Pack Workflow Analysis

* Run determinism and traceability:

  * Running a pack produces `.oraclepack/runs/<pack_id>/{run.json,steps.json,outputs/}` with stable `pack_id` and stable output naming.

  * `run.json` includes required metadata fields; `steps.json` includes required per-step fields and statuses.

        Oracle Pack Workflow Analysis

* Resume/rerun:

  * Interrupting a run mid-way and rerunning resumes without re-executing completed steps (validated via output hashes and statuses).

  * `--rerun failed|all|<step list>` behaves as specified.

        Oracle Pack Workflow Analysis

* Concurrency/rate limiting:

  * `--max-parallel N` bounds concurrency; transient failures (e.g., throttling/timeouts) are retried with backoff within a retry budget and recorded in step status.

        Oracle Pack Workflow Analysis

* Caching (if implemented):

  * Rerunning on unchanged inputs (same prompt, same attached file digests, same flags/model) results in zero provider calls and identical outputs.

        Oracle Pack Workflow Analysis

* Stage 3 “Actionizer”:

  * `oraclepack actionize --run-dir ...` generates deterministic artifacts under `actionizer/` (`normalized.jsonl`, `backlog.md`, `change-plan.md`).

  * Reruns do not duplicate tasks (stable IDs) and produce byte-identical output when inputs unchanged.

  * Missing/contradictory answers produce explicit `blocked`/`conflict` tasks with required evidence patterns.

* CI mode:

  * `--ci --non-interactive --json-log` runs without TUI/interaction and uses structured logs and policy-driven exit codes.

        Oracle Pack Workflow Analysis

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* enhancement, cli, determinism, validation, resume, caching, concurrency, workflow, stage3-actionizer

---
