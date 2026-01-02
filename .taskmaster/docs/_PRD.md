## 1) Overview

### Problem

`oraclepack` is a Bubble Tea TUI for running “oracle packs” (bash steps embedded in Markdown) with state/reporting and automatic flag injection into `oracle` commands.
Today, changing **oracle CLI flags** or the **ChatGPT project URL/endpoint** typically requires editing configs or re-invoking the CLI differently per run, and there is no TUI flow to (a) pick overrides, (b) choose which steps they apply to, and (c) validate the override set before running.

### Who has it (target users)

* Developers using `oraclepack` to run repeatable AI-assisted workflows but needing per-run customization (different ChatGPT project folders, different oracle flags for debugging, different remotes).
* Pack authors who want packs to remain stable while allowing users to safely tweak runtime behavior.

### Why current solutions fail

* Current injection is global and line-based (`InjectFlags` rewrites lines starting with `oracle`), which can break common multi-line `oracle \` formatting and cannot selectively apply overrides per step.
* No TUI affordance exists to stage overrides and confirm/validate them before returning to the main run screen.

### Success metrics

* Override workflow completion rate: ≥ 95% of users can (flags + URL + step targeting) and return to main run screen without errors.
* Preflight prevention: ≥ 90% of invalid flag/url combinations are caught by validation before any step execution.
* UX efficiency: Configure overrides in ≤ 30 seconds for typical cases (multi-select flags + pick steps + confirm).
* Execution correctness: Overrides apply only to selected steps; non-selected steps run with baseline behavior.

Assumptions (explicit):

* Upstream `oracle` supports `--dry-run` and `--chatgpt-url` for safe validation and URL targeting. ([GitHub][1])
* Overrides are **ephemeral per TUI session/run** (no persistence beyond the run) unless added later (out of scope).

---

## 2) Capability Tree (Functional Decomposition)

### Capability: Runtime Overrides Management (MVP)

Manage a staged set of runtime overrides (flags + project URL) and their step targets.

#### Feature: Overrides data model

* **Description**: Represent user-selected flags, selected project URL, and selected step targets.
* **Inputs**: Baseline runner flags (current), user selections, pack step IDs.
* **Outputs**: `RuntimeOverrides` object usable by TUI + execution.
* **Behavior**: Store “added flags”, “removed flags”, `chatgptURL` (optional), and `applyToSteps` set.

#### Feature: Merge baseline flags + runtime overrides (step-aware)

* **Description**: Compute effective flags per step.
* **Inputs**: Baseline flags (existing `Runner.OracleFlags`), overrides, step ID.
* **Outputs**: Effective flag list and any key-value flag pairs (e.g., `--chatgpt-url <url>`).
* **Behavior**: If step not targeted, return baseline; else baseline + adds − removes, plus URL flag injection when set.

---

### Capability: Flag Picker UI (MVP)

Allow choosing additional oracle flags at runtime.

#### Feature: Multi-select flags picker

* **Description**: Pick known/common flags (multi-select) and optionally remove baseline flags.
* **Inputs**: Baseline flags, a curated list of known flags, current overrides.
* **Outputs**: Updated overrides (added/removed flags).
* **Behavior**: Toggle selection; show “added” vs “removed” vs “baseline unchanged”.

---

### Capability: Project URL (Endpoint) UI (MVP)

Allow entering/selecting a ChatGPT project URL for browser runs.

#### Feature: Project URL input screen

* **Description**: Enter a ChatGPT project URL string.
* **Inputs**: Text input (URL), existing selected URL.
* **Outputs**: Updated overrides `chatgptURL`.
* **Behavior**: Basic validation (non-empty, looks like URL); store exactly as entered.

#### Feature: Project URL selection menu (optional within MVP if simple list is kept in-memory)

* **Description**: Choose among recently entered URLs within the same session.
* **Inputs**: Session list, selected value.
* **Outputs**: Updated overrides `chatgptURL`.
* **Behavior**: Select one; allow “clear”.

Upstream basis: `oracle` supports `--chatgpt-url` and config `browser.chatgptUrl` for targeting a ChatGPT project folder. ([GitHub][1])

---

### Capability: Step Targeting UI (MVP)

Choose which steps receive overrides.

#### Feature: Multi-select step picker

* **Description**: Select pack steps that should receive runtime overrides.
* **Inputs**: Pack steps (IDs/titles), current selection set.
* **Outputs**: Updated overrides `applyToSteps`.
* **Behavior**: Toggle step selection; “select all / none”.

---

### Capability: Confirmation + Mode 2 Validation (MVP)

Validate selected overrides against upstream `oracle` before returning to main run screen.

#### Feature: Confirmation summary screen

* **Description**: Show flags added/removed, URL, and affected steps before applying.
* **Inputs**: Overrides + pack step list.
* **Outputs**: User confirmation or cancel.
* **Behavior**: Render a diff-style summary; confirm triggers validation.

#### Feature: Mode 2 validation runner

* **Description**: Validate that the override set is accepted by upstream `oracle` CLI.
* **Inputs**: Pack steps targeted, effective flags per targeted step, extracted oracle invocations.
* **Outputs**: Pass/fail with actionable error text including step + failing invocation.
* **Behavior**: For each targeted step, run each `oracle ...` invocation with overrides plus `--dry-run summary` to ensure CLI parsing succeeds without spending tokens. ([GitHub][1])

Requirement basis: “Mode 2 validation” must run before returning to run screen; failures must block readiness and show actionable errors.

---

### Capability: Execution Integration (MVP)

Apply overrides during actual run.

#### Feature: Step execution uses effective flags (step-aware)

* **Description**: Inject computed flags only into targeted steps.
* **Inputs**: Step code, step ID, effective flags.
* **Outputs**: Executed step with correct oracle invocations.
* **Behavior**: Replace current global injection behavior with per-step effective flags. Current injection is via `Runner.RunStep` → `InjectFlags`.

#### Feature: Hardened oracle injection (multi-line tolerant)

* **Description**: Inject flags without breaking common `oracle \` multi-line formatting.
* **Inputs**: Step script text, flags to inject.
* **Outputs**: Rewritten script.
* **Behavior**: Detect oracle invocations across line continuations and inject before trailing `\` on the invocation line when present.

Risk basis: current line-based approach can break multi-line commands.

---

## 3) Repository Structure + Module Definitions (Structural Decomposition)

Current structure includes `internal/exec`, `internal/tui`, `internal/app`, etc.

### Proposed additions/changes

```
internal/
  overrides/
    types.go
    merge.go
    validate.go
  exec/
    inject.go          (extend/replace with multiline-safe injector)
    oracle_scan.go     (new: extract oracle invocations)
    oracle_validate.go (new: mode-2 validation runner)
  tui/
    overrides_flow.go  (new: state machine for overrides screens)
    overrides_flags.go (new: flags picker model)
    overrides_steps.go (new: step picker model)
    overrides_url.go   (new: project URL input/selector)
    overrides_confirm.go (new: summary + validation screen)
