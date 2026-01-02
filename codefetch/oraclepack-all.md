<filetree>
Project Structure:
├── .config
│   ├── commands
│   │   └── oracle-pack_v2.toml
│   ├── completion
│   │   └── oraclepack.completion.sh
│   ├── scripts
│   │   ├── build_install_oraclepack.md
│   │   ├── build_install_oraclepack.sh
│   │   ├── install-global.ps1
│   │   └── install-global.sh
│   └── skill
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
│       └── update-codefetch-outputs.yml
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
│   │   ├── inject.go
│   │   ├── inject_test.go
│   │   ├── runner.go
│   │   ├── runner_test.go
│   │   └── stream.go
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
│       ├── filter_test.go
│       ├── tui.go
│       └── tui_test.go
├── working-examples
│   └── strategist-questions-oracle-pack-2025-12-30.md
├── .browser-echo-mcp.json
├── .goreleaser.yaml
├── TUI Runtime Overrides Implementation.md
├── go.mod
└── package.json

</filetree>

<source_code>
.browser-echo-mcp.json
```
{"url":"http://127.0.0.1:37599","route":"/__client-logs","timestamp":1767299128472,"pid":34018}
```

.goreleaser.yaml
```
# .goreleaser.yaml
project_name: oraclepack

before:
  hooks:
    - go mod tidy

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    main: ./cmd/oraclepack
    ldflags:
      - -s -w -X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}}

archives:
  - name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"
    format_overrides:
      - goos: windows
        format: zip

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{ .Tag }}-next"

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
```

