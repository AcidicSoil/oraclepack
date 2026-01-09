<filetree>
Project Structure:
└── docs
    └── oracle-questions-2026-01-08
        ├── actions
        │   ├── 01-contracts-interfaces-ticket-surface.md
        │   ├── 02-contracts-interfaces-integration-points.md
        │   ├── 03-invariants-invariant-map.md
        │   ├── 04-invariants-validation-boundaries.md
        │   ├── 05-caching-state-state-artifacts.md
        │   ├── 06-caching-state-cache-keys.md
        │   ├── 07-background-jobs-job-model.md
        │   ├── 08-background-jobs-queue-failure.md
        │   ├── 09-observability-logging-metrics.md
        │   ├── 10-observability-tracing.md
        │   ├── 11-permissions-authz-gaps.md
        │   ├── 13-migrations-schema-migrations.md
        │   ├── 14-migrations-backfill-plan.md
        │   ├── 15-ux-flows-user-journeys.md
        │   ├── 16-ux-flows-edge-cases.md
        │   ├── 17-failure-modes-timeouts-retries.md
        │   ├── 18-failure-modes-rollback-plan.md
        │   ├── 19-feature-flags-flag-plan.md
        │   └── 20-feature-flags-compat-rollout.md
        ├── packs
        │   ├── actions.md
        │   ├── mcp.md
        │   ├── misc.md
        │   ├── other.md
        │   └── prd-tui.md
        ├── _groups.json
        └── manifest.json

</filetree>

<source_code>
docs/oracle-questions-2026-01-08/_groups.json
```
{
  "PRD-TUI": [
    ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
    ".tickets/PRD-TUI/PRD-generator URL routing.md"
  ],
  "actions": [
    ".tickets/actions/Enable Action Packs Dispatch.md",
    ".tickets/actions/Improving Oraclepack Workflow.md",
    ".tickets/actions/Oraclepack Action Pack Integration.md",
    ".tickets/actions/Oraclepack Action Pack Issue.md",
    ".tickets/actions/Oraclepack Action Packs.md",
    ".tickets/actions/Oraclepack Compatibility Issues.md"
  ],
  "mcp": [
    ".tickets/mcp/Expose Oraclepack as MCP.md",
    ".tickets/mcp/MCP Server for Oraclepack.md",
    ".tickets/mcp/gaps-still-not-covered.md",
    ".tickets/mcp/gaps_part2-mcp-builder.md",
    ".tickets/mcp/oraclepack-MCP.md",
    ".tickets/mcp/oraclepack_mcp_server.md"
  ],
  "misc": [
    ".tickets/Oraclepack File Storage.md",
    ".tickets/Oraclepack Schema Approach.md",
    ".tickets/Oraclepack bash fix.md",
    ".tickets/Publish OraclePack MCP.md"
  ],
  "other": [
    ".tickets/other/Oraclepack Pipeline Improvements.md",
    ".tickets/other/Oraclepack Prompt Generator.md",
    ".tickets/other/Oraclepack Workflow Enhancement.md",
    ".tickets/other/Verbose Payload Rendering TUI.md"
  ]
}
```

docs/oracle-questions-2026-01-08/manifest.json
```
{
  "groups": [
    {
      "attached_paths": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "group": "PRD-TUI",
      "original_tickets": [
        ".tickets/PRD-TUI/Oraclepack TUI Integration.md",
        ".tickets/PRD-TUI/PRD-generator URL routing.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-08/prd-tui",
      "pack_path": "docs/oracle-questions-2026-01-08/packs/prd-tui.md",
      "part": 1,
      "slug": "prd-tui"
    },
    {
      "attached_paths": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "group": "actions",
      "original_tickets": [
        ".tickets/actions/Enable Action Packs Dispatch.md",
        ".tickets/actions/Improving Oraclepack Workflow.md",
        ".tickets/actions/Oraclepack Action Pack Integration.md",
        ".tickets/actions/Oraclepack Action Pack Issue.md",
        ".tickets/actions/Oraclepack Action Packs.md",
        ".tickets/actions/Oraclepack Compatibility Issues.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-08/actions",
      "pack_path": "docs/oracle-questions-2026-01-08/packs/actions.md",
      "part": 1,
      "slug": "actions"
    },
    {
      "attached_paths": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "group": "mcp",
      "original_tickets": [
        ".tickets/mcp/Expose Oraclepack as MCP.md",
        ".tickets/mcp/MCP Server for Oraclepack.md",
        ".tickets/mcp/gaps-still-not-covered.md",
        ".tickets/mcp/gaps_part2-mcp-builder.md",
        ".tickets/mcp/oraclepack-MCP.md",
        ".tickets/mcp/oraclepack_mcp_server.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-08/mcp",
      "pack_path": "docs/oracle-questions-2026-01-08/packs/mcp.md",
      "part": 1,
      "slug": "mcp"
    },
    {
      "attached_paths": [
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "group": "misc",
      "original_tickets": [
        ".tickets/Oraclepack File Storage.md",
        ".tickets/Oraclepack Schema Approach.md",
        ".tickets/Oraclepack bash fix.md",
        ".tickets/Publish OraclePack MCP.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-08/misc",
      "pack_path": "docs/oracle-questions-2026-01-08/packs/misc.md",
      "part": 1,
      "slug": "misc"
    },
    {
      "attached_paths": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "group": "other",
      "original_tickets": [
        ".tickets/other/Oraclepack Pipeline Improvements.md",
        ".tickets/other/Oraclepack Prompt Generator.md",
        ".tickets/other/Oraclepack Workflow Enhancement.md",
        ".tickets/other/Verbose Payload Rendering TUI.md"
      ],
      "out_dir": "docs/oracle-questions-2026-01-08/other",
      "pack_path": "docs/oracle-questions-2026-01-08/packs/other.md",
      "part": 1,
      "slug": "other"
    }
  ]
}
```

docs/oracle-questions-2026-01-08/actions/01-contracts-interfaces-ticket-surface.md
```
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
```

docs/oracle-questions-2026-01-08/actions/02-contracts-interfaces-integration-points.md
```
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
```

docs/oracle-questions-2026-01-08/actions/03-invariants-invariant-map.md
```
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
```

docs/oracle-questions-2026-01-08/actions/04-invariants-validation-boundaries.md
```
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
```

docs/oracle-questions-2026-01-08/actions/05-caching-state-state-artifacts.md
```
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
```

