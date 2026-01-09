Parent Ticket:

* Title: Formalize enforceable LLM decision points for oraclepack generation/execution across Skills and MCP
* Summary:

  * The ticket enumerates a comprehensive set of pre-generation and runtime decision points (DP-01..DP-60) intended to improve pack generation and step execution quality. It also notes a preferred way to make these decision points “real, enforceable” by emitting small structured artifacts (JSON/YAML) that deterministic scripts consume, while keeping the emitted pack Markdown schema unchanged.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “an ‘extensive’ set of additional, unique LLM decision points you can introduce (or explicitly formalize)…”
* Global Constraints:

  * “keeping the emitted pack Markdown schema unchanged”
* Global Environment:

  * Unknown
* Global Evidence:

  * OpenAI Codex “Custom Prompts” documentation. ([OpenAI Developers][1])
  * Gemini CLI “Agent Skills” documentation. ([Gemini CLI][2])
  * MCP specification revision 2025-11-25. ([Model Context Protocol][3])
  * MCP tools specification (server/tools) revision 2025-06-18. ([Model Context Protocol][4])

Split Plan:

* Coverage Map:

  * “In Codex/Gemini CLI, the ‘live agent’ inference…” → Info-only
  * “In an MCP setup, the ‘live agent’ inference…” → Info-only
  * “Below is an ‘extensive’ set…” → Info-only
  * DP-01 Choose generator mode (tickets-grouped vs codebase-grouped vs gold) → T2
  * DP-02 Select root(s) to scan (ticket_root/code_root) → T2
  * DP-03 Decide include/exclude glob rules → T2
  * DP-04 Decide whether to treat “loose” items as first-class candidates → T2
  * DP-05 Decide whether to require “strict schema mode” for outputs → T2
  * DP-06 Choose max pack size strategy (by tokens/bytes/files) → T2
  * DP-07 Choose per-pack cap: number of tickets/files → T2
  * DP-08 Decide ticket/file canonical title extraction rule → T3
  * DP-09 Choose text preprocessing (stopwords, stemming, code tokens) → T3
  * DP-10 Decide duplicate threshold policy (strict vs lenient) → T3
  * DP-11 Decide duplicate merge strategy (merge vs link vs pick canonical) → T3
  * DP-12 Decide what “differences” to preserve when merging duplicates → T3
  * DP-13 Decide hierarchical topic taxonomy vs flat groups → T4
  * DP-14 Decide group count target → T4
  * DP-15 Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid) → T4
  * DP-16 Decide “ambiguous” assignment policy (multi-home vs best-fit) → T4
  * DP-17 Decide whether to create an “Unsorted / Needs triage” pack → T4
  * DP-18 Generate group names optimized for human scanning → T4
  * DP-19 Decide pack order (dependency, ROI, urgency, confidence) → T4
  * DP-20 Decide ticket order within pack (chronological vs dependency graph) → T4
  * DP-21 Select “context bundle” files to attach per pack → T5
  * DP-22 Decide whether to attach full files vs excerpts/summaries → T5
  * DP-23 Decide whether to generate a synthetic “pack brief” doc → T5
  * DP-24 Decide whether to include prior run artifacts (outputs) as inputs → T5
  * DP-25 Choose template variant (tickets vs codebase vs mixed) → T5
  * DP-26 Decide whether to inject organization standards into prompts → T5
  * DP-27 Decide step allocation across subtopics (steps per subdomain) → T5
  * DP-28 Choose prompt “stance” (audit-first vs implement-first vs design-first) → T5
  * DP-29 Choose “evidence bar” (strict citations vs lightweight) → T5
  * DP-30 Decide per-step required outputs (files changed, diffs, JSON, etc.) → T5
  * DP-31 Decide self-contained steps vs relying on previous outputs → T5
  * DP-32 Decide whether to add “ask-user” gates for missing critical inputs → T5
  * DP-33 Choose which MCP tools to call during generation (list/validate/generate) → T6
  * DP-34 Decide oracle model/engine selection and parameters → T6
  * DP-35 Decide whether to preflight oracle invocations (“--dry-run”) → T6
  * DP-36 Decide redaction policy for sensitive strings → T6
  * DP-37 Decide trace/correlation ID scheme for packs/steps → T6
  * DP-38 Decide which metrics/log events must be emitted per stage → T6
  * DP-39 Decide run strategy (run-all vs selective) → T7
  * DP-40 Decide concurrency / rate-limiting policy → T7
  * DP-41 Decide fail-fast vs continue collecting failures → T7
  * DP-42 Decide retry policy per failure type → T7
  * DP-43 Decide “prompt patch” for retries → T7
  * DP-44 Decide when to escalate to user for clarification → T7
  * DP-45 Decide whether to re-run earlier steps on contradictions → T7
  * DP-46 Decide acceptance criteria for a step output (format, completeness) → T8
  * DP-47 Decide whether to auto-validate produced artifacts (lint/tests/validate) → T8
  * DP-48 Decide how to synthesize step outputs into a final report → T9
  * DP-49 Decide PR/patch vs documentation-only output → T9
  * DP-50 Decide how to resolve conflicting recommendations across steps → T9
  * DP-51 Decide whether outputs meet policy/security standards before writing → T9
  * DP-52 Decide whether to reuse cached groupings/packs vs regenerate → T10
  * DP-53 Decide cache invalidation scope (one pack vs all) → T10
  * DP-54 Decide incident-style annotation on failures → T10
  * DP-55 Decide what artifacts to persist (manifests, intermediate summaries) → T10
  * DP-56 Decide how to shard outputs into “mini-packs” for follow-on runs → T10
  * DP-57 Decide naming/versioning for generated packs → T10
  * DP-58 Decide “next best tool call” in MCP (validate vs list vs run vs regenerate) → T10
  * DP-59 Decide whether to present diffs, file lists, or narrative only → T10
  * DP-60 Decide whether to extract new org heuristics into a reusable profile → T10
  * “make each decision point produce a small structured artifact (JSON/YAML)…” → T1