TUI Runtime Overrides Implementation.md
```
Title:

* Add TUI “Runtime Overrides” picker to inject Oracle flags + project URLs per selected oraclepack steps (with Mode 2 validation)

Summary:

* Extend the `oraclepack` Go TUI to let users (1) choose/enter ChatGPT project URLs (endpoints) and (2) add/remove Oracle CLI flags at runtime, then (3) select which oraclepack steps should receive those overrides, and (4) confirm before returning to the main run screen. Validation should use the “mode 2” approach (per user preference) to ensure selected flags/overrides are accepted by the upstream `oracle` CLI before running.

    Branch · Extending OraclePack F…

Background / Context:

* `oraclepack` is a wrapper around upstream `oracle` ([https://github.com/steipete/oracle](https://github.com/steipete/oracle)) and currently injects `Runner.OracleFlags` into steps by rewriting step shell scripts via `InjectFlags` in `internal/exec/Runner.RunStep` (per assistant).

    Branch · Extending OraclePack F…

* User goal: allow runtime customization for different use cases, specifically cherry-picking flags and selecting/entering different project endpoints (project URLs) without editing config files per run.

    Branch · Extending OraclePack F…

* TUI flow requested by user: flags picker → steps-to-modify picker → confirm → return to initial run screen; plus a separate menu for project URL input/injection; a broader config menu is explicitly deferred to the future.

    Branch · Extending OraclePack F…

Current Behavior (Actual):

* Runtime flag/endpoint selection via TUI pickers is not available (per user request implying missing functionality).

    Branch · Extending OraclePack F…

* Current injection mechanism is line-based regex append-at-EOL; assistant notes it can break common multi-line `oracle \` command formatting (risk/limitation in current approach).

    Branch · Extending OraclePack F…

Expected Behavior:

* From the main run screen, a user can open a TUI flow to:

  * Add/remove Oracle flags (multi-select).

  * Choose which oraclepack steps will be affected (multi-select).

  * Enter/select project URL(s) to be injected (separate menu in TUI).

  * Confirm changes and return to the initial run screen to proceed with normal step selection and run.

        Branch · Extending OraclePack F…

* “Mode 2” validation runs before returning to the run screen, rejecting invalid flag combinations/overrides and showing an actionable error.

    Branch · Extending OraclePack F…

Requirements:

* TUI picker flow:

  * Multi-select UI to add/remove flags for a run.

        Branch · Extending OraclePack F…

  * Multi-select UI to choose which oraclepack steps are modified by the selected flag changes.

        Branch · Extending OraclePack F…

  * Confirmation step summarizing selected flags and affected steps; then return to the initial run screen (no immediate run required).

        Branch · Extending OraclePack F…

* Project URL injection:

  * Provide a separate TUI menu for users to input project URL(s) that oraclepacks get sent to (endpoint customization).

        Branch · Extending OraclePack F…

* Validation:

  * Use “mode 2” validation (user preference) to verify the injected flags/overrides are valid for `oracle` before proceeding.

        Branch · Extending OraclePack F…

  * Assistant proposed “dry-run parse per affected step” as the practical mode-2 mechanism, executing affected `oracle ...` invocations in `--dry-run` and failing fast with step + invocation + error output.

        Branch · Extending OraclePack F…

* Constraints:

  * Do not build the “full customizable config menu” now; stick to the runtime pickers and project URL input described above.

        Branch · Extending OraclePack F…

Out of Scope:

* Full configurable config/settings menu beyond project URL injection and runtime flag/step selection (explicitly deferred).

    Branch · Extending OraclePack F…

Reproduction Steps:

* Not provided.

Environment:

* OS: Unknown

* `oraclepack` version/commit: Unknown

* TUI framework: Bubble Tea (implied by assistant discussion; exact versions unknown).

    Branch · Extending OraclePack F…

Evidence:

* Conversation export: /mnt/data/Branch · Extending OraclePack Functionality.md

    Branch · Extending OraclePack F…

* Notable referenced components/files (content not included in the export):

  * `oraclepack-all.md` (referenced as code base context)

  * `config.json` (referenced for project URL / remote fields)

        Branch · Extending OraclePack F…

* Key quoted intent (per user, from the export): “picker in the TUI… cherry-pick additional flags… cherry-pick which oraclepack steps… confirm changes… separate menu… project urls… In the future… full customizable config menu… but for now…”

    Branch · Extending OraclePack F…

Decisions / Agreements:

* Use a picker-based TUI flow (not arbitrary runtime command injection UI as the primary UX) (per user).

    Branch · Extending OraclePack F…

* Prefer “mode 2” validation (per user).

    Branch · Extending OraclePack F…

* Defer full config menu; implement only the specified pickers + project URL input now (per user).

    Branch · Extending OraclePack F…

Open Items / Unknowns:

* Definition of “mode 2” validation is partially ambiguous:

  * Assistant earlier described mode 2 as help-based preflight (`oracle --help` / `oracle <subcommand> --help`), and later described mode 2 as `--dry-run` validation per invocation; user only confirmed “mode 2” generally.

        Branch · Extending OraclePack F…

* Where and how project URLs should be stored (ephemeral per run vs persisted registry) is not specified by the user.

    Branch · Extending OraclePack F…

* Exact oraclepack step model (IDs/titles, where steps are defined) and where in the TUI state machine to add the new screens is not provided in the export (assistant referenced `internal/tui/tui.go`, but code is not attached).

    Branch · Extending OraclePack F…

Risks / Dependencies:

* Dependency on upstream `oracle` CLI behavior and flags (validation and injected overrides must match supported CLI semantics).

    Branch · Extending OraclePack F…

* Injection mechanism risk: current line-based `InjectFlags` approach may break multi-line commands with `\` continuations; may require hardening or more robust parsing to safely inject flags (per assistant).

    Branch · Extending OraclePack F…

Acceptance Criteria:

* A new TUI menu entry (from the main run screen) launches a runtime overrides flow that:

  * Lets the user add/remove flags (multi-select) and proceeds to a step-selection picker (multi-select).

        Branch · Extending OraclePack F…

  * Includes a project URL input/selection menu to control where oraclepacks are sent.

        Branch · Extending OraclePack F…

  * Shows a confirmation screen summarizing: selected flags (added/removed), affected steps, and selected/entered project URL(s).

        Branch · Extending OraclePack F…

  * On confirmation, performs “mode 2” validation; if validation fails, the UI displays the failure and does not proceed back to the main run screen as “ready”.

        Branch · Extending OraclePack F…

  * If validation succeeds, returns to the main run screen with overrides staged for the run, and the user can run steps as normal.

        Branch · Extending OraclePack F…

Priority & Severity (if inferable from text):

* Not provided.

Labels (optional):

* enhancement

* tui

* flags

* validation

* config

* endpoints

* oraclepack

---
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

package.json
```
{
  "devDependencies": {
    "codefetch": "^2.2.0"
  },
  "scripts": {
    "code:tui": "codefetch -t 5 --include-dir cmd,internal -o oraclepack.md --max-tokens 50000 --token-limiter truncated",
    "code:all": "codefetch -t 5 -o oraclepack-all.md",
    "code": "pnpm code:tui && pnpm code:all"
  }
}
```

working-examples/strategist-questions-oracle-pack-2025-12-30.md
```
# strategist questions — oracle pack

