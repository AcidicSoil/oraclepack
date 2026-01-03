Title:

* Fix oraclepack pack execution error: `/bin/bash: line 2: GenerateReport: command not found`

Summary:

* Running an `oraclepack` Markdown pack (referenced as `oraclepack.md`) produces a bash error indicating `GenerateReport` is being executed as a shell command.

    Bash error explanation

* The reported cause is that the pack’s generated bash script contains a bare `GenerateReport` token (e.g., on line 2), which bash interprets as a command name that does not exist in the shell environment.

    Bash error explanation

Background / Context:

* `oraclepack` executes each prelude/step by invoking `/bin/bash -lc "<your step script>"` (per assistant).

    Bash error explanation

* The token `GenerateReport` is also a Go function name in the codebase (`internal/report/generate.go`), but it is not a shell command (per assistant).

    Bash error explanation

Current Behavior (Actual):

* Pack execution emits: `"/bin/bash: line 2: GenerateReport: command not found"`.

    Bash error explanation

* Error may appear twice if both a prelude and a step include `GenerateReport` in the executed script (per assistant).

    Bash error explanation

Expected Behavior:

* Pack prelude/steps execute without attempting to run `GenerateReport` as a shell command.

    Bash error explanation

* If `GenerateReport` is intended as a visual header/label, it should be rendered as output (e.g., via `echo`) or commented out, not executed.

    Bash error explanation

Requirements:

* Remove or comment out any bare `GenerateReport` line inside the pack’s bash code block (or generated script) so bash does not treat it as a command.

    Bash error explanation

* If a visible section header is needed, replace the bare token with `echo "GenerateReport"` (per assistant).

    Bash error explanation

* Ensure both prelude and step scripts (if both exist) are updated to avoid duplicate occurrences of the error.

    Bash error explanation

Out of Scope:

* Not provided.

Reproduction Steps:

1. Run `oraclepack` on a pack whose executed bash script contains a bare `GenerateReport` token on line 2 (per assistant).

    Bash error explanation

2. Observe `/bin/bash: line 2: GenerateReport: command not found`.

    Bash error explanation

3. (Optional) If both prelude and step contain it, observe the error twice (per assistant).

    Bash error explanation

Environment:

* Shell: `/bin/bash -lc` execution path (per assistant).

    Bash error explanation

* OS: Unknown

* oraclepack version: Unknown

* Pack file name/path: `oraclepack.md` (exact path unknown).

    Bash error explanation

Evidence:

* Error text (verbatim): `/bin/bash: line 2: GenerateReport: command not found`.

    Bash error explanation

* Referenced code location: `internal/report/generate.go` (Go function `GenerateReport`, per assistant).

    Bash error explanation

* Conversation source: “Bash error explanation.md”.

    Bash error explanation

Decisions / Agreements:

* Treat the error as originating from bash interpreting `GenerateReport` as a command (per assistant).

    Bash error explanation

* Fix by removing/commenting the bare token or converting it to `echo` output (per assistant).

    Bash error explanation

Open Items / Unknowns:

* Exact contents of `oraclepack.md` bash fence and the specific line where `GenerateReport` appears.

    Bash error explanation

* Whether the token appears in the prelude, the step(s), or both.

    Bash error explanation

* Exact oraclepack command used and full console output beyond the single error line.

Risks / Dependencies:

* Pack correctness depends on how `oraclepack` generates/executes the script via `/bin/bash -lc` and on the pack author keeping non-commands commented/escaped.

    Bash error explanation

Acceptance Criteria:

* Running `oraclepack` on the pack no longer produces `command not found` for `GenerateReport`.

    Bash error explanation

* Any intended section labels are either commented (`# GenerateReport`) or printed (`echo "GenerateReport"`) and do not affect execution.

    Bash error explanation

* If both prelude and steps exist, neither emits the error (no duplicate occurrences).

    Bash error explanation

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* bash

* oraclepack

* pack-authoring

* execution

* command-not-found

---
