1) Direct answer (top integration points, minimal defining files/locations)

Shell execution boundary (runs arbitrary scripts via OS process execution)

Integration: os/exec launches a shell with -lc (login shell semantics), sets working dir, env, and streams stdout/stderr. This is the core “external system” boundary for running everything.

Contract/config: RunnerOptions{Shell, WorkDir, Env, OracleFlags, Overrides, ChatGPTURL} → Runner (constructed in NewRunner).

Defined in: internal/exec/runner.go (NewRunner, RunPrelude, RunStep, run). Evidence: exec.CommandContext(ctx, r.Shell, "-lc", script), cmd.Dir = r.WorkDir, cmd.Env = r.Env, cmd.Stdout/Stderr = logWriter.

Oracle CLI invocation boundary (flag injection + per-step overrides + optional dry-run validation)

Integration: Oracle commands are expected to appear inside step code; runner injects flags into those commands before execution, and validation can run oracle --dry-run summary for overridden steps.

Contract/config: baseline Runner.OracleFlags + optional overrides.RuntimeOverrides (scoped by step.ID) + optional ChatGPTURL mapped into flags (ApplyChatGPTURL); validation is scoped by over.ApplyToSteps[step.ID].

Defined in:

internal/exec/runner.go (RunStep: ApplyChatGPTURL, Overrides.EffectiveFlags, InjectFlags).

internal/exec/oracle_validate.go (ValidateOverrides: ExtractOracleInvocations, --dry-run summary, execDryRun).

Note: The actual parsing/injection logic lives elsewhere (functions referenced but not shown): ExtractOracleInvocations, InjectFlags, ApplyChatGPTURL.

Environment-variable boundary (inherit + append + PATH propagation during validation)

Integration: Environment is inherited from the parent process and extended with caller-provided entries; validation explicitly re-exports PATH into the command string to ensure the shell sees it.

Contract/config: RunnerOptions.Env merged with os.Environ(); findEnvValue(env, "PATH") drives PATH export behavior in dry-run validation.

Defined in:

internal/exec/runner.go (NewRunner: Env: append(os.Environ(), opts.Env...)).

internal/exec/oracle_validate.go (env := append(os.Environ(), opts.Env...), execDryRun PATH export + quoting via shellQuote).

Filesystem outputs boundary (working directory as the implicit artifact root)

Integration: All script execution happens with cmd.Dir = r.WorkDir, so any relative file reads/writes done by steps land under that directory (and any output directories are created by the scripts themselves).

Contract/config: RunnerOptions.WorkDir / Runner.WorkDir.

Defined in: internal/exec/runner.go (run: cmd.Dir = r.WorkDir), plus va