---

## run (exactly 20)

```bash
# shellcheck shell=bash
set -euo pipefail

out_dir="artifacts/oracle/strategist-questions/2025-12-30"
mkdir -p "$out_dir"

# 01) ROI=2.80 | Immediate | invariants | src/registry_config/models.py:validate_unique_ids
oracle \
  --files-report \
  --write-output "$out_dir/01-immediate-invariants-unique-ids.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, invariants, ref: src/registry_config/models.py:validate_unique_ids): Does validate_unique_ids correctly detect duplicate source IDs during manifest load and make the error actionable for maintainers? Answer format: Reference: src/registry_config/models.py:validate_unique_ids; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/registry_config/models.py"

# 02) ROI=2.40 | Immediate | invariants | src/registry_config/models.py:validate_slug
oracle \
  --files-report \
  --write-output "$out_dir/02-immediate-invariants-slug-validation.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, invariants, ref: src/registry_config/models.py:validate_slug): Does validate_slug and SLUG_REGEX accept all current source IDs and doc folder names, or will it reject any real entries? Answer format: Reference: src/registry_config/models.py:validate_slug; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/registry_config/models.py" \
  -f "sources.json"

# 03) ROI=2.00 | Immediate | contracts/interfaces | src/registry_config/loader.py:load_manifest
oracle \
  --files-report \
  --write-output "$out_dir/03-immediate-contracts-load-manifest.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, contracts/interfaces, ref: src/registry_config/loader.py:load_manifest): When sources.json is missing or malformed, does load_manifest behave as intended (empty Manifest vs error), and do callers handle that outcome safely? Answer format: Reference: src/registry_config/loader.py:load_manifest; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/registry_config/loader.py" \
  -f "sources.json"

# 04) ROI=1.87 | Immediate | contracts/interfaces | src/registry_config/models.py:Source
oracle \
  --files-report \
  --write-output "$out_dir/04-immediate-contracts-source-schema.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, contracts/interfaces, ref: src/registry_config/models.py:Source): Does the Source model capture all fields actually used by the refresh workflow (enabled/profile/last_refreshed/last_model_used), and are defaults aligned with scripts/refresh.py expectations? Answer format: Reference: src/registry_config/models.py:Source; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/registry_config/models.py" \
  -f "scripts/refresh.py"

# 05) ROI=1.80 | Immediate | observability | src/generator_runner/runner.py:logger
oracle \
  --files-report \
  --write-output "$out_dir/05-immediate-observability-run-logging.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, observability, ref: src/generator_runner/runner.py:logger): Are current log statements sufficient to trace per-source execution, including command, duration, and errors, or do we need additional structured logging? Answer format: Reference: src/generator_runner/runner.py:logger; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/generator_runner/runner.py"

# 06) ROI=1.58 | Immediate | failure modes | src/generator_runner/runner.py:run_source
oracle \
  --files-report \
  --write-output "$out_dir/06-immediate-failure-run-source.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, failure modes, ref: src/generator_runner/runner.py:run_source): Does run_source surface subprocess failures and timeouts with enough context (stderr/stdout/returncode) to diagnose generator issues quickly? Answer format: Reference: src/generator_runner/runner.py:run_source; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/generator_runner/runner.py"

# 07) ROI=1.40 | Immediate | feature flags | src/generator_runner/runner.py:env["ENABLE_CTX"]
oracle \
  --files-report \
  --write-output "$out_dir/07-immediate-feature-flags-enable-ctx.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, feature flags, ref: src/generator_runner/runner.py:env[ENABLE_CTX]): ENABLE_CTX is hard-coded to 1; should this be configurable via CLI or manifest profile, and what is the safest default? Answer format: Reference: src/generator_runner/runner.py:env[ENABLE_CTX]; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/generator_runner/runner.py" \
  -f "scripts/refresh.py"

# 08) ROI=1.40 | Immediate | contracts/interfaces | src/reporting/report.py:RunReport
oracle \
  --files-report \
  --write-output "$out_dir/08-immediate-contracts-run-report.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, contracts/interfaces, ref: src/reporting/report.py:RunReport): Does RunReport capture all fields needed for auditing a refresh run (per-source status, model_used, artifacts, duration), or are key fields missing? Answer format: Reference: src/reporting/report.py:RunReport; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/reporting/report.py"

# 09) ROI=1.40 | Immediate | invariants | src/git_sync/status.py:check_stale_artifacts
oracle \
  --files-report \
  --write-output "$out_dir/09-immediate-invariants-stale-check.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, invariants, ref: src/git_sync/status.py:check_stale_artifacts): Does check_stale_artifacts reliably detect stale docs updates when sources.json changes, or can it miss or misflag common workflows? Answer format: Reference: src/git_sync/status.py:check_stale_artifacts; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/git_sync/status.py"

# 10) ROI=1.40 | Immediate | contracts/interfaces | src/artifact_ingest/ingest.py:ingest_artifacts
oracle \
  --files-report \
  --write-output "$out_dir/10-immediate-contracts-ingest-artifacts.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, contracts/interfaces, ref: src/artifact_ingest/ingest.py:ingest_artifacts): Does ingest_artifacts handle normalization and copying without overwriting unintended files, and should it enforce deduping or cleanup semantics? Answer format: Reference: src/artifact_ingest/ingest.py:ingest_artifacts; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/artifact_ingest/ingest.py"

# 11) ROI=1.20 | Immediate | contracts/interfaces | src/artifact_ingest/index.py:generate_registry_index
oracle \
  --files-report \
  --write-output "$out_dir/11-immediate-contracts-registry-index.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, contracts/interfaces, ref: src/artifact_ingest/index.py:generate_registry_index): Does generate_registry_index produce stable ordering and include all required artifact patterns/metadata for downstream consumers? Answer format: Reference: src/artifact_ingest/index.py:generate_registry_index; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/artifact_ingest/index.py"

# 12) ROI=1.20 | Immediate | background jobs | scripts/refresh.py:main
oracle \
  --files-report \
  --write-output "$out_dir/12-immediate-jobs-refresh-main.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Immediate, background jobs, ref: scripts/refresh.py:main): Is refresh.py main idempotent and safe to re-run on the same source set, and are outputs deterministic across runs? Answer format: Reference: scripts/refresh.py:main; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "scripts/refresh.py"

# 13) ROI=0.84 | Strategic | observability | src/reporting/report.py:to_json
oracle \
  --files-report \
  --write-output "$out_dir/13-strategic-observability-report-json.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, observability, ref: src/reporting/report.py:to_json): Is the RunReport JSON format stable enough for long-term observability and metrics ingestion, or do we need schema versioning and retention strategy? Answer format: Reference: src/reporting/report.py:to_json; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/reporting/report.py"

# 14) ROI=0.75 | Strategic | UX flows | scripts/refresh.py:argparse.ArgumentParser
oracle \
  --files-report \
  --write-output "$out_dir/14-strategic-ux-cli-args.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, UX flows, ref: scripts/refresh.py:argparse.ArgumentParser): Is the CLI UX for refresh.py adequate (help text, error clarity, flags), and would adding a dry-run or list-sources command materially reduce operator error? Answer format: Reference: scripts/refresh.py:argparse.ArgumentParser; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "scripts/refresh.py"

# 15) ROI=0.72 | Strategic | contracts/interfaces | scripts/populate_manifest.py:scan_and_populate
oracle \
  --files-report \
  --write-output "$out_dir/15-strategic-contracts-populate-manifest.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, contracts/interfaces, ref: scripts/populate_manifest.py:scan_and_populate): Does scan_and_populate safely rename doc directories and update sources.json without collisions or data loss, and should it support preview/dry-run? Answer format: Reference: scripts/populate_manifest.py:scan_and_populate; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "scripts/populate_manifest.py"

# 16) ROI=0.70 | Strategic | feature flags | src/registry_config/models.py:GeneratorProfile
oracle \
  --files-report \
  --write-output "$out_dir/16-strategic-feature-profiles.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, feature flags, ref: src/registry_config/models.py:GeneratorProfile): Are GeneratorProfile fields sufficient for future provider support (headers, timeout, model), or do we need additional per-provider flags/config? Answer format: Reference: src/registry_config/models.py:GeneratorProfile; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/registry_config/models.py" \
  -f "sources.json"

# 17) ROI=0.57 | Strategic | background jobs | src/generator_runner/runner.py:GeneratorRunner
oracle \
  --files-report \
  --write-output "$out_dir/17-strategic-jobs-runner-scale.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, background jobs, ref: src/generator_runner/runner.py:GeneratorRunner): Should GeneratorRunner support parallel execution or rate limiting for large source sets, and where is the safest insertion point? Answer format: Reference: src/generator_runner/runner.py:GeneratorRunner; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "src/generator_runner/runner.py"

# 18) ROI=0.53 | Strategic | caching/state | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/18-strategic-cache-unknown.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, caching/state, ref: Unknown; missing patterns: **/*cache*.*, **/redis/**, **/*session*.*, **/*store*.*): Is a caching/state layer needed for artifact generation or downloads, and where should it live if introduced? If reference is Unknown, confirm absence in attached files and suggest the most likely location in this repo. Answer format: Reference: Unknown; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "README.md" \
  -f "sources.json"

# 19) ROI=0.36 | Strategic | permissions | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/19-strategic-permissions-unknown.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, permissions, ref: Unknown; missing patterns: **/auth/**, **/*rbac*.*, **/*permission*.*, **/*policy*.*): Does registry ingestion require access control or URL allowlists, and where would that enforcement live? If reference is Unknown, confirm absence in attached files and suggest the most likely location in this repo. Answer format: Reference: Unknown; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "README.md"

# 20) ROI=0.35 | Strategic | migrations | Unknown
oracle \
  --files-report \
  --write-output "$out_dir/20-strategic-migrations-unknown.md" \
  -p "Repo: llms-txt registry tooling in Python; core code in src/ and scripts/, artifacts in docs/. Constraints: Unknown. Non-goals: Unknown. Team size: Unknown. Deadline: Unknown. Question (Strategic, migrations, ref: Unknown; missing patterns: **/migrations/**, **/schema.* , **/prisma/**): How should sources.json schema evolution be managed, and do we need explicit migrations or versioned schema docs? If reference is Unknown, confirm absence in attached files and suggest the most likely location in this repo. Answer format: Reference: Unknown; Direct answer; Evidence (cite attached paths/symbols); Smallest experiment today; Risks/assumptions (max 3)." \
  -f "README.md" \
  -f "sources.json"
```
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

