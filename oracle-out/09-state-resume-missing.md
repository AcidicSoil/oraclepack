1. Direct answer (1–4 bullets, evidence-cited)

* Yes—when the user passes `--resume`, a missing state file should typically *degrade to a fresh `RunState`* (with a clear warning) rather than aborting, because “missing file” is a common benign condition (first run, cleaned dir, different `--out-dir`, etc.).
* Do **not** change `LoadState` to silently create new state unconditionally; keep it “strict” and handle the fallback **at the `--resume` call site**, so other callers still get a hard error when they truly require a state file.
* Evidence: `LoadState` currently treats `os.IsNotExist(err)` as an error and returns `nil, fmt.Errorf("state file not found: %w", err)`, which will bubble up and abort unless the caller intercepts it.

2. Risks/unknowns (bullets)

* Silent fallback can hide real configuration mistakes (wrong state path, wrong `--out-dir`, permission issues misdiagnosed as “missing”), leading users to think they resumed when they didn’t.
* If fallback creates a fresh state and then writes it, you may overwrite expectations around “resume” semantics (e.g., user expected prior step results to be honored).
* Unknown: the exact `--resume` semantics and how state path is derived (needs the caller code to ensure warnings/UX are correct).

3. Next smallest concrete experiment (1 action)

* Add a CLI-level test for `--resume` where the state path does not exist: assert the command **does not error**, prints a **warning like “resume state not found; starting fresh”**, and proceeds with a newly initialized `RunState`.

4. If evidence is insufficient, exact missing file/path pattern(s) to attach next

* The `--resume` handling call site and state-path derivation, likely in one or more of:

  * `internal/cli/run.go` (flag parsing / resume logic)
  * `internal/app/run.go` (wiring runner/state into execution)
  * Any helper constructing state paths, e.g. `internal/state/*.go` (besides `persist.go`)
