Parent Ticket:

* Title: Stop oraclepack from writing run state/config JSON files into the project working directory
* Summary: oraclepack currently writes per-pack `*.state.json`, `*.report.json`, and `*.chatgpt-urls.json` files into the repo/working directory. The requested change is to store these as config/state/cache outside the repo root (prefer XDG base dirs) and/or under a dedicated project-local `.oraclepack/` directory to avoid clutter.
* Source:

  * Link/ID: Not provided
  * Original ticket excerpt (≤25 words) capturing the overall theme: “Move state/report outputs out of CWD by default… Stop producing per-pack `*.chatgpt-urls.json` by default.”
* Global Constraints:

  * Treat outputs as **config/state/cache** and store outside repo root using XDG base dirs (per ticket text).
  * Use Go `os.UserConfigDir()` / `os.UserCacheDir()` for cross-platform defaults (per ticket text).
  * No `UserStateDir()` in Go stdlib; implement `$XDG_STATE_HOME` fallback (per ticket text).
* Global Environment:

  * Unknown
* Global Evidence:

  * Current filenames mentioned: `<packBase>.state.json`, `<packBase>.report.json`, `<sameBase>.chatgpt-urls.json`.
  * XDG Base Directory spec reference (background).
  * Go `os.UserConfigDir` / `os.UserCacheDir` reference (background).

Split Plan:

* Coverage Map:

  * Original item: “`oraclepack run` … derives filenames from the pack basename and writes them to the current working directory: `statePath := <packBase>.state.json`, `reportPath := <packBase>.report.json`”

    * Assigned Ticket ID: T2
  * Original item: “The TUI ‘ChatGPT URL picker’ then creates `<sameBase>.chatgpt-urls.json` next to the state file (or next to the pack file if statePath is empty).”

    * Assigned Ticket ID: T3
  * Original item: “It also defaults edits to **project scope**, so it will keep generating project-scoped stores unless the user explicitly switches to global.”

    * Assigned Ticket ID: T3
  * Original item: “Treat these as **config/state/cache** and store them outside the repo root using standard base dirs: … `$XDG_CONFIG_HOME` … `$XDG_STATE_HOME` … `$XDG_CACHE_HOME`…”

    * Assigned Ticket ID: T1
  * Original item: “In Go, you should use `os.UserConfigDir()` / `os.UserCacheDir()`… (There’s no `UserStateDir()`… implement XDG_STATE_HOME fallback…)”

    * Assigned Ticket ID: T1
  * Original item: “Move state/report outputs out of CWD by default… Update `internal/cli/run.go`… Make the directory overridable with a flag/env (e.g., `--state-dir` / `ORACLEPACK_STATE_DIR`).”

    * Assigned Ticket ID: T2
  * Original item: “Stop producing per-pack `*.chatgpt-urls.json` by default… Best UX default: change … default save scope to **global**…”

    * Assigned Ticket ID: T3
  * Original item: “Keep ‘project scope’ as an opt-in mode, but write it to a single per-project location (e.g., `<repo>/.oraclepack/chatgpt-urls.json`), not `<packName>.chatgpt-urls.json`.”

    * Assigned Ticket ID: T3
  * Original item: “Acceptable alternative (project-local…): `<repo>/.oraclepack/state/*.state.json` … `<repo>/.oraclepack/chatgpt-urls.json` … add `.oraclepack/` to `.gitignore`.”

    * Assigned Ticket ID: T4
  * Original item: “Immediate workaround (no code changes): Add these to `.gitignore`: `*.state.json`, `*.report.json`, `*.chatgpt-urls.json`.”

    * Assigned Ticket ID: T4
* Dependencies:

  * T2 depends on T1 because T2 needs an agreed/default “oraclepack state dir” location strategy (XDG-based) to write into.
  * T3 depends on T1 because T3 needs a global config location strategy (XDG-based) for URL persistence.
* Split Tickets:

