<filetree>
Project Structure:
├── .config
│   ├── commands
│   │   └── oracle-pack_v2.toml
│   ├── completion
│   │   └── oraclepack.completion.sh
│   ├── mcp
│   │   ├── mcp-builder
│   │   │   ├── reference
│   │   │   │   ├── evaluation.md
│   │   │   │   ├── mcp_best_practices.md
│   │   │   │   ├── node_mcp_server.md
│   │   │   │   └── python_mcp_server.md
│   │   │   ├── scripts
│   │   │   │   ├── connections.py
│   │   │   │   ├── evaluation.py
│   │   │   │   ├── example_evaluation.xml
│   │   │   │   └── requirements.txt
│   │   │   └── SKILL.md
│   │   ├── oraclepack-gold-pack
│   │   │   ├── references
│   │   │   │   ├── attachment-minimization.md
│   │   │   │   ├── inference-first-discovery.md
│   │   │   │   ├── oracle-pack-template.md
│   │   │   │   └── oracle-scratch-format.md
│   │   │   ├── scripts
│   │   │   │   ├── lint_attachments.py
│   │   │   │   └── validate_pack.py
│   │   │   └── SKILL.md
│   │   └── oraclepack-taskify
│   │       ├── assets
│   │       │   ├── action-pack-template.md
│   │       │   ├── actions-json-schema.md
│   │       │   └── prd-synthesis-prompt.md
│   │       ├── references
│   │       │   ├── determinism-and-safety.md
│   │       │   ├── task-master-cli-cheatsheet.md
│   │       │   └── workflow-overview.md
│   │       ├── scripts
│   │       │   ├── detect-oracle-outputs.sh
│   │       │   └── validate-action-pack.sh
│   │       └── SKILL.md
│   ├── scripts
│   │   ├── build_install_oraclepack.md
│   │   ├── build_install_oraclepack.sh
│   │   ├── install-global.ps1
│   │   ├── install-global.sh
│   │   └── tag-release.sh
│   └── skills
│       ├── oracle-pack
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   └── oracle-scratch-format.md
│       │   └── SKILL.md
│       ├── oracle-pack-rpg-infused
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   └── oracle-pack-spec.md
│       │   ├── scripts
│       │   │   └── validate_oracle_pack.py
│       │   └── SKILL.md
│       ├── oracle-strategist-commands
│       │   └── SKILL.md
│       ├── oraclepack-gold-pack
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   ├── oracle-pack-template.md
│       │   │   └── oracle-scratch-format.md
│       │   ├── scripts
│       │   │   ├── lint_attachments.py
│       │   │   └── validate_pack.py
│       │   └── SKILL.md
│       ├── oraclepack-taskify
│       │   ├── assets
│       │   │   ├── action-pack-template.md
│       │   │   ├── actions-json-schema.md
│       │   │   └── prd-synthesis-prompt.md
│       │   ├── references
│       │   │   ├── determinism-and-safety.md
│       │   │   ├── task-master-cli-cheatsheet.md
│       │   │   └── workflow-overview.md
│       │   ├── scripts
│       │   │   ├── detect-oracle-outputs.sh
│       │   │   └── validate-action-pack.sh
│       │   └── SKILL.md
│       ├── oraclepack-usecase-rpg
│       │   ├── references
│       │   │   └── oracle-pack-template.md
│       │   ├── scripts
│       │   │   └── validate_oraclepack_pack.py
│       │   └── SKILL.md
│       └── strategist-questions-oracle-pack
│           ├── assets
│           │   └── oracle-pack-template.md
│           ├── references
│           │   ├── attachment-minimization.md
│           │   ├── inference-first-discovery.md
│           │   └── oracle-scratch-format.md
│           └── SKILL.md
├── .github
│   └── workflows
│       ├── ci.yml
│       └── release.yml
├── .ruler
│   ├── AGENTS.md
│   ├── ruler.toml
│   └── tm-AGENTS.md
├── .rules
│   ├── dev_workflow.md
│   ├── rules.md
│   ├── self_improve.md
│   └── taskmaster.md
├── .tickets
│   ├── mcp
│   │   ├── Expose Oraclepack as MCP.md
│   │   ├── MCP Server for Oraclepack.md
│   │   ├── oraclepack-MCP.md
│   │   └── oraclepack_mcp_server.md
│   ├── Enable Action Packs Dispatch.md
│   ├── Improving Oraclepack Workflow.md
│   ├── Oraclepack Action Pack Integration.md
│   ├── Oraclepack Action Pack Issue.md
│   ├── Oraclepack Action Packs.md
│   ├── Oraclepack Pipeline Improvements.md
│   ├── Oraclepack Prompt Generator.md
│   ├── Oraclepack Workflow Enhancement.md
│   └── Verbose Payload Rendering TUI.md
├── docs
│   ├── oracle-questions-2026-01-07
│   │   ├── 01-contracts-interfaces-surface.md
│   │   ├── 02-contracts-interfaces-integration.md
│   │   ├── 03-invariants-domain.md
│   │   ├── 04-invariants-validation.md
│   │   ├── 05-caching-state-layers.md
│   │   ├── 06-caching-state-consistency.md
│   │   ├── 07-background-jobs-discovery.md
│   │   ├── 08-background-jobs-reliability.md
│   │   ├── 09-observability-signals.md
│   │   ├── 10-observability-gaps.md
│   │   ├── 11-permissions-model.md
│   │   ├── 12-permissions-enforcement.md
│   │   ├── 13-migrations-schema.md
│   │   ├── 14-migrations-compat.md
│   │   ├── 15-ux-flows-primary.md
│   │   ├── 16-ux-flows-edgecases.md
│   │   ├── 17-failure-modes-taxonomy.md
│   │   ├── 18-failure-modes-resilience.md
│   │   ├── 19-feature-flags-inventory.md
│   │   └── 20-feature-flags-rollout.md
│   ├── oracle-actions-pack-2026-01-07.md
│   └── oracle-pack-2026-01-07.md
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── app_test.go
│   │   ├── run.go
│   │   └── run_test.go
│   ├── cli
│   │   ├── cmds.go
│   │   ├── root.go
│   │   └── run.go
│   ├── errors
│   │   ├── errors.go
│   │   └── errors_test.go
│   ├── exec
│   │   ├── flags.go
│   │   ├── inject.go
│   │   ├── inject_test.go
│   │   ├── oracle_scan.go
│   │   ├── oracle_scan_test.go
│   │   ├── oracle_validate.go
│   │   ├── oracle_validate_test.go
│   │   ├── runner.go
│   │   ├── runner_test.go
│   │   ├── sanitize.go
│   │   ├── sanitize_test.go
│   │   └── stream.go
│   ├── overrides
│   │   ├── merge.go
│   │   ├── merge_test.go
│   │   └── types.go
│   ├── pack
│   │   ├── parser.go
│   │   ├── parser_test.go
│   │   └── types.go
│   ├── render
│   │   ├── render.go
│   │   └── render_test.go
│   ├── report
│   │   ├── generate.go
│   │   ├── report_test.go
│   │   └── types.go
│   ├── state
│   │   ├── persist.go
│   │   ├── state_test.go
│   │   └── types.go
│   └── tui
│       ├── clipboard.go
│       ├── filter_test.go
│       ├── overrides_confirm.go
│       ├── overrides_flags.go
│       ├── overrides_flow.go
│       ├── overrides_steps.go
│       ├── overrides_url.go
│       ├── preview_test.go
│       ├── tui.go
│       ├── tui_test.go
│       ├── url_picker.go
│       ├── url_store.go
│       └── url_store_test.go
├── .browser-echo-mcp.json
├── .goreleaser.yaml
├── go.mod
├── oracle-pack-2026-01-02.chatgpt-urls.json
├── oracle-pack-2026-01-02.state.json
└── package.json

</filetree>

<source_code>
.browser-echo-mcp.json
```
{"url":"http://127.0.0.1:37553","route":"/__client-logs","timestamp":1767829000345,"pid":79099}
```

.goreleaser.yaml
```
# path: .goreleaser.yaml
version: 2

project_name: oraclepack

before:
  hooks:
    - go mod download

builds:
  - id: oraclepack
    main: ./cmd/oraclepack
    binary: oraclepack
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    flags:
      - -trimpath

archives:
  - id: default
    builds:
      - oraclepack
    name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"
    format: tar.gz
    format_overrides:
      - goos: windows
        format: zip

checksum:
  name_template: "checksums.txt"

changelog:
  sort: asc
  filters:
    exclude:
      - "^docs:"
      - "^test:"
      - "^chore:"
```

go.mod
```
module github.com/user/oraclepack

go 1.24.0

toolchain go1.24.11

require (
	github.com/charmbracelet/bubbles v0.21.0
	github.com/charmbracelet/bubbletea v1.3.10
	github.com/charmbracelet/glamour v0.10.0
	github.com/charmbracelet/lipgloss v1.1.1-0.20250404203927-76690c660834
	github.com/spf13/cobra v1.10.2
)

require (
	github.com/alecthomas/chroma/v2 v2.14.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/charmbracelet/colorprofile v0.2.3-0.20250311203215-f60798e515dc // indirect
	github.com/charmbracelet/x/ansi v0.10.1 // indirect
	github.com/charmbracelet/x/cellbuf v0.0.13 // indirect
	github.com/charmbracelet/x/exp/slice v0.0.0-20250327172914-2fdc97757edf // indirect
	github.com/charmbracelet/x/term v0.2.1 // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/microcosm-cc/bluemonday v1.0.27 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.16.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sahilm/fuzzy v0.1.1 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/yuin/goldmark v1.7.8 // indirect
	github.com/yuin/goldmark-emoji v1.0.5 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/term v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)
```

oracle-pack-2026-01-02.chatgpt-urls.json
```
{
  "default": "",
  "items": [
    {
      "name": "oracle",
      "url": "https://chatgpt.com/g/g-p-694ed925bab08191acaf80aefaf27dfc-oracle/project",
      "lastUsed": "2026-01-02T22:40:05Z"
    }
  ]
}
```

oracle-pack-2026-01-02.state.json
```
{
  "schema_version": 1,
  "pack_hash": "",
  "start_time": "0001-01-01T00:00:00Z",
  "step_statuses": {},
  "roi_threshold": 1.6,
  "roi_mode": "over"
}
```

package.json
```
{
  "devDependencies": {
    "codefetch": "^2.2.0"
  },
  "scripts": {
    "code": "pnpm code:tui && pnpm code:all && pnpm code:t && pnpm code:op && pnpm code:opt",
    "code:tui": "codefetch -t 5 --include-dir cmd,internal -o oraclepack-tui.md --max-tokens 50000 --token-limiter truncated",
    "code:all": "codefetch -t 5 -o oraclepack-all.md",
    "code:temp": "codefetch -t 5 --include-dir .rules -o rules.md --max-tokens 50000",
    "build": "bash .config/scripts/build_install_oraclepack.sh && exec bash",
    "tag": "bash .config/scripts/tag-release.sh",
    "code:op": "codefetch -t 5 --include-files \"**/oracle*\",\"**/oracle*/**/*\" -o oracle_SKILL_and_PROMPTS.md --max-tokens 75000",
    "code:t": "codefetch -t 5 --include-dir .tickets -o oraclepack-tickets.md --max-tokens 50000 --token-limiter truncated",
    "code:mcp": "codefetch -t 5 --include-dir .config/mcp -o oraclepack-mcp.md --max-tokens 50000 --token-limiter truncated",
    "code:tm": "codefetch -t 5 --include-files \"**/tm-*\",\"**/tm-*/**/*\" -o tm_SKILL_and_PROMPTS.md --max-tokens 75000",
    "code:opt": "codefetch -t 5 --include-files \".config/skills/oraclepack-taskify/**/*.md\",\".config/skills/oraclepack-gold-pack/**/*.md\" -o oraclepack-goldpack-taskify_SKILL.md --max-tokens 75000 --token-limiter truncated"
  }
}
```

.ruler/AGENTS.md
```
# AGENTS.md

Centralised AI agent instructions. Add coding guidelines, style guides, and project context here.

Ruler concatenates all .md files in this directory (and subdirectories), starting with AGENTS.md (if present), then remaining files in sorted order.
```

.ruler/ruler.toml
```
# Ruler Configuration File
# See https://ai.intellectronica.net/ruler for documentation.

# To specify which agents are active by default when --agents is not used,
# uncomment and populate the following line. If omitted, all agents are active.
default_agents = ["codex"]

# Enable nested rule loading from nested .ruler directories
# When enabled, ruler will search for and process .ruler directories throughout the project hierarchy
nested = true

# --- Agent Specific Configurations ---
# You can enable/disable agents and override their default output paths here.
# Use lowercase agent identifiers: amp, copilot, claude, codex, cursor, windsurf, cline, aider, kilocode

# [agents.copilot]
# enabled = true
# output_path = ".github/copilot-instructions.md"

# [agents.aider]
# enabled = true
# output_path_instructions = "AGENTS.md"
# output_path_config = ".aider.conf.yml"

# [agents.gemini-cli]
# enabled = true

# --- MCP Servers ---
# Define Model Context Protocol servers here. Two examples:
# 1. A stdio server (local executable)
# 2. A remote server (HTTP-based)

# [mcp_servers.example_stdio]
# command = "node"
# args = ["scripts/your-mcp-server.js"]
# env = { API_KEY = "replace_me" }

# [mcp_servers.example_remote]
# url = "https://api.example.com/mcp"
# headers = { Authorization = "Bearer REPLACE_ME" }
```

.ruler/tm-AGENTS.md
```
# Agent Instructions

## Task Master AI and Workflow Instructions
**Import Task Master's development workflow commands and guidelines, treat as if import is in the main AGENT.md file.**
@./.taskmaster/AGENTS.md @./.rules/dev_workflow.md @./.rules/rules.md @./.rules/self_improve.md @./.rules/taskmaster.md
```

.rules/dev_workflow.md
```
---
description: Guide for using Taskmaster to manage task-driven development workflows
---

# Taskmaster Development Workflow

This guide outlines the standard process for using Taskmaster to manage software development projects. It is written as a set of instructions for you, the AI agent.

- **Your Default Stance**: For most projects, the user can work directly within the `master` task context. Your initial actions should operate on this default context unless a clear pattern for multi-context work emerges.
- **Your Goal**: Your role is to elevate the user's workflow by intelligently introducing advanced features like **Tagged Task Lists** when you detect the appropriate context. Do not force tags on the user; suggest them as a helpful solution to a specific need.

## The Basic Loop
The fundamental development cycle you will facilitate is:
1.  **`list`**: Show the user what needs to be done.
2.  **`next`**: Help the user decide what to work on.
3.  **`show <id>`**: Provide details for a specific task.
4.  **`expand <id>`**: Break down a complex task into smaller, manageable subtasks.
5.  **Implement**: The user writes the code and tests.
6.  **`update-subtask`**: Log progress and findings on behalf of the user.
7.  **`set-status`**: Mark tasks and subtasks as `done` as work is completed.
8.  **Repeat**.

All your standard command executions should operate on the user's current task context, which defaults to `master`.

---

## Standard Development Workflow Process

### Simple Workflow (Default Starting Point)

For new projects or when users are getting started, operate within the `master` tag context:

-   Start new projects by running `initialize_project` tool / `task-master init` or `parse_prd` / `task-master parse-prd --input='<prd-file.txt>'` (see @`taskmaster.md`) to generate initial tasks.json with tagged structure
-   Configure rule sets during initialization with `--rules` flag (e.g., `task-master init --rules <AGENT>,windsurf`) or manage them later with `task-master rules add/remove` commands
-   Begin coding sessions with `get_tasks` / `task-master list` (see @`taskmaster.md`) to see current tasks, status, and IDs
-   Determine the next task to work on using `next_task` / `task-master next` (see @`taskmaster.md`)
-   Analyze task complexity with `analyze_project_complexity` / `task-master analyze-complexity --research` (see @`taskmaster.md`) before breaking down tasks
-   Review complexity report using `complexity_report` / `task-master complexity-report` (see @`taskmaster.md`)
-   Select tasks based on dependencies (all marked 'done'), priority level, and ID order
-   View specific task details using `get_task` / `task-master show <id>` (see @`taskmaster.md`) to understand implementation requirements
-   Break down complex tasks using `expand_task` / `task-master expand --id=<id> --force --research` (see @`taskmaster.md`) with appropriate flags like `--force` (to replace existing subtasks) and `--research`
-   Implement code following task details, dependencies, and project standards
-   Mark completed tasks with `set_task_status` / `task-master set-status --id=<id> --status=done` (see @`taskmaster.md`)
-   Update dependent tasks when implementation differs from original plan using `update` / `task-master update --from=<id> --prompt="..."` or `update_task` / `task-master update-task --id=<id> --prompt="..."` (see @`taskmaster.md`)

---

## Leveling Up: Agent-Led Multi-Context Workflows

While the basic workflow is powerful, your primary opportunity to add value is by identifying when to introduce **Tagged Task Lists**. These patterns are your tools for creating a more organized and efficient development environment for the user, especially if you detect agentic or parallel development happening across the same session.

**Critical Principle**: Most users should never see a difference in their experience. Only introduce advanced workflows when you detect clear indicators that the project has evolved beyond simple task management.

### When to Introduce Tags: Your Decision Patterns

Here are the patterns to look for. When you detect one, you should propose the corresponding workflow to the user.

#### Pattern 1: Simple Git Feature Branching
This is the most common and direct use case for tags.

- **Trigger**: The user creates a new git branch (e.g., `git checkout -b feature/user-auth`).
- **Your Action**: Propose creating a new tag that mirrors the branch name to isolate the feature's tasks from `master`.
- **Your Suggested Prompt**: *"I see you've created a new branch named 'feature/user-auth'. To keep all related tasks neatly organized and separate from your main list, I can create a corresponding task tag for you. This helps prevent merge conflicts in your `tasks.json` file later. Shall I create the 'feature-user-auth' tag?"*
- **Tool to Use**: `task-master add-tag --from-branch`

#### Pattern 2: Team Collaboration
- **Trigger**: The user mentions working with teammates (e.g., "My teammate Alice is handling the database schema," or "I need to review Bob's work on the API.").
- **Your Action**: Suggest creating a separate tag for the user's work to prevent conflicts with shared master context.
- **Your Suggested Prompt**: *"Since you're working with Alice, I can create a separate task context for your work to avoid conflicts. This way, Alice can continue working with the master list while you have your own isolated context. When you're ready to merge your work, we can coordinate the tasks back to master. Shall I create a tag for your current work?"*
- **Tool to Use**: `task-master add-tag my-work --copy-from-current --description="My tasks while collaborating with Alice"`

#### Pattern 3: Experiments or Risky Refactors
- **Trigger**: The user wants to try something that might not be kept (e.g., "I want to experiment with switching our state management library," or "Let's refactor the old API module, but I want to keep the current tasks as a reference.").
- **Your Action**: Propose creating a sandboxed tag for the experimental work.
- **Your Suggested Prompt**: *"This sounds like a great experiment. To keep these new tasks separate from our main plan, I can create a temporary 'experiment-zustand' tag for this work. If we decide not to proceed, we can simply delete the tag without affecting the main task list. Sound good?"*
- **Tool to Use**: `task-master add-tag experiment-zustand --description="Exploring Zustand migration"`

#### Pattern 4: Large Feature Initiatives (PRD-Driven)
This is a more structured approach for significant new features or epics.

- **Trigger**: The user describes a large, multi-step feature that would benefit from a formal plan.
- **Your Action**: Propose a comprehensive, PRD-driven workflow.
- **Your Suggested Prompt**: *"This sounds like a significant new feature. To manage this effectively, I suggest we create a dedicated task context for it. Here's the plan: I'll create a new tag called 'feature-xyz', then we can draft a Product Requirements Document (PRD) together to scope the work. Once the PRD is ready, I'll automatically generate all the necessary tasks within that new tag. How does that sound?"*
- **Your Implementation Flow**:
    1.  **Create an empty tag**: `task-master add-tag feature-xyz --description "Tasks for the new XYZ feature"`. You can also start by creating a git branch if applicable, and then create the tag from that branch.
    2.  **Collaborate & Create PRD**: Work with the user to create a detailed PRD file (e.g., `.taskmaster/docs/feature-xyz-prd.txt`).
    3.  **Parse PRD into the new tag**: `task-master parse-prd .taskmaster/docs/feature-xyz-prd.txt --tag feature-xyz`
    4.  **Prepare the new task list**: Follow up by suggesting `analyze-complexity` and `expand-all` for the newly created tasks within the `feature-xyz` tag.

#### Pattern 5: Version-Based Development
Tailor your approach based on the project maturity indicated by tag names.

- **Prototype/MVP Tags** (`prototype`, `mvp`, `poc`, `v0.x`):
  - **Your Approach**: Focus on speed and functionality over perfection
  - **Task Generation**: Create tasks that emphasize "get it working" over "get it perfect"
  - **Complexity Level**: Lower complexity, fewer subtasks, more direct implementation paths
  - **Research Prompts**: Include context like "This is a prototype - prioritize speed and basic functionality over optimization"
  - **Example Prompt Addition**: *"Since this is for the MVP, I'll focus on tasks that get core functionality working quickly rather than over-engineering."*

- **Production/Mature Tags** (`v1.0+`, `production`, `stable`):
  - **Your Approach**: Emphasize robustness, testing, and maintainability
  - **Task Generation**: Include comprehensive error handling, testing, documentation, and optimization
  - **Complexity Level**: Higher complexity, more detailed subtasks, thorough implementation paths
  - **Research Prompts**: Include context like "This is for production - prioritize reliability, performance, and maintainability"
  - **Example Prompt Addition**: *"Since this is for production, I'll ensure tasks include proper error handling, testing, and documentation."*

### Advanced Workflow (Tag-Based & PRD-Driven)

**When to Transition**: Recognize when the project has evolved (or has initiated a project which existing code) beyond simple task management. Look for these indicators:
- User mentions teammates or collaboration needs
- Project has grown to 15+ tasks with mixed priorities
- User creates feature branches or mentions major initiatives
- User initializes Taskmaster on an existing, complex codebase
- User describes large features that would benefit from dedicated planning

**Your Role in Transition**: Guide the user to a more sophisticated workflow that leverages tags for organization and PRDs for comprehensive planning.

#### Master List Strategy (High-Value Focus)
Once you transition to tag-based workflows, the `master` tag should ideally contain only:
- **High-level deliverables** that provide significant business value
- **Major milestones** and epic-level features
- **Critical infrastructure** work that affects the entire project
- **Release-blocking** items

**What NOT to put in master**:
- Detailed implementation subtasks (these go in feature-specific tags' parent tasks)
- Refactoring work (create dedicated tags like `refactor-auth`)
- Experimental features (use `experiment-*` tags)
- Team member-specific tasks (use person-specific tags)

#### PRD-Driven Feature Development

**For New Major Features**:
1. **Identify the Initiative**: When user describes a significant feature
2. **Create Dedicated Tag**: `add_tag feature-[name] --description="[Feature description]"`
3. **Collaborative PRD Creation**: Work with user to create comprehensive PRD in `.taskmaster/docs/feature-[name]-prd.txt`
4. **Parse & Prepare**:
   - `parse_prd .taskmaster/docs/feature-[name]-prd.txt --tag=feature-[name]`
   - `analyze_project_complexity --tag=feature-[name] --research`
   - `expand_all --tag=feature-[name] --research`
5. **Add Master Reference**: Create a high-level task in `master` that references the feature tag

**For Existing Codebase Analysis**:
When users initialize Taskmaster on existing projects:
1. **Codebase Discovery**: Use your native tools for producing deep context about the code base. You may use `research` tool with `--tree` and `--files` to collect up to date information using the existing architecture as context.
2. **Collaborative Assessment**: Work with user to identify improvement areas, technical debt, or new features
3. **Strategic PRD Creation**: Co-author PRDs that include:
   - Current state analysis (based on your codebase research)
   - Proposed improvements or new features
   - Implementation strategy considering existing code
4. **Tag-Based Organization**: Parse PRDs into appropriate tags (`refactor-api`, `feature-dashboard`, `tech-debt`, etc.)
5. **Master List Curation**: Keep only the most valuable initiatives in master

The parse-prd's `--append` flag enables the user to parse multiple PRDs within tags or across tags. PRDs should be focused and the number of tasks they are parsed into should be strategically chosen relative to the PRD's complexity and level of detail.

### Workflow Transition Examples

**Example 1: Simple → Team-Based**
```
User: "Alice is going to help with the API work"
Your Response: "Great! To avoid conflicts, I'll create a separate task context for your work. Alice can continue with the master list while you work in your own context. When you're ready to merge, we can coordinate the tasks back together."
Action: add_tag my-api-work --copy-from-current --description="My API tasks while collaborating with Alice"
```

**Example 2: Simple → PRD-Driven**
```
User: "I want to add a complete user dashboard with analytics, user management, and reporting"
Your Response: "This sounds like a major feature that would benefit from detailed planning. Let me create a dedicated context for this work and we can draft a PRD together to ensure we capture all requirements."
Actions:
1. add_tag feature-dashboard --description="User dashboard with analytics and management"
2. Collaborate on PRD creation
3. parse_prd dashboard-prd.txt --tag=feature-dashboard
4. Add high-level "User Dashboard" task to master
```

**Example 3: Existing Project → Strategic Planning**
```
User: "I just initialized Taskmaster on my existing React app. It's getting messy and I want to improve it."
Your Response: "Let me research your codebase to understand the current architecture, then we can create a strategic plan for improvements."
Actions:
1. research "Current React app architecture and improvement opportunities" --tree --files=src/
2. Collaborate on improvement PRD based on findings
3. Create tags for different improvement areas (refactor-components, improve-state-management, etc.)
4. Keep only major improvement initiatives in master
```

---

## Primary Interaction: MCP Server vs. CLI

Taskmaster offers two primary ways to interact:

1.  **MCP Server (Recommended for Integrated Tools)**:
    - For AI agents and integrated development environments (like <AGENT>), interacting via the **MCP server is the preferred method**.
    - The MCP server exposes Taskmaster functionality through a set of tools (e.g., `get_tasks`, `add_subtask`).
    - This method offers better performance, structured data exchange, and richer error handling compared to CLI parsing.
    - Refer to @`mcp.md` for details on the MCP architecture and available tools.
    - A comprehensive list and description of MCP tools and their corresponding CLI commands can be found in @`taskmaster.md`.
    - **Restart the MCP server** if core logic in `scripts/modules` or MCP tool/direct function definitions change.
    - **Note**: MCP tools fully support tagged task lists with complete tag management capabilities.

2.  **`task-master` CLI (For Users & Fallback)**:
    - The global `task-master` command provides a user-friendly interface for direct terminal interaction.
    - It can also serve as a fallback if the MCP server is inaccessible or a specific function isn't exposed via MCP.
    - Install globally with `npm install -g task-master-ai` or use locally via `npx task-master-ai ...`.
    - The CLI commands often mirror the MCP tools (e.g., `task-master list` corresponds to `get_tasks`).
    - Refer to @`taskmaster.md` for a detailed command reference.
    - **Tagged Task Lists**: CLI fully supports the new tagged system with seamless migration.

## How the Tag System Works (For Your Reference)

- **Data Structure**: Tasks are organized into separate contexts (tags) like "master", "feature-branch", or "v2.0".
- **Silent Migration**: Existing projects automatically migrate to use a "master" tag with zero disruption.
- **Context Isolation**: Tasks in different tags are completely separate. Changes in one tag do not affect any other tag.
- **Manual Control**: The user is always in control. There is no automatic switching. You facilitate switching by using `use-tag <name>`.
- **Full CLI & MCP Support**: All tag management commands are available through both the CLI and MCP tools for you to use. Refer to @`taskmaster.md` for a full command list.

---

## Task Complexity Analysis

-   Run `analyze_project_complexity` / `task-master analyze-complexity --research` (see @`taskmaster.md`) for comprehensive analysis
-   Review complexity report via `complexity_report` / `task-master complexity-report` (see @`taskmaster.md`) for a formatted, readable version.
-   Focus on tasks with highest complexity scores (8-10) for detailed breakdown
-   Use analysis results to determine appropriate subtask allocation
-   Note that reports are automatically used by the `expand_task` tool/command

## Task Breakdown Process

-   Use `expand_task` / `task-master expand --id=<id>`. It automatically uses the complexity report if found, otherwise generates default number of subtasks.
-   Use `--num=<number>` to specify an explicit number of subtasks, overriding defaults or complexity report recommendations.
-   Add `--research` flag to leverage Perplexity AI for research-backed expansion.
-   Add `--force` flag to clear existing subtasks before generating new ones (default is to append).
-   Use `--prompt="<context>"` to provide additional context when needed.
-   Review and adjust generated subtasks as necessary.
-   Use `expand_all` tool or `task-master expand --all` to expand multiple pending tasks at once, respecting flags like `--force` and `--research`.
-   If subtasks need complete replacement (regardless of the `--force` flag on `expand`), clear them first with `clear_subtasks` / `task-master clear-subtasks --id=<id>`.

## Implementation Drift Handling

-   When implementation differs significantly from planned approach
-   When future tasks need modification due to current implementation choices
-   When new dependencies or requirements emerge
-   Use `update` / `task-master update --from=<futureTaskId> --prompt='<explanation>\nUpdate context...' --research` to update multiple future tasks.
-   Use `update_task` / `task-master update-task --id=<taskId> --prompt='<explanation>\nUpdate context...' --research` to update a single specific task.

## Task Status Management

-   Use 'pending' for tasks ready to be worked on
-   Use 'done' for completed and verified tasks
-   Use 'deferred' for postponed tasks
-   Add custom status values as needed for project-specific workflows

## Task Structure Fields

- **id**: Unique identifier for the task (Example: `1`, `1.1`)
- **title**: Brief, descriptive title (Example: `"Initialize Repo"`)
- **description**: Concise summary of what the task involves (Example: `"Create a new repository, set up initial structure."`)
- **status**: Current state of the task (Example: `"pending"`, `"done"`, `"deferred"`)
- **dependencies**: IDs of prerequisite tasks (Example: `[1, 2.1]`)
    - Dependencies are displayed with status indicators (✅ for completed, ⏱️ for pending)
    - This helps quickly identify which prerequisite tasks are blocking work
- **priority**: Importance level (Example: `"high"`, `"medium"`, `"low"`)
- **details**: In-depth implementation instructions (Example: `"Use GitHub client ID/secret, handle callback, set session token."`)
- **testStrategy**: Verification approach (Example: `"Deploy and call endpoint to confirm 'Hello World' response."`)
- **subtasks**: List of smaller, more specific tasks (Example: `[{"id": 1, "title": "Configure OAuth", ...}]`)
- Refer to task structure details (previously linked to `tasks.md`).

## Configuration Management (Updated)

Taskmaster configuration is managed through two main mechanisms:

1.  **`.taskmaster/config.json` File (Primary):**
    *   Located in the project root directory.
    *   Stores most configuration settings: AI model selections (main, research, fallback), parameters (max tokens, temperature), logging level, default subtasks/priority, project name, etc.
    *   **Tagged System Settings**: Includes `global.defaultTag` (defaults to "master") and `tags` section for tag management configuration.
    *   **Managed via `task-master models --setup` command.** Do not edit manually unless you know what you are doing.
    *   **View/Set specific models via `task-master models` command or `models` MCP tool.**
    *   Created automatically when you run `task-master models --setup` for the first time or during tagged system migration.

2.  **Environment Variables (`.env` / `mcp.json`):**
    *   Used **only** for sensitive API keys and specific endpoint URLs.
    *   Place API keys (one per provider) in a `.env` file in the project root for CLI usage.
    *   For MCP/<AGENT> integration, configure these keys in the `env` section of `.<AGENT>/mcp.json`.
    *   Available keys/variables: See `assets/env.example` or the Configuration section in the command reference (previously linked to `taskmaster.md`).

3.  **`.taskmaster/state.json` File (Tagged System State):**
    *   Tracks current tag context and migration status.
    *   Automatically created during tagged system migration.
    *   Contains: `currentTag`, `lastSwitched`, `migrationNoticeShown`.

**Important:** Non-API key settings (like model selections, `MAX_TOKENS`, `TASKMASTER_LOG_LEVEL`) are **no longer configured via environment variables**. Use the `task-master models` command (or `--setup` for interactive configuration) or the `models` MCP tool.
**If AI commands FAIL in MCP** verify that the API key for the selected provider is present in the `env` section of `.<AGENT>/mcp.json`.
**If AI commands FAIL in CLI** verify that the API key for the selected provider is present in the `.env` file in the root of the project.

## Rules Management

Taskmaster supports multiple AI coding assistant rule sets that can be configured during project initialization or managed afterward:

- **Available Profiles**: Claude Code, <AGENT>, Codex, <AGENT>, Roo Code, Trae, Windsurf (claude, <AGENT>, codex, <AGENT>, roo, trae, windsurf)
- **During Initialization**: Use `task-master init --rules <AGENT>,windsurf` to specify which rule sets to include
- **After Initialization**: Use `task-master rules add <profiles>` or `task-master rules remove <profiles>` to manage rule sets
- **Interactive Setup**: Use `task-master rules setup` to launch an interactive prompt for selecting rule profiles
- **Default Behavior**: If no `--rules` flag is specified during initialization, all available rule profiles are included
- **Rule Structure**: Each profile creates its own directory (e.g., `.<AGENT>/rules`, `.roo/rules`) with appropriate configuration files

## Determining the Next Task

- Run `next_task` / `task-master next` to show the next task to work on.
- The command identifies tasks with all dependencies satisfied
- Tasks are prioritized by priority level, dependency count, and ID
- The command shows comprehensive task information including:
    - Basic task details and description
    - Implementation details
    - Subtasks (if they exist)
    - Contextual suggested actions
- Recommended before starting any new development work
- Respects your project's dependency structure
- Ensures tasks are completed in the appropriate sequence
- Provides ready-to-use commands for common task actions

## Viewing Specific Task Details

- Run `get_task` / `task-master show <id>` to view a specific task.
- Use dot notation for subtasks: `task-master show 1.2` (shows subtask 2 of task 1)
- Displays comprehensive information similar to the next command, but for a specific task
- For parent tasks, shows all subtasks and their current status
- For subtasks, shows parent task information and relationship
- Provides contextual suggested actions appropriate for the specific task
- Useful for examining task details before implementation or checking status

## Managing Task Dependencies

- Use `add_dependency` / `task-master add-dependency --id=<id> --depends-on=<id>` to add a dependency.
- Use `remove_dependency` / `task-master remove-dependency --id=<id> --depends-on=<id>` to remove a dependency.
- The system prevents circular dependencies and duplicate dependency entries
- Dependencies are checked for existence before being added or removed
- Task files are automatically regenerated after dependency changes
- Dependencies are visualized with status indicators in task listings and files

## Task Reorganization

- Use `move_task` / `task-master move --from=<id> --to=<id>` to move tasks or subtasks within the hierarchy
- This command supports several use cases:
  - Moving a standalone task to become a subtask (e.g., `--from=5 --to=7`)
  - Moving a subtask to become a standalone task (e.g., `--from=5.2 --to=7`)
  - Moving a subtask to a different parent (e.g., `--from=5.2 --to=7.3`)
  - Reordering subtasks within the same parent (e.g., `--from=5.2 --to=5.4`)
  - Moving a task to a new, non-existent ID position (e.g., `--from=5 --to=25`)
  - Moving multiple tasks at once using comma-separated IDs (e.g., `--from=10,11,12 --to=16,17,18`)
- The system includes validation to prevent data loss:
  - Allows moving to non-existent IDs by creating placeholder tasks
  - Prevents moving to existing task IDs that have content (to avoid overwriting)
  - Validates source tasks exist before attempting to move them
- The system maintains proper parent-child relationships and dependency integrity
- Task files are automatically regenerated after the move operation
- This provides greater flexibility in organizing and refining your task structure as project understanding evolves
- This is especially useful when dealing with potential merge conflicts arising from teams creating tasks on separate branches. Solve these conflicts very easily by moving your tasks and keeping theirs.

## Iterative Subtask Implementation

Once a task has been broken down into subtasks using `expand_task` or similar methods, follow this iterative process for implementation:

1.  **Understand the Goal (Preparation):**
    *   Use `get_task` / `task-master show <subtaskId>` (see @`taskmaster.md`) to thoroughly understand the specific goals and requirements of the subtask.

2.  **Initial Exploration & Planning (Iteration 1):**
    *   This is the first attempt at creating a concrete implementation plan.
    *   Explore the codebase to identify the precise files, functions, and even specific lines of code that will need modification.
    *   Determine the intended code changes (diffs) and their locations.
    *   Gather *all* relevant details from this exploration phase.

3.  **Log the Plan:**
    *   Run `update_subtask` / `task-master update-subtask --id=<subtaskId> --prompt='<detailed plan>'`.
    *   Provide the *complete and detailed* findings from the exploration phase in the prompt. Include file paths, line numbers, proposed diffs, reasoning, and any potential challenges identified. Do not omit details. The goal is to create a rich, timestamped log within the subtask's `details`.

4.  **Verify the Plan:**
    *   Run `get_task` / `task-master show <subtaskId>` again to confirm that the detailed implementation plan has been successfully appended to the subtask's details.

5.  **Begin Implementation:**
    *   Set the subtask status using `set_task_status` / `task-master set-status --id=<subtaskId> --status=in-progress`.
    *   Start coding based on the logged plan.

6.  **Refine and Log Progress (Iteration 2+):**
    *   As implementation progresses, you will encounter challenges, discover nuances, or confirm successful approaches.
    *   **Before appending new information**: Briefly review the *existing* details logged in the subtask (using `get_task` or recalling from context) to ensure the update adds fresh insights and avoids redundancy.
    *   **Regularly** use `update_subtask` / `task-master update-subtask --id=<subtaskId> --prompt='<update details>\n- What worked...\n- What didn't work...'` to append new findings.
    *   **Crucially, log:**
        *   What worked ("fundamental truths" discovered).
        *   What didn't work and why (to avoid repeating mistakes).
        *   Specific code snippets or configurations that were successful.
        *   Decisions made, especially if confirmed with user input.
        *   Any deviations from the initial plan and the reasoning.
    *   The objective is to continuously enrich the subtask's details, creating a log of the implementation journey that helps the AI (and human developers) learn, adapt, and avoid repeating errors.

7.  **Review & Update Rules (Post-Implementation):**
    *   Once the implementation for the subtask is functionally complete, review all code changes and the relevant chat history.
    *   Identify any new or modified code patterns, conventions, or best practices established during the implementation.
    *   Create new or update existing rules following internal guidelines (previously linked to `cursor_rules.md` and `self_improve.md`).

8.  **Mark Task Complete:**
    *   After verifying the implementation and updating any necessary rules, mark the subtask as completed: `set_task_status` / `task-master set-status --id=<subtaskId> --status=done`.

9.  **Commit Changes (If using Git):**
    *   Stage the relevant code changes and any updated/new rule files (`git add .`).
    *   Craft a comprehensive Git commit message summarizing the work done for the subtask, including both code implementation and any rule adjustments.
    *   Execute the commit command directly in the terminal (e.g., `git commit -m 'feat(module): Implement feature X for subtask <subtaskId>\n\n- Details about changes...\n- Updated rule Y for pattern Z'`).
    *   Consider if a Changeset is needed according to internal versioning guidelines (previously linked to `changeset.md`). If so, run `npm run changeset`, stage the generated file, and amend the commit or create a new one.

10. **Proceed to Next Subtask:**
    *   Identify the next subtask (e.g., using `next_task` / `task-master next`).

## Code Analysis & Refactoring Techniques

- **Top-Level Function Search**:
    - Useful for understanding module structure or planning refactors.
    - Use grep/ripgrep to find exported functions/constants:
      `rg "export (async function|function|const) \w+"` or similar patterns.
    - Can help compare functions between files during migrations or identify potential naming conflicts.

---
*This workflow provides a general guideline. Adapt it based on your specific project needs and team practices.*
```

.rules/rules.md
```
---
description: Guidelines for creating and maintaining AGENT rules to ensure consistency and effectiveness.
---

- **Required Rule Structure:**
  ```markdown
  ---
  description: Clear, one-line description of what the rule enforces
  ---

  - **Main Points in Bold**
    - Sub-points with details
    - Examples and explanations
  ```

- **File References:**
  - Use `[filename](md:path/to/file)` ([filename](md:filename)) to reference files
  - Example: [prisma.md](.ruler/prisma.md) for rule references
  - Example: [schema.prisma](md:prisma/schema.prisma) for code references

- **Code Examples:**
  - Use language-specific code blocks
  ```typescript
  // ✅ DO: Show good examples
  const goodExample = true;

  // ❌ DON'T: Show anti-patterns
  const badExample = false;
  ```

- **Rule Content Guidelines:**
  - Start with high-level overview
  - Include specific, actionable requirements
  - Show examples of correct implementation
  - Reference existing code when possible
  - Keep rules DRY by referencing other rules

- **Rule Maintenance:**
  - Update rules when new patterns emerge
  - Add examples from actual codebase
  - Remove outdated patterns
  - Cross-reference related rules

- **Best Practices:**
  - Use bullet points for clarity
  - Keep descriptions concise
  - Include both DO and DON'T examples
  - Reference actual code over theoretical examples
  - Use consistent formatting across rules
```

.rules/self_improve.md
```
---
description: Guidelines for continuously improving  rules based on emerging code patterns and best practices.
---

- **Rule Improvement Triggers:**
  - New code patterns not covered by existing rules
  - Repeated similar implementations across files
  - Common error patterns that could be prevented
  - New libraries or tools being used consistently
  - Emerging best practices in the codebase

- **Analysis Process:**
  - Compare new code with existing rules
  - Identify patterns that should be standardized
  - Look for references to external documentation
  - Check for consistent error handling patterns
  - Monitor test patterns and coverage

- **Rule Updates:**
  - **Add New Rules When:**
    - A new technology/pattern is used in 3+ files
    - Common bugs could be prevented by a rule
    - Code reviews repeatedly mention the same feedback
    - New security or performance patterns emerge

  - **Modify Existing Rules When:**
    - Better examples exist in the codebase
    - Additional edge cases are discovered
    - Related rules have been updated
    - Implementation details have changed

- **Example Pattern Recognition:**
  ```typescript
  // If you see repeated patterns like:
  const data = await prisma.user.findMany({
    select: { id: true, email: true },
    where: { status: 'ACTIVE' }
  });

  // Consider adding to [prisma.md](.ruler/prisma.md):
  // - Standard select fields
  // - Common where conditions
  // - Performance optimization patterns
  ```

- **Rule Quality Checks:**
  - Rules should be actionable and specific
  - Examples should come from actual code
  - References should be up to date
  - Patterns should be consistently enforced

- **Continuous Improvement:**
  - Monitor code review comments
  - Track common development questions
  - Update rules after major refactors
  - Add links to relevant documentation
  - Cross-reference related rules

- **Rule Deprecation:**
  - Mark outdated patterns as deprecated
  - Remove rules that no longer apply
  - Update references to deprecated rules
  - Document migration paths for old patterns

- **Documentation Updates:**
  - Keep examples synchronized with code
  - Update references to external docs
  - Maintain links between related rules
  - Document breaking changes
Follow [.ruler.md](.ruler/rules.md) for proper rule formatting and structure.
```

.rules/taskmaster.md
```
---
description: Comprehensive reference for Taskmaster MCP tools and CLI commands.
---

# Taskmaster Tool & Command Reference

This document provides a detailed reference for interacting with Taskmaster, covering both the recommended MCP tools, suitable for integrations like <AGENT>, and the corresponding `task-master` CLI commands, designed for direct user interaction or fallback.

**Note:** For interacting with Taskmaster programmatically or via integrated tools, using the **MCP tools is strongly recommended** due to better performance, structured data, and error handling. The CLI commands serve as a user-friendly alternative and fallback.

**Important:** Several MCP tools involve AI processing... The AI-powered tools include `parse_prd`, `analyze_project_complexity`, `update_subtask`, `update_task`, `update`, `expand_all`, `expand_task`, and `add_task`.

**🏷️ Tagged Task Lists System:** Task Master now supports **tagged task lists** for multi-context task management. This allows you to maintain separate, isolated lists of tasks for different features, branches, or experiments. Existing projects are seamlessly migrated to use a default "master" tag. Most commands now support a `--tag <name>` flag to specify which context to operate on. If omitted, commands use the currently active tag.

---

## Initialization & Setup

### 1. Initialize Project (`init`)

*   **MCP Tool:** `initialize_project`
*   **CLI Command:** `task-master init [options]`
*   **Description:** `Set up the basic Taskmaster file structure and configuration in the current directory for a new project.`
*   **Key CLI Options:**
    *   `--name <name>`: `Set the name for your project in Taskmaster's configuration.`
    *   `--description <text>`: `Provide a brief description for your project.`
    *   `--version <version>`: `Set the initial version for your project, e.g., '0.1.0'.`
    *   `-y, --yes`: `Initialize Taskmaster quickly using default settings without interactive prompts.`
*   **Usage:** Run this once at the beginning of a new project.
*   **MCP Variant Description:** `Set up the basic Taskmaster file structure and configuration in the current directory for a new project by running the 'task-master init' command.`
*   **Key MCP Parameters/Options:**
    *   `projectName`: `Set the name for your project.` (CLI: `--name <name>`)
    *   `projectDescription`: `Provide a brief description for your project.` (CLI: `--description <text>`)
    *   `projectVersion`: `Set the initial version for your project, e.g., '0.1.0'.` (CLI: `--version <version>`)
    *   `authorName`: `Author name.` (CLI: `--author <author>`)
    *   `skipInstall`: `Skip installing dependencies. Default is false.` (CLI: `--skip-install`)
    *   `addAliases`: `Add shell aliases tm, taskmaster, hamster, and ham. Default is false.` (CLI: `--aliases`)
    *   `yes`: `Skip prompts and use defaults/provided arguments. Default is false.` (CLI: `-y, --yes`)
*   **Usage:** Run this once at the beginning of a new project, typically via an integrated tool like <AGENT>. Operates on the current working directory of the MCP server.
*   **Important:** Once complete, you *MUST* parse a prd in order to generate tasks. There will be no tasks files until then. The next step after initializing should be to create a PRD using the example PRD in .taskmaster/templates/example_prd.txt.
*   **Tagging:** Use the `--tag` option to parse the PRD into a specific, non-default tag context. If the tag doesn't exist, it will be created automatically. Example: `task-master parse-prd spec.txt --tag=new-feature`.

### 2. Parse PRD (`parse_prd`)

*   **MCP Tool:** `parse_prd`
*   **CLI Command:** `task-master parse-prd [file] [options]`
*   **Description:** `Parse a Product Requirements Document, PRD, or text file with Taskmaster to automatically generate an initial set of tasks in tasks.json.`
*   **Key Parameters/Options:**
    *   `input`: `Path to your PRD or requirements text file that Taskmaster should parse for tasks.` (CLI: `[file]` positional or `-i, --input <file>`)
    *   `output`: `Specify where Taskmaster should save the generated 'tasks.json' file. Defaults to '.taskmaster/tasks/tasks.json'.` (CLI: `-o, --output <file>`)
    *   `numTasks`: `Approximate number of top-level tasks Taskmaster should aim to generate from the document.` (CLI: `-n, --num-tasks <number>`)
    *   `force`: `Use this to allow Taskmaster to overwrite an existing 'tasks.json' without asking for confirmation.` (CLI: `-f, --force`)
*   **Usage:** Useful for bootstrapping a project from an existing requirements document.
*   **Notes:** Task Master will strictly adhere to any specific requirements mentioned in the PRD, such as libraries, database schemas, frameworks, tech stacks, etc., while filling in any gaps where the PRD isn't fully specified. Tasks are designed to provide the most direct implementation path while avoiding over-engineering.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress. If the user does not have a PRD, suggest discussing their idea and then use the example PRD in `.taskmaster/templates/example_prd.txt` as a template for creating the PRD based on their idea, for use with `parse-prd`.

---

## AI Model Configuration

### 2. Manage Models (`models`)
*   **MCP Tool:** `models`
*   **CLI Command:** `task-master models [options]`
*   **Description:** `View the current AI model configuration or set specific models for different roles (main, research, fallback). Allows setting custom model IDs for Ollama and OpenRouter.`
*   **Key MCP Parameters/Options:**
    *   `setMain <model_id>`: `Set the primary model ID for task generation/updates.` (CLI: `--set-main <model_id>`)
    *   `setResearch <model_id>`: `Set the model ID for research-backed operations.` (CLI: `--set-research <model_id>`)
    *   `setFallback <model_id>`: `Set the model ID to use if the primary fails.` (CLI: `--set-fallback <model_id>`)
    *   `ollama <boolean>`: `Indicates the set model ID is a custom Ollama model.` (CLI: `--ollama`)
    *   `openrouter <boolean>`: `Indicates the set model ID is a custom OpenRouter model.` (CLI: `--openrouter`)
    *   `listAvailableModels <boolean>`: `If true, lists available models not currently assigned to a role.` (CLI: No direct equivalent; CLI lists available automatically)
    *   `projectRoot <string>`: `Optional. Absolute path to the project root directory.` (CLI: Determined automatically)
*   **Key CLI Options:**
    *   `--set-main <model_id>`: `Set the primary model.`
    *   `--set-research <model_id>`: `Set the research model.`
    *   `--set-fallback <model_id>`: `Set the fallback model.`
    *   `--ollama`: `Specify that the provided model ID is for Ollama (use with --set-*).`
    *   `--openrouter`: `Specify that the provided model ID is for OpenRouter (use with --set-*). Validates against OpenRouter API.`
    *   `--bedrock`: `Specify that the provided model ID is for AWS Bedrock (use with --set-*).`
    *   `--setup`: `Run interactive setup to configure models, including custom Ollama/OpenRouter IDs.`
*   **Usage (MCP):** Call without set flags to get current config. Use `setMain`, `setResearch`, or `setFallback` with a valid model ID to update the configuration. Use `listAvailableModels: true` to get a list of unassigned models. To set a custom model, provide the model ID and set `ollama: true` or `openrouter: true`.
*   **Usage (CLI):** Run without flags to view current configuration and available models. Use set flags to update specific roles. Use `--setup` for guided configuration, including custom models. To set a custom model via flags, use `--set-<role>=<model_id>` along with either `--ollama` or `--openrouter`.
*   **Notes:** Configuration is stored in `.taskmaster/config.json` in the project root. This command/tool modifies that file. Use `listAvailableModels` or `task-master models` to see internally supported models. OpenRouter custom models are validated against their live API. Ollama custom models are not validated live.
*   **API note:** API keys for selected AI providers (based on their model) need to exist in the mcp.json file to be accessible in MCP context. The API keys must be present in the local .env file for the CLI to be able to read them.
*   **Model costs:** The costs in supported models are expressed in dollars. An input/output value of 3 is $3.00. A value of 0.8 is $0.80.
*   **Warning:** DO NOT MANUALLY EDIT THE .taskmaster/config.json FILE. Use the included commands either in the MCP or CLI format as needed. Always prioritize MCP tools when available and use the CLI as a fallback.

---

## Task Listing & Viewing

### 3. Get Tasks (`get_tasks`)

*   **MCP Tool:** `get_tasks`
*   **CLI Command:** `task-master list [options]`
*   **Description:** `List your Taskmaster tasks, optionally filtering by status and showing subtasks.`
*   **Key Parameters/Options:**
    *   `status`: `Show only Taskmaster tasks matching this status (or multiple statuses, comma-separated), e.g., 'pending' or 'done,in-progress'.` (CLI: `-s, --status <status>`)
    *   `withSubtasks`: `Include subtasks indented under their parent tasks in the list.` (CLI: `--with-subtasks`)
    *   `tag`: `Specify which tag context to list tasks from. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
    *   `watch`: `Watch for changes and auto-refresh the list in real-time. Works with file storage (fs.watch) and API storage (Supabase Realtime).` (CLI: `-w, --watch`)
*   **Usage:** Get an overview of the project status, often used at the start of a work session. Use `--watch` to keep the list live-updating as tasks change.

### 4. Get Next Task (`next_task`)

*   **MCP Tool:** `next_task`
*   **CLI Command:** `task-master next [options]`
*   **Description:** `Ask Taskmaster to show the next available task you can work on, based on status and completed dependencies.`
*   **Key Parameters/Options:**
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
    *   `tag`: `Specify which tag context to use. Defaults to the current active tag.` (CLI: `--tag <name>`)
*   **Usage:** Identify what to work on next according to the plan.

### 5. Get Task Details (`get_task`)

*   **MCP Tool:** `get_task`
*   **CLI Command:** `task-master show [id] [options]`
*   **Description:** `Display detailed information for one or more specific Taskmaster tasks or subtasks by ID.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID of the Taskmaster task (e.g., '15'), subtask (e.g., '15.2'), or a comma-separated list of IDs ('1,5,10.2') you want to view.` (CLI: `[id]` positional or `-i, --id <id>`)
    *   `tag`: `Specify which tag context to get the task(s) from. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Understand the full details for a specific task. When multiple IDs are provided, a summary table is shown.
*   **CRITICAL INFORMATION** If you need to collect information from multiple tasks, use comma-separated IDs (i.e. 1,2,3) to receive an array of tasks. Do not needlessly get tasks one at a time if you need to get many as that is wasteful.

---

## Task Creation & Modification

### 6. Add Task (`add_task`)

*   **MCP Tool:** `add_task`
*   **CLI Command:** `task-master add-task [options]`
*   **Description:** `Add a new task to Taskmaster by describing it; AI will structure it.`
*   **Key Parameters/Options:**
    *   `prompt`: `Required. Describe the new task you want Taskmaster to create, e.g., "Implement user authentication using JWT".` (CLI: `-p, --prompt <text>`)
    *   `dependencies`: `Specify the IDs of any Taskmaster tasks that must be completed before this new one can start, e.g., '12,14'.` (CLI: `-d, --dependencies <ids>`)
    *   `priority`: `Set the priority for the new task: 'high', 'medium', or 'low'. Default is 'medium'.` (CLI: `--priority <priority>`)
    *   `research`: `Enable Taskmaster to use the research role for potentially more informed task creation.` (CLI: `-r, --research`)
    *   `tag`: `Specify which tag context to add the task to. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Quickly add newly identified tasks during development.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 7. Add Subtask (`add_subtask`)

*   **MCP Tool:** `add_subtask`
*   **CLI Command:** `task-master add-subtask [options]`
*   **Description:** `Add a new subtask to a Taskmaster parent task, or convert an existing task into a subtask.`
*   **Key Parameters/Options:**
    *   `id` / `parent`: `Required. The ID of the Taskmaster task that will be the parent.` (MCP: `id`, CLI: `-p, --parent <id>`)
    *   `taskId`: `Use this if you want to convert an existing top-level Taskmaster task into a subtask of the specified parent.` (CLI: `-i, --task-id <id>`)
    *   `title`: `Required if not using taskId. The title for the new subtask Taskmaster should create.` (CLI: `-t, --title <title>`)
    *   `description`: `A brief description for the new subtask.` (CLI: `-d, --description <text>`)
    *   `details`: `Provide implementation notes or details for the new subtask.` (CLI: `--details <text>`)
    *   `dependencies`: `Specify IDs of other tasks or subtasks, e.g., '15' or '16.1', that must be done before this new subtask.` (CLI: `--dependencies <ids>`)
    *   `status`: `Set the initial status for the new subtask. Default is 'pending'.` (CLI: `-s, --status <status>`)
    *   `generate`: `Enable Taskmaster to regenerate markdown task files after adding the subtask.` (CLI: `--generate`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Break down tasks manually or reorganize existing tasks.

### 8. Update Tasks (`update`)

*   **MCP Tool:** `update`
*   **CLI Command:** `task-master update [options]`
*   **Description:** `Update multiple upcoming tasks in Taskmaster based on new context or changes, starting from a specific task ID.`
*   **Key Parameters/Options:**
    *   `from`: `Required. The ID of the first task Taskmaster should update. All tasks with this ID or higher that are not 'done' will be considered.` (CLI: `--from <id>`)
    *   `prompt`: `Required. Explain the change or new context for Taskmaster to apply to the tasks, e.g., "We are now using React Query instead of Redux Toolkit for data fetching".` (CLI: `-p, --prompt <text>`)
    *   `research`: `Enable Taskmaster to use the research role for more informed updates. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Handle significant implementation changes or pivots that affect multiple future tasks. Example CLI: `task-master update --from='18' --prompt='Switching to React Query.\nNeed to refactor data fetching...'`
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 9. Update Task (`update_task`)

*   **MCP Tool:** `update_task`
*   **CLI Command:** `task-master update-task [options]`
*   **Description:** `Modify a specific Taskmaster task by ID, incorporating new information or changes. By default, this replaces the existing task details.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The specific ID of the Taskmaster task, e.g., '15', you want to update.` (CLI: `-i, --id <id>`)
    *   `prompt`: `Required. Explain the specific changes or provide the new information Taskmaster should incorporate into this task.` (CLI: `-p, --prompt <text>`)
    *   `append`: `If true, appends the prompt content to the task's details with a timestamp, rather than replacing them. Behaves like update-subtask.` (CLI: `--append`)
    *   `research`: `Enable Taskmaster to use the research role for more informed updates. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `tag`: `Specify which tag context the task belongs to. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Refine a specific task based on new understanding. Use `--append` to log progress without creating subtasks.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 10. Update Subtask (`update_subtask`)

*   **MCP Tool:** `update_subtask`
*   **CLI Command:** `task-master update-subtask [options]`
*   **Description:** `Append timestamped notes or details to a specific Taskmaster subtask without overwriting existing content. Intended for iterative implementation logging.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID of the Taskmaster subtask, e.g., '5.2', to update with new information.` (CLI: `-i, --id <id>`)
    *   `prompt`: `Required. The information, findings, or progress notes to append to the subtask's details with a timestamp.` (CLI: `-p, --prompt <text>`)
    *   `research`: `Enable Taskmaster to use the research role for more informed updates. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `tag`: `Specify which tag context the subtask belongs to. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Log implementation progress, findings, and discoveries during subtask development. Each update is timestamped and appended to preserve the implementation journey.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 11. Set Task Status (`set_task_status`)

*   **MCP Tool:** `set_task_status`
*   **CLI Command:** `task-master set-status [options]`
*   **Description:** `Update the status of one or more Taskmaster tasks or subtasks, e.g., 'pending', 'in-progress', 'done'.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID(s) of the Taskmaster task(s) or subtask(s), e.g., '15', '15.2', or '16,17.1', to update.` (CLI: `-i, --id <id>`)
    *   `status`: `Required. The new status to set, e.g., 'done', 'pending', 'in-progress', 'review', 'cancelled'.` (CLI: `-s, --status <status>`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Mark progress as tasks move through the development cycle.

### 12. Remove Task (`remove_task`)

*   **MCP Tool:** `remove_task`
*   **CLI Command:** `task-master remove-task [options]`
*   **Description:** `Permanently remove a task or subtask from the Taskmaster tasks list.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID of the Taskmaster task, e.g., '5', or subtask, e.g., '5.2', to permanently remove.` (CLI: `-i, --id <id>`)
    *   `yes`: `Skip the confirmation prompt and immediately delete the task.` (CLI: `-y, --yes`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Permanently delete tasks or subtasks that are no longer needed in the project.
*   **Notes:** Use with caution as this operation cannot be undone. Consider using 'blocked', 'cancelled', or 'deferred' status instead if you just want to exclude a task from active planning but keep it for reference. The command automatically cleans up dependency references in other tasks.

---

## Task Structure & Breakdown

### 13. Expand Task (`expand_task`)

*   **MCP Tool:** `expand_task`
*   **CLI Command:** `task-master expand [options]`
*   **Description:** `Use Taskmaster's AI to break down a complex task into smaller, manageable subtasks. Appends subtasks by default.`
*   **Key Parameters/Options:**
    *   `id`: `The ID of the specific Taskmaster task you want to break down into subtasks.` (CLI: `-i, --id <id>`)
    *   `num`: `Optional: Suggests how many subtasks Taskmaster should aim to create. Uses complexity analysis/defaults otherwise.` (CLI: `-n, --num <number>`)
    *   `research`: `Enable Taskmaster to use the research role for more informed subtask generation. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `prompt`: `Optional: Provide extra context or specific instructions to Taskmaster for generating the subtasks.` (CLI: `-p, --prompt <text>`)
    *   `force`: `Optional: If true, clear existing subtasks before generating new ones. Default is false (append).` (CLI: `--force`)
    *   `tag`: `Specify which tag context the task belongs to. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Generate a detailed implementation plan for a complex task before starting coding. Automatically uses complexity report recommendations if available and `num` is not specified.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 14. Expand All Tasks (`expand_all`)

*   **MCP Tool:** `expand_all`
*   **CLI Command:** `task-master expand --all [options]` (Note: CLI uses the `expand` command with the `--all` flag)
*   **Description:** `Tell Taskmaster to automatically expand all eligible pending/in-progress tasks based on complexity analysis or defaults. Appends subtasks by default.`
*   **Key Parameters/Options:**
    *   `num`: `Optional: Suggests how many subtasks Taskmaster should aim to create per task.` (CLI: `-n, --num <number>`)
    *   `research`: `Enable research role for more informed subtask generation. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `prompt`: `Optional: Provide extra context for Taskmaster to apply generally during expansion.` (CLI: `-p, --prompt <text>`)
    *   `force`: `Optional: If true, clear existing subtasks before generating new ones for each eligible task. Default is false (append).` (CLI: `--force`)
    *   `tag`: `Specify which tag context to expand. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Useful after initial task generation or complexity analysis to break down multiple tasks at once.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 15. Clear Subtasks (`clear_subtasks`)

*   **MCP Tool:** `clear_subtasks`
*   **CLI Command:** `task-master clear-subtasks [options]`
*   **Description:** `Remove all subtasks from one or more specified Taskmaster parent tasks.`
*   **Key Parameters/Options:**
    *   `id`: `The ID(s) of the Taskmaster parent task(s) whose subtasks you want to remove, e.g., '15' or '16,18'. Required unless using 'all'.` (CLI: `-i, --id <ids>`)
    *   `all`: `Tell Taskmaster to remove subtasks from all parent tasks.` (CLI: `--all`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Used before regenerating subtasks with `expand_task` if the previous breakdown needs replacement.

### 16. Remove Subtask (`remove_subtask`)

*   **MCP Tool:** `remove_subtask`
*   **CLI Command:** `task-master remove-subtask [options]`
*   **Description:** `Remove a subtask from its Taskmaster parent, optionally converting it into a standalone task.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID(s) of the Taskmaster subtask(s) to remove, e.g., '15.2' or '16.1,16.3'.` (CLI: `-i, --id <id>`)
    *   `convert`: `If used, Taskmaster will turn the subtask into a regular top-level task instead of deleting it.` (CLI: `-c, --convert`)
    *   `generate`: `Enable Taskmaster to regenerate markdown task files after removing the subtask.` (CLI: `--generate`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Delete unnecessary subtasks or promote a subtask to a top-level task.

### 17. Move Task (`move_task`)

*   **MCP Tool:** `move_task`
*   **CLI Command:** `task-master move [options]`
*   **Description:** `Move a task or subtask to a new position within the task hierarchy.`
*   **Key Parameters/Options:**
    *   `from`: `Required. ID of the task/subtask to move (e.g., "5" or "5.2"). Can be comma-separated for multiple tasks.` (CLI: `--from <id>`)
    *   `to`: `Required. ID of the destination (e.g., "7" or "7.3"). Must match the number of source IDs if comma-separated.` (CLI: `--to <id>`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Reorganize tasks by moving them within the hierarchy. Supports various scenarios like:
    *   Moving a task to become a subtask
    *   Moving a subtask to become a standalone task
    *   Moving a subtask to a different parent
    *   Reordering subtasks within the same parent
    *   Moving a task to a new, non-existent ID (automatically creates placeholders)
    *   Moving multiple tasks at once with comma-separated IDs
*   **Validation Features:**
    *   Allows moving tasks to non-existent destination IDs (creates placeholder tasks)
    *   Prevents moving to existing task IDs that already have content (to avoid overwriting)
    *   Validates that source tasks exist before attempting to move them
    *   Maintains proper parent-child relationships
*   **Example CLI:** `task-master move --from=5.2 --to=7.3` to move subtask 5.2 to become subtask 7.3.
*   **Example Multi-Move:** `task-master move --from=10,11,12 --to=16,17,18` to move multiple tasks to new positions.
*   **Common Use:** Resolving merge conflicts in tasks.json when multiple team members create tasks on different branches.

---

## Dependency Management

### 18. Add Dependency (`add_dependency`)

*   **MCP Tool:** `add_dependency`
*   **CLI Command:** `task-master add-dependency [options]`
*   **Description:** `Define a dependency in Taskmaster, making one task a prerequisite for another.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID of the Taskmaster task that will depend on another.` (CLI: `-i, --id <id>`)
    *   `dependsOn`: `Required. The ID of the Taskmaster task that must be completed first, the prerequisite.` (CLI: `-d, --depends-on <id>`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <path>`)
*   **Usage:** Establish the correct order of execution between tasks.

### 19. Remove Dependency (`remove_dependency`)

*   **MCP Tool:** `remove_dependency`
*   **CLI Command:** `task-master remove-dependency [options]`
*   **Description:** `Remove a dependency relationship between two Taskmaster tasks.`
*   **Key Parameters/Options:**
    *   `id`: `Required. The ID of the Taskmaster task you want to remove a prerequisite from.` (CLI: `-i, --id <id>`)
    *   `dependsOn`: `Required. The ID of the Taskmaster task that should no longer be a prerequisite.` (CLI: `-d, --depends-on <id>`)
    *   `tag`: `Specify which tag context to operate on. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Update task relationships when the order of execution changes.

### 20. Validate Dependencies (`validate_dependencies`)

*   **MCP Tool:** `validate_dependencies`
*   **CLI Command:** `task-master validate-dependencies [options]`
*   **Description:** `Check your Taskmaster tasks for dependency issues (like circular references or links to non-existent tasks) without making changes.`
*   **Key Parameters/Options:**
    *   `tag`: `Specify which tag context to validate. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Audit the integrity of your task dependencies.

### 21. Fix Dependencies (`fix_dependencies`)

*   **MCP Tool:** `fix_dependencies`
*   **CLI Command:** `task-master fix-dependencies [options]`
*   **Description:** `Automatically fix dependency issues (like circular references or links to non-existent tasks) in your Taskmaster tasks.`
*   **Key Parameters/Options:**
    *   `tag`: `Specify which tag context to fix dependencies in. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Clean up dependency errors automatically.

---

## Analysis & Reporting

### 22. Analyze Project Complexity (`analyze_project_complexity`)

*   **MCP Tool:** `analyze_project_complexity`
*   **CLI Command:** `task-master analyze-complexity [options]`
*   **Description:** `Have Taskmaster analyze your tasks to determine their complexity and suggest which ones need to be broken down further.`
*   **Key Parameters/Options:**
    *   `output`: `Where to save the complexity analysis report. Default is '.taskmaster/reports/task-complexity-report.json' (or '..._tagname.json' if a tag is used).` (CLI: `-o, --output <file>`)
    *   `threshold`: `The minimum complexity score (1-10) that should trigger a recommendation to expand a task.` (CLI: `-t, --threshold <number>`)
    *   `research`: `Enable research role for more accurate complexity analysis. Requires appropriate API key.` (CLI: `-r, --research`)
    *   `tag`: `Specify which tag context to analyze. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Used before breaking down tasks to identify which ones need the most attention.
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. Please inform users to hang tight while the operation is in progress.

### 23. View Complexity Report (`complexity_report`)

*   **MCP Tool:** `complexity_report`
*   **CLI Command:** `task-master complexity-report [options]`
*   **Description:** `Display the task complexity analysis report in a readable format.`
*   **Key Parameters/Options:**
    *   `tag`: `Specify which tag context to show the report for. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to the complexity report (default: '.taskmaster/reports/task-complexity-report.json').` (CLI: `-f, --file <file>`)
*   **Usage:** Review and understand the complexity analysis results after running analyze-complexity.

---

## File Management

### 24. Generate Task Files (`generate`)

*   **MCP Tool:** `generate`
*   **CLI Command:** `task-master generate [options]`
*   **Description:** `Create or update individual Markdown files for each task based on your tasks.json.`
*   **Key Parameters/Options:**
    *   `output`: `The directory where Taskmaster should save the task files (default: in a 'tasks' directory).` (CLI: `-o, --output <directory>`)
    *   `tag`: `Specify which tag context to generate files for. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
*   **Usage:** Run this after making changes to tasks.json to keep individual task files up to date. This command is now manual and no longer runs automatically.

---

## AI-Powered Research

### 25. Research (`research`)

*   **MCP Tool:** `research`
*   **CLI Command:** `task-master research [options]`
*   **Description:** `Perform AI-powered research queries with project context to get fresh, up-to-date information beyond the AI's knowledge cutoff.`
*   **Key Parameters/Options:**
    *   `query`: `Required. Research query/prompt (e.g., "What are the latest best practices for React Query v5?").` (CLI: `[query]` positional or `-q, --query <text>`)
    *   `taskIds`: `Comma-separated list of task/subtask IDs from the current tag context (e.g., "15,16.2,17").` (CLI: `-i, --id <ids>`)
    *   `filePaths`: `Comma-separated list of file paths for context (e.g., "src/api.js,docs/readme.md").` (CLI: `-f, --files <paths>`)
    *   `customContext`: `Additional custom context text to include in the research.` (CLI: `-c, --context <text>`)
    *   `includeProjectTree`: `Include project file tree structure in context (default: false).` (CLI: `--tree`)
    *   `detailLevel`: `Detail level for the research response: 'low', 'medium', 'high' (default: medium).` (CLI: `--detail <level>`)
    *   `saveTo`: `Task or subtask ID (e.g., "15", "15.2") to automatically save the research conversation to.` (CLI: `--save-to <id>`)
    *   `saveFile`: `If true, saves the research conversation to a markdown file in '.taskmaster/docs/research/'.` (CLI: `--save-file`)
    *   `noFollowup`: `Disables the interactive follow-up question menu in the CLI.` (CLI: `--no-followup`)
    *   `tag`: `Specify which tag context to use for task-based context gathering. Defaults to the current active tag.` (CLI: `--tag <name>`)
    *   `projectRoot`: `The directory of the project. Must be an absolute path.` (CLI: Determined automatically)
*   **Usage:** **This is a POWERFUL tool that agents should use FREQUENTLY** to:
    *   Get fresh information beyond knowledge cutoff dates
    *   Research latest best practices, library updates, security patches
    *   Find implementation examples for specific technologies
    *   Validate approaches against current industry standards
    *   Get contextual advice based on project files and tasks
*   **When to Consider Using Research:**
    *   **Before implementing any task** - Research current best practices
    *   **When encountering new technologies** - Get up-to-date implementation guidance (libraries, apis, etc)
    *   **For security-related tasks** - Find latest security recommendations
    *   **When updating dependencies** - Research breaking changes and migration guides
    *   **For performance optimization** - Get current performance best practices
    *   **When debugging complex issues** - Research known solutions and workarounds
*   **Research + Action Pattern:**
    *   Use `research` to gather fresh information
    *   Use `update_subtask` to commit findings with timestamps
    *   Use `update_task` to incorporate research into task details
    *   Use `add_task` with research flag for informed task creation
*   **Important:** This MCP tool makes AI calls and can take up to a minute to complete. The research provides FRESH data beyond the AI's training cutoff, making it invaluable for current best practices and recent developments.

---

## Tag Management

This new suite of commands allows you to manage different task contexts (tags).

### 26. List Tags (`tags`)

*   **MCP Tool:** `list_tags`
*   **CLI Command:** `task-master tags [options]`
*   **Description:** `List all available tags with task counts, completion status, and other metadata.`
*   **Key Parameters/Options:**
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)
    *   `--show-metadata`: `Include detailed metadata in the output (e.g., creation date, description).` (CLI: `--show-metadata`)

### 27. Add Tag (`add_tag`)

*   **MCP Tool:** `add_tag`
*   **CLI Command:** `task-master add-tag <tagName> [options]`
*   **Description:** `Create a new, empty tag context, or copy tasks from another tag.`
*   **Key Parameters/Options:**
    *   `tagName`: `Name of the new tag to create (alphanumeric, hyphens, underscores).` (CLI: `<tagName>` positional)
    *   `--from-branch`: `Creates a tag with a name derived from the current git branch, ignoring the <tagName> argument.` (CLI: `--from-branch`)
    *   `--copy-from-current`: `Copy tasks from the currently active tag to the new tag.` (CLI: `--copy-from-current`)
    *   `--copy-from <tag>`: `Copy tasks from a specific source tag to the new tag.` (CLI: `--copy-from <tag>`)
    *   `--description <text>`: `Provide an optional description for the new tag.` (CLI: `-d, --description <text>`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)

### 28. Delete Tag (`delete_tag`)

*   **MCP Tool:** `delete_tag`
*   **CLI Command:** `task-master delete-tag <tagName> [options]`
*   **Description:** `Permanently delete a tag and all of its associated tasks.`
*   **Key Parameters/Options:**
    *   `tagName`: `Name of the tag to delete.` (CLI: `<tagName>` positional)
    *   `--yes`: `Skip the confirmation prompt.` (CLI: `-y, --yes`)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)

### 29. Use Tag (`use_tag`)

*   **MCP Tool:** `use_tag`
*   **CLI Command:** `task-master use-tag <tagName>`
*   **Description:** `Switch your active task context to a different tag.`
*   **Key Parameters/Options:**
    *   `tagName`: `Name of the tag to switch to.` (CLI: `<tagName>` positional)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)

### 30. Rename Tag (`rename_tag`)

*   **MCP Tool:** `rename_tag`
*   **CLI Command:** `task-master rename-tag <oldName> <newName>`
*   **Description:** `Rename an existing tag.`
*   **Key Parameters/Options:**
    *   `oldName`: `The current name of the tag.` (CLI: `<oldName>` positional)
    *   `newName`: `The new name for the tag.` (CLI: `<newName>` positional)
    *   `file`: `Path to your Taskmaster 'tasks.json' file. Default relies on auto-detection.` (CLI: `-f, --file <file>`)

### 31. Copy Tag (`copy_tag`)

*   **MCP Tool:** `copy_tag`
*   **CLI Command:** `task-master copy-tag <sourceName> <targetName> [options]`
*   **Description:** `Copy an entire tag context, including all its tasks and metadata, to a new tag.`
*   **Key Parameters/Options:**
    *   `sourceName`: `Name of the tag to copy from.` (CLI: `<sourceName>` positional)
    *   `targetName`: `Name of the new tag to create.` (CLI: `<targetName>` positional)
    *   `--description <text>`: `Optional description for the new tag.` (CLI: `-d, --description <text>`)

---

## Miscellaneous

### 32. Sync Readme (`sync-readme`) -- experimental

*   **MCP Tool:** N/A
*   **CLI Command:** `task-master sync-readme [options]`
*   **Description:** `Exports your task list to your project's README.md file, useful for showcasing progress.`
*   **Key Parameters/Options:**
    *   `status`: `Filter tasks by status (e.g., 'pending', 'done').` (CLI: `-s, --status <status>`)
    *   `withSubtasks`: `Include subtasks in the export.` (CLI: `--with-subtasks`)
    *   `tag`: `Specify which tag context to export from. Defaults to the current active tag.` (CLI: `--tag <name>`)

---

## Environment Variables Configuration (Updated)

Taskmaster primarily uses the **`.taskmaster/config.json`** file (in project root) for configuration (models, parameters, logging level, etc.), managed via `task-master models --setup`.

[TRUNCATED]
```

.tickets/Enable Action Packs Dispatch.md
```
Title:

* Enable oraclepack Action Packs to dispatch to non-oracle executors (codex/gemini/task-master/tm)

Summary:

* Current oraclepack usage feels “oracle-only” because certain UX features (flag injection and overrides validation) are hard-coded to `oracle` invocations, while the user needs Action Packs that actually execute work via other agents/tools (e.g., `codex exec`, `gemini`, `task-master`, `tm`). Update the Stage-3 Action Pack generation and/or oraclepack runner logic so Action Packs can deterministically run the correct executor commands for each action item.

Background / Context:

* User concern: “oraclepack is a wrapper around `oracle`” and adding more `oracle` calls won’t implement tasks; Action Packs must run the real tools/agents used in their workflow (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Current behavior explanation: oraclepack executes shell steps but has special handling only for lines starting with `oracle` (detection regex, flag injection, and overrides validation via `oracle --dry-run summary`).
* Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) and performs guarded checks for autopilot.
* Referenced repos/tools: Gemini CLI, Claude Task Master, OpenAI Codex, steipete/oracle.
* Referenced code/assets: `oraclepack-tui.md`, `oracle_pack_and_taskify-skills.md` (not included in transcript).

Current Behavior (Actual):

* Oraclepack “nice UX features” are oracle-specific:

  * Detects invocations using a regex anchored to literal `oracle`.
  * Injects selected flags only into `oracle …` lines.
  * Validates overrides by running `oracle --dry-run summary` only for detected oracle invocations.
* Action Packs can run arbitrary shell commands, but oraclepack’s overrides/validation UX does not generalize to other tools (codex/gemini/task-master/tm).

Expected Behavior:

* Generated Action Packs can deterministically dispatch execution to the intended executor per action item (e.g., `codex exec …`, `gemini -p …`, `task-master …`, `tm …`) rather than relying on more `oracle` calls.
* Optional: oraclepack’s overrides/flag-injection UX can recognize additional command prefixes beyond `oracle` (if desired).

Requirements:

* Update the Stage-3 generator (“oraclepack-taskify” / Stage-3 Action Pack template) to support configurable tool command strings beyond existing `oracle_cmd`, `task_master_cmd`, `tm_cmd`:

  * Add `codex_cmd` (default `codex`)
  * Add `gemini_cmd` (default `gemini`)
  * Optional: add `autopilot_cmd` (default `${tm_cmd} autopilot`).
* Extend `_actions.json` schema to include an executor plan:

  * `tooling`: include `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per action item: `executor` (`"codex" | "gemini" | "tm" | "manual"`), `exec_prompt` (deterministic; no code fences), optional `inputs` (paths/globs).
* Add a new Action Pack execution path for implementation:

  * Either add `mode=implement`, or add a new “Step 09” gated by `MODE == implement`.
  * Step reads `<out_dir>/_actions.json`, selects top N items (using existing `top_n`), then dispatches:

    * `codex exec …` when `executor == codex`
    * `gemini -p …` when `executor == gemini`.
* Safety constraint:

  * Keep defaults strict (avoid “yolo” execution by default); Gemini approval/tool execution should remain conservative unless explicitly opted in.
* Optional (nice-to-have): generalize oraclepack’s oracle-specific detection/injection:

  * Generalize `ExtractOracleInvocations` / `InjectFlags` to recognize a registry of prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`).
  * Add per-tool override sets (so “Oracle Flags” aren’t incorrectly applied to other tools).

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* OS: Unknown
* oraclepack version/commit: Unknown
* Shell/runner context: Unknown
* Stage-3 generator version/commit: Unknown

Evidence:

* User statement of need: actionpacks must call their agents/tools (examples: `codex exec ...`, `tm ...`, `gemini ...`).
* Oraclepack oracle-specific UX behavior (regex detection, flag injection, `oracle --dry-run summary` validation).
* Proposed schema + implement mode/Step 09 dispatcher design.
* Attachment: Oraclepack Action Pack Integration.md

Decisions / Agreements:

* “Fastest path” identified in transcript: upgrade Stage-3 Action Packs to include an executor dispatch step and extend `_actions.json` with `executor` metadata; modifying oraclepack core is optional.

Open Items / Unknowns:

* Current `_actions.json` schema (exact fields/types) is not provided.
* Current Stage-3 Action Pack template structure (exact steps and modes) is not provided beyond `backlog|pipelines|autopilot`.
* Exact locations/implementations of `ExtractOracleInvocations` / `InjectFlags` in `oraclepack-tui.md` are not provided in this transcript.
* Expected non-interactive invocation patterns/flags for each tool in this project (codex/gemini/task-master/tm) beyond the examples are not provided.

Risks / Dependencies:

* Dependency on external CLIs and their execution/approval modes (especially Gemini CLI) with safety implications; defaults must remain conservative.
* If oraclepack overrides/validation stays oracle-only, users may expect those UX features to apply to non-oracle commands; requires clear separation or per-tool override support.

Acceptance Criteria:

* `_actions.json` produced by the Stage-3 generator includes:

  * `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  * Per-item `executor` and `exec_prompt` (and optional `inputs`).
* Action Pack supports `implement` execution (via `mode=implement` or Step 09) that:

  * Reads `<out_dir>/_actions.json`
  * Selects top N actions
  * Runs `codex exec …` for `executor=codex` and `gemini -p …` for `executor=gemini` deterministically.
* Defaults do not enable unrestricted/automatic tool execution without opt-in (conservative approval/safety posture).
* (Optional) oraclepack runner recognizes non-oracle prefixes for invocation detection and does not incorrectly apply oracle-specific overrides to other tools.

Priority & Severity (if inferable from text):

* Priority: Not provided
* Severity: Not provided

Labels (optional):

* enhancement
* action-pack
* executor-dispatch
* cli-integration
* oraclepack
* taskify
```

.tickets/Improving Oraclepack Workflow.md
```
Title:

* Add deterministic chaining + structured outputs to oraclepack (fix prelude semantics and enable stage-3 “actions” workflow)

Summary:

* The current oraclepack workflow is a 2-stage pipeline (pack generation → oraclepack execution) but has a “disconnect” after execution: the 20 Markdown outputs are not machine-consumable for automated follow-on work, blocking a seamless next stage for actionable implementation.

    Workflow Improvement Suggestions

* Additionally, the runner executes the pack prelude and each step in separate `bash -lc` processes, so prelude-defined variables do not persist into steps, creating a mismatch between template expectations and runtime behavior.

    Workflow Improvement Suggestions

Background / Context:

* Stage 1: a Codex “skill” or Gemini CLI “command” generates a single Markdown “oracle-pack” document whose machine-critical content is a single fenced `bash` block containing 20 numbered steps that call `oracle` with `--write-output ...`.

    Workflow Improvement Suggestions

* Stage 2: oraclepack parses that Markdown by extracting the first \`\`\`bash fence and splitting steps by a step-header pattern, then runs a “prelude” and each step in separate shells, producing 20 output files.

    Workflow Improvement Suggestions

* Goal: improve workflow productivity and longer runtimes with minimal human interaction, including automatically passing the final 20 results into a next stage that yields actionable implementation steps.

    Workflow Improvement Suggestions

Current Behavior (Actual):

* Prelude variables do not persist into step execution because prelude and steps run in separate `bash -lc` processes; packs must use explicit paths instead of relying on prelude variables like `$out_dir`.

    Workflow Improvement Suggestions

* The 20 `oracle` outputs are human-readable Markdown but lack a machine-friendly handoff artifact (e.g., JSON/YAML), making automated downstream processing brittle.

    Workflow Improvement Suggestions

* Parser/run is vulnerable to “format drift” from pack generators (extra code fences, missing/incorrect headers, multiple `bash` fences—parser grabs the first).

    Workflow Improvement Suggestions

Expected Behavior:

* Pack prelude semantics should match user/template expectations (either reliably shared across steps or explicitly non-shared with enforced guidance).

    Workflow Improvement Suggestions

* After a run, oraclepack should produce a deterministic, machine-readable “handoff” that enables an immediate next stage without manual intervention (“without missing a beat”).

    Workflow Improvement Suggestions

* Pack ingestion should be resilient and self-validating (fail fast on drift and contract violations).

    Workflow Improvement Suggestions

Requirements:

* Prelude semantics (choose and make official):

    Workflow Improvement Suggestions

  * Option A: “Prelude is prep-only” (no shared vars): update pack template guidance accordingly.

  * Option B: “Prelude is sourced into every step”: implement by prefixing each step script with prelude content (or run all steps in a single long-lived shell session).

* Stage-1 → Stage-2 contract hardening (“self-healing”):

    Workflow Improvement Suggestions

  * Standardize step headers to the conservative form `# NN)`.

  * Enforce exactly one fenced `bash` block per pack (or validation error).

  * Run `oraclepack validate` immediately after pack generation (wrapper/scriptable convention).

* Add a machine-readable “run index” artifact per run:

    Workflow Improvement Suggestions

  * Must include per-step `--write-output` paths, exit codes, timestamps.

  * Include parsed metadata when available (ROI/category/reference from step header line).

* Add a first-class “chain” capability to generate an “actions” next stage:

    Workflow Improvement Suggestions

  * Proposed interface: `oraclepack chain <pack.md> --mode actions`.

  * Must synthesize: `oracle-out/_summary.md` (human) and `oracle-out/_actions.json` (machine).

  * `_actions.json` should normalize each item with at least: `category`, `roi`, `reference`, `recommended_action` (“Next smallest concrete experiment”), `missing_artifacts[]`, `risk_notes[]`.

        Workflow Improvement Suggestions

  * Emit a follow-on pack: `docs/oracle-actions-YYYY-MM-DD.md`.

        Workflow Improvement Suggestions

* Execution/runtime considerations:

  * Keep compatibility with non-interactive operation (`--no-tui`, `--run-all`) and stop-on-fail behavior so chaining can run in CI.

        Workflow Improvement Suggestions

  * Optional: opt-in parallel execution with a small concurrency cap if provider/rate limits allow.

        Workflow Improvement Suggestions

* Improve pack input/context usage:

  * Support “focus + exclusion” inputs (e.g., `focus_categories=permissions,migrations`, `exclude_paths=docs,node_modules,dist`) without changing pack shape.

        Workflow Improvement Suggestions

  * Treat “extra\_files” as a deliberate context bundle (org standards/ADRs/threat model/coding conventions) appended to commands.

        Workflow Improvement Suggestions

Out of Scope:

* Not provided

Reproduction Steps:

1. Generate an oracle-pack whose prelude defines `out_dir="..."` and steps reference `$out_dir/...`.

2. Run oraclepack on the pack.

3. Observe that step shells do not see prelude-defined variables (because each step runs in a separate `bash -lc`), requiring explicit paths instead.

    Workflow Improvement Suggestions

Environment:

* OS: Unknown

* oraclepack: Unknown version (Go wrapper around `oracle`)

    Workflow Improvement Suggestions

* Shell execution model: `bash -lc` per step (per assistant analysis)

    Workflow Improvement Suggestions

* Pack generators: Codex skill and Gemini CLI command workflows

    Workflow Improvement Suggestions

Evidence:

* “oraclepack executes **prelude and steps in separate `bash -lc` processes**, so variables defined in the prelude do **not** persist into the step shells.”

    Workflow Improvement Suggestions

* Format drift risks listed: extra code fences, missing/incorrect step headers, multiple `bash` fences (parser grabs the first).

    Workflow Improvement Suggestions

* Proposed chain command + structured artifacts: `_summary.md`, `_actions.json`, and `docs/oracle-actions-YYYY-MM-DD.md`.

    Workflow Improvement Suggestions

Decisions / Agreements:

* None explicitly finalized; two alternative fixes for prelude semantics were presented (runner fix vs template fix), but no selection recorded.

    Workflow Improvement Suggestions

Open Items / Unknowns:

* Which prelude semantic is intended as the official contract (“prep-only” vs “sourced into every step”).

    Workflow Improvement Suggestions

* Exact current parser rules (regex specifics), validation behavior, and current report/state outputs (whether a run index already exists).

    Workflow Improvement Suggestions

* Whether Task Master integration is desired as a required dependency or just an optional downstream consumer.

    Workflow Improvement Suggestions

Risks / Dependencies:

* Dependency on consistent pack formatting across multiple generators (Codex/Gemini); drift can break parsing/validation.

    Workflow Improvement Suggestions

* If parallelism is added, provider rate limits and error handling may complicate deterministic runs.

    Workflow Improvement Suggestions

* Downstream automation quality depends on having a structured index/JSON handoff rather than parsing free-form Markdown answers.

    Workflow Improvement Suggestions

Acceptance Criteria:

* Running a pack that defines variables in the prelude and references them in steps behaves according to the selected official contract:

    *   If “sourced prelude”: steps can reference prelude-defined variables successfully.

    *   If “prep-only”: validation or guidance prevents packs from relying on prelude vars (or template guidance is updated and enforced).

        Workflow Improvement Suggestions

* `oraclepack validate` reliably fails on packs with:

    *   multiple `bash` fences,

    *   missing/incorrect `# NN)` step headers (per enforced convention),

    *   other contract-breaking drift conditions.

        Workflow Improvement Suggestions

* After a run, oraclepack emits a machine-readable run index that includes per-step output paths, exit codes, and timestamps.

    Workflow Improvement Suggestions

* `oraclepack chain <pack.md> --mode actions` produces:

    *   `oracle-out/_summary.md`,

    *   `oracle-out/_actions.json` with the specified normalized fields,

    *   `docs/oracle-actions-YYYY-MM-DD.md` suitable for immediate stage-2 execution.

        Workflow Improvement Suggestions

* Chaining works in non-interactive mode and respects stop-on-fail behavior to support CI usage.

    Workflow Improvement Suggestions

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* enhancement

* workflow

* automation

* cli

* parsing

* schema

* ci

---
```

.tickets/Oraclepack Action Pack Integration.md
```
## You asked

oracle\_pack\_and\_taskify-skills.md

File

oraclepack-tui.md

File

Our question is since oraclepack is a wrapper around `oracle`, we can not see how running it as is with more oracle calls will help us implement the tasks we created. I provided a list of links of the tools and agents we use, so what we need is to make it so our actionpacks will perform the correct calls to the agents i.e., codex and/or tools i.e., task-master-ai. Basically making it so our oraclepack can run action packs that contain the correct commands to command the agents and/or tools we will be using here. Example 'codex exec ...', 'tm ...', 'gemini ...' etc... A screenshot of the oraclepack consuming oraclepack-taskify artifacts. oraclepack source code is in `oraclepack-tui.md` `https://github.com/google-gemini/gemini-cli` `https://github.com/eyaltoledano/claude-task-master` `https://github.com/openai/codex` `https://github.com/steipete/oracle`

---

## ChatGPT Replied

Thought for 4m 14s

You’re right that **adding more `oracle` calls doesn’t implement tasks**—it only produces more analysis/synthesis. In your current system, implementation only happens when the **Action Pack runs real tool commands** (e.g., `task-master …`, `tm autopilot`, etc.). The key is to make the Action Pack (and optionally oraclepack’s runner UX) **tool-agnostic**, so it can dispatch to `codex`, `gemini`, `tm`, etc.

What oraclepack does today (why it “feels like oracle-only”)
------------------------------------------------------------

* Oraclepack executes each step as shell, but it has **special logic only for lines that start with `oracle`**:

  * It detects invocations using a regex anchored to the literal `oracle` command.

        oraclepack-tui

  * It injects selected flags only into `oracle …` lines.

        oraclepack-tui

  * Its overrides validation runs `oracle --dry-run summary` only for detected oracle invocations.

        oraclepack-tui

* Your Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) and does guarded branch checks for autopilot.

So: **oraclepack can run any command**, but its “nice UX features” (flag overrides + validation) are currently **oracle-specific**.

The practical solution: make Stage-3 Action Packs dispatch to your real agents/tools
------------------------------------------------------------------------------------

Your Stage-3 skill already supports configurable tool command strings:

* `oracle_cmd`, `task_master_cmd`, `tm_cmd`

    oracle\_pack\_and\_taskify-skills

Extend that same pattern for:

* `codex_cmd` (default `codex`)

* `gemini_cmd` (default `gemini`)

* optionally `autopilot_cmd` (default `${tm_cmd} autopilot`)

### Why this works

* **Codex** has an explicit non-interactive execution entrypoint: `codex exec` (alias `codex e`). [OpenAI Developers+1](https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com)

* **Gemini CLI** supports headless prompting via `--prompt` / `-p`, and has explicit controls for approval/tooling. [Google Gemini+2Gemini CLI+2](https://google-gemini.github.io/gemini-cli/docs/cli/?utm_source=chatgpt.com)

* **Task Master** provides deterministic CLI flows (`parse-prd`, `analyze-complexity`, `expand`). [Medium+1](https://medium.com/google-cloud/gemini-cli-tutorial-series-part-2-gemini-cli-command-line-parameters-e64e21b157be?utm_source=chatgpt.com)

Recommended architecture: keep oracle as “planner”, make actionpacks do “executor dispatch”
-------------------------------------------------------------------------------------------

### 1) Extend `_actions.json` to include an executor plan per item

Your canonical actions schema today has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action`, `acceptance_criteria`, etc.

oracle\_pack\_and\_taskify-skills

Add fields like:

* `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`

* per item:

  * `executor`: `"codex" | "gemini" | "tm" | "manual"`

  * `exec_prompt`: a short instruction string (no code fences; deterministic)

  * `inputs`: optional list of paths/globs (from `missing_artifacts` / repo anchors)

This lets you generate a fully deterministic “what to run” plan without guessing repo specifics.

### 2) Add an “implement” mode (or a Step 09) to the Action Pack template

Your Action Pack already has modes `backlog|pipelines|autopilot`.

oracle\_pack\_and\_taskify-skills

Add either:

* `mode=implement` (clean separation), or

* a new **Step 09** guarded by `if MODE == implement`.

That step:

* reads `<out_dir>/_actions.json`

* selects the top N items (you already have `top_n`)

    oracle\_pack\_and\_taskify-skills

* dispatches deterministically:

  * `codex exec …` for items with executor `codex`

  * `gemini -p …` for items with executor `gemini`

### 3) Keep safety defaults strict (do not “yolo” by default)

Gemini CLI explicitly supports different approval modes; keep defaults conservative unless a user opts in. [Gemini CLI+1](https://geminicli.com/docs/get-started/configuration-v1/?utm_source=chatgpt.com)
Also be aware Gemini CLI has had real-world safety footguns around command/tool execution; sandboxing and explicit approvals matter. [IT Pro](https://www.itpro.com/security/a-flaw-in-googles-new-gemini-cli-tool-couldve-allowed-hackers-to-exfiltrate-data?utm_source=chatgpt.com)
Codex similarly has concepts like exec policies and JSON event streaming (`--json`) that are useful for controlled automation. [OpenAI Developers+1](https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com)

Minimal changes you should make (best ROI)
------------------------------------------

### A) Update **oraclepack-taskify** (Stage 3 generator), not oraclepack core

1. **Add args**: `codex_cmd`, `gemini_cmd`, `autopilot_cmd` (optional) alongside existing tool cmd args.

    oracle\_pack\_and\_taskify-skills

2. **Update Prompt A** (canonical actions synthesis) to emit `executor` + `exec_prompt` per action item.

    oracle\_pack\_and\_taskify-skills

3. **Update Action Pack template** to include the dispatch step that runs the chosen CLI tools.

This gives you an Action Pack that actually runs:

* `codex exec …` (non-interactive) [OpenAI Developers+1](https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com)

* `gemini -p …` (headless prompt) [Google Gemini](https://google-gemini.github.io/gemini-cli/docs/cli/?utm_source=chatgpt.com)

* `task-master …` for task ingestion/expansion [Medium+1](https://medium.com/google-cloud/gemini-cli-tutorial-series-part-2-gemini-cli-command-line-parameters-e64e21b157be?utm_source=chatgpt.com)

### B) Optional: improve oraclepack UX to “understand” non-oracle commands

If you want oraclepack’s TUI/overrides to apply to other tools:

* Generalize `ExtractOracleInvocations` / `InjectFlags` to a small registry of command prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`). Right now it is hard-coded to `oracle` only.

* Add per-tool override sets (otherwise “Oracle Flags” won’t make sense for codex/gemini).

This is **nice-to-have**; it’s not required to make actionpacks actually execute your agents.

Concrete command patterns you can standardize on
------------------------------------------------

* Codex implementation call (headless): `codex exec …` [OpenAI Developers+1](https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com)

* Gemini implementation call (headless): `gemini -p "…"`, plus optional `--approval-mode …` / `--allowed-tools …` controls [Google Gemini+1](https://google-gemini.github.io/gemini-cli/docs/cli/?utm_source=chatgpt.com)

* Task Master pipeline:

  * `task-master parse-prd <prd>`

  * `task-master analyze-complexity --output <out_dir>/tm-complexity.json`

  * `task-master expand --all` [Medium+1](https://medium.com/google-cloud/gemini-cli-tutorial-series-part-2-gemini-cli-command-line-parameters-e64e21b157be?utm_source=chatgpt.com)

Bottom line
-----------

* Oraclepack is fine as the runner; it’s “oracle-centric” only in its flag-injection/validation UX.

* The fastest path is: **upgrade Stage-3 Action Packs to include an executor dispatch step** and extend `_actions.json` with `executor` metadata so the pack can deterministically run `codex`, `gemini`, etc.

If you want, I can propose the exact schema delta for `_actions.json` and the exact additional Step 09 bash logic (dispatcher) in a way that preserves your determinism rules (self-contained steps, stable ordering, fail-fast preflight).

---
```

.tickets/Oraclepack Action Pack Issue.md
```
Parent Ticket:

* Title: Oraclepack Action Pack compatibility and non-`oracle` command handling clarity
* Summary: Clarify what happens when running `oraclepack-taskify`-generated Action Packs in current oraclepack, especially the difference between “executing the pack” vs “dispatching/wrapping non-`oracle` commands.” Capture current limitations: oraclepack’s special handling (flag injection + override validation) targets `oracle` invocations only; other CLIs run as plain shell commands and may fail/block depending on PATH and interactivity.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Two meanings of ‘run these actionpack artifacts’… execute the pack file vs dispatch non-`oracle` commands.”
* Global Constraints:

  * Action Pack is described as “oraclepack-ingestible” with “single ```bash fence” and “# NN) step headers.”
* Global Environment:

  * Unknown
* Global Evidence:

  * docs.task-master.dev link(s) referenced in ticket text
  * developers.openai.com Codex CLI link(s) referenced in ticket text
  * docs.cloud.google.com / google-gemini.github.io Gemini CLI link(s) referenced in ticket text

Split Plan:

* Coverage Map:

  * Original item: “Can our current oraclepack run these actionpack artifacts… generated from the oraclepack-taskify skill?” → Assigned Ticket ID: T1
  * Original item: “Action Pack… ‘oraclepack-ingestible’ (single ```bash fence, # NN) step headers, deterministic paths…)” → Assigned Ticket ID: T1
  * Original item: “How you’d run it… `oraclepack validate …` / `oraclepack run …`” → Assigned Ticket ID: T1
  * Original item: “Actionpacks calling other agents/tools (codex exec, tm autopilot, gemini)… will run them as long as installed and usable non-interactively” → Assigned Ticket ID: T1
  * Original item: “Why did another reviewer say otherwise?” → Assigned Ticket ID: T1
  * Original item: “oraclepack runs each step as a bash script (bash -lc <step script>)” → Assigned Ticket ID: T1
  * Original item: “injects flags into commands that begin with oracle… does not match tm/task-master/codex/gemini” → Assigned Ticket ID: T3
  * Original item: “TUI ‘override validation’ only targets oracle… steps without oracle invocations are skipped” → Assigned Ticket ID: T2
  * Original item: “If binary isn’t installed/on PATH → step fails; if CLI is interactive → it will block” → Assigned Ticket ID: T1
  * Original item: “Only way it ‘gets confused’ is if you expect oracle output text to magically invoke Codex/Gemini” → Assigned Ticket ID: T1
  * Original item: “Two meanings… execute vs dispatch/apply oraclepack overrides to non-oracle commands” → Assigned Ticket ID: T1
  * Original item: “Reviewer answered ‘No’ because broader goal is dispatcher behavior… not interpret actions” → Assigned Ticket ID: Info-only
  * Original item: “Reconciliation: both statements can be true” → Assigned Ticket ID: Info-only
* Dependencies:

  * Not provided
* Split Tickets:

````ticket T1
T# Title: Document current Action Pack execution semantics and operator expectations
Type: docs
Target Area: oraclepack documentation / runbook for Action Packs
Summary:
  Clarify what “running an Action Pack” means in current oraclepack: steps execute as shell via `bash -lc`, and any `tm`/`codex`/`gemini` lines run as plain shell commands. Document the two meanings that caused reviewer disagreement: executing the pack vs dispatching/applying oracle-specific override logic to non-`oracle` commands. Include operator-facing notes on common failure modes already described (PATH missing binaries, interactive CLIs blocking, environment guardrails for autopilot).
In Scope:
  - Explain: “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - Explain: non-`oracle` commands (`tm`, `task-master`, `codex`, `gemini`) are executed “directly as normal shell commands.”
  - Provide the exact run examples already given (`oraclepack validate …` and `oraclepack run …`).
  - Capture the limitation: oraclepack’s oracle-specific overrides/validation do not apply to non-`oracle` commands.
  - Document noted failure/blocking conditions:
    - Missing binaries not on PATH.
    - Interactive CLIs blocking waiting for input.
    - Autopilot “fail fast” environment issues (e.g., not in git repo, dirty tree) as stated.
Out of Scope:
  - Implementing dispatcher behavior for non-`oracle` tools (not specified beyond “broader goal” mention).
Current Behavior (Actual):
  - Steps are run as `bash -lc <step script>`.
  - Non-`oracle` commands run directly; no automatic dispatch/wrapping is applied.
Expected Behavior:
  - Operators can correctly predict:
    - What will execute (literal commands in the pack).
    - What will not happen automatically (no “magic” invocation of Codex/Gemini from oracle output text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Preserve stated Action Pack constraints: “single ```bash fence” and “# NN) step headers” (as described).
Evidence:
  - Included links mentioned in ticket text:
    - docs.task-master.dev references
    - developers.openai.com Codex CLI references
    - Gemini CLI references (docs.cloud.google.com / google-gemini.github.io)
Open Items / Unknowns:
  - Exact location(s) where this documentation should live (Not provided).
  - Whether this should be surfaced in CLI help text vs README vs TUI (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - Documentation explicitly states:
    - Packs execute step bodies via `bash -lc`.
    - Non-`oracle` commands run as-is and are not routed through oracle-specific logic.
    - Common failure modes listed in the ticket text (PATH missing, interactive blocking, autopilot environment guards).
  - Documentation includes the exact run commands already provided (`oraclepack validate …`, `oraclepack run …`).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “oraclepack executes each step’s body as shell via `bash -lc <command>`.”
  - “Will oraclepack dispatch non-`oracle` commands…? No… only targets commands that start with `oracle`.”
  - “If the CLI is interactive → it will block waiting for input.”
````

```ticket T2
T# Title: Make TUI override validation behavior explicit for steps without `oracle` invocations
Type: enhancement
Target Area: TUI overrides flow / override validation messaging
Summary:
  The ticket text states that TUI “override validation” runs `oracle --dry-run summary` on detected `oracle` invocations and skips steps without `oracle` calls. Make this behavior explicit in the TUI so operators do not misinterpret skipped steps as “validated,” especially when packs include `tm`/`codex`/`gemini` commands.
In Scope:
  - Surface an explicit note/state in the overrides validation flow indicating:
    - Validation applies only to detected `oracle` invocations.
    - Steps without `oracle` invocations are skipped by this validator.
Out of Scope:
  - Adding validation implementations for `tm`/`codex`/`gemini` (not described in ticket text).
Current Behavior (Actual):
  - “Override validation… only targets `oracle` commands… Steps without oracle invocations are skipped.”
Expected Behavior:
  - TUI clearly communicates when steps are skipped (and why), avoiding operator confusion.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must preserve current behavior described (validate only `oracle` invocations) unless changed elsewhere (Not provided).
Evidence:
  - “The TUI ‘override validation’ also only targets `oracle` commands… Steps without oracle invocations are skipped…”
Open Items / Unknowns:
  - Exact UI copy/placement and which screen(s) in the TUI should show this (Not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - When override validation runs, the UI explicitly indicates:
    - It validates only `oracle` invocations (via `oracle --dry-run summary`, as stated).
    - Steps without `oracle` invocations are skipped (and shown as skipped).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “The TUI ‘override validation’ also only targets `oracle` commands… runs `oracle --dry-run summary`…”
  - “Steps without oracle invocations are skipped by that validator.”
```

```ticket T3
T# Title: Extend runtime flag-injection matching beyond `oracle` invocations (configurable prefixes)
Type: enhancement
Target Area: command rewriting / runtime override injection
Summary:
  The ticket text states oraclepack injects flags only into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`) and does not match `tm`, `task-master`, `codex`, or `gemini`. Add support for matching additional command prefixes (or a configurable list) so override injection is not limited to `oracle` only, aligning with packs that include other CLIs.
In Scope:
  - Expand the “inject flags” behavior beyond `oracle`-only matching, as motivated by:
    - “does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
  - Preserve anchored/prefix-based matching semantics as described (no broad substring matching implied).
Out of Scope:
  - Defining tool-specific semantics for what flags should be injected for each CLI (Not provided).
  - Implementing dispatcher logic that changes execution from “literal shell command” to “interpreted actions” (Not provided).
Current Behavior (Actual):
  - “Injects flags into commands that begin with `oracle`… regex anchored to `^(\s*)(oracle)\b`.”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Expected Behavior:
  - Flag injection can apply to non-`oracle` command prefixes as configured/defined (details not provided in ticket text).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Must not break execution of packs that rely on current `oracle`-only behavior (Not provided).
Evidence:
  - “injects flags into commands that begin with `oracle` (regex anchored to `^(\s*)(oracle)\b`). It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
Open Items / Unknowns:
  - Which non-`oracle` commands should be included first (list not provided beyond examples).
  - Where configuration should live (Not provided).
  - Whether injection should be opt-in per pack/step or global (Not provided).
Risks / Dependencies:
  - Risk: unintended rewriting of commands if prefix matching is overly broad (mitigate via anchored matching; exact approach not provided).
Acceptance Criteria:
  - There is a documented/configured way to include additional command prefixes for injection beyond `oracle`.
  - Existing `oracle` prefix injection continues to work unchanged.
  - Demonstrably, a command beginning with one added prefix is recognized for injection (exact flags and CLI semantics: Not provided).
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “injects flags into commands that begin with `oracle` (regex is anchored to `^(\\s*)(oracle)\\b`).”
  - “It does not match `tm`, `task-master`, `codex`, `gemini`, etc.”
```
```

.tickets/Oraclepack Action Packs.md
```
Parent Ticket:

* Title: Oraclepack Action Packs: tool-agnostic execution (Codex/Gemini/Task Master) instead of oracle-only flows
* Summary:

  * Current pain point: oraclepack is a wrapper around `oracle`, and “more oracle calls” won’t implement taskified work; Action Packs must run real tool commands (e.g., `codex`, `gemini`, `tm`, `task-master`).
  * Key idea: keep oracle as “planner” and make Action Packs do deterministic executor dispatch via `_actions.json` metadata and a new implement mode/step.
* Source:

  * Link/ID (if present) or “Not provided”: Uploaded file: Oraclepack Action Pack Integration.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “make it so our actionpacks will perform the correct calls… Example ‘codex exec …’, ‘tm …’, ‘gemini …’”.
* Global Constraints:

  * Action Packs should be “tool-agnostic” (dispatch to `codex`, `gemini`, `tm`, etc.).
  * `exec_prompt` should be short and deterministic (explicitly “no code fences”).
  * Safety defaults must be strict (“do not yolo by default”); conservative approval/tool execution unless opted-in.
  * Optional/“nice-to-have”: generalize oraclepack’s oracle-specific overrides UX for other command prefixes.
* Global Environment:

  * Unknown
* Global Evidence:

  * Mentioned repos/tools:

    * `https://github.com/google-gemini/gemini-cli`
    * `https://github.com/eyaltoledano/claude-task-master`
    * `https://github.com/openai/codex`
    * `https://github.com/steipete/oracle`
  * Reference docs cited/used in the ticket text:

    * Codex CLI reference (supports “codex exec”). ([OpenAI Developers][1])
    * Gemini CLI docs (tools/approval concepts referenced by the ticket). ([Gemini CLI][2])
    * Claude Task Master repository. ([GitHub][3])

Split Plan:

* Coverage Map:

  * “make it so our actionpacks will perform the correct calls… ‘codex exec …’, ‘tm …’, ‘gemini …’” → T3
  * Oraclepack has special logic only for lines starting with `oracle` (regex detection / flag injection / validation) → T5
  * “Stage-3 Action Pack template already runs non-oracle tools (`task-master …` and `tm autopilot`) … guarded branch checks” → T3
  * Stage-3 skill supports `oracle_cmd`, `task_master_cmd`, `tm_cmd` → T1
  * “Extend that same pattern for `codex_cmd`, `gemini_cmd`, optionally `autopilot_cmd`” → T1
  * “Extend `_actions.json` … `executor`, `exec_prompt` (no code fences), `inputs`, plus expanded `tooling`” → T2
  * “Add an ‘implement’ mode (or Step 09)… reads `<out_dir>/_actions.json`… selects top N (`top_n`)… dispatches” → T3
  * “Keep safety defaults strict (do not ‘yolo’ by default)” → T4
  * “Minimal changes… add args… update Prompt A… update Action Pack template” → T1, T2, T3
  * “Optional: improve oraclepack UX… registry of command prefixes… per-tool override sets” → T5
  * “Concrete command patterns… `codex exec`, `gemini -p`, Task Master pipeline commands” → Info-only
  * “If you want, I can propose the exact schema delta… Step 09 bash logic” → Info-only
* Dependencies:

  * T3 depends on T2 because the implement/dispatcher step reads `_actions.json` and needs `executor`/`exec_prompt` metadata.
  * T2 depends on T1 because `_actions.json.tooling` expansion references new tool command fields (`codex_cmd`, `gemini_cmd`, optional `autopilot_cmd`).
  * T4 depends on T3 because safety defaults/opt-ins apply to the implement/dispatcher execution path.
  * T5 is independent (optional) but may follow T3 if you want the TUI/overrides UX to apply to non-oracle commands.
* Split Tickets:

```ticket T1
T1 Title:
- Extend oraclepack-taskify Stage-3 generator to accept and propagate executor CLI command configs (codex/gemini/autopilot)

Type:
- enhancement

Target Area:
- oraclepack-taskify (Stage-3 generator inputs/args and emitted configs)

Summary:
- The Stage-3 generator already supports configurable command strings for `oracle_cmd`, `task_master_cmd`, and `tm_cmd`.
- Extend the same configuration pattern to include `codex_cmd` and `gemini_cmd`, and optionally `autopilot_cmd`, so Action Packs can invoke the intended executors without hard-coding tool names.

In Scope:
- Add generator inputs/args for:
  - `codex_cmd` (default `codex`)
  - `gemini_cmd` (default `gemini`)
  - optional `autopilot_cmd` (default `${tm_cmd} autopilot`)
- Ensure generated artifacts carry these command strings for later use by the Action Pack execution steps.

Out of Scope:
- Modifying oraclepack core/TUI behavior (handled in T5)
- Implement-mode dispatcher logic (handled in T3)

Current Behavior (Actual):
- Stage-3 skill supports configurable tool command strings:
  - `oracle_cmd`, `task_master_cmd`, `tm_cmd`

Expected Behavior:
- Stage-3 generator also supports `codex_cmd`, `gemini_cmd`, and optionally `autopilot_cmd`, using the same configuration pattern.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Preserve the existing “configurable tool command strings” pattern already used by the Stage-3 generator.
- Defaults as stated in the ticket text.

Evidence:
- References: “Your Stage-3 skill already supports… `oracle_cmd`, `task_master_cmd`, `tm_cmd`… Extend… `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`…”

Open Items / Unknowns:
- Where these args are defined/passed in the current generator (file paths not provided).
- Whether additional executors beyond codex/gemini/tm are needed (not provided).

Risks / Dependencies:
- Depends on T2 if `_actions.json.tooling` is expanded to include these command strings.

Acceptance Criteria:
- Stage-3 generator accepts `codex_cmd` and `gemini_cmd` (and optional `autopilot_cmd`) inputs.
- Defaults match: `codex`, `gemini`, and `${tm_cmd} autopilot` (optional).
- Generated outputs expose these command strings for downstream Action Pack steps.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Your Stage-3 skill already supports configurable tool command strings: `oracle_cmd`, `task_master_cmd`, `tm_cmd`”
- “Extend that same pattern for: `codex_cmd`… `gemini_cmd`… optionally `autopilot_cmd`”
```

```ticket T2
T2 Title:
- Extend `_actions.json` schema + Prompt A to emit per-item executor metadata (executor/exec_prompt/inputs) and expanded tooling

Type:
- enhancement

Target Area:
- Canonical actions schema (`_actions.json`) and “Prompt A” (canonical actions synthesis)

Summary:
- The current actions schema has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action` and `acceptance_criteria`.
- Add executor planning fields per action item so Action Packs can deterministically select and run the correct executor (`codex`, `gemini`, `tm`, or manual), and include relevant inputs.

In Scope:
- Update `_actions.json` to add:
  - `tooling`: `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`
  - per-item fields:
    - `executor`: `"codex" | "gemini" | "tm" | "manual"`
    - `exec_prompt`: short instruction string (explicitly “no code fences; deterministic”)
    - `inputs`: optional list of paths/globs (from `missing_artifacts` / repo anchors)
- Update “Prompt A” to emit the above fields for each action item.

Out of Scope:
- Implement-mode dispatcher logic that consumes `_actions.json` (handled in T3)
- oraclepack core UX changes (handled in T5)

Current Behavior (Actual):
- `_actions.json` has `tooling` (oracle/task-master) and per-item fields like `recommended_next_action`, `acceptance_criteria`.

Expected Behavior:
- `_actions.json` includes explicit executor plan per item (`executor`, `exec_prompt`, optional `inputs`) and expanded `tooling` including codex/gemini command strings.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- `exec_prompt` must be short and deterministic; explicitly “no code fences”.
- Executor enum values and tooling fields as written in the ticket text.

Evidence:
- References: “Extend `_actions.json` to include an executor plan per item… `executor`… `exec_prompt`… `inputs`…”

Open Items / Unknowns:
- The exact existing `_actions.json` schema shape and where it’s validated (not provided).
- Whether `missing_artifacts` / repo anchors already exist in the schema (not provided).

Risks / Dependencies:
- Depends on T1 if `tooling` is expected to include the new `codex_cmd`/`gemini_cmd` command strings.
- Required by T3 since the dispatcher reads `_actions.json`.

Acceptance Criteria:
- `_actions.json` schema includes `tooling` with `{ oracle_cmd, task_master_cmd, codex_cmd, gemini_cmd }`.
- Each action item can include:
  - `executor` with allowed values: `codex`, `gemini`, `tm`, `manual`
  - `exec_prompt` (no code fences requirement captured)
  - optional `inputs` list
- Prompt A output includes these fields per item.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Extend `_actions.json` to include an executor plan per item”
- “`executor`: ‘codex’ | ‘gemini’ | ‘tm’ | ‘manual’”
- “`exec_prompt`: a short instruction string (no code fences; deterministic)”
```

```ticket T3
T3 Title:
- Add Action Pack “implement” mode (or Step 09) to dispatch executor commands based on `_actions.json`

Type:
- enhancement

Target Area:
- Stage-3 Action Pack template (runner-ingestible actionpack artifacts)

Summary:
- Action Packs must execute real tool commands to implement tasks; adding more `oracle` calls only produces more analysis/synthesis.
- Add an implement execution path that reads `_actions.json`, selects top N actions (`top_n`), and dispatches to the specified executor (e.g., `codex exec …`, `gemini -p …`) using per-item metadata.

In Scope:
- Add either:
  - `mode=implement`, or
  - a new Step 09 guarded by `if MODE == implement`
- Implement-mode behavior:
  - Read `<out_dir>/_actions.json`
  - Select the top N items (uses existing `top_n`)
  - Dispatch deterministically:
    - `codex exec …` for items with `executor=codex`
    - `gemini -p …` for items with `executor=gemini`
- Keep existing modes (`backlog|pipelines|autopilot`) intact.

Out of Scope:
- Changing oraclepack core logic for overrides/flag injection/validation (handled in T5)

Current Behavior (Actual):
- Implementation happens only when the Action Pack runs real tool commands (e.g., `task-master …`, `tm autopilot`).
- The Action Pack already runs non-oracle tools (`task-master …` and `tm autopilot`) and includes guarded branch checks for autopilot.

Expected Behavior:
- An implement mode/step exists that consumes `_actions.json` and runs executor-specific commands deterministically based on per-item metadata.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must read `<out_dir>/_actions.json`.
- Must use existing `top_n` to select items.
- Must dispatch based on `executor` field.
- Preserve deterministic behavior as described (stable ordering / fail-fast preflight referenced as desired properties).

Evidence:
- References: “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically…”

Open Items / Unknowns:
- The exact Action Pack template file path and step numbering constraints (not provided).
- How “top N” is currently computed/ordered (not provided).

Risks / Dependencies:
- Depends on T2 since implement mode reads `_actions.json` executor metadata.
- Safety defaults for tool execution are addressed in T4.

Acceptance Criteria:
- Action Pack supports `mode=implement` (or an equivalent Step 09 guarded behavior).
- Implement mode:
  - Reads `<out_dir>/_actions.json`
  - Selects top N via `top_n`
  - Runs `codex exec …` for `executor=codex`
  - Runs `gemini -p …` for `executor=gemini`
- Existing `backlog|pipelines|autopilot` behavior remains unchanged.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “implementation only happens when the Action Pack runs real tool commands”
- “Add an ‘implement’ mode (or a Step 09)… reads `<out_dir>/_actions.json`… selects the top N… dispatches deterministically”
```

```ticket T4
T4 Title:
- Enforce conservative safety defaults for executor dispatch (no auto-approve “yolo” behavior unless opted-in)

Type:
- chore

Target Area:
- Action Pack implement/dispatcher execution path (gemini/codex dispatch configuration)

Summary:
- The ticket explicitly calls out safety: executor tools may run commands or perform actions; defaults should be conservative.
- Add/confirm guardrails so the implement/dispatcher mode does not auto-approve tool execution by default, and requires explicit opt-in for riskier behaviors.

In Scope:
- Ensure implement/dispatcher mode defaults are “strict” and not “yolo by default”.
- Ensure any approval/tool-execution modes are conservative unless a user opts in (mechanism not specified in the ticket text).

Out of Scope:
- Defining new security models or sandboxing systems beyond what’s stated (not provided).

Current Behavior (Actual):
- Not provided (safety behavior in implement/dispatcher mode is not yet implemented per ticket text).

Expected Behavior:
- Implement/dispatcher execution has conservative defaults around tool approvals/execution; opt-in required for less restrictive modes.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- “Keep safety defaults strict (do not ‘yolo’ by default)”
- “keep defaults conservative unless a user opts in”

Evidence:
- References: “Keep safety defaults strict (do not ‘yolo’ by default)… sandboxing and explicit approvals matter…”

Open Items / Unknowns:
- Which specific flags/options are used to control approvals for each executor (not provided).
- How/where opt-in is configured (env var, arg, config) (not provided).

Risks / Dependencies:
- Depends on T3 because the implement/dispatcher path is where safety defaults apply.

Acceptance Criteria:
- Implement/dispatcher mode does not default to auto-approving tool execution.
- Any non-conservative approval behavior requires explicit user opt-in (as defined by existing config patterns; specifics not provided).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Keep safety defaults strict (do not ‘yolo’ by default)”
- “keep defaults conservative unless a user opts in”
```

```ticket T5
T5 Title:
- Optional: Generalize oraclepack “oracle-only” overrides UX to support multiple command prefixes (codex/gemini/task-master/tm)

Type:
- enhancement

Target Area:
- oraclepack core/TUI overrides and validation logic (oracle-specific detection/injection/validation)

Summary:
- Oraclepack can run any shell command, but its “nice UX features” (flag overrides + validation) are oracle-specific.
- Optionally generalize the command detection and flag-injection mechanisms so overrides/validation can apply to non-oracle tools used by Action Packs.

In Scope:
- Generalize `ExtractOracleInvocations` / `InjectFlags` to a small registry of command prefixes:
  - `oracle`, `codex`, `gemini`, `task-master`, `tm`
- Add per-tool override sets (so “Oracle Flags” semantics don’t incorrectly apply to other tools).
- Keep current oracle behavior intact.

Out of Scope:
- Implement-mode executor dispatch itself (handled in T3)
- Any new validation semantics beyond “oracle --dry-run summary” analogy (not provided)

Current Behavior (Actual):
- Special logic only for lines that start with `oracle`:
  - Invocation detection via regex anchored to literal `oracle`
  - Flag injection only into `oracle …` lines
  - Overrides validation runs `oracle --dry-run summary` only for detected oracle invocations

Expected Behavior:
- oraclepack recognizes multiple command prefixes for the purposes of overrides/flag handling and (where applicable) validation hooks, without breaking existing oracle behavior.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must not regress oracle detection/injection/validation behavior.
- Registry-based handling for additional tool prefixes as listed in the ticket text.

Evidence:
- References: “special logic only for lines that start with `oracle`… detects… injects… validation runs `oracle --dry-run summary`…”
- References: “Optional: improve oraclepack UX… registry of command prefixes… Add per-tool override sets…”

Open Items / Unknowns:
- Exact code locations for `ExtractOracleInvocations` / `InjectFlags` and override sets in the repo (paths not provided in this ticket text).

Risks / Dependencies:
- Not required to make Action Packs execute non-oracle tools; explicitly described as “nice-to-have”.
- Potential semantic mismatch if oracle-style overrides are applied to other CLIs without per-tool override sets.

Acceptance Criteria:
- oraclepack supports a registry of command prefixes including: `oracle`, `codex`, `gemini`, `task-master`, `tm`.
- Overrides/flag injection is not hard-coded to only `oracle` command lines.
- Per-tool override sets exist (or equivalent structure) so overrides are not incorrectly treated as “Oracle Flags” for non-oracle tools.
- Existing oracle-only behavior remains functional.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Oraclepack executes each step as shell, but it has special logic only for lines that start with `oracle`”
- “Generalize `ExtractOracleInvocations` / `InjectFlags`… registry of command prefixes (`oracle`, `codex`, `gemini`, `task-master`, `tm`)”
- “Add per-tool override sets…”
```

[1]: https://developers.openai.com/codex/cli/reference/?utm_source=chatgpt.com "Command line options"
[2]: https://geminicli.com/docs/tools/?utm_source=chatgpt.com "Gemini CLI tools"
[3]: https://github.com/eyaltoledano/claude-task-master?utm_source=chatgpt.com "eyaltoledano/claude-task-master"
```

.tickets/Oraclepack Pipeline Improvements.md
```
Title:

* Implement deterministic oraclepack pipeline improvements (strict validation, run manifests, resume/caching, Stage 3 “Actionizer”)

Summary:

* The current two-stage oraclepack workflow (Stage 1 pack generation → Stage 2 execution) is “weakly connected” and lacks deterministic handoff metadata, robust resume/retry, and an automated Stage 3 that converts 20 outputs into actionable engineering work.

    Oracle Pack Workflow Analysis

* This ticket proposes additive, backward-compatible enhancements to oraclepack and the Stage 1 generator prompts so runs are reproducible, CI-friendly, and produce machine-readable artifacts suitable for automation.

    Oracle Pack Workflow Analysis

Background / Context:

* Workflow context:

  * Stage 1: Codex skill or Gemini CLI slash command generates a single Markdown oracle question pack under `docs/*oracle-pack*.md`, following a strict oraclepack schema and containing exactly 20 `oracle ...` commands.

        Oracle Pack Workflow Analysis

  * Stage 2: oraclepack (Go wrapper around `@steipete/oracle`) executes the 20 commands and writes per-question outputs (via `--write-output`).

        Oracle Pack Workflow Analysis

  * Stage 3 is currently missing: outputs are not automatically turned into actionable implementation work.

        Oracle Pack Workflow Analysis

* Non-negotiable constraints:

  * No schema-breaking changes to the oraclepack Markdown pack schema without a backward-compatible migration path and validator-safe proof.

  * Automation must be deterministic and reproducible (no interactive steps in the critical path).

  * Stage 1 output must remain a single-pack deliverable that oraclepack can validate/run (no extra blocks/headers; no schema drift).

  * Prefer minimal file attachments per question; avoid broad globs unless unavoidable.

  * Optimize for longer runtimes with minimal human interaction (batching, resume/retry, caching, stable outputs, CI-friendly).

        Oracle Pack Workflow Analysis

Current Behavior (Actual):

* Stage 1 (generation) failure modes / friction points:

  * Packs can drift from the strict schema (extra fenced blocks, step-like headers, missing fields, wrong ordering, wrong count ≠ 20), causing ingestion/validation issues.

        Oracle Pack Workflow Analysis

  * Attachments may be bloated (broad globs, “just in case” files), increasing token cost and reducing signal-to-noise.

        Oracle Pack Workflow Analysis

  * ROI scoring can be inconsistent (unstable prioritization vs stated rationale).

        Oracle Pack Workflow Analysis

  * Coverage duplication across 20 questions (overlapping targets) wastes runs/budget.

        Oracle Pack Workflow Analysis

* Stage 2 (execution) failure modes / friction points:

  * Resume/retry semantics are weak (reruns may re-execute completed steps; partial failures require manual selection).

        Oracle Pack Workflow Analysis

  * Output determinism gaps: inconsistent output paths/slugs/out\_dir naming undermine CI diffs and Stage 3 discovery.

        Oracle Pack Workflow Analysis

  * Concurrency/rate limiting is not first-class (provider throttling/timeouts lead to nondeterministic failures).

        Oracle Pack Workflow Analysis

* Cross-stage handoff issues:

  * Missing traceability between pack ↔ outputs (no explicit manifest tying outputs to pack hash, git SHA, tool versions, provider/model settings).

        Oracle Pack Workflow Analysis

  * Stage 2 may be bypassed (pack executed “by hand”), losing wrapper state/report and consistent run directory.

        Oracle Pack Workflow Analysis

Expected Behavior:

* Stage 1 packs are always validator-safe and deterministic (single pack, exactly 20 oracle invocations, no schema drift).

    Oracle Pack Workflow Analysis

* Stage 2 produces stable, discoverable, machine-readable run artifacts that bind pack ↔ outputs and enable idempotent resume/rerun.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) exists and deterministically converts the 20 outputs into actionable engineering work artifacts (backlog + change plan + optional issue export), without duplicating work on reruns.

    Oracle Pack Workflow Analysis

* CI can run validate → run → actionize non-interactively with structured outputs and policy-driven exit codes.

    Oracle Pack Workflow Analysis

Requirements:

* Validation / linting (additive, backward-compatible):

  * Add `oraclepack validate --strict --json` that reports counts (steps=20, bash\_blocks=1, oracle\_invocations=20), ordering checks (ROI desc; ties effort asc), and required fields presence.

        Oracle Pack Workflow Analysis

* Deterministic run directory + manifest:

  * On `run`, create `.oraclepack/runs/<pack_id>/` and emit `run.json` + `steps.json`.

        Oracle Pack Workflow Analysis

  * `pack_id = YYYY-MM-DD__<gitshort>__<packhash8>`.

        Oracle Pack Workflow Analysis

  * `run.json` must include: `pack_path`, `pack_hash`, `git_sha`, `oraclepack_version`, `oracle_version`, `created_at`.

        Oracle Pack Workflow Analysis

  * `steps.json` must include: `step_id` (01..20), `reference`, `category`, `roi`, `command_hash`, `output_path`, `output_hash`, `status` (pending|ok|failed|skipped).

        Oracle Pack Workflow Analysis

* Resume/rerun semantics:

  * Make resume default: if `run.json` exists, skip steps whose output exists and matches recorded hash.

  * Support explicit overrides: `--rerun all|failed|01,03,07`.

        Oracle Pack Workflow Analysis

* Concurrency and rate limiting:

  * Add global `--max-parallel N` and optionally per-provider caps via config.

  * Implement exponential backoff + jitter on transient errors (e.g., 429/503) with a retry budget.

        Oracle Pack Workflow Analysis

* Deterministic caching (optional initially):

  * Implement caching keyed by `sha256(prompt + attached_file_hashes + oracle_flags + model)`, stored under `.oraclepack/cache/<invocation_key>.md`; rerun reuses cached outputs when key matches.

        Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) design and artifacts:

  * Implement `oraclepack actionize --run-dir .oraclepack/runs/<pack_id>`.

  * Inputs: `run.json` + 20 outputs under `.oraclepack/runs/<pack_id>/outputs/`.

        Oracle Pack Workflow Analysis

  * Deterministic processing: normalize outputs → dedupe → categorize via fixed taxonomy → generate action tasks, including blocked/conflict handling.

        Oracle Pack Workflow Analysis

  * Outputs under `.oraclepack/runs/<pack_id>/actionizer/`:

    * `normalized.jsonl`, `backlog.md`, `change-plan.md`

    * Optional: `github-issues.json`, `taskmaster.json`.

            Oracle Pack Workflow Analysis

  * Idempotency: stable IDs derived from `pack_hash` (e.g., `OP3-<packhash8>-<issue_index>-<task_index>`), stable paths, byte-identical regeneration when inputs unchanged.

        Oracle Pack Workflow Analysis

* Stage 1 prompt/skill improvements (without breaking schema):

  * Embed structured mini-metadata inside each `-p` prompt text (not new pack headers), e.g., `QuestionId`, `Category`, `Reference`, `ExpectedArtifacts`.

        Oracle Pack Workflow Analysis

  * Enforce deterministic attachment minimization heuristics (reference file + 0–2 neighbors; avoid broad globs unless evidence demands).

        Oracle Pack Workflow Analysis

  * Standardize generator prompt across Codex skills and Gemini CLI commands using a single canonical prompt file in repo.

        Oracle Pack Workflow Analysis

* CI-native mode:

  * Provide `oraclepack run --ci --non-interactive --json-log` and `oraclepack actionize --ci`.

  * CI policy can fail build if validation fails, completion rate below threshold, or action yield below threshold (threshold values: Not provided).

        Oracle Pack Workflow Analysis

* Security/safety:

  * Path safety: prevent `--write-output` from escaping run dir (reject `..` traversal).

        Oracle Pack Workflow Analysis

Out of Scope:

* Breaking changes to the existing oraclepack Markdown pack schema (unless a backward-compatible migration path and validator-safe proof are provided).

    Oracle Pack Workflow Analysis

Reproduction Steps:

1. Generate a pack via Stage 1 and save to `docs/oracle-pack-YYYY-MM-DD.md`.

    Oracle Pack Workflow Analysis

2. Run `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md` and observe schema drift / strictness gaps (exact current validator behavior: Unknown).

    Oracle Pack Workflow Analysis

3. Execute the pack, interrupt mid-run, rerun, and observe whether completed steps are skipped (current behavior: weak/unclear).

    Oracle Pack Workflow Analysis

4. Compare two runs on the same commit and observe output path/slug stability and traceability artifacts (manifest missing today).

    Oracle Pack Workflow Analysis

Environment:

* Tooling:

  * oraclepack (Go wrapper around `@steipete/oracle`).

        Oracle Pack Workflow Analysis

  * Stage 1 generators: Codex skills or Gemini CLI slash commands.

        Oracle Pack Workflow Analysis

* Repository/OS/versions: Unknown (git SHA, oraclepack version, oracle version, provider/model settings not provided in the conversation; also identified as missing traceability today).

    Oracle Pack Workflow Analysis

Evidence:

* Proposed stable artifact layout and handoff contract:

    Oracle Pack Workflow Analysis

  * `docs/oracle-pack-YYYY-MM-DD.md`

  * `.oraclepack/runs/<pack_id>/run.json`

  * `.oraclepack/runs/<pack_id>/steps.json`

  * `.oraclepack/runs/<pack_id>/outputs/01-<slug>.md … 20-<slug>.md`

  * `.oraclepack/runs/<pack_id>/actionizer/{normalized.jsonl, backlog.md, change-plan.md}`

* Proposed commands (some flag names explicitly “proposed where not already present”):

    Oracle Pack Workflow Analysis

  * `oraclepack validate --strict docs/oracle-pack-YYYY-MM-DD.md --json > .oraclepack/validate.json`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --max-parallel 4 --resume --ci`

  * `oraclepack actionize --run-dir .oraclepack/runs/<pack_id> --ci`

* Example Stage 3 output record shape (JSONL line):

    Oracle Pack Workflow Analysis

Decisions / Agreements:

* Do not break the oraclepack Markdown pack schema; any change must be backward-compatible with a validator-safe proof.

    Oracle Pack Workflow Analysis

* Stage 3 (“Actionizer”) is required and should be implemented as a first-class oraclepack subcommand (`actionize`) producing deterministic artifacts with idempotent reruns.

    Oracle Pack Workflow Analysis

* Traceability and determinism should be achieved via additive sidecar files (e.g., `run.json`, `steps.json`) rather than pack schema changes.

    Oracle Pack Workflow Analysis

Open Items / Unknowns:

* Current oraclepack CLI surface area:

  * Whether `validate --strict`, `--json`, `run --ci`, `--resume`, `--json-log`, and `actionize` exist today vs need implementation (conversation notes some flags are “proposed”).

        Oracle Pack Workflow Analysis

* Current on-disk state/report formats and locations (“state lives today (intended): oraclepack state/report + per-step outputs”; exact paths not provided).

    Oracle Pack Workflow Analysis

* Threshold definitions for CI policy (“completion rate < threshold”, “action yield < threshold”): Not provided.

    Oracle Pack Workflow Analysis

* Exact strict pack schema invariants enforced today (beyond “strict output contract” and “exactly 20” requirement): Not provided in this conversation (referenced as external inputs).

    Oracle Pack Workflow Analysis

Risks / Dependencies:

* Risk: filesystem layout changes may affect existing users; mitigation is additive behavior that preserves current out\_dir behavior.

    Oracle Pack Workflow Analysis

* Risk: caching correctness depends on hashing all attached file contents; incomplete hashing risks “cache poisoning.”

    Oracle Pack Workflow Analysis

* Risk: provider throttling/timeouts require robust transient-error classification for backoff/retry behavior.

    Oracle Pack Workflow Analysis

* Dependency: Stage 3 quality depends on stable, parseable structure in per-question outputs; mitigated by deterministic normalization heuristics and improved Stage 1 prompt shaping.

Acceptance Criteria:

* Validation:

  * `oraclepack validate --strict --json` deterministically reports schema invariants (20 steps, 20 oracle invocations, schema drift detection) and can gate CI.

        Oracle Pack Workflow Analysis

* Run determinism and traceability:

  * Running a pack produces `.oraclepack/runs/<pack_id>/{run.json,steps.json,outputs/}` with stable `pack_id` and stable output naming.

  * `run.json` includes required metadata fields; `steps.json` includes required per-step fields and statuses.

        Oracle Pack Workflow Analysis

* Resume/rerun:

  * Interrupting a run mid-way and rerunning resumes without re-executing completed steps (validated via output hashes and statuses).

  * `--rerun failed|all|<step list>` behaves as specified.

        Oracle Pack Workflow Analysis

* Concurrency/rate limiting:

  * `--max-parallel N` bounds concurrency; transient failures (e.g., throttling/timeouts) are retried with backoff within a retry budget and recorded in step status.

        Oracle Pack Workflow Analysis

* Caching (if implemented):

  * Rerunning on unchanged inputs (same prompt, same attached file digests, same flags/model) results in zero provider calls and identical outputs.

        Oracle Pack Workflow Analysis

* Stage 3 “Actionizer”:

  * `oraclepack actionize --run-dir ...` generates deterministic artifacts under `actionizer/` (`normalized.jsonl`, `backlog.md`, `change-plan.md`).

  * Reruns do not duplicate tasks (stable IDs) and produce byte-identical output when inputs unchanged.

  * Missing/contradictory answers produce explicit `blocked`/`conflict` tasks with required evidence patterns.

* CI mode:

  * `--ci --non-interactive --json-log` runs without TUI/interaction and uses structured logs and policy-driven exit codes.

        Oracle Pack Workflow Analysis

Priority & Severity (if inferable from text):

* Priority: Not provided

* Severity: Not provided

Labels (optional):

* enhancement, cli, determinism, validation, resume, caching, concurrency, workflow, stage3-actionizer

---
```

.tickets/Oraclepack Prompt Generator.md
```
Parent Ticket:

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
```

.tickets/Oraclepack Workflow Enhancement.md
```
Title:

* Stabilize oraclepack “oracle-pack” pipeline and add profile-based context + Stage-3 synthesis for actionable follow-through

Summary:

* The current two-step workflow generates an `oracle-pack` Markdown file (20 `oracle` calls) via Codex skills/Gemini CLI commands, then runs it through the `oraclepack` Go wrapper to produce outputs and state/report artifacts.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* A key failure mode is schema/format drift in the pack file (human-doc + machine-ingest combined), which can break ingestion; an example drift is step headers using an em dash (`# 01 — ROI=…`) while the documented contract expects `# NN)`.

    Workflow Optimization for Oracl…

* Requested outcome: improve workflow continuity, enable richer context injection without breaking the strict pack contract, and add an automatic next stage that turns the “final twenty questions/answers” into actionable implementation steps with minimal human interaction.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Background / Context:

* Workflow stages:

  * Stage 1: LLM authoring creates `docs/oracle-pack-YYYY-MM-DD.md` containing 20 `oracle` commands (with ROI metadata and per-step output paths).

        Workflow Optimization for Oracl…

  * Stage 2: `oraclepack` executes the pack to produce 20 outputs under `oracle-out/...` plus state/report JSON artifacts.

        Workflow Optimization for Oracl…

* `oraclepack` is a wrapper around `oracle` intended for batched/bulk requests.

    Workflow Optimization for Oracl…

* Core concern: “disconnection” after the 20-question output; desire to chain into a useful, actionable implementation stage.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

Current Behavior (Actual):

* Pack file acts as both:

  * Human documentation, and

  * A strict machine-ingest contract, making formatting drift a pipeline-breaking event.

        Workflow Optimization for Oracl…

* Documented/expected step-header schema (`# NN)`) can drift to alternative formats (example: `# 01 — ROI=…`), risking parse/validation failures.

    Workflow Optimization for Oracl…

* High-risk edits include adding additional code fences (especially additional bash fences) or introducing lines that accidentally match the step-header pattern.

    Workflow Optimization for Oracl…

Expected Behavior:

* Packs remain schema-stable and reliably parse/validate/run across providers (Codex skills, Gemini CLI commands).

    Workflow Optimization for Oracl…

* Richer “skill context” can be injected without changing the pack’s ingest shape (no added code fences / no header drift).

    Workflow Optimization for Oracl…

* After Stage 2 produces 20 outputs + report JSON, a subsequent stage can automatically convert results into actionable implementation steps.

    Workflow Optimization for Oracl…

Requirements:

* Preserve the non-negotiable pack contract:

  * Pack is Markdown containing exactly one `bash` code block; the first bash block is executed.

        Workflow Optimization for Oracl…

  * Steps are identified via header pattern `# NN)` with sequential numbering starting at `01`.

        Workflow Optimization for Oracl…

  * Prelude content before the first step header executes once.

        Workflow Optimization for Oracl…

* Standardize Stage-1 generation to the strict header form `# NN)` (avoid em dash variants).

    Workflow Optimization for Oracl…

* Add a hard gate between Stage 1 and Stage 2:

  * Make `oraclepack validate` mandatory before `oraclepack run` (prevent schema drift reaching execution).

        Workflow Optimization for Oracl…

* Provide schema-safe extensibility for context injection:

  * Allow context to be injected via `oracle -p` prompt text and/or `oracle -f` file/directory attachments (preferred for larger context).

        Workflow Optimization for Oracl…

  * Use prelude variables and templating only if it does not interfere with header parsing.

        Workflow Optimization for Oracl…

  * Avoid adding extra code fences or lines resembling step headers.

        Workflow Optimization for Oracl…

* Implement “Context Profiles” as file-backed bundles:

  * Add `skills/oracle-pack/references/profiles/<name>.md` and inject via `oracle -f "$profile_file"` plus a short prompt preamble (“Follow the attached profile standards”).

  * Add an optional `profile` input to the Stage-1 skill/command, with backwards-compatible behavior when absent.

        Workflow Optimization for Oracl…

* Add a first-class Stage 3 synthesis step:

  * Provide a command shape such as `oraclepack synthesize --in oracle-out --report pack.report.json --out docs/implementation-pack-YYYY-MM-DD.md` that reads the 20 outputs, extracts proposed changes/file targets/tests, and emits a new validated pack for implementation.

        Workflow Optimization for Oracl…

  * Support minimal-interaction execution for Stage 3 (e.g., headless usage via Codex/Gemini CLI).

        Workflow Optimization for Oracl…

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* `oraclepack` Go program wrapping `oracle` CLI.

* Stage-1 generation tools mentioned: Codex skills, Gemini CLI commands.

    Workflow Optimization for Oracl…

* OS/CI details: Unknown.

Evidence:

* Attachment: “Workflow Optimization for Oraclepack.md”.

    Workflow Optimization for Oracl…

    Workflow Optimization for Oracl…

* Example schema drift called out: step headers using `# 01 — ROI=…` vs documented `# NN)`.

    Workflow Optimization for Oracl…

* Proposed validation/run sequence:

  * `oraclepack validate docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack list docs/oracle-pack-YYYY-MM-DD.md`

  * `oraclepack run docs/oracle-pack-YYYY-MM-DD.md --no-tui --run-all --stop-on-fail=true --out-dir oracle-out`

        Workflow Optimization for Oracl…

Decisions / Agreements:

* Treat the pack as a stable intermediate representation (IR) and keep context flowing only through `-p` and `-f` to avoid breaking the ingest contract.

* Prefer “Context Profiles” as a file-backed mechanism located under `skills/oracle-pack/references/profiles/`.

* Add a validation gate (`validate` before `run`) to reduce pipeline breakage from formatting drift.

    Workflow Optimization for Oracl…

Open Items / Unknowns:

* Exact current parser/validator behavior regarding em dash header variants (whether it currently accepts them) is not provided; only that it is avoidable schema drift.

    Workflow Optimization for Oracl…

* Exact filenames/paths for current `SKILL.md` and template files in the repo are referenced conceptually but not provided in full.

    Workflow Optimization for Oracl…

* Whether `oraclepack synthesize` already exists or is a new feature request is not provided; it is described as a proposed product shape.

    Workflow Optimization for Oracl…

Risks / Dependencies:

* Dependency on `oracle` CLI flags and behavior (`-p/--prompt`, `-f/--file`, `--write-output`, `--files-report`, `--dry-run`).

    Workflow Optimization for Oracl…

* Risk of pack invalidation from added code fences, additional bash blocks, or accidental header-like lines.

    Workflow Optimization for Oracl…

* Cross-provider consistency risk (Codex skills vs Gemini CLI commands) unless Stage 1 is standardized around a shared template/profile mechanism.

    Workflow Optimization for Oracl…

Acceptance Criteria:

* Pack schema stability

  * Packs validate when they contain exactly one bash block and step headers are strictly `# NN)` starting at `01`.

  * Stage-1 generation output uses `# NN)` (no em dash variant) across providers.

        Workflow Optimization for Oracl…

* Validation gate

  * Workflow includes a required `oraclepack validate` pass before any `oraclepack run`.

        Workflow Optimization for Oracl…

* Context Profiles

  * A `profile` selection results in `oracle -f "$profile_file"` being attached per step without adding new code fences or breaking parsing.

        Workflow Optimization for Oracl…

  * Absence of `profile` preserves current behavior (backwards compatible).

        Workflow Optimization for Oracl…

* Stage 3 synthesis

  * A synthesis step can consume `oracle-out` outputs + report JSON and emit a follow-on implementation pack intended to be validated and run.

        Workflow Optimization for Oracl…

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* workflow

* cli

* parsing

* validation

* context-bundles

* automation

---
```

.tickets/Verbose Payload Rendering TUI.md
```
Title:

* Add verbose payload rendering in TUI to display full prepared scripts/flags for oraclepack runs

Summary:

* The TUI should support a verbose mode that prints the full “prepared payload” being executed for oraclepack runs, including effective flags (post overrides and `--chatgpt-url` injection) and the entire script passed to execution.

    Verbose TUI Payload Rendering

* This is needed to verify exactly what payloads are being sent/executed during oraclepack runs and to support tests that confirm the rendered payload contents.

    Verbose TUI Payload Rendering

Background / Context:

* Proposed approach: add a reusable “prepared payload” layer to `internal/exec.Runner` (prepare prelude/step scripts after overrides + flag injection + sanitization), then have the TUI emit these prepared payload blocks to its log viewport immediately before execution.

    Verbose TUI Payload Rendering

* Files implicated by the proposal include `internal/exec/runner.go`, `internal/tui/tui.go`, `internal/cli/run.go`, plus new helpers/tests under `internal/tui/` and `internal/exec/`.

    Verbose TUI Payload Rendering

Current Behavior (Actual):

* The TUI does not provide a verbose rendering that shows the full prepared payload (full script + effective flags + extracted `oracle …` invocations) being executed for oraclepack runs.

    Verbose TUI Payload Rendering

Expected Behavior:

* When verbose payload logging is enabled, the TUI log viewport prints a payload block before each step runs that includes: effective oracle flags, extracted oracle invocations (full lines), and the full prepared script that will be executed.

    Verbose TUI Payload Rendering

* Verbose payload logging can be enabled via CLI flag (e.g., `--verbose-payload` with `-v`) and toggled in the TUI via a keybinding (proposed: `p`).

    Verbose TUI Payload Rendering

Requirements:

* Exec runner: expose “prepared payload” APIs:

  * `PreparePreludePayload(p *pack.Prelude) PreparedPreludePayload`

  * `PrepareStepPayload(s *pack.Step) PreparedStepPayload`

  * `RunPreparedPrelude(...)` / `RunPreparedStep(...)` to execute the prepared scripts.

        Verbose TUI Payload Rendering

* Prepared payload structures must include:

  * `Script` (sanitized, post injection),

  * `EffectiveFlags` (for steps),

  * `OracleInvocations` extracted from the prepared script,

  * sanitizer `Warnings`.

        Verbose TUI Payload Rendering

* TUI formatting helper:

  * Add `internal/tui/verbose_payload.go` to format payload blocks (header, effective flags, oracle invocations, then full script).

        Verbose TUI Payload Rendering

* TUI integration:

  * Add a `verbosePayload bool` toggle to the TUI model.

  * In the run flow, call `PrepareStepPayload` and, when enabled, push formatted payload lines into `logChan` before `RunPreparedStep`.

  * Add keybinding `p` to toggle `verbosePayload`.

        Verbose TUI Payload Rendering

* CLI wiring:

  * Add `--verbose-payload` / `-v` flag and pass it into `tui.NewModel(..., verbosePayload)`.

        Verbose TUI Payload Rendering

* Tests:

  * New `internal/exec/runner_payload_test.go` verifying prepared payload includes effective flags and injected oracle command text.

  * New `internal/tui/verbose_payload_test.go` verifying formatted lines include flags, invocation, and script content.

  * Update existing TUI tests to include the new `NewModel` arg.

        Verbose TUI Payload Rendering

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Go (per referenced `.go` files and packages).

    Verbose TUI Payload Rendering

* TUI framework: Bubble Tea (`tea.NewProgram(...)` referenced).

    Verbose TUI Payload Rendering

* OS / terminal / versions: Unknown.

Evidence:

* Proposed change list and implementation sketch in: /mnt/data/Verbose TUI Payload Rendering.md

    Verbose TUI Payload Rendering

* Proposed file tree changes:

  * `internal/exec/runner.go` (modify)

  * `internal/exec/runner_payload_test.go` (new)

  * `internal/tui/verbose_payload.go` (new)

  * `internal/tui/verbose_payload_test.go` (new)

  * `internal/tui/tui.go` (modify)

  * `internal/cli/run.go` (modify)

  * Update TUI tests to pass new model arg.

        Verbose TUI Payload Rendering

Decisions / Agreements:

* Adopt a “prepared payload” abstraction in `exec.Runner` to ensure the TUI logs exactly what will run after overrides, injection, and sanitization.

    Verbose TUI Payload Rendering

* Add both CLI flag control (`--verbose-payload` / `-v`) and an in-TUI toggle (proposed key: `p`).

    Verbose TUI Payload Rendering

Open Items / Unknowns:

* Exact existing TUI run flow for prelude execution (whether/where prelude runs in TUI) is not provided; proposal notes “if you also execute a prelude… do the same.”

    Verbose TUI Payload Rendering

* Exact current `NewModel(...)` signature call sites and all affected tests/files beyond those listed are not fully enumerated (some are referenced as examples).

    Verbose TUI Payload Rendering

Risks / Dependencies:

* Not provided.

Acceptance Criteria:

* Running the TUI with `--verbose-payload` causes each executed step to prepend a log block that includes:

  * “payload (step <id>)” header,

  * “effective oracle flags: …” line,

  * extracted “oracle invocations:” section (or explicit none found),

  * full “script:” content (not truncated),

  * “end payload” footer.

        Verbose TUI Payload Rendering

* Toggling `p` in the TUI flips payload logging on/off for subsequent step executions.

    Verbose TUI Payload Rendering

* `Runner.PrepareStepPayload` produces:

  * effective flags reflecting overrides and `--chatgpt-url`,

  * a prepared script containing the injected oracle invocation with those flags.

        Verbose TUI Payload Rendering

* New tests (`runner_payload_test.go`, `verbose_payload_test.go`) pass, and existing TUI tests compile and pass after updating `NewModel` call signature.

    Verbose TUI Payload Rendering

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* logging

* exec-runner

* cli

* testing

* go

---
```

docs/oracle-actions-pack-2026-01-07.md
```
# Oraclepack Stage 3 — Action Pack (Taskify)

Generated: 2026-01-07

Parsed args (resolved):
- stage2_path: auto
- out_dir: auto
- pack_path: docs/oracle-actions-pack-2026-01-07.md
- actions_json: auto/_actions.json
- actions_md: auto/_actions.md
- prd_path: .taskmaster/docs/oracle-actions-prd.md
- tag: oraclepack
- mode: autopilot
- top_n: 10
- oracle_cmd: oracle
- task_master_cmd: task-master
- tm_cmd: tm
- extra_files: (embedded where applicable)

This document must contain exactly one bash code fence, and no other code fences.

```bash
# 01) Preflight: resolve Stage 2 inputs and verify tools (fail fast)
set -euo pipefail

STAGE2_PATH="auto"
OUT_DIR="auto"
MODE="autopilot"
TASK_MASTER_CMD="task-master"
ORACLE_CMD="oracle"
TM_CMD="tm"
ACTIONS_JSON="auto/_actions.json"
ACTIONS_MD="auto/_actions.md"
PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"

check_dir_complete() {
  local d="$1"
  [ -d "$d" ] || return 1
  local n
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    local matches=( "$d/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -ne 1 ]; then
      return 1
    fi
  done
  return 0
}

resolve_stage2_auto() {
  local candidate

  candidate="oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi

  candidate="docs/oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi

  for candidate in $(ls -1d docs/oracle-questions-*/oracle-out/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done

  for candidate in $(ls -1d docs/oracle-questions-*/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done

  candidate=$(ls -1 docs/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oracle-questions-*/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  candidate=$(ls -1 docs/oracle-questions-*/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi

  return 1
}

resolve_out_dir() {
  local stage2="$1"
  if [ "$OUT_DIR" != "auto" ]; then
    echo "$OUT_DIR"
    return 0
  fi
  if [ -d "$stage2" ]; then
    echo "$stage2"
    return 0
  fi
  if [[ "$stage2" == docs/oracle-questions-*/* ]]; then
    local base
    base=$(echo "$stage2" | awk -F/ '{print $1"/"$2}')
    echo "$base/oracle-out"
    return 0
  fi
  echo "oracle-out"
}

STAGE2_SOURCE=""
if [ "$STAGE2_PATH" != "auto" ]; then
  if [ -d "$STAGE2_PATH" ] || [ -f "$STAGE2_PATH" ]; then
    STAGE2_SOURCE="$STAGE2_PATH"
  else
    echo "ERROR: stage2_path does not exist: $STAGE2_PATH" >&2
    exit 2
  fi
else
  if ! STAGE2_SOURCE=$(resolve_stage2_auto); then
    echo "ERROR: could not resolve Stage 2 outputs via auto search." >&2
    echo "Searched in order: oracle-out/, docs/oracle-out/, docs/oracle-questions-*/oracle-out/, docs/oracle-questions-*/, docs/oracle-pack-*.md, docs/oraclepacks/oracle-pack-*.md, docs/oracle-questions-*/oracle-pack-*.md, docs/oracle-questions-*/oraclepacks/oracle-pack-*.md" >&2
    exit 3
  fi
fi

if [ -d "$STAGE2_SOURCE" ]; then
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    matches=( "$STAGE2_SOURCE/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -eq 0 ]; then
      echo "ERROR: missing oracle output for prefix ${n}: expected ${STAGE2_SOURCE}/${n}-*.md" >&2
      exit 4
    fi
    if [ "${#matches[@]}" -gt 1 ]; then
      echo "ERROR: multiple oracle outputs for prefix ${n}; expected exactly one file." >&2
      printf '%s\n' "${matches[@]}" >&2
      exit 5
    fi
  done
fi

OUT_DIR_RESOLVED=$(resolve_out_dir "$STAGE2_SOURCE")

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -d "$OUT_DIR_RESOLVED" ]; then
    ACTIONS_JSON="$OUT_DIR_RESOLVED/_actions.json"
    ACTIONS_MD="$OUT_DIR_RESOLVED/_actions.md"
  else
    ACTIONS_JSON="docs/_actions.json"
    ACTIONS_MD="docs/_actions.md"
  fi
fi

if ! command -v "${TASK_MASTER_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${TASK_MASTER_CMD%% *} (from task_master_cmd='${TASK_MASTER_CMD}')" >&2
  exit 10
fi

if ! command -v "${ORACLE_CMD%% *}" >/dev/null 2>&1; then
  echo "ERROR: required tool missing: ${ORACLE_CMD%% *} (from oracle_cmd='${ORACLE_CMD}')" >&2
  exit 11
fi

if [ "$MODE" = "autopilot" ]; then
  if ! command -v "${TM_CMD%% *}" >/dev/null 2>&1; then
    echo "ERROR: autopilot mode requires tm_cmd, but tool missing: ${TM_CMD%% *} (from tm_cmd='${TM_CMD}')" >&2
    echo "HINT: rerun Stage 3 with mode=backlog if you only want tasks generated." >&2
    exit 12
  fi
fi

mkdir -p "$(dirname "$ACTIONS_JSON")" "$(dirname "$ACTIONS_MD")" "$(dirname "$PRD_PATH")" "docs"
if [ -d "$OUT_DIR_RESOLVED" ]; then mkdir -p "$OUT_DIR_RESOLVED"; fi

cat > "_actions.schema.md" <<'SCHEMA'
# Canonical Actions JSON Schema (human-readable)
[... schema details omitted for brevity, see asset ...]
SCHEMA
mv "_actions.schema.md" "docs/_actions.schema.md"

echo "OK: Stage2 source: ${STAGE2_SOURCE}"
echo "OK: Out dir: ${OUT_DIR_RESOLVED}"
echo "OK: Mode: ${MODE}"


# 02) Synthesize canonical actions JSON + summary MD
set -euo pipefail

STAGE2_PATH="auto"
OUT_DIR="auto"
ACTIONS_JSON="auto/_actions.json"
ACTIONS_MD="auto/_actions.md"

check_dir_complete() {
  local d="$1"
  [ -d "$d" ] || return 1
  local n
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    shopt -s nullglob
    local matches=( "$d/${n}-"*.md )
    shopt -u nullglob
    if [ "${#matches[@]}" -ne 1 ]; then
      return 1
    fi
  done
  return 0
}

resolve_stage2_auto() {
  local candidate
  candidate="oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  candidate="docs/oracle-out"
  if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  for candidate in $(ls -1d docs/oracle-questions-*/oracle-out/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done
  for candidate in $(ls -1d docs/oracle-questions-*/ 2>/dev/null | sort -r); do
    if check_dir_complete "$candidate"; then echo "$candidate"; return 0; fi
  done
  candidate=$(ls -1 docs/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oracle-questions-*/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  candidate=$(ls -1 docs/oracle-questions-*/oraclepacks/oracle-pack-*.md 2>/dev/null | sort | tail -n 1 || true)
  if [ -n "$candidate" ]; then echo "$candidate"; return 0; fi
  return 1
}

resolve_out_dir() {
  local stage2="$1"
  if [ "$OUT_DIR" != "auto" ]; then
    echo "$OUT_DIR"
    return 0
  fi
  if [ -d "$stage2" ]; then
    echo "$stage2"
    return 0
  fi
  if [[ "$stage2" == docs/oracle-questions-*/* ]]; then
    local base
    base=$(echo "$stage2" | awk -F/ '{print $1"/"$2}')
    echo "$base/oracle-out"
    return 0
  fi
  echo "oracle-out"
}

STAGE2_SOURCE=""
if [ "$STAGE2_PATH" != "auto" ]; then
  if [ -d "$STAGE2_PATH" ] || [ -f "$STAGE2_PATH" ]; then
    STAGE2_SOURCE="$STAGE2_PATH"
  else
    echo "ERROR: stage2_path does not exist: $STAGE2_PATH" >&2
    exit 20
  fi
else
  if ! STAGE2_SOURCE=$(resolve_stage2_auto); then
    echo "ERROR: could not resolve Stage 2 outputs via auto search." >&2
    exit 21
  fi
fi

OUT_DIR_RESOLVED=$(resolve_out_dir "$STAGE2_SOURCE")

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -d "$OUT_DIR_RESOLVED" ]; then
    ACTIONS_JSON="$OUT_DIR_RESOLVED/_actions.json"
    ACTIONS_MD="$OUT_DIR_RESOLVED/_actions.md"
  else
    ACTIONS_JSON="docs/_actions.json"
    ACTIONS_MD="docs/_actions.md"
  fi
fi

oracle_file_flags=()
if [ -f "$STAGE2_SOURCE" ]; then
  oracle_file_flags+=( -f "$STAGE2_SOURCE" )
else
  shopt -s nullglob
  for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
    matches=( "$STAGE2_SOURCE/${n}-"*.md )
    if [ "${#matches[@]}" -ne 1 ]; then
      echo "ERROR: expected exactly one match for ${STAGE2_SOURCE}/${n}-*.md, got ${#matches[@]}" >&2
      printf '%s\n' "${matches[@]:-}" >&2
      exit 22
    fi
    oracle_file_flags+=( -f "${matches[0]}" )
  done
  shopt -u nullglob
fi

mkdir -p "$(dirname "$ACTIONS_JSON")" "$(dirname "$ACTIONS_MD")"

oracle \
  --write-output "$ACTIONS_JSON" \
  "${oracle_file_flags[@]}" \
  -p "$(cat <<'PROMPT'
You are producing a SINGLE JSON document and nothing else.

Task: Normalize oraclepack answers into a canonical actionable plan.

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
- pack_date: 2026-01-07
- source_out_dir: auto
- top_n: 10
- tag: oraclepack

Output hygiene:
- Do not include backticks or fenced code blocks anywhere.
- Keep strings concise and specific.

Now produce the JSON.
PROMPT
)"

oracle \
  --write-output "$ACTIONS_MD" \
  -f "$ACTIONS_JSON" \
  -p "$(cat <<'PROMPT'
Write a human-readable Markdown summary of the canonical actions JSON.

Hard requirements:
- Output MUST be Markdown text with headings/bullets only (no code fences).
- Keep ordering aligned with items id ascending (01..20).
- Include:
  - short executive summary (5–10 bullets)
  - top 10 prioritized list (with id, title inferred from recommended_next_action, category, and why)
  - per-item: recommended_next_action + acceptance_criteria bullets + missing_artifacts bullets
- Do not invent facts; reflect only what is present in the JSON.

Now write the summary Markdown.
PROMPT
)"

echo "OK: Wrote $ACTIONS_JSON"
echo "OK: Wrote $ACTIONS_MD"


# 03) Generate PRD for Task Master
set -euo pipefail

ACTIONS_JSON="auto/_actions.json"
PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"

if [[ "$ACTIONS_JSON" == auto/* ]]; then
  if [ -f "docs/_actions.json" ]; then
    ACTIONS_JSON="docs/_actions.json"
  else
    found=$(find docs -name "_actions.json" | head -n 1 || true)
    if [ -n "$found" ]; then ACTIONS_JSON="$found"; fi
  fi
fi

mkdir -p "$(dirname "$PRD_PATH")"

oracle \
  --write-output "$PRD_PATH" \
  -f "$ACTIONS_JSON" \
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
- Use the tag value "oraclepack" in the PRD text where helpful for grouping.

Constants:
- TOP_N=10
- TAG=oraclepack

Now produce the PRD.
PROMPT
)"

echo "OK: Wrote $PRD_PATH"


# 04) Task Master: parse PRD into tasks
set -euo pipefail

PRD_PATH=".taskmaster/docs/oracle-actions-prd.md"
TASK_MASTER_CMD="task-master"
TAG="oraclepack"

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

TASK_MASTER_CMD="task-master"
OUT_DIR="auto"

if [[ "$OUT_DIR" == "auto" ]] || [[ ! -e "$OUT_DIR" ]]; then
  OUT_DIR="docs"
fi
if [ -f "$OUT_DIR" ]; then OUT_DIR="$(dirname "$OUT_DIR")"; fi

mkdir -p "$OUT_DIR"

"${TASK_MASTER_CMD}" analyze-complexity --output "$OUT_DIR/tm-complexity.json"
echo "OK: Wrote $OUT_DIR/tm-complexity.json"


# 06) Task Master: expand tasks
set -euo pipefail

TASK_MASTER_CMD="task-master"

"${TASK_MASTER_CMD}" expand --all
echo "OK: Expanded tasks."


# 07) Pipelines (pipelines mode only): generate deterministic pipelines from tasks.json
set -euo pipefail

MODE="autopilot"

if [ "$MODE" != "pipelines" ]; then
  echo "SKIP: mode=$MODE (pipelines step runs only when mode=pipelines)."
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

  oracle \
    --write-output "docs/oracle-actions-pipelines.md" \
    -f "$tasks_path" \
    -p "$(cat <<'PROMPT'
Generate deterministic command pipelines from tasks.json.

Hard requirements:
- Output MUST be Markdown (no code fences).
- Include 3–6 pipelines, each a numbered list of shell commands.
- Each pipeline must be tests-first and avoid destructive operations.
- Commands should be generic and repo-agnostic (no invented scripts).
- Include a short resume strategy section explaining how to re-run pipelines safely.

Now write docs/oracle-actions-pipelines.md content.
PROMPT
)"

  echo "OK: Wrote docs/oracle-actions-pipelines.md"
fi


# 08) Autopilot (autopilot mode only): branch safety + guarded entrypoint
set -euo pipefail

MODE="autopilot"
TM_CMD="tm"
OUT_DIR="auto"
PACK_DATE="2026-01-07"
TAG="oraclepack"

if [ "$MODE" != "autopilot" ]; then
  echo "SKIP: mode=$MODE (autopilot step runs only when mode=autopilot)."
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
  if [ "$current_branch" = "main" ] || [ "$current_branch" = "master" ]; then
    new_branch="oraclepack/${PACK_DATE}-${TAG}"
    echo "INFO: on default-like branch (${current_branch}); creating work branch: ${new_branch}"
    git checkout -b "$new_branch"
  else
    echo "OK: current branch is ${current_branch}"
  fi

  if [[ "$OUT_DIR" == "auto" ]] || [[ ! -d "$OUT_DIR" ]]; then OUT_DIR="docs"; fi
  mkdir -p "$OUT_DIR"

  cat > "$OUT_DIR/tm-autopilot.state.json" <<STATE
{"pack_date":"${PACK_DATE}","tag":"${TAG}","mode":"autopilot","notes":"State file created by Stage-3 Action Pack. Autopilot tooling should resume from this file if supported."}
STATE

  echo "OK: Wrote $OUT_DIR/tm-autopilot.state.json"
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

docs/oracle-pack-2026-01-07.md
```
# Oracle Pack — Unknown (Gold Stage 1)

## Parsed args
- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown
- out_dir: docs/oracle-questions-2026-01-07
- oracle_cmd: oracle
- oracle_flags: --files-report
- engine: None
- model: None
- extra_files: None

Notes:
- Template is the contract. Keep the pack runner-ingestible.
- Exactly one fenced bash block in this whole document.
- Exactly 20 steps, numbered 01..20.
- Each step includes: ROI= impact= confidence= effort= horizon= category= reference=
- Categories must be exactly the fixed set used in Coverage check.

## Commands
```bash
# Optional preflight pattern:
# - Add `--dry-run summary` to preview what will be sent, and keep `--files-report` enabled when available.

# 01) ROI=8.2 impact=9 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/cli/root.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md" \
  --file "internal/cli/root.go" \
  --file "internal/cli/cmds.go" \
  -p "$(cat <<'PROMPT'
Strategist question #01

Reference: internal/cli/root.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 8.2 (impact=9, confidence=0.9, effort=1)

Question:
Identify the primary public interface(s) of this system (CLI commands and the oracle-pack Markdown contract). For each, list the key inputs/outputs and where they are defined in code.

Rationale (one sentence):
We need a trustworthy map of the system’s “outside-facing contract” before deeper planning.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ROI=7.5 impact=8 confidence=0.9 effort=1 horizon=Immediate category=contracts/interfaces reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md" \
  --file "internal/exec/runner.go" \
  --file "internal/exec/oracle_validate.go" \
  -p "$(cat <<'PROMPT'
Strategist question #02

Reference: internal/exec/runner.go
Category: contracts/interfaces
Horizon: Immediate
ROI: 7.5 (impact=8, confidence=0.9, effort=1)

Question:
What are the top integration points with external systems (oracle CLI invocation, filesystem outputs, environment variables), and what contract(s) or config declare them? Provide the minimal list of files/locations that define each integration.

Rationale (one sentence):
Integration boundaries drive risk, deployment needs, and test strategy.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–6 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 03) ROI=6.8 impact=8 confidence=0.85 effort=1 horizon=NearTerm category=invariants reference=internal/pack/parser.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/03-invariants-domain.md" \
  --file "internal/pack/parser.go" \
  --file "internal/pack/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #03

Reference: internal/pack/parser.go
Category: invariants
Horizon: NearTerm
ROI: 6.8 (impact=8, confidence=0.85, effort=1)

Question:
List the system’s most important invariants about pack structure and step ordering. For each, show where it is enforced (or where it should be enforced but currently is not).

Rationale (one sentence):
Invariants define correctness and are the backbone of reliable changes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 04) ROI=6.2 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=invariants reference=internal/app/app.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/04-invariants-validation.md" \
  --file "internal/app/app.go" \
  --file "internal/pack/parser.go" \
  -p "$(cat <<'PROMPT'
Strategist question #04

Reference: internal/app/app.go
Category: invariants
Horizon: NearTerm
ROI: 6.2 (impact=7, confidence=0.8, effort=1)

Question:
Where does validation happen (pack parse/validate, config resolution, runtime checks)? Identify validation boundaries and the most likely gaps that could cause inconsistent state.

Rationale (one sentence):
Knowing validation boundaries prevents regressions and reduces correctness risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 05) ROI=6.0 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=caching/state reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/05-caching-state-layers.md" \
  --file "internal/state/types.go" \
  --file "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #05

Reference: internal/state/types.go
Category: caching/state
Horizon: NearTerm
ROI: 6.0 (impact=7, confidence=0.8, effort=1)

Question:
What stateful components exist (run state, status tracking, persisted state files)? For each, describe lifecycle, read/write points, and where it is implemented.

Rationale (one sentence):
State handling is a common source of subtle bugs and operational issues.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 06) ROI=5.8 impact=7 confidence=0.75 effort=1 horizon=NearTerm category=caching/state reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/06-caching-state-consistency.md" \
  --file "internal/app/run.go" \
  --file "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #06

Reference: internal/app/run.go
Category: caching/state
Horizon: NearTerm
ROI: 5.8 (impact=7, confidence=0.75, effort=1)

Question:
Identify the top consistency risks between in-memory run state and persisted state/report outputs (stale writes, partial updates, missing flushes). Where are the knobs/configs for state behavior?

Rationale (one sentence):
Consistency failure modes often surface as “random bugs” and are expensive to debug.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 07) ROI=5.0 impact=6 confidence=0.7 effort=1 horizon=NearTerm category=background jobs reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/07-background-jobs-discovery.md" \
  --file "internal/app/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #07

Reference: internal/app/run.go
Category: background jobs
Horizon: NearTerm
ROI: 5.0 (impact=6, confidence=0.7, effort=1)

Question:
What background jobs/workers/scheduled tasks exist, if any? For each, identify trigger mechanism, payload, retries, idempotency, and where it is defined. If none, confirm based on evidence.

Rationale (one sentence):
Background work affects reliability, cost, and operational complexity.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 08) ROI=4.2 impact=5 confidence=0.7 effort=1 horizon=NearTerm category=background jobs reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/08-background-jobs-reliability.md" \
  -p "$(cat <<'PROMPT'
Strategist question #08

Reference: Unknown
Category: background jobs
Horizon: NearTerm
ROI: 4.2 (impact=5, confidence=0.7, effort=1)

Question:
Where are the main reliability controls for background work (dead-lettering, backoff, concurrency limits, reprocessing), and what is missing or inconsistent? If there are no background jobs, explicitly state that and point to confirming evidence.

Rationale (one sentence):
Reliability controls prevent incident loops and data corruption.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 09) ROI=6.4 impact=8 confidence=0.8 effort=1 horizon=Immediate category=observability reference=internal/report/generate.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/09-observability-signals.md" \
  --file "internal/report/generate.go" \
  --file "internal/state/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #09

Reference: internal/report/generate.go
Category: observability
Horizon: Immediate
ROI: 6.4 (impact=8, confidence=0.8, effort=1)

Question:
What observability signals exist (reports, warnings, log lines), and what are the primary identifiers for correlating a run/step across components? Point to the code/config that defines them.

Rationale (one sentence):
You can’t operate or improve what you can’t measure or debug quickly.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 10) ROI=5.6 impact=7 confidence=0.8 effort=1 horizon=Immediate category=observability reference=internal/exec/runner.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/10-observability-gaps.md" \
  --file "internal/exec/runner.go" \
  --file "internal/errors/errors.go" \
  -p "$(cat <<'PROMPT'
Strategist question #10

Reference: internal/exec/runner.go
Category: observability
Horizon: Immediate
ROI: 5.6 (impact=7, confidence=0.8, effort=1)

Question:
Where are the biggest observability gaps (missing structured logs around key decisions, missing metrics for run outcomes, missing correlation IDs)? Recommend the smallest additions that would most improve debugging.

Rationale (one sentence):
Targeted observability improvements compound across all future changes.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 11) ROI=5.2 impact=7 confidence=0.75 effort=1 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/11-permissions-model.md" \
  -p "$(cat <<'PROMPT'
Strategist question #11

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 5.2 (impact=7, confidence=0.75, effort=1)

Question:
What is the permission model (roles/scopes/ACLs), and where is it defined? Provide the minimal set of files that encode “who can do what.” If no permission model exists, confirm that with evidence.

Rationale (one sentence):
Permission rules are a high-risk area with security and product impact.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 12) ROI=4.8 impact=6 confidence=0.8 effort=1 horizon=Immediate category=permissions reference=Unknown
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/12-permissions-enforcement.md" \
  -p "$(cat <<'PROMPT'
Strategist question #12

Reference: Unknown
Category: permissions
Horizon: Immediate
ROI: 4.8 (impact=6, confidence=0.8, effort=1)

Question:
Where are permissions enforced (middleware/guards/policies/service-layer checks), and where are likely bypass risks? Identify enforcement chokepoints and any inconsistent patterns.

Rationale (one sentence):
Enforcement consistency prevents privilege escalation and policy drift.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 13) ROI=4.5 impact=6 confidence=0.75 effort=1 horizon=NearTerm category=migrations reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/13-migrations-schema.md" \
  --file "internal/state/types.go" \
  --file "internal/state/persist.go" \
  -p "$(cat <<'PROMPT'
Strategist question #13

Reference: internal/state/types.go
Category: migrations
Horizon: NearTerm
ROI: 4.5 (impact=6, confidence=0.75, effort=1)

Question:
How are schema/config migrations handled (state schema versioning, report format evolution)? Identify the tooling, files, and how migrations are applied or would be applied.

Rationale (one sentence):
Migration mechanics are critical for safe releases and rollbacks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 14) ROI=4.1 impact=5 confidence=0.8 effort=1 horizon=NearTerm category=migrations reference=internal/state/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/14-migrations-compat.md" \
  --file "internal/state/types.go" \
  --file "internal/report/types.go" \
  -p "$(cat <<'PROMPT'
Strategist question #14

Reference: internal/state/types.go
Category: migrations
Horizon: NearTerm
ROI: 4.1 (impact=5, confidence=0.8, effort=1)

Question:
What are the backward/forward compatibility expectations for state/report schemas? Identify where compatibility is ensured or currently risky.

Rationale (one sentence):
Compatibility strategy prevents outages during upgrades and rollbacks.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–8 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 15) ROI=5.7 impact=7 confidence=0.8 effort=1 horizon=NearTerm category=UX flows reference=internal/tui/tui.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/15-ux-flows-primary.md" \
  --file "internal/tui/tui.go" \
  --file "internal/cli/run.go" \
  -p "$(cat <<'PROMPT'
Strategist question #15

Reference: internal/tui/tui.go
Category: UX flows
Horizon: NearTerm
ROI: 5.7 (impact=7, confidence=0.8, effort=1)

Question:
What are the primary user/operator flows in the TUI (run steps, run all, overrides wizard, URL selection)? Map each to main components/modules and key state transitions.

Rationale (one sentence):
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=5.1 impact=6 confidence=0.85 effort=1 horizon=NearTerm category=UX flows reference=internal/tui/overrides_flow.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/16-ux-flows-edgecases.md" \
  --file "internal/tui/overrides_flow.go" \
  --file "internal/tui/overrides_flags.go" \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: internal/tui/overrides_flow.go
Category: UX flows
Horizon: NearTerm
ROI: 5.1 (impact=6, confidence=0.85, effort=1)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, cancel/back navigation, partial completion, timeouts)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=6.1 impact=8 confidence=0.76 effort=1 horizon=Immediate category=failure modes reference=internal/app/run.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/17-failure-modes-taxonomy.md" \
  --file "internal/app/run.go" \
  --file "internal/errors/errors.go" \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: internal/app/run.go
Category: failure modes
Horizon: Immediate
ROI: 6.1 (impact=8, confidence=0.76, effort=1)

Question:
What is the failure-mode taxonomy of this system (pack validation errors, execution failures, state/report write failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=5.4 impact=7 confidence=0.77 effort=1 horizon=Immediate category=failure modes reference=internal/exec/sanitize.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/18-failure-modes-resilience.md" \
  --file "internal/exec/sanitize.go" \
  --file "internal/exec/runner.go" \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: internal/exec/sanitize.go
Category: failure modes
Horizon: Immediate
ROI: 5.4 (impact=7, confidence=0.77, effort=1)

Question:
What resilience mechanisms exist (sanitization, error wrapping, stop-on-fail behavior), and which critical paths lack them? Where are they configured?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=4.7 impact=6 confidence=0.78 effort=1 horizon=NearTerm category=feature flags reference=internal/overrides/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/19-feature-flags-inventory.md" \
  --file "internal/overrides/types.go" \
  --file "internal/tui/overrides_flags.go" \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: internal/overrides/types.go
Category: feature flags
Horizon: NearTerm
ROI: 4.7 (impact=6, confidence=0.78, effort=1)

Question:
What feature-flag-like controls exist (oracle flag overrides, per-step targeting, ROI filtering)? Inventory the flags/controls and identify where they are defined, evaluated, and documented.

Rationale (one sentence):
Flags and runtime controls enable safe rollout and experimentation and reduce release risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=4.3 impact=5 confidence=0.86 effort=1 horizon=NearTerm category=feature flags reference=internal/overrides/types.go
oracle \
  --files-report \
  --write-output "docs/oracle-questions-2026-01-07/20-feature-flags-rollout.md" \
  --file "internal/overrides/types.go" \
  --file "internal/overrides/merge.go" \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: internal/overrides/types.go
Category: feature flags
Horizon: NearTerm
ROI: 4.3 (impact=5, confidence=0.86, effort=1)

Question:
Describe the flag/rollout lifecycle for runtime overrides (how flags are selected, applied per-step, validated, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: None
Non-goals: None

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check
--------------

Mark each as OK or Missing(<which step ids>):

* contracts/interfaces: OK

* invariants: OK

* caching/state: OK

* background jobs: OK

* observability: OK

* permissions: OK

* migrations: OK

* UX flows: OK

* failure modes: OK

* feature flags: OK
```

.config/commands/oracle-pack_v2.toml
```
# path: commands/oraclepack/oracle-pack_v2.toml
# invoked via: /oraclepack:oracle-pack_v2 {{args}}

description = "Generate a strategist-questions Oracle pack that validates/runs under oraclepack (```bash fence + # NN) headers), writes oracle outputs under docs/, and preserves the Codex-style template (parsed args + ROI ordering + coverage check)."

prompt = """
You are generating output for a STRICT downstream CLI parser (oraclepack).
Your output MUST be parseable by oraclepack’s Markdown parser:
- The commands MUST be inside a ```bash fenced code block (triple backticks), not ~~~.
- Step headers MUST be exactly: # NN) ... (e.g., # 01) ...). Do NOT use em dashes (—) in place of the ')'.

User args (raw): {{args}}

------------------------------------------------------------
A) ARG PARSING (no follow-ups; apply defaults)
Extract from {{args}} if present (key=value or similar), else default:

- codebase_name: Unknown
- constraints: None
- non_goals: None
- team_size: Unknown
- deadline: Unknown

Oracle pack controls:
- out_dir: docs/oracle/strategist-questions/$(date +%F)
- oracle_cmd: oracle
- oracle_flags: --browser-attachments always --files-report
- extra_files: (empty; if present, treat as comma-separated file paths/globs to include as additional -f entries in EVERY command)

You MUST render these values into the final output under the “parsed args” section.

------------------------------------------------------------
B) INFERENCE-FIRST DISCOVERY (adaptive; evidence-driven; deterministic)
Goal: infer what exists before searching broadly.

1) Prefer cheap “index” artifacts first:
   - README / docs index (if present)
   - primary manifests (package.json, pyproject.toml, go.mod, Cargo.toml, etc.)
   - obvious entrypoints referenced by scripts/manifests

2) Derive a search plan from evidence:
   - follow imports/registrations from entrypoints into:
     routing/handlers, auth/permissions, jobs/queues, data layer/migrations,
     feature flags, observability/logging, caching/state, failure modes.

3) Harvest >= 20 candidate anchors, each as ONE of:
   - {path}:{symbol} (preferred)
   - {endpoint}
   - {event}
If a category cannot be evidenced, keep its anchor as Unknown AND record the missing artifact pattern.

Required coverage categories (must appear across your 20 questions):
- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

------------------------------------------------------------
C) QUESTION GENERATION (exactly 20; 12 Immediate + 8 Strategic)
For each question, produce:
- Reference: {path}:{symbol} OR {endpoint} OR {event} OR Unknown
- Category: one of the required categories (or a sensible subcategory that still maps to them)
- Horizon: Immediate or Strategic
- Question: focused, feasibility-first
- Rationale: EXACTLY one sentence
- Smallest experiment today: EXACTLY one concrete action

No duplicates by intent or reference.

------------------------------------------------------------
D) ROI SCORING + ORDERING (required)
For each question, choose:
- impact, confidence, effort ∈ {0.1..1.0} with one decimal
Compute:
- ROI = (impact * confidence) / effort

Sort ALL 20 commands by:
1) ROI descending
2) tie-breaker: lower effort first

Numbering remains 01..20 in this final sorted order.

------------------------------------------------------------
E) ORACLE COMMAND EMISSION + ATTACHMENTS (minimal; evidence-driven)
For each command:
- Use exactly:
  <oracle_cmd> \
    <oracle_flags> \
    --write-output "<out_dir>/<nn>-<slug>.md" \
    -p "<prompt>" \
    -f "<file 1>" \
    -f "<file 2>" \
    ... \
    <extra_files...>

Attachment minimization:
- Prefer 1–3 attachments per command.
- If Reference is {path}:{symbol}: attach that file; optionally ONE more upstream entrypoint/router/config if needed.
- If Reference is {endpoint}: attach route map + handler file (if different).
- If Reference is {event}: attach job registration + worker implementation (if different).
- If Reference is Unknown: attach only “index” files (README + primary manifest + 1 best-guess entrypoint if found).
- NEVER attach files you did not discover or cannot justify.

extra_files behavior:
- If extra_files was provided, append those as additional -f entries at the end of EVERY command (after the minimal evidence attachments).

Slug rules:
- <slug> lowercase a-z0-9 and hyphens only.
- Derive from category + a hint of reference (fallback: category only).

------------------------------------------------------------
F) SAVE / WRITE SEMANTICS (required; docs/)
- The oracle outputs are written by oracle itself via --write-output "<out_dir>/...".
  Because out_dir defaults under docs/, oracle outputs will land under docs/.
- This slash-command response (the pack Markdown) is printed to stdout.
  The caller MUST save it to the path printed on the first line:
  Output file: docs/strategist-questions-oracle-pack-YYYY-MM-DD.md
  (If you cannot determine date, use: docs/strategist-questions-oracle-pack.md)

------------------------------------------------------------
G) OUTPUT FORMAT (STRICT; MUST MATCH THIS SHAPE)
1) First line of the assistant response MUST be:
   Output file: docs/strategist-questions-oracle-pack-YYYY-MM-DD.md
   (Use today’s date if known; otherwise: docs/strategist-questions-oracle-pack.md)

2) After that line, output EXACTLY ONE Markdown document matching the template below.
3) Do NOT add any extra prose, headings, or fences beyond what the template requires.
4) The ```bash code fence MUST be triple-backtick fenced, and MUST close with triple-backticks.
5) Step headers inside the bash fence MUST be: # NN) ... (to satisfy oraclepack).

TEMPLATE (MUST MATCH):
<!-- begin template -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <docs/oracle/strategist-questions/$(date +%F)|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
out_dir="<out_dir>"
mkdir -p "$out_dir"

# 01) ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02) ...
# ...
# 20) ...
````

---

## coverage check (must be satisfied)

* contracts/interfaces: <OK|Missing (state missing artifact pattern)>

* invariants: <OK|Missing (...)>

* caching/state: <OK|Missing (...)>

* background jobs: <OK|Missing (...)>

* observability: <OK|Missing (...)>

* permissions: <OK|Missing (...)>

* migrations: <OK|Missing (...)>

* UX flows: <OK|Missing (...)>

* failure modes: <OK|Missing (...)>

* feature flags: <OK|Missing (...)>

<!-- end template -->

Now generate the final output that matches the template exactly.
"""

```

.config/completion/oraclepack.completion.sh
```
# bash completion V2 for oraclepack                           -*- shell-script -*-

__oraclepack_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE-} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Macs have bash3 for which the bash-completion package doesn't include
# _init_completion. This is a minimal version of that function.
__oraclepack_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

# This function calls the oraclepack program to obtain the completion
# results and the directive.  It fills the 'out' and 'directive' vars.
__oraclepack_get_completion_results() {
    local requestComp lastParam lastChar args

    # Prepare the command to request completions for the program.
    # Calling ${words[0]} instead of directly oraclepack allows handling aliases
    args=("${words[@]:1}")
    requestComp="${words[0]} __complete ${args[*]}"

    lastParam=${words[$((${#words[@]}-1))]}
    lastChar=${lastParam:$((${#lastParam}-1)):1}
    __oraclepack_debug "lastParam ${lastParam}, lastChar ${lastChar}"

    if [[ -z ${cur} && ${lastChar} != = ]]; then
        # If the last parameter is complete (there is a space following it)
        # We add an extra empty parameter so we can indicate this to the go method.
        __oraclepack_debug "Adding extra empty parameter"
        requestComp="${requestComp} ''"
    fi

    # When completing a flag with an = (e.g., oraclepack -n=<TAB>)
    # bash focuses on the part after the =, so we need to remove
    # the flag part from $cur
    if [[ ${cur} == -*=* ]]; then
        cur="${cur#*=}"
    fi

    __oraclepack_debug "Calling ${requestComp}"
    # Use eval to handle any environment variables and such
    out=$(eval "${requestComp}" 2>/dev/null)

    # Extract the directive integer at the very end of the output following a colon (:)
    directive=${out##*:}
    # Remove the directive
    out=${out%:*}
    if [[ ${directive} == "${out}" ]]; then
        # There is not directive specified
        directive=0
    fi
    __oraclepack_debug "The completion directive is: ${directive}"
    __oraclepack_debug "The completions are: ${out}"
}

__oraclepack_process_completion_results() {
    local shellCompDirectiveError=1
    local shellCompDirectiveNoSpace=2
    local shellCompDirectiveNoFileComp=4
    local shellCompDirectiveFilterFileExt=8
    local shellCompDirectiveFilterDirs=16
    local shellCompDirectiveKeepOrder=32

    if (((directive & shellCompDirectiveError) != 0)); then
        # Error code.  No completion.
        __oraclepack_debug "Received error from custom completion go code"
        return
    else
        if (((directive & shellCompDirectiveNoSpace) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                __oraclepack_debug "Activating no space"
                compopt -o nospace
            else
                __oraclepack_debug "No space directive not supported in this version of bash"
            fi
        fi
        if (((directive & shellCompDirectiveKeepOrder) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                # no sort isn't supported for bash less than < 4.4
                if [[ ${BASH_VERSINFO[0]} -lt 4 || ( ${BASH_VERSINFO[0]} -eq 4 && ${BASH_VERSINFO[1]} -lt 4 ) ]]; then
                    __oraclepack_debug "No sort directive not supported in this version of bash"
                else
                    __oraclepack_debug "Activating keep order"
                    compopt -o nosort
                fi
            else
                __oraclepack_debug "No sort directive not supported in this version of bash"
            fi
        fi
        if (((directive & shellCompDirectiveNoFileComp) != 0)); then
            if [[ $(type -t compopt) == builtin ]]; then
                __oraclepack_debug "Activating no file completion"
                compopt +o default
            else
                __oraclepack_debug "No file completion directive not supported in this version of bash"
            fi
        fi
    fi

    # Separate activeHelp from normal completions
    local completions=()
    local activeHelp=()
    __oraclepack_extract_activeHelp

    if (((directive & shellCompDirectiveFilterFileExt) != 0)); then
        # File extension filtering
        local fullFilter="" filter filteringCmd

        # Do not use quotes around the $completions variable or else newline
        # characters will be kept.
        for filter in ${completions[*]}; do
            fullFilter+="$filter|"
        done

        filteringCmd="_filedir $fullFilter"
        __oraclepack_debug "File filtering command: $filteringCmd"
        $filteringCmd
    elif (((directive & shellCompDirectiveFilterDirs) != 0)); then
        # File completion for directories only

        local subdir
        subdir=${completions[0]}
        if [[ -n $subdir ]]; then
            __oraclepack_debug "Listing directories in $subdir"
            pushd "$subdir" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1 || return
        else
            __oraclepack_debug "Listing directories in ."
            _filedir -d
        fi
    else
        __oraclepack_handle_completion_types
    fi

    __oraclepack_handle_special_char "$cur" :
    __oraclepack_handle_special_char "$cur" =

    # Print the activeHelp statements before we finish
    __oraclepack_handle_activeHelp
}

__oraclepack_handle_activeHelp() {
    # Print the activeHelp statements
    if ((${#activeHelp[*]} != 0)); then
        if [ -z $COMP_TYPE ]; then
            # Bash v3 does not set the COMP_TYPE variable.
            printf "\n";
            printf "%s\n" "${activeHelp[@]}"
            printf "\n"
            __oraclepack_reprint_commandLine
            return
        fi

        # Only print ActiveHelp on the second TAB press
        if [ $COMP_TYPE -eq 63 ]; then
            printf "\n"
            printf "%s\n" "${activeHelp[@]}"

            if ((${#COMPREPLY[*]} == 0)); then
                # When there are no completion choices from the program, file completion
                # may kick in if the program has not disabled it; in such a case, we want
                # to know if any files will match what the user typed, so that we know if
                # there will be completions presented, so that we know how to handle ActiveHelp.
                # To find out, we actually trigger the file completion ourselves;
                # the call to _filedir will fill COMPREPLY if files match.
                if (((directive & shellCompDirectiveNoFileComp) == 0)); then
                    __oraclepack_debug "Listing files"
                    _filedir
                fi
            fi

            if ((${#COMPREPLY[*]} != 0)); then
                # If there are completion choices to be shown, print a delimiter.
                # Re-printing the command-line will automatically be done
                # by the shell when it prints the completion choices.
                printf -- "--"
            else
                # When there are no completion choices at all, we need
                # to re-print the command-line since the shell will
                # not be doing it itself.
                __oraclepack_reprint_commandLine
            fi
        elif [ $COMP_TYPE -eq 37 ] || [ $COMP_TYPE -eq 42 ]; then
            # For completion type: menu-complete/menu-complete-backward and insert-completions
            # the completions are immediately inserted into the command-line, so we first
            # print the activeHelp message and reprint the command-line since the shell won't.
            printf "\n"
            printf "%s\n" "${activeHelp[@]}"

            __oraclepack_reprint_commandLine
        fi
    fi
}

__oraclepack_reprint_commandLine() {
    # The prompt format is only available from bash 4.4.
    # We test if it is available before using it.
    if (x=${PS1@P}) 2> /dev/null; then
        printf "%s" "${PS1@P}${COMP_LINE[@]}"
    else
        # Can't print the prompt.  Just print the
        # text the user had typed, it is workable enough.
        printf "%s" "${COMP_LINE[@]}"
    fi
}

# Separate activeHelp lines from real completions.
# Fills the $activeHelp and $completions arrays.
__oraclepack_extract_activeHelp() {
    local activeHelpMarker="_activeHelp_ "
    local endIndex=${#activeHelpMarker}

    while IFS='' read -r comp; do
        [[ -z $comp ]] && continue

        if [[ ${comp:0:endIndex} == $activeHelpMarker ]]; then
            comp=${comp:endIndex}
            __oraclepack_debug "ActiveHelp found: $comp"
            if [[ -n $comp ]]; then
                activeHelp+=("$comp")
            fi
        else
            # Not an activeHelp line but a normal completion
            completions+=("$comp")
        fi
    done <<<"${out}"
}

__oraclepack_handle_completion_types() {
    __oraclepack_debug "__oraclepack_handle_completion_types: COMP_TYPE is $COMP_TYPE"

    case $COMP_TYPE in
    37|42)
        # Type: menu-complete/menu-complete-backward and insert-completions
        # If the user requested inserting one completion at a time, or all
        # completions at once on the command-line we must remove the descriptions.
        # https://github.com/spf13/cobra/issues/1508

        # If there are no completions, we don't need to do anything
        (( ${#completions[@]} == 0 )) && return 0

        local tab=$'\t'

        # Strip any description and escape the completion to handled special characters
        IFS=$'\n' read -ra completions -d '' < <(printf "%q\n" "${completions[@]%%$tab*}")

        # Only consider the completions that match
        IFS=$'\n' read -ra COMPREPLY -d '' < <(IFS=$'\n'; compgen -W "${completions[*]}" -- "${cur}")

        # compgen looses the escaping so we need to escape all completions again since they will
        # all be inserted on the command-line.
        IFS=$'\n' read -ra COMPREPLY -d '' < <(printf "%q\n" "${COMPREPLY[@]}")
        ;;

    *)
        # Type: complete (normal completion)
        __oraclepack_handle_standard_completion_case
        ;;
    esac
}

__oraclepack_handle_standard_completion_case() {
    local tab=$'\t'

    # If there are no completions, we don't need to do anything
    (( ${#completions[@]} == 0 )) && return 0

    # Short circuit to optimize if we don't have descriptions
    if [[ "${completions[*]}" != *$tab* ]]; then
        # First, escape the completions to handle special characters
        IFS=$'\n' read -ra completions -d '' < <(printf "%q\n" "${completions[@]}")
        # Only consider the completions that match what the user typed
        IFS=$'\n' read -ra COMPREPLY -d '' < <(IFS=$'\n'; compgen -W "${completions[*]}" -- "${cur}")

        # compgen looses the escaping so, if there is only a single completion, we need to
        # escape it again because it will be inserted on the command-line.  If there are multiple
        # completions, we don't want to escape them because they will be printed in a list
        # and we don't want to show escape characters in that list.
        if (( ${#COMPREPLY[@]} == 1 )); then
            COMPREPLY[0]=$(printf "%q" "${COMPREPLY[0]}")
        fi
        return 0
    fi

    local longest=0
    local compline
    # Look for the longest completion so that we can format things nicely
    while IFS='' read -r compline; do
        [[ -z $compline ]] && continue

        # Before checking if the completion matches what the user typed,
        # we need to strip any description and escape the completion to handle special
        # characters because those escape characters are part of what the user typed.
        # Don't call "printf" in a sub-shell because it will be much slower
        # since we are in a loop.
        printf -v comp "%q" "${compline%%$tab*}" &>/dev/null || comp=$(printf "%q" "${compline%%$tab*}")

        # Only consider the completions that match
        [[ $comp == "$cur"* ]] || continue

        # The completions matches.  Add it to the list of full completions including
        # its description.  We don't escape the completion because it may get printed
        # in a list if there are more than one and we don't want show escape characters
        # in that list.
        COMPREPLY+=("$compline")

        # Strip any description before checking the length, and again, don't escape
        # the completion because this length is only used when printing the completions
        # in a list and we don't want show escape characters in that list.
        comp=${compline%%$tab*}
        if ((${#comp}>longest)); then
            longest=${#comp}
        fi
    done < <(printf "%s\n" "${completions[@]}")

    # If there is a single completion left, remove the description text and escape any special characters
    if ((${#COMPREPLY[*]} == 1)); then
        __oraclepack_debug "COMPREPLY[0]: ${COMPREPLY[0]}"
        COMPREPLY[0]=$(printf "%q" "${COMPREPLY[0]%%$tab*}")
        __oraclepack_debug "Removed description from single completion, which is now: ${COMPREPLY[0]}"
    else
        # Format the descriptions
        __oraclepack_format_comp_descriptions $longest
    fi
}

__oraclepack_handle_special_char()
{
    local comp="$1"
    local char=$2
    if [[ "$comp" == *${char}* && "$COMP_WORDBREAKS" == *${char}* ]]; then
        local word=${comp%"${comp##*${char}}"}
        local idx=${#COMPREPLY[*]}
        while ((--idx >= 0)); do
            COMPREPLY[idx]=${COMPREPLY[idx]#"$word"}
        done
    fi
}

__oraclepack_format_comp_descriptions()
{
    local tab=$'\t'
    local comp desc maxdesclength
    local longest=$1

    local i ci
    for ci in ${!COMPREPLY[*]}; do
        comp=${COMPREPLY[ci]}
        # Properly format the description string which follows a tab character if there is one
        if [[ "$comp" == *$tab* ]]; then
            __oraclepack_debug "Original comp: $comp"
            desc=${comp#*$tab}
            comp=${comp%%$tab*}

            # $COLUMNS stores the current shell width.
            # Remove an extra 4 because we add 2 spaces and 2 parentheses.
            maxdesclength=$(( COLUMNS - longest - 4 ))

            # Make sure we can fit a description of at least 8 characters
            # if we are to align the descriptions.
            if ((maxdesclength > 8)); then
                # Add the proper number of spaces to align the descriptions
                for ((i = ${#comp} ; i < longest ; i++)); do
                    comp+=" "
                done
            else
                # Don't pad the descriptions so we can fit more text after the completion
                maxdesclength=$(( COLUMNS - ${#comp} - 4 ))
            fi

            # If there is enough space for any description text,
            # truncate the descriptions that are too long for the shell width
            if ((maxdesclength > 0)); then
                if ((${#desc} > maxdesclength)); then
                    desc=${desc:0:$(( maxdesclength - 1 ))}
                    desc+="…"
                fi
                comp+="  ($desc)"
            fi
            COMPREPLY[ci]=$comp
            __oraclepack_debug "Final comp: $comp"
        fi
    done
}

__start_oraclepack()
{
    local cur prev words cword split

    COMPREPLY=()

    # Call _init_completion from the bash-completion package
    # to prepare the arguments properly
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -n =: || return
    else
        __oraclepack_init_completion -n =: || return
    fi

    __oraclepack_debug
    __oraclepack_debug "========= starting completion logic =========="
    __oraclepack_debug "cur is ${cur}, words[*] is ${words[*]}, #words[@] is ${#words[@]}, cword is $cword"

    # The user could have moved the cursor backwards on the command-line.
    # We need to trigger completion from the $cword location, so we need
    # to truncate the command-line ($words) up to the $cword location.
    words=("${words[@]:0:$cword+1}")
    __oraclepack_debug "Truncated words[*]: ${words[*]},"

    local out directive
    __oraclepack_get_completion_results
    __oraclepack_process_completion_results
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_oraclepack oraclepack
else
    complete -o default -o nospace -F __start_oraclepack oraclepack
fi

# ex: ts=4 sw=4 et filetype=sh
```

.config/scripts/build_install_oraclepack.md
```
```bash
# path: scripts/build_install_oraclepack.sh
#!/usr/bin/env bash
set -euo pipefail

# Builds oraclepack for:
# - WSL/Linux (oraclepack)
# - Windows amd64 (oraclepack.exe)
# Then installs:
# - WSL binary -> ~/.local/bin/
# - Windows exe -> /mnt/c/Users/<winuser>/.local/bin/
# And writes a Git Bash wrapper:
# - /mnt/c/Users/<winuser>/bin/oraclepack  (calls WSL binary via wsl.exe)

usage() {
  cat <<'USAGE'
Usage:
  scripts/build_install_oraclepack.sh [options]

Options:
  --repo-root <path>      Repo root (default: git toplevel, else current dir)
  --win-user <name>       Windows username (default: auto-detect, else "user")
  --wsl-distro <name>     WSL distro name for wsl.exe -d (default: $WSL_DISTRO_NAME or "Ubuntu-24.04")
  --wsl-user <name>       WSL username for wsl.exe -u (default: current user)
  --no-linux              Skip building/installing Linux binary
  --no-windows            Skip building/installing Windows .exe
  --no-wrapper            Skip writing the Git Bash wrapper script
  --cgo <0|1>             Set CGO_ENABLED (default: leave unchanged)

Examples:
  scripts/build_install_oraclepack.sh
  scripts/build_install_oraclepack.sh --win-user Alice --wsl-distro Ubuntu-24.04
USAGE
}

die() { echo "error: $*" >&2; exit 1; }

have() { command -v "$1" >/dev/null 2>&1; }

detect_repo_root() {
  if have git && git rev-parse --show-toplevel >/dev/null 2>&1; then
    git rev-parse --show-toplevel
  else
    pwd
  fi
}

detect_win_user() {
  # Best-effort: ask Windows for %USERNAME% via cmd.exe (works in most WSL setups).
  if have cmd.exe; then
    local u
    u="$(cmd.exe /c "echo %USERNAME%" 2>/dev/null | tr -d '\r' | tail -n 1 || true)"
    if [[ -n "${u:-}" && "${u:-}" != "%USERNAME%" ]]; then
      echo "$u"
      return
    fi
  fi
  echo "user"
}

REPO_ROOT=""
WIN_USER=""
WSL_DISTRO="${WSL_DISTRO_NAME:-Ubuntu-24.04}"
WSL_USER="$(whoami)"
DO_LINUX=1
DO_WINDOWS=1
DO_WRAPPER=1
CGO=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo-root)  REPO_ROOT="${2:-}"; shift 2;;
    --win-user)   WIN_USER="${2:-}"; shift 2;;
    --wsl-distro) WSL_DISTRO="${2:-}"; shift 2;;
    --wsl-user)   WSL_USER="${2:-}"; shift 2;;
    --no-linux)   DO_LINUX=0; shift;;
    --no-windows) DO_WINDOWS=0; shift;;
    --no-wrapper) DO_WRAPPER=0; shift;;
    --cgo)        CGO="${2:-}"; shift 2;;
    -h|--help)    usage; exit 0;;
    *)            die "unknown option: $1 (use --help)";;
  esac
done

[[ -n "$REPO_ROOT" ]] || REPO_ROOT="$(detect_repo_root)"
[[ -d "$REPO_ROOT" ]] || die "repo root not found: $REPO_ROOT"

[[ -n "$WIN_USER" ]] || WIN_USER="$(detect_win_user)"

have go || die "go not found in PATH"

# Paths
LINUX_OUT="$REPO_ROOT/oraclepack"
WIN_OUT="$REPO_ROOT/oraclepack.exe"

LINUX_INSTALL_DIR="$HOME/.local/bin"
LINUX_INSTALL_PATH="$LINUX_INSTALL_DIR/oraclepack"

WIN_HOME="/mnt/c/Users/$WIN_USER"
WIN_LOCAL_BIN_DIR="$WIN_HOME/.local/bin"
WIN_LOCAL_BIN_PATH="$WIN_LOCAL_BIN_DIR/oraclepack.exe"

WIN_GITBASH_BIN_DIR="$WIN_HOME/bin"
WIN_GITBASH_WRAPPER_PATH="$WIN_GITBASH_BIN_DIR/oraclepack"

# Optional CGO toggle
if [[ -n "$CGO" ]]; then
  [[ "$CGO" == "0" || "$CGO" == "1" ]] || die "--cgo must be 0 or 1"
  export CGO_ENABLED="$CGO"
fi

cd "$REPO_ROOT"

# Build binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Building Linux (WSL) binary: $LINUX_OUT"
  go build -o "$LINUX_OUT" ./cmd/oraclepack
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Building Windows amd64 exe: $WIN_OUT"
  GOOS=windows GOARCH=amd64 go build -o "$WIN_OUT" ./cmd/oraclepack
fi

# Install binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Installing Linux binary -> $LINUX_INSTALL_PATH"
  mkdir -p "$LINUX_INSTALL_DIR"
  cp -f "$LINUX_OUT" "$LINUX_INSTALL_PATH"
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Installing Windows exe -> $WIN_LOCAL_BIN_PATH"
  mkdir -p "$WIN_LOCAL_BIN_DIR"
  cp -f "$WIN_OUT" "$WIN_LOCAL_BIN_PATH"
fi

# Write Git Bash wrapper (stored on Windows filesystem)
if [[ "$DO_WRAPPER" -eq 1 ]]; then
  echo "==> Writing Git Bash wrapper -> $WIN_GITBASH_WRAPPER_PATH"
  mkdir -p "$WIN_GITBASH_BIN_DIR"

  cat > "$WIN_GITBASH_WRAPPER_PATH" <<EOF
#!/usr/bin/env bash
set -euo pipefail

# Git for Windows (Git Bash) path-conversion off for this exec call.
# Required so /home/... is not rewritten into C:/Program Files/Git/...
MSYS_NO_PATHCONV=1 exec wsl.exe -d ${WSL_DISTRO@Q} -u ${WSL_USER@Q} -- ${LINUX_INSTALL_PATH@Q} "\$@"
EOF

  # Ensure LF line endings (in case editor/tool wrote CRLF) and best-effort executable bit.
  sed -i 's/\r$//' "$WIN_GITBASH_WRAPPER_PATH" || true
  chmod +x "$WIN_GITBASH_WRAPPER_PATH" 2>/dev/null || true

  echo "==> Note: In Git Bash, ensure ~/bin is in PATH (so 'oraclepack' resolves to this wrapper)."
fi

echo "==> Done."
echo "    WSL binary:     $LINUX_INSTALL_PATH"
echo "    Windows exe:    $WIN_LOCAL_BIN_PATH"
echo "    Git Bash wrap:  $WIN_GITBASH_WRAPPER_PATH"
```

Run (from repo root in WSL):

```bash
chmod +x scripts/build_install_oraclepack.sh
scripts/build_install_oraclepack.sh
```

Optional example:

```bash
scripts/build_install_oraclepack.sh --win-user user --wsl-distro Ubuntu-24.04 --wsl-user user
```
```

.config/scripts/build_install_oraclepack.sh
```
# path: scripts/build_install_oraclepack.sh
#!/usr/bin/env bash
set -euo pipefail

# Builds oraclepack for:
# - WSL/Linux (oraclepack)
# - Windows amd64 (oraclepack.exe)
# Then installs:
# - WSL binary -> ~/.local/bin/
# - Windows exe -> /mnt/c/Users/<winuser>/.local/bin/
# And writes a Git Bash wrapper:
# - /mnt/c/Users/<winuser>/bin/oraclepack  (calls WSL binary via wsl.exe)

usage() {
  cat <<'USAGE'
Usage:
  scripts/build_install_oraclepack.sh [options]

Options:
  --repo-root <path>      Repo root (default: git toplevel, else current dir)
  --win-user <name>       Windows username (default: auto-detect, else "user")
  --wsl-distro <name>     WSL distro name for wsl.exe -d (default: $WSL_DISTRO_NAME or "Ubuntu-24.04")
  --wsl-user <name>       WSL username for wsl.exe -u (default: current user)
  --no-linux              Skip building/installing Linux binary
  --no-windows            Skip building/installing Windows .exe
  --no-wrapper            Skip writing the Git Bash wrapper script
  --cgo <0|1>             Set CGO_ENABLED (default: leave unchanged)

Examples:
  scripts/build_install_oraclepack.sh
  scripts/build_install_oraclepack.sh --win-user Alice --wsl-distro Ubuntu-24.04
USAGE
}

die() { echo "error: $*" >&2; exit 1; }

have() { command -v "$1" >/dev/null 2>&1; }

detect_repo_root() {
  if have git && git rev-parse --show-toplevel >/dev/null 2>&1; then
    git rev-parse --show-toplevel
  else
    pwd
  fi
}

detect_win_user() {
  # Best-effort: ask Windows for %USERNAME% via cmd.exe (works in most WSL setups).
  if have cmd.exe; then
    local u
    u="$(cmd.exe /c "echo %USERNAME%" 2>/dev/null | tr -d '\r' | tail -n 1 || true)"
    if [[ -n "${u:-}" && "${u:-}" != "%USERNAME%" ]]; then
      echo "$u"
      return
    fi
  fi
  echo "user"
}

REPO_ROOT=""
WIN_USER=""
WSL_DISTRO="${WSL_DISTRO_NAME:-Ubuntu-24.04}"
WSL_USER="$(whoami)"
DO_LINUX=1
DO_WINDOWS=1
DO_WRAPPER=1
CGO=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo-root)  REPO_ROOT="${2:-}"; shift 2;;
    --win-user)   WIN_USER="${2:-}"; shift 2;;
    --wsl-distro) WSL_DISTRO="${2:-}"; shift 2;;
    --wsl-user)   WSL_USER="${2:-}"; shift 2;;
    --no-linux)   DO_LINUX=0; shift;;
    --no-windows) DO_WINDOWS=0; shift;;
    --no-wrapper) DO_WRAPPER=0; shift;;
    --cgo)        CGO="${2:-}"; shift 2;;
    -h|--help)    usage; exit 0;;
    *)            die "unknown option: $1 (use --help)";;
  esac
done

[[ -n "$REPO_ROOT" ]] || REPO_ROOT="$(detect_repo_root)"
[[ -d "$REPO_ROOT" ]] || die "repo root not found: $REPO_ROOT"

[[ -n "$WIN_USER" ]] || WIN_USER="$(detect_win_user)"

have go || die "go not found in PATH"

# Paths
LINUX_OUT="$REPO_ROOT/oraclepack"
WIN_OUT="$REPO_ROOT/oraclepack.exe"

LINUX_INSTALL_DIR="$HOME/.local/bin"
LINUX_INSTALL_PATH="$LINUX_INSTALL_DIR/oraclepack"

WIN_HOME="/mnt/c/Users/$WIN_USER"
WIN_LOCAL_BIN_DIR="$WIN_HOME/.local/bin"
WIN_LOCAL_BIN_PATH="$WIN_LOCAL_BIN_DIR/oraclepack.exe"

WIN_GITBASH_BIN_DIR="$WIN_HOME/bin"
WIN_GITBASH_WRAPPER_PATH="$WIN_GITBASH_BIN_DIR/oraclepack"

# Optional CGO toggle
if [[ -n "$CGO" ]]; then
  [[ "$CGO" == "0" || "$CGO" == "1" ]] || die "--cgo must be 0 or 1"
  export CGO_ENABLED="$CGO"
fi

cd "$REPO_ROOT"

# Build binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Building Linux (WSL) binary: $LINUX_OUT"
  go build -o "$LINUX_OUT" ./cmd/oraclepack
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Building Windows amd64 exe: $WIN_OUT"
  GOOS=windows GOARCH=amd64 go build -o "$WIN_OUT" ./cmd/oraclepack
fi

# Install binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Installing Linux binary -> $LINUX_INSTALL_PATH"
  mkdir -p "$LINUX_INSTALL_DIR"
  cp -f "$LINUX_OUT" "$LINUX_INSTALL_PATH"
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Installing Windows exe -> $WIN_LOCAL_BIN_PATH"
  mkdir -p "$WIN_LOCAL_BIN_DIR"
  cp -f "$WIN_OUT" "$WIN_LOCAL_BIN_PATH"
fi

# Write Git Bash wrapper (stored on Windows filesystem)
if [[ "$DO_WRAPPER" -eq 1 ]]; then
  echo "==> Writing Git Bash wrapper -> $WIN_GITBASH_WRAPPER_PATH"
  mkdir -p "$WIN_GITBASH_BIN_DIR"

  cat > "$WIN_GITBASH_WRAPPER_PATH" <<EOF
#!/usr/bin/env bash
set -euo pipefail

# Git for Windows (Git Bash) path-conversion off for this exec call.
# Required so /home/... is not rewritten into C:/Program Files/Git/...
MSYS_NO_PATHCONV=1 exec wsl.exe -d ${WSL_DISTRO@Q} -u ${WSL_USER@Q} -- ${LINUX_INSTALL_PATH@Q} "\$@"
EOF

  # Ensure LF line endings (in case editor/tool wrote CRLF) and best-effort executable bit.
  sed -i 's/\r$//' "$WIN_GITBASH_WRAPPER_PATH" || true
  chmod +x "$WIN_GITBASH_WRAPPER_PATH" 2>/dev/null || true

  echo "==> Note: In Git Bash, ensure ~/bin is in PATH (so 'oraclepack' resolves to this wrapper)."
fi

echo "==> Done."
echo "    WSL binary:     $LINUX_INSTALL_PATH"
echo "    Windows exe:    $WIN_LOCAL_BIN_PATH"
echo "    Git Bash wrap:  $WIN_GITBASH_WRAPPER_PATH"
```

.config/scripts/install-global.ps1
```
# path: scripts/install-global.ps1
[CmdletBinding()]
param(
  [string]$InstallDir = "C:\Tools",
  [switch]$AddToPath,
  [switch]$SkipTests
)

$ErrorActionPreference = "Stop"

function Get-RepoRoot {
  try {
    return (git rev-parse --show-toplevel).Trim()
  } catch {
    throw "Not inside a git repo (expected oraclepack repo)."
  }
}

$RepoRoot = Get-RepoRoot
Set-Location $RepoRoot

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
  throw "Go not found in PATH."
}

Write-Host "==> Repo: $RepoRoot"
Write-Host "==> InstallDir: $InstallDir"

if (-not $SkipTests) {
  Write-Host "==> go test ./..."
  go test ./...
}

# Match README build target: ./cmd/oraclepack
Write-Host "==> go build -o oraclepack.exe ./cmd/oraclepack"
go build -o oraclepack.exe ./cmd/oraclepack

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
Copy-Item -Force -Path (Join-Path $RepoRoot "oraclepack.exe") -Destination (Join-Path $InstallDir "oraclepack.exe")

if ($AddToPath) {
  $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
  if ($userPath -notmatch [Regex]::Escape($InstallDir)) {
    $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) { $InstallDir } else { "$userPath;$InstallDir" }
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "==> Added to User PATH: $InstallDir (restart terminal to pick up PATH changes)"
  } else {
    Write-Host "==> InstallDir already in User PATH"
  }
}

Write-Host "==> Installed: $InstallDir\oraclepack.exe"
& (Join-Path $InstallDir "oraclepack.exe") --version 2>$null
```

.config/scripts/install-global.sh
```
# path: scripts/install-global.sh
#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'USAGE'
install-global.sh — rebuild oraclepack and install it to a PATH directory.

Usage:
  ./scripts/install-global.sh [--prefix DIR] [--sudo] [--skip-tests]

Options:
  --prefix DIR     Install directory (default: ~/.local/bin; macOS prefers /usr/local/bin if writable)
  --sudo           Use sudo when copying into --prefix (for system dirs like /usr/local/bin)
  --skip-tests     Skip `go test ./...`

Examples:
  ./scripts/install-global.sh
  ./scripts/install-global.sh --prefix "$HOME/.local/bin"
  ./scripts/install-global.sh --prefix /usr/local/bin --sudo
USAGE
}

PREFIX=""
USE_SUDO=0
SKIP_TESTS=0

while [[ $# -gt 0 ]]; do
  case "$1" in
    --prefix) PREFIX="${2:-}"; shift 2 ;;
    --sudo) USE_SUDO=1; shift ;;
    --skip-tests) SKIP_TESTS=1; shift ;;
    -h|--help) usage; exit 0 ;;
    *) echo "Unknown arg: $1" >&2; usage; exit 2 ;;
  esac
done

# Repo root (so script works from anywhere)
REPO_ROOT="$(git rev-parse --show-toplevel 2>/dev/null || true)"
if [[ -z "${REPO_ROOT}" ]]; then
  echo "Error: not inside a git repo (expected oraclepack repo)." >&2
  exit 1
fi
cd "$REPO_ROOT"

# Build output
BUILD_DIR="$REPO_ROOT/.build"
mkdir -p "$BUILD_DIR"
OUT_BIN="$BUILD_DIR/oraclepack"

# Default install prefix
if [[ -z "${PREFIX}" ]]; then
  case "$(uname -s)" in
    Darwin)
      if [[ -w "/usr/local/bin" ]]; then
        PREFIX="/usr/local/bin"
      else
        PREFIX="$HOME/.local/bin"
      fi
      ;;
    *)
      PREFIX="$HOME/.local/bin"
      ;;
  esac
fi

echo "==> Repo: $REPO_ROOT"
echo "==> Building: $OUT_BIN"
echo "==> Installing to: $PREFIX"

command -v go >/dev/null 2>&1 || { echo "Error: go not found in PATH." >&2; exit 1; }

if [[ "$SKIP_TESTS" -eq 0 ]]; then
  echo "==> go test ./..."
  go test ./...
fi

# Match README build target: ./cmd/oraclepack :contentReference[oaicite:0]{index=0}
go build -o "$OUT_BIN" ./cmd/oraclepack

mkdir -p "$PREFIX"

# Install (overwrite existing)
if [[ "$USE_SUDO" -eq 1 ]]; then
  sudo install -m 0755 "$OUT_BIN" "$PREFIX/oraclepack"
else
  install -m 0755 "$OUT_BIN" "$PREFIX/oraclepack"
fi

echo "==> Installed: $PREFIX/oraclepack"
echo "==> Version (if supported):"
"$PREFIX/oraclepack" --version 2>/dev/null || true

# Best-effort: refresh shell hash table (won't error if unsupported)
hash -r 2>/dev/null || true
```

.config/scripts/tag-release.sh
```
# path: scripts/tag-release.sh
#!/usr/bin/env bash
set -euo pipefail

read -r -p "Input tag version (e.g., 0.2.0): " version
version="${version//[[:space:]]/}"
[[ -n "${version}" ]] || { echo "Error: version cannot be empty." >&2; exit 1; }

tag="v${version}"

git rev-parse -q --verify "refs/tags/${tag}" >/dev/null && {
  echo "Error: tag ${tag} already exists locally." >&2
  exit 1
}

git ls-remote --tags origin "refs/tags/${tag}" | grep -q . && {
  echo "Error: tag ${tag} already exists on origin." >&2
  exit 1
}

git tag "${tag}"
git push origin "${tag}"
```

.github/workflows/ci.yml
```
# path: .github/workflows/ci.yml
name: ci

on:
  pull_request:
  push:
    branches:
      - main

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v5

      - name: Set up Go
        uses: actions/setup-go@v6
        with:
          go-version-file: go.mod
          cache: true

      # oraclepack tests expect `oracle` to be available on PATH
      - name: Set up Node (for oracle CLI)
        uses: actions/setup-node@v4
        with:
          node-version: "22"

      - name: Install oracle CLI
        run: npm install -g @steipete/oracle

      - name: Smoke test oracle CLI
        run: oracle --help

      - name: Test
        run: go test ./...

      - name: Vet
        run: go vet ./...
```

.github/workflows/release.yml
```
# path: .github/workflows/release.yml
name: release

on:
  push:
    tags:
      - "v*"

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v5
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v6
        with:
          # Reads the Go version from go.mod ("go 1.xx")
          go-version-file: go.mod
          cache: true

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v6
        with:
          distribution: goreleaser
          version: "~> v2"
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

.tickets/mcp/Expose Oraclepack as MCP.md
```
Parent Ticket:

* Title: Expose oraclepack as MCP tools (with Taskify Stage-2/Stage-3 helpers)
* Summary: Provide an MCP server that exposes `oraclepack` CLI capabilities (validate/list/run) plus helper tools for Stage-2 detection/validation and Stage-3 action-pack validation/execution/artifact summarization, with secure-by-default controls (allowlisted filesystem roots, execution gating, timeouts, truncation) and support for stdio + streamable-http transports.
* Source:

  * Link/ID (if present) or “Not provided”: /mnt/data/MCP tools for Oraclepack.md
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Expose `oraclepack` … as **MCP tools**, so an agent can … run packs … validate Stage-2 … validate Stage-3 … summarize artifacts.”
* Global Constraints:

  * Support MCP transports: “stdio” and “streamable-http”.
  * Security defaults: “deny-by-default execution”, “allowlisted roots”, “stdout/stderr truncation and timeouts”.
* Global Environment:

  * Unknown
* Global Evidence:

  * MCP tool list and env var config list.
  * Reference implementation structure (README/requirements and Python modules).

Split Plan:

* Coverage Map:

  * “Expose `oraclepack` (validate/list/run) … as **MCP tools**” → T1
  * “run packs non-interactively (`--no-tui --yes --run-all`)” → T5
  * “validate Stage-2 outputs (01-*.md..20-*.md)” → T4
  * “validate Stage-3 Action Packs (single ```bash fence, step headers…)” → T7
  * “summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)” → T7
  * “Tools: oraclepack_validate_pack / oraclepack_list_steps / oraclepack_run_pack …” → T5
  * “Tools: oraclepack_read_file …” → T5
  * “Tools: … taskify_detect_stage2 / taskify_validate_stage2 …” → T5
  * “Tools: … taskify_validate_action_pack / taskify_artifacts_summary …” → T5
  * “Tools: … taskify_run_action_pack …” → T5
  * “Transports: stdio … streamable-http …” → T6
  * “Tool annotations: readOnlyHint / destructiveHint / openWorldHint …” → T6
  * “Security defaults: ORACLEPACK_ENABLE_EXEC=1 gating …” → T2
  * “Security defaults: allowlisted filesystem roots …” → T2
  * “Security defaults: truncation and timeouts …” → T3
  * “Env vars: ORACLEPACK_ALLOWED_ROOTS / BIN / WORKDIR / ENABLE_EXEC / CHARACTER_LIMIT / MAX_READ_BYTES” → T2
  * “Typical Stage-3 usage: detect/validate → validate action pack → run → summarize” → Info-only
  * “Reference implementation tree (README, requirements.txt, modules list)” → Info-only
  * Links to MCP specs / python-sdk repo mentioned → Info-only
* Dependencies:

  * T5 depends on T2 because server tools must enforce allowed roots and execution gating.
  * T5 depends on T3 because `oraclepack_*run*` tools need subprocess execution with timeouts/truncation.
  * T5 depends on T4 because `oraclepack_taskify_*stage2*` tools call Stage-2 detection/validation.
  * T5 depends on T7 because `oraclepack_taskify_*action_pack*` tools call action-pack validation/summarization helpers.
  * T6 depends on T5 because annotations/transport-hardening apply to the MCP server surface.
* Split Tickets:

```ticket T1
T# Title: Scaffold oraclepack MCP server project (README + packaging entrypoints)
Type: chore
Target Area: oraclepack-mcp-server repo scaffolding (README.md, requirements.txt, __init__.py, __main__.py)
Summary:
- Create the MCP server project structure that exposes `oraclepack` + Taskify helpers as MCP tools, including installation and run instructions and the tool list.
- Ensure the package has an executable entrypoint to start the MCP server with selectable transport.
In Scope:
- Create/maintain project tree with:
  - README describing purpose, install, configuration env vars, run modes, tools list, and typical Stage-3 usage.
  - requirements.txt listing dependencies.
  - Python package layout with `oraclepack_mcp_server/__init__.py` and `oraclepack_mcp_server/__main__.py`.
- CLI args in `__main__.py` to accept `--transport` with choices `stdio` and `streamable-http`.
Out of Scope:
- Implementing tool logic (handled in other tickets).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- A runnable package that starts an MCP server with a chosen transport.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support `--transport` choices: `stdio`, `streamable-http`.
Evidence:
- Project tree and entrypoint snippet showing `choices=["stdio", "streamable-http"]`. :contentReference[oaicite:26]{index=26}
Open Items / Unknowns:
- Exact repository root / packaging approach (pip package vs repo-local module): Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Repository includes README.md, requirements.txt, and `oraclepack_mcp_server` package directory.
- [ ] Running `python -m oraclepack_mcp_server --transport stdio` is supported (starts server process).
- [ ] Running `python -m oraclepack_mcp_server --transport streamable-http` is supported (starts server process).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers) … tree … README.md … requirements.txt … __main__.py” :contentReference[oaicite:27]{index=27}
- “choices=[‘stdio’, ‘streamable-http’] … mcp.run(transport=args.transport)” :contentReference[oaicite:28]{index=28}
```

```ticket T2
T# Title: Implement config + filesystem security controls (allowed roots, exec gating, max read bytes)
Type: chore
Target Area: oraclepack_mcp_server/config.py and oraclepack_mcp_server/security.py
Summary:
- Implement secure-by-default configuration for the MCP server, driven by environment variables, including allowlisted filesystem roots and explicit execution enablement.
- Provide safe path resolution under allowed roots and bounded file reads for tool operations.
In Scope:
- Env-driven config including:
  - `ORACLEPACK_ALLOWED_ROOTS` (colon-separated roots)
  - `ORACLEPACK_BIN`
  - `ORACLEPACK_WORKDIR`
  - `ORACLEPACK_ENABLE_EXEC`
  - `ORACLEPACK_CHARACTER_LIMIT`
  - `ORACLEPACK_MAX_READ_BYTES`
- Path allowlisting:
  - Resolve requested file paths and ensure they live under at least one allowed root.
  - Raise an explicit “path not allowed” error on violation.
- Safe file reads:
  - Read text/bytes with max size enforcement and “truncated” indicator.
Out of Scope:
- Subprocess execution and output truncation (handled in T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Server loads config from env and enforces filesystem access boundaries for read tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Execution must be deny-by-default unless `ORACLEPACK_ENABLE_EXEC=1`.
- Filesystem access must be restricted to allowlisted roots.
Evidence:
- Env var list and semantics. :contentReference[oaicite:29]{index=29}
- Security guidance: “deny-by-default execution … allowlisted roots”. :contentReference[oaicite:30]{index=30}
Open Items / Unknowns:
- Whether Windows path separator support is required for `ORACLEPACK_ALLOWED_ROOTS`: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Config loader reads the listed env vars and applies defaults as documented.
- [ ] Path resolution rejects paths outside allowed roots.
- [ ] File reads enforce max bytes and indicate truncation.
- [ ] Exec gating flag is available for run tools to check.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Environment variables: ORACLEPACK_ALLOWED_ROOTS … ORACLEPACK_ENABLE_EXEC … ORACLEPACK_MAX_READ_BYTES …” :contentReference[oaicite:31]{index=31}
- “Hard deny-by-default execution … Restrict filesystem access to allowlisted roots …” :contentReference[oaicite:32]{index=32}
```

```ticket T3
T# Title: Implement oraclepack subprocess runner with timeouts and stdout/stderr truncation
Type: chore
Target Area: oraclepack_mcp_server/oraclepack_cli.py (subprocess execution)
Summary:
- Provide a subprocess wrapper to invoke the `oraclepack` CLI with a hard timeout and bounded stdout/stderr capture to prevent wedging the MCP server.
- Return structured results including exit code, duration, and truncation indicators.
In Scope:
- Async subprocess execution wrapper (create subprocess, capture stdout/stderr).
- Timeout behavior:
  - Kill process on timeout and return an explicit “Timed out after {timeout}s” error result.
- Character-limit truncation for stdout/stderr based on configured limit.
Out of Scope:
- MCP tool registration (handled in T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Running oraclepack commands yields deterministic, bounded outputs suitable for returning via MCP tools.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must enforce “stdout/stderr truncation and timeouts”.
Evidence:
- Guidance: “Enforce stdout/stderr truncation and timeouts …”. :contentReference[oaicite:33]{index=33}
- Runner timeout error example snippet. :contentReference[oaicite:34]{index=34}
Open Items / Unknowns:
- Default timeout values per tool beyond examples (3600/7200) are not provided outside snippets.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Runner returns: ok, exit_code, duration_s, stdout, stderr, stdout_truncated, stderr_truncated.
- [ ] Timeout produces exit_code=124 (or equivalent) and includes a timeout message.
- [ ] Outputs are truncated to configured character limit and flags are set accordingly.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.” :contentReference[oaicite:35]{index=35}
- “Timed out after {timeout_s}s: …” return structure snippet. :contentReference[oaicite:36]{index=36}
```

```ticket T4
T# Title: Implement Taskify Stage-2 detection and validation (01-*.md..20-*.md + single-pack form)
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-2 detection + validation)
Summary:
- Implement deterministic detection of Stage-2 outputs for agents, supporting both directory-form outputs (01-*.md..20-*.md) and a single-pack input file form.
- Provide validation that ensures exactly one match per prefix 01..20 and returns missing/ambiguous details.
In Scope:
- `validate_stage2_dir(out_dir)`:
  - For each prefix 01..20, glob `{pfx}-*.md`
  - Return missing patterns and ambiguous prefix matches.
- `detect_stage2(stage2_path, repo_root)`:
  - Support explicit dir path.
  - Support explicit file path with out_dir rules:
    - If under `docs/oracle-questions-YYYY-MM-DD/…`, use sibling `oracle-out` under that docs subtree; else default `repo_root/oracle-out`.
  - Support “auto” discovery (best-effort, deterministic ordering), including checking `repo_root/oracle-out`.
Out of Scope:
- Stage-3 action pack validation (handled in T7).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Agents can resolve and validate Stage-2 outputs deterministically for downstream Stage-3 workflows.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form)”.
- “Validate Stage-2 outputs (exactly one match per prefix 01..20)”.
Evidence:
- Requirements text for Stage-2 detection/validation. :contentReference[oaicite:37]{index=37}
- Validation logic snippet for 01..20 and ambiguous/missing. :contentReference[oaicite:38]{index=38}
- Out-dir rule snippet referencing `docs/oracle-questions-YYYY-MM-DD/…` → `oracle-out`. :contentReference[oaicite:39]{index=39}
Open Items / Unknowns:
- Full “auto discovery” search order beyond checking `repo_root/oracle-out`: Not provided in visible excerpts.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation returns ok=false with `missing` when any prefix has no matches.
- [ ] Validation returns ok=false with `ambiguous` when any prefix has >1 match.
- [ ] Validation returns ok=true only when exactly one match exists for every prefix 01..20.
- [ ] Detection supports explicit dir and explicit file resolution and produces an `out_dir`.
- [ ] Detection supports “auto” and returns deterministic results for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Detect Stage-2 outputs (dir-form 01-*.md..20-*.md OR single-pack form) … Validate … exactly one match per prefix 01..20.” :contentReference[oaicite:40]{index=40}
- “out_dir rules … if under docs/oracle-questions-YYYY-MM-DD/ … then …/oracle-out else oracle-out” :contentReference[oaicite:41]{index=41}
```

```ticket T5
T# Title: Implement MCP tools for oraclepack and taskify helper operations
Type: enhancement
Target Area: oraclepack_mcp_server/server.py (MCP tool registration + schemas + formatting)
Summary:
- Implement the MCP server surface that maps `oraclepack` CLI operations and Taskify helper functions into callable MCP tools.
- Provide consistent response formatting (markdown/json) and ensure run tools respect execution gating.
In Scope:
- Tool schemas/inputs covering:
  - `oraclepack_validate_pack`
  - `oraclepack_list_steps`
  - `oraclepack_run_pack` (gated)
  - `oraclepack_read_file`
  - `oraclepack_taskify_detect_stage2`
  - `oraclepack_taskify_validate_stage2`
  - `oraclepack_taskify_validate_action_pack`
  - `oraclepack_taskify_artifacts_summary`
  - `oraclepack_taskify_run_action_pack` (gated)
- Response formatting:
  - Support JSON and Markdown result formats (including stdout/stderr blocks and truncation notes).
- CLI arg mapping for oraclepack operations (including references to flags like `--no-tui`, `--out-dir`, `--oracle-bin` as inputs or internal argv composition).
- Ensure `ORACLEPACK_ENABLE_EXEC=1` gating is enforced for run tools.
Out of Scope:
- Implementing the underlying Stage-2 detection/validation logic (T4) and Stage-3 validation/summary logic (T7), except wiring them in.
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients can invoke oraclepack and taskify workflows end-to-end via tools and receive bounded, formatted results.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must expose the tool list as stated.
- Run tools must be disabled by default unless enabled via env flag.
Evidence:
- Tool list and gating note. :contentReference[oaicite:42]{index=42}
- Server schema snippets (e.g., ReadFileInput, Taskify*Input, timeout defaults). :contentReference[oaicite:43]{index=43}
Open Items / Unknowns:
- Exact argument surface for `oraclepack_validate_pack` / `oraclepack_list_steps` (flags and required params) is not fully specified in excerpts.
Risks / Dependencies:
- Depends on config/security/runner/taskify modules existing and being wired correctly.
Acceptance Criteria:
- [ ] All tools listed in the parent ticket are registered and callable.
- [ ] `oraclepack_read_file` enforces allowed roots and max read bytes.
- [ ] `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack` refuse execution unless `ORACLEPACK_ENABLE_EXEC=1`.
- [ ] Response formatter supports markdown and json outputs and includes truncation indicators.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Tools — oraclepack_validate_pack … oraclepack_taskify_run_action_pack (requires ORACLEPACK_ENABLE_EXEC=1)” :contentReference[oaicite:44]{index=44}
- “Map … oraclepack CLI capabilities (validate/list/run + flags like --no-tui, --out-dir, --oracle-bin) into MCP tools” :contentReference[oaicite:45]{index=45}
```

```ticket T6
T# Title: Add MCP transport hardening guidance + tool UX annotations (readOnly/destructive/openWorld)
Type: enhancement
Target Area: MCP server transport configuration and tool metadata (server.py / deployment guidance)
Summary:
- Ensure the MCP server supports stdio and streamable-http in a way suitable for “real-time” usage, including the recommended security posture for HTTP transport.
- Add MCP tool annotations so clients can present appropriate approval UX for read-only vs execution tools.
In Scope:
- Transport support considerations:
  - stdio: ensure logs go to stderr (per guidance in ticket text).
  - streamable-http: implement recommended protections (“Origin validation”, “bind to localhost + auth”).
- Tool annotations:
  - Mark validate/list/read tools as `readOnlyHint: true`.
  - Mark run tools as `destructiveHint: true` and `openWorldHint: true`.
Out of Scope:
- Implementing tool business logic (T5).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- MCP clients see proper tool risk hints and HTTP transport is protected as recommended for local/real-time usage.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Must support “stdio” and “streamable-http”.
- Must provide tool annotations as described.
Evidence:
- Transport choices and HTTP security recommendations. :contentReference[oaicite:46]{index=46}
- Tool annotation guidance (readOnlyHint/destructiveHint/openWorldHint). :contentReference[oaicite:47]{index=47}
Open Items / Unknowns:
- Exact auth mechanism for streamable-http (token, mTLS, etc.): Not provided.
Risks / Dependencies:
- Depends on MCP SDK capabilities available in the chosen implementation.
Acceptance Criteria:
- [ ] Running with `--transport stdio` is supported and does not interleave logs on stdout.
- [ ] Running with `--transport streamable-http` includes Origin validation and uses localhost binding + authentication (mechanism documented/implemented).
- [ ] Validate/list/read tools are annotated as read-only; run tools are annotated as destructive/open-world.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Streamable HTTP … implement Origin validation and bind to localhost + auth to avoid DNS rebinding …” :contentReference[oaicite:48]{index=48}
- “mark validate/list/read tools as readOnlyHint … mark run tools as destructiveHint, openWorldHint …” :contentReference[oaicite:49]{index=49}
```

````ticket T7
T# Title: Implement Stage-3 Action Pack validation + artifact summarization helpers
Type: chore
Target Area: oraclepack_mcp_server/taskify.py (Stage-3 helpers) and wiring into server tools
Summary:
- Implement helper functions to validate Stage-3 “Action Pack” markdown constraints before execution and to summarize key Stage-3 artifacts produced by running action packs.
- These helpers support deterministic agent workflows around Taskify outputs.
In Scope:
- Action Pack validation logic:
  - Enforce “single ```bash fence” and “step headers” constraints as stated in the parent ticket.
- Artifacts summary logic:
  - Summarize outputs such as `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, and “pipelines doc” when present.
- Integration points:
  - Provide outputs suitable for `oraclepack_taskify_validate_action_pack` and `oraclepack_taskify_artifacts_summary` tools (registered in T5).
Out of Scope:
- Stage-2 detection/validation (T4).
- Actual execution of action packs (tool wiring and subprocess invocation handled in T5/T3).
Current Behavior (Actual):
- Not provided.
Expected Behavior:
- Action packs can be validated for structural correctness prior to execution, and resulting artifacts can be summarized for quick agent consumption.
Reproduction Steps:
- Not provided.
Requirements / Constraints:
- Validate Stage-3 Action Pack “single ```bash fence, step headers `# NN)`”.
- Summarize Stage-3 artifacts: `_actions.json`, PRD path, `tm-complexity.json`, pipelines doc, etc.
Evidence:
- Stage-3 validation requirement text. :contentReference[oaicite:50]{index=50}
- Artifact summary examples list. :contentReference[oaicite:51]{index=51}
Open Items / Unknowns:
- Exact, formal grammar for “step headers” beyond the example text: Not provided.
- Exact artifact filenames/paths beyond examples: Not provided.
Risks / Dependencies:
- Not provided.
Acceptance Criteria:
- [ ] Validation fails with a clear error when the action pack violates the “single bash fence” constraint.
- [ ] Validation fails with a clear error when step headers do not meet the stated expectations.
- [ ] Artifact summarizer reports presence/absence of the example artifacts and returns a readable summary.
- [ ] Tool outputs are deterministic for the same filesystem state.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Validate Stage-3 Action Pack … (single ```bash fence, step headers `# NN)` …)” :contentReference[oaicite:52]{index=52}
- “Summarize Stage-3 artifacts (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc)” :contentReference[oaicite:53]{index=53}
````
```

.tickets/mcp/MCP Server for Oraclepack.md
```
Title:

* Add MCP server that exposes `oraclepack` + Taskify Stage-2/Stage-3 helpers as tools for agents

Summary:

* Agents need real-time access to `oraclepack` capabilities via MCP so they can validate, inspect, and run packs, then consume Taskify artifacts produced by the `oracle_pack_and_taskify-skills.md` workflow.
* Implement a secure-by-default Python MCP server that wraps the `oraclepack` CLI and adds deterministic helpers for Stage-2 detection/validation and Stage-3 Action Pack validation/execution + artifact summarization.

Background / Context:

* Request: “give access to agents/assistants the following oraclepack tool as a MCP tool… so they can perform actions using the artifacts generated from `oracle_pack_and_taskify-skills.md` … in real time.”
* Proposed reference implementation (from the conversation) is a Python project named `oraclepack-mcp-server` using FastMCP (MCP Python SDK), supporting `stdio` and `streamable-http` transports.
* Stage-2 outputs are expected to be 20 markdown files matching `01-*.md` … `20-*.md`; Stage-2 “single-pack form” needs out-dir resolution rules when the pack lives under `docs/oracle-questions-YYYY-MM-DD/...`.

Current Behavior (Actual):

* No MCP tool surface is available for `oraclepack` and Taskify workflows (agents cannot call validated tools to run/inspect packs and artifacts). (per user)

Expected Behavior:

* Agents can use MCP tools to:

  * Validate and inspect packs.
  * Run packs non-interactively to generate artifacts.
  * Detect/validate Stage-2 outputs and validate Stage-3 Action Packs before execution.
  * Summarize Stage-3 artifacts for immediate downstream consumption.

Requirements:

* MCP server implementation

  * Provide a Python MCP server project structure (e.g., `oraclepack-mcp-server/` with `oraclepack_mcp_server/` package).
  * Support transports:

    * `stdio`
    * `streamable-http`
* Tool surface (MCP tools)

  * `oraclepack_validate_pack`
  * `oraclepack_list_steps`
  * `oraclepack_run_pack` (execution gated)
  * `oraclepack_read_file`
  * `oraclepack_taskify_detect_stage2`
  * `oraclepack_taskify_validate_stage2`
  * `oraclepack_taskify_validate_action_pack`
  * `oraclepack_taskify_artifacts_summary`
  * `oraclepack_taskify_run_action_pack` (execution gated)
* Execution + safety controls

  * Deny-by-default execution; require `ORACLEPACK_ENABLE_EXEC=1` to enable “run” tools.
  * Restrict filesystem reads to allowlisted roots via `ORACLEPACK_ALLOWED_ROOTS` (colon-separated); block paths outside allowed roots.
  * Enforce timeouts and truncate stdout/stderr (`ORACLEPACK_CHARACTER_LIMIT`) and cap file read sizes (`ORACLEPACK_MAX_READ_BYTES`).
* Stage-2 reliability helpers

  * Validate Stage-2 directory contains exactly one match per prefix `01`..`20` (missing/ambiguous detection).
  * Deterministic Stage-2 detection:

    * Accept explicit dir or file.
    * If file is under `docs/oracle-questions-YYYY-MM-DD/...`, set out-dir to `docs/oracle-questions-YYYY-MM-DD/oracle-out`; otherwise default `repo_root/oracle-out`.
* Stage-3 reliability helpers

  * Validate Action Pack structure constraints before executing (e.g., “single ```bash fence, step headers”).
  * Summarize produced artifacts (examples cited: `_actions.json`, PRD path, Task Master outputs).
* Agent UX metadata

  * Apply MCP tool annotations:

    * validate/list/read: `readOnlyHint: true`
    * run: `destructiveHint: true`, `openWorldHint: true`

Out of Scope:

* Not provided.

Reproduction Steps:

* Not provided.

Environment:

* Language/runtime: Python (MCP server), wraps external `oraclepack` CLI.
* OS: Unknown
* Deployment: Unknown (local stdio vs HTTP service)
* MCP SDK version: Unknown (example uses `mcp>=1.0.0`, `pydantic>=2.0.0`).

Evidence:

* Conversation transcript + proposed reference implementation: `/mnt/data/MCP tools for Oraclepack.md`.
* Proposed env vars and tool list (deny-by-default exec, allowed roots, transports).
* Stage-2 directory validation (`01-*.md..20-*.md`) and Stage-2 out-dir resolution logic for `docs/oracle-questions-YYYY-MM-DD`.

Decisions / Agreements:

* Use a Python MCP server (FastMCP / MCP Python SDK) to expose `oraclepack` CLI + Taskify helpers.
* Support both `stdio` and `streamable-http` transports.
* Default-secure posture: execution gated by `ORACLEPACK_ENABLE_EXEC`, filesystem access constrained by allowlisted roots, truncation + timeout enforced.

Open Items / Unknowns:

* Where this MCP server should live (same repo as `oraclepack` vs separate repo) is not provided.
* Authentication / origin validation requirements for `streamable-http` deployment are mentioned conceptually but concrete requirements are not provided.
* Exact definition of “artifacts summary” contents/format beyond examples (`_actions.json`, PRD, Task Master outputs) is not provided.
* Whether `oraclepack_run_pack` must always use `--no-tui --yes --run-all` vs configurable flags is not provided (example suggests non-interactive flags).

Risks / Dependencies:

* Dependency on external `oraclepack` binary path/config (`ORACLEPACK_BIN`) and correct working directory (`ORACLEPACK_WORKDIR`).
* Security risk if exec is enabled without strict path/root controls and timeouts; mitigations are required as above.
* Stage-2 ambiguity risk when multiple `NN-*.md` match the same prefix; must report ambiguity deterministically and fail validation.

Acceptance Criteria:

* [ ] MCP server starts successfully in both `stdio` and `streamable-http` modes.
* [ ] All listed tools are exposed with the documented names.
* [ ] When `ORACLEPACK_ENABLE_EXEC!=1`, run tools refuse to execute and return a clear error; validate/list/read tools still work.
* [ ] `oraclepack_read_file` rejects paths outside `ORACLEPACK_ALLOWED_ROOTS`.
* [ ] Stage-2 validation enforces exactly one file per prefix `01`..`20` and returns `missing` and `ambiguous` sets when invalid.
* [ ] Stage-2 detection resolves out-dir correctly for both:

  * explicit dir inputs
  * single-pack file under `docs/oracle-questions-YYYY-MM-DD/...` → `docs/oracle-questions-YYYY-MM-DD/oracle-out`
  * otherwise default `repo_root/oracle-out`.
* [ ] stdout/stderr truncation and timeouts are enforced on CLI subprocess execution.
* [ ] Tool annotations are applied as specified for read-only vs destructive tools.

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement
* mcp
* oraclepack
* cli-wrapper
* taskify
* security
* tooling
```

.tickets/mcp/oraclepack-MCP.md
```
## MCP surface for `oraclepack` (so agents can act on Taskify artifacts)

### What to expose as MCP tools

Map the existing `oraclepack` CLI capabilities (validate/list/run + flags like `--no-tui`, `--out-dir`, `--oracle-bin`) into MCP tools so an agent can run packs non-interactively and then inspect the resulting state/report/artifacts.

Add a small “taskify helper” layer to make the **Stage-2 → Stage-3** workflow deterministic for agents:

* **Detect Stage-2 outputs** (dir-form `01-*.md..20-*.md` OR single-pack form) using the ordered resolver rules described in your skill.
* **Validate Stage-2 outputs** (exactly one match per prefix 01..20).
* **Validate Stage-3 Action Pack** structure constraints (single ```bash fence, step headers `# NN)`, etc.) before executing anything.
* **Summarize Stage-3 artifacts** (e.g., `_actions.json`, `_actions.md`, PRD path, `tm-complexity.json`, pipelines doc) so agents can immediately consume them.

### Transport choices (real-time vs local)

* **stdio** is simplest for local agent runtimes; it requires MCP messages only on stdout (logs must go to stderr). ([Model Context Protocol][1])
* **Streamable HTTP** is better for “real-time” multi-client usage; implement Origin validation and bind to localhost + auth to avoid DNS rebinding and local-network abuse. ([Model Context Protocol][1])

### Tool UX metadata (important for agents)

Use MCP **tool annotations** so clients can correctly present approval UX:

* mark validate/list/read tools as `readOnlyHint: true`
* mark run tools as `destructiveHint: true`, `openWorldHint: true` (they can touch filesystem, run processes, etc.) ([Model Context Protocol][2])

### Security defaults (recommended)

* Hard **deny-by-default execution**: require an env flag (e.g. `ORACLEPACK_ENABLE_EXEC=1`) to enable “run” tools.
* Restrict filesystem access to **allowlisted roots** to avoid path traversal and accidental exfiltration of secrets.
* Enforce stdout/stderr truncation and timeouts so a pack can’t wedge the server process.

## Reference implementation

I provided a complete, drop-in **Python MCP server** (FastMCP from the official MCP Python SDK) that exposes:

* `oraclepack_validate_pack`, `oraclepack_list_steps`, `oraclepack_run_pack`
* `oraclepack_taskify_detect_stage2`, `oraclepack_taskify_validate_stage2`
* `oraclepack_taskify_validate_action_pack`, `oraclepack_taskify_artifacts_summary`
* `oraclepack_taskify_run_action_pack`

It supports **stdio** and **streamable-http** transports per the official SDK. ([GitHub][3])

The full multi-file project is in the canvas titled **“oraclepack-mcp-server”**.

[1]: https://modelcontextprotocol.io/specification/2025-03-26/basic/transports "Transports - Model Context Protocol"
[2]: https://modelcontextprotocol.io/legacy/concepts/tools "Tools - Model Context Protocol"
[3]: https://github.com/modelcontextprotocol/python-sdk "GitHub - modelcontextprotocol/python-sdk: The official Python SDK for Model Context Protocol servers and clients"
```

.tickets/mcp/oraclepack_mcp_server.md
```
# oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

## tree

```text
oraclepack-mcp-server/
  README.md
  requirements.txt
  oraclepack_mcp_server/
    __init__.py
    __main__.py
    config.py
    security.py
    oraclepack_cli.py
    taskify.py
    server.py
```

---

```md
<!-- path: oraclepack-mcp-server/README.md -->
# oraclepack-mcp-server

Expose `oraclepack` (validate/list/run) plus a few Stage-3 (oraclepack-taskify) helpers as **MCP tools**, so an agent can:
- validate and inspect packs
- run packs non-interactively (`--no-tui --yes --run-all`) to produce artifacts
- validate Stage-2 outputs (01-*.md..20-*.md)
- validate Stage-3 Action Packs (single ```bash fence, step headers, etc.)
- summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)

## Install

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Configure (recommended)

Environment variables:

- `ORACLEPACK_ALLOWED_ROOTS` (default: current working directory)
  - Colon-separated list of allowed filesystem roots the server may read from.
  - Example: `ORACLEPACK_ALLOWED_ROOTS=/repo:/tmp/oracle-out`
- `ORACLEPACK_BIN` (default: `oraclepack`) – path to the oraclepack CLI
- `ORACLEPACK_WORKDIR` (default: current working directory)
  - Where packs are executed from (important for relative paths).
- `ORACLEPACK_ENABLE_EXEC` (default: `0`)
  - Must be `1` to enable `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack`.
- `ORACLEPACK_CHARACTER_LIMIT` (default: `25000`) – truncate large stdout/stderr
- `ORACLEPACK_MAX_READ_BYTES` (default: `500000`) – max bytes read from a file

## Run (stdio)

```bash
# Stdio transport is the simplest local integration.
python -m oraclepack_mcp_server --transport stdio
```

## Run (Streamable HTTP)

```bash
python -m oraclepack_mcp_server --transport streamable-http
```

## Tools

- `oraclepack_validate_pack`
- `oraclepack_list_steps`
- `oraclepack_run_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)
- `oraclepack_read_file`
- `oraclepack_taskify_detect_stage2`
- `oraclepack_taskify_validate_stage2`
- `oraclepack_taskify_validate_action_pack`
- `oraclepack_taskify_artifacts_summary`
- `oraclepack_taskify_run_action_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)

## Typical Stage-3 usage

1) Detect/validate Stage-2 outputs (directory or single-pack)
2) Validate the Action Pack markdown
3) Run the Action Pack via `oraclepack run ...`
4) Summarize produced artifacts
```

```txt
# path: oraclepack-mcp-server/requirements.txt
mcp>=1.0.0
pydantic>=2.0.0
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
__all__ = []
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/config.py
from __future__ import annotations

import os
from dataclasses import dataclass
from pathlib import Path


def _truthy(value: str | None) -> bool:
    if value is None:
        return False
    return value.strip().lower() in {"1", "true", "yes", "y", "on"}


@dataclass(frozen=True)
class ServerConfig:
    allowed_roots: tuple[Path, ...]
    oraclepack_bin: str
    workdir: Path
    enable_exec: bool
    character_limit: int
    max_read_bytes: int


def load_config() -> ServerConfig:
    # Allowed roots: colon-separated. Default to CWD.
    roots_raw = os.environ.get("ORACLEPACK_ALLOWED_ROOTS")
    if roots_raw:
        roots = tuple(Path(p).expanduser().resolve() for p in roots_raw.split(":") if p.strip())
    else:
        roots = (Path.cwd().resolve(),)

    oraclepack_bin = os.environ.get("ORACLEPACK_BIN", "oraclepack")
    workdir = Path(os.environ.get("ORACLEPACK_WORKDIR", str(Path.cwd()))).expanduser().resolve()

    enable_exec = _truthy(os.environ.get("ORACLEPACK_ENABLE_EXEC", "0"))

    character_limit = int(os.environ.get("ORACLEPACK_CHARACTER_LIMIT", "25000"))
    max_read_bytes = int(os.environ.get("ORACLEPACK_MAX_READ_BYTES", "500000"))

    return ServerConfig(
        allowed_roots=roots,
        oraclepack_bin=oraclepack_bin,
        workdir=workdir,
        enable_exec=enable_exec,
        character_limit=character_limit,
        max_read_bytes=max_read_bytes,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/security.py
from __future__ import annotations

from pathlib import Path


class PathNotAllowedError(ValueError):
    pass


def resolve_under_roots(path: Path, allowed_roots: tuple[Path, ...]) -> Path:
    """Resolve a path and enforce it lives under at least one allowed root."""
    p = path.expanduser().resolve()

    for root in allowed_roots:
        r = root.expanduser().resolve()
        try:
            p.relative_to(r)
            return p
        except ValueError:
            continue

    raise PathNotAllowedError(
        f"Path not allowed (outside allowed roots). path={p} roots={[str(r) for r in allowed_roots]}"
    )


def safe_read_text(path: Path, max_bytes: int) -> tuple[str, bool]:
    """Read up to max_bytes from a file as UTF-8 (replace errors)."""
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data.decode("utf-8", errors="replace"), truncated


def safe_read_bytes(path: Path, max_bytes: int) -> tuple[bytes, bool]:
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data, truncated
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
from __future__ import annotations

import asyncio
import os
import time
from dataclasses import dataclass
from pathlib import Path


@dataclass
class CmdResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool


def _truncate(s: str, limit: int) -> tuple[str, bool]:
    if limit <= 0:
        return s, False
    if len(s) <= limit:
        return s, False
    return s[:limit], True


async def run_cmd(
    argv: list[str],
    cwd: Path,
    timeout_s: int,
    env: dict[str, str] | None,
    character_limit: int,
) -> CmdResult:
    start = time.time()

    proc = await asyncio.create_subprocess_exec(
        *argv,
        cwd=str(cwd),
        env=(os.environ | (env or {})),
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )

    try:
        out_b, err_b = await asyncio.wait_for(proc.communicate(), timeout=timeout_s)
    except asyncio.TimeoutError:
        proc.kill()
        await proc.communicate()
        duration = time.time() - start
        return CmdResult(
            ok=False,
            exit_code=124,
            duration_s=duration,
            stdout="",
            stderr=f"Timed out after {timeout_s}s: {' '.join(argv)}",
            stdout_truncated=False,
            stderr_truncated=False,
        )

    duration = time.time() - start
    out = out_b.decode("utf-8", errors="replace") if out_b else ""
    err = err_b.decode("utf-8", errors="replace") if err_b else ""

    out, out_tr = _truncate(out, character_limit)
    err, err_tr = _truncate(err, character_limit)

    exit_code = proc.returncode if proc.returncode is not None else 1
    return CmdResult(
        ok=(exit_code == 0),
        exit_code=exit_code,
        duration_s=duration,
        stdout=out,
        stderr=err,
        stdout_truncated=out_tr,
        stderr_truncated=err_tr,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
from __future__ import annotations

import re
from dataclasses import dataclass
from datetime import date
from pathlib import Path


@dataclass
class Stage2Resolution:
    kind: str  # "dir" | "file"
    stage2_path: Path
    out_dir: Path
    notes: list[str]


PREFIXES = [f"{i:02d}" for i in range(1, 21)]


def _is_iso_date(s: str) -> bool:
    # Minimal heuristic for YYYY-MM-DD
    return bool(re.fullmatch(r"\d{4}-\d{2}-\d{2}", s))


def validate_stage2_dir(out_dir: Path) -> dict:
    missing: list[str] = []
    ambiguous: dict[str, list[str]] = {}
    selected: dict[str, str] = {}

    for pfx in PREFIXES:
        matches = sorted(out_dir.glob(f"{pfx}-*.md"))
        if len(matches) == 0:
            missing.append(f"{pfx}-*.md")
        elif len(matches) > 1:
            ambiguous[pfx] = [m.name for m in matches]
        else:
            selected[pfx] = matches[0].name

    ok = (not missing) and (not ambiguous)
    return {
        "ok": ok,
        "out_dir": str(out_dir),
        "selected": selected,
        "missing": missing,
        "ambiguous": ambiguous,
    }


def _lexi_newest(paths: list[Path]) -> Path | None:
    # Deterministic: lexicographic max
    return sorted(paths, key=lambda p: p.name)[-1] if paths else None


def detect_stage2(stage2_path: str, repo_root: Path) -> Stage2Resolution:
    notes: list[str] = []

    if stage2_path != "auto":
        p = (repo_root / stage2_path).expanduser()
        if p.exists() and p.is_dir():
            return Stage2Resolution(kind="dir", stage2_path=p.resolve(), out_dir=p.resolve(), notes=["explicit dir"])
        if p.exists() and p.is_file():
            # out_dir rules from skill: if under docs/oracle-questions-YYYY-MM-DD/ then parent/oracle-out else oracle-out
            p_res = p.resolve()
            out = repo_root / "oracle-out"
            parts = list(p_res.parts)
            if "docs" in parts:
                try:
                    idx = parts.index("docs")
                    # docs/oracle-questions-YYYY-MM-DD/.../oracle-pack-YYYY-MM-DD.md
                    if idx + 1 < len(parts) and parts[idx + 1].startswith("oracle-questions-"):
                        out = Path(*parts[: idx + 2]) / "oracle-out"
                except ValueError:
                    pass
            return Stage2Resolution(kind="file", stage2_path=p_res, out_dir=out.resolve(), notes=["explicit file"])
        raise FileNotFoundError(f"stage2_path not found: {p}")

    # auto discovery (best-effort, deterministic ordering)
    searched: list[str] = []

    # 1) repo_root/oracle-out
    candidate = repo_root / "oracle-out"
    searched.append(str(candidate))
    if candidate.is_dir():
        v = validate_stage2_dir(candidate)
        if v["ok"]:
            notes.append("auto: selected repo_root/oracle-out")
            return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 2) docs/oracle-questions-YYYY-MM-DD/oracle-out (newest by lexicographic date suffix)
    docs = repo_root / "docs"
    if docs.is_dir():
        qdirs = [p for p in docs.glob("oracle-questions-*") if p.is_dir()]
        # deterministic: sort by name and take newest
        newest_q = _lexi_newest(qdirs)
        if newest_q:
            candidate = newest_q / "oracle-out"
            searched.append(str(candidate))
            if candidate.is_dir():
                v = validate_stage2_dir(candidate)
                if v["ok"]:
                    notes.append(f"auto: selected {candidate}")
                    return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 3) single-pack form (newer): look for docs/oracle-pack-*.md or docs/oraclepacks/oracle-pack-*.md
    file_candidates: list[Path] = []
    if docs.is_dir():
        file_candidates += list(docs.glob("oracle-pack-*.md"))
        file_candidates += list((docs / "oraclepacks").glob("oracle-pack-*.md"))
        file_candidates += list(docs.glob("oracle-questions-*/oraclepacks/oracle-pack-*.md"))

    newest_file = _lexi_newest(sorted([p for p in file_candidates if p.is_file()], key=lambda p: p.name))
    if newest_file:
        notes.append(f"auto: selected single-pack {newest_file}")
        out = repo_root / "oracle-out"
        # If under docs/oracle-questions-YYYY-MM-DD/..., default out_dir there.
        if "docs" in newest_file.parts:
            try:
                idx = newest_file.parts.index("docs")
                if idx + 1 < len(newest_file.parts) and newest_file.parts[idx + 1].startswith("oracle-questions-"):
                    out = Path(*newest_file.parts[: idx + 2]) / "oracle-out"
            except ValueError:
                pass
        return Stage2Resolution(kind="file", stage2_path=newest_file.resolve(), out_dir=out.resolve(), notes=notes)

    raise FileNotFoundError(
        "stage2_path=auto could not resolve Stage-2 outputs. Searched:\n- " + "\n- ".join(searched)
    )


def validate_action_pack(pack_path: Path) -> dict:
    text = pack_path.read_text(encoding="utf-8", errors="replace")

    bash_fence = re.findall(r"(?m)^\s*```bash\s*$", text)
    any_fence = re.findall(r"(?m)^\s*```\s*", text)

    errors: list[str] = []
    if len(bash_fence) != 1:
        errors.append(f"expected exactly one ```bash fence; found {len(bash_fence)}")
    if len(any_fence) != 2:
        # One opening and one closing fence expected, and it must be a bash fence.
        errors.append(f"expected no other code fences; found {len(any_fence)} total fences")

    # Extract bash block content if possible
    bash_block = ""
    m = re.search(r"```bash\s*\n(?P<body>[\s\S]*?)\n```\s*", text)
    if m:
        bash_block = m.group("body")

    # Validate step headers inside bash fence
    step_headers = re.findall(r"(?m)^\s*#\s*(\d{2})\)\s+.*$", bash_block)
    if not step_headers:
        errors.append("no step headers found inside the bash fence (expected lines like '# 01) ...')")
    else:
        # Ensure they start at 01 and are strictly increasing by 1.
        nums = [int(x) for x in step_headers]
        if nums[0] != 1:
            errors.append(f"first step must be 01; got {nums[0]:02d}")
        for prev, cur in zip(nums, nums[1:]):
            if cur != prev + 1:
                errors.append(f"step numbers must be sequential; got {prev:02d} then {cur:02d}")

    return {
        "ok": len(errors) == 0,
        "pack_path": str(pack_path),
        "step_count": len(step_headers),
        "errors": errors,
    }


def default_pack_path(today: date | None = None) -> str:
    d = today or date.today()
    return f"docs/oracle-actions-pack-{d.isoformat()}.md"
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/server.py
from __future__ import annotations

import json
from enum import Enum
from pathlib import Path
from typing import Any

from pydantic import BaseModel, Field
from mcp.server.fastmcp import FastMCP

from .config import load_config
from .security import resolve_under_roots, safe_read_text, PathNotAllowedError
from .oraclepack_cli import run_cmd
from .taskify import detect_stage2, validate_stage2_dir, validate_action_pack


class ResponseFormat(str, Enum):
    markdown = "markdown"
    json = "json"


class PackPathInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class RunPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    out_dir: str | None = Field(default=None, description="Output directory for step execution (passes --out-dir).")

    no_tui: bool = Field(default=True, description="If true, pass --no-tui")
    yes: bool = Field(default=True, description="If true, pass --yes")
    run_all: bool = Field(default=True, description="If true, pass --run-all")

    resume: bool = Field(default=False, description="If true, pass --resume")
    stop_on_fail: bool = Field(default=True, description="If true, pass --stop-on-fail (default true)")

    roi_threshold: float = Field(default=0.0, description="Pass --roi-threshold")
    roi_mode: str = Field(default="over", description="Pass --roi-mode ('over' or 'under')")

    timeout_s: int = Field(default=3600, description="Hard timeout for the oraclepack process")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class ReadFileInput(BaseModel):
    path: str = Field(..., description="Path to a file within ORACLEPACK_ALLOWED_ROOTS")
    max_bytes: int | None = Field(default=None, description="Override max bytes read")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyDetectStage2Input(BaseModel):
    stage2_path: str = Field(default="auto", description="Dir or file, or 'auto'")
    repo_root: str = Field(default=".", description="Repo root for relative resolution")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateStage2Input(BaseModel):
    out_dir: str = Field(..., description="Directory that should contain 01-*.md..20-*.md")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to Stage-3 Action Pack markdown")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyArtifactsSummaryInput(BaseModel):
    out_dir: str = Field(..., description="Stage-3 out_dir (where _actions.json etc are written)")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyRunActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the Stage-3 Action Pack markdown")
    out_dir: str | None = Field(default=None, description="Pass --out-dir for execution")
    timeout_s: int = Field(default=7200)
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


cfg = load_config()

mcp = FastMCP(
    name="oraclepack-mcp-server",
    # For production Streamable HTTP deployments, stateless_http + json_response is recommended.
    # Clients may override by running behind an ASGI app if needed.
    stateless_http=True,
    json_response=True,
)


def _md_codeblock(lang: str, content: str) -> str:
    return f"```{lang}\n{content}\n```"


def _format_cmd_result(result: Any, response_format: ResponseFormat) -> Any:
    if response_format == ResponseFormat.json:
        return {
            "ok": result.ok,
            "exit_code": result.exit_code,
            "duration_s": result.duration_s,
            "stdout": result.stdout,
            "stderr": result.stderr,
            "stdout_truncated": result.stdout_truncated,
            "stderr_truncated": result.stderr_truncated,
        }

    lines = []
    lines.append(f"**ok**: {result.ok}")
    lines.append(f"**exit_code**: {result.exit_code}")
    lines.append(f"**duration_s**: {result.duration_s:.2f}")
    lines.append("")

    if result.stdout:
        lines.append("## stdout")
        lines.append(_md_codeblock("text", result.stdout))
        if result.stdout_truncated:
            lines.append(f"(stdout truncated to {cfg.character_limit} chars)")

    if result.stderr:
        lines.append("## stderr")
        lines.append(_md_codeblock("text", result.stderr))
        if result.stderr_truncated:
            lines.append(f"(stderr truncated to {cfg.character_limit} chars)")

    return "\n".join(lines)


def _ensure_exec_enabled() -> None:
    if not cfg.enable_exec:
        raise PermissionError(
            "Execution is disabled. Set ORACLEPACK_ENABLE_EXEC=1 to enable pack execution tools."
        )


@mcp.tool(
    name="oraclepack_validate_pack",
    annotations={
        "title": "Validate oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_validate_pack(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "validate", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_list_steps",
    annotations={
        "title": "List steps in an oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_list_steps(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "list", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_run_pack",
    annotations={
        "title": "Run an oraclepack pack (non-interactive)",
        "readOnlyHint": False,
        "destructiveHint": True,
        "idempotentHint": False,
        "openWorldHint": True,
    },
)
async def oraclepack_run_pack(params: RunPackInput) -> Any:
    _ensure_exec_enabled()

    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv: list[str] = [cfg.oraclepack_bin]
    if params.no_tui:
        argv += ["--no-tui"]
    if params.out_dir:
        out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
        argv += ["--out-dir", str(out_dir)]

    argv += ["run"]

    if params.yes:
        argv += ["--yes"]
    if params.run_all:
        argv += ["--run-all"]

    if params.resume:
        argv += ["--resume"]

    # stop-on-fail is default true in oraclepack; pass explicitly for clarity.
    argv += ["--stop-on-fail", "true" if params.stop_on_fail else "false"]

    argv += ["--roi-threshold", str(params.roi_threshold), "--roi-mode", params.roi_mode]

    argv += [str(pack_path)]

    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=params.timeout_s, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_read_file",
    annotations={
        "title": "Read a file under allowed roots",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_read_file(params: ReadFileInput) -> Any:
    p = resolve_under_roots(Path(params.path), cfg.allowed_roots)
    if not p.exists() or not p.is_file():
        raise FileNotFoundError(f"file not found: {p}")

    max_bytes = params.max_bytes if params.max_bytes is not None else cfg.max_read_bytes
    text, truncated = safe_read_text(p, max_bytes=max_bytes)

    if params.response_format == ResponseFormat.json:
        return {"path": str(p), "truncated": truncated, "content": text}

    note = f"\n\n(content truncated to {max_bytes} bytes)" if truncated else ""
    return f"# {p.name}\n\n{_md_codeblock('text', text)}{note}"


@mcp.tool(
    name="oraclepack_taskify_detect_stage2",
    annotations={
        "title": "Detect Stage-2 oraclepack outputs",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_detect_stage2(params: TaskifyDetectStage2Input) -> Any:
    repo_root = resolve_under_roots(Path(params.repo_root), cfg.allowed_roots)
    res = detect_stage2(stage2_path=params.stage2_path, repo_root=repo_root)

    payload = {
        "kind": res.kind,
        "stage2_path": str(res.stage2_path),
        "out_dir": str(res.out_dir),
        "notes": res.notes,
    }

    if params.response_format == ResponseFormat.json:
        return payload

    lines = ["# Stage-2 resolution", ""]
    lines.append(f"- **kind**: `{payload['kind']}`")
    lines.append(f"- **stage2_path**: `{payload['stage2_path']}`")
    lines.append(f"- **out_dir**: `{payload['out_dir']}`")
    if payload["notes"]:
        lines.append("- **notes**:")
        lines += [f"  - {n}" for n in payload["notes"]]
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_validate_stage2",
    annotations={
        "title": "Validate Stage-2 directory outputs (01-*.md..20-*.md)",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_validate_stage2(params: TaskifyValidateStage2Input) -> Any:
    out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
    if not out_dir.is_dir():
        raise FileNotFoundError(f"out_dir is not a directory: {out_dir}")

    report = validate_stage2_dir(out_dir)
    if params.response_format == ResponseFormat.json:
        return report

    lines = [f"# Stage-2 validation: `{out_dir}`", ""]
    lines.append(f"- **ok**: {report['ok']}")
    if report["missing"]:
        lines.append("- **missing**:")
        lines += [f"  - {m}" for m in report["missing"]]
    if report["ambiguous"]:
        lines.append("- **ambiguous**:")
        for k, v in report["ambiguous"].items():
            lines.append(f"  - {k}: {v}")
    if report["selected"]:
        lines.append("- **selected**:")
        for k in sorted(report["selected"].keys()):
            lines.append(f"  - {k}: {report['selected'][k]}")
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_validate_action_pack",
    annotations={
        "title": "Validate Stage-3 Action Pack structure",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_validate_action_pack(params: TaskifyValidateActionPackInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)
    if not pack_path.is_file():
        raise FileNotFoundError(f"pack not found: {pack_path}")

    report = validate_action_pack(pack_path)
    if params.response_format == ResponseFormat.json:
        return report

    lines = [f"# Action Pack validation: `{pack_path}`", ""]
    lines.append(f"- **ok**: {report['ok']}")
    lines.append(f"- **step_count**: {report['step_count']}")
    if report["errors"]:
        lines.append("- **errors**:")
        lines += [f"  - {e}" for e in report["errors"]]
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_artifacts_summary",
    annotations={
        "title": "Summarize expected Stage-3 artifacts",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_artifacts_summary(params: TaskifyArtifactsSummaryInput) -> Any:
    out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)

    # Known artifacts from the skill
    expected = {
        "actions_json": out_dir / "_actions.json",
        "actions_md": out_dir / "_actions.md",
        "tm_complexity": out_dir / "tm-complexity.json",
        "prd": Path(cfg.workdir) / ".taskmaster" / "docs" / "oracle-actions-prd.md",
        "pipelines_doc": Path(cfg.workdir) / "docs" / "oracle-actions-pipelines.md",
        "taskmaster_tasks_local": Path(cfg.workdir) / "tasks.json",
        "taskmaster_tasks_dot": Path(cfg.workdir) / ".taskmaster" / "tasks.json",
    }

    found: dict[str, str] = {}
    missing: dict[str, str] = {}
    for k, p in expected.items():
        # Some paths may not be under allowed_roots; report as "not allowed" rather than crashing.
        try:
            p2 = resolve_under_roots(p, cfg.allowed_roots)
        except PathNotAllowedError:
            missing[k] = f"not allowed by ORACLEPACK_ALLOWED_ROOTS: {p}"
            continue

        if p2.exists():
            found[k] = str(p2)
        else:
            missing[k] = str(p2)

    payload = {"out_dir": str(out_dir), "found": found, "missing": missing}

    if params.response_format == ResponseFormat.json:
        return payload

    lines = [f"# Stage-3 artifacts summary: `{out_dir}`", ""]
    if found:
        lines.append("## found")
        for k, v in found.items():
            lines.append(f"- **{k}**: `{v}`")
    if missing:
        lines.append("\n## missing")
        for k, v in missing.items():
            lines.append(f"- **{k}**: `{v}`")
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_run_action_pack",
    annotations={
        "title": "Run a Stage-3 Action Pack via oraclepack",
        "readOnlyHint": False,
        "destructiveHint": True,
        "idempotentHint": False,
        "openWorldHint": True,
    },
)
async def oraclepack_taskify_run_action_pack(params: TaskifyRunActionPackInput) -> Any:
    _ensure_exec_enabled()

    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    # Validate pack structure before executing
    v = validate_action_pack(pack_path)
    if not v["ok"]:
        raise ValueError(f"Refusing to execute invalid Action Pack: {json.dumps(v, indent=2)}")

    argv: list[str] = [cfg.oraclepack_bin]
    argv += ["--no-tui"]
    if params.out_dir:
        out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
        argv += ["--out-dir", str(out_dir)]

    argv += ["run", "--yes", "--run-all", "--stop-on-fail", "true", str(pack_path)]

    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=params.timeout_s, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__main__.py
from __future__ import annotations

import argparse

from .server import mcp


def main() -> None:
    parser = argparse.ArgumentParser(description="oraclepack MCP server")
    parser.add_argument(
        "--transport",
        default="stdio",
        choices=["stdio", "streamable-http"],
        help="MCP transport (stdio or streamable-http)",
    )
    args = parser.parse_args()

    # The official Python SDK supports stdio and streamable-http transports.
    mcp.run(transport=args.transport)


if __name__ == "__main__":
    main()
```

```

docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md
```
1) Direct answer (1–6 bullets, evidence-cited)

Primary public interface #1: CLI root command oraclepack (Cobra)

Inputs: persistent flags

--no-tui (bool, default false) to disable TUI

--oracle-bin (string, default "oracle") path to oracle binary

--out-dir / -o (string, default "") output directory for step execution

Outputs: process exit code is mapped via errors.ExitCode(err) when rootCmd.Execute() fails.

Defined in: internal/cli/root.go (root command + persistent flags + Execute error handling).

Primary public interface #2: CLI command validate [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: on success prints Pack is valid.; on failure returns an error from a.LoadPack() (then surfaced by root Execute path).

Defined in: internal/cli/cmds.go (validate command + wiring into root).

Primary public interface #3: CLI command list [pack.md]

Inputs: one positional arg pack.md (cobra.ExactArgs(1)), passed as app.Config{PackPath: args[0]}

Outputs: prints one line per step: "<step.ID>: <step.OriginalLine>", iterating a.Pack.Steps.

Defined in: internal/cli/cmds.go (list command + wiring into root).

Primary public interface #4: “oracle-pack” Markdown contract (pack file)

Inputs (implied by CLI): a Markdown file path (pack.md) that can be loaded by a.LoadPack() from app.New(cfg) where cfg.PackPath is set.

Outputs (observable via list): the parsed pack must produce a.Pack.Steps, and each step must at least have fields ID and OriginalLine (since they’re printed).

Where defined in code: only the usage is shown in internal/cli/cmds.go; the actual contract/schema + parser are in the internal/app package (not included in the provided files).

2) Risks/unknowns (bullets)

Additional CLI commands/flags may exist in other internal/cli/*.go files; only validate and list are visible here.

The actual oracle-pack Markdown schema (required headings, step syntax, validation rules, parsing behavior, versioning) is not visible; we only know LoadPack() yields Pack.Steps[] with ID and OriginalLine.

Exit-code mapping and error taxonomy depend on internal/errors (only referenced, not shown).

--oracle-bin / --out-dir are declared at root-level but not used in the shown commands; actual execution interface likely exists in other commands not included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for the contract types and parser entrypoints: rg -n "type Pack|type Step|LoadPack\\(|PackPath|OriginalLine" internal/ to locate the canonical oracle-pack Markdown schema definitions.

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/app/*.go (especially files defining Config, New, LoadPack, and any Pack / Step structs)

internal/errors/*.go (to confirm public error*
```

docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md
```
1) Direct answer (top integration points, minimal defining files/locations)

Shell execution boundary (runs arbitrary scripts via OS process execution)

Integration: os/exec launches a shell with -lc (login shell semantics), sets working dir, env, and streams stdout/stderr. This is the core “external system” boundary for running everything.

Contract/config: RunnerOptions{Shell, WorkDir, Env, OracleFlags, Overrides, ChatGPTURL} → Runner (constructed in NewRunner).

Defined in: internal/exec/runner.go (NewRunner, RunPrelude, RunStep, run). Evidence: exec.CommandContext(ctx, r.Shell, "-lc", script), cmd.Dir = r.WorkDir, cmd.Env = r.Env, cmd.Stdout/Stderr = logWriter.

Oracle CLI invocation boundary (flag injection + per-step overrides + optional dry-run validation)

Integration: Oracle commands are expected to appear inside step code; runner injects flags into those commands before execution, and validation can run oracle --dry-run summary for overridden steps.

Contract/config: baseline Runner.OracleFlags + optional overrides.RuntimeOverrides (scoped by step.ID) + optional ChatGPTURL mapped into flags (ApplyChatGPTURL); validation is scoped by over.ApplyToSteps[step.ID].

Defined in:

internal/exec/runner.go (RunStep: ApplyChatGPTURL, Overrides.EffectiveFlags, InjectFlags).

internal/exec/oracle_validate.go (ValidateOverrides: ExtractOracleInvocations, --dry-run summary, execDryRun).

Note: The actual parsing/injection logic lives elsewhere (functions referenced but not shown): ExtractOracleInvocations, InjectFlags, ApplyChatGPTURL.

Environment-variable boundary (inherit + append + PATH propagation during validation)

Integration: Environment is inherited from the parent process and extended with caller-provided entries; validation explicitly re-exports PATH into the command string to ensure the shell sees it.

Contract/config: RunnerOptions.Env merged with os.Environ(); findEnvValue(env, "PATH") drives PATH export behavior in dry-run validation.

Defined in:

internal/exec/runner.go (NewRunner: Env: append(os.Environ(), opts.Env...)).

internal/exec/oracle_validate.go (env := append(os.Environ(), opts.Env...), execDryRun PATH export + quoting via shellQuote).

Filesystem outputs boundary (working directory as the implicit artifact root)

Integration: All script execution happens with cmd.Dir = r.WorkDir, so any relative file reads/writes done by steps land under that directory (and any output directories are created by the scripts themselves).

Contract/config: RunnerOptions.WorkDir / Runner.WorkDir.

Defined in: internal/exec/runner.go (run: cmd.Dir = r.WorkDir), plus va
```

docs/oracle-questions-2026-01-07/03-invariants-domain.md
```
1) Direct answer (1–8 bullets, evidence-cited)

Pack must contain a fenced ```bash ... ``` code block (the pack is parsed only from that block).
Enforced: Parse() searches bashFenceRegex and errors if none found (no bash code block found).
Not enforced / edge: if multiple bash blocks exist, only the first match is used (FindSubmatch), others are ignored.
Evidence: bashFenceRegex = ... "```bash\n(.*?)\n```" and if match == nil ... return ... ErrInvalidPack in internal/pack/parser.go.

Steps are defined only by Markdown-style headers inside the bash block, and must start with # + a 2-digit step id plus an accepted separator.
Enforced (implicitly by parsing): step boundaries are recognized only when stepHeaderRegex matches a line; everything else is either prelude (before first step) or step body.
Invariants implied: step IDs must be exactly two digits (\d{2}) and separator must be one of ), —, or - (with whitespace variants).
Evidence: stepHeaderRegex = regexp.MustCompile(^#\s*(\d{2})(?:\)|[\s]+[—-])) and the headerMatch := ...; if len(headerMatch) > 1 { ... } logic in internal/pack/parser.go.

All lines before the first step header become “Prelude”; all lines after step parsing begins are assigned to the current step’s Code.
Enforced: inSteps toggles true at first header; before that lines go to preludeLines; after that lines append to currentStep.Code.
Should be considered: any configuration intended for metadata extraction must appear in the prelude, or it won’t be discovered.
Evidence: var preludeLines []string; var inSteps bool + if inSteps { currentStep.Code += ... } else { preludeLines = append(...) } and pack.Prelude.Code = strings.Join(preludeLines, "\n") in internal/pack/parser.go.

At least one step is required for a valid pack.
Enforced: Validate() rejects packs with len(p.Steps) == 0. Note: Parse() itself does not call Validate()—caller must.
Evidence: if len(p.Steps) == 0 { return ... ErrInvalidPack } in internal/pack/parser.go.

Step numbers must be positive, unique, and sequential starting at 1 in the encountered order (i.e., the file order is the order).
Enforced: Validate() checks step.Number > 0, no duplicates (seen map), and step.Number == i+1. This implies the first step must be 01 (because 00 parses to 0 and fails).
Evidence: seen := make(map[int]bool) + duplicate check + if step.Number != i+1 { ... sequential starting from 1 ... } in internal/pack/parser.go.

Metadata invariants: out_dir=... and --write-output are derived only from the Prelude, with strict-ish patterns.
Enforced (extraction behavior): DeriveMetadata() looks only at p.Prelude.Code; out_dir must match ^out_dir=... with no spaces around =, and --write-output is a simple substring/regex presence test.
Not enforced: packs can place these in steps and still “work” at runtime, but metadata will be wrong; multiple out_dir= lines are not resolved explicitly (first match_
```

docs/oracle-questions-2026-01-07/04-invariants-validation.md
```
1) Direct answer (validation boundaries + likely gaps)

Pack parse/structural validation happens in App.LoadPack(): it reads the pack file, calls pack.Parse(data), then p.Validate() and fails fast on any error (internal/app/app.go:LoadPack).

Boundary: “bytes on disk” → “in-memory Pack that is considered valid enough to run.”

Parse-time validation is minimal and mostly syntactic: pack.Parse() rejects packs without a fenced ```bash block (bashFenceRegex) and only recognizes steps whose headers match stepHeaderRegex (internal/pack/parser.go:Parse).

Gap: a Markdown pack can be “accepted” but silently miss intended steps if headers don’t match the regex (e.g., # 1 - ... vs # 01 - ...), leading to Validate() failing later with “at least one step is required” or producing an unexpected step set.

Pack semantic validation is narrow (step numbering only): Pack.Validate() enforces: at least one step, positive numbers, no duplicates, and sequential numbering starting from 1 in parse order (internal/pack/parser.go:Validate).

Gap: it does not validate step code presence/emptiness, command safety, required prelude directives, or cross-step dependencies—so many “bad” packs pass validation and fail at runtime.

Metadata extraction happens without validation: after parsing, DeriveMetadata() extracts out_dir=... and --write-output from the prelude using regexes (internal/pack/parser.go:DeriveMetadata).

Gap: extracted OutDir is not validated for emptiness, path traversal, or consistency with CLI overrides (it’s just a string captured by regex).

Config resolution validation is effectively “filesystem-only”: App.Prepare() resolves outDir by precedence CLI cfg.OutDir > pack p.OutDir > ".", then only validates by attempting os.MkdirAll(outDir, 0755) (internal/app/app.go:Prepare).

Gap: there’s no validation of other config fields here (PackPath, StatePath, ROIMode values, ROIThreshold range, etc.), so invalid config can survive until deeper runtime code.

State validation boundary is weak and can mask corruption: App.LoadState() on Resume tries state.LoadState(...); if it errors, it silently falls back to a fresh RunState{SchemaVersion:1, StepStatuses: map[...]} (internal/app/app.go:LoadState).

Most likely inconsistent-state gap: partial outputs on disk + corrupted/old state file ⇒ the system “resumes” from an empty state and may rerun steps unexpectedly or skip intended recovery paths.

Runtime checks shown are limited to environment shaping: Prepare() appends out_dir=<resolved> into Runner.Env but does not persist the resolved outDir back into Config/State, and does not de-dup/override existing out_dir entries (internal/app/app.go:Prepare).

Gap: scripts may see multiple out_dir values depending on how Runner.Env is constructed elsewhere; plus “where outputs went” is not recorded in state, increasing resume/report inconsistency risk.

Parser robustness gap that can cause silent truncation: pack.Parse() uses bufio.Scanner but does not call scanner.Err() after the scan loop, nor increase the scanner buffer (internal/pack/parser.go:Parse).

Likely inconsistent-state failure mode: long lines in bash blocks can exceed Scanner’s token lim
```

docs/oracle-questions-2026-01-07/05-caching-state-layers.md
```
Direct answer (1–8 bullets, evidence-cited)

Persisted run-state file (RunState JSON)

What it is: A single JSON document representing an oracle-pack run: schema version, pack hash, start time, per-step status map, ROI settings, and warnings. (internal/state/types.go: RunState{SchemaVersion, PackHash, StartTime, StepStatuses, ROIThreshold, ROIMode, Warnings})

Lifecycle (as implemented in this package): In-memory RunState is serialized to disk and later deserialized back into memory. (internal/state/persist.go: SaveStateAtomic, LoadState)

Read/write points: Writes via SaveStateAtomic(path, state); reads via LoadState(path). (internal/state/persist.go)

Per-step status tracking (StepStatus)

What it is: A per-step record with state machine values pending|running|success|failed|skipped, plus exit code, timestamps, and error string. (internal/state/types.go: type Status, constants; StepStatus{Status, ExitCode, StartedAt, EndedAt, Error})

Lifecycle (represented, not enforced here): A step typically transitions pending → running → (success|failed) or pending → skipped; timestamps support “running window” and duration computations. The code shown defines the model but not the transition logic. (internal/state/types.go)

Run warnings (non-fatal notes)

What it is: Structured warnings attached to run state, scoped and optionally tied to a specific step and line/token. (internal/state/types.go: Warning{Scope, StepID, Line, Token, Message}; RunState.Warnings)

Lifecycle / read/write: Persisted as part of RunState JSON alongside statuses. (internal/state/types.go; internal/state/persist.go)

Atomic persistence mechanism

What it does: Writes path + ".tmp" then renames to path, removing the temp file on rename failure. (internal/state/persist.go: tempPath := path + ".tmp", os.WriteFile, os.Rename, os.Remove)

Operational behavior: Output is pretty-printed JSON (json.MarshalIndent) which aids debugging but can be larger/slower than compact JSON. (internal/state/persist.go)

State “identity” / cache key fields

What it is: PackHash suggests the run state is intended to be tied to a specific pack content hash, likely to detect “pack changed vs. old state”. (internal/state/types.go: RunState.PackHash)

Lifecycle: The code provided stores the value but does not validate it during load; any reconciliation logic must be elsewhere. (internal/state/types.go; internal/state/persist.go)

Risks/unknowns (bullets)

Unknown where state is created/updated. The snippets define the data model and persistence, but not when StepStatuses are initialized, advanced, or saved (e.g., after each step, only at end, on error). That logic is likely in the executor/runner layer, not shown.

Missing “no state yet” behavior. LoadState returns an error when the state file does not exist ("state file not found"), which may force callers to treat “first run” as an error unless they special-case it. (internal/state/persist.go: os.IsNotExist(err))

Map key ambiguity. StepStatuses map[string]StepStatus does not guarantee ordering; the key format (step ID? label?) and collision rules are unknown without the parser/runner code that populates it. (internal/state/types.go)

Cross-platform atomicity. os.Rename(tempPath, path) may not overwrite an existing destination file on some p
```

docs/oracle-questions-2026-01-07/06-caching-state-consistency.md
```
1) Direct answer (top consistency risks + knobs)

Silent state-write failures → persisted state can lag in-memory indefinitely. saveState() discards errors (_ = state.SaveStateAtomic(...)), so the run can proceed while the on-disk state remains stale. Evidence: internal/app/run.go: saveState().

Non-transactional step transitions (RUNNING → SUCCESS/FAILED) can leave “stuck RUNNING” on disk. Each step does: set StatusRunning + saveState(), then later set EndedAt + final status + saveState(); any crash or write failure between these creates divergence (disk shows RUNNING; memory/report show final). Evidence: internal/app/run.go: RunPlain() loop setting StatusRunning then later StatusFailed/StatusSuccess.

Warnings can be persisted out of sync with step status. recordWarnings() drains and appends warnings and immediately calls saveState(); it runs after RunStep but before the final status save, so the state file can contain warnings from a step while still showing that step as RUNNING if the later status save fails. Evidence: internal/app/run.go: recordWarnings() + call sites around RunStep.

ROI-based skipping is not reflected in persisted state/report as “skipped”. ROI filter continues before writing any StepStatus; yet the state model has StatusSkipped available. This makes outputs inconsistent with console logs (“Skipping step…”) and loses auditability. Evidence: internal/app/run.go: ROIThreshold/ROIMode filter continues + internal/state/types.go: StatusSkipped exists.

StartTime is not guaranteed to flush. StartTime is set in-memory (if IsZero() { StartTime = time.Now() }) but there’s no immediate saveState(); if there are zero steps (and no warnings) you can generate a report with the updated StartTime while the persisted state file never updates. Evidence: internal/app/run.go: StartTime set; finalize() can run without any prior save.

Report write success is claimed even if the write fails. finalize() ignores json.MarshalIndent and os.WriteFile errors and prints “Report written to …” unconditionally, so report output can be missing/partial while stdout says it exists. Evidence: internal/app/run.go: finalize().

Resume behavior depends on persisted success markers and can misbehave with stale state. Resume skips only when StepStatuses[step.ID].Status == success; if state saves fail, already-completed steps can rerun (or previously-successful steps might not be recognized). Evidence: internal/app/run.go: resume check.

State behavior knobs/configs (observable here):

Config.StatePath: enables/disables persisted run state (saveState() is a no-op if empty). Evidence: internal/app/run.go: saveState().

Config.ReportPath: enables/disables report writing. Evidence: internal/app/run.go: finalize().

Config.StopOnFail: controls early termination path (still calls finalize() on failure). Evidence: internal/app/run.go: StopOnFail branch.

Config.ROIThreshold + Config.ROIMode: controls skip filter semantics (“under” is strict <; otherwise >=). Evidence: internal/app/run.go: ROI filter.

2) Risks/unknowns

Atomicity guarantees are unknown. state.SaveStateAtomic is assumed to be atomic, but wi
```

docs/oracle-questions-2026-01-07/07-background-jobs-discovery.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No explicit background jobs/workers/scheduled tasks are defined in internal/app/run.go. The execution model is a single synchronous control flow in (*App).RunPlain: optional prelude, then a for-loop over steps, each executed via a.Runner.RunStep(...), followed by state updates and writes. Evidence: RunPlain(...) calls RunPrelude(...) and RunStep(...) inline (no goroutines, tickers, or async constructs shown).

Triggers: Only direct, immediate calls from RunPlain (prelude runs if a.Pack.Prelude.Code != ""; steps run when they pass ROI/resume filters). Evidence: prelude conditional and for _, step := range a.Pack.Steps { ... err := a.Runner.RunStep(...) ... }.

Payloads: The “work” payload is the current step (&step) and out writer passed into Runner.RunStep, plus the prelude (&a.Pack.Prelude). Evidence: a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out) and a.Runner.RunStep(ctx, &step, out).

Retries: No retry loop/backoff exists here. The only “try again” behavior is manual restart/resume: successful steps are skipped on subsequent runs. Evidence: resume check if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess { ... continue }.

Idempotency: “At-least-once with resume-skip on success” at the step level. Steps are marked running, persisted, then marked success/failed and persisted; rerun skips StatusSuccess steps. Evidence: status set to StatusRunning + a.saveState(), then on success sets StatusSuccess + a.saveState(), and skip logic above.

Scheduled behavior: None (no timers/cron). The only time usage is stamping StartTime, StartedAt, EndedAt via time.Now(). Evidence: a.State.StartTime = time.Now(), StartedAt: time.Now(), status.EndedAt = time.Now().

2) Risks/unknowns

a.Runner.RunPrelude, a.Runner.RunStep, and a.Runner.DrainWarnings() could hide concurrency (goroutines, worker pools, background process supervision) inside internal/exec or other packages; this file alone cannot confirm that.

state.SaveStateAtomic(...) could include locking/FS semantics or side effects relevant to “background-like” behavior (e.g., fsync patterns), though it is still invoked synchronously here.

CLI/TUI entrypoints may spawn background refreshers (status polling, UI redraw loops) outside run.go.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background/scheduling primitives:
rg -n --hidden "time\.NewTicker|time\.AfterFunc|cron|scheduler|worker|goroutine|go\s+func|errgroup|WaitGroup|queue|async" internal/

4) If evidence is insufficient, missing file/path pattern(s) to attach next

internal/exec/**/*.go (especially internal/exec/runner.go and any process-management helpers)

internal/cli/**/*.go (entrypoints/TUI loops can introduce background refresh/polling)

internal/state/**/*.go (e.g., SaveStateAtomic implementation)
```

docs/oracle-questions-2026-01-07/08-background-jobs-reliability.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No dedicated background job system is evident in the Go code shown (no queue/worker process, scheduler, or job runner). The primary execution path appears to be a synchronous CLI (Cobra) that runs work inline rather than enqueuing background tasks. Evidence: CLI entry is Cobra-based with flags (internal/cli/root.go), and execution uses context + os/exec to run an external command (internal/exec/runner.go).

Because there’s no background job substrate, classic reliability controls (DLQ/dead-lettering, queue backoff policies, concurrency limits, reprocessing UI) are not implemented as first-class “background work” features—failures would surface as immediate CLI errors rather than being quarantined/retried by a worker. Evidence: inline execution model suggested by os/exec runner (internal/exec/runner.go) and lack of worker constructs in the CLI surface (internal/cli/root.go).

What does exist is “run reliability” via status/state tracking, which is the closest analogue to reprocessing controls: the system models step/run outcomes with explicit statuses (pending/running/success/failed/skipped). Evidence: Status* constants in internal/state/types.go.

State/report paths are configurable at the app level, which suggests persistence of outcomes to disk (useful for restart/manual re-run), but this is not the same as automated retries/backoff. Evidence: app-wide config includes StatePath and Report... fields (internal/app/app.go), and run orchestration imports internal/state + internal/report (internal/app/run.go snippet).

Concurrency limiting is implicitly “1 run = 1 process” (single CLI invocation) unless internal/app/run.go spawns goroutines (not shown). No explicit worker pool/limit controls are visible in the provided excerpts. Evidence: only synchronous CLI + runner excerpts provided (internal/cli/root.go, internal/exec/runner.go, partial internal/app/run.go).

2) Risks/unknowns (bullets)

internal/app/run.go is truncated; it could still spawn goroutines, use tickers, or implement internal retry/backoff that isn’t visible in the snippet.

The persistence semantics (when state is flushed, atomicity, resume behavior) can’t be confirmed without the full implementations of internal/state/* and internal/report/*.

There may be background work outside the shown files (e.g., separate cmd/* binaries, build tags, or a different package) that wasn’t included in the excerpts.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search for background-work primitives and libraries, e.g. rg -n "(time\\.NewTicker|time\\.Tick|cron|scheduler|queue|worker|goroutine|go\\s+func|asynq|machinery|temporal|celery|bull)" .

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/app/run.go (full file)

internal/state/*.go

internal/report/*.go

cmd/** and internal/cli/*.go (full set, not snippets)

Any directories named or resembling jobs/, worker/, background/, scheduler/, queue/, cron/ (if present)
```

docs/oracle-questions-2026-01-07/09-observability-signals.md
```
1) Direct answer (observability + correlation identifiers)

Primary structured observability output is the run “report”: internal/report/GenerateReport(s *state.RunState, packName string) *ReportV1 builds a ReportV1 containing PackInfo{Name, Hash}, GeneratedAt, per-step results (Steps: []StepReport), and a Summary with counts + total duration. (Evidence: internal/report/generate.go, GenerateReport populating PackInfo, Steps, Summary.)

Per-step signals captured (and thus observable via the report) include: StepReport.ID (step identifier), Status, ExitCode, Duration/DurationMs, and Error. These are derived from RunState.StepStatuses map[string]StepStatus and each StepStatus’s timestamps/exit code/error. (Evidence: internal/report/generate.go builds StepReport from the for id, status := range s.StepStatuses loop; internal/state/types.go defines RunState.StepStatuses and StepStatus fields.)

Warnings are first-class, structured observability signals: RunState.Warnings []Warning are copied into report.Warnings with fields {Scope, StepID, Line, Token, Message}. (Evidence: internal/state/types.go defines Warning; internal/report/generate.go copies s.Warnings into report.Warnings.)

Primary correlation identifiers across components (based on what’s present in state/report):

pack_hash: RunState.PackHash and ReportV1.PackInfo.Hash (ties a run to a specific pack content). (Evidence: internal/state/types.go PackHash; internal/report/generate.go assigns Hash: s.PackHash.)

step_id: the map key in RunState.StepStatuses and StepReport.ID; also optionally present on warnings (Warning.StepID). This is the main join key between step execution, step status, and warnings. (Evidence: internal/state/types.go StepStatuses map[string]StepStatus, Warning.StepID; internal/report/generate.go uses loop variable id → StepReport.ID.)

start_time / generated_at: RunState.StartTime and ReportV1.GeneratedAt provide coarse run-level temporal correlation (e.g., match logs/events by time window). (Evidence: internal/state/types.go StartTime; internal/report/generate.go sets GeneratedAt: time.Now().)

There is no explicit “run_id”/trace id in the provided state/report schema; correlation appears to rely on (pack_hash, start_time) + step IDs, rather than a stable UUID. (Evidence: internal/state/types.go RunState fields shown do not include a run identifier other than PackHash/StartTime.)

Log lines: no logging primitives (log.*, fmt.Printf, etc.) are shown in the provided generate.go or state/types.go, so log-based observability and log correlation keys cannot be confirmed from these files alone. (Evidence: the provided snippets contain only struct composition / time math.)

2) Risks / unknowns

Run correlation fragility: without a dedicated run_id, correlating across components (state file writes, executor output, CLI output, report artifacts) will be brittle when multiple runs share the same pack hash or occur close in time.

Non-deterministic step ordering in reports: iterating for id, status := range s.StepStatuses over a Go map yields randomized order, which can create noisy diffs and make “compare reports across runs” harder.

Missing visibility into logging: cannot verify
```

docs/oracle-questions-2026-01-07/10-observability-gaps.md
```
1) Direct answer (1–8 bullets, evidence-cited)

No correlation IDs in execution logs, so step output cannot be tied back to a step/run reliably. run() wires cmd.Stdout/cmd.Stderr directly to logWriter with no prefix/header/trailer, and it receives no stepID/runID to stamp lines (run(ctx, script, logWriter) only gets the script + writer). Evidence: RunPrelude/RunStep call r.run(ctx, script, logWriter); run() sets cmd.Stdout = logWriter and cmd.Stderr = logWriter and emits no logs of its own. (internal/exec/runner.go: RunPrelude, RunStep, run)

Key decision points are silent (missing “what config did we actually use?” logs). Shell defaulting (/bin/bash), override application, ChatGPTURL flag application, and flag injection happen without any structured record, making misconfig/debugging difficult. Evidence: NewRunner defaults Shell; RunStep conditionally calls Overrides.EffectiveFlags(...) and ApplyChatGPTURL(...), then InjectFlags(...), with no logging. (internal/exec/runner.go: NewRunner, RunStep)

Errors lose high-value context (step, command, exit code, duration). On failure, run() returns fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err) (or ctx.Err()), which preserves a category but omits step ID, elapsed time, and process exit code/signal—often the first things needed to diagnose. Evidence: run() error path and wrapping; error taxonomy only maps categories to exit codes. (internal/exec/runner.go: run; internal/errors/errors.go: ExitCode)

Sanitizer warnings are emitted as human text only and are not emitted at “decision boundaries.” Warnings are printed as a single formatted line and accumulated in-memory, but there’s no summary count per step/run, no machine-parseable form, and no “this step was sanitized” header at step start/end. Evidence: recordWarnings prints fmt.Fprintf(... "⚠ oraclepack: sanitized label ...") and appends to r.warnings; DrainWarnings() returns them but nothing here enforces they’re logged/serialized elsewhere. (internal/exec/runner.go: recordWarnings, DrainWarnings)

No run/step outcome metrics exist at the point where the outcome is known. The runner is the natural place to emit: step_success/step_failure counts, durations, and warning counts; currently it emits none. Evidence: run() measures no time and publishes no counters; it only runs the command and returns error/nil. (internal/exec/runner.go: run)

Smallest, highest-leverage additions (minimal surface area):

Add a tiny correlation context passed into run() (at least step_id, optionally run_id) and emit two boundary lines: step.start and step.end (include duration_ms, status, exit_code if available). This alone makes raw stdout/stderr interpretable. (Hook point: RunPrelude/RunStep just before/after r.run(...); data already present: s.ID in RunStep.) (internal/exec/runner.go: RunStep, run)

On error, detect *exec.ExitError and include exit_code/`s_*
```

docs/oracle-questions-2026-01-07/11-permissions-model.md
```
Direct answer (1–8 bullets, evidence-cited)

No first-class permission model is evident in the provided code: none of the shown entrypoints/config/state types include a notion of user identity, roles, scopes, ACLs, or authorization checks (e.g., the CLI root command only defines local flags like noTUI, oracleBin, outDir). Evidence: internal/cli/root.go (flag surface only; no auth/identity inputs shown).

Execution appears to be “whoever can run the CLI can do everything the CLI does.” The runner shells out to an external binary and interacts with the local filesystem; no gating by role/scope is shown. Evidence: internal/exec/runner.go (uses os/exec + os without any authz hooks shown).

“Who can do what” is therefore effectively determined by OS-level permissions and environment constraints, not by application-level RBAC/ACL logic. Evidence: internal/app/app.go (config is paths like PackPath/StatePath/Report…; no policy/role configuration shown) + internal/app/run.go (reads/writes JSON/state/report to local paths; no authz shown).

Any real access control is likely delegated to external dependencies (e.g., the invoked oracle binary and filesystem permissions), not implemented in oraclepack itself, based on the visible integration boundary. Evidence: internal/exec/runner.go (external command invocation) + internal/report/generate.go (report generation from run state; no authz metadata).

Minimal set of files that encode “who can do what” (based on what’s been provided)

None (no app-level permission model is shown in the provided files). Practically, behavior is governed by:

OS user permissions (read/write/execute) for the invoked binary and directories (PackPath, StatePath, report/output dirs). Evidence: internal/app/app.go, internal/app/run.go.

External tool behavior (the oracle binary). Evidence: internal/exec/runner.go.

Risks/unknowns (bullets)

The repository may contain an unshared permissions/auth subsystem (e.g., internal/auth/*, HTTP server middleware, policy engine) that wasn’t included in the snippets.

The external oracle binary may implement its own authn/authz, which would materially change the true permission model even if oraclepack itself has none.

There may be implicit “permissions” via overrides/config resolution (e.g., internal/overrides) that constrain execution; not enough code was provided to verify.

Next smallest concrete experiment (exactly one action)

Run a repo-wide search for authorization primitives: rg -n "(rbac|role|roles|scope|scopes|acl|permission|authorize|authn|authz|oauth|jwt|policy|casbin)" .

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

internal/**/auth*.go

internal/**/perm*.go

internal/**/{rbac,acl,policy}*.go

internal/**/middleware*.go

internal/**/http*/**/*.go and/or internal/**/server*/**/*.go

cmd/** (any additional CLIs/entrypoints beyond internal/cli/root.go)

Any config/schema files that might define roles/scopes (e.g., **/*config*.*, **/*schema*.*, **/*policy*.*)
```

docs/oracle-questions-2026-01-07/12-permissions-enforcement.md
```
1) Direct answer (1–8 bullets, evidence-cited)

There is no explicit permission enforcement layer in the surfaced code (no authn/authz context, no roles/scopes/ACL objects, no guards/middleware imports). The CLI entrypoint is Cobra-based with flag globals (noTUI, oracleBin, outDir) and does not introduce any identity/permission concept. Evidence: internal/cli/root.go snippet shows only Cobra + basic flags; no auth-related imports/types.

Effective “permissions” are the OS user’s privileges at runtime. The system executes external commands via os/exec, so whatever the current process user can do (filesystem, network, process execution) the tool can do. Evidence: internal/exec/runner.go imports and uses os/exec and os.

Primary enforcement chokepoint is the execution boundary (Runner) — but it enforces execution, not authorization. If you need “who can run what,” the only practical choke is where commands are spawned (and where env/args are assembled). Evidence: internal/exec/runner.go is the dedicated Runner module and pulls overrides.

Pack parsing enables embedding runnable shell content; there is no shown policy layer restricting it. The parser extracts bash fenced blocks (regex), which strongly implies step payloads can contain shell scripts. **Evidence:** `internal/pack/parser.go` defines `bashFenceRegex = ... bash\n(.*?)\n``` ...`.

State/report generation has no permission gates; it simply materializes outputs. So any caller who can run the CLI can cause state/report writes wherever configured. Evidence: internal/report/generate.go constructs a report from RunState; internal/app/app.go config includes StatePath/Report… fields (from snippet) with no auth checks.

Likely bypass risks are “configuration-as-authority” rather than missing middleware. If a caller can control oracleBin, pack content, overrides, output/state paths, they effectively control what runs and where it writes—without an authorization decision point. Evidence: internal/cli/root.go exposes oracleBin/outDir; internal/exec/runner.go consumes overrides + executes commands; internal/pack/parser.go extracts bash payloads.

Inconsistent pattern: business logic lives in internal/app/*, but enforcement is absent at service-layer boundaries. There’s no “service-layer check” because there’s no identity/claims passed through app -> exec -> filesystem; everything is implicitly trusted input. Evidence: internal/app/app.go wires exec, pack, state with no auth context types.

2) Risks/unknowns (bullets)

Unknown whether any permission checks exist elsewhere (e.g., an HTTP server, MCP server wrapper, or API layer not included in the snippets). If those exist, they may be the real enforcement point.

Unknown how overrides are applied (env var allow/deny lists, argument escaping, path validation). internal/overrides/* was referenced but not provided.

Unknown whether pack execution is sandboxed (chroot, containers, seccomp, allowlisted commands). Nothing in the provided snippets suggests sandboxing, but absence of evidence isn’t proof.

Unknown whether state/output paths are validated (path traversal, writing outside workspace, overwriting). Config shows paths, but validation code wasn’t included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search to locate any authz/authn/permission enforcement and all chokepoints where requests/commands enter:

bash
Copy code
rg -n --hidden --no-ignore-vcs "(auth|authz|authorize|authorization|permission|rbac|acl|role|scope|policy|guard|middleware|claims|principal|subject|tenant)" .

4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

Attach any of the following if they exist (these are the most likely places where enforcement would live, and we haven’
```

docs/oracle-questions-2026-01-07/13-migrations-schema.md
```
1) Direct answer (1–8 bullets, evidence-cited)

State schema has an explicit version field but no migration logic is implemented. RunState includes SchemaVersion int \json:"schema_version"`ininternal/state/types.go, but internal/state/persist.go::LoadStatesimplyjson.Unmarshals into RunStateand returns it without checking or upgradingSchemaVersion`.

Evidence: internal/state/types.go (RunState.SchemaVersion) and internal/state/persist.go (LoadState unmarshal + return; no branching on version).

State evolution currently relies on Go JSON’s “additive compatibility” behavior. New optional fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty in internal/state/types.go, meaning:

older state files missing these fields will still load (they default to zero values),

newer state files with extra fields would be ignored by older binaries (Go encoding/json ignores unknown fields by default).

Evidence: internal/state/types.go JSON tags on ROIThreshold, ROIMode, Warnings.

There is no tooling or “migration runner” in the shown code paths. The persistence layer only provides atomic write + read (SaveStateAtomic, LoadState) and does not do version stamping, backfills, or transforms.

Evidence: internal/state/persist.go contains only marshal/write/rename and read/unmarshal.

Practical implication: only additive changes are safe; renames/removals are breaking. Because no upgrade step exists, any future schema changes that rename/remove fields, change meaning, or change requiredness will either silently degrade (defaults) or break logic downstream.

Where migrations would be applied (natural chokepoint): internal/state/persist.go::LoadState (or the caller immediately after it) is the obvious place to:

read → unmarshal

check SchemaVersion against a CurrentSchemaVersion constant

run sequential migrations migrateV1toV2, migrateV2toV3, etc.

set SchemaVersion = CurrentSchemaVersion

optionally rewrite the upgraded file via SaveStateAtomic

Evidence for chokepoint existence: LoadState/SaveStateAtomic are the only provided persistence entry points in internal/state/persist.go.

Report format evolution cannot be confirmed from the provided snippets. Earlier context shows GenerateReport returns a *ReportV1, which suggests intent to version report structures by type name, but there’s no evidence (in the snippets here) of report persistence, report version fields, or report migrations.

Evidence: internal/report/generate.go shows GenerateReport(...) *ReportV1 (from prior snippet), but no internal/report/*.go definitions/persistence were provided here.

2) Risks/unknowns

Unknown whether SchemaVersion is ever set/non-zero (e.g., on initial state creation), and what the “current” version is meant to be (no constant or assignment shown).

Unknown whether there is any hidden migration handling in the caller (likely internal/app/run.go or similar) before/after LoadState.

Unknown report persistence/read path (where report files are written/read, whether a version field exists, whether multiple report versions are supported).

Unknown whether pack/config/overrides have versioning (separate from state schema; not shown in the provided files).

3) Next smallest concrete experiment (exactly one action)

Run and paste the output of:
rg -n "SchemaVersion|schema_version|CurrentSchema|migrat(e|ion)|ReportV[0-9]|report_version" internal/

4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

internal/report/*.go (especially the file defining ReportV1 and any report write/read functions)

internal/app/run.go (or whichever file calls state.LoadState / state.SaveStateAtomic, and whi
```

docs/oracle-questions-2026-01-07/14-migrations-compat.md
```
1) Direct answer (backward/forward compatibility expectations; where ensured vs risky)

Run state is explicitly versioned; expectation is “migrate-by-version when needed.” RunState carries SchemaVersion int \json:"schema_version"`which is the primary hook for backward/forward compatibility across persisted state formats. Evidence:internal/state/types.go (RunState.SchemaVersion`).

State schema is designed to be additive (backward-compatible) via optional fields. Fields like ROIThreshold, ROIMode, and Warnings are tagged omitempty, implying they can be absent in older saved state and should default safely in newer code. Evidence: internal/state/types.go (ROIThreshold ...omitempty, ROIMode ...omitempty, Warnings ...omitempty).

Forward compatibility for state/report within Go likely relies on default encoding/json behavior (unknown fields ignored). The structs are plain JSON-tagged types with no visible strict decoding controls here; in Go, older binaries typically ignore new fields when unmarshalling (unless DisallowUnknownFields is used elsewhere). Evidence: presence of standard JSON tags across structs in internal/state/types.go and internal/report/types.go (no strictness shown).

Report schema is implicitly versioned by type name, not by an on-the-wire version field (riskier for evolvability). The report root type is ReportV1, but there is no schema_version/report_version JSON field on ReportV1. This suggests compatibility depends on file naming/conventions or external knowledge, not self-describing payloads. Evidence: internal/report/types.go (type ReportV1 struct { ... } with no version field).

Report is partially “consumer-friendly” via redundant duration fields, suggesting an evolution/compat strategy. Summary.TotalDuration (a time.Duration) is paired with TotalDurationMs int64, and StepReport.Duration is paired with DurationMs int64. This looks like an intentional compatibility measure for consumers that can’t/won’t parse Go time.Duration encoding. Evidence: internal/report/types.go (TotalDuration + TotalDurationMs, Duration + DurationMs).

Two distinct Warning types exist (state vs report), duplicating schema; compatibility depends on keeping them in lockstep. Both packages define Warning with the same JSON fields; this is currently compatible, but any future divergence becomes a schema drift risk. Evidence: internal/state/types.go type Warning struct { Scope, StepID, Line, Token, Message } and internal/report/types.go type Warning struct { ...same fields... }.

Map-based step status storage is tolerant of step additions but sensitive to step ID churn. RunState.StepStatuses map[string]StepStatus supports new step IDs appearing over time, but renames/removals can break continuity of resumed/inspected runs unless migration logic remaps keys. Evidence: internal/state/types.go (StepStatuses map[string]StepStatus).

2) Risks/unknowns

Unknown whether reads/writes are strict or tolerant. If any loader uses json.Decoder.DisallowUnknownFields(), forward compatibility (new fields read by old binaries) will break. Missing evidence: state/report load code.

No visible migration dispatcher. SchemaVersion exists, but we don’t see code that applies migrations on read (or stamps versions on write), so rollback/upgrade safety is unproven from the provided files alone.

Report payload is not self-describing. Without a report_version, consumers can’t robustly branch on schema, making forward/backward compatibility brittle as ReportV2 emerges.

G
```

docs/oracle-questions-2026-01-07/15-ux-flows-primary.md
```
1) Direct answer (primary TUI user/operator flows, mapped to components + key state transitions)

App entry → TUI boot (CLI “run” command): internal/cli/run.go loads/initializes app (pack, runner, state), then constructs tui.Model via tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll) and runs Bubble Tea (tea.NewProgram(..., tea.WithAltScreen())). The --run-all flag is passed as autoRun into the TUI model. (internal/cli/run.go: runCmd.RunE; internal/tui/tui.go: NewModel, Init)

Browse/select steps (default “Steps” screen): Primary UI is list.Model (left) + viewport.Model (right preview). Selection changes update preview content via setListPreviewContent(newID) when not running and viewState == ViewSteps. (internal/tui/tui.go: Model.View, Update list handling, setListPreviewContent, selectedItemID)

Run a single selected step (“Enter”): From ViewSteps, pressing Enter sets running=true, viewState=ViewRunning, runAll=false, clears logLines, and triggers runStep(stepID) + waitForLogs() + spinner.Tick. Completion arrives as FinishedMsg{Err,...}; success sets running=false, viewState=ViewDone; failure sets err, running=false, runAll=false, viewState=ViewDone. (internal/tui/tui.go: Update key "enter", runStep, waitForLogs, FinishedMsg handler)

Run all filtered steps sequentially (“a”) + auto-run on init: Pressing a in ViewSteps (when not running) sets running=true, viewState=ViewRunning, runAll=true, currentIdx=0, selects first filtered list item, then executes runStep(firstID). On each FinishedMsg success while runAll=true, it increments currentIdx, selects next item, and calls runStep(nextID) until exhausted, then sets running=false, runAll=false, viewState=ViewDone. Auto-run uses the same path via StartAutoRunMsg emitted from Init() when autoRun is true. (internal/tui/tui.go: Update key "a", StartAutoRunMsg handling, FinishedMsg sequential logic; Init)

ROI filtering flow (“f” and “m”):

f enters filter input modal: isFiltering=true, filterInput.Focus(), view switches to centered modal via View() (still viewState remains Steps, but UI is driven by isFiltering). Enter parses float, updates roiThreshold, calls refreshList(), and persists into state (m.state.ROIThreshold, m.state.ROIMode) via state.SaveStateAtomic. Esc cancels.

m toggles roiMode between "under" and "over", refreshes list, persists state.
This directly changes what “run all” iterates over because it uses m.list.Items() (filtered). (internal/tui/tui.go: Update keys "f", "m", refreshList, persistFilterState, View modal branches)

Step preview (“v”) with wrap/copy/back: v sets previewID, viewState=ViewStepPreview, resets viewport offset, and renders Markdown (render.RenderMarkdown) into the viewport. In preview: t toggles wrap (previewWrap) and re-renders; c copies step plain text (clipboard fallback writes to file); b/esc returns to ViewSteps and restores list preview. (internal/tui/tui.go: Update key "v", `ViewSte
```

docs/oracle-questions-2026-01-07/16-ux-flows-edgecases.md
```
1) Direct answer (top edge cases / gotchas, where handled vs missing)

Cancel is “hard cancel” everywhere, with no confirm and no “go back one step” on Esc: OverridesFlowModel.Update maps "esc" directly to OverridesCancelledMsg{} regardless of wizard step; there’s no “Are you sure?” or “Back” on Esc. (Handled: cancel exit path exists; Missing: confirm/discard + step-back on Esc.) (OverridesFlowModel.Update: case "esc": ... OverridesCancelledMsg{})

Back navigation conflicts with normal text editing / list filtering: the parent model binds "backspace" (and "shift+tab") to step-back (m.step--) whenever m.step > OverridesFlags. If StepsPickerModel (and/or confirm UI) uses inputs or list filtering, Backspace is commonly needed for editing—this will unexpectedly navigate back instead. (Missing: scoped key handling / focus-aware routing.) (OverridesFlowModel.Update: case "shift+tab", "backspace": if m.step > OverridesFlags { m.step-- })

Next navigation conflicts with child controls: the parent captures "enter" and "tab" to advance steps (or trigger validation on confirm). If the list widgets use Enter/Tab for selection, filtering, or focus, those interactions can be overridden by step progression. (Missing: “only advance when not filtering / not editing” guard, or explicit Next button/key distinct from list controls.) (OverridesFlowModel.Update: case "enter", "tab": ... m.step++)

Validation failures are displayed, but recovery UX is minimal: confirm step handles ValidationResultMsg by setting confirm.validating=false, storing confirm.errors, and setting a short errMsg (“N validation errors detected.”). It does not automatically move the user back to the step that likely caused the errors, nor does it highlight which selection caused which validation error (depends on OverridesConfirmModel.View, not shown). (Handled: error surfaced; Missing: guided remediation / jump-back.) (OverridesFlowModel.Update: case ValidationResultMsg: ... m.confirm.errors = v.Errors ... errMsg ... return m, nil)

Validation can hang indefinitely (no timeout/cancel), making the UI feel frozen: validateCmd uses context.Background() and calls exec.ValidateOverrides(...) with no deadline and no cancel path. While confirm.validating prevents duplicate submissions, there’s no way to abort a stuck validation. (Missing: context timeout + Esc to cancel validation.) (validateCmd: exec.ValidateOverrides(context.Background(), ...))

Baseline flags appear incorrectly wired into the flags picker (likely a functional bug): NewOverridesFlowModel constructs flags: NewFlagsPickerModel(nil) even though it has a baseline argument and separately stores baseline: exec.ApplyChatGPTURL(baseline, ...). This means flags that are already in baseline won’t be marked (base) / [*], won’t auto-check, and won’t be protected from toggling. (Missing: pass the baseline into NewFlagsPickerModel(baseline).) (NewOverridesFlowModel: NewFlagsPickerModel(nil) vs stored baseline: ...)

Users cannot remove baseline flags (only add): currentOverrides() sets RemovedFlags: nil always, and FlagsPickerModel.SelectedFlags() only returns selected flags where !fi.IsBaseline. Even if you wanted removal, there’s no path in the
```

docs/oracle-questions-2026-01-07/17-failure-modes-taxonomy.md
```
1) Direct answer (failure-mode taxonomy; where classified/handled; what’s surfaced)

Precondition / lifecycle failures (internal “programmer error” / wrong call order): RunPlain hard-fails if a.Pack == nil or a.State == nil with plain fmt.Errorf("pack not loaded") / "state not loaded". These are not classified into internal/errors sentinel types, so callers will likely treat them as generic failures. (Evidence: internal/app/run.go:RunPlain)

Prelude execution failures (treated as fatal): If Pack.Prelude.Code != "", it runs a.Runner.RunPrelude(...); on error it records warnings, then returns fmt.Errorf("prelude failed: %w", err) without writing a report (no finalize() call on this path). (Evidence: internal/app/run.go:RunPlain)

Step execution failures (classified at the step-status layer, optionally fatal at the run layer):

Before executing a step, state is set to StatusRunning with StartedAt, then persisted via a.saveState().

On RunStep error: step status becomes StatusFailed, status.Error = err.Error(), state is persisted, and:

if StopOnFail is true: finalize() runs and RunPlain returns the underlying err;

otherwise it continues to the next step.
This is the core “execution failed” mode in practice. (Evidence: internal/app/run.go:RunPlain)

“Stop-on-fail” semantics create two user-visible failure shapes: with StopOnFail=true, the run returns immediately on the first failed step (after writing the report); with StopOnFail=false, the run returns nil even if one or more steps failed, relying on state/report for operators to discover failures. (Evidence: internal/app/run.go:RunPlain)

Skips are treated as non-fail outcomes, but are not consistently represented in state:

ROI-based skips print Skipping step... and continue without writing any step status (no StatusSkipped assignment here).

Resume-based skips require an existing StatusSuccess entry; they print and continue leaving the prior success intact.
Result: reports/operators may not be able to distinguish “intentionally skipped by ROI filter” vs “never attempted / missing status” unless report generation infers it elsewhere. (Evidence: internal/app/run.go:RunPlain)

Warnings are a parallel channel, not failures: recordWarnings() drains runner warnings and appends them to a.State.Warnings, then persists state. They are not printed directly by App (beyond whatever the runner prints), so visibility is primarily via persisted state/report. (Evidence: internal/app/run.go:recordWarnings, saveState)

State persistence failures are explicitly ignored (silent data-loss risk): saveState() discards errors from state.SaveStateAtomic(...) via _ = .... If state writes fail, the run proceeds and operators may later see stale/missing progress. (Evidence: internal/app/run.go:saveState)

Report generation / write failures are explicitly ignored, but success is always announced: finalize() discards errors from json.MarshalIndent(...) and os.WriteFile(...), yet always prints Report written to ... when ReportPath != "". This can falsely signal success to operators even when no report was produced. (Evidence: internal/app/run.go:finalize)

System-wide error classification exists, but RunPlain largely relies on upstream to use it: internal/errors/errors.go defines sentinel errors (ErrInvalidPack, ErrExecutionFailed, ErrConfigInvalid) and an ExitCode(err) mapping using `errors.Is(
```

docs/oracle-questions-2026-01-07/18-failure-modes-resilience.md
```
1) Direct answer (resilience mechanisms; gaps; where configured)

Script “label line” sanitization exists and is applied to both preludes and steps. SanitizeScript converts single-token, label-shaped lines into echo "<token>" to avoid accidental command execution (e.g., a bare Setup line). This is invoked from Runner.RunPrelude and Runner.RunStep before execution. (internal/exec/sanitize.go: SanitizeScript; internal/exec/runner.go: RunPrelude, RunStep)

Sanitization is deliberately scoped to reduce false positives. It skips blank lines and comments (#...), skips anything inside heredocs (tracks start via heredocStartToken and ignores lines until the terminator), and only rewrites lines where strings.Fields(trimmed) yields exactly 1 field. (internal/exec/sanitize.go: SanitizeScript, heredocStartToken)

Sanitization avoids rewriting real shell constructs and real executables. Tokens are rewritten only if they match labelTokenRegex (^[A-Za-z][A-Za-z0-9_-]*$), are not shell builtins/keywords, and are not found on $PATH (osexec.LookPath(token) must fail). (internal/exec/sanitize.go: labelTokenRegex, shellBuiltins, shellKeywords, LookPath check)

Sanitization produces explicit warnings and they are surfaced at least to the run log stream. Warnings are stored on the runner and also written to logWriter as ⚠ oraclepack: sanitized label ..., and can be retrieved via DrainWarnings(). (internal/exec/runner.go: recordWarnings, DrainWarnings)

Execution is cancellation-aware via context. exec.CommandContext(ctx, ...) is used; on failure, if ctx.Err() != nil it returns the context error directly. This is a resilience mechanism for user cancel / deadline propagation. (internal/exec/runner.go: run)

Execution errors are wrapped with a sentinel classification. Non-context execution errors return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err), enabling upstream code to recognize “execution failed” as a class. (internal/exec/runner.go: run)

Configurable execution knobs exist, but they’re mostly “wiring,” not resilience policies. RunnerOptions configures Shell (default /bin/bash), WorkDir, Env (appended to os.Environ()), OracleFlags, Overrides, and ChatGPTURL. These influence behavior but do not add retries/limits/strict-mode by themselves. (internal/exec/runner.go: RunnerOptions, NewRunner)

Critical path gap: “stop-on-fail” within a script is not enforced. The runner executes bash -lc <script> without forcing set -e / -euo pipefail. So a failing command earlier in a script may not fail the step unless the script author ensured it (depends on the script). This is a major resilience gap for “partial failure looks like success.” (internal/exec/runner.go: run uses r.Shell, "-lc", script)

Critical path gap: sanitization misses common “label + trailing comment” patterns. Because only len(fields) == 1 is sanitized, a line like Setup # do thing becomes two fields and is not rewritten; bash will attempt to execute Setup and fail (or worse, execute if Setup exists on PATH). (internal/exec/sanitize.go: Fields/len(fields) != 1 early-continue)

Critical path gap: environment
```

docs/oracle-questions-2026-01-07/19-feature-flags-inventory.md
```
Direct answer (inventory + where defined/evaluated/documented)

Runtime “flag override” mechanism exists via RuntimeOverrides: you can append flags, remove flags, and inject a ChatGPT URL (as --chatgpt-url) at runtime. Defined in internal/overrides/types.go as AddedFlags, RemovedFlags, and ChatGPTURL. (internal/overrides/types.go: RuntimeOverrides)

Per-step targeting exists via ApplyToSteps: a set keyed by step ID that determines which steps the overrides apply to. Defined in internal/overrides/types.go as ApplyToSteps map[string]bool with comment indicating “If empty, applies to none.” (internal/overrides/types.go: RuntimeOverrides.ApplyToSteps)

TUI exposes a curated list of toggleable “Oracle Flags” (feature-flag-like user controls):
--files-report, --render, --render-plain, --copy, --wait. Defined in internal/tui/overrides_flags.go in curated := []FlagItem{...}. (internal/tui/overrides_flags.go: curated)

Baseline vs. user-selected behavior is enforced: baseline flags are pre-selected and cannot be toggled off; only non-baseline flags can be toggled with spacebar. Defined/evaluated in NewFlagsPickerModel (baseline set, IsBaseline, initial Selected) and Update (toggle only if !item.IsBaseline). (internal/tui/overrides_flags.go: NewFlagsPickerModel, Update)

UI-level “documentation” exists for these flags via the Desc field displayed alongside each curated flag (e.g., “Show per-file token usage”, “Print assembled markdown bundle”). Implemented in flagsDelegate.Render. (internal/tui/overrides_flags.go: FlagItem.Desc, Render)

Discoverability control exists in the TUI via filtering: the list enables filtering (l.SetFilteringEnabled(true)), which is a UX control for narrowing/selecting flags (not ROI, but “runtime control”). (internal/tui/overrides_flags.go: SetFilteringEnabled(true))

Selection output is explicitly “delta-only”: SelectedFlags() returns only flags that are selected and not baseline, i.e., the override payload is “changes from baseline,” which is consistent with safe rollout semantics. (internal/tui/overrides_flags.go: SelectedFlags)

Risks/unknowns

Where overrides are actually applied/evaluated is not shown in the provided snippets (i.e., where AddedFlags/RemovedFlags/ChatGPTURL/ApplyToSteps mutate the final oracle CLI invocation per step).

ROI filtering is not evidenced in the provided files; if it exists, it’s elsewhere (selection of steps by ROI, filtering packs/steps, etc.).

Documentation scope is currently “in-TUI only” for curated flags; no evidence here of CLI --help docs, README references, or pack schema docs for these controls.

Next smallest concrete experiment (exactly one action)

Run: rg -n "RuntimeOverrides|AddedFlags|RemovedFlags|ChatGPTURL|ApplyToSteps|--chatgpt-url|ROI|roi" internal/ and attach the matching snippets (especially any code that builds the exec command / selects steps).

If evidence is insufficient, attach these missing file/path pattern(s) next

internal/overrides/*.go (anything beyond types.go, especially functions that apply overrides)

The portion(s) of internal/exec/runner.go (or related exec builder) that incorporate overrides into the command/args

Any step-selection or filtering logic files, likely under: internal/app/*.go, internal/pack/*.go, internal/tui/*flow*.go (for ROI or step filtering)
```

docs/oracle-questions-2026-01-07/20-feature-flags-rollout.md
```
## 1) Direct answer (flag/rollout lifecycle, evidence-cited)

* **Selection is an explicit, user-supplied “diff” over baseline flags**: runtime overrides are modeled as `AddedFlags`, `RemovedFlags`, and an optional `ChatGPTURL` injection (`--chatgpt-url <url>`). Evidence: `internal/overrides/types.go` (`RuntimeOverrides` fields) and `internal/overrides/merge.go` (injection logic in `EffectiveFlags`).
* **Targeting is allowlist-by-step and defaults to “off”**: overrides apply only when `ApplyToSteps[stepID] == true`; if `ApplyToSteps` is `nil` (or the step key is absent/false), the baseline flags are returned unchanged. Evidence: `internal/overrides/merge.go` early-return guard in `EffectiveFlags`; `internal/overrides/types.go` comment “If empty, applies to none.”
* **Application happens at flag-materialization time, per step**: the lifecycle is “start from baseline → remove exact-match flags → append added flags → optionally append `--chatgpt-url <ChatGPTURL>`”. Evidence: `internal/overrides/merge.go` (`removed` map filter, then `append(r.AddedFlags...)`, then URL injection).
* **Validation is not implemented in the shown override module**: there is no validation method or error path in `RuntimeOverrides` or `EffectiveFlags`; invalid flags, conflicting flags, malformed URLs, or nonexistent step IDs are not rejected here. Evidence: `internal/overrides/types.go` contains only fields; `internal/overrides/merge.go` is pure transformation with no errors/validation.
* **Retirement is purely process-driven with current code**: there is no TTL/expiry, metadata, audit, or “unused override” detection in the override types/merge logic. Evidence: `internal/overrides/types.go` has no lifecycle metadata; `internal/overrides/merge.go` has no tracking/telemetry hooks.

**Minimum guardrails to prevent “flag debt” (smallest set that actually closes the main failure modes):**

* **Central validation gate** before any run: validate (a) every `ApplyToSteps` key exists in the pack’s step IDs, (b) flags conform to an allowlisted pattern (e.g., `^--[a-z0-9-]+(=.*)?$`), (c) no duplicates/conflicts after merge (e.g., two `--model` variants), and (d) `ChatGPTURL` is a valid URL. (Missing today in `types.go`/`merge.go`.)
* **Add lifecycle metadata** to each override bundle: `owner`, `reason`, `created_at`, `expires_at` (or `ttl_days`), and `issue/link`. Enforce “expired overrides fail closed.” (Not present today in `RuntimeOverrides`.)
* **Auditability**: emit a structured record per step of `{baseline_flags, removed, added, url_injected, effective_flags}` so operators can see what actually happened and spot stale overrides. (Nothing in `EffectiveFlags` records decisions.)
* **Retirement enforcement**: CI/test-time lints that fail if (a) overrides have expired, (b) an override hasn’t been applied to any step in N runs, or (c) overrides reference removed steps/flags. (No mechanism exists in the shown code; must be added around it.)
* **Prefer semantic removal** (optional but high leverage): today removal is exact-string match (`removed[flag]`), which doesn’t catch `--flag=value` vs `--flag` variants and encourages drift/debt. Either normalize flags or support key-based removal. Evidence of exact-match removal: `internal/overrides/merge.go` baseline filter uses `removed[flag]` only.

## 2) Risks/unknowns

* **Where overrides are created/loaded/persisted is not shown**, so it’s unclear whether any validation exists upstream (TUI wizard, config loader, etc.) or whether overrides can be injected via environment/CLI.
* **No evidence of observability/logging around overrides**: if something breaks due to an override, operators may not be able to quickly attribute the effective flag set to the override bundle.
* **Exact-match removal is brittle** and can lead to long-lived “ineffective” `RemovedFlags` entries (flag debt), because they silently don’t remove anything when baseline formatting changes. Evidence: `internal/overrides/merge.go`.
* **No conflict resolution**: if baseline already includes a flag that `AddedFlags` adds again, the effective list can contain duplicates, and downstream flag parsing semantics may become ambiguous. Evidence: `internal/overrides/merge.go` simply appends.

## 3) Next smallest concrete experiment (exactly one action)

* **Add `(*RuntimeOverrides) Validate(stepIDs []string) error` and call it once before execution begins**, rejecting unknown step IDs, malformed flags, duplicates/conflicts, and invalid `ChatGPTURL`; include a unit test table for the edge cases above.

## 4) If evidence is insufficient, exact missing file/path pattern(s) to attach next

* `internal/tui/overrides_flow.go` (how users select/serialize overrides)
* `internal/exec/runner.go` (where `EffectiveFlags` is invoked and how effective flags are executed/logged)
* Any override config IO paths, e.g. `internal/overrides/*.go` beyond `types.go`/`merge.go`, and any config structs/files such as `internal/app/*.go`, `internal/pack/*.go`, or `internal/config/*.go` that load/store runtime overrides
```

internal/app/app.go
```
package app

import (
	"fmt"
	"os"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

// Config holds application-wide configuration.
type Config struct {
	PackPath     string
	StatePath    string
	ReportPath   string
	StopOnFail   bool
	Resume       bool
	Verbose      bool
	DryRun       bool
	OracleFlags  []string
	WorkDir      string
	OutDir       string // CLI override for output directory
	ROIThreshold float64
	ROIMode      string // "over" or "under"
}

// App orchestrates the execution flow.
type App struct {
	Config Config
	Pack   *pack.Pack
	State  *state.RunState
	Runner *exec.Runner
}

// New creates a new application instance.
func New(cfg Config) *App {
	return &App{
		Config: cfg,
		Runner: exec.NewRunner(exec.RunnerOptions{
			WorkDir:     cfg.WorkDir,
			OracleFlags: cfg.OracleFlags,
		}),
	}
}

// LoadPack loads and validates the pack.
func (a *App) LoadPack() error {
	data, err := os.ReadFile(a.Config.PackPath)
	if err != nil {
		return err
	}

	p, err := pack.Parse(data)
	if err != nil {
		return err
	}

	if err := p.Validate(); err != nil {
		return err
	}

	a.Pack = p
	a.Pack.Source = a.Config.PackPath
	return nil
}

// LoadState loads or initializes the state.
func (a *App) LoadState() error {
	if a.Config.Resume {
		s, err := state.LoadState(a.Config.StatePath)
		if err == nil {
			a.State = s
			return nil
		}
	}

	a.State = &state.RunState{
		SchemaVersion: 1,
		StepStatuses:  make(map[string]state.StepStatus),
	}
	return nil
}

// Prepare resolves configuration and prepares the runtime environment.
func (a *App) Prepare() error {
	if a.Pack == nil {
		if err := a.LoadPack(); err != nil {
			return err
		}
	}

	// Resolve Output Directory
	// Precedence: CLI > Pack > Default (.)
	outDir := a.Config.OutDir
	if outDir == "" && a.Pack.OutDir != "" {
		outDir = a.Pack.OutDir
	}
	if outDir == "" {
		outDir = "."
	}

	// Provision Directory
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outDir, err)
	}

	// Update Runner
	// We do NOT set WorkDir to outDir, so execution happens in the project root.
	// This preserves relative path resolution for -f flags.
	// a.Runner.WorkDir = outDir 
	
	// Add out_dir to Env so scripts can reference it
	a.Runner.Env = append(a.Runner.Env, fmt.Sprintf("out_dir=%s", outDir))

	return nil
}
```

internal/app/app_test.go
```
package app

import (
	"bytes"
	"context"
	"os"
	"testing"
)

func TestApp_RunPlain(t *testing.T) {
	packContent := `
# Test Pack
` + "```" + `bash
# 01)
echo "step 1"
# 02)
echo "step 2"
` + "```" + `
`
	packFile := "test.md"
	stateFile := "test_state.json"
	reportFile := "test_report.json"
	defer os.Remove(packFile)
	defer os.Remove(stateFile)
	defer os.Remove(reportFile)

	os.WriteFile(packFile, []byte(packContent), 0644)

	cfg := Config{
		PackPath:   packFile,
		StatePath:  stateFile,
		ReportPath: reportFile,
	}

	a := New(cfg)
	if err := a.Prepare(); err != nil {
		t.Fatalf("Prepare failed: %v", err)
	}
	if err := a.LoadState(); err != nil {
		t.Fatalf("LoadState failed: %v", err)
	}
	
	var out bytes.Buffer
	err := a.RunPlain(context.Background(), &out)
	if err != nil {
		t.Fatalf("RunPlain failed: %v", err)
	}

	output := out.String()
	if !contains(output, "step 1") || !contains(output, "step 2") {
		t.Errorf("output missing steps: %s", output)
	}

	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		t.Error("state file was not created")
	}

	if _, err := os.Stat(reportFile); os.IsNotExist(err) {
		t.Error("report file was not created")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}
```

internal/app/run.go
```
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/user/oraclepack/internal/report"
	"github.com/user/oraclepack/internal/state"
)

func (a *App) RunPlain(ctx context.Context, out io.Writer) error {
	// Assumes a.Prepare() and a.LoadState() have been called by the CLI entrypoint.
	if a.Pack == nil {
		return fmt.Errorf("pack not loaded")
	}
	if a.State == nil {
		return fmt.Errorf("state not loaded")
	}

	if a.State.StartTime.IsZero() {
		a.State.StartTime = time.Now()
	}

	fmt.Fprintf(out, "Running pack: %s\n", a.Config.PackPath)
	fmt.Fprintf(out, "Output directory: %s\n", a.Runner.WorkDir)

	// Prelude
	if a.Pack.Prelude.Code != "" {
		fmt.Fprintln(out, "Executing prelude...")
		err := a.Runner.RunPrelude(ctx, &a.Pack.Prelude, out)
		a.recordWarnings()
		if err != nil {
			return fmt.Errorf("prelude failed: %w", err)
		}
	}

	for _, step := range a.Pack.Steps {
		// Filter by ROI
		if a.Config.ROIThreshold > 0 {
			if a.Config.ROIMode == "under" {
				// "under" is strictly less than
				if step.ROI >= a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f >= %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			} else {
				// "over" is greater than or equal to (3.3 or higher)
				if step.ROI < a.Config.ROIThreshold {
					fmt.Fprintf(out, "Skipping step %s (ROI %.2f < %.2f)\n", step.ID, step.ROI, a.Config.ROIThreshold)
					continue
				}
			}
		}

		// Check resume
		if s, ok := a.State.StepStatuses[step.ID]; ok && s.Status == state.StatusSuccess {
			fmt.Fprintf(out, "Skipping step %s (already succeeded)\n", step.ID)
			continue
		}

		fmt.Fprintf(out, "\n>>> Step %s: %s\n", step.ID, step.OriginalLine)

		status := state.StepStatus{
			Status:    state.StatusRunning,
			StartedAt: time.Now(),
		}
		a.State.StepStatuses[step.ID] = status
		a.saveState()

		// Execute
		err := a.Runner.RunStep(ctx, &step, out)
		a.recordWarnings()

		status.EndedAt = time.Now()
		if err != nil {
			status.Status = state.StatusFailed
			status.Error = err.Error()
			a.State.StepStatuses[step.ID] = status
			a.saveState()

			if a.Config.StopOnFail {
				a.finalize(out)
				return err
			}
			continue
		}

		status.Status = state.StatusSuccess
		status.ExitCode = 0
		a.State.StepStatuses[step.ID] = status
		a.saveState()
	}

	a.finalize(out)
	return nil
}

func (a *App) recordWarnings() {
	if a.State == nil || a.Runner == nil {
		return
	}
	warnings := a.Runner.DrainWarnings()
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		a.State.Warnings = append(a.State.Warnings, state.Warning{
			Scope:   w.Scope,
			StepID:  w.StepID,
			Line:    w.Line,
			Token:   w.Token,
			Message: w.Message,
		})
	}
	a.saveState()
}

func (a *App) saveState() {
	if a.Config.StatePath != "" {
		_ = state.SaveStateAtomic(a.Config.StatePath, a.State)
	}
}

func (a *App) finalize(out io.Writer) {
	if a.Config.ReportPath != "" {
		rep := report.GenerateReport(a.State, filepath.Base(a.Config.PackPath))
		data, _ := json.MarshalIndent(rep, "", "  ")
		_ = os.WriteFile(a.Config.ReportPath, data, 0644)
		fmt.Fprintf(out, "\nReport written to %s\n", a.Config.ReportPath)
	}
}
```

internal/app/run_test.go
```
package app

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
)

func TestApp_RunPlain_ROI(t *testing.T) {
	packContent := `
# ROI Test Pack
` + "```" + `bash
# 01) ROI=5.0
echo "high"
# 02) ROI=3.3
echo "threshold"
# 03) ROI=1.0
echo "low"
` + "```" + `
`
	packFile := "roi_test.md"
	defer os.Remove(packFile)
	os.WriteFile(packFile, []byte(packContent), 0644)

	// Test Case 1: Filter OVER 3.3 (Should run 5.0 and 3.3)
	t.Run("Filter Over 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "over",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if !strings.Contains(output, "Step 01") {
			t.Error("expected Step 01 (5.0) to run")
		}
		if !strings.Contains(output, "Step 02") {
			t.Error("expected Step 02 (3.3) to run (inclusive)")
		}
		if strings.Contains(output, "Step 03") && !strings.Contains(output, "Skipping step 03") {
			t.Error("expected Step 03 (1.0) to be skipped")
		}
	})

	// Test Case 2: Filter UNDER 3.3 (Should run 1.0 only)
	t.Run("Filter Under 3.3", func(t *testing.T) {
		var out bytes.Buffer
		cfg := Config{
			PackPath:     packFile,
			ROIThreshold: 3.3,
			ROIMode:      "under",
		}
		app := New(cfg)
		if err := app.Prepare(); err != nil {
			t.Fatalf("Prepare failed: %v", err)
		}
		if err := app.LoadState(); err != nil {
			t.Fatalf("LoadState failed: %v", err)
		}
		if err := app.RunPlain(context.Background(), &out); err != nil {
			t.Fatalf("RunPlain failed: %v", err)
		}
		output := out.String()
		if strings.Contains(output, "Step 01") && !strings.Contains(output, "Skipping step 01") {
			t.Error("expected Step 01 (5.0) to be skipped")
		}
		if strings.Contains(output, "Step 02") && !strings.Contains(output, "Skipping step 02") {
			t.Error("expected Step 02 (3.3) to be skipped (exclusive)")
		}
		if !strings.Contains(output, "Step 03") {
			t.Error("expected Step 03 (1.0) to run")
		}
	})
}
```

internal/cli/cmds.go
```
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
)

var validateCmd = &cobra.Command{
	Use:   "validate [pack.md]",
	Short: "Validate an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := app.Config{PackPath: args[0]}
		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		fmt.Println("Pack is valid.")
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list [pack.md]",
	Short: "List steps in an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := app.Config{PackPath: args[0]}
		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		for _, s := range a.Pack.Steps {
			fmt.Printf("%s: %s\n", s.ID, s.OriginalLine)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(listCmd)
}
```

internal/cli/root.go
```
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/errors"
)

var (
	noTUI     bool
	oracleBin string
	outDir    string
)

var rootCmd = &cobra.Command{
	Use:   "oraclepack",
	Short: "Oracle Pack Runner",
	Long:  `A polished TUI-driven runner for oracle-based interactive bash steps.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(errors.ExitCode(err))
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noTUI, "no-tui", false, "Disable the TUI and run in plain terminal mode")
	rootCmd.PersistentFlags().StringVar(&oracleBin, "oracle-bin", "oracle", "Path to the oracle binary")
	rootCmd.PersistentFlags().StringVarP(&outDir, "out-dir", "o", "", "Output directory for step execution")
}
```

internal/cli/run.go
```
package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/tui"
)

var (
	yes          bool
	resume       bool
	stopOnFail   bool
	roiThreshold float64
	roiMode      string
	runAll       bool
)

var runCmd = &cobra.Command{
	Use:   "run [pack.md]",
	Short: "Run an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		packPath := args[0]
		
		// Setup paths
		base := strings.TrimSuffix(filepath.Base(packPath), filepath.Ext(packPath))
		statePath := base + ".state.json"
		reportPath := base + ".report.json"

		cfg := app.Config{
			PackPath:     packPath,
			StatePath:    statePath,
			ReportPath:   reportPath,
			Resume:       resume,
			StopOnFail:   stopOnFail,
			WorkDir:      ".",
			OutDir:       outDir,
			ROIThreshold: roiThreshold,
			ROIMode:      roiMode,
		}

		a := app.New(cfg)
		// Prepare the application (loads pack, resolves out_dir, provisions env)
		if err := a.Prepare(); err != nil {
			return err
		}
		
		if err := a.LoadState(); err != nil {
			return err
		}

		if noTUI {
			return a.RunPlain(context.Background(), os.Stdout)
		}

		m := tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err := p.Run()
		return err
	},
}

func init() {
	runCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Auto-approve all steps")
	runCmd.Flags().BoolVar(&resume, "resume", false, "Resume from last successful step")
	runCmd.Flags().BoolVar(&stopOnFail, "stop-on-fail", true, "Stop execution if a step fails")
	runCmd.Flags().Float64Var(&roiThreshold, "roi-threshold", 0.0, "Filter steps by ROI threshold")
	runCmd.Flags().StringVar(&roiMode, "roi-mode", "over", "ROI filter mode ('over' or 'under')")
	runCmd.Flags().BoolVar(&runAll, "run-all", false, "Automatically run all steps sequentially on start")
	rootCmd.AddCommand(runCmd)
}
```

internal/errors/errors.go
```
package errors

import (
	"errors"
)

var (
	// ErrInvalidPack is returned when the Markdown pack is malformed.
	ErrInvalidPack = errors.New("invalid pack structure")
	// ErrExecutionFailed is returned when a shell command fails.
	ErrExecutionFailed = errors.New("execution failed")
	// ErrConfigInvalid is returned when CLI flags or environment variables are incorrect.
	ErrConfigInvalid = errors.New("invalid configuration")
)

// ExitCode returns the appropriate exit code for a given error.
func ExitCode(err error) int {
	if err == nil {
		return 0
	}

	if errors.Is(err, ErrConfigInvalid) {
		return 2
	}

	if errors.Is(err, ErrInvalidPack) {
		return 3
	}

	if errors.Is(err, ErrExecutionFailed) {
		return 4
	}

	return 1 // Generic error
}
```

internal/errors/errors_test.go
```
package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestExitCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{"nil error", nil, 0},
		{"generic error", errors.New("generic"), 1},
		{"invalid pack", ErrInvalidPack, 3},
		{"execution failed", ErrExecutionFailed, 4},
		{"config invalid", ErrConfigInvalid, 2},
		{"wrapped invalid pack", fmt.Errorf("wrap: %w", ErrInvalidPack), 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExitCode(tt.err); got != tt.expected {
				t.Errorf("ExitCode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
```

internal/exec/flags.go
```
package exec

import "strings"

// ApplyChatGPTURL ensures a single --chatgpt-url flag is present when url is set.
// It removes any existing --chatgpt-url/--browser-url flags and their values.
func ApplyChatGPTURL(flags []string, url string) []string {
	var out []string
	skipNext := false
	for _, f := range flags {
		if skipNext {
			skipNext = false
			continue
		}
		if f == "--chatgpt-url" || f == "--browser-url" {
			skipNext = true
			continue
		}
		if strings.HasPrefix(f, "--chatgpt-url=") || strings.HasPrefix(f, "--browser-url=") {
			continue
		}
		out = append(out, f)
	}
	if url != "" {
		out = append(out, "--chatgpt-url", url)
	}
	return out
}
```

internal/exec/inject.go
```
package exec

import "strings"

// InjectFlags scans a script and appends flags to any 'oracle' command invocation.
func InjectFlags(script string, flags []string) string {
	if len(flags) == 0 {
		return script
	}

	flagStr := strings.Join(flags, " ")

	lines := strings.Split(script, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		insertIdx := oracleInsertIndex(line)
		if insertIdx == -1 {
			continue
		}

		lines[i] = insertFlagsInLine(line, insertIdx, flagStr)
	}

	return strings.Join(lines, "\n")
}

func oracleInsertIndex(line string) int {
	i := 0
	for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
		i++
	}

	if !strings.HasPrefix(line[i:], "oracle") {
		return -1
	}

	end := i + len("oracle")
	if end < len(line) {
		next := line[end]
		if next != ' ' && next != '\t' {
			return -1
		}
	}

	return end
}

func insertFlagsInLine(line string, insertIdx int, flags string) string {
	prefix := line[:insertIdx]
	rest := line[insertIdx:]
	if rest == "" {
		return prefix + " " + flags
	}
	if rest[0] == ' ' || rest[0] == '\t' {
		return prefix + " " + flags + rest
	}
	return prefix + " " + flags + " " + rest
}
```

internal/exec/inject_test.go
```
package exec

import (
	"testing"
)

func TestInjectFlags(t *testing.T) {
	tests := []struct {
		name     string
		script   string
		flags    []string
		expected string
	}{
		{
			"simple injection",
			"oracle query 'hello'",
			[]string{"--verbose"},
			"oracle --verbose query 'hello'",
		},
		{
			"indented injection",
			"  oracle query 'hello'",
			[]string{"--verbose"},
			"  oracle --verbose query 'hello'",
		},
		{
			"no injection needed",
			"echo 'hello'",
			[]string{"--verbose"},
			"echo 'hello'",
		},
		{
			"multiple lines",
			"echo 'start'\noracle query\necho 'end'",
			[]string{"--debug"},
			"echo 'start'\noracle --debug query\necho 'end'",
		},
		{
			"multiline with continuation",
			"oracle \\\n  --json \\\n  --files",
			[]string{"--flag"},
			"oracle --flag \\\n  --json \\\n  --files",
		},
		{
			"multiline with args and continuation",
			"  oracle arg \\\n  --json",
			[]string{"--flag"},
			"  oracle --flag arg \\\n  --json",
		},
		{
			"commented command",
			"# oracle --json",
			[]string{"--verbose"},
			"# oracle --json",
		},
		{
			"oracle as part of word",
			"coracle query",
			[]string{"--verbose"},
			"coracle query",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InjectFlags(tt.script, tt.flags)
			if got != tt.expected {
				t.Errorf("InjectFlags() = %q, want %q", got, tt.expected)
			}
		})
	}
}
```

internal/exec/oracle_scan.go
```
package exec

import (
	"regexp"
	"strings"
)

var oracleCmdRegex = regexp.MustCompile(`^(\s*)(oracle)\b`)

// OracleInvocation represents a detected oracle command in a script.
type OracleInvocation struct {
	StartLine   int    // 0-based start line index
	EndLine     int    // 0-based end line index (inclusive)
	Raw         string // The full command string (joined if multi-line)
	Display     string // A trimmed version for UI display
	Indentation string // The leading whitespace
}

// ExtractOracleInvocations extracts oracle invocations from a script.
func ExtractOracleInvocations(script string) []OracleInvocation {
	var invocations []OracleInvocation
	lines := strings.Split(script, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Skip comments
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Check for oracle command
		loc := oracleCmdRegex.FindStringSubmatchIndex(line)
		if loc != nil {
			startLine := i
			// Group 1 is the indentation
			indentation := line[loc[2]:loc[3]]

			var cmdBuilder strings.Builder
			cmdBuilder.WriteString(line)

			endLine := i
			// Handle line continuations
			// Check if line ends with backslash (ignoring trailing whitespace)
			for {
				if endLine+1 >= len(lines) {
					break
				}

				// Check current line for continuation
				currTrimmed := strings.TrimRight(lines[endLine], " \t")
				if !strings.HasSuffix(currTrimmed, "\\") {
					break
				}

				endLine++
				cmdBuilder.WriteString("\n")
				cmdBuilder.WriteString(lines[endLine])
			}

			raw := cmdBuilder.String()
			invocations = append(invocations, OracleInvocation{
				StartLine:   startLine,
				EndLine:     endLine,
				Raw:         raw,
				Display:     strings.TrimSpace(raw),
				Indentation: indentation,
			})

			i = endLine // Advance loop
		}
	}
	return invocations
}
```

internal/exec/oracle_scan_test.go
```
package exec

import (
	"reflect"
	"testing"
)

func TestExtractOracleInvocations(t *testing.T) {
	tests := []struct {
		name   string
		script string
		want   []OracleInvocation
	}{
		{
			name:   "Simple command",
			script: "oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "oracle --json", Display: "oracle --json", Indentation: ""},
			},
		},
		{
			name:   "Indented command",
			script: "  oracle --json",
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 0, Raw: "  oracle --json", Display: "oracle --json", Indentation: "  "},
			},
		},
		{
			name: "Multiline command",
			script: `oracle \
  --json \
  --files`,
			want: []OracleInvocation{
				{StartLine: 0, EndLine: 2, Raw: `oracle \
  --json \
  --files`, Display: `oracle \
  --json \
  --files`, Indentation: ""},
			},
		},
		{
			name: "Commented command",
			script: `# oracle --json
oracle --real`,
			want: []OracleInvocation{
				{StartLine: 1, EndLine: 1, Raw: "oracle --real", Display: "oracle --real", Indentation: ""},
			},
		},
		{
			name: "Multiple commands",
			script: `
echo start
oracle --one
echo mid
oracle --two
echo end
`,
			want: []OracleInvocation{
				{StartLine: 2, EndLine: 2, Raw: "oracle --one", Display: "oracle --one", Indentation: ""},
				{StartLine: 4, EndLine: 4, Raw: "oracle --two", Display: "oracle --two", Indentation: ""},
			},
		},
		{
			name:   "Oraclepack prefix (should not match)",
			script: "oraclepack run",
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractOracleInvocations(tt.script)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractOracleInvocations() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
```

internal/exec/oracle_validate.go
```
package exec

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

// ValidationError captures a failed oracle validation for a step.
type ValidationError struct {
	StepID       string
	Command      string
	ErrorMessage string
}

// ValidateOverrides runs oracle --dry-run summary for targeted steps.
func ValidateOverrides(
	ctx context.Context,
	steps []pack.Step,
	over *overrides.RuntimeOverrides,
	baseline []string,
	opts RunnerOptions,
) ([]ValidationError, error) {
	if over == nil || over.ApplyToSteps == nil {
		return nil, nil
	}

	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}
	env := append(os.Environ(), opts.Env...)

	var results []ValidationError
	for _, step := range steps {
		if !over.ApplyToSteps[step.ID] {
			continue
		}

		invocations := ExtractOracleInvocations(step.Code)
		if len(invocations) == 0 {
			continue
		}

		flags := over.EffectiveFlags(step.ID, baseline)
		flags = append(flags, "--dry-run", "summary")

		for _, inv := range invocations {
			cmdStr := InjectFlags(inv.Raw, flags)
			msg, err := execDryRun(ctx, shell, opts.WorkDir, env, cmdStr)
			if err == nil {
				continue
			}

			results = append(results, ValidationError{
				StepID:       step.ID,
				Command:      cmdStr,
				ErrorMessage: msg,
			})
		}
	}

	return results, nil
}

func execDryRun(ctx context.Context, shell, workDir string, env []string, command string) (string, error) {
	if pathVal := findEnvValue(env, "PATH"); pathVal != "" {
		command = "export PATH=" + shellQuote(pathVal) + "; " + command
	}

	cmd := exec.CommandContext(ctx, shell, "-lc", command)
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Env = env

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		return stdout.String(), nil
	}
	if stderr.Len() > 0 {
		return strings.TrimSpace(stderr.String()), err
	}
	if stdout.Len() > 0 {
		return strings.TrimSpace(stdout.String()), err
	}
	return err.Error(), err
}

func findEnvValue(env []string, key string) string {
	prefix := key + "="
	for _, entry := range env {
		if strings.HasPrefix(entry, prefix) {
			return strings.TrimPrefix(entry, prefix)
		}
	}
	return ""
}

func shellQuote(value string) string {
	if value == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}
```

internal/exec/oracle_validate_test.go
```
package exec

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

func TestValidateOverrides_Success(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --ok"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	_, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		[]string{"--base"},
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
}

func TestValidateOverrides_Error(t *testing.T) {
	dir := t.TempDir()
	writeOracleStub(t, dir)

	steps := []pack.Step{
		{ID: "01", Code: "oracle --bad"},
	}
	over := &overrides.RuntimeOverrides{
		ApplyToSteps: map[string]bool{"01": true},
	}

	errs, err := ValidateOverrides(
		context.Background(),
		steps,
		over,
		nil,
		RunnerOptions{
			WorkDir: dir,
			Env:     []string{"PATH=" + dir + string(os.PathListSeparator) + os.Getenv("PATH")},
		},
	)
	if err != nil {
		t.Fatalf("ValidateOverrides failed: %v", err)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 validation error, got %d", len(errs))
	}
	msg := errs[0].ErrorMessage
	if !strings.Contains(msg, "invalid flag") && !strings.Contains(msg, "unknown option") {
		t.Fatalf("unexpected error message: %q", msg)
	}
	if !strings.Contains(errs[0].Command, "--dry-run summary") {
		t.Fatalf("expected command to include --dry-run summary, got %q", errs[0].Command)
	}
}

func writeOracleStub(t *testing.T, dir string) {
	t.Helper()
	stub := `#!/bin/sh
has_dry=0
has_summary=0
for arg in "$@"; do
  if [ "$arg" = "--dry-run" ]; then has_dry=1; fi
  if [ "$arg" = "summary" ]; then has_summary=1; fi
  if [ "$arg" = "--bad" ]; then echo "invalid flag" 1>&2; exit 1; fi
done
if [ $has_dry -eq 0 ] || [ $has_summary -eq 0 ]; then
  echo "missing dry run" 1>&2
  exit 1
fi
exit 0
`
	path := filepath.Join(dir, "oracle")
	if err := os.WriteFile(path, []byte(stub), 0o755); err != nil {
		t.Fatalf("write oracle stub: %v", err)
	}
}
```

internal/exec/runner.go
```
package exec

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/user/oraclepack/internal/errors"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

// Runner handles the execution of shell scripts.
type Runner struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
	warnings    []SanitizeWarning
}

// RunnerOptions configures a Runner.
type RunnerOptions struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
	Overrides   *overrides.RuntimeOverrides
	ChatGPTURL  string
}

// NewRunner creates a new Runner with options.
func NewRunner(opts RunnerOptions) *Runner {
	shell := opts.Shell
	if shell == "" {
		shell = "/bin/bash"
	}

	return &Runner{
		Shell:       shell,
		WorkDir:     opts.WorkDir,
		Env:         append(os.Environ(), opts.Env...),
		OracleFlags: opts.OracleFlags,
		Overrides:   opts.Overrides,
		ChatGPTURL:  opts.ChatGPTURL,
	}
}

// RunPrelude executes the prelude code.
func (r *Runner) RunPrelude(ctx context.Context, p *pack.Prelude, logWriter io.Writer) error {
	script, warnings := SanitizeScript(p.Code, "prelude", "")
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

// RunStep executes a single step's code.
func (r *Runner) RunStep(ctx context.Context, s *pack.Step, logWriter io.Writer) error {
	flags := ApplyChatGPTURL(r.OracleFlags, r.ChatGPTURL)
	if r.Overrides != nil {
		flags = r.Overrides.EffectiveFlags(s.ID, r.OracleFlags)
		flags = ApplyChatGPTURL(flags, r.ChatGPTURL)
	}
	code := InjectFlags(s.Code, flags)
	script, warnings := SanitizeScript(code, "step", s.ID)
	r.recordWarnings(warnings, logWriter)
	return r.run(ctx, script, logWriter)
}

func (r *Runner) recordWarnings(warnings []SanitizeWarning, logWriter io.Writer) {
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		r.warnings = append(r.warnings, w)
		if logWriter != nil {
			scope := w.Scope
			if scope == "" {
				scope = "script"
			}
			step := ""
			if w.StepID != "" {
				step = " step " + w.StepID
			}
			_, _ = fmt.Fprintf(logWriter, "⚠ oraclepack: sanitized label in %s%s line %d: %s\n", scope, step, w.Line, w.Token)
		}
	}
}

// DrainWarnings returns any sanitizer warnings collected since the last call.
func (r *Runner) DrainWarnings() []SanitizeWarning {
	if len(r.warnings) == 0 {
		return nil
	}
	out := make([]SanitizeWarning, len(r.warnings))
	copy(out, r.warnings)
	r.warnings = nil
	return out
}

func (r *Runner) run(ctx context.Context, script string, logWriter io.Writer) error {
	// We use bash -lc to ensure login shell (paths, aliases, etc)
	cmd := exec.CommandContext(ctx, r.Shell, "-lc", script)
	cmd.Dir = r.WorkDir
	cmd.Env = r.Env

	// Standardize stdout and stderr to the logWriter
	cmd.Stdout = logWriter
	cmd.Stderr = logWriter

	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return fmt.Errorf("%w: %v", errors.ErrExecutionFailed, err)
	}

	return nil
}
```

internal/exec/runner_test.go
```
package exec

import (
	"context"
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/pack"
)

func TestRunner_RunStep(t *testing.T) {
	r := NewRunner(RunnerOptions{})
	
	var lines []string
	lw := &LineWriter{
		Callback: func(line string) {
			lines = append(lines, line)
		},
	}

	step := &pack.Step{
		Code: "echo 'hello world'",
	}

	err := r.RunStep(context.Background(), step, lw)
	if err != nil {
		t.Fatalf("RunStep failed: %v", err)
	}
	lw.Close()

	found := false
	for _, l := range lines {
		if strings.TrimSpace(l) == "hello world" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected 'hello world' in output, got: %v", lines)
	}
}

func TestRunner_ContextCancellation(t *testing.T) {
	r := NewRunner(RunnerOptions{})
	
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	step := &pack.Step{
		Code: "sleep 10",
	}

	err := r.RunStep(ctx, step, nil)
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}
```

internal/exec/sanitize.go
```
package exec

import (
	osexec "os/exec"
	"regexp"
	"strings"
)

// SanitizeWarning captures a label line that was converted to a safe echo.
type SanitizeWarning struct {
	Scope   string
	StepID  string
	Line    int
	Token   string
	Message string
}

var (
	labelTokenRegex   = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*$`)
	heredocStartRegex = regexp.MustCompile(`<<-?\s*['"]?([A-Za-z0-9_]+)['"]?`)
)

var shellBuiltins = map[string]bool{
	"alias":    true,
	"bg":       true,
	"break":    true,
	"cd":       true,
	"command":  true,
	"continue": true,
	"declare":  true,
	"dirs":     true,
	"echo":     true,
	"eval":     true,
	"exec":     true,
	"exit":     true,
	"export":   true,
	"fg":       true,
	"getopts":  true,
	"hash":     true,
	"help":     true,
	"jobs":     true,
	"local":    true,
	"popd":     true,
	"printf":   true,
	"pushd":    true,
	"pwd":      true,
	"readonly": true,
	"return":   true,
	"set":      true,
	"shift":    true,
	"source":   true,
	"test":     true,
	"trap":     true,
	"true":     true,
	"type":     true,
	"ulimit":   true,
	"umask":    true,
	"unalias":  true,
	"unset":    true,
	"wait":     true,
	"false":    true,
}

var shellKeywords = map[string]bool{
	"case":     true,
	"do":       true,
	"done":     true,
	"elif":     true,
	"else":     true,
	"esac":     true,
	"fi":       true,
	"for":      true,
	"function": true,
	"if":       true,
	"in":       true,
	"select":   true,
	"then":     true,
	"time":     true,
	"until":    true,
	"while":    true,
}

// SanitizeScript converts bare label-like lines into safe echo statements.
func SanitizeScript(script, scope, stepID string) (string, []SanitizeWarning) {
	if script == "" {
		return script, nil
	}

	lines := strings.Split(script, "\n")
	var warnings []SanitizeWarning
	var heredocEnd string

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if heredocEnd != "" {
			if trimmed == heredocEnd {
				heredocEnd = ""
			}
			continue
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		if end := heredocStartToken(trimmed); end != "" {
			heredocEnd = end
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) != 1 {
			continue
		}
		token := fields[0]
		if !labelTokenRegex.MatchString(token) {
			continue
		}
		lower := strings.ToLower(token)
		if shellBuiltins[lower] || shellKeywords[lower] {
			continue
		}
		if _, err := osexec.LookPath(token); err == nil {
			continue
		}

		indent := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
		lines[i] = indent + "echo \"" + token + "\""
		warnings = append(warnings, SanitizeWarning{
			Scope:   scope,
			StepID:  stepID,
			Line:    i + 1,
			Token:   token,
			Message: "Converted bare label to echo",
		})
	}

	return strings.Join(lines, "\n"), warnings
}

func heredocStartToken(line string) string {
	match := heredocStartRegex.FindStringSubmatch(line)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}
```

internal/exec/sanitize_test.go
```
package exec

import "testing"

func TestSanitizeScript_LabelLine(t *testing.T) {
	input := "GenerateReport\noracle --help\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(warnings))
	}
	if warnings[0].Token != "GenerateReport" {
		t.Fatalf("expected token GenerateReport, got %s", warnings[0].Token)
	}
	wantPrefix := "echo \"GenerateReport\""
	if got[:len(wantPrefix)] != wantPrefix {
		t.Fatalf("expected sanitized line to start with %q, got %q", wantPrefix, got)
	}
}

func TestSanitizeScript_BuiltinUnchanged(t *testing.T) {
	input := "echo\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected script unchanged, got %q", got)
	}
}

func TestSanitizeScript_HeredocUnchanged(t *testing.T) {
	input := "cat <<'EOF'\nGenerateReport\nEOF\n"
	got, warnings := SanitizeScript(input, "step", "01")
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %d", len(warnings))
	}
	if got != input {
		t.Fatalf("expected heredoc unchanged, got %q", got)
	}
}
```

internal/exec/stream.go
```
package exec

import (
	"io"
)

// LineWriter is an io.Writer that splits output into lines and calls a callback.
type LineWriter struct {
	Callback func(string)
	buffer   []byte
}

func (w *LineWriter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			w.Callback(string(w.buffer))
			w.buffer = w.buffer[:0]
		} else {
			w.buffer = append(w.buffer, b)
		}
	}
	return len(p), nil
}

// Close flushes any remaining data in the buffer.
func (w *LineWriter) Close() error {
	if len(w.buffer) > 0 {
		w.Callback(string(w.buffer))
		w.buffer = w.buffer[:0]
	}
	return nil
}

// MultiWriter handles multiple writers efficiently.
func MultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}
```

internal/overrides/merge.go
```
package overrides

// EffectiveFlags calculates the final flags for a step.
func (r *RuntimeOverrides) EffectiveFlags(stepID string, baseline []string) []string {
	if r == nil || r.ApplyToSteps == nil || !r.ApplyToSteps[stepID] {
		return baseline
	}

	var effective []string

	// Map for removed flags
	removed := make(map[string]bool)
	for _, f := range r.RemovedFlags {
		removed[f] = true
	}

	// Filter baseline
	for _, flag := range baseline {
		if !removed[flag] {
			effective = append(effective, flag)
		}
	}

	// Append added flags
	effective = append(effective, r.AddedFlags...)

	// Inject ChatGPTURL
	if r.ChatGPTURL != "" {
		effective = append(effective, "--chatgpt-url", r.ChatGPTURL)
	}

	return effective
}
```

internal/overrides/merge_test.go
```
package overrides

import (
	"reflect"
	"testing"
)

func TestEffectiveFlags(t *testing.T) {
	tests := []struct {
		name      string
		overrides *RuntimeOverrides
		stepID    string
		baseline  []string
		want      []string
	}{
		{
			name:      "No overrides (nil)",
			overrides: nil,
			stepID:    "01",
			baseline:  []string{"--json"},
			want:      []string{"--json"},
		},
		{
			name: "Step not targeted",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"02": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json"},
		},
		{
			name: "Step targeted: Add flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--verbose"},
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--verbose"},
		},
		{
			name: "Step targeted: Remove flags",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				RemovedFlags: []string{"--json"},
			},
			stepID:   "01",
			baseline: []string{"--json", "--other"},
			want:     []string{"--other"},
		},
		{
			name: "Step targeted: Add and Remove",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				AddedFlags:   []string{"--new"},
				RemovedFlags: []string{"--old"},
			},
			stepID:   "01",
			baseline: []string{"--old", "--keep"},
			want:     []string{"--keep", "--new"},
		},
		{
			name: "Step targeted: Inject ChatGPT URL",
			overrides: &RuntimeOverrides{
				ApplyToSteps: map[string]bool{"01": true},
				ChatGPTURL:   "https://chat.openai.com/share/123",
			},
			stepID:   "01",
			baseline: []string{"--json"},
			want:     []string{"--json", "--chatgpt-url", "https://chat.openai.com/share/123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.overrides.EffectiveFlags(tt.stepID, tt.baseline)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EffectiveFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

internal/overrides/types.go
```
package overrides

// RuntimeOverrides holds configuration for runtime flag modifications.
type RuntimeOverrides struct {
	AddedFlags   []string        // Flags to append (e.g., "--model=gpt-4")
	RemovedFlags []string        // Flags to remove (e.g., "--json")
	ChatGPTURL   string          // Optional URL to inject via --chatgpt-url
	ApplyToSteps map[string]bool // Set of step IDs to apply overrides to. If empty, applies to none.
}
```

internal/pack/parser.go
```
package pack

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/user/oraclepack/internal/errors"
)

var (
	bashFenceRegex = regexp.MustCompile("(?s)```bash\n(.*?)\n```")
	// Updated regex to support ")", " —", and " -" separators
	stepHeaderRegex = regexp.MustCompile(`^#\s*(\d{2})(?:\)|[\s]+[—-])`)
	roiRegex        = regexp.MustCompile(`ROI=(\d+(\.\d+)?)`)
	outDirRegex    = regexp.MustCompile(`(?m)^out_dir=["']?([^"'\s]+)["']?`)
	writeOutputRegex = regexp.MustCompile(`(?m)--write-output`)
)

// Parse reads a Markdown content and returns a Pack.
func Parse(content []byte) (*Pack, error) {
	match := bashFenceRegex.FindSubmatch(content)
	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("%w: no bash code block found", errors.ErrInvalidPack)
	}

	bashCode := string(match[1])
	pack := &Pack{}
	
	scanner := bufio.NewScanner(strings.NewReader(bashCode))
	var currentStep *Step
	var preludeLines []string
	var inSteps bool

	for scanner.Scan() {
		line := scanner.Text()
		headerMatch := stepHeaderRegex.FindStringSubmatch(strings.TrimSpace(line))

		if len(headerMatch) > 1 {
			inSteps = true
			if currentStep != nil {
				pack.Steps = append(pack.Steps, *currentStep)
			}
			num, _ := strconv.Atoi(headerMatch[1])
			
			// Extract ROI if present
			var roi float64
			cleanedLine := line
			roiMatch := roiRegex.FindStringSubmatch(line)
			if len(roiMatch) > 1 {
				val, err := strconv.ParseFloat(roiMatch[1], 64)
				if err == nil {
					roi = val
					// Remove ROI tag from display title, but keep original line intact?
					// The task says "strip from Step.Title". Step struct currently has `OriginalLine`.
					// I'll assume OriginalLine is what is displayed, or I should add a Title field.
					// Looking at Step struct: ID, Number, Code, OriginalLine.
					// I'll remove it from OriginalLine for now or add a Title field.
					// The existing TUI uses OriginalLine as description. 
					// Let's clean OriginalLine for display purposes or add a dedicated Title field.
					// Adding a dedicated Title field seems cleaner but requires struct change.
					// For now, I'll strip it from OriginalLine to match the prompt requirement "cleaner UI display".
					cleanedLine = strings.Replace(cleanedLine, roiMatch[0], "", 1)
					cleanedLine = strings.TrimSpace(cleanedLine)
					// Fix any double spaces or trailing separators if needed, but simple replace is a good start.
				}
			}

			currentStep = &Step{
				ID:           headerMatch[1],
				Number:       num,
				OriginalLine: cleanedLine,
				ROI:          roi,
			}
			continue
		}

		if inSteps {
			currentStep.Code += line + "\n"
		} else {
			preludeLines = append(preludeLines, line)
		}
	}

	if currentStep != nil {
		pack.Steps = append(pack.Steps, *currentStep)
	}

	pack.Prelude.Code = strings.Join(preludeLines, "\n")
	pack.DeriveMetadata()

	return pack, nil
}

// DeriveMetadata extracts configuration from the prelude.
func (p *Pack) DeriveMetadata() {
	outDirMatch := outDirRegex.FindStringSubmatch(p.Prelude.Code)
	if len(outDirMatch) > 1 {
		p.OutDir = outDirMatch[1]
	}

	if writeOutputRegex.MatchString(p.Prelude.Code) {
		p.WriteOutput = true
	}
}

// Validate checks if the pack follows all rules.
func (p *Pack) Validate() error {
	if len(p.Steps) == 0 {
		return fmt.Errorf("%w: at least one step is required", errors.ErrInvalidPack)
	}

	seen := make(map[int]bool)
	for i, step := range p.Steps {
		if step.Number <= 0 {
			return fmt.Errorf("%w: invalid step number %d", errors.ErrInvalidPack, step.Number)
		}
		if seen[step.Number] {
			return fmt.Errorf("%w: duplicate step number %d", errors.ErrInvalidPack, step.Number)
		}
		seen[step.Number] = true

		// Optional: Ensure sequential starting from 1
		if step.Number != i+1 {
			return fmt.Errorf("%w: steps must be sequential starting from 1 (expected %d, got %d)", errors.ErrInvalidPack, i+1, step.Number)
		}
	}

	return nil
}
```

internal/pack/parser_test.go
```
package pack

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	content := []byte(`
# My Pack
Some description.

` + "```" + `bash
out_dir="dist"
--write-output

# 01)
echo "hello"

# 02)
echo "world"
` + "```" + `
`)

	p, err := Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if p.OutDir != "dist" {
		t.Errorf("expected OutDir dist, got %s", p.OutDir)
	}

	if !p.WriteOutput {
		t.Errorf("expected WriteOutput true, got false")
	}

	if len(p.Steps) != 2 {
		t.Errorf("expected 2 steps, got %d", len(p.Steps))
	}

	if p.Steps[0].ID != "01" || p.Steps[0].Number != 1 {
		t.Errorf("step 1 mismatch: %+v", p.Steps[0])
	}

	if err := p.Validate(); err != nil {
		t.Errorf("Validate failed: %v", err)
	}
}

func TestParseVariants(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			"em dash",
			`
` + "```" + `bash
# 01 — ROI=...
echo "step 1"
` + "```" + `
`,
		},
		{
			"hyphen",
			`
` + "```" + `bash
# 01 - ROI=...
echo "step 1"
` + "```" + `
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := Parse([]byte(tt.content))
			if err != nil {
				t.Fatalf("Parse failed: %v", err)
			}
			if len(p.Steps) != 1 {
				t.Errorf("expected 1 step, got %d", len(p.Steps))
			}
		})
	}
}

func TestParseROI(t *testing.T) {
	content := []byte(`
` + "```" + `bash
# 01) ROI=4.5 clean me
echo "high value"

# 02) ROI=0.5
echo "low value"

# 03) No ROI
echo "default"
` + "```" + `
`)

	p, err := Parse(content)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(p.Steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(p.Steps))
	}

	if p.Steps[0].ROI != 4.5 {
		t.Errorf("step 1 ROI mismatch: expected 4.5, got %f", p.Steps[0].ROI)
	}
	if strings.Contains(p.Steps[0].OriginalLine, "ROI=4.5") {
		t.Errorf("step 1 title was not cleaned: %q", p.Steps[0].OriginalLine)
	}

	if p.Steps[1].ROI != 0.5 {
		t.Errorf("step 2 ROI mismatch: expected 0.5, got %f", p.Steps[1].ROI)
	}

	if p.Steps[2].ROI != 0.0 {
		t.Errorf("step 3 ROI mismatch: expected 0.0, got %f", p.Steps[2].ROI)
	}
}

func TestValidateErrors(t *testing.T) {
	tests := []struct {
		name    string
		pack    *Pack
		wantErr string
	}{
		{
			"no steps",
			&Pack{},
			"at least one step is required",
		},
		{
			"duplicate steps",
			&Pack{
				Steps: []Step{
					{Number: 1, ID: "01"},
					{Number: 1, ID: "01"},
				},
			},
			"duplicate step number 1",
		},
		{
			"non-sequential",
			&Pack{
				Steps: []Step{
					{Number: 1, ID: "01"},
					{Number: 3, ID: "03"},
				},
			},
			"steps must be sequential starting from 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pack.Validate()
			if err == nil {
				t.Error("expected error, got nil")
			} else if !contains(err.Error(), tt.wantErr) {
				t.Errorf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr))))
}
```

internal/pack/types.go
```
package pack

// Pack represents a parsed oracle pack.
type Pack struct {
	Prelude     Prelude
	Steps       []Step
	Source      string
	OutDir      string
	WriteOutput bool
}

// Prelude contains the shell code that runs before any steps.
type Prelude struct {
	Code string
}

// Step represents an individual executable step within the pack.
type Step struct {
	ID           string  // e.g., "01"
	Number       int     // e.g., 1
	Code         string  // The bash code
	OriginalLine string  // The header line, e.g., "# 01)"
	ROI          float64 // Return on Investment value extracted from header
}
```

internal/render/render.go
```
package render

import (
	"sync"

	"github.com/charmbracelet/glamour"
	"github.com/user/oraclepack/internal/pack"
)

const (
	DefaultStyle = "dark"
	DefaultWidth = 80
)

type rendererKey struct {
	width int
	style string
}

var (
	rendererMu    sync.Mutex
	rendererCache = map[rendererKey]*glamour.TermRenderer{}
)

// RenderMarkdown renders markdown text as ANSI-styled text.
func RenderMarkdown(text string, width int, style string) (string, error) {
	if width <= 0 {
		width = DefaultWidth
	}
	if style == "" {
		style = DefaultStyle
	}

	r, err := rendererFor(width, style)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}

// RenderStepCode renders a step's code block for preview.
func RenderStepCode(s pack.Step, width int, style string) (string, error) {
	md := "```bash\n" + s.Code + "\n```"
	return RenderMarkdown(md, width, style)
}

func rendererFor(width int, style string) (*glamour.TermRenderer, error) {
	key := rendererKey{width: width, style: style}

	rendererMu.Lock()
	r := rendererCache[key]
	rendererMu.Unlock()
	if r != nil {
		return r, nil
	}

	opts := []glamour.TermRendererOption{glamour.WithWordWrap(width)}
	if style == "auto" {
		opts = append(opts, glamour.WithAutoStyle())
	} else {
		opts = append(opts, glamour.WithStandardStyle(style))
	}

	r, err := glamour.NewTermRenderer(opts...)
	if err != nil {
		return nil, err
	}

	rendererMu.Lock()
	rendererCache[key] = r
	rendererMu.Unlock()
	return r, nil
}
```

internal/render/render_test.go
```
package render

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	text := "# Hello\n**bold**"
	got, err := RenderMarkdown(text, 40, DefaultStyle)
	if err != nil {
		t.Fatalf("RenderMarkdown failed: %v", err)
	}

	// ANSI escape codes start with \x1b[
	if !strings.Contains(got, "\x1b[") {
		t.Errorf("expected ANSI codes in output, got: %q", got)
	}
}
```

internal/report/generate.go
```
package report

import (
	"time"

	"github.com/user/oraclepack/internal/state"
)

// GenerateReport creates a ReportV1 from a RunState.
func GenerateReport(s *state.RunState, packName string) *ReportV1 {
	report := &ReportV1{
		PackInfo: PackInfo{
			Name: packName,
			Hash: s.PackHash,
		},
		GeneratedAt: time.Now(),
		Steps:       []StepReport{},
	}

	var totalDuration time.Duration
	success, failure, skipped := 0, 0, 0

	for id, status := range s.StepStatuses {
		duration := status.EndedAt.Sub(status.StartedAt)
		if status.EndedAt.IsZero() || status.StartedAt.IsZero() {
			duration = 0
		}

		totalDuration += duration

		sr := StepReport{
			ID:         id,
			Status:     string(status.Status),
			ExitCode:   status.ExitCode,
			Duration:   duration,
			DurationMs: duration.Milliseconds(),
			Error:      status.Error,
		}
		report.Steps = append(report.Steps, sr)

		switch status.Status {
		case state.StatusSuccess:
			success++
		case state.StatusFailed:
			failure++
		case state.StatusSkipped:
			skipped++
		}
	}

	report.Summary = Summary{
		TotalSteps:      len(s.StepStatuses),
		SuccessCount:    success,
		FailureCount:    failure,
		SkippedCount:    skipped,
		TotalDuration:   totalDuration,
		TotalDurationMs: totalDuration.Milliseconds(),
	}

	if len(s.Warnings) > 0 {
		report.Warnings = make([]Warning, 0, len(s.Warnings))
		for _, w := range s.Warnings {
			report.Warnings = append(report.Warnings, Warning{
				Scope:   w.Scope,
				StepID:  w.StepID,
				Line:    w.Line,
				Token:   w.Token,
				Message: w.Message,
			})
		}
	}

	return report
}
```

internal/report/report_test.go
```
package report

import (
	"testing"
	"time"

	"github.com/user/oraclepack/internal/state"
)

func TestGenerateReport(t *testing.T) {
	s := &state.RunState{
		PackHash: "hash123",
		StepStatuses: map[string]state.StepStatus{
			"01": {
				Status:    state.StatusSuccess,
				StartedAt: time.Now().Add(-1 * time.Second),
				EndedAt:   time.Now(),
			},
		},
	}

	rep := GenerateReport(s, "my-pack")

	if rep.PackInfo.Name != "my-pack" {
		t.Errorf("expected name my-pack, got %s", rep.PackInfo.Name)
	}

	if rep.Summary.TotalSteps != 1 {
		t.Errorf("expected 1 total step, got %d", rep.Summary.TotalSteps)
	}

	if rep.Summary.SuccessCount != 1 {
		t.Errorf("expected 1 success, got %d", rep.Summary.SuccessCount)
	}
}
```

internal/report/types.go
```
package report

import (
	"time"
)

// ReportV1 represents the final machine-readable summary.
type ReportV1 struct {
	Summary     Summary      `json:"summary"`
	PackInfo    PackInfo     `json:"pack_info"`
	Steps       []StepReport `json:"steps"`
	Warnings    []Warning    `json:"warnings,omitempty"`
	GeneratedAt time.Time    `json:"generated_at"`
}

type Summary struct {
	TotalSteps      int           `json:"total_steps"`
	SuccessCount    int           `json:"success_count"`
	FailureCount    int           `json:"failure_count"`
	SkippedCount    int           `json:"skipped_count"`
	TotalDuration   time.Duration `json:"total_duration"`
	TotalDurationMs int64         `json:"total_duration_ms"`
}

type PackInfo struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type StepReport struct {
	ID         string        `json:"id"`
	Status     string        `json:"status"`
	ExitCode   int           `json:"exit_code"`
	Duration   time.Duration `json:"duration"`
	DurationMs int64         `json:"duration_ms"`
	Error      string        `json:"error,omitempty"`
}

// Warning captures non-fatal execution notes surfaced during a run.
type Warning struct {
	Scope   string `json:"scope"`
	StepID  string `json:"step_id,omitempty"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
```

internal/state/persist.go
```
package state

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveStateAtomic saves the state to a file atomically.
func SaveStateAtomic(path string, state *RunState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal state: %w", err)
	}

	tempPath := path + ".tmp"
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}

	if err := os.Rename(tempPath, path); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("rename temp file: %w", err)
	}

	return nil
}

// LoadState loads the state from a file.
func LoadState(path string) (*RunState, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("state file not found: %w", err)
		}
		return nil, fmt.Errorf("read state file: %w", err)
	}

	var state RunState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, fmt.Errorf("unmarshal state: %w", err)
	}

	return &state, nil
}
```

internal/state/state_test.go
```
package state

import (
	"os"
	"testing"
)

func TestStatePersistence(t *testing.T) {
	tmpFile := "test_state.json"
	defer os.Remove(tmpFile)

	s := &RunState{
		SchemaVersion: 1,
		PackHash:      "abc",
		StepStatuses: map[string]StepStatus{
			"01": {Status: StatusSuccess, ExitCode: 0},
		},
	}

	if err := SaveStateAtomic(tmpFile, s); err != nil {
		t.Fatalf("SaveStateAtomic failed: %v", err)
	}

	loaded, err := LoadState(tmpFile)
	if err != nil {
		t.Fatalf("LoadState failed: %v", err)
	}

	if loaded.PackHash != s.PackHash {
		t.Errorf("expected hash %s, got %s", s.PackHash, loaded.PackHash)
	}

	if loaded.StepStatuses["01"].Status != StatusSuccess {
		t.Errorf("expected status success, got %s", loaded.StepStatuses["01"].Status)
	}
}
```

internal/state/types.go
```
package state

import (
	"time"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
	StatusSkipped Status = "skipped"
)

// RunState tracks the execution progress of an oracle pack.
type RunState struct {
	SchemaVersion int                   `json:"schema_version"`
	PackHash      string                `json:"pack_hash"`
	StartTime     time.Time             `json:"start_time"`
	StepStatuses  map[string]StepStatus `json:"step_statuses"`
	ROIThreshold  float64               `json:"roi_threshold,omitempty"`
	ROIMode       string                `json:"roi_mode,omitempty"`
	Warnings      []Warning             `json:"warnings,omitempty"`
}

// StepStatus holds the outcome of an individual step.
type StepStatus struct {
	Status    Status    `json:"status"`
	ExitCode  int       `json:"exit_code"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Error     string    `json:"error,omitempty"`
}

// Warning captures a non-fatal execution note (e.g., sanitized labels).
type Warning struct {
	Scope   string `json:"scope"`
	StepID  string `json:"step_id,omitempty"`
	Line    int    `json:"line"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
```

internal/tui/clipboard.go
```
package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func copyToClipboard(content string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else {
			return err
		}
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return exec.ErrNotFound
	}

	cmd.Stdin = strings.NewReader(content)
	return cmd.Run()
}

func writeClipboardFallback(content string) (string, error) {
	file, err := os.CreateTemp("", "oraclepack-step-*.txt")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return "", fmt.Errorf("write temp file: %w", err)
	}
	return file.Name(), nil
}
```

internal/tui/filter_test.go
```
package tui

import (
	"os"
	"path/filepath"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

func TestFilterLogic(t *testing.T) {
	// Setup pack with steps having different ROI
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
			{ID: "02", ROI: 5.0, OriginalLine: "Step 2"},
			{ID: "03", ROI: 10.0, OriginalLine: "Step 3"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Initialize model with no filter (threshold 0)
	m := NewModel(p, r, s, "", 0, "over", false)

	if len(m.list.Items()) != 3 {
		t.Fatalf("expected 3 items initially, got %d", len(m.list.Items()))
	}

	// Apply filter: ROI >= 5.0
	m.roiThreshold = 5.0
	m.roiMode = "over"
	m = m.refreshList()

	if len(m.list.Items()) != 2 {
		t.Errorf("expected 2 items after filtering >= 5.0, got %d", len(m.list.Items()))
	}

	// Verify items are 02 and 03
	items := m.list.Items()
	if items[0].(item).id != "02" {
		t.Errorf("expected first item to be 02, got %s", items[0].(item).id)
	}
	if items[1].(item).id != "03" {
		t.Errorf("expected second item to be 03, got %s", items[1].(item).id)
	}

	// Apply filter: ROI < 5.0 ("under")
	m.roiThreshold = 5.0
	m.roiMode = "under"
	m = m.refreshList()

	if len(m.list.Items()) != 1 {
		t.Errorf("expected 1 item after filtering < 5.0, got %d", len(m.list.Items()))
	}
	if m.list.Items()[0].(item).id != "01" {
		t.Errorf("expected item to be 01, got %s", m.list.Items()[0].(item).id)
	}
}

func TestROIModeTogglePersists(t *testing.T) {
	dir := t.TempDir()
	statePath := filepath.Join(dir, "state.json")
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", ROI: 1.0, OriginalLine: "Step 1"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{SchemaVersion: 1}

	m := NewModel(p, r, s, statePath, 0, "over", false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	m2 := updated.(Model)
	if m2.roiMode != "under" {
		t.Fatalf("expected roiMode to toggle to under, got %s", m2.roiMode)
	}

	loaded, err := state.LoadState(statePath)
	if err != nil {
		t.Fatalf("failed to load state: %v", err)
	}
	if loaded.ROIMode != "under" {
		t.Fatalf("expected persisted roiMode under, got %s", loaded.ROIMode)
	}

	if err := os.Remove(statePath); err != nil {
		t.Fatalf("failed to cleanup state file: %v", err)
	}
}
```

internal/tui/overrides_confirm.go
```
package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
)

type ValidationResultMsg struct {
	Errors []exec.ValidationError
	Err    error
}

type OverridesConfirmModel struct {
	validating bool
	errMsg     string
	errors     []exec.ValidationError
}

func (m OverridesConfirmModel) View(over overrides.RuntimeOverrides, baseline []string) string {
	added := strings.Join(over.AddedFlags, ", ")
	if added == "" {
		added = "(none)"
	}
	removed := strings.Join(over.RemovedFlags, ", ")
	if removed == "" {
		removed = "(none)"
	}
	targeted := len(over.ApplyToSteps)
	targetList := formatTargetList(over.ApplyToSteps, 5)
	effective := effectiveFlagsSummary(over, baseline)
	lines := []string{
		"Summary:",
		fmt.Sprintf("Added flags: %s", added),
		fmt.Sprintf("Removed flags: %s", removed),
		fmt.Sprintf("Targeted steps: %d%s", targeted, targetList),
		fmt.Sprintf("Effective flags: %s", effective),
		"",
		"[Enter] Validate  [Esc] Cancel",
	}

	if m.validating {
		lines = append(lines, "", "Validating overrides...")
	}
	if m.errMsg != "" {
		lines = append(lines, "", "Validation failed:", m.errMsg)
	}
	if len(m.errors) > 0 {
		lines = append(lines, "", fmt.Sprintf("Validation errors (%d):", len(m.errors)))
		lines = append(lines, formatValidationErrors(m.errors, 6)...)
	}

	return strings.Join(lines, "\n")
}

func formatTargetList(targets map[string]bool, limit int) string {
	if len(targets) == 0 || limit <= 0 {
		return ""
	}
	ids := make([]string, 0, len(targets))
	for id := range targets {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	if len(ids) <= limit {
		return fmt.Sprintf(" (%s)", strings.Join(ids, ", "))
	}
	return fmt.Sprintf(" (%s, +%d more)", strings.Join(ids[:limit], ", "), len(ids)-limit)
}

func effectiveFlagsSummary(over overrides.RuntimeOverrides, baseline []string) string {
	if len(over.ApplyToSteps) == 0 {
		return "(no steps targeted)"
	}
	var first string
	for id := range over.ApplyToSteps {
		first = id
		break
	}
	flags := over.EffectiveFlags(first, baseline)
	if len(flags) == 0 {
		return "(none)"
	}
	return strings.Join(flags, " ")
}

func formatValidationErrors(errors []exec.ValidationError, limit int) []string {
	if limit <= 0 {
		return nil
	}
	lines := []string{}
	for i, err := range errors {
		if i >= limit {
			lines = append(lines, fmt.Sprintf("- (+%d more)", len(errors)-limit))
			break
		}
		msg := strings.TrimSpace(err.ErrorMessage)
		if msg == "" {
			msg = "(no error message)"
		}
		lines = append(lines, fmt.Sprintf("- Step %s: %s", err.StepID, firstLine(msg)))
	}
	return lines
}

func firstLine(msg string) string {
	if idx := strings.IndexByte(msg, '\n'); idx != -1 {
		return msg[:idx]
	}
	return msg
}
```

internal/tui/overrides_flags.go
```
package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FlagItem struct {
	Flag       string
	Desc       string
	IsBaseline bool
	Selected   bool
}

func (i FlagItem) Title() string       { return i.Flag }
func (i FlagItem) Description() string { return i.Desc }
func (i FlagItem) FilterValue() string { return i.Flag }

type FlagsPickerModel struct {
	list list.Model
}

func NewFlagsPickerModel(baseline []string) FlagsPickerModel {
	baselineSet := make(map[string]bool, len(baseline))
	for _, f := range baseline {
		baselineSet[f] = true
	}

	curated := []FlagItem{
		{Flag: "--files-report", Desc: "Show per-file token usage"},
		{Flag: "--render", Desc: "Print assembled markdown bundle"},
		{Flag: "--render-plain", Desc: "Render markdown without ANSI"},
		{Flag: "--copy", Desc: "Copy assembled markdown bundle"},
		{Flag: "--wait", Desc: "Wait for background API runs"},
	}

	items := make([]list.Item, 0, len(curated))
	for _, c := range curated {
		c.IsBaseline = baselineSet[c.Flag]
		if c.IsBaseline {
			c.Selected = true
		}
		items = append(items, c)
	}

	delegate := newFlagsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Oracle Flags"
	l.SetFilteringEnabled(true)

	return FlagsPickerModel{list: l}
}

func (m FlagsPickerModel) Init() tea.Cmd {
	return nil
}

func (m FlagsPickerModel) Update(msg tea.Msg) (FlagsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(FlagItem)
			if ok && !item.IsBaseline {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *FlagsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m FlagsPickerModel) View() string {
	return m.list.View()
}

func (m FlagsPickerModel) SelectedFlags() []string {
	var flags []string
	for _, item := range m.list.Items() {
		if fi, ok := item.(FlagItem); ok && fi.Selected && !fi.IsBaseline {
			flags = append(flags, fi.Flag)
		}
	}
	return flags
}

type flagsDelegate struct {
	list.DefaultDelegate
}

func newFlagsDelegate() flagsDelegate {
	d := list.NewDefaultDelegate()
	return flagsDelegate{DefaultDelegate: d}
}

func (d flagsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	fi, ok := item.(FlagItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	checked := fi.Selected || fi.IsBaseline
	marker := "[ ]"
	if checked {
		marker = "[x]"
	}
	if fi.IsBaseline {
		marker = "[*]"
	}

	label := fi.Flag
	if fi.Desc != "" {
		label = fmt.Sprintf("%s - %s", fi.Flag, fi.Desc)
	}
	if fi.IsBaseline {
		label = label + " (base)"
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}
	if fi.IsBaseline {
		line = lipgloss.NewStyle().Faint(true).Render(line)
	}

	fmt.Fprintln(w, line)
}
```

internal/tui/overrides_flow.go
```
package tui

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
)

type OverridesStep int

const (
	OverridesFlags OverridesStep = iota
	OverridesSteps
	OverridesConfirm
)

type OverridesStartedMsg struct{}

type OverridesAppliedMsg struct {
	Overrides overrides.RuntimeOverrides
}

type OverridesCancelledMsg struct{}

type OverridesFlowModel struct {
	step    OverridesStep
	flags   FlagsPickerModel
	steps   StepsPickerModel
	confirm OverridesConfirmModel

	packSteps        []pack.Step
	baseline         []string
	runnerOpts       exec.RunnerOptions
	pendingOverrides overrides.RuntimeOverrides
}

func NewOverridesFlowModel(steps []pack.Step, baseline []string, opts exec.RunnerOptions) OverridesFlowModel {
	return OverridesFlowModel{
		step:       OverridesFlags,
		flags:      NewFlagsPickerModel(nil),
		steps:      NewStepsPickerModel(steps),
		confirm:    OverridesConfirmModel{},
		packSteps:  steps,
		baseline:   exec.ApplyChatGPTURL(baseline, opts.ChatGPTURL),
		runnerOpts: opts,
	}
}

func (m OverridesFlowModel) Init() tea.Cmd {
	return nil
}

func (m OverridesFlowModel) Update(msg tea.Msg) (OverridesFlowModel, tea.Cmd) {
	var cmd tea.Cmd
	if m.step == OverridesFlags {
		m.flags, cmd = m.flags.Update(msg)
	}
	if m.step == OverridesSteps {
		m.steps, cmd = m.steps.Update(msg)
	}
	if m.step == OverridesConfirm {
		switch v := msg.(type) {
		case ValidationResultMsg:
			m.confirm.validating = false
			m.confirm.errors = v.Errors
			if v.Err != nil {
				m.confirm.errMsg = v.Err.Error()
				return m, nil
			}
			if len(v.Errors) > 0 {
				m.confirm.errMsg = fmt.Sprintf("%d validation errors detected.", len(v.Errors))
				return m, nil
			}
			return m, func() tea.Msg { return OverridesAppliedMsg{Overrides: m.pendingOverrides} }
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return OverridesCancelledMsg{} }
		case "shift+tab", "backspace":
			if m.step > OverridesFlags {
				m.step--
			}
		case "enter", "tab":
			if m.step == OverridesConfirm {
				if m.confirm.validating {
					return m, nil
				}
				m.pendingOverrides = m.currentOverrides()
				m.confirm.validating = true
				m.confirm.errMsg = ""
				m.confirm.errors = nil
				return m, m.validateCmd(m.pendingOverrides)
			}
			m.step++
		}
	}

	return m, cmd
}

func (m OverridesFlowModel) View(width, height int) string {
	title := lipgloss.NewStyle().Bold(true).Render("Overrides Wizard")
	step := fmt.Sprintf("Step %d/3", int(m.step)+1)
	body := fmt.Sprintf("Current step: %s\n\n[Enter] Next  [Esc] Cancel", overridesStepName(m.step))

	var content string
	if m.step == OverridesFlags {
		m.flags.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.flags.View(),
			"",
			body,
		)
	} else if m.step == OverridesSteps {
		m.steps.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.steps.View(),
			"",
			body,
		)
	} else if m.step == OverridesConfirm {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.confirm.View(m.currentOverrides(), m.baseline),
		)
	} else {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			body,
		)
	}

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}

func (m OverridesFlowModel) currentOverrides() overrides.RuntimeOverrides {
	return overrides.RuntimeOverrides{
		AddedFlags:   m.flags.SelectedFlags(),
		RemovedFlags: nil,
		ApplyToSteps: m.steps.SelectedSteps(),
	}
}

func (m OverridesFlowModel) validateCmd(over overrides.RuntimeOverrides) tea.Cmd {
	return func() tea.Msg {
		errs, err := exec.ValidateOverrides(context.Background(), m.packSteps, &over, m.baseline, m.runnerOpts)
		return ValidationResultMsg{Errors: errs, Err: err}
	}
}

func overridesStepName(step OverridesStep) string {
	switch step {
	case OverridesFlags:
		return "Flags"
	case OverridesSteps:
		return "Target Steps"
	case OverridesConfirm:
		return "Confirm"
	default:
		return "Unknown"
	}
}
```

internal/tui/overrides_steps.go
```
package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/pack"
)

type StepItem struct {
	ID       string
	TitleTxt string
	DescTxt  string
	Selected bool
}

func (i StepItem) Title() string       { return i.TitleTxt }
func (i StepItem) Description() string { return i.DescTxt }
func (i StepItem) FilterValue() string { return i.TitleTxt }

type StepsPickerModel struct {
	list list.Model
}

func NewStepsPickerModel(steps []pack.Step) StepsPickerModel {
	items := make([]list.Item, 0, len(steps))
	for _, s := range steps {
		items = append(items, StepItem{
			ID:       s.ID,
			TitleTxt: fmt.Sprintf("Step %s", s.ID),
			DescTxt:  s.OriginalLine,
			Selected: true,
		})
	}

	delegate := newStepsDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Target Steps"
	l.SetFilteringEnabled(true)

	return StepsPickerModel{list: l}
}

func (m StepsPickerModel) Init() tea.Cmd {
	return nil
}

func (m StepsPickerModel) Update(msg tea.Msg) (StepsPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a", "A":
			m = m.setAll(true)
			return m, nil
		case "i":
			m = m.invert()
			return m, nil
		case "n":
			m = m.setAll(false)
			return m, nil
		case " ":
			idx := m.list.Index()
			item, ok := m.list.SelectedItem().(StepItem)
			if ok {
				item.Selected = !item.Selected
				_ = m.list.SetItem(idx, item)
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *StepsPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height)
}

func (m StepsPickerModel) View() string {
	help := lipgloss.NewStyle().Faint(true).Render("[space] toggle  [a] all  [i] invert  [n] none")
	return m.list.View() + "\n" + help
}

func (m StepsPickerModel) SelectedSteps() map[string]bool {
	selected := make(map[string]bool)
	for _, item := range m.list.Items() {
		if si, ok := item.(StepItem); ok && si.Selected {
			selected[si.ID] = true
		}
	}
	return selected
}

func (m StepsPickerModel) setAll(value bool) StepsPickerModel {
	for idx, item := range m.list.Items() {
		si, ok := item.(StepItem)
		if !ok {
			continue
		}
		si.Selected = value
		_ = m.list.SetItem(idx, si)
	}
	return m
}

func (m StepsPickerModel) invert() StepsPickerModel {
	for idx, item := range m.list.Items() {
		si, ok := item.(StepItem)
		if !ok {
			continue
		}
		si.Selected = !si.Selected
		_ = m.list.SetItem(idx, si)
	}
	return m
}

type stepsDelegate struct {
	list.DefaultDelegate
}

func newStepsDelegate() stepsDelegate {
	d := list.NewDefaultDelegate()
	return stepsDelegate{DefaultDelegate: d}
}

func (d stepsDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	si, ok := item.(StepItem)
	if !ok {
		d.DefaultDelegate.Render(w, m, index, item)
		return
	}

	marker := "[ ]"
	if si.Selected {
		marker = "[x]"
	}

	label := si.TitleTxt
	if si.DescTxt != "" {
		label = fmt.Sprintf("%s - %s", si.TitleTxt, si.DescTxt)
	}

	line := fmt.Sprintf("%s %s", marker, label)
	if index == m.Index() {
		line = d.Styles.SelectedTitle.Render(line)
	} else {
		line = d.Styles.NormalTitle.Render(line)
	}

	fmt.Fprintln(w, line)
}
```

internal/tui/overrides_url.go
```
package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLInputModel struct {
	input textinput.Model
	err   string
}

func NewURLInputModel() URLInputModel {
	ti := textinput.New()
	ti.Placeholder = "https://chat.openai.com/project/..."
	ti.CharLimit = 200
	ti.Width = 50

	return URLInputModel{input: ti}
}

func (m URLInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m URLInputModel) Update(msg tea.Msg) (URLInputModel, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.err = ""
	if !m.IsValid() {
		m.err = "Invalid URL (must start with http:// or https://)"
	}
	return m, cmd
}

func (m URLInputModel) Value() string {
	return strings.TrimSpace(m.input.Value())
}

func (m URLInputModel) IsValid() bool {
	v := m.Value()
	if v == "" {
		return true
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}

func (m URLInputModel) View() string {
	body := m.input.View()
	if m.err != "" {
		body = body + "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.err)
	}
	return body
}

func (m *URLInputModel) SetValue(v string) {
	m.input.SetValue(v)
}

func (m *URLInputModel) Focus() {
	m.input.Focus()
}

func (m *URLInputModel) Blur() {
	m.input.Blur()
}
```

internal/tui/preview_test.go
```
package tui

import (
	"strings"
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

func TestStepPreviewContentUnwrapped(t *testing.T) {
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", OriginalLine: "Step 1", Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}
	m := NewModel(p, r, s, "", 0, "over", false)
	m.width = 80
	m.previewID = "01"
	m.previewWrap = false

	content := m.stepPreviewContent()
	if !strings.Contains(content, "Step 01") {
		t.Fatalf("expected header to include step id, got %q", content)
	}
	if !strings.Contains(content, "echo hello") {
		t.Fatalf("expected content to include code, got %q", content)
	}
}
```

internal/tui/tui.go
```
package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/render"
	"github.com/user/oraclepack/internal/state"
)

type ViewState int

const (
	ViewSteps ViewState = iota
	ViewRunning
	ViewDone
	ViewOverrides
	ViewStepPreview
)

type item struct {
	id     string
	title  string
	desc   string
	status state.Status
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list        list.Model
	viewport    viewport.Model
	spinner     spinner.Model
	filterInput textinput.Model
	urlInput    URLInputModel
	urlPicker   URLPickerModel
	pack        *pack.Pack
	runner      *exec.Runner
	state       *state.RunState
	statePath   string

	width  int
	height int

	viewState     ViewState
	running       bool
	runAll        bool // State for sequential execution
	currentIdx    int
	autoRun       bool // Config to auto-start on init
	previewID     string
	previewWrap   bool
	previewNotice string

	// Filtering state
	allSteps     []item // Store all items to support dynamic filtering
	roiThreshold float64
	roiMode      string
	isFiltering  bool
	isEditingURL bool
	isPickingURL bool

	overridesFlow    OverridesFlowModel
	appliedOverrides *overrides.RuntimeOverrides
	chatGPTURL       string

	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, statePath string, roiThreshold float64, roiMode string, autoRun bool) Model {
	if s != nil {
		if s.ROIThreshold > 0 {
			roiThreshold = s.ROIThreshold
		}
		if s.ROIMode != "" {
			roiMode = s.ROIMode
		}
	}
	var allItems []item
	for _, step := range p.Steps {
		allItems = append(allItems, item{
			id:    step.ID,
			title: fmt.Sprintf("Step %s", step.ID),
			desc:  step.OriginalLine,
		})
	}

	ti := textinput.New()
	ti.Placeholder = "Enter ROI threshold (e.g. 2.5)"
	ti.CharLimit = 10
	ti.Width = 20

	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Oracle Pack Steps"

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	vp := viewport.New(0, 0)
	vp.SetContent("Press Enter to run selected, 'a' to run all filtered steps, 'f' to set ROI threshold, 'm' to toggle ROI mode, 'v' to view step, 'o' to configure overrides, 'u' for ChatGPT URL, 'U' to pick a saved URL.")

	projectPath := ProjectURLStorePath(statePath, p.Source)
	globalPath := GlobalURLStorePath()
	urlPicker := NewURLPickerModel(projectPath, globalPath)
	resolvedURL := r.ChatGPTURL
	if resolvedURL == "" {
		resolvedURL = urlPicker.DefaultURL()
	}
	if resolvedURL != "" {
		r.ChatGPTURL = resolvedURL
	}

	m := Model{
		list:          l,
		viewport:      vp,
		spinner:       sp,
		filterInput:   ti,
		urlInput:      NewURLInputModel(),
		urlPicker:     urlPicker,
		pack:          p,
		runner:        r,
		state:         s,
		statePath:     statePath,
		autoRun:       autoRun,
		allSteps:      allItems,
		roiThreshold:  roiThreshold,
		roiMode:       roiMode,
		logChan:       make(chan string, 100),
		viewState:     ViewSteps,
		overridesFlow: NewOverridesFlowModel(p.Steps, r.OracleFlags, RunnerOptionsFromRunner(r)),
		chatGPTURL:    resolvedURL,
		previewWrap:   true,
	}
	m.urlInput.SetValue(resolvedURL)
	m.urlInput.Blur()

	// Apply initial filter
	return m.refreshList()
}

func (m Model) refreshList() Model {
	var filtered []list.Item
	for _, it := range m.allSteps {
		// Find the original step to check ROI
		var step *pack.Step
		for _, s := range m.pack.Steps {
			if s.ID == it.id {
				step = &s
				break
			}
		}
		if step == nil {
			continue
		}

		if m.roiThreshold > 0 {
			if m.roiMode == "under" {
				if step.ROI >= m.roiThreshold {
					continue
				}
			} else {
				if step.ROI < m.roiThreshold {
					continue
				}
			}
		}
		filtered = append(filtered, it)
	}
	m.list.SetItems(filtered)
	return m
}

type StartAutoRunMsg struct{}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	if m.autoRun {
		cmds = append(cmds, func() tea.Msg { return StartAutoRunMsg{} })
	}
	cmds = append(cmds, textinput.Blink)
	return tea.Batch(cmds...)
}

type LogMsg string
type FinishedMsg struct {
	Err error
	ID  string
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Global keys (Quit)
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	if msg, ok := msg.(OverridesStartedMsg); ok {
		_ = msg
		m.viewState = ViewOverrides
		m.overridesFlow = NewOverridesFlowModel(m.pack.Steps, m.runner.OracleFlags, RunnerOptionsFromRunner(m.runner))
		return m, nil
	}
	if msg, ok := msg.(OverridesAppliedMsg); ok {
		over := msg.Overrides
		m.appliedOverrides = &over
		if m.runner != nil {
			m.runner.Overrides = &over
		}
		m.viewState = ViewSteps
		return m, nil
	}
	if msg, ok := msg.(OverridesCancelledMsg); ok {
		_ = msg
		m.appliedOverrides = nil
		if m.runner != nil {
			m.runner.Overrides = nil
		}
		m.viewState = ViewSteps
		return m, nil
	}
	if msg, ok := msg.(URLPickedMsg); ok {
		m.chatGPTURL = msg.URL
		if m.runner != nil {
			m.runner.ChatGPTURL = m.chatGPTURL
		}
		m.urlInput.SetValue(m.chatGPTURL)
		m.isPickingURL = false
		return m, nil
	}
	if _, ok := msg.(URLPickerCancelledMsg); ok {
		m.isPickingURL = false
		return m, nil
	}

	if m.viewState == ViewOverrides {
		var cmd tea.Cmd
		m.overridesFlow, cmd = m.overridesFlow.Update(msg)
		return m, cmd
	}

	if m.viewState == ViewStepPreview {
		switch msg := msg.(type) {
		case clearPreviewNoticeMsg:
			m.previewNotice = ""
			return m, nil
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "b", "esc":
				m.previewID = ""
				m.previewNotice = ""
				m.viewState = ViewSteps
				m.setListPreviewContent(m.selectedItemID())
				return m, nil
			case "t":
				m.previewWrap = !m.previewWrap
				m.viewport.SetContent(m.stepPreviewContent())
				return m, nil
			case "c":
				content := m.stepPlainTextFor(m.previewID)
				if err := copyToClipboard(content); err != nil {
					path, fallbackErr := writeClipboardFallback(content)
					if fallbackErr != nil {
						m.previewNotice = "Copy failed: " + err.Error()
					} else {
						m.previewNotice = "Copy failed; saved to " + path
					}
				} else {
					m.previewNotice = "Copied to clipboard"
				}
				return m, tea.Tick(2*time.Second, func(time.Time) tea.Msg {
					return clearPreviewNoticeMsg{}
				})
			}
		}
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	switch m.viewState {
	case ViewDone:
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "b":
				m.viewState = ViewSteps
				m.setListPreviewContent(m.selectedItemID())
				return m, nil
			case "n":
				m.resetState()
				return m, nil
			case "r":
				// Rerun selected step (if we have one selected in list)
				// Or rerun the whole sequence if that was the context?
				// Requirement says "rerun a step ('r')". Assuming selected step.
				// We need to transition to ViewSteps logic or trigger run directly.
				m.viewState = ViewSteps // Go back to steps view? Or Running?
				// To trigger run, we can fall through or simulate Enter.
				// Let's just switch to steps and let user press Enter, or trigger run immediately?
				// "trigger the execution logic for the specific failed step"
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.running = true
					m.viewState = ViewRunning
					m.logLines = nil
					m.viewport.SetContent("Re-running execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
			}
		}
		// In Done view, we might still want to handle window size?
		if msg, ok := msg.(tea.WindowSizeMsg); ok {
			m.handleWindowSize(msg)
		}
		return m, nil
	}

	// Filter Input Mode
	if m.isFiltering {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				var val float64
				_, err := fmt.Sscanf(m.filterInput.Value(), "%f", &val)
				if err == nil {
					m.roiThreshold = val
					m = m.refreshList()
					m.persistFilterState()
				}
				m.isFiltering = false
				m.filterInput.Blur()
				return m, nil
			case "esc":
				m.isFiltering = false
				m.filterInput.Blur()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.filterInput, cmd = m.filterInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	// ChatGPT URL Input Mode
	if m.isEditingURL {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.urlInput.IsValid() {
					m.chatGPTURL = m.urlInput.Value()
					if m.runner != nil {
						m.runner.ChatGPTURL = m.chatGPTURL
					}
					m.isEditingURL = false
					m.urlInput.Blur()
					return m, nil
				}
			case "esc":
				m.isEditingURL = false
				m.urlInput.Blur()
				return m, nil
			}
		}
		var cmd tea.Cmd
		m.urlInput, cmd = m.urlInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	// URL Picker Mode
	if m.isPickingURL {
		var cmd tea.Cmd
		m.urlPicker, cmd = m.urlPicker.Update(msg)
		return m, cmd
	}

	// Normal Steps View / Running
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if !m.running {
				return m, tea.Quit
			}
		case "enter":
			if !m.running {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.running = true
					m.viewState = ViewRunning
					m.runAll = false
					m.logLines = nil
					m.viewport.SetContent("Starting execution...")
					return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
				}
			}
		case "a":
			if !m.running && len(m.list.Items()) > 0 {
				m.running = true
				m.viewState = ViewRunning
				m.runAll = true
				m.currentIdx = 0
				m.logLines = nil
				m.list.Select(0)
				i := m.list.Items()[0].(item)
				m.viewport.SetContent(fmt.Sprintf("Starting sequential run (Step 1/%d)...", len(m.list.Items())))
				return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
			}
		case "f":
			if !m.running {
				m.isFiltering = true
				m.filterInput.Focus()
				m.filterInput.SetValue(fmt.Sprintf("%.1f", m.roiThreshold))
				return m, textinput.Blink
			}
		case "m":
			if !m.running {
				if m.roiMode == "under" {
					m.roiMode = "over"
				} else {
					m.roiMode = "under"
				}
				m = m.refreshList()
				m.persistFilterState()
				return m, nil
			}
		case "v":
			if !m.running {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.previewID = i.id
					m.viewState = ViewStepPreview
					m.viewport.YOffset = 0
					m.viewport.SetContent(m.stepPreviewContent())
					return m, nil
				}
			}
		case "u":
			if !m.running {
				m.isEditingURL = true
				m.urlInput.SetValue(m.chatGPTURL)
				m.urlInput.Focus()
				return m, textinput.Blink
			}
		case "U":
			if !m.running {
				m.isPickingURL = true
				return m, nil
			}
		case "o":
			if !m.running {
				return m, func() tea.Msg { return OverridesStartedMsg{} }
			}
		}

	case StartAutoRunMsg:
		if !m.running && len(m.list.Items()) > 0 {
			m.running = true
			m.viewState = ViewRunning
			m.runAll = true
			m.currentIdx = 0
			m.logLines = nil
			m.list.Select(0)
			i := m.list.Items()[0].(item)
			m.viewport.SetContent(fmt.Sprintf("Auto-running all steps (Step 1/%d)...", len(m.list.Items())))
			return m, tea.Batch(m.runStep(i.id), m.waitForLogs(), m.spinner.Tick)
		}

	case tea.WindowSizeMsg:
		m.handleWindowSize(msg)

	case LogMsg:
		m.logLines = append(m.logLines, string(msg))
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()
		return m, m.waitForLogs()

	case FinishedMsg:
		m.recordWarnings()
		if msg.Err != nil {
			m.err = msg.Err
			m.logLines = append(m.logLines, fmt.Sprintf("\n❌ ERROR: %v", msg.Err))
			m.running = false
			m.runAll = false
			m.viewState = ViewDone // Or stay in steps? Requirement says ViewDone on completion?
			// If error, maybe stay on steps or go to done with error?
			// "Failed at step X" is a summary state.
			m.viewState = ViewDone
		} else {
			m.logLines = append(m.logLines, "\n✅ SUCCESS")

			if m.runAll {
				m.currentIdx++
				if m.currentIdx < len(m.list.Items()) {
					m.list.Select(m.currentIdx)
					i := m.list.Items()[m.currentIdx].(item)
					m.logLines = append(m.logLines, fmt.Sprintf("\n--- Starting Step %d/%d ---\n", m.currentIdx+1, len(m.list.Items())))
					return m, m.runStep(i.id)
				} else {
					m.logLines = append(m.logLines, "\n🏁 ALL STEPS COMPLETED")
					m.running = false
					m.runAll = false
					m.viewState = ViewDone
				}
			} else {
				m.running = false
				m.viewState = ViewDone // Single step done
			}
		}
		m.viewport.SetContent(strings.Join(m.logLines, "\n"))
		m.viewport.GotoBottom()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if !m.running && !m.isFiltering && m.viewState == ViewSteps {
		prevID := m.selectedItemID()
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
		newID := m.selectedItemID()
		if newID != "" && newID != prevID {
			m.viewport.YOffset = 0
			m.setListPreviewContent(newID)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) handleWindowSize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
	if m.viewState == ViewStepPreview {
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 6
		if m.viewport.Height < 1 {
			m.viewport.Height = 1
		}
		m.viewport.SetContent(m.stepPreviewContent())
		m.viewport.GotoTop()
		return
	}
	contentHeight := msg.Height - 5
	if contentHeight < 1 {
		contentHeight = 1
	}
	m.list.SetSize(msg.Width/3, contentHeight)
	m.viewport.Width = msg.Width - (msg.Width / 3) - 6
	m.viewport.Height = contentHeight
	if !m.running && m.viewState == ViewSteps {
		m.setListPreviewContent(m.selectedItemID())
	}
}

func (m *Model) resetState() {
	// Reset RunState
	m.state.StartTime = time.Now()
	m.state.StepStatuses = make(map[string]state.StepStatus)

	// Save cleared state to disk
	if m.statePath != "" {
		_ = state.SaveStateAtomic(m.statePath, m.state)
	}

	// Reset UI
	m.logLines = nil
	m.viewport.SetContent("State reset. Ready for new run.")
	m.list.Select(0)
	m.viewState = ViewSteps
	m.running = false
	m.runAll = false
	m.appliedOverrides = nil
	if m.runner != nil {
		m.runner.Overrides = nil
	}
}

func (m *Model) persistFilterState() {
	if m.state == nil || m.statePath == "" {
		return
	}
	m.state.ROIThreshold = m.roiThreshold
	m.state.ROIMode = m.roiMode
	_ = state.SaveStateAtomic(m.statePath, m.state)
}

func (m *Model) recordWarnings() {
	if m.state == nil || m.statePath == "" || m.runner == nil {
		return
	}
	warnings := m.runner.DrainWarnings()
	if len(warnings) == 0 {
		return
	}
	for _, w := range warnings {
		m.state.Warnings = append(m.state.Warnings, state.Warning{
			Scope:   w.Scope,
			StepID:  w.StepID,
			Line:    w.Line,
			Token:   w.Token,
			Message: w.Message,
		})
	}
	_ = state.SaveStateAtomic(m.statePath, m.state)
}

func (m *Model) setLogContent() {
	if len(m.logLines) == 0 {
		return
	}
	m.viewport.SetContent(strings.Join(m.logLines, "\n"))
	m.viewport.GotoBottom()
}

func (m *Model) stepPreviewContent() string {
	return m.stepPreviewContentFor(m.previewID)
}

func (m *Model) stepPreviewContentFor(id string) string {
	md, ok := m.stepMarkdownFor(id)
	if !ok {
		return md
	}
	width := m.previewRenderWidth()
	rendered, err := render.RenderMarkdown(md, width, "auto")
	if err != nil {
		return m.stepPlainTextFor(id)
	}
	return rendered
}

func (m *Model) stepMarkdownFor(id string) (string, bool) {
	if id == "" {
		return "No step selected.", false
	}
	step := m.stepForID(id)
	if step == nil {
		return "Step not found.", false
	}
	header := fmt.Sprintf("## Step %s\n%s\n\n", step.ID, step.OriginalLine)
	md := header + "```bash\n" + step.Code + "\n```\n"
	return md, true
}

func (m *Model) stepPlainTextFor(id string) string {
	if id == "" {
		return "No step selected."
	}
	step := m.stepForID(id)
	if step == nil {
		return "Step not found."
	}
	header := fmt.Sprintf("Step %s\n%s\n", step.ID, step.OriginalLine)
	return header + "\n" + step.Code
}

func (m *Model) stepForID(id string) *pack.Step {
	for i := range m.pack.Steps {
		if m.pack.Steps[i].ID == id {
			return &m.pack.Steps[i]
		}
	}
	return nil
}

func (m *Model) previewRenderWidth() int {
	width := m.viewport.Width
	if width <= 0 {
		width = render.DefaultWidth
	}
	if !m.previewWrap {
		if width < render.DefaultWidth {
			width = render.DefaultWidth
		}
		width = width * 4
	}
	return width
}

func (m *Model) selectedItemID() string {
	it, ok := m.list.SelectedItem().(item)
	if !ok {
		return ""
	}
	return it.id
}

func (m *Model) setListPreviewContent(id string) {
	if id == "" {
		m.viewport.SetContent("No step selected.")
		return
	}
	m.viewport.SetContent(m.stepPreviewContentFor(id))
	m.viewport.GotoTop()
}

type clearPreviewNoticeMsg struct{}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}

	if m.viewState == ViewDone {
		return m.viewDone()
	}

	if m.viewState == ViewOverrides {
		return m.overridesFlow.View(m.width, m.height)
	}

	if m.isFiltering {
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center,
				"Enter ROI Threshold:",
				m.filterInput.View(),
				"(Enter to apply, Esc to cancel)",
			),
		)
	}

	if m.isEditingURL {
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center,
				"ChatGPT URL (browser mode):",
				m.urlInput.View(),
				"(Enter to apply, Esc to cancel)",
			),
		)
	}

	if m.isPickingURL {
		m.urlPicker.SetSize(m.width-4, m.height-4)
		return lipgloss.Place(m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			m.urlPicker.View(),
		)
	}

	if m.viewState == ViewStepPreview {
		m.viewport.Width = m.width - 4
		m.viewport.Height = m.height - 6
		title := lipgloss.NewStyle().Bold(true).Render("Step Preview")
		help := "[b] Back  [q] Quit  [t] Wrap  [c] Copy  (scroll with ↑↓ / PgUp/PgDn)"
		notice := ""
		if m.previewNotice != "" {
			notice = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Render(m.previewNotice)
		}
		content := lipgloss.JoinVertical(lipgloss.Left,
			title,
			help,
			notice,
			"",
			m.viewport.View(),
		)
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
	}

	left := m.list.View()
	right := m.viewport.View()

	if m.running {
		status := "Running..."
		if m.runAll {
			status = fmt.Sprintf("Running All (%d/%d)...", m.currentIdx+1, len(m.list.Items()))
		}
		right = m.spinner.View() + " " + status + "\n" + right
	} else {
		filterStatus := ""
		if m.roiThreshold > 0 {
			modeSym := ">="
			if m.roiMode == "under" {
				modeSym = "<"
			}
			filterStatus = fmt.Sprintf(" [Filter: ROI %s %.1f]", modeSym, m.roiThreshold)
		}
		if filterStatus == "" {
			modeSym := ">="
			if m.roiMode == "under" {
				modeSym = "<"
			}
			filterStatus = fmt.Sprintf(" [Filter: ROI %s ∞]", modeSym)
		}
		overrideStatus := ""
		if m.appliedOverrides != nil {
			added := len(m.appliedOverrides.AddedFlags)
			removed := len(m.appliedOverrides.RemovedFlags)
			targeted := len(m.appliedOverrides.ApplyToSteps)
			overrideStatus = fmt.Sprintf(" [Overrides: +%d -%d steps:%d]", added, removed, targeted)
		}
		urlStatus := ""
		if m.chatGPTURL != "" {
			urlStatus = " [ChatGPT URL: set]"
		} else {
			urlStatus = " [ChatGPT URL: none]"
		}
		statusLine := strings.TrimSpace(filterStatus + overrideStatus + urlStatus)
		if statusLine != "" {
			right = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render(statusLine) + "\n" + right
		}
	}

	help := m.stepsHelpBar(m.width)
	rightWidth := m.viewport.Width
	if rightWidth < 1 {
		rightWidth = 1
	}
	right = lipgloss.NewStyle().Width(rightWidth).Render(right)
	main := lipgloss.JoinHorizontal(lipgloss.Top, left, " | ", right)
	return lipgloss.JoinVertical(lipgloss.Left, main, help)
}

func (m Model) viewDone() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).Render("Execution Complete")
	if m.err != nil {
		title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196")).Render("Execution Failed")
	}

	help := "[n] New Run  [r] Rerun  [b] Back to List  [q] Quit  [m] ROI Mode"

	// Show the log viewport in the done screen too? Or just a summary?
	// Requirement says "displays a summary".
	// But viewing the logs is useful.
	// I'll show the viewport in the center/bottom.

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		"",
		m.viewport.View(),
		"",
		help,
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

func (m Model) stepsHelpBar(width int) string {
	help := "[enter] run  [a] run all  [f] filter ROI  [m] ROI mode  [v] view  [o] overrides  [u] url  [U] url picker  [q] quit"
	if m.running {
		help = "[q] quit  [running] wait for completion"
	}
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	if width > 0 {
		style = style.Width(width)
	}
	return style.Render(help)
}

func (m Model) waitForLogs() tea.Cmd {
	return func() tea.Msg {
		line, ok := <-m.logChan
		if !ok {
			return nil
		}
		return LogMsg(line)
	}
}

func (m Model) runStep(id string) tea.Cmd {
	return func() tea.Msg {
		var step *pack.Step
		for _, s := range m.pack.Steps {
			if s.ID == id {
				step = &s
				break
			}
		}

		if step == nil {
			return FinishedMsg{Err: fmt.Errorf("step not found"), ID: id}
		}

		lw := &exec.LineWriter{
			Callback: func(line string) {
				m.logChan <- line
			},
		}

		err := m.runner.RunStep(context.Background(), step, lw)
		lw.Close()
		return FinishedMsg{Err: err, ID: id}
	}
}

func RunnerOptionsFromRunner(r *exec.Runner) exec.RunnerOptions {
	if r == nil {
		return exec.RunnerOptions{}
	}
	return exec.RunnerOptions{
		Shell:       r.Shell,
		WorkDir:     r.WorkDir,
		Env:         r.Env,
		OracleFlags: r.OracleFlags,
		Overrides:   r.Overrides,
		ChatGPTURL:  r.ChatGPTURL,
	}
}
```

internal/tui/tui_test.go
```
package tui

import (
	"testing"

	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

func TestInitAutoRun(t *testing.T) {
	p := &pack.Pack{
		Steps: []pack.Step{
			{ID: "01", Number: 1, Code: "echo hello"},
		},
	}
	r := exec.NewRunner(exec.RunnerOptions{})
	s := &state.RunState{}

	// Test case 1: autoRun = true
	modelAuto := NewModel(p, r, s, "", 0, "over", true)
	cmdAuto := modelAuto.Init()
	
	if cmdAuto == nil {
		t.Fatal("expected Init cmd to be non-nil when autoRun is true")
	}
	// Note: We can't easily assert the content of a Batch command in a unit test.

	// Test case 2: autoRun = false
	modelManual := NewModel(p, r, s, "", 0, "over", false)
	// Even with autoRun false, we have textinput.Blink, so Init is not nil.
	cmdManual := modelManual.Init()
	if cmdManual == nil {
		t.Fatal("expected Init cmd to be non-nil due to textinput.Blink")
	}
}
```

internal/tui/url_picker.go
```
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type URLPickedMsg struct {
	URL string
}

type URLPickerCancelledMsg struct{}

type urlItem struct {
	name  string
	url   string
	scope string
}

func (i urlItem) Title() string       { return i.name }
func (i urlItem) Description() string { return fmt.Sprintf("%s • %s", i.scope, i.url) }
func (i urlItem) FilterValue() string { return i.name }

type URLPickerModel struct {
	list list.Model

	projectPath string
	globalPath  string
	project     URLStore
	global      URLStore

	editing   bool
	editName  textinput.Model
	editURL   textinput.Model
	editScope string
	editIdx   int
	editIsNew bool

	errMsg string
}

func NewURLPickerModel(projectPath, globalPath string) URLPickerModel {
	project, _ := LoadURLStore(projectPath)
	global, _ := LoadURLStore(globalPath)

	items := makeURLItems(project, global)
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "ChatGPT Project URLs"
	l.SetFilteringEnabled(true)
	selectDefault(&l, project, global)

	name := textinput.New()
	name.Placeholder = "Name (e.g., Core Project)"
	name.CharLimit = 60
	name.Width = 40

	url := textinput.New()
	url.Placeholder = "https://chatgpt.com/g/.../project"
	url.CharLimit = 200
	url.Width = 60

	return URLPickerModel{
		list:        l,
		projectPath: projectPath,
		globalPath:  globalPath,
		project:     project,
		global:      global,
		editName:    name,
		editURL:     url,
	}
}

func (m URLPickerModel) Init() tea.Cmd {
	return nil
}

func (m URLPickerModel) Update(msg tea.Msg) (URLPickerModel, tea.Cmd) {
	if m.editing {
		return m.updateEdit(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return URLPickerCancelledMsg{} }
		case "enter":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.touch(item)
			return m, func() tea.Msg { return URLPickedMsg{URL: item.url} }
		case "a":
			m.startEdit(urlScopeProject, "", "", true)
			return m, nil
		case "e":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.startEdit(item.scope, item.name, item.url, false)
			return m, nil
		case "d":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.delete(item)
			return m, nil
		case "s":
			item, ok := m.list.SelectedItem().(urlItem)
			if !ok {
				return m, nil
			}
			m.setDefault(item)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *URLPickerModel) SetSize(width, height int) {
	m.list.SetSize(width, height-4)
}

func (m URLPickerModel) View() string {
	if m.editing {
		return m.editView()
	}

	help := lipgloss.NewStyle().Faint(true).Render("[enter] use  [a] add  [e] edit  [d] delete  [s] default  [esc] cancel")
	return m.list.View() + "\n" + help
}

func makeURLItems(project URLStore, global URLStore) []list.Item {
	var items []list.Item
	for _, it := range project.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeProject})
	}
	for _, it := range global.Items {
		items = append(items, urlItem{name: it.Name, url: it.URL, scope: urlScopeGlobal})
	}
	return items
}

func selectDefault(l *list.Model, project URLStore, global URLStore) {
	if l == nil {
		return
	}
	name, scope := defaultNameScope(project, global)
	if name == "" {
		return
	}
	for idx, item := range l.Items() {
		if it, ok := item.(urlItem); ok && it.name == name && it.scope == scope {
			l.Select(idx)
			return
		}
	}
}

func defaultNameScope(project URLStore, global URLStore) (string, string) {
	if project.Default != "" {
		return project.Default, urlScopeProject
	}
	if global.Default != "" {
		return global.Default, urlScopeGlobal
	}
	return "", ""
}

func (m URLPickerModel) DefaultURL() string {
	name, scope := defaultNameScope(m.project, m.global)
	if name == "" {
		return ""
	}
	store := m.storeFor(scope)
	if store == nil {
		return ""
	}
	for _, it := range store.Items {
		if it.Name == name {
			return it.URL
		}
	}
	return ""
}

func (m *URLPickerModel) refresh() {
	m.list.SetItems(makeURLItems(m.project, m.global))
	selectDefault(&m.list, m.project, m.global)
}

func (m *URLPickerModel) touch(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	for i := range store.Items {
		if store.Items[i].Name == item.name {
			store.Items[i].LastUsed = nowRFC3339()
			break
		}
	}
	_ = m.saveStores()
}

func (m *URLPickerModel) delete(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	var out []URLItem
	for _, it := range store.Items {
		if it.Name == item.name {
			continue
		}
		out = append(out, it)
	}
	store.Items = out
	if store.Default == item.name {
		store.Default = ""
	}
	_ = m.saveStores()
	m.refresh()
}

func (m *URLPickerModel) setDefault(item urlItem) {
	store := m.storeFor(item.scope)
	if store == nil {
		return
	}
	store.Default = item.name
	_ = m.saveStores()
}

func (m *URLPickerModel) startEdit(scope, name, url string, isNew bool) {
	m.editing = true
	m.editScope = scope
	m.editIsNew = isNew
	m.editName.SetValue(name)
	m.editURL.SetValue(url)
	m.editName.Focus()
	m.editURL.Blur()
	m.errMsg = ""
}

func (m URLPickerModel) editView() string {
	scopeLabel := fmt.Sprintf("Scope: %s (g=global, p=project)", m.editScope)
	if m.globalPath == "" && m.projectPath != "" {
		scopeLabel = "Scope: project"
		m.editScope = urlScopeProject
	}
	if m.projectPath == "" && m.globalPath != "" {
		scopeLabel = "Scope: global"
		m.editScope = urlScopeGlobal
	}
	lines := []string{
		"Add / Edit ChatGPT URL",
		scopeLabel,
		"",
		"Name:",
		m.editName.View(),
		"",
		"URL:",
		m.editURL.View(),
		"",
		"[tab] switch field  [enter] save  [esc] cancel",
	}
	if m.errMsg != "" {
		lines = append(lines, "", lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.errMsg))
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m URLPickerModel) updateEdit(msg tea.Msg) (URLPickerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.editing = false
			return m, nil
		case "tab":
			if m.editName.Focused() {
				m.editName.Blur()
				m.editURL.Focus()
			} else {
				m.editURL.Blur()
				m.editName.Focus()
			}
			return m, nil
		case "g":
			if m.globalPath != "" {
				m.editScope = urlScopeGlobal
			}
		case "p":
			if m.projectPath != "" {
				m.editScope = urlScopeProject
			}
		case "enter":
			name := strings.TrimSpace(m.editName.Value())
			url := strings.TrimSpace(m.editURL.Value())
			if name == "" || !isValidURL(url) {
				m.errMsg = "Name and a valid URL are required."
				return m, nil
			}
			m.saveEdit(name, url)
			m.editing = false
			m.refresh()
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.editName.Focused() {
		m.editName, cmd = m.editName.Update(msg)
	} else {
		m.editURL, cmd = m.editURL.Update(msg)
	}
	return m, cmd
}

func (m *URLPickerModel) saveEdit(name, url string) {
	scope := m.editScope
	if scope == "" {
		scope = urlScopeProject
	}

	// remove from other store if scope changed
	m.removeByName(name, urlScopeProject)
	m.removeByName(name, urlScopeGlobal)

	store := m.storeFor(scope)
	if store == nil {
		return
	}
	updated := false
	for i := range store.Items {
		if store.Items[i].Name == name {
			store.Items[i].URL = url
			store.Items[i].LastUsed = nowRFC3339()
			updated = true
			break
		}
	}
	if !updated {
		store.Items = append(store.Items, URLItem{Name: name, URL: url, LastUsed: nowRFC3339()})
	}
	_ = m.saveStores()
}

func (m *URLPickerModel) removeByName(name, scope string) {
	store := m.storeFor(scope)
	if store == nil {
		return
	}
	var out []URLItem
	for _, it := range store.Items {
		if it.Name == name {
			continue
		}
		out = append(out, it)
	}
	store.Items = out
}

func (m *URLPickerModel) storeFor(scope string) *URLStore {
	switch scope {
	case urlScopeProject:
		if m.projectPath == "" {
			return nil
		}
		return &m.project
	case urlScopeGlobal:
		if m.globalPath == "" {
			return nil
		}
		return &m.global
	default:
		return nil
	}
}

func (m *URLPickerModel) saveStores() error {
	if err := SaveURLStore(m.projectPath, m.project); err != nil {
		return err
	}
	if err := SaveURLStore(m.globalPath, m.global); err != nil {
		return err
	}
	return nil
}
```

internal/tui/url_store.go
```
package tui

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	urlScopeProject = "project"
	urlScopeGlobal  = "global"
)

type URLItem struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	LastUsed string `json:"lastUsed,omitempty"`
}

type URLStore struct {
	Default string    `json:"default"`
	Items   []URLItem `json:"items"`
}

func LoadURLStore(path string) (URLStore, error) {
	if path == "" {
		return URLStore{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return URLStore{}, nil
		}
		return URLStore{}, err
	}
	var store URLStore
	if err := json.Unmarshal(data, &store); err != nil {
		return URLStore{}, err
	}
	return store, nil
}

func SaveURLStore(path string, store URLStore) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func ProjectURLStorePath(statePath, packSource string) string {
	if statePath != "" {
		base := strings.TrimSuffix(statePath, ".state.json")
		return base + ".chatgpt-urls.json"
	}
	if packSource == "" {
		return ""
	}
	return packSource + ".chatgpt-urls.json"
}

func GlobalURLStorePath() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".oraclepack", "chatgpt-urls.json")
}

func nowRFC3339() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func isValidURL(value string) bool {
	v := strings.TrimSpace(value)
	if v == "" {
		return false
	}
	return strings.HasPrefix(v, "http://") || strings.HasPrefix(v, "https://")
}
```

internal/tui/url_store_test.go
```
package tui

import (
	"path/filepath"
	"testing"
)

func TestURLStoreSaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "urls.json")

	store := URLStore{
		Default: "Primary",
		Items: []URLItem{{
			Name: "Primary",
			URL:  "https://chatgpt.com/g/primary",
		}},
	}

	if err := SaveURLStore(path, store); err != nil {
		t.Fatalf("failed to save store: %v", err)
	}

	loaded, err := LoadURLStore(path)
	if err != nil {
		t.Fatalf("failed to load store: %v", err)
	}

	if loaded.Default != store.Default {
		t.Fatalf("expected default %q, got %q", store.Default, loaded.Default)
	}
	if len(loaded.Items) != 1 || loaded.Items[0].URL != store.Items[0].URL {
		t.Fatalf("loaded items mismatch: %+v", loaded.Items)
	}
}

func TestURLPickerDefaultURLPrefersProject(t *testing.T) {
	dir := t.TempDir()
	projectPath := filepath.Join(dir, "project.json")
	globalPath := filepath.Join(dir, "global.json")

	project := URLStore{
		Default: "Project",
		Items: []URLItem{{
			Name: "Project",
			URL:  "https://chatgpt.com/g/project",
		}},
	}
	global := URLStore{
		Default: "Global",
		Items: []URLItem{{
			Name: "Global",
			URL:  "https://chatgpt.com/g/global",
		}},
	}

	if err := SaveURLStore(projectPath, project); err != nil {
		t.Fatalf("failed to save project store: %v", err)
	}
	if err := SaveURLStore(globalPath, global); err != nil {
		t.Fatalf("failed to save global store: %v", err)
	}

	picker := NewURLPickerModel(projectPath, globalPath)
	if got := picker.DefaultURL(); got != project.Items[0].URL {
		t.Fatalf("expected project default URL %q, got %q", project.Items[0].URL, got)
	}
}
```

.config/mcp/mcp-builder/SKILL.md
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

.config/mcp/oraclepack-gold-pack/SKILL.md
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

.config/mcp/oraclepack-taskify/SKILL.md
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

.config/skills/oracle-pack/SKILL.md
```
---
name: oracle-pack
description: Generate exactly 20 evidence-cited, ROI-ranked strategist questions for the current repository (12 immediate + 8 strategic), then emit them as runnable @steipete/oracle CLI commands with minimal targeted file attachments and deterministic per-question output paths (a copy/paste-ready “oracle question pack” in the same scratch-style format as oracle-scratch.md).
---

## Quick start

Use this skill when the user wants strategist questions **and** wants to run each question through Oracle as a separate command (so a second model can answer with repo context).

Invocation:

- `$oracle-pack <free-text args>`

Primary output: a Markdown doc whose main content is a **single** fenced `bash` block containing **exactly 20** `oracle ...` commands (copy/paste-ready), in the same “scratch-style” format shown in `references/oracle-scratch-format.md`.

## Inputs

- Repository contents (current working directory).
- Free-text args (may include): `codebase_name`, `constraints`, `non_goals`, `team_size`, `deadline`, plus Oracle pack controls:
  - `out_dir` (default: `oracle-out`)
  - `oracle_cmd` (default: `oracle`)
  - `oracle_flags` (default: `--files-report`)
  - `extra_files` (optional: comma-separated extra `-f/--file` entries to include in *every* command; default: none)

## Deterministic search/tooling rules (must follow)

- Prefer deterministic, non-interactive commands so runs are reproducible.
- Before using `ck` in a session, **always run**: `ck --help` and only rely on flags that `ck --help` confirms.
- Default search tool is `ck`:
  - Conceptual: `ck --sem "<query>"`
  - Literal: `ck --regex "<pattern>"`
  - Best of both: `ck --hybrid "<query>"`
  - For post-processing: prefer `ck --jsonl | jq ... | sort | head`
- Structural search: use `ast-grep` when syntax/shape matters (see `references/inference-first-discovery.md`).
- File finding: prefer `fd` (filename/path/extension filters).

## Args parsing (no follow-ups)

Extract values if present; otherwise set defaults:

- `codebase_name`: `Unknown`
- `constraints`: `None`
- `non_goals`: `None`
- `team_size`: `Unknown`
- `deadline`: `Unknown`
- `out_dir`: `oracle-out`
- `oracle_cmd`: `oracle`
- `oracle_flags`: `--files-report`
- `extra_files`: (empty)

Honor `constraints`. Explicitly exclude `non_goals`.

## Workflow (inference-first discovery → evidence → questions → oracle pack)

### 1) Inference-first discovery (adaptive, codebase-search compliant)

Goal: infer *what this repo is* and *where the truth likely lives* before broad sweeps.

Do this in order:

1. **Discover the search interface (required)**
   - Run: `ck --help`
   - If `ck` indicates it needs setup/indexing, follow the instructions printed by `ck`.

2. **Map the repo surface (cheap, high-signal)**
   - Use `fd` to list likely roots deterministically:
     - `fd . -d 2`
     - `fd -p README.md`
     - `fd -p package.json`
     - `fd -p pyproject.toml`
     - `fd -p go.mod`
     - `fd -p Cargo.toml`
   - Read the smallest set of “index” files that describe structure:
     - README / docs index if present
     - primary manifests/config
     - obvious entrypoint referenced by scripts/manifests

3. **Infer stack + conventions from evidence**
   - From manifests/config, infer runtime + framework signals (dependencies, filenames, directory conventions).
   - From entrypoints, infer wiring patterns (router, DI/container, job scheduler, migration tool).

4. **Derive a “search plan” from inferred signals**
   - Prefer following imports/references from entrypoints into:
     - routing/handlers
     - auth middleware/policies
     - job/queue registration
     - data layer/migrations
     - feature flags config
     - observability bootstrap
   - Use `ck` in the smallest inferred app roots first:
     - Conceptual: `ck --sem "<subsystem intent>"`
     - Hybrid when unsure: `ck --hybrid "<subsystem intent>"`
   - If results are broad/noisy:
     - Narrow to directories/files from the top hits, then re-run `ck`.
     - Use `ast-grep` when you need structure, not text matches.

This improves efficiency vs. hard-targeting a long list of patterns up front: you start with the repo’s own “self-description” and expand only as needed.

### 2) Evidence harvesting (collect anchors before writing questions)

Collect **at least 20 candidate anchors** as one of:

- `{path}:{symbol}` (preferred: nearest function/class/type/handler)
- `{endpoint}` (literal endpoint string)
- `{event}` (job/queue/event name)

For each required category, attempt to identify at least one anchor:

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

Use deterministic selection of candidates:
- If using `ck --jsonl`, post-process to stable top-N:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- If using plain output, cap deterministically with `head`.

If evidence for a category cannot be found:

- still generate a question for it with reference `Unknown`
- explicitly state which missing artifact pattern would likely provide it (e.g., `**/migrations/**` not found)

### 3) Generate exactly 20 strategist questions (evidence-cited + minimal experiments)

Produce exactly:

- **12** Immediate Exploration
- **8** Strategic Planning

Each question must include:

- **Reference**: `{path}:{symbol}` OR `{endpoint}` OR `{event}` OR `Unknown`
- **Question**: incisive, feasibility-focused, no scope creep
- **Rationale**: exactly one sentence
- **Smallest experiment today**: exactly one concrete action runnable today (targeted `ck` query, targeted `ast-grep` pattern, trace a codepath, add a log line, minimal unit/integration test, run migration dry-run, execute a single endpoint, etc.)

No duplicates (by intent or reference).

### 4) Score and order by near-term ROI (required math + sorting)

For each question compute:

- `ROI = (impact * confidence) / effort`
- `impact`, `confidence`, `effort` ∈ `{0.1..1.0}` (one decimal)

Sort all 20 descending by ROI; break ties by lower effort.

### 5) Convert the 20 questions into an Oracle “question pack” (final deliverable)

For each question (in final sorted order), emit a runnable command using:

- command: `{{oracle_cmd}}` (default `oracle`)
- include `{{oracle_flags}}` (default `--files-report`)
- deterministic output file: `--write-output "<out_dir>/<nn>-<slug>.md"`
  - `<nn>` = zero-padded 2-digit index (01..20)
  - `<slug>` = short, filesystem-safe slug derived from `{category}` + a hint of `{reference}` (fallback to category only)
- prompt: a structured prompt that includes the strategist question fields and constraints:
  - reference, question, rationale, smallest experiment today
  - constraints + non_goals (explicit)
  - requested Oracle answer shape (concise, evidence-cited, and actionable)

#### Attachment selection (must be minimal, evidence-driven)

For each command, attach only the smallest set of files that lets Oracle answer correctly:

- If reference is `{path}:{symbol}`:
  - attach that `path`
  - optionally attach **one** upstream config/router/entrypoint file that shows how it’s invoked (only if needed)
- If reference is `{endpoint}` or `{event}`:
  - attach the file(s) where you found the literal endpoint/event definition
- If reference is `Unknown`:
  - attach only “index” files (README + manifest + 1 best-guess entrypoint), and in the prompt state what artifact pattern is missing

Follow `references/attachment-minimization.md`.

Also:
- If `extra_files` was provided in args, include those `-f` entries in every command (after the evidence-minimal attachments).

### 6) Render final output using the pack template

Use `assets/oracle-pack-template.md` as the strict shape for the final Markdown deliverable.

## Output contract (strict)

Produce exactly one Markdown deliverable that:

1. Matches the structure in `assets/oracle-pack-template.md`
2. Contains **exactly 20** `oracle` invocations total
3. Commands are sorted by ROI (desc; ties by lower effort)
4. Each command:
   - includes `--write-output "<out_dir>/<nn>-<slug>.md"`
   - includes minimal targeted `-f/--file` attachments
   - embeds the strategist question (reference, question, rationale, experiment) in the `-p/--prompt` text

## Failure modes

- Missing/insufficient evidence:
  - use reference `Unknown`
  - state the missing artifact pattern that would provide evidence
  - keep the question actionable and experiment minimal
  - keep attachments limited to index files (don’t attach huge globs)
- Search tool limitations (`ck` missing/setup required):
  - run `ck --help` and follow printed setup; if unavailable, fall back to reading index files + minimal `fd`-based locating
  - proceed with partial evidence; mark affected references `Unknown`
- Ambiguous args:
  - apply defaults without follow-ups

## References

- `assets/oracle-pack-template.md` — exact final output shape
- `references/inference-first-discovery.md` — inference-driven search plan aligned to `ck`/`ast-grep`/`fd`
- `references/attachment-minimization.md` — deterministic rules for minimal `-f` selections
- `references/oracle-scratch-format.md` — the target “scratch-style” formatting to mimic

## Invocation examples

- “Generate an oracle question pack for this repo. codebase_name=Payments API; constraints=no new infra; non_goals=mobile app; team_size=3; deadline=2 weeks; out_dir=artifacts/oracle”
- “Create oracle strategist questions; constraints=ship bugfixes only; deadline=Friday; out_dir=oracle-out”
- “Produce an oracle pack: focus on auth + jobs + migrations; non_goals=refactor; out_dir=oracle-q”
- “Generate strategist oracle commands; constraints=keep DB schema stable; team_size=2; deadline=1 week”
- “Make an Oracle question pack; oracle_cmd='npx -y @steipete/oracle'; oracle_flags='--engine browser  --files-report'; out_dir=oracle-out”

### Write output to `docs/` (required)

- Ensure a `docs/` directory exists at the repo root (create it if missing).
- Write the complete final deliverable to:
  - `docs/oracle-pack-YYYY-MM-DD.md` (use today’s date; ISO 8601)
  - Fallback if date is unavailable: `docs/oracle-pack.md`
- The file must contain only the deliverable Markdown (no extra preamble/epilogue).
- In the assistant response, print the chosen path first on a single line: `Output file: <path>`, then print the same Markdown content.
```

.config/skills/oracle-pack-rpg-infused/SKILL.md
```
---
name: oracle-pack-rpg-infused
description: Generate a strictly oraclepack-ingestible “oracle strategist question pack” Markdown file under docs/oracle-pack-YYYY-MM-DD.md containing exactly one fenced bash block with exactly 20 numbered oracle commands (01..20), ROI-scored/sorted and category-covered, with an RPG block (WHAT/HOW + dependency edges) embedded in each -p prompt for graph extraction.
---

# Oracle Pack (RPG-infused)

## Quick start

1) Read `assets/oracle-pack-template.md` and keep its structure exactly.
2) Produce exactly one Markdown deliverable and output it verbatim in your final response:
   - First line: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
   - Then the exact Markdown content (no extra commentary, no additional code fences).
3) Ensure the bash block contains exactly 20 steps, numbered `01..20`, each with:
   - One header comment line in the required format
   - Exactly one runnable oracle command line

For the full spec, follow `references/oracle-pack-spec.md`.

## Inputs and defaults

Interpret any user-provided details as “args” and apply these defaults when missing:

- `codebase_name=Unknown`
- `constraints=None`
- `non_goals=None`
- `team_size=Unknown`
- `deadline=Unknown`
- `out_dir=oracle-out`
- `oracle_cmd=oracle`
- `oracle_flags=--files-report`
- `extra_files=empty` (only use if the user provides an explicit mechanism/flag; otherwise omit and keep evidence handling inside the prompt body as text)

If the user supplies their own values, use them verbatim unless they would break the strict output contract.

## Workflow

1) Establish minimal repo understanding (inference-first)
   - Prefer reading index artifacts (e.g., README, top-level config/manifests) before proposing file sweeps.
   - If you cannot access the repo contents, keep `reference=Unknown` and make the “missing artifact pattern(s)” explicit in the prompt body.

2) Plan the 20 questions (before writing any command lines)
   - Use only the required categories (no additions, no removals).
   - Ensure horizon mix: exactly 12 `Immediate` and 8 `Strategic`.
   - Ensure RPG infusion in every step’s `-p` prompt (WHAT/HOW + DependsOn + Phase).
   - Ensure dependencies only point backward (DependsOn may reference only earlier step IDs).
   - Ensure coverage: across all 20 steps, each required category appears at least once.

3) Assign scoring and compute ROI
   - Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
   - Compute `ROI = (impact * confidence) / effort` (one decimal).
   - Sort all 20 steps by ROI descending; tie-break by lower effort.

4) Emit the pack Markdown
   - Copy `assets/oracle-pack-template.md` structure exactly.
   - Fill in metadata values and the 20 steps.
   - Keep exactly ONE fenced bash block.
   - Keep exactly 20 steps (01..20) inside that block.

5) Emit the coverage check section
   - List the 10 required categories exactly once each with `OK` or `Missing(...)`.

6) Validate (optional but recommended when execution is available)
   - Run `scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md` and fix any reported issues.

## Output contract (must hold)

- Exactly one Markdown deliverable.
- Matches `assets/oracle-pack-template.md` structure exactly.
- Exactly one fenced bash block.
- Bash block contains exactly 20 steps, numbered 01..20.
- Each step has:
  - Header line with: `ROI=... impact=... confidence=... effort=... horizon=... category=... reference=...`
  - Exactly one runnable oracle command including:
    - `${oracle_cmd}` (default `oracle`)
    - `${oracle_flags}` (default `--files-report`)
    - `--write-output "<out_dir>/<nn>-<slug>.md"`
- Steps sorted by ROI desc; tie-break by lower effort.
- Coverage check section present and correctly formatted.

## Failure modes (do not guess)

- Unknown subsystem/file references: set `reference=Unknown` and include explicit missing file/path patterns in section (4) of the prompt body.
- Missing or unclear user args: apply defaults; do not invent tool flags or repo-specific details.
- Potential contract violations: prioritize strict compliance over richer prose.

## Invocation examples

1) “Generate an RPG-infused oracle strategist pack for this repo. out_dir=oracle-out.”
2) “Create the oracle-pack for a monorepo; prioritize caching/state, observability, and permissions. Use oracle_flags=--files-report.”
3) “Make a pack for codebase_name=AcmeWeb, constraints=‘No DB migrations this sprint’, non_goals=‘No UI redesign’.”
4) “Produce a pack that emphasizes feature flags and invariants, but still covers all required categories.”
5) “Generate the pack with conservative confidence values (0.2–0.6) and document missing artifact patterns explicitly.”
```

.config/skills/oracle-strategist-commands/SKILL.md
```
---

name: oracle-strategist-commands
description: Generate exactly 20 evidence-anchored strategist questions as runnable @steipete/oracle bash commands (12 immediate + 8 strategic), using inference-first repo discovery to avoid wasted searching.
metadata:
  short-description: 20 ROI-ranked Oracle CLI question commands for a repo (inference-first discovery)
  output: bash
  count: 20
---

## Use when

Use this skill when the user wants incisive, feasibility-focused strategist questions, but **as an executable pack** of **exactly 20** `oracle` CLI commands.

## Deliverable (strict)

Output MUST be:

- A single fenced code block: ```bash
- Optional header lines (env vars, `mkdir -p`, shared preamble)
- Then **exactly 20** Oracle invocations (lines beginning with `oracle` or `npx -y @steipete/oracle`)
- Counts:
  - **12** commands that are “Immediate” (01–12)
  - **8** commands that are “Strategic” (13–20)
- No prose outside the bash code fence.

Each command MUST include:

- `-p/--prompt`
- at least one `-f/--file`
- `--files-report`
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- `--wait` (unless user explicitly requests omitting it)

Optional but recommended:

- a commented preflight line using `--dry-run summary` before each real command.

## Inputs (do not ask follow-ups)

Interpret free-text args as optionally including:

- `codebase_name` (default `Unknown`)
- `constraints` (default `None`)
- `non_goals` (default `None`)
- `team_size` (default `Unknown`)
- `deadline` (default `Unknown`)
- `engine` (`api|browser`; if unspecified, omit)
- `model` (if unspecified, omit)
- `out_dir` (default `docs/oracle-questions-YYYY-MM-DD`)

If missing, apply defaults and proceed.

## Core principle: inference-first discovery

Avoid “hard-targeting” specific artifact types up front (e.g., hunting for migrations/auth/jobs) because many repos won’t have them. Instead:

1) Gather cheap high-signal evidence (anchors).
2) Infer the stack and architecture hypotheses.
3) Run small, evidence-gated probes to confirm.
4) Only search/read deeper for categories that now have supporting evidence.

This reduces wasted probing and keeps attachments minimal.

## Workflow

### 1) Build a repo fingerprint (anchors only)

Read the smallest set of files that often exist and immediately reveal stack/shape:

- `README*` (first ~200 lines)
- root manifests that exist (one or more): `package.json`, `pyproject.toml`, `go.mod`, `Cargo.toml`, `pom.xml`, `build.gradle`, `composer.json`, `Gemfile`
- `Dockerfile`, `docker-compose.*` (if present)
- `.github/workflows/*` (names + one representative workflow file)
- `Makefile` / `Taskfile.*` (if present)

Also capture a lightweight file index summary (without reading everything):

- top-level directories
- dominant file extensions
- any “obvious entrypoint” indicators from scripts/config (e.g., package.json scripts, go `cmd/`, rust `src/main.rs`)

### 2) Form hypotheses (stack + entrypoints)

From the fingerprint, infer:

- primary language/runtime
- framework candidates (e.g., Express/Nest/Next, FastAPI/Django, chi/gin, Rails, Spring)
- likely entrypoints and routing surfaces
- where “truth” lives (schema, API contracts, config)

Write down 3–6 candidate entrypoints (paths) you will validate.

### 3) Evidence-gated probing loop (bounded)

Run a small loop of probes; each probe must have a clear purpose and a stop condition.

**Budgets (defaults):**

- max probes: 12
- max file reads: 25
- max search passes: 6

**Allowed probes (choose the cheapest next step):**

- Open an inferred entrypoint and follow imports to one routing/handler layer.
- Read a config file referenced by an entrypoint.
- Perform a targeted search ONLY for confirmed-framework idioms (e.g., “router”, “app.get”, “FastAPI()”, “chi.NewRouter”) after framework is evidenced.
- If you see a strong signal (e.g., `prisma` in deps), then and only then probe its likely files (`prisma/schema.prisma`, migrations).

**Do NOT** scan for categories (auth/jobs/migrations/feature flags/observability) unless you have at least one of:

- a dependency/config signal (library name, config key)
- an import path signal
- a directory naming signal found naturally via the fingerprint or followed imports

### 4) Build the evidence index (mandatory)

Create an internal evidence index with **>= 20** references that actually exist, as one of:

- `path:symbol` (function/class/type/handler name)
- explicit endpoint strings (e.g., `/api/users`)
- job/event names (queue/topic/schedule identifiers)
- config keys/flags (feature gates, env var keys)

For each reference, also record a “best attachment set” (1–6 files) that makes it intelligible.

If the repo is too small to reach 20, pad with `Unknown` entries; those questions must explicitly call out missing evidence.

### 5) Draft exactly 20 strategist questions (12 immediate + 8 strategic)

Each question must:

- be feasibility-focused and grounded in repo reality
- respect `constraints`
- explicitly exclude `non_goals`
- be actionable for `team_size` and `deadline`
- cite at least one evidence reference (or `Unknown` with a clear gap)
- fall into one of these areas (only when evidenced; otherwise treat as a “risk/unknowns” question):
  - contracts/interfaces
  - invariants/data consistency
  - auth/permissions
  - migrations/schema
  - request/UX flows (if applicable)
  - failure modes & reliability
  - feature flags/config gating
  - observability (logs/metrics/tracing)

Guidance:

- Immediate (01–12): unblockers, correctness, runtime risk, deployment friction, quick wins.
- Strategic (13–20): architectural constraints, roadmap leverage, risk retirement, scalability, maintainability.

### 6) Emit the question pack as bash commands (final output)

#### Common command shape

Each command MUST:

- call Oracle with `-p/--prompt` and `-f/--file` (repeat `-f` as needed)
- include `--files-report`
- include `--write-output "<out_dir>/<nn>-<slug>.md"`
- include `--wait` unless user opted out
- attach minimal files:
  - small base context: 1–3 anchor files (entrypoint + config)
  - focus context: files containing the cited evidence reference(s)

Avoid broad `**/*` globs. Prefer explicit files and small globs with negation excludes.

#### Prompt micro-format (required in every `-p`)

Each `-p` must require Oracle to respond using exactly:

- **Answer**
- **Evidence** (bullets; cite concrete refs)
- **Risks / unknowns**
- **Smallest experiment** (runnable today)

Keep prompts short and specific.

#### Optional preflight (recommended)

For each command, optionally emit a commented “dry-run” line immediately above it:

- same command + `--dry-run summary`
This lets the user validate attachments/token weights without sending.

## Failure handling rules

- Missing evidence for a category:
  - do NOT hunt for it
  - convert into a “risk/unknowns” question
  - attach best available anchors + nearest relevant files
- Very large repos:
  - tighten budgets
  - reduce base context to 1–2 files
  - prefer following import chains over directory scans
- Unclear entrypoint:
  - use build scripts/config to infer; otherwise choose the most central “composition root” file you can find and proceed

## Write output to docs (required)

When running this skill, also write the complete deliverable (the single `bash` fence) to:

- `docs/oracle-strategist-commands-YYYY-MM-DD.md` (ISO date)
Fallback:
- `docs/oracle-strategist-commands.md`

In the assistant response:

1) Print: `Output file: <path>`
2) Then print the exact same `bash` fenced deliverable.
```

.config/skills/oraclepack-gold-pack/SKILL.md
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

.config/skills/oraclepack-taskify/SKILL.md
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

.config/skills/oraclepack-usecase-rpg/SKILL.md
```
---
name: oraclepack-usecase-rpg
description: Generate a strict, oraclepack-ingestible 20-step Oracle strategist pack (single ```bash fence) with ROI-ranked step headers and RPG (Repository Planning Graph) metadata embedded in each -p prompt body.
metadata:
  short-description: Strict oraclepack pack generator (RPG-infused)
---

## Quick start

Use this skill when the user asks for any of the following:
- an “oraclepack pack”, “oracle strategist question pack”, or “20-step ROI-ranked pack”
- “oraclepack use cases” that must keep the strict runner-ingestible output
- the same, but with RPG/RPG graph/Repository Planning Graph infusion

What you produce:
- Exactly one Markdown deliverable (intended path: `docs/oracle-pack-YYYY-MM-DD.md`)
- The deliverable contains **exactly one** fenced code block that starts with **exactly** ` ```bash` then a newline (no extra tokens after `bash`)
- The fenced `bash` block contains **exactly 20** sequential steps, `01..20`

Reference template: `references/oracle-pack-template.md`

Optional validator:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

## Workflow

### 1) Parse inputs (free-text args)

Interpret the user’s trailing text as conceptual `{{args}}` and extract:
- `codebase_name` (default: `Unknown`)
- `constraints` (default: `None`)
- `non_goals` (default: `None`)
- `team_size` (default: `Unknown`)
- `deadline` (default: `Unknown`)
- `out_dir` (default: `oracle-out`)
- `oracle_cmd` (default: `oracle`)
- `oracle_flags` (default: `--files-report`)
- `extra_files` (default: empty; append **literally** to every command if provided)

If any value is missing or ambiguous, use the defaults above (do not ask questions).

### 2) Repo discovery (inference-first, deterministic)

Read small, high-signal “index” artifacts first (if they exist):
- README / overview docs
- manifests / dependency lockfiles
- entrypoints (main binaries, CLI, server startup)
- config examples
- CI configuration

Then do targeted discovery (only as needed):
- Prefer deterministic searches (e.g., `ck`, `ast-grep`, `fd`) if available.
- Cap results deterministically (fixed max hits).
- If a file/path does not exist: record `Unknown` and continue.

### 3) Plan the 20 probes (what to cover)

You must keep categories **exactly** to this set (do not add/remove/rename):
- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

Across the 20 steps, ensure at least one step explicitly targets each sub-layer below (embed it in the *Question* text; map to the best required category above):
- supply chain/dependencies
- secrets/credentials
- CI/CD or release process
- runtime configuration/IaC (if present)
- privacy/PII/retention (if applicable)
- SLO/alerting/runbooks
- cost/performance hotspots
- testing strategy/coverage

Horizon mix requirement:
- Exactly **12** steps with `horizon=Immediate`
- Exactly **8** steps with `horizon=Strategic`

### 4) RPG infusion (non-breaking)

Inside each step’s `-p` prompt body, include an `RPG:` block with:
- `FunctionalNode: <Capability>::<Feature> (WHAT)`
- `StructuralNode: <module-or-dir> :: <public surface> (HOW)`
- `DependsOn: [<prior step ids, optional>]`
- `Phase: <P0|P1|P2|P3>`

Dependency rule (to preserve ROI-sorted executability):
- `DependsOn` may reference **only earlier step IDs**.

Phase mapping guidance (do not create new categories):
- `P0`: contracts/interfaces, invariants, permissions, observability
- `P1`: caching/state, background jobs, failure modes
- `P2`: migrations, feature flags
- `P3`: UX flows

### 5) ROI scoring and ordering

For each step choose:
- `impact`, `confidence`, `effort` as **one decimal** in `0.1..1.0`

Compute:
- `ROI = (impact * confidence) / effort`

Ordering requirement:
- Sort all 20 steps by ROI descending
- Tie-break by lower effort
- After sorting, assign IDs `01..20` in that sorted order

If ordering conflicts with RPG dependencies:
- Adjust impact/confidence/effort (and thus ROI) so prerequisite/foundation steps float earlier.
- Keep `DependsOn` pointing backwards only.

### 6) Render the pack (strict output)

You must output exactly one Markdown deliverable matching the structure in `references/oracle-pack-template.md`:
- Title section
- Parsed args list
- Commands section with **exactly one** fenced code block:
  - opening fence must be exactly: <code>```bash</code>
  - closing fence must be exactly: <code>```</code>
- Coverage check section

Inside the fenced `bash` block:
- Include a small prelude setting `out_dir="..."` (so pack metadata can be derived).
- Each of the 20 steps must start with a header line matching one of:
  - `# 01) ...`
  - `# 01 — ...`
  - `# 01 - ...`

Each step header line must include all tokens:
- `ROI=<number>`
- `impact=<i>`
- `confidence=<c>`
- `effort=<e>`
- `horizon=<Immediate|Strategic>`
- `category=<one of required categories>`
- `reference=<...>` (use `Unknown` if needed)

Each step must invoke the Oracle CLI once and must include:
- `--write-output "<out_dir>/<nn>-<slug>.md"`
- A `-p` prompt body that includes (in this exact section order):
  - `Strategist question #NN`
  - `RPG:` block (fields listed above)
  - `Reference: ...`
  - `Category: ...`
  - `Horizon: ...`
  - `ROI: ... (impact=..., confidence=..., effort=...)`
  - `Question: ...`
  - `Rationale: ...` (exactly one sentence)
  - `Smallest experiment today: ...` (exactly one action)
  - `Constraints: ...`
  - `Non-goals: ...`
  - `Answer format:` with the 4 required numbered items

Quoting guidance for `-p` (recommended, robust):
- Use a heredoc in command substitution so the prompt stays readable while remaining a single oracle invocation:
  - `-p "$(cat <<'PROMPT' ... PROMPT)"`

### 7) Validate before finalizing

Run the included validator if possible:
- `python3 scripts/validate_oraclepack_pack.py docs/oracle-pack-YYYY-MM-DD.md`

Fix any failures by editing the pack until it passes.

## Output contract

When invoked, produce:

1) A single line:
- `Output file: docs/oracle-pack-YYYY-MM-DD.md`

2) Immediately after, the **exact** Markdown content for that file (no extra commentary).

If the current date is not available in the execution context:
- Use `docs/oracle-pack.md`.

## Failure modes (do not guess)

- Missing repo evidence → use `Reference: Unknown` and in the prompt’s item (4) name the exact missing file/path patterns to attach next.
- Missing args → apply defaults; do not ask follow-ups.
- Conflicting constraints (ordering vs dependencies vs horizon mix) → preserve strict ingestion first:
  1) single `bash` fence
  2) 20 sequential steps
  3) required header tokens
  4) horizon mix 12/8
  Then adjust ROI inputs and dependency edges.

## Invocation examples

1) “Generate an RPG-infused oraclepack strategist pack for this repo; focus on permissions boundaries and CI/CD; keep strict oraclepack output.”

2) “Create a 20-step ROI-ranked oraclepack pack for `promptbench` with constraints=None and non_goals=None; default out_dir; strict format.”

3) “Produce an oraclepack pack emphasizing caching/state correctness and observability; include at least one step about cost/performance hotspots; strict template; RPG nodes.”

4) “Generate an oraclepack use-case pack for a Go CLI; include secrets/credentials handling and release process; keep categories fixed; strict single ```bash fence.”

5) “Make an oraclepack pack with explicit dependency edges (DependsOn) that only point backwards; preserve ROI-sorted order; strict output.”
```

.config/skills/strategist-questions-oracle-pack/SKILL.md
```
---
name: strategist-questions-oracle-pack
description: Generate exactly 20 evidence-cited, ROI-ranked strategist questions for the current repository (12 immediate + 8 strategic), then emit them as runnable @steipete/oracle CLI commands with minimal targeted file attachments and deterministic per-question output paths (a copy/paste-ready “oracle question pack” in the same scratch-style format as oracle-scratch.md).
---

## Quick start

Use this skill when the user wants strategist questions **and** wants to run each question through Oracle as a separate command (so a second model can answer with repo context).

Invocation:

- `$strategist-questions-oracle-pack <free-text args>`

Primary output: a Markdown doc whose main content is a **single** fenced `bash` block containing **exactly 20** `oracle ...` commands (copy/paste-ready), in the same “scratch-style” format shown in `references/oracle-scratch-format.md`.

## Inputs

- Repository contents (current working directory).
- Free-text args (may include): `codebase_name`, `constraints`, `non_goals`, `team_size`, `deadline`, plus Oracle pack controls:
  - `out_dir` (default: `oracle-out`)
  - `oracle_cmd` (default: `oracle`)
  - `oracle_flags` (default: `--browser-attachments always --files-report`)
  - `extra_files` (optional: comma-separated extra `-f/--file` entries to include in *every* command; default: none)

## Deterministic search/tooling rules (must follow)

- Prefer deterministic, non-interactive commands so runs are reproducible.
- Before using `ck` in a session, **always run**: `ck --help` and only rely on flags that `ck --help` confirms.
- Default search tool is `ck`:
  - Conceptual: `ck --sem "<query>"`
  - Literal: `ck --regex "<pattern>"`
  - Best of both: `ck --hybrid "<query>"`
  - For post-processing: prefer `ck --jsonl | jq ... | sort | head`
- Structural search: use `ast-grep` when syntax/shape matters (see `references/inference-first-discovery.md`).
- File finding: prefer `fd` (filename/path/extension filters).

## Args parsing (no follow-ups)

Extract values if present; otherwise set defaults:

- `codebase_name`: `Unknown`
- `constraints`: `None`
- `non_goals`: `None`
- `team_size`: `Unknown`
- `deadline`: `Unknown`
- `out_dir`: `oracle-out`
- `oracle_cmd`: `oracle`
- `oracle_flags`: `--browser-attachments always --files-report`
- `extra_files`: (empty)

Honor `constraints`. Explicitly exclude `non_goals`.

## Workflow (inference-first discovery → evidence → questions → oracle pack)

### 1) Inference-first discovery (adaptive, codebase-search compliant)

Goal: infer *what this repo is* and *where the truth likely lives* before broad sweeps.

Do this in order:

1. **Discover the search interface (required)**
   - Run: `ck --help`
   - If `ck` indicates it needs setup/indexing, follow the instructions printed by `ck`.

2. **Map the repo surface (cheap, high-signal)**
   - Use `fd` to list likely roots deterministically:
     - `fd . -d 2`
     - `fd -p README.md`
     - `fd -p package.json`
     - `fd -p pyproject.toml`
     - `fd -p go.mod`
     - `fd -p Cargo.toml`
   - Read the smallest set of “index” files that describe structure:
     - README / docs index if present
     - primary manifests/config
     - obvious entrypoint referenced by scripts/manifests

3. **Infer stack + conventions from evidence**
   - From manifests/config, infer runtime + framework signals (dependencies, filenames, directory conventions).
   - From entrypoints, infer wiring patterns (router, DI/container, job scheduler, migration tool).

4. **Derive a “search plan” from inferred signals**
   - Prefer following imports/references from entrypoints into:
     - routing/handlers
     - auth middleware/policies
     - job/queue registration
     - data layer/migrations
     - feature flags config
     - observability bootstrap
   - Use `ck` in the smallest inferred app roots first:
     - Conceptual: `ck --sem "<subsystem intent>"`
     - Hybrid when unsure: `ck --hybrid "<subsystem intent>"`
   - If results are broad/noisy:
     - Narrow to directories/files from the top hits, then re-run `ck`.
     - Use `ast-grep` when you need structure, not text matches.

This improves efficiency vs. hard-targeting a long list of patterns up front: you start with the repo’s own “self-description” and expand only as needed.

### 2) Evidence harvesting (collect anchors before writing questions)

Collect **at least 20 candidate anchors** as one of:

- `{path}:{symbol}` (preferred: nearest function/class/type/handler)
- `{endpoint}` (literal endpoint string)
- `{event}` (job/queue/event name)

For each required category, attempt to identify at least one anchor:

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

Use deterministic selection of candidates:
- If using `ck --jsonl`, post-process to stable top-N:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- If using plain output, cap deterministically with `head`.

If evidence for a category cannot be found:

- still generate a question for it with reference `Unknown`
- explicitly state which missing artifact pattern would likely provide it (e.g., `**/migrations/**` not found)

### 3) Generate exactly 20 strategist questions (evidence-cited + minimal experiments)

Produce exactly:

- **12** Immediate Exploration
- **8** Strategic Planning

Each question must include:

- **Reference**: `{path}:{symbol}` OR `{endpoint}` OR `{event}` OR `Unknown`
- **Question**: incisive, feasibility-focused, no scope creep
- **Rationale**: exactly one sentence
- **Smallest experiment today**: exactly one concrete action runnable today (targeted `ck` query, targeted `ast-grep` pattern, trace a codepath, add a log line, minimal unit/integration test, run migration dry-run, execute a single endpoint, etc.)

No duplicates (by intent or reference).

### 4) Score and order by near-term ROI (required math + sorting)

For each question compute:

- `ROI = (impact * confidence) / effort`
- `impact`, `confidence`, `effort` ∈ `{0.1..1.0}` (one decimal)

Sort all 20 descending by ROI; break ties by lower effort.

### 5) Convert the 20 questions into an Oracle “question pack” (final deliverable)

For each question (in final sorted order), emit a runnable command using:

- command: `{{oracle_cmd}}` (default `oracle`)
- include `{{oracle_flags}}` (default `--browser-attachments always --files-report`)
- deterministic output file: `--write-output "<out_dir>/<nn>-<slug>.md"`
  - `<nn>` = zero-padded 2-digit index (01..20)
  - `<slug>` = short, filesystem-safe slug derived from `{category}` + a hint of `{reference}` (fallback to category only)
- prompt: a structured prompt that includes the strategist question fields and constraints:
  - reference, question, rationale, smallest experiment today
  - constraints + non_goals (explicit)
  - requested Oracle answer shape (concise, evidence-cited, and actionable)

#### Attachment selection (must be minimal, evidence-driven)

For each command, attach only the smallest set of files that lets Oracle answer correctly:

- If reference is `{path}:{symbol}`:
  - attach that `path`
  - optionally attach **one** upstream config/router/entrypoint file that shows how it’s invoked (only if needed)
- If reference is `{endpoint}` or `{event}`:
  - attach the file(s) where you found the literal endpoint/event definition
- If reference is `Unknown`:
  - attach only “index” files (README + manifest + 1 best-guess entrypoint), and in the prompt state what artifact pattern is missing

Follow `references/attachment-minimization.md`.

Also:
- If `extra_files` was provided in args, include those `-f` entries in every command (after the evidence-minimal attachments).

### 6) Render final output using the pack template

Use `assets/oracle-pack-template.md` as the strict shape for the final Markdown deliverable.

## Output contract (strict)

Produce exactly one Markdown deliverable that:

1. Matches the structure in `assets/oracle-pack-template.md`
2. Contains **exactly 20** `oracle` invocations total
3. Commands are sorted by ROI (desc; ties by lower effort)
4. Each command:
   - includes `--write-output "<out_dir>/<nn>-<slug>.md"`
   - includes minimal targeted `-f/--file` attachments
   - embeds the strategist question (reference, question, rationale, experiment) in the `-p/--prompt` text

## Failure modes

- Missing/insufficient evidence:
  - use reference `Unknown`
  - state the missing artifact pattern that would provide evidence
  - keep the question actionable and experiment minimal
  - keep attachments limited to index files (don’t attach huge globs)
- Search tool limitations (`ck` missing/setup required):
  - run `ck --help` and follow printed setup; if unavailable, fall back to reading index files + minimal `fd`-based locating
  - proceed with partial evidence; mark affected references `Unknown`
- Ambiguous args:
  - apply defaults without follow-ups

## References

- `assets/oracle-pack-template.md` — exact final output shape
- `references/inference-first-discovery.md` — inference-driven search plan aligned to `ck`/`ast-grep`/`fd`
- `references/attachment-minimization.md` — deterministic rules for minimal `-f` selections
- `references/oracle-scratch-format.md` — the target “scratch-style” formatting to mimic

## Invocation examples

- “Generate an oracle question pack for this repo. codebase_name=Payments API; constraints=no new infra; non_goals=mobile app; team_size=3; deadline=2 weeks; out_dir=artifacts/oracle”
- “Create oracle strategist questions; constraints=ship bugfixes only; deadline=Friday; out_dir=oracle-out”
- “Produce an oracle pack: focus on auth + jobs + migrations; non_goals=refactor; out_dir=oracle-q”
- “Generate strategist oracle commands; constraints=keep DB schema stable; team_size=2; deadline=1 week”
- “Make an Oracle question pack; oracle_cmd='npx -y @steipete/oracle'; oracle_flags='--engine browser --browser-attachments always --files-report'; out_dir=oracle-out”

### Write output to `docs/` (required)

- Ensure a `docs/` directory exists at the repo root (create it if missing).
- Write the complete final deliverable to:
  - `docs/strategist-questions-oracle-pack-YYYY-MM-DD.md` (use today’s date; ISO 8601)
  - Fallback if date is unavailable: `docs/strategist-questions-oracle-pack.md`
- The file must contain only the deliverable Markdown (no extra preamble/epilogue).
- In the assistant response, print the chosen path first on a single line: `Output file: <path>`, then print the same Markdown content.
```

.config/mcp/mcp-builder/reference/evaluation.md
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
- Doesn't require deep exploration or understanding
- No synthesis or analysis needed

**Example 3: Ambiguous answer format**

```xml
<qa_pair>
   <question>List all the repositories that have Python as their primary language.</question>
   <answer>repo1, repo2, repo3, data-pipeline, ml-tools</answer>
</qa_pair>
```

This question is poor because:

- Answer is a list that could be returned in any order
- Difficult to verify with direct string comparison
- LLM might format differently (JSON array, comma-separated, newline-separated)
- Better to ask for a specific aggregate (count) or superlative (most stars)

## Verification Process

After creating evaluations:

1. **Examine the XML file** to understand the schema
2. **Load each task instruction** and in parallel using the MCP server and tools, identify the correct answer by attempting to solve the task YOURSELF
3. **Flag any operations** that require WRITE or DESTRUCTIVE operations
4. **Accumulate all CORRECT answers** and replace any incorrect answers in the document
5. **Remove any `<qa_pair>`** that require WRITE or DESTRUCTIVE operations

Remember to parallelize solving tasks to avoid running out of context, then accumulate all answers and make changes to the file at the end.

## Tips for Creating Quality Evaluations

1. **Think Hard and Plan Ahead** before generating tasks
2. **Parallelize Where Opportunity Arises** to speed up the process and manage context
3. **Focus on Realistic Use Cases** that humans would actually want to accomplish
4. **Create Challenging Questions** that test the limits of the MCP server's capabilities
5. **Ensure Stability** by using historical data and closed concepts
6. **Verify Answers** by solving the questions yourself using the MCP server tools
7. **Iterate and Refine** based on what you learn during the process

---

# Running Evaluations

After creating your evaluation file, you can use the provided evaluation harness to test your MCP server.

## Setup

1. **Install Dependencies**

   ```bash
   pip install -r scripts/requirements.txt
   ```

   Or install manually:

   ```bash
   pip install anthropic mcp
   ```

2. **Set API Key**

   ```bash
   export ANTHROPIC_API_KEY=your_api_key_here
   ```

## Evaluation File Format

Evaluation files use XML format with `<qa_pair>` elements:

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
</evaluation>
```

## Running Evaluations

The evaluation script (`scripts/evaluation.py`) supports three transport types:

**Important:**

- **stdio transport**: The evaluation script automatically launches and manages the MCP server process for you. Do not run the server manually.
- **sse/http transports**: You must start the MCP server separately before running the evaluation. The script connects to the already-running server at the specified URL.

### 1. Local STDIO Server

For locally-run MCP servers (script launches the server automatically):

```bash
python scripts/evaluation.py \
  -t stdio \
  -c python \
  -a my_mcp_server.py \
  evaluation.xml
```

With environment variables:

```bash
python scripts/evaluation.py \
  -t stdio \
  -c python \
  -a my_mcp_server.py \
  -e API_KEY=abc123 \
  -e DEBUG=true \
  evaluation.xml
```

### 2. Server-Sent Events (SSE)

For SSE-based MCP servers (you must start the server first):

```bash
python scripts/evaluation.py \
  -t sse \
  -u https://example.com/mcp \
  -H "Authorization: Bearer token123" \
  -H "X-Custom-Header: value" \
  evaluation.xml
```

### 3. HTTP (Streamable HTTP)

For HTTP-based MCP servers (you must start the server first):

```bash
python scripts/evaluation.py \
  -t http \
  -u https://example.com/mcp \
  -H "Authorization: Bearer token123" \
  evaluation.xml
```

## Command-Line Options

```
usage: evaluation.py [-h] [-t {stdio,sse,http}] [-m MODEL] [-c COMMAND]
                     [-a ARGS [ARGS ...]] [-e ENV [ENV ...]] [-u URL]
                     [-H HEADERS [HEADERS ...]] [-o OUTPUT]
                     eval_file

positional arguments:
  eval_file             Path to evaluation XML file

optional arguments:
  -h, --help            Show help message
  -t, --transport       Transport type: stdio, sse, or http (default: stdio)
  -m, --model           codex model to use (default: codex-3-7-sonnet-20250219)
  -o, --output          Output file for report (default: print to stdout)

stdio options:
  -c, --command         Command to run MCP server (e.g., python, node)
  -a, --args            Arguments for the command (e.g., server.py)
  -e, --env             Environment variables in KEY=VALUE format

sse/http options:
  -u, --url             MCP server URL
  -H, --header          HTTP headers in 'Key: Value' format
```

## Output

The evaluation script generates a detailed report including:

- **Summary Statistics**:
  - Accuracy (correct/total)
  - Average task duration
  - Average tool calls per task
  - Total tool calls

- **Per-Task Results**:
  - Prompt and expected response
  - Actual response from the agent
  - Whether the answer was correct (✅/❌)
  - Duration and tool call details
  - Agent's summary of its approach
  - Agent's feedback on the tools

### Save Report to File

```bash
python scripts/evaluation.py \
  -t stdio \
  -c python \
  -a my_server.py \
  -o evaluation_report.md \
  evaluation.xml
```

## Complete Example Workflow

Here's a complete example of creating and running an evaluation:

1. **Create your evaluation file** (`my_evaluation.xml`):

```xml
<evaluation>
   <qa_pair>
      <question>Find the user who created the most issues in January 2024. What is their username?</question>
      <answer>alice_developer</answer>
   </qa_pair>
   <qa_pair>
      <question>Among all pull requests merged in Q1 2024, which repository had the highest number? Provide the repository name.</question>
      <answer>backend-api</answer>
   </qa_pair>
   <qa_pair>
      <question>Find the project that was completed in December 2023 and had the longest duration from start to finish. How many days did it take?</question>
      <answer>127</answer>
   </qa_pair>
</evaluation>
```

2. **Install dependencies**:

```bash
pip install -r scripts/requirements.txt
export ANTHROPIC_API_KEY=your_api_key
```

3. **Run evaluation**:

```bash
python scripts/evaluation.py \
  -t stdio \
  -c python \
  -a github_mcp_server.py \
  -e GITHUB_TOKEN=ghp_xxx \
  -o github_eval_report.md \
  my_evaluation.xml
```

4. **Review the report** in `github_eval_report.md` to:
   - See which questions passed/failed
   - Read the agent's feedback on your tools
   - Identify areas for improvement
   - Iterate on your MCP server design

## Troubleshooting

### Connection Errors

If you get connection errors:

- **STDIO**: Verify the command and arguments are correct
- **SSE/HTTP**: Check the URL is accessible and headers are correct
- Ensure any required API keys are set in environment variables or headers

### Low Accuracy

If many evaluations fail:

- Review the agent's feedback for each task
- Check if tool descriptions are clear and comprehensive
- Verify input parameters are well-documented
- Consider whether tools return too much or too little data
- Ensure error messages are actionable

### Timeout Issues

If tasks are timing out:

- Use a more capable model (e.g., `codex-3-7-sonnet-20250219`)
- Check if tools are returning too much data
- Verify pagination is working correctly
- Consider simplifying complex questions
```

.config/mcp/mcp-builder/reference/mcp_best_practices.md
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

.config/mcp/mcp-builder/reference/node_mcp_server.md
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
async function getUser(id: string): Promise<any> {
  return await apiCall(`/users/${id}`);  // No type safety
}
```

## Package Configuration

### package.json

```json
{
  "name": "{service}-mcp-server",
  "version": "1.0.0",
  "description": "MCP server for {Service} API integration",
  "type": "module",
  "main": "dist/index.js",
  "scripts": {
    "start": "node dist/index.js",
    "dev": "tsx watch src/index.ts",
    "build": "tsc",
    "clean": "rm -rf dist"
  },
  "engines": {
    "node": ">=18"
  },
  "dependencies": {
    "@modelcontextprotocol/sdk": "^1.6.1",
    "axios": "^1.7.9",
    "zod": "^3.23.8"
  },
  "devDependencies": {
    "@types/node": "^22.10.0",
    "tsx": "^4.19.2",
    "typescript": "^5.7.2"
  }
}
```

### tsconfig.json

```json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "Node16",
    "moduleResolution": "Node16",
    "lib": ["ES2022"],
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "declaration": true,
    "declarationMap": true,
    "sourceMap": true,
    "allowSyntheticDefaultImports": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

## Complete Example

```typescript
#!/usr/bin/env node
/**
 * MCP Server for Example Service.
 *
 * This server provides tools to interact with Example API, including user search,
 * project management, and data export capabilities.
 */

import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { z } from "zod";
import axios, { AxiosError } from "axios";

// Constants
const API_BASE_URL = "https://api.example.com/v1";
const CHARACTER_LIMIT = 25000;

// Enums
enum ResponseFormat {
  MARKDOWN = "markdown",
  JSON = "json"
}

// Zod schemas
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

type UserSearchInput = z.infer<typeof UserSearchInputSchema>;

// Shared utility functions
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

// Create MCP server instance
const server = new McpServer({
  name: "example-mcp",
  version: "1.0.0"
});

// Register tools
server.registerTool(
  "example_search_users",
  {
    title: "Search Example Users",
    description: `[Full description as shown above]`,
    inputSchema: UserSearchInputSchema,
    annotations: {
      readOnlyHint: true,
      destructiveHint: false,
      idempotentHint: true,
      openWorldHint: true
    }
  },
  async (params: UserSearchInput) => {
    // Implementation as shown above
  }
);

// Main function
// For stdio (local):
async function runStdio() {
  if (!process.env.EXAMPLE_API_KEY) {
    console.error("ERROR: EXAMPLE_API_KEY environment variable is required");
    process.exit(1);
  }

  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("MCP server running via stdio");
}

// For streamable HTTP (remote):
async function runHTTP() {
  if (!process.env.EXAMPLE_API_KEY) {
    console.error("ERROR: EXAMPLE_API_KEY environment variable is required");
    process.exit(1);
  }

  const app = express();
  app.use(express.json());

  app.post('/mcp', async (req, res) => {
    const transport = new StreamableHTTPServerTransport({
      sessionIdGenerator: undefined,
      enableJsonResponse: true
    });
    res.on('close', () => transport.close());
    await server.connect(transport);
    await transport.handleRequest(req, res, req.body);
  });

  const port = parseInt(process.env.PORT || '3000');
  app.listen(port, () => {
    console.error(`MCP server running on http://localhost:${port}/mcp`);
  });
}

// Choose transport based on environment
const transport = process.env.TRANSPORT || 'stdio';
if (transport === 'http') {
  runHTTP().catch(error => {
    console.error("Server error:", error);
    process.exit(1);
  });
} else {
  runStdio().catch(error => {
    console.error("Server error:", error);
    process.exit(1);
  });
}
```

---

## Advanced MCP Features

### Resource Registration

Expose data as resources for efficient, URI-based access:

```typescript
import { ResourceTemplate } from "@modelcontextprotocol/sdk/types.js";

// Register a resource with URI template
server.registerResource(
  {
    uri: "file://documents/{name}",
    name: "Document Resource",
    description: "Access documents by name",
    mimeType: "text/plain"
  },
  async (uri: string) => {
    // Extract parameter from URI
    const match = uri.match(/^file:\/\/documents\/(.+)$/);
    if (!match) {
      throw new Error("Invalid URI format");
    }

    const documentName = match[1];
    const content = await loadDocument(documentName);

    return {
      contents: [{
        uri,
        mimeType: "text/plain",
        text: content
      }]
    };
  }
);

// List available resources dynamically
server.registerResourceList(async () => {
  const documents = await getAvailableDocuments();
  return {
    resources: documents.map(doc => ({
      uri: `file://documents/${doc.name}`,
      name: doc.name,
      mimeType: "text/plain",
      description: doc.description
    }))
  };
});
```

**When to use Resources vs Tools:**
- **Resources**: For data access with simple URI-based parameters
- **Tools**: For complex operations requiring validation and business logic
- **Resources**: When data is relatively static or template-based
- **Tools**: When operations have side effects or complex workflows

### Transport Options

The TypeScript SDK supports two main transport mechanisms:

#### Streamable HTTP (Recommended for Remote Servers)

```typescript
import { StreamableHTTPServerTransport } from "@modelcontextprotocol/sdk/server/streamableHttp.js";
import express from "express";

const app = express();
app.use(express.json());

app.post('/mcp', async (req, res) => {
  // Create new transport for each request (stateless, prevents request ID collisions)
  const transport = new StreamableHTTPServerTransport({
    sessionIdGenerator: undefined,
    enableJsonResponse: true
  });

  res.on('close', () => transport.close());

  await server.connect(transport);
  await transport.handleRequest(req, res, req.body);
});

app.listen(3000);
```

#### stdio (For Local Integrations)

```typescript
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";

const transport = new StdioServerTransport();
await server.connect(transport);
```

**Transport selection:**
- **Streamable HTTP**: Web services, remote access, multiple clients
- **stdio**: Command-line tools, local development, subprocess integration

### Notification Support

Notify clients when server state changes:

```typescript
// Notify when tools list changes
server.notification({
  method: "notifications/tools/list_changed"
});

// Notify when resources change
server.notification({
  method: "notifications/resources/list_changed"
});
```

Use notifications sparingly - only when server capabilities genuinely change.

---

## Code Best Practices

### Code Composability and Reusability

Your implementation MUST prioritize composability and code reuse:

1. **Extract Common Functionality**:
   - Create reusable helper functions for operations used across multiple tools
   - Build shared API clients for HTTP requests instead of duplicating code
   - Centralize error handling logic in utility functions
   - Extract business logic into dedicated functions that can be composed
   - Extract shared markdown or JSON field selection & formatting functionality

2. **Avoid Duplication**:
   - NEVER copy-paste similar code between tools
   - If you find yourself writing similar logic twice, extract it into a function
   - Common operations like pagination, filtering, field selection, and formatting should be shared
   - Authentication/authorization logic should be centralized

## Building and Running

Always build your TypeScript code before running:

```bash
# Build the project
npm run build

# Run the server
npm start

# Development with auto-reload
npm run dev
```

Always ensure `npm run build` completes successfully before considering the implementation complete.

## Quality Checklist

Before finalizing your Node/TypeScript MCP server implementation, ensure:

### Strategic Design
- [ ] Tools enable complete workflows, not just API endpoint wrappers
- [ ] Tool names reflect natural task subdivisions
- [ ] Response formats optimize for agent context efficiency
- [ ] Human-readable identifiers used where appropriate
- [ ] Error messages guide agents toward correct usage

### Implementation Quality
- [ ] FOCUSED IMPLEMENTATION: Most important and valuable tools implemented
- [ ] All tools registered using `registerTool` with complete configuration
- [ ] All tools include `title`, `description`, `inputSchema`, and `annotations`
- [ ] Annotations correctly set (readOnlyHint, destructiveHint, idempotentHint, openWorldHint)
- [ ] All tools use Zod schemas for runtime input validation with `.strict()` enforcement
- [ ] All Zod schemas have proper constraints and descriptive error messages
- [ ] All tools have comprehensive descriptions with explicit input/output types
- [ ] Descriptions include return value examples and complete schema documentation
- [ ] Error messages are clear, actionable, and educational

### TypeScript Quality
- [ ] TypeScript interfaces are defined for all data structures
- [ ] Strict TypeScript is enabled in tsconfig.json
- [ ] No use of `any` type - use `unknown` or proper types instead
- [ ] All async functions have explicit Promise<T> return types
- [ ] Error handling uses proper type guards (e.g., `axios.isAxiosError`, `z.ZodError`)

### Advanced Features (where applicable)
- [ ] Resources registered for appropriate data endpoints
- [ ] Appropriate transport configured (stdio or streamable HTTP)
- [ ] Notifications implemented for dynamic server capabilities
- [ ] Type-safe with SDK interfaces

### Project Configuration
- [ ] Package.json includes all necessary dependencies
- [ ] Build script produces working JavaScript in dist/ directory
- [ ] Main entry point is properly configured as dist/index.js
- [ ] Server name follows format: `{service}-mcp-server`
- [ ] tsconfig.json properly configured with strict mode

### Code Quality
- [ ] Pagination is properly implemented where applicable
- [ ] Large responses check CHARACTER_LIMIT constant and truncate with clear messages
- [ ] Filtering options are provided for potentially large result sets
- [ ] All network operations handle timeouts and connection errors gracefully
- [ ] Common functionality is extracted into reusable functions
- [ ] Return types are consistent across similar operations

### Testing and Build
- [ ] `npm run build` completes successfully without errors
- [ ] dist/index.js created and executable
- [ ] Server runs: `node dist/index.js --help`
- [ ] All imports resolve correctly
- [ ] Sample tool calls work as expected
```

.config/mcp/mcp-builder/reference/python_mcp_server.md
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
            lines.append("")

            for user in users:
                lines.append(f"## {user['name']} ({user['id']})")
                lines.append(f"- **Email**: {user['email']}")
                if user.get('team'):
                    lines.append(f"- **Team**: {user['team']}")
                lines.append("")

            return "\n".join(lines)

        else:
            # Machine-readable JSON format
            import json
            response = {
                "total": total,
                "count": len(users),
                "offset": params.offset,
                "users": users
            }
            return json.dumps(response, indent=2)

    except Exception as e:
        return _handle_api_error(e)

if __name__ == "__main__":
    mcp.run()
```

---

## Advanced FastMCP Features

### Context Parameter Injection

FastMCP can automatically inject a `Context` parameter into tools for advanced capabilities like logging, progress reporting, resource reading, and user interaction:

```python
from mcp.server.fastmcp import FastMCP, Context

mcp = FastMCP("example_mcp")

@mcp.tool()
async def advanced_search(query: str, ctx: Context) -> str:
    '''Advanced tool with context access for logging and progress.'''

    # Report progress for long operations
    await ctx.report_progress(0.25, "Starting search...")

    # Log information for debugging
    await ctx.log_info("Processing query", {"query": query, "timestamp": datetime.now()})

    # Perform search
    results = await search_api(query)
    await ctx.report_progress(0.75, "Formatting results...")

    # Access server configuration
    server_name = ctx.fastmcp.name

    return format_results(results)

@mcp.tool()
async def interactive_tool(resource_id: str, ctx: Context) -> str:
    '''Tool that can request additional input from users.'''

    # Request sensitive information when needed
    api_key = await ctx.elicit(
        prompt="Please provide your API key:",
        input_type="password"
    )

    # Use the provided key
    return await api_call(resource_id, api_key)
```

**Context capabilities:**
- `ctx.report_progress(progress, message)` - Report progress for long operations
- `ctx.log_info(message, data)` / `ctx.log_error()` / `ctx.log_debug()` - Logging
- `ctx.elicit(prompt, input_type)` - Request input from users
- `ctx.fastmcp.name` - Access server configuration
- `ctx.read_resource(uri)` - Read MCP resources

### Resource Registration

Expose data as resources for efficient, template-based access:

```python
@mcp.resource("file://documents/{name}")
async def get_document(name: str) -> str:
    '''Expose documents as MCP resources.

    Resources are useful for static or semi-static data that doesn't
    require complex parameters. They use URI templates for flexible access.
    '''
    document_path = f"./docs/{name}"
    with open(document_path, "r") as f:
        return f.read()

@mcp.resource("config://settings/{key}")
async def get_setting(key: str, ctx: Context) -> str:
    '''Expose configuration as resources with context.'''
    settings = await load_settings()
    return json.dumps(settings.get(key, {}))
```

**When to use Resources vs Tools:**
- **Resources**: For data access with simple parameters (URI templates)
- **Tools**: For complex operations with validation and business logic

### Structured Output Types

FastMCP supports multiple return types beyond strings:

```python
from typing import TypedDict
from dataclasses import dataclass
from pydantic import BaseModel

# TypedDict for structured returns
class UserData(TypedDict):
    id: str
    name: str
    email: str

@mcp.tool()
async def get_user_typed(user_id: str) -> UserData:
    '''Returns structured data - FastMCP handles serialization.'''
    return {"id": user_id, "name": "John Doe", "email": "john@example.com"}

# Pydantic models for complex validation
class DetailedUser(BaseModel):
    id: str
    name: str
    email: str
    created_at: datetime
    metadata: Dict[str, Any]

@mcp.tool()
async def get_user_detailed(user_id: str) -> DetailedUser:
    '''Returns Pydantic model - automatically generates schema.'''
    user = await fetch_user(user_id)
    return DetailedUser(**user)
```

### Lifespan Management

Initialize resources that persist across requests:

```python
from contextlib import asynccontextmanager

@asynccontextmanager
async def app_lifespan():
    '''Manage resources that live for the server's lifetime.'''
    # Initialize connections, load config, etc.
    db = await connect_to_database()
    config = load_configuration()

    # Make available to all tools
    yield {"db": db, "config": config}

    # Cleanup on shutdown
    await db.close()

mcp = FastMCP("example_mcp", lifespan=app_lifespan)

@mcp.tool()
async def query_data(query: str, ctx: Context) -> str:
    '''Access lifespan resources through context.'''
    db = ctx.request_context.lifespan_state["db"]
    results = await db.query(query)
    return format_results(results)
```

### Transport Options

FastMCP supports two main transport mechanisms:

```python
# stdio transport (for local tools) - default
if __name__ == "__main__":
    mcp.run()

# Streamable HTTP transport (for remote servers)
if __name__ == "__main__":
    mcp.run(transport="streamable_http", port=8000)
```

**Transport selection:**
- **stdio**: Command-line tools, local integrations, subprocess execution
- **Streamable HTTP**: Web services, remote access, multiple clients

---

## Code Best Practices

### Code Composability and Reusability

Your implementation MUST prioritize composability and code reuse:

1. **Extract Common Functionality**:
   - Create reusable helper functions for operations used across multiple tools
   - Build shared API clients for HTTP requests instead of duplicating code
   - Centralize error handling logic in utility functions
   - Extract business logic into dedicated functions that can be composed
   - Extract shared markdown or JSON field selection & formatting functionality

2. **Avoid Duplication**:
   - NEVER copy-paste similar code between tools
   - If you find yourself writing similar logic twice, extract it into a function
   - Common operations like pagination, filtering, field selection, and formatting should be shared
   - Authentication/authorization logic should be centralized

### Python-Specific Best Practices

1. **Use Type Hints**: Always include type annotations for function parameters and return values
2. **Pydantic Models**: Define clear Pydantic models for all input validation
3. **Avoid Manual Validation**: Let Pydantic handle input validation with constraints
4. **Proper Imports**: Group imports (standard library, third-party, local)
5. **Error Handling**: Use specific exception types (httpx.HTTPStatusError, not generic Exception)
6. **Async Context Managers**: Use `async with` for resources that need cleanup
7. **Constants**: Define module-level constants in UPPER_CASE

## Quality Checklist

Before finalizing your Python MCP server implementation, ensure:

### Strategic Design
- [ ] Tools enable complete workflows, not just API endpoint wrappers
- [ ] Tool names reflect natural task subdivisions
- [ ] Response formats optimize for agent context efficiency
- [ ] Human-readable identifiers used where appropriate
- [ ] Error messages guide agents toward correct usage

### Implementation Quality
- [ ] FOCUSED IMPLEMENTATION: Most important and valuable tools implemented
- [ ] All tools have descriptive names and documentation
- [ ] Return types are consistent across similar operations
- [ ] Error handling is implemented for all external calls
- [ ] Server name follows format: `{service}_mcp`
- [ ] All network operations use async/await
- [ ] Common functionality is extracted into reusable functions
- [ ] Error messages are clear, actionable, and educational
- [ ] Outputs are properly validated and formatted

### Tool Configuration
- [ ] All tools implement 'name' and 'annotations' in the decorator
- [ ] Annotations correctly set (readOnlyHint, destructiveHint, idempotentHint, openWorldHint)
- [ ] All tools use Pydantic BaseModel for input validation with Field() definitions
- [ ] All Pydantic Fields have explicit types and descriptions with constraints
- [ ] All tools have comprehensive docstrings with explicit input/output types
- [ ] Docstrings include complete schema structure for dict/JSON returns
- [ ] Pydantic models handle input validation (no manual validation needed)

### Advanced Features (where applicable)
- [ ] Context injection used for logging, progress, or elicitation
- [ ] Resources registered for appropriate data endpoints
- [ ] Lifespan management implemented for persistent connections
- [ ] Structured output types used (TypedDict, Pydantic models)
- [ ] Appropriate transport configured (stdio or streamable HTTP)

### Code Quality
- [ ] File includes proper imports including Pydantic imports
- [ ] Pagination is properly implemented where applicable
- [ ] Filtering options are provided for potentially large result sets
- [ ] All async functions are properly defined with `async def`
- [ ] HTTP client usage follows async patterns with proper context managers
- [ ] Type hints are used throughout the code
- [ ] Constants are defined at module level in UPPER_CASE

### Testing
- [ ] Server runs successfully: `python your_server.py --help`
- [ ] All imports resolve correctly
- [ ] Sample tool calls work as expected
- [ ] Error scenarios handled gracefully
```

.config/mcp/mcp-builder/scripts/connections.py
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

.config/mcp/mcp-builder/scripts/evaluation.py
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

.config/mcp/mcp-builder/scripts/example_evaluation.xml
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

.config/mcp/mcp-builder/scripts/requirements.txt
```
anthropic>=0.39.0
mcp>=1.1.0
```

.config/mcp/oraclepack-gold-pack/references/attachment-minimization.md
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

.config/mcp/oraclepack-gold-pack/references/inference-first-discovery.md
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

.config/mcp/oraclepack-gold-pack/references/oracle-pack-template.md
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
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edgecases.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, partial completion, retries, timeouts)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-taxonomy.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the failure-mode taxonomy of this system (timeouts, retries, partial failures, inconsistent state, dependency failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-resilience.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What resilience mechanisms exist (circuit breakers, bulkheads, retries with jitter, rate limiting, graceful degradation)? Where are they configured, and which critical path lacks them?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-inventory.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What feature-flag system exists (or how are conditional rollouts handled)? Inventory the flags (or equivalents) and identify where flags are defined, evaluated, and documented.

Rationale (one sentence):
Flags enable safe rollout and experimentation and reduce release risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-rollout.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Describe the flag/rollout lifecycle (how flags are created, tested, ramped, monitored, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check
--------------

Mark each as `OK` or `Missing(<which step ids>)`:

*   contracts/interfaces: OK

*   invariants: OK

*   caching/state: OK

*   background jobs: OK

*   observability: OK

*   permissions: OK

*   migrations: OK

*   UX flows: OK

*   failure modes: OK

*   feature flags: OK
```

.config/mcp/oraclepack-gold-pack/references/oracle-scratch-format.md
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

.config/mcp/oraclepack-gold-pack/scripts/lint_attachments.py
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

.config/mcp/oraclepack-gold-pack/scripts/validate_pack.py
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

.config/mcp/oraclepack-taskify/assets/action-pack-template.md
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

.config/mcp/oraclepack-taskify/assets/actions-json-schema.md
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

.config/mcp/oraclepack-taskify/assets/prd-synthesis-prompt.md
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

.config/mcp/oraclepack-taskify/references/determinism-and-safety.md
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

.config/mcp/oraclepack-taskify/references/task-master-cli-cheatsheet.md
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

.config/mcp/oraclepack-taskify/references/workflow-overview.md
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

.config/mcp/oraclepack-taskify/scripts/detect-oracle-outputs.sh
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

.config/mcp/oraclepack-taskify/scripts/validate-action-pack.sh
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

.config/skills/oracle-pack/assets/oracle-pack-template.md
```
<!-- path: ~/.codex/skills/oracle-pack/assets/oracle-pack-template.md -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <oracle-out|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02 ...
# ...
# 20 ...
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: <OK|Missing (state missing artifact pattern)>

*   invariants: <OK|Missing (...)>

*   caching/state: <OK|Missing (...)>

*   background jobs: <OK|Missing (...)>

*   observability: <OK|Missing (...)>

*   permissions: <OK|Missing (...)>

*   migrations: <OK|Missing (...)>

*   UX flows: <OK|Missing (...)>

*   failure modes: <OK|Missing (...)>

*   feature flags: <OK|Missing (...)>
```
```

.config/skills/oracle-pack/references/attachment-minimization.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/attachment-minimization.md -->
# Attachment minimization rules (evidence-driven)

Objective: keep Oracle inputs small and high-signal. Attach only what Oracle needs to answer the single question.

## Always exclude

- Secrets (`.env`, key files, tokens)
- Large generated dirs unless explicitly needed (`node_modules`, `dist`, `build`, `coverage`, etc.)
- Broad globs like `src/**` unless the question has `Unknown` reference and no better anchor exists

## Core mapping: reference → files to attach

### Reference = `{path}:{symbol}`

Attach:
1) `path` (the file containing the symbol)

Optionally attach **one** more file (only if it materially affects correctness):
- nearest router/entrypoint that calls into `path`
- config file that wires the subsystem (`config/*`, `settings.*`, `env` loader)
- interface/schema referenced by the symbol (e.g., `openapi.*`, `schema.*`)

Avoid attaching more than 2 files unless the repo uses very small modules and it’s clearly required.

### Reference = `{endpoint}`

Attach:
1) the route map / router file where the endpoint is declared
2) the handler file (if different)

Prefer literal files over globs.

### Reference = `{event}`

Attach:
1) the job/worker registration file where the event is scheduled/declared
2) the worker implementation file (if different)

### Reference = `Unknown`

Attach only “index” evidence:
- README (or closest equivalent)
- primary manifest (`package.json`, `pyproject.toml`, etc.)
- one best-guess entrypoint file (if you found one)

In the prompt, explicitly request the missing artifact pattern to attach next (e.g., “attach `**/migrations/**`”).

## Optional global extras (`extra_files`)

If the user provided `extra_files`, append those attachments to every command *after* the minimal evidence attachments. Use this sparingly.
```

.config/skills/oracle-pack/references/inference-first-discovery.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/inference-first-discovery.md -->
# Inference-first discovery (adaptive search ladder, ck-first)

Goal: avoid wasting time on hard-targeted globs/greps that may not exist by deriving the search plan from repo evidence,
using the codebase-search rules (ck → ast-grep when structure matters → deterministic output shaping).

## Step 0 — Discover the search interface (required)

- Always run: `ck --help`
- Use only flags confirmed by `ck --help`.
- If `ck` indicates indexing/setup is needed, follow its printed instructions.

## Step 1 — Read “index” artifacts first (cheap, high-signal)

Use `fd` to locate these quickly:

- Repo docs/index:
  - `fd -p README.md`
  - `fd -e md . docs | head`
- Manifests/config (choose the ones that exist):
  - `fd -p package.json`
  - `fd -p pyproject.toml`
  - `fd -p go.mod`
  - `fd -p Cargo.toml`
  - `fd -p pom.xml`
  - `fd -p build.gradle`

Read the smallest set of files that explain structure and entrypoints.

## Step 2 — Infer subsystem locations from what you actually saw

Build a short “signals” list from evidence:
- Router/framework name (from dependencies and scripts)
- Job system
- Migration tool
- Feature flag system
- Observability stack
- Cache layer

Then: follow imports/registration code rather than searching the whole repo.

## Step 3 — Derive targeted `ck` queries (conceptual-first)

Start conceptual (meaning-based), scoped to inferred roots:

- Conceptual: `ck --sem "<intent phrase>"`
  - Examples of intent phrases (adapt to the repo’s vocabulary):
    - “route registration” / “router setup”
    - “auth middleware” / “permission check”
    - “queue worker registration”
    - “migration runner” / “schema change”
    - “feature flag evaluation”
    - “telemetry initialization” / “logger setup”

When unsure, use hybrid:
- `ck --hybrid "<intent phrase>"`

For literal hits (exact tokens you already saw), use regex:
- `ck --regex "<literal-or-regex>"`

### Deterministic narrowing

If results are broad/noisy, narrow deterministically:

- Use `ck --jsonl` (if available) and cap results:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- Otherwise cap with `head` and re-run `ck` constrained to the top-hit directories/files.

## Step 4 — Structural search with `ast-grep` (when structure matters)

Use `ast-grep` when you need syntax-aware matching (AST-level), such as:
- finding a specific call pattern (e.g., permission checks around handlers)
- finding registration shapes (e.g., route definitions)
- reducing false positives vs text search

Examples:

- Find code structure:
  - `ast-grep --lang <language> -p '<pattern>'`
- List matching files (cap output):
  - `ast-grep -l --lang <language> -p '<pattern>' | head -n 10`

Use `ck` first to locate candidate files/modules, then `ast-grep` to enforce structure inside those.

## Step 5 — Progressive widening (fallback ladder)

If inference cannot locate a subsystem:

1) Run a small set of conceptual `ck --sem` / `ck --hybrid` queries using generic intent phrases.
2) If still nothing, broaden to repo-wide `ck --hybrid` for that intent.
3) If still nothing, mark evidence for that category as missing and record the most likely missing artifact pattern.

## Step 6 — Early stopping (don’t over-search)

Stop harvesting once:
- you have >= 20 candidate anchors total, and
- every required category has at least one candidate (or a clearly documented “missing artifact pattern”)

Then generate questions; do not keep scanning “just in case”.
```

.config/skills/oracle-pack/references/oracle-scratch-format.md
```
<!-- path: ~/.codex/skills/oracle-pack/references/oracle-scratch-format.md -->
# oracle usage examples

---

## add attachments

```bash
oracle \
  --browser-attachments always \
  --browser-input-timeout 5s \
  -p "Run the UI smoke test." \
  -f "packages/"
```

## copy to clipboard

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

---

## oracle/codefetch

```bash
oracle \
  --browser-attachments always \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

---

## other scratch examples

```bash
oracle \
  -p "Walk through the UI smoke test" \
  --file "src/**/*.ts"
```

```bash
oracle \
  -p "Review these changes and propose fixes" \
  -f "src/**/*.ts,!**/*.test.ts"
```

```bash
oracle \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

```bash
oracle  \
  --prompt "Read packages.md and explain: (1) what each package does, (2) how workflows are executed end-to-end, (3) where to add a new workflow, and (4) the top 5 refactors to reduce coupling." \
  --file "codefetch/packages.md"
```
```

.config/skills/oracle-pack-rpg-infused/assets/oracle-pack-template.md
```
# oracle-pack (RPG-infused)

PRIMARY GOAL
Generate an oracle strategist question pack that is strictly ingestible by the oracle-pack template and oraclepack runner, while embedding an RPG (Repository Planning Graph) fragment in each prompt for graph extraction.

## Pack metadata

- codebase_name: <fill or Unknown>
- constraints: <fill or None>
- non_goals: <fill or None>
- team_size: <fill or Unknown>
- deadline: <fill or Unknown>
- out_dir: <fill or oracle-out>
- oracle_cmd: <fill or oracle>
- oracle_flags: <fill or --files-report>

## Steps (exactly one fenced bash block; exactly 20 steps inside)

```bash
# 01 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/01-<slug>.md" -p $'Strategist question #01\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: []\nPhase: <P0|P1|P2|P3>\n\nReference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>\nCategory: <one of required categories>\nHorizon: <Immediate|Strategic>\nROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)\nQuestion: <question text>\nRationale: <exactly one sentence>\nSmallest experiment today: <exactly one action>\n\nConstraints: <constraints or None>\nNon-goals: <non_goals or None>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 02 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/02-<slug>.md" -p $'Strategist question #02\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: [01]\nPhase: <P0|P1|P2|P3>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 03 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/03-<slug>.md" -p $'Strategist question #03\nRPG:\nFunctionalNode: <Capability>::<Feature> (WHAT)\nStructuralNode: <module-or-dir> :: <public surface> (HOW)\nDependsOn: [01, 02]\nPhase: <P0|P1|P2|P3>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 04 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/04-<slug>.md" -p $'Strategist question #04\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<prior step ids>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 05 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/05-<slug>.md" -p $'Strategist question #05\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 06 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/06-<slug>.md" -p $'Strategist question #06\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 07 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/07-<slug>.md" -p $'Strategist question #07\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 08 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/08-<slug>.md" -p $'Strategist question #08\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 09 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/09-<slug>.md" -p $'Strategist question #09\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 10 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/10-<slug>.md" -p $'Strategist question #10\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 11 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/11-<slug>.md" -p $'Strategist question #11\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 12 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/12-<slug>.md" -p $'Strategist question #12\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 13 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/13-<slug>.md" -p $'Strategist question #13\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 14 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/14-<slug>.md" -p $'Strategist question #14\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 15 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/15-<slug>.md" -p $'Strategist question #15\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 16 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/16-<slug>.md" -p $'Strategist question #16\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 17 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/17-<slug>.md" -p $'Strategist question #17\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 18 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/18-<slug>.md" -p $'Strategist question #18\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 19 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/19-<slug>.md" -p $'Strategist question #19\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
# 20 ROI=<r> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>
<oracle_cmd> <oracle_flags> --write-output "<out_dir>/20-<slug>.md" -p $'Strategist question #20\nRPG:\nFunctionalNode: <...>\nStructuralNode: <...>\nDependsOn: [<...>]\nPhase: <...>\n\nReference: <...>\nCategory: <...>\nHorizon: <...>\nROI: <...>\nQuestion: <...>\nRationale: <...>\nSmallest experiment today: <...>\n\nConstraints: <...>\nNon-goals: <...>\n\nAnswer format:\n\n1) Direct answer (1–4 bullets, evidence-cited)\n2) Risks/unknowns (bullets)\n3) Next smallest concrete experiment (1 action)\n4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.\n'
```

## Coverage check (list required categories exactly; each is OK or Missing(...))

- contracts/interfaces: <OK or Missing(...)>
- invariants: <OK or Missing(...)>
- caching/state: <OK or Missing(...)>
- background jobs: <OK or Missing(...)>
- observability: <OK or Missing(...)>
- permissions: <OK or Missing(...)>
- migrations: <OK or Missing(...)>
- UX flows: <OK or Missing(...)>
- failure modes: <OK or Missing(...)>
- feature flags: <OK or Missing(...)>

```
```

.config/skills/oracle-pack-rpg-infused/references/oracle-pack-spec.md
```
# Oracle Pack (RPG-infused) — Spec

## Strict output contract

1) Produce exactly one Markdown deliverable matching `assets/oracle-pack-template.md` structure exactly.
2) The deliverable must contain exactly ONE fenced bash block.
3) The bash block must contain EXACTLY 20 steps, numbered 01..20 sequentially.
4) Each step header line must match:

`# NN ROI=<roi> impact=<i> confidence=<c> effort=<e> horizon=<Immediate|Strategic> category=<required> reference=<...>`

5) For each step, emit exactly one runnable oracle command:
   - command: `<oracle_cmd>` (default `oracle`)
   - include `<oracle_flags>` (default `--files-report`)
   - include deterministic output file:

`--write-output "<out_dir>/<nn>-<slug>.md"`

6) Compute ROI:

- Pick `impact`, `confidence`, `effort` each as one decimal in `[0.1..1.0]`.
- Compute `ROI = (impact * confidence) / effort` and round to one decimal.

7) Sort all 20 steps by ROI descending; tie-break by lower effort.

8) Include the coverage check section listing the 10 required categories exactly (OK or Missing(...)).

## Required categories (do not add/remove)

- contracts/interfaces
- invariants
- caching/state
- background jobs
- observability
- permissions
- migrations
- UX flows
- failure modes
- feature flags

## RPG infusion requirements (non-breaking)

A) Each strategist prompt (`-p`) MUST include this RPG block:

- FunctionalNode: capability + feature (WHAT)
- StructuralNode: module/file boundary + interface points (HOW)
- DependsOn: a list of prior step IDs, e.g. [01, 03]
- Phase: P0|P1|P2|P3

B) Dependencies MUST ONLY point backwards:
- DependsOn may reference only earlier step IDs.

C) Across the 20 steps, ensure coverage of BOTH semantics:
- At least 10 steps are functional-first (capability/feature discovery).
- At least 10 steps are structural-first (module/file/interface localization).

D) Ensure at least one step explicitly targets each RPG dimension:
- capabilities/features
- file/module boundaries
- public interfaces/contracts
- data flows/state
- dependency edges/topological build order

## Horizon mix (unchanged)

Exactly 12 Immediate and 8 Strategic.

## Attachment minimization (unchanged)

- Read index artifacts first; infer subsystem locations before sweeping.
- Use deterministic ck/ast-grep/fd queries if available; cap results deterministically.
- Attach only minimal evidence files needed for the single question.
- If reference is Unknown: attach only index files (if possible) and state the missing artifact pattern(s).

## Prompt body format inside -p (must include these sections)

Strategist question #NN
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: [<prior step ids, optional>]
Phase: <P0|P1|P2|P3>

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:

1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

## Defaults

If args are missing, use:

- codebase_name=Unknown
- constraints=None
- non_goals=None
- team_size=Unknown
- deadline=Unknown
- out_dir=oracle-out
- oracle_cmd=oracle
- oracle_flags=--files-report
- extra_files=empty

## Final response requirement

In the assistant response:

- Print: `Output file: docs/oracle-pack-YYYY-MM-DD.md`
- Then print the exact same Markdown content (no extra commentary).
```

.config/skills/oracle-pack-rpg-infused/scripts/validate_oracle_pack.py
```
#!/usr/bin/env python3
"""
Validate an oracle-pack (RPG-infused) markdown file against the strict contract.

Usage:
  scripts/validate_oracle_pack.py docs/oracle-pack-YYYY-MM-DD.md
"""

from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import List, Optional, Tuple


REQUIRED_CATEGORIES: List[str] = [
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


@dataclass(frozen=True)
class Step:
    nn: int
    roi: float
    impact: float
    confidence: float
    effort: float
    horizon: str
    category: str
    reference: str
    cmd_line: str


def die(msg: str) -> None:
    print(f"[ERROR] {msg}", file=sys.stderr)
    raise SystemExit(1)


def ok(msg: str) -> None:
    print(f"[OK] {msg}")


def read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8")
    except Exception as e:
        die(f"Failed to read {path}: {e}")
        raise  # unreachable


def find_single_bash_block(md: str) -> str:
    blocks = re.findall(r"```bash\s*\n(.*?)\n```", md, flags=re.DOTALL)
    if len(blocks) != 1:
        die(f"Expected exactly 1 fenced bash block; found {len(blocks)}.")
    # Ensure no other fenced blocks exist
    fence_count = len(re.findall(r"```", md))
    if fence_count != 2:
        die(f"Expected exactly 1 fenced code block total (2 fences); found {fence_count} fences.")
    return blocks[0]


def parse_out_dir(md: str) -> str:
    m = re.search(r"(?m)^\s*-\s*out_dir:\s*([^\s]+)\s*$", md)
    if not m:
        return "oracle-out"
    return m.group(1).strip()


HEADER_RE = re.compile(
    r"""
    ^\n    #\s*\n    (?P<nn>\d{2})
    \s+
    ROI=(?P<roi>\d+(?:\.\d)?)
    \s+
    impact=(?P<impact>[01]\.\d)
    \s+
    confidence=(?P<confidence>[01]\.\d)
    \s+
    effort=(?P<effort>[01]\.\d)
    \s+
    horizon=(?P<horizon>Immediate|Strategic)
    \s+
    category=(?P<category>.+?)
    \s+
    reference=(?P<reference>.+)
    $
    """,
    flags=re.VERBOSE,
)


def parse_steps(bash: str) -> List[Step]:
    lines = [ln.rstrip("\n") for ln in bash.splitlines()]
    steps: List[Step] = []

    i = 0
    while i < len(lines):
        ln = lines[i].strip()
        if not ln:
            i += 1
            continue

        m = HEADER_RE.match(ln)
        if not m:
            die(f"Unexpected line in bash block (expected step header): {lines[i]}")
        nn = int(m.group("nn"))
        roi = float(m.group("roi"))
        impact = float(m.group("impact"))
        confidence = float(m.group("confidence")),
        effort = float(m.group("effort")),
        horizon = m.group("horizon")
        category = m.group("category").strip()
        reference = m.group("reference").strip()

        # Next non-empty, non-comment line must be the single oracle command
        j = i + 1
        cmd_line: Optional[str] = None
        while j < len(lines):
            nxt = lines[j].strip()
            if not nxt:
                j += 1
                continue
            if nxt.startswith("#"):
                die(f"Found comment line where oracle command line was expected after step {nn:02d}.")
            cmd_line = lines[j].strip()
            j += 1
            break
        if cmd_line is None:
            die(f"Missing oracle command line after step header {nn:02d}.")

        # Ensure there are no extra non-empty, non-comment lines before next header
        k = j
        while k < len(lines):
            peek = lines[k].strip()
            if not peek:
                k += 1
                continue
            if HEADER_RE.match(peek):
                break
            if peek.startswith("#"):
                die(f"Unexpected comment line inside step body for {nn:02d}; each step must be exactly two lines.")
            die(f"Unexpected extra command/content line inside step {nn:02d}: {lines[k]}")
        i = k

        steps.append(
            Step(
                nn=nn,
                roi=roi,
                impact=impact,
                confidence=confidence,
                effort=effort,
                horizon=horizon,
                category=category,
                reference=reference,
                cmd_line=cmd_line,
            )
        )

    return steps


def require_20_steps(steps: List[Step]) -> None:
    if len(steps) != 20:
        die(f"Expected exactly 20 steps; found {len(steps)}.")

    nns = [s.nn for s in steps]
    expected = list(range(1, 21))
    if nns != expected:
        die(f"Step numbers must be 01..20 in order; found: {[f'{n:02d}' for n in nns]}")


def validate_scores(steps: List[Step]) -> None:
    for s in steps:
        for name, v in [("impact", s.impact), ("confidence", s.confidence), ("effort", s.effort)]:
            if v < 0.1 or v > 1.0:
                die(f"Step {s.nn:02d} {name} must be in [0.1..1.0]; got {v}.")
            # One decimal check (string form after rounding to 1 decimal should match)
            if abs(v - round(v, 1)) > 1e-9:
                die(f"Step {s.nn:02d} {name} must have exactly one decimal; got {v}.")

        expected_roi = round((s.impact * s.confidence) / s.effort, 1)
        if abs(s.roi - expected_roi) > 0.05:
            die(
                f"Step {s.nn:02d} ROI mismatch: header ROI={s.roi} but computed ROI={expected_roi} "
                f"from (impact*confidence)/effort."
            )
        if abs(s.roi - round(s.roi, 1)) > 1e-9:
            die(f"Step {s.nn:02d} ROI must have exactly one decimal; got {s.roi}.")


def validate_sort_order(steps: List[Step]) -> None:
    for a, b in zip(steps, steps[1:]):
        if b.roi > a.roi + 1e-9:
            die(
                f"Steps must be sorted by ROI descending; step {a.nn:02d} ROI={a.roi} "
                f"precedes step {b.nn:02d} ROI={b.roi}."
            )
        if abs(b.roi - a.roi) < 1e-9 and b.effort < a.effort - 1e-9:
            die(
                f"ROI tie-break must prefer lower effort first; step {a.nn:02d} effort={a.effort} "
                f"precedes step {b.nn:02d} effort={b.effort} with equal ROI={a.roi}."
            )


def validate_horizon_mix(steps: List[Step]) -> None:
    immediate = sum(1 for s in steps if s.horizon == "Immediate")
    strategic = sum(1 for s in steps if s.horizon == "Strategic")
    if immediate != 12 or strategic != 8:
        die(f"Horizon mix must be exactly 12 Immediate and 8 Strategic; got {immediate} Immediate, {strategic} Strategic.")


def validate_categories(steps: List[Step]) -> None:
    allowed = set(REQUIRED_CATEGORIES)
    for s in steps:
        if s.category not in allowed:
            die(f"Step {s.nn:02d} category must be one of the required categories; got: {s.category!r}")

    covered = {s.category for s in steps}
    missing = [c for c in REQUIRED_CATEGORIES if c not in covered]
    if missing:
        die(f"All required categories must appear at least once across the 20 steps. Missing: {missing}")


def validate_commands(steps: List[Step], out_dir: str) -> None:
    for s in steps:
        if "--files-report" not in s.cmd_line:
            die(f"Step {s.nn:02d} oracle command must include --files-report (or oracle_flags equivalent).")

        # --write-output "<out_dir>/<nn>-<slug>.md"
        m = re.search(r'--write-output\s+"([^"]+)"', s.cmd_line)
        if not m:
            die(f"Step {s.nn:02d} oracle command missing --write-output \"...\".")
        out_path = m.group(1)
        expected_prefix = f"{out_dir}/{s.nn:02d}-"
        if not out_path.startswith(expected_prefix) or not out_path.endswith(".md"):
            die(
                f"Step {s.nn:02d} write-output path must start with {expected_prefix!r} and end with '.md'; got {out_path!r}."
            )
        slug = out_path[len(expected_prefix) : -3]
        if not re.fullmatch(r"[a-z0-9]+(?:-[a-z0-9]+)*", slug or ""):
            die(f"Step {s.nn:02d} slug must be kebab-case [a-z0-9-]; got {slug!r}.")

        if " -p " not in f" {s.cmd_line} " and not re.search(r"\s-p\s", s.cmd_line):
            die(f"Step {s.nn:02d} oracle command must include -p with the strategist prompt body.")

        # Minimal prompt-structure checks (best-effort; do not overfit quoting style)
        if f"Strategist question #{s.nn:02d}" not in s.cmd_line:
            die(f"Step {s.nn:02d} -p prompt must include 'Strategist question #{s.nn:02d}'.")
        for required_snip in ["RPG:", "FunctionalNode:", "StructuralNode:", "DependsOn:", "Phase:", "Answer format:"]:
            if required_snip not in s.cmd_line:
                die(f"Step {s.nn:02d} -p prompt missing required section marker: {required_snip}")


def validate_coverage_check_section(md: str) -> None:
    m = re.search(r"(?ms)^## Coverage check.*?\n(.*?)(?:\n## |\Z)", md)
    if not m:
        die("Missing '## Coverage check' section.")

    body = m.group(1).strip("\n")
    lines = [ln.strip() for ln in body.splitlines() if ln.strip()]
    items: List[Tuple[str, str]] = []
    for ln in lines:
        if not ln.startswith("- "):
            continue
        rest = ln[2:]
        if ":" not in rest:
            continue
        cat, status = [x.strip() for x in rest.split(":", 1)]
        items.append((cat, status))

    if not items:
        die("Coverage check section did not contain any '- category: status' items.")

    cats = [c for c, _ in items]
    if cats != REQUIRED_CATEGORIES:
        die(
            "Coverage check categories must appear exactly once each and in the required order.\n"
            f"Expected: {REQUIRED_CATEGORIES}\n"
            f"Found:    {cats}"
        )

    for cat, status in items:
        if status == "OK":
            continue
        if not status.startswith("Missing(") or not status.endswith(")"):
            die(f"Coverage check status for {cat!r} must be 'OK' or 'Missing(...)'; got {status!r}.")


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("markdown_path", help="Path to docs/oracle-pack-YYYY-MM-DD.md")
    args = ap.parse_args()

    path = Path(args.markdown_path)
    if not path.exists():
        die(f"File not found: {path}")

    md = read_text(path)

    bash = find_single_bash_block(md)
    out_dir = parse_out_dir(md)

    steps = parse_steps(bash)
    require_20_steps(steps)
    validate_scores(steps)
    validate_sort_order(steps)
    validate_horizon_mix(steps)
    validate_categories(steps)
    validate_commands(steps, out_dir)
    validate_coverage_check_section(md)

    ok("oracle-pack is valid.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
```

.config/skills/oraclepack-gold-pack/references/attachment-minimization.md
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

.config/skills/oraclepack-gold-pack/references/inference-first-discovery.md
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

.config/skills/oraclepack-gold-pack/references/oracle-pack-template.md
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
Flow maps reveal critical paths and help prioritize work with user impact.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 16) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=UX flows reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/16-ux-flows-edgecases.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #16

Reference: Unknown
Category: UX flows
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
For the primary flows, what are the top edge cases and “gotchas” (validation failures, partial completion, retries, timeouts)? Identify where these cases are handled and where they are missing.

Rationale (one sentence):
Edge-case handling is where many UX and reliability issues originate.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 17) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/17-failure-modes-taxonomy.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #17

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What is the failure-mode taxonomy of this system (timeouts, retries, partial failures, inconsistent state, dependency failures)? Identify where failures are classified/handled and what is surfaced to users/operators.

Rationale (one sentence):
Explicit failure handling prevents incidents and reduces user-facing errors.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 18) ROI=... impact=... confidence=... effort=... horizon=Immediate category=failure modes reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/18-failure-modes-resilience.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #18

Reference: Unknown
Category: failure modes
Horizon: Immediate
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What resilience mechanisms exist (circuit breakers, bulkheads, retries with jitter, rate limiting, graceful degradation)? Where are they configured, and which critical path lacks them?

Rationale (one sentence):
Resilience patterns determine real-world reliability under stress.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 19) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/19-feature-flags-inventory.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #19

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
What feature-flag system exists (or how are conditional rollouts handled)? Inventory the flags (or equivalents) and identify where flags are defined, evaluated, and documented.

Rationale (one sentence):
Flags enable safe rollout and experimentation and reduce release risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 20) ROI=... impact=... confidence=... effort=... horizon=NearTerm category=feature flags reference=Unknown
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/20-feature-flags-rollout.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #20

Reference: Unknown
Category: feature flags
Horizon: NearTerm
ROI: ... (impact=..., confidence=..., effort=...)

Question:
Describe the flag/rollout lifecycle (how flags are created, tested, ramped, monitored, and retired). Identify the minimum guardrails needed to prevent “flag debt.”

Rationale (one sentence):
A disciplined rollout lifecycle reduces long-term complexity and operational risk.

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–10 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (exactly one action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"
```

Coverage check
--------------

Mark each as `OK` or `Missing(<which step ids>)`:

*   contracts/interfaces: OK

*   invariants: OK

*   caching/state: OK

*   background jobs: OK

*   observability: OK

*   permissions: OK

*   migrations: OK

*   UX flows: OK

*   failure modes: OK

*   feature flags: OK
```

.config/skills/oraclepack-gold-pack/references/oracle-scratch-format.md
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

.config/skills/oraclepack-gold-pack/scripts/lint_attachments.py
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
    header_re = re.compile(r"^#\s*(\d{2})\)?\s+")
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

.config/skills/oraclepack-gold-pack/scripts/validate_pack.py
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
    header_re = re.compile(r"^#\s*(\d{2})\)?\s+")
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
    m = re.search(r"--write-output\s+(\"([^\"]+)\"|'([^']+)'|(\S+))", joined)
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

.config/skills/oraclepack-taskify/assets/action-pack-template.md
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

.config/skills/oraclepack-taskify/assets/actions-json-schema.md
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

.config/skills/oraclepack-taskify/assets/prd-synthesis-prompt.md
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

.config/skills/oraclepack-taskify/references/determinism-and-safety.md
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

.config/skills/oraclepack-taskify/references/task-master-cli-cheatsheet.md
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

.config/skills/oraclepack-taskify/references/workflow-overview.md
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

.config/skills/oraclepack-taskify/scripts/detect-oracle-outputs.sh
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

.config/skills/oraclepack-taskify/scripts/validate-action-pack.sh
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

.config/skills/oraclepack-usecase-rpg/references/oracle-pack-template.md
```
<!-- # path: oraclepack-usecase-rpg/references/oracle-pack-template.md -->
# Oracle Pack — {{codebase_name}} (RPG-infused)

## Parsed args
- codebase_name: {{codebase_name}}
- constraints: {{constraints}}
- non_goals: {{non_goals}}
- team_size: {{team_size}}
- deadline: {{deadline}}
- out_dir: {{out_dir}}
- oracle_cmd: {{oracle_cmd}}
- oracle_flags: {{oracle_flags}}
- extra_files: {{extra_files}}

## Commands
```bash
out_dir="{{out_dir}}"

# 01) ROI=... impact=... confidence=... effort=... horizon=Immediate category=contracts/interfaces reference=...
{{oracle_cmd}} \
  {{oracle_flags}} \
  --write-output "{{out_dir}}/01-<slug>.md" \
  {{extra_files}} \
  -p "$(cat <<'PROMPT'
Strategist question #01
RPG:
FunctionalNode: <Capability>::<Feature> (WHAT)
StructuralNode: <module-or-dir> :: <public surface> (HOW)
DependsOn: []
Phase: P0

Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: contracts/interfaces
Horizon: Immediate
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>

Constraints: {{constraints}}
Non-goals: {{non_goals}}

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action)
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.
PROMPT
)"

# 02) ...
# ...
# 20) ...
```

## Coverage check

Mark each as `OK` or `Missing(<which step ids>)`:

*   contracts/interfaces: OK
*   invariants: OK
*   caching/state: OK
*   background jobs: OK
*   observability: OK
*   permissions: OK
*   migrations: OK
*   UX flows: OK
*   failure modes: OK
*   feature flags: OK

```
```

.config/skills/oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
```
# path: oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
from __future__ import annotations

import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple


REQUIRED_CATEGORIES = [
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

BASH_FENCE_RE = re.compile(r"(?s)```bash\n(.*?)\n```")
STEP_HEADER_RE = re.compile(r"^#\s*(\d{2})(?:\)|[\s]+[—-])")
TOKEN_RE = {
    "roi": re.compile(r"\bROI=(\d+(?:\.\d+)?)\b"),
    "impact": re.compile(r"\bimpact=(\d+(?:\.\d+)?)\b"),
    "confidence": re.compile(r"\bconfidence=(\d+(?:\.\d+)?)\b"),
    "effort": re.compile(r"\beffort=(\d+(?:\.\d+)?)\b"),
    "horizon": re.compile(r"\bhorizon=(Immediate|Strategic)\b"),
    "category": re.compile(r"\bcategory=(contracts/interfaces|invariants|caching/state|background jobs|observability|permissions|migrations|UX flows|failure modes|feature flags)\b"),
    "reference": re.compile(r"\breference=([^\s]+)\b"),
}
ORACLE_INVOCATION_RE = re.compile(r"(?m)^\s*(oracle|\$oracle_cmd|\$\{oracle_cmd\})\b")
WRITE_OUTPUT_RE = re.compile(r'--write-output\s+["\\]?([^"\\]+)["\\]?')
PROMPT_FLAG_RE = re.compile(r"\s-\s*p\s|\s-p\s")


@dataclass
class Step:
    step_id: str  # "01"
    header_line: str
    body: str


def die(msg: str) -> None:
    print(f"ERROR: {msg}", file=sys.stderr)
    sys.exit(1)


def warn(msg: str) -> None:
    print(f"WARNING: {msg}", file=sys.stderr)


def parse_bash_block(md: str) -> str:
    matches = list(BASH_FENCE_RE.finditer(md))
    if len(matches) != 1:
        die(f"Expected exactly one fenced bash block starting with ```bash, found {len(matches)}.")
    return matches[0].group(1)


def split_steps(bash: str) -> List[Step]:
    lines = bash.splitlines(True)  # keep \n
    headers: List[Tuple[int, str, str]] = []  # (line_idx, step_id, header_line)
    for i, line in enumerate(lines):
        m = STEP_HEADER_RE.match(line.strip())
        if m:
            headers.append((i, m.group(1), line.rstrip("\n")))

    if len(headers) != 20:
        die(f"Expected exactly 20 step headers, found {len(headers)}.")

    # Ensure sequential 01..20 by appearance
    expected = [f"{i:02d}" for i in range(1, 21)]
    got = [sid for _, sid, _ in headers]
    if got != expected:
        die(f"Step IDs are not sequential 01..20 in order. Got: {got}")

    steps: List[Step] = []
    for idx, (line_i, sid, header_line) in enumerate(headers):
        start = line_i + 1
        end = headers[idx + 1][0] if idx + 1 < len(headers) else len(lines)
        body = "".join(lines[start:end]).strip("\n")
        steps.append(Step(step_id=sid, header_line=header_line, body=body))
        return steps


def parse_float(token: str, s: str, step_id: str) -> float:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return float(m.group(1))


def parse_str(token: str, s: str, step_id: str) -> str:
    m = TOKEN_RE[token].search(s)
    if not m:
        die(f"Step {step_id}: missing token {token} in header.")
    return m.group(1)


def validate_header_tokens(steps: List[Step]) -> Tuple[int, int, Dict[str, int]]:
    immediate = 0
    strategic = 0
    category_counts: Dict[str, int] = {c: 0 for c in REQUIRED_CATEGORIES}

    for st in steps:
        header = st.header_line

        roi = parse_float("roi", header, st.step_id)
        impact = parse_float("impact", header, st.step_id)
        confidence = parse_float("confidence", header, st.step_id)
        effort = parse_float("effort", header, st.step_id)
        horizon = parse_str("horizon", header, st.step_id)
        category = parse_str("category", header, st.step_id)
        _ = parse_str("reference", header, st.step_id)

        for name, val in [("impact", impact), ("confidence", confidence), ("effort", effort)]:
            if not (0.1 <= val <= 1.0):
                die(f"Step {st.step_id}: {name} must be in 0.1..1.0, got {val}.")
            # One decimal requirement (tolerate float formatting issues)
            if round(val, 1) != val:
                die(f"Step {st.step_id}: {name} must have one decimal, got {val}.")

        if effort == 0.0:
            die(f"Step {st.step_id}: effort must be non-zero.")

        expected_roi = (impact * confidence) / effort
        if abs(expected_roi - roi) > 0.02:
            die(
                f"Step {st.step_id}: ROI mismatch. Header ROI={roi}, expected {(impact*confidence)/effort:.2f} "
                f"from impact={impact}, confidence={confidence}, effort={effort}."
            )

        if horizon == "Immediate":
            immediate += 1
        elif horizon == "Strategic":
            strategic += 1
        else:
            die(f"Step {st.step_id}: invalid horizon={horizon}.")

        if category not in category_counts:
            die(f"Step {st.step_id}: category not in required set: {category}")
            category_counts[category] += 1

    return immediate, strategic, category_counts


def validate_step_bodies(steps: List[Step], out_dir_hint: str | None) -> None:
    for st in steps:
        body = st.body

        invocations = ORACLE_INVOCATION_RE.findall(body)
        if len(invocations) < 1:
            die(f"Step {st.step_id}: no oracle invocation detected in body.")
        if len(invocations) > 1:
            warn(f"Step {st.step_id}: multiple oracle-like invocations detected; expected one. Matches={invocations}")

        if "--write-output" not in body:
            die(f"Step {st.step_id}: missing --write-output in oracle invocation.")
        if "-p" not in body:
            die(f"Step {st.step_id}: missing -p prompt body in oracle invocation.")

        m = WRITE_OUTPUT_RE.search(body)
        if not m:
            die(f"Step {st.step_id}: could not parse --write-output path.")
        out_path = m.group(1)

        # Ensure step number is reflected in write-output file path
        if f"/{st.step_id}-" not in out_path and f"{st.step_id}-" not in out_path:
            die(
                f"Step {st.step_id}: --write-output path should include '{st.step_id}-<slug>'. Got: {out_path}"
            )

        # If out_dir provided in prelude, encourage usage (non-fatal)
        if out_dir_hint and out_dir_hint not in out_path and "$out_dir" not in out_path and "${out_dir}" not in out_path:
            warn(f"Step {st.step_id}: write-output path does not appear to use out_dir='{out_dir_hint}' (ok if intentional).")


def extract_out_dir_from_prelude(bash: str) -> str | None:
    # Match: out_dir="oracle-out" or out_dir='oracle-out' or out_dir=oracle-out
    m = re.search(r'(?m)^\s*out_dir\s*=\s*["\\]?([^"\\]+)["\\]?\s*$', bash)
    return m.group(1) if m else None


def validate_coverage(steps: List[Step], category_counts: Dict[str, int], md: str) -> None:
    missing = [c for c, n in category_counts.items() if n == 0]
    if missing:
        die(f"Missing required categories in step headers: {missing}")

    # Coverage check section is a strict requirement for this skill; verify presence.
    if re.search(r'(?im)^##\s*Coverage check\s*$', md) is None:
        warn("Coverage check section heading (## Coverage check) not found (skill expects it).")
        return

    # Light validation that each category appears in the coverage section.
    # (Do not over-parse; formatting may vary.)
    coverage_tail = md.split("## Coverage check", 1)[-1]
    for c in REQUIRED_CATEGORIES:
        if c not in coverage_tail:
            warn(f"Coverage check section does not mention category '{c}' (skill expects OK/Missing lines).")


def main() -> None:
    if len(sys.argv) != 2:
        die("Usage: python3 validate_oraclepack_pack.py <pack.md>")

    path = Path(sys.argv[1])
    if not path.exists():
        die(f"File not found: {path}")

    md = path.read_text(encoding="utf-8")
    bash = parse_bash_block(md)
    steps = split_steps(bash)

    immediate, strategic, category_counts = validate_header_tokens(steps)

    if immediate != 12 or strategic != 8:
        die(f"Horizon mix must be exactly 12 Immediate and 8 Strategic; got Immediate={immediate}, Strategic={strategic}")

    out_dir_hint = extract_out_dir_from_prelude(bash)
    if not out_dir_hint:
        warn('Prelude does not define out_dir="...". oraclepack can derive metadata from this if present.')

    validate_step_bodies(steps, out_dir_hint)
    validate_coverage(steps, category_counts, md)

    print("OK: Pack matches strict oraclepack-usecase-rpg validation rules.")


if __name__ == "__main__":
    main()
```

.config/skills/strategist-questions-oracle-pack/assets/oracle-pack-template.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/assets/oracle-pack-template.md -->
# oracle strategist question pack

---

## parsed args

- codebase_name: <Unknown|value>
- constraints: <None|value>
- non_goals: <None|value>
- team_size: <Unknown|value>
- deadline: <Unknown|value>
- out_dir: <oracle-out|value>
- oracle_cmd: <oracle|value>
- oracle_flags: <--browser-attachments always --files-report|value>
- extra_files: <empty|value>

---

## commands (exactly 20; sorted by ROI desc; ties by lower effort)

```bash
# 01 — ROI=<..> impact=<..> confidence=<..> effort=<..> horizon=<Immediate|Strategic> category=<...> reference=<...>
<oracle_cmd> \
  <oracle_flags> \
  --write-output "<out_dir>/01-<slug>.md" \
  -p "Strategist question #01
Reference: <{path}:{symbol} OR {endpoint} OR {event} OR Unknown>
Category: <one of required categories>
Horizon: <Immediate|Strategic>
ROI: <roi> (impact=<i>, confidence=<c>, effort=<e>)
Question: <question text>
Rationale: <exactly one sentence>
Smallest experiment today: <exactly one action>
Constraints: <constraints or None>
Non-goals: <non_goals or None>

Answer format:
1) Direct answer (1–4 bullets, evidence-cited)
2) Risks/unknowns (bullets)
3) Next smallest concrete experiment (1 action) — may differ from the suggested one if you justify it
4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next." \
  -f "<minimal evidence file 1>" \
  -f "<optional supporting file 2>" \
  <optional extra_files entries...>

# 02 ...
# ...
# 20 ...
```

---

## coverage check (must be satisfied)

*   contracts/interfaces: <OK|Missing (state missing artifact pattern)>

*   invariants: <OK|Missing (...)>

*   caching/state: <OK|Missing (...)>

*   background jobs: <OK|Missing (...)>

*   observability: <OK|Missing (...)>

*   permissions: <OK|Missing (...)>

*   migrations: <OK|Missing (...)>

*   UX flows: <OK|Missing (...)>

*   failure modes: <OK|Missing (...)>

*   feature flags: <OK|Missing (...)>
```
```

.config/skills/strategist-questions-oracle-pack/references/attachment-minimization.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/attachment-minimization.md -->
# Attachment minimization rules (evidence-driven)

Objective: keep Oracle inputs small and high-signal. Attach only what Oracle needs to answer the single question.

## Always exclude

- Secrets (`.env`, key files, tokens)
- Large generated dirs unless explicitly needed (`node_modules`, `dist`, `build`, `coverage`, etc.)
- Broad globs like `src/**` unless the question has `Unknown` reference and no better anchor exists

## Core mapping: reference → files to attach

### Reference = `{path}:{symbol}`

Attach:
1) `path` (the file containing the symbol)

Optionally attach **one** more file (only if it materially affects correctness):
- nearest router/entrypoint that calls into `path`
- config file that wires the subsystem (`config/*`, `settings.*`, `env` loader)
- interface/schema referenced by the symbol (e.g., `openapi.*`, `schema.*`)

Avoid attaching more than 2 files unless the repo uses very small modules and it’s clearly required.

### Reference = `{endpoint}`

Attach:
1) the route map / router file where the endpoint is declared
2) the handler file (if different)

Prefer literal files over globs.

### Reference = `{event}`

Attach:
1) the job/worker registration file where the event is scheduled/declared
2) the worker implementation file (if different)

### Reference = `Unknown`

Attach only “index” evidence:
- README (or closest equivalent)
- primary manifest (`package.json`, `pyproject.toml`, etc.)
- one best-guess entrypoint file (if you found one)

In the prompt, explicitly request the missing artifact pattern to attach next (e.g., “attach `**/migrations/**`”).

## Optional global extras (`extra_files`)

If the user provided `extra_files`, append those attachments to every command *after* the minimal evidence attachments. Use this sparingly.
```

.config/skills/strategist-questions-oracle-pack/references/inference-first-discovery.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/inference-first-discovery.md -->
# Inference-first discovery (adaptive search ladder, ck-first)

Goal: avoid wasting time on hard-targeted globs/greps that may not exist by deriving the search plan from repo evidence,
using the codebase-search rules (ck → ast-grep when structure matters → deterministic output shaping).

## Step 0 — Discover the search interface (required)

- Always run: `ck --help`
- Use only flags confirmed by `ck --help`.
- If `ck` indicates indexing/setup is needed, follow its printed instructions.

## Step 1 — Read “index” artifacts first (cheap, high-signal)

Use `fd` to locate these quickly:

- Repo docs/index:
  - `fd -p README.md`
  - `fd -e md . docs | head`
- Manifests/config (choose the ones that exist):
  - `fd -p package.json`
  - `fd -p pyproject.toml`
  - `fd -p go.mod`
  - `fd -p Cargo.toml`
  - `fd -p pom.xml`
  - `fd -p build.gradle`

Read the smallest set of files that explain structure and entrypoints.

## Step 2 — Infer subsystem locations from what you actually saw

Build a short “signals” list from evidence:
- Router/framework name (from dependencies and scripts)
- Job system
- Migration tool
- Feature flag system
- Observability stack
- Cache layer

Then: follow imports/registration code rather than searching the whole repo.

## Step 3 — Derive targeted `ck` queries (conceptual-first)

Start conceptual (meaning-based), scoped to inferred roots:

- Conceptual: `ck --sem "<intent phrase>"`
  - Examples of intent phrases (adapt to the repo’s vocabulary):
    - “route registration” / “router setup”
    - “auth middleware” / “permission check”
    - “queue worker registration”
    - “migration runner” / “schema change”
    - “feature flag evaluation”
    - “telemetry initialization” / “logger setup”

When unsure, use hybrid:
- `ck --hybrid "<intent phrase>"`

For literal hits (exact tokens you already saw), use regex:
- `ck --regex "<literal-or-regex>"`

### Deterministic narrowing

If results are broad/noisy, narrow deterministically:

- Use `ck --jsonl` (if available) and cap results:
  - `ck --jsonl ... | jq ... | sort | head -n 20`
- Otherwise cap with `head` and re-run `ck` constrained to the top-hit directories/files.

## Step 4 — Structural search with `ast-grep` (when structure matters)

Use `ast-grep` when you need syntax-aware matching (AST-level), such as:
- finding a specific call pattern (e.g., permission checks around handlers)
- finding registration shapes (e.g., route definitions)
- reducing false positives vs text search

Examples:

- Find code structure:
  - `ast-grep --lang <language> -p '<pattern>'`
- List matching files (cap output):
  - `ast-grep -l --lang <language> -p '<pattern>' | head -n 10`

Use `ck` first to locate candidate files/modules, then `ast-grep` to enforce structure inside those.

## Step 5 — Progressive widening (fallback ladder)

If inference cannot locate a subsystem:

1) Run a small set of conceptual `ck --sem` / `ck --hybrid` queries using generic intent phrases.
2) If still nothing, broaden to repo-wide `ck --hybrid` for that intent.
3) If still nothing, mark evidence for that category as missing and record the most likely missing artifact pattern.

## Step 6 — Early stopping (don’t over-search)

Stop harvesting once:
- you have >= 20 candidate anchors total, and
- every required category has at least one candidate (or a clearly documented “missing artifact pattern”)

Then generate questions; do not keep scanning “just in case”.
```

.config/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md
```
<!-- path: ~/.codex/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md -->
# oracle usage examples

---

## add attachments

```bash
oracle \
  --browser-attachments always \
  --browser-input-timeout 5s \
  -p "Run the UI smoke test." \
  -f "packages/"
```

## copy to clipboard

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

---

## oracle/codefetch

```bash
oracle \
  --browser-attachments always \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

---

## other scratch examples

```bash
oracle \
  -p "Walk through the UI smoke test" \
  --file "src/**/*.ts"
```

```bash
oracle \
  -p "Review these changes and propose fixes" \
  -f "src/**/*.ts,!**/*.test.ts"
```

```bash
oracle \
  -p "Using packages.md + attached files, give me a dependency graph and a build order. Then propose 5 repo hygiene improvements." \
  -f "codefetch/packages.md" \
  -f "packages/**/package.json" \
  -f "packages/**/tsconfig.json"
```

```bash
oracle --render --copy \
  -p "Explain the CLI architecture and control flow: index.ts -> registry init -> list/run commands. Identify error-handling gaps, async pitfalls, and how you'd improve UX (help text, exit codes, structured output). Provide proposed code changes." \
  -f "codefetch/packages.md" \
  -f "packages/cli/src/index.ts" \
  -f "packages/cli/src/commands/list.ts" \
  -f "packages/cli/src/commands/run.ts"
```

```bash
oracle  \
  --prompt "Read packages.md and explain: (1) what each package does, (2) how workflows are executed end-to-end, (3) where to add a new workflow, and (4) the top 5 refactors to reduce coupling." \
  --file "codefetch/packages.md"
```

```

</source_code>