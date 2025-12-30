## 1) Overview

### Problem

You have “oracle question packs” stored as Markdown that contain a single fenced `bash` block with:

* a **prelude** (e.g., `set -euo pipefail`, `out_dir=...`, `mkdir -p ...`)
* **numbered steps** `# 01)` … `# 20)` that each run an `oracle ...` invocation

Today you can run these with an interactive Bash script, but you want a **polished, first-class CLI/TUI** built with Charm’s ecosystem.

### Target users

* Engineers / operators running packs repeatedly, wanting safe confirm-before-run, logs, resume, and good terminal UX.
* Maintainers who want pack validation and predictable behavior across machines/CI.

### Success metrics

* **Safety**: zero accidental executions (every step is explicitly confirmed unless `--yes` is set).
* **Reliability**: deterministic step parsing; clear non-zero exit codes on failure.
* **Observability**: each step yields a timestamped log file + a machine-readable run summary (JSON).
* **Usability**: fast navigation/filtering, preview of command and `--write-output`, resume after quit/crash.

### Constraints / assumptions

* Pack format is “one ` ```bash ` fence + numbered `# NN)` headers” as in your current pack output .
* `oracle` is installed and available in `PATH`, or provided via `--oracle-bin`.
* Default shell is `bash`; Windows support may require WSL/Git Bash or using `--shell` explicitly.

### “How hard is it?”

* **Low** effort for a “nice enough” wrapper using **gum** prompts around your existing script. Gum is explicitly designed to add “glamorous” prompts to shell scripts without writing Go. ([GitHub][1])
* **Moderate** effort for a **production-grade compiled CLI** with a full-screen TUI using **Bubble Tea** + **Bubbles** + **Lip Gloss**, with robust parsing, streaming logs, resume/state, and packaging polish. Bubble Tea is a Go TUI framework (Elm-style architecture) suited for full-window apps. ([GitHub][2])

**This PRD specifies the “most polished” option (compiled Go CLI + Bubble Tea TUI).**

---

## 2) Capability Tree (Functional Decomposition)

### Capability: Pack ingestion

#### Feature: Load pack from Markdown (MVP)

* **Description**: Read a Markdown pack file and extract the first ` ```bash ` fenced block.
* **Inputs**: `pack_path`
* **Outputs**: `bash_block_text`
* **Behavior**: Find the first matching fence start/end; error if missing/empty.

#### Feature: Parse steps + prelude (MVP)

* **Description**: Split extracted bash block into prelude and ordered step blocks based on `# NN)` headers.
* **Inputs**: `bash_block_text`
* **Outputs**: `Prelude`, `[]Step{number,title,body}`
* **Behavior**: Everything before first `# NN)` is prelude; each subsequent header starts a new step.

#### Feature: Detect derived metadata (MVP)

* **Description**: Best-effort parse of `out_dir="..."` (prelude) and `--write-output "..."` (step body).
* **Inputs**: `Prelude`, `Step.body`
* **Outputs**: `Pack.outDir?`, `Step.writeOutput?`
* **Behavior**: Regex-based extraction; never blocks execution if missing.

#### Feature: Validate pack structure (MVP)

* **Description**: Validate required structure and produce actionable errors.
* **Inputs**: `Pack`
* **Outputs**: `[]ValidationIssue`
* **Behavior**: Enforce “>=1 step”, step numbers are two-digit, no duplicates; warn on missing `--write-output`.

---

### Capability: Execution orchestration

#### Feature: Execute prelude once per run (MVP)

* **Description**: Run prelude in the chosen shell before running step bodies.
* **Inputs**: `Prelude`, `shell`, `workdir`, `env`
* **Outputs**: `PreludeResult`
* **Behavior**: If prelude fails, stop and mark run failed.

#### Feature: Execute a single step with streaming logs (MVP)

* **Description**: Run one step in a subprocess and capture stdout/stderr lines.
* **Inputs**: `Step`, `shell`, `workdir`, `env`
* **Outputs**: `StepResult{exitCode, startedAt, endedAt, logPath}`
* **Behavior**: Stream lines to UI/console and persist to a log file.

#### Feature: Inject flags into oracle invocation (MVP)

* **Description**: Optionally inject user-provided flags after `oracle` in each step.
* **Inputs**: `Step.body`, `injectFlags`
* **Outputs**: transformed `Step.body`
* **Behavior**: Rewrite only lines beginning with `oracle` (conservative, predictable).

