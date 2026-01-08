<filetree>
Project Structure:
└── .tickets
    ├── mcp
    │   ├── Expose Oraclepack as MCP.md
    │   ├── MCP Server for Oraclepack.md
    │   ├── oraclepack-MCP.md
    │   └── oraclepack_mcp_server.md
    ├── Enable Action Packs Dispatch.md
    ├── Improving Oraclepack Workflow.md
    ├── Oraclepack Action Pack Integration.md
    ├── Oraclepack Action Pack Issue.md
    ├── Oraclepack Action Packs.md
    ├── Oraclepack Compatibility Issues.md
    ├── Oraclepack Pipeline Improvements.md
    ├── Oraclepack Prompt Generator.md
    ├── Oraclepack TUI Integration.md
    ├── Oraclepack Workflow Enhancement.md
    └── Verbose Payload Rendering TUI.md

</filetree>

<source_code>
.tickets/Enable Action Packs Dispatch.md
```
Title:

* Enable oraclepack Action Packs to dispatch to non-oracle executors (codex/gemini/task-master/tm)

Summary:

* Current oraclepack usage feels “oracle-only” because certain UX features (flag injection and overrides validation) are hard-coded to `oracle` invocations, while the user needs Action Packs that actually execute work via other agents/tools (e.g., `codex exec`, `gemini`, `task-master`, `tm`). Update the Stage-3 Action Pack generation and/or oraclepack runner logic so Action Packs can deterministically run the correct executor commands for each action item.

Background / Context:

* User concern: “oraclepack is a wrapper around `oracle`” and adding more `oracle` calls won’t implement tasks; Action Packs must run the real tools/agents used in their workflow (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Current behavior explanation: oraclepack executes shell steps but has special handling only for lines starting with `oracle` (detection regex, flag injection, and overrides validation via `oracle --dry-run summary`).
* Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) and performs guarded checks for autopilot.
* Referenced repos/tools: Gemini CLI, Claude Task Master, OpenAI Codex, steipete/oracle.
* Referenced code/assets: `oraclepack-tui.md`, `oracle_pack_and_taskify-skills.md` (not included in transcript).

Current Behavior (Actual):

* Oraclepack “nice UX features” are oracle-specific:

  * Detects invocations using a regex anchored to literal `oracle`.
  * Injects selected flags only into `oracle …` lines.
  * Validates overrides by running `oracle --dry-run summary` only for detected oracle invocations.
* Action Packs can run arbitrary shell commands, but oraclepack’s overrides/validation UX does not generalize to other tools (codex/gemini/task-master/tm).

Expected Behavior:

* Generated Action Packs can deterministically dispatch execution to the intended executor per action item (e.g., `codex exec …`, `gemini -p …`, `task-master …`, `tm …`) rather than relying on more `oracle` calls.
* Optional: oraclepack’s overrides/flag-injection UX can recognize additional command prefixes beyond `oracle` (if desired).

Requirements:

* Update the Stage-3 generator (“oraclepack-taskify” / Stage-3 Action Pack template) to support configurable tool command strings beyond existing `oracle_cmd`, `task_master_cmd`, `tm_cmd`:

  * Add `codex_cmd` (default `codex`)
  * Add `gemini_cmd` (default `gemini`)
  * Optional: add `autopilot_cmd` (default `${tm_cmd} autopilot`).
* Extend `_actions.json` schema to include an executor plan:

  * `tooling`: include `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per action item: `executor` (`"codex" | "gemini" | "tm" | "manual"`), `exec_prompt` (deterministic; no code fences), optional `inputs` (paths/globs).
* Add a new Action Pack execution path for implementation:

  * Either add `mode=implement`, or add a new “Step 09” gated by `MODE == implement`.
  * Step reads `<out_dir>/_actions.json`, selects top N items (using existing `top_n`), then dispatches:

    * `codex exec …` when `executor == codex`
    * `gemini -p …` when `executor == gemini`.
* Safety constraint:

  * Keep defaults strict (avoid “yolo” execution by default); Gemini approval/tool execution should remain conservative unless explicitly opted in.
* Optional (nice-to-have): generalize oraclepack’s oracle-specific detection/injection:

  * Generalize `ExtractOracleInvocations` / `InjectFlags` to recognize a registry of prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`).
  * Add per-tool override sets (so “Oracle Flags” aren’t incorrectly applied to other tools).

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* OS: Unknown
* oraclepack version/commit: Unknown
* Shell/runner context: Unknown
* Stage-3 generator version/commit: Unknown

Evidence:

* User statement of need: actionpacks must call their agents/tools (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Oraclepack oracle-specific UX behavior (regex detection, flag injection, `oracle --dry-run summary` validation).
* Proposed schema + implement mode/Step 09 dispatcher design.
* Attachment: Oraclepack Action Pack Integration.md

Decisions / Agreements:

* “Fastest path” identified in transcript: upgrade Stage-3 Action Packs to include an executor dispatch step and extend `_actions.json` with `executor` metadata; modifying oraclepack core is optional.

Open Items / Unknowns:

* Current `_actions.json` schema (exact fields/types) is not provided.
* Current Stage-3 Action Pack template structure (exact steps and modes) is not provided beyond `backlog|pipelines|autopilot`.
* Exact locations/implementations of `ExtractOracleInvocations` / `InjectFlags` in `oraclepack-tui.md` are not provided in this transcript.
* Expected non-interactive invocation patterns/flags for each tool in this project (codex/gemini/task-master/tm) beyond the examples are not provided.

Risks / Dependencies:

* Dependency on external CLIs and their execution/approval modes (especially Gemini CLI) with safety implications; defaults must remain conservative.
* If oraclepack overrides/validation stays oracle-only, users may expect those UX features to apply to non-oracle commands; requires clear separation or per-tool override support.

Acceptance Criteria:

* `_actions.json` produced by the Stage-3 generator includes:

  * `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per-item `executor` and `exec_prompt` (and optional `inputs`).
* Action Pack supports `implement` execution (via `mode=implement` or Step 09) that:

  * Reads `<out_dir>/_actions.json`
  * Selects top N actions
  * Runs `codex exec …` for `executor=codex` and `gemini -p …` for `executor=gemini` deterministically.
* Defaults do not enable unrestricted/automatic tool execution without opt-in (conservative approval/safety posture).
* (Optional) oraclepack runner recognizes non-oracle prefixes for invocation detection and does not incorrectly apply oracle-specific overrides to other tools.

Priority & Severity (if inferable from text):

* Priority: Not provided
* Severity: Not provided

Labels (optional):

* enhancement
* action-pack
* executor-dispatch
* cli-integration
* oraclepack
* taskify
```

.tickets/Improving Oraclepack Workflow.md
```
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

    *   If “sourced prelude”: steps can reference prelude-defined variables successfully.

    *   If “prep-only”: validation or guidance prevents packs from relying on prelude vars (or template guidance is updated and enforced).

        Workflow Improvement Suggestions

* `oraclepack validate` reliably fails on packs with:

    *   multiple `bash` fences,

    *   missing/incorrect `# NN)` step headers (per enforced convention),

    *   other contract-breaking drift conditions.

        Workflow Improvement Suggestions

* After a run, oraclepack emits a machine-readable run index that includes per-step output paths, exit codes, and timestamps.

    Workflow Improvement Suggestions

* `oraclepack chain <pack.md> --mode actions` produces:

    *   `oracle-out/_summary.md`,

    *   `oracle-out/_actions.json` with the specified normalized fields,

    *   `docs/oracle-actions-YYYY-MM-DD.md` suitable for immediate stage-2 execution.

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
```

.tickets/Oraclepack Action Pack Integration.md
```
Parent Ticket:

* Title: Oraclepack Action Pack Integration: dispatch Action Packs to external agents/tools (Codex/Gemini/Task Master)
* Summary:

  * Current concern: oraclepack is perceived as “oracle-only,” and adding more `oracle` calls won’t implement tasks.
  * Desired outcome: Action Packs should run the correct agent/tool commands (e.g., `codex exec …`, `tm …`, `gemini …`) and be tool-agnostic in how they dispatch work.
  * Optional scope: extend oraclepack’s oracle-specific UX (flag injection / validation) to support non-`oracle` commands.
* Source:

  * Link/ID (if present) or “Not provided”: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make it so our actionpacks will perform the correct calls to the agents… Example ‘codex exec …’, ‘tm …’, ‘gemini …’”
* Global Constraints:

  * Not provided
* Global Environment:

  * Tools referenced as available/on PATH in discussion: `oracle`, `codex`, `gemini`, `task-master`, `tm`
  * Action Pack modes referenced: `backlog|pipelines|autopilot` (and proposed `implement`)
  * Runner behavior referenced: steps execute as `bash`; stdin/TTY not attached (impacting interactive CLIs)
* Global Evidence:

  * Referenced files: `oracle_pack_and_taskify-skills.md`, `oraclepack-tui.md`
  * Referenced tool repos: `https://github.com/google-gemini/gemini-cli`, `https://github.com/eyaltoledano/claude-task-master`, `https://github.com/openai/codex`, `https://github.com/steipete/oracle`
  * Mentioned artifacts/screens: “screenshot of oraclepack consuming oraclepack-taskify artifacts”; “oracle-actions-pack-2026-01-07.md”

Split Plan:

* Coverage Map:

  * “oraclepack is a wrapper around `oracle`… can not see how… more oracle calls will help us implement” → Info-only
  * “make it so our actionpacks will perform the correct calls… ‘codex exec …’, ‘tm …’, ‘gemini …’” → T3
  * “oraclepack… special logic only for lines that start with `oracle`… detect… inject flags… overrides validation” → T5
  * “Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`… Extend… `codex_cmd`, `gemini_cmd`, optionally `autopilot_cmd`” → T1
  * “Extend `_actions.json`… include… `tooling`… per item `executor`, `exec_prompt`, `inputs`” → T2
  * “Add an ‘implement’ mode (or a Step 09)… reads… `_actions.json`… selects top N… dispatches… `codex exec…` / `gemini -p…`” → T3
  * “Keep safety defaults strict (do not ‘yolo’ by default)” → T4
  * “interactive CLIs… may fail/hang because oraclepack doesn’t provide stdin/TTY” → T4
  * “This particular Action Pack does not call `codex`, `gemini`…” → Info-only
  * “Optional: improve oraclepack UX to ‘understand’ non-oracle commands… registry of prefixes… per-tool override sets” → T5
* Dependencies:

  * T3 depends on T2 because the proposed dispatcher reads `_actions.json` fields like `executor` / `exec_prompt`.
  * T2 depends on T1 because the proposed `_actions.json.tooling` includes `codex_cmd` / `gemini_cmd` command strings.
* Split Tickets:

```ticket T1
T# Title:
- Add Codex/Gemini command configuration to Stage-3 generator (alongside oracle/task-master/tm)

Type:
- enhancement

Target Area:
- Stage-3 generator / skill template that currently supports `oracle_cmd`, `task_master_cmd`, `tm_cmd`

Summary:
- Extend the Stage-3 generation inputs to support additional tool command strings so Action Packs can invoke Codex and Gemini explicitly.
- This is intended to mirror the existing configurable command pattern already used for `oracle` and Task Master tools.

In Scope:
- Add `codex_cmd` (default `codex`) to the generator inputs.
- Add `gemini_cmd` (default `gemini`) to the generator inputs.
- Add optional `autopilot_cmd` (default `${tm_cmd} autopilot`) to the generator inputs.
- Ensure generated Action Pack uses these command variables where relevant.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Generator supports configurable tool command strings: `oracle_cmd`, `task_master_cmd`, `tm_cmd`.

Expected Behavior:
- Generator also supports configurable tool command strings: `codex_cmd`, `gemini_cmd` (and optional `autopilot_cmd`), enabling Action Packs to call these tools directly.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Preserve the existing configurable command pattern already used for `oracle_cmd`, `task_master_cmd`, `tm_cmd`.
- Defaults as stated in the ticket text (`codex`, `gemini`, `${tm_cmd} autopilot`).

Evidence:
- References to existing pattern: `oracle_cmd`, `task_master_cmd`, `tm_cmd`
- Proposed additions: `codex_cmd`, `gemini_cmd`, `autopilot_cmd`

Open Items / Unknowns:
- Where the Stage-3 generator defines/declares these args (exact file/path not provided).
- How generated artifacts currently surface tool command strings (exact schema not provided).

Risks / Dependencies:
- Depends on T2 if these command strings must also be emitted into `_actions.json.tooling`.

Acceptance Criteria:
- Stage-3 generator accepts `codex_cmd` and `gemini_cmd` (and optional `autopilot_cmd`) as inputs.
- Defaults match the ticket text when values are not provided.
- Existing inputs (`oracle_cmd`, `task_master_cmd`, `tm_cmd`) remain supported and unchanged.
- Generated Action Pack artifacts include/use these variables (where the template expects tool invocations).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Your Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`”
- “Extend that same pattern for: `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`”
```

```ticket T2
T# Title:
- Extend `_actions.json` to include per-item executor metadata and tool command strings

Type:
- enhancement

Target Area:
- Canonical actions artifacts (`_actions.json` / `_actions.md`) generated from Stage-2 outputs

Summary:
- Add explicit executor metadata to each action item so downstream Action Pack steps can deterministically decide which tool to run.
- Add a `tooling` object to carry tool command strings (including Codex/Gemini) for use by the dispatcher.

In Scope:
- Extend `_actions.json` to include `tooling` with: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`.
- For each action item, add:
  - `executor`: `"codex" | "gemini" | "tm" | "manual"`
  - `exec_prompt`: short instruction string (no code fences; deterministic)
  - `inputs`: optional list of paths/globs (as referenced by ticket text)
- Ensure `_actions.md` can be produced from the JSON after the schema extension (format details not provided in ticket).

Out of Scope:
- Not provided

Current Behavior (Actual):
- `_actions.json` exists and includes `tooling` and per-item fields such as `recommended_next_action`, `acceptance_criteria` (as described in ticket text).

Expected Behavior:
- `_actions.json` includes the new `tooling` and per-item executor fields so execution can be dispatched without guessing.
- `_actions.md` can still be generated from `_actions.json`.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- `exec_prompt` must be short and deterministic (“no code fences” stated in ticket text).
- `inputs` is optional and can be derived from “missing_artifacts / repo anchors” (source details not provided).

Evidence:
- Proposed schema fields: `tooling`, `executor`, `exec_prompt`, `inputs`
- Existing per-item fields referenced: `recommended_next_action`, `acceptance_criteria`

Open Items / Unknowns:
- Exact current `_actions.json` schema and where it is defined (not provided).
- How `_actions.md` is rendered from `_actions.json` (not provided).

Risks / Dependencies:
- T3 depends on these fields being present to implement the dispatcher logic.

Acceptance Criteria:
- `_actions.json` includes `tooling` with the listed command keys.
- Each action item includes `executor` and `exec_prompt`; `inputs` is present when available and omitted/empty when not.
- The executor enum matches the ticket text: `codex|gemini|tm|manual`.
- `_actions.md` generation continues to work with the extended JSON structure.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Extend `_actions.json` to include an executor plan per item”
- “Add fields like… `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`”
- “per item… `executor`… `exec_prompt`… `inputs`”
```

```ticket T3
T# Title:
- Add `implement` dispatch mode (or Step 09) to Action Pack template to run Codex/Gemini based on `_actions.json`

Type:
- enhancement

Target Area:
- Stage-3 Action Pack template (modes: `backlog|pipelines|autopilot`) and step execution flow

Summary:
- Introduce an execution path that reads `_actions.json`, selects the top N actions, and dispatches to the appropriate tool based on `executor`.
- This is intended to make Action Packs perform actual implementation calls (e.g., `codex exec …`, `gemini -p …`) instead of only producing planning artifacts.

In Scope:
- Add either:
  - `mode=implement`, or
  - a new Step 09 gated by `MODE == implement`.
- In that mode/step:
  - Read `<out_dir>/_actions.json`.
  - Select the top N items (based on existing `top_n` concept referenced in ticket text).
  - Dispatch:
    - `codex exec …` for items with executor `codex`
    - `gemini -p …` for items with executor `gemini`
  - (Other executor values referenced in ticket text: `tm`, `manual` — behavior not specified beyond inclusion in enum.)

Out of Scope:
- Not provided

Current Behavior (Actual):
- Action Pack modes exist: `backlog|pipelines|autopilot`.
- Current example pack “does not call `codex`, `gemini`, etc. at all” (per ticket text).

Expected Behavior:
- When `MODE=implement`, the Action Pack executes tool commands for actions according to `_actions.json.executor`.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Use `_actions.json` as the source of truth for executor decisions (as described).
- Top-N selection should use the existing `top_n` referenced in ticket text.

Evidence:
- Proposed mode/step: “Add an ‘implement’ mode (or a Step 09)”
- Proposed dispatch: “`codex exec …`… `gemini -p …`”
- Top-N: “selects the top N items (you already have `top_n`)”

Open Items / Unknowns:
- Where `top_n` is currently defined and how ranking is determined (not provided).
- Exact location/structure for `<out_dir>/_actions.json` (not provided).

Risks / Dependencies:
- Depends on T2 because dispatcher requires `executor`/`exec_prompt` data in `_actions.json`.
- Depends on T1 if dispatcher uses `codex_cmd` / `gemini_cmd` variables.

Acceptance Criteria:
- Action Pack supports a distinct execution path for `implement` (as a mode or gated step).
- In `implement`, the Action Pack reads `_actions.json` and dispatches at least `codex` and `gemini` executors using the stated command forms.
- Existing modes (`backlog|pipelines|autopilot`) remain unchanged in behavior.
- If `_actions.json` is missing/unreadable, the implement path fails clearly (exact messaging not provided).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`”
- “dispatches deterministically: `codex exec …`… `gemini -p …`”
- “This particular Action Pack does not call `codex`, `gemini`, etc. at all”
```

```ticket T4
T# Title:
- Add non-interactive and safety guardrails for Codex/Gemini execution in Action Packs

Type:
- enhancement

Target Area:
- Action Pack execution behavior when running non-`oracle` tools (Codex/Gemini), including runner constraints around stdin/TTY

Summary:
- The ticket notes that oraclepack runs step scripts without attaching stdin/TTY, which can break or hang interactive CLIs.
- Add guardrails so Action Packs use non-interactive invocation patterns and keep defaults safe (explicitly: do not “yolo” by default).

In Scope:
- Ensure Codex invocation is non-interactive (ticket references `codex exec …` as the intended entrypoint).
- Ensure Gemini invocation is non-interactive (ticket references `gemini -p/--prompt` as intended).
- Apply safety defaults: “Keep safety defaults strict (do not ‘yolo’ by default)” (flag specifics not mandated beyond this phrase).
- Reflect/acknowledge runner constraint: “does not attach stdin/TTY” and the resulting failure mode for interactive tools.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Runner behavior described: executes scripts via `bash -lc`, does not attach stdin/TTY; interactive CLIs may fail or stall.

Expected Behavior:
- Action Packs avoid interactive-only command forms and use non-interactive command forms for Codex/Gemini.
- Defaults remain conservative and do not enable “yolo” behavior by default.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must account for “oraclepack doesn’t provide stdin/TTY to subprocesses” (per ticket text).
- Must keep safety defaults strict; do not enable “yolo” by default.

Evidence:
- “oraclepack will execute each step… but it does not attach stdin/TTY to the subprocess”
- “interactive CLIs… fail immediately… or stall waiting for input”
- “Keep safety defaults strict (do not ‘yolo’ by default)”

Open Items / Unknowns:
- Exact approval/safety flags to use for Gemini/Codex (not provided in ticket text beyond “yolo” mention).
- Where to place guardrails (template preflight vs dispatcher vs documentation) not specified.

Risks / Dependencies:
- Works alongside T3 (implement dispatcher); guardrails should apply to that path.

Acceptance Criteria:
- Implement path uses non-interactive forms (`codex exec …`, `gemini -p/--prompt …`) as referenced in ticket text.
- Default behavior does not enable “yolo” execution.
- Action Pack behavior is compatible with “no stdin/TTY” runner constraint (i.e., does not require interactive prompts to proceed).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “oraclepack… does not attach stdin/TTY… interactive CLIs… fail… or stall”
- “Keep safety defaults strict (do not ‘yolo’ by default)”
- “Codex… `codex exec`… Gemini CLI… `--prompt` / `-p`”
```

