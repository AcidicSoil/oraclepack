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
