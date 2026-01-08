Direct answer (1–10 bullets, evidence-cited)

Today, “background” work in these Action Packs is not a separate job system: oraclepack executes each step as a shell script (bash -lc ...), so failures are synchronous step failures (the step exits non-zero) unless the step body explicitly catches/handles them. 

Oraclepack_Compatibility_Issues

For non-oracle tooling (tm/task-master, codex, gemini), oraclepack does not apply oracle-specific override/validation; those commands run directly, so any retries/idempotency must be implemented either inside the tool itself or in the step script wrapper. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

The tickets explicitly call out “missing binary on PATH fails the step” and “interactive CLI blocks”, and recommend command-availability guards (command -v ...) plus “skip” behavior to avoid hard failures when tools are missing; treat this as the baseline failure-handling pattern for codex/gemini. 

Oraclepack_Compatibility_Issues

Retries: where API/background polling exists (Oracle API “background response” flow), the client logic already uses automatic retry for retryable transport errors with exponential backoff and a max-wait timeout; mirror this pattern in step-level wrappers for transient failures (network/API flaps) when invoking headless gemini / codex exec. 

oraclepack-llms-full

Idempotency: the automation steps are required to write deterministic artifacts at fixed paths (e.g., .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md); to make reruns safe, write outputs atomically (tmp file + rename) and overwrite in place, since downstream steps depend on those canonical paths existing. 

Oraclepack_Compatibility_Issues

Poison messages (step-runner analogue): define a “poison step” as one that fails repeatedly with the same non-transient reason (e.g., missing CLI, interactive prompt, invalid config). The ticket’s prescribed approach is to convert some of these into “skip” outcomes (guard + skip) rather than repeated hard-fails; for non-skippable steps, persist a failure record in the pack’s state/report outputs so it’s visible and resumable. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Observability hooks should be file-first because the workflow already hinges on emitted artifacts: capture (a) tool-availability checks and skip reasons, (b) step status (started/completed/failed/skipped), (c) timings and exit codes, and (d) pointers to generated artifacts like .oraclepack/ticketify/* and the pack state/report JSON mentioned in the ticket. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

For long-running “background-like” operations (API background responses), log status transitions (“queued”/“in_progress”/reconnected) and retry delays; this is already present in the background polling logic and is the concrete model to replicate for gemini/codex headless steps (emit to stdout and/or a JSONL events file). 

oraclepack-llms-full

Risks/unknowns (bullets)

The tickets do not specify whether oraclepack should implement runner-level retries (re-run a failing step automatically) vs step-script retries only; adding runner-level retries could change failure semantics. 

Oraclepack_Compatibility_Issues

No explicit policy is provided for which failures are “skippable” vs “fatal” across Steps 09–13 and 16; incorrect classification could silently skip important work. 

Oraclepack_Compatibility_Issues

“Poison message” handling is underspecified for this architecture (step-based execution, not
