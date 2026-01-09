Direct answer (top integration points; contract/config that declares it)

External tool prefixes (oracle / tm / task-master / codex / gemini): declared as a closed set of “tool kinds” in internal/tools/types.go via ToolKind + the registry map (canonical binary name, plus default args like codex exec), and bounded by PresenceChecker.DetectBinary(name) as the abstraction for “is this binary available and where?” (reference: internal/tools/types.go, in the excerpt you attached).

Oracle CLI (primary runtime dependency): oraclepack is explicitly a wrapper/runner for packs whose steps invoke the oracle CLI; docs call out the oracle binary as a core external integration and that it must be on PATH (or otherwise supplied via configuration/flags). 

oraclepack-llms-full

 

oraclepack-llms-full

 

oraclepack-llms-full

Shell environment / subprocess execution: step bodies are executed as shell scripts (default bash), so the runtime boundary includes the user’s shell semantics (login shell env, PATH/aliases, quoting) and subprocess behavior. The docs enumerate “shell runtime (bash default)” as a first-class external integration point. 

oraclepack-llms-full

 

oraclepack-llms-full

Filesystem (inputs + persistent artifacts): oraclepack consumes a Markdown pack file and produces on-disk artifacts (logs/state/report). The pack contract also commonly uses variables like out_dir and step-level --write-output paths, which become filesystem writes during execution. 

oraclepack-llms-full

 

oraclepack-llms-full

 

oraclepack-llms-full

ChatGPT/remote target selection “by URL” (via Oracle flags/config): the integration boundary is “oraclepack → oracle flags/config”, especially --chatgpt-url / “project URL” targeting. Oracle’s own config surface includes browser.chatgptUrl (and precedence rules), and oraclepack’s UX explicitly calls out project URL management and runtime flag overrides. 

oraclepack-llms-full

 

oracle-llms-full

Risks/unknowns

The excerpted internal/tools/types.go shows tool-kind registration, but the evidence bundle does not include the call sites: it’s unclear whether these tool kinds affect (a) pack parsing, (b) oracle-invocation scanning/validation, (c) execution rewriting, or (d) only UX/diagnostics.

The exact “contract” for overriding tool binaries (e.g., --oracle-bin, per-tool env vars, config files) isn’t evidenced in the attached corpus; docs mention override concepts, but the authoritative flag/config declarations aren’t included.

Shell execution correctness is highly sensitive to quoting/multiline continuation; docs acknowledge injection/rewriting heuristics exist, but the exact transform rules and edge-case handling aren’t in the provided evidence.

Filesystem safety boundaries (path traversal prevention, atomic writes, permissions, log redaction) can’t be verified without the state/report/log writer code and CLI flag definitions.

Next smallest concrete experiment (exactly one action)

Run a targeted symbol trace to locate the real integration call sites:

bash
Copy code
rg -n "ToolKind|ToolOracle|ToolTM|ToolTaskMaster|ToolCodex|ToolGemini|PresenceChecker|DetectBinary|InjectFlags|ApplyChatGPTURL|oracle_cmd|oracle_flags|--oracle" internal/ cmd/


If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

Tool-prefix parsing + usage: internal/tools/*.go (especially anything that parses a command/prefix string and any PresenceChecker implementation).

CLI flag/config declarations that wire binaries/shell/env/out-dir: internal/cli/*.go and cmd/oraclepack/*.go.

Execution rewriting
