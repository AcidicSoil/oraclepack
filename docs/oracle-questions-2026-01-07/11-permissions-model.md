Direct answer (1–8 bullets, evidence-cited)

No first-class permission model is evident in the provided code: none of the shown entrypoints/config/state types include a notion of user identity, roles, scopes, ACLs, or authorization checks (e.g., the CLI root command only defines local flags like noTUI, oracleBin, outDir). Evidence: internal/cli/root.go (flag surface only; no auth/identity inputs shown).

Execution appears to be “whoever can run the CLI can do everything the CLI does.” The runner shells out to an external binary and interacts with the local filesystem; no gating by role/scope is shown. Evidence: internal/exec/runner.go (uses os/exec + os without any authz hooks shown).

“Who can do what” is therefore effectively determined by OS-level permissions and environment constraints, not by application-level RBAC/ACL logic. Evidence: internal/app/app.go (config is paths like PackPath/StatePath/Report…; no policy/role configuration shown) + internal/app/run.go (reads/writes JSON/state/report to local paths; no authz shown).

Any real access control is likely delegated to external dependencies (e.g., the invoked oracle binary and filesystem permissions), not implemented in oraclepack itself, based on the visible integration boundary. Evidence: internal/exec/runner.go (external command invocation) + internal/report/generate.go (report generation from run state; no authz metadata).

Minimal set of files that encode “who can do what” (based on what’s been provided)

None (no app-level permission model is shown in the provided files). Practically, behavior is governed by:

OS user permissions (read/write/execute) for the invoked binary and directories (PackPath, StatePath, report/output dirs). Evidence: internal/app/app.go, internal/app/run.go.

External tool behavior (the oracle binary). Evidence: internal/exec/runner.go.

Risks/unknowns (bullets)

The repository may contain an unshared permissions/auth subsystem (e.g., internal/auth/*, HTTP server middleware, policy engine) that wasn’t included in the snippets.

The external oracle binary may implement its own authn/authz, which would materially change the true permission model even if oraclepack itself has none.

There may be implicit “permissions” via overrides/config resolution (e.g., internal/overrides) that constrain execution; not enough code was provided to verify.

Next smallest concrete experiment (exactly one action)

Run a repo-wide search for authorization primitives: rg -n "(rbac|role|roles|scope|scopes|acl|permission|authorize|authn|authz|oauth|jwt|policy|casbin)" .

If evidence is insufficient, name the exact missing file/path pattern(s) to attach next

internal/**/auth*.go

internal/**/perm*.go

internal/**/{rbac,acl,policy}*.go

internal/**/middleware*.go

internal/**/http*/**/*.go and/or internal/**/server*/**/*.go

cmd/** (any additional CLIs/entrypoints beyond internal/cli/root.go)

Any config/schema files that might define roles/scopes (e.g., **/*config*.*, **/*schema*.*, **/*policy*.*)
