<filetree>
Project Structure:
├── mcp-builder
│   ├── reference
│   │   ├── evaluation.md
│   │   ├── mcp_best_practices.md
│   │   ├── node_mcp_server.md
│   │   └── python_mcp_server.md
│   ├── scripts
│   │   ├── connections.py
│   │   ├── evaluation.py
│   │   ├── example_evaluation.xml
│   │   └── requirements.txt
│   └── SKILL.md
├── oraclepack-gold-pack
│   ├── references
│   │   ├── attachment-minimization.md
│   │   ├── inference-first-discovery.md
│   │   ├── oracle-pack-template.md
│   │   └── oracle-scratch-format.md
│   ├── scripts
│   │   ├── lint_attachments.py
│   │   └── validate_pack.py
│   └── SKILL.md
└── oraclepack-taskify
    ├── assets
    │   ├── action-pack-template.md
    │   ├── actions-json-schema.md
    │   └── prd-synthesis-prompt.md
    ├── references
    │   ├── determinism-and-safety.md
    │   ├── task-master-cli-cheatsheet.md
    │   └── workflow-overview.md
    ├── scripts
    │   ├── detect-oracle-outputs.sh
    │   └── validate-action-pack.sh
    └── SKILL.md

</filetree>

<source_code>
mcp-builder/SKILL.md
```
---
name: mcp-builder
description: Guide for creating high-quality MCP (Model Context Protocol) servers that enable LLMs to interact with external services through well-designed tools. Use when building MCP servers to integrate external APIs or services, whether in Python (FastMCP) or Node/TypeScript (MCP SDK).
license: Complete terms in LICENSE.txt
---

# MCP Server Development Guide

## Overview

Create MCP (Model Context Protocol) servers that enable LLMs to interact with external services through well-designed tools. The quality of an MCP server is measured by how well it enables LLMs to accomplish real-world tasks.

---

# Process

## 🚀 High-Level Workflow

Creating a high-quality MCP server involves four main phases:

### Phase 1: Deep Research and Planning

#### 1.1 Understand Modern MCP Design

**API Coverage vs. Workflow Tools:**
Balance comprehensive API endpoint coverage with specialized workflow tools. Workflow tools can be more convenient for specific tasks, while comprehensive coverage gives agents flexibility to compose operations. Performance varies by client—some clients benefit from code execution that combines basic tools, while others work better with higher-level workflows. When uncertain, prioritize comprehensive API coverage.

**Tool Naming and Discoverability:**
Clear, descriptive tool names help agents find the right tools quickly. Use consistent prefixes (e.g., `github_create_issue`, `github_list_repos`) and action-oriented naming.

**Context Management:**
Agents benefit from concise tool descriptions and the ability to filter/paginate results. Design tools that return focused, relevant data. Some clients support code execution which can help agents filter and process data efficiently.

**Actionable Error Messages:**
Error messages should guide agents toward solutions with specific suggestions and next steps.

#### 1.2 Study MCP Protocol Documentation

**Navigate the MCP specification:**

Start with the sitemap to find relevant pages: `https://modelcontextprotocol.io/sitemap.xml`

Then fetch specific pages with `.md` suffix for markdown format (e.g., `https://modelcontextprotocol.io/specification/draft.md`).

Key pages to review:
- Specification overview and architecture
- Transport mechanisms (streamable HTTP, stdio)
- Tool, resource, and prompt definitions

#### 1.3 Study Framework Documentation

**Recommended stack:**
- **Language**: TypeScript (high-quality SDK support and good compatibility in many execution environments e.g. MCPB. Plus AI models are good at generating TypeScript code, benefiting from its broad usage, static typing and good linting tools)
- **Transport**: Streamable HTTP for remote servers, using stateless JSON (simpler to scale and maintain, as opposed to stateful sessions and streaming responses). stdio for local servers.

**Load framework documentation:**

- **MCP Best Practices**: [📋 View Best Practices](./reference/mcp_best_practices.md) - Core guidelines

**For TypeScript (recommended):**
- **TypeScript SDK**: Use WebFetch to load `https://raw.githubusercontent.com/modelcontextprotocol/typescript-sdk/main/README.md`
- [⚡ TypeScript Guide](./reference/node_mcp_server.md) - TypeScript patterns and examples

**For Python:**
- **Python SDK**: Use WebFetch to load `https://raw.githubusercontent.com/modelcontextprotocol/python-sdk/main/README.md`
- [🐍 Python Guide](./reference/python_mcp_server.md) - Python patterns and examples

#### 1.4 Plan Your Implementation

**Understand the API:**
Review the service's API documentation to identify key endpoints, authentication requirements, and data models. Use web search and WebFetch as needed.

**Tool Selection:**
Prioritize comprehensive API coverage. List endpoints to implement, starting with the most common operations.

---

### Phase 2: Implementation

#### 2.1 Set Up Project Structure

See language-specific guides for project setup:
- [⚡ TypeScript Guide](./reference/node_mcp_server.md) - Project structure, package.json, tsconfig.json
- [🐍 Python Guide](./reference/python_mcp_server.md) - Module organization, dependencies

#### 2.2 Implement Core Infrastructure

Create shared utilities:
- API client with authentication
- Error handling helpers
- Response formatting (JSON/Markdown)
- Pagination support

#### 2.3 Implement Tools

For each tool:

**Input Schema:**
- Use Zod (TypeScript) or Pydantic (Python)
- Include constraints and clear descriptions
- Add examples in field descriptions

**Output Schema:**
- Define `outputSchema` where possible for structured data
- Use `structuredContent` in tool responses (TypeScript SDK feature)
- Helps clients understand and process tool outputs

**Tool Description:**
- Concise summary of functionality
- Parameter descriptions
- Return type schema

**Implementation:**
- Async/await for I/O operations
- Proper error handling with actionable messages
- Support pagination where applicable
- Return both text content and structured data when using modern SDKs

**Annotations:**
- `readOnlyHint`: true/false
- `destructiveHint`: true/false
- `idempotentHint`: true/false
- `openWorldHint`: true/false

---

### Phase 3: Review and Test

#### 3.1 Code Quality

Review for:
- No duplicated code (DRY principle)
- Consistent error handling
- Full type coverage
- Clear tool descriptions

#### 3.2 Build and Test

**TypeScript:**
- Run `npm run build` to verify compilation
- Test with MCP Inspector: `npx @modelcontextprotocol/inspector`

**Python:**
- Verify syntax: `python -m py_compile your_server.py`
- Test with MCP Inspector

See language-specific guides for detailed testing approaches and quality checklists.

---

### Phase 4: Create Evaluations

After implementing your MCP server, create comprehensive evaluations to test its effectiveness.

**Load [✅ Evaluation Guide](./reference/evaluation.md) for complete evaluation guidelines.**

#### 4.1 Understand Evaluation Purpose

Use evaluations to test whether LLMs can effectively use your MCP server to answer realistic, complex questions.

#### 4.2 Create 10 Evaluation Questions

To create effective evaluations, follow the process outlined in the evaluation guide:

1. **Tool Inspection**: List available tools and understand their capabilities
2. **Content Exploration**: Use READ-ONLY operations to explore available data
3. **Question Generation**: Create 10 complex, realistic questions
4. **Answer Verification**: Solve each question yourself to verify answers

#### 4.3 Evaluation Requirements

Ensure each question is:
- **Independent**: Not dependent on other questions
- **Read-only**: Only non-destructive operations required
- **Complex**: Requiring multiple tool calls and deep exploration
- **Realistic**: Based on real use cases humans would care about
- **Verifiable**: Single, clear answer that can be verified by string comparison
- **Stable**: Answer won't change over time

#### 4.4 Output Format

Create an XML file with this structure:

```xml
<evaluation>
  <qa_pair>
    <question>Find discussions about AI model launches with animal codenames. One model needed a specific safety designation that uses the format ASL-X. What number X was being determined for the model named after a spotted wild cat?</question>
    <answer>3</answer>
  </qa_pair>
<!-- More qa_pairs... -->
</evaluation>
```

---

# Reference Files

## 📚 Documentation Library

Load these resources as needed during development:

### Core MCP Documentation (Load First)
- **MCP Protocol**: Start with sitemap at `https://modelcontextprotocol.io/sitemap.xml`, then fetch specific pages with `.md` suffix
- [📋 MCP Best Practices](./reference/mcp_best_practices.md) - Universal MCP guidelines including:
  - Server and tool naming conventions
  - Response format guidelines (JSON vs Markdown)
  - Pagination best practices
  - Transport selection (streamable HTTP vs stdio)
  - Security and error handling standards

### SDK Documentation (Load During Phase 1/2)
- **Python SDK**: Fetch from `https://raw.githubusercontent.com/modelcontextprotocol/python-sdk/main/README.md`
- **TypeScript SDK**: Fetch from `https://raw.githubusercontent.com/modelcontextprotocol/typescript-sdk/main/README.md`

### Language-Specific Implementation Guides (Load During Phase 2)
- [🐍 Python Implementation Guide](./reference/python_mcp_server.md) - Complete Python/FastMCP guide with:
  - Server initialization patterns
  - Pydantic model examples
  - Tool registration with `@mcp.tool`
  - Complete working examples
  - Quality checklist

- [⚡ TypeScript Implementation Guide](./reference/node_mcp_server.md) - Complete TypeScript guide with:
  - Project structure
  - Zod schema patterns
  - Tool registration with `server.registerTool`
  - Complete working examples
  - Quality checklist

### Evaluation Guide (Load During Phase 4)
- [✅ Evaluation Guide](./reference/evaluation.md) - Complete evaluation creation guide with:
  - Question creation guidelines
  - Answer verification strategies
  - XML format specifications
  - Example questions and answers
  - Running an evaluation with the provided scripts
```

oraclepack-gold-pack/SKILL.md
```
---
name: oraclepack-gold-pack
description: Generate a single canonical Stage-1 oraclepack question pack as Markdown:exactly one ```bash fence containing exactly 20 steps (01..20) with strict ROI header tokens, per-step --write-output, fixed categories + coverage check. Use when you need a gold, runner-ingestible pack template (Stage 1 only; not Stage 3 taskify).
metadata:
  short-description: Gold Stage-1 oraclepack pack generator + validators
---

# Oraclepack Gold Pack (Stage 1)

This skill produces the **canonical Stage-1** oraclepack question pack (20 Oracle CLI calls). It is intentionally strict to prevent schema drift.

**Non-negotiable contract (pack output):**
- Exactly **one** fenced code block: starts with exactly ` ```bash` on its own line, and ends with exactly ` ````
- No other fenced code blocks anywhere in the pack.
- Exactly **20** steps, numbered **01..20** in order.
- Every step has a header line matching:
  - `# NN) ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
- Every step includes `--write-output "<out_dir>/<nn>-<slug>.md"`.
- Categories are fixed to this exact set (no additions/renames):
  - `contracts/interfaces`
  - `invariants`
  - `caching/state`
  - `background jobs`
  - `observability`
  - `permissions`
  - `migrations`
  - `UX flows`
  - `failure modes`
  - `feature flags`
- Pack ends with a **Coverage check** section listing all 10 categories as `OK` or `Missing(<step ids>)`.

The pack template is the contract. The scratch doc is **not** a pack format.

References:
- Contract template: `references/oracle-pack-template.md`
- Repo discovery: `references/inference-first-discovery.md`
- Attachment rules: `references/attachment-minimization.md`
- Scratch playbook (not pack): `references/oracle-scratch-format.md`

## Quick start

1) Generate a pack file (intended path):
- `docs/oracle-pack-YYYY-MM-DD.md`

1) Validate it (recommended before running oraclepack):
- `python3 scripts/validate_pack.py docs/oracle-pack-YYYY-MM-DD.md`

1) Optional attachment lint:
- `python3 scripts/lint_attachments.py docs/oracle-pack-YYYY-MM-DD.md`

## Inputs (do not ask follow-ups)

Interpret the user’s trailing text as conceptual `{{args}}`. Extract:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD`)
- `oracle_cmd` (default `oracle`)
- `oracle_flags` (default `--files-report`)
- `engine` (`api|browser`; optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(engine flag unknown)`)
- `model` (optional; if provided, append to flags *only if your oracle CLI supports it*, else omit and record `TODO(model flag unknown)`)
- `extra_files` (default empty; if provided, append **literally** to every command)

If any value is missing: use defaults and proceed.

## Workflow (deterministic)

### 1) Read the contract template first
Open `references/oracle-pack-template.md` and treat it as the **single source of truth** for formatting.

### 2) Repo discovery (inference-first)
Follow `references/inference-first-discovery.md`:
- Read a small set of “anchors” first.
- Infer what’s present in the repo.
- Only then choose the best 1–2 attachments per step.

### 3) Plan the 20 probes (2 per category)
Use the fixed categories and produce **exactly 2 steps per category** (20 total). Keep each step’s prompt focused and non-overlapping.

For each step:
- Pick a **reference anchor** (`reference=` token): `{path}:{symbol}` OR `{path}` OR `Unknown`.
- Pick ≤2 attachments (or fewer) using `references/attachment-minimization.md`.
- Ensure the prompt asks for:
  - Direct answer (bullets)
  - Risks/unknowns
  - Next smallest concrete experiment
  - Missing artifact patterns to request if evidence is insufficient

### 4) Emit the pack (single file)
Produce exactly one Markdown document with:
- Title + parsed args section (plain markdown; no code fences)
- Exactly one ` ```bash` fence containing the 20 steps
- A Coverage check section after the fence

### 5) Validate and correct drift
Run:
- `python3 scripts/validate_pack.py <pack.md>`
If it fails, fix the pack until it passes.

Optionally run:
- `python3 scripts/lint_attachments.py <pack.md>`
If it fails, reduce attachments to ≤2 per step (before any literal `extra_files`).

## Output contract

When invoked, you produce:
- One runner-ingestible Markdown pack (intended filename: `docs/oracle-pack-YYYY-MM-DD.md`)

You do **not**:
- Run oraclepack
- Generate Stage-3 “action packs” (that is `oraclepack-taskify`)

## Failure modes (do not guess)

- Missing repo evidence → set `reference=Unknown`, attach fewer files, and explicitly request missing file/path patterns in the prompt.
- Uncertain CLI flag support (`engine`, `model`) → omit flags and write `TODO(engine/model flag unknown)` in the pack’s parsed args notes (do not invent flags).
- Any schema drift → fix until `scripts/validate_pack.py` passes.

## Invocation examples

1) “Generate a gold oraclepack Stage-1 pack for this repo. out_dir=docs/oracle-questions-2026-01-06”
2) “Make the strict 20-step pack for codebase_name=AcmeAPI constraints=‘no DB changes’”
3) “Create the canonical pack; engine=browser model=gpt-5.2-pro (if supported)”
4) “Produce the gold pack; add extra_files='-f docs/ARCHITECTURE.md -f docs/API.md'”
5) “Regenerate this pack but fix headers and coverage check so it validates.”
```

oraclepack-taskify/SKILL.md
```
---
name: oraclepack-taskify
description: Generate a Stage-3 Action Pack from oraclepack output (oracle-out 01–20 .md answers) that synthesizes a canonical actions plan + Task Master PRD, runs Task Master to create/expand tasks, and by default starts a guarded autopilot to begin implementation.
metadata:
  short-description: Stage 3:answers → tasks → pipelines/autopilot
---

# oraclepack-taskify

## Use when

Use this skill only after **oraclepack Stage 2** has produced **20 answer files** under an output directory (default `oracle-out/`), and the user wants Stage 3 work products such as:

- “Stage 3” / “taskify” / “actionize” / “turn oracle outputs into tasks”
- “Task Master follow-up” / “PRD from oracle-out” / “implementation plan”
- “Start work automatically from oracle-out” / “autopilot top tasks”

## Deliverable

When invoked, produce exactly one primary deliverable:

- A single Markdown doc at `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD.md`)

The generated Action Pack MUST be oraclepack-ingestible and MUST obey:

- Exactly **one** fenced code block labeled `bash` in the entire document.
- No other code fences anywhere in the document.
- Inside the bash fence, include sequential step headers exactly like: `# NN)` where `NN` is `01, 02, 03...`
- Each step is self-contained and must **not rely on shell variables created in previous steps**.
- Each step writes artifacts to explicit, deterministic paths.

