Parent Ticket:

* Title: Keep oraclepack in parity with upstream oracle updates
* Summary: Define and automate a process so oraclepack (wrapper) stays compatible as upstream oracle changes. Approach centers on pinning the upstream oracle version, minimizing wrapper coupling, adding contract tests, and automating upgrade PRs + canary checks.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “keep parity with updates with a target repo as their repo updates/changes… wrapper around oracle… keep… up to date”
* Global Constraints:

  * Not provided
* Global Environment:

  * Node 22+ (per referenced upstream distribution note)
  * GitHub Actions (Dependabot + scheduled workflows)
  * Go toolchain used in CI example (version shown as 1.24.x)
* Global Evidence:

  * Dependabot configuration requires `.github/dependabot.yml` and supports npm + GitHub Actions ecosystems. ([GitHub Docs][1])
  * Renovate npm manager documentation (alternative automation). ([Renovate Docs][2])
  * peter-evans/create-pull-request GitHub Action (PR automation). ([GitHub][3])
  * Upstream Sync marketplace action (fork-sync scenario). ([GitHub][4])
  * Git submodule/subtree background (vendor option scenario). ([Atlassian][5])
  * Full original discussion text.

Split Plan:

* Coverage Map:

  * Original item: “treat upstream as a versioned dependency… automate detection/compat testing/PRs/releases” → Assigned Ticket ID: T1
  * Original item: “pins `@steipete/oracle`… runs deterministically” → Assigned Ticket ID: T1
  * Original item: “log/record the exact `oracle` version used into its `.state.json`/`.report.json`” → Assigned Ticket ID: T1
  * Original item: “Vendor via Node workspace… execute `node_modules/.bin/oracle`” → Assigned Ticket ID: T1
  * Original item: “Or run `npx -y @steipete/oracle@<pinned>`” → Assigned Ticket ID: T1
  * Original item: “Reduce parity surface area… pass-through custom args/flags forwarded unchanged” → Assigned Ticket ID: T2
  * Original item: “Only model/enumerate upstream flags where you must (Overrides Wizard)” → Assigned Ticket ID: T2
  * Original item: “If you do enumerate, auto-discover rather than hardcode” → Assigned Ticket ID: T3
  * Original item: “Add a contract test suite that detects upstream CLI surface changes early” → Assigned Ticket ID: T4
  * Original item: “Snapshot tests… capture `oracle --help` and `oracle <critical-subcommand> --help`… diff snapshots in CI” → Assigned Ticket ID: T4
  * Original item: “Behavioral ‘golden path’ tests… run a small pack fixture… assert only on wrapper-owned artifacts” → Assigned Ticket ID: T4
  * Original item: “Automate updates: Dependabot/Renovate” → Assigned Ticket ID: T5
  * Original item: “Scheduled ‘canary’ workflow… run against `@steipete/oracle@latest`” → Assigned Ticket ID: T7
  * Original item: “Release discipline: compatibility statement… gate releases on contract tests + canary” → Assigned Ticket ID: T8
  * Original item: “Concrete GitHub config examples: `.github/dependabot.yml`” → Assigned Ticket ID: T5
  * Original item: “Concrete workflow example: `.github/workflows/bump-oracle.yml` using `peter-evans/create-pull-request`” → Assigned Ticket ID: T6
  * Original item: “If you want ‘sync the fork with upstream’… use fork-sync/upstream-sync workflow/action” → Assigned Ticket ID: Info-only
  * Original item: “If you ever decide to vendor upstream code… use subtree/submodule… submodules don’t auto-track branches” → Assigned Ticket ID: Info-only
  * Original item: “Minimal next step… pinned version file + CI contract suite + enable Dependabot weekly updates” → Assigned Ticket ID: T1 (pin), T4 (contract suite), T5 (Dependabot)
