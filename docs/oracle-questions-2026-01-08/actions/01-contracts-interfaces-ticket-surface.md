Direct answer (public surface changes; backwards-compat called out)

Clarify/document Action Pack execution semantics: steps run as bash -lc ..., and oraclepack’s “special handling” (flag injection/override validation) only applies to commands that begin with oracle (today). This is a user-facing CLI/TUI documentation surface change (README and/or oraclepack-tui.md-style docs). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Dispatcher/command-detection contract expansion: update command detection from an oracle-anchored regex (noted as ^(\\s*)(oracle)\\b) so non-oracle tools (tm/task-master, codex, gemini) can be treated as first-class step commands for override/dispatch purposes. Back-compat constraint: preserve existing behavior for oracle ... commands. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Override validation behavior change (TUI/CI surface): today validation runs oracle --dry-run summary on detected oracle invocations and skips steps without oracle invocations; tickets imply extending or restructuring validation so steps containing tm/task-master, codex, gemini are not silently excluded. Back-compat constraint: do not regress the current oracle-only validation flow. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

ticket-action-pack.md content contract change: replace placeholder steps (explicitly 09–13 and 16) with headless gemini + non-interactive codex exec automation, while keeping the pack ingestible (single bash fence, 20-step structure). Back-compat constraint: keep Steps 01–07 semantics unchanged. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

New output artifact interface for the action workflow: introduce/standardize generated files under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md and/or gemini-review.json, PR.md) as “expected outputs” of those steps. Back-compat constraint: paths must match the step examples; don’t move or rename without versioning. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Tool-availability/timeout behavior requirements become part of the pack’s operational contract: add command -v ... guards and documented “skip” behavior to avoid hard failures when codex/gemini are missing; also explicitly call out “interactive CLI blocks” risk. Back-compat constraint: default runs shouldn’t newly fail just because optional tools aren’t installed. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Taskify Action Pack generator surface: add an “agent-mode” option for taskify-generated packs (suggested mode=codex / mode=gemini) that swaps the existing autopilot entrypoint step with an agent implementation step, without changing the 20-step contract. Back-compat constraint: default mode should remain the current behavior; agent-mode must be opt-in. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

“Pack schema” remains a hard compatibility boundary: multiple tickets restate that the Action Pack must remain oraclepack-ingestible (single `bas