## Skill interface (args)

Parse user trailing text as whitespace-separated `KEY=value` tokens (last-one-wins). Do not ask follow-ups.

Supported args (all optional):

- `out_dir` (default `oracle-out`)
- `pack_path` (default `docs/oracle-actions-pack-YYYY-MM-DD.md`)
- `actions_json` (default `<out_dir>/_actions.json`)
- `actions_md` (default `<out_dir>/_actions.md`)
- `prd_path` (default `.taskmaster/docs/oracle-actions-prd.md`)
- `tag` (default `oraclepack`)
- `mode` (default `autopilot`; allowed `backlog|pipelines|autopilot`)
- `top_n` (default `10`; clamp to `1..20`)
- `oracle_cmd` (default `oracle`; allow a multi-word command like `npx -y @steipete/oracle` only if the user requests it)
- `task_master_cmd` (default `task-master`)
- `tm_cmd` (default `tm`; used only in autopilot mode)
- `extra_files` (default empty; if provided, treat as a comma-separated list of file paths; attach them wherever the synthesis step accepts file inputs)

If any referenced file/path does not exist, the skill still outputs the Action Pack, but includes an early step that prints a clear error and exits non-zero.

## Workflow (what to do when invoked)

1) Parse args from `KEY=value` tokens (no follow-ups; last-one-wins; unknown keys ignored).

2) Resolve defaults deterministically:
   - Compute `YYYY-MM-DD` using the local date at generation time.
   - Apply defaults and clamp `top_n` to `1..20`.
   - Default `mode=autopilot` unless explicitly overridden.
   - Derive:
     - `actions_json = <out_dir>/_actions.json` unless user overrides
     - `actions_md = <out_dir>/_actions.md` unless user overrides

3) Render the Action Pack:
   - Start from `assets/action-pack-template.md`.
   - Substitute placeholders:
     - `{{pack_date}}`, `{{out_dir}}`, `{{pack_path}}`, `{{actions_json}}`, `{{actions_md}}`, `{{prd_path}}`, `{{tag}}`, `{{mode}}`, `{{top_n}}`, `{{oracle_cmd}}`, `{{task_master_cmd}}`, `{{tm_cmd}}`
   - Expand `extra_files` into literal bash lines of the form:
     - `oracle_file_flags+=( -f "<path>" )`
     - and insert them only where the template indicates “Extra attachments”.

4) Ensure Action Pack contract:
   - Exactly one `bash` code fence in the document.
   - No other code fences.
   - Step headers `# 01)`.. are sequential and stable.
   - Each step re-declares its variables and does not depend on variables from earlier steps.

5) Write the final pack to `pack_path` (create parent dirs).

## Modes

### mode=backlog

Action Pack should:
- Synthesize canonical `_actions.json` + `_actions.md`
- Write PRD to `prd_path`
- Run:
  - `task-master parse-prd <prd_path>` (attempt with tag scoping if supported)
  - `task-master analyze-complexity --output <out_dir>/tm-complexity.json`
  - `task-master expand --all`

### mode=pipelines

Do everything in `backlog`, then:
- Generate deterministic pipelines from tasks.json
- Write: `docs/oracle-actions-pipelines.md`

### mode=autopilot (default)

Do everything in `backlog`, then:
- Enforce branch safety and tests-first guardrails
- Start a guarded autopilot entrypoint (expects `tm`-style autopilot tooling)
- Never commit to the default branch (do not run `git commit` on main/master; create a work branch first)

Important: if the environment does not provide a compatible `tm autopilot`, Step 08 will fail fast with a clear error. To avoid that, run Stage 3 with `mode=backlog` or `mode=pipelines`.

## Determinism + safety rules

- No interactive prompts in the generated pack.
- Stable ordering: select exactly the 20 outputs by filename prefix ordering `01`..`20`.
- Fail fast when required tools are missing:
  - `command -v <task_master_cmd first word>`
  - `command -v <oracle_cmd first word>`
  - `command -v <tm_cmd first word>` only in autopilot mode
- Always create directories before writing files.
- Avoid destructive commands:
  - Do not delete files.
  - Do not force-push.
  - Do not commit to main/master (autopilot mode creates a new branch).
- If multiple outputs exist for a prefix (e.g., `01-*.md` expands to more than one file), exit non-zero with an explicit error listing the matches.

## Failure modes (handle without questions)

- Missing `out_dir` → pack Step 01 exits non-zero with a clear message.
- Missing any of `01-*.md`..`20-*.md` → Step 01 exits non-zero (and prints which one is missing).
- Missing required tools → Step 01 exits non-zero (and prints which command is missing).
- Unknown `mode` → treat as `autopilot` (clamp at generation time) and render `mode=autopilot` in the pack.

## Resources

- `assets/action-pack-template.md`: base template for the Action Pack deliverable.
- `assets/actions-json-schema.md`: canonical schema spec for `_actions.json` normalization.
- `assets/prd-synthesis-prompt.md`: exact prompt text embedded into the Action Pack for synthesis.
- `references/*`: Stage 3 overview and guardrails for future maintenance.
- `scripts/*`: optional helpers (standalone); may also be embedded verbatim into packs if desired.

## Invocation examples

- `$oraclepack-taskify out_dir=oracle-out` (defaults to mode=autopilot)
- `$oraclepack-taskify out_dir=oracle-out mode=backlog`
- `$oraclepack-taskify out_dir=oracle-out mode=pipelines tag=oraclepack-top top_n=10`
- `$oraclepack-taskify out_dir=oracle-out mode=autopilot tag=oraclepack-top pack_path=docs/oracle-actions-pack-2026-01-05.md`
- `$oraclepack-taskify out_dir=oracle-out extra_files=README.md,package.json`
```

mcp-builder/reference/evaluation.md
```
# MCP Server Evaluation Guide

## Overview

This document provides guidance on creating comprehensive evaluations for MCP servers. Evaluations test whether LLMs can effectively use your MCP server to answer realistic, complex questions using only the tools provided.

---

## Quick Reference

### Evaluation Requirements

