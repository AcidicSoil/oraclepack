1) Summary

----------

* Use this skill when you want to extend Codex with a repeatable, task-specific workflow packaged as a “skill” (instructions + optional scripts/resources), rather than re-prompting the same multi-step process each time. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Best fit for generating multiple, deterministic, per-domain/per-topic “mini-packs” for a codebase (grouped output + validation), using a skill folder that includes `SKILL.md` plus optional scripts/references.

    oraclepack-codebase-pack-groupe…

    [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Improves reliability via progressive disclosure: Codex loads only skill metadata upfront, then reads full instructions/resources only when invoked (explicitly or implicitly). [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

* Supports portability/versioning: skills are folders of instructions/scripts/resources that can be reused across compatible agents and checked into repo/user/admin scopes. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

1) Optimal usage moments mapped to lifecycle phases

---------------------------------------------------

Discovery / Scoping: When you need fast, repeatable extraction of “what’s here” across a codebase, but want results split by domain/topic to avoid monolithic context.

Architecture / Design: When you want deterministic grouping and standardized pack structure so downstream analysis is consistent and auditable.

Implementation: When code changes frequently and you want a one-command regeneration of grouped packs with the same schema and constraints.

Testing / Validation: When you need enforceable checks (schema, attachments, size limits) before trusting generated packs or using them in CI.

Release / Rollout: When you want a stable, version-controlled artifact set (per group) to gate releases, document surfaces, or capture invariants.

Operations / Maintenance: When incidents or regressions are localized to a subsystem and you want to quickly re-slice context by directory/domain and re-run analysis consistently.

(Invocation model note: Codex can use a skill by explicit invocation (e.g., selecting from `/skills` or typing `$…`) or by implicitly selecting it when the task matches the skill description.) [OpenAI Developers](https://developers.openai.com/codex/skills/)

1) Example library (12 examples)

--------------------------------

| Phase | Goal | Inputs required (files, constraints, audience, format) | Raw prompt (before) | Optimized prompt (after) | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Discovery / Scoping | Generate grouped packs for an unfamiliar repo | {context} repo path; constraints on size; audience = eng; format = commands + outputs | “Make oracle packs for this repo.” | “Use $oraclepack-codebase-pack-grouped to generate Stage-1 grouped packs for {codebase\_name}. {inputs}: code\_root={code\_root}, out\_dir={out\_dir}, group\_mode=subdir+infer. {constraints}: cap group\_max\_files={group\_max\_files}, group\_max\_chars={group\_max\_chars}, include\_exts={include\_exts}. {acceptance\_criteria}: produce packs/\*.md + \_groups.json and list largest group. {format}: steps + exact commands. {deadline}: {deadline}.” | Mentions explicit skill use; provides concrete KEY=value inputs; defines caps; requests exact artifacts and where they land. |
| Discovery / Scoping | Identify public surface area per subsystem | {context} desired subsystems; constraints; audience = platform team; format = pack set + “how to run” | “Analyze the API surface.” | “Use $oraclepack-codebase-pack-grouped to split the codebase into per-domain packs. Then describe which pack(s) cover public surfaces (CLI/TUI/API). {inputs}: code\_root={code\_root}, exclude\_glob={exclude\_glob}. {format}: (1) generated pack list by group (2) recommended run order (3) what each pack will answer. {acceptance\_criteria}: each group mapped to expected surface area.” | Output ties groups to surfaces, not just “generated packs”; includes exclusions to avoid noise. |
| Architecture / Design | Tune grouping rules to match team domains | {context} desired domain map; constraints on determinism; audience = tech leads; format = updated parameters | “Group these files better.” | “Using the existing grouping approach, propose a deterministic parameter set for $oraclepack-codebase-pack-grouped that aligns with {context} domain boundaries. {inputs}: current repo layout summary, known domain dirs, known ‘loose files’. {constraints}: deterministic ordering; no manual per-file curation unless via code\_paths. {format}: recommended KEY=value args + rationale. {acceptance\_criteria}: explains how loose files resolve; avoids non-deterministic heuristics.” | Provides a concrete args profile and explains deterministic behavior for “loose” files. |
| Architecture / Design | Define skill storage/scope strategy for a mono-repo | {context} mono-repo layout; constraints on precedence; audience = tooling; format = recommended locations | “Where should we put this skill?” | “Recommend where to store this skill so teams can override safely. {inputs}: {context} repo structure + ownership boundaries. {constraints}: follow Codex skill scopes/precedence; prefer version-controlled repo skills. {format}: path recommendations and why (repo root vs nested), plus override plan. {acceptance\_criteria}: uses the documented skill locations and precedence model.” [OpenAI Developers](https://developers.openai.com/codex/skills/) | Uses Codex scope model (repo/user/admin/system) and explains override/precedence. [OpenAI Developers](https://developers.openai.com/codex/skills/) |
| Implementation | Regenerate packs after refactor with minimal noise | {context} refactor areas; constraints on excluding build outputs; audience = eng; format = commands | “Re-run packs after changes.” | “Use $oraclepack-codebase-pack-grouped to regenerate packs for {codebase\_name} focusing on {context}. {inputs}: code\_root={code\_root}, include\_exts={include\_exts}, ignore\_dirs={ignore\_dirs}, exclude\_glob={exclude\_glob}. {constraints}: code\_max\_files={code\_max\_files}; deterministic sort. {format}: exact python command + expected output paths. {acceptance\_criteria}: excludes build/venv artifacts; outputs stable pack slugs.” | Shows exact invocation, includes ignore/exclude controls, and expects stable outputs. |
| Implementation | Create a narrow “single-area” pack for a hotfix | {context} explicit paths; constraints; audience = eng; format = one pack | “Make a pack just for these files.” | “Use $oraclepack-codebase-pack-grouped but force explicit file selection. {inputs}: code\_paths={code\_paths\_csv}, out\_dir={out\_dir}, group\_mode=subdir (or infer off). {constraints}: group\_max\_files={group\_max\_files}. {format}: one generated pack plus how to validate it. {acceptance\_criteria}: pack attaches only the provided files; no extra discovery.” | Uses code\_paths to override globbing; verifies only specified files are included. |
| Testing / Validation | Validate every pack is schema-safe and directly attachable | {context} out\_dir; constraints; audience = CI/tooling; format = command list | “Check these packs.” | “For packs in {out\_dir}/packs, run the skill’s validators and report failures. {inputs}: {out\_dir}. {constraints}: must fail fast; include exact failing step(s). {format}: shell commands + pass/fail summary. {acceptance\_criteria}: every pack checked; failures actionable.” | Includes concrete validation actions and expects actionable failure output. |
| Testing / Validation | Enforce size/token caps by splitting groups | {context} largest group; constraints = token budgets; audience = eng; format = recommended caps + rerun | “This pack is too big.” | “Use $oraclepack-codebase-pack-grouped and adjust caps to meet {constraints} size limits. {inputs}: current group stats, group\_max\_files, group\_max\_chars. {format}: new KEY=value args + rerun command + expected part naming. {acceptance\_criteria}: largest group splits into part 1..N deterministically; no missing files.” | Produces a parameter change and deterministic split strategy, not vague advice. |
| Release / Rollout | Produce release-ready, versioned analysis artifacts | {context} release tag/branch; constraints on reproducibility; audience = release mgr; format = artifact checklist | “Generate release docs.” | “Generate grouped Stage-1 packs for {context} release candidate using $oraclepack-codebase-pack-grouped. {inputs}: code\_root={code\_root}, out\_dir={out\_dir}. {constraints}: reproducible outputs; no env-specific paths beyond out\_dir. {format}: artifact checklist (packs + \_groups.json) + how to run them in order. {acceptance\_criteria}: re-running yields same pack set and slugs for same tree.” | Emphasizes reproducibility and a concrete artifact set. |
| Release / Rollout | Gate release on invariants/contract deltas per domain | {context} “what changed”; constraints; audience = eng/release; format = pack mapping + run plan | “Tell me if contracts changed.” | “Use $oraclepack-codebase-pack-grouped to slice by domain, then outline how each pack will be used to detect contracts/interfaces changes across {context}. {inputs}: baseline vs head refs (described), out\_dir. {format}: (1) which groups matter (2) run sequence (3) what output files to diff. {acceptance\_criteria}: explicit diff targets per group; minimal false positives via excludes.” | Produces a concrete diff plan tied to group outputs. |
| Operations / Maintenance | Incident triage localized to one subsystem | {context} incident area; constraints = speed; audience = oncall; format = narrow regeneration + run | “Help me debug this module.” | “Use $oraclepack-codebase-pack-grouped to generate packs limited to {context} subsystem. {inputs}: code\_root={code\_root}, exclude\_glob={exclude\_glob}, include\_exts={include\_exts}. {constraints}: smaller caps for speed. {format}: exact command + identify which pack(s) to run first. {acceptance\_criteria}: minimal scope; outputs point to likely failure modes and missing signals.” | Keeps scope tight and yields a prioritized run order for the most relevant packs. |
| Operations / Maintenance | Maintain the skill as a reusable capability across teams | {context} ownership model; constraints = portability; audience = platform; format = maintenance playbook | “How do we keep this skill usable?” | “Propose a maintenance playbook for this skill folder so it stays portable and discoverable. {inputs}: {context} team boundaries + repo layout. {constraints}: skill must remain a folder of instructions/scripts/resources; version-controlled; documented invocation. {format}: checklist for updates + review gates. {acceptance\_criteria}: aligns to Agent Skills concept (discoverable folder bundle) and Codex skill loading/scopes.” [Agent Skills+1](https://agentskills.io/home) | Playbook explicitly reflects the “skill folder” model and Codex scope/precedence rules. [Agent Skills+1](https://agentskills.io/home) |

1) Non-goals / anti-patterns (do not use this skill)

----------------------------------------------------

1. Ticket-driven packaging: if the primary inputs are `.tickets/` or issue threads (not repo code), use a ticket-pack skill instead of a codebase-grouped skill.

    oraclepack-codebase-pack-groupe…

2. One-off questions where a full grouped pack set is overkill (e.g., “what does this function do?”).

3. Non-deterministic or manual grouping requirements (e.g., “group by whatever seems interesting today”) where reproducibility is required.

4. Situations where you cannot provide/attach or point to the underlying codebase paths (no evidence to pack).

5. Tasks that primarily require editing/building code rather than generating/validating grouped analysis artifacts (use a coding/implementation workflow instead of a packaging workflow).

---
