# oraclepack-mcp-server (MCP wrapper for oraclepack + taskify helpers)

## tree

```text
oraclepack-mcp-server/
  README.md
  requirements.txt
  oraclepack_mcp_server/
    __init__.py
    __main__.py
    config.py
    security.py
    oraclepack_cli.py
    taskify.py
    server.py
```

---

```md
<!-- path: oraclepack-mcp-server/README.md -->
# oraclepack-mcp-server

Expose `oraclepack` (validate/list/run) plus a few Stage-3 (oraclepack-taskify) helpers as **MCP tools**, so an agent can:
- validate and inspect packs
- run packs non-interactively (`--no-tui --yes --run-all`) to produce artifacts
- validate Stage-2 outputs (01-*.md..20-*.md)
- validate Stage-3 Action Packs (single ```bash fence, step headers, etc.)
- summarize Stage-3 artifacts (`_actions.json`, PRD, Task Master outputs, etc.)

## Install

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

## Configure (recommended)

Environment variables:

- `ORACLEPACK_ALLOWED_ROOTS` (default: current working directory)
  - Colon-separated list of allowed filesystem roots the server may read from.
  - Example: `ORACLEPACK_ALLOWED_ROOTS=/repo:/tmp/oracle-out`
- `ORACLEPACK_BIN` (default: `oraclepack`) – path to the oraclepack CLI
- `ORACLEPACK_WORKDIR` (default: current working directory)
  - Where packs are executed from (important for relative paths).
- `ORACLEPACK_ENABLE_EXEC` (default: `0`)
  - Must be `1` to enable `oraclepack_run_pack` and `oraclepack_taskify_run_action_pack`.
- `ORACLEPACK_CHARACTER_LIMIT` (default: `25000`) – truncate large stdout/stderr
- `ORACLEPACK_MAX_READ_BYTES` (default: `500000`) – max bytes read from a file

## Run (stdio)

```bash
# Stdio transport is the simplest local integration.
python -m oraclepack_mcp_server --transport stdio
```

## Run (Streamable HTTP)

```bash
python -m oraclepack_mcp_server --transport streamable-http
```

## Tools

- `oraclepack_validate_pack`
- `oraclepack_list_steps`
- `oraclepack_run_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)
- `oraclepack_read_file`
- `oraclepack_taskify_detect_stage2`
- `oraclepack_taskify_validate_stage2`
- `oraclepack_taskify_validate_action_pack`
- `oraclepack_taskify_artifacts_summary`
- `oraclepack_taskify_run_action_pack` (requires `ORACLEPACK_ENABLE_EXEC=1`)

## Typical Stage-3 usage

1) Detect/validate Stage-2 outputs (directory or single-pack)
2) Validate the Action Pack markdown
3) Run the Action Pack via `oraclepack run ...`
4) Summarize produced artifacts
```

