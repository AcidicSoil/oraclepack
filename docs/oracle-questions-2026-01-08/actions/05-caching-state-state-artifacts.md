Direct answer (1–10 bullets, evidence-cited)

Preserve the ticket discovery index: .oraclepack/ticketify/_tickets_index.json (produced by ticket-action-pack.md Step 01). It is explicitly called out as an output artifact and is the natural cache boundary for downstream “actions” inference. 

Oraclepack_Compatibility_Issues

Preserve the canonical actions outputs: _actions.json and _actions.md (produced by Step 02). These are explicitly named outputs; _actions.json is expected to follow a stable, machine-readable schema (metadata + up to 20 items) to support later automation. 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

Preserve the Task Master PRD artifact: .taskmaster/docs/tickets_prd.md (produced by Step 03). This is explicitly named as a written output and is a key handoff artifact into Task Master parsing/expansion. 

Oraclepack_Compatibility_Issues

Preserve Task Master analysis/expansion artifacts, especially .oraclepack/ticketify/tm-complexity.json; and preserve tasks.json as the post-expansion state boundary (agent-mode insertion is specified “after Task Master expands tasks” when these exist). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Preserve oraclepack run state and reporting outputs: ticket-action-pack.state.json and ticket-action-pack.report.json. These are explicitly named outputs and are the audit/resume boundary (the broader design expectation is “persist run state” + “machine-readable summary report” with stable schema/versioning). 

Oraclepack_Compatibility_Issues

 

oraclepack-llms-full

Add and preserve new headless “agent” artifacts under .oraclepack/ticketify/ with deterministic paths: next.json (Step 09), codex-implement.md (Step 10), codex-verify.md and/or gemini-review.json (Step 11), and PR.md (Step 16). These paths are hard requirements in the ticket’s acceptance criteria. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Preserve the pack’s ingestible shape as a compatibility constraint: action packs must remain “oraclepack-ingestible” (single bash fence, # NN) step headers), and taskify “agent-mode” must keep the “20-step contract intact” by swapping a step slot rather than adding steps. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Treat “project-root execution” as a stability requirement for artifact paths: steps run via bash -lc in the project root and oraclepack does not chdir to out_dir, so all preserved artifacts must be written/read using repo-root-relative deterministic paths (not relying on CWD changes). 

Oraclepack_Compatibility_Issues

Back-compat requirement for overrides/validation: oraclepack’s special injection/validation currently targets only commands beginning with oracle (regex anchored to ^(\\s*)(oracle)\\b), so any stateful Codex/Gemini/Task Master artifacts must not depend on oracle-specific overrides unless dispatcher logic is explicitly extended. 

Oraclepack_Compatibility_Issues

ROI metadata stability affects whether artifacts exist at all: ROI filtering may skip steps without ROI= when thresholds are used, so steps that must produce the preserved artifacts should keep ROI headers (and not become “placeholder/notes” steps) to prevent accidental non-execution. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

Exact schemas are not specified (in the ticket) for .oraclepack/ticketify/_tickets_index.json, .oraclepack/ticketify/tm-complexity.json, .oraclepack/ticketify/next.json, and the exact fields/versions of ticket-action-pack.state.json / .report.json; without these, schema drift ris

Oraclepack_Compatibility_Issues
