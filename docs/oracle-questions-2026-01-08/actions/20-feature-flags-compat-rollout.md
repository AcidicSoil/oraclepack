Direct answer (minimal compat-safe rollout plan + guardrails, fallback, monitoring gates)

Phase 0 (baseline invariants): keep current execution semantics unchanged: oraclepack executes each step via bash -lc ... in the project root, and only applies special handling (flag injection + override validation) to commands that begin with oracle. This is the compatibility baseline you must not break. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Phase 1 (shadow / observe-only): add “dispatcher v2 shadow” that detects steps containing tm/task-master, codex, gemini for reporting/UX warnings, but does not modify execution, inject flags, or expand validation beyond the current oracle-only behavior. This directly addresses the “silently excluded” concern while staying compat-safe. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Phase 2 (opt-in, guarded automation in packs): update ticket-action-pack.md placeholder steps by inserting headless Gemini + non-interactive codex exec only behind an explicit pack mode / flag, and require command-availability guards (command -v ...) with documented “skip” behavior to prevent hard failures when binaries are missing or interactive. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Guardrail: never allow “interactive CLI blocks the run” by default; any Codex/Gemini invocation used in packs must be explicitly non-interactive/headless, and if that cannot be guaranteed in a given environment, the step must skip (and write a skip note) rather than hang. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Guardrail: preserve step-count/schema contracts; for taskify-generated packs, implement mode=codex / mode=gemini by swapping the existing “autopilot entrypoint” step slot (not adding steps), keeping the “20-step contract intact,” and inserting only after Task Master expansion artifacts exist (tasks.json, tm-complexity.json). 

Oraclepack_Compatibility_Issues

Fallback behavior (runtime): if non-oracle overrides/validation are not implemented or misconfigured, fall back to today’s behavior (oracle-only injection/validation) and treat non-oracle tools as literal shell commands; do not attempt partial wrapping that changes semantics unpredictably. 

Oraclepack_Compatibility_Issues

Fallback behavior (pack outputs): for the new automated ticket-action flow, require the artifact-based fallbacks: if gemini is missing, Step 09/16 skip and do not create .oraclepack/ticketify/next.json / PR.md; if codex is missing, Step 10 and Codex-based Step 11 skip and do not create .oraclepack/ticketify/codex-implement.md / codex-verify.md (but the run remains usable for the non-agent steps 01–07). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Monitoring gates (acceptance-as-gates): treat the ticket’s “artifact existence” as rollout gates in CI/canary: Step 09 must produce .oraclepack/ticketify/next.json; Step 10 must produce codex-implement.md; Step 11 must produce at least one of codex-verify.md or gemini-review.json; Step 16 must produce PR.md—or, if tools are missing, must produce an explicit “skipped due to missing binary” record instead of failing. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Monitoring gates (operational): explicitly track and alert on the known failure modes: (a) step failures due to missing binaries on PATH, (b) runs that stall due to interactivity, and (c) ROI-filter skipping of steps without ROI= when thresholds are used (rollout guard: ensure any newly “required” steps include ROI= or document running with ROI threshold = 0 during canary). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Risks/unknowns

The ticket set does not specify which overrides/flags should apply to each non-oracle tool (tm, codex, gemini) or what “validation” means for them; implementing anything beyond shadow-detection without this risks incorrect behavior. 

Oraclepack_Compatibility_Issues
