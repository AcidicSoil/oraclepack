#!/usr/bin/env python3
from __future__ import annotations

import datetime as _dt
import fnmatch
import json
import re
import sys
from pathlib import Path
from typing import Dict, Iterable, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "repo", "repos", "code", "codebase", "project", "oraclepack", "oracle", "pack", "packs",
}

DEFAULT_IGNORE_DIRS = {
    ".git",
    ".hg",
    ".svn",
    "node_modules",
    "dist",
    "build",
    ".next",
    ".venv",
    "venv",
    "coverage",
    "target",
}

DEFAULT_INCLUDE_EXTS = {
    ".py", ".ts", ".tsx", ".js", ".jsx", ".go", ".rs", ".java", ".kt", ".cpp", ".c",
    ".h", ".hpp", ".cs", ".rb", ".php", ".swift", ".scala", ".sql", ".md", ".yaml",
    ".yml", ".json", ".toml", ".ini", ".sh", ".ps1", ".tf", ".proto",
}

def _read_gitignore(root: Path) -> List[str]:
    path = root / ".gitignore"
    if not path.exists():
        return []
    lines: List[str] = []
    for ln in path.read_text(encoding="utf-8", errors="replace").splitlines():
        s = ln.strip()
        if not s or s.startswith("#"):
            continue
        lines.append(s)
    return lines


def _gitignore_match(rel_posix: str, name: str, pattern: str) -> bool:
    neg = pattern.startswith("!")
    if neg:
        pattern = pattern[1:]
    if not pattern:
        return False

    anchored = pattern.startswith("/")
    if anchored:
        pattern = pattern.lstrip("/")

    dir_only = pattern.endswith("/")
    if dir_only:
        pattern = pattern.rstrip("/")

    # Match basename-only patterns
    if "/" not in pattern:
        if dir_only:
            # Any parent dir name match
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    # Path patterns (allow **)
    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    # Unanchored path pattern: match anywhere
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def _is_gitignored(rel_posix: str, patterns: List[str]) -> bool:
    ignored = False
    parts = rel_posix.split("/")
    subpaths = []
    cur = ""
    for part in parts:
        cur = f"{cur}/{part}" if cur else part
        subpaths.append((cur, part))
    # Evaluate patterns in order for each subpath and full path
    for pat in patterns:
        neg = pat.startswith("!")
        for subpath, name in subpaths:
            if _gitignore_match(subpath, name, pat):
                ignored = not neg
    return ignored


def _parse_kv_args(argv: List[str]) -> Dict[str, str]:
    args: Dict[str, str] = {}
    for raw in argv:
        if "=" not in raw:
            continue
        k, v = raw.split("=", 1)
        args[k.strip()] = v.strip()
    return args


def _today() -> str:
    return _dt.date.today().isoformat()


def _now_slug() -> str:
    return _dt.datetime.now().strftime("%Y-%m-%d-%H%M%S")


def _slugify(s: str) -> str:
    s = s.strip().lower()
    s = re.sub(r"[^a-z0-9]+", "-", s)
    s = re.sub(r"-+", "-", s).strip("-")
    return s or "group"


def _ensure_run_dir(out_dir: str, allow_overwrite: bool) -> str:
    base = Path(out_dir)
    if not base.exists():
        base.mkdir(parents=True, exist_ok=True)
        return str(base)
    if allow_overwrite:
        return str(base)
    ts = _now_slug()
    for i in range(0, 1000):
        suffix = f"-run-{ts}" if i == 0 else f"-run-{ts}-{i:02d}"
        cand = Path(f"{out_dir}{suffix}")
        if not cand.exists():
            cand.mkdir(parents=True, exist_ok=True)
            return str(cand)
    raise SystemExit("[ERROR] could not allocate unique run directory")


def _tokenize(text: str) -> List[str]:
    text = text.lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    toks = [t for t in text.split() if len(t) >= 3 and t not in STOPWORDS]
    return toks


def _group_by_subdir(paths: Iterable[Path], code_root: str) -> Tuple[Dict[str, List[Path]], List[Path]]:
    root = Path(code_root).resolve()
    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    for p in paths:
        try:
            rel = p.resolve().relative_to(root)
        except ValueError:
            loose.append(p)
            continue
        if len(rel.parts) >= 2:
            key = rel.parts[0]
            groups.setdefault(key, []).append(p)
        else:
            loose.append(p)
    return groups, loose


