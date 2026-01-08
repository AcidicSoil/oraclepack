Direct answer (validation boundaries that must exist; with minimal validation plan)

Pack structure boundary (parser-level): reject any “action pack” that is not oraclepack-ingestible: exactly one fenced bash block, and step headers in the # NN) form (sequential). This is explicitly treated as a global constraint in the ticket content and is the assumed execution contract. 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

Execution semantics boundary (runner-level): treat every step body as literal shell executed via bash -lc … from the project root, with no special routing unless the command begins with oracle. This must be documented and enforced as a “you get what you wrote” boundary to prevent false assumptions about dispatching/wrapping non-oracle tools. 

Oraclepack_Compatibility_Issues

Oracle-only transform/override boundary (validator + injector-level): override injection and “override validation” must only apply to commands that begin with oracle (regex anchored to oracle per ticket), and steps without oracle invocations must not be assumed validated by oracle dry-runs. Validation must therefore (a) detect oracle invocations precisely, and (b) warn (or at least report) that non-oracle steps are outside oracle override validation. 

Oraclepack_Compatibility_Issues

Tooling availability boundary (pack authoring + runtime guard-level): because non-oracle CLIs (tm/task-master, codex, gemini) run directly, packs must either (a) require them explicitly (hard fail), or (b) guard with command -v … and implement “skip” behavior. The ticket explicitly calls out “missing binary on PATH fails” and “interactive blocks” as common failure modes; validation should catch missing guards for steps intended to be optional/headless. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Ticket parsing boundary (ticketify step-level): before any “ticketify” logic runs, validate .tickets/ exists (and contains readable ticket files) because missing .tickets/ is called out as a primary failure note; then validate the expected index artifact(s) are produced (e.g., .oraclepack/ticketify/_tickets_index.json). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Artifact contract boundary (step-output-level): steps that claim to produce .oraclepack/ticketify/* outputs must be validated for deterministic paths and for actually writing the promised files (e.g., next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md). The ticket’s acceptance criteria is already framed this way; validation should mirror it. 

Oraclepack_Compatibility_Issues

ROI filtering boundary (filter-level): if ROI-based filtering is enabled, steps without ROI= may be skipped; packs that rely on certain steps must include ROI tags (or validation must fail/warn when ROI thresholding would skip required steps). 

Oraclepack_Compatibility_Issues

Minimal validation plan (smallest set that covers the boundaries above):

Static pack validation: run oraclepack validate <pack.md> for representative packs (gold pack, ticket-action-pack, and a mixed-tool pack). 

oraclepack-llms-full

Fixture-based parser regression: validate packs with (a) multiple code fences, (b) missing/duplicated/unsorted # NN) headers, (c) non-sequential numbering, and confirm deterministic errors (prevents schema drift). 

oraclepack-llms-full

Mixed-tool step cove