.github/workflows/update-codefetch-outputs.yml
```
# path: .github/workflows/update-codefetch-outputs.yml
name: Update codefetch outputs

on:
  push:
    # Run on every push, but ignore pushes that only update the generated docs
    paths-ignore:
      - "codefetch/oraclepack.md"
      - "codefetch/oraclepack-all.md"
  workflow_dispatch: {}

permissions:
  contents: write

concurrency:
  group: update-codefetch-outputs-${{ github.ref }}
  cancel-in-progress: true

jobs:
  generate-and-commit:
    runs-on: ubuntu-latest

    # Avoid running on the bot's own commit if you ever change paths-ignore later
    if: github.actor != 'github-actions[bot]'

    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Setup Node
        uses: actions/setup-node@v4
        with:
          node-version: "22"
          cache: "pnpm"

      - name: Install dependencies
        shell: bash
        run: |
          set -euo pipefail
          if [[ -f pnpm-lock.json ]]; then
            pnpm ci
          else
            pnpm install
          fi

      - name: Run code generation
        run: |
          pnpm code
          pnpm codea

      - name: Commit and push changes (if any)
        shell: bash
        run: |
          set -euo pipefail

          if [[ -z "$(git status --porcelain)" ]]; then
            echo "No changes to commit."
            exit 0
          fi

          git config user.name "github-actions[bot]"
          git config user.email "github-actions[bot]@users.noreply.github.com"

          git add oraclepack.md oraclepack-all.md
          git commit -m "chore(codefetch): update generated outputs [skip ci]"
          git push
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

internal/exec/inject.go
```
package exec

