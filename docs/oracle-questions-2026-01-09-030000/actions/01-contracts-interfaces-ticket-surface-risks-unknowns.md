Risks/unknowns

The tickets do not specify which overrides/flags should apply to each non-oracle tool (tm/task-master, codex, gemini) or what “validation” means for them; implementing anything beyond broadened detection risks incorrect/partial behavior. 

Oraclepack_Compatibility_Issues

“Headless/non-interactive” invocation details for codex/gemini are not specified; interactivity is explicitly called out as a risk (runs can block), and missing binaries on PATH can hard-fail steps without guards/skip semantics. 

Oraclepack_Compatibility_Issues

The intended default for verification is underspecified: Step 11 could be codex exec, Gemini diff review, or both; similarly, whether Steps 12–13 should be modified is not stated. 

Oraclepack_Compatibility_Issues

The user-facing selection surface for taskify “agent-mode” is not specified (CLI flag vs TUI option vs config), and the exact step slot to swap is not specified. 

Oraclepack_Compatibility_Issues

Where the “clarify current execution semantics/failure modes” documentation should live (README vs oraclepack-tui.md vs TUI help text) is not provided, which affects discoverability and backward-compat messaging. 

Oraclepack_Compatibility_Issues

The tickets do not provide exact code locations for the current detection/override-validation machinery (they describe behavior like an oracle-anchored matcher and oracle-only validation, but not file paths), so the scope of required CLI/TUI/internal refactors is uncertain. 

Oraclepack_Compatibility_Issues

Inference: Several on-disk artifact paths are named as expected outputs (e.g., .oraclepack/ticketify/*, ticket-action-pack.state.json, ticket-action-pack.report.json), but the tickets don’t define schemas/versions for these artifacts, increasing drift risk for any downstream tooling that consumes them. 

Oraclepack_Compatibility_Issues

Reproduction steps are not provided across the split tickets, making regression testing of new/changed public behavior (CLI/TUI validation, pack execution outcomes) harder to anchor. 

Oraclepack_Compatibility_Issues
