Parent Ticket:

* Title: Prevent oraclepack pack failures caused by orphaned `-p/--prompt` lines in generated bash steps
* Summary: Generated oraclepack markdown packs can emit a multiline `oracle ...` command where `-p "$(cat <<'PROMPT' ...)"` starts on a new line without a continuation, causing Bash to treat `-p` as a standalone command and fail (`exit status 127`). The fix requires making pack generation structurally safe and adding validator guardrails that fail fast on regressions.
* Source:

  * Link/ID: Bash command syntax fix.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make the generator/template structurally unable to emit orphaned flags… and make oraclepack validate fail fast”
* Global Constraints:

  * “never put `-p/--prompt` (or any flag) on its own line”
  * “no inline comments at end of an `oracle ...` line”
* Global Environment:

  * Unknown
* Global Evidence:

  * Error: `bash: line 59: -p: command not found` + `exit status 127`
  * Reference pack: `oracle-pack-2026-01-08-tickets-direct.md` (pattern repeated across steps)

Split Plan:

* Coverage Map:

  * “`bash: line 59: -p: command not found` + `exit status 127`” → T1
  * “`-p "$(cat <<'PROMPT' ...)"` is on the next line without `\` … repeated in others” → T1
  * “Minimal fix: add a continuation backslash, or put `-p` on the same line” → T1
  * “Optional comment goes ABOVE the command, not inline” → T1
  * “Wherever you render `oracle ... "${ticket_args[@]}" # extra_files ...` then newline then `-p ...` … change it” → T1
  * “Permanent template rule: never put `-p/--prompt` on its own line; build prompt first, then call oracle on a single command line” → T1
  * “Enforce: no inline comments at end of an `oracle ...` line” → T1
  * “If you must wrap long lines, require explicit `\` continuations and disallow comments on continued lines” → T1
  * “Add checks to `oraclepack validate` after extracting the single `bash` fence” → T2
  * “Add `bash -n` syntax check” → T2
  * “Add `shellcheck` static analysis” → T2
  * “Custom ‘orphaned flag line’ detector (regex + continuation exceptions)” → T2
  * “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI Run/Rerun paths)” → T3
  * “Add CI/pre-commit: run `oraclepack validate` on any generated/modified pack” → T3
  * “Offer: point at exact rendering pattern + canonical snippet” → Non-actionable / Info-only
* Dependencies:

  * T3 depends on T2 because “Make validate unavoidable” is intended to enforce the added validator checks that catch regressions.
* Split Tickets:

```ticket T1
T# Title: Make pack generation structurally safe (no orphaned `-p/--prompt` lines)
Type: bug
Target Area: Pack generator/template that emits oraclepack Markdown steps (tickets-direct pack generation)
Summary:
- Generated packs can split an `oracle ...` invocation across lines such that `-p "$(cat <<'PROMPT' ...)"` starts on a new line without a continuation.
- Bash then executes `-p` as a standalone command, causing `command not found` and `exit status 127`.
- Update generation patterns so prompts are built safely and the `oracle` command remains a single logical command (or uses correct continuations without inline comment footguns).
In Scope:
- Eliminate multiline `oracle` invocations that place `-p/--prompt` on its own line.
- Apply the “minimal fix” pattern where multiline is unavoidable: add a line-continuation `\` (and ensure comments do not break continuation).
- Enforce generator rule: no inline trailing comments on `oracle ...` lines (comments/newlines can terminate the command unexpectedly).
- Adopt canonical “build prompt first, then call oracle” step shape as the standard emission pattern.
Out of Scope:
- Not provided
Current Behavior (Actual):
- `oracle ...` command is terminated by a newline, then `-p "$(cat <<'PROMPT' ...)"` appears on the next line without `\`, so Bash treats `-p` as a command and fails.
Expected Behavior:
- Generated bash steps never emit orphaned flag lines (e.g., `-p`, `-f`, `--prompt`) that can be interpreted as standalone shell commands.
- Generated `oracle` invocations are either a single logical line or correctly continued (without inline comments breaking continuation).
Reproduction Steps:
1. Run the generated pack `oracle-pack-2026-01-08-tickets-direct.md`.
2. Observe the step where `-p` begins a new line without a continuation and the shell errors.
Requirements / Constraints:
- “never put `-p/--prompt` (or any flag) on its own line”
- “no inline comments at end of an `oracle ...` line”
- If wrapping is necessary: require explicit `\` continuations and disallow comments on continued lines.
Evidence:
- Error: `bash: line 59: -p: command not found` + `exit status 127`
- Pattern described: `-p "$(cat <<'PROMPT' ...)"` on next line without `\` in `oracle-pack-2026-01-08-tickets-direct.md`
Open Items / Unknowns:
- Exact location(s) of the emitting template(s): Unknown / Not provided
- Whether multiple generators/templates emit the pattern beyond tickets-direct: Unknown / Not provided
Risks / Dependencies:
- Risk: Partial fixes (only adding `\`) may regress if inline comments or formatting are reintroduced.
Acceptance Criteria:
- Generated packs do not contain any step where a line begins with `-p` / `--prompt` intended as a continuation of `oracle` without an explicit safe structure.
- Running the regenerated tickets-direct pack no longer produces `-p: command not found` / `exit status 127` for the previously failing steps.
- Generator output follows one of:
  - prompt built first + `oracle ... --prompt "$prompt"` as a single logical command, OR
  - explicit `\` continuation with no inline trailing comments on continued lines.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “`bash: line 59: -p: command not found` + `exit status 127`”
- “the `-p "$(cat <<'PROMPT' ...)"` part is on the next line without a line-continuation (`\`)”
- “Permanent template rule: never put `-p/--prompt` (or any flag) on its own line”
```