def _group_tokens(group_name: str, paths: Iterable[Path]) -> set:
    tokens = set(_tokenize(group_name))
    for p in paths:
        tokens.update(_tokenize(p.stem))
    return tokens


def _file_tokens(p: Path) -> set:
    toks = set(_tokenize(p.stem))
    toks.update(_tokenize(str(p.parent.name)))
    return toks


def _jaccard(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    union = a.union(b)
    return float(len(inter)) / float(len(union))


def _collect_paths(
    code_root: str,
    code_glob: str,
    code_paths: str,
    include_exts: str,
    exclude_glob: str,
    ignore_dirs: str,
) -> List[Path]:
    if code_paths:
        parts = [p.strip() for p in code_paths.split(",") if p.strip()]
        return [Path(p) for p in parts]

    root = Path(code_root)
    if not root.exists():
        return []

    gitignore_patterns = _read_gitignore(root)
    ignore = {p.strip() for p in ignore_dirs.split(",") if p.strip()}
    ignore = ignore.union(DEFAULT_IGNORE_DIRS)

    include = {e.strip().lower() for e in include_exts.split(",") if e.strip()}
    if not include_exts.strip():
        include = set(DEFAULT_INCLUDE_EXTS)

    excludes = [g.strip() for g in exclude_glob.split(",") if g.strip()]

    out: List[Path] = []
    for p in root.glob(code_glob):
        if p.is_dir():
            continue
        parts = set(p.parts)
        rel = p.relative_to(root)
        rel_posix = rel.as_posix()
        if _is_gitignored(rel_posix, gitignore_patterns):
            continue
        if parts.intersection(ignore):
            continue
        if excludes and any(fnmatch.fnmatch(str(p), g) for g in excludes):
            continue
        if include:
            ext = p.suffix.lower()
            if ext not in include:
                continue
        out.append(p)

    return out


def _cap_group(paths: List[Path], max_files: int, max_chars: int) -> List[List[Path]]:
    chunks: List[List[Path]] = []
    current: List[Path] = []
    size = 0

    for p in paths:
        p_size = 0
        try:
            p_size = p.stat().st_size
        except FileNotFoundError:
            p_size = 0

        if max_files and len(current) >= max_files:
            chunks.append(current)
            current = []
            size = 0

        if max_chars and current and size + p_size > max_chars:
            chunks.append(current)
            current = []
            size = 0

        current.append(p)
        size += p_size

    if current:
        chunks.append(current)

    return chunks


def main() -> None:
    args = _parse_kv_args(sys.argv[1:])

    codebase_name = args.get("codebase_name", "Unknown")
    out_dir = args.get("out_dir", f"docs/oracle-questions-{_now_slug()}")
    oracle_cmd = args.get("oracle_cmd", "oracle")
    oracle_flags = args.get("oracle_flags", "--files-report")
    extra_files = args.get("extra_files", "")
    code_root = args.get("code_root", ".")
    code_glob = args.get("code_glob", "**/*")
    code_paths = args.get("code_paths", "")
    code_max_files = int(args.get("code_max_files", "200"))
    group_mode = args.get("group_mode", "subdir+infer")
    group_min_score = float(args.get("group_min_score", "0.10"))
    group_max_files = int(args.get("group_max_files", "200"))
    group_max_chars = int(args.get("group_max_chars", "200000"))
    ignore_dirs = args.get("ignore_dirs", "")
    include_exts = args.get("include_exts", "")
    exclude_glob = args.get("exclude_glob", "")
    mode = args.get("mode", "codebase-grouped-direct")
    allow_overwrite = args.get("allow_overwrite", "false").lower() in ("1", "true", "yes")

    template_path = Path(__file__).resolve().parents[1] / "references" / "codebase-pack-template.md"
    if not template_path.exists():
        raise SystemExit(f"[ERROR] Template not found: {template_path}")

    paths = _collect_paths(code_root, code_glob, code_paths, include_exts, exclude_glob, ignore_dirs)
    paths = sorted(paths, key=lambda p: str(p))
    if code_max_files and code_max_files > 0:
        paths = paths[:code_max_files]

    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []

    if "subdir" in group_mode:
        groups, loose = _group_by_subdir(paths, code_root)
    else:
        loose = list(paths)

    if "infer" in group_mode and loose:
        group_tokens = {k: _group_tokens(k, v) for k, v in groups.items()}
        for p in loose:
            best = None
            best_score = 0.0
            pt = _file_tokens(p)
            for g, gt in group_tokens.items():
                score = _jaccard(pt, gt)
                if score > best_score:
                    best_score = score
                    best = g
            if best and best_score >= group_min_score:
                groups.setdefault(best, []).append(p)
            else:
                groups.setdefault("root", []).append(p)
    else:
        if loose:
            groups.setdefault("root", []).extend(loose)

    if not groups:
        groups = {"root": []}

    out_dir = _ensure_run_dir(out_dir, allow_overwrite)
    out_dir_path = Path(out_dir)
    packs_dir = out_dir_path / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    rendered_groups: Dict[str, List[str]] = {}

    template = template_path.read_text(encoding="utf-8")
    for group_name in sorted(groups.keys()):
        files = sorted(groups[group_name], key=lambda p: str(p))
        chunks = _cap_group(files, group_max_files, group_max_chars)
        for idx, chunk in enumerate(chunks, start=1):
            part_suffix = f" part {idx}" if len(chunks) > 1 else ""
            full_group_name = f"{group_name}{part_suffix}"
            group_slug = _slugify(full_group_name)
            pack_path = packs_dir / f"{group_slug}.md"

            rendered = template
            rendered = rendered.replace("{{codebase_name}}", codebase_name)
            rendered = rendered.replace("{{out_dir}}", str(out_dir))
            rendered = rendered.replace("{{oracle_cmd}}", oracle_cmd)
            rendered = rendered.replace("{{oracle_flags}}", oracle_flags)
            rendered = rendered.replace("{{extra_files}}", extra_files)
            rendered = rendered.replace("{{code_root}}", code_root)
            rendered = rendered.replace("{{code_glob}}", code_glob)
            rendered = rendered.replace("{{code_paths}}", code_paths)
            rendered = rendered.replace("{{code_max_files}}", str(code_max_files))
            rendered = rendered.replace("{{group_name}}", full_group_name)
            rendered = rendered.replace("{{group_slug}}", group_slug)
            rendered = rendered.replace("{{group_mode}}", group_mode)
            rendered = rendered.replace("{{group_min_score}}", str(group_min_score))
            rendered = rendered.replace("{{group_max_files}}", str(group_max_files))
            rendered = rendered.replace("{{group_max_chars}}", str(group_max_chars))
            rendered = rendered.replace("{{ignore_dirs}}", ignore_dirs)
            rendered = rendered.replace("{{include_exts}}", include_exts)
            rendered = rendered.replace("{{exclude_glob}}", exclude_glob)
            rendered = rendered.replace("{{mode}}", mode)
            rendered = rendered.replace(
                "{{group_files_json}}",
                json.dumps([str(p) for p in chunk], indent=2),
            )
            unresolved = sorted(set(re.findall(r"\{\{([^}]+)\}\}", rendered)))
            if unresolved:
                raise SystemExit(f"[ERROR] Unresolved template placeholders: {unresolved}")

            pack_path.write_text(rendered, encoding="utf-8")
            rendered_groups.setdefault(full_group_name, []).append(str(pack_path))

    groups_json = {
        "code_root": code_root,
        "groups": {k: [str(p) for p in v] for k, v in groups.items()},
        "packs": rendered_groups,
    }
    (out_dir_path / "_groups.json").write_text(json.dumps(groups_json, indent=2), encoding="utf-8")
    manifest = {
        "groups": [
            {
                "group": g,
                "packs": rendered_groups.get(g, []),
                "files": [str(p) for p in groups.get(g, [])],
            }
            for g in sorted(groups.keys())
        ]
    }
    (out_dir_path / "manifest.json").write_text(json.dumps(manifest, indent=2), encoding="utf-8")


if __name__ == "__main__":
    main()