#### Feature: Stop/continue policy on failures (MVP)

* **Description**: Define behavior when a step fails.
* **Inputs**: `StepResult`, `policy`
* **Outputs**: next action (stop / prompt / continue)
* **Behavior**: Default: stop and require explicit user action to continue.

---

### Capability: User interaction (TUI + CLI)

#### Feature: Full-screen TUI step list (MVP)

* **Description**: Show steps with status (Pending/Running/OK/Failed/Skipped) and metadata.
* **Inputs**: `Pack`, `RunState`
* **Outputs**: terminal UI
* **Behavior**: Navigate, filter, select; show `--write-output` and command preview.

#### Feature: Confirm-before-run per step (MVP)

* **Description**: Require confirmation before executing a step (unless `--yes`).
* **Inputs**: selected `Step`, `mode`
* **Outputs**: run / skip decision
* **Behavior**: Prompt with clear consequences and output path.

#### Feature: Plain mode (no TUI) (MVP)

* **Description**: Provide `--no-tui` mode that prompts in-line (or runs with `--yes`).
* **Inputs**: same as TUI
* **Outputs**: console output
* **Behavior**: Deterministic text output and exit codes.

#### Feature: Markdown rendering in preview (Non-MVP)

* **Description**: Render relevant Markdown (e.g., step prompt text) in a styled terminal view.
* **Inputs**: markdown snippet
* **Outputs**: ANSI-rendered text
* **Behavior**: Use Glamour for stylesheet-based terminal Markdown rendering. ([GitHub][3])

---

### Capability: State, resume, and reporting

#### Feature: Persist run state (MVP)

* **Description**: Save step statuses and pointers so a run can be resumed.
* **Inputs**: `RunState`
* **Outputs**: `state.json` in a deterministic directory
* **Behavior**: Atomic writes; load on startup when `--resume` is set.

#### Feature: Machine-readable summary report (MVP)

* **Description**: Write `run.json` with per-step results and log paths.
* **Inputs**: `RunState`, `Pack`
* **Outputs**: JSON file
* **Behavior**: Stable schema version; include environment metadata (shell, cwd, start time).

---

### Capability: Packaging and operational polish

#### Feature: Shell completions + man/help text (Non-MVP)

* **Description**: Provide completions and high-quality help UX.
* **Inputs**: CLI command model
* **Outputs**: completion scripts, man page
* **Behavior**: Generated during release; consistent flag naming.

#### Feature: Release automation (Non-MVP)

* **Description**: Produce multi-platform binaries and package manifests (brew/scoop).
* **Inputs**: tags/releases
* **Outputs**: release artifacts
* **Behavior**: Use GoReleaser capabilities for Homebrew/Scoop publishing. ([GoReleaser][4])

---

## 3) Repository Structure + Module Definitions (Structural Decomposition)

### Proposed repository layout

```
orpack/
  cmd/orpack/
    main.go
  internal/
    app/            # application composition (wiring)
    cli/            # command model, flags, help text
    pack/           # markdown fence extraction + step parsing + validation
    exec/           # subprocess execution + streaming + log persistence
    state/          # run state model + persistence
    report/         # run summary JSON schema + writer
    tui/            # Bubble Tea model/view; Bubbles components usage
    render/         # markdown-to-ANSI rendering (Glamour)
    errors/         # typed errors + exit codes
```

### Module: `internal/pack`

* **Responsibility**: Parse pack Markdown into `Pack{Prelude, Steps}` and validate format.
* **Exports**:

  * `Load(path string) (Pack, error)`
  * `Validate(p Pack) []Issue`
  * `DeriveMetadata(p *Pack)` (fills `outDir`, `writeOutput` best-effort)

### Module: `internal/exec`

* **Responsibility**: Execute prelude/steps in a shell; stream output; write log files.
* **Exports**:

  * `Runner{Shell, WorkDir, Env, InjectFlags}`
  * `RunPrelude(ctx, prelude string) (Result, error)`
  * `RunStep(ctx, prelude string, step Step, logDir string, onLine func(Line)) (Result, error)`

### Module: `internal/state`