- Create 10 human-readable questions
- Questions must be READ-ONLY, INDEPENDENT, NON-DESTRUCTIVE
- Each question requires multiple tool calls (potentially dozens)
- Answers must be single, verifiable values
- Answers must be STABLE (won't change over time)

### Output Format

```xml
<evaluation>
   <qa_pair>
      <question>Your question here</question>
      <answer>Single verifiable answer</answer>
   </qa_pair>
</evaluation>
```

---

## Purpose of Evaluations

The measure of quality of an MCP server is NOT how well or comprehensively the server implements tools, but how well these implementations (input/output schemas, docstrings/descriptions, functionality) enable LLMs with no other context and access ONLY to the MCP servers to answer realistic and difficult questions.

## Evaluation Overview

Create 10 human-readable questions requiring ONLY READ-ONLY, INDEPENDENT, NON-DESTRUCTIVE, and IDEMPOTENT operations to answer. Each question should be:

- Realistic
- Clear and concise
- Unambiguous
- Complex, requiring potentially dozens of tool calls or steps
- Answerable with a single, verifiable value that you identify in advance

## Question Guidelines

### Core Requirements

1. **Questions MUST be independent**
   - Each question should NOT depend on the answer to any other question
   - Should not assume prior write operations from processing another question

2. **Questions MUST require ONLY NON-DESTRUCTIVE AND IDEMPOTENT tool use**
   - Should not instruct or require modifying state to arrive at the correct answer

3. **Questions must be REALISTIC, CLEAR, CONCISE, and COMPLEX**
   - Must require another LLM to use multiple (potentially dozens of) tools or steps to answer

### Complexity and Depth

4. **Questions must require deep exploration**
   - Consider multi-hop questions requiring multiple sub-questions and sequential tool calls
   - Each step should benefit from information found in previous questions

5. **Questions may require extensive paging**
   - May need paging through multiple pages of results
   - May require querying old data (1-2 years out-of-date) to find niche information
   - The questions must be DIFFICULT

6. **Questions must require deep understanding**
   - Rather than surface-level knowledge
   - May pose complex ideas as True/False questions requiring evidence
   - May use multiple-choice format where LLM must search different hypotheses

7. **Questions must not be solvable with straightforward keyword search**
   - Do not include specific keywords from the target content
   - Use synonyms, related concepts, or paraphrases
   - Require multiple searches, analyzing multiple related items, extracting context, then deriving the answer

### Tool Testing

8. **Questions should stress-test tool return values**
   - May elicit tools returning large JSON objects or lists, overwhelming the LLM
   - Should require understanding multiple modalities of data:
     - IDs and names
     - Timestamps and datetimes (months, days, years, seconds)
     - File IDs, names, extensions, and mimetypes
     - URLs, GIDs, etc.
   - Should probe the tool's ability to return all useful forms of data

9. **Questions should MOSTLY reflect real human use cases**
   - The kinds of information retrieval tasks that HUMANS assisted by an LLM would care about

10. **Questions may require dozens of tool calls**
    - This challenges LLMs with limited context
    - Encourages MCP server tools to reduce information returned

11. **Include ambiguous questions**
    - May be ambiguous OR require difficult decisions on which tools to call
    - Force the LLM to potentially make mistakes or misinterpret
    - Ensure that despite AMBIGUITY, there is STILL A SINGLE VERIFIABLE ANSWER

### Stability

12. **Questions must be designed so the answer DOES NOT CHANGE**
    - Do not ask questions that rely on "current state" which is dynamic
    - For example, do not count:
      - Number of reactions to a post
      - Number of replies to a thread
      - Number of members in a channel

13. **DO NOT let the MCP server RESTRICT the kinds of questions you create**
    - Create challenging and complex questions
    - Some may not be solvable with the available MCP server tools
    - Questions may require specific output formats (datetime vs. epoch time, JSON vs. MARKDOWN)
    - Questions may require dozens of tool calls to complete

## Answer Guidelines

### Verification

1. **Answers must be VERIFIABLE via direct string comparison**
   - If the answer can be re-written in many formats, clearly specify the output format in the QUESTION
   - Examples: "Use YYYY/MM/DD.", "Respond True or False.", "Answer A, B, C, or D and nothing else."
   - Answer should be a single VERIFIABLE value such as:
     - User ID, user name, display name, first name, last name
     - Channel ID, channel name
     - Message ID, string
     - URL, title
     - Numerical quantity
     - Timestamp, datetime
     - Boolean (for True/False questions)
     - Email address, phone number
     - File ID, file name, file extension
     - Multiple choice answer
   - Answers must not require special formatting or complex, structured output
   - Answer will be verified using DIRECT STRING COMPARISON

### Readability

2. **Answers should generally prefer HUMAN-READABLE formats**
   - Examples: names, first name, last name, datetime, file name, message string, URL, yes/no, true/false, a/b/c/d
   - Rather than opaque IDs (though IDs are acceptable)
   - The VAST MAJORITY of answers should be human-readable

### Stability

3. **Answers must be STABLE/STATIONARY**
   - Look at old content (e.g., conversations that have ended, projects that have launched, questions answered)
   - Create QUESTIONS based on "closed" concepts that will always return the same answer
   - Questions may ask to consider a fixed time window to insulate from non-stationary answers
   - Rely on context UNLIKELY to change
   - Example: if finding a paper name, be SPECIFIC enough so answer is not confused with papers published later

4. **Answers must be CLEAR and UNAMBIGUOUS**
   - Questions must be designed so there is a single, clear answer
   - Answer can be derived from using the MCP server tools

### Diversity

5. **Answers must be DIVERSE**
   - Answer should be a single VERIFIABLE value in diverse modalities and formats
   - User concept: user ID, user name, display name, first name, last name, email address, phone number
   - Channel concept: channel ID, channel name, channel topic
   - Message concept: message ID, message string, timestamp, month, day, year

6. **Answers must NOT be complex structures**
   - Not a list of values
   - Not a complex object
   - Not a list of IDs or strings
   - Not natural language text
   - UNLESS the answer can be straightforwardly verified using DIRECT STRING COMPARISON
   - And can be realistically reproduced
   - It should be unlikely that an LLM would return the same list in any other order or format

## Evaluation Process

### Step 1: Documentation Inspection

Read the documentation of the target API to understand:

- Available endpoints and functionality
- If ambiguity exists, fetch additional information from the web
- Parallelize this step AS MUCH AS POSSIBLE
- Ensure each subagent is ONLY examining documentation from the file system or on the web

### Step 2: Tool Inspection

List the tools available in the MCP server:

- Inspect the MCP server directly
- Understand input/output schemas, docstrings, and descriptions
- WITHOUT calling the tools themselves at this stage

### Step 3: Developing Understanding

Repeat steps 1 & 2 until you have a good understanding:

- Iterate multiple times
- Think about the kinds of tasks you want to create
- Refine your understanding
- At NO stage should you READ the code of the MCP server implementation itself
- Use your intuition and understanding to create reasonable, realistic, but VERY challenging tasks

### Step 4: Read-Only Content Inspection

After understanding the API and tools, USE the MCP server tools:

- Inspect content using READ-ONLY and NON-DESTRUCTIVE operations ONLY
- Goal: identify specific content (e.g., users, channels, messages, projects, tasks) for creating realistic questions
- Should NOT call any tools that modify state
- Will NOT read the code of the MCP server implementation itself
- Parallelize this step with individual sub-agents pursuing independent explorations
- Ensure each subagent is only performing READ-ONLY, NON-DESTRUCTIVE, and IDEMPOTENT operations
- BE CAREFUL: SOME TOOLS may return LOTS OF DATA which would cause you to run out of CONTEXT
- Make INCREMENTAL, SMALL, AND TARGETED tool calls for exploration
- In all tool call requests, use the `limit` parameter to limit results (<10)
- Use pagination

### Step 5: Task Generation

After inspecting the content, create 10 human-readable questions:

- An LLM should be able to answer these with the MCP server
- Follow all question and answer guidelines above

## Output Format

Each QA pair consists of a question and an answer. The output should be an XML file with this structure:

```xml
<evaluation>
   <qa_pair>
      <question>Find the project created in Q2 2024 with the highest number of completed tasks. What is the project name?</question>
      <answer>Website Redesign</answer>
   </qa_pair>
   <qa_pair>
      <question>Search for issues labeled as "bug" that were closed in March 2024. Which user closed the most issues? Provide their username.</question>
      <answer>sarah_dev</answer>
   </qa_pair>
   <qa_pair>
      <question>Look for pull requests that modified files in the /api directory and were merged between January 1 and January 31, 2024. How many different contributors worked on these PRs?</question>
      <answer>7</answer>
   </qa_pair>
   <qa_pair>
      <question>Find the repository with the most stars that was created before 2023. What is the repository name?</question>
      <answer>data-pipeline</answer>
   </qa_pair>
</evaluation>
```

## Evaluation Examples

### Good Questions

**Example 1: Multi-hop question requiring deep exploration (GitHub MCP)**

```xml
<qa_pair>
   <question>Find the repository that was archived in Q3 2023 and had previously been the most forked project in the organization. What was the primary programming language used in that repository?</question>
   <answer>Python</answer>
</qa_pair>
```

This question is good because:

- Requires multiple searches to find archived repositories
- Needs to identify which had the most forks before archival
- Requires examining repository details for the language
- Answer is a simple, verifiable value
- Based on historical (closed) data that won't change

**Example 2: Requires understanding context without keyword matching (Project Management MCP)**

```xml
<qa_pair>
   <question>Locate the initiative focused on improving customer onboarding that was completed in late 2023. The project lead created a retrospective document after completion. What was the lead's role title at that time?</question>
   <answer>Product Manager</answer>
</qa_pair>
```

This question is good because:

- Doesn't use specific project name ("initiative focused on improving customer onboarding")
- Requires finding completed projects from specific timeframe
- Needs to identify the project lead and their role
- Requires understanding context from retrospective documents
- Answer is human-readable and stable
- Based on completed work (won't change)

**Example 3: Complex aggregation requiring multiple steps (Issue Tracker MCP)**

```xml
<qa_pair>
   <question>Among all bugs reported in January 2024 that were marked as critical priority, which assignee resolved the highest percentage of their assigned bugs within 48 hours? Provide the assignee's username.</question>
   <answer>alex_eng</answer>
</qa_pair>
```

This question is good because:

- Requires filtering bugs by date, priority, and status
- Needs to group by assignee and calculate resolution rates
- Requires understanding timestamps to determine 48-hour windows
- Tests pagination (potentially many bugs to process)
- Answer is a single username
- Based on historical data from specific time period

**Example 4: Requires synthesis across multiple data types (CRM MCP)**

```xml
<qa_pair>
   <question>Find the account that upgraded from the Starter to Enterprise plan in Q4 2023 and had the highest annual contract value. What industry does this account operate in?</question>
   <answer>Healthcare</answer>
</qa_pair>
```

This question is good because:

- Requires understanding subscription tier changes
- Needs to identify upgrade events in specific timeframe
- Requires comparing contract values
- Must access account industry information
- Answer is simple and verifiable
- Based on completed historical transactions

### Poor Questions

**Example 1: Answer changes over time**

```xml
<qa_pair>
   <question>How many open issues are currently assigned to the engineering team?</question>
   <answer>47</answer>
</qa_pair>
```

This question is poor because:

- The answer will change as issues are created, closed, or reassigned
- Not based on stable/stationary data
- Relies on "current state" which is dynamic

**Example 2: Too easy with keyword search**

```xml
<qa_pair>
   <question>Find the pull request with title "Add authentication feature" and tell me who created it.</question>
   <answer>developer123</answer>
</qa_pair>
```

This question is poor because:

- Can be solved with a straightforward keyword search for exact title
[TRUNCATED]
```

mcp-builder/reference/mcp_best_practices.md
```
# MCP Server Best Practices

## Quick Reference

### Server Naming
- **Python**: `{service}_mcp` (e.g., `slack_mcp`)
- **Node/TypeScript**: `{service}-mcp-server` (e.g., `slack-mcp-server`)

### Tool Naming
- Use snake_case with service prefix
- Format: `{service}_{action}_{resource}`
- Example: `slack_send_message`, `github_create_issue`

### Response Formats
- Support both JSON and Markdown formats
- JSON for programmatic processing
- Markdown for human readability

### Pagination
- Always respect `limit` parameter
- Return `has_more`, `next_offset`, `total_count`
- Default to 20-50 items

### Transport
- **Streamable HTTP**: For remote servers, multi-client scenarios
- **stdio**: For local integrations, command-line tools
- Avoid SSE (deprecated in favor of streamable HTTP)

---

## Server Naming Conventions

Follow these standardized naming patterns:

**Python**: Use format `{service}_mcp` (lowercase with underscores)
- Examples: `slack_mcp`, `github_mcp`, `jira_mcp`

**Node/TypeScript**: Use format `{service}-mcp-server` (lowercase with hyphens)
- Examples: `slack-mcp-server`, `github-mcp-server`, `jira-mcp-server`

The name should be general, descriptive of the service being integrated, easy to infer from the task description, and without version numbers.

---

## Tool Naming and Design

### Tool Naming

1. **Use snake_case**: `search_users`, `create_project`, `get_channel_info`
2. **Include service prefix**: Anticipate that your MCP server may be used alongside other MCP servers
   - Use `slack_send_message` instead of just `send_message`
   - Use `github_create_issue` instead of just `create_issue`
3. **Be action-oriented**: Start with verbs (get, list, search, create, etc.)
4. **Be specific**: Avoid generic names that could conflict with other servers

### Tool Design

- Tool descriptions must narrowly and unambiguously describe functionality
- Descriptions must precisely match actual functionality
- Provide tool annotations (readOnlyHint, destructiveHint, idempotentHint, openWorldHint)
- Keep tool operations focused and atomic

---

## Response Formats

All tools that return data should support multiple formats:

### JSON Format (`response_format="json"`)
- Machine-readable structured data
- Include all available fields and metadata
- Consistent field names and types
- Use for programmatic processing

### Markdown Format (`response_format="markdown"`, typically default)
- Human-readable formatted text
- Use headers, lists, and formatting for clarity
- Convert timestamps to human-readable format
- Show display names with IDs in parentheses
- Omit verbose metadata

---

## Pagination

For tools that list resources:

- **Always respect the `limit` parameter**
- **Implement pagination**: Use `offset` or cursor-based pagination
- **Return pagination metadata**: Include `has_more`, `next_offset`/`next_cursor`, `total_count`
- **Never load all results into memory**: Especially important for large datasets
- **Default to reasonable limits**: 20-50 items is typical

Example pagination response:
```json
{
  "total": 150,
  "count": 20,
  "offset": 0,
  "items": [...],
  "has_more": true,
  "next_offset": 20
}
```

---

## Transport Options

### Streamable HTTP

**Best for**: Remote servers, web services, multi-client scenarios

**Characteristics**:
- Bidirectional communication over HTTP
- Supports multiple simultaneous clients
- Can be deployed as a web service
- Enables server-to-client notifications

**Use when**:
- Serving multiple clients simultaneously
- Deploying as a cloud service
- Integration with web applications

### stdio

**Best for**: Local integrations, command-line tools

**Characteristics**:
- Standard input/output stream communication
- Simple setup, no network configuration needed
- Runs as a subprocess of the client

**Use when**:
- Building tools for local development environments
- Integrating with desktop applications
- Single-user, single-session scenarios

**Note**: stdio servers should NOT log to stdout (use stderr for logging)

### Transport Selection

| Criterion | stdio | Streamable HTTP |
|-----------|-------|-----------------|
| **Deployment** | Local | Remote |
| **Clients** | Single | Multiple |
| **Complexity** | Low | Medium |
| **Real-time** | No | Yes |

---

## Security Best Practices

### Authentication and Authorization

**OAuth 2.1**:
- Use secure OAuth 2.1 with certificates from recognized authorities
- Validate access tokens before processing requests
- Only accept tokens specifically intended for your server

**API Keys**:
- Store API keys in environment variables, never in code
- Validate keys on server startup
- Provide clear error messages when authentication fails

### Input Validation

- Sanitize file paths to prevent directory traversal
- Validate URLs and external identifiers
- Check parameter sizes and ranges
- Prevent command injection in system calls
- Use schema validation (Pydantic/Zod) for all inputs

### Error Handling

- Don't expose internal errors to clients
- Log security-relevant errors server-side
- Provide helpful but not revealing error messages
- Clean up resources after errors

### DNS Rebinding Protection

For streamable HTTP servers running locally:
- Enable DNS rebinding protection
- Validate the `Origin` header on all incoming connections
- Bind to `127.0.0.1` rather than `0.0.0.0`

---

## Tool Annotations

Provide annotations to help clients understand tool behavior:

| Annotation | Type | Default | Description |
|-----------|------|---------|-------------|
| `readOnlyHint` | boolean | false | Tool does not modify its environment |
| `destructiveHint` | boolean | true | Tool may perform destructive updates |
| `idempotentHint` | boolean | false | Repeated calls with same args have no additional effect |
| `openWorldHint` | boolean | true | Tool interacts with external entities |

**Important**: Annotations are hints, not security guarantees. Clients should not make security-critical decisions based solely on annotations.

---

## Error Handling

- Use standard JSON-RPC error codes
- Report tool errors within result objects (not protocol-level errors)
- Provide helpful, specific error messages with suggested next steps
- Don't expose internal implementation details
- Clean up resources properly on errors

Example error handling:
```typescript
try {
  const result = performOperation();
  return { content: [{ type: "text", text: result }] };
} catch (error) {
  return {
    isError: true,
    content: [{
      type: "text",
      text: `Error: ${error.message}. Try using filter='active_only' to reduce results.`
    }]
  };
}
```

---

## Testing Requirements

Comprehensive testing should cover:

- **Functional testing**: Verify correct execution with valid/invalid inputs
- **Integration testing**: Test interaction with external systems
- **Security testing**: Validate auth, input sanitization, rate limiting
- **Performance testing**: Check behavior under load, timeouts
- **Error handling**: Ensure proper error reporting and cleanup

---

## Documentation Requirements

- Provide clear documentation of all tools and capabilities
- Include working examples (at least 3 per major feature)
- Document security considerations
- Specify required permissions and access levels
- Document rate limits and performance characteristics
```

mcp-builder/reference/node_mcp_server.md
```
# Node/TypeScript MCP Server Implementation Guide

## Overview

This document provides Node/TypeScript-specific best practices and examples for implementing MCP servers using the MCP TypeScript SDK. It covers project structure, server setup, tool registration patterns, input validation with Zod, error handling, and complete working examples.

---

## Quick Reference

### Key Imports
```typescript
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StreamableHTTPServerTransport } from "@modelcontextprotocol/sdk/server/streamableHttp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import express from "express";
import { z } from "zod";
```

### Server Initialization
```typescript
const server = new McpServer({
  name: "service-mcp-server",
  version: "1.0.0"
});
```

### Tool Registration Pattern
```typescript
server.registerTool(
  "tool_name",
  {
    title: "Tool Display Name",
    description: "What the tool does",
    inputSchema: { param: z.string() },
    outputSchema: { result: z.string() }
  },
  async ({ param }) => {
    const output = { result: `Processed: ${param}` };
    return {
      content: [{ type: "text", text: JSON.stringify(output) }],
      structuredContent: output // Modern pattern for structured data
    };
  }
);
```

---

## MCP TypeScript SDK

The official MCP TypeScript SDK provides:
- `McpServer` class for server initialization
- `registerTool` method for tool registration
- Zod schema integration for runtime input validation
- Type-safe tool handler implementations

**IMPORTANT - Use Modern APIs Only:**
- **DO use**: `server.registerTool()`, `server.registerResource()`, `server.registerPrompt()`
- **DO NOT use**: Old deprecated APIs such as `server.tool()`, `server.setRequestHandler(ListToolsRequestSchema, ...)`, or manual handler registration
- The `register*` methods provide better type safety, automatic schema handling, and are the recommended approach

See the MCP SDK documentation in the references for complete details.

## Server Naming Convention

Node/TypeScript MCP servers must follow this naming pattern:
- **Format**: `{service}-mcp-server` (lowercase with hyphens)
- **Examples**: `github-mcp-server`, `jira-mcp-server`, `stripe-mcp-server`

The name should be:
- General (not tied to specific features)
- Descriptive of the service/API being integrated
- Easy to infer from the task description
- Without version numbers or dates

## Project Structure

Create the following structure for Node/TypeScript MCP servers:

```
{service}-mcp-server/
├── package.json
├── tsconfig.json
├── README.md
├── src/
│   ├── index.ts          # Main entry point with McpServer initialization
│   ├── types.ts          # TypeScript type definitions and interfaces
│   ├── tools/            # Tool implementations (one file per domain)
│   ├── services/         # API clients and shared utilities
│   ├── schemas/          # Zod validation schemas
│   └── constants.ts      # Shared constants (API_URL, CHARACTER_LIMIT, etc.)
└── dist/                 # Built JavaScript files (entry point: dist/index.js)
```

## Tool Implementation

### Tool Naming

Use snake_case for tool names (e.g., "search_users", "create_project", "get_channel_info") with clear, action-oriented names.

**Avoid Naming Conflicts**: Include the service context to prevent overlaps:
- Use "slack_send_message" instead of just "send_message"
- Use "github_create_issue" instead of just "create_issue"
- Use "asana_list_tasks" instead of just "list_tasks"

### Tool Structure

Tools are registered using the `registerTool` method with the following requirements:
- Use Zod schemas for runtime input validation and type safety
- The `description` field must be explicitly provided - JSDoc comments are NOT automatically extracted
- Explicitly provide `title`, `description`, `inputSchema`, and `annotations`
- The `inputSchema` must be a Zod schema object (not a JSON schema)
- Type all parameters and return values explicitly

```typescript
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { z } from "zod";

const server = new McpServer({
  name: "example-mcp",
  version: "1.0.0"
});

// Zod schema for input validation
const UserSearchInputSchema = z.object({
  query: z.string()
    .min(2, "Query must be at least 2 characters")
    .max(200, "Query must not exceed 200 characters")
    .describe("Search string to match against names/emails"),
  limit: z.number()
    .int()
    .min(1)
    .max(100)
    .default(20)
    .describe("Maximum results to return"),
  offset: z.number()
    .int()
    .min(0)
    .default(0)
    .describe("Number of results to skip for pagination"),
  response_format: z.nativeEnum(ResponseFormat)
    .default(ResponseFormat.MARKDOWN)
    .describe("Output format: 'markdown' for human-readable or 'json' for machine-readable")
}).strict();

// Type definition from Zod schema
type UserSearchInput = z.infer<typeof UserSearchInputSchema>;

server.registerTool(
  "example_search_users",
  {
    title: "Search Example Users",
    description: `Search for users in the Example system by name, email, or team.

This tool searches across all user profiles in the Example platform, supporting partial matches and various search filters. It does NOT create or modify users, only searches existing ones.

Args:
  - query (string): Search string to match against names/emails
  - limit (number): Maximum results to return, between 1-100 (default: 20)
  - offset (number): Number of results to skip for pagination (default: 0)
  - response_format ('markdown' | 'json'): Output format (default: 'markdown')

Returns:
  For JSON format: Structured data with schema:
  {
    "total": number,           // Total number of matches found
    "count": number,           // Number of results in this response
    "offset": number,          // Current pagination offset
    "users": [
      {
        "id": string,          // User ID (e.g., "U123456789")
        "name": string,        // Full name (e.g., "John Doe")
        "email": string,       // Email address
        "team": string,        // Team name (optional)
        "active": boolean      // Whether user is active
      }
    ],
    "has_more": boolean,       // Whether more results are available
    "next_offset": number      // Offset for next page (if has_more is true)
  }

Examples:
  - Use when: "Find all marketing team members" -> params with query="team:marketing"
  - Use when: "Search for John's account" -> params with query="john"
  - Don't use when: You need to create a user (use example_create_user instead)

Error Handling:
  - Returns "Error: Rate limit exceeded" if too many requests (429 status)
  - Returns "No users found matching '<query>'" if search returns empty`,
    inputSchema: UserSearchInputSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: true
    }
  },
  async (params: UserSearchInput) => {
    try {
      // Input validation is handled by Zod schema
      // Make API request using validated parameters
      const data = await makeApiRequest<any>(
        "users/search",
        "GET",
        undefined,
        {
          q: params.query,
          limit: params.limit,
          offset: params.offset
        }
      );

      const users = data.users || [];
      const total = data.total || 0;

      if (!users.length) {
        return {
          content: [{
            type: "text",
            text: `No users found matching '${params.query}'`
          }]
        };
      }

      // Prepare structured output
      const output = {
        total,
        count: users.length,
        offset: params.offset,
        users: users.map((user: any) => ({
          id: user.id,
          name: user.name,
          email: user.email,
          ...(user.team ? { team: user.team } : {}),
          active: user.active ?? true
        })),
        has_more: total > params.offset + users.length,
        ...(total > params.offset + users.length ? {
          next_offset: params.offset + users.length
        } : {})
      };

      // Format text representation based on requested format
      let textContent: string;
      if (params.response_format === ResponseFormat.MARKDOWN) {
        const lines = [`# User Search Results: '${params.query}'`, "",
          `Found ${total} users (showing ${users.length})`, ""];
        for (const user of users) {
          lines.push(`## ${user.name} (${user.id})`);
          lines.push(`- **Email**: ${user.email}`);
          if (user.team) lines.push(`- **Team**: ${user.team}`);
          lines.push("");
        }
        textContent = lines.join("\n");
      } else {
        textContent = JSON.stringify(output, null, 2);
      }

      return {
        content: [{ type: "text", text: textContent }],
        structuredContent: output // Modern pattern for structured data
      };
    } catch (error) {
      return {
        content: [{
          type: "text",
          text: handleApiError(error)
        }]
      };
    }
  }
);
```

## Zod Schemas for Input Validation

Zod provides runtime type validation:

```typescript
import { z } from "zod";

// Basic schema with validation
const CreateUserSchema = z.object({
  name: z.string()
    .min(1, "Name is required")
    .max(100, "Name must not exceed 100 characters"),
  email: z.string()
    .email("Invalid email format"),
  age: z.number()
    .int("Age must be a whole number")
    .min(0, "Age cannot be negative")
    .max(150, "Age cannot be greater than 150")
}).strict();  // Use .strict() to forbid extra fields

