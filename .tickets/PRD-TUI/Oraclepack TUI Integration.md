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