docs/oracle-questions-2026-01-08/actions/06-caching-state-cache-keys.md
```
Direct answer (caching opportunities/risks; keys; invalidation; correctness)

Ticket discovery cache: Step 01’s .oraclepack/ticketify/_tickets_index.json is a natural deterministic cache of “what tickets exist + extracted metadata.” Cache key = hash of all ticket file paths+contents under the discovered roots (notably .tickets/), plus discovery algorithm/version, plus run CWD (since steps run in project root and oraclepack does not chdir to out_dir). Invalidate on any ticket add/remove/edit/rename, discovery logic change, or different root/CWD. Correctness risk: stale index leads to downstream actions/PRD being generated against the wrong ticket set; missing .tickets/ is a known failure point. 

Oraclepack_Compatibility_Issues

Action synthesis cache: Step 02’s _actions.json + _actions.md can be cached as a pure function of _tickets_index.json + the prompt/template used to derive actions. Cache key = _tickets_index.json content hash + “actions prompt/template version” + toolchain identity (e.g., oracle CLI flags if oracle is used here). Invalidate if tickets index changes, prompt/template changes, or toolchain flags change. Correctness risk: reusing _actions.json after tickets change causes Step 09 selection and Step 10 implementation to target obsolete work.

Task Master derived artifacts cache: Step 03 .taskmaster/docs/tickets_prd.md and Steps 05–07 outputs (including .oraclepack/ticketify/tm-complexity.json) are expensive and can be cached with key = _actions.json hash + Task Master version/config + PRD generation prompt/template. Invalidate on _actions.json change, Task Master upgrade/config change, or PRD template change. Correctness risk: stale complexity/expanded tasks drives “agent-mode” (post-expansion) to implement the wrong tasks.

“Next target” selection cache: Step 09’s .oraclepack/ticketify/next.json should be treated as a cacheable decision record. Cache key = _actions.json hash + repo state fingerprint (recommended: git rev-parse HEAD + git diff hash) + selection heuristic version. Invalidate on actions change or repo state change. Correctness risk: using an old next.json after code changes can produce irrelevant patches or merge conflicts in Step 10.

Codex/Gemini implementation/review outputs cache: Step 10 .oraclepack/ticketify/codex-implement.md and Step 11 .oraclepack/ticketify/codex-verify.md / .oraclepack/ticketify/gemini-review.json should be cached only as immutable run artifacts, not as reusable inputs, unless the repo-state key matches exactly. Cache key = next.json hash + repo-state fingerprint + codex/gemini invocation flags + tool versions. Invalidate on any mismatch. Correctness risk: reusing “implementation” content across different repo states can apply incorrect changes; interactivity/availability are known failure modes, so caches must not mask missing-tool failures.

Oracle command output cache (per-step --write-output): each oracle invocation’s output markdown is cacheable if (and only if) inputs are fully captured. Cache key = (full prompt string, including embedded strategist question) + ordered list of attached files with content hashes + oracle flags (the injected/overridden flags matter) + oracle CLI version/engine/model selection. Invalidate on any of those changing. Correctness risk: if override injection changes the effective flags but the cache key only uses the pre-injection command string, you can serve the wrong output.
```

docs/oracle-questions-2026-01-08/actions/07-background-jobs-job-model.md
```
Direct answer (1–10 bullets, evidence-cited)

There is no explicit “oraclepack-managed” background job system implied; packs execute step bodies as shell (bash -lc ...) in-order, so any async/long-running behavior is delegated to the invoked tools inside a step (e.g., oracle, tm, codex, gemini). 

Oraclepack_Compatibility_Issues

Oracle CLI “background mode” (Responses-style async) is an implied long-running operation: requests may return queued/in_progress and Oracle must poll until completed, with retry/backoff and a max-wait timeout. Interface boundary is effectively pollBackgroundResponse(params) with client, responseId, initialResponse, log, wait, now, maxWaitMs (plus retrieveBackgroundResponseWithRetry). 

oraclepack-llms-full

Task Master “autopilot” is an implied long-running (and potentially interactive) operation executed as a pack step. The step is responsible for branch safety + preflight checks, writing a resumable state file (tm-autopilot.state.json), then launching ${TM_CMD} autopilot. Inputs are MODE/TM_CMD/OUT_DIR/PACK_DATE/TAG; output contract includes the state file path. 

oraclepack-llms-full

The tickets imply “agent-mode” replacement of autopilot with headless gemini + non-interactive codex exec steps (select → implement → verify → draft PR). These are long-running operations that must be non-interactive to avoid blocking oraclepack run. The interface is file-based artifacts under .oraclepack/ticketify/ (e.g., next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md) with command-availability guards (command -v ...) and “skip” semantics when tools are missing.

If you want oraclepack-level responsibility for async, the implied seam is the “dispatcher/override/validation pipeline” extension: broaden command detection beyond oracle so tm/codex/gemini steps can participate in consistent validation/guardrails (but the ticket does not define the per-tool override/validation semantics). 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

Non-oracle tool steps can block indefinitely if interactive; the ticket explicitly flags interactivity as a failure mode and requires headless/non-interactive variants. 

Oraclepack_Compatibility_Issues

The override/validation behavior for tm/codex/gemini is underspecified (which flags apply, what “validation” means per tool), so implementing oraclepack-managed async/dispatch risks being partial or incorrect. 

Oraclepack_Compatibility_Issues

It’s unclear whether the repo contains any first-class job/queue implementations beyond “tool-driven” long-running steps; the existing coverage checklist explicitly says “background jobs: Missing” and requests specific directories if they exist. 

oraclepack-llms-full

Autopilot resumability is only implied (“state file created… should resume if supported”); whether tm autopilot actually supports resume-from-state is unknown. 

oraclepack-llms-full

Next smallest concrete experiment (exactly one action)

Add a “resumable long-running step wrapper” to the autopilot step (write/update a state JSON before/after launch; wrap the tm autopilot invocation with a hard timeout and clear exit codes), then run oraclepack validate + oraclepack run --no-tui on that pack to confirm (a) no indefinite hangs, and (b) reruns observe the state file and behave deterministically. 

oraclepack-llms-full

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

**/jobs/** (explicitly requested by the current coverage checklist) 

oraclepack-llms-full

**/queue/**, **/worker/**, **/cron/**, **/*scheduler*/**, **/*background*/**

Any runtime/service configs that define scheduled work: **/*.yaml / **/*.yml / **/*.toml containing cron, schedule, worker, queue

If Task Master autopilot integration
```