import (
	"bufio"
	"regexp"
	"strings"
)

var (
	oracleCmdRegex = regexp.MustCompile(`^(\s*)oracle(\s+|$)`)
)

// InjectFlags scans a script and appends flags to any 'oracle' command invocation.
func InjectFlags(script string, flags []string) string {
	if len(flags) == 0 {
		return script
	}

	flagStr := " " + strings.Join(flags, " ")
	
	var result []string
	scanner := bufio.NewScanner(strings.NewReader(script))
	for scanner.Scan() {
		line := scanner.Text()
		if oracleCmdRegex.MatchString(line) {
			// Append flags to the line
			line += flagStr
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
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
			"oracle query 'hello' --verbose",
		},
		{
			"indented injection",
			"  oracle query 'hello'",
			[]string{"--verbose"},
			"  oracle query 'hello' --verbose",
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
			"echo 'start'\noracle query --debug\necho 'end'",
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
	"github.com/user/oraclepack/internal/pack"
)

// Runner handles the execution of shell scripts.
type Runner struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
}

// RunnerOptions configures a Runner.
type RunnerOptions struct {
	Shell       string
	WorkDir     string
	Env         []string
	OracleFlags []string
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
	}
}

// RunPrelude executes the prelude code.
func (r *Runner) RunPrelude(ctx context.Context, p *pack.Prelude, logWriter io.Writer) error {
	return r.run(ctx, p.Code, logWriter)
}