```ticket T2
T# Title: Add validator guardrails (bash-lint + orphaned-flag detection) to fail fast
Type: enhancement
Target Area: `oraclepack validate` (after extracting the single `bash` fence)
Summary:
- Even with a safer generator, regressions can reintroduce orphaned flag lines that only fail at runtime.
- Add validation checks that detect bash syntax issues and the specific “orphaned flag line” class before execution.
- Validation should clearly fail on suspicious standalone flag lines unless safely continued.
In Scope:
- Run `bash -n` against the extracted bash step script(s) as a syntax sanity check.
- Run `shellcheck` static analysis on the extracted bash script(s).
- Implement the custom “orphaned flag line” detector:
  - Fail if a non-heredoc line matches `^\s*-(p|f)\b` or `^\s*--(prompt|file|write-output)\b`
  - Unless the previous non-empty line ends with a legal continuation (`\`, `|`, `&&`, `||`, `(`, etc.)
Out of Scope:
- Making `validate` mandatory in all run paths (handled separately)
Current Behavior (Actual):
- Not provided
Expected Behavior:
- `oraclepack validate` fails fast with a clear error when a pack contains likely orphaned flag lines (e.g., `-p ...`) outside permitted continuation contexts.
- `oraclepack validate` reports bash syntax issues before execution.
Reproduction Steps:
1. Create/modify a pack step where `-p` is on its own line without a valid continuation.
2. Run `oraclepack validate`.
Requirements / Constraints:
- Checks are added “after extracting the single `bash` fence”.
- Orphaned-flag detector must ignore heredoc bodies (“non-heredoc line”).
Evidence:
- “Add these checks to `oraclepack validate` after extracting the single `bash` fence”
- Detector specification (regex + continuation exceptions) provided in ticket text.
Open Items / Unknowns:
- Availability/installation method for `shellcheck` in the execution environment: Unknown / Not provided
- Exact current structure of `oraclepack validate` and how it extracts bash fence: Unknown / Not provided
Risks / Dependencies:
- Risk: False positives if continuation heuristics are too strict; must match the specified allowed continuations.
Acceptance Criteria:
- `oraclepack validate` includes `bash -n` and fails on invalid bash syntax in the extracted script(s).
- `oraclepack validate` runs `shellcheck` and surfaces failures per project policy (pass/fail behavior not specified in ticket text).
- `oraclepack validate` fails when a non-heredoc line begins with `-p`, `-f`, `--prompt`, `--file`, or `--write-output` and the previous non-empty line does not end with an allowed continuation token.
- `oraclepack validate` does not falsely flag valid heredoc prompt bodies.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Add these checks to `oraclepack validate` after extracting the single `bash` fence”
- “`bash -n` syntax check (cheap sanity)”
- “Custom ‘orphaned flag line’ detector… `^\s*-(p|f)\b` … unless… ends with… (`\`, `|`, `&&`, `||`, `(`, etc.)”
```

```ticket T3
T# Title: Make validation unavoidable in normal use (run/TUI) and add CI/pre-commit gate
Type: chore
Target Area: `oraclepack run` execution flow, TUI “Run/Rerun” paths, and repo automation (CI/pre-commit)
Summary:
- Validation guardrails are only effective if they run consistently before pack execution.
- Ensure `oraclepack run` calls `validate` first (or at minimum in TUI Run/Rerun paths).
- Add automated gating so modified/generated packs are validated before being executed/merged.
In Scope:
- Ensure `oraclepack run` always calls `validate` first.
- Ensure TUI “Run/Rerun” paths invoke `validate` first (at minimum).
- Add CI/pre-commit step to run `oraclepack validate` on generated/modified packs.
Out of Scope:
- Implementing the validator checks themselves (handled separately)
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Running a pack via CLI or TUI triggers validation first, preventing execution of invalid packs.
- CI/pre-commit blocks changes that introduce invalid pack structure detectable by `oraclepack validate`.
Reproduction Steps:
1. Introduce a known-invalid pattern (e.g., orphaned `-p` line) into a pack.
2. Attempt to run via `oraclepack run` and via TUI Run/Rerun.
3. Attempt to commit/CI-run with the invalid pack present.
Requirements / Constraints:
- “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI ‘Run/Rerun’ paths).”
- “Add CI/pre-commit: run `oraclepack validate` on any generated/modified pack.”
Evidence:
- The ticket text specifies making validation unavoidable and adding CI/pre-commit gating.
Open Items / Unknowns:
- Existing CI/pre-commit tooling and where to hook validation: Unknown / Not provided
- Exact TUI entrypoints for Run/Rerun: Unknown / Not provided
Risks / Dependencies:
- Depends on `oraclepack validate` providing the intended guardrails to justify making it mandatory.
Acceptance Criteria:
- `oraclepack run` invokes `validate` before executing pack steps.
- TUI Run/Rerun paths invoke `validate` before execution (at minimum).
- CI/pre-commit configuration exists to run `oraclepack validate` on generated/modified packs and fails on validation errors.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Make validate unavoidable in normal use”
- “Ensure `oraclepack run` always calls `validate` first (or at minimum in TUI ‘Run/Rerun’ paths).”
- “Add CI/pre-commit: run `oraclepack validate` on any generated/modified pack.”
```