docs/oracle-questions-2026-01-08/actions/08-background-jobs-queue-failure.md
```
Direct answer (1–10 bullets, evidence-cited)

Today, “background” work in these Action Packs is not a separate job system: oraclepack executes each step as a shell script (bash -lc ...), so failures are synchronous step failures (the step exits non-zero) unless the step body explicitly catches/handles them. 

Oraclepack_Compatibility_Issues

For non-oracle tooling (tm/task-master, codex, gemini), oraclepack does not apply oracle-specific override/validation; those commands run directly, so any retries/idempotency must be implemented either inside the tool itself or in the step script wrapper. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

The tickets explicitly call out “missing binary on PATH fails the step” and “interactive CLI blocks”, and recommend command-availability guards (command -v ...) plus “skip” behavior to avoid hard failures when tools are missing; treat this as the baseline failure-handling pattern for codex/gemini. 

Oraclepack_Compatibility_Issues

Retries: where API/background polling exists (Oracle API “background response” flow), the client logic already uses automatic retry for retryable transport errors with exponential backoff and a max-wait timeout; mirror this pattern in step-level wrappers for transient failures (network/API flaps) when invoking headless gemini / codex exec. 

oraclepack-llms-full

Idempotency: the automation steps are required to write deterministic artifacts at fixed paths (e.g., .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md, gemini-review.json, PR.md); to make reruns safe, write outputs atomically (tmp file + rename) and overwrite in place, since downstream steps depend on those canonical paths existing. 

Oraclepack_Compatibility_Issues

Poison messages (step-runner analogue): define a “poison step” as one that fails repeatedly with the same non-transient reason (e.g., missing CLI, interactive prompt, invalid config). The ticket’s prescribed approach is to convert some of these into “skip” outcomes (guard + skip) rather than repeated hard-fails; for non-skippable steps, persist a failure record in the pack’s state/report outputs so it’s visible and resumable. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Observability hooks should be file-first because the workflow already hinges on emitted artifacts: capture (a) tool-availability checks and skip reasons, (b) step status (started/completed/failed/skipped), (c) timings and exit codes, and (d) pointers to generated artifacts like .oraclepack/ticketify/* and the pack state/report JSON mentioned in the ticket. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

For long-running “background-like” operations (API background responses), log status transitions (“queued”/“in_progress”/reconnected) and retry delays; this is already present in the background polling logic and is the concrete model to replicate for gemini/codex headless steps (emit to stdout and/or a JSONL events file). 

oraclepack-llms-full

Risks/unknowns (bullets)

The tickets do not specify whether oraclepack should implement runner-level retries (re-run a failing step automatically) vs step-script retries only; adding runner-level retries could change failure semantics. 

Oraclepack_Compatibility_Issues

No explicit policy is provided for which failures are “skippable” vs “fatal” across Steps 09–13 and 16; incorrect classification could silently skip important work. 

Oraclepack_Compatibility_Issues

“Poison message” handling is underspecified for this architecture (step-based execution, not
```

docs/oracle-questions-2026-01-08/actions/09-observability-logging-metrics.md
```
Direct answer (logging/metrics + minimal instrumentation plan)

Pack-generation logs must emit: (a) input discovery summary (ticket roots searched, tickets found), (b) deterministic output paths written (e.g., .oraclepack/ticketify/_tickets_index.json, _actions.json, _actions.md, .taskmaster/docs/tickets_prd.md, .oraclepack/ticketify/tm-complexity.json), and (c) “workdir = project root” (since execution does not switch to out_dir). 

Oraclepack_Compatibility_Issues

Pack-parse/validate logs must include: detected “single bash fence” + parsed step headers (# NN)), step count/order, and any rejection reason (because ingestion is strict and mis-detection causes downstream execution surprises). 

Oraclepack_Compatibility_Issues

Step-execution logs must be per-step, persisted to disk, and include: step number/title, full script body, bash -lc invocation, cwd, start/end timestamps, exit code, and captured stdout+stderr (streamed + saved). This is required because oraclepack executes literal step bodies as shell. 

Oraclepack_Compatibility_Issues

“Skip reason” logging must be first-class: (a) ROI-filter skips when steps lack ROI= and threshold > 0, (b) tool-availability skips via command -v ... guards, and (c) mode-gated steps (e.g., agent-mode/autopilot slots). Without explicit skip reasons, headless runs look like silent no-ops.

Command-detection / override-injection / validation logs must explicitly show: (a) which lines matched the current oracle-anchored regex, (b) what flags were injected, and (c) the result of override validation (oracle --dry-run summary)—because today only oracle-prefixed commands participate, while tm/codex/gemini do not.

External-tool execution logs must record “missing binary” vs “interactive block risk” for tm/codex/gemini, since the tickets call out both as primary failure modes (missing PATH hard-fails unless guarded; interactive CLIs can hang headless runs).

Artifact-verification logging must run after relevant steps and record existence/size/hash for required outputs (at minimum: .oraclepack/ticketify/next.json, codex-implement.md, codex-verify.md and/or gemini-review.json, PR.md). These are explicit acceptance criteria; when absent, the logs should say “expected but missing” and point to the producing step.

Minimal metrics should be derivable from the same run report: counters for steps_total/ran/skipped/failed, plus durations (pack_parse_ms, prelude_ms, step_duration_ms by step/tool), and counts of oracle_detected/oracle_validated/non_oracle_steps. The repo already targets state+report artifacts for traceability; extend them to carry these fields rather than inventing a separate telemetry stack.

Minimal instrumentation plan (implementation shape): write (1) per-step log files and (2) a single machine-readable run summary (JSON) that includes per-step Result{exitCode, startedAt, endedAt, logPath} and environment metadata (shell, cwd), enabling headless debugging and CI parsing.

Risks/unknowns

Current ticket bundle provides desired behaviors and failure modes, but not the existing logger/report schema actually implemented (so some recommendations may already exist partially). 

Oraclepack_Compatibility_Issues

Exact override behavior for non-oracle tools (tm/codex/gemini) is explicitly unspecified; metrics/logging can expose gaps, but cannot resolve semantics without a spec.

Step 11 default (“Codex verify vs Gemini review vs both”) remains undecided; instrumentation should capture which path ran and why, but product behavior still needs a default decision.
```

