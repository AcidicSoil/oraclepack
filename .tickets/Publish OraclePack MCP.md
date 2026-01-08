Parent Ticket:

* Title: Publish/distribute `oraclepack-mcp-server` to avoid long-form MCP client configuration
* Summary: Replace the hardcoded venv interpreter path in MCP client configs with a portable, short config, and optionally enable one-click installation for supported clients.
* Source:

  * Link/ID (if present) or “Not provided”: `oraclepack-op-mcp.md`
  * Original ticket excerpt (≤25 words) capturing the overall theme: “publish this so we do not have to use the long form configuration for configuring mcp clients”
* Global Constraints:

  * Eliminate reliance on an absolute venv interpreter path in MCP client configuration.
  * Preserve required env variables (`ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`) in examples.
* Global Environment:

  * Unknown
* Global Evidence:

  * Current MCP client config example (shows venv path + args + env).

Split Plan:

* Coverage Map:

  * Original item: “how do I publish this so we do not have to use the long form configuration for configuring mcp clients?”

    * Assigned Ticket ID: T1
  * Original item: Current config uses venv interpreter path: `"command": "/home/user/.../venv/bin/python"`

    * Assigned Ticket ID: T1
  * Original item: Current args: `["-m", "oraclepack_mcp_server", "--transport", "stdio"]`

    * Assigned Ticket ID: T1
  * Original item: Current env vars: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`

    * Assigned Ticket ID: T1
  * Original item: Option 1: “Publish a Python package so the MCP command is just a PATH executable” + `uv build`, `uv publish`, `uv tool install ...`

    * Assigned Ticket ID: T1
  * Original item: Option 1 config snippet (command becomes `oraclepack-mcp-server`, args `--transport stdio`, env preserved)

    * Assigned Ticket ID: T1
  * Original item: Note: “If you want to reduce env too, prefer absolute paths…”

    * Assigned Ticket ID: T1
  * Original item: Option 2: “No install short config: run via uvx” + config snippet using `"command": "uvx"`

    * Assigned Ticket ID: T2
  * Original item: Option 2 note: aligns with `server.json` PyPI example using `runtimeHint: "uvx"`

    * Assigned Ticket ID: T4
  * Original item: Option 3a: “Claude Desktop: ship a .mcpb bundle” + `mcpb init`, `mcpb pack` + distribute `.mcpb`

    * Assigned Ticket ID: T3
  * Original item: Option 3b: “publish to Official MCP Registry (and GitHub MCP Registry)” via `server.json` and `mcp-publisher` steps

    * Assigned Ticket ID: T4
  * Original item: Recommendation section (choose Option 1 for no venv path; Option 3 for no manual config)

    * Assigned Ticket ID: Info-only
  * Original item: Note: “standardize the executable name to match the PyPI identifier”

    * Assigned Ticket ID: T1
* Dependencies:

  * T2 depends on T1 because the `uvx` approach runs the published package name (`oraclepack-mcp-server`).
  * T4 depends on T1 because the described registry publishing path references a PyPI stdio server example.
* Split Tickets:

```ticket T1
T# Title: Publish `oraclepack-mcp-server` as a PATH executable (PyPI + uv tools) and update config example
Type: enhancement
Target Area: Distribution/packaging for MCP server (`oraclepack-mcp-server`) + MCP client config examples
Summary:
- Publish the MCP server as a Python package so MCP clients can invoke it via a normal command on PATH instead of a venv interpreter path.
- Provide the shorter MCP client configuration example that uses `command: "oraclepack-mcp-server"` and preserves required env vars.
In Scope:
- Publish steps using uv:
  - `uv build`
  - `uv publish`
- Install guidance via uv tools:
  - `uv tool install oraclepack-mcp-server`
- Update the MCP client config example to:
  - `"command": "oraclepack-mcp-server"`
  - `"args": ["--transport", "stdio"]`
  - Keep env: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
