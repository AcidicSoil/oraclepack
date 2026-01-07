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