// RunStep executes a single step's code.
func (r *Runner) RunStep(ctx context.Context, s *pack.Step, logWriter io.Writer) error {
	code := InjectFlags(s.Code, r.OracleFlags)
	return r.run(ctx, code, logWriter)
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
	"github.com/charmbracelet/glamour"
	"github.com/user/oraclepack/internal/pack"
)

// RenderMarkdown renders markdown text as ANSI-styled text.
func RenderMarkdown(text string) (string, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return "", err
	}

	return r.Render(text)
}

// RenderStepCode renders a step's code block for preview.
func RenderStepCode(s pack.Step) (string, error) {
	md := "```bash\n" + s.Code + "\n```"
	return RenderMarkdown(md)
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
	got, err := RenderMarkdown(text)
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
	Summary   Summary      `json:"summary"`
	PackInfo  PackInfo     `json:"pack_info"`
	Steps     []StepReport `json:"steps"`
	GeneratedAt time.Time  `json:"generated_at"`
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
	ID        string `json:"id"`
	Status    string `json:"status"`
	ExitCode  int    `json:"exit_code"`
	Duration  time.Duration `json:"duration"`
	DurationMs int64         `json:"duration_ms"`
	Error     string `json:"error,omitempty"`
}
```

internal/tui/filter_test.go
```
package tui