* **Responsibility**: Track statuses and persist/restore run state.
* **Exports**:

  * `RunState` (pack hash, start time, per-step status/result)
  * `LoadState(path string) (RunState, error)`
  * `SaveStateAtomic(path string, s RunState) error`

### Module: `internal/report`

* **Responsibility**: Write stable JSON summary for automation/CI.
* **Exports**:

  * `ReportV1`
  * `WriteReport(path string, r ReportV1) error`

### Module: `internal/tui`

* **Responsibility**: Full-screen terminal UX using Bubble Tea + Bubbles + Lip Gloss.
* **Exports**:

  * `Run(p Pack, runner *exec.Runner, opts Options) error`
* **Notes**: Bubble Tea provides the stateful TUI architecture; Bubbles provides ready components; Lip Gloss provides styling/layout. ([GitHub][2])

### Module: `internal/render`

* **Responsibility**: Render Markdown snippets to ANSI for previews/help panes.
* **Exports**:

  * `RenderMarkdown(md string, style Style) (string, error)`
* **Notes**: Use Glamour for stylesheet-based terminal markdown rendering. ([GitHub][3])

### Module: `internal/cli`

* **Responsibility**: Parse args, define subcommands, map flags to app behavior.
* **Exports**:

  * `Execute(argv []string) int` (returns exit code)

### Module: `internal/errors`

* **Responsibility**: Centralize error types and exit code mapping.
* **Exports**:

  * `type Code int`
  * `ExitCode(err error) int`

### Module: `internal/app`

* **Responsibility**: Wire modules into a runnable application.
* **Exports**:

  * `New(...) *App`
  * `RunTUI(...) error`
  * `RunPlain(...) error`
  * `Validate(...) error`

---

## 4) Dependency Chain (layers, explicit “Depends on: […]”)

### Foundation Layer

* **errors**: exit codes and typed errors. Depends on: []
* **pack**: parsing + validation. Depends on: [errors]
* **state**: state model + persistence. Depends on: [errors]
* **report**: report schema + writer. Depends on: [errors]
* **render**: markdown rendering adapter. Depends on: [errors]

### Execution Layer

* **exec**: shell execution + streaming + logs. Depends on: [errors, pack]

### Application Layer

* **app**: composition and run orchestration. Depends on: [pack, exec, state, report, render, errors]

### Interaction Layer

* **cli**: command/flags and entrypoints. Depends on: [app, errors]
* **tui**: Bubble Tea UI. Depends on: [app, pack, state, render, errors]

(Acyclic by construction; `app` is the orchestration hub, `errors` is the base.)

---

## 5) Development Phases (Phase 0…N; entry/exit criteria; tasks with dependencies + acceptance criteria + test strategy)

### Phase 0: Foundation

**Entry criteria**: repo initialized; CI runs `go test ./...`

* **Task**: Implement `internal/errors`

  * Depends on: []
  * Acceptance criteria: stable exit codes; unit tests for mapping
  * Test strategy: unit tests
* **Task**: Implement `internal/pack` parsing + validation

  * Depends on: [errors]
  * Acceptance criteria: parses packs shaped like your current output ; clear issues for missing fence/steps
  * Test strategy: golden tests with fixture Markdown files; fuzz-ish tests for malformed inputs
* **Task**: Implement `internal/state` persistence

  * Depends on: [errors]
  * Acceptance criteria: atomic save/load; schema version field present
  * Test strategy: unit tests with temp dirs; corruption handling tests
* **Exit criteria**: `pack.Load + Validate`, `state.Save/Load`, exit codes proven by tests.

### Phase 1: Execution core

**Entry criteria**: Phase 0 complete

* **Task**: Implement `internal/exec` runner (prelude + step)

  * Depends on: [errors, pack]
  * Acceptance criteria: streams stdout/stderr; writes log file; returns accurate exit code
  * Test strategy: integration tests using small deterministic shell scripts (`printf`, `exit 1`)
* **Task**: Implement flag injection transform

  * Depends on: [pack]
  * Acceptance criteria: only modifies intended `oracle` lines; does not rewrite arbitrary commands
  * Test strategy: unit tests with representative step bodies
* **Exit criteria**: can execute a single parsed step and produce logs deterministically.

### Phase 2: Non-TUI CLI (usable early)

**Entry criteria**: Phase 1 complete