```ticket T5
T# Title:
- Generalize oraclepack override injection/validation beyond `oracle` to a multi-tool command registry (optional)

Type:
- enhancement

Target Area:
- oraclepack core: command detection for overrides injection and overrides validation (currently `oracle`-anchored)

Summary:
- The ticket identifies that oraclepack’s “nice UX features” (flag overrides + validation) are currently oracle-specific because detection/injection targets only `oracle` commands.
- Optional work: generalize detection to support additional tool prefixes and add per-tool override sets.

In Scope:
- Generalize command detection/injection from `oracle`-only to a small registry of command prefixes (explicitly listed in ticket text): `oracle`, `codex`, `gemini`, `task-master`, `tm`.
- Add per-tool override sets so “Oracle Flags” are not the only concept when tools differ.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Special logic targets only lines that start with `oracle`: detects invocations, injects flags into `oracle …`, and validates overrides using `oracle --dry-run summary` (as described in ticket text).

Expected Behavior:
- The system can detect/inject overrides for additional tool prefixes listed above.
- Overrides configuration is tool-specific (per-tool override sets).

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must preserve the existing oracle behavior while extending to additional command prefixes.

Evidence:
- “special logic only for lines that start with `oracle`… detects… injects… overrides validation”
- “Generalize… to a small registry of command prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`)”
- “Add per-tool override sets”

Open Items / Unknowns:
- Exact internal functions/files to modify (names mentioned but paths not provided in this ticket text).
- How validation should work for non-`oracle` tools (not specified beyond “per-tool override sets”).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- Non-`oracle` command prefixes listed in the ticket are recognized by the override injection mechanism.
- Overrides can be configured per tool (not only “Oracle Flags”).
- Existing `oracle` override injection and validation behavior continues to function.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Oraclepack… special logic only for lines that start with `oracle`”
- “Generalize… to… prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`)”
- “Add per-tool override sets”
```
```

.tickets/Oraclepack Action Pack Issue.md
```
Parent Ticket:

* Title: Oraclepack Action Pack compatibility and non-`oracle` command handling clarity
* Summary: Clarify what happens when running `oraclepack-taskify`-generated Action Packs in current oraclepack, especially the difference between “executing the pack” vs “dispatching/wrapping non-`oracle` commands.” Capture current limitations: oraclepack’s special handling (flag injection + override validation) targets `oracle` invocations only; other CLIs run as plain shell commands and may fail/block depending on PATH and interactivity.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Two meanings of ‘run these actionpack artifacts’… execute the pack file vs dispatch non-`oracle` commands.”
* Global Constraints:

  * Action Pack is described as “oraclepack-ingestible” with “single ```bash fence” and “# NN) step headers.”
* Global Environment:

  * Unknown
* Global Evidence:

  * docs.task-master.dev link(s) referenced in ticket text
  * developers.openai.com Codex CLI link(s) referenced in ticket text
  * docs.cloud.google.com / google-gemini.github.io Gemini CLI link(s) referenced in ticket text

Split Plan:

* Coverage Map:

  * Original item: “Can our current oraclepack run these actionpack artifacts… generated from the oraclepack-taskify skill?” → Assigned Ticket ID: T1
  * Original item: “Action Pack… ‘oraclepack-ingestible’ (single ```bash fence, # NN) step headers, deterministic paths…)” → Assigned Ticket ID: T1
  * Original item: “How you’d run it… `oraclepack validate …` / `oraclepack run …`” → Assigned Ticket ID: T1
  * Original item: “Actionpacks calling other agents/tools (codex exec, tm autopilot, gemini)… will run them as long as installed and usable non-interactively” → Assigned Ticket ID: T1
  * Original item: “Why did another reviewer say otherwise?” → Assigned Ticket ID: T1
  * Original item: “oraclepack runs each step as a bash script (bash -lc <step script>)” → Assigned Ticket ID: T1
  * Original item: “injects flags into commands that begin with oracle… does not match tm/task-master/codex/gemini” → Assigned Ticket ID: T3
  * Original item: “TUI ‘override validation’ only targets oracle… steps without oracle invocations are skipped” → Assigned Ticket ID: T2
  * Original item: “If binary isn’t installed/on PATH → step fails; if CLI is interactive → it will block” → Assigned Ticket ID: T1
  * Original item: “Only way it ‘gets confused’ is if you expect oracle output text to magically invoke Codex/Gemini” → Assigned Ticket ID: T1
  * Original item: “Two meanings… execute vs dispatch/apply oraclepack overrides to non-oracle commands” → Assigned Ticket ID: T1
  * Original item: “Reviewer answered ‘No’ because broader goal is dispatcher behavior… not interpret actions” → Assigned Ticket ID: Info-only
  * Original item: “Reconciliation: both statements can be true” → Assigned Ticket ID: Info-only
* Dependencies:

  * Not provided
* Split Tickets:

````ticket T1
T# Title: Document current Action Pack execution semantics and operator expectations
Type: docs
Target Area: oraclepack documentation / runbook for Action Packs
Summary:
  Clarify what “running an Action Pack” means in current oraclepack: steps execute as shell via `bash -lc`, and any `tm`/`codex`/`gemini` lines run as plain shell commands. Document the two meanings that caused reviewer disagreement: executing the pack vs dispatching/applying oracle-specific override logic to non-`oracle` commands. Include operator-facing notes on common failure modes already described (PATH missing binaries, interactive CLIs blocking, environment guardrails for autopilot).
In Scope:
  - Explain: “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - Explain: non-`oracle` commands (`tm`, `task-master`, `codex`, `gemini`) are executed “directly as normal shell commands.”
  - Provide the exact run examples already given (`oraclepack validate …` and `oraclepack run …`).
  - Capture the limitation: oraclepack’s oracle-specific overrides/validation do not apply to non-`oracle` commands.
  - Document noted failure/blocking conditions:
    - Missing binaries not on PATH.
    - Interactive CLIs blocking waiting for input.
    - Autopilot “fail fast” environment issues (e.g., not in git repo, dirty tree) as stated.
Out of Scope:
  - Implementing dispatcher behavior for non-`oracle` tools (not specified beyond “broader goal” mention).
Current Behavior (Actual):
  - Steps are run as `bash -lc <step script>`.
  - Non-`oracle` commands run directly; no automatic dispatch/wrapping is applied.
Expected Behavior:
  - Operators can correctly predict:
    - What will execute (literal commands in the pack).
    - What will not happen automatically (no “magic” invocation of Codex/Gemini from oracle output text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Preserve stated Action Pack constraints: “single ```bash fence” and “# NN) step headers” (as described).
Evidence:
  - Included links mentioned in ticket text:
    - docs.task-master.dev references
    - developers.openai.com Codex CLI references
    - Gemini CLI references (docs.cloud.google.com / google-gemini.github.io)
Open Items / Unknowns:
  - Exact location(s) where this documentation should live (Not provided).
  - Whether this should be surfaced in CLI help text vs README vs TUI (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Documentation explicitly states:
    - Packs execute step bodies via `bash -lc`.
    - Non-`oracle` commands run as-is and are not routed through oracle-specific logic.
    - Common failure modes listed in the ticket text (PATH missing, interactive blocking, autopilot environment guards).
  - Documentation includes the exact run commands already provided (`oraclepack validate …`, `oraclepack run …`).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - “Will oraclepack dispatch non-`oracle` commands…? No… only targets commands that start with `oracle`.”
  - “If the CLI is interactive → it will block waiting for input.”
````

```ticket T2
T# Title: Make TUI override validation behavior explicit for steps without `oracle` invocations
Type: enhancement
Target Area: TUI overrides flow / override validation messaging
Summary:
  The ticket text states that TUI “override validation” runs `oracle --dry-run summary` on detected `oracle` invocations and skips steps without `oracle` calls. Make this behavior explicit in the TUI so operators do not misinterpret skipped steps as “validated,” especially when packs include `tm`/`codex`/`gemini` commands.
In Scope:
  - Surface an explicit note/state in the overrides validation flow indicating:
    - Validation applies only to detected `oracle` invocations.
    - Steps without `oracle` invocations are skipped by this validator.
Out of Scope:
  - Adding validation implementations for `tm`/`codex`/`gemini` (not described in ticket text).
Current Behavior (Actual):
  - “Override validation… only targets `oracle` commands… Steps without oracle invocations are skipped.”
Expected Behavior:
  - TUI clearly communicates when steps are skipped (and why), avoiding operator confusion.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must preserve current behavior described (validate only `oracle` invocations) unless changed elsewhere (Not provided).
Evidence:
  - “The TUI ‘override validation’ also only targets `oracle` commands… Steps without oracle invocations are skipped…”
Open Items / Unknowns:
  - Exact UI copy/placement and which screen(s) in the TUI should show this (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - When override validation runs, the UI explicitly indicates:
    - It validates only `oracle` invocations (via `oracle --dry-run summary`, as stated).
    - Steps without `oracle` invocations are skipped (and shown as skipped).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “The TUI ‘override validation’ also only targets `oracle` commands… runs `oracle --dry-run summary`…”
  - “Steps without oracle invocations are skipped by that validator.”
```

```ticket T3
T# Title: Extend runtime flag-injection matching beyond `oracle` invocations (configurable prefixes)
Type: enhancement
Target Area: command rewriting / runtime override injection
Summary:
  The ticket text states oraclepack injects flags only into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`) and does not match `tm`, `task-master`, `codex`, or `gemini`. Add support for matching additional command prefixes (or a configurable list) so override injection is not limited to `oracle` only, aligning with packs that include other CLIs.
In Scope:
  - Expand the “inject flags” behavior beyond `oracle`-only matching, as motivated by:
    - “does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
  - Preserve anchored/prefix-based matching semantics as described (no broad substring matching implied).
Out of Scope:
  - Defining tool-specific semantics for what flags should be injected for each CLI (Not provided).
  - Implementing dispatcher logic that changes execution from “literal shell command” to “interpreted actions” (Not provided).
Current Behavior (Actual):
  - “Injects flags into commands that begin with `oracle`… regex anchored to `^(\s*)(oracle)\b`.”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Expected Behavior:
  - Flag injection can apply to non-`oracle` command prefixes as configured/defined (details not provided in ticket text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must not break execution of packs that rely on current `oracle`-only behavior (Not provided).
Evidence:
  - “injects flags into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`). It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Open Items / Unknowns:
  - Which non-`oracle` commands should be included first (list not provided beyond examples).
  - Where configuration should live (Not provided).
  - Whether injection should be opt-in per pack/step or global (Not provided).
Risks / Dependencies:
  - Risk: unintended rewriting of commands if prefix matching is overly broad (mitigate via anchored matching; exact approach not provided).
Acceptance Criteria:
  - There is a documented/configured way to include additional command prefixes for injection beyond `oracle`.
  - Existing `oracle` prefix injection continues to work unchanged.
  - Demonstrably, a command beginning with one added prefix is recognized for injection (exact flags and CLI semantics: Not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “injects flags into commands that begin with `oracle` (regex is anchored to `^(\\s*)(oracle)\\b`).”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
```
```

.tickets/Oraclepack Action Packs.md
```
Parent Ticket:

* Title: Oraclepack Action Packs: tool-agnostic execution (Codex/Gemini/Task Master) instead of oracle-only flows
* Summary:

  * Current pain point: oraclepack is a wrapper around `oracle`, and “more oracle calls” won’t implement taskified work; Action Packs must run real tool commands (e.g., `codex`, `gemini`, `tm`, `task-master`).
  * Key idea: keep oracle as “planner” and make Action Packs do deterministic executor dispatch via `_actions.json` metadata and a new implement mode/step.
* Source:

  * Link/ID (if present) or “Not provided”: Uploaded file: Oraclepack Action Pack Integration.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make it so our actionpacks will perform the correct calls… Example ‘codex exec …’, ‘tm …’, ‘gemini …’”.
* Global Constraints:

  * Action Packs should be “tool-agnostic” (dispatch to `codex`, `gemini`, `tm`, etc.).
  * `exec_prompt` should be short and deterministic (explicitly “no code fences”).
  * Safety defaults must be strict (“do not yolo by default”); conservative approval/tool execution unless opted-in.
  * Optional/“nice-to-have”: generalize oraclepack’s oracle-specific overrides UX for other command prefixes.
* Global Environment:

  * Unknown
* Global Evidence:

  * Mentioned repos/tools:

    * `https://github.com/google-gemini/gemini-cli`
    * `https://github.com/eyaltoledano/claude-task-master`
    * `https://github.com/openai/codex`
    * `https://github.com/steipete/oracle`
  * Reference docs cited/used in the ticket text:

    * Codex CLI reference (supports “codex exec”). ([OpenAI Developers][1])
    * Gemini CLI docs (tools/approval concepts referenced by the ticket). ([Gemini CLI][2])
    * Claude Task Master repository. ([GitHub][3])

Split Plan:

* Coverage Map:

  * “make it so our actionpacks will perform the correct calls… ‘codex exec …’, ‘tm …’, ‘gemini …’” → T3
  * Oraclepack has special logic only for lines starting with `oracle` (regex detection / flag injection / validation) → T5
  * “Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) … guarded branch checks” → T3
  * Stage-3 skill supports `oracle_cmd`, `task_master_cmd`, `tm_cmd` → T1
  * “Extend that same pattern for `codex_cmd`, `gemini_cmd`, optionally `autopilot_cmd`” → T1
  * “Extend `_actions.json` … `executor`, `exec_prompt` (no code fences), `inputs`, plus expanded `tooling`” → T2
  * “Add an ‘implement’ mode (or Step 09)… reads `<out_dir>/_actions.json`… selects top N (`top_n`)… dispatches” → T3
  * “Keep safety defaults strict (do not ‘yolo’ by default)” → T4
  * “Minimal changes… add args… update Prompt A… update Action Pack template” → T1, T2, T3
  * “Optional: improve oraclepack UX… registry of command prefixes… per-tool override sets” → T5
  * “Concrete command patterns… `codex exec`, `gemini -p`, Task Master pipeline commands” → Info-only
  * “If you want, I can propose the exact schema delta… Step 09 bash logic” → Info-only
* Dependencies:

  * T3 depends on T2 because the implement/dispatcher step reads `_actions.json` and needs `executor`/`exec_prompt` metadata.
  * T2 depends on T1 because `_actions.json.tooling` expansion references new tool command fields (`codex_cmd`, `gemini_cmd`, optional `autopilot_cmd`).
  * T4 depends on T3 because safety defaults/opt-ins apply to the implement/dispatcher execution path.
  * T5 is independent (optional) but may follow T3 if you want the TUI/overrides UX to apply to non-oracle commands.
* Split Tickets:

```ticket T1
T1 Title:
- Extend oraclepack-taskify Stage-3 generator to accept and propagate executor CLI command configs (codex/gemini/autopilot)

Type:
- enhancement

Target Area:
- oraclepack-taskify (Stage-3 generator inputs/args and emitted configs)

Summary:
- The Stage-3 generator already supports configurable command strings for `oracle_cmd`, `task_master_cmd`, and `tm_cmd`.
- Extend the same configuration pattern to include `codex_cmd` and `gemini_cmd`, and optionally `autopilot_cmd`, so Action Packs can invoke the intended executors without hard-coding tool names.

In Scope:
- Add generator inputs/args for:
  - `codex_cmd` (default `codex`)
  - `gemini_cmd` (default `gemini`)
  - optional `autopilot_cmd` (default `${tm_cmd} autopilot`)
- Ensure generated artifacts carry these command strings for later use by the Action Pack execution steps.

Out of Scope:
- Modifying oraclepack core/TUI behavior (handled in T5)
- Implement-mode dispatcher logic (handled in T3)

Current Behavior (Actual):
- Stage-3 skill supports configurable tool command strings:
  - `oracle_cmd`, `task_master_cmd`, `tm_cmd`

Expected Behavior:
- Stage-3 generator also supports `codex_cmd`, `gemini_cmd`, and optionally `autopilot_cmd`, using the same configuration pattern.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Preserve the existing “configurable tool command strings” pattern already used by the Stage-3 generator.
- Defaults as stated in the ticket text.

Evidence:
- References: “Your Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`… Extend… `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`…”

Open Items / Unknowns:
- Where these args are defined/passed in the current generator (file paths not provided).
- Whether additional executors beyond codex/gemini/tm are needed (not provided).

Risks / Dependencies:
- Depends on T2 if `_actions.json.tooling` is expanded to include these command strings.

Acceptance Criteria:
- Stage-3 generator accepts `codex_cmd` and `gemini_cmd` (and optional `autopilot_cmd`) inputs.
- Defaults match: `codex`, `gemini`, and `${tm_cmd} autopilot` (optional).
- Generated outputs expose these command strings for downstream Action Pack steps.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Your Stage-3 skill already supports configurable tool command strings: `oracle_cmd`, `task_master_cmd`, `tm_cmd`”
- “Extend that same pattern for: `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`”
```

```ticket T2
T2 Title:
- Extend `_actions.json` schema + Prompt A to emit per-item executor metadata (executor/exec_prompt/inputs) and expanded tooling

Type:
- enhancement

Target Area:
- Canonical actions schema (`_actions.json`) and “Prompt A” (canonical actions synthesis)

Summary:
- The current actions schema has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action` and `acceptance_criteria`.
- Add executor planning fields per action item so Action Packs can deterministically select and run the correct executor (`codex`, `gemini`, `tm`, or manual), and include relevant inputs.

In Scope:
- Update `_actions.json` to add:
  - `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  - per-item fields:
    - `executor`: `"codex" | "gemini" | "tm" | "manual"`
    - `exec_prompt`: short instruction string (explicitly “no code fences; deterministic”)
    - `inputs`: optional list of paths/globs (from `missing_artifacts` / repo anchors)
- Update “Prompt A” to emit the above fields for each action item.

Out of Scope:
- Implement-mode dispatcher logic that consumes `_actions.json` (handled in T3)
- oraclepack core UX changes (handled in T5)

Current Behavior (Actual):
- `_actions.json` has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action`, `acceptance_criteria`.

Expected Behavior:
- `_actions.json` includes explicit executor plan per item (`executor`, `exec_prompt`, optional `inputs`) and expanded `tooling` including codex/gemini command strings.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- `exec_prompt` must be short and deterministic; explicitly “no code fences”.
- Executor enum values and tooling fields as written in the ticket text.

Evidence:
- References: “Extend `_actions.json` to include an executor plan per item… `executor`… `exec_prompt`… `inputs`…”

Open Items / Unknowns:
- The exact existing `_actions.json` schema shape and where it’s validated (not provided).
- Whether `missing_artifacts` / repo anchors already exist in the schema (not provided).

Risks / Dependencies:
- Depends on T1 if `tooling` is expected to include the new `codex_cmd`/`gemini_cmd` command strings.
- Required by T3 since the dispatcher reads `_actions.json`.

Acceptance Criteria:
- `_actions.json` schema includes `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`.
- Each action item can include:
  - `executor` with allowed values: `codex`, `gemini`, `tm`, `manual`
  - `exec_prompt` (no code fences requirement captured)
  - optional `inputs` list
- Prompt A output includes these fields per item.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Extend `_actions.json` to include an executor plan per item”
- “`executor`: ‘codex’ | ‘gemini’ | ‘tm’ | ‘manual’”
- “`exec_prompt`: a short instruction string (no code fences; deterministic)”
```

```ticket T3
T3 Title:
- Add Action Pack “implement” mode (or Step 09) to dispatch executor commands based on `_actions.json`

Type:
- enhancement

Target Area:
- Stage-3 Action Pack template (runner-ingestible actionpack artifacts)

Summary:
- Action Packs must execute real tool commands to implement tasks; adding more `oracle` calls only produces more analysis/synthesis.
- Add an implement execution path that reads `_actions.json`, selects top N actions (`top_n`), and dispatches to the specified executor (e.g., `codex exec …`, `gemini -p …`) using per-item metadata.

In Scope:
- Add either:
  - `mode=implement`, or
  - a new Step 09 guarded by `if MODE == implement`
- Implement-mode behavior:
  - Read `<out_dir>/_actions.json`
  - Select the top N items (uses existing `top_n`)
  - Dispatch deterministically:
    - `codex exec …` for items with `executor=codex`
    - `gemini -p …` for items with `executor=gemini`
- Keep existing modes (`backlog|pipelines|autopilot`) intact.

Out of Scope:
- Changing oraclepack core logic for overrides/flag injection/validation (handled in T5)

Current Behavior (Actual):
- Implementation happens only when the Action Pack runs real tool commands (e.g., `task-master …`, `tm autopilot`).
- The Action Pack already runs non-oracle tools (`task-master …` and `tm autopilot`) and includes guarded branch checks for autopilot.

Expected Behavior:
- An implement mode/step exists that consumes `_actions.json` and runs executor-specific commands deterministically based on per-item metadata.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must read `<out_dir>/_actions.json`.
- Must use existing `top_n` to select items.
- Must dispatch based on `executor` field.
- Preserve deterministic behavior as described (stable ordering / fail-fast preflight referenced as desired properties).

Evidence:
- References: “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically…”

Open Items / Unknowns:
- The exact Action Pack template file path and step numbering constraints (not provided).
- How “top N” is currently computed/ordered (not provided).

Risks / Dependencies:
- Depends on T2 since implement mode reads `_actions.json` executor metadata.
- Safety defaults for tool execution are addressed in T4.

Acceptance Criteria:
- Action Pack supports `mode=implement` (or an equivalent Step 09 guarded behavior).
- Implement mode:
  - Reads `<out_dir>/_actions.json`
  - Selects top N via `top_n`
  - Runs `codex exec …` for `executor=codex`
  - Runs `gemini -p …` for `executor=gemini`
- Existing `backlog|pipelines|autopilot` behavior remains unchanged.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “implementation only happens when the Action Pack runs real tool commands”
- “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically”
```

```ticket T4
T4 Title:
- Enforce conservative safety defaults for executor dispatch (no auto-approve “yolo” behavior unless opted-in)

Type:
- chore

Target Area:
- Action Pack implement/dispatcher execution path (gemini/codex dispatch configuration)

Summary:
- The ticket explicitly calls out safety: executor tools may run commands or perform actions; defaults should be conservative.
- Add/confirm guardrails so the implement/dispatcher mode does not auto-approve tool execution by default, and requires explicit opt-in for riskier behaviors.

In Scope:
- Ensure implement/dispatcher mode defaults are “strict” and not “yolo by default”.
- Ensure any approval/tool-execution modes are conservative unless a user opts in (mechanism not specified in the ticket text).

Out of Scope:
- Defining new security models or sandboxing systems beyond what’s stated (not provided).

Current Behavior (Actual):
- Not provided (safety behavior in implement/dispatcher mode is not yet implemented per ticket text).

Expected Behavior:
- Implement/dispatcher execution has conservative defaults around tool approvals/execution; opt-in required for less restrictive modes.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- “Keep safety defaults strict (do not ‘yolo’ by default)”
- “keep defaults conservative unless a user opts in”

Evidence:
- References: “Keep safety defaults strict (do not ‘yolo’ by default)… sandboxing and explicit approvals matter…”

Open Items / Unknowns:
- Which specific flags/options are used to control approvals for each executor (not provided).
- How/where opt-in is configured (env var, arg, config) (not provided).

Risks / Dependencies:
- Depends on T3 because the implement/dispatcher path is where safety defaults apply.

Acceptance Criteria:
- Implement/dispatcher mode does not default to auto-approving tool execution.
- Any non-conservative approval behavior requires explicit user opt-in (as defined by existing config patterns; specifics not provided).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Keep safety defaults strict (do not ‘yolo’ by default)”
- “keep defaults conservative unless a user opts in”
```

```ticket T5
T5 Title:
- Optional: Generalize oraclepack “oracle-only” overrides UX to support multiple command prefixes (codex/gemini/task-master/tm)

Type:
- enhancement

Target Area:
- oraclepack core/TUI overrides and validation logic (oracle-specific detection/injection/validation)

Summary:
- Oraclepack can run any shell command, but its “nice UX features” (flag overrides + validation) are oracle-specific.
- Optionally generalize the command detection and flag-injection mechanisms so overrides/validation can apply to non-oracle tools used by Action Packs.

In Scope:
- Generalize `ExtractOracleInvocations` / `InjectFlags` to a small registry of command prefixes:
  - `oracle`, `codex`, `gemini`, `task-master`, `tm`
- Add per-tool override sets (so “Oracle Flags” semantics don’t incorrectly apply to other tools).
- Keep current oracle behavior intact.

Out of Scope:
- Implement-mode executor dispatch itself (handled in T3)
- Any new validation semantics beyond “oracle --dry-run summary” analogy (not provided)

Current Behavior (Actual):
- Special logic only for lines that start with `oracle`:
  - Invocation detection via regex anchored to literal `oracle`
  - Flag injection only into `oracle …` lines
  - Overrides validation runs `oracle --dry-run summary` only for detected oracle invocations

Expected Behavior:
- oraclepack recognizes multiple command prefixes for the purposes of overrides/flag handling and (where applicable) validation hooks, without breaking existing oracle behavior.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must not regress oracle detection/injection/validation behavior.
- Registry-based handling for additional tool prefixes as listed in the ticket text.

Evidence:
- References: “special logic only for lines that start with `oracle`… detects… injects… validation runs `oracle --dry-run summary`…”
- References: “Optional: improve oraclepack UX… registry of command prefixes… Add per-tool override sets…”

Open Items / Unknowns:
- Exact code locations for `ExtractOracleInvocations` / `InjectFlags` and override sets in the repo (paths not provided in this ticket text).

Risks / Dependencies:
- Not required to make Action Packs execute non-oracle tools; explicitly described as “nice-to-have”.
- Potential semantic mismatch if oracle-style overrides are applied to other CLIs without per-tool override sets.

Acceptance Criteria:
- oraclepack supports a registry of command prefixes including: `oracle`, `codex`, `gemini`, `task-master`, `tm`.
- Overrides/flag injection is not hard-coded to only `oracle` command lines.
- Per-tool override sets exist (or equivalent structure) so overrides are not incorrectly treated as “Oracle Flags” for non-oracle tools.
- Existing oracle-only behavior remains functional.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Oraclepack executes each step as shell, but it has special logic only for lines that start with `oracle`”
- “Generalize `ExtractOracleInvocations` / `InjectFlags`… registry of command prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`)”
- “Add per-tool override sets…”
```

[1]: https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com "Command line options"
[2]: https://geminicli.com/docs/tools/?utm_source=chatgpt.com "Gemini CLI tools"
[3]: https://github.com/eyaltoledano/claude-task-master?utm_source=chatgpt.com "eyaltoledano/claude-task-master"
```

