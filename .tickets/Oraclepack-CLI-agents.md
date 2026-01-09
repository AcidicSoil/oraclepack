
In Codex/Gemini CLI, the “live agent” inference is the model deciding what actions to take given the skill context (skills are discovered/activated and the model can inspect skill assets). [OpenAI Developers+2OpenAI Developers+2](https://developers.openai.com/codex/custom-prompts/?utm_source=chatgpt.com) In an MCP setup, the “live agent” inference is the model deciding which tools to call (and with what arguments) based on the tools’ schemas. [Model Context Protocol+3Model Context Protocol+3Model Context Protocol+3](https://modelcontextprotocol.io/specification/2025-11-25?utm_source=chatgpt.com)

Below is an “extensive” set of additional, unique LLM decision points you can introduce (or explicitly formalize) beyond “classify tickets/domains then feed into grouping”.

| ID | Stage | Class | LLM decision point | Inputs considered | Output / action produced |
| --- | --- | --- | --- | --- | --- |
| DP-01 | Pre-gen | Routing | Choose generator mode (tickets-grouped vs codebase-grouped vs gold) | User goal, repo state, available skills/tools | Selected generator + params |
| DP-02 | Pre-gen | Scope | Select root(s) to scan (ticket\_root/code\_root) | Working dir tree, config defaults | Root paths |
| DP-03 | Pre-gen | Scope | Decide include/exclude glob rules | Repo conventions, noise directories | Include/exclude patterns |
| DP-04 | Pre-gen | Scope | Decide whether to treat “loose” items as first-class candidates | Ticket placement quality, density | Loose-ticket policy |
| DP-05 | Pre-gen | Governance | Decide whether to require “strict schema mode” for outputs | Target environment strictness | Enforce extra validation gates |
| DP-06 | Pre-gen | Budgeting | Choose max pack size strategy (by tokens/bytes/files) | Model context limits, file sizes | Sharding plan |
| DP-07 | Pre-gen | Budgeting | Choose per-pack cap: number of tickets/files | Expected coherence vs coverage | Caps per pack |
| DP-08 | Pre-gen | Normalization | Decide ticket/file canonical title extraction rule | Noisy headings, filenames | Canonical titles for clustering |
| DP-09 | Pre-gen | Normalization | Choose text preprocessing (stopwords, stemming, code tokens) | Domain vocabulary | Feature extraction policy |
| DP-10 | Pre-gen | Dedup | Decide duplicate threshold policy (strict vs lenient) | Similarity signals | Threshold values |
| DP-11 | Pre-gen | Dedup | Decide duplicate merge strategy (merge vs link vs pick canonical) | Diff size, recency, metadata | Merge plan + canonical selection |
| DP-12 | Pre-gen | Dedup | Decide what “differences” to preserve when merging duplicates | Patch-like deltas | Delta retention format |
| DP-13 | Pre-gen | Clustering | Decide whether to use hierarchical topic taxonomy vs flat groups | Repo size, domain diversity | Taxonomy mode |
| DP-14 | Pre-gen | Clustering | Decide group count target | Ticket count, entropy | K groups target |
| DP-15 | Pre-gen | Clustering | Decide grouping algorithm choice (heuristic vs LLM-labeled vs hybrid) | Determinism needs, accuracy | Algorithm selection |
| DP-16 | Pre-gen | Clustering | Decide “ambiguous” assignment policy (multi-home vs best-fit) | Similarity ties | Assignment rule |
| DP-17 | Pre-gen | Clustering | Decide whether to create an “Unsorted / Needs triage” pack | Low-confidence items | Extra pack creation |
| DP-18 | Pre-gen | Naming | Generate group names optimized for human scanning | Group summaries | Group name strings |
| DP-19 | Pre-gen | Ordering | Decide pack order (by dependency, ROI, urgency, confidence) | Ticket metadata, heuristics | Pack sequence |
| DP-20 | Pre-gen | Ordering | Decide ticket order within pack (chronological vs dependency graph) | References among tickets | Ordered ticket list |
| DP-21 | Pre-gen | Context | Select “context bundle” files to attach per pack | Code references, paths mentioned | Attachment list |
| DP-22 | Pre-gen | Context | Decide whether to attach full files vs excerpts/summaries | Token budget, file size | Attachment granularity |
| DP-23 | Pre-gen | Context | Decide whether to generate a synthetic “pack brief” doc | Need for shared framing | Brief content |
| DP-24 | Pre-gen | Context | Decide whether to include prior run artifacts (outputs) as inputs | Existing out\_dir contents | Reuse plan |
| DP-25 | Pre-gen | Template | Choose template variant (tickets vs codebase vs mixed) | Pack type, audience | Template choice |
| DP-26 | Pre-gen | Template | Decide whether to inject organization standards into prompts (not pack shape) | Policy bundles, style rules | Prompt preamble content |
| DP-27 | Pre-gen | Template | Decide step allocation across subtopics (how many steps per subdomain) | Group composition | Step budget map |
| DP-28 | Pre-gen | Prompting | Choose prompt “stance” (audit-first vs implement-first vs design-first) | Risk, stage | Prompt style per step |
| DP-29 | Pre-gen | Prompting | Choose “evidence bar” (strict citations vs lightweight) | Audience and stakes | Evidence requirements |
| DP-30 | Pre-gen | Prompting | Decide per-step required outputs (files changed, diffs, JSON, etc.) | Tooling integration, downstream parser | Output spec per step |
| DP-31 | Pre-gen | Prompting | Decide if each step should be self-contained or rely on previous step outputs | Runtime environment, caching | Dependency policy |
| DP-32 | Pre-gen | Prompting | Decide whether to add “ask-user” gates for missing critical inputs | Missing file detection | Gate instructions in prompts |
| DP-33 | Pre-gen | Tooling | Choose which MCP tools to call during generation (list/validate/generate) | Tool availability, cost | Tool call plan [Model Context Protocol+1](https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com) |
| DP-34 | Pre-gen | Tooling | Decide oracle model/engine selection and parameters | Cost, latency, quality | `oracle` flags (model/temperature) |
| DP-35 | Pre-gen | Tooling | Decide whether to preflight `oracle` invocations (`--dry-run` summary) | Compatibility risk | Preflight on/off |
| DP-36 | Pre-gen | Compliance | Decide redaction policy for sensitive strings in prompts/attachments | Secrets risk | Redaction rules |
| DP-37 | Pre-gen | Observability | Decide trace/correlation ID scheme for packs/steps | Downstream logging needs | ID format + propagation plan |
| DP-38 | Pre-gen | Observability | Decide which metrics/log events must be emitted per stage | Debuggability requirements | Instrumentation checklist |
| DP-39 | Runtime | Execution planning | Decide run strategy (run-all vs selective) | Confidence map, cost | Step subset |
| DP-40 | Runtime | Execution planning | Decide concurrency / rate-limiting policy | Provider limits | Parallelism level |
| DP-41 | Runtime | Execution gating | Decide whether to halt on first failure vs continue collecting failures | Failure criticality | Fail-fast/continue |
| DP-42 | Runtime | Recovery | Decide retry policy per failure type (tool error vs content error) | Error classification | Retry plan |
| DP-43 | Runtime | Recovery | Decide “prompt patch” for retries (tighten constraints vs broaden context) | Failure analysis | Revised prompt text |
| DP-44 | Runtime | Recovery | Decide when to escalate to user for clarification | Low confidence / missing inputs | User question(s) |
| DP-45 | Runtime | Consistency | Decide whether to re-run earlier steps when later steps reveal contradictions | Cross-step inconsistency | Re-run selection |
| DP-46 | Runtime | Quality control | Decide acceptance criteria for a step output (format, completeness) | Output parsing/validation | Pass/fail verdict |
| DP-47 | Runtime | Quality control | Decide whether to auto-validate produced artifacts (lint/tests/validate) | Available checks | Validation commands |
| DP-48 | Runtime | Post-processing | Decide how to synthesize step outputs into a final report | Audience, required structure | Summary artifact |
| DP-49 | Runtime | Post-processing | Decide whether to generate PR/patch vs documentation-only output | Repo permissions, user intent | Output mode |
| DP-50 | Runtime | Post-processing | Decide how to resolve conflicting recommendations across steps | Conflicts detected | Resolution rationale |
| DP-51 | Runtime | Governance | Decide whether outputs meet policy/security standards before writing | Policy bundle | Block/allow + edits |
| DP-52 | Runtime | Caching | Decide whether to reuse cached groupings/packs vs regenerate | Cache keys, repo diffs | Cache decision |
| DP-53 | Runtime | Caching | Decide cache invalidation scope (one pack vs all) | Changed inputs | Invalidation plan |
| DP-54 | Runtime | Observability | Decide incident-style annotation (root cause, repro, next action) on failures | Error logs, outputs | Failure note artifact |
| DP-55 | Runtime | Observability | Decide what artifacts to persist (manifests, intermediate summaries) | Debug needs | Persist list |
| DP-56 | Runtime | Packaging | Decide how to shard outputs into “mini-packs” for follow-on runs | Token limits, dependencies | Mini-pack plan |
| DP-57 | Runtime | Packaging | Decide naming/versioning for generated packs | Date, domain, stability | Pack names + version tags |
| DP-58 | Runtime | UX | Decide “next best tool call” in MCP (validate vs list vs run vs regenerate) | Current state | Tool invocation [Model Context Protocol+1](https://modelcontextprotocol.io/docs/develop/build-server?utm_source=chatgpt.com) |
| DP-59 | Runtime | UX | Decide whether to present diffs, file lists, or narrative only | Reviewer preference | Presentation format |
| DP-60 | Runtime | Learning loop | Decide whether to extract new org heuristics from this run into a reusable profile | Repeated patterns | Proposed profile snippet |

If you want these to be real, enforceable decision points (not ad hoc), the typical pattern is: make each decision point produce a small structured artifact (JSON/YAML) and then make the deterministic scripts consume it, keeping the emitted pack Markdown schema unchanged. This fits both “skills” (model chooses actions using skill assets) [Gemini CLI+1](https://geminicli.com/docs/cli/skills/?utm_source=chatgpt.com) and MCP (model-controlled tool calls with typed schemas). [Model Context Protocol+1](https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com)

---

