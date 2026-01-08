Direct answer (system invariants + where to enforce them)

Invariant: Action Packs must be “oraclepack-ingestible”: a single bash code fence containing numbered step headers # NN) (and the broader workflow expects a fixed step contract). Enforce in a pack parser/validator (structural validation before execution) and in pack generators/templates (to prevent emitting invalid shapes). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: Step execution semantics are “literal shell”: each step is executed via bash -lc ..., in the project root; oraclepack does not change working directory to out_dir. Enforce by documenting as a non-changing contract, and by adding a “preflight” check in pack Step 01 if any later step assumes paths relative to out_dir. 

Oraclepack_Compatibility_Issues

Invariant: Oraclepack’s override/flag injection applies only to commands that begin with oracle (regex anchored to ^(\\s*)(oracle)\\b); tm/task-master, codex, gemini run directly and do not inherit oracle-specific transforms. Enforce in the override injection stage (explicitly scope the matcher to oracle unless/until dispatcher is extended) and in docs/TUI messaging to prevent incorrect user expectations. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: Override validation behavior in the TUI targets only detected oracle invocations (it runs oracle --dry-run summary and skips steps with no oracle invocations). Enforce in the TUI validation pipeline (make the “oracle-only validation” rule explicit) and ensure validate output clearly distinguishes “not validated (non-oracle)” vs “validated”. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: ROI filtering can skip steps that lack ROI= metadata when a threshold > 0 is used. Enforce in pack generation (always emit ROI metadata if packs are intended to be filtered) and in runtime (warn when filtering is enabled but steps have no ROI fields). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: Placeholder steps do not “magically dispatch” external agents—Steps 08–20 are effectively notes unless the step body contains real commands; if you add codex/gemini, oraclepack will attempt to run them as-is. Enforce via template hygiene (placeholders clearly marked; “agent-mode” packs replace placeholders with real commands) and via validation/docs (set expectations: “no commands, nothing happens”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: Non-oracle tool availability and interactivity are execution blockers: missing binaries fail steps; interactive CLIs will block unless headless/non-interactive flags are used. Packs that include codex/gemini must include availability guards and skip behavior (e.g., command -v ...). Enforce inside the pack steps themselves (guards + deterministic skip semantics) and optionally in oraclepack preflight (detect missing binaries before running). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Invariant: Certain steps are expected to produce specific on-disk artifacts (e.g., .oraclepack/ticketify/_tickets_index.json, _actions.json, _actions.md, .taskmaster/docs/tickets_prd.md, plus later .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md/gemini-review.json, PR.md). Enforce by (a) pack-level preflight asserting prerequisites exi_