.tickets/Oraclepack Compatibility Issues.md
```
Parent Ticket:

* Title: Oraclepack Actionpack Compatibility: non-`oracle` tools execution, dispatcher/overrides gaps, and adding Codex/Gemini headless steps
* Summary:

  * There is confusion about whether oraclepack can run Action Packs that include non-`oracle` commands (e.g., `task-master` / `tm`, `codex`, `gemini`). The current behavior is that oraclepack executes each step as shell (`bash -lc ...`) and only applies oracle-specific injection/validation to commands that begin with `oracle`. The request also includes adding headless `gemini` + non-interactive `codex exec` automation into the placeholder steps of `ticket-action-pack.md`, and optionally extending this pattern to taskify-generated packs.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`… won’t dispatch/wrap non-`oracle` commands” (per file)
* Global Constraints:

  * Action Pack is “oraclepack-ingestible” (single `bash` fence, `# NN)` steps) (per file)
  * Do not assume oraclepack overrides apply to non-`oracle` commands (per file)
* Global Environment:

  * Steps run via `bash -lc ...` in the project root; oraclepack does not change `WorkDir` to `out_dir` (per file)
  * ROI filter behavior may skip steps with no `ROI=` if a threshold > 0 is used (per file)
* Global Evidence:

  * References: `oracle_pack_and_taskify-skills.md`, `oraclepack-tui.md`, `ticket-action-pack.md` (per file)
  * Noted behaviors: oracle-only regex targeting; override validation runs `oracle --dry-run summary` on detected oracle invocations (per file)
  * Mentioned outputs: `.oraclepack/ticketify/_tickets_index.json`, `_actions.json`, `_actions.md`, `.taskmaster/docs/tickets_prd.md`, `.oraclepack/ticketify/tm-complexity.json`, `ticket-action-pack.state.json`, `ticket-action-pack.report.json` (per file)

Split Plan:

* Coverage Map:

  * “oraclepack runs each step as a bash script (`bash -lc <step script>`).” → T1
  * “injects flags into commands that begin with `oracle`… regex anchored to `^(\\s*)(oracle)\\b`.” → T2
  * “TUI override validation… runs `oracle --dry-run summary`… skips steps without oracle invocations.” → T2
  * “`tm` / `task-master` run directly… not routed through oracle.” → T1
  * “If you manually add `codex` / `gemini` lines… oraclepack will try to run them directly.” → T1
  * “If CLI isn’t installed/on PATH → step fails; if interactive → blocks.” → T3
  * “If you don’t add those commands… pack mainly uses `oracle` + Task Master; won’t ‘implement’ via Codex/Gemini.” → T1
  * `ticket-action-pack.md` Step 01 writes `.oraclepack/ticketify/_tickets_index.json` → Info-only
  * `ticket-action-pack.md` Step 02 writes `_actions.json` + `_actions.md` → Info-only
  * `ticket-action-pack.md` Step 03 writes `.taskmaster/docs/tickets_prd.md` → Info-only
  * `ticket-action-pack.md` Steps 05–07 run `task-master parse-prd`, `analyze-complexity`, `expand --all` and write `.oraclepack/ticketify/tm-complexity.json` → Info-only
  * “Steps 08–20 are placeholders/notes (echo guidance).” → T3
  * “Best insertion points… placeholder steps (09–13 and 16).” → T3
  * Step 09: Gemini headless selects next target, writes `.oraclepack/ticketify/next.json` → T3
  * Step 10: `codex exec` implements selected action, writes `.oraclepack/ticketify/codex-implement.md` → T3
  * Step 11: verification via `codex exec` and/or Gemini diff review, writes `.oraclepack/ticketify/codex-verify.md` and/or `.oraclepack/ticketify/gemini-review.json` → T3
  * Step 16: Gemini drafts PR body, writes `.oraclepack/ticketify/PR.md` → T3
  * “Optional… add an agent-mode to oraclepack-taskify packs… keep 20-step contract intact.” → T4
  * “Key constraint… overrides only target commands that begin with `oracle`; codex/gemini won’t inherit unless wrap/extend oraclepack.” → T2
  * Failure notes: missing `.tickets/`, missing `task-master` / provider keys, ROI filter gotcha → T1
* Dependencies:

  * T3 depends on T2 because codex/gemini steps will not participate in oraclepack override injection/validation unless oraclepack is extended beyond `oracle`-prefixed commands (per file).
  * T4 depends on T2 for the same reason (per file).
* Split Tickets:

```ticket T1
T# Title:
- Clarify current oraclepack Action Pack execution semantics (and common failure modes)

Type:
- docs

Target Area:
- oraclepack CLI/TUI user-facing documentation (exact file(s) not provided)

Summary:
- Document the current behavior that oraclepack executes Action Pack steps as `bash -lc ...` and only applies oracle-specific behavior to `oracle`-prefixed commands. Capture practical implications for running packs containing `task-master`/`tm`, `codex`, and `gemini`, including common failure modes and the ROI filter gotcha noted in the ticket content.

In Scope:
- Document that steps execute as shell (`bash -lc ...`) and whatever commands appear in the step body are executed.
- Document that non-`oracle` commands (`task-master`/`tm`, `codex`, `gemini`) run directly and are not routed through oracle.
- Document the “interactive CLI can block” and “missing binary on PATH fails the step” implications.
- Document `ticket-action-pack.md` likely failure points called out: missing `.tickets/`, missing `task-master`/provider configuration/API keys, ROI filter gotcha.
- Document that Steps 08–20 are placeholders unless replaced with real commands.

Out of Scope:
- Changing oraclepack dispatcher / override injection logic.
- Editing `ticket-action-pack.md` steps to add new automation.

Current Behavior (Actual):
- Confusion among reviewers/users about whether oraclepack “runs everything through oracle.”
- Running packs without `oracle ...` commands results in no oracle-specific override behavior being applied.

Expected Behavior:
- A clear, discoverable doc section explains what oraclepack does/does not do with non-`oracle` commands and how to interpret failures.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Preserve the current semantics described in the ticket content (no implied change to execution model).

Evidence:
- “oraclepack runs each step as a bash script (`bash -lc <step script>`).”
- “If the CLI is interactive → it will block waiting for input.”
- “ROI filter gotcha… steps with no `ROI=`… may be skipped.”

Open Items / Unknowns:
- Where documentation should live (README, `oraclepack-tui.md`, or other) is not provided.
- Whether this should be shown in TUI help text vs repository docs is not provided.

Risks / Dependencies:
- Not provided.

Acceptance Criteria:
- Documentation explicitly states:
  - Steps execute via `bash -lc ...` and run the literal commands present.
  - Only `oracle`-prefixed commands receive oraclepack’s special handling.
  - Non-`oracle` tools (`tm`/`task-master`, `codex`, `gemini`) run directly (PATH/interactive caveats included).
  - The listed failure modes and ROI filter gotcha are described with practical guidance.
- Documentation includes a short “What to expect after running `ticket-action-pack.md`” section referencing the artifact paths named in the ticket content.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “oraclepack runs each step as a bash script (`bash -lc <step script>`).”
- “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`.”
- “ROI filter gotcha… steps with no `ROI=`… may be skipped.”
```

```ticket T2
T# Title:
- Extend oraclepack override injection/validation beyond `oracle`-prefixed commands (dispatcher changes)

Type:
- enhancement

Target Area:
- oraclepack command detection + overrides/validation pipeline (regex/dispatch behavior; exact file paths not provided)

Summary:
- The current oraclepack behavior applies oracle-specific transforms only to commands that begin with `oracle` and the TUI override validation only targets oracle invocations. Implement “dispatcher changes” so that non-`oracle` commands (explicitly referenced: `tm`/`task-master`, `codex`, `gemini`) can participate in the same override/validation flow, or otherwise be handled explicitly as first-class command targets.

In Scope:
- Update command detection so it is not limited to `oracle` (currently anchored to `^(\\s*)(oracle)\\b` per the ticket text).
- Update override validation so it does not only run `oracle --dry-run summary` on detected oracle invocations and skip steps without oracle invocations.
- Ensure steps containing `tm`/`task-master`, `codex`, and/or `gemini` can be detected for dispatcher/validation purposes (as described in the ticket content).

Out of Scope:
- Adding new `codex exec` / `gemini` automation steps to specific packs (handled in T3).
- Changing Task Master’s behavior or requirements.

Current Behavior (Actual):
- Flag injection “only… injects flags into commands that begin with `oracle`… does not match `tm`, `task-master`, `codex`, `gemini`.”
- TUI override validation “only targets `oracle` commands… runs `oracle --dry-run summary`… steps without oracle invocations are skipped.”

Expected Behavior:
- Dispatcher/override handling is not limited to `oracle`-prefixed commands for the explicitly mentioned tool commands, so non-`oracle` steps are not silently excluded from override/validation.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Must preserve existing `oracle` command behavior.
- Must address the limitation called out: overrides/validation currently only target `oracle` commands.

Evidence:
- “injects flags into commands that begin with `oracle` (regex… `^(\\s*)(oracle)\\b`).”
- “override validation… runs `oracle --dry-run summary`… steps without oracle invocations are skipped.”
- “codex/gemini won’t inherit oraclepack overrides unless you wrap them yourself or extend oraclepack.”

Open Items / Unknowns:
- Exact desired behavior for applying overrides to `tm`/`task-master`, `codex`, and `gemini` is not provided (which flags apply, how validation works).
- Whether the dispatcher should “interpret actions” vs only broaden prefix-based detection is not provided.

Risks / Dependencies:
- Risk: unclear spec for how overrides should apply to each non-`oracle` tool could lead to partial/incorrect behavior.

Acceptance Criteria:
- A pack step containing at least one of the referenced non-`oracle` command prefixes (`tm`/`task-master`, `codex`, `gemini`) is no longer automatically excluded from the override/validation pipeline solely due to not starting with `oracle`.
- Existing behavior for `oracle`-prefixed commands remains unchanged.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “injects flags into commands that begin with `oracle`… does not match `tm`, `codex`, `gemini`.”
- “override validation… only targets `oracle` commands… steps without oracle invocations are skipped.”
- “codex/gemini won’t inherit oraclepack overrides unless you wrap them yourself or extend oraclepack.”
```

```ticket T3
T# Title:
- Replace `ticket-action-pack.md` placeholder steps with headless Gemini + non-interactive Codex automation

Type:
- enhancement

Target Area:
- `ticket-action-pack.md` (ticketify Action Pack content/template; exact generator location not provided)

Summary:
- Steps 08–20 are described as placeholders/notes that only echo guidance. Replace specific placeholder steps (explicitly called out: 09–13 and 16) to add end-to-end automation using headless `gemini` and `codex exec`, producing machine-readable and human-readable artifacts under `.oraclepack/ticketify/`.

In Scope:
- Step 09: Add headless `gemini` selection that writes `.oraclepack/ticketify/next.json`.
- Step 10: Add non-interactive `codex exec` implementation that consumes `next.json` and writes `.oraclepack/ticketify/codex-implement.md`.
- Step 11: Add verification automation via `codex exec` and/or Gemini diff review:
  - `.oraclepack/ticketify/codex-verify.md` and/or `.oraclepack/ticketify/gemini-review.json`.
- Step 16: Add PR draft automation that writes `.oraclepack/ticketify/PR.md`.
- Include command-availability guards and “skip” behavior as shown in the referenced step snippets (e.g., `command -v ...` checks) to avoid hard failures when tools are missing.

Out of Scope:
- Changing Steps 01–07 semantics (ticket discovery/actions/PRD/Task Master parse/complexity/expand).
- Extending oraclepack’s override injection/validation to cover `codex`/`gemini` (handled in T2).

Current Behavior (Actual):
- Steps 08–20 “are effectively placeholders/notes (echo guidance)” and “don’t dispatch Codex/Gemini or implement code unless the step body explicitly contains those commands.”
- If `codex`/`gemini` are added and the binary is missing → step fails; if interactive → blocks.

Expected Behavior:
- Running the updated steps produces the specified `.oraclepack/ticketify/*` artifacts and enables a headless “select → implement → verify → draft PR” workflow driven by the earlier-generated ticketify outputs.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Must write the artifacts to the paths specified in the step examples (e.g., `.oraclepack/ticketify/next.json`, `codex-implement.md`, `codex-verify.md`, `gemini-review.json`, `PR.md`).
- Must tolerate missing `codex`/`gemini` binaries via skip behavior (per the example snippets).

Evidence:
- “Best insertion points are the placeholder steps… (09–13 and 16).”
- Step outputs in examples:
  - “Wrote .oraclepack/ticketify/next.json”
  - “Wrote .oraclepack/ticketify/codex-implement.md”
  - “Wrote .oraclepack/ticketify/codex-verify.md”
  - “Wrote .oraclepack/ticketify/gemini-review.json”
  - “Wrote .oraclepack/ticketify/PR.md”
