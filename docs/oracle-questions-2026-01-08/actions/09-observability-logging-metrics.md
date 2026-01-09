Direct answer (logging/metrics + minimal instrumentation plan)

Pack-generation logs must emit: (a) input discovery summary (ticket roots searched, tickets found), (b) deterministic output paths written (e.g., .oraclepack/ticketify/_tickets_index.json, _actions.json, _actions.md, .taskmaster/docs/tickets_prd.md, .oraclepack/ticketify/tm-complexity.json), and (c) “workdir = project root” (since execution does not switch to out_dir). 

Oraclepack_Compatibility_Issues

Pack-parse/validate logs must include: detected “single bash fence” + parsed step headers (# NN)), step count/order, and any rejection reason (because ingestion is strict and mis-detection causes downstream execution surprises). 

Oraclepack_Compatibility_Issues

Step-execution logs must be per-step, persisted to disk, and include: step number/title, full script body, bash -lc invocation, cwd, start/end timestamps, exit code, and captured stdout+stderr (streamed + saved). This is required because oraclepack executes literal step bodies as shell. 

Oraclepack_Compatibility_Issues

“Skip reason” logging must be first-class: (a) ROI-filter skips when steps lack ROI= and threshold > 0, (b) tool-availability skips via command -v ... guards, and (c) mode-gated steps (e.g., agent-mode/autopilot slots). Without explicit skip reasons, headless runs look like silent no-ops.

Command-detection / override-injection / validation logs must explicitly show: (a) which lines matched the current oracle-anchored regex, (b) what flags were injected, and (c) the result of override validation (oracle --dry-run summary)—because today only oracle-prefixed commands participate, while tm/codex/gemini do not.

External-tool execution logs must record “missing binary” vs “interactive block risk” for tm/codex/gemini, since the tickets call out both as primary failure modes (missing PATH hard-fails unless guarded; interactive CLIs can hang headless runs).

Artifact-verification logging must run after relevant steps and record existence/size/hash for required outputs (at minimum: .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md and/or gemini-review.json, PR.md). These are explicit acceptance criteria; when absent, the logs should say “expected but missing” and point to the producing step.

Minimal metrics should be derivable from the same run report: counters for steps_total/ran/skipped/failed, plus durations (pack_parse_ms, prelude_ms, step_duration_ms by step/tool), and counts of oracle_detected/oracle_validated/non_oracle_steps. The repo already targets state+report artifacts for traceability; extend them to carry these fields rather than inventing a separate telemetry stack.

Minimal instrumentation plan (implementation shape): write (1) per-step log files and (2) a single machine-readable run summary (JSON) that includes per-step Result{exitCode, startedAt, endedAt, logPath} and environment metadata (shell, cwd), enabling headless debugging and CI parsing.

Risks/unknowns

Current ticket bundle provides desired behaviors and failure modes, but not the existing logger/report schema actually implemented (so some recommendations may already exist partially). 

Oraclepack_Compatibility_Issues

Exact override behavior for non-oracle tools (tm/codex/gemini) is explicitly unspecified; metrics/logging can expose gaps, but cannot resolve semantics without a spec.

Step 11 default (“Codex verify vs Gemini review vs both”) remains undecided; instrumentation should capture which path ran and why, but product behavior still needs a default decision.