import (
	"testing"

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
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/state"
)

type ViewState int

const (
	ViewSteps ViewState = iota
	ViewRunning
	ViewDone
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
	pack        *pack.Pack
	runner      *exec.Runner
	state       *state.RunState
	statePath   string
	
	width    int
	height   int
	
	viewState ViewState
	running   bool
	runAll    bool // State for sequential execution
	currentIdx int
	autoRun   bool // Config to auto-start on init
	
	// Filtering state
	allSteps     []item // Store all items to support dynamic filtering
	roiThreshold float64
	roiMode      string
	isFiltering  bool

	err      error
	logLines []string
	logChan  chan string
}

func NewModel(p *pack.Pack, r *exec.Runner, s *state.RunState, statePath string, roiThreshold float64, roiMode string, autoRun bool) Model {
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
	vp.SetContent("Press Enter to run selected, 'a' to run all filtered steps, 'f' to filter.")

	m := Model{
		list:         l,
		viewport:     vp,
		spinner:      sp,
		filterInput:  ti,
		pack:         p,
		runner:       r,
		state:        s,
		statePath:    statePath,
		autoRun:      autoRun,
		allSteps:     allItems,
		roiThreshold: roiThreshold,
		roiMode:      roiMode,
		logChan:      make(chan string, 100),
		viewState:    ViewSteps,
	}

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
type FinishedMsg struct{ Err error; ID string }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Global keys (Quit)
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.viewState {
	case ViewDone:
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "b":
				m.viewState = ViewSteps
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
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) handleWindowSize(msg tea.WindowSizeMsg) {
	m.width = msg.Width
	m.height = msg.Height
	m.list.SetSize(msg.Width/3, msg.Height-4)
	m.viewport.Width = msg.Width - (msg.Width / 3) - 6
	m.viewport.Height = msg.Height - 4
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
}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}

	if m.viewState == ViewDone {
		return m.viewDone()
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
		if filterStatus != "" {
			right = lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render(filterStatus) + "\n" + right
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, left, " | ", right)
}

func (m Model) viewDone() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).Render("Execution Complete")
	if m.err != nil {
		title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196")).Render("Execution Failed")
	}

	help := "[n] New Run  [r] Rerun  [b] Back to List  [q] Quit"
	
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
	StatusPending  Status = "pending"
	StatusRunning  Status = "running"
	StatusSuccess  Status = "success"
	StatusFailed   Status = "failed"
	StatusSkipped  Status = "skipped"
)

// RunState tracks the execution progress of an oracle pack.
type RunState struct {
	SchemaVersion int                   `json:"schema_version"`
	PackHash      string                `json:"pack_hash"`
	StartTime     time.Time             `json:"start_time"`
	StepStatuses  map[string]StepStatus `json:"step_statuses"`
}

// StepStatus holds the outcome of an individual step.
type StepStatus struct {
	Status    Status    `json:"status"`
	ExitCode  int       `json:"exit_code"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Error     string    `json:"error,omitempty"`
}
```

.config/skill/strategist-questions-oracle-pack/SKILL.md
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

.config/skill/strategist-questions-oracle-pack/assets/oracle-pack-template.md
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

.config/skill/strategist-questions-oracle-pack/references/attachment-minimization.md
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

.config/skill/strategist-questions-oracle-pack/references/inference-first-discovery.md
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

.config/skill/strategist-questions-oracle-pack/references/oracle-scratch-format.md
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