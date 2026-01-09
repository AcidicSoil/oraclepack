Direct answer (1–10 bullets, evidence-cited)

Define a single “trace context” for every pack run: run_id (unique per run + stable across resume), pack_path, pack_hash, and started_at; persist it in the existing state/report artifacts (e.g., ticket-action-pack.state.json, ticket-action-pack.report.json) so every downstream artifact can be joined back to one run. 

Oraclepack_Compatibility_Issues

Use pack_hash + step_number as the canonical step key, since the pack model already has Pack{path, hash} and Step{number, title} and the run model already tracks per-step results and logPath; extend those records with run_id and attempt (increment on retries). 

oraclepack-llms-full

Propagate correlation across all tools by setting environment variables for every step execution (works because steps execute as bash -lc ...):
ORACLEPACK_RUN_ID, ORACLEPACK_PACK_HASH, ORACLEPACK_PACK_PATH, ORACLEPACK_STEP_NO, ORACLEPACK_STEP_TITLE, ORACLEPACK_STEP_ATTEMPT, ORACLEPACK_LOG_PATH, ORACLEPACK_OUT_DIR. This avoids relying on tool-specific wrappers. 

Oraclepack_Compatibility_Issues

For oracle invocations, additionally inject the trace context into the prompt text (so the LLM output is self-identifying) and/or into the output file preamble; this is especially important because oraclepack’s current special handling/override pipeline only targets commands beginning with oracle. 

Oraclepack_Compatibility_Issues

For non-oracle tools (task-master/tm, codex, gemini), treat the step log (logPath) as the ground-truth join point, because they currently run “directly” and do not inherit oraclepack’s oracle-only injection/validation behavior.

Require every step that produces a machine-readable artifact to embed {run_id, pack_hash, step_no} inside the artifact (JSON metadata block for .json, a short header for .md). This is aligned with the tickets’ emphasis on deterministic, path-specific artifacts under .oraclepack/ticketify/* (e.g., next.json, codex-implement.md, gemini-review.json, PR.md). 

Oraclepack_Compatibility_Issues

When/if you implement the dispatcher changes requested in the tickets (extend detection/overrides beyond oracle), define a per-tool adapter that maps the same trace context into each tool’s native flags (where available) while keeping env-var propagation as the baseline. 

Oraclepack_Compatibility_Issues

Minimum required IDs and how they join:
run_id (primary join key across everything), pack_hash (stable pack identity), step_no (01..20), attempt (retry/resume), invocation_index (nth tool call within a step, if needed), plus logPath (filesystem pointer to raw evidence). 

oraclepack-llms-full

Risks/unknowns (bullets)

The tickets don’t specify whether codex/gemini expose native run/session IDs or non-interactive flags sufficient to avoid blocking; without that, correlation may rely entirely on logs + embedded metadata.

Current oraclepack behavior applies special handling only to oracle-prefixed commands; if you expect uniform injection/validation/tracing for tm/codex/gemini, you need the dispatcher work (and a spec for per-tool override behavior).

The provided materials describe desired state/report models but not the concrete on-disk schema used by your current runner; adding fields requires a versioning/back-compat plan.

Next smallest concrete experiment (exactly one action)

Implement runner-level env-var injection + add run_id to the state/report JSON (persisted once per run, reused on resume), then run oraclepack run --no-tui on ticket-action-pack.md and confirm each step log contains the printed ORACLEPACK_RUN_ID and each produced .oraclepack/ticketify/* artifact embeds {run_id, pack_hash, step_no}.

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

cmd/oraclepack/** (entryp**