* Dependencies:

  * T2 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T3 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T4 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T5 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T6 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T7 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T8 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T9 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.
  * T10 depends on T1 because decision outputs should be emitted/consumed as “small structured artifact (JSON/YAML)”.

```ticket T1
T# Title: Make decision points enforceable via structured decision artifacts while keeping pack Markdown schema unchanged
Type: chore
Target Area: Pack generation/execution control plane (decision capture + deterministic consumption)
Summary:
- Establish the “typical pattern” described: each decision point emits a small structured artifact (JSON/YAML) and deterministic scripts consume it. Ensure the emitted pack Markdown schema remains unchanged. This provides a consistent way to formalize the DP-01..DP-60 decisions for both Skills-based flows and MCP tool-calling flows.
In Scope:
- “make each decision point produce a small structured artifact (JSON/YAML)”
- “make the deterministic scripts consume it”
- “keeping the emitted pack Markdown schema unchanged”
- Applicability across “skills” and “MCP”
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Decision points emit a small structured artifact (JSON/YAML).
- Deterministic scripts consume the artifact.
- Emitted pack Markdown schema remains unchanged.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Keep emitted pack Markdown schema unchanged.
Evidence:
- Text: “make each decision point produce a small structured artifact (JSON/YAML)… deterministic scripts consume it… pack Markdown schema unchanged.”
Open Items / Unknowns:
- Exact schema/fields for the structured artifact (JSON vs YAML specifics not provided).
- Where deterministic consumption is implemented (locations not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A structured decision artifact format exists that can represent DP outputs.
- Deterministic scripts consume the decision artifact to drive generation/execution flow.
- No change to the emitted pack Markdown schema is required to adopt the pattern.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “make each decision point produce a small structured artifact (JSON/YAML)”
- “make the deterministic scripts consume it”
- “keeping the emitted pack Markdown schema unchanged”
```

