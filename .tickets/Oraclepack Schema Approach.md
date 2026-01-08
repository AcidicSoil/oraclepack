Parent Ticket:

* Title: Adopt a schema-driven approach to prevent oraclepack run failures
* Summary:

  * Current runs fail because structure is inferred from Markdown heuristics (e.g., exactly one ```bash fence, sequential step headers, exactly 20 steps).
  * Proposal: generate a machine-validated **manifest** (JSON Schema) and **deterministically render** the Markdown pack; optionally add stricter linting for Markdown-only packs.
* Source:

  * Link/ID (if present) or “Not provided”
  * Original ticket excerpt (≤25 words) capturing the overall theme

    * “separate ‘machine-validated structure’ from ‘human-readable Markdown’ … generate only a JSON manifest … then a deterministic renderer produces the Markdown pack.”
* Global Constraints:

  * Keep existing oraclepack Markdown contract / backward-compatible (“keep the existing Markdown contract for oraclepack execution”).
  * Steps must be exactly 20; step IDs must be sequential 01..20.
* Global Environment:

  * Unknown
* Global Evidence:

  * Error text: “invalid pack structure: no bash code block found”.
  * Pack constraints referenced: “Exactly one ```bash fence”, “Exactly 20 steps”, “sequential step numbers”.

Split Plan:

* Coverage Map:

  * Original item: “separate ‘machine-validated structure’ from ‘human-readable Markdown.’”

    * Assigned Ticket ID: T1
  * Original item: “AI generates only a JSON manifest that must validate against a JSON Schema; renderer produces Markdown pack.”

    * Assigned Ticket ID: T1
  * Original item: “Prevent missing/multiple ```bash fences (root cause of ‘invalid pack structure: no bash code block found’).”

    * Assigned Ticket ID: T3
  * Original item: “Prevent non-sequential steps (Go validator requires sequential step numbers).”

    * Assigned Ticket ID: T1
  * Original item: “Prevent wrong step count (enforce exactly 20 in schema).”

    * Assigned Ticket ID: T1
  * Original item: “Minimal ‘Pack Manifest v1’ JSON Schema (Draft 2020-12) with schema_version/kind/out_dir/write_output/steps; step fields id/title/bash plus roi/impact/confidence/effort/horizon/category/reference.”

    * Assigned Ticket ID: T1
  * Original item: “Rendering rule (deterministic): one ```bash fence; prelude out_dir=…; optional --write-output; each step ‘# 01) …’ with body.”

    * Assigned Ticket ID: T2
  * Original item: “If Markdown-only: add explicit schema/lint mode (exactly one ```bash fence; exactly 20 steps; sequential 01..20; optional header tokens).”

    * Assigned Ticket ID: T3
  * Original item: “Stage-2 directory contract: exactly one file per prefix 01-*.md … 20-*.md.”

    * Assigned Ticket ID: T3
  * Original item: “Action pack lint (Stage 3): one ```bash fence; enforce 01..20 exact count.”

    * Assigned Ticket ID: T3
  * Original item: “CI checks: validate(manifest.json) → render(pack.md) → oraclepack validate pack.md → (optional) dry-run checks.”

    * Assigned Ticket ID: T4
* Dependencies:

  * T2 depends on T1 because the renderer needs the validated “Pack Manifest v1” structure as input.
  * T4 depends on T1 and T2 because CI runs “validate(manifest.json) → render(pack.md)”.
* Split Tickets:

```ticket T1
T# Title: Define and validate “Pack Manifest v1” schema (manifest-first)
Type: chore
Target Area: Pack authoring contract (manifest JSON + JSON Schema validation)
Summary:
- Introduce a manifest-first source of truth: the AI produces a JSON manifest that must validate against a JSON Schema.
- The schema enforces step count (exactly 20) and step IDs (01..20) to prevent structural runner failures.
- This separates machine-validated structure from the Markdown pack to reduce malformed packs.
In Scope:
- Define “Pack Manifest v1” JSON Schema (Draft 2020-12) with required fields: schema_version (const 1), kind (enum), out_dir, steps (min/max 20).
- Define step object constraints: required id/title/bash; id pattern for 01..20; optional roi/impact/confidence/effort/horizon/category/reference.
- Validate manifests against the schema before rendering/using them.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Runner infers structure from Markdown heuristics; malformed structure can cause run-time validation errors.
- Step count and sequential numbering can be violated if not enforced early.
Expected Behavior:
- A manifest that does not conform (wrong count, wrong IDs, missing required fields) is rejected by schema validation.
- Manifests accepted by validation always contain exactly 20 steps with valid IDs and required fields.
Reproduction Steps:
1) Provide a manifest with fewer than 20 steps.
2) Provide a manifest with a non-matching step id (e.g., "21" or "1").
3) Validate manifest and confirm it fails schema validation.
Requirements / Constraints:
- schema_version must be 1.
- steps must be exactly 20 (minItems=20, maxItems=20).
- step id must match 01..20 via pattern.
Evidence:
- “the AI generates only a JSON manifest that must validate against a JSON Schema” (ticket text)
- “Wrong step count (you can enforce exactly 20 in schema rather than ‘hoping’ the model did it).”
Open Items / Unknowns:
- Exact location/path conventions for storing manifest.json are not provided.
- How/where validation is invoked (CLI, CI, library) is not provided.
Risks / Dependencies:
- Risk: keeping backward compatibility requires rendering to the existing Markdown pack contract (handled in T2).
Acceptance Criteria:
- [ ] A JSON Schema exists for “Pack Manifest v1” with required fields and constraints as described.
- [ ] A manifest with != 20 steps fails validation.
- [ ] A manifest with an invalid step id fails validation.
- [ ] A manifest missing required fields (schema_version/kind/out_dir/steps) fails validation.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Manifest-first + JSON Schema (then render Markdown)”
- “minItems: 20, maxItems: 20”
- “id … pattern: ^(0[1-9]|1[0-9]|20)$”
```

````ticket T2
T# Title: Implement deterministic renderer from manifest → oraclepack Markdown pack
Type: enhancement
Target Area: Pack rendering (manifest → Markdown pack)
Summary:
- Add a deterministic rendering rule that converts a validated manifest into a Markdown pack that always satisfies oraclepack’s structural expectations.
- This prevents issues like missing/multiple bash fences and malformed step formatting by making Markdown a compiled artifact.
In Scope:
- Render exactly one fenced code block labeled `bash` in the entire document.
- Render prelude lines including: `out_dir="..."` and optional `--write-output` as described.
- Render each step with header `# NN) ...` and step body from `bash` content.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Markdown is the primary artifact; structure can be malformed by generation, causing downstream failures.
Expected Behavior:
- Renderer output always includes exactly one `bash` fence and emits all 20 steps in order.
- Pack contains the required prelude line(s) described in the ticket text.
Reproduction Steps:
1) Validate a manifest (per T1).
2) Render to Markdown.
3) Confirm output contains exactly one `bash` fence and step headers for 01..20 in sequence.
Requirements / Constraints:
- Output must be runner-ingestible per the described structural rules (single bash fence, 20 steps, sequential).
Evidence:
- “Rendering rule (deterministic) … emits exactly: one ```bash fence … prelude lines … then each step: # 01) … Step body = bash”
Open Items / Unknowns:
- Exact step title formatting beyond “# NN) …” is not provided.
- Whether additional header tokens (ROI=…) are required at render time is not provided.
Risks / Dependencies:
- Depends on T1 (renderer assumes manifest structure/constraints).
Acceptance Criteria:
- [ ] Given a valid manifest, renderer produces a Markdown pack with exactly one ```bash fenced block.
- [ ] Output contains 20 step headers numbered 01..20 in order.
- [ ] Output includes the `out_dir="..."` prelude line.
- [ ] Renderer can conditionally include the optional `--write-output` line when present in the manifest.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “From this manifest, your renderer emits exactly: One ```bash fence”
- “Prelude lines: out_dir="..." … optional --write-output”
- “Then each step: # 01) … Step body = bash”
````

````ticket T3
T# Title: Add stricter lint/validation for Markdown-only packs and Stage-2/Stage-3 outputs
Type: chore
Target Area: Pack linting/validation (Markdown packs + output directory contract)
Summary:
- If the project keeps Markdown-only as a supported input path, add an explicit lint/validation mode that enforces the same structural contract.
- Extend checks to Stage-2 output directory naming expectations and Stage-3 action pack constraints to reduce “runner infers structure” failures.
In Scope:
- Pack-level lint (Stage 1) enforcing:
  - Exactly one ```bash fence.
  - Exactly 20 steps.
  - Step IDs exactly 01..20 and sequential.
  - Optional enforcement of required header tokens (ROI= impact= confidence= … reference=) as described.