```ticket T1
T# Title: Define XDG-based directory strategy for oraclepack config/state/cache
Type: chore
Target Area: Config/state path resolution (shared utility / helpers)
Summary:
- Define the standard locations where oraclepack stores user config, run state, and cache so outputs stop polluting the repo root.
- The ticket requires using XDG base dirs and Go’s cross-platform helpers where applicable, with an explicit fallback for state.
In Scope:
- Adopt XDG directory categories as the guiding model:
  - Config: `$XDG_CONFIG_HOME` (default `~/.config`)
  - State: `$XDG_STATE_HOME` (default `~/.local/state`)
  - Cache: `$XDG_CACHE_HOME` (default `~/.cache`)
- Use Go `os.UserConfigDir()` / `os.UserCacheDir()` for cross-platform defaults (per ticket text).
- Implement a state-dir resolver that honors `$XDG_STATE_HOME` and falls back when not set (since Go stdlib has no `UserStateDir()`).
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- oraclepack has a single, consistent mechanism to determine:
  - “oraclepack config dir” (for user prefs like URL lists)
  - “oraclepack state dir” (for resume/run state)
  - “oraclepack cache dir” (for non-essential cached data)
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Treat outputs as config/state/cache; store outside repo root using standard base dirs (per ticket text).
- Use `os.UserConfigDir()` / `os.UserCacheDir()` where applicable (per ticket text).
- Implement `$XDG_STATE_HOME` fallback logic (per ticket text).
Evidence:
- “Treat these as config/state/cache and store them outside the repo root using standard base dirs…” (parent ticket)
- “In Go, you should use os.UserConfigDir() / os.UserCacheDir()… There’s no UserStateDir()…” (parent ticket)
Open Items / Unknowns:
- Exact package/file locations for where to place the shared directory-resolution logic: Unknown
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- A single directory-resolution mechanism exists for config/state/cache categories as described in scope.
- The state-dir resolution honors `$XDG_STATE_HOME` when set and has a documented fallback when unset.
- The config/cache resolution uses Go’s `os.UserConfigDir()` / `os.UserCacheDir()` (or equivalent wrapper) per ticket text.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Treat these as config/state/cache and store them outside the repo root using standard base dirs…”
- “In Go, you should use os.UserConfigDir() / os.UserCacheDir()… There’s no UserStateDir()…”
```

```ticket T2
T# Title: Move run-generated state/report JSON outputs out of CWD and add state-dir override
Type: enhancement
Target Area: Run command output paths (`internal/cli/run.go`)
Summary:
- oraclepack currently writes `<packBase>.state.json` and `<packBase>.report.json` into the current working directory.
- Update the run pathing so these files are written under a dedicated “oraclepack state dir” by default, with an override via flag/env.
In Scope:
- Change default output location for:
  - `<packBase>.state.json`
  - `<packBase>.report.json`
  from current working directory to a dedicated “oraclepack state dir”.
- Update `internal/cli/run.go` to compute state/report paths under that state dir (per ticket text).
- Add override via flag and env:
  - `--state-dir`
  - `ORACLEPACK_STATE_DIR`
Out of Scope:
- Not provided
Current Behavior (Actual):
- `<packBase>.state.json` and `<packBase>.report.json` are written to the current working directory.
Expected Behavior:
- By default, running oraclepack does not create `*.state.json` / `*.report.json` in the repo root / CWD.
- By default, state/report files are written under the dedicated oraclepack state dir.
- Setting `--state-dir` or `ORACLEPACK_STATE_DIR` writes state/report files under the specified directory.
Reproduction Steps:
1) Run `oraclepack run` from a repo root (or any working directory).
2) Observe creation of `<packBase>.state.json` and `<packBase>.report.json` in the working directory.
Requirements / Constraints:
- Must be overridable by `--state-dir` / `ORACLEPACK_STATE_DIR` (per ticket text).
- Should use the state-dir strategy defined in T1 for the default state dir.
Evidence:
- “`oraclepack run` … writes them to the current working directory: `statePath := <packBase>.state.json`, `reportPath := <packBase>.report.json`”
- “Update `internal/cli/run.go` … Make the directory overridable with a flag/env (e.g., `--state-dir` / `ORACLEPACK_STATE_DIR`).”
Open Items / Unknowns:
- Whether state/report filenames must remain exactly `<packBase>.state.json` / `<packBase>.report.json` or can change: Not provided
Risks / Dependencies:
- Depends on T1 for default state-dir resolution strategy.
Acceptance Criteria:
- Running oraclepack with no overrides does not create `*.state.json` or `*.report.json` in the current working directory.
- With no overrides, state/report files are written under the resolved oraclepack state dir.
- With `--state-dir=<dir>`, state/report files are written under `<dir>`.
- With `ORACLEPACK_STATE_DIR=<dir>`, state/report files are written under `<dir>`.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “`oraclepack run` … writes … `<packBase>.state.json` … `<packBase>.report.json` … to the current working directory”
- “Move state/report outputs out of CWD by default… Update internal/cli/run.go… `--state-dir` / `ORACLEPACK_STATE_DIR`”
```

