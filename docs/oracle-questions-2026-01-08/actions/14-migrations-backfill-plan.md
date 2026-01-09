1. Direct answer (needed backfill / one-time transforms; verification plan)

* One-time transform: patch/regenerate existing `ticket-action-pack.md` instances so placeholder steps are replaced at the specified insertion points (09–13 and 16) with real headless `gemini` + non-interactive `codex exec` commands, while preserving the existing 20-step contract. 
* One-time transform: ensure the updated steps write the new artifacts exactly at the ticketed paths under `.oraclepack/ticketify/` (`next.json`, `codex-implement.md`, `codex-verify.md` and/or `gemini-review.json`, `PR.md`). 
* One-time transform: add command-availability guards (`command -v ...`) plus “skip” behavior for missing `codex`/`gemini` so older environments don’t hard-fail when the new steps land. 
* One-time transform (taskify packs): add a mode switch (suggested `mode=codex` / `mode=gemini`) that swaps the existing “autopilot entrypoint” step slot with “agent implementation,” keeping the 20-step contract intact and placing it after Task Master expansion artifacts exist. 
* Potential one-time transform (only if you implement dispatcher changes): widen command detection beyond `oracle` (currently described as anchored to `^(\\s*)(oracle)\\b`) so steps containing `tm`/`task-master`, `codex`, `gemini` are not excluded from override/validation. Any backfill here is likely “re-run override selection / re-validate packs” rather than rewriting pack files, but the exact per-tool override semantics are unspecified. 
* Verification plan: for the updated `ticket-action-pack.md`, assert the concrete acceptance criteria via file existence checks after running the relevant steps: Step 09 produces `.oraclepack/ticketify/next.json`; Step 10 produces `codex-implement.md`; Step 11 produces `codex-verify.md` and/or `gemini-review.json`; Step 16 produces `PR.md`; and missing `codex`/`gemini` triggers the documented skip behavior. 
* Verification plan: run pack validation + execution in a “no tools installed” environment to confirm the core pack still runs through earlier steps and that new steps degrade gracefully (skip vs block), aligning with the documented “interactive CLI can block” and PATH-missing failure modes. 
* Verification plan (dispatcher change): create a minimal test pack step that contains only `tm`/`codex`/`gemini` prefixes and confirm it is detected and included in the override/validation pipeline (i.e., no longer skipped solely because it lacks an `oracle ...` invocation). 

2. Risks/unknowns

* Unknown: what “overrides” should mean for `tm`/`task-master`, `codex`, `gemini` (which flags, how validation works). This directly affects whether dispatcher changes require any state/config backfill. 
* Risk: headless flags may still be insufficient and could block runs if either CLI remains interactive in edge cases. 
* Risk: environments lacking `codex`/`gemini` on PATH will fail unless the guarded skip behavior is implemented exactly as intended. 
* Risk: ROI filtering can skip steps without `ROI=` when thresholds are used, which can make verification look like “it didn’t run” unless tests pin ROI behavior. 
* Unknown: exact generator/template file locations for `ticket-action-pack.md` and taskify pack generation are not provided, so the transform/backfill automation can’t be precisely scoped yet. 

3. Next smallest concrete experiment (exactly one action)

* Create a single “golden” `ticket-action-pack.md` fixture that implements the Step 09/10/11/16 replacements with `command -v` guards, then run `oraclepack validate` followed by `oraclepack run --no-tui` for steps 09–11 and 16 in two environments (with and without `codex`/`gemini` on PATH) and assert the artifact-path acceptance criteria.

4. If evidence is insufficient, exact missing file/path pattern(s) to attach next

* The generator/template source for `ticket-action-pack.md` (e.g., `**/*ticket-action-pack*.md`, `**/*ticketify*pack*.md`, `**/assets/**`, `**/templates/**`).
* The oraclepack command detection + override injection + validation implementation (e.g., `**/cmd/oraclepack/**`, `**/internal/**`, `**/*dispatch*.*`, `**/*override*.*`, `**/*validate*.*`).
* The persisted override/state/report schema definitions and writers/readers (e.g., `**/*.state.json*` schema code, `**/*report*.go`, `**/*state*.go`).
