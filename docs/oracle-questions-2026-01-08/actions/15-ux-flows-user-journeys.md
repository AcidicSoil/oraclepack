Direct answer (UX/TUI workflows, journey states, expected outputs)

Core “run pack” journey: user selects an Action Pack, oraclepack parses it, then executes each step as literal shell via bash -lc ... (in repo root; no automatic cd to out_dir). Journey states: Selected → Parsed/Validated → Ready → Running(step N) → Step {Succeeded|Failed|Skipped|Blocked} → Completed/Resumable. Expected outputs: per-step logs plus persistent run artifacts like ticket-action-pack.state.json and ticket-action-pack.report.json (names explicitly referenced), alongside whatever files each step writes. 

Oraclepack_Compatibility_Issues

Overrides/validation journey is “oracle-only” today: in TUI, an “Overrides Wizard”/override-validation flow should only attach to steps containing oracle ... lines, because override injection/detection is anchored to oracle-prefixed commands (regex anchored to ^(\\s*)(oracle)\\b), and override validation runs oracle --dry-run summary on detected oracle invocations while skipping steps with no oracle invocations. Journey states: Overrides entry → Step classification {Oracle-managed|Shell-direct} → Overrides selected (Oracle-managed only) → Dry-run validation (Oracle-managed only) → Apply during execution. 

Oraclepack_Compatibility_Issues

ROI filtering journey: users can filter which steps appear/run by ROI; a known UX gotcha is that steps without ROI= may be skipped when an ROI threshold > 0 is used. Journey states: Filter configured → Steps partitioned {Included|Excluded (reason: ROI missing/too low)} → Run included set. Expected output: UI should surface counts and exclusion reasons to avoid “missing steps” confusion. 

Oraclepack_Compatibility_Issues

Ticketify “analysis pipeline” journey (Steps 01–07): user runs early steps to discover tickets and generate action/PRD scaffolding. Expected outputs include .oraclepack/ticketify/_tickets_index.json (Step 01), _actions.json + _actions.md (Step 02), .taskmaster/docs/tickets_prd.md (Step 03), and .oraclepack/ticketify/tm-complexity.json after Task Master parse/complexity/expand steps. Journey states: Inputs present (.tickets/) → Index/actions generated → PRD generated → TM expanded tasks available. 

Oraclepack_Compatibility_Issues

Ticketify “agent loop” journey (replacing placeholders 09–13 + 16): implied end-to-end headless workflow is Select → Implement → Verify → Draft PR. Expected outputs: Step 09 writes .oraclepack/ticketify/next.json; Step 10 writes .oraclepack/ticketify/codex-implement.md; Step 11 writes .oraclepack/ticketify/codex-verify.md and/or .oraclepack/ticketify/gemini-review.json; Step 16 writes .oraclepack/ticketify/PR.md. Journey states: Candidate chosen → Implementation attempted → Verification completed → PR draft produced.

Non-oracle tool execution journey (tm/codex/gemini): these tools run directly (not routed through oracle overrides). The UX implication is preflight detection and “blocked vs skipped” semantics: if the binary is missing on PATH, the step fails unless the step implements guard+skip logic; if the tool is interactive, it can block waiting for input. Journey states: Preflight {Available|Missing|Interactive-risk} → Execute or Skip (if guarded) → Surface outcome clearly.

“Agent-mode” selection for taskify-generated packs: tickets imply a mode switch (e.g., mode=codex / mode=gemini) that swaps an existing “autopilot entrypoint” step slot for agent implementation while keeping the 20-step contract intact. UX journey states: Mode selected → Step slot behavior changes (autopilot vs agent loop) → Same step-count/structure preserved. 

Oraclepack_Compatibility_Issues

Risks/unknowns

Exact TUI affordances are underspecified: where to present “oracle-managed vs shell-direct” labeling (step list badge, override wizard copy, help text) is not defined in the tickets. 

Oraclepack_Compatibility_Issues

Dispatcher semantics for non-oracle commands are unspecified: what “overrides” would mean for tm/codex/gemini (flags, validation method, dry-run equivalents) is explicitly unknown. 

Oraclepack_Compatibility_Issues

Step 11 default verification path is undecided (Codex vs Gemini vs both), which impacts the expected UX branching and success criteria. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Add a TUI-visible per-step classification (“Oracle-managed” iff the step contains an oracle invocation; otherwise “Shell-direct”), then run a fixture pack containing one oracle ... step and one tm/codex/gemini step t