* Dependencies:

  * T4 depends on T1 because contract tests are described as running against a pinned/deterministic upstream version.
  * T6 depends on T4 because the bump workflow is described as running unit + contract tests before opening a PR.
  * T7 depends on T4 because the canary job is described as running the integration/contract suite.
  * T8 depends on T4 and T7 because release gating is described as using contract tests and the canary signal.
* Split Tickets:

```ticket T1
T# Title: Pin upstream oracle version and record the runtime oracle version used
Type: enhancement
Target Area: Dependency boundary / runtime invocation / run artifacts (.state.json/.report.json)
Summary:
  Establish a deterministic upstream boundary by pinning the upstream `@steipete/oracle` version used by oraclepack runs. Ensure the exact oracle version used is recorded in wrapper-owned artifacts so reports/bugs are attributable to a specific upstream version.
In Scope:
  - Add a single source of truth for the pinned upstream oracle version (example given: `tools/oracle-version.txt`).
  - Ensure oraclepack uses the pinned `@steipete/oracle` version deterministically at runtime.
  - Record/log the exact oracle version used into `.state.json` and/or `.report.json` (as referenced).
  - Support one of the described execution modes:
    - Node workspace vendoring (`node_modules/.bin/oracle`), or
    - `npx -y @steipete/oracle@<pinned>` execution.
Out of Scope:
  - Upstream fork syncing and vendoring upstream source code (subtree/submodule) approaches.
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - oraclepack runs use a deterministic pinned oracle version rather than an implicitly installed/global/latest version.
  - oraclepack run artifacts include the upstream oracle version used.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Upstream oracle is referenced as an npm package `@steipete/oracle` and expects Node 22+ (as noted).
Evidence:
  - Not provided
Open Items / Unknowns:
  - Whether oraclepack currently shells out to a global oracle, vendors a local dependency, or uses npx (not provided).
  - Whether both `.state.json` and `.report.json` exist and their current schema/paths (not provided).
Risks / Dependencies:
  - Depends on CI/runtime environment having Node available if execution uses node workspace or npx.
Acceptance Criteria:
  - A pinned oracle version is stored in a single source-of-truth file (example: `tools/oracle-version.txt`).
  - oraclepack execution uses the pinned version deterministically (no implicit “latest”).
  - `.state.json` and/or `.report.json` emitted by a run includes the oracle version used.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “pins `@steipete/oracle`… runs that version deterministically”
  - “logs/records the exact `oracle` version used into its `.state.json`/`.report.json`”
  - “Vendor via Node workspace… or run `npx -y @steipete/oracle@<pinned>`”
```

```ticket T2
T# Title: Add pass-through forwarding for upstream oracle args/flags to reduce wrapper coupling
Type: enhancement
Target Area: CLI wrapper interface (oraclepack → oracle invocation)
Summary:
  Reduce breakage risk by defaulting to a “pass-through” model where oraclepack forwards custom args/flags to upstream oracle unchanged. Only model/enumerate upstream flags where oraclepack explicitly needs structured UX (e.g., an overrides wizard).
In Scope:
  - Provide a mechanism for users/config to supply freeform upstream args/flags that oraclepack forwards to oracle unchanged.
  - Define/implement the boundary: which flags are wrapper-owned vs upstream pass-through.
  - Keep modeled/enumerated upstream flags limited to where explicitly required (example given: Overrides Wizard).
Out of Scope:
  - Auto-discovery of upstream flags/commands for modeled UX (handled in T3).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - oraclepack can forward arbitrary upstream oracle args/flags without needing wrapper updates for each new upstream option.
  - oraclepack only hard-couples to upstream flag semantics where required for wrapper UX.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Not provided
Evidence:
  - Not provided
Open Items / Unknowns:
  - Current oraclepack surface area that models upstream flags vs forwards arguments (not provided).
  - “Overrides Wizard” location/implementation details in oraclepack (not provided).
Risks / Dependencies:
  - Not provided
Acceptance Criteria:
  - A user-visible way exists to pass through upstream oracle args/flags unchanged.
  - Wrapper-owned flags are not silently forwarded (clear separation in behavior/docs/help output).
  - Existing wrapper features that depend on specific upstream flags continue to function as before.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Reduce ‘parity surface area’… Prefer: ‘pass-through’ custom args/flags… forwards… unchanged”
  - “Only model/enumerate upstream flags where you must (e.g., your Overrides Wizard)”
```

