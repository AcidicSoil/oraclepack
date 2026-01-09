
Oraclepack\_Compatibility\_Issues.md

File

Strategist question #01 (ticket-driven, group: actions) Reference: actions Category: contracts/interfaces Horizon: Immediate ROI: 4.8 (impact=6, confidence=0.80, effort=1) Question: Using the attached tickets as the primary context, list the public surface changes implied by the tickets (CLI/TUI/API/interfaces/contracts); call out backwards-compat constraints. Constraints: None Non-goals: None Answer format: 1) Direct answer (1–10 bullets, evidence-cited) 2) Risks/unknowns (bullets) 3) Next smallest concrete experiment (exactly one action) 4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next.

---

1. Direct answer (public surface changes; backwards-compat called out)

* Clarify/document Action Pack execution semantics: steps run as `bash -lc ...`, and oraclepack’s “special handling” (flag injection/override validation) only applies to commands that begin with `oracle` (today). This is a user-facing CLI/TUI documentation surface change (README and/or `oraclepack-tui.md`\-style docs).

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Dispatcher/command-detection contract expansion: update command detection from an `oracle`\-anchored regex (noted as `^(\\s*)(oracle)\\b`) so non-`oracle` tools (`tm`/`task-master`, `codex`, `gemini`) can be treated as first-class step commands for override/dispatch purposes. Back-compat constraint: preserve existing behavior for `oracle ...` commands.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Override validation behavior change (TUI/CI surface): today validation runs `oracle --dry-run summary` on detected oracle invocations and skips steps without oracle invocations; tickets imply extending or restructuring validation so steps containing `tm`/`task-master`, `codex`, `gemini` are not silently excluded. Back-compat constraint: do not regress the current oracle-only validation flow.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* `ticket-action-pack.md` content contract change: replace placeholder steps (explicitly 09–13 and 16) with headless `gemini` + non-interactive `codex exec` automation, while keeping the pack ingestible (single bash fence, 20-step structure). Back-compat constraint: keep Steps 01–07 semantics unchanged.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* New output artifact interface for the action workflow: introduce/standardize generated files under `.oraclepack/ticketify/` (e.g., `next.json`, `codex-implement.md`, `codex-verify.md` and/or `gemini-review.json`, `PR.md`) as “expected outputs” of those steps. Back-compat constraint: paths must match the step examples; don’t move or rename without versioning.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Tool-availability/timeout behavior requirements become part of the pack’s operational contract: add `command -v ...` guards and documented “skip” behavior to avoid hard failures when `codex`/`gemini` are missing; also explicitly call out “interactive CLI blocks” risk. Back-compat constraint: default runs shouldn’t newly fail just because optional tools aren’t installed.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Taskify Action Pack generator surface: add an “agent-mode” option for taskify-generated packs (suggested `mode=codex` / `mode=gemini`) that swaps the existing autopilot entrypoint step with an agent implementation step, without changing the 20-step contract. Back-compat constraint: default mode should remain the current behavior; agent-mode must be opt-in.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* “Pack schema” remains a hard compatibility boundary: multiple tickets restate that the Action Pack must remain oraclepack-ingestible (single `bash` fence, fixed step structure). Any added behavior must fit inside existing step slots.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* User-facing guidance around ROI filtering becomes part of the public contract (docs and/or TUI hints): ROI filtering may skip steps lacking `ROI=` when a threshold is used; tickets call this out as a common gotcha to document. Back-compat constraint: don’t change ROI semantics silently; document them.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

1. Risks/unknowns

* Exact CLI/TUI flag names, config keys, and where mode selection lives are not specified (ticket explicitly leaves mode-selection mechanism open).

    Oraclepack\_Compatibility\_Issues

* Applying “overrides” to `tm`/`codex`/`gemini` is underspecified (which flags, how validation works, whether wrapping is required).

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Exact repository file locations for: the regex/dispatcher, TUI validation logic, `ticket-action-pack.md`, and the taskify pack generator are not provided, so interface impact cannot be mapped to concrete codepaths yet.

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

* Headless/non-interactive guarantees for `codex` and `gemini` are assumed; if either tool still prompts, runs can hang unless enforced/timeouted at the step level.

    Oraclepack\_Compatibility\_Issues

1. Next smallest concrete experiment (exactly one action)

* Add a tiny fixture Action Pack with one `oracle ...` line and one `tm` (or `codex`/`gemini`) line, then run `oraclepack run --no-tui` to confirm (a) both commands execute under `bash -lc`, and (b) only the `oracle ...` line receives oraclepack override injection/validation today (baseline behavior before changing dispatcher logic).

    Oraclepack\_Compatibility\_Issues

    Oraclepack\_Compatibility\_Issues

1. If evidence is insufficient, exact missing file/path pattern(s) to attach next

* The actual Action Pack template: `**/ticket-action-pack.md` (or `**/*ticket-action-pack*.md`)

* Taskify Action Pack generator/template: `**/*taskify*pack*.md`, `**/*taskify*template*.md`, and/or the generator code under `internal/**/taskify/**` (or equivalent)

* Dispatcher/override injection + detection code: `internal/**/{override,overrides,inject,dispatch}*.go` and `internal/**/tui/**` (for validation logic)

* Current user-facing docs that must be updated: `README*.md`, `docs/**oraclepack*tui*.md`, `docs/**action*pack*.md`

---
