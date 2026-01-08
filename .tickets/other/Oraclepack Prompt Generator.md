* Title: Create oraclepack-style prompt/skill generator for tickets and .tickets
* Summary:

  * Need a reusable prompt (and/or “skill” template) that can generate an oraclepack-style prompt/skill specifically for “tickets” and/or “.tickets”.
  * Must support the existing placeholder-driven wrapper pattern (e.g., `{user-idea}`, `{project-in-question}`, `{PAIN-POINT}`, `{REFERENCE-FILE}`, `{CAPABILITY}`, `{TARGET-AGENT}`, `{OPTIMIZE-PROMPT}`), including optional fields and “infer from context” behavior as described.
  * Also need guidance on what to change in the current skill and what other viable integration options exist (within the constraints already used in prior work).
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt: “prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets”
* Global Constraints:

  * Optional inputs may be omitted; proceed by inferring from context or requesting missing info within the generated prompt template (“Not always provided…”).
  * “Pain point” is optional; proceed without it if absent.
  * `{REFERENCE-FILE}` may be provided as additional constraints/spec content.
* Global Environment:

  * Unknown
* Global Evidence:

  * Existing wrapper pattern + MCP prompt/tool/resource publication precedent captured in: `/mnt/data/MCP server implementation.md`

Split Plan:

* Coverage Map:

  * Original item: “We need a prompt that can create an oraclepack prompt/skill but for our tickets and/or .tickets.”

    * Assigned Ticket ID: T2
  * Original item: “What could we do to our current skill…”

    * Assigned Ticket ID: T3
  * Original item: “…and/or what else are our options for this request?”

    * Assigned Ticket ID: T4
  * Original item: Wrapper placeholders + optionality rules (“Not always provided…”, “Our pain point…”, `{REFERENCE-FILE}`, `{TARGET-AGENT}`, `{CAPABILITY}`, `{OPTIMIZE-PROMPT}`)

    * Assigned Ticket ID: T1
  * Original item: “optimized prompt that will get the {TARGET-AGENT} to find us a solution for adding capabilities…”

    * Assigned Ticket ID: T2
* Dependencies:

  * T2 depends on T1 because the prompt/skill generator must align to the placeholder schema + optionality rules.
  * T3 depends on T2 because “current skill” changes should incorporate the finalized ticket prompt/skill template.
  * T5 depends on T2 and T3 because examples/validation need the final template and integration approach.

````ticket T1
T1 Title: Define ticket/.tickets prompt input schema and placeholder mapping
Type: docs
Target Area: Ticket input model (tickets and/or .tickets) + wrapper placeholders
Summary:
  - Define the canonical set of inputs and placeholders required to generate an oraclepack-style ticket prompt/skill.
  - Preserve the existing wrapper’s rules around optional inputs and context inference.
  - Provide a clear mapping between “tickets/.tickets” fields (if any) and wrapper placeholders without inventing unspecified fields.
In Scope:
  - Enumerate required vs optional placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Document handling rules explicitly stated: optional fields, “infer from context”, and behavior when pain point is absent.
  - Clarify what “tickets” vs “.tickets” means in this system using “Unknown/Not provided” where details are missing.
Out of Scope:
  - Defining new ticket fields beyond what is provided.
  - Implementing tooling or code changes (covered elsewhere).
Current Behavior (Actual):
  - Placeholder set and optionality rules exist in the wrapper pattern, but ticket/.tickets-specific mapping is not defined.
Expected Behavior:
  - A documented, stable mapping that the ticket prompt/skill generator can follow.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not add new placeholders or required fields beyond what is already used in the wrapper.
  - Preserve optionality rules: proceed safely when PAIN-POINT or additional info is absent.
Evidence:
  - Reference wrapper placeholders and prompt-engineer wrapper structure as precedent. (/mnt/data/MCP server implementation.md) :contentReference[oaicite:1]{index=1}
Open Items / Unknowns:
  - Exact structure/format of “tickets” and “.tickets” (not provided).
  - Whether “.tickets” is a file extension, folder convention, or schema name (not provided).
Risks / Dependencies:
  - Risk of mismatch between ticket data shape and placeholder mapping if .tickets format is not standardized.
