1) Summary

----------

* Use this skill to generate multiple **runner-ingestible “Stage-1” oracle packs** from `.tickets/`, split into focused **topic/domain mini-packs** with **direct ticket attachments** (no bundle dependency).

    oraclepack-tickets-pack-grouped…

* Improves reliability by enforcing a **deterministic workflow**, a strict **20-step pack schema**, and mandatory **pack + attachment lint validation**.

    oraclepack-tickets-pack-grouped…

* Reduces noise by optionally detecting **possible duplicate tickets** (Jaccard + overlap) and applying a chosen dedupe mode (`report/prune/merge`).

    oraclepack-tickets-pack-grouped…

* Fits the Agent Skills model: skills package instructions/resources/scripts so agents can load procedural knowledge “on demand” and execute repeatable workflows consistently. [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

1) Optimal usage moments (mapped to lifecycle phases)

-----------------------------------------------------

Intake / Backlog triage
Use when you have many tickets (often across subdirs) and need them organized into coherent, runnable analysis packs per domain/topic. The skill’s deterministic grouping (`subdir+infer`) and optional dedupe are designed for this.

oraclepack-tickets-pack-grouped…

Requirements / Discovery
Use when you need to derive interfaces, invariants, state, caching, etc. from tickets as primary evidence, but want results separated by domain to avoid cross-contamination. The produced packs standardize the question set across 20 steps/categories.

oraclepack-tickets-pack-grouped…

Architecture / Design
Use when teams want per-area “design brief” outputs (contracts, validation boundaries, rollout flags) generated consistently for each domain group, enabling parallel review.

oraclepack-tickets-pack-grouped…

Implementation planning / Task execution setup
Use when you’re preparing an agent workflow that depends on stable, validated artifacts (pack files + manifests) and needs predictable inputs/outputs. This aligns with “skills as packaged workflows + scripts/resources.” [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

QA / Validation gates
Use when you need a guardrail that packs are schema-safe before execution (exact fence expectations, step count, direct-ticket attachment checks).

oraclepack-tickets-pack-grouped…

Release / Ops / Maintenance
Use when tickets evolve continuously and you need regeneration with deterministic behavior (lexicographic ordering, stable tie-breaks, explicit caps) to keep outputs comparable across runs.

oraclepack-tickets-pack-grouped…

1) Example library

------------------

