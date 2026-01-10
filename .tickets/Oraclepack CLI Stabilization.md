Parent Ticket:

* Title: Stabilize oraclepack CLI/MCP by versioning and pinning skills
* Summary: oraclepack CLI/MCP runs against rapidly changing skills without version pinning, compatibility gates, or automated contract tests. Introduce a versioned skill bundle model with manifests, lockfiles, install/update workflows, unified runner paths, MCP tool surfacing, and CI contract tests.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “moving target (skills) with no version pinning, no compatibility gates, and no automated contract testing.”
* Global Constraints:

  * Skill updates should be explicit (“update … explicit, not automatic”).
  * Skills should run from an immutable installed path per version.
  * Skill execution should respect an exec safety gate similar to pack execution (`ORACLEPACK_ENABLE_EXEC`).
* Global Environment:

  * Go CLI: `oraclepack validate`, `oraclepack run`
  * MCP server shells out to Go CLI
  * Skills invoked via trailing `KEY=value` args
  * Existing skill scripts/validators referenced: `generate_grouped_packs.py`, `validate_pack.py`, `lint_attachments.py`
* Global Evidence:

  * Reference links in ticket:

    * Model Context Protocol tools spec
    * Semantic Versioning
    * Python plugin packaging guidance

Split Plan:

* Coverage Map:

  * “integrating skills into the CLI/MCP… does not solve the core failure mode…” → T1
  * “moving target (skills) with no version pinning, no compatibility gates, and no automated contract testing” → T4
  * “make skills a versioned, discoverable runtime dependency… compatibility rules and a lockable install” → T1
  * “Skills already define… trailing `KEY=value` args… deterministic scripts… validators…” → T2
  * “MCP today shells out… safety gate (`ORACLEPACK_ENABLE_EXEC`)” → T5
  * “MCP tools… discoverable via `tools/list`… but… need versioning/compatibility” → T5
  * “Define a skill bundle manifest… `name`, `version`, `oraclepack_requires`, entrypoints, args schema, content hash” → T1
  * “Add a lockfile… records skill name+version (or git sha), manifest digest, oraclepack CLI version” → T4
  * “On `oraclepack run` / `oraclepack skill run`, verify… installed skill matches lock (or fetches)… `oraclepack_requires` matches… fail fast…” → T4
  * “Ship an install/update mechanism… `oraclepack skill list/install/run/update`” → T3
  * “Skills can be distributed as… GitHub release tarballs… or Python package… entry points” → T3
  * “Store installed bundles under `~/.oraclepack/skills/<name>/<version>/…`… immutable path” → T3
  * “Unify execution so CLI and MCP cannot drift… one ‘skill runner’ code path” → T2
  * “Expose via MCP tools: `oraclepack_skill_list/describe/run`” → T5
  * “Add contract tests… CI runs… `oraclepack skill run …`… then `oraclepack validate`…” → T6
  * “Why ‘just integrate’ isn’t sufficient…” → Info-only
  * “Minimal implementation order…” → Info-only
  * “If you want, I can draft…” → Info-only
  * Footnote links [1][2][3] → Info-only (captured in Global Evidence)

* Dependencies:

  * T2 depends on T1 because the runner uses the manifest’s declared entrypoints/metadata.
  * T3 depends on T1 because install layout/versioning uses `name`/`version` and bundle identity.
  * T4 depends on T1 because compatibility checks rely on `oraclepack_requires` and manifest digest.
  * T4 depends on T2 because lock enforcement is exercised by `oraclepack skill run`/run paths.
  * T5 depends on T2 because MCP tools must call the same runner code path.
  * T6 depends on T2 because CI contract tests require `oraclepack skill run`.

* Split Tickets:

