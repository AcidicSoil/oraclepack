Title:

* Add deterministic chaining + structured outputs to oraclepack (fix prelude semantics and enable stage-3 “actions” workflow)

Summary:

* The current oraclepack workflow is a 2-stage pipeline (pack generation → oraclepack execution) but has a “disconnect” after execution: the 20 Markdown outputs are not machine-consumable for automated follow-on work, blocking a seamless next stage for actionable implementation.

    Workflow Improvement Suggestions

* Additionally, the runner executes the pack prelude and each step in separate `bash -lc` processes, so prelude-defined variables do not persist into steps, creating a mismatch between template expectations and runtime behavior.

    Workflow Improvement Suggestions

Background / Context:

* Stage 1: a Codex “skill” or Gemini CLI “command” generates a single Markdown “oracle-pack” document whose machine-critical content is a single fenced `bash` block containing 20 numbered steps that call `oracle` with `--write-output ...`.

    Workflow Improvement Suggestions

* Stage 2: oraclepack parses that Markdown by extracting the first \`\`\`bash fence and splitting steps by a step-header pattern, then runs a “prelude” and each step in separate shells, producing 20 output files.

    Workflow Improvement Suggestions

* Goal: improve workflow productivity and longer runtimes with minimal human interaction, including automatically passing the final 20 results into a next stage that yields actionable implementation steps.

    Workflow Improvement Suggestions

Current Behavior (Actual):

* Prelude variables do not persist into step execution because prelude and steps run in separate `bash -lc` processes; packs must use explicit paths instead of relying on prelude variables like `$out_dir`.

    Workflow Improvement Suggestions

* The 20 `oracle` outputs are human-readable Markdown but lack a machine-friendly handoff artifact (e.g., JSON/YAML), making automated downstream processing brittle.

    Workflow Improvement Suggestions

* Parser/run is vulnerable to “format drift” from pack generators (extra code fences, missing/incorrect headers, multiple `bash` fences—parser grabs the first).

    Workflow Improvement Suggestions

Expected Behavior:

* Pack prelude semantics should match user/template expectations (either reliably shared across steps or explicitly non-shared with enforced guidance).

    Workflow Improvement Suggestions

* After a run, oraclepack should produce a deterministic, machine-readable “handoff” that enables an immediate next stage without manual intervention (“without missing a beat”).

    Workflow Improvement Suggestions

* Pack ingestion should be resilient and self-validating (fail fast on drift and contract violations).

    Workflow Improvement Suggestions

Requirements:

* Prelude semantics (choose and make official):

    Workflow Improvement Suggestions

  * Option A: “Prelude is prep-only” (no shared vars): update pack template guidance accordingly.

  * Option B: “Prelude is sourced into every step”: implement by prefixing each step script with prelude content (or run all steps in a single long-lived shell session).

* Stage-1 → Stage-2 contract hardening (“self-healing”):

    Workflow Improvement Suggestions

  * Standardize step headers to the conservative form `# NN)`.

  * Enforce exactly one fenced `bash` block per pack (or validation error).

  * Run `oraclepack validate` immediately after pack generation (wrapper/scriptable convention).

* Add a machine-readable “run index” artifact per run:

    Workflow Improvement Suggestions

  * Must include per-step `--write-output` paths, exit codes, timestamps.

  * Include parsed metadata when available (ROI/category/reference from step header line).

* Add a first-class “chain” capability to generate an “actions” next stage:

    Workflow Improvement Suggestions

  * Proposed interface: `oraclepack chain <pack.md> --mode actions`.

  * Must synthesize: `oracle-out/_summary.md` (human) and `oracle-out/_actions.json` (machine).

  * `_actions.json` should normalize each item with at least: `category`, `roi`, `reference`, `recommended_action` (“Next smallest concrete experiment”), `missing_artifacts[]`, `risk_notes[]`.

        Workflow Improvement Suggestions

  * Emit a follow-on pack: `docs/oracle-actions-YYYY-MM-DD.md`.

        Workflow Improvement Suggestions

