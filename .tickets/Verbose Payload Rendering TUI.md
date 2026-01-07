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