* **Task**: Implement `internal/app` orchestration (plain mode)

  * Depends on: [pack, exec, state, report, errors]
  * Acceptance criteria: `orpack run --no-tui` supports confirm/skip, stop-on-fail policy, writes state + report
  * Test strategy: integration tests running the CLI against fixture packs (mock oracle via dummy commands)
* **Task**: Implement `internal/cli` with subcommands

  * Depends on: [app, errors]
  * Acceptance criteria: `run`, `validate`, `list` (optional) produce predictable text output
  * Test strategy: CLI snapshot tests (stdout/stderr) + exit code assertions
* **Exit criteria**: tool is useful without TUI.

### Phase 3: TUI (polish)

**Entry criteria**: Phase 2 complete

* **Task**: Implement `internal/tui` step list + status rendering

  * Depends on: [app, pack, state, errors]
  * Acceptance criteria: navigate/filter steps; run/skip; show per-step metadata
  * Test strategy: model-level tests for update messages and state transitions
* **Task**: Live log viewport

  * Depends on: [tui, exec]
  * Acceptance criteria: streaming output appears during execution; persisted logs still written
  * Test strategy: integration-ish tests using a fake line stream injected into model
* **Exit criteria**: full-screen TUI supports the primary workflow.

### Phase 4: Rendering + reporting polish

**Entry criteria**: Phase 3 complete

* **Task**: Implement `internal/render` (Markdown preview)

  * Depends on: [errors]
  * Acceptance criteria: renders markdown snippets to ANSI using Glamour. ([GitHub][3])
  * Test strategy: golden tests (ANSI output normalized)
* **Task**: Stabilize report schema v1

  * Depends on: [report, state]
  * Acceptance criteria: `run.json` includes per-step results, log paths, pack hash, schema version
  * Test strategy: schema conformance tests
* **Exit criteria**: consistent artifacts for both humans and automation.

### Phase 5: Packaging and distribution

**Entry criteria**: stable CLI behavior + docs

* **Task**: Add release pipeline (GoReleaser)

  * Depends on: [cli]
  * Acceptance criteria: reproducible builds; generates Homebrew/Scoop metadata as configured. ([GoReleaser][4])
  * Test strategy: CI dry-run of release config
* **Exit criteria**: users can install via release artifacts with minimal friction.

---

## 6) User Experience

### Personas

* **Operator**: runs packs, wants minimal mistakes; values confirmations, resume, clear failure context.
* **Maintainer**: authors packs; values `validate`, actionable parse errors, predictable numbering.

### Key flows

1. **Validate a pack**

   * `orpack validate pack.md` → prints issues + derived summary (`out_dir`, step count, detected outputs).
2. **Run in TUI**

   * `orpack run pack.md` → list steps → select → preview command and `--write-output` → confirm → stream logs → mark OK/Failed.
3. **Resume**

   * `orpack run --resume pack.md` → loads saved state → cursor moves to next pending step.
4. **Plain mode / CI**

   * `orpack run --no-tui --yes pack.md` → runs remaining steps deterministically; writes `run.json`.

### UX notes tied to capabilities

* Always display the detected `--write-output` prominently before confirmation.
* Provide explicit “Stop on failure / Continue?” behavior; never silently continue after errors.
* Keep keybindings discoverable (help line), consistent with common Bubble Tea apps. ([GitHub][2])

---

## 7) Technical Architecture

### System components

* **Parser**: Markdown → bash fence → prelude/steps.
* **Executor**: builds script = prelude + step; runs via `bash -lc` (default); streams lines.
* **State/Report**: write `state.json` (resume) + `run.json` (audit).
* **TUI**: Bubble Tea program + Bubbles list/viewport/spinner; Lip Gloss layout/styling. ([GitHub][2])
* **Markdown rendering**: Glamour (optional) for previews/help. ([GitHub][3])

### Data models

* `Pack{path, hash, outDir?, prelude, steps[]}`
* `Step{number, title, body, writeOutput?}`
* `RunState{schemaVersion, packHash, startedAt, steps: map[number]StepState}`
* `StepState{status, lastResult?}`
* `Result{exitCode, startedAt, endedAt, logPath}`

### External integrations

* `oracle` binary (invoked as-is)
* shell runtime (`bash` default)
* filesystem for logs/state/report

### Key decisions and trade-offs