```ticket T3
T# Title: Auto-discover upstream oracle CLI flags/commands used by modeled wrapper UX (avoid hardcoding)
Type: enhancement
Target Area: Wrapper UX that enumerates upstream options (e.g., Overrides Wizard)
Summary:
  Where oraclepack must enumerate upstream oracle flags/commands for structured UX, implement auto-discovery so upstream additions/changes do not require manual hardcoded updates. This specifically targets the stated risk of wrappers breaking when they re-model upstream too strictly.
In Scope:
  - Implement auto-discovery of upstream CLI surface for any wrapper UX that enumerates upstream flags/commands.
  - Ensure the enumerated list updates when the pinned upstream oracle version changes (ties to T1).
Out of Scope:
  - General pass-through forwarding (handled in T2).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - Wrapper UX that enumerates upstream options derives its list from upstream oracle, not a hardcoded list.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Not provided
Evidence:
  - Not provided
Open Items / Unknowns:
  - Which wrapper flows enumerate upstream options (only “Overrides Wizard” is referenced; details not provided).
  - The discovery mechanism source (help output vs another interface) is not specified.
Risks / Dependencies:
  - Depends on T1 for a stable/pinned upstream version target during discovery.
Acceptance Criteria:
  - For any wrapper UX that enumerates upstream flags/commands, the list is derived from upstream oracle rather than hardcoded constants.
  - Updating the pinned upstream oracle version results in the enumerated UX reflecting new/changed upstream options without manual list edits.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “If you do enumerate, auto-discover rather than hardcode”
  - “The main way wrappers break is when they re-model upstream flags/commands too strictly”
```

```ticket T4
T# Title: Add contract test suite for upstream oracle CLI surface changes (snapshot + golden path)
Type: tests
Target Area: CI tests / compatibility checks between oraclepack and upstream oracle
Summary:
  Add tests that fail when upstream oracle changes in ways that could break oraclepack. Include snapshot diffs for help output and a small “golden path” integration fixture that asserts on stable, wrapper-owned artifacts.
In Scope:
  - Snapshot tests capturing `oracle --help` output and `oracle <critical-subcommand> --help` output for a small set of relied-upon commands, then diff snapshots in CI.
  - A “golden path” test that runs a small pack fixture to exercise oraclepack integration points (dry-run path; API path if keys exist; browser path optionally skipped in CI as described).
  - Assertions focus on wrapper-owned artifacts (exit codes, oraclepack report/state schema) rather than upstream human text output.
Out of Scope:
  - Automation that updates the pinned upstream version (handled in T5/T6).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - CI provides early signal when upstream oracle help/CLI surface changes.
  - CI validates at least one end-to-end wrapper “golden path” behavior against the pinned upstream version.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Tests should target the pinned upstream oracle version (deterministic) rather than “latest”.
Evidence:
  - Not provided
Open Items / Unknowns:
  - The “critical subcommands” list oraclepack relies on (examples referenced, but exact dependency set not provided).
  - Where oraclepack stores `.state.json`/`.report.json` and current schema stability guarantees (not provided).
Risks / Dependencies:
  - Depends on T1 for deterministic upstream version selection in tests.
Acceptance Criteria:
  - Snapshot tests exist for `oracle --help` and at least one `oracle <subcommand> --help`, and are executed in CI with diffs detected.
  - A runnable golden-path fixture exists and asserts only on wrapper-owned artifacts (exit code and/or `.state.json`/`.report.json`).
  - CI fails when snapshots change or golden-path behavior breaks against the pinned upstream version.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Snapshot tests… Capture `oracle --help` and `oracle <critical-subcommand> --help`… Diff these snapshots in CI”
  - “Behavioral ‘golden path’ tests… Assert only on stable, wrapper-owned artifacts”
```