docs/oracle-questions-2026-01-08/actions/10-observability-tracing.md
```
Direct answer (1–10 bullets, evidence-cited)

Define a single “trace context” for every pack run: run_id (unique per run + stable across resume), pack_path, pack_hash, and started_at; persist it in the existing state/report artifacts (e.g., ticket-action-pack.state.json, ticket-action-pack.report.json) so every downstream artifact can be joined back to one run. 

Oraclepack_Compatibility_Issues

Use pack_hash + step_number as the canonical step key, since the pack model already has Pack{path, hash} and Step{number, title} and the run model already tracks per-step results and logPath; extend those records with run_id and attempt (increment on retries). 

oraclepack-llms-full

Propagate correlation across all tools by setting environment variables for every step execution (works because steps execute as bash -lc ...):
ORACLEPACK_RUN_ID, ORACLEPACK_PACK_HASH, ORACLEPACK_PACK_PATH, ORACLEPACK_STEP_NO, ORACLEPACK_STEP_TITLE, ORACLEPACK_STEP_ATTEMPT, ORACLEPACK_LOG_PATH, ORACLEPACK_OUT_DIR. This avoids relying on tool-specific wrappers. 

Oraclepack_Compatibility_Issues

For oracle invocations, additionally inject the trace context into the prompt text (so the LLM output is self-identifying) and/or into the output file preamble; this is especially important because oraclepack’s current special handling/override pipeline only targets commands beginning with oracle. 

Oraclepack_Compatibility_Issues

For non-oracle tools (task-master/tm, codex, gemini), treat the step log (logPath) as the ground-truth join point, because they currently run “directly” and do not inherit oraclepack’s oracle-only injection/validation behavior.

Require every step that produces a machine-readable artifact to embed {run_id, pack_hash, step_no} inside the artifact (JSON metadata block for .json, a short header for .md). This is aligned with the tickets’ emphasis on deterministic, path-specific artifacts under .oraclepack/ticketify/* (e.g., next.json, codex-implement.md, gemini-review.json, PR.md). 

Oraclepack_Compatibility_Issues

When/if you implement the dispatcher changes requested in the tickets (extend detection/overrides beyond oracle), define a per-tool adapter that maps the same trace context into each tool’s native flags (where available) while keeping env-var propagation as the baseline. 

Oraclepack_Compatibility_Issues

Minimum required IDs and how they join:
run_id (primary join key across everything), pack_hash (stable pack identity), step_no (01..20), attempt (retry/resume), invocation_index (nth tool call within a step, if needed), plus logPath (filesystem pointer to raw evidence). 

oraclepack-llms-full

Risks/unknowns (bullets)

The tickets don’t specify whether codex/gemini expose native run/session IDs or non-interactive flags sufficient to avoid blocking; without that, correlation may rely entirely on logs + embedded metadata.

Current oraclepack behavior applies special handling only to oracle-prefixed commands; if you expect uniform injection/validation/tracing for tm/codex/gemini, you need the dispatcher work (and a spec for per-tool override behavior).

The provided materials describe desired state/report models but not the concrete on-disk schema used by your current runner; adding fields requires a versioning/back-compat plan.

Next smallest concrete experiment (exactly one action)

Implement runner-level env-var injection + add run_id to the state/report JSON (persisted once per run, reused on resume), then run oraclepack run --no-tui on ticket-action-pack.md and confirm each step log contains the printed ORACLEPACK_RUN_ID and each produced .oraclepack/ticketify/* artifact embeds {run_id, pack_hash, step_no}.

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

cmd/oraclepack/** (entryp**
```

docs/oracle-questions-2026-01-08/actions/11-permissions-authz-gaps.md
```
Direct answer (1–10 bullets, evidence-cited)

Command execution boundary: oraclepack runs each step by executing the step body as shell via bash -lc ..., i.e., whatever commands appear will run with the oraclepack process’s OS privileges. Safe default: require an explicit “exec enable” gate (flag/env) plus per-step confirmation unless --yes; otherwise operate in validate/list-only mode. 

Oraclepack_Compatibility_Issues

Tool routing/authz boundary: oraclepack’s “special handling” (override injection/validation) only applies to commands that begin with oracle; tm/task-master, codex, and gemini run directly and are not protected by oracle-specific checks. Safe default: treat any non-oracle command as “raw exec” and block-by-default unless an allowlist (by prefix) is configured.

Interactivity boundary: tickets call out that interactive CLIs can block waiting for input. Safe default: force non-interactive execution (environment like CI=1, plus only “headless/non-interactive” subcommands) and add step timeouts so a blocked CLI cannot hang the run indefinitely.

Binary availability boundary: missing binaries on PATH cause step failures; tickets explicitly recommend command -v ... guards with “skip” behavior. Safe default: (a) preflight required tools per pack (oracle/tm/codex/gemini), (b) default to “skip with a clear message” for optional tools, and (c) provide a strict mode that turns skips into hard failures.

File read boundary: tickets list common failure points like missing .tickets/ and emphasize steps run in the project root (not an isolated out_dir). Safe default: restrict file reads to the workspace root + explicitly allowed subtrees (e.g., .tickets/, tickets/) and deny absolute-path reads or .. traversal outside allowed roots.

File write boundary: expected artifacts are written into deterministic project paths like .oraclepack/ticketify/* and .taskmaster/docs/* (and per-step outputs). Safe default: restrict writes to a small allowlist of directories (e.g., .oraclepack/**, .taskmaster/**, and the declared pack out_dir), deny writes outside the repo, and reject absolute paths / parent traversal.

Network/egress boundary (implied by tool choices): the proposed automation inserts headless gemini and codex exec steps and notes missing provider/API key configuration as a failure mode. Safe default: make “network-enabled execution” an explicit mode (opt-in), and otherwise assume offline (or at least warn and require confirmation before running steps that invoke these tools).

Secrets/logging boundary (implied by provider keys + artifact writing): the workflow depends on external provider configuration/API keys and produces multiple on-disk reports/artifacts. Safe default: never dump environment variables in logs, redact likely secret patterns in streamed output before writing report JSON/MD, and prohibit packs from writing secret material into .oraclepack/ticketify/* outputs.

Risks/unknowns (bullets)

The desired override/validation semantics for tm/task-master, codex, and gemini are explicitly unspecified (which flags apply, how validation works), so any “first-class tool authz” policy could be incomplete or wrong.

It’s not specified where the authoritative user-facing permission model should live (repo docs vs TUI help), which increases the chance the real boundary behavior is misunderstood by users. 

Oraclepack_Compatibility_Issues

“Agent-mode” selection mechanism (CLI flag vs TUI option vs config) is not defined, which affects how you safely expose network/exec enablement without surprising defaults. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Implement a “preflight policy report” in oraclepack (or a companion script) that parses a pack and prints:
```

docs/oracle-questions-2026-01-08/actions/13-migrations-schema-migrations.md
```
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
```

docs/oracle-questions-2026-01-08/actions/14-migrations-backfill-plan.md
```
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
```

