Direct answer (rollback plan, artifact preservation, safe-to-re-run semantics)

Treat “rollback” as step-scoped, not transactional: oraclepack executes each step as a literal shell script (bash -lc ...), so the only reliable boundary for failure/rollback is “step N succeeded/failed,” not partial command-level undo within a step. 

Oraclepack_Compatibility_Issues

On any failure, default behavior should be “stop, preserve everything written so far,” because steps may include non-oracle tools (tm, codex, gemini) that oraclepack does not currently wrap/normalize, and those tools can have side effects oraclepack can’t safely undo. 

Oraclepack_Compatibility_Issues

Define “owned artifacts” that oraclepack must always preserve across partial runs: .oraclepack/ticketify/_tickets_index.json, _actions.json, _actions.md, .taskmaster/docs/tickets_prd.md, .oraclepack/ticketify/tm-complexity.json, plus the run-level ticket-action-pack.state.json and ticket-action-pack.report.json mentioned in the tickets. These become the canonical resume inputs/outputs. 

Oraclepack_Compatibility_Issues

Require atomic writes for owned artifacts (write temp → rename) so a killed/partial step cannot leave corrupted JSON/MD that later steps consume (especially the “select → implement” handoff via .oraclepack/ticketify/next.json). 

Oraclepack_Compatibility_Issues

Persist a run state file after every step with: step number, status (pending/running/succeeded/failed/skipped), timestamps, and pointers to logs; load it on resume. This aligns with the project’s stated “state/resume/reporting” capability and avoids guessing what succeeded after a crash. 

oraclepack-llms-full

Preserve logs per step even on failure (stdout/stderr), and record their paths in the report JSON so operators can diagnose why a partial run stopped without rerunning destructive steps. 

oraclepack-llms-full

Safe-to-re-run semantics (default): a step is re-runnable if either (a) its expected outputs are missing, or (b) state says it failed/incomplete. If state says “succeeded” and expected outputs exist, skip the step by default (idempotent resume). For ticketify, expected outputs include .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md / gemini-review.json, and PR.md where applicable.

Safe-to-re-run semantics (non-oracle tools): every codex/gemini/tm step must have command-availability guards and explicit “SKIP” behavior when missing, so reruns don’t hard-fail solely due to environment drift. 

Oraclepack_Compatibility_Issues

Define “rollback” actions as reversible-only for owned artifacts: on rerun of a failed step, move any prior step-owned outputs to a quarantine path (e.g., .oraclepack/ticketify/_failed/<run_id>/step-10/*) rather than deleting them; never attempt to auto-revert repo changes (operators use VCS for that). This is necessary because steps run in the repo root and may touch arbitrary paths beyond out_dir. 

Oraclepack_Compatibility_Issues

For Task Master involvement, preserve its generated outputs as durable inputs (PRD/tasks/complexity) rather than regenerating blindly; the tool’s purpose is workflow automation with generated task artifacts, so losing them on partial runs defeats the pipeline. 

claude-task-master-llms

Risks / unknowns

The tickets name ticket-action-pack.state.json / ticket-action-pack.report.json but do not specify schema, location rules (cwd vs out_dir), or whether they already exist in current oraclepack builds. 

Oraclepack_Compatibility_Issues

“Safe-to-re-run” for steps that modify the working tree (e.g., Codex implementation) is underspecified: without explicit git/patch guards, rerunning may duplicate edits or create conflicts. The tickets only call out interactivity/PATH failures, not repo-mutation policy. 

Oraclepack_Compatibility_Issues

Because oraclepack special handling cu