```ticket T5
T# Title: Enable automated dependency update PRs for pinned upstream oracle (Dependabot; Renovate as alternative)
Type: chore
Target Area: Repo automation / dependency update configuration
Summary:
  Set up automated PRs to bump the pinned `@steipete/oracle` dependency. The provided example uses Dependabot; Renovate is referenced as an alternative option.
In Scope:
  - Add `.github/dependabot.yml` to enable weekly npm dependency update PRs.
  - Add Dependabot configuration to update GitHub Actions versions as well (as shown in the example).
  - Ensure dependency update PRs run oraclepack’s test suite in CI (ties to T4).
Out of Scope:
  - Implementing Renovate configuration (treated as an alternative approach; not specified beyond reference).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - The repo receives automated PRs that bump npm dependencies, including the pinned upstream oracle dependency, on a schedule.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Dependabot configuration file must live under `.github/dependabot.yml`.
Evidence:
  - Example config snippet references npm + GitHub Actions ecosystems and weekly interval.
Open Items / Unknowns:
  - Whether the project prefers Dependabot vs Renovate (not provided).
Risks / Dependencies:
  - Depends on T4 so update PRs can be validated by contract/golden-path tests.
Acceptance Criteria:
  - `.github/dependabot.yml` exists and is configured for npm and GitHub Actions updates on a weekly schedule.
  - Dependency update PRs trigger CI and run unit + contract tests.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Automate updates: Dependabot/Renovate”
  - “Dependabot version updates… adds PRs based on a `dependabot.yml`”
  - “A) Dependabot config (npm + actions)… schedule interval: weekly”
```

```ticket T6
T# Title: Add scheduled workflow to bump pinned upstream oracle version and open a PR
Type: chore
Target Area: GitHub Actions workflow automation
Summary:
  Add a scheduled GitHub Actions workflow that computes the latest upstream oracle version, updates the pinned version source-of-truth file, runs tests, and opens a PR automatically using `peter-evans/create-pull-request`.
In Scope:
  - Add `.github/workflows/bump-oracle.yml` workflow (as exemplified) that:
    - checks out the repo,
    - sets up Node for `npm view`,
    - computes latest `@steipete/oracle` version,
    - updates the pinned version file (example: `tools/oracle-version.txt`),
    - sets up Go and runs unit + contract tests,
    - creates a PR via `peter-evans/create-pull-request`.
Out of Scope:
  - Canary job running tests against `@latest` without bumping pins (handled in T7).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - On schedule (and manually via workflow dispatch), the repo opens a PR proposing a pinned upstream oracle version bump after tests pass.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Workflow uses `npm view @steipete/oracle version` per example.
  - Workflow runs `go test ./.` per provided snippet.
Evidence:
  - Provided workflow example includes checkout, setup-node, setup-go, test, and create PR steps.
Open Items / Unknowns:
  - Whether the repo wants Dependabot-only vs also maintaining this custom bump workflow (not provided).
Risks / Dependencies:
  - Depends on T4 so the workflow’s “unit + contract tests” exist and can gate PR creation.
Acceptance Criteria:
  - `.github/workflows/bump-oracle.yml` exists and runs on a schedule plus workflow_dispatch.
  - The workflow updates the pinned version file and only opens a PR after tests succeed.
  - The created PR includes the pinned version bump change.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “B) Scheduled workflow that bumps pinned `@steipete/oracle` and opens a PR”
  - “uses… `peter-evans/create-pull-request`… Run unit + contract tests… Create PR”
```

