Directly “integrating skills into the CLI/MCP” helps, but it does not solve the core failure mode by itself. The real problem is that the CLI/MCP is running against a moving target (skills) with no version pinning, no compatibility gates, and no automated contract testing. The fix is to make skills a versioned, discoverable runtime dependency of oraclepack, with explicit compatibility rules and a lockable install.

What you have today

* Skills already define a stable-ish invocation surface: trailing `KEY=value` args, deterministic scripts like `generate_grouped_packs.py`, and mandatory validators (`validate_pack.py`, `lint_attachments.py`).
* MCP today shells out to the Go CLI (`oraclepack validate`, `oraclepack run`) and has a clear “exec enabled” safety gate (`ORACLEPACK_ENABLE_EXEC`).
* MCP tools are schema-described and discoverable via `tools/list`; that same pattern can be used for skills-as-tools, but you still need versioning/compatibility or you’ll just move the breakage boundary. ([Model Context Protocol][1])

A solution that actually keeps oraclepack in sync with rapidly changing skills

1. Make “skills” a versioned plugin/bundle, not “whatever is on disk”

* Define a skill bundle manifest (per skill, or per bundle) with:

  * `name`, `version`
  * `oraclepack_requires` (a version range)
  * entrypoints to scripts (`generate`, `validate`, `lint`, etc.)
  * an args schema (even if you keep `KEY=value` internally)
  * a content hash (sha256 tree digest) for reproducibility
* This is the same “declare a public API + version it” principle behind SemVer: you need an explicit contract before versions mean anything. ([Semantic Versioning][2])

1. Pin skill versions per run (lockfile), and enforce compatibility at runtime

* Add a lockfile (e.g., `oraclepack.lock.json`) that records:

  * skill name + exact version (or git sha)
  * manifest digest
  * oraclepack CLI version
* On `oraclepack run` / `oraclepack skill run`, oraclepack verifies:

  * installed skill version matches the lock (or fetches it)
  * `oraclepack_requires` matches the current binary
  * otherwise: fail fast with a clear message: “skill X@Y requires oraclepack >=A,<B”
* This eliminates “skills changed under me” failures.

1. Ship an “install/update” mechanism so the CLI can stay stable while skills move fast

* CLI commands:

  * `oraclepack skill list`
  * `oraclepack skill install oraclepack-tickets-pack-grouped@1.4.2` (or `@<gitsha>`)
  * `oraclepack skill run <name> KEY=value...`
  * `oraclepack skill update <name>` (explicit, not automatic)
* Skills can be distributed as:

  * GitHub release tarballs, or
  * a Python package with entry points for discovery (clean plugin model). ([Python Packaging][3])
* Store installed bundles under `~/.oraclepack/skills/<name>/<version>/…` and always run from that immutable path.

1. Unify execution so CLI and MCP cannot drift

* Implement one “skill runner” code path:

  * CLI subcommand uses it
  * MCP tool calls the same runner
* Your MCP server already uses subprocess patterns + safety gates; add “skills” under the same gate as pack execution because skills write files and can invoke tooling.
* Expose via MCP tools:

  * `oraclepack_skill_list`
  * `oraclepack_skill_describe`
  * `oraclepack_skill_run`
    MCP tools are designed to be discoverable with explicit schemas. ([Model Context Protocol][1])

1. Add contract tests that run on every skills change

* When any skill changes, CI runs:

  * `oraclepack skill run oraclepack-tickets-pack-grouped …` against a fixture repo
  * then `oraclepack validate` on every generated pack
* This converts “skills broke runs” into “PR fails before merge.”

Why “just integrate” isn’t sufficient

* If oraclepack always runs “whatever skill checkout you happen to have,” integration simply makes the breakage happen at runtime inside oraclepack instead of outside it.
* What prevents breakage is: versioned artifacts + compatibility gating + lockfiles + contract tests.

Minimal implementation order (fastest path to stability)

1. Add skill manifest + digest + `oraclepack_requires` range.
2. Add `oraclepack skill run` that executes the existing deterministic scripts referenced in the skill (no rewriting).
3. Add lockfile support + install cache.
4. Add MCP `oraclepack_skill_*` tools that call the same runner.
5. Add CI contract tests on skills changes.

If you want, I can draft the exact manifest schema and the CLI/MCP command/tool surfaces to match what your skills already expose (`KEY=value` args, deterministic generators, validators), without forcing you to change how skills are authored.

[1]: https://modelcontextprotocol.io/specification/2025-06-18/server/tools?utm_source=chatgpt.com "Tools"
[2]: https://semver.org/?utm_source=chatgpt.com "Semantic Versioning 2.0.0 | Semantic Versioning"
[3]: https://packaging.python.org/guides/creating-and-discovering-plugins/?utm_source=chatgpt.com "Creating and discovering plugins"