// Enums
enum ResponseFormat {
  MARKDOWN = "markdown",
  JSON = "json"
}

const SearchSchema = z.object({
  response_format: z.nativeEnum(ResponseFormat)
    .default(ResponseFormat.MARKDOWN)
    .describe("Output format")
});

// Optional fields with defaults
const PaginationSchema = z.object({
  limit: z.number()
    .int()
    .min(1)
    .max(100)
    .default(20)
    .describe("Maximum results to return"),
  offset: z.number()
    .int()
    .min(0)
    .default(0)
    .describe("Number of results to skip")
});
```

## Response Format Options

Support multiple output formats for flexibility:

```typescript
enum ResponseFormat {
  MARKDOWN = "markdown",
  JSON = "json"
}

const inputSchema = z.object({
  query: z.string(),
  response_format: z.nativeEnum(ResponseFormat)
    .default(ResponseFormat.MARKDOWN)
    .describe("Output format: 'markdown' for human-readable or 'json' for machine-readable")
});
```

**Markdown format**:
- Use headers, lists, and formatting for clarity
- Convert timestamps to human-readable format
- Show display names with IDs in parentheses
- Omit verbose metadata
- Group related information logically

**JSON format**:
- Return complete, structured data suitable for programmatic processing
- Include all available fields and metadata
- Use consistent field names and types

## Pagination Implementation

For tools that list resources:

```typescript
const ListSchema = z.object({
  limit: z.number().int().min(1).max(100).default(20),
  offset: z.number().int().min(0).default(0)
});

async function listItems(params: z.infer<typeof ListSchema>) {
  const data = await apiRequest(params.limit, params.offset);

  const response = {
    total: data.total,
    count: data.items.length,
    offset: params.offset,
    items: data.items,
    has_more: data.total > params.offset + data.items.length,
    next_offset: data.total > params.offset + data.items.length
      ? params.offset + data.items.length
      : undefined
  };

  return JSON.stringify(response, null, 2);
}
```

## Character Limits and Truncation

Add a CHARACTER_LIMIT constant to prevent overwhelming responses:

```typescript
// At module level in constants.ts
export const CHARACTER_LIMIT = 25000;  // Maximum response size in characters

async function searchTool(params: SearchInput) {
  let result = generateResponse(data);

  // Check character limit and truncate if needed
  if (result.length > CHARACTER_LIMIT) {
    const truncatedData = data.slice(0, Math.max(1, data.length / 2));
    response.data = truncatedData;
    response.truncated = true;
    response.truncation_message =
      `Response truncated from ${data.length} to ${truncatedData.length} items. ` +
      `Use 'offset' parameter or add filters to see more results.`;
    result = JSON.stringify(response, null, 2);
  }

  return result;
}
```

## Error Handling

Provide clear, actionable error messages:

```typescript
import axios, { AxiosError } from "axios";

function handleApiError(error: unknown): string {
  if (error instanceof AxiosError) {
    if (error.response) {
      switch (error.response.status) {
        case 404:
          return "Error: Resource not found. Please check the ID is correct.";
        case 403:
          return "Error: Permission denied. You don't have access to this resource.";
        case 429:
          return "Error: Rate limit exceeded. Please wait before making more requests.";
        default:
          return `Error: API request failed with status ${error.response.status}`;
      }
    } else if (error.code === "ECONNABORTED") {
      return "Error: Request timed out. Please try again.";
    }
  }
  return `Error: Unexpected error occurred: ${error instanceof Error ? error.message : String(error)}`;
}
```

## Shared Utilities

Extract common functionality into reusable functions:

```typescript
// Shared API request function
async function makeApiRequest<T>(
  endpoint: string,
  method: "GET" | "POST" | "PUT" | "DELETE" = "GET",
  data?: any,
  params?: any
): Promise<T> {
  try {
    const response = await axios({
      method,
      url: `${API_BASE_URL}/${endpoint}`,
      data,
      params,
      timeout: 30000,
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json"
      }
    });
    return response.data;
  } catch (error) {
    throw error;
  }
}
```

## Async/Await Best Practices

Always use async/await for network requests and I/O operations:

```typescript
// Good: Async network request
async function fetchData(resourceId: string): Promise<ResourceData> {
  const response = await axios.get(`${API_URL}/resource/${resourceId}`);
  return response.data;
}

// Bad: Promise chains
function fetchData(resourceId: string): Promise<ResourceData> {
  return axios.get(`${API_URL}/resource/${resourceId}`)
    .then(response => response.data);  // Harder to read and maintain
}
```

## TypeScript Best Practices

1. **Use Strict TypeScript**: Enable strict mode in tsconfig.json
2. **Define Interfaces**: Create clear interface definitions for all data structures
3. **Avoid `any`**: Use proper types or `unknown` instead of `any`
4. **Zod for Runtime Validation**: Use Zod schemas to validate external data
5. **Type Guards**: Create type guard functions for complex type checking
6. **Error Handling**: Always use try-catch with proper error type checking
7. **Null Safety**: Use optional chaining (`?.`) and nullish coalescing (`??`)

```typescript
// Good: Type-safe with Zod and interfaces
interface UserResponse {
  id: string;
  name: string;
  email: string;
  team?: string;
  active: boolean;
}

const UserSchema = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string().email(),
  team: z.string().optional(),
  active: z.boolean()
});

type User = z.infer<typeof UserSchema>;

async function getUser(id: string): Promise<User> {
  const data = await apiCall(`/users/${id}`);
  return UserSchema.parse(data);  // Runtime validation
}

// Bad: Using any
[TRUNCATED]
```

mcp-builder/reference/python_mcp_server.md
```
# Python MCP Server Implementation Guide

## Overview

This document provides Python-specific best practices and examples for implementing MCP servers using the MCP Python SDK. It covers server setup, tool registration patterns, input validation with Pydantic, error handling, and complete working examples.

---

## Quick Reference

### Key Imports
```python
from mcp.server.fastmcp import FastMCP
from pydantic import BaseModel, Field, field_validator, ConfigDict
from typing import Optional, List, Dict, Any
from enum import Enum
import httpx
```

### Server Initialization
```python
mcp = FastMCP("service_mcp")
```

### Tool Registration Pattern
```python
@mcp.tool(name="tool_name", annotations={...})
async def tool_function(params: InputModel) -> str:
    # Implementation
    pass
```

---

## MCP Python SDK and FastMCP

The official MCP Python SDK provides FastMCP, a high-level framework for building MCP servers. It provides:
- Automatic description and inputSchema generation from function signatures and docstrings
- Pydantic model integration for input validation
- Decorator-based tool registration with `@mcp.tool`

**For complete SDK documentation, use WebFetch to load:**
`https://raw.githubusercontent.com/modelcontextprotocol/python-sdk/main/README.md`

## Server Naming Convention

Python MCP servers must follow this naming pattern:
- **Format**: `{service}_mcp` (lowercase with underscores)
- **Examples**: `github_mcp`, `jira_mcp`, `stripe_mcp`

The name should be:
- General (not tied to specific features)
- Descriptive of the service/API being integrated
- Easy to infer from the task description
- Without version numbers or dates

## Tool Implementation

### Tool Naming

Use snake_case for tool names (e.g., "search_users", "create_project", "get_channel_info") with clear, action-oriented names.

**Avoid Naming Conflicts**: Include the service context to prevent overlaps:
- Use "slack_send_message" instead of just "send_message"
- Use "github_create_issue" instead of just "create_issue"
- Use "asana_list_tasks" instead of just "list_tasks"

### Tool Structure with FastMCP

Tools are defined using the `@mcp.tool` decorator with Pydantic models for input validation:

```python
from pydantic import BaseModel, Field, ConfigDict
from mcp.server.fastmcp import FastMCP

# Initialize the MCP server
mcp = FastMCP("example_mcp")

# Define Pydantic model for input validation
class ServiceToolInput(BaseModel):
    '''Input model for service tool operation.'''
    model_config = ConfigDict(
        str_strip_whitespace=True,  # Auto-strip whitespace from strings
        validate_assignment=True,    # Validate on assignment
        extra='forbid'              # Forbid extra fields
    )

    param1: str = Field(..., description="First parameter description (e.g., 'user123', 'project-abc')", min_length=1, max_length=100)
    param2: Optional[int] = Field(default=None, description="Optional integer parameter with constraints", ge=0, le=1000)
    tags: Optional[List[str]] = Field(default_factory=list, description="List of tags to apply", max_items=10)

@mcp.tool(
    name="service_tool_name",
    annotations={
        "title": "Human-Readable Tool Title",
        "readOnlyHint": True,     # Tool does not modify environment
        "destructiveHint": False,  # Tool does not perform destructive operations
        "idempotentHint": True,    # Repeated calls have no additional effect
        "openWorldHint": False     # Tool does not interact with external entities
    }
)
async def service_tool_name(params: ServiceToolInput) -> str:
    '''Tool description automatically becomes the 'description' field.

    This tool performs a specific operation on the service. It validates all inputs
    using the ServiceToolInput Pydantic model before processing.

    Args:
        params (ServiceToolInput): Validated input parameters containing:
            - param1 (str): First parameter description
            - param2 (Optional[int]): Optional parameter with default
            - tags (Optional[List[str]]): List of tags

    Returns:
        str: JSON-formatted response containing operation results
    '''
    # Implementation here
    pass
```

## Pydantic v2 Key Features

- Use `model_config` instead of nested `Config` class
- Use `field_validator` instead of deprecated `validator`
- Use `model_dump()` instead of deprecated `dict()`
- Validators require `@classmethod` decorator
- Type hints are required for validator methods

```python
from pydantic import BaseModel, Field, field_validator, ConfigDict

class CreateUserInput(BaseModel):
    model_config = ConfigDict(
        str_strip_whitespace=True,
        validate_assignment=True
    )

    name: str = Field(..., description="User's full name", min_length=1, max_length=100)
    email: str = Field(..., description="User's email address", pattern=r'^[\w\.-]+@[\w\.-]+\.\w+$')
    age: int = Field(..., description="User's age", ge=0, le=150)

    @field_validator('email')
    @classmethod
    def validate_email(cls, v: str) -> str:
        if not v.strip():
            raise ValueError("Email cannot be empty")
        return v.lower()
```

## Response Format Options

Support multiple output formats for flexibility:

```python
from enum import Enum

class ResponseFormat(str, Enum):
    '''Output format for tool responses.'''
    MARKDOWN = "markdown"
    JSON = "json"

class UserSearchInput(BaseModel):
    query: str = Field(..., description="Search query")
    response_format: ResponseFormat = Field(
        default=ResponseFormat.MARKDOWN,
        description="Output format: 'markdown' for human-readable or 'json' for machine-readable"
    )
```

**Markdown format**:
- Use headers, lists, and formatting for clarity
- Convert timestamps to human-readable format (e.g., "2024-01-15 10:30:00 UTC" instead of epoch)
- Show display names with IDs in parentheses (e.g., "@john.doe (U123456)")
- Omit verbose metadata (e.g., show only one profile image URL, not all sizes)
- Group related information logically

**JSON format**:
- Return complete, structured data suitable for programmatic processing
- Include all available fields and metadata
- Use consistent field names and types

## Pagination Implementation

For tools that list resources:

```python
class ListInput(BaseModel):
    limit: Optional[int] = Field(default=20, description="Maximum results to return", ge=1, le=100)
    offset: Optional[int] = Field(default=0, description="Number of results to skip for pagination", ge=0)

async def list_items(params: ListInput) -> str:
    # Make API request with pagination
    data = await api_request(limit=params.limit, offset=params.offset)

    # Return pagination info
    response = {
        "total": data["total"],
        "count": len(data["items"]),
        "offset": params.offset,
        "items": data["items"],
        "has_more": data["total"] > params.offset + len(data["items"]),
        "next_offset": params.offset + len(data["items"]) if data["total"] > params.offset + len(data["items"]) else None
    }
    return json.dumps(response, indent=2)
```

## Error Handling

Provide clear, actionable error messages:

```python
def _handle_api_error(e: Exception) -> str:
    '''Consistent error formatting across all tools.'''
    if isinstance(e, httpx.HTTPStatusError):
        if e.response.status_code == 404:
            return "Error: Resource not found. Please check the ID is correct."
        elif e.response.status_code == 403:
            return "Error: Permission denied. You don't have access to this resource."
        elif e.response.status_code == 429:
            return "Error: Rate limit exceeded. Please wait before making more requests."
        return f"Error: API request failed with status {e.response.status_code}"
    elif isinstance(e, httpx.TimeoutException):
        return "Error: Request timed out. Please try again."
    return f"Error: Unexpected error occurred: {type(e).__name__}"
```

## Shared Utilities

Extract common functionality into reusable functions:

```python
# Shared API request function
async def _make_api_request(endpoint: str, method: str = "GET", **kwargs) -> dict:
    '''Reusable function for all API calls.'''
    async with httpx.AsyncClient() as client:
        response = await client.request(
            method,
            f"{API_BASE_URL}/{endpoint}",
            timeout=30.0,
            **kwargs
        )
        response.raise_for_status()
        return response.json()
```

## Async/Await Best Practices

Always use async/await for network requests and I/O operations:

```python
# Good: Async network request
async def fetch_data(resource_id: str) -> dict:
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{API_URL}/resource/{resource_id}")
        response.raise_for_status()
        return response.json()

# Bad: Synchronous request
def fetch_data(resource_id: str) -> dict:
    response = requests.get(f"{API_URL}/resource/{resource_id}")  # Blocks
    return response.json()
```

## Type Hints

Use type hints throughout:

```python
from typing import Optional, List, Dict, Any

async def get_user(user_id: str) -> Dict[str, Any]:
    data = await fetch_user(user_id)
    return {"id": data["id"], "name": data["name"]}
```

## Tool Docstrings

Every tool must have comprehensive docstrings with explicit type information:

```python
async def search_users(params: UserSearchInput) -> str:
    '''
    Search for users in the Example system by name, email, or team.

    This tool searches across all user profiles in the Example platform,
    supporting partial matches and various search filters. It does NOT
    create or modify users, only searches existing ones.

    Args:
        params (UserSearchInput): Validated input parameters containing:
            - query (str): Search string to match against names/emails (e.g., "john", "@example.com", "team:marketing")
            - limit (Optional[int]): Maximum results to return, between 1-100 (default: 20)
            - offset (Optional[int]): Number of results to skip for pagination (default: 0)

    Returns:
        str: JSON-formatted string containing search results with the following schema:

        Success response:
        {
            "total": int,           # Total number of matches found
            "count": int,           # Number of results in this response
            "offset": int,          # Current pagination offset
            "users": [
                {
                    "id": str,      # User ID (e.g., "U123456789")
                    "name": str,    # Full name (e.g., "John Doe")
                    "email": str,   # Email address (e.g., "john@example.com")
                    "team": str     # Team name (e.g., "Marketing") - optional
                }
            ]
        }

        Error response:
        "Error: <error message>" or "No users found matching '<query>'"

    Examples:
        - Use when: "Find all marketing team members" -> params with query="team:marketing"
        - Use when: "Search for John's account" -> params with query="john"
        - Don't use when: You need to create a user (use example_create_user instead)
        - Don't use when: You have a user ID and need full details (use example_get_user instead)

    Error Handling:
        - Input validation errors are handled by Pydantic model
        - Returns "Error: Rate limit exceeded" if too many requests (429 status)
        - Returns "Error: Invalid API authentication" if API key is invalid (401 status)
        - Returns formatted list of results or "No users found matching 'query'"
    '''
```

## Complete Example

See below for a complete Python MCP server example:

