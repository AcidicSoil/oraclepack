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
