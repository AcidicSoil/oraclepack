1. Direct answer (1–4 bullets, evidence-cited)

* **Yes** — `RunPlain` should print the *resolved output directory actually used for writes*, not `Runner.WorkDir`, because `RunPlain` currently reports `a.Runner.WorkDir` (`fmt.Fprintf(out, "Output directory: %s\n", a.Runner.WorkDir)`), which can be wrong when `--out-dir` (or any config override) redirects outputs elsewhere.
* Prefer printing a single canonical value (ideally an absolute, cleaned path) sourced from the same field the runner uses to decide where artifacts are written (e.g., `Runner.OutDir` / `Config.OutDirResolved`), so the log becomes a reliable debugging anchor rather than a hint that can be misleading.

2. Risks/unknowns (bullets)

* `Runner.WorkDir` might currently *also* be the resolved out dir in some implementations; changing output could appear “no-op” in many cases and hide the real mismatch unless tests cover overrides.
* If external tooling parses the exact “Output directory:” line and expects `WorkDir`, changing it could be a breaking change (mitigation: print both, or keep the label explicit: “Runner work dir:” vs “Resolved output dir:”).
* Without seeing where `--out-dir` is applied, there’s a chance the *real* resolved directory is computed later than this print, or per-step.

3. Next smallest concrete experiment (1 action)

* Add a test that runs `RunPlain` with a temp `--out-dir` override and asserts the printed “Output directory” equals the directory that actually gets created/used (and that `Runner.WorkDir` is different in the setup to prove the test would have failed before).

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* The config/runner resolution points where `--out-dir` is parsed and applied, e.g.:

  * `internal/config/*.go` (or wherever `Config` is defined; fields like `OutDir`, `WorkDir`, etc.)
  * `internal/app/prepare*.go` (or wherever `a.Prepare()` resolves directories)
  * `internal/runner/*.go` (definition of `Runner`, especially where `WorkDir` is set and where outputs are written)
  * Any CLI binding file that wires `--out-dir` into `a.Config` (likely `internal/cli/*.go`).