```ticket T2
T# Title: Implement pre-generation routing, scope, governance, and budgeting decision points (DP-01..DP-07)
Type: enhancement
Target Area: Pre-gen decision layer (routing/scope/governance/budgeting)
Summary:
- Add explicit, formalized pre-generation decisions for selecting generator mode, scan roots, include/exclude rules, loose-item handling, strict schema mode, and pack sizing/caps. Each decision produces the specified output/action for use by pack generation.
In Scope:
- DP-01 Choose generator mode (tickets-grouped vs codebase-grouped vs gold)
- DP-02 Select root(s) to scan (ticket_root/code_root)
- DP-03 Decide include/exclude glob rules
- DP-04 Decide whether to treat “loose” items as first-class candidates
- DP-05 Decide whether to require “strict schema mode” for outputs
- DP-06 Choose max pack size strategy (by tokens/bytes/files)
- DP-07 Choose per-pack cap: number of tickets/files
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Generator mode can be selected and recorded.
- Scan roots and include/exclude patterns can be selected and recorded.
- Loose-item handling policy and strict schema mode selection can be decided and recorded.
- Pack sizing strategy and per-pack caps can be decided and recorded.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Selected generator + params”, “Root paths”, “Include/exclude patterns”, “Sharding plan”, “Caps per pack”).
Evidence:
- DP-01..DP-07 rows (routing/scope/governance/budgeting)
Open Items / Unknowns:
- What “gold” mode entails (not provided).
- Default config values and where they are defined (not provided).
- Token/byte/file limits used for sizing (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-01 emits “Selected generator + params”.
- DP-02 emits “Root paths”.
- DP-03 emits “Include/exclude patterns”.
- DP-04 emits a “Loose-ticket policy”.
- DP-05 emits an “Enforce extra validation gates” decision.
- DP-06 emits a “Sharding plan”.
- DP-07 emits “Caps per pack”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Choose generator mode (tickets-grouped vs codebase-grouped vs gold)”
- “Select root(s) to scan (ticket_root/code_root)”
- “Choose max pack size strategy (by tokens/bytes/files)”
```

```ticket T3
T# Title: Implement pre-generation normalization and deduplication decision points (DP-08..DP-12)
Type: enhancement
Target Area: Pre-gen text normalization + dedup decisions
Summary:
- Add explicit, formalized decisions for canonical title extraction, preprocessing policy, duplicate thresholds, merge strategy, and delta preservation when merging duplicates. Each decision produces the specified output/action for use during grouping/packing.
In Scope:
- DP-08 Decide ticket/file canonical title extraction rule
- DP-09 Choose text preprocessing (stopwords, stemming, code tokens)
- DP-10 Decide duplicate threshold policy (strict vs lenient)
- DP-11 Decide duplicate merge strategy (merge vs link vs pick canonical)
- DP-12 Decide what “differences” to preserve when merging duplicates
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Canonical title extraction rules can be selected and recorded.
- Preprocessing policy can be selected and recorded.
- Duplicate threshold and merge strategy can be selected and recorded.
- Delta retention format/policy can be selected and recorded.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Canonical titles for clustering”, “Feature extraction policy”, “Threshold values”, “Merge plan + canonical selection”, “Delta retention format”).
Evidence:
- DP-08..DP-12 rows (normalization/dedup)
Open Items / Unknowns:
- Exact similarity signals used for deduplication (not provided).
- Definition of “diff size, recency, metadata” and their source (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-08 emits “Canonical titles for clustering”.
- DP-09 emits a “Feature extraction policy”.
- DP-10 emits “Threshold values”.
- DP-11 emits a “Merge plan + canonical selection”.
- DP-12 emits a “Delta retention format”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide ticket/file canonical title extraction rule”
- “Decide duplicate merge strategy (merge vs link vs pick canonical)”
- “Decide what ‘difdferences’ to preserve when merging duplicates”
```