```txt
# path: oraclepack-mcp-server/requirements.txt
mcp>=1.0.0
pydantic>=2.0.0
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
__all__ = []
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/config.py
from __future__ import annotations

import os
from dataclasses import dataclass
from pathlib import Path


def _truthy(value: str | None) -> bool:
    if value is None:
        return False
    return value.strip().lower() in {"1", "true", "yes", "y", "on"}


@dataclass(frozen=True)
class ServerConfig:
    allowed_roots: tuple[Path, ...]
    oraclepack_bin: str
    workdir: Path
    enable_exec: bool
    character_limit: int
    max_read_bytes: int


def load_config() -> ServerConfig:
    # Allowed roots: colon-separated. Default to CWD.
    roots_raw = os.environ.get("ORACLEPACK_ALLOWED_ROOTS")
    if roots_raw:
        roots = tuple(Path(p).expanduser().resolve() for p in roots_raw.split(":") if p.strip())
    else:
        roots = (Path.cwd().resolve(),)

    oraclepack_bin = os.environ.get("ORACLEPACK_BIN", "oraclepack")
    workdir = Path(os.environ.get("ORACLEPACK_WORKDIR", str(Path.cwd()))).expanduser().resolve()

    enable_exec = _truthy(os.environ.get("ORACLEPACK_ENABLE_EXEC", "0"))

    character_limit = int(os.environ.get("ORACLEPACK_CHARACTER_LIMIT", "25000"))
    max_read_bytes = int(os.environ.get("ORACLEPACK_MAX_READ_BYTES", "500000"))

    return ServerConfig(
        allowed_roots=roots,
        oraclepack_bin=oraclepack_bin,
        workdir=workdir,
        enable_exec=enable_exec,
        character_limit=character_limit,
        max_read_bytes=max_read_bytes,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/security.py
from __future__ import annotations

from pathlib import Path


class PathNotAllowedError(ValueError):
    pass


def resolve_under_roots(path: Path, allowed_roots: tuple[Path, ...]) -> Path:
    """Resolve a path and enforce it lives under at least one allowed root."""
    p = path.expanduser().resolve()

    for root in allowed_roots:
        r = root.expanduser().resolve()
        try:
            p.relative_to(r)
            return p
        except ValueError:
            continue

    raise PathNotAllowedError(
        f"Path not allowed (outside allowed roots). path={p} roots={[str(r) for r in allowed_roots]}"
    )


def safe_read_text(path: Path, max_bytes: int) -> tuple[str, bool]:
    """Read up to max_bytes from a file as UTF-8 (replace errors)."""
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data.decode("utf-8", errors="replace"), truncated


def safe_read_bytes(path: Path, max_bytes: int) -> tuple[bytes, bool]:
    data = path.read_bytes()
    truncated = False
    if len(data) > max_bytes:
        data = data[:max_bytes]
        truncated = True
    return data, truncated
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
from __future__ import annotations

import asyncio
import os
import time
from dataclasses import dataclass
from pathlib import Path


@dataclass
class CmdResult:
    ok: bool
    exit_code: int
    duration_s: float
    stdout: str
    stderr: str
    stdout_truncated: bool
    stderr_truncated: bool


def _truncate(s: str, limit: int) -> tuple[str, bool]:
    if limit <= 0:
        return s, False
    if len(s) <= limit:
        return s, False
    return s[:limit], True


async def run_cmd(
    argv: list[str],
    cwd: Path,
    timeout_s: int,
    env: dict[str, str] | None,
    character_limit: int,
) -> CmdResult:
    start = time.time()

    proc = await asyncio.create_subprocess_exec(
        *argv,
        cwd=str(cwd),
        env=(os.environ | (env or {})),
        stdout=asyncio.subprocess.PIPE,
        stderr=asyncio.subprocess.PIPE,
    )

    try:
        out_b, err_b = await asyncio.wait_for(proc.communicate(), timeout=timeout_s)
    except asyncio.TimeoutError:
        proc.kill()
        await proc.communicate()
        duration = time.time() - start
        return CmdResult(
            ok=False,
            exit_code=124,
            duration_s=duration,
            stdout="",
            stderr=f"Timed out after {timeout_s}s: {' '.join(argv)}",
            stdout_truncated=False,
            stderr_truncated=False,
        )

    duration = time.time() - start
    out = out_b.decode("utf-8", errors="replace") if out_b else ""
    err = err_b.decode("utf-8", errors="replace") if err_b else ""

    out, out_tr = _truncate(out, character_limit)
    err, err_tr = _truncate(err, character_limit)

    exit_code = proc.returncode if proc.returncode is not None else 1
    return CmdResult(
        ok=(exit_code == 0),
        exit_code=exit_code,
        duration_s=duration,
        stdout=out,
        stderr=err,
        stdout_truncated=out_tr,
        stderr_truncated=err_tr,
    )
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
from __future__ import annotations

import re
from dataclasses import dataclass
from datetime import date
from pathlib import Path


@dataclass
class Stage2Resolution:
    kind: str  # "dir" | "file"
    stage2_path: Path
    out_dir: Path
    notes: list[str]


PREFIXES = [f"{i:02d}" for i in range(1, 21)]


def _is_iso_date(s: str) -> bool:
    # Minimal heuristic for YYYY-MM-DD
    return bool(re.fullmatch(r"\d{4}-\d{2}-\d{2}", s))


def validate_stage2_dir(out_dir: Path) -> dict:
    missing: list[str] = []
    ambiguous: dict[str, list[str]] = {}
    selected: dict[str, str] = {}

    for pfx in PREFIXES:
        matches = sorted(out_dir.glob(f"{pfx}-*.md"))
        if len(matches) == 0:
            missing.append(f"{pfx}-*.md")
        elif len(matches) > 1:
            ambiguous[pfx] = [m.name for m in matches]
        else:
            selected[pfx] = matches[0].name

    ok = (not missing) and (not ambiguous)
    return {
        "ok": ok,
        "out_dir": str(out_dir),
        "selected": selected,
        "missing": missing,
        "ambiguous": ambiguous,
    }


def _lexi_newest(paths: list[Path]) -> Path | None:
    # Deterministic: lexicographic max
    return sorted(paths, key=lambda p: p.name)[-1] if paths else None


def detect_stage2(stage2_path: str, repo_root: Path) -> Stage2Resolution:
    notes: list[str] = []

    if stage2_path != "auto":
        p = (repo_root / stage2_path).expanduser()
        if p.exists() and p.is_dir():
            return Stage2Resolution(kind="dir", stage2_path=p.resolve(), out_dir=p.resolve(), notes=["explicit dir"])
        if p.exists() and p.is_file():
            # out_dir rules from skill: if under docs/oracle-questions-YYYY-MM-DD/ then parent/oracle-out else oracle-out
            p_res = p.resolve()
            out = repo_root / "oracle-out"
            parts = list(p_res.parts)
            if "docs" in parts:
                try:
                    idx = parts.index("docs")
                    # docs/oracle-questions-YYYY-MM-DD/.../oracle-pack-YYYY-MM-DD.md
                    if idx + 1 < len(parts) and parts[idx + 1].startswith("oracle-questions-"):
                        out = Path(*parts[: idx + 2]) / "oracle-out"
                except ValueError:
                    pass
            return Stage2Resolution(kind="file", stage2_path=p_res, out_dir=out.resolve(), notes=["explicit file"])
        raise FileNotFoundError(f"stage2_path not found: {p}")

    # auto discovery (best-effort, deterministic ordering)
    searched: list[str] = []

    # 1) repo_root/oracle-out
    candidate = repo_root / "oracle-out"
    searched.append(str(candidate))
    if candidate.is_dir():
        v = validate_stage2_dir(candidate)
        if v["ok"]:
            notes.append("auto: selected repo_root/oracle-out")
            return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 2) docs/oracle-questions-YYYY-MM-DD/oracle-out (newest by lexicographic date suffix)
    docs = repo_root / "docs"
    if docs.is_dir():
        qdirs = [p for p in docs.glob("oracle-questions-*") if p.is_dir()]
        # deterministic: sort by name and take newest
        newest_q = _lexi_newest(qdirs)
        if newest_q:
            candidate = newest_q / "oracle-out"
            searched.append(str(candidate))
            if candidate.is_dir():
                v = validate_stage2_dir(candidate)
                if v["ok"]:
                    notes.append(f"auto: selected {candidate}")
                    return Stage2Resolution(kind="dir", stage2_path=candidate.resolve(), out_dir=candidate.resolve(), notes=notes)

    # 3) single-pack form (newer): look for docs/oracle-pack-*.md or docs/oraclepacks/oracle-pack-*.md
    file_candidates: list[Path] = []
    if docs.is_dir():
        file_candidates += list(docs.glob("oracle-pack-*.md"))
        file_candidates += list((docs / "oraclepacks").glob("oracle-pack-*.md"))
        file_candidates += list(docs.glob("oracle-questions-*/oraclepacks/oracle-pack-*.md"))

    newest_file = _lexi_newest(sorted([p for p in file_candidates if p.is_file()], key=lambda p: p.name))
    if newest_file:
        notes.append(f"auto: selected single-pack {newest_file}")
        out = repo_root / "oracle-out"
        # If under docs/oracle-questions-YYYY-MM-DD/..., default out_dir there.
        if "docs" in newest_file.parts:
            try:
                idx = newest_file.parts.index("docs")
                if idx + 1 < len(newest_file.parts) and newest_file.parts[idx + 1].startswith("oracle-questions-"):
                    out = Path(*newest_file.parts[: idx + 2]) / "oracle-out"
            except ValueError:
                pass
        return Stage2Resolution(kind="file", stage2_path=newest_file.resolve(), out_dir=out.resolve(), notes=notes)

    raise FileNotFoundError(
        "stage2_path=auto could not resolve Stage-2 outputs. Searched:\n- " + "\n- ".join(searched)
    )


def validate_action_pack(pack_path: Path) -> dict:
    text = pack_path.read_text(encoding="utf-8", errors="replace")

    bash_fence = re.findall(r"(?m)^\s*```bash\s*$", text)
    any_fence = re.findall(r"(?m)^\s*```\s*", text)

    errors: list[str] = []
    if len(bash_fence) != 1:
        errors.append(f"expected exactly one ```bash fence; found {len(bash_fence)}")
    if len(any_fence) != 2:
        # One opening and one closing fence expected, and it must be a bash fence.
        errors.append(f"expected no other code fences; found {len(any_fence)} total fences")

    # Extract bash block content if possible
    bash_block = ""
    m = re.search(r"```bash\s*\n(?P<body>[\s\S]*?)\n```\s*", text)
    if m:
        bash_block = m.group("body")

    # Validate step headers inside bash fence
    step_headers = re.findall(r"(?m)^\s*#\s*(\d{2})\)\s+.*$", bash_block)
    if not step_headers:
        errors.append("no step headers found inside the bash fence (expected lines like '# 01) ...')")
    else:
        # Ensure they start at 01 and are strictly increasing by 1.
        nums = [int(x) for x in step_headers]
        if nums[0] != 1:
            errors.append(f"first step must be 01; got {nums[0]:02d}")
        for prev, cur in zip(nums, nums[1:]):
            if cur != prev + 1:
                errors.append(f"step numbers must be sequential; got {prev:02d} then {cur:02d}")

    return {
        "ok": len(errors) == 0,
        "pack_path": str(pack_path),
        "step_count": len(step_headers),
        "errors": errors,
    }


