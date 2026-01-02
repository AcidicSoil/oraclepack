1. Direct answer (1–4 bullets, evidence-cited)

* **Yes — add an explicit JSON-to-stdout mode for CI**, because the current `--no-tui` path writes human output to `os.Stdout` (`return a.RunPlain(..., os.Stdout)`), which is awkward to parse in pipelines. (See `if noTUI { return a.RunPlain(context.Background(), os.Stdout) }` in `internal/cli/run.go`.)
* **Keep stdout “clean” when JSON is enabled**: route all human/log/step output to **stderr** and reserve **stdout** for a single JSON object/line. Otherwise, JSON will be interleaved with run logs because `RunPlain` currently targets stdout. (Same `RunPlain(..., os.Stdout)` evidence.)
* **Reuse the existing report schema rather than inventing a new one**, since you already compute `reportPath := base + ".report.json"` and pass it via `cfg.ReportPath`. A stdout JSON mode can emit the same structure (or a small stable “summary” subset) for consistency with the report file. (See `reportPath` and `cfg.ReportPath` setup in `internal/cli/run.go`.)
* **Make it opt-in via a flag (`--json` or `--output=json`) and implicitly disable TUI if set**, because the TUI program path (`tea.NewProgram(...)`) assumes interactive terminal rendering and shouldn’t share stdout with machine output. (See the `noTUI` gate vs. TUI branch in `internal/cli/run.go`.)

2. Risks/unknowns (bullets)

* **Unknown what `app.RunPlain` prints and when** (e.g., step stdout vs. structured status). If it streams step output to the provided writer, switching that writer to stderr may be a behavior change for users who relied on stdout piping.
* **Early failures** (`Prepare()` / `LoadState()`) may happen before a report is generated; JSON mode needs a deterministic error payload (and still return a non-zero exit code).
* **Schema stability/versioning**: if stdout JSON is used in CI, you’ll want an explicit version field and compatibility guarantees to avoid breaking downstream parsers.
* **Duplication/confusion with `ReportPath`**: if both a report file and stdout JSON exist, users may expect them to match byte-for-byte; that implies emitting the same model at the same stage of execution.

3. Next smallest concrete experiment (1 action)

* Implement `--json` such that in `--no-tui` mode you run with **logs to stderr**, then **emit a single JSON object to stdout** derived from the same report model (preferably the existing `ReportV1` or a `Summary`-only struct), and add one integration test that runs a minimal pack and asserts stdout is valid JSON with expected keys (and that stderr contains the human log lines).

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* `internal/app/run.go` (especially `RunPlain` implementation and what it writes)
* `internal/report/types.go` and `internal/report/generate.go` (to confirm the canonical schema to emit)
* `internal/state/types.go` (to confirm what is available in-memory for a failure JSON when report generation doesn’t occur)
* Any CLI/output helpers if present (e.g., `internal/app/*.go` or `internal/cli/*.go`) that already handle printing or report writing beyond what’s shown here