```ticket T1
T# Title: Define versioned skill bundle manifest and compatibility contract
Type: enhancement
Target Area: Skills packaging / manifest schema
Summary:
- Introduce a skill bundle manifest that turns “skills on disk” into a versioned plugin/bundle with an explicit contract.
- The manifest must capture versioning and compatibility requirements so oraclepack can validate and gate usage over time.

In Scope:
- Define a skill bundle manifest (per skill or per bundle) including:
  - `name`, `version`
  - `oraclepack_requires` (version range)
  - entrypoints to scripts (`generate`, `validate`, `lint`, etc.)
  - args schema (even if `KEY=value` remains the internal interface)
  - content hash (sha256 tree digest) for reproducibility

Out of Scope:
- Implementing install/update flows
- Implementing lockfile enforcement
- MCP tool exposure
- CI contract tests

Current Behavior (Actual):
- Skills are effectively “whatever is on disk,” and compatibility is not explicitly declared.

Expected Behavior:
- Each skill (or bundle) has a manifest that declares version, compatibility (`oraclepack_requires`), runnable entrypoints, argument schema, and a reproducibility hash.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must support existing `KEY=value` calling convention (schema may exist alongside it).
- Must include a content hash (sha256 tree digest).

Evidence:
- References in ticket:
  - “trailing `KEY=value` args”
  - “entrypoints to scripts (`generate`, `validate`, `lint`, etc.)”
  - “content hash (sha256 tree digest)”

Open Items / Unknowns:
- Manifest location/convention (per skill vs per bundle) not provided.
- Exact args schema format not provided.

Risks / Dependencies:
- Not provided

Acceptance Criteria:
- A manifest schema is specified that includes: `name`, `version`, `oraclepack_requires`, script entrypoints, args schema, and a sha256 tree digest.
- The manifest contract is sufficient for downstream tickets to enforce compatibility and reproducibility.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Define a skill bundle manifest… `name`, `version`… `oraclepack_requires`… entrypoints… args schema… content hash (sha256 tree digest)”
- “Make ‘skills’ a versioned plugin/bundle, not ‘whatever is on disk’”
```

```ticket T2
T# Title: Implement unified skill runner in CLI (`oraclepack skill run`)
Type: enhancement
Target Area: oraclepack CLI (skill runner shared code path)
Summary:
- Add a single “skill runner” execution path that the CLI uses and that MCP can call to prevent drift between interfaces.
- The runner should execute the existing deterministic scripts referenced by skills, without rewriting skills.

In Scope:
- Implement one “skill runner” code path.
- Add CLI support for: `oraclepack skill run <name> KEY=value...`
- Ensure the runner executes existing deterministic scripts referenced in the skill.

Out of Scope:
- Install/update mechanism for distributing skills
- Lockfile creation/enforcement
- MCP tool surfacing (covered separately)
- CI contract tests (covered separately)

Current Behavior (Actual):
- Skills exist with `KEY=value` invocation and deterministic scripts/validators, but there is no described `oraclepack skill run` runner path.

Expected Behavior:
- `oraclepack skill run <name> KEY=value...` runs the skill using manifest-declared entrypoints and supports trailing `KEY=value` args.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- “executes the existing deterministic scripts referenced in the skill (no rewriting)”
- Must preserve existing `KEY=value` usage pattern.

Evidence:
- Script/validator examples referenced: `generate_grouped_packs.py`, `validate_pack.py`, `lint_attachments.py`
- CLI command shape referenced: `oraclepack skill run <name> KEY=value...`

Open Items / Unknowns:
- How skills are located/resolved (installed vs local checkout) not provided here (install is a separate ticket).

Risks / Dependencies:
- Depends on T1 for manifest/entrypoint definitions.

Acceptance Criteria:
- CLI supports `oraclepack skill run <name> KEY=value...`.
- Runner executes the deterministic scripts referenced by the skill (no rewriting requirement preserved).
- A single runner path exists that can be called by both CLI and (later) MCP tools.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Implement one ‘skill runner’ code path”
- “Add `oraclepack skill run` that executes the existing deterministic scripts… (no rewriting)”
- “trailing `KEY=value` args… deterministic scripts… validators”
```