```python
#!/usr/bin/env python3
'''
MCP Server for Example Service.

This server provides tools to interact with Example API, including user search,
project management, and data export capabilities.
'''

from typing import Optional, List, Dict, Any
from enum import Enum
import httpx
from pydantic import BaseModel, Field, field_validator, ConfigDict
from mcp.server.fastmcp import FastMCP

# Initialize the MCP server
mcp = FastMCP("example_mcp")

# Constants
API_BASE_URL = "https://api.example.com/v1"

# Enums
class ResponseFormat(str, Enum):
    '''Output format for tool responses.'''
    MARKDOWN = "markdown"
    JSON = "json"

# Pydantic Models for Input Validation
class UserSearchInput(BaseModel):
    '''Input model for user search operations.'''
    model_config = ConfigDict(
        str_strip_whitespace=True,
        validate_assignment=True
    )

    query: str = Field(..., description="Search string to match against names/emails", min_length=2, max_length=200)
    limit: Optional[int] = Field(default=20, description="Maximum results to return", ge=1, le=100)
    offset: Optional[int] = Field(default=0, description="Number of results to skip for pagination", ge=0)
    response_format: ResponseFormat = Field(default=ResponseFormat.MARKDOWN, description="Output format")

    @field_validator('query')
    @classmethod
    def validate_query(cls, v: str) -> str:
        if not v.strip():
            raise ValueError("Query cannot be empty or whitespace only")
        return v.strip()

# Shared utility functions
async def _make_api_request(endpoint: str, method: str = "GET", **kwargs) -> dict:
    '''Reusable function for all API calls.'''
    async with httpx.AsyncClient() as client:
        response = await client.request(
            method,
            f"{API_BASE_URL}/{endpoint}",
            timeout=30.0,
            **kwargs
        )
        response.raise_for_status()
        return response.json()

def _handle_api_error(e: Exception) -> str:
    '''Consistent error formatting across all tools.'''
    if isinstance(e, httpx.HTTPStatusError):
        if e.response.status_code == 404:
            return "Error: Resource not found. Please check the ID is correct."
        elif e.response.status_code == 403:
            return "Error: Permission denied. You don't have access to this resource."
        elif e.response.status_code == 429:
            return "Error: Rate limit exceeded. Please wait before making more requests."
        return f"Error: API request failed with status {e.response.status_code}"
    elif isinstance(e, httpx.TimeoutException):
        return "Error: Request timed out. Please try again."
    return f"Error: Unexpected error occurred: {type(e).__name__}"

# Tool definitions
@mcp.tool(
    name="example_search_users",
    annotations={
        "title": "Search Example Users",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": True
    }
)
async def example_search_users(params: UserSearchInput) -> str:
    '''Search for users in the Example system by name, email, or team.

    [Full docstring as shown above]
    '''
    try:
        # Make API request using validated parameters
        data = await _make_api_request(
            "users/search",
            params={
                "q": params.query,
                "limit": params.limit,
                "offset": params.offset
            }
        )

        users = data.get("users", [])
        total = data.get("total", 0)

        if not users:
            return f"No users found matching '{params.query}'"

        # Format response based on requested format
        if params.response_format == ResponseFormat.MARKDOWN:
            lines = [f"# User Search Results: '{params.query}'", ""]
            lines.append(f"Found {total} users (showing {len(users)})")
[TRUNCATED]
```

mcp-builder/scripts/connections.py
```
"""Lightweight connection handling for MCP servers."""

from abc import ABC, abstractmethod
from contextlib import AsyncExitStack
from typing import Any

from mcp import ClientSession, StdioServerParameters
from mcp.client.sse import sse_client
from mcp.client.stdio import stdio_client
from mcp.client.streamable_http import streamablehttp_client


class MCPConnection(ABC):
    """Base class for MCP server connections."""

    def __init__(self):
        self.session = None
        self._stack = None

    @abstractmethod
    def _create_context(self):
        """Create the connection context based on connection type."""

    async def __aenter__(self):
        """Initialize MCP server connection."""
        self._stack = AsyncExitStack()
        await self._stack.__aenter__()

        try:
            ctx = self._create_context()
            result = await self._stack.enter_async_context(ctx)

            if len(result) == 2:
                read, write = result
            elif len(result) == 3:
                read, write, _ = result
            else:
                raise ValueError(f"Unexpected context result: {result}")

            session_ctx = ClientSession(read, write)
            self.session = await self._stack.enter_async_context(session_ctx)
            await self.session.initialize()
            return self
        except BaseException:
            await self._stack.__aexit__(None, None, None)
            raise

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        """Clean up MCP server connection resources."""
        if self._stack:
            await self._stack.__aexit__(exc_type, exc_val, exc_tb)
        self.session = None
        self._stack = None

    async def list_tools(self) -> list[dict[str, Any]]:
        """Retrieve available tools from the MCP server."""
        response = await self.session.list_tools()
        return [
            {
                "name": tool.name,
                "description": tool.description,
                "input_schema": tool.inputSchema,
            }
            for tool in response.tools
        ]

    async def call_tool(self, tool_name: str, arguments: dict[str, Any]) -> Any:
        """Call a tool on the MCP server with provided arguments."""
        result = await self.session.call_tool(tool_name, arguments=arguments)
        return result.content


class MCPConnectionStdio(MCPConnection):
    """MCP connection using standard input/output."""

    def __init__(self, command: str, args: list[str] = None, env: dict[str, str] = None):
        super().__init__()
        self.command = command
        self.args = args or []
        self.env = env

    def _create_context(self):
        return stdio_client(
            StdioServerParameters(command=self.command, args=self.args, env=self.env)
        )


class MCPConnectionSSE(MCPConnection):
    """MCP connection using Server-Sent Events."""

    def __init__(self, url: str, headers: dict[str, str] = None):
        super().__init__()
        self.url = url
        self.headers = headers or {}

    def _create_context(self):
        return sse_client(url=self.url, headers=self.headers)


class MCPConnectionHTTP(MCPConnection):
    """MCP connection using Streamable HTTP."""

    def __init__(self, url: str, headers: dict[str, str] = None):
        super().__init__()
        self.url = url
        self.headers = headers or {}

    def _create_context(self):
        return streamablehttp_client(url=self.url, headers=self.headers)


def create_connection(
    transport: str,
    command: str = None,
    args: list[str] = None,
    env: dict[str, str] = None,
    url: str = None,
    headers: dict[str, str] = None,
) -> MCPConnection:
    """Factory function to create the appropriate MCP connection.

    Args:
        transport: Connection type ("stdio", "sse", or "http")
        command: Command to run (stdio only)
        args: Command arguments (stdio only)
        env: Environment variables (stdio only)
        url: Server URL (sse and http only)
        headers: HTTP headers (sse and http only)

    Returns:
        MCPConnection instance
    """
    transport = transport.lower()

    if transport == "stdio":
        if not command:
            raise ValueError("Command is required for stdio transport")
        return MCPConnectionStdio(command=command, args=args, env=env)

    elif transport == "sse":
        if not url:
            raise ValueError("URL is required for sse transport")
        return MCPConnectionSSE(url=url, headers=headers)

    elif transport in ["http", "streamable_http", "streamable-http"]:
        if not url:
            raise ValueError("URL is required for http transport")
        return MCPConnectionHTTP(url=url, headers=headers)

    else:
        raise ValueError(f"Unsupported transport type: {transport}. Use 'stdio', 'sse', or 'http'")
```

mcp-builder/scripts/evaluation.py
```
"""MCP Server Evaluation Harness

This script evaluates MCP servers by running test questions against them using codex-cli.
"""

import argparse
import asyncio
import json
import re
import sys
import time
import traceback
import xml.etree.ElementTree as ET
from pathlib import Path
from typing import Any

from anthropic import Anthropic

from connections import create_connection

EVALUATION_PROMPT = """You are an AI assistant with access to tools.

When given a task, you MUST:
1. Use the available tools to complete the task
2. Provide summary of each step in your approach, wrapped in <summary> tags
3. Provide feedback on the tools provided, wrapped in <feedback> tags
4. Provide your final response, wrapped in <response> tags

Summary Requirements:
- In your <summary> tags, you must explain:
  - The steps you took to complete the task
  - Which tools you used, in what order, and why
  - The inputs you provided to each tool
  - The outputs you received from each tool
  - A summary for how you arrived at the response

Feedback Requirements:
- In your <feedback> tags, provide constructive feedback on the tools:
  - Comment on tool names: Are they clear and descriptive?
  - Comment on input parameters: Are they well-documented? Are required vs optional parameters clear?
  - Comment on descriptions: Do they accurately describe what the tool does?
  - Comment on any errors encountered during tool usage: Did the tool fail to execute? Did the tool return too many tokens?
  - Identify specific areas for improvement and explain WHY they would help
  - Be specific and actionable in your suggestions

Response Requirements:
- Your response should be concise and directly address what was asked
- Always wrap your final response in <response> tags
- If you cannot solve the task return <response>NOT_FOUND</response>
- For numeric responses, provide just the number
- For IDs, provide just the ID
- For names or text, provide the exact text requested
- Your response should go last"""


def parse_evaluation_file(file_path: Path) -> list[dict[str, Any]]:
    """Parse XML evaluation file with qa_pair elements."""
    try:
        tree = ET.parse(file_path)
        root = tree.getroot()
        evaluations = []

        for qa_pair in root.findall(".//qa_pair"):
            question_elem = qa_pair.find("question")
            answer_elem = qa_pair.find("answer")

            if question_elem is not None and answer_elem is not None:
                evaluations.append({
                    "question": (question_elem.text or "").strip(),
                    "answer": (answer_elem.text or "").strip(),
                })

        return evaluations
    except Exception as e:
        print(f"Error parsing evaluation file {file_path}: {e}")
        return []


def extract_xml_content(text: str, tag: str) -> str | None:
    """Extract content from XML tags."""
    pattern = rf"<{tag}>(.*?)</{tag}>"
    matches = re.findall(pattern, text, re.DOTALL)
    return matches[-1].strip() if matches else None


async def agent_loop(
    client: Anthropic,
    model: str,
    question: str,
    tools: list[dict[str, Any]],
    connection: Any,
) -> tuple[str, dict[str, Any]]:
    """Run the agent loop with MCP tools."""
    messages = [{"role": "user", "content": question}]

    response = await asyncio.to_thread(
        client.messages.create,
        model=model,
        max_tokens=4096,
        system=EVALUATION_PROMPT,
        messages=messages,
        tools=tools,
    )

    messages.append({"role": "assistant", "content": response.content})

    tool_metrics = {}

    while response.stop_reason == "tool_use":
        tool_use = next(block for block in response.content if block.type == "tool_use")
        tool_name = tool_use.name
        tool_input = tool_use.input

        tool_start_ts = time.time()
        try:
            tool_result = await connection.call_tool(tool_name, tool_input)
            tool_response = json.dumps(tool_result) if isinstance(tool_result, (dict, list)) else str(tool_result)
        except Exception as e:
            tool_response = f"Error executing tool {tool_name}: {str(e)}\n"
            tool_response += traceback.format_exc()
        tool_duration = time.time() - tool_start_ts

        if tool_name not in tool_metrics:
            tool_metrics[tool_name] = {"count": 0, "durations": []}
        tool_metrics[tool_name]["count"] += 1
        tool_metrics[tool_name]["durations"].append(tool_duration)

        messages.append({
            "role": "user",
            "content": [{
                "type": "tool_result",
                "tool_use_id": tool_use.id,
                "content": tool_response,
            }]
        })

        response = await asyncio.to_thread(
            client.messages.create,
            model=model,
            max_tokens=4096,
            system=EVALUATION_PROMPT,
            messages=messages,
            tools=tools,
        )
        messages.append({"role": "assistant", "content": response.content})

    response_text = next(
        (block.text for block in response.content if hasattr(block, "text")),
        None,
    )
    return response_text, tool_metrics


async def evaluate_single_task(
    client: Anthropic,
    model: str,
    qa_pair: dict[str, Any],
    tools: list[dict[str, Any]],
    connection: Any,
    task_index: int,
) -> dict[str, Any]:
    """Evaluate a single QA pair with the given tools."""
    start_time = time.time()

    print(f"Task {task_index + 1}: Running task with question: {qa_pair['question']}")
    response, tool_metrics = await agent_loop(client, model, qa_pair["question"], tools, connection)

    response_value = extract_xml_content(response, "response")
    summary = extract_xml_content(response, "summary")
    feedback = extract_xml_content(response, "feedback")

    duration_seconds = time.time() - start_time

    return {
        "question": qa_pair["question"],
        "expected": qa_pair["answer"],
        "actual": response_value,
        "score": int(response_value == qa_pair["answer"]) if response_value else 0,
        "total_duration": duration_seconds,
        "tool_calls": tool_metrics,
        "num_tool_calls": sum(len(metrics["durations"]) for metrics in tool_metrics.values()),
        "summary": summary,
        "feedback": feedback,
    }


REPORT_HEADER = """
# Evaluation Report

## Summary

- **Accuracy**: {correct}/{total} ({accuracy:.1f}%)
- **Average Task Duration**: {average_duration_s:.2f}s
- **Average Tool Calls per Task**: {average_tool_calls:.2f}
- **Total Tool Calls**: {total_tool_calls}

---
"""

TASK_TEMPLATE = """
### Task {task_num}

**Question**: {question}
**Ground Truth Answer**: `{expected_answer}`
**Actual Answer**: `{actual_answer}`
**Correct**: {correct_indicator}
**Duration**: {total_duration:.2f}s
**Tool Calls**: {tool_calls}

**Summary**
{summary}

**Feedback**
{feedback}

---
"""


async def run_evaluation(
    eval_path: Path,
    connection: Any,
    model: str = "codex-cli-3-7-sonnet-20250219",
) -> str:
    """Run evaluation with MCP server tools."""
    print("🚀 Starting Evaluation")

    client = Anthropic()

    tools = await connection.list_tools()
    print(f"📋 Loaded {len(tools)} tools from MCP server")

    qa_pairs = parse_evaluation_file(eval_path)
    print(f"📋 Loaded {len(qa_pairs)} evaluation tasks")

    results = []
    for i, qa_pair in enumerate(qa_pairs):
        print(f"Processing task {i + 1}/{len(qa_pairs)}")
        result = await evaluate_single_task(client, model, qa_pair, tools, connection, i)
        results.append(result)

    correct = sum(r["score"] for r in results)
    accuracy = (correct / len(results)) * 100 if results else 0
    average_duration_s = sum(r["total_duration"] for r in results) / len(results) if results else 0
    average_tool_calls = sum(r["num_tool_calls"] for r in results) / len(results) if results else 0
    total_tool_calls = sum(r["num_tool_calls"] for r in results)

    report = REPORT_HEADER.format(
        correct=correct,
        total=len(results),
        accuracy=accuracy,
        average_duration_s=average_duration_s,
        average_tool_calls=average_tool_calls,
        total_tool_calls=total_tool_calls,
    )

    report += "".join([
        TASK_TEMPLATE.format(
            task_num=i + 1,
            question=qa_pair["question"],
            expected_answer=qa_pair["answer"],
            actual_answer=result["actual"] or "N/A",
            correct_indicator="✅" if result["score"] else "❌",
            total_duration=result["total_duration"],
            tool_calls=json.dumps(result["tool_calls"], indent=2),
            summary=result["summary"] or "N/A",
            feedback=result["feedback"] or "N/A",
        )
        for i, (qa_pair, result) in enumerate(zip(qa_pairs, results))
    ])

    return report


def parse_headers(header_list: list[str]) -> dict[str, str]:
    """Parse header strings in format 'Key: Value' into a dictionary."""
    headers = {}
    if not header_list:
        return headers

    for header in header_list:
        if ":" in header:
            key, value = header.split(":", 1)
            headers[key.strip()] = value.strip()
        else:
            print(f"Warning: Ignoring malformed header: {header}")
    return headers


def parse_env_vars(env_list: list[str]) -> dict[str, str]:
    """Parse environment variable strings in format 'KEY=VALUE' into a dictionary."""
    env = {}
    if not env_list:
        return env

    for env_var in env_list:
        if "=" in env_var:
            key, value = env_var.split("=", 1)
            env[key.strip()] = value.strip()
        else:
            print(f"Warning: Ignoring malformed environment variable: {env_var}")
    return env


async def main():
    parser = argparse.ArgumentParser(
        description="Evaluate MCP servers using test questions",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Evaluate a local stdio MCP server
  python evaluation.py -t stdio -c python -a my_server.py eval.xml

  # Evaluate an SSE MCP server
  python evaluation.py -t sse -u https://example.com/mcp -H "Authorization: Bearer token" eval.xml

  # Evaluate an HTTP MCP server with custom model
  python evaluation.py -t http -u https://example.com/mcp -m codex-cli-3-5-sonnet-20241022 eval.xml
        """,
    )

    parser.add_argument("eval_file", type=Path, help="Path to evaluation XML file")
    parser.add_argument("-t", "--transport", choices=["stdio", "sse", "http"], default="stdio", help="Transport type (default: stdio)")
    parser.add_argument("-m", "--model", default="codex-cli-3-7-sonnet-20250219", help="codex-cli model to use (default: codex-cli-3-7-sonnet-20250219)")

    stdio_group = parser.add_argument_group("stdio options")
    stdio_group.add_argument("-c", "--command", help="Command to run MCP server (stdio only)")
    stdio_group.add_argument("-a", "--args", nargs="+", help="Arguments for the command (stdio only)")
    stdio_group.add_argument("-e", "--env", nargs="+", help="Environment variables in KEY=VALUE format (stdio only)")

    remote_group = parser.add_argument_group("sse/http options")
    remote_group.add_argument("-u", "--url", help="MCP server URL (sse/http only)")
    remote_group.add_argument("-H", "--header", nargs="+", dest="headers", help="HTTP headers in 'Key: Value' format (sse/http only)")

    parser.add_argument("-o", "--output", type=Path, help="Output file for evaluation report (default: stdout)")

    args = parser.parse_args()

    if not args.eval_file.exists():
        print(f"Error: Evaluation file not found: {args.eval_file}")
        sys.exit(1)

    headers = parse_headers(args.headers) if args.headers else None
    env_vars = parse_env_vars(args.env) if args.env else None

    try:
        connection = create_connection(
            transport=args.transport,
            command=args.command,
            args=args.args,
            env=env_vars,
            url=args.url,
            headers=headers,
        )
    except ValueError as e:
        print(f"Error: {e}")
        sys.exit(1)

    print(f"🔗 Connecting to MCP server via {args.transport}...")

    async with connection:
        print("✅ Connected successfully")
        report = await run_evaluation(args.eval_file, connection, args.model)

        if args.output:
            args.output.write_text(report)
            print(f"\n✅ Report saved to {args.output}")
        else:
            print("\n" + report)


if __name__ == "__main__":
    asyncio.run(main())
```