* **Compiled Go + Bubble Tea** over shell-only: enables richer UX, better structure, cross-platform releases; more implementation surface area. Bubble Tea is explicitly intended for rich TUIs. ([GitHub][2])
* **Gum as fallback/alternative**: fastest path but limited extensibility; still valuable for “shell-first” users. ([GitHub][1])
* **Glamour for Markdown rendering**: consistent ANSI rendering with stylesheets for preview panes. ([GitHub][3])

---

## 8) Test Strategy

### Test pyramid targets

* **Unit**: 70% (parser, validators, injection transform, state/report serialization)
* **Integration**: 25% (executor running controlled shell snippets; CLI end-to-end with fixture packs)
* **E2E**: 5% (smoke tests in CI for major OS targets as feasible)

### Coverage minimums

* Pack parsing + validation: near-100% branch coverage (high risk surface)
* Executor: focus on error paths (timeouts, non-zero exits, stderr handling)

### Critical scenarios per module

* **pack**

  * Missing/empty fence; multiple fences; malformed headers; duplicate step numbers; no steps.
* **exec**

  * Prelude fails; step fails; long output; stderr-only output; non-UTF8 handling strategy (define behavior).
* **state/report**

  * Partial runs; resume from saved state; atomic write; schema version mismatch.
* **tui**

  * Status transitions; run/skip flows; failure stop policy; log rendering under high volume.

### Integration points

* Parsing → execution uses the same step bodies as read from pack fixtures derived from your known format.

---

## 9) Risks and Mitigations

### Risk: Pack format drift

* **Impact**: High
* **Likelihood**: Medium
* **Mitigation**: strict `validate` with clear errors; keep parser conservative; add fixture packs as regression tests.
* **Fallback**: allow `--format=legacy|strict` modes if needed.

### Risk: Executing arbitrary shell content (safety/security)

* **Impact**: High
* **Likelihood**: Medium
* **Mitigation**: confirm-by-default; show command preview; provide `--dry-run` and `--print-script` modes; never auto-run without `--yes`.
* **Fallback**: sandboxing is out-of-scope; document threat model.

### Risk: Cross-platform shell assumptions

* **Impact**: Medium
* **Likelihood**: Medium
* **Mitigation**: `--shell` flag; detect Windows and warn; document WSL/Git Bash expectations.
* **Fallback**: support `pwsh` only if pack bodies are compatible (likely not for bash-specific prelude).

### Risk: High-volume output overwhelms TUI

* **Impact**: Medium
* **Likelihood**: Medium
* **Mitigation**: ring buffer for in-memory viewport; always persist full logs to disk; provide “open log file path” UX.

---

## 10) Appendix

### Source format baseline

* Your existing pack structure and interactive-run expectations are exemplified in the uploaded script/pack discussion.

### Charm ecosystem references

* Bubble Tea (TUI framework; Elm-style architecture). ([GitHub][2])
* Bubbles (production-used components for Bubble Tea). ([GitHub][5])
* Lip Gloss (terminal styling/layout). ([GitHub][6])
* Gum (glamorous shell prompts without writing Go). ([GitHub][1])
* Glow (Markdown in terminal) and Glamour (Markdown rendering library for CLI apps). ([GitHub][7])
* GoReleaser packaging docs (Homebrew/Scoop/Actions). ([GoReleaser][4])

[1]: https://github.com/charmbracelet/gum?utm_source=chatgpt.com "charmbracelet/gum: A tool for glamorous shell scripts"
[2]: https://github.com/charmbracelet/bubbletea?utm_source=chatgpt.com "charmbracelet/bubbletea: A powerful little TUI framework"
[3]: https://github.com/charmbracelet/glamour?utm_source=chatgpt.com "charmbracelet/glamour: Stylesheet-based markdown ..."
[4]: https://goreleaser.com/customization/homebrew_casks/?utm_source=chatgpt.com "Homebrew Casks"
[5]: https://github.com/charmbracelet/bubbles?utm_source=chatgpt.com "charmbracelet/bubbles: TUI components for Bubble Tea"
[6]: https://github.com/charmbracelet/lipgloss?utm_source=chatgpt.com "charmbracelet/lipgloss: Style definitions for nice terminal ..."
[7]: https://github.com/charmbracelet/glow?utm_source=chatgpt.com "charmbracelet/glow: Render markdown on the CLI, with ..."