```

### Module: `internal/overrides`

* **Maps to capability**: Runtime Overrides Management
* **Responsibility**: Own the runtime override data model and step-aware flag resolution.
* **Exports**:

  * `type RuntimeOverrides`
  * `func (o RuntimeOverrides) EffectiveFlags(stepID string, baseline []string) []string`
  * `func (o RuntimeOverrides) Targeted(stepID string) bool`
  * `func (o RuntimeOverrides) Summary(packSteps []pack.Step) OverridesSummary`

### Module: `internal/exec/oracle_scan.go`

* **Maps to capability**: Mode 2 Validation (oracle invocation extraction)
* **Responsibility**: Extract oracle invocations from bash step code robustly enough for validation + injection.
* **Exports**:

  * `type OracleInvocation { Raw string; Display string }`
  * `func ExtractOracleInvocations(script string) []OracleInvocation`

### Module: `internal/exec/inject.go` (updated)

* **Maps to capability**: Execution Integration + Hardened Injection
* **Responsibility**: Inject flags into oracle invocations without breaking multi-line formatting.
* **Exports**:

  * `func InjectFlags(script string, flags []string) string` (same signature; improved implementation)

### Module: `internal/exec/oracle_validate.go`

* **Maps to capability**: Mode 2 validation runner
* **Responsibility**: Execute safe validations (`--dry-run summary`) for extracted invocations with overrides.
* **Exports**:

  * `type ValidationError { StepID, Invocation string; Output string }`
  * `func ValidateOverrides(ctx context.Context, shell string, workDir string, env []string, steps []pack.Step, baseline []string, ov overrides.RuntimeOverrides) error`

### Module: `internal/tui/overrides_flow.go` (+ related files)

* **Maps to capability**: Flag Picker UI, Project URL UI, Step Targeting UI, Confirmation UX
* **Responsibility**: Bubble Tea models for the overrides wizard and integration back to main TUI model.
* **Exports**:

  * `func NewOverridesFlowModel(...) tea.Model` (or integrated into existing `tui.Model` as substates)
  * Messages: `OverridesAppliedMsg`, `OverridesCancelledMsg`, `OverridesValidationFailedMsg`

### Module changes: `internal/tui/tui.go`

* **Maps to capability**: Execution Integration + UX integration
* **Responsibility change**: Add a new view state for overrides flow; show current overrides status on main screen; pass overrides into runner step execution path.
* **Exports**: unchanged (existing `NewModel`), but internal state extended.

---

## 4) Dependency Chain (layers, explicit “Depends on: […]”)

### Foundation Layer (Phase 0)

* **internal/overrides**: override types + merge logic.
  Depends on: []
* **internal/exec/oracle_scan**: extract oracle invocations.
  Depends on: []
* **internal/exec/inject (improved)**: multiline-safe injection.
  Depends on: [internal/exec/oracle_scan] (for shared detection rules)

### Validation Layer (Phase 1)

* **internal/exec/oracle_validate**: mode-2 validation runner (`--dry-run summary`). ([GitHub][1])
  Depends on: [internal/overrides, internal/exec/oracle_scan]

### TUI Flow Layer (Phase 2)

* **internal/tui/overrides_* models**: flags picker, url input, step picker, confirm/validate screen.
  Depends on: [internal/overrides, internal/exec/oracle_validate]

### Integration Layer (Phase 3)

* **internal/tui/tui.go integration**: entrypoint keybinding/menu, apply overrides to step execution, display summary.
  Depends on: [internal/tui/overrides_*]
* **internal/app / internal/exec runner integration**: ensure `RunStep` can accept per-step effective flags (or TUI updates Runner flags before each step).
  Depends on: [internal/overrides, internal/exec/inject]

No cycles: all dependencies flow from overrides/scan → validate → tui flow → integration.

---

## 5) Development Phases (Phase 0…N; entry/exit criteria; tasks with dependencies + acceptance criteria + test strategy)

### Phase 0: Foundations (Overrides model + scanning + injection)

**Entry criteria**: Current pack parsing and TUI build clean.
**Tasks**:

* [ ] Implement `internal/overrides` data model + step-aware merge

  * **Depends on**: []
  * **Acceptance criteria**: Given baseline flags and overrides, `EffectiveFlags(stepID)` matches expected for targeted vs non-targeted steps.
  * **Test strategy**: Unit tests for add/remove precedence, empty sets, step targeting.

* [ ] Implement `internal/exec/oracle_scan.ExtractOracleInvocations`

  * **Depends on**: []
  * **Acceptance criteria**: Extracts oracle invocations from (a) single-line `oracle ...`, (b) `oracle \` + continued lines, (c) indented variants.
  * **Test strategy**: Unit tests with representative scripts; snapshot extracted `Display`.

* [ ] Upgrade `internal/exec/InjectFlags` to be multiline-tolerant

  * **Depends on**: [internal/exec/oracle_scan]
  * **Acceptance criteria**: Injects flags into multiline `oracle \` form without breaking trailing backslash; preserves non-oracle lines.
  * **Test strategy**: Extend `inject_test.go` with multiline continuation fixtures (current tests exist).

**Exit criteria**: Per-step effective flag computation and injection behave correctly under unit tests.

---

### Phase 1: Mode 2 validation runner

**Goal**: Validate overrides safely before returning to run screen.
**Entry criteria**: Phase 0 complete.

**Tasks**:

* [ ] Implement `internal/exec/oracle_validate.ValidateOverrides` using `--dry-run summary`

  * **Depends on**: [internal/overrides, internal/exec/oracle_scan]
  * **Acceptance criteria**:

    * For each targeted step, each extracted oracle invocation is executed with overrides + `--dry-run summary` and fails fast on first error.
    * Error includes step ID, invocation, and captured output.
  * **Test strategy**:

    * Unit tests for command construction (pure functions).
    * Integration test with a fake `oracle` binary on PATH (test harness script) that asserts received args.

Upstream basis for dry run: `oracle --dry-run [summary|json|full]` previews without sending. ([GitHub][1])

**Exit criteria**: Validation produces deterministic pass/fail with actionable error payload.

---

### Phase 2: Overrides wizard TUI

**Goal**: Provide the requested picker flow and confirmation/validation gating.
**Entry criteria**: Phase 1 complete.

**Tasks**:

* [ ] Add overrides flow state machine + entrypoint from main run screen

  * **Depends on**: [internal/overrides, internal/exec/oracle_validate]
  * **Acceptance criteria**: From main steps screen, user can enter overrides flow and return (cancel or apply).
  * **Test strategy**: Bubble Tea model tests for state transitions (message-driven).

* [ ] Implement flags picker (multi-select)

  * **Depends on**: [overrides flow state machine]
  * **Acceptance criteria**: Select/deselect flags; mark removed baseline flags; navigation works.
  * **Test strategy**: Model tests simulate keypresses; verify overrides state.

* [ ] Implement step picker (multi-select)

  * **Depends on**: [overrides flow state machine]
  * **Acceptance criteria**: Can select any subset; “all/none” supported.
  * **Test strategy**: Model tests; verify `applyToSteps` set.

* [ ] Implement project URL input/selection menu

  * **Depends on**: [overrides flow state machine]
  * **Acceptance criteria**: Enter URL; clear URL; appears in summary.
  * **Test strategy**: Model tests for textinput.

* [ ] Implement confirmation screen that runs validation on confirm

  * **Depends on**: [flags picker, step picker, url input, internal/exec/oracle_validate]
  * **Acceptance criteria**:

    * Summary includes flags added/removed, affected steps, selected URL.
    * Confirm triggers validation; failure shows error and does not apply.
  * **Test strategy**: Model tests with stubbed validator (inject interface).

**Exit criteria**: Full wizard: flags → steps → confirm/validate → return to run screen.

---

### Phase 3: Execution integration + UX polish

**Goal**: Overrides affect actual execution exactly as selected.

**Tasks**:

* [ ] Apply step-aware effective flags during `runStep` execution path

  * **Depends on**: [Phase 0 injection, Phase 2 applied overrides messages]
  * **Acceptance criteria**: Running a non-targeted step uses baseline flags only; targeted step includes overrides.
  * **Test strategy**: Integration test with fake `oracle` that logs args for selected step only.

* [ ] Display staged overrides summary on main run screen

  * **Depends on**: [Phase 2]
  * **Acceptance criteria**: Main view indicates overrides active and shows brief summary (e.g., “Overrides: +2 flags, URL set, 3 steps”).
  * **Test strategy**: View snapshot tests (string contains markers).

**Exit criteria**: End-to-end usable: configure overrides in TUI, validated, then run steps with correct behavior.

---

## 6) User Experience

### Personas

* **Pack runner**: Wants to run a known pack but direct output into a specific ChatGPT project folder and enable debug flags temporarily.
* **Pack author**: Wants users to safely override runtime settings without editing the pack or shared config.

### Key flows

1. **Open overrides flow** from main steps screen (new keybinding/menu entry).
2. **Flags picker**: multi-select adds/removes.
3. **Steps picker**: choose steps receiving overrides.
4. **Project URL menu**: enter/select/clear URL (separate screen within flow).
5. **Confirm**: show full summary; confirm runs validation.
6. **Return to main run screen**: overrides staged; run steps as normal.

### UI/UX notes (tied to capabilities)

* Always show an “Overrides staged” indicator on the main screen to avoid hidden behavior changes.
* Validation errors must include: step ID + the failing oracle invocation + raw output.
* Keep “full config menu” out of scope; only runtime pickers + project URL.

---

## 7) Technical Architecture

### System components

* **TUI (Bubble Tea)**: collects overrides and step targets; triggers validation; applies overrides to execution.
* **Overrides core**: step-aware merge + summary formatting.
* **Exec layer**:

  * Oracle invocation extraction
  * Multiline-safe injection
  * Mode 2 validation runner (`oracle --dry-run summary ...`) ([GitHub][1])
* **Runner**: executes rewritten step scripts (current `Runner.RunStep` injects flags).

### Data models

* `RuntimeOverrides`:

  * `AddedFlags []string`
  * `RemovedFlags []string`
  * `ChatGPTURL string` (optional)
  * `ApplyToSteps map[string]bool`
* `OverridesSummary`:

  * counts + formatted lists for confirm screen and main screen indicator

### Key decision: validate via `--dry-run summary` (Mode 2)

* **Rationale**: Upstream `oracle` provides a dry-run mode specifically to preview without sending tokens, so it exercises real CLI parsing. ([GitHub][1])
* **Trade-offs**: Requires executing `oracle` during validation; may still read files and can fail if referenced paths/vars are unresolved.
* **Alternative**: `oracle --help`-based validation (weaker; doesn’t validate combos). Not chosen for MVP.

### Key decision: heuristic multiline-safe injection (not full shell parsing)

* **Rationale**: Addresses the most common breakage (`oracle \` continuation) without introducing a full bash AST parser.
* **Trade-offs**: Some exotic shell constructs may not be perfectly handled; mitigate with tests and clear error output.

---

## 8) Test Strategy

### Test pyramid targets

* Unit: ~70% (overrides merge, scan, injection)
* Integration: ~25% (validation runner + fake oracle binary, runner integration)
* E2E: ~5% (optional: simulate TUI flow at message level)

### Coverage minimums

* Unit-tested modules (overrides, scan, inject): ≥ 90% line coverage
* Validation runner: ≥ 80% line coverage with integration fixtures

### Critical scenarios (by module)

* **internal/overrides**

  * Happy: targeted step gets baseline+adds−removes + url flag
  * Edge: empty overrides; remove non-existent flag; applyToSteps empty
* **internal/exec/oracle_scan**

  * Happy: single-line oracle
  * Edge: multiline with backslashes and indentation
  * Error: no oracle invocations → validation is a no-op success (if targeted steps contain none)
* **internal/exec/inject**

  * Multiline injection preserves `\`
  * Does not inject into non-oracle commands
* **internal/exec/oracle_validate**

  * Fails fast and returns structured error with output
  * Uses `--dry-run summary` consistently ([GitHub][1])
* **internal/tui overrides flow**

  * Cancel returns without applying
  * Confirm runs validator; failure stays in flow with readable error
  * Apply sends message to main model and updates indicator

---

## 9) Risks and Mitigations

### Risk: Injection breaks uncommon shell patterns

* **Impact**: Medium (step execution could change meaning)
* **Likelihood**: Medium
* **Mitigation**: Focus on well-defined heuristics (start-of-command + multiline `\`), add fixture tests for real packs, provide clear docs/limitations.
* **Fallback**: Allow disabling runtime injection for a step (not in MVP; could be added later).

### Risk: Validation executes `oracle` but fails due to environment/path assumptions

* **Impact**: Medium
* **Likelihood**: Medium
* **Mitigation**: Validation runs only extracted oracle invocations; capture and display full output; allow user to adjust overrides/targets.
* **Fallback**: Provide “skip validation” escape hatch (not in MVP unless necessary; conflicts with requirement).

### Risk: Upstream `oracle` flags change

* **Impact**: Medium
* **Likelihood**: Low/Medium
* **Mitigation**: Treat the flags list in the picker as curated (best-effort) and allow manual entry for advanced flags; keep validation authoritative.
* **Fallback**: Update curated list independently; validation remains source of truth.

---

## 10) Appendix

### References

* Runtime overrides requirements and acceptance criteria.
* Current `oraclepack` code structure, current injection path (`Runner.RunStep` → `InjectFlags`).
* `oraclepack` product description and existing CLI behavior.
* Upstream `oracle` CLI supports `--dry-run` and `--chatgpt-url` configuration patterns. ([GitHub][1])

### Glossary

* **Oracle Pack**: Markdown file containing one `bash` block with numbered steps executed by `oraclepack`.
* **Mode 2 validation**: For this PRD, defined as executing extracted oracle invocations with overrides using `--dry-run summary` to validate CLI parsing without sending. ([GitHub][1])

### Open questions

* Should project URL injection also support remote browser service flags (`--remote-host/--remote-token`) in addition to `--chatgpt-url`? Upstream supports these. ([GitHub][1])
* Should entered project URLs be persisted across runs (state/config) or remain session-only? (Currently assumed session-only per scope.)

[1]: https://github.com/steipete/oracle "GitHub - steipete/oracle: Ask the oracle when you're stuck. Invoke GPT-5 Pro with a custom context and files."