mcp-builder/scripts/example_evaluation.xml
```
<evaluation>
   <qa_pair>
      <question>Calculate the compound interest on $10,000 invested at 5% annual interest rate, compounded monthly for 3 years. What is the final amount in dollars (rounded to 2 decimal places)?</question>
      <answer>11614.72</answer>
   </qa_pair>
   <qa_pair>
      <question>A projectile is launched at a 45-degree angle with an initial velocity of 50 m/s. Calculate the total distance (in meters) it has traveled from the launch point after 2 seconds, assuming g=9.8 m/s². Round to 2 decimal places.</question>
      <answer>87.25</answer>
   </qa_pair>
   <qa_pair>
      <question>A sphere has a volume of 500 cubic meters. Calculate its surface area in square meters. Round to 2 decimal places.</question>
      <answer>304.65</answer>
   </qa_pair>
   <qa_pair>
      <question>Calculate the population standard deviation of this dataset: [12, 15, 18, 22, 25, 30, 35]. Round to 2 decimal places.</question>
      <answer>7.61</answer>
   </qa_pair>
   <qa_pair>
      <question>Calculate the pH of a solution with a hydrogen ion concentration of 3.5 × 10^-5 M. Round to 2 decimal places.</question>
      <answer>4.46</answer>
   </qa_pair>
</evaluation>
```

mcp-builder/scripts/requirements.txt
```
anthropic>=0.39.0
mcp>=1.1.0
```

oraclepack-gold-pack/references/attachment-minimization.md
```
# Attachment minimization rules (Stage 1 packs)

Objective: keep oracle calls fast, portable, and deterministic by attaching the minimum evidence per step.

## Hard limits

- Default: **0–2 attachments per step** (`-f/--file`).
- If you need more than 2, the step is not scoped tightly enough: split or reduce.
- Any `extra_files` the user provides must be appended **literally** (do not reinterpret), but you should still keep the step’s own attachments ≤2.

## What to attach (rule of thumb)

For a given step, prefer:
1) One file that *defines* the concept (contract/schema/config/type)
2) One file that *enforces/uses* the concept (handler/service/policy)

If you can’t find both confidently, attach only the “definition” file.

## Common attachment choices by category (patterns, not requirements)

Use these as **patterns** to recognize likely candidates in a repo; do not assume these paths exist.

- contracts/interfaces:
  - route registration, API schema/spec, public type definitions, CLI command registry

- invariants:
  - domain model definitions, validation layer, schema types, critical service functions that enforce rules

- caching/state:
  - cache client config, state container/store, session manager, any TTL/invalidation logic

- background jobs:
  - worker entrypoint, job registry, scheduler configuration, queue client config

- observability:
  - logger initialization/config, metrics/tracing setup, middleware that injects correlation ids

- permissions:
  - authn/authz middleware, policy definitions, role/scope mapping, guard functions

- migrations:
  - migrations folder index, migration runner config, schema definition file (if present)

- UX flows:
  - key UI/router flow code, top-level workflow orchestrator, controller/handler representing the flow

- failure modes:
  - error handling utilities, retry/backoff config, boundary middleware, circuit breaker wrappers (if any)

- feature flags:
  - flag config/registry, evaluation hook, rollout/targeting logic

## If you cannot find good attachments

- Attach nothing or only 1 file.
- Set `reference=Unknown`.
- Make the prompt request “exact missing file/path pattern(s) to attach next.”

## Avoid these attachment anti-patterns

- Attaching entire directories when one file is enough.
- Attaching multiple duplicates (e.g., the same config in three copies).
- Attaching generated/lock files unless the question is explicitly about them.
- Attaching secrets.
```

oraclepack-gold-pack/references/inference-first-discovery.md
```
# Inference-first discovery (Stage 1 packs)

Goal: pick the *right* 1–2 attachments per step without over-attaching, by inferring repo shape from a small set of anchors.

## Principles

- Prefer **evidence** (actual files) over assumptions.
- Start broad with **cheap, high-signal** files; only then zoom in.
- If a file/path doesn’t exist: record `Unknown` and continue.
- Keep steps self-contained: each step should succeed without relying on shared shell setup.

## Deterministic discovery order

1) **Repo identity + entrypoints**
- `README*` (first ~200 lines)
- top-level manifests (language/framework/package)
- main entrypoints (server start, CLI main, app bootstrap) if obvious from tree

2) **Configuration + environment**
- example config files
- `.env.example` or equivalent (if present)
- CI config files (to infer build/test and deploy steps)

3) **Public surface**
- routing tables / controllers / handlers
- schema/contract definitions (API specs, message schemas, type definitions)
- CLI command registration (if applicable)

4) **Data + jobs + operations**
- data models and storage adapters
- migrations directory (if present)
- background job definitions and worker entrypoints (if present)
- logging/metrics/tracing configuration (if present)

## Turning discovery into step-specific attachments

For each planned step:
- Choose 1 “definition” file (where the thing is declared).
- Optionally choose 1 “use-site” file (where the thing is used/enforced).
- If you can’t confidently pick: attach fewer files and use `reference=Unknown`.

Then write the prompt so the oracle can request missing artifact patterns when needed.

## What to do when evidence is insufficient

- Set `reference=Unknown` in the step header.
- In the prompt, explicitly ask for:
  - “the exact missing file/path pattern(s) to attach next”
- Keep attachments minimal; do not guess file paths.
```

oraclepack-gold-pack/references/oracle-pack-template.md
```
# Oracle Pack — {{codebase_name}} (Gold Stage 1)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- engine: {{engine}}
- model: {{model}}
- extra_files: {{extra_files}}

Notes:
- Template is the **contract**. Keep the pack runner-ingestible.
- Exactly one fenced `bash` block in this whole document.
- Exactly 20 steps, numbered `01..20`.
- Each step includes: `ROI= impact= confidence= effort= horizon= category= reference=`
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-contracts-interfaces-surface.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the primary public interface(s) of this system (API endpoints, CLI commands, public SDK surface, event contracts). For each, list the key request/response (or input/output) shapes and where they are defined in the code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/02-contracts-interfaces-integration.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: Unknown
Category: contracts/interfaces
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the top integration points with external systems (databases, queues, third-party APIs, SSO, storage), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/03-invariants-domain.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
List the system’s most important invariants (business rules, correctness properties, “must always be true” conditions). For each, show where it is enforced (or where it should be enforced but currently is not).

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=invariants reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/04-invariants-validation.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: Unknown
Category: invariants
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where does validation happen (input validation, schema validation, domain validation)? Identify the validation boundaries and the most likely gaps that could cause inconsistent state.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces security/correctness risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/05-caching-state-layers.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What stateful components exist (in-memory state, caches, sessions, client-side state, persisted state)? For each, describe lifecycle, invalidation/expiry, and where it is implemented.

Rationale (one sentence):
State and caching are common sources of subtle bugs and performance issues.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=caching/state reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/06-caching-state-consistency.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: Unknown
Category: caching/state
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Identify the top consistency risks between caches/state layers and the source of truth (stale reads, write skew, missing invalidation). Where are the knobs/configs for cache behavior?

Rationale (one sentence):
Consistency failure modes often surface as “random bugs” and are expensive to debug.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/07-background-jobs-discovery.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What background jobs/workers/scheduled tasks exist? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined.

Rationale (one sentence):
Background work affects reliability, cost, and operational complexity.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=background jobs reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/08-background-jobs-reliability.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the main reliability controls for background work (dead-lettering, backoff, concurrency limits, reprocessing), and what is missing or inconsistent?

Rationale (one sentence):
Reliability controls prevent incident loops and data corruption.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/09-observability-signals.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What observability signals exist (logs/metrics/traces/events), and what are the primary identifiers for correlating a request/job across components? Point to the code/config that defines them.

Rationale (one sentence):
You can’t operate or improve what you can’t measure or debug quickly.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=... impact=... confidence=... effort=... horizon=Immediate category=observability reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/10-observability-gaps.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: Unknown
Category: observability
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are the biggest observability gaps (missing logs around key decisions, missing metrics for SLOs, missing trace spans)? Recommend the smallest additions that would most improve debugging.

Rationale (one sentence):
Targeted observability improvements compound across all future changes.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/11-permissions-model.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the permission model (roles/scopes/claims/ACLs), and where is it defined? Provide the minimal set of files that encode “who can do what.”

Rationale (one sentence):
Permission rules are a high-risk area with security and product impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=... impact=... confidence=... effort=... horizon=Immediate category=permissions reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/12-permissions-enforcement.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Where are permissions enforced (middleware/guards/policies/service-layer checks), and where are likely bypass risks? Identify the enforcement chokepoints and any inconsistent patterns.

Rationale (one sentence):
Enforcement consistency prevents privilege escalation and policy drift.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/13-migrations-schema.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
How are schema/config migrations handled (DB migrations, data backfills, versioned configs)? Identify the tooling, directories, and how migrations are applied in CI/deploy.

Rationale (one sentence):
Migration mechanics are critical for safe releases and rollbacks.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=migrations reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/14-migrations-compat.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: Unknown
Category: migrations
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the backward/forward compatibility expectations during migrations (rolling deploys, dual-read/dual-write, feature-flagged schema use)? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during deployments.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/15-ux-flows-primary.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What are the primary user flows (or primary operator workflows) and their steps? Map each to the main components/modules involved, and note the key state transitions.

Rationale (one sentence):
[TRUNCATED]
```

oraclepack-gold-pack/references/oracle-scratch-format.md
```
# Oracle scratch playbook (NOT a pack format)

This document is for **manual debugging / one-off oracle runs**. It is **not** runner-ingestible oraclepack pack format.

## When to use scratch vs pack

Use the **pack** (`references/oracle-pack-template.md`) when:
- You need a strict 20-step Stage-1 pack for oraclepack ingestion.
- You want deterministic execution and validation via `scripts/validate_pack.py`.

Use **scratch** when:
- You need a single oracle call to explore something quickly.
- You are iterating on prompt wording before committing it into the pack.
- You want to test attachment choices with `--dry-run`.

## Scratch workflow

1) Start with one focused question.
2) Attach 0–2 high-signal files.
3) Use a quoted heredoc prompt to avoid shell-escaping issues.
4) If results are weak, add *one* more attachment (or refine the question).

## Pack-adjacent scratch example (single run)

Example pattern (edit paths/flags to match your environment):

- Uses the quoted heredoc prompt style.
- Shows the optional `--dry-run summary` style (if supported).

```bash
# (optional) preview:
# oracle --dry-run summary --files-report -f "README.md" -p "$(cat <<'PROMPT'
# Explain the repo’s main entrypoint and how requests flow through the system.
# If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
# PROMPT
# )"

oracle \
  --files-report \
  -f "README.md" \
  -p "$(cat <<'PROMPT'
Goal: Understand the repo’s main entrypoints and primary request flow.

Answer format:
1) Direct answer (bullets, evidence-cited)
2) Risks/unknowns
3) Next smallest concrete experiment (one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

## Promoting scratch into the pack

When a scratch run looks good:
- Convert it into a numbered step in the pack.
- Add the strict header tokens (`ROI= impact= confidence= effort= horizon= category= reference=`).
- Add `--write-output "{{out_dir}}/NN-<slug>.md"`.
- Ensure category is one of the fixed 10 and update Coverage check accordingly.

```
```

oraclepack-gold-pack/scripts/lint_attachments.py
```
# path: oraclepack-gold-pack/scripts/lint_attachments.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple


@dataclass
class Step:
    n: str
    header: str
    lines: List[str]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_bash_fence(lines: List[str]) -> List[str]:
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]
    if len(fence_idxs) != 2:
        raise ValueError(f"Expected exactly one fenced block (2 fence lines). Found {len(fence_idxs)}.")
    open_i, close_i = fence_idxs
    if lines[open_i].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash.")
    if lines[close_i].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ```.")
    return [ln.rstrip("\n") for ln in lines[open_i + 1 : close_i]]


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\)\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if not header_idxs:
        raise ValueError("No step headers found inside bash fence.")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(Step(n=n, header=block[0], lines=block))
    return steps


def _count_attachments(step: Step) -> int:
    count = 0
    for ln in step.lines[1:]:
        s = ln.strip()
        if not s or s.startswith("#"):
            continue
        # Count occurrences in non-comment lines
        count += len(re.findall(r"(?<!\S)(-f|--file)(?!\S)", ln))
    return count


def _is_unknown_reference(step: Step) -> bool:
    # Step header token format contains reference=...
    m = re.search(r"\breference=([^\s]+)", step.header)
    if not m:
        return False
    val = m.group(1).strip()
    return val.lower() == "unknown"


def _has_missing_artifact_request(step: Step) -> bool:
    hay = "\n".join(step.lines).lower()
    # Accept several common phrasings, keep it simple and robust.
    patterns = [
        r"missing file/path pattern",
        r"missing file.*pattern",
        r"attach next",
        r"name the exact missing.*pattern",
    ]
    return any(re.search(p, hay) for p in patterns)

def lint(path: Path) -> None:
    raw = _read_text(path)
    lines = raw.splitlines(True)
    fence = _extract_bash_fence(lines)
    steps = _parse_steps(fence)

    errors: List[str] = []
    for step in steps:
        attachments = _count_attachments(step)
        if attachments > 2:
            errors.append(f"Step {step.n}: has {attachments} attachments; must be <= 2 (minimal attachments rule).")

        if _is_unknown_reference(step) and not _has_missing_artifact_request(step):
            errors.append(
                f"Step {step.n}: reference=Unknown but prompt does not request missing file/path pattern(s) to attach next."
            )

    if errors:
        for e in errors:
            print(f"[ERROR] {e}", file=sys.stderr)
        sys.exit(1)

    print("[OK] Attachment lint passed (<=2 attachments per step; Unknown-reference prompts request missing patterns).")