```ticket T4
T# Title: Implement pre-generation clustering, naming, and ordering decision points (DP-13..DP-20)
Type: enhancement
Target Area: Pre-gen clustering + group naming + ordering decisions
Summary:
- Add explicit, formalized clustering decisions including taxonomy mode, group count, algorithm choice, ambiguous assignment policy, unsorted pack creation, group naming, and ordering of packs/tickets. Each decision produces the specified output/action for use in pack generation.
In Scope:
- DP-13 Decide hierarchical topic taxonomy vs flat groups
- DP-14 Decide group count target
- DP-15 Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid)
- DP-16 Decide “ambiguous” assignment policy (multi-home vs best-fit)
- DP-17 Decide whether to create an “Unsorted / Needs triage” pack
- DP-18 Generate group names optimized for human scanning
- DP-19 Decide pack order (dependency, ROI, urgency, confidence)
- DP-20 Decide ticket order within pack (chronological vs dependency graph)
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Taxonomy mode and group count target can be decided and recorded.
- Grouping algorithm and ambiguous assignment policy can be decided and recorded.
- Unsorted/triage pack creation can be decided and recorded.
- Group names and ordering (pack and ticket) can be decided and recorded.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Taxonomy mode”, “K groups target”, “Algorithm selection”, “Assignment rule”, “Extra pack creation”, “Group name strings”, “Pack sequence”, “Ordered ticket list”).
Evidence:
- DP-13..DP-20 rows (clustering/naming/ordering)
Open Items / Unknowns:
- How ROI/urgency/confidence are derived (not provided).
- Dependency detection inputs (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-13 emits a “Taxonomy mode” selection.
- DP-14 emits a “K groups target”.
- DP-15 emits an “Algorithm selection”.
- DP-16 emits an “Assignment rule”.
- DP-17 emits an “Extra pack creation” decision.
- DP-18 emits “Group name strings”.
- DP-19 emits a “Pack sequence”.
- DP-20 emits an “Ordered ticket list”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid)”
- “Decide whether to create an ‘Unsorted / Needs triage’ pack”
- “Decide pack order (by dependency, ROI, urgency, confidence)”
```

```ticket T5
T# Title: Implement pre-generation context, template, and prompting decision points (DP-21..DP-32)
Type: enhancement
Target Area: Pre-gen context selection + template/prompt shaping decisions
Summary:
- Add explicit, formalized decisions for context attachment (bundles, excerpts vs full, prior artifacts), template selection, step allocation, and prompting constraints (stance, evidence bar, per-step outputs, step dependency policy, ask-user gates). Each decision produces the specified output/action for use during pack generation and step prompting.
In Scope:
- DP-21 Select “context bundle” files to attach per pack
- DP-22 Decide full files vs excerpts/summaries
- DP-23 Decide whether to generate a synthetic “pack brief” doc
- DP-24 Decide whether to include prior run artifacts (outputs) as inputs
- DP-25 Choose template variant (tickets vs codebase vs mixed)
- DP-26 Decide whether to inject organization standards into prompts
- DP-27 Decide step allocation across subtopics
- DP-28 Choose prompt “stance” (audit-first vs implement-first vs design-first)
- DP-29 Choose “evidence bar” (strict citations vs lightweight)
- DP-30 Decide per-step required outputs (files changed, diffs, JSON, etc.)
- DP-31 Decide self-contained steps vs relying on previous step outputs
- DP-32 Decide whether to add “ask-user” gates for missing critical inputs
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Context selection decisions are captured and drive what is attached per pack.
- Template and step allocation decisions are captured and drive pack/step structure.
- Prompting decisions (stance/evidence/required outputs/dependency policy/ask-user gates) are captured and applied to step prompts.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Attachment list”, “Attachment granularity”, “Brief content”, “Reuse plan”, “Template choice”, “Prompt preamble content”, “Step budget map”, “Prompt style per step”, “Evidence requirements”, “Output spec per step”, “Dependency policy”, “Gate instructions in prompts”).
Evidence:
- DP-21..DP-32 rows (context/template/prompting)
Open Items / Unknowns:
- How “policy bundles/style rules” are sourced (not provided).
- What constitutes “missing critical inputs” detection (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-21 emits an “Attachment list”.
- DP-22 emits an “Attachment granularity” choice.
- DP-23 emits “Brief content” (or a decision to not generate it).
- DP-24 emits a “Reuse plan”.
- DP-25 emits a “Template choice”.
- DP-26 emits “Prompt preamble content” (or decision not to inject).
- DP-27 emits a “Step budget map”.
- DP-28 emits “Prompt style per step”.
- DP-29 emits “Evidence requirements”.
- DP-30 emits an “Output spec per step”.
- DP-31 emits a “Dependency policy”.
- DP-32 emits “Gate instructions in prompts”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Select ‘context bundle’ files to attach per pack”
- “Choose prompt ‘stance’ (audit-first vs implement-first vs design-first)”
- “Decide whether to add ‘ask-user’ gates for missing critical inputs”
```