- “If the CLI is interactive → it will block waiting for input.”

Open Items / Unknowns:
- Whether Step 11 should use Codex execution, Gemini review, or both by default is not provided.
- Whether Steps 12–13 should be modified (they are within the suggested 09–13 range but specifics are not included in the provided snippets).

Risks / Dependencies:
- Depends on T2 if these steps must participate in oraclepack’s overrides/validation system (otherwise, they run as direct shell commands).
- Risk: tool interactivity can block runs if headless/non-interactive flags are not sufficient (noted in the ticket content).

Acceptance Criteria:
- After running Step 09, `.oraclepack/ticketify/next.json` exists.
- After running Step 10 (with Step 09 completed), `.oraclepack/ticketify/codex-implement.md` exists.
- After running Step 11, at least one of:
  - `.oraclepack/ticketify/codex-verify.md` or `.oraclepack/ticketify/gemini-review.json`
  exists (as configured by the updated pack).
- After running Step 16, `.oraclepack/ticketify/PR.md` exists.
- If `codex` is not available on PATH, Step 10 and any Codex-based Step 11 behavior performs the documented skip behavior (no hard crash beyond what the step defines).
- If `gemini` is not available on PATH, Step 09/16 and any Gemini-based Step 11 behavior performs the documented skip behavior.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Steps 08–20… placeholders/notes… don’t dispatch Codex/Gemini… unless… contains the actual `codex` / `gemini` commands.”
- “Best insertion points… placeholder steps… (09–13 and 16).”
- “If the CLI is interactive → it will block waiting for input.”
```

```ticket T4
T# Title:
- Add “agent-mode” to oraclepack-taskify Action Pack generation (Codex/Gemini path in place of autopilot)

Type:
- enhancement

Target Area:
- oraclepack-taskify Action Pack template/generator (exact file paths not provided)

Summary:
- Extend taskify-generated Action Packs so they can optionally use an “agent-mode” (e.g., `mode=codex` / `mode=gemini`) in the phase after Task Master has expanded tasks (when `tasks.json` and `tm-complexity.json` exist). The ticket text specifies keeping the 20-step contract intact by swapping the existing autopilot entrypoint step slot with agent implementation.

In Scope:
- Add a mode switch for taskify-generated packs (explicitly suggested: `mode=codex` / `mode=gemini`).
- Place the agent-mode insertion “right after Task Master expands tasks” (per ticket content).
- Keep the “20-step contract intact” by swapping “autopilot entrypoint” with “agent implementation” using the same step slot (per ticket content).

Out of Scope:
- Modifying `ticket-action-pack.md` (handled in T3).
- Defining new Task Master workflows beyond what is described.

Current Behavior (Actual):
- Taskify Action Pack can include a “guarded `tm autopilot` entrypoint” (per ticket content).
- Without dispatcher/agent commands in steps, no Codex/Gemini implementation occurs (per ticket content).

Expected Behavior:
- A taskify-generated Action Pack can be generated in an agent-mode that uses Codex/Gemini implementation in the appropriate step slot while retaining the existing step-count/schema contract.

Reproduction Steps:
- Not provided.

Requirements / Constraints:
- Maintain the existing pack schema/contract: “keeping the 20-step contract intact” (per ticket content).
- Agent-mode placement occurs after Task Master expansion (“the point where you have `tasks.json` and `tm-complexity.json`”) (per ticket content).

Evidence:
- “Optional… add an agent-mode to oraclepack-taskify packs… right after Task Master expands tasks… keep the 20-step contract intact.”
- “swap ‘autopilot entrypoint’ with ‘agent implementation’ using the same step slot.”

Open Items / Unknowns:
- How the mode is selected (CLI flag, TUI option, config) is not provided in the included content.
- Exact step number/slot to replace is not provided.

Risks / Dependencies:
- Depends on T2 if Codex/Gemini calls must receive oraclepack overrides/validation (per ticket’s “won’t inherit unless wrap/extend” note).

Acceptance Criteria:
- There is a documented/implemented way to generate a taskify Action Pack in `mode=codex` and/or `mode=gemini`.
- In agent-mode, the pack still conforms to the same step-count contract described in the ticket content (20 steps; autopilot entrypoint swapped rather than expanded beyond contract).
- Agent-mode insertion occurs after Task Master task expansion artifacts exist (as described in the ticket content).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Optional… add an agent-mode to oraclepack-taskify packs… right after Task Master expands tasks… keep the 20-step contract intact.”
- “swap ‘autopilot entrypoint’ with ‘agent implementation’ using the same step slot.”
```
```

.tickets/Oraclepack Pipeline Improvements.md
```
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
```

.tickets/Oraclepack Prompt Generator.md
```
Parent Ticket:

* Title: Create oraclepack-style prompt/skill generator for tickets and .tickets
* Summary:

  * Need a reusable prompt (and/or “skill” template) that can generate an oraclepack-style prompt/skill specifically for “tickets” and/or “.tickets”.
  * Must support the existing placeholder-driven wrapper pattern (e.g., `{user-idea}`, `{project-in-question}`, `{PAIN-POINT}`, `{REFERENCE-FILE}`, `{CAPABILITY}`, `{TARGET-AGENT}`, `{OPTIMIZE-PROMPT}`), including optional fields and “infer from context” behavior as described.
  * Also need guidance on what to change in the current skill and what other viable integration options exist (within the constraints already used in prior work).
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt: “prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets”
* Global Constraints:

  * Optional inputs may be omitted; proceed by inferring from context or requesting missing info within the generated prompt template (“Not always provided…”).
  * “Pain point” is optional; proceed without it if absent.
  * `{REFERENCE-FILE}` may be provided as additional constraints/spec content.
* Global Environment:

  * Unknown
* Global Evidence:

  * Existing wrapper pattern + MCP prompt/tool/resource publication precedent captured in: `/mnt/data/MCP server implementation.md`

Split Plan:

* Coverage Map:

  * Original item: “We need a prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets.”

    * Assigned Ticket ID: T2
  * Original item: “What could we do to our current skill…”

    * Assigned Ticket ID: T3
  * Original item: “…and/or what else are our options for this request?”

    * Assigned Ticket ID: T4
  * Original item: Wrapper placeholders + optionality rules (“Not always provided…”, “Our pain point…”, `{REFERENCE-FILE}`, `{TARGET-AGENT}`, `{CAPABILITY}`, `{OPTIMIZE-PROMPT}`)

    * Assigned Ticket ID: T1
  * Original item: “optimized prompt that will get the {TARGET-AGENT} to find us a solution for adding capabilities…”

    * Assigned Ticket ID: T2
* Dependencies:

  * T2 depends on T1 because the prompt/skill generator must align to the placeholder schema + optionality rules.
  * T3 depends on T2 because “current skill” changes should incorporate the finalized ticket prompt/skill template.
  * T5 depends on T2 and T3 because examples/validation need the final template and integration approach.

````ticket T1
T1 Title: Define ticket/.tickets prompt input schema and placeholder mapping
Type: docs
Target Area: Ticket input model (tickets and/or .tickets) + wrapper placeholders
Summary:
  - Define the canonical set of inputs and placeholders required to generate an oraclepack-style ticket prompt/skill.
  - Preserve the existing wrapper’s rules around optional inputs and context inference.
  - Provide a clear mapping between “tickets/.tickets” fields (if any) and wrapper placeholders without inventing unspecified fields.
In Scope:
  - Enumerate required vs optional placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Document handling rules explicitly stated: optional fields, “infer from context”, and behavior when pain point is absent.
  - Clarify what “tickets” vs “.tickets” means in this system using “Unknown/Not provided” where details are missing.
Out of Scope:
  - Defining new ticket fields beyond what is provided.
  - Implementing tooling or code changes (covered elsewhere).
Current Behavior (Actual):
  - Placeholder set and optionality rules exist in the wrapper pattern, but ticket/.tickets-specific mapping is not defined.
Expected Behavior:
  - A documented, stable mapping that the ticket prompt/skill generator can follow.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not add new placeholders or required fields beyond what is already used in the wrapper.
  - Preserve optionality rules: proceed safely when PAIN-POINT or additional info is absent.
Evidence:
  - Reference wrapper placeholders and prompt-engineer wrapper structure as precedent. (/mnt/data/MCP server implementation.md) :contentReference[oaicite:1]{index=1}
Open Items / Unknowns:
  - Exact structure/format of “tickets” and “.tickets” (not provided).
  - Whether “.tickets” is a file extension, folder convention, or schema name (not provided).
Risks / Dependencies:
  - Risk of mismatch between ticket data shape and placeholder mapping if .tickets format is not standardized.
Acceptance Criteria:
  - A single written spec exists that lists:
    - All placeholders and which are optional.
    - Rules for missing fields (“infer from context” as described).
    - How ticket/.tickets inputs populate placeholders (or explicitly “Unknown” where not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Not always provided, inference from context…”
  - “Our pain point: {PAIN-POINT} … if not just continue…”
  - “```md {REFERENCE-FILE}.md”
````

```ticket T2
T2 Title: Author oraclepack-style prompt/skill template for ticket and .tickets generation
Type: enhancement
Target Area: Prompt/skill template content (oraclepack-style) for tickets/.tickets
Summary:
  - Create the actual prompt/skill template that produces an oraclepack-style prompt/skill when given a ticket or .tickets input.
  - The template must use the existing wrapper structure and placeholders, and must instruct the TARGET-AGENT to generate the desired capability for the project/tool in question.
  - Ensure the template explicitly supports optional inputs and reference-file injection as described.
In Scope:
  - Produce the “ticket prompt-engineer wrapper” template that mirrors the existing wrapper pattern but targets tickets/.tickets.
  - Include all placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Ensure the prompt text includes the “optimized prompt that will get the {TARGET-AGENT}…” requirement, scoped to tickets/.tickets.
Out of Scope:
  - Any new MCP server requirements, tools, or resource URI schemes not explicitly requested for tickets/.tickets.
Current Behavior (Actual):
  - There is no ticket/.tickets-specific oraclepack-style prompt/skill generator template.
Expected Behavior:
  - A single reusable prompt/skill template exists that can be filled with placeholders to drive a TARGET-AGENT to create ticket/.tickets capabilities.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must follow the wrapper’s optionality rules and placeholder usage.
  - Must treat {REFERENCE-FILE} content as “additional constraints/spec” when present.
Evidence:
  - Wrapper structure and placeholder set captured in existing reference prompt material. :contentReference[oaicite:2]{index=2}
Open Items / Unknowns:
  - Where this template will live (file path/naming) in the current repo/tooling (not provided).
Risks / Dependencies:
  - Depends on T1 for a stable placeholder-to-ticket mapping.
Acceptance Criteria:
  - Template includes:
    - All stated placeholders.
    - Explicit instruction to proceed when optional fields are missing.
    - A clearly stated “question to the prompt-engineer: {OPTIMIZE-PROMPT}” section.
  - Template is copy/paste ready and self-contained.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “create an oraclepack prompt/skill but for our tickets and/or .tickets”
  - “optimized prompt that will get the {TARGET-AGENT}… giving it {CAPABILITY}”
  - “Our question to the prompt-engineer: {OPTIMIZE-PROMPT}”
```

```ticket T3
T3 Title: Update current skill to support ticket/.tickets prompt-skill generation
Type: enhancement
Target Area: Existing “current skill” (location/name not provided)
Summary:
  - Identify changes required to the existing skill so it can generate or host the new tickets/.tickets oraclepack-style prompt/skill template.
  - Ensure the current skill can accept the ticket/.tickets inputs and populate the standardized placeholders.
  - Preserve existing behavior for non-ticket use cases (if any), since only ticket support is being added.
In Scope:
  - Incorporate the finalized template from T2 into the current skill workflow.
  - Add/adjust input handling so the current skill can be driven by “tickets and/or .tickets” as the source material.
  - Ensure optional inputs (pain point, additional information, reference file) remain optional in the workflow.
Out of Scope:
  - Designing a brand-new system if the current skill can be extended (unless extension is impossible; not provided).
Current Behavior (Actual):
  - Current skill does not explicitly support generating oraclepack-style prompts/skills for tickets/.tickets (per request).
Expected Behavior:
  - Current skill can produce the tickets/.tickets oraclepack-style prompt/skill using the same wrapper placeholder mechanism.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not remove or break existing skill behavior (implied by “current skill” extension request).
Evidence:
  - “What could we do to our current skill…”
Open Items / Unknowns:
  - Current skill name, file path, and execution context (not provided).
  - How tickets/.tickets are currently stored or passed into the system (not provided).
Risks / Dependencies:
  - Depends on T2 for the template content.
Acceptance Criteria:
  - Current skill supports a ticket/.tickets-driven flow that results in the T2 template with placeholders populated (or explicitly left blank when optional).
  - No regression to existing skill behaviors (validation method not provided; document what was exercised).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “prompt… for our tickets and/or .tickets”
```

```ticket T4
T4 Title: Document integration options for delivering the tickets/.tickets prompt-skill capability
Type: docs
Target Area: Delivery/integration approach (within existing patterns)
Summary:
  - Provide a concise options write-up for how to deliver and reuse the tickets/.tickets prompt/skill generator, aligned to the existing approach patterns already in use.
  - Focus on the two explicitly requested dimensions: changes to the current skill and “other options” for fulfilling the request.
  - Avoid committing to new systems; frame as documented options with constraints and unknowns.
In Scope:
  - Option 1: Extend current skill to include tickets/.tickets support (ties to T3).
  - Option 2: Provide a standalone tickets/.tickets prompt/skill template artifact that can be consumed independently (ties to T2).
  - List constraints/unknowns impacting option choice (e.g., unknown .tickets format, unknown current-skill location).
Out of Scope:
  - Implementing the chosen option (handled by T3 and/or T2).
Current Behavior (Actual):
  - No documented approach exists for how tickets/.tickets prompt/skill generation should be delivered.
Expected Behavior:
  - A short decision-ready document exists describing the options and what each requires, without adding new requirements.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Options must stay within what’s already requested (modify current skill and/or alternative ways to package the prompt/skill).
Evidence:
  - “what else are our options for this request?”
Open Items / Unknowns:
  - Whether the user prefers a single consolidated skill vs multiple dedicated skills (not provided).
Risks / Dependencies:
  - Depends on T1/T2 clarity to accurately describe what each option would deliver.
Acceptance Criteria:
  - Document lists at least:
    - “Modify current skill” option (summary, prerequisites, impact).
    - “Standalone template” option (summary, prerequisites, impact).
    - Explicit unknowns that block a final choice.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “…what else are our options for this request?”
```

````ticket T5
T5 Title: Add examples and validation checks for ticket/.tickets prompt-skill generation
Type: tests
Target Area: Examples + validation of generated prompt/skill output
Summary:
  - Provide concrete example inputs (ticket and/or .tickets) and the expected generated prompt/skill output shape for validation.
  - Ensure examples exercise optional fields (missing PAIN-POINT, missing ADDITIONAL-INFORMATION, with/without REFERENCE-FILE).
  - Add lightweight validation criteria to confirm generated output preserves placeholders and wrapper structure.
In Scope:
  - Example cases covering:
    - Only {user-idea} + {project-in-question}
    - With {PAIN-POINT}
    - With {REFERENCE-FILE}
  - Validation checklist for generated output structure (placeholders present; optional fields handled).
Out of Scope:
  - End-to-end integration tests that require specific repo tooling not provided.
Current Behavior (Actual):
  - No examples/validation for tickets/.tickets prompt-skill generation are defined.
Expected Behavior:
  - Examples exist and can be used to validate that the template and current-skill integration behave as intended.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must preserve the wrapper structure and placeholders as-is.
Evidence:
  - Placeholder and wrapper expectations referenced in the existing wrapper pattern. :contentReference[oaicite:3]{index=3}
Open Items / Unknowns:
  - Exact acceptance mechanism for “validation checks” in the existing system (not provided).
Risks / Dependencies:
  - Depends on T2 (template) and T3 (integration) for meaningful expected outputs.
Acceptance Criteria:
  - At least 3 example inputs exist (covering optionality cases).
  - Each example includes an expected output outline that confirms:
    - Placeholders are present.
    - Optional fields can be omitted without breaking structure.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Not always provided…”
  - “Our pain point: {PAIN-POINT} … continue without needing it.”
  - “```md {REFERENCE-FILE}.md”
````
```

.tickets/Oraclepack TUI Integration.md
```
Parent Ticket:

* Title: Oraclepack TUI integration for PRD Generator Project URL + ticketify PRD run flow
* Summary: Enable oraclepack to route a ticket-derived PRD artifact (`tickets_prd.md`) to a specific “PRD Generator” ChatGPT Project URL via the TUI and automated runs. Address the current failure when attempting to run `tickets_prd.md` as a pack (“no bash code block found”) by introducing supported execution paths (micro-pack and/or single-shot call), and optionally generating a richer PRD context bundle + dedicated PRD-generator pack from ticketify outputs.
* Source:

  * Link/ID: Oraclepack TUI Integration (1).md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “utilize our specific prd-generator from one of many gpt project urls… add that project url… as an option in the tui.”
* Global Constraints:

  * `tickets_prd.md` is content/artifact and cannot be executed as a pack unless wrapped in a valid oraclepack structure (requires a `bash` fenced block).
  * Pack parsing requires a ` ```bash … ``` ` fenced code block; otherwise error: “invalid pack structure: no bash code block found”.
  * Avoid hardcoding ChatGPT Project URLs into packs/repos; prefer selecting/storing via the TUI URL picker/store.
  * Support “one of many project urls” including per-step targeting for PRD generation steps.
  * “Simple oracle calls” should be possible without sending entire multi-step packs.
* Global Environment:

  * Unknown
* Global Evidence:

  * Error string: `Error: invalid pack structure: no bash code block found`.
  * “Projects in ChatGPT” (Projects retain chats/files within a project). ([OpenAI Help Center][1])
  * CommonMark: fenced code blocks support an “info string” after the opening fence (language identifier). ([CommonMark Spec][2])

Split Plan:

* Coverage Map:

  1. Original item: “add that project url and send it with our oraclepack in an automated manner… option in the tui.”
     Assigned Ticket ID: T1
  2. Original item: “simple way of utilizing oracle for simpler calls… do not require entire packs being sent.”
     Assigned Ticket ID: T4
  3. Original item: “oraclepack will not allow the `tickets_prd.md`… `invalid pack structure: no bash code block found`.”
     Assigned Ticket ID: T3
  4. Original item: “Add a new entry in the URL picker… Name: `PRD Generator`… Scope: `project`… or `global`.”
     Assigned Ticket ID: T1
  5. Original item: “Set it as default (`s`)…”
     Assigned Ticket ID: T1
  6. Original item: “Headless/CI: add a CLI flag… `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`).”
     Assigned Ticket ID: T2
  7. Original item: “Using multiple project URLs… `RuntimeOverrides` supports `ChatGPTURL` and `ApplyToSteps`… missing piece is Overrides Wizard UI.”
     Assigned Ticket ID: T5
  8. Original item: “Add a new wizard step: ‘ChatGPT URL’… reuse URLPickerModel… write to `pendingOverrides.ChatGPTURL`.”
     Assigned Ticket ID: T5
  9. Original item: “Option A: ‘micro-pack’… attach `tickets_prd.md`… run `oracle` once…”
     Assigned Ticket ID: T3
  10. Original item: “Option B: add `oraclepack call`… pick URL preset… pick files… run one `oracle` invocation… bypass `internal/pack/parser.go`…”
      Assigned Ticket ID: T4
  11. Original item: “Better idea: `tickets_prd.md` artifact parsed into a valid oraclepack… sent to PRD-generator project url… add missing context as part of stage.”
      Assigned Ticket ID: T6
  12. Original item: “Generate deterministic `prd_context.md`… feature summary, prioritized requirements, user stories + AC, constraints/deps/out-of-scope/risks/open questions, keep vs rewrite.”
      Assigned Ticket ID: T6
  13. Original item: “Generate `.oraclepack/ticketify/prd-generator.pack.md`… attach `tickets_prd.md` + `prd_context.md`… `--write-output ".taskmaster/docs/final_prd.md"`.”
      Assigned Ticket ID: T6
  14. Original item: “Wire it into the TUI… toggle `[ ] Generate enhanced PRD via PRD Generator Project`… prompt pick URL… run pack… apply `RuntimeOverrides{ChatGPTURL: <picked>, ApplyToSteps: {"01": true}}`.”
      Assigned Ticket ID: T7
  15. Original item: “Static context in ChatGPT Project; dynamic context in `prd_context.md` attachment.”
      Assigned Ticket ID: T6
* Dependencies:

  * T7 depends on T6 because the TUI flow runs the generated `prd-generator.pack.md` and attaches `prd_context.md`.
* Split Tickets:

```ticket T1
T# Title:
- Add/select PRD Generator ChatGPT Project URL in TUI URL picker (project/global scope + default)