docs/oracle-questions-2026-01-08/actions/15-ux-flows-user-journeys.md
```
Direct answer (UX/TUI workflows, journey states, expected outputs)

Core “run pack” journey: user selects an Action Pack, oraclepack parses it, then executes each step as literal shell via bash -lc ... (in repo root; no automatic cd to out_dir). Journey states: Selected → Parsed/Validated → Ready → Running(step N) → Step {Succeeded|Failed|Skipped|Blocked} → Completed/Resumable. Expected outputs: per-step logs plus persistent run artifacts like ticket-action-pack.state.json and ticket-action-pack.report.json (names explicitly referenced), alongside whatever files each step writes. 

Oraclepack_Compatibility_Issues

Overrides/validation journey is “oracle-only” today: in TUI, an “Overrides Wizard”/override-validation flow should only attach to steps containing oracle ... lines, because override injection/detection is anchored to oracle-prefixed commands (regex anchored to ^(\\s*)(oracle)\\b), and override validation runs oracle --dry-run summary on detected oracle invocations while skipping steps with no oracle invocations. Journey states: Overrides entry → Step classification {Oracle-managed|Shell-direct} → Overrides selected (Oracle-managed only) → Dry-run validation (Oracle-managed only) → Apply during execution. 

Oraclepack_Compatibility_Issues

ROI filtering journey: users can filter which steps appear/run by ROI; a known UX gotcha is that steps without ROI= may be skipped when an ROI threshold > 0 is used. Journey states: Filter configured → Steps partitioned {Included|Excluded (reason: ROI missing/too low)} → Run included set. Expected output: UI should surface counts and exclusion reasons to avoid “missing steps” confusion. 

Oraclepack_Compatibility_Issues

Ticketify “analysis pipeline” journey (Steps 01–07): user runs early steps to discover tickets and generate action/PRD scaffolding. Expected outputs include .oraclepack/ticketify/_tickets_index.json (Step 01), _actions.json + _actions.md (Step 02), .taskmaster/docs/tickets_prd.md (Step 03), and .oraclepack/ticketify/tm-complexity.json after Task Master parse/complexity/expand steps. Journey states: Inputs present (.tickets/) → Index/actions generated → PRD generated → TM expanded tasks available. 

Oraclepack_Compatibility_Issues

Ticketify “agent loop” journey (replacing placeholders 09–13 + 16): implied end-to-end headless workflow is Select → Implement → Verify → Draft PR. Expected outputs: Step 09 writes .oraclepack/ticketify/next.json; Step 10 writes .oraclepack/ticketify/codex-implement.md; Step 11 writes .oraclepack/ticketify/codex-verify.md and/or .oraclepack/ticketify/gemini-review.json; Step 16 writes .oraclepack/ticketify/PR.md. Journey states: Candidate chosen → Implementation attempted → Verification completed → PR draft produced.

Non-oracle tool execution journey (tm/codex/gemini): these tools run directly (not routed through oracle overrides). The UX implication is preflight detection and “blocked vs skipped” semantics: if the binary is missing on PATH, the step fails unless the step implements guard+skip logic; if the tool is interactive, it can block waiting for input. Journey states: Preflight {Available|Missing|Interactive-risk} → Execute or Skip (if guarded) → Surface outcome clearly.

“Agent-mode” selection for taskify-generated packs: tickets imply a mode switch (e.g., mode=codex / mode=gemini) that swaps an existing “autopilot entrypoint” step slot for agent implementation while keeping the 20-step contract intact. UX journey states: Mode selected → Step slot behavior changes (autopilot vs agent loop) → Same step-count/structure preserved. 

Oraclepack_Compatibility_Issues

Risks/unknowns

Exact TUI affordances are underspecified: where to present “oracle-managed vs shell-direct” labeling (step list badge, override wizard copy, help text) is not defined in the tickets. 

Oraclepack_Compatibility_Issues

Dispatcher semantics for non-oracle commands are unspecified: what “overrides” would mean for tm/codex/gemini (flags, validation method, dry-run equivalents) is explicitly unknown. 

Oraclepack_Compatibility_Issues

Step 11 default verification path is undecided (Codex vs Gemini vs both), which impacts the expected UX branching and success criteria. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Add a TUI-visible per-step classification (“Oracle-managed” iff the step contains an oracle invocation; otherwise “Shell-direct”), then run a fixture pack containing one oracle ... step and one tm/codex/gemini step t
```

docs/oracle-questions-2026-01-08/actions/16-ux-flows-edge-cases.md
```
Direct answer (edge cases + minimal UX behavior)

Cancel during a step (Ctrl+C / interrupt): treat as “interrupted”, persist partial logs, and persist run state so the user can resume without guessing what ran. This is required because each step is executed as a shell script (bash -lc ...), so interrupts will happen at the process level and must be reflected in state/report, not just in the UI. 

Oraclepack_Compatibility_Issues

Cancel while a non-oracle CLI is running (codex/gemini/tm): surface a specific warning that these tools are executed “as-is” (not routed through oraclepack’s oracle-only dispatcher), and that interactive tools may block indefinitely; cancellation must not leave the TUI thinking the step “completed”. 

Oraclepack_Compatibility_Issues

Resume after cancel: the resume UX must clearly indicate what will re-run (only failed/interrupted steps, not already-successful ones), and it must work even when some steps are “non-oracle” and therefore not part of oraclepack override/validation. 

Oraclepack_Compatibility_Issues

Partial runs caused by missing binaries (codex/gemini not installed): steps that are intended to be optional automation should “skip” (not hard-fail) when command -v checks fail, and the UI should label them as “skipped: tool not available” to keep the run resumable and comprehensible. 

Oraclepack_Compatibility_Issues

Partial runs caused by interactivity (tool prompts for input): minimal UX is (a) warn before running a step that is likely to block, and (b) provide a consistent cancel path that records the step as “blocked/interrupted” rather than “running forever”. 

Oraclepack_Compatibility_Issues

Resume when ROI filtering changes what’s “visible”: steps can be silently excluded when ROI thresholding is used, especially if headers lack ROI=. Minimal UX must explicitly show “skipped by filter” (and why) so users don’t confuse filtering with completion. 

Oraclepack_Compatibility_Issues

Resume when overrides wizard was used: because override injection/validation is currently oracle-only, the UI must make it impossible to assume codex/gemini/tm steps were validated/overridden, and must preserve that distinction on resume (otherwise users will infer safety guarantees that do not exist). 

Oraclepack_Compatibility_Issues

Placeholder/“note” steps vs executable steps: steps that only echo guidance should be visually distinguishable from steps that produce required artifacts; otherwise “run-all” followed by “resume” becomes confusing (many “successful” steps that did nothing). This matters because the current plan explicitly calls out placeholder steps (08–20) and replacing some with real automation. 

Oraclepack_Compatibility_Issues

Minimal required status taxonomy for UX: success / failed / skipped (with reason: ROI filter, tool-missing, user-skip) / interrupted (user cancel, crash). This is the minimum needed to support “cancel → resume” and “partial runs” without ambiguity given the documented failure modes.

Risks/unknowns

The tickets describe execution semantics and failure modes, but do not specify current signal handling details (does SIGINT propagate to the child process group; does the runner always flush logs/state atomically). 

Oraclepack_Compatibility_Issues

The desired “resume contract” is not fully specified (resume from last success vs resume selected step; whether prelude is re-run; how to treat interrupted prelude). 

Oraclepack_Compatibility_Issues

Override/validation behavior for non-oracle tools is explicitly an open design gap; UX must avoid implying guarantees until dispatcher changes exist.

The “skip behavior” for missing codex/gemini is referenced as required, but the exact skip contract (exit code, status file semantics, report fields) is not provided. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Add a single integration test fixture pack with a long-running step (e.g., sleep 30), run it under oraclepack run in a way the test can send SIGINT mid-step, assert the state/report mark the step as “interrupted”, then run oraclepack run --resume and assert it continues from the interrupted step (and that ROI-filtered steps are labeled “skipped by filter”, not “pending”).

I
```