def default_pack_path(today: date | None = None) -> str:
    d = today or date.today()
    return f"docs/oracle-actions-pack-{d.isoformat()}.md"
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/server.py
from __future__ import annotations

import json
from enum import Enum
from pathlib import Path
from typing import Any

from pydantic import BaseModel, Field
from mcp.server.fastmcp import FastMCP

from .config import load_config
from .security import resolve_under_roots, safe_read_text, PathNotAllowedError
from .oraclepack_cli import run_cmd
from .taskify import detect_stage2, validate_stage2_dir, validate_action_pack


class ResponseFormat(str, Enum):
    markdown = "markdown"
    json = "json"


class PackPathInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class RunPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the pack markdown file")
    out_dir: str | None = Field(default=None, description="Output directory for step execution (passes --out-dir).")

    no_tui: bool = Field(default=True, description="If true, pass --no-tui")
    yes: bool = Field(default=True, description="If true, pass --yes")
    run_all: bool = Field(default=True, description="If true, pass --run-all")

    resume: bool = Field(default=False, description="If true, pass --resume")
    stop_on_fail: bool = Field(default=True, description="If true, pass --stop-on-fail (default true)")

    roi_threshold: float = Field(default=0.0, description="Pass --roi-threshold")
    roi_mode: str = Field(default="over", description="Pass --roi-mode ('over' or 'under')")

    timeout_s: int = Field(default=3600, description="Hard timeout for the oraclepack process")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class ReadFileInput(BaseModel):
    path: str = Field(..., description="Path to a file within ORACLEPACK_ALLOWED_ROOTS")
    max_bytes: int | None = Field(default=None, description="Override max bytes read")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyDetectStage2Input(BaseModel):
    stage2_path: str = Field(default="auto", description="Dir or file, or 'auto'")
    repo_root: str = Field(default=".", description="Repo root for relative resolution")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateStage2Input(BaseModel):
    out_dir: str = Field(..., description="Directory that should contain 01-*.md..20-*.md")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyValidateActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to Stage-3 Action Pack markdown")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyArtifactsSummaryInput(BaseModel):
    out_dir: str = Field(..., description="Stage-3 out_dir (where _actions.json etc are written)")
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