Type:
- enhancement

Target Area:
- oraclepack TUI URL picker/store (“ChatGPT Project URLs”)

Summary:
- Provide a standard way to store and select a “PRD Generator” ChatGPT Project URL from the existing URL picker/store, with project vs global scope guidance and the ability to set it as the default. This enables routing PRD-generation runs through the intended ChatGPT Project without hardcoding URLs into packs.

In Scope:
- Ensure the TUI supports creating/selecting an entry named `PRD Generator` with a ChatGPT Project URL.
- Support scope selection: `project` (repo-specific) vs `global` (shared), as described.
- Support setting the selected entry as the default (`s`) for the relevant scope.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Not provided

Expected Behavior:
- User can add/select `PRD Generator` as a named URL entry.
- User can set it as default for the current workflow/repo.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Do not hardcode URLs into packs/repos (prefer picker/store selection).
- Preserve support for project vs global URL scoping.

Evidence:
- “Add a new entry in the URL picker… Name: `PRD Generator`… Scope: `project`… `global`…”
- “Set it as default (`s`)…”

Open Items / Unknowns:
- Actual PRD Generator ChatGPT Project URL value(s) (not provided).
- Whether a predefined/seeded entry is required vs user-created entry (not provided).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] TUI allows creating/selecting a `PRD Generator` URL entry.
- [ ] TUI allows choosing `project` vs `global` scope for the entry.
- [ ] TUI allows setting the entry as default (`s`) and that default is subsequently used when selected.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Add a new entry in the URL picker: … Name: `PRD Generator`”
- “Scope: `project`… `global`…”
- “Set it as default (`s`)…”
```

```ticket T2
T# Title:
- Add headless/CI CLI override for ChatGPT URL selection during runs

Type:
- enhancement

Target Area:
- oraclepack CLI run command (headless / `--no-tui` runs)

Summary:
- Add a CLI override so headless/CI executions can force a ChatGPT URL (or saved URL name) rather than relying on interactive selection. This supports automated routing to the PRD Generator project URL.

In Scope:
- Add CLI support equivalent to: `oraclepack run --chatgpt-url <url>` and/or `--chatgpt-url-name <saved-name>`.
- Ensure the provided value sets the runner’s ChatGPT URL for the run.

Out of Scope:
- Not provided

Current Behavior (Actual):
- “Right now, the TUI resolves a URL; but there isn’t a CLI flag … that forces it in `--no-tui` runs.”

Expected Behavior:
- Headless/CI runs can specify a ChatGPT URL (or saved name) and have it applied for the run.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must work without requiring the interactive TUI URL picker.
- Must not require editing packs to include URLs.

Evidence:
- “Headless/CI: add a CLI flag to override the picker”
- “Add `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`)...”

Open Items / Unknowns:
- Exact CLI flag naming/shape (both `--chatgpt-url` and `--chatgpt-url-name` are suggested; final selection not provided).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] `oraclepack run --chatgpt-url <url>` applies the provided URL for the run.
- [ ] If `--chatgpt-url-name <saved-name>` is implemented, it resolves to a stored URL and applies it for the run.
- [ ] Works in non-interactive/headless mode.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Headless/CI: add a CLI flag to override the picker”
- “Add `oraclepack run --chatgpt-url <url>` (or `--chatgpt-url-name <saved-name>`)...”
```

````ticket T3
T# Title:
- Provide a micro-pack wrapper to run PRD generation using tickets_prd.md (valid bash fence pack)

Type:
- enhancement

Target Area:
- oraclepack pack inputs (micro-pack file used to call `oracle` with `tickets_prd.md`)

Summary:
- Introduce a minimal, valid oraclepack “micro-pack” that wraps PRD generation as an executable pack step. This avoids attempting to run `tickets_prd.md` directly (which fails pack validation) while enabling attaching `tickets_prd.md` to a single `oracle` call.

In Scope:
- Provide a one-step (or 2–3 step) pack that:
  - Is valid for oraclepack parsing (contains a `bash` fenced code block).
  - Attaches `tickets_prd.md`.
  - Runs `oracle` once to generate a PRD output.
- Ensure this flow can be routed to the selected PRD Generator ChatGPT Project URL (via existing URL selection/overrides mechanism).

Out of Scope:
- Full “single-shot call” mode that bypasses pack parsing (see T4).

Current Behavior (Actual):
- Running `tickets_prd.md` as a pack fails with: “Error: invalid pack structure: no bash code block found”.

Expected Behavior:
- A micro-pack wrapper can be run successfully by oraclepack and performs the PRD-generation oracle call using `tickets_prd.md` as an attachment.

Reproduction Steps:
1) Attempt to run `tickets_prd.md` as a pack and observe “no bash code block found”.
2) Run the micro-pack wrapper pack and observe it parses and executes.

Requirements / Constraints:
- Pack must include a ` ```bash … ``` ` fenced block to be parseable.
- Must attach `tickets_prd.md` to the oracle invocation.

Evidence:
- Error: “invalid pack structure: no bash code block found”
- “Generate a 1-step pack… attach `tickets_prd.md`… run `oracle` once…”

Open Items / Unknowns:
- Final location/name of the micro-pack file (example name provided: `prd-generator-call.pack.md`).
- Exact prompt text and output path conventions for the oracle call (not fully specified).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] A micro-pack file exists and is parseable by oraclepack (contains a `bash` fenced block).
- [ ] Running the micro-pack executes a single `oracle` call that includes `tickets_prd.md` as an attachment.
- [ ] The flow does not require running `tickets_prd.md` directly as a pack.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Generate a 1-step pack… attach `tickets_prd.md`…”
- “Create a new file… valid pack… run `oracle` while attaching `tickets_prd.md`.”
- “Error: invalid pack structure: no bash code block found”
````

```ticket T4
T# Title:
- Add single-shot oracle invocation mode (CLI/TUI) that bypasses pack parsing

Type:
- enhancement

Target Area:
- oraclepack CLI + TUI (“single-call” mode)

Summary:
- Add a new execution path for “simple oracle calls” that does not require a full pack file or pack parsing. The user can select a ChatGPT URL preset, attach files such as `tickets_prd.md`, provide a prompt/template, and run exactly one `oracle` invocation.

In Scope:
- Implement a new subcommand such as `oraclepack call` (or `oraclepack oracle`) that:
  - Lets the user pick a ChatGPT URL preset.
  - Lets the user specify attachments (e.g., `tickets_prd.md`).
  - Runs one `oracle …` invocation.
- Implement a corresponding TUI flow/screen for “Single Oracle Call” with:
  - URL preset selection
  - attachments selection
  - prompt/template input
  - run
- Ensure this path bypasses pack parsing requirements (no need for a `bash` fenced block).

Out of Scope:
- Generating ticket-derived context bundle (`prd_context.md`) and a dedicated PRD generator pack (see T6/T7).

Current Behavior (Actual):
- “Simple calls” currently imply using packs; running `tickets_prd.md` directly fails pack validation.

Expected Behavior:
- Users can perform a single oracle call via oraclepack without needing a pack file format.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must bypass `internal/pack/parser.go` (as described) so it does not require a `bashFenceRegex`-parseable pack structure.
- Must support selecting a ChatGPT URL preset and attaching files.

Evidence:
- “Option B (best UX): add a new CLI/TUI mode for single-shot calls”
- “Add a subcommand like: `oraclepack call`… pick ChatGPT URL preset… files to attach…”

Open Items / Unknowns:
- Exact CLI UX (flags for attachments, prompt, output path) not fully specified.
- Whether output-writing is required (“--write-output …” is shown in examples elsewhere but not mandated here).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] `oraclepack call` (or equivalent) runs a single oracle invocation without requiring a pack file.
- [ ] User can select a ChatGPT URL preset and attach `tickets_prd.md`.
- [ ] A TUI “Single Oracle Call” flow exists with URL selection, attachments selection, prompt/template entry, and run.
- [ ] Running this path does not trigger “no bash code block found” pack-structure errors.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Option B (best UX): add a new CLI/TUI mode for single-shot calls”
- “Add a subcommand like: `oraclepack call` (or `oraclepack oracle`)”
- “This avoids the structural requirement that triggers your error…”
```

```ticket T5
T# Title:
- Extend Overrides Wizard to support per-step ChatGPT URL selection (RuntimeOverrides.ChatGPTURL)

Type:
- enhancement

Target Area:
- oraclepack TUI Overrides Wizard UI

Summary:
- Enable selecting a ChatGPT URL in the Overrides Wizard and applying it to specific steps via per-step targeting. This supports “one of many project urls” where only PRD generation steps use the PRD Generator project while others use the default.

In Scope:
- Add a wizard step: “ChatGPT URL”.
- Reuse the existing URL picker/store UI model (URLPickerModel) to choose a URL.
- Store the selection into `pendingOverrides.ChatGPTURL`.
- Ensure per-step targeting continues to work via `ApplyToSteps`, so only selected steps get the override.

Out of Scope:
- Not provided

Current Behavior (Actual):
- Overrides Wizard “currently picks flags + target steps, but doesn’t let you pick a ChatGPT URL to apply for those targeted steps.”

Expected Behavior:
- Overrides Wizard allows selecting a ChatGPT URL override and applying it only to selected steps.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must integrate with existing `RuntimeOverrides` structure (supports `ChatGPTURL` and `ApplyToSteps`).
- Must not force a single global URL for the entire run when per-step targeting is desired.

Evidence:
- “`RuntimeOverrides` supports `ChatGPTURL` and `ApplyToSteps` targeting.”
- “Add a new wizard step: ‘ChatGPT URL’… Write the chosen value into `pendingOverrides.ChatGPTURL`.”

Open Items / Unknowns:
- Exact UX flow placement/order in the wizard (not provided).
- How conflicts are resolved between default URL and step override when both are present (not provided).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] Overrides Wizard includes a “ChatGPT URL” selection step.
- [ ] Selected URL is stored in `pendingOverrides.ChatGPTURL`.
- [ ] When overrides are applied to a subset of steps, only those steps use the overridden ChatGPT URL; other steps use the default.
- [ ] Existing step-targeting selection remains functional.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “`RuntimeOverrides` supports `ChatGPTURL` … `ApplyToSteps` targeting.”
- “Add a new wizard step: **‘ChatGPT URL’**”
- “Write the chosen value into `pendingOverrides.ChatGPTURL`.”
```

````ticket T6
T# Title:
- Enhance ticketify outputs: generate prd_context.md + prd-generator.pack.md for PRD Generator Project

Type:
- enhancement

Target Area:
- oraclepack-ticketify outputs/artifacts generation

Summary:
- Generate a deterministic PRD context bundle alongside `tickets_prd.md`, then generate a dedicated, valid micro-pack that calls `oracle` against the PRD Generator ChatGPT Project URL using both `tickets_prd.md` and the new context artifact. This supplies the missing context needed for higher-quality PRD generation while keeping the PRD Generator’s stable instructions in the ChatGPT Project.

In Scope:
- Generate `.oraclepack/ticketify/prd_context.md` containing ticket-derived context, including:
  - product/feature summary inferred from tickets
  - prioritized requirements (functional + non-functional)
  - user stories + acceptance criteria extracted from tickets
  - constraints, dependencies, out-of-scope, risks, open questions
  - explicit “what to keep vs rewrite” instructions for the PRD generator
- Generate `.oraclepack/ticketify/prd-generator.pack.md` that:
  - Is a valid pack with a `bash` fenced code block.
  - Attaches BOTH `tickets_prd.md` and `prd_context.md`.
  - Calls `oracle … --write-output ".taskmaster/docs/final_prd.md"` (as shown in the example).
- Preserve the intended split of context:
  - Static context lives in the PRD Generator ChatGPT Project.
  - Dynamic context is attached via `prd_context.md`.

Out of Scope:
- Not provided

Current Behavior (Actual):
- `tickets_prd.md` alone “does not include the context it would need” for the PRD generator.

Expected Behavior:
- ticketify produces `prd_context.md` and `prd-generator.pack.md` so PRD generation can run with complete, deterministic inputs.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Outputs should be deterministic and ticket-derived (“stable ordering/determinism” is implied by “deterministic… ticket-derived”).
- Generated PRD pack must be valid for oraclepack parsing (contains ` ```bash … ``` `).
- Avoid hardcoding ChatGPT Project URLs into the generated pack; selection/override handled externally (TUI/overrides).

Evidence:
- “Generate a deterministic ‘PRD context bundle’ artifact…”
- “Emit a second markdown file… `.oraclepack/ticketify/prd-generator.pack.md`”
- “Attach BOTH… `tickets_prd.md`… `prd_context.md`… `--write-output ".taskmaster/docs/final_prd.md"`”

Open Items / Unknowns:
- Exact prompt content to send to PRD Generator project (not fully specified).
- Exact source inputs from ticketify flow used to derive `prd_context.md` (“index/actions/etc.” referenced but not enumerated).

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- [ ] ticketify produces `.oraclepack/ticketify/prd_context.md` with the listed context sections/bullets.
- [ ] ticketify produces `.oraclepack/ticketify/prd-generator.pack.md` that is parseable by oraclepack (contains a `bash` fence).
- [ ] The generated pack attaches both `tickets_prd.md` and `prd_context.md`.
- [ ] The generated pack writes output to `.taskmaster/docs/final_prd.md` (as specified in the example).
- [ ] Generated artifacts are deterministic (same ticket inputs produce stable output ordering/content structure).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Generate a deterministic ‘PRD context bundle’ artifact”
- “Emit a second markdown file… `.oraclepack/ticketify/prd-generator.pack.md`”
- “Attach BOTH… `tickets_prd.md`… `prd_context.md`”
````

```ticket T7
T# Title:
- Add TUI option to run ticketify → PRD Generator flow using selected ChatGPT Project URL

Type:
- enhancement

Target Area:
- oraclepack TUI (new toggle/flow for PRD generation)

Summary:
- Add a TUI toggle/option to run the generated PRD generator pack and route it to the PRD Generator ChatGPT Project URL. The TUI should prompt for (or auto-select) the PRD Generator URL entry, execute `prd-generator.pack.md`, and apply per-step ChatGPT URL overrides to the PRD generation step.

In Scope:
- Add a TUI toggle like: `[ ] Generate enhanced PRD via PRD Generator Project`.
- When enabled:
  1) Prompt user to pick the PRD-generator URL from the existing URL picker (or auto-select “PRD Generator” if present).
  2) Run the generated `prd-generator.pack.md`.
  3) Apply `RuntimeOverrides{ChatGPTURL: <picked>, ApplyToSteps: {"01": true}}` (or whichever step calls oracle).

Out of Scope:
- Not provided

Current Behavior (Actual):
- Not provided

Expected Behavior:
- TUI provides an explicit entry point for the PRD Generator run that applies the correct ChatGPT URL specifically to PRD generation step(s).

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must use existing URL picker/store (no hardcoded URL in pack/repo).
- Must apply ChatGPT URL override to targeted PRD step(s), not necessarily globally.

Evidence:
- “Add a TUI toggle like: `[ ] Generate enhanced PRD via PRD Generator Project`”
- “prompt user to pick the PRD-generator URL… run `prd-generator.pack.md`… apply `RuntimeOverrides{ChatGPTURL: <picked>, ApplyToSteps: {"01": true}}`…”

Open Items / Unknowns:
- Whether step ID is always `01` or determined dynamically (“or whichever step calls oracle”).
- Whether auto-selecting “PRD Generator” is required behavior or best-effort fallback.

Risks / Dependencies:
- Depends on ticketify generating `prd-generator.pack.md` and `prd_context.md` (T6).

Acceptance Criteria:
- [ ] TUI displays a toggle/option labeled as specified (or equivalent).
- [ ] When enabled, user can select (or auto-select) the PRD Generator ChatGPT URL entry.
- [ ] TUI runs `prd-generator.pack.md`.
- [ ] The PRD generation step(s) run with the selected ChatGPT URL override applied via `ApplyToSteps`.
- [ ] Non-PRD steps (if any) are not forced to use the PRD Generator URL unless targeted.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “[ ] Generate enhanced PRD via PRD Generator Project”
- “prompt user to pick the PRD-generator URL… run… apply `RuntimeOverrides{ChatGPTURL: <picked>, ApplyToSteps: {"01": true}}`”
- “without hardcoding URLs into packs or repos…”
```

[1]: https://help.openai.com/en/articles/10169521-using-projects-in-chatgpt?utm_source=chatgpt.com "Projects in ChatGPT"
[2]: https://spec.commonmark.org/0.12/?utm_source=chatgpt.com "Fenced code blocks"
```

.tickets/Oraclepack Workflow Enhancement.md
```
Title:

* Stabilize oraclepack “oracle-pack” pipeline and add profile-based context + Stage-3 synthesis for actionable follow-through

Summary:

* The current two-step workflow generates an `oracle-pack` Markdown file (20 `oracle` calls) via Codex skills/Gemini CLI commands, then runs it through the `oraclepack` Go wrapper to produce outputs and state/report artifacts.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* A key failure mode is schema/format drift in the pack file (human-doc + machine-ingest combined), which can break ingestion; an example drift is step headers using an em dash (`# 01 — ROI=…`) while the documented contract expects `# NN)`.

    Workflow Optimization for Oracl…

* Requested outcome: improve workflow continuity, enable richer context injection without breaking the strict pack contract, and add an automatic next stage that turns the “final twenty questions/answers” into actionable implementation steps with minimal human interaction.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Background / Context:

* Workflow stages:

  * Stage 1: LLM authoring creates `docs/oracle-pack-YYYY-MM-DD.md` containing 20 `oracle` commands (with ROI metadata and per-step output paths).

        Workflow Optimization for Oracl…

  * Stage 2: `oraclepack` executes the pack to produce 20 outputs under `oracle-out/...` plus state/report JSON artifacts.

        Workflow Optimization for Oracl…

* `oraclepack` is a wrapper around `oracle` intended for batched/bulk requests.

    Workflow Optimization for Oracl…

* Core concern: “disconnection” after the 20-question output; desire to chain into a useful, actionable implementation stage.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Current Behavior (Actual):

* Pack file acts as both:

  * Human documentation, and

  * A strict machine-ingest contract, making formatting drift a pipeline-breaking event.

        Workflow Optimization for Oracl…

* Documented/expected step-header schema (`# NN)`) can drift to alternative formats (example: `# 01 — ROI=…`), risking parse/validation failures.

    Workflow Optimization for Oracl…

* High-risk edits include adding additional code fences (especially additional bash fences) or introducing lines that accidentally match the step-header pattern.

    Workflow Optimization for Oracl…

Expected Behavior:

* Packs remain schema-stable and reliably parse/validate/run across providers (Codex skills, Gemini CLI commands).

    Workflow Optimization for Oracl…