* Execution/runtime considerations:

  * Keep compatibility with non-interactive operation (`--no-tui`, `--run-all`) and stop-on-fail behavior so chaining can run in CI.

        Workflow Improvement Suggestions

  * Optional: opt-in parallel execution with a small concurrency cap if provider/rate limits allow.

        Workflow Improvement Suggestions

* Improve pack input/context usage:

  * Support “focus + exclusion” inputs (e.g., `focus_categories=permissions,migrations`, `exclude_paths=docs,node_modules,dist`) without changing pack shape.

        Workflow Improvement Suggestions

  * Treat “extra\_files” as a deliberate context bundle (org standards/ADRs/threat model/coding conventions) appended to commands.

        Workflow Improvement Suggestions

Out of Scope:

* Not provided

Reproduction Steps:

1. Generate an oracle-pack whose prelude defines `out_dir="..."` and steps reference `$out_dir/...`.

2. Run oraclepack on the pack.

3. Observe that step shells do not see prelude-defined variables (because each step runs in a separate `bash -lc`), requiring explicit paths instead.

    Workflow Improvement Suggestions

Environment:

* OS: Unknown

* oraclepack: Unknown version (Go wrapper around `oracle`)

    Workflow Improvement Suggestions

* Shell execution model: `bash -lc` per step (per assistant analysis)

    Workflow Improvement Suggestions

* Pack generators: Codex skill and Gemini CLI command workflows

    Workflow Improvement Suggestions

Evidence:

* “oraclepack executes **prelude and steps in separate `bash -lc` processes**, so variables defined in the prelude do **not** persist into the step shells.”

    Workflow Improvement Suggestions

* Format drift risks listed: extra code fences, missing/incorrect step headers, multiple `bash` fences (parser grabs the first).

    Workflow Improvement Suggestions

* Proposed chain command + structured artifacts: `_summary.md`, `_actions.json`, and `docs/oracle-actions-YYYY-MM-DD.md`.

    Workflow Improvement Suggestions

Decisions / Agreements:

* None explicitly finalized; two alternative fixes for prelude semantics were presented (runner fix vs template fix), but no selection recorded.

    Workflow Improvement Suggestions

Open Items / Unknowns:

* Which prelude semantic is intended as the official contract (“prep-only” vs “sourced into every step”).

    Workflow Improvement Suggestions

* Exact current parser rules (regex specifics), validation behavior, and current report/state outputs (whether a run index already exists).

    Workflow Improvement Suggestions

* Whether Task Master integration is desired as a required dependency or just an optional downstream consumer.

    Workflow Improvement Suggestions

Risks / Dependencies:

* Dependency on consistent pack formatting across multiple generators (Codex/Gemini); drift can break parsing/validation.

    Workflow Improvement Suggestions

* If parallelism is added, provider rate limits and error handling may complicate deterministic runs.

    Workflow Improvement Suggestions

* Downstream automation quality depends on having a structured index/JSON handoff rather than parsing free-form Markdown answers.

    Workflow Improvement Suggestions

Acceptance Criteria:

* Running a pack that defines variables in the prelude and references them in steps behaves according to the selected official contract:

  * If “sourced prelude”: steps can reference prelude-defined variables successfully.

  * If “prep-only”: validation or guidance prevents packs from relying on prelude vars (or template guidance is updated and enforced).

        Workflow Improvement Suggestions

* `oraclepack validate` reliably fails on packs with:

  * multiple `bash` fences,

  * missing/incorrect `# NN)` step headers (per enforced convention),

  * other contract-breaking drift conditions.

        Workflow Improvement Suggestions

* After a run, oraclepack emits a machine-readable run index that includes per-step output paths, exit codes, and timestamps.

    Workflow Improvement Suggestions

* `oraclepack chain <pack.md> --mode actions` produces:

  * `oracle-out/_summary.md`,

  * `oracle-out/_actions.json` with the specified normalized fields,

  * `docs/oracle-actions-YYYY-MM-DD.md` suitable for immediate stage-2 execution.

        Workflow Improvement Suggestions

* Chaining works in non-interactive mode and respects stop-on-fail behavior to support CI usage.

    Workflow Improvement Suggestions

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* enhancement

* workflow

* automation

* cli

* parsing

* schema

* ci

---