class TaskifyRunActionPackInput(BaseModel):
    pack_path: str = Field(..., description="Path to the Stage-3 Action Pack markdown")
    out_dir: str | None = Field(default=None, description="Pass --out-dir for execution")
    timeout_s: int = Field(default=7200)
    response_format: ResponseFormat = Field(default=ResponseFormat.markdown)


cfg = load_config()

mcp = FastMCP(
    name="oraclepack-mcp-server",
    # For production Streamable HTTP deployments, stateless_http + json_response is recommended.
    # Clients may override by running behind an ASGI app if needed.
    stateless_http=True,
    json_response=True,
)


def _md_codeblock(lang: str, content: str) -> str:
    return f"```{lang}\n{content}\n```"


def _format_cmd_result(result: Any, response_format: ResponseFormat) -> Any:
    if response_format == ResponseFormat.json:
        return {
            "ok": result.ok,
            "exit_code": result.exit_code,
            "duration_s": result.duration_s,
            "stdout": result.stdout,
            "stderr": result.stderr,
            "stdout_truncated": result.stdout_truncated,
            "stderr_truncated": result.stderr_truncated,
        }

    lines = []
    lines.append(f"**ok**: {result.ok}")
    lines.append(f"**exit_code**: {result.exit_code}")
    lines.append(f"**duration_s**: {result.duration_s:.2f}")
    lines.append("")

    if result.stdout:
        lines.append("## stdout")
        lines.append(_md_codeblock("text", result.stdout))
        if result.stdout_truncated:
            lines.append(f"(stdout truncated to {cfg.character_limit} chars)")

    if result.stderr:
        lines.append("## stderr")
        lines.append(_md_codeblock("text", result.stderr))
        if result.stderr_truncated:
            lines.append(f"(stderr truncated to {cfg.character_limit} chars)")

    return "\n".join(lines)


