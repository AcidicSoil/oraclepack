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