```ticket T6
T# Title: Implement pre-generation tooling, compliance, and observability decision points (DP-33..DP-38)
Type: enhancement
Target Area: Pre-gen tool selection + redaction + tracing/metrics decisions
Summary:
- Add explicit, formalized decisions for which MCP tools to call during generation, oracle model/parameter selection, oracle preflight behavior, redaction policy, trace/correlation ID scheme, and required metrics/log events. Each decision produces the specified output/action for use during generation and downstream execution.
In Scope:
- DP-33 Choose which MCP tools to call during generation (list/validate/generate)
- DP-34 Decide oracle model/engine selection and parameters
- DP-35 Decide whether to preflight oracle invocations (“--dry-run” summary)
- DP-36 Decide redaction policy for sensitive strings
- DP-37 Decide trace/correlation ID scheme for packs/steps
- DP-38 Decide which metrics/log events must be emitted per stage
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Tool-call plan, oracle flags, and preflight decisions are captured and drive generation behavior.
- Redaction rules are captured and applied when handling sensitive strings.
- Trace/correlation IDs and per-stage metrics/log requirements are captured and used for observability.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Tool call plan”, “oracle flags (model/temperature)”, “Preflight on/off”, “Redaction rules”, “ID format + propagation plan”, “Instrumentation checklist”).
Evidence:
- DP-33..DP-38 rows (tooling/compliance/observability)
Open Items / Unknowns:
- Available MCP tool set and schemas (not provided here beyond “list/validate/generate”).
- Error/PII/secret patterns for redaction (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-33 emits a “Tool call plan”.
- DP-34 emits “oracle flags (model/temperature)” (or equivalent parameter selection decision).
- DP-35 emits “Preflight on/off”.
- DP-36 emits “Redaction rules”.
- DP-37 emits an “ID format + propagation plan”.
- DP-38 emits an “Instrumentation checklist”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Choose which MCP tools to call during generation (list/validate/generate)”
- “Decide redaction policy for sensitive strings in prompts/attachments”
- “Decide trace/correlation ID scheme for packs/steps”
```

