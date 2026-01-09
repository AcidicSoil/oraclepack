Direct answer (1–10 bullets, evidence-cited)

Command execution boundary: oraclepack runs each step by executing the step body as shell via bash -lc ..., i.e., whatever commands appear will run with the oraclepack process’s OS privileges. Safe default: require an explicit “exec enable” gate (flag/env) plus per-step confirmation unless --yes; otherwise operate in validate/list-only mode. 

Oraclepack_Compatibility_Issues

Tool routing/authz boundary: oraclepack’s “special handling” (override injection/validation) only applies to commands that begin with oracle; tm/task-master, codex, and gemini run directly and are not protected by oracle-specific checks. Safe default: treat any non-oracle command as “raw exec” and block-by-default unless an allowlist (by prefix) is configured.

Interactivity boundary: tickets call out that interactive CLIs can block waiting for input. Safe default: force non-interactive execution (environment like CI=1, plus only “headless/non-interactive” subcommands) and add step timeouts so a blocked CLI cannot hang the run indefinitely.

Binary availability boundary: missing binaries on PATH cause step failures; tickets explicitly recommend command -v ... guards with “skip” behavior. Safe default: (a) preflight required tools per pack (oracle/tm/codex/gemini), (b) default to “skip with a clear message” for optional tools, and (c) provide a strict mode that turns skips into hard failures.

File read boundary: tickets list common failure points like missing .tickets/ and emphasize steps run in the project root (not an isolated out_dir). Safe default: restrict file reads to the workspace root + explicitly allowed subtrees (e.g., .tickets/, tickets/) and deny absolute-path reads or .. traversal outside allowed roots.

File write boundary: expected artifacts are written into deterministic project paths like .oraclepack/ticketify/* and .taskmaster/docs/* (and per-step outputs). Safe default: restrict writes to a small allowlist of directories (e.g., .oraclepack/**, .taskmaster/**, and the declared pack out_dir), deny writes outside the repo, and reject absolute paths / parent traversal.

Network/egress boundary (implied by tool choices): the proposed automation inserts headless gemini and codex exec steps and notes missing provider/API key configuration as a failure mode. Safe default: make “network-enabled execution” an explicit mode (opt-in), and otherwise assume offline (or at least warn and require confirmation before running steps that invoke these tools).

Secrets/logging boundary (implied by provider keys + artifact writing): the workflow depends on external provider configuration/API keys and produces multiple on-disk reports/artifacts. Safe default: never dump environment variables in logs, redact likely secret patterns in streamed output before writing report JSON/MD, and prohibit packs from writing secret material into .oraclepack/ticketify/* outputs.

Risks/unknowns (bullets)

The desired override/validation semantics for tm/task-master, codex, and gemini are explicitly unspecified (which flags apply, how validation works), so any “first-class tool authz” policy could be incomplete or wrong.

It’s not specified where the authoritative user-facing permission model should live (repo docs vs TUI help), which increases the chance the real boundary behavior is misunderstood by users. 

Oraclepack_Compatibility_Issues

“Agent-mode” selection mechanism (CLI flag vs TUI option vs config) is not defined, which affects how you safely expose network/exec enablement without surprising defaults. 

Oraclepack_Compatibility_Issues

Next smallest concrete experiment (exactly one action)

Implement a “preflight policy report” in oraclepack (or a companion script) that parses a pack and prints:
