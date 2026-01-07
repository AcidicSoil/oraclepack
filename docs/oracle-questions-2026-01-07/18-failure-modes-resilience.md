1) Direct answer (resilience mechanisms; gaps; where configured)

Script “label line” sanitization exists and is applied to both preludes and steps. SanitizeScript converts single-token, label-shaped lines into echo "<token>" to avoid accidental command execution (e.g., a bare Setup line). This is invoked from Runner.RunPrelude and Runner.RunStep before execution. (internal/exec/sanitize.go: SanitizeScript; internal/exec/runner.go: RunPrelude, RunStep)

Sanitization is deliberately scoped to reduce false positives. It skips blank lines and comments (#...), skips anything inside heredocs (tracks start via heredocStartToken and ignores lines until the terminator), and only rewrites lines where strings.Fields(trimmed) yields exactly 1 field. (internal/exec/sanitize.go: SanitizeScript, heredocStartToken)

Sanitization avoids rewriting real shell constructs and real executables. Tokens are rewritten only if they match labelTokenRegex (^[A-Za-z][A-Za-z0-9_-]*$), are not shell builtins/keywords, and are not found on $PATH (osexec.LookPath(token) must fail). (internal/exec/sanitize.go: labelTokenRegex, shellBuiltins, shellKeywords, LookPath check)

Sanitization produces explicit warnings and they are surfaced at least to the run log stream. Warnings are stored on the runner and also written to logWriter as ⚠ oraclepack: sanitized label ..., and can be retrieved via DrainWarnings(). (internal/exec/runner.go: recordWarnings, DrainWarnings)

Execution is cancellation-aware via context. exec.CommandContext(ctx, ...) is used; on failure, if ctx.Err() != nil it returns the context error directly. This is a resilience mechanism for user cancel / deadline propagation. (internal/exec/runner.go: run)

Execution errors are wrapped with a sentinel classification. Non-context execution errors return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err), enabling upstream code to recognize “execution failed” as a class. (internal/exec/runner.go: run)

Configurable execution knobs exist, but they’re mostly “wiring,” not resilience policies. RunnerOptions configures Shell (default /bin/bash), WorkDir, Env (appended to os.Environ()), OracleFlags, Overrides, and ChatGPTURL. These influence behavior but do not add retries/limits/strict-mode by themselves. (internal/exec/runner.go: RunnerOptions, NewRunner)

Critical path gap: “stop-on-fail” within a script is not enforced. The runner executes bash -lc <script> without forcing set -e / -euo pipefail. So a failing command earlier in a script may not fail the step unless the script author ensured it (depends on the script). This is a major resilience gap for “partial failure looks like success.” (internal/exec/runner.go: run uses r.Shell, "-lc", script)

Critical path gap: sanitization misses common “label + trailing comment” patterns. Because only len(fields) == 1 is sanitized, a line like Setup # do thing becomes two fields and is not rewritten; bash will attempt to execute Setup and fail (or worse, execute if Setup exists on PATH). (internal/exec/sanitize.go: Fields/len(fields) != 1 early-continue)

Critical path gap: environment
