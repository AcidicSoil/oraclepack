Direct answer (1–10 bullets, evidence-cited)

Treat “feature flags” as compat-preserving runtime gates for the three ticketed deltas: (a) dispatcher/override/validation expansion beyond oracle-prefixed commands, (b) Action Pack placeholder replacement with headless gemini + non-interactive codex exec, and (c) taskify “agent-mode” generation that swaps the same step slot while keeping the contract unchanged.

Define scopes explicitly and separately (so you can roll out safely):

Pack generation scope (taskify/template toggles): taskify.agent_mode / actionpack.enable_agent_steps

Execution/dispatcher scope (step parsing + injection): dispatcher.v2_enabled, overrides.targets

Validation scope (what gets “mode-2 validated”): validation.targets, validation.strict

Defaults must preserve current behavior for existing users: dispatcher.v2_enabled=false, overrides.targets=["oracle"], validation.targets=["oracle"], and any new Codex/Gemini automation in packs must be opt-in at generation time. This aligns with the explicit constraint “must preserve existing oracle command behavior” and the current oracle-only injection/validation behavior.

Rollout the dispatcher expansion in three stages behind flags (reduces blast radius given unclear non-oracle semantics):

“detect-only” (collect telemetry on tm/task-master/codex/gemini occurrences; no injection/validation changes),

“validate-only” (include non-oracle steps in the override/validation pipeline as first-class targets but keep behavior non-blocking unless validation.strict=true),

“enforce” (apply overrides to configured tool prefixes, with tool-specific validators). This matches the ticket’s ask to stop silently excluding non-oracle steps.

For Action Pack headless automation (T3), gate pack content changes with a generation-time flag (e.g., actionpack.enable_agent_steps=true) and require command-availability guards with “skip” semantics (not hard fail) when codex/gemini aren’t on PATH.

For taskify “agent-mode” (T4), add a generation-time mode switch (ticket-suggested mode=codex / mode=gemini) and implement it as a swap of the existing autopilot step slot, keeping the “20-step contract intact” and placing the insertion after Task Master expansion artifacts exist. 

Oraclepack_Compatibility_Issues

Telemetry: always emit the resolved feature-flag set and derived “effective targets” into the machine-readable run report and persisted run state, including counts of detected tool prefixes per step, which steps were targeted, which were skipped (and why), and whether validation ran. This is consistent with the existing direction to write deterministic state.json and a machine-readable summary report (run.json) with stable schema/versioning.

TUI/UX: expose these flags as a single “Experimental / Beta” panel that expands into (1) dispatcher mode, (2) override targets, and (3) validation strictness; keep them off by default and require explicit opt-in. This avoids the current confusion about oraclepack “running everything through oracle” and makes the scope visible.

Guardrails: because steps execute via bash -lc and interactive CLIs can block, require “headless/non-interactive” sub-flags for any Codex/Gemini execution paths and default to “skip with recorded reason” when headless requirements aren’t met.

Risks/unknowns (bullets)

The tickets explicitly do not specify the exact desired override semantics for tm/task-master, codex, and gemini (which flags apply, how validation should work), so an “enforce overrides for non-oracle tools” flag risks doing the wrong thing without a tool-specific contract. 

Oraclepack_Compatibility_Issues

Interactive behavior and missing binaries remain a runtime risk for any headless automation steps unless every path is guarded and defaults to skip.

The selection mechanism for taskify agent-mode (CLI flag vs TUI vs config) is not specified in the ticket summary; inconsistent selection surfaces can fragment behavior and telemetry. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Implement --dispatcher=v2-detect (or equivalent) that onl