```ticket T7
T# Title: Implement runtime execution planning, gating, recovery, and consistency decision points (DP-39..DP-45)
Type: enhancement
Target Area: Runtime execution control (planning/gating/retry/escalation/consistency)
Summary:
- Add explicit, formalized runtime decisions for selecting which steps to run, concurrency/rate limiting, fail-fast vs continue behavior, retry policy, retry prompt patching, escalation to user for clarification, and re-running earlier steps when contradictions are found. Each decision produces the specified output/action for use during execution.
In Scope:
- DP-39 Decide run strategy (run-all vs selective)
- DP-40 Decide concurrency / rate-limiting policy
- DP-41 Decide fail-fast vs continue collecting failures
- DP-42 Decide retry policy per failure type (tool error vs content error)
- DP-43 Decide “prompt patch” for retries (tighten constraints vs broaden context)
- DP-44 Decide when to escalate to user for clarification
- DP-45 Decide re-run earlier steps when later steps reveal contradictions
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Execution planning decisions select the step subset and parallelism behavior.
- Execution gating decisions define halt/continue behavior on failure.
- Recovery decisions define retries and retry prompt modifications.
- Consistency decisions define when to re-run earlier steps.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Step subset”, “Parallelism level”, “Fail-fast/continue”, “Retry plan”, “Revised prompt text”, “User question(s)”, “Re-run selection”).
Evidence:
- DP-39..DP-45 rows (execution planning/gating/recovery/consistency)
Open Items / Unknowns:
- Failure classification taxonomy beyond “tool error vs content error” (not provided).
- What constitutes a “contradiction” signal (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-39 emits a “Step subset” selection.
- DP-40 emits a “Parallelism level”.
- DP-41 emits “Fail-fast/continue”.
- DP-42 emits a “Retry plan”.
- DP-43 emits “Revised prompt text”.
- DP-44 emits “User question(s)” escalation decision.
- DP-45 emits a “Re-run selection”.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide retry policy per failure type (tool error vs content error)”
- “Decide ‘prompt patch’ for retries (tighten constraints vs broaden context)”
- “Decide whether to re-run earlier steps when later steps reveal contradictions”
```

```ticket T8
T# Title: Implement runtime quality control and validation decision points (DP-46..DP-47)
Type: enhancement
Target Area: Runtime QC (output acceptance + optional validation)
Summary:
- Add explicit, formalized decisions for step-output acceptance criteria and whether to auto-validate produced artifacts via lint/tests/validate. Each decision produces the specified output/action for use in runtime quality control.
In Scope:
- DP-46 Decide acceptance criteria for a step output (format, completeness)
- DP-47 Decide whether to auto-validate produced artifacts (lint/tests/validate)
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Step-output acceptance criteria can be decided and applied for pass/fail.
- Auto-validation decision can be made and applied using available checks.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Pass/fail verdict”, “Validation commands”).
Evidence:
- DP-46..DP-47 rows (quality control)
Open Items / Unknowns:
- Which validators are available and what commands they require (not provided).
- What “format, completeness” requirements are for outputs (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-46 yields a “Pass/fail verdict” based on decided acceptance criteria.
- DP-47 yields a decision on validation and the associated “Validation commands” (or a no-validation decision).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide acceptance criteria for a step output (format, completeness)”
- “Decide whether to auto-validate produced artifacts (lint/tests/validate)”
```

```ticket T9
T# Title: Implement runtime post-processing and governance decision points (DP-48..DP-51)
Type: enhancement
Target Area: Runtime post-processing + governance checks
Summary:
- Add explicit, formalized decisions for synthesizing step outputs into a final report, selecting PR/patch vs documentation-only mode, resolving conflicting recommendations, and checking outputs against policy/security standards before writing. Each decision produces the specified output/action for use in post-processing and governance.
In Scope:
- DP-48 Decide how to synthesize step outputs into a final report
- DP-49 Decide PR/patch vs documentation-only output
- DP-50 Decide how to resolve conflicting recommendations across steps
- DP-51 Decide whether outputs meet policy/security standards before writing
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- A synthesis approach can be decided and applied to generate a final report artifact.
- Output mode (PR/patch vs docs-only) can be decided and applied.
- Conflicts can be detected and a resolution rationale can be decided and applied.
- Policy/security gating decision can be made before writing outputs.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Summary artifact”, “Output mode”, “Resolution rationale”, “Block/allow + edits”).
Evidence:
- DP-48..DP-51 rows (post-processing/governance)
Open Items / Unknowns:
- How conflicts are detected (not provided).
- Which policy bundle applies and how it is provided (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-48 yields a “Summary artifact” decision and corresponding synthesis result.
- DP-49 yields an “Output mode” decision.
- DP-50 yields a “Resolution rationale” for detected conflicts.
- DP-51 yields a “Block/allow + edits” governance decision prior to writing outputs.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide how to synthesize step outputs into a final report”
- “Decide whether to generate PR/patch vs documentation-only output”
- “Decide whether outputs meet policy/security standards before writing”
```

