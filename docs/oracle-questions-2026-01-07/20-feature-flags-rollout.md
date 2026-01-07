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
