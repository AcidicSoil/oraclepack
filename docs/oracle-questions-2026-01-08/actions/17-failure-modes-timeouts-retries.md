Direct answer (1–10 bullets, evidence-cited)

Oraclepack executes each pack step as a literal shell script (bash -lc ...), so timeout/retry behavior must be implemented either (a) inside the called tool (oracle, task-master, codex, gemini) or (b) explicitly wrapped in the step body; oraclepack does not “route everything through oracle,” and only oracle-prefixed commands receive oraclepack’s injection/validation behavior. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

For oracle ... steps: do not add an outer “rerun the whole command” retry loop by default (billing/cost risk); instead rely on Oracle’s internal handling: background responses are polled until completion or maxWaitMs, with retry-on-transport during background retrieval and a hard client-timeout when the overall wait budget is exceeded. 

oraclepack-llms-full

Set wall-clock timeouts to respect Oracle’s existing defaults: browser runs show a default timeoutMs of 1,200,000ms (20m) and an inputTimeoutMs on the order of tens of seconds; if you add a shell timeout, set it ≥ 25m for oracle browser/background steps so you don’t preempt Oracle’s own timeout. 

oraclepack-llms-full

For non-oracle CLIs (tm/task-master, codex, gemini): treat “missing binary on PATH” as a first-class, non-retryable Dependency failure; the ticket explicitly calls for command -v ... guards plus “skip” behavior to avoid hard failures when these tools are optional in placeholder steps. Operator action: install tool / fix PATH, then re-run the step. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Treat “interactive CLI blocks waiting for input” as a non-retryable Execution-Mode failure; enforce headless/non-interactive flags and add a shell-level timeout for these steps (e.g., timeout 15m …) so runs cannot hang indefinitely. Operator action: re-run with correct headless/non-interactive configuration. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Classify failures into: (A) Dependency (binary missing), (B) Configuration/Input (missing .tickets/, missing provider keys/config), (C) Interactive/Blocked, (D) Transient Transport (network/connection/timeouts), (E) Remote Throttle/5xx (retryable with backoff), (F) Remote Permanent (auth/validation/permission; non-retryable), (G) Postcondition/Artifact missing, (H) Skipped/Filtered. The ticket calls out .tickets/ missing and provider key/config gaps as common failure points. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

 

claude-task-master-llms

Retry policy: only for (D) and (E). Use bounded exponential backoff with a hard cap (and never exceed the step’s wall-clock budget); do not retry (A)(B)(C)(F). This matches the “retry transport, but stop on non-retryable” shape shown in Oracle’s background retrieval loop. Operator action: if retries exhaust, inspect logs/state, fix network/provider, rerun the step once. 

oraclepack-llms-full

Operator actions for (G) Postcondition failures: treat “command exited 0 but expected artifact path missing” as actionable because the ticket has explicit acceptance criteria for .oraclepack/ticketify/* outputs; action is to open the step output/log, confirm working directory assumptions (oraclepack runs in project root), then rerun step after fixing paths/permissions. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Operator actions for (H) Skipped/Filtered: if ROI filtering is enabled, steps without ROI= may be skipped; treat that as “not executed,” not a failure. Operator action: re-run without the ROI threshold/filter or ensure steps include ROI metadata. 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

The tickets do not specify numeric timeout budgets or retry counts per tool (task-master vs codex vs gemini), only that interactivity can block and missing binaries should be handled with guards/skip behavior. 

Oraclepack_Compatibility_Issues