- Stage-2 directory contract lint:
  - Exactly one file per prefix 01-*.md … 20-*.md.
- Stage-3 action pack lint:
  - Exactly one ```bash fence.
  - Enforce 01..20 and exact count.
Out of Scope:
- Not provided
Current Behavior (Actual):
- Common failure mode noted: “invalid pack structure: no bash code block found.”
- Existing checks may be incomplete (“your current check only ensures ‘some’ step headers exist” per ticket text).
Expected Behavior:
- Markdown packs that violate the contract are rejected early with lint errors before execution.
- Stage-2 outputs and Stage-3 action packs are validated against exact-count and naming/structure rules.
Reproduction Steps:
1) Create a Markdown pack with no `bash` fenced block → lint should fail.
2) Create a Markdown pack with 19 steps or non-sequential IDs → lint should fail.
3) Create an output directory missing `07-*.md` or containing duplicates for a prefix → lint should fail.
Requirements / Constraints:
- Enforce: one ```bash fence, exactly 20 steps, sequential 01..20.
Evidence:
- “Missing / multiple ```bash fences (common root cause of ‘invalid pack structure: no bash code block found’).”
- “Add an explicit schema/lint mode … Exactly one ```bash fence … Exactly 20 steps … Step IDs exactly 01..20”
- “Stage-2 directory contract … Exactly one file per prefix 01-*.md … 20-*.md”
- “Action pack lint (Stage 3) … Enforce 01..20 and exact count”
Open Items / Unknowns:
- Exact current validator behaviors and what already exists vs missing are not provided.
Risks / Dependencies:
- Risk: enforcing optional header tokens could break existing packs if not already standardized.
Acceptance Criteria:
- [ ] Lint fails when no `bash` fence exists and surfaces a clear error.
- [ ] Lint fails when step count != 20.
- [ ] Lint fails when step IDs are not exactly 01..20 sequential.
- [ ] Stage-2 lint fails when any step output prefix 01..20 is missing or duplicated.
- [ ] Stage-3 lint fails when action pack does not have exactly one `bash` fence or correct 01..20 steps.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “invalid pack structure: no bash code block found”
- “Pack-level lint (Stage 1) … Exactly one ```bash fence … Exactly 20 steps”
- “Stage-2 directory contract … Exactly one file per prefix 01-*.md … 20-*.md”
````

```ticket T4
T# Title: Add CI validation pipeline for manifest-first workflow (validate → render → oraclepack validate → optional dry-run)
Type: chore
Target Area: CI checks / pipeline gating
Summary:
- Add CI checks that gate merges/runs on structural correctness by validating the manifest, rendering Markdown deterministically, and validating the rendered pack with oraclepack tooling.
- This formalizes the “Markdown is compiled artifact” approach and reduces runtime surprises.
In Scope:
- CI sequence as described:
  - validate(manifest.json)
  - render(pack.md)
  - oraclepack validate pack.md
  - optional dry-run checks
