Direct answer (1–10 bullets, evidence-cited)

There is no explicit “oraclepack-managed” background job system implied; packs execute step bodies as shell (bash -lc ...) in-order, so any async/long-running behavior is delegated to the invoked tools inside a step (e.g., oracle, tm, codex, gemini). 

Oraclepack_Compatibility_Issues

Oracle CLI “background mode” (Responses-style async) is an implied long-running operation: requests may return queued/in_progress and Oracle must poll until completed, with retry/backoff and a max-wait timeout. Interface boundary is effectively pollBackgroundResponse(params) with client, responseId, initialResponse, log, wait, now, maxWaitMs (plus retrieveBackgroundResponseWithRetry). 

oraclepack-llms-full

Task Master “autopilot” is an implied long-running (and potentially interactive) operation executed as a pack step. The step is responsible for branch safety + preflight checks, writing a resumable state file (tm-autopilot.state.json), then launching ${TM_CMD} autopilot. Inputs are MODE/TM_CMD/OUT_DIR/PACK_DATE/TAG; output contract includes the state file path. 

oraclepack-llms-full

The tickets imply “agent-mode” replacement of autopilot with headless gemini + non-interactive codex exec steps (select → implement → verify → draft PR). These are long-running operations that must be non-interactive to avoid blocking oraclepack run. The interface is file-based artifacts under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md) with command-availability guards (command -v ...) and “skip” semantics when tools are missing.

If you want oraclepack-level responsibility for async, the implied seam is the “dispatcher/override/validation pipeline” extension: broaden command detection beyond oracle so tm/codex/gemini steps can participate in consistent validation/guardrails (but the ticket does not define the per-tool override/validation semantics). 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

Non-oracle tool steps can block indefinitely if interactive; the ticket explicitly flags interactivity as a failure mode and requires headless/non-interactive variants. 

Oraclepack_Compatibility_Issues

The override/validation behavior for tm/codex/gemini is underspecified (which flags apply, what “validation” means per tool), so implementing oraclepack-managed async/dispatch risks being partial or incorrect. 

Oraclepack_Compatibility_Issues

It’s unclear whether the repo contains any first-class job/queue implementations beyond “tool-driven” long-running steps; the existing coverage checklist explicitly says “background jobs: Missing” and requests specific directories if they exist. 

oraclepack-llms-full

Autopilot resumability is only implied (“state file created… should resume if supported”); whether tm autopilot actually supports resume-from-state is unknown. 

oraclepack-llms-full

Next smallest concrete experiment (exactly one action)

Add a “resumable long-running step wrapper” to the autopilot step (write/update a state JSON before/after launch; wrap the tm autopilot invocation with a hard timeout and clear exit codes), then run oraclepack validate + oraclepack run --no-tui on that pack to confirm (a) no indefinite hangs, and (b) reruns observe the state file and behave deterministically. 

oraclepack-llms-full

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

**/jobs/** (explicitly requested by the current coverage checklist) 

oraclepack-llms-full

**/queue/**, **/worker/**, **/cron/**, **/*scheduler*/**, **/*background*/**

Any runtime/service configs that define scheduled work: **/*.yaml / **/*.yml / **/*.toml containing cron, schedule, worker, queue

If Task Master autopilot integration
