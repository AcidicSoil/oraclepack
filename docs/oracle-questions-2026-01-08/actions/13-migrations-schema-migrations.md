Direct answer (1–10 bullets, evidence-cited)

No “pack format” migration is required (and should be avoided): the Action Pack must remain oraclepack-ingestible (single bash fence, # NN) steps), and the taskify “agent-mode” must keep the “20-step contract intact” by swapping a step slot rather than adding steps. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

A behavioral/compat migration is required in oraclepack’s dispatcher: today, override injection and override validation only target commands that begin with oracle (regex anchored to ^(\\s*)(oracle)\\b) and validation runs oracle --dry-run summary, skipping steps without oracle invocations; the ticket requires extending this beyond oracle to tm/task-master, codex, and gemini while preserving existing oracle behavior. 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

That dispatcher expansion likely forces a state/config schema migration: oraclepack tracks persistent run state (with schemaVersion) and per-step results; once “overrides/validation” is multi-tool, the persisted state needs room for tool-specific override selections/validation outcomes, with a schema bump + auto-upgrade defaults for old state files. 

oraclepack-llms-full

 

Oraclepack_Compatibility_Issues

A CLI/config surface migration is required for taskify generation: add a mode switch (explicitly suggested mode=codex / mode=gemini) that selects an alternative implementation path after Task Master expansion, while keeping the pack schema/step count stable and defaulting to current behavior when unset. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

A content/template migration is required for ticket-action-pack.md: replace placeholder steps (called out as 09–13 and 16) with headless Gemini selection, non-interactive codex exec implementation, verification, and PR drafting, writing the specified .oraclepack/ticketify/* artifacts, and guarding for missing binaries via “skip” behavior. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

A compat adjustment is required for ROI filtering: because ROI thresholding may skip steps that lack an ROI= tag, any new/repurposed Codex/Gemini steps should either (a) include ROI tags in their headers, or (b) be documented as “may be skipped under ROI filtering,” to avoid surprising partial runs. 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

A path/working-directory compatibility constraint must be carried into updated templates: steps run via bash -lc ... in the project root and oraclepack does not change WorkDir to out_dir, so any new steps must use explicit paths (e.g., .oraclepack/ticketify/...) and not assume cd "$out_dir" happened. 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

The tickets do not define what “override injection” means for each non-oracle tool (which flags are injectable/valid, and how to validate them), so a naive “apply oracle overrides everywhere” approach risks breaking tool invocations. 

Oraclepack_Compatibility_Issues

Codex/Gemini “headless/non-interactive” invocation details are not specified; interactivity is explicitly a risk (blocking the run). 

Oraclepack_Compatibility_Issues

It’s unclear whether non-oracle commands should participate in the same overrides wizard UX, or be handled as separate “tool adapters” with distinct validation semantics. 

Oraclepack_Compatibility_Issues

Exact code locations for the current detection regex, override validation pipeline, and state/report serialization are not provided in the tickets, so migration impact on persisted state cannot be precisely scoped from the provided material. 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

Next smallest concrete experiment (exactly one action)

Add a single “mode” switch to the taskify Action Pack generator (defaulting to current behavior), generate one pack in mode=codex, and run oraclepack validate on it to prove (a) the 20-step contract is preserved and (b) the pack remains ingestible even with codex/gemini commands present. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

The oraclepack runner implementation where command detection / override injection / override validation lives (e.g., cmd/oraclepack/**, internal/**/overrides*, internal/**/dispatch*, internal/**/validate*).

The state/report schema and serializers (e.g., `intern
