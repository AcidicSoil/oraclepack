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