docs/oracle-questions-2026-01-08/actions/17-failure-modes-timeouts-retries.md
```
Direct answer (1–10 bullets, evidence-cited)

Oraclepack executes each pack step as a literal shell script (bash -lc ...), so timeout/retry behavior must be implemented either (a) inside the called tool (oracle, task-master, codex, gemini) or (b) explicitly wrapped in the step body; oraclepack does not “route everything through oracle,” and only oracle-prefixed commands receive oraclepack’s injection/validation behavior. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

For oracle ... steps: do not add an outer “rerun the whole command” retry loop by default (billing/cost risk); instead rely on Oracle’s internal handling: background responses are polled until completion or maxWaitMs, with retry-on-transport during background retrieval and a hard client-timeout when the overall wait budget is exceeded. 

oraclepack-llms-full

Set wall-clock timeouts to respect Oracle’s existing defaults: browser runs show a default timeoutMs of 1,200,000ms (20m) and an inputTimeoutMs on the order of tens of seconds; if you add a shell timeout, set it ≥ 25m for oracle browser/background steps so you don’t preempt Oracle’s own timeout. 

oraclepack-llms-full

For non-oracle CLIs (tm/task-master, codex, gemini): treat “missing binary on PATH” as a first-class, non-retryable Dependency failure; the ticket explicitly calls for command -v ... guards plus “skip” behavior to avoid hard failures when these tools are optional in placeholder steps. Operator action: install tool / fix PATH, then re-run the step. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Treat “interactive CLI blocks waiting for input” as a non-retryable Execution-Mode failure; enforce headless/non-interactive flags and add a shell-level timeout for these steps (e.g., timeout 15m …) so runs cannot hang indefinitely. Operator action: re-run with correct headless/non-interactive configuration. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Classify failures into: (A) Dependency (binary missing), (B) Configuration/Input (missing .tickets/, missing provider keys/config), (C) Interactive/Blocked, (D) Transient Transport (network/connection/timeouts), (E) Remote Throttle/5xx (retryable with backoff), (F) Remote Permanent (auth/validation/permission; non-retryable), (G) Postcondition/Artifact missing, (H) Skipped/Filtered. The ticket calls out .tickets/ missing and provider key/config gaps as common failure points. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

 

claude-task-master-llms

Retry policy: only for (D) and (E). Use bounded exponential backoff with a hard cap (and never exceed the step’s wall-clock budget); do not retry (A)(B)(C)(F). This matches the “retry transport, but stop on non-retryable” shape shown in Oracle’s background retrieval loop. Operator action: if retries exhaust, inspect logs/state, fix network/provider, rerun the step once. 

oraclepack-llms-full

Operator actions for (G) Postcondition failures: treat “command exited 0 but expected artifact path missing” as actionable because the ticket has explicit acceptance criteria for .oraclepack/ticketify/* outputs; action is to open the step output/log, confirm working directory assumptions (oraclepack runs in project root), then rerun step after fixing paths/permissions. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Operator actions for (H) Skipped/Filtered: if ROI filtering is enabled, steps without ROI= may be skipped; treat that as “not executed,” not a failure. Operator action: re-run without the ROI threshold/filter or ensure steps include ROI metadata. 

Oraclepack_Compatibility_Issues

Risks/unknowns (bullets)

The tickets do not specify numeric timeout budgets or retry counts per tool (task-master vs codex vs gemini), only that interactivity can block and missing binaries should be handled with guards/skip behavior. 

Oraclepack_Compatibility_Issues
```

docs/oracle-questions-2026-01-08/actions/18-failure-modes-rollback-plan.md
```
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
```

