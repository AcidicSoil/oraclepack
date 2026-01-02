1. Direct answer (1–4 bullets, evidence-cited)

* Yes—`NewRunner` should not hardcode `"/bin/bash"` as the implicit default; it will reliably fail on Windows (and some Unix variants) because the path won’t exist, and the current error path will look like an execution failure rather than a configuration/platform issue. Evidence: `NewRunner` sets `shell = "/bin/bash"` when `opts.Shell == ""`, and `run()` always executes `exec.CommandContext(ctx, r.Shell, "-lc", script)`.
* Prefer **OS-aware default selection** plus a **fail-fast / preflight check** when the shell is missing. This turns “confusing runtime exec error” into “actionable configuration error” before any step runs. Evidence: the runner assumes “bash -lc” semantics (“We use bash -lc…”) but doesn’t verify the shell exists or is compatible with `-lc`.
* On Windows specifically, either (a) default to a discoverable `bash` on `PATH` (Git Bash / MSYS2) if present, otherwise (b) fail fast with a clear message (“Windows requires WSL/Git Bash or set --shell …”). This matches the intent of having `RunnerOptions.Shell` as an override while avoiding a default that is guaranteed-wrong on Windows. Evidence: `RunnerOptions` already exposes `Shell`, but the default overrides to `"/bin/bash"` when unset.

2. Risks/unknowns (bullets)

* **Windows support contract is unclear**: if oraclepack is intentionally Unix-only, the correct behavior may be to fail immediately on `GOOS=windows` with an explicit unsupported-platform error, rather than attempting to auto-detect Git Bash. (No repo-level support statement provided.)
* **Shell compatibility**: `-lc` is not universally supported across all shells; if you ever default to something other than bash (e.g., `sh`), behavior may change subtly (login shell init files, PATH, etc.). The comment explicitly expects bash-like behavior.
* **API change cost**: making `NewRunner` return `(*Runner, error)` is the cleanest fail-fast design, but may cascade to multiple call sites. If you keep the current signature, you’ll need to fail fast in `run()` (or add a separate validation step) instead.

3. Next smallest concrete experiment (1 action)

* Extract a small pure function (e.g., `defaultShell(goos string) (string, error)` or `defaultShellCandidates(goos string) []string`) and add a unit test table that asserts:

  * `windows` → either returns `"bash"` (preferred if you intend to support Git Bash) **or** returns an error indicating unsupported/missing shell;
  * `linux/darwin` → returns `"/bin/bash"` (or `"bash"` if you choose PATH-based detection).
    This lets you test OS-specific selection without needing cross-OS test execution.

4. If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

* Call site(s) for `NewRunner` and any CLI flags affecting shell selection: `internal/app/*.go`, `internal/cli/*.go` (especially where `RunnerOptions` is constructed).
* Any documented platform support or install expectations: `README.md`, `docs/**`, or `internal/**/doc.go` if present.