```ticket T10
T# Title: Implement runtime caching, observability artifacts, packaging/versioning, UX, and learning-loop decision points (DP-52..DP-60)
Type: enhancement
Target Area: Runtime caching + artifact persistence + packaging/versioning + UX + learning loop
Summary:
- Add explicit, formalized decisions for cache reuse/invalidation, incident-style failure annotation, artifact persistence, sharding outputs into mini-packs, naming/versioning, selecting next best MCP tool call, presentation format (diffs/files/narrative), and extracting reusable heuristics for future runs. Each decision produces the specified output/action for use in runtime operations and operator/user experience.
In Scope:
- DP-52 Decide whether to reuse cached groupings/packs vs regenerate
- DP-53 Decide cache invalidation scope (one pack vs all)
- DP-54 Decide incident-style annotation (root cause, repro, next action) on failures
- DP-55 Decide what artifacts to persist (manifests, intermediate summaries)
- DP-56 Decide how to shard outputs into “mini-packs” for follow-on runs
- DP-57 Decide naming/versioning for generated packs
- DP-58 Decide “next best tool call” in MCP (validate vs list vs run vs regenerate)
- DP-59 Decide whether to present diffs, file lists, or narrative only
- DP-60 Decide whether to extract new org heuristics into a reusable profile
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Cache reuse and invalidation decisions can be made and applied.
- Failure annotation and persistence decisions can be made and applied.
- Output sharding and naming/versioning decisions can be made and applied.
- MCP “next best tool call” and presentation format decisions can be made and applied.
- Learning-loop extraction decision can be made and applied.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Decisions should produce the “Output / action produced” listed for each DP (e.g., “Cache decision”, “Invalidation plan”, “Failure note artifact”, “Persist list”, “Mini-pack plan”, “Pack names + version tags”, “Tool invocation”, “Presentation format”, “Proposed profile snippet”).
Evidence:
- DP-52..DP-60 rows (caching/observability/packaging/UX/learning loop)
Open Items / Unknowns:
- Cache key design and what inputs constitute “repo diffs” (not provided).
- Failure note format and where it is stored (not provided).
- Version tag format specifics (not provided).
Risks / Dependencies:
- Depends on T1 if adopting the structured decision artifact pattern.
Acceptance Criteria:
- DP-52 yields a “Cache decision”.
- DP-53 yields an “Invalidation plan”.
- DP-54 yields a “Failure note artifact”.
- DP-55 yields a “Persist list”.
- DP-56 yields a “Mini-pack plan”.
- DP-57 yields “Pack names + version tags”.
- DP-58 yields a “Tool invocation” decision.
- DP-59 yields a “Presentation format” decision.
- DP-60 yields a “Proposed profile snippet” decision.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Decide cache invalidation scope (one pack vs all)”
- “Decide how to shard outputs into ‘mini-packs’ for follow-on runs”
- “Decide whether to extract new org heuristics from this run into a reusable profile”
```

[1]: https://developers.openai.com/codex/custom-prompts/?utm_source=chatgpt.com "Custom Prompts"
[2]: https://geminicli.com/docs/cli/skills/?utm_source=chatgpt.com "Agent Skills | Gemini CLI"
[3]: https://modelcontextprotocol.io/specification/2025-11-25?utm_source=chatgpt.com "Specification"
[4]: https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com "Tools"
