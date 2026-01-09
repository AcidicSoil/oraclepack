Direct answer (edge cases + minimal UX behavior)

Cancel during a step (Ctrl+C / interrupt): treat as “interrupted”, persist partial logs, and persist run state so the user can resume without guessing what ran. This is required because each step is executed as a shell script (bash -lc ...), so interrupts will happen at the process level and must be reflected in state/report, not just in the UI. 

Oraclepack_Compatibility_Issues

Cancel while a non-oracle CLI is running (codex/gemini/tm): surface a specific warning that these tools are executed “as-is” (not routed through oraclepack’s oracle-only dispatcher), and that interactive tools may block indefinitely; cancellation must not leave the TUI thinking the step “completed”. 

Oraclepack_Compatibility_Issues

Resume after cancel: the resume UX must clearly indicate what will re-run (only failed/interrupted steps, not already-successful ones), and it must work even when some steps are “non-oracle” and therefore not part of oraclepack override/validation. 

Oraclepack_Compatibility_Issues

Partial runs caused by missing binaries (codex/gemini not installed): steps that are intended to be optional automation should “skip” (not hard-fail) when command -v checks fail, and the UI should label them as “skipped: tool not available” to keep the run resumable and comprehensible. 

Oraclepack_Compatibility_Issues

Partial runs caused by interactivity (tool prompts for input): minimal UX is (a) warn before running a step that is likely to block, and (b) provide a consistent cancel path that records the step as “blocked/interrupted” rather than “running forever”. 

Oraclepack_Compatibility_Issues

Resume when ROI filtering changes what’s “visible”: steps can be silently excluded when ROI thresholding is used, especially if headers lack ROI=. Minimal UX must explicitly show “skipped by filter” (and why) so users don’t confuse filtering with completion. 

Oraclepack_Compatibility_Issues

Resume when overrides wizard was used: because override injection/validation is currently oracle-only, the UI must make it impossible to assume codex/gemini/tm steps were validated/overridden, and must preserve that distinction on resume (otherwise users will infer safety guarantees that do not exist). 

Oraclepack_Compatibility_Issues

Placeholder/“note” steps vs executable steps: steps that only echo guidance should be visually distinguishable from steps that produce required artifacts; otherwise “run-all” followed by “resume” becomes confusing (many “successful” steps that did nothing). This matters because the current plan explicitly calls out placeholder steps (08–20) and replacing some with real automation. 

Oraclepack_Compatibility_Issues

Minimal required status taxonomy for UX: success / failed / skipped (with reason: ROI filter, tool-missing, user-skip) / interrupted (user cancel, crash). This is the minimum needed to support “cancel → resume” and “partial runs” without ambiguity given the documented failure modes.

Risks/unknowns

The tickets describe execution semantics and failure modes, but do not specify current signal handling details (does SIGINT propagate to the child process group; does the runner always flush logs/state atomically). 

Oraclepack_Compatibility_Issues

The desired “resume contract” is not fully specified (resume from last success vs resume selected step; whether prelude is re-run; how to treat interrupted prelude). 

Oraclepack_Compatibility_Issues

Override/validation behavior for non-oracle tools is explicitly an open design gap; UX must avoid implying guarantees until dispatcher changes exist.

The “skip behavior” for missing codex/gemini is referenced as required, but the exact skip contract (exit code, status file semantics, report fields) is not provided. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Add a single integration test fixture pack with a long-running step (e.g., sleep 30), run it under oraclepack run in a way the test can send SIGINT mid-step, assert the state/report mark the step as “interrupted”, then run oraclepack run --resume and assert it continues from the interrupted step (and that ROI-filtered steps are labeled “skipped by filter”, not “pending”).

I
