Direct answer

Document (CLI/TUI user-facing) the current Action Pack execution contract: each step runs as shell via bash -lc ..., and oraclepack’s special handling (override injection/validation) applies only to commands beginning with oracle (non-oracle tools run directly). Evidence: Oraclepack_Compatibility_Issues.md (“oraclepack executes each step as shell (bash -lc ...) and only applies oracle-specific behavior to oracle-prefixed commands… non-oracle tools (tm/task-master, codex, gemini) run directly”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Extend the dispatcher’s command-detection surface beyond oracle-prefixed commands (currently anchored to ^(\\s*)(oracle)\\b) so steps containing tm/task-master, codex, and/or gemini are no longer excluded from the override/validation pipeline solely due to prefix mismatch; preserve existing oracle behavior. Evidence: Oraclepack_Compatibility_Issues.md (T2: “injects flags into commands that begin with oracle… regex anchored to ^(\\s*)(oracle)\\b… does not match tm, codex, gemini… Must preserve existing oracle command behavior”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Extend/adjust TUI override validation so it doesn’t only run oracle --dry-run summary for detected oracle invocations and skip steps without oracle commands (i.e., make override validation aware of non-oracle tool steps, per the dispatcher changes). Evidence: Oraclepack_Compatibility_Issues.md (“TUI override validation… runs oracle --dry-run summary… steps without oracle invocations are skipped”; T2 scope calls for updating override validation beyond oracle-only targeting). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Replace placeholder steps in ticket-action-pack.md (explicitly 09–13 and 16) with headless gemini + non-interactive codex exec automation, standardizing new output artifacts under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md and/or gemini-review.json, PR.md). Back-compat constraint: the artifact paths are part of the step contract (don’t rename/move without versioning). Evidence: Oraclepack_Compatibility_Issues.md (T3 “Replace… placeholder steps… Step 09 writes .oraclepack/ticketify/next.json… Step 10 writes codex-implement.md… Step 11 writes codex-verify.md and/or gemini-review.json… Step 16 writes PR.md; “Must write the artifacts to the paths specified”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Make tool-availability guards and “skip” behavior part of the public operational contract for these new non-oracle steps (e.g., command -v ... checks), and explicitly treat “interactive CLI blocks” as a first-class hazard; back-compat constraint: default runs should not start failing/hanging just because optional tools aren’t installed or are interactive. Evidence: Oraclepack_Compatibility_Issues.md (“If CLI isn’t installed/on PATH → step fails; if interactive → blocks”; T3 requires “command-availability guards and ‘skip’ behavior… to avoid hard failures when tools are missing”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Add an “agent-mode” option to taskify-generated Action Packs (suggested mode=codex / mode=gemini) that swaps the existing autopilot entrypoint step with an agent-implementation step, without changing the schema/step-count contract; back-compat constraint: keep the “20-step contract intact” and keep existing default behavior unchanged unless mode is explicitly selected. Evidence: Oraclepack_Compatibility_Issues.md (T4 “Add ‘agent-mode’… mode=codex / mode=gemini… keep the 20-step contract intact by swapping… step slot”). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Preserve the Action Pack ingestibility contract as a hard interface boundary (single ```bash fence, # NN) step structure / fixed-step schema), even as new automation is inserted into placeholder slots. Evidence: Oraclepack_Compatibility_Issues.md (“Action Pack is ‘oraclepack-ingestible’ (single bash fence, # NN) steps)”; “keep the 20-step contract intact”).

Oraclepack_Compatibility_Issues