- Naming guidance: standardize executable name to match the PyPI identifier (e.g., `oraclepack-mcp-server`).
- Guidance note: prefer absolute paths for env values if trying to reduce/env-stabilize in hosts with undefined working directory.
Out of Scope:
- “One-click install” packaging and registry publishing (handled in other tickets).
Current Behavior (Actual):
- MCP client config points at a venv interpreter path and runs `-m oraclepack_mcp_server`:
  - `"command": "/home/user/.../venv/bin/python"`
Expected Behavior:
- MCP client config can run a PATH command directly:
  - `"command": "oraclepack-mcp-server"`
  - No venv absolute path required.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Preserve required env variables in examples:
  - `ORACLEPACK_BIN`
  - `ORACLEPACK_ALLOWED_ROOTS`
  - `ORACLEPACK_ENABLE_EXEC`
Evidence:
- Current config snippet includes venv interpreter path and env vars:
  - `"command": "/home/user/projects/temp/oraclepack/oraclepack-mcp-server/venv/bin/python"`
  - `"args": ["-m", "oraclepack_mcp_server", "--transport", "stdio"]`
  - `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
Open Items / Unknowns:
- Package metadata and repository details for publishing (Unknown / Not provided).
- Desired final executable name if it differs from `oraclepack-mcp-server` (Unknown / Not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A published distribution path exists that does not require MCP clients to reference a venv interpreter path.
- Documentation/config example shows:
  - `"command": "oraclepack-mcp-server"`
  - `"args": ["--transport", "stdio"]`
  - env variables preserved as shown in the source text.
- Executable naming guidance is documented (“match the PyPI identifier”).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “publish/distribute paths that eliminate the venv absolute path”
- “Publish a Python package so the MCP `command` is just a PATH executable”
- “uv build / uv publish … uv tool install oraclepack-mcp-server”
```

```ticket T2
T# Title: Add “no install” MCP config option using `uvx`
Type: docs
Target Area: MCP client configuration documentation/examples
Summary:
- Provide a short MCP client configuration that runs the server via `uvx` so users don’t need a pre-created local venv path in config.
- Keep required environment variables in the example config.
In Scope:
- Document the `uvx`-based MCP client config example:
  - `"command": "uvx"`
  - `"args": ["oraclepack-mcp-server", "--transport", "stdio"]`
  - env: `ORACLEPACK_BIN`, `ORACLEPACK_ALLOWED_ROOTS`, `ORACLEPACK_ENABLE_EXEC`
Out of Scope:
- Publishing to PyPI (handled in T1).
- Registry publishing via `server.json`/`mcp-publisher` (handled in T4).
Current Behavior (Actual):
- Config is “long-form” due to a venv interpreter path.
Expected Behavior:
- Users can use a short config that invokes `uvx` with the package name and stdio transport args.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Preserve required env variables in the example config.
Evidence:
- Option 2 config snippet:
  - `"command": "uvx"`
  - `"args": ["oraclepack-mcp-server", "--transport", "stdio"]`
Open Items / Unknowns:
- Whether target MCP clients/hosts support `uvx` invocation in their MCP server configuration (Unknown / Not provided).
Risks / Dependencies:
- Depends on T1 (published package name referenced by `uvx`).
Acceptance Criteria:
- Documentation includes the `uvx` config snippet exactly as described in the source text.
- Documentation explicitly retains the required env variable keys used in the source text.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “No install short config: run via `uvx`”
- `"command": "uvx", "args": ["oraclepack-mcp-server", "--transport", "stdio"]`
```