```ticket T3
T# Title: Stop generating per-pack `*.chatgpt-urls.json`; default URL picker persistence to global store
Type: enhancement
Target Area: TUI “ChatGPT URL picker” persistence
Summary:
- The TUI URL picker currently creates `<sameBase>.chatgpt-urls.json` near the pack/state file and defaults edits to project scope.
- Change it so the default save scope is global (one file), while keeping project scope as an opt-in that writes to a single stable per-project path.
In Scope:
- Remove/avoid creating `<sameBase>.chatgpt-urls.json` “next to the state file (or next to the pack file…)” (per ticket text).
- Change the URL picker default save scope to **global** (per ticket text).
- Keep “project scope” as opt-in, but store at a single stable path:
  - `<repo>/.oraclepack/chatgpt-urls.json`
  rather than `<packName>.chatgpt-urls.json` (per ticket text).
- Persist the global URL store to a single global location:
  - Per ticket text, an existing global store path is referenced: `~/.oraclepack/chatgpt-urls.json`.
Out of Scope:
- Not provided
Current Behavior (Actual):
- URL picker creates `<sameBase>.chatgpt-urls.json` next to the state file (or pack file).
- URL picker defaults edits to project scope.
Expected Behavior:
- Using the URL picker does not create `<packBase>.chatgpt-urls.json` files.
- Default persistence is global (one stable file).
- Project scope, if selected, writes only to `<repo>/.oraclepack/chatgpt-urls.json`.
Reproduction Steps:
1) Use the TUI “ChatGPT URL picker” during a run.
2) Observe `<sameBase>.chatgpt-urls.json` being created near pack/state file.
Requirements / Constraints:
- Default save scope should be global (per ticket text).
- Project scope must be opt-in and must not create per-pack URL json files (per ticket text).
- Global persistence location:
  - Conflicting guidance exists: ticket recommends XDG config dir generally, but also references existing `~/.oraclepack/chatgpt-urls.json` path.
Evidence:
- “The TUI ‘ChatGPT URL picker’ then creates `<sameBase>.chatgpt-urls.json`…”
- “It also defaults edits to project scope…”
- “Stop producing per-pack `*.chatgpt-urls.json` by default… change … default save scope to global…”
- “Keep ‘project scope’ as an opt-in… write it to `<repo>/.oraclepack/chatgpt-urls.json`… not `<packName>.chatgpt-urls.json`.”
Open Items / Unknowns:
- Whether to keep `~/.oraclepack/chatgpt-urls.json` as the global path or migrate to `$XDG_CONFIG_HOME/...` (both appear in the parent ticket guidance).
- Exact file/path of the “URL picker” implementation: Not provided
Risks / Dependencies:
- Depends on T1 if migrating global storage to XDG config dir.
Acceptance Criteria:
- After using the URL picker, no `<packBase>.chatgpt-urls.json` is created near the pack/state/CWD.
- Default behavior persists URLs to exactly one global store (stable path; not per-pack).
- When “project scope” is selected, URLs persist to `<repo>/.oraclepack/chatgpt-urls.json` (single per-project file).
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “The TUI ‘ChatGPT URL picker’ then creates `<sameBase>.chatgpt-urls.json`…”
- “Best UX default: change … default save scope to global…”
- “Keep ‘project scope’ … `<repo>/.oraclepack/chatgpt-urls.json`, not `<packName>.chatgpt-urls.json`.”
```

```ticket T4
T# Title: Add project-local `.oraclepack/` layout guidance and `.gitignore` patterns to prevent repo pollution
Type: docs
Target Area: Repo hygiene (docs / templates / ignore rules)
Summary:
- The parent ticket proposes an acceptable project-local layout under `<repo>/.oraclepack/` and an immediate workaround via `.gitignore`.
- Capture these as documented guidance (and/or provide default ignore patterns) so repos remain clean even before code changes land.
In Scope:
- Document (or provide recommended structure for) project-local layout:
  - `<repo>/.oraclepack/state/*.state.json`
  - `<repo>/.oraclepack/state/*.report.json`
  - `<repo>/.oraclepack/chatgpt-urls.json`
- Add guidance to add `.oraclepack/` to `.gitignore` when adopting that structure.
- Add the immediate workaround ignore patterns:
  - `*.state.json`
  - `*.report.json`
  - `*.chatgpt-urls.json`
Out of Scope:
- Not provided
Current Behavior (Actual):
- Not provided
Expected Behavior:
- Repos can adopt a single project-local `.oraclepack/` directory and ignore it.
- Repos can immediately ignore current output filenames to avoid noise.
Reproduction Steps:
- Not provided
Requirements / Constraints:
- Must preserve the exact patterns and structure described in the parent ticket text.
Evidence:
- “Acceptable alternative (project-local…): `<repo>/.oraclepack/state/*.state.json` … add `.oraclepack/` to `.gitignore`.”
- “Immediate workaround (no code changes): Add these to `.gitignore`: `*.state.json`, `*.report.json`, `*.chatgpt-urls.json`.”
Open Items / Unknowns:
- Where this guidance should live (README, docs page, template): Not provided
Risks / Dependencies:
- Not provided
Acceptance Criteria:
- Documentation includes the described `.oraclepack/` directory layout and explicitly recommends ignoring `.oraclepack/` when using it.
- Documentation includes the three immediate-workaround `.gitignore` patterns exactly as specified.
Priority & Severity (if inferable from input text):
- Priority: Not provided
- Severity: Not provided
Source:
- “Acceptable alternative (project-local…): `<repo>/.oraclepack/state/*.state.json`… add `.oraclepack/` to `.gitignore`.”
- “Immediate workaround… Add these to `.gitignore`: `*.state.json`…”
```