* Richer “skill context” can be injected without changing the pack’s ingest shape (no added code fences / no header drift).

    Workflow Optimization for Oracl…

* After Stage 2 produces 20 outputs + report JSON, a subsequent stage can automatically convert results into actionable implementation steps.

    Workflow Optimization for Oracl…

Requirements:

* Preserve the non-negotiable pack contract:

  * Pack is Markdown containing exactly one `bash` code block; the first bash block is executed.

        Workflow Optimization for Oracl…

  * Steps are identified via header pattern `# NN)` with sequential numbering starting at `01`.

        Workflow Optimization for Oracl…

  * Prelude content before the first step header executes once.

        Workflow Optimization for Oracl…

* Standardize Stage-1 generation to the strict header form `# NN)` (avoid em dash variants).

    Workflow Optimization for Oracl…

* Add a hard gate between Stage 1 and Stage 2:

  * Make `oraclepack validate` mandatory before `oraclepack run` (prevent schema drift reaching execution).

        Workflow Optimization for Oracl…

* Provide schema-safe extensibility for context injection:

  * Allow context to be injected via `oracle -p` prompt text and/or `oracle -f` file/directory attachments (preferred for larger context).

        Workflow Optimization for Oracl…

  * Use prelude variables and templating only if it does not interfere with header parsing.

        Workflow Optimization for Oracl…

  * Avoid adding extra code fences or lines resembling step headers.

        Workflow Optimization for Oracl…

* Implement “Context Profiles” as file-backed bundles:

  * Add `skills/oracle-pack/references/profiles/<name>.md` and inject via `oracle -f "$profile_file"` plus a short prompt preamble (“Follow the attached profile standards”).

  * Add an optional `profile` input to the Stage-1 skill/command, with backwards-compatible behavior when absent.

        Workflow Optimization for Oracl…

* Add a first-class Stage 3 synthesis step:

  * Provide a command shape such as `oraclepack synthesize --in oracle-out --report pack.report.json --out docs/implementation-pack-YYYY-MM-DD.md` that reads the 20 outputs, extracts proposed changes/file targets/tests, and emits a new validated pack for implementation.

        Workflow Optimization for Oracl…

  * Support minimal-interaction execution for Stage 3 (e.g., headless usage via Codex/Gemini CLI).

        Workflow Optimization for Oracl…

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* `oraclepack` Go program wrapping `oracle` CLI.

* Stage-1 generation tools mentioned: Codex skills, Gemini CLI commands.

    Workflow Optimization for Oracl…

* OS/CI details: Unknown.

Evidence:

* Attachment: “Workflow Optimization for Oraclepack.md”.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* Example schema drift called out: step headers using `# 01 — ROI=…` vs documented `# NN)`.

    Workflow Optimization for Oracl…

* Proposed validation/run sequence:

  * `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack list docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --no-tui --run-all --stop-on-fail=true --out-dir oracle-out`

        Workflow Optimization for Oracl…

Decisions / Agreements:

* Treat the pack as a stable intermediate representation (IR) and keep context flowing only through `-p` and `-f` to avoid breaking the ingest contract.

* Prefer “Context Profiles” as a file-backed mechanism located under `skills/oracle-pack/references/profiles/`.

* Add a validation gate (`validate` before `run`) to reduce pipeline breakage from formatting drift.

    Workflow Optimization for Oracl…

Open Items / Unknowns:

* Exact current parser/validator behavior regarding em dash header variants (whether it currently accepts them) is not provided; only that it is avoidable schema drift.

    Workflow Optimization for Oracl…

* Exact filenames/paths for current `SKILL.md` and template files in the repo are referenced conceptually but not provided in full.

    Workflow Optimization for Oracl…

* Whether `oraclepack synthesize` already exists or is a new feature request is not provided; it is described as a proposed product shape.

    Workflow Optimization for Oracl…

Risks / Dependencies:

* Dependency on `oracle` CLI flags and behavior (`-p/--prompt`, `-f/--file`, `--write-output`, `--files-report`, `--dry-run`).

    Workflow Optimization for Oracl…

* Risk of pack invalidation from added code fences, additional bash blocks, or accidental header-like lines.

    Workflow Optimization for Oracl…

* Cross-provider consistency risk (Codex skills vs Gemini CLI commands) unless Stage 1 is standardized around a shared template/profile mechanism.

    Workflow Optimization for Oracl…

Acceptance Criteria:

* Pack schema stability

  * Packs validate when they contain exactly one bash block and step headers are strictly `# NN)` starting at `01`.

  * Stage-1 generation output uses `# NN)` (no em dash variant) across providers.

        Workflow Optimization for Oracl…

* Validation gate

  * Workflow includes a required `oraclepack validate` pass before any `oraclepack run`.

        Workflow Optimization for Oracl…

* Context Profiles

  * A `profile` selection results in `oracle -f "$profile_file"` being attached per step without adding new code fences or breaking parsing.

        Workflow Optimization for Oracl…

  * Absence of `profile` preserves current behavior (backwards compatible).

        Workflow Optimization for Oracl…

* Stage 3 synthesis

  * A synthesis step can consume `oracle-out` outputs + report JSON and emit a follow-on implementation pack intended to be validated and run.

        Workflow Optimization for Oracl…

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* workflow

* cli

* parsing

* validation

* context-bundles

* automation

---
```

.tickets/Verbose Payload Rendering TUI.md
```
Title:

* Add verbose payload rendering in TUI to display full prepared scripts/flags for oraclepack runs

Summary:

* The TUI should support a verbose mode that prints the full “prepared payload” being executed for oraclepack runs, including effective flags (post overrides and `--chatgpt-url` injection) and the entire script passed to execution.

    Verbose TUI Payload Rendering

* This is needed to verify exactly what payloads are being sent/executed during oraclepack runs and to support tests that confirm the rendered payload contents.

    Verbose TUI Payload Rendering

Background / Context:

* Proposed approach: add a reusable “prepared payload” layer to `internal/exec.Runner` (prepare prelude/step scripts after overrides + flag injection + sanitization), then have the TUI emit these prepared payload blocks to its log viewport immediately before execution.

    Verbose TUI Payload Rendering

* Files implicated by the proposal include `internal/exec/runner.go`, `internal/tui/tui.go`, `internal/cli/run.go`, plus new helpers/tests under `internal/tui/` and `internal/exec/`.

    Verbose TUI Payload Rendering

Current Behavior (Actual):

* The TUI does not provide a verbose rendering that shows the full prepared payload (full script + effective flags + extracted `oracle …` invocations) being executed for oraclepack runs.

    Verbose TUI Payload Rendering

Expected Behavior:

* When verbose payload logging is enabled, the TUI log viewport prints a payload block before each step runs that includes: effective oracle flags, extracted oracle invocations (full lines), and the full prepared script that will be executed.

    Verbose TUI Payload Rendering

* Verbose payload logging can be enabled via CLI flag (e.g., `--verbose-payload` with `-v`) and toggled in the TUI via a keybinding (proposed: `p`).

    Verbose TUI Payload Rendering

Requirements:

* Exec runner: expose “prepared payload” APIs:

  * `PreparePreludePayload(p *pack.Prelude) PreparedPreludePayload`

  * `PrepareStepPayload(s *pack.Step) PreparedStepPayload`

  * `RunPreparedPrelude(...)` / `RunPreparedStep(...)` to execute the prepared scripts.

        Verbose TUI Payload Rendering

* Prepared payload structures must include:

  * `Script` (sanitized, post injection),

  * `EffectiveFlags` (for steps),

  * `OracleInvocations` extracted from the prepared script,

  * sanitizer `Warnings`.

        Verbose TUI Payload Rendering

* TUI formatting helper:

  * Add `internal/tui/verbose_payload.go` to format payload blocks (header, effective flags, oracle invocations, then full script).

        Verbose TUI Payload Rendering

* TUI integration:

  * Add a `verbosePayload bool` toggle to the TUI model.

  * In the run flow, call `PrepareStepPayload` and, when enabled, push formatted payload lines into `logChan` before `RunPreparedStep`.

  * Add keybinding `p` to toggle `verbosePayload`.

        Verbose TUI Payload Rendering

* CLI wiring:

  * Add `--verbose-payload` / `-v` flag and pass it into `tui.NewModel(..., verbosePayload)`.

        Verbose TUI Payload Rendering

* Tests:

  * New `internal/exec/runner_payload_test.go` verifying prepared payload includes effective flags and injected oracle command text.

  * New `internal/tui/verbose_payload_test.go` verifying formatted lines include flags, invocation, and script content.

  * Update existing TUI tests to include the new `NewModel` arg.

        Verbose TUI Payload Rendering

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Go (per referenced `.go` files and packages).

    Verbose TUI Payload Rendering

* TUI framework: Bubble Tea (`tea.NewProgram(...)` referenced).

    Verbose TUI Payload Rendering

* OS / terminal / versions: Unknown.

Evidence:

* Proposed change list and implementation sketch in: /mnt/data/Verbose TUI Payload Rendering.md

    Verbose TUI Payload Rendering

* Proposed file tree changes:

  * `internal/exec/runner.go` (modify)

  * `internal/exec/runner_payload_test.go` (new)

  * `internal/tui/verbose_payload.go` (new)

  * `internal/tui/verbose_payload_test.go` (new)

  * `internal/tui/tui.go` (modify)

  * `internal/cli/run.go` (modify)

  * Update TUI tests to pass new model arg.

        Verbose TUI Payload Rendering

Decisions / Agreements:

* Adopt a “prepared payload” abstraction in `exec.Runner` to ensure the TUI logs exactly what will run after overrides, injection, and sanitization.

    Verbose TUI Payload Rendering

* Add both CLI flag control (`--verbose-payload` / `-v`) and an in-TUI toggle (proposed key: `p`).

    Verbose TUI Payload Rendering

Open Items / Unknowns:

* Exact existing TUI run flow for prelude execution (whether/where prelude runs in TUI) is not provided; proposal notes “if you also execute a prelude… do the same.”

    Verbose TUI Payload Rendering

* Exact current `NewModel(...)` signature call sites and all affected tests/files beyond those listed are not fully enumerated (some are referenced as examples).

    Verbose TUI Payload Rendering

Risks / Dependencies:

* Not provided.

Acceptance Criteria:

* Running the TUI with `--verbose-payload` causes each executed step to prepend a log block that includes:

  * “payload (step <id>)” header,

  * “effective oracle flags: …” line,

  * extracted “oracle invocations:” section (or explicit none found),

  * full “script:” content (not truncated),

  * “end payload” footer.

        Verbose TUI Payload Rendering

* Toggling `p` in the TUI flips payload logging on/off for subsequent step executions.

    Verbose TUI Payload Rendering

* `Runner.PrepareStepPayload` produces:

  * effective flags reflecting overrides and `--chatgpt-url`,

  * a prepared script containing the injected oracle invocation with those flags.

        Verbose TUI Payload Rendering

* New tests (`runner_payload_test.go`, `verbose_payload_test.go`) pass, and existing TUI tests compile and pass after updating `NewModel` call signature.

    Verbose TUI Payload Rendering

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* logging

* exec-runner

* cli

* testing

* go

---
```

.tickets/mcp/Expose Oraclepack as MCP.md
```
Parent Ticket:

* Title: Expose oraclepack as MCP tools (with Taskify Stage-2/Stage-3 helpers)
* Summary: Provide an MCP server that exposes `oraclepack` CLI capabilities (validate/list/run) plus helper tools for Stage-2 detection/validation and Stage-3 action-pack validation/execution/artifact summarization, with secure-by-default controls (allowlisted filesystem roots, execution gating, timeouts, truncation) and support for stdio + streamable-http transports.
* Source:

  * Link/ID (if present) or “Not provided”: /mnt/data/MCP tools for Oraclepack.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Expose `oraclepack` … as **MCP tools**, so an agent can … run packs … validate Stage-2 … validate Stage-3 … summarize artifacts.”
* Global Constraints:

  * Support MCP transports: “stdio” and “streamable-http”.
  * Security defaults: “deny-by-default execution”, “allowlisted roots”, “stdout/stderr truncation and timeouts”.
* Global Environment:

  * Unknown
* Global Evidence:

  * MCP tool list and env var config list.
  * Reference implementation structure (README/requirements and Python modules).

Split Plan:

* Coverage Map:

  * “Expose `oraclepack` (validate/list/run) … as **MCP tools**” → T1
  * “run packs non-interactively (`--no-tui --yes --run-all`)” → T5
  * “validate Stage-2 outputs (01-*.md..20-*.md)” → T4
  * “validate Stage-3 Action Packs (single ```bash fence, step headers…)” → T7
  * “summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)” → T7
  * “Tools: oraclepack_validate_pack / oraclepack_list_steps / oraclepack_run_pack …” → T5
  * “Tools: oraclepack_read_file …” → T5
  * “Tools: … taskify_detect_stage2 / taskify_validate_stage2 …” → T5
  * “Tools: … taskify_validate_action_pack / taskify_artifacts_summary …” → T5
  * “Tools: … taskify_run_action_pack …” → T5
  * “Transports: stdio … streamable-http …” → T6
  * “Tool annotations: readOnlyHint / destructiveHint / openWorldHint …” → T6
  * “Security defaults: ORACLEPACK_ENABLE_EXEC=1 gating …” → T2
  * “Security defaults: allowlisted filesystem roots …” → T2
  * “Security defaults: truncation and timeouts …” → T3
  * “Env vars: ORACLEPACK_ALLOWED_ROOTS / BIN / WORKDIR / ENABLE_EXEC / CHARACTER_LIMIT / MAX_READ_BYTES” → T2
  * “Typical Stage-3 usage: detect/validate → validate action pack → run → summarize” → Info-only
  * “Reference implementation tree (README, requirements.txt, modules list)” → Info-only
  * Links to MCP specs / python-sdk repo mentioned → Info-only
* Dependencies:

  * T5 depends on T2 because server tools must enforce allowed roots and execution gating.
  * T5 depends on T3 because `oraclepack_*run*` tools need subprocess execution with timeouts/truncation.
  * T5 depends on T4 because `oraclepack_taskify_*stage2*` tools call Stage-2 detection/validation.
  * T5 depends on T7 because `oraclepack_taskify_*action_pack*` tools call action-pack validation/summarization helpers.
  * T6 depends on T5 because annotations/transport-hardening apply to the MCP server surface.
* Split Tickets:

```ticket T1
T# Title: Scaffold oraclepack MCP server project (README + packaging entrypoints)
Type: chore
Target Area: oraclepack-mcp-server repo scaffolding (README.md, requirements.txt, __init__.py, __main__.py)
Summary:
- Create the MCP server project structure that exposes `oraclepack` + Taskify helpers as MCP tools, including installation and run instructions and the tool list.
- Ensure the package has an executable entrypoint to start the MCP server with selectable transport.
In Scope:
- Create/maintain project tree with:
  - README describing purpose, install, configuration env vars, run modes, tools list, and typical Stage-3 usage.
  - requirements.txt listing dependencies.
  - Python package layout with `oraclepack_mcp_server/__init__.py` and `oraclepack_mcp_server/__main__.py`.
- CLI args in `__main__.py` to accept `--transport` with choices `stdio` and `streamable-http`.
Out of Scope:
- Implementing tool logic (handled in other tickets).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- A runnable package that starts an MCP server with a chosen transport.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support `--transport` choices: `stdio`, `streamable-http`.
Evidence:
- Project tree and entrypoint snippet showing `choices=["stdio", "streamable-http"]`. :contentReference[oaicite:26]{index=26}
Open Items / Unknowns:
- Exact repository root / packaging approach (pip package vs repo-local module): Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Repository includes README.md, requirements.txt, and `oraclepack_mcp_server` package directory.
- [ ] Running `python -m oraclepack_mcp_server --transport stdio` is supported (starts server process).
- [ ] Running `python -m oraclepack_mcp_server --transport streamable-http` is supported (starts server process).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers) … tree … README.md … requirements.txt … __main__.py” :contentReference[oaicite:27]{index=27}
- “choices=[‘stdio’, ‘streamable-http’] … mcp.run(transport=args.transport)” :contentReference[oaicite:28]{index=28}
```

```ticket T2
T# Title: Implement config + filesystem security controls (allowed roots, exec gating, max read bytes)
Type: chore
Target Area: oraclepack_mcp_server/config.py and oraclepack_mcp_server/security.py
Summary:
- Implement secure-by-default configuration for the MCP server, driven by environment variables, including allowlisted filesystem roots and explicit execution enablement.
- Provide safe path resolution under allowed roots and bounded file reads for tool operations.
In Scope:
- Env-driven config including:
  - `ORACLEPACK_ALLOWED_ROOTS` (colon-separated roots)
  - `ORACLEPACK_BIN`
  - `ORACLEPACK_WORKDIR`
  - `ORACLEPACK_ENABLE_EXEC`
  - `ORACLEPACK_CHARACTER_LIMIT`
  - `ORACLEPACK_MAX_READ_BYTES`
- Path allowlisting:
  - Resolve requested file paths and ensure they live under at least one allowed root.
  - Raise an explicit “path not allowed” error on violation.
- Safe file reads:
  - Read text/bytes with max size enforcement and “truncated” indicator.
