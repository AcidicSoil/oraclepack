Title:

* Add TUI “Runtime Overrides” picker to inject Oracle flags + project URLs per selected oraclepack steps (with Mode 2 validation)

Summary:

* Extend the `oraclepack` Go TUI to let users (1) choose/enter ChatGPT project URLs (endpoints) and (2) add/remove Oracle CLI flags at runtime, then (3) select which oraclepack steps should receive those overrides, and (4) confirm before returning to the main run screen. Validation should use the “mode 2” approach (per user preference) to ensure selected flags/overrides are accepted by the upstream `oracle` CLI before running.

    Branch · Extending OraclePack F…

Background / Context:

* `oraclepack` is a wrapper around upstream `oracle` ([https://github.com/steipete/oracle](https://github.com/steipete/oracle)) and currently injects `Runner.OracleFlags` into steps by rewriting step shell scripts via `InjectFlags` in `internal/exec/Runner.RunStep` (per assistant).

    Branch · Extending OraclePack F…

* User goal: allow runtime customization for different use cases, specifically cherry-picking flags and selecting/entering different project endpoints (project URLs) without editing config files per run.

    Branch · Extending OraclePack F…

* TUI flow requested by user: flags picker → steps-to-modify picker → confirm → return to initial run screen; plus a separate menu for project URL input/injection; a broader config menu is explicitly deferred to the future.

    Branch · Extending OraclePack F…

Current Behavior (Actual):

* Runtime flag/endpoint selection via TUI pickers is not available (per user request implying missing functionality).

    Branch · Extending OraclePack F…

* Current injection mechanism is line-based regex append-at-EOL; assistant notes it can break common multi-line `oracle \` command formatting (risk/limitation in current approach).

    Branch · Extending OraclePack F…

Expected Behavior:

* From the main run screen, a user can open a TUI flow to:

  * Add/remove Oracle flags (multi-select).

  * Choose which oraclepack steps will be affected (multi-select).

  * Enter/select project URL(s) to be injected (separate menu in TUI).

  * Confirm changes and return to the initial run screen to proceed with normal step selection and run.

        Branch · Extending OraclePack F…

* “Mode 2” validation runs before returning to the run screen, rejecting invalid flag combinations/overrides and showing an actionable error.

    Branch · Extending OraclePack F…

Requirements:

* TUI picker flow:

  * Multi-select UI to add/remove flags for a run.

        Branch · Extending OraclePack F…

  * Multi-select UI to choose which oraclepack steps are modified by the selected flag changes.

        Branch · Extending OraclePack F…

  * Confirmation step summarizing selected flags and affected steps; then return to the initial run screen (no immediate run required).

        Branch · Extending OraclePack F…

* Project URL injection:

  * Provide a separate TUI menu for users to input project URL(s) that oraclepacks get sent to (endpoint customization).

        Branch · Extending OraclePack F…

* Validation:

  * Use “mode 2” validation (user preference) to verify the injected flags/overrides are valid for `oracle` before proceeding.

        Branch · Extending OraclePack F…

  * Assistant proposed “dry-run parse per affected step” as the practical mode-2 mechanism, executing affected `oracle ...` invocations in `--dry-run` and failing fast with step + invocation + error output.

        Branch · Extending OraclePack F…

* Constraints:

  * Do not build the “full customizable config menu” now; stick to the runtime pickers and project URL input described above.

        Branch · Extending OraclePack F…

Out of Scope:

* Full configurable config/settings menu beyond project URL injection and runtime flag/step selection (explicitly deferred).

    Branch · Extending OraclePack F…

Reproduction Steps:

* Not provided.

Environment:

* OS: Unknown

* `oraclepack` version/commit: Unknown

* TUI framework: Bubble Tea (implied by assistant discussion; exact versions unknown).

    Branch · Extending OraclePack F…

Evidence:

* Conversation export: /mnt/data/Branch · Extending OraclePack Functionality.md

    Branch · Extending OraclePack F…

* Notable referenced components/files (content not included in the export):

  * `oraclepack-all.md` (referenced as code base context)

  * `config.json` (referenced for project URL / remote fields)

        Branch · Extending OraclePack F…

* Key quoted intent (per user, from the export): “picker in the TUI… cherry-pick additional flags… cherry-pick which oraclepack steps… confirm changes… separate menu… project urls… In the future… full customizable config menu… but for now…”

    Branch · Extending OraclePack F…

Decisions / Agreements:

* Use a picker-based TUI flow (not arbitrary runtime command injection UI as the primary UX) (per user).

    Branch · Extending OraclePack F…

* Prefer “mode 2” validation (per user).

    Branch · Extending OraclePack F…

* Defer full config menu; implement only the specified pickers + project URL input now (per user).

    Branch · Extending OraclePack F…

Open Items / Unknowns:

* Definition of “mode 2” validation is partially ambiguous:

  * Assistant earlier described mode 2 as help-based preflight (`oracle --help` / `oracle <subcommand> --help`), and later described mode 2 as `--dry-run` validation per invocation; user only confirmed “mode 2” generally.

        Branch · Extending OraclePack F…

* Where and how project URLs should be stored (ephemeral per run vs persisted registry) is not specified by the user.

    Branch · Extending OraclePack F…

* Exact oraclepack step model (IDs/titles, where steps are defined) and where in the TUI state machine to add the new screens is not provided in the export (assistant referenced `internal/tui/tui.go`, but code is not attached).

    Branch · Extending OraclePack F…

Risks / Dependencies:

* Dependency on upstream `oracle` CLI behavior and flags (validation and injected overrides must match supported CLI semantics).

    Branch · Extending OraclePack F…

* Injection mechanism risk: current line-based `InjectFlags` approach may break multi-line commands with `\` continuations; may require hardening or more robust parsing to safely inject flags (per assistant).

    Branch · Extending OraclePack F…

Acceptance Criteria:

* A new TUI menu entry (from the main run screen) launches a runtime overrides flow that:

  * Lets the user add/remove flags (multi-select) and proceeds to a step-selection picker (multi-select).

        Branch · Extending OraclePack F…

  * Includes a project URL input/selection menu to control where oraclepacks are sent.

        Branch · Extending OraclePack F…

  * Shows a confirmation screen summarizing: selected flags (added/removed), affected steps, and selected/entered project URL(s).

        Branch · Extending OraclePack F…

  * On confirmation, performs “mode 2” validation; if validation fails, the UI displays the failure and does not proceed back to the main run screen as “ready”.

        Branch · Extending OraclePack F…

  * If validation succeeds, returns to the main run screen with overrides staged for the run, and the user can run steps as normal.

        Branch · Extending OraclePack F…

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* flags

* validation

* config

* endpoints

* oraclepack

---