| Phase | Goal | Inputs required (files, constraints, audience, format) | Raw prompt (before) | Optimized prompt (after) | Acceptance criteria |
| --- | --- | --- | --- | --- | --- |
| Intake / triage | Split `.tickets/` into domain packs automatically | Files: `.tickets/**/*.md`; Constraints: deterministic grouping; Format: packs + manifest | “Group my tickets and make oracle packs.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets ticket\_glob=\*\*/_.md group\_mode=subdir+infer dedupe\_mode=report out\_dir={out\_dir}. {constraints}: group\_max\_files=25 group\_max\_chars=200000. {format}: produce packs/_.md + \_groups.json + manifest.json. {acceptance\_criteria}: deterministic grouping, no bundle attachments, packs validate. {deadline}” | Multiple packs created per group; `_groups.json` + `manifest.json` exist; validation/lint passes. |
| Intake / triage | Detect possible duplicates before grouping | Files: `.tickets/**/*.md`; Constraints: report-only; Audience: maintainers | “Find duplicate tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets dedupe\_mode=report dedupe\_jaccard=0.55 dedupe\_overlap\_hi=0.80 dedupe\_overlap\_lo=0.70 dedupe\_body\_chars=2000. {constraints}: do not drop tickets; emit duplicate reports. {format}: ensure \_dupes\_possible.json + \_duplicates.json + \_dedupe\_plan.json. {acceptance\_criteria}: clusters + canonical + delta/redundant identified. {deadline}” | Duplicate artifacts emitted; clusters include canonical and classification; no tickets lost from assignments. |
| Requirements / discovery | Generate per-domain “public surface changes” outputs | Files: `.tickets/**`; Audience: product/eng; Format: per-pack outputs | “Answer strategist Q1 from these tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets group\_mode=subdir+infer out\_dir={out\_dir}. {constraints}: each pack must be 20 steps; direct `-f` ticket attachments only. {format}: one pack per domain producing 01–20 outputs under each group out\_dir. {acceptance\_criteria}: step 01 answers contracts/interfaces with evidence-cited bullets. {deadline}” | Packs exist per domain; each pack contains step 01 prompt; outputs pathing is group-specific. |
| Requirements / discovery | Keep groups small enough for model limits | Files: `.tickets/**`; Constraints: strict caps | “Make sure it fits context windows.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets group\_max\_files=15 group\_max\_chars=120000. {constraints}: enforce size control; split oversize groups into -part-XX packs. {format}: packs per group/part + manifest updated. {acceptance\_criteria}: no group exceeds caps; parts are deterministically chunked. {deadline}” | Large groups are split; pack names include part suffix; manifest records attached\_paths per part. |
| Architecture / design | Produce validation boundary plan per domain | Files: `.tickets/**`; Audience: platform team | “Propose validation boundaries.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: keep non-ticket attachments to minimum; direct tickets only. {format}: ensure step 04 (validation boundaries) exists in each domain pack. {acceptance\_criteria}: step 04 returns boundaries + minimal plan + one concrete experiment. {deadline}” | Each pack includes step 04; step output format matches required 4 sections; content is domain-scoped. |
| Architecture / design | Produce caching/state model and cache invalidation per domain | Files: `.tickets/**`; Audience: infra | “Analyze caching and state.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: preserve determinism; use dedupe\_mode=prune to reduce redundancy. {format}: ensure steps 05 and 06 are present and write outputs. {acceptance\_criteria}: defines cache keys + invalidation + correctness risks; evidence-cited. {deadline}” | Step 05/06 exist for every pack; dedupe reduces redundant attachments; outputs are written per group. |
| Implementation planning | Generate packs for only a subset of tickets (explicit list) | Files: explicit paths list; Constraint: ignore glob | “Only use these tickets.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_paths={ticket\_list\_csv} (comma-separated), ticket\_root=.tickets, ticket\_glob ignored. {constraints}: lexicographic ordering; ticket\_max\_files={n}. {format}: packs generated from explicit list only. {acceptance\_criteria}: no other tickets included; manifest lists only provided paths. {deadline}” | Packs include only the explicit files; order is stable; validation passes. |
| Implementation planning | Append consistent non-ticket evidence files to each oracle call | Files: repo evidence file(s); Constraint: 0–1 extra per step | “Add these extra files everywhere.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets extra\_files={extra\_files\_flags}. {constraints}: extra\_files appended literally, on its own line comment; keep explicit non-ticket attachments minimal. {format}: regenerate packs and confirm extra\_files included in each step invocation. {acceptance\_criteria}: every step includes extra\_files line; direct tickets still attached. {deadline}” | Every step includes extra files exactly as provided; no bundle references introduced; lint passes. |
| QA / validation | Validate every generated pack before running | Files: generated packs; Constraint: fail fast | “Make sure the packs are valid.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: out\_dir={out\_dir}. {constraints}: run validate\_pack.py and lint\_attachments.py on every packs/\*.md; fail on first error. {format}: report pass/fail per pack with file paths. {acceptance\_criteria}: all packs pass schema + direct-attach lint. {deadline}” | Validator and linter run; failures identify exact step and issue; all packs pass or errors are actionable. |
| QA / validation | Enforce “no bundle dependency” in direct-ticket mode | Files: packs/\*.md | “Make sure it doesn’t use a tickets bundle.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: out\_dir={out\_dir}. {constraints}: direct-ticket mode only; reject any `_tickets_bundle` references. {format}: lint output listing any violating steps/lines. {acceptance\_criteria}: zero `_tickets_bundle` matches across packs. {deadline}” | Lint reports zero bundle references; any violation is pinpointed to step and offending line. |
| Release / ops | Regenerate packs deterministically on a schedule (same grouping) | Files: `.tickets/**`; Constraint: stable outputs | “Rebuild the packs after ticket changes.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets out\_dir={out\_dir}. {constraints}: deterministic ordering only; no mtime/size-based heuristics; stable tie-breaks. {format}: regenerate packs + update manifest; keep group slugs stable except when splits change. {acceptance\_criteria}: repeated runs with unchanged tickets produce identical pack content. {deadline}” | With unchanged inputs, outputs are byte-stable (or explainable diffs only); grouping remains consistent. |
| Release / ops | Switch dedupe strategy to reduce attachment load | Files: `.tickets/**`; Constraint: minimize attachments | “Reduce duplicates in packs.” | “Use **$oraclepack-tickets-pack-grouped**. {context} {inputs}: ticket\_root=.tickets dedupe\_mode=merge dedupe\_body\_chars=2000. {constraints}: create \_ticket\_merges/cluster-XXXX.md and attach merged file only. {format}: manifest must show attached\_paths reflect merge outputs. {acceptance\_criteria}: merged clusters created; redundant tickets no longer attached directly. {deadline}” | Merge files exist; packs attach merge outputs; dedupe plan reflects merge mode; validation passes. |

1) Non-goals / anti-patterns (do NOT use this skill)

----------------------------------------------------

1. You want a single monolithic pack across all tickets (use a non-grouped tickets pack skill instead).

    oraclepack-tickets-pack-grouped…

2. You need non-deterministic or ML/service-based clustering (this skill explicitly requires deterministic rules and no external ML services).

    oraclepack-tickets-pack-grouped…

3. You want to rewrite ticket contents or author new tickets as the primary output (this skill’s contract is pack generation + validation).

    oraclepack-tickets-pack-grouped…

4. You can’t tolerate the fixed “Stage-1 / 20-step” schema (this skill enforces a strict 20-step contract and validation).

    oraclepack-tickets-pack-grouped…

5. You need interactive, ad-hoc investigation rather than a repeatable workflow package (skills are meant to standardize repeatable workflows via instructions/scripts/resources). [OpenAI Developers+1](https://developers.openai.com/codex/skills/)

---