```ticket T7
T# Title: Add scheduled canary CI job that tests oraclepack against upstream oracle “latest”
Type: tests
Target Area: CI early-warning compatibility checks
Summary:
  Add a scheduled “canary” workflow that runs oraclepack’s integration/contract suite against `@steipete/oracle@latest` without changing the pinned version. This provides early warning of upstream changes that may break the wrapper.
In Scope:
  - Add a scheduled CI workflow/job that installs/runs against upstream oracle “latest” and executes oraclepack’s integration/contract test suite.
  - Ensure the canary job is clearly distinct from pinned-version gating tests.
Out of Scope:
  - Automatically updating the pinned oracle version (handled in T6/T5).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - CI periodically reports whether upstream oracle “latest” would break oraclepack, even before adopting the version.
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Canary runs oraclepack’s existing contract/integration suite (as described).
Evidence:
  - Not provided
Open Items / Unknowns:
  - How to run oraclepack tests against “latest” while keeping normal CI pinned (mechanism not specified).
Risks / Dependencies:
  - Depends on T4 because the canary job is defined as running the contract/integration suite.
Acceptance Criteria:
  - A scheduled workflow/job exists that runs oraclepack’s contract/integration tests against `@steipete/oracle@latest`.
  - Canary results are visible in CI and fail when breakage is detected.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Lane B (early warning): a scheduled CI job… against `@steipete/oracle@latest`… tells you ‘breakage is coming’”
```

```ticket T8
T# Title: Document and enforce compatibility matrix + release gating for oraclepack vs upstream oracle
Type: docs
Target Area: Release discipline / documentation / CI release gates
Summary:
  Publish a compatibility statement for which upstream oracle versions oraclepack supports (or is tested against). Gate oraclepack releases based on contract test results on the pinned version and review/health of the canary signal.
In Scope:
  - Add a compatibility statement (example phrasing provided: “oraclepack vX.Y supports oracle vA.B.C–vA.B.*” or “tested against”).
  - Define release gating expectations tied to:
    - contract tests passing on pinned oracle version, and
    - canary job not failing (or at least reviewed).
Out of Scope:
  - Implementing the tests/workflows themselves (handled in T4/T7).
Current Behavior (Actual):
  - Not provided
Expected Behavior:
  - Users can determine which upstream oracle versions are supported/tested.
  - Releases are gated by defined CI signals (pinned contract suite + canary).
Reproduction Steps:
  - Not provided
Requirements / Constraints:
  - Not provided
Evidence:
  - Not provided
Open Items / Unknowns:
  - Where releases are performed/published and what existing release process exists (not provided).
Risks / Dependencies:
  - Depends on T4 and T7 because gating references those CI signals.
Acceptance Criteria:
  - Repo documentation includes a compatibility statement for oraclepack vs upstream oracle versions.
  - Release gating criteria are documented and reference the pinned contract suite and canary status/review.
Priority & Severity (if inferable from input text):
  - Priority: Not provided
  - Severity: Not provided
Source:
  - “Release discipline: compatibility matrix + gating”
  - “Gate oraclepack releases on: Contract tests… Canary job not failing (or at least reviewed)”
```

Non-actionable / Info-only:

* “If oraclepack were literally a fork of oracle… use a fork-sync/upstream-sync workflow/action… In your case it’s a wrapper… dependency+tests is usually the better fit.”
* “If you ever decide to vendor upstream code into the wrapper… Use subtree/submodule… Submodules… don’t auto-track branches… Subtree vendors code directly…”

[1]: https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuring-dependabot-version-updates?utm_source=chatgpt.com "Configuring Dependabot version updates"
[2]: https://docs.renovatebot.com/modules/manager/npm/?utm_source=chatgpt.com "Automated Dependency Updates for npm - Renovate Docs"
[3]: https://github.com/peter-evans/create-pull-request?utm_source=chatgpt.com "peter-evans/create-pull-request"
[4]: https://github.com/marketplace/actions/upstream-sync?utm_source=chatgpt.com "Upstream Sync · Actions · GitHub Marketplace"
[5]: https://www.atlassian.com/git/tutorials/git-subtree?utm_source=chatgpt.com "Git Subtree: Alternative to Git Submodule"