```ticket T3
T# Title: Add skill install/update workflow and immutable per-version install cache
Type: enhancement
Target Area: oraclepack CLI (skill distribution + local cache)
Summary:
- Provide explicit skill install/update commands so oraclepack can remain stable while skills move fast.
- Ensure installed skills are stored by name/version under an immutable path and executed from that location.

In Scope:
- Add CLI commands:
  - `oraclepack skill list`
  - `oraclepack skill install <name>@<version|gitsha>`
  - `oraclepack skill update <name>` (explicit, not automatic)
- Store installed bundles under: `~/.oraclepack/skills/<name>/<version>/…`
- Ensure skill execution runs from the immutable installed path.

Out of Scope:
- Lockfile compatibility enforcement (separate ticket)
- MCP tool exposure (separate ticket)
- CI contract tests (separate ticket)

Current Behavior (Actual):
- Not provided (ticket text implies skills are not managed as versioned installs).

Expected Behavior:
- Skills can be installed and updated explicitly, and executions reference the installed, versioned, immutable location.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Updates must be explicit (“not automatic”).
- Distribution options mentioned:
  - GitHub release tarballs, or
  - Python package with entry points for discovery (approach not specified)

Evidence:
- Command list in ticket:
  - `oraclepack skill list/install/run/update`
- Install path in ticket:
  - `~/.oraclepack/skills/<name>/<version>/…`

Open Items / Unknowns:
- Which distribution mechanism(s) must be supported (GitHub tarballs vs Python package entry points) not specified.
- Whether install supports version only, git sha only, or both beyond the examples is not specified.

Risks / Dependencies:
- Depends on T1 for `name`/`version` manifest metadata alignment.

Acceptance Criteria:
- CLI provides `oraclepack skill list`, `oraclepack skill install <name>@…`, and `oraclepack skill update <name>`.
- Installed skills are stored under `~/.oraclepack/skills/<name>/<version>/…`.
- Skill runs use the immutable installed path (no “whatever is on disk” execution for installed skills).

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Ship an ‘install/update’ mechanism… `oraclepack skill list/install/run/update`”
- “Store installed bundles under `~/.oraclepack/skills/<name>/<version>/…` and always run from that immutable path.”
- “update… (explicit, not automatic)”
```

```ticket T4
T# Title: Implement skill lockfile pinning and runtime compatibility enforcement
Type: enhancement
Target Area: oraclepack runtime (lockfile + compatibility gates)
Summary:
- Eliminate “skills changed under me” failures by pinning skill versions per run and enforcing compatibility at runtime.
- Introduce a lockfile that records the exact skill version (or git sha), manifest digest, and the oraclepack CLI version.

In Scope:
- Add lockfile support (example name given: `oraclepack.lock.json`) that records:
  - skill name + exact version (or git sha)
  - manifest digest
  - oraclepack CLI version
- On `oraclepack run` / `oraclepack skill run`, verify:
  - installed skill version matches the lock (or fetches it)
  - `oraclepack_requires` matches the current binary
  - fail fast with a clear compatibility message when mismatched

Out of Scope:
- Defining the manifest schema itself (separate ticket)
- MCP tool exposure (separate ticket)
- CI contract tests (separate ticket)

Current Behavior (Actual):
- No version pinning, compatibility gates, or automated contract testing; the runtime can drift as skills change.

Expected Behavior:
- Runs are reproducible against pinned skill versions/digests and are blocked when compatibility constraints are violated.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must validate `oraclepack_requires` against current binary.
- Must provide a clear failure message on mismatch (example included in ticket text).

Evidence:
- “Add a lockfile… records… skill name + exact version (or git sha)… manifest digest… oraclepack CLI version”
- “fail fast… ‘skill X@Y requires oraclepack >=A,<B’”
- “no version pinning, no compatibility gates”

Open Items / Unknowns:
- Exact lockfile schema/format beyond listed fields not provided.
- Definition of “fetches it” behavior (sources, auth, etc.) not provided.

Risks / Dependencies:
- Depends on T1 for `oraclepack_requires` and manifest digest definition.
- Depends on T2 because enforcement occurs on `oraclepack skill run`/run paths.

Acceptance Criteria:
- A lockfile exists and records: skill name, exact version or git sha, manifest digest, and oraclepack CLI version.
- On run, oraclepack verifies installed skill version matches lock (or fetches it) and checks `oraclepack_requires` compatibility.
- On incompatibility, oraclepack fails fast with a clear message indicating required oraclepack version range.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Pin skill versions per run (lockfile), and enforce compatibility at runtime”
- “Add a lockfile (e.g., `oraclepack.lock.json`) that records…”
- “This eliminates ‘skills changed under me’ failures.”
```

