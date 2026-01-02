1. Direct answer (1–4 bullets, evidence-cited)

* **Some demand is likely, but it’s narrow**: people running long “packs” (build/test/deploy chains) will want to keep using the TUI while a step runs, or reduce total wall-clock time; however the current product positioning is explicitly **sequential** (“Run All (Sequential)”) with **resume/state** semantics that assume a single linear progression. (README.md: “Run All (Sequential)”, “State Persistence… `.state.json` … `--resume`”, “Real-time Streaming”, “Machine Readable Reports”)
* **Parallel execution is risky for this pack format**: packs are “interactive bash scripts embedded within Markdown” and the format implies shared context (prelude variables “available to all steps”), so parallelizing arbitrary shell steps can break ordering, shared filesystem assumptions, environment reuse, and determinism. (README.md: “interactive bash scripts…”, “Prelude… Variables defined here are available to all steps”)
* **Background (detached) execution inside the TUI is the higher-probability win than true parallelism**: it improves UX for long-running steps without inventing a job/worker subsystem and without changing the sequential contract. (README.md: TUI + real-time streaming + run-all sequential framing)

2. Risks/unknowns (bullets)

* **Unknown user demand**: README/go.mod don’t give any signals about pack sizes, typical runtimes, or user requests; there’s no evidence of issues/telemetry or “CI scale” focus in what’s provided.
* **Execution model uncertainty**: whether steps run in one shared shell/session or as separate execs with prelude re-injected changes the feasibility of safe parallelism and “backgrounding.” (Need to see internal/exec + internal/tui run loop.)
* **State/report invariants**: `.state.json` + `.report.json` likely assume a single “current step”; adding concurrent “running” states can create race conditions and complicate resume semantics.

3. Next smallest concrete experiment (1 action)

* **Add a “detach step” UX in TUI (not parallelism):** implement a hotkey (e.g., `d`) that starts the selected step in the background **but keeps execution sequential** (only one step may run at a time), returning the user to the step list with a “running” indicator and allowing them to re-open the live output viewport when they reselect it. This validates the strongest “long pack” pain point (waiting) without a job/queue system.

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* `internal/tui/**` (especially the step execution loop and output viewport plumbing; e.g., `internal/tui/tui.go` around `runStep` / run-all logic)
* `internal/exec/**` (runner model, shell/session reuse, process lifecycle; e.g., `internal/exec/runner.go`, any command execution helpers)
* `internal/state/**` and `internal/report/**` (schema assumptions about “current step”, concurrency safety, resume behavior)
* Repo tree proof for absence/presence of job infra: `internal/**/jobs/**`, `internal/**/worker/**`, plus any `queue`/`scheduler` mentions via grep output if you have it