Acceptance Criteria:
  - A single written spec exists that lists:
    - All placeholders and which are optional.
    - Rules for missing fields (“infer from context” as described).
    - How ticket/.tickets inputs populate placeholders (or explicitly “Unknown” where not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Not always provided, inference from context…”
  - “Our pain point: {PAIN-POINT} … if not just continue…”
  - “```md {REFERENCE-FILE}.md”
````

```ticket T2
T2 Title: Author oraclepack-style prompt/skill template for ticket and .tickets generation
Type: enhancement
Target Area: Prompt/skill template content (oraclepack-style) for tickets/.tickets
Summary:
  - Create the actual prompt/skill template that produces an oraclepack-style prompt/skill when given a ticket or .tickets input.
  - The template must use the existing wrapper structure and placeholders, and must instruct the TARGET-AGENT to generate the desired capability for the project/tool in question.
  - Ensure the template explicitly supports optional inputs and reference-file injection as described.
In Scope:
  - Produce the “ticket prompt-engineer wrapper” template that mirrors the existing wrapper pattern but targets tickets/.tickets.
  - Include all placeholders: {user-idea}, {project-in-question}, {ADDITIONAL-INFORMATION}, {PAIN-POINT}, {REFERENCE-FILE}, {CAPABILITY}, {TARGET-AGENT}, {OPTIMIZE-PROMPT}.
  - Ensure the prompt text includes the “optimized prompt that will get the {TARGET-AGENT}…” requirement, scoped to tickets/.tickets.
Out of Scope:
  - Any new MCP server requirements, tools, or resource URI schemes not explicitly requested for tickets/.tickets.
Current Behavior (Actual):
  - There is no ticket/.tickets-specific oraclepack-style prompt/skill generator template.
Expected Behavior:
  - A single reusable prompt/skill template exists that can be filled with placeholders to drive a TARGET-AGENT to create ticket/.tickets capabilities.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must follow the wrapper’s optionality rules and placeholder usage.
  - Must treat {REFERENCE-FILE} content as “additional constraints/spec” when present.
Evidence:
  - Wrapper structure and placeholder set captured in existing reference prompt material. :contentReference[oaicite:2]{index=2}
Open Items / Unknowns:
  - Where this template will live (file path/naming) in the current repo/tooling (not provided).
Risks / Dependencies:
  - Depends on T1 for a stable placeholder-to-ticket mapping.
Acceptance Criteria:
  - Template includes:
    - All stated placeholders.
    - Explicit instruction to proceed when optional fields are missing.
    - A clearly stated “question to the prompt-engineer: {OPTIMIZE-PROMPT}” section.
  - Template is copy/paste ready and self-contained.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “create an oraclepack prompt/skill but for our tickets and/or .tickets”
  - “optimized prompt that will get the {TARGET-AGENT}… giving it {CAPABILITY}”
  - “Our question to the prompt-engineer: {OPTIMIZE-PROMPT}”
```

```ticket T3
T3 Title: Update current skill to support ticket/.tickets prompt-skill generation
Type: enhancement
Target Area: Existing “current skill” (location/name not provided)
Summary:
  - Identify changes required to the existing skill so it can generate or host the new tickets/.tickets oraclepack-style prompt/skill template.
  - Ensure the current skill can accept the ticket/.tickets inputs and populate the standardized placeholders.
  - Preserve existing behavior for non-ticket use cases (if any), since only ticket support is being added.
In Scope:
  - Incorporate the finalized template from T2 into the current skill workflow.
  - Add/adjust input handling so the current skill can be driven by “tickets and/or .tickets” as the source material.
  - Ensure optional inputs (pain point, additional information, reference file) remain optional in the workflow.
Out of Scope:
  - Designing a brand-new system if the current skill can be extended (unless extension is impossible; not provided).
Current Behavior (Actual):
  - Current skill does not explicitly support generating oraclepack-style prompts/skills for tickets/.tickets (per request).
Expected Behavior:
  - Current skill can produce the tickets/.tickets oraclepack-style prompt/skill using the same wrapper placeholder mechanism.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Do not remove or break existing skill behavior (implied by “current skill” extension request).
Evidence:
  - “What could we do to our current skill…”
Open Items / Unknowns:
  - Current skill name, file path, and execution context (not provided).
  - How tickets/.tickets are currently stored or passed into the system (not provided).
Risks / Dependencies:
  - Depends on T2 for the template content.
Acceptance Criteria:
  - Current skill supports a ticket/.tickets-driven flow that results in the T2 template with placeholders populated (or explicitly left blank when optional).
  - No regression to existing skill behaviors (validation method not provided; document what was exercised).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “prompt… for our tickets and/or .tickets”
```

```ticket T4
T4 Title: Document integration options for delivering the tickets/.tickets prompt-skill capability
Type: docs
Target Area: Delivery/integration approach (within existing patterns)
Summary:
  - Provide a concise options write-up for how to deliver and reuse the tickets/.tickets prompt/skill generator, aligned to the existing approach patterns already in use.
  - Focus on the two explicitly requested dimensions: changes to the current skill and “other options” for fulfilling the request.
  - Avoid committing to new systems; frame as documented options with constraints and unknowns.
In Scope:
  - Option 1: Extend current skill to include tickets/.tickets support (ties to T3).
  - Option 2: Provide a standalone tickets/.tickets prompt/skill template artifact that can be consumed independently (ties to T2).
  - List constraints/unknowns impacting option choice (e.g., unknown .tickets format, unknown current-skill location).
Out of Scope:
  - Implementing the chosen option (handled by T3 and/or T2).
Current Behavior (Actual):
  - No documented approach exists for how tickets/.tickets prompt/skill generation should be delivered.
Expected Behavior:
  - A short decision-ready document exists describing the options and what each requires, without adding new requirements.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Options must stay within what’s already requested (modify current skill and/or alternative ways to package the prompt/skill).
Evidence:
  - “what else are our options for this request?”
Open Items / Unknowns:
  - Whether the user prefers a single consolidated skill vs multiple dedicated skills (not provided).
Risks / Dependencies:
  - Depends on T1/T2 clarity to accurately describe what each option would deliver.
Acceptance Criteria:
  - Document lists at least:
    - “Modify current skill” option (summary, prerequisites, impact).
    - “Standalone template” option (summary, prerequisites, impact).
    - Explicit unknowns that block a final choice.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “What could we do to our current skill…”
  - “…what else are our options for this request?”
```

````ticket T5
T5 Title: Add examples and validation checks for ticket/.tickets prompt-skill generation
Type: tests
Target Area: Examples + validation of generated prompt/skill output
Summary:
  - Provide concrete example inputs (ticket and/or .tickets) and the expected generated prompt/skill output shape for validation.
  - Ensure examples exercise optional fields (missing PAIN-POINT, missing ADDITIONAL-INFORMATION, with/without REFERENCE-FILE).
  - Add lightweight validation criteria to confirm generated output preserves placeholders and wrapper structure.
In Scope:
  - Example cases covering:
    - Only {user-idea} + {project-in-question}
    - With {PAIN-POINT}
    - With {REFERENCE-FILE}
  - Validation checklist for generated output structure (placeholders present; optional fields handled).
Out of Scope:
  - End-to-end integration tests that require specific repo tooling not provided.
Current Behavior (Actual):
  - No examples/validation for tickets/.tickets prompt-skill generation are defined.
Expected Behavior:
  - Examples exist and can be used to validate that the template and current-skill integration behave as intended.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must preserve the wrapper structure and placeholders as-is.
Evidence:
  - Placeholder and wrapper expectations referenced in the existing wrapper pattern. :contentReference[oaicite:3]{index=3}
Open Items / Unknowns:
  - Exact acceptance mechanism for “validation checks” in the existing system (not provided).
Risks / Dependencies:
  - Depends on T2 (template) and T3 (integration) for meaningful expected outputs.
Acceptance Criteria:
  - At least 3 example inputs exist (covering optionality cases).
  - Each example includes an expected output outline that confirms:
    - Placeholders are present.
    - Optional fields can be omitted without breaking structure.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Not always provided…”
  - “Our pain point: {PAIN-POINT} … continue without needing it.”
  - “```md {REFERENCE-FILE}.md”
````