Out of Scope:
- Not provided
Current Behavior (Actual):
- Pack structural issues can slip into execution time if not validated earlier.
Expected Behavior:
- CI fails fast when manifest validation fails or rendered pack fails oraclepack validation.
Reproduction Steps:
1) Commit a manifest with 19 steps; CI should fail at validate(manifest.json).
2) Commit a manifest that renders an invalid pack (if possible); CI should fail at oraclepack validate.
Requirements / Constraints:
- CI must preserve existing oraclepack Markdown contract (rendered pack is what oraclepack consumes).
Evidence:
- “Add CI checks: validate(manifest.json) → render(pack.md) → oraclepack validate pack.md → (optional) dry-run checks”
Open Items / Unknowns:
- Where CI runs (provider/tooling) is not provided.
- Whether “dry-run checks” exist and what they check is not provided.
Risks / Dependencies:
- Depends on T1 and T2 to provide validate+render steps.
Acceptance Criteria:
- [ ] CI runs schema validation on manifest.json and fails on invalid manifests.
- [ ] CI renders pack.md deterministically from the manifest.
- [ ] CI runs oraclepack validation on pack.md and fails if invalid.
- [ ] Optional dry-run step is present if supported; otherwise omitted without breaking the sequence intent.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Add CI checks: validate(manifest.json) → render(pack.md) → oraclepack validate pack.md → (optional) dry-run checks”
- “Treat Markdown packs as a compiled artifact, not the source of truth.”
```