def main() -> None:
    p = argparse.ArgumentParser(description="Lint oraclepack Stage-1 gold pack attachments (<=2 per step) and Unknown-reference prompt behavior.")
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        print(f"[ERROR] File not found: {path}", file=sys.stderr)
        sys.exit(1)

    lint(path)


if __name__ == "__main__":
    main()
```

oraclepack-gold-pack/scripts/validate_pack.py
```
# path: oraclepack-gold-pack/scripts/validate_pack.py
import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Tuple, Optional, Dict


ALLOWED_CATEGORIES = [
    "contracts/interfaces",
    "invariants",
    "caching/state",
    "background jobs",
    "observability",
    "permissions",
    "migrations",
    "UX flows",
    "failure modes",
    "feature flags",
]

REQUIRED_TOKENS = [
    "ROI=",
    "impact=",
    "confidence=",
    "effort=",
    "horizon=",
    "category=",
    "reference=",
]


@dataclass
class Step:
    n: str
    header_line_no: int
    header_line: str
    block_lines: List[str]


def _fail(errors: List[str]) -> None:
    for e in errors:
        print(f"[ERROR] {e}", file=sys.stderr)
    sys.exit(1)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except UnicodeDecodeError:
        # fallback
        return path.read_text(encoding="utf-8", errors="replace")


def _extract_single_bash_fence(lines: List[str]) -> Tuple[int, int, List[str], List[str]]:
    """
    Enforces:
      - exactly two fence lines in entire document
      - first fence line must be exactly ```bash
      - closing fence line must be exactly ```
    Returns: (open_idx, close_idx, fence_lines, outside_lines)
    """
    fence_idxs = [i for i, ln in enumerate(lines) if ln.startswith("```")]

    if len(fence_idxs) != 2:
        raise ValueError(
            f"Expected exactly 1 fenced code block (2 fence lines), found {len(fence_idxs)} fence line(s)."
        )

    open_idx, close_idx = fence_idxs
    if lines[open_idx].rstrip("\n") != "```bash":
        raise ValueError("Opening fence must be exactly ```bash (no extra tokens/spaces).")
    if lines[close_idx].rstrip("\n") != "```":
        raise ValueError("Closing fence must be exactly ``` (no extra tokens/spaces).")
    if close_idx <= open_idx:
        raise ValueError("Closing fence appears before opening fence.")

    fence_lines = lines[open_idx + 1 : close_idx]
    outside_lines = lines[:open_idx] + lines[close_idx + 1 :]
    return open_idx, close_idx, fence_lines, outside_lines


def _parse_steps(fence_lines: List[str]) -> List[Step]:
    header_re = re.compile(r"^#\s*(\d{2})\s+")
    header_idxs: List[Tuple[int, str]] = []
    for i, ln in enumerate(fence_lines):
        m = header_re.match(ln)
        if m:
            header_idxs.append((i, m.group(1)))

    if len(header_idxs) != 20:
        raise ValueError(f"Expected exactly 20 step headers inside bash fence, found {len(header_idxs)}.")

    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [n for _, n in header_idxs]
    if got != expected:
        raise ValueError(f"Step numbering must be sequential 01..20. Got: {got}")

    steps: List[Step] = []
    for idx, (start_i, n) in enumerate(header_idxs):
        end_i = header_idxs[idx + 1][0] if idx + 1 < len(header_idxs) else len(fence_lines)
        block = fence_lines[start_i:end_i]
        steps.append(
            Step(
                n=n,
                header_line_no=start_i + 1,  # 1-based within fence
                header_line=block[0].rstrip("\n"),
                block_lines=[b.rstrip("\n") for b in block],
            )
        )
    return steps


def _validate_header(step: Step, errors: List[str]) -> None:
    header = step.header_line

    for tok in REQUIRED_TOKENS:
        if tok not in header:
            errors.append(f"Step {step.n}: missing required token '{tok}' in header: {header}")

    # ensure each token has a value (non-empty, non-space)
    for key in ["ROI", "impact", "confidence", "effort", "horizon", "category", "reference"]:
        m = re.search(rf"\b{re.escape(key)}=([^\s]+)", header)
        if not m:
            errors.append(f"Step {step.n}: header missing/empty '{key}=' value: {header}")

    cat_m = re.search(r"\bcategory=([^\s]+(?:\s+[^\s]+)?)", header)
    # Category can contain a space (e.g., "background jobs", "UX flows", "failure modes", "feature flags")
    # The regex above captures up to two words; handle 2-word categories explicitly by matching allowed set.
    if cat_m:
        # Extract by checking allowed set presence after 'category='
        after = header.split("category=", 1)[1]
        # category value ends at next token start or end of line
        end = len(after)
        for token in [" reference=", " ROI=", " impact=", " confidence=", " effort=", " horizon="]:
            pos = after.find(token)
            if pos != -1:
                end = min(end, pos)
        cat_val = after[:end].strip()
        if cat_val not in ALLOWED_CATEGORIES:
            errors.append(
                f"Step {step.n}: invalid category='{cat_val}'. Must be one of: {ALLOWED_CATEGORIES}"
            )
    else:
        errors.append(f"Step {step.n}: could not parse category=... from header: {header}")


def _validate_write_output(step: Step, errors: List[str]) -> None:
    # Find first --write-output in the step block
    joined = "\n".join(step.block_lines)
    m = re.search(r'--write-output\s+(["(\S+)"|"(\S+)"|(\S+)])', joined)
    if not m:
        errors.append(f"Step {step.n}: missing --write-output in step block.")
        return

    path = m.group(2) or m.group(3) or m.group(4) or ""
    if "/" not in path:
        errors.append(
            f"Step {step.n}: --write-output path must look like <out_dir>/{step.n}-<slug>.md; got: {path}"
        )
        return

    filename = path.split("/")[-1]
    if not filename.startswith(f"{step.n}-"):
        errors.append(
            f"Step {step.n}: --write-output filename must start with '{step.n}-'; got: {filename}"
        )
    if not filename.endswith(".md"):
        errors.append(f"Step {step.n}: --write-output filename must end with .md; got: {filename}")


def _validate_coverage_check(outside_lines: List[str], errors: List[str]) -> None:
    text = "\n".join(outside_lines)
    # Require a "Coverage check" heading (case-insensitive)
    if re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE) is None:
        errors.append('Missing "## Coverage check" section (must be outside the bash fence).')
        return

    # Ensure every category appears as "<category>:` somewhere after the heading
    # Find slice after the heading
    m = re.search(r"^##\s+Coverage check\s*$", text, flags=re.IGNORECASE | re.MULTILINE)
    assert m is not None
    after = text[m.end() :]

    missing = []
    for cat in ALLOWED_CATEGORIES:
        # Escape special chars (/, etc.)
        if re.search(rf"^\s*[*-]\s+{re.escape(cat)}\s*:", after, flags=re.MULTILINE) is None:
            missing.append(cat)
    if missing:
        errors.append(f'Coverage check missing category lines for: {missing}')


def validate_pack(path: Path) -> None:
    errors: List[str] = []
    raw = _read_text(path)
    lines = raw.splitlines(True)

    try:
        _, _, fence_lines, outside_lines = _extract_single_bash_fence(lines)
    except ValueError as e:
        _fail([str(e)])

    try:
        steps = _parse_steps(fence_lines)
    except ValueError as e:
        _fail([str(e)])

    for step in steps:
        _validate_header(step, errors)
        _validate_write_output(step, errors)

    _validate_coverage_check(outside_lines, errors)

    if errors:
        _fail(errors)

    print("[OK] Pack validates against gold Stage-1 contract.")


def main() -> None:
    p = argparse.ArgumentParser(
        description="Validate an oraclepack Stage-1 gold pack (single bash fence, 20 steps, strict headers, coverage check)."
    )
    p.add_argument("pack_path", help="Path to the Markdown pack file")
    args = p.parse_args()

    path = Path(args.pack_path)
    if not path.exists():
        _fail([f"File not found: {path}"])

    validate_pack(path)


if __name__ == "__main__":
    main()
```

oraclepack-taskify/assets/action-pack-template.md
```
# Oraclepack Stage 3 — Action Pack (Taskify)

Generated: {{pack_date}}

Parsed args (resolved):
- out_dir: {{out_dir}}
- pack_path: {{pack_path}}
- actions_json: {{actions_json}}
- actions_md: {{actions_md}}
- prd_path: {{prd_path}}
- tag: {{tag}}
- mode: {{mode}}
- top_n: {{top_n}}
- oracle_cmd: {{oracle_cmd}}
- task_master_cmd: {{task_master_cmd}}
- tm_cmd: {{tm_cmd}}
- extra_files: (embedded where applicable)

This document must contain exactly one bash code fence, and no other code fences.

```bash
# 01) Preflight: verify inputs and tools (fail fast)
set -euo pipefail

OUT_DIR="{{out_dir}}"
MODE="{{mode}}"
TASK_MASTER_CMD="{{task_master_cmd}}"
ORACLE_CMD="{{oracle_cmd}}"
TM_CMD="{{tm_cmd}}"

if [ ! -d "${OUT_DIR}" ]; then
  echo "ERROR: out_dir does not exist: ${OUT_DIR}" >&2
  exit 2
fi

shopt -s nullglob
for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${OUT_DIR}/${n}-"*.md )
  if [ "${#matches[@]}" -eq 0 ]; then
    echo "ERROR: missing oracle output for prefix ${n}: expected ${OUT_DIR}/${n}-*.md" >&2
    exit 3
  fi
  if [ "${#matches[@]}" -gt 1 ]; then
    echo "ERROR: multiple oracle outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
done

if ! command -v "${TASK_MASTER_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${TASK_MASTER_CMD%% *} (from task_master_cmd='${TASK_MASTER_CMD}')" >&2
  exit 10
fi

if ! command -v "${ORACLE_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${ORACLE_CMD%% *} (from oracle_cmd='${ORACLE_CMD}')" >&2
  exit 11
fi

if [ "${MODE}" = "autopilot" ]; then
  if ! command -v "${TM_CMD%% *}" >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires tm_cmd, but tool missing: ${TM_CMD%% *} (from tm_cmd='${TM_CMD}')" >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 12
  fi
fi

mkdir -p "$(dirname "{{actions_json}}")" "$(dirname "{{actions_md}}")" "$(dirname "{{prd_path}}")" "docs" "${OUT_DIR}"

cat > "${OUT_DIR}/_actions.schema.md" <<'SCHEMA'
# Canonical Actions JSON Schema (human-readable)

This file describes the required structure of `_actions.json` produced in Step 02.

## Root object

- `metadata` (object, required)
  - `generated_at` (string, ISO-8601 recommended)
  - `pack_date` (string, YYYY-MM-DD)
  - `source_out_dir` (string)
  - `repo` (object)
    - `name` (string, optional)
    - `root` (string, optional)
    - `head_sha` (string, optional)
  - `tooling` (object)
    - `oracle_cmd` (string)
    - `task_master_cmd` (string)
  - `top_n` (number)

- `items` (array, required; max 20)
  - Each item is normalized and must include:

## Item fields (required unless marked optional)

- `id` (string): "01".."20"
- `source_file` (string): the exact answer file path used for this item
- `category` (string): a stable label (e.g., `contracts/interfaces`, `permissions`, `observability`, etc.)
- `priority_score` (number): higher means more important (can be ROI if available)
- `recommended_next_action` (string): a single imperative sentence
- `missing_artifacts` (array of strings): file/path globs to locate or create
- `acceptance_criteria` (array of strings): testable, objective conditions
- `risk_notes` (array of strings): specific risks/unknowns, no generic filler
- `estimated_effort` (string): one of `XS|S|M|L|XL` (or a short consistent scale)

## Optional item fields

- `dependencies` (array of strings): other `id` values that should precede this item

## Normalization rules

- Keep `items` sorted by `id` ascending (01..20), regardless of priority.
- Keep all arrays stably ordered (most important first; ties by lexical order).
- Do not include code fences in any string values.
SCHEMA

cat > "${OUT_DIR}/_prd_synthesis_prompt.md" <<'PROMPT'
See the canonical prompt text in the skill asset: assets/prd-synthesis-prompt.md.

This repo-local copy exists for traceability and to keep the Action Pack portable.
PROMPT

echo "OK: Preflight passed."
echo "OK: Inputs: ${OUT_DIR}/01-*.md .. ${OUT_DIR}/20-*.md"
echo "OK: Mode: ${MODE}"


# 02) Synthesize canonical actions JSON + summary MD
set -euo pipefail

OUT_DIR="{{out_dir}}"
ACTIONS_JSON="{{actions_json}}"
ACTIONS_MD="{{actions_md}}"

shopt -s nullglob
oracle_file_flags=()
for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${OUT_DIR}/${n}-"*.md )
  if [ "${#matches[@]}" -ne 1 ]; then
    echo "ERROR: expected exactly one match for ${OUT_DIR}/${n}-*.md, got ${#matches[@]}" >&2
    printf '%s\n' "${matches[@]:-}" >&2
    exit 20
  fi
  oracle_file_flags+=( -f "${matches[0]}" )
done

# Extra attachments (auto-expanded at pack generation time)
# (If none, this section is empty)
{{EXTRA_FILES_LINES}}

mkdir -p "$(dirname "${ACTIONS_JSON}")" "$(dirname "${ACTIONS_MD}")"

{{oracle_cmd}} \
  --write-output "${ACTIONS_JSON}" \
  "${oracle_file_flags[@]}" \
  -p "$(cat <<'PROMPT'
You are producing a SINGLE JSON document and nothing else.

Task: Normalize 20 oraclepack answer files into a canonical actionable plan.

Hard requirements:
- Output MUST be valid JSON (no markdown, no prose, no code fences).
- Output MUST follow the schema described below.
- Output MUST be deterministic in ordering:
  - items sorted by id ascending: 01..20
  - arrays use stable ordering (highest priority first; ties lexical)
- Extract only actionable work. Do not invent repo facts. If evidence is missing, record it in missing_artifacts/risk_notes explicitly.

Schema (summarized):
Root object:
- metadata: { generated_at, pack_date, source_out_dir, repo?, tooling?, top_n }
- items: array (max 20)
Each item:
- id: "01".."20"
- source_file: string
- category: string
- priority_score: number
- recommended_next_action: string (single imperative sentence)
- missing_artifacts: string[]
- acceptance_criteria: string[] (testable)
- risk_notes: string[]
- estimated_effort: "XS"|"S"|"M"|"L"|"XL"
- dependencies?: string[] of ids

Repo/run context:
- pack_date: {{pack_date}}
- source_out_dir: {{out_dir}}
- top_n: {{top_n}}
- tag: {{tag}}

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

Now produce the JSON.
PROMPT
)"

{{oracle_cmd}} \
  --write-output "${ACTIONS_MD}" \
  -f "${ACTIONS_JSON}" \
  -p "$(cat <<'PROMPT'
Write a human-readable Markdown summary of the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown text with headings/bullets only (no code fences).
- Keep ordering aligned with items id ascending (01..20).
- Include:
  - short executive summary (5–10 bullets)
  - top {{top_n}} prioritized list (with id, title inferred from recommended_next_action, category, and why)
  - per-item: recommended_next_action + acceptance_criteria bullets + missing_artifacts bullets
- Do not invent facts; reflect only what is present in the JSON.

Now write the summary Markdown.
PROMPT
)"

echo "OK: Wrote ${ACTIONS_JSON}"
echo "OK: Wrote ${ACTIONS_MD}"


# 03) Generate PRD for Task Master
set -euo pipefail

ACTIONS_JSON="{{actions_json}}"
PRD_PATH="{{prd_path}}"

mkdir -p "$(dirname "${PRD_PATH}")"

{{oracle_cmd}} \
  --write-output "${PRD_PATH}" \
  -f "${ACTIONS_JSON}" \
  -p "$(cat <<'PROMPT'
Write a Task Master-compatible PRD (Markdown) derived from the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Be dependency-aware (use dependencies if present; otherwise infer minimal dependencies cautiously).
- Prioritize focus: select the top N items by priority_score (N = TOP_N), but keep a traceability appendix mapping all ids 01..20.
- Every selected item must become an implementation-ready PRD section with:
  - Goal
  - Scope
  - Non-goals
  - Constraints
  - Acceptance criteria (testable)
  - Risks/unknowns
  - Dependencies (explicit)
- Use the tag value "{{tag}}" in the PRD text where helpful for grouping.

Constants:
- TOP_N={{top_n}}
- TAG={{tag}}

Now produce the PRD.
PROMPT
)"

echo "OK: Wrote ${PRD_PATH}"


# 04) Task Master: parse PRD into tasks
set -euo pipefail

PRD_PATH="{{prd_path}}"
TASK_MASTER_CMD="{{task_master_cmd}}"
TAG="{{tag}}"

if "${TASK_MASTER_CMD}" parse-prd "${PRD_PATH}" --tag "${TAG}" 2>/dev/null; then
  echo "OK: task-master parse-prd (tagged) succeeded."
else
  echo "INFO: task-master parse-prd did not accept --tag; retrying without tag."
  "${TASK_MASTER_CMD}" parse-prd "${PRD_PATH}"
fi

if [ -f ".taskmaster/tasks.json" ]; then
  echo "OK: Found .taskmaster/tasks.json"
elif [ -f "tasks.json" ]; then
  echo "OK: Found tasks.json"
else
  echo "WARN: tasks.json not found at .taskmaster/tasks.json or tasks.json. Check your Task Master configuration/output path."
fi


# 05) Task Master: analyze complexity and save report
set -euo pipefail

TASK_MASTER_CMD="{{task_master_cmd}}"
OUT_DIR="{{out_dir}}"

mkdir -p "${OUT_DIR}"

"${TASK_MASTER_CMD}" analyze-complexity --output "${OUT_DIR}/tm-complexity.json"
echo "OK: Wrote ${OUT_DIR}/tm-complexity.json"


# 06) Task Master: expand tasks
set -euo pipefail

TASK_MASTER_CMD="{{task_master_cmd}}"

"${TASK_MASTER_CMD}" expand --all
echo "OK: Expanded tasks."


# 07) Pipelines (pipelines mode only): generate deterministic pipelines from tasks.json
set -euo pipefail

MODE="{{mode}}"

if [ "${MODE}" != "pipelines" ]; then
  echo "SKIP: mode=${MODE} (pipelines step runs only when mode=pipelines)."
else
  tasks_path=""
  if [ -f ".taskmaster/tasks.json" ]; then
    tasks_path=".taskmaster/tasks.json"
  elif [ -f "tasks.json" ]; then
    tasks_path="tasks.json"
  else
    echo "ERROR: tasks.json not found; cannot generate pipelines." >&2
    exit 70
  fi

  mkdir -p "docs"

  {{oracle_cmd}} \
    --write-output "docs/oracle-actions-pipelines.md" \
    -f "${tasks_path}" \
    -p "$(cat <<'PROMPT'
Generate deterministic command pipelines from tasks.json.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Include 3–6 pipelines, each a numbered list of shell commands.
- Each pipeline must be tests-first and avoid destructive operations.
- Commands should be generic and repo-agnostic (no invented scripts).
- Include a short “resume strategy” section explaining how to re-run pipelines safely.

Now write docs/oracle-actions-pipelines.md content.
PROMPT
)"

  echo "OK: Wrote docs/oracle-actions-pipelines.md"
fi


# 08) Autopilot (autopilot mode only): branch safety + guarded entrypoint
set -euo pipefail

MODE="{{mode}}"
TM_CMD="{{tm_cmd}}"
OUT_DIR="{{out_dir}}"
PACK_DATE="{{pack_date}}"
TAG="{{tag}}"

if [ "${MODE}" != "autopilot" ]; then
  echo "SKIP: mode=${MODE} (autopilot step runs only when mode=autopilot)."
else
  if ! command -v git >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires git on PATH." >&2
    exit 80
  fi

  if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "ERROR: not inside a git work tree; autopilot mode requires a git repo." >&2
    exit 81
  fi

  if ! git diff --quiet || ! git diff --cached --quiet; then
    echo "ERROR: working tree not clean. Commit/stash before autopilot." >&2
    exit 82
  fi

  current_branch="$(git rev-parse --abbrev-ref HEAD)"
  if [ "${current_branch}" = "main" ] || [ "${current_branch}" = "master" ]; then
    new_branch="oraclepack/${PACK_DATE}-${TAG}"
    echo "INFO: on default-like branch (${current_branch}); creating work branch: ${new_branch}"
    git checkout -b "${new_branch}"
  else
    echo "OK: current branch is ${current_branch}"
  fi

  mkdir -p "${OUT_DIR}"
  cat > "${OUT_DIR}/tm-autopilot.state.json" <<STATE
{"pack_date":"${PACK_DATE}","tag":"${TAG}","mode":"autopilot","notes":"State file created by Stage-3 Action Pack. Autopilot tooling should resume from this file if supported."}
STATE

  echo "OK: Wrote ${OUT_DIR}/tm-autopilot.state.json"
  echo "INFO: Starting autopilot via: ${TM_CMD} autopilot"
  echo "INFO: If your tm tool uses a different subcommand, edit this step accordingly."

  if ! "${TM_CMD}" --help 2>&1 | grep -qi "autopilot"; then
    echo "ERROR: '${TM_CMD}' does not advertise 'autopilot' in --help output." >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 83
  fi

  "${TM_CMD}" autopilot
fi
```
```

oraclepack-taskify/assets/actions-json-schema.md
```
# Canonical Actions JSON Schema (human-readable)