Out of Scope:
- Subprocess execution and output truncation (handled in T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Server loads config from env and enforces filesystem access boundaries for read tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Execution must be deny-by-default unless `ORACLEPACK_ENABLE_EXEC=1`.
- Filesystem access must be restricted to allowlisted roots.
Evidence:
- Env var list and semantics. :contentReference[oaicite:29]{index=29}
- Security guidance: “deny-by-default execution … allowlisted roots”. :contentReference[oaicite:30]{index=30}
Open Items / Unknowns:
- Whether Windows path separator support is required for `ORACLEPACK_ALLOWED_ROOTS`: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Config loader reads the listed env vars and applies defaults as documented.
- [ ] Path resolution rejects paths outside allowed roots.
- [ ] File reads enforce max bytes and indicate truncation.
- [ ] Exec gating flag is available for run tools to check.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Environment variables: ORACLEPACK_ALLOWED_ROOTS … ORACLEPACK_ENABLE_EXEC … ORACLEPACK_MAX_READ_BYTES …” :contentReference[oaicite:31]{index=31}
- “Hard deny-by-default execution … Restrict filesystem access to allowlisted roots …” :contentReference[oaicite:32]{index=32}
```

```ticket T3
T# Title: Implement oraclepack subprocess runner with timeouts and stdout/stderr truncation
Type: chore
Target Area: oraclepack_mcp_server/oraclepack_cli.py (subprocess execution)
Summary:
- Provide a subprocess wrapper to invoke the `oraclepack` CLI with a hard timeout and bounded stdout/stderr capture to prevent wedging the MCP server.
- Return structured results including exit code, duration, and truncation indicators.
In Scope:
- Async subprocess execution wrapper (create subprocess, capture stdout/stderr).
- Timeout behavior:
  - Kill process on timeout and return an explicit “Timed out after {timeout}s” error result.
- Character-limit truncation for stdout/stderr based on configured limit.
Out of Scope:
- MCP tool registration (handled in T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Running oraclepack commands yields deterministic, bounded outputs suitable for returning via MCP tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must enforce “stdout/stderr truncation and timeouts”.
Evidence:
- Guidance: “Enforce stdout/stderr truncation and timeouts …”. :contentReference[oaicite:33]{index=33}
- Runner timeout error example snippet. :contentReference[oaicite:34]{index=34}
Open Items / Unknowns:
- Default timeout values per tool beyond examples (3600/7200) are not provided outside snippets.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Runner returns: ok, exit_code, duration_s, stdout, stderr, stdout_truncated, stderr_truncated.
- [ ] Timeout produces exit_code=124 (or equivalent) and includes a timeout message.
- [ ] Outputs are truncated to configured character limit and flags are set accordingly.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.” :contentReference[oaicite:35]{index=35}
- “Timed out after {timeout_s}s: …” return structure snippet. :contentReference[oaicite:36]{index=36}
```

```ticket T4
T# Title: Implement Taskify Stage-2 detection and validation (01-*.md..20-*.md + single-pack form)
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-2 detection + validation)
Summary:
- Implement deterministic detection of Stage-2 outputs for agents, supporting both directory-form outputs (01-*.md..20-*.md) and a single-pack input file form.
- Provide validation that ensures exactly one match per prefix 01..20 and returns missing/ambiguous details.
In Scope:
- `validate_stage2_dir(out_dir)`:
  - For each prefix 01..20, glob `{pfx}-*.md`
  - Return missing patterns and ambiguous prefix matches.
- `detect_stage2(stage2_path, repo_root)`:
  - Support explicit dir path.
  - Support explicit file path with out_dir rules:
    - If under `docs/oracle-questions-YYYY-MM-DD/…`, use sibling `oracle-out` under that docs subtree; else default `repo_root/oracle-out`.
  - Support “auto” discovery (best-effort, deterministic ordering), including checking `repo_root/oracle-out`.
Out of Scope:
- Stage-3 action pack validation (handled in T7).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Agents can resolve and validate Stage-2 outputs deterministically for downstream Stage-3 workflows.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form)”.
- “Validate Stage-2 outputs (exactly one match per prefix 01..20)”.
Evidence:
- Requirements text for Stage-2 detection/validation. :contentReference[oaicite:37]{index=37}
- Validation logic snippet for 01..20 and ambiguous/missing. :contentReference[oaicite:38]{index=38}
- Out-dir rule snippet referencing `docs/oracle-questions-YYYY-MM-DD/…` → `oracle-out`. :contentReference[oaicite:39]{index=39}
Open Items / Unknowns:
- Full “auto discovery” search order beyond checking `repo_root/oracle-out`: Not provided in visible excerpts.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation returns ok=false with `missing` when any prefix has no matches.
- [ ] Validation returns ok=false with `ambiguous` when any prefix has >1 match.
- [ ] Validation returns ok=true only when exactly one match exists for every prefix 01..20.
- [ ] Detection supports explicit dir and explicit file resolution and produces an `out_dir`.
- [ ] Detection supports “auto” and returns deterministic results for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form) … Validate … exactly one match per prefix 01..20.” :contentReference[oaicite:40]{index=40}
- “out_dir rules … if under docs/oracle-questions-YYYY-MM-DD/ … then …/oracle-out else oracle-out” :contentReference[oaicite:41]{index=41}
```

```ticket T5
T# Title: Implement MCP tools for oraclepack and taskify helper operations
Type: enhancement
Target Area: oraclepack_mcp_server/server.py (MCP tool registration + schemas + formatting)
Summary:
- Implement the MCP server surface that maps `oraclepack` CLI operations and Taskify helper functions into callable MCP tools.
- Provide consistent response formatting (markdown/json) and ensure run tools respect execution gating.
In Scope:
- Tool schemas/inputs covering:
  - `oraclepack_validate_pack`
  - `oraclepack_list_steps`
  - `oraclepack_run_pack` (gated)
  - `oraclepack_read_file`
  - `oraclepack_taskify_detect_stage2`
  - `oraclepack_taskify_validate_stage2`
  - `oraclepack_taskify_validate_action_pack`
  - `oraclepack_taskify_artifacts_summary`
  - `oraclepack_taskify_run_action_pack` (gated)
- Response formatting:
  - Support JSON and Markdown result formats (including stdout/stderr blocks and truncation notes).
- CLI arg mapping for oraclepack operations (including references to flags like `--no-tui`, `--out-dir`, `--oracle-bin` as inputs or internal argv composition).
- Ensure `ORACLEPACK_ENABLE_EXEC=1` gating is enforced for run tools.
Out of Scope:
- Implementing the underlying Stage-2 detection/validation logic (T4) and Stage-3 validation/summary logic (T7), except wiring them in.
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients can invoke oraclepack and taskify workflows end-to-end via tools and receive bounded, formatted results.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must expose the tool list as stated.
- Run tools must be disabled by default unless enabled via env flag.
Evidence:
- Tool list and gating note. :contentReference[oaicite:42]{index=42}
- Server schema snippets (e.g., ReadFileInput, Taskify*Input, timeout defaults). :contentReference[oaicite:43]{index=43}
Open Items / Unknowns:
- Exact argument surface for `oraclepack_validate_pack` / `oraclepack_list_steps` (flags and required params) is not fully specified in excerpts.
Risks / Dependencies:
- Depends on config/security/runner/taskify modules existing and being wired correctly.
Acceptance Criteria:
- [ ] All tools listed in the parent ticket are registered and callable.
- [ ] `oraclepack_read_file` enforces allowed roots and max read bytes.
- [ ] `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack` refuse execution unless `ORACLEPACK_ENABLE_EXEC=1`.
- [ ] Response formatter supports markdown and json outputs and includes truncation indicators.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Tools — oraclepack_validate_pack … oraclepack_taskify_run_action_pack (requires ORACLEPACK_ENABLE_EXEC=1)” :contentReference[oaicite:44]{index=44}
- “Map … oraclepack CLI capabilities (validate/list/run + flags like --no-tui, --out-dir, --oracle-bin) into MCP tools” :contentReference[oaicite:45]{index=45}
```

```ticket T6
T# Title: Add MCP transport hardening guidance + tool UX annotations (readOnly/destructive/openWorld)
Type: enhancement
Target Area: MCP server transport configuration and tool metadata (server.py / deployment guidance)
Summary:
- Ensure the MCP server supports stdio and streamable-http in a way suitable for “real-time” usage, including the recommended security posture for HTTP transport.
- Add MCP tool annotations so clients can present appropriate approval UX for read-only vs execution tools.
In Scope:
- Transport support considerations:
  - stdio: ensure logs go to stderr (per guidance in ticket text).
  - streamable-http: implement recommended protections (“Origin validation”, “bind to localhost + auth”).
- Tool annotations:
  - Mark validate/list/read tools as `readOnlyHint: true`.
  - Mark run tools as `destructiveHint: true` and `openWorldHint: true`.
Out of Scope:
- Implementing tool business logic (T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients see proper tool risk hints and HTTP transport is protected as recommended for local/real-time usage.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support “stdio” and “streamable-http”.
- Must provide tool annotations as described.
Evidence:
- Transport choices and HTTP security recommendations. :contentReference[oaicite:46]{index=46}
- Tool annotation guidance (readOnlyHint/destructiveHint/openWorldHint). :contentReference[oaicite:47]{index=47}
Open Items / Unknowns:
- Exact auth mechanism for streamable-http (token, mTLS, etc.): Not provided.
Risks / Dependencies:
- Depends on MCP SDK capabilities available in the chosen implementation.
Acceptance Criteria:
- [ ] Running with `--transport stdio` is supported and does not interleave logs on stdout.
- [ ] Running with `--transport streamable-http` includes Origin validation and uses localhost binding + authentication (mechanism documented/implemented).
- [ ] Validate/list/read tools are annotated as read-only; run tools are annotated as destructive/open-world.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Streamable HTTP … implement Origin validation and bind to localhost + auth to avoid DNS rebinding …” :contentReference[oaicite:48]{index=48}
- “mark validate/list/read tools as readOnlyHint … mark run tools as destructiveHint, openWorldHint …” :contentReference[oaicite:49]{index=49}
```

````ticket T7
T# Title: Implement Stage-3 Action Pack validation + artifact summarization helpers
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-3 helpers) and wiring into server tools
Summary:
- Implement helper functions to validate Stage-3 “Action Pack” markdown constraints before execution and to summarize key Stage-3 artifacts produced by running action packs.
- These helpers support deterministic agent workflows around Taskify outputs.
In Scope:
- Action Pack validation logic:
  - Enforce “single ```bash fence” and “step headers” constraints as stated in the parent ticket.
- Artifacts summary logic:
  - Summarize outputs such as `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, and “pipelines doc” when present.
- Integration points:
  - Provide outputs suitable for `oraclepack_taskify_validate_action_pack` and `oraclepack_taskify_artifacts_summary` tools (registered in T5).
Out of Scope:
- Stage-2 detection/validation (T4).
- Actual execution of action packs (tool wiring and subprocess invocation handled in T5/T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Action packs can be validated for structural correctness prior to execution, and resulting artifacts can be summarized for quick agent consumption.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Validate Stage-3 Action Pack “single ```bash fence, step headers `# NN)`”.
- Summarize Stage-3 artifacts: `_actions.json`, PRD path, `tm-complexity.json`, pipelines doc, etc.
Evidence:
- Stage-3 validation requirement text. :contentReference[oaicite:50]{index=50}
- Artifact summary examples list. :contentReference[oaicite:51]{index=51}
Open Items / Unknowns:
- Exact, formal grammar for “step headers” beyond the example text: Not provided.
- Exact artifact filenames/paths beyond examples: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation fails with a clear error when the action pack violates the “single bash fence” constraint.
- [ ] Validation fails with a clear error when step headers do not meet the stated expectations.
- [ ] Artifact summarizer reports presence/absence of the example artifacts and returns a readable summary.
- [ ] Tool outputs are deterministic for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Validate Stage-3 Action Pack … (single ```bash fence, step headers `# NN)` …)” :contentReference[oaicite:52]{index=52}
- “Summarize Stage-3 artifacts (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc)” :contentReference[oaicite:53]{index=53}
````
```

.tickets/mcp/MCP Server for Oraclepack.md
```
Title:

* Add MCP server that exposes `oraclepack` + Taskify Stage-2/Stage-3 helpers as tools for agents

Summary:

* Agents need real-time access to `oraclepack` capabilities via MCP so they can validate, inspect, and run packs, then consume Taskify artifacts produced by the `oracle_pack_and_taskify-skills.md` workflow.
* Implement a secure-by-default Python MCP server that wraps the `oraclepack` CLI and adds deterministic helpers for Stage-2 detection/validation and Stage-3 Action Pack validation/execution + artifact summarization.

Background / Context:

* Request: “give access to agents/assistants the following oraclepack tool as a MCP tool… so they can perform actions using the artifacts generated from `oracle_pack_and_taskify-skills.md` … in real time.”
* Proposed reference implementation (from the conversation) is a Python project named `oraclepack-mcp-server` using FastMCP (MCP Python SDK), supporting `stdio` and `streamable-http` transports.
* Stage-2 outputs are expected to be 20 markdown files matching `01-*.md` … `20-*.md`; Stage-2 “single-pack form” needs out-dir resolution rules when the pack lives under `docs/oracle-questions-YYYY-MM-DD/...`.

Current Behavior (Actual):

* No MCP tool surface is available for `oraclepack` and Taskify workflows (agents cannot call validated tools to run/inspect packs and artifacts). (per user)

Expected Behavior:

* Agents can use MCP tools to:

  * Validate and inspect packs.
  * Run packs non-interactively to generate artifacts.
  * Detect/validate Stage-2 outputs and validate Stage-3 Action Packs before execution.
  * Summarize Stage-3 artifacts for immediate downstream consumption.

Requirements:

* MCP server implementation

  * Provide a Python MCP server project structure (e.g., `oraclepack-mcp-server/` with `oraclepack_mcp_server/` package).
  * Support transports:

    * `stdio`
    * `streamable-http`
* Tool surface (MCP tools)

  * `oraclepack_validate_pack`
  * `oraclepack_list_steps`
  * `oraclepack_run_pack` (execution gated)
  * `oraclepack_read_file`
  * `oraclepack_taskify_detect_stage2`
  * `oraclepack_taskify_validate_stage2`
  * `oraclepack_taskify_validate_action_pack`
  * `oraclepack_taskify_artifacts_summary`
  * `oraclepack_taskify_run_action_pack` (execution gated)
* Execution + safety controls

  * Deny-by-default execution; require `ORACLEPACK_ENABLE_EXEC=1` to enable “run” tools.
  * Restrict filesystem reads to allowlisted roots via `ORACLEPACK_ALLOWED_ROOTS` (colon-separated); block paths outside allowed roots.
  * Enforce timeouts and truncate stdout/stderr (`ORACLEPACK_CHARACTER_LIMIT`) and cap file read sizes (`ORACLEPACK_MAX_READ_BYTES`).
* Stage-2 reliability helpers

  * Validate Stage-2 directory contains exactly one match per prefix `01`..`20` (missing/ambiguous detection).
  * Deterministic Stage-2 detection:

    * Accept explicit dir or file.
    * If file is under `docs/oracle-questions-YYYY-MM-DD/...`, set out-dir to `docs/oracle-questions-YYYY-MM-DD/oracle-out`; otherwise default `repo_root/oracle-out`.
* Stage-3 reliability helpers

  * Validate Action Pack structure constraints before executing (e.g., “single ```bash fence, step headers”).
  * Summarize produced artifacts (examples cited: `_actions.json`, PRD path, Task Master outputs).
* Agent UX metadata

  * Apply MCP tool annotations:

    * validate/list/read: `readOnlyHint: true`
    * run: `destructiveHint: true`, `openWorldHint: true`

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Python (MCP server), wraps external `oraclepack` CLI.
* OS: Unknown
* Deployment: Unknown (local stdio vs HTTP service)
* MCP SDK version: Unknown (example uses `mcp>=1.0.0`, `pydantic>=2.0.0`).

Evidence:

* Conversation transcript + proposed reference implementation: `/mnt/data/MCP tools for Oraclepack.md`.
* Proposed env vars and tool list (deny-by-default exec, allowed roots, transports).
* Stage-2 directory validation (`01-*.md..20-*.md`) and Stage-2 out-dir resolution logic for `docs/oracle-questions-YYYY-MM-DD`.

Decisions / Agreements:

* Use a Python MCP server (FastMCP / MCP Python SDK) to expose `oraclepack` CLI + Taskify helpers.
* Support both `stdio` and `streamable-http` transports.
* Default-secure posture: execution gated by `ORACLEPACK_ENABLE_EXEC`, filesystem access constrained by allowlisted roots, truncation + timeout enforced.

Open Items / Unknowns:

* Where this MCP server should live (same repo as `oraclepack` vs separate repo) is not provided.
* Authentication / origin validation requirements for `streamable-http` deployment are mentioned conceptually but concrete requirements are not provided.
* Exact definition of “artifacts summary” contents/format beyond examples (`_actions.json`, PRD, Task Master outputs) is not provided.
* Whether `oraclepack_run_pack` must always use `--no-tui --yes --run-all` vs configurable flags is not provided (example suggests non-interactive flags).

Risks / Dependencies:

* Dependency on external `oraclepack` binary path/config (`ORACLEPACK_BIN`) and correct working directory (`ORACLEPACK_WORKDIR`).
* Security risk if exec is enabled without strict path/root controls and timeouts; mitigations are required as above.
* Stage-2 ambiguity risk when multiple `NN-*.md` match the same prefix; must report ambiguity deterministically and fail validation.

Acceptance Criteria:

* [ ] MCP server starts successfully in both `stdio` and `streamable-http` modes.
* [ ] All listed tools are exposed with the documented names.
* [ ] When `ORACLEPACK_ENABLE_EXEC!=1`, run tools refuse to execute and return a clear error; validate/list/read tools still work.
* [ ] `oraclepack_read_file` rejects paths outside `ORACLEPACK_ALLOWED_ROOTS`.
* [ ] Stage-2 validation enforces exactly one file per prefix `01`..`20` and returns `missing` and `ambiguous` sets when invalid.
* [ ] Stage-2 detection resolves out-dir correctly for both:

  * explicit dir inputs
  * single-pack file under `docs/oracle-questions-YYYY-MM-DD/...` → `docs/oracle-questions-YYYY-MM-DD/oracle-out`
  * otherwise default `repo_root/oracle-out`.
* [ ] stdout/stderr truncation and timeouts are enforced on CLI subprocess execution.
* [ ] Tool annotations are applied as specified for read-only vs destructive tools.

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement
* mcp
* oraclepack
* cli-wrapper
* taskify
* security
* tooling
```

.tickets/mcp/oraclepack-MCP.md
```
## MCP surface for `oraclepack` (so agents can act on Taskify artifacts)

### What to expose as MCP tools

Map the existing `oraclepack` CLI capabilities (validate/list/run + flags like `--no-tui`, `--out-dir`, `--oracle-bin`) into MCP tools so an agent can run packs non-interactively and then inspect the resulting state/report/artifacts.

Add a small “taskify helper” layer to make the **Stage-2 → Stage-3** workflow deterministic for agents:

* **Detect Stage-2 outputs** (dir-form `01-*.md..20-*.md` OR single-pack form) using the ordered resolver rules described in your skill.
* **Validate Stage-2 outputs** (exactly one match per prefix 01..20).
* **Validate Stage-3 Action Pack** structure constraints (single ```bash fence, step headers `# NN)`, etc.) before executing anything.
* **Summarize Stage-3 artifacts** (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc) so agents can immediately consume them.

### Transport choices (real-time vs local)

* **stdio** is simplest for local agent runtimes; it requires MCP messages only on stdout (logs must go to stderr). ([Model Context Protocol][1])
* **Streamable HTTP** is better for “real-time” multi-client usage; implement Origin validation and bind to localhost + auth to avoid DNS rebinding and local-network abuse. ([Model Context Protocol][1])

### Tool UX metadata (important for agents)

Use MCP **tool annotations** so clients can correctly present approval UX:

* mark validate/list/read tools as `readOnlyHint: true`
* mark run tools as `destructiveHint: true`, `openWorldHint: true` (they can touch filesystem, run processes, etc.) ([Model Context Protocol][2])

### Security defaults (recommended)

* Hard **deny-by-default execution**: require an env flag (e.g. `ORACLEPACK_ENABLE_EXEC=1`) to enable “run” tools.
* Restrict filesystem access to **allowlisted roots** to avoid path traversal and accidental exfiltration of secrets.
* Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.

## Reference implementation

I provided a complete, drop-in **Python MCP server** (FastMCP from the official MCP Python SDK) that exposes:

* `oraclepack_validate_pack`, `oraclepack_list_steps`, `oraclepack_run_pack`
* `oraclepack_taskify_detect_stage2`, `oraclepack_taskify_validate_stage2`
* `oraclepack_taskify_validate_action_pack`, `oraclepack_taskify_artifacts_summary`
* `oraclepack_taskify_run_action_pack`

It supports **stdio** and **streamable-http** transports per the official SDK. ([GitHub][3])

The full multi-file project is in the canvas titled **“oraclepack-mcp-server”**.

[1]: https://modelcontextprotocol.io/specification/2025-03-26/basic/transports "Transports - Model Context Protocol"
[2]: https://modelcontextprotocol.io/legacy/concepts/tools "Tools - Model Context Protocol"
[3]: https://github.com/modelcontextprotocol/python-sdk "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
```

.tickets/mcp/oraclepack_mcp_server.md
```
# oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

## tree

```text
oraclepack-mcp-server/
  README.md
  requirements.txt
  oraclepack_mcp_server/
    __init__.py
    __main__.py
    config.py
    security.py
    oraclepack_cli.py
    taskify.py
    server.py
```

---

```md
<!-- path: oraclepack-mcp-server/README.md -->
# oraclepack-mcp-server

Expose `oraclepack` (validate/list/run) plus a few Stage-3 (oraclepack-taskify) helpers as **MCP tools**, so an agent can:
- validate and inspect packs
- run packs non-interactively (`--no-tui --yes --run-all`) to produce artifacts
- validate Stage-2 outputs (01-*.md..20-*.md)
- validate Stage-3 Action Packs (single ```bash fence, step headers, etc.)
- summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)

## Install

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Configure (recommended)

Environment variables:

- `ORACLEPACK_ALLOWED_ROOTS` (default: current working directory)
  - Colon-separated list of allowed filesystem roots the server may read from.
  - Example: `ORACLEPACK_ALLOWED_ROOTS=/repo:/tmp/oracle-out`