def _ensure_exec_enabled() -> None:
    if not cfg.enable_exec:
        raise PermissionError(
            "Execution is disabled. Set ORACLEPACK_ENABLE_EXEC=1 to enable pack execution tools."
        )


@mcp.tool(
    name="oraclepack_validate_pack",
    annotations={
        "title": "Validate oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_validate_pack(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "validate", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_list_steps",
    annotations={
        "title": "List steps in an oraclepack pack",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_list_steps(params: PackPathInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv = [cfg.oraclepack_bin, "list", str(pack_path)]
    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=120, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_run_pack",
    annotations={
        "title": "Run an oraclepack pack (non-interactive)",
        "readOnlyHint": False,
        "destructiveHint": True,
        "idempotentHint": False,
        "openWorldHint": True,
    },
)
async def oraclepack_run_pack(params: RunPackInput) -> Any:
    _ensure_exec_enabled()

    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    argv: list[str] = [cfg.oraclepack_bin]
    if params.no_tui:
        argv += ["--no-tui"]
    if params.out_dir:
        out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
        argv += ["--out-dir", str(out_dir)]

    argv += ["run"]

    if params.yes:
        argv += ["--yes"]
    if params.run_all:
        argv += ["--run-all"]

    if params.resume:
        argv += ["--resume"]

    # stop-on-fail is default true in oraclepack; pass explicitly for clarity.
    argv += ["--stop-on-fail", "true" if params.stop_on_fail else "false"]

    argv += ["--roi-threshold", str(params.roi_threshold), "--roi-mode", params.roi_mode]

    argv += [str(pack_path)]

    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=params.timeout_s, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)


@mcp.tool(
    name="oraclepack_read_file",
    annotations={
        "title": "Read a file under allowed roots",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_read_file(params: ReadFileInput) -> Any:
    p = resolve_under_roots(Path(params.path), cfg.allowed_roots)
    if not p.exists() or not p.is_file():
        raise FileNotFoundError(f"file not found: {p}")

    max_bytes = params.max_bytes if params.max_bytes is not None else cfg.max_read_bytes
    text, truncated = safe_read_text(p, max_bytes=max_bytes)

    if params.response_format == ResponseFormat.json:
        return {"path": str(p), "truncated": truncated, "content": text}

    note = f"\n\n(content truncated to {max_bytes} bytes)" if truncated else ""
    return f"# {p.name}\n\n{_md_codeblock('text', text)}{note}"


@mcp.tool(
    name="oraclepack_taskify_detect_stage2",
    annotations={
        "title": "Detect Stage-2 oraclepack outputs",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_detect_stage2(params: TaskifyDetectStage2Input) -> Any:
    repo_root = resolve_under_roots(Path(params.repo_root), cfg.allowed_roots)
    res = detect_stage2(stage2_path=params.stage2_path, repo_root=repo_root)

    payload = {
        "kind": res.kind,
        "stage2_path": str(res.stage2_path),
        "out_dir": str(res.out_dir),
        "notes": res.notes,
    }

    if params.response_format == ResponseFormat.json:
        return payload

    lines = ["# Stage-2 resolution", ""]
    lines.append(f"- **kind**: `{payload['kind']}`")
    lines.append(f"- **stage2_path**: `{payload['stage2_path']}`")
    lines.append(f"- **out_dir**: `{payload['out_dir']}`")
    if payload["notes"]:
        lines.append("- **notes**:")
        lines += [f"  - {n}" for n in payload["notes"]]
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_validate_stage2",
    annotations={
        "title": "Validate Stage-2 directory outputs (01-*.md..20-*.md)",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_validate_stage2(params: TaskifyValidateStage2Input) -> Any:
    out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
    if not out_dir.is_dir():
        raise FileNotFoundError(f"out_dir is not a directory: {out_dir}")

    report = validate_stage2_dir(out_dir)
    if params.response_format == ResponseFormat.json:
        return report

    lines = [f"# Stage-2 validation: `{out_dir}`", ""]
    lines.append(f"- **ok**: {report['ok']}")
    if report["missing"]:
        lines.append("- **missing**:")
        lines += [f"  - {m}" for m in report["missing"]]
    if report["ambiguous"]:
        lines.append("- **ambiguous**:")
        for k, v in report["ambiguous"].items():
            lines.append(f"  - {k}: {v}")
    if report["selected"]:
        lines.append("- **selected**:")
        for k in sorted(report["selected"].keys()):
            lines.append(f"  - {k}: {report['selected'][k]}")
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_validate_action_pack",
    annotations={
        "title": "Validate Stage-3 Action Pack structure",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_validate_action_pack(params: TaskifyValidateActionPackInput) -> Any:
    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)
    if not pack_path.is_file():
        raise FileNotFoundError(f"pack not found: {pack_path}")

    report = validate_action_pack(pack_path)
    if params.response_format == ResponseFormat.json:
        return report

    lines = [f"# Action Pack validation: `{pack_path}`", ""]
    lines.append(f"- **ok**: {report['ok']}")
    lines.append(f"- **step_count**: {report['step_count']}")
    if report["errors"]:
        lines.append("- **errors**:")
        lines += [f"  - {e}" for e in report["errors"]]
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_artifacts_summary",
    annotations={
        "title": "Summarize expected Stage-3 artifacts",
        "readOnlyHint": True,
        "destructiveHint": False,
        "idempotentHint": True,
        "openWorldHint": False,
    },
)
async def oraclepack_taskify_artifacts_summary(params: TaskifyArtifactsSummaryInput) -> Any:
    out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)

    # Known artifacts from the skill
    expected = {
        "actions_json": out_dir / "_actions.json",
        "actions_md": out_dir / "_actions.md",
        "tm_complexity": out_dir / "tm-complexity.json",
        "prd": Path(cfg.workdir) / ".taskmaster" / "docs" / "oracle-actions-prd.md",
        "pipelines_doc": Path(cfg.workdir) / "docs" / "oracle-actions-pipelines.md",
        "taskmaster_tasks_local": Path(cfg.workdir) / "tasks.json",
        "taskmaster_tasks_dot": Path(cfg.workdir) / ".taskmaster" / "tasks.json",
    }

    found: dict[str, str] = {}
    missing: dict[str, str] = {}
    for k, p in expected.items():
        # Some paths may not be under allowed_roots; report as "not allowed" rather than crashing.
        try:
            p2 = resolve_under_roots(p, cfg.allowed_roots)
        except PathNotAllowedError:
            missing[k] = f"not allowed by ORACLEPACK_ALLOWED_ROOTS: {p}"
            continue

        if p2.exists():
            found[k] = str(p2)
        else:
            missing[k] = str(p2)

    payload = {"out_dir": str(out_dir), "found": found, "missing": missing}

    if params.response_format == ResponseFormat.json:
        return payload

    lines = [f"# Stage-3 artifacts summary: `{out_dir}`", ""]
    if found:
        lines.append("## found")
        for k, v in found.items():
            lines.append(f"- **{k}**: `{v}`")
    if missing:
        lines.append("\n## missing")
        for k, v in missing.items():
            lines.append(f"- **{k}**: `{v}`")
    return "\n".join(lines)


@mcp.tool(
    name="oraclepack_taskify_run_action_pack",
    annotations={
        "title": "Run a Stage-3 Action Pack via oraclepack",
        "readOnlyHint": False,
        "destructiveHint": True,
        "idempotentHint": False,
        "openWorldHint": True,
    },
)
async def oraclepack_taskify_run_action_pack(params: TaskifyRunActionPackInput) -> Any:
    _ensure_exec_enabled()

    pack_path = resolve_under_roots(Path(params.pack_path), cfg.allowed_roots)

    # Validate pack structure before executing
    v = validate_action_pack(pack_path)
    if not v["ok"]:
        raise ValueError(f"Refusing to execute invalid Action Pack: {json.dumps(v, indent=2)}")

    argv: list[str] = [cfg.oraclepack_bin]
    argv += ["--no-tui"]
    if params.out_dir:
        out_dir = resolve_under_roots(Path(params.out_dir), cfg.allowed_roots)
        argv += ["--out-dir", str(out_dir)]

    argv += ["run", "--yes", "--run-all", "--stop-on-fail", "true", str(pack_path)]

    result = await run_cmd(argv, cwd=cfg.workdir, timeout_s=params.timeout_s, env={}, character_limit=cfg.character_limit)
    return _format_cmd_result(result, params.response_format)
```

```python
# path: oraclepack-mcp-server/oraclepack_mcp_server/__main__.py
from __future__ import annotations

import argparse

from .server import mcp


def main() -> None:
    parser = argparse.ArgumentParser(description="oraclepack MCP server")
    parser.add_argument(
        "--transport",
        default="stdio",
        choices=["stdio", "streamable-http"],
        help="MCP transport (stdio or streamable-http)",
    )
    args = parser.parse_args()

    # The official Python SDK supports stdio and streamable-http transports.
    mcp.run(transport=args.transport)


if __name__ == "__main__":
    main()
```