docs/oracle-questions-2026-01-08/actions/19-feature-flags-flag-plan.md
```
Direct answer (1–10 bullets, evidence-cited)

Treat “feature flags” as compat-preserving runtime gates for the three ticketed deltas: (a) dispatcher/override/validation expansion beyond oracle-prefixed commands, (b) Action Pack placeholder replacement with headless gemini + non-interactive codex exec, and (c) taskify “agent-mode” generation that swaps the same step slot while keeping the contract unchanged.

Define scopes explicitly and separately (so you can roll out safely):

Pack generation scope (taskify/template toggles): taskify.agent_mode / actionpack.enable_agent_steps

Execution/dispatcher scope (step parsing + injection): dispatcher.v2_enabled, overrides.targets

Validation scope (what gets “mode-2 validated”): validation.targets, validation.strict

Defaults must preserve current behavior for existing users: dispatcher.v2_enabled=false, overrides.targets=["oracle"], validation.targets=["oracle"], and any new Codex/Gemini automation in packs must be opt-in at generation time. This aligns with the explicit constraint “must preserve existing oracle command behavior” and the current oracle-only injection/validation behavior.

Rollout the dispatcher expansion in three stages behind flags (reduces blast radius given unclear non-oracle semantics):

“detect-only” (collect telemetry on tm/task-master/codex/gemini occurrences; no injection/validation changes),

“validate-only” (include non-oracle steps in the override/validation pipeline as first-class targets but keep behavior non-blocking unless validation.strict=true),

“enforce” (apply overrides to configured tool prefixes, with tool-specific validators). This matches the ticket’s ask to stop silently excluding non-oracle steps.

For Action Pack headless automation (T3), gate pack content changes with a generation-time flag (e.g., actionpack.enable_agent_steps=true) and require command-availability guards with “skip” semantics (not hard fail) when codex/gemini aren’t on PATH.

For taskify “agent-mode” (T4), add a generation-time mode switch (ticket-suggested mode=codex / mode=gemini) and implement it as a swap of the existing autopilot step slot, keeping the “20-step contract intact” and placing the insertion after Task Master expansion artifacts exist. 

Oraclepack_Compatibility_Issues

Telemetry: always emit the resolved feature-flag set and derived “effective targets” into the machine-readable run report and persisted run state, including counts of detected tool prefixes per step, which steps were targeted, which were skipped (and why), and whether validation ran. This is consistent with the existing direction to write deterministic state.json and a machine-readable summary report (run.json) with stable schema/versioning.

TUI/UX: expose these flags as a single “Experimental / Beta” panel that expands into (1) dispatcher mode, (2) override targets, and (3) validation strictness; keep them off by default and require explicit opt-in. This avoids the current confusion about oraclepack “running everything through oracle” and makes the scope visible.

Guardrails: because steps execute via bash -lc and interactive CLIs can block, require “headless/non-interactive” sub-flags for any Codex/Gemini execution paths and default to “skip with recorded reason” when headless requirements aren’t met.

Risks/unknowns (bullets)

The tickets explicitly do not specify the exact desired override semantics for tm/task-master, codex, and gemini (which flags apply, how validation should work), so an “enforce overrides for non-oracle tools” flag risks doing the wrong thing without a tool-specific contract. 

Oraclepack_Compatibility_Issues

Interactive behavior and missing binaries remain a runtime risk for any headless automation steps unless every path is guarded and defaults to skip.

The selection mechanism for taskify agent-mode (CLI flag vs TUI vs config) is not specified in the ticket summary; inconsistent selection surfaces can fragment behavior and telemetry. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Implement --dispatcher=v2-detect (or equivalent) that onl
```

docs/oracle-questions-2026-01-08/actions/20-feature-flags-compat-rollout.md
```
Direct answer (minimal compat-safe rollout plan + guardrails, fallback, monitoring gates)

Phase 0 (baseline invariants): keep current execution semantics unchanged: oraclepack executes each step via bash -lc ... in the project root, and only applies special handling (flag injection + override validation) to commands that begin with oracle. This is the compatibility baseline you must not break. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Phase 1 (shadow / observe-only): add “dispatcher v2 shadow” that detects steps containing tm/task-master, codex, gemini for reporting/UX warnings, but does not modify execution, inject flags, or expand validation beyond the current oracle-only behavior. This directly addresses the “silently excluded” concern while staying compat-safe. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Phase 2 (opt-in, guarded automation in packs): update ticket-action-pack.md placeholder steps by inserting headless Gemini + non-interactive codex exec only behind an explicit pack mode / flag, and require command-availability guards (command -v ...) with documented “skip” behavior to prevent hard failures when binaries are missing or interactive. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Guardrail: never allow “interactive CLI blocks the run” by default; any Codex/Gemini invocation used in packs must be explicitly non-interactive/headless, and if that cannot be guaranteed in a given environment, the step must skip (and write a skip note) rather than hang. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Guardrail: preserve step-count/schema contracts; for taskify-generated packs, implement mode=codex / mode=gemini by swapping the existing “autopilot entrypoint” step slot (not adding steps), keeping the “20-step contract intact,” and inserting only after Task Master expansion artifacts exist (tasks.json, tm-complexity.json). 

Oraclepack_Compatibility_Issues

Fallback behavior (runtime): if non-oracle overrides/validation are not implemented or misconfigured, fall back to today’s behavior (oracle-only injection/validation) and treat non-oracle tools as literal shell commands; do not attempt partial wrapping that changes semantics unpredictably. 

Oraclepack_Compatibility_Issues

Fallback behavior (pack outputs): for the new automated ticket-action flow, require the artifact-based fallbacks: if gemini is missing, Step 09/16 skip and do not create .oraclepack/ticketify/next.json / PR.md; if codex is missing, Step 10 and Codex-based Step 11 skip and do not create .oraclepack/ticketify/codex-implement.md / codex-verify.md (but the run remains usable for the non-agent steps 01–07). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Monitoring gates (acceptance-as-gates): treat the ticket’s “artifact existence” as rollout gates in CI/canary: Step 09 must produce .oraclepack/ticketify/next.json; Step 10 must produce codex-implement.md; Step 11 must produce at least one of codex-verify.md or gemini-review.json; Step 16 must produce PR.md—or, if tools are missing, must produce an explicit “skipped due to missing binary” record instead of failing. 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Monitoring gates (operational): explicitly track and alert on the known failure modes: (a) step failures due to missing binaries on PATH, (b) runs that stall due to interactivity, and (c) ROI-filter skipping of steps without ROI= when thresholds are used (rollout guard: ensure any newly “required” steps include ROI= or document running with ROI threshold = 0 during canary). 

Oraclepack_Compatibility_Issues

 

Oraclepack_Compatibility_Issues

Risks/unknowns

The ticket set does not specify which overrides/flags should apply to each non-oracle tool (tm, codex, gemini) or what “validation” means for them; implementing anything beyond shadow-detection without this risks incorrect behavior. 

Oraclepack_Compatibility_Issues
```

docs/oracle-questions-2026-01-08/packs/actions.md
```
# Oracle Pack — oraclepack (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-08/actions
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md
- ticket_max_files: 6
- group_name: actions
- group_slug: actions
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/actions/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08/actions"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/actions/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/actions/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: actions)

Reference: actions
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/actions/03-invariants-invariant-map.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/actions/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: actions)

Reference: actions
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/actions/Enable Action Packs Dispatch.md,.tickets/actions/Improving Oraclepack Workflow.md,.tickets/actions/Oraclepack Action Pack Integration.md,.tickets/actions/Oraclepack Action Pack Issue.md,.tickets/actions/Oraclepack Action Packs.md,.tickets/actions/Oraclepack Compatibility Issues.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'actions'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/actions/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: actions)

Reference: actions
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=actions

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
[TRUNCATED]
```

docs/oracle-questions-2026-01-08/packs/mcp.md
```
# Oracle Pack — oraclepack (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-08/mcp
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md
- ticket_max_files: 6
- group_name: mcp
- group_slug: mcp
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/mcp/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08/mcp"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/mcp/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/mcp/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: mcp)

Reference: mcp
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/mcp/03-invariants-invariant-map.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/mcp/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: mcp)

Reference: mcp
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/mcp/Expose Oraclepack as MCP.md,.tickets/mcp/MCP Server for Oraclepack.md,.tickets/mcp/gaps-still-not-covered.md,.tickets/mcp/gaps_part2-mcp-builder.md,.tickets/mcp/oraclepack-MCP.md,.tickets/mcp/oraclepack_mcp_server.md".strip()
MAX = int("6")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'mcp'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/mcp/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: mcp)

Reference: mcp
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=mcp

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
[TRUNCATED]
```