- `ORACLEPACK_BIN` (default: `oraclepack`) – path to the oraclepack CLI
- `ORACLEPACK_WORKDIR` (default: current working directory)
  - Where packs are executed from (important for relative paths).
- `ORACLEPACK_ENABLE_EXEC` (default: `0`)
  - Must be `1` to enable `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack`.
- `ORACLEPACK_CHARACTER_LIMIT` (default: `25000`) – truncate large stdout/stderr
- `ORACLEPACK_MAX_READ_BYTES` (default: `500000`) – max bytes read from a file

## Run (stdio)

```bash
# Stdio transport is the simplest local integration.
python -m oraclepack_mcp_server --transport stdio
```

## Run (Streamable HTTP)

```bash
python -m oraclepack_mcp_server --transport streamable-http
```

## Tools

- `oraclepack_validate_pack`
- `oraclepack_list_steps`
- `oraclepack_run_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)
- `oraclepack_read_file`
- `oraclepack_taskify_detect_stage2`
- `oraclepack_taskify_validate_stage2`
- `oraclepack_taskify_validate_action_pack`
- `oraclepack_taskify_artifacts_summary`
- `oraclepack_taskify_run_action_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)

## Typical Stage-3 usage

1) Detect/validate Stage-2 outputs (directory or single-pack)
2) Validate the Action Pack markdown
3) Run the Action Pack via `oraclepack run ...`
4) Summarize produced artifacts
```

```txt
# path: oraclepack-mcp-server/requirements.txt
mcp>=1.0.0
pydantic>=2.0.0
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
__all__ = []
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/config.py
from __future__ import annotations

import os
from dataclasses import dataclass
from pathlib import Path


def _truthy(value: str | None) -> bool:
    if value is None:
        return False
    return value.strip().lower() in {"1", "true", "yes", "y", "on"}


@dataclass(frozen=True)
class ServerConfig:
    allowed_roots: tuple[Path, ...]
    oraclepack_bin: str
    workdir: Path
    enable_exec: bool
    character_limit: int
    max_read_bytes: int


def load_config() -> ServerConfig:
    # Allowed roots: colon-separated. Default to CWD.
    roots_raw = os.environ.get("ORACLEPACK_ALLOWED_ROOTS")
    if roots_raw:
        roots = tuple(Path(p).expanduser().resolve() for p in roots_raw.split(":") if p.strip())
    else:
        roots = (Path.cwd().resolve(),)

    oraclepack_bin = os.environ.get("ORACLEPACK_BIN", "oraclepack")
    workdir = Path(os.environ.get("ORACLEPACK_WORKDIR", str(Path.cwd()))).expanduser().resolve()

    enable_exec = _truthy(os.environ.get("ORACLEPACK_ENABLE_EXEC", "0"))

    character_limit = int(os.environ.get("ORACLEPACK_CHARACTER_LIMIT", "25000"))
    max_read_bytes = int(os.environ.get("ORACLEPACK_MAX_READ_BYTES", "500000"))

    return ServerConfig(
        allowed_roots=roots,
        oraclepack_bin=oraclepack_bin,
        workdir=workdir,
        enable_exec=enable_exec,
        character_limit=character_limit,
        max_read_bytes=max_read_bytes,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/security.py
from __future__ import annotations

from pathlib import Path


class PathNotAllowedError(ValueError):
    pass


def resolve_under_roots(path: Path, allowed_roots: tuple[Path, ...]) -> Path:
    """Resolve a path and enforce it lives under at least one allowed root."""
    p = path.expanduser().resolve()

    for root in allowed_roots:
        r = root.expanduser().resolve()
        try:
            p.relative_to(r)
            return p
        except ValueError:
            continue

    raise PathNotAllowedError(
        f"Path not allowed (outside allowed roots). path={p} roots={[str(r) for r in allowed_roots]}"
    )


def safe_read_text(path: Path, max_bytes: int) -> tuple[str, bool]:
    """Read up to max_bytes from a file as UTF-8 (replace errors)."""
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data.decode("utf-8", errors="replace"), truncated


def safe_read_bytes(path: Path, max_bytes: int) -> tuple[bytes, bool]:
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data, truncated
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
from __future__ import annotations

import asyncio
import os
import time
from dataclasses import dataclass
from pathlib import Path


@dataclass
class CmdResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool


def _truncate(s: str, limit: int) -> tuple[str, bool]:
    if limit <= 0:
        return s, False
    if len(s) <= limit:
        return s, False
    return s[:limit], True


async def run_cmd(
    argv: list[str],
    cwd: Path,
    timeout_s: int,
    env: dict[str, str] | None,
    character_limit: int,
) -> CmdResult:
    start = time.time()

    proc = await asyncio.create_subprocess_exec(
        *argv,
        cwd=str(cwd),
        env=(os.environ | (env or {})),
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )

    try:
        out_b, err_b = await asyncio.wait_for(proc.communicate(), timeout=timeout_s)
    except asyncio.TimeoutError:
        proc.kill()
        await proc.communicate()
        duration = time.time() - start
        return CmdResult(
            ok=False,
            exit_code=124,
            duration_s=duration,
            stdout="",
            stderr=f"Timed out after {timeout_s}s: {' '.join(argv)}",
            stdout_truncated=False,
            stderr_truncated=False,
        )

    duration = time.time() - start
    out = out_b.decode("utf-8", errors="replace") if out_b else ""
    err = err_b.decode("utf-8", errors="replace") if err_b else ""

    out, out_tr = _truncate(out, character_limit)
    err, err_tr = _truncate(err, character_limit)

    exit_code = proc.returncode if proc.returncode is not None else 1
    return CmdResult(
        ok=(exit_code == 0),
        exit_code=exit_code,
        duration_s=duration,
        stdout=out,
        stderr=err,
        stdout_truncated=out_tr,
        stderr_truncated=err_tr,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
from __future__ import annotations

import re
from dataclasses import dataclass
from datetime import date
from pathlib import Path


@dataclass
class Stage2Resolution:
    kind: str  # "dir" | "file"
    stage2_path: Path
    out_dir: Path
    notes: list[str]


PREFIXES = [f"{i:02d}" for i in range(1, 21)]


def _is_iso_date(s: str) -> bool:
    # Minimal heuristic for YYYY-MM-DD
    return bool(re.fullmatch(r"\d{4}-\d{2}-\d{2}", s))


def validate_stage2_dir(out_dir: Path) -> dict:
    missing: list[str] = []
    ambiguous: dict[str, list[str]] = {}
    selected: dict[str, str] = {}

    for pfx in PREFIXES:
        matches = sorted(out_dir.glob(f"{pfx}-*.md"))
        if len(matches) == 0:
            missing.append(f"{pfx}-*.md")
        elif len(matches) > 1:
            ambiguous[pfx] = [m.name for m in matches]
        else:
            selected[pfx] = matches[0].name

    ok = (not missing) and (not ambiguous)
    return {
        "ok": ok,
        "out_dir": str(out_dir),
        "selected": selected,
        "missing": missing,
        "ambiguous": ambiguous,
    }


def _lexi_newest(paths: list[Path]) -> Path | None:
    # Deterministic: lexicographic max
    return sorted(paths, key=lambda p: p.name)[-1] if paths else None


def detect_stage2(stage2_path: str, repo_root: Path) -> Stage2Resolution:
    notes: list[str] = []

    if stage2_path != "auto":
        p = (repo_root / stage2_path).expanduser()
        if p.exists() and p.is_dir():
            return Stage2Resolution(kind="dir", stage2_path=p.resolve(), out_dir=p.resolve(), notes=["explicit dir"])
        if p.exists() and p.is_file():
            # out_dir rules from skill: if under docs/oracle-questions-YYYY-MM-DD/ then parent/oracle-out else oracle-out
            p_res = p.resolve()
            out = repo_root / "oracle-out"
            parts = list(p_res.parts)
            if "docs" in parts:
                try:
                    idx = parts.index("docs")
                    # docs/oracle-questions-YYYY-MM-DD/.../oracle-pack-YYYY-MM-DD.md
                    if idx + 1 < len(parts) and parts[idx + 1].startswith("oracle-questions-"):
                        out = Path(*parts[: idx + 2]) / "oracle-out"
                except ValueError:
                    pass
            return Stage2Resolution(kind="file", stage2_path=p_res, out_dir=out.resolve(), notes=["explicit file"])
        raise FileNotFoundError(f"stage2_path not found: {p}")

    # auto discovery (best-effort, deterministic ordering)
    searched: list[str] = []

    # 1) repo_root/oracle-out
    candidate = repo_root / "oracle-out"
    searched.append(str(candidate))
    if candidate.is_dir():
        v = validate_stage2_dir(candidate)
        if v["ok"]:
            notes.append("auto: selected repo_root/oracle-out")
            return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 2) docs/oracle-questions-YYYY-MM-DD/oracle-out (newest by lexicographic date suffix)
    docs = repo_root / "docs"
    if docs.is_dir():
        qdirs = [p for p in docs.glob("oracle-questions-*") if p.is_dir()]
        # deterministic: sort by name and take newest
        newest_q = _lexi_newest(qdirs)
        if newest_q:
            candidate = newest_q / "oracle-out"
            searched.append(str(candidate))
            if candidate.is_dir():
                v = validate_stage2_dir(candidate)
                if v["ok"]:
                    notes.append(f"auto: selected {candidate}")
                    return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 3) single-pack form (newer): look for docs/oracle-pack-*.md or docs/oraclepacks/oracle-pack-*.md
    file_candidates: list[Path] = []
    if docs.is_dir():
        file_candidates += list(docs.glob("oracle-pack-*.md"))
        file_candidates += list((docs / "oraclepacks").glob("oracle-pack-*.md"))
        file_candidates += list(docs.glob("oracle-questions-*/oraclepacks/oracle-pack-*.md"))

    newest_file = _lexi_newest(sorted([p for p in file_candidates if p.is_file()], key=lambda p: p.name))
    if newest_file:
        notes.append(f"auto: selected single-pack {newest_file}")
        out = repo_root / "oracle-out"
        # If under docs/oracle-questions-YYYY-MM-DD/..., default out_dir there.
        if "docs" in newest_file.parts:
            try:
                idx = newest_file.parts.index("docs")
                if idx + 1 < len(newest_file.parts) and newest_file.parts[idx + 1].startswith("oracle-questions-"):
                    out = Path(*newest_file.parts[: idx + 2]) / "oracle-out"
            except ValueError:
                pass
        return Stage2Resolution(kind="file", stage2_path=newest_file.resolve(), out_dir=out.resolve(), notes=notes)

    raise FileNotFoundError(
        "stage2_path=auto could not resolve Stage-2 outputs. Searched:\n- " + "\n- ".join(searched)
    )


def validate_action_pack(pack_path: Path) -> dict:
    text = pack_path.read_text(encoding="utf-8", errors="replace")

    bash_fence = re.findall(r"(?m)^\s*```bash\s*$", text)
    any_fence = re.findall(r"(?m)^\s*```\s*", text)

    errors: list[str] = []
    if len(bash_fence) != 1:
        errors.append(f"expected exactly one ```bash fence; found {len(bash_fence)}")
    if len(any_fence) != 2:
        # One opening and one closing fence expected, and it must be a bash fence.
        errors.append(f"expected no other code fences; found {len(any_fence)} total fences")

    # Extract bash block content if possible
    bash_block = ""
    m = re.search(r"```bash\s*\n(?P<body>[\s\S]*?)\n```\s*", text)
    if m:
        bash_block = m.group("body")

    # Validate step headers inside bash fence
    step_headers = re.findall(r"(?m)^\s*#\s*(\d{2})\)\s+.*$", bash_block)
    if not step_headers:
        errors.append("no step headers found inside the bash fence (expected lines like '# 01) ...')")
    else:
        # Ensure they start at 01 and are strictly increasing by 1.
        nums = [int(x) for x in step_headers]
        if nums[0] != 1:
            errors.append(f"first step must be 01; got {nums[0]:02d}")
        for prev, cur in zip(nums, nums[1:]):
            if cur != prev + 1:
                errors.append(f"step numbers must be sequential; got {prev:02d} then {cur:02d}")

    return {
        "ok": len(errors) == 0,
        "pack_path": str(pack_path),
        "step_count": len(step_headers),
        "errors": errors,
    }


def default_pack_path(today: date | None = None) -> str:
    d = today or date.today()
    return f"docs/oracle-actions-pack-{d.isoformat()}.md"
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/server.py
from __future__ import annotations

import json
from enum import Enum
from pathlib import Path
from typing import Any

from pydantic import BaseModel, Field
from mcp.server.fastmcp import FastMCP

from .config import load_config
from .security import resolve_under_roots, safe_read_text, PathNotAllowedError
from .oraclepack_cli import run_cmd
from .taskify import detect_stage2, validate_stage2_dir, validate_action_pack


class ResponseFormat(str, Enum):
    markdown = "markdown"
    json = "json"


class PackPathInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class RunPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    out_dir: str | None = Field(default=None, description="Output directory for step execution (passes --out-dir).")

    no_tui: bool = Field(default=True, description="If true, pass --no-tui")
    yes: bool = Field(default=True, description="If true, pass --yes")
    run_all: bool = Field(default=True, description="If true, pass --run-all")

    resume: bool = Field(default=False, description="If true, pass --resume")
    stop_on_fail: bool = Field(default=True, description="If true, pass --stop-on-fail (default true)")

    roi_threshold: float = Field(default=0.0, description="Pass --roi-threshold")
    roi_mode: str = Field(default="over", description="Pass --roi-mode ('over' or 'under')")

    timeout_s: int = Field(default=3600, description="Hard timeout for the oraclepack process")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class ReadFileInput(BaseModel):
    path: str = Field(..., description="Path to a file within ORACLEPACK_ALLOWED_ROOTS")
    max_bytes: int | None = Field(default=None, description="Override max bytes read")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyDetectStage2Input(BaseModel):
    stage2_path: str = Field(default="auto", description="Dir or file, or 'auto'")
    repo_root: str = Field(default=".", description="Repo root for relative resolution")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateStage2Input(BaseModel):
    out_dir: str = Field(..., description="Directory that should contain 01-*.md..20-*.md")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to Stage-3 Action Pack markdown")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyArtifactsSummaryInput(BaseModel):
    out_dir: str = Field(..., description="Stage-3 out_dir (where _actions.json etc are written)")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyRunActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the Stage-3 Action Pack markdown")
    out_dir: str | None = Field(default=None, description="Pass --out-dir for execution")
    timeout_s: int = Field(default=7200)
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


cfg = load_config()

mcp = FastMCP(
    name="oraclepack-mcp-server",
    # For production Streamable HTTP deployments, stateless_http + json_response is recommended.
    # Clients may override by running behind an ASGI app if needed.
    stateless_http=True,
    json_response=True,
)


def _md_codeblock(lang: str, content: str) -> str:
    return f"```{lang}\n{content}\n```"


def _format_cmd_result(result: Any, response_format: ResponseFormat) -> Any:
    if response_format == ResponseFormat.json:
        return {
            "ok": result.ok,
            "exit_code": result.exit_code,
            "duration_s": result.duration_s,
            "stdout": result.stdout,
            "stderr": result.stderr,
            "stdout_truncated": result.stdout_truncated,
            "stderr_truncated": result.stderr_truncated,
        }

    lines = []
    lines.append(f"**ok**: {result.ok}")
    lines.append(f"**exit_code**: {result.exit_code}")
    lines.append(f"**duration_s**: {result.duration_s:.2f}")
    lines.append("")

    if result.stdout:
        lines.append("## stdout")
        lines.append(_md_codeblock("text", result.stdout))
        if result.stdout_truncated:
            lines.append(f"(stdout truncated to {cfg.character_limit} chars)")

    if result.stderr:
        lines.append("## stderr")
        lines.append(_md_codeblock("text", result.stderr))
        if result.stderr_truncated:
            lines.append(f"(stderr truncated to {cfg.character_limit} chars)")

    return "\n".join(lines)


def _ensure_exec_enabled() -> None:
    if not cfg.enable_exec:
        raise PermissionError(
            "Execution is disabled. Set ORACLEPACK_ENABLE_EXEC=1 to enable pack execution tools."
        )


@mcp.tool(
    name="oraclepack_validate_pack",
    annotations={
        "title": "Validate oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_validate_pack(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "validate", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_list_steps",
    annotations={
        "title": "List steps in an oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_list_steps(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "list", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_run_pack",
    annotations={
        "title": "Run an oraclepack pack (non-interactive)",
        "readOnlyHint": False,
        "destructiveHint": True,
        "idempotentHint": False,
        "openWorldHint": True,
    },
)
async def oraclepack_run_pack(params: RunPackInput) -> Any:
    _ensure_exec_enabled()

    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv: list[str] = [cfg.oraclepack_bin]
    if params.no_tui:
        argv += ["--no-tui"]
    if params.out_dir:
        out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
        argv += ["--out-dir", str(out_dir)]

    argv += ["run"]

    if params.yes:
        argv += ["--yes"]
    if params.run_all:
        argv += ["--run-all"]

    if params.resume:
        argv += ["--resume"]

    # stop-on-fail is default true in oraclepack; pass explicitly for clarity.
    argv += ["--stop-on-fail", "true" if params.stop_on_fail else "false"]

    argv += ["--roi-threshold", str(params.roi_threshold), "--roi-mode", params.roi_mode]

    argv += [str(pack_path)]

    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=params.timeout_s, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_read_file",
    annotations={
        "title": "Read a file under allowed roots",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_read_file(params: ReadFileInput) -> Any:
    p = resolve_under_roots(Path(params.path), cfg.allowed_roots)
    if not p.exists() or not p.is_file():
        raise FileNotFoundError(f"file not found: {p}")

    max_bytes = params.max_bytes if params.max_bytes is not None else cfg.max_read_bytes
    text, truncated = safe_read_text(p, max_bytes=max_bytes)

    if params.response_format == ResponseFormat.json:
        return {"path": str(p), "truncated": truncated, "content": text}

    note = f"\n\n(content truncated to {max_bytes} bytes)" if truncated else ""
    return f"# {p.name}\n\n{_md_codeblock('text', text)}{note}"


@mcp.tool(
    name="oraclepack_taskify_detect_stage2",
    annotations={
        "title": "Detect Stage-2 oraclepack outputs",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_detect_stage2(params: TaskifyDetectStage2Input) -> Any:
    repo_root = resolve_under_roots(Path(params.repo_root), cfg.allowed_roots)
    res = detect_stage2(stage2_path=params.stage2_path, repo_root=repo_root)

    payload = {
        "kind": res.kind,
        "stage2_path": str(res.stage2_path),
        "out_dir": str(res.out_dir),
        "notes": res.notes,
    }

    if params.response_format == ResponseFormat.json:
        return payload

    lines = ["# Stage-2 resolution", ""]
    lines.append(f"- **kind**: `{payload['kind']}`")
    lines.append(f"- **stage2_path**: `{payload['stage2_path']}`")
    lines.append(f"- **out_dir**: `{payload['out_dir']}`")
    if payload["notes"]:
        lines.append("- **notes**:")
        lines += [f"  - {n}" for n in payload["notes"]]
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_validate_stage2",
    annotations={
        "title": "Validate Stage-2 directory outputs (01-*.md..20-*.md)",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_validate_stage2(params: TaskifyValidateStage2Input) -> Any:
    out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
    if not out_dir.is_dir():
        raise FileNotFoundError(f"out_dir is not a directory: {out_dir}")

    report = validate_stage2_dir(out_dir)
    if params.response_format == ResponseFormat.json:
        return report

    lines = [f"# Stage-2 validation: `{out_dir}`", ""]
    lines.append(f"- **ok**: {report['ok']}")
    if report["missing"]:
        lines.append("- **missing**:")
        lines += [f"  - {m}" for m in report["missing"]]
    if report["ambiguous"]:
[TRUNCATED]
```

</source_code>