```ticket T3
T# Title: Create and distribute a `.mcpb` bundle for Claude Desktop installation
Type: enhancement
Target Area: MCP Bundle packaging for Claude Desktop distribution
Summary:
- Package the local MCP server as a `.mcpb` bundle so users can install via a UI flow in supported clients (Claude Desktop mentioned).
- Document the bundle creation commands and distribution approach.
In Scope:
- Use MCPB tooling steps as described:
  - `npm install -g @anthropic-ai/mcpb`
  - `mcpb init`
  - `mcpb pack`
- Distribute the resulting `.mcpb` (example given: GitHub Releases).
- Document that users install via Claude Desktop Extensions UI flow (per source text).
Out of Scope:
- PyPI publishing and `uv` tooling approach (handled in T1).
- Official/GitHub MCP registry publishing (handled in T4).
Current Behavior (Actual):
- Users must configure MCP clients manually with JSON.
Expected Behavior:
- Users can install the server via a `.mcpb` bundle in clients that support MCP bundles (Claude Desktop mentioned).
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Follow the described `.mcpb` workflow (init + pack).
Evidence:
- “Claude Desktop: ship a `.mcpb` bundle … mcpb init … mcpb pack … Distribute the resulting `.mcpb`”
Open Items / Unknowns:
- Bundle manifest contents and exact server entrypoints required by MCPB for this server (Unknown / Not provided).
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A `.mcpb` bundle can be produced using the documented CLI steps.
- Documentation explains how to obtain the `.mcpb` (distribution channel mentioned) and install it in Claude Desktop (Extensions UI flow mentioned).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Claude Desktop: ship a `.mcpb` bundle”
- “mcpb init … mcpb pack”
- “Distribute the resulting `.mcpb` (e.g., GitHub Releases)”
```

```ticket T4
T# Title: Publish `server.json` via `mcp-publisher` for MCP Registry / GitHub MCP Registry discovery
Type: enhancement
Target Area: Registry publishing metadata (`server.json`) + publishing workflow (`mcp-publisher`)
Summary:
- Enable “one-click install” in supported clients by publishing a `server.json` descriptor via `mcp-publisher`, targeting the Official MCP Registry and GitHub MCP Registry (as described).
- Document the high-level publishing steps and ownership proof requirement as stated.
In Scope:
- Generate `server.json` using `mcp-publisher init`.
- Follow the publishing sequence as described:
  1) Install `mcp-publisher`
  2) `mcp-publisher init` to generate `server.json`
  3) Prove package ownership (PyPI: add `mcp-name: ...` to README)
  4) `mcp-publisher login github`
  5) `mcp-publisher publish`
- Ensure `server.json` aligns with the described PyPI stdio example capabilities (mentions `environmentVariables` and `runtimeHint: "uvx"`).
Out of Scope:
- `.mcpb` bundling (handled in T3).
- Core PyPI package publishing steps (handled in T1).
Current Behavior (Actual):
- Users must manually configure MCP clients using JSON and local paths.
Expected Behavior:
- Server is discoverable/installable via registry mechanisms described (VS Code / ecosystem via registry).
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Include the described package ownership proof metadata for PyPI (README `mcp-name: ...`).
Evidence:
- “publish a `server.json` via `mcp-publisher`”
- Step list including `mcp-publisher init`, README `mcp-name: ...`, `login github`, `publish`
- Note about PyPI example supporting `environmentVariables` and `runtimeHint: "uvx"`
Open Items / Unknowns:
- Final server identifier/name to use for `mcp-name: ...` (Unknown / Not provided).
- Which registries/clients are in scope beyond “VS Code / ecosystem” phrasing (Unknown / Not provided).
Risks / Dependencies:
- Depends on T1 if the published registry entry targets a PyPI package distribution (as described).
Acceptance Criteria:
- A `server.json` exists generated/maintained via `mcp-publisher init` per the described workflow.
- Documentation includes the stated publishing steps and the PyPI ownership proof requirement (`mcp-name: ...` in README).
- Documentation notes the described `runtimeHint: "uvx"` alignment for the PyPI stdio example.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “publish to the Official MCP Registry (and GitHub MCP Registry)”
- “publish a `server.json` via `mcp-publisher`”
- “Prove package ownership … add `mcp-name: ...` to your README”
```