```ticket T5
T# Title: Expose skill management and execution via MCP tools using the shared runner
Type: enhancement
Target Area: MCP server (skills-as-tools)
Summary:
- Add MCP tools for listing/describing/running skills, ensuring the MCP path calls the same runner as the CLI to prevent drift.
- Ensure skill execution is governed by an exec-enabled safety gate comparable to pack execution.

In Scope:
- Expose MCP tools:
  - `oraclepack_skill_list`
  - `oraclepack_skill_describe`
  - `oraclepack_skill_run`
- Ensure MCP tools call the same shared runner code path as the CLI.
- Add “skills” under the same exec safety gate approach used for pack execution, since skills can write files and invoke tooling.

Out of Scope:
- Implementing the manifest schema (separate ticket)
- Implementing install/update mechanism (separate ticket)
- Lockfile enforcement (separate ticket)
- CI contract tests (separate ticket)

Current Behavior (Actual):
- MCP shells out to Go CLI (`oraclepack validate`, `oraclepack run`) and uses an exec-enabled safety gate (`ORACLEPACK_ENABLE_EXEC`).
- MCP tools are schema-described and discoverable via `tools/list`.

Expected Behavior:
- MCP exposes skill list/describe/run as discoverable tools and uses the same runner implementation as the CLI.
- Skill execution via MCP is gated similarly to pack execution.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must reuse the shared runner to avoid CLI/MCP drift.
- Must use an exec safety gate model (explicitly referenced as existing for pack execution).

Evidence:
- “MCP today shells out… and has… safety gate (`ORACLEPACK_ENABLE_EXEC`).”
- “Expose via MCP tools… `oraclepack_skill_list/describe/run`”
- “Implement one ‘skill runner’ code path… MCP tool calls the same runner”
- “MCP tools… discoverable via `tools/list`”

Open Items / Unknowns:
- Exact MCP tool schemas/arguments not provided.

Risks / Dependencies:
- Depends on T2 for the shared runner functionality.

Acceptance Criteria:
- MCP server exposes `oraclepack_skill_list`, `oraclepack_skill_describe`, and `oraclepack_skill_run`.
- MCP skill execution calls the same runner code path as the CLI.
- Skill execution is gated under an exec-enabled safety control comparable to existing pack execution gating.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “Expose via MCP tools: `oraclepack_skill_list`… `oraclepack_skill_describe`… `oraclepack_skill_run`”
- “Implement one ‘skill runner’ code path… MCP tool calls the same runner”
- “add ‘skills’ under the same gate as pack execution”
```

```ticket T6
T# Title: Add CI contract tests to detect skill-induced run breakages before merge
Type: tests
Target Area: CI / automated contract testing
Summary:
- Convert “skills broke runs” into “PR fails before merge” by running contract tests whenever skills change.
- Tests should run a representative skill against a fixture repo and validate generated packs.

In Scope:
- When any skill changes, CI runs:
  - `oraclepack skill run oraclepack-tickets-pack-grouped …` against a fixture repo
  - `oraclepack validate` on every generated pack

Out of Scope:
- Defining fixture repo content (not specified)
- Implementing skill runner/install/lockfile mechanisms (handled in other tickets)

Current Behavior (Actual):
- No automated contract testing is in place; breakages can surface at runtime when skills change.

Expected Behavior:
- Skill changes are validated via CI by executing a skill and validating outputs, failing the PR on regressions.

Reproduction Steps:
- Not provided

Requirements / Constraints:
- Must run on “any skill changes” (triggering mechanism details not provided).
- Must run both skill execution and pack validation steps.

Evidence:
- “Add contract tests that run on every skills change”
- “CI runs… `oraclepack skill run …`… then `oraclepack validate`…”

Open Items / Unknowns:
- Fixture repo location and exact arguments for the skill run not provided.
- How CI detects “any skill changes” not specified.

Risks / Dependencies:
- Depends on T2 for `oraclepack skill run` availability.

Acceptance Criteria:
- CI executes the specified flow on skill changes:
  - runs the referenced skill against a fixture repo
  - validates every generated pack
- A failing skill run or validation fails the PR/CI job.

Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided

Source:
- “When any skill changes, CI runs… `oraclepack skill run …`… then `oraclepack validate`…”
- “This converts ‘skills broke runs’ into ‘PR fails before merge.’”
```
