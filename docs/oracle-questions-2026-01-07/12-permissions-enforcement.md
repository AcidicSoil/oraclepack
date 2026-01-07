1) Direct answer (1–8 bullets, evidence-cited)

There is no explicit permission enforcement layer in the surfaced code (no authn/authz context, no roles/scopes/ACL objects, no guards/middleware imports). The CLI entrypoint is Cobra-based with flag globals (noTUI, oracleBin, outDir) and does not introduce any identity/permission concept. Evidence: internal/cli/root.go snippet shows only Cobra + basic flags; no auth-related imports/types.

Effective “permissions” are the OS user’s privileges at runtime. The system executes external commands via os/exec, so whatever the current process user can do (filesystem, network, process execution) the tool can do. Evidence: internal/exec/runner.go imports and uses os/exec and os.

Primary enforcement chokepoint is the execution boundary (Runner) — but it enforces execution, not authorization. If you need “who can run what,” the only practical choke is where commands are spawned (and where env/args are assembled). Evidence: internal/exec/runner.go is the dedicated Runner module and pulls overrides.

Pack parsing enables embedding runnable shell content; there is no shown policy layer restricting it. The parser extracts bash fenced blocks (regex), which strongly implies step payloads can contain shell scripts. **Evidence:** `internal/pack/parser.go` defines `bashFenceRegex = ... bash\n(.*?)\n``` ...`.

State/report generation has no permission gates; it simply materializes outputs. So any caller who can run the CLI can cause state/report writes wherever configured. Evidence: internal/report/generate.go constructs a report from RunState; internal/app/app.go config includes StatePath/Report… fields (from snippet) with no auth checks.

Likely bypass risks are “configuration-as-authority” rather than missing middleware. If a caller can control oracleBin, pack content, overrides, output/state paths, they effectively control what runs and where it writes—without an authorization decision point. Evidence: internal/cli/root.go exposes oracleBin/outDir; internal/exec/runner.go consumes overrides + executes commands; internal/pack/parser.go extracts bash payloads.

Inconsistent pattern: business logic lives in internal/app/*, but enforcement is absent at service-layer boundaries. There’s no “service-layer check” because there’s no identity/claims passed through app -> exec -> filesystem; everything is implicitly trusted input. Evidence: internal/app/app.go wires exec, pack, state with no auth context types.

2) Risks/unknowns (bullets)

Unknown whether any permission checks exist elsewhere (e.g., an HTTP server, MCP server wrapper, or API layer not included in the snippets). If those exist, they may be the real enforcement point.

Unknown how overrides are applied (env var allow/deny lists, argument escaping, path validation). internal/overrides/* was referenced but not provided.

Unknown whether pack execution is sandboxed (chroot, containers, seccomp, allowlisted commands). Nothing in the provided snippets suggests sandboxing, but absence of evidence isn’t proof.

Unknown whether state/output paths are validated (path traversal, writing outside workspace, overwriting). Config shows paths, but validation code wasn’t included.

3) Next smallest concrete experiment (exactly one action)

Run a repo-wide search to locate any authz/authn/permission enforcement and all chokepoints where requests/commands enter:

bash
Copy code
rg -n --hidden --no-ignore-vcs "(auth|authz|authorize|authorization|permission|rbac|acl|role|scope|policy|guard|middleware|claims|principal|subject|tenant)" .

4) If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

Attach any of the following if they exist (these are the most likely places where enforcement would live, and we haven’
