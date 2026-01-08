Direct answer (external integrations, required config/contract changes, failure/timeout, rollout)

Task Master CLI integration is assumed: Action Packs invoke task-master / tm for PRD parsing, complexity analysis, and task expansion; this implies the runtime environment must have Task Master installed and configured, and the pack writes/reads Task Master artifacts (e.g., .taskmaster/docs/tickets_prd.md, .oraclepack/ticketify/tm-complexity.json). 

Oraclepack_Compatibility_Issues

Codex CLI integration is implied for “implementation” and “verification”: placeholder steps (notably Step 10 and optionally Step 11) are intended to run codex exec non-interactively and emit .oraclepack/ticketify/codex-implement.md and .oraclepack/ticketify/codex-verify.md. 

Oraclepack_Compatibility_Issues

Gemini CLI integration is implied for “selection” and “PR drafting”: placeholder steps (notably Step 09 and Step 16, and optionally Step 11) are intended to run headless gemini and write .oraclepack/ticketify/next.json, .oraclepack/ticketify/PR.md, and optionally .oraclepack/ticketify/gemini-review.json. 

Oraclepack_Compatibility_Issues

Oracle CLI integration is the only integration that currently receives oraclepack’s special handling: oraclepack injects flags / performs validation only for commands beginning with oracle (regex anchored to ^(\\s*)(oracle)\\b), while tm/task-master, codex, gemini run as raw shell commands. 

Oraclepack_Compatibility_Issues

Required contract change (dispatcher/validation): extend oraclepack’s command detection + override/validation pipeline beyond oracle-prefixed commands so steps containing tm/task-master, codex, gemini are no longer excluded from override/validation purely due to prefix mismatch; must preserve existing oracle behavior. 

Oraclepack_Compatibility_Issues

Required pack-template change (ticketify Action Pack): replace the placeholder/echo-only steps (08–20) with real, headless commands in the suggested slots (09–13 and 16) and enforce deterministic output paths under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md). 

Oraclepack_Compatibility_Issues

Required generator change (taskify Action Packs): add an “agent-mode” switch (suggested mode=codex / mode=gemini) that swaps an existing step-slot (not an added step) “right after Task Master expands tasks,” keeping the 20-step contract intact. 

Oraclepack_Compatibility_Issues

Failure/timeout behavior implied by tickets: every step is executed as a literal bash -lc ... script; if a CLI is interactive it can block waiting for input, and if a binary is missing from PATH the step fails unless the step explicitly guards/skips via command -v ... checks. 

Oraclepack_Compatibility_Issues

Failure behavior implied by tickets: ROI filtering can skip steps that lack an ROI= annotation when a threshold > 0 is used (“ROI filter gotcha”), which can silently skip non-annotated automation steps unless you standardize ROI tags for them. 

Oraclepack_Compatibility_Issues

Minimal compat-safe rollout sequence: (1) document current semantics + failure modes (no behavior change), then (2) update ticket-action-pack.md placeholder steps using guarded command -v + non-interactive/headless flags (no oraclepack core change), then (3) introduce dispatcher/validation broadening behind an opt-in flag or “mode” to preserve default oracle-only semantics.

Risks/unknowns

The tickets explicitly note the spec gap: “exact desired behavior for applying overrides” to tm/task-master, codex, gemini (which flags, what validation) is not provided, so d