docs/oracle-questions-2026-01-08/packs/misc.md
```
# Oracle Pack — oraclepack (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-08/misc
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md
- ticket_max_files: 4
- group_name: misc
- group_slug: misc
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/misc/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08/misc"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/misc/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/misc/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: misc)

Reference: misc
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/misc/03-invariants-invariant-map.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/misc/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: misc)

Reference: misc
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/misc/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: misc)

Reference: misc
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=misc

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/Oraclepack File Storage.md,.tickets/Oraclepack Schema Approach.md,.tickets/Oraclepack bash fix.md,.tickets/Publish OraclePack MCP.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'misc'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
[TRUNCATED]
```

docs/oracle-questions-2026-01-08/packs/other.md
```
# Oracle Pack — oraclepack (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-08/other
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md
- ticket_max_files: 4
- group_name: other
- group_slug: other
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/other/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08/other"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/other/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/other/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: other)

Reference: other
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/other/03-invariants-invariant-map.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/other/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: other)

Reference: other
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'other'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/other/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: other)

Reference: other
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=other

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/other/Oraclepack Pipeline Improvements.md,.tickets/other/Oraclepack Prompt Generator.md,.tickets/other/Oraclepack Workflow Enhancement.md,.tickets/other/Verbose Payload Rendering TUI.md".strip()
MAX = int("4")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

[TRUNCATED]
```

docs/oracle-questions-2026-01-08/packs/prd-tui.md
```
# Oracle Pack — oraclepack (Grouped Tickets Stage 1 — Direct Attach)

## Parsed args
- codebase_name: oraclepack
- out_dir: docs/oracle-questions-2026-01-08/prd-tui
- oracle_cmd: oracle
- oracle_flags: --files-report
- extra_files: 
- ticket_root: .tickets
- ticket_glob: **/*.md
- ticket_paths: .tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md
- ticket_max_files: 2
- group_name: PRD-TUI
- group_slug: prd-tui
- mode: tickets-grouped-direct

Notes (contract):
- Exactly one fenced `bash` block in this document.
- No other ``` fences anywhere.
- Exactly 20 steps, numbered 01..20 in order.
- Step header: `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes: `--write-output "docs/oracle-questions-2026-01-08/prd-tui/NN-<slug>.md"` (double quotes required).
- Steps must be self-contained and must not rely on shell variables created in previous steps.
- Each step must attach tickets directly (no `_tickets_bundle.md` dependency).
- Pack ends with a Coverage check section listing all 10 categories as OK or Missing(<step ids>).

```bash
set -euo pipefail

mkdir -p "docs/oracle-questions-2026-01-08/prd-tui"

# 01) ROI=4.8 impact=6 confidence=0.80 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/prd-tui/01-contracts-interfaces-ticket-surface.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #01  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.80, effort=1)

Question:
Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=4.6 impact=6 confidence=0.78 effort=1 horizon=Immediate category=contracts/interfaces reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/prd-tui/02-contracts-interfaces-integration-points.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #02  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: contracts/interfaces
Horizon: Immediate
ROI: 4.6 (impact=6, confidence=0.78, effort=1)

Question:
Using the attached tickets as the primary context, identify external integrations implied by the tickets; required config/contract changes; failure/timeout behavior; minimal compat-safe rollout.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=5.1 impact=7 confidence=0.74 effort=1 horizon=NearTerm category=invariants reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/prd-tui/03-invariants-invariant-map.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #03  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.1 (impact=7, confidence=0.74, effort=1)

Question:
Using the attached tickets as the primary context, extract system invariants implied by tickets (inputs/outputs, pack schema rules, step execution rules) and where to enforce them.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=5.0 impact=7 confidence=0.72 effort=2 horizon=NearTerm category=invariants reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/prd-tui/04-invariants-validation-boundaries.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #04  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: invariants
Horizon: NearTerm
ROI: 5.0 (impact=7, confidence=0.72, effort=2)

Question:
Using the attached tickets as the primary context, identify validation boundaries that must exist (ticket parsing, pack generation, pack validation); propose minimal validation plan.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=4.4 impact=6 confidence=0.78 effort=2 horizon=NearTerm category=caching/state reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
oracle   --files-report   --write-output "docs/oracle-questions-2026-01-08/prd-tui/05-caching-state-state-artifacts.md"   "${ticket_args[@]}"      -p "$(cat <<'PROMPT'
Strategist question #05  (ticket-driven, group: PRD-TUI)

Reference: prd-tui
Category: caching/state
Horizon: NearTerm
ROI: 4.4 (impact=6, confidence=0.78, effort=2)

Question:
Using the attached tickets as the primary context, identify state/artifacts that must be produced and preserved; schema/format expectations; stability/back-compat requirements.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=4.2 impact=6 confidence=0.75 effort=2 horizon=NearTerm category=caching/state reference=prd-tui

# tickets attached directly (deterministic; self-contained)
mapfile -t __tickets < <(python3 - <<'PY'
from __future__ import annotations
from pathlib import Path

TICKET_ROOT = ".tickets"
TICKET_GLOB = "**/*.md"
TICKET_PATHS = ".tickets/PRD-TUI/Oraclepack TUI Integration.md,.tickets/PRD-TUI/PRD-generator URL routing.md".strip()
MAX = int("2")

root = Path(TICKET_ROOT)

def lex_sorted(ps):
    return sorted((str(p) for p in ps), key=lambda s: s)

if TICKET_PATHS:
    tickets = [Path(p.strip()) for p in TICKET_PATHS.split(",") if p.strip()]
else:
    tickets = list(root.glob(TICKET_GLOB)) if root.exists() else []

tickets = [Path(p) for p in lex_sorted(tickets)]
if MAX and MAX > 0:
    tickets = tickets[:MAX]

for p in tickets:
    print(str(p))
PY
)

ticket_args=()
for p in "${__tickets[@]}"; do
  ticket_args+=(-f "$p")
done

if [ "${#ticket_args[@]}" -eq 0 ]; then
  echo "WARNING: no tickets resolved for group 'PRD-TUI'." >&2
fi

# extra_files appended literally (may be empty; may include -f/--file):
[TRUNCATED]
```

</source_code>