This schema defines the required structure of the canonical actions file:

- Default path: `<out_dir>/_actions.json`

## Root object

The root MUST be a JSON object with:

### metadata (required)

- `generated_at` (string): generation timestamp (ISO-8601 recommended)
- `pack_date` (string): `YYYY-MM-DD`
- `source_out_dir` (string): the oraclepack output dir used (e.g., `oracle-out`)
- `repo` (object, optional):
  - `name` (string, optional)
  - `root` (string, optional)
  - `head_sha` (string, optional)
- `tooling` (object, optional):
  - `oracle_cmd` (string)
  - `task_master_cmd` (string)
- `top_n` (number): the top-N focus value used to build PRD/pipelines

### items (required; max 20)

`items` MUST be an array with up to 20 objects. Each item MUST include:

- `id` (string): `"01"`..`"20"`
- `source_file` (string): the specific answer file used for this item
- `category` (string): stable label describing the domain area
- `priority_score` (number): higher means higher priority (can be ROI if present)
- `recommended_next_action` (string): single imperative sentence
- `missing_artifacts` (array of strings): file/path globs or concrete paths
- `acceptance_criteria` (array of strings): testable, objective conditions
- `risk_notes` (array of strings): risks/unknowns grounded in evidence gaps
- `estimated_effort` (string): use a consistent scale such as `XS|S|M|L|XL`

Optional:

- `dependencies` (array of strings): other ids (e.g., `["03","07"]`) that should precede this item

## Normalization rules

- Items MUST be sorted by `id` ascending (`01..20`) for machine stability.
- Within each item:
  - `missing_artifacts`, `acceptance_criteria`, and `risk_notes` MUST be stably ordered.
- Do not include fenced code blocks in any string values.
```

oraclepack-taskify/assets/prd-synthesis-prompt.md
```
# Stage 3 Synthesis Prompts (exact text)

Use these prompts verbatim in the Action Pack.

## Prompt A — Canonical actions JSON (_actions.json)

You are producing a SINGLE JSON document and nothing else.

Task: Normalize 20 oraclepack answer files into a canonical actionable plan.

Hard requirements:
- Output MUST be valid JSON (no markdown, no prose, no code fences).
- Output MUST follow the schema described below.
- Output MUST be deterministic in ordering:
  - items sorted by id ascending: 01..20
  - arrays use stable ordering (highest priority first; ties lexical)
- Extract only actionable work. Do not invent repo facts. If evidence is missing, record it in missing_artifacts/risk_notes explicitly.

Schema (summarized):
Root object:
- metadata: { generated_at, pack_date, source_out_dir, repo?, tooling?, top_n }
- items: array (max 20)
Each item:
- id: "01".."20"
- source_file: string
- category: string
- priority_score: number
- recommended_next_action: string (single imperative sentence)
- missing_artifacts: string[]
- acceptance_criteria: string[] (testable)
- risk_notes: string[]
- estimated_effort: "XS"|"S"|"M"|"L"|"XL"
- dependencies?: string[] of ids

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

## Prompt B — Task Master PRD (oracle-actions-prd.md)

Write a Task Master-compatible PRD (Markdown) derived from the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Be dependency-aware (use dependencies if present; otherwise infer minimal dependencies cautiously).
- Prioritize focus: select the top N items by priority_score (N = TOP_N), but keep a traceability appendix mapping all ids 01..20.
- Every selected item must become an implementation-ready PRD section with:
  - Goal
  - Scope
  - Non-goals
  - Constraints
  - Acceptance criteria (testable)
  - Risks/unknowns
  - Dependencies (explicit)

Hygiene:
- Do not invent scripts/paths; use missing_artifacts when you need the repo to supply something.
- Keep acceptance criteria objective and testable.
```

oraclepack-taskify/references/determinism-and-safety.md
```
# Determinism and safety guardrails

## Determinism rules

- Always select inputs by prefix ordering:
  - exactly one match for each: `01-*.md` … `20-*.md`
- If any prefix is missing or has multiple matches, exit non-zero with a precise error.
- Keep generated JSON normalized:
  - items sorted by id ascending (01..20)
  - stable ordering for arrays
- Keep output paths explicit and stable:
  - do not rely on shared environment variables across steps
  - each step re-declares its constants

## Safety rules

- No interactive prompts in the Action Pack.
- Fail fast when prerequisites are missing:
  - `task-master` (or override)
  - `oracle` (or override)
  - `tm` only in autopilot mode (default)
- Always `mkdir -p` parent directories before writing files.
- Avoid destructive operations:
  - do not delete
  - do not force push
  - do not commit to main/master in autopilot mode
- Autopilot mode:
  - require a clean working tree
  - if on main/master, create a new work branch before starting autopilot
  - write a state file to support resumption

## Failure behavior

- Prefer explicit, early errors over partial or ambiguous outputs.
- If Task Master output paths differ from defaults, print warnings but keep the pack deterministic.
```

oraclepack-taskify/references/task-master-cli-cheatsheet.md
```
# Task Master CLI cheatsheet (minimal)

This skill assumes only these Task Master commands:

## Parse PRD into tasks

- `task-master parse-prd <prd_path>`
- If tag scoping is supported in your setup, the Action Pack attempts:
  - `task-master parse-prd <prd_path> --tag <tag>`
  - and falls back to the untagged command if the flag is not accepted.

## Analyze complexity

- `task-master analyze-complexity --output <out_dir>/tm-complexity.json`

## Expand tasks

- `task-master expand --all`

## Autopilot (default mode behavior)

- The Action Pack attempts: `tm autopilot` (via `tm_cmd`, default `tm`)
- If your `tm` tool does not support autopilot, run Stage 3 with `mode=backlog` or `mode=pipelines`.

Notes:
- The Action Pack checks for `.taskmaster/tasks.json` or `tasks.json` after parsing, but Task Master may be configured differently. If neither file exists, the pack prints a warning.
```

oraclepack-taskify/references/workflow-overview.md
```
# Stage 3 (oraclepack-taskify) — Workflow overview

## What this stage solves

Stage 1 produces a 20-question oracle pack.
Stage 2 runs oraclepack and produces 20 answer files.

Stage 2 outputs are answers, not work. Stage 3 creates the deterministic bridge from answers to executable planning artifacts and (by default) starts a guarded autopilot to begin implementation.

## Inputs

- A completed oraclepack output directory containing exactly:
  - `01-*.md` … `20-*.md` (one file per prefix)
- Optional additional files to improve synthesis fidelity (extra attachments)

## Primary output (this skill generates)

- An “Action Pack” markdown file at:
  - default: `docs/oracle-actions-pack-YYYY-MM-DD.md`
  - override: `pack_path=...`

The Action Pack is designed to be executed as a deterministic pipeline.

## Artifacts the Action Pack produces when executed

- Canonical actions:
  - `<out_dir>/_actions.json` (machine-consumable)
  - `<out_dir>/_actions.md` (human summary)
- PRD/spec suitable for Task Master:
  - `.taskmaster/docs/oracle-actions-prd.md`
- Task Master outputs:
  - tasks created/expanded by `task-master`
  - complexity report: `<out_dir>/tm-complexity.json`
- Optional:
  - pipelines doc: `docs/oracle-actions-pipelines.md` (pipelines mode)
  - autopilot entrypoint + state file (autopilot mode, default)

## Execution modes

- backlog: actions → PRD → tasks
- pipelines: backlog + pipelines generation
- autopilot (default): backlog + guarded autopilot entrypoint
```

oraclepack-taskify/scripts/detect-oracle-outputs.sh
```
#!/usr/bin/env bash
set -euo pipefail

out_dir="${1:-oracle-out}"

if [[ ! -d "${out_dir}" ]]; then
  echo "ERROR: out_dir does not exist: ${out_dir}" >&2
  exit 2
fi

shopt -s nullglob

for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${out_dir}/${n}-"*.md )
  if [[ "${#matches[@]}" -eq 0 ]]; then
    echo "ERROR: missing output for prefix ${n}: expected ${out_dir}/${n}-*.md" >&2
    exit 3
  fi
  if [[ "${#matches[@]}" -gt 1 ]]; then
    echo "ERROR: multiple outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
  printf '%s\n' "${matches[0]}"
done
```

oraclepack-taskify/scripts/validate-action-pack.sh
```
#!/usr/bin/env bash
set -euo pipefail

pack_path="${1:-}"
if [[ -z "${pack_path}" ]]; then
  echo "Usage: validate-action-pack.sh <path/to/oracle-actions-pack.md>" >&2
  exit 2
fi

if [[ ! -f "${pack_path}" ]]; then
  echo "ERROR: file not found: ${pack_path}" >&2
  exit 3
fi

# Rule: exactly one bash fence, and no other fences.
bash_fence_count="$(grep -cE '^[[:space:]]*```bash[[:space:]]*$' "${pack_path}" || true)"
any_fence_count="$(grep -cE '^[[:space:]]*```' "${pack_path}" || true)"

if [[ "${bash_fence_count}" -ne 1 ]]; then
  echo "ERROR: expected exactly one '```bash' fence; found ${bash_fence_count}" >&2
  exit 10
fi

if [[ "${any_fence_count}" -ne 2 ]]; then
  echo "ERROR: expected exactly 2 total fences (start+end); found ${any_fence_count}" >&2
  echo "Fences found:" >&2
  grep -nE '^[[:space:]]*```' "${pack_path}" >&2 || true
  exit 11
fi

# Extract lines within the bash fence and validate step headers.
in_bash=0
headers=()
while IFS= read -r line; do
  if [[ "${line}" =~ ^[[:space:]]*```bash[[:space:]]*$ ]]; then
    in_bash=1
    continue
  fi
  if [[ "${line}" =~ ^[[:space:]]*```[[:space:]]*$ ]]; then
    in_bash=0
    continue
  fi
  if [[ "${in_bash}" -eq 1 ]]; then
    if [[ "${line}" =~ ^#\ ([0-9]{2})\) ]]; then
      headers+=( "${BASH_REMATCH[1]}" )
    fi
  fi
done < "${pack_path}"

if [[ "${#headers[@]}" -lt 1 ]]; then
  echo "ERROR: no step headers found inside bash fence (expected '# NN)')" >&2
  exit 20
fi

# Validate strict sequential: 01..N with no gaps and no duplicates.
seen=""
expected=1
for h in "${headers[@]}"; do
  # Check duplicates
  if [[ " ${seen} " == *" ${h} "* ]]; then
    echo "ERROR: duplicate step header: ${h}" >&2
    exit 21
  fi
  seen="${seen} ${h}"

  exp="$(printf '%02d' "${expected}")"
  if [[ "${h}" != "${exp}" ]]; then
    echo "ERROR: non-sequential step header. Expected ${exp}, got ${h}" >&2
    echo "All headers: ${headers[*]}" >&2
    exit 22
  fi
  expected=$((expected + 1))
done

echo "OK: Action Pack validation passed."
```

</source_code>