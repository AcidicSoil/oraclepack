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
