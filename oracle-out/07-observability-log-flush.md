Direct answer (evidence-cited)

Yes—make LineWriter.Close() guaranteed via defer, and ensure Close() flushes any buffered partial line. Right now runStep calls lw.Close() after RunStep returns (err := m.runner.RunStep(...); lw.Close()), so it will run on normal error returns, but it is not guaranteed to run if RunStep panics or if future edits add early returns before the close. Evidence: runStep constructs lw, calls RunStep, then calls lw.Close() (not deferred) in internal/tui/tui.go:runStep.

If you’re already “losing the last line,” the most likely root cause is that exec.LineWriter buffers until \n and does not flush remainder on Close(). The TUI is already structured to append each callback line into logLines and render immediately on LogMsg, so flushing the final partial line at Close() is the key place to fix this (same code path: LineWriter{ Callback: func(line string) { m.logChan <- line } }, then RunStep, then Close()).

Risks/unknowns

We can’t confirm buffering/flush semantics without internal/exec/linewriter.go (or wherever exec.LineWriter is implemented). If it already flushes on Close(), then the symptom points elsewhere (e.g., RunStep panics, or writes occur after Close() from background goroutines).

Potential double-close / write-after-close behavior if RunStep also closes the writer internally or writes asynchronously after returning; Close() should be idempotent and safe.

Backpressure risk: the callback writes into m.logChan (buffer 100). If a step emits logs faster than the UI consumes, the callback can block and affect the runner’s IO pumping; that can look like “missing last line” or hangs rather than just flush issues.

Next smallest concrete experiment (1 action)

Add a unit test in the exec package for LineWriter: write a string with no trailing newline, call Close(), and assert the callback receives that final buffered fragment exactly once (and that calling Close() twice doesn’t duplicate it).

If evidence is insufficient, attach next

internal/exec/linewriter.go (or the file that defines type LineWriter and its Write/Close methods).

internal/exec/runner.go (or wherever RunStep is implemented), especially any goroutines reading stdout/stderr and writing to the LineWriter.
