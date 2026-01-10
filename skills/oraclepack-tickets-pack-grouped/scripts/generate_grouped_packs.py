#!/usr/bin/env python3
from __future__ import annotations

import datetime as _dt
import math
import json
import re
import sys
import fnmatch
from pathlib import Path
from typing import Dict, Iterable, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "ticket", "tickets", "oraclepack", "oracle", "pack", "packs",
}


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
    # create a new run directory
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


def _normalize_title(text: str) -> str:
    text = text.strip().lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    text = re.sub(r"\s+", " ", text).strip()
    return text


def _read_gitignore(start: Path) -> Tuple[Path, List[str]]:
    cur = start.resolve()
    if cur.is_file():
        cur = cur.parent
    while True:
        p = cur / ".gitignore"
        if p.exists():
            lines: List[str] = []
            for ln in p.read_text(encoding="utf-8", errors="replace").splitlines():
                s = ln.strip()
                if not s or s.startswith("#"):
                    continue
                lines.append(s)
            return cur, lines
        if cur.parent == cur:
            return start.resolve(), []
        cur = cur.parent


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

    if "/" not in pattern:
        if dir_only:
            return any(part == pattern for part in rel_posix.split("/"))
        return fnmatch.fnmatch(name, pattern)

    target = rel_posix
    if anchored:
        if dir_only:
            return target.startswith(pattern + "/")
        return fnmatch.fnmatch(target, pattern)
    if dir_only:
        return f"/{pattern}/" in f"/{target}/"
    return fnmatch.fnmatch(target, f"**/{pattern}") or fnmatch.fnmatch(target, pattern)


def _is_gitignored(path: Path, git_root: Path, patterns: List[str]) -> bool:
    try:
        rel = path.resolve().relative_to(git_root)
    except Exception:
        rel = path
    rel_posix = rel.as_posix()
    parts = rel_posix.split("/")
    subpaths = []
    cur = ""
    for part in parts:
        cur = f"{cur}/{part}" if cur else part
        subpaths.append((cur, part))
    ignored = False
    for pat in patterns:
        neg = pat.startswith("!")
        for subpath, name in subpaths:
            if _gitignore_match(subpath, name, pat):
                ignored = not neg
    return ignored


def _read_heading(path: Path) -> str:
    try:
        for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
            if line.startswith("#"):
                return line.lstrip("#").strip()
    except FileNotFoundError:
        return ""
    return ""


def _collect_ticket_paths(ticket_root: str, ticket_glob: str, ticket_paths: str) -> List[Path]:
    if ticket_paths:
        parts = [p.strip() for p in ticket_paths.split(",") if p.strip()]
        return [Path(p) for p in parts]
    root = Path(ticket_root)
    if not root.exists():
        return []
    git_root, patterns = _read_gitignore(root)
    out: List[Path] = []
    for p in root.glob(ticket_glob):
        if _is_gitignored(Path(p), git_root, patterns):
            continue
        out.append(Path(p))
    return out


def _read_signature(path: Path, max_lines: int = 40) -> Tuple[str, str]:
    heading = ""
    lines: List[str] = []
    try:
        for line in path.read_text(encoding="utf-8", errors="replace").splitlines():
            if not heading and line.startswith("#"):
                heading = line.lstrip("#").strip()
            if line.strip():
                lines.append(line.strip())
            if len(lines) >= max_lines:
                break
    except FileNotFoundError:
        pass
    return heading, " ".join(lines)


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def _group_by_subdir(paths: Iterable[Path], ticket_root: str) -> Tuple[Dict[str, List[Path]], List[Path]]:
    root = Path(ticket_root)
    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    for p in paths:
        try:
            rel = p.relative_to(root)
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
        tokens.update(_tokenize(_read_heading(p)))
    return tokens


def _ticket_tokens(p: Path) -> set:
    toks = set(_tokenize(p.stem))
    heading, snippet = _read_signature(p)
    toks.update(_tokenize(heading))
    toks.update(_tokenize(snippet))
    return toks


def _signature_tokens(p: Path, body_chars: int) -> set:
    heading = _read_heading(p)
    body = _read_text(p)
    body = body[:body_chars]
    toks = set(_tokenize(p.stem))
    toks.update(_tokenize(heading))
    toks.update(_tokenize(body))
    return toks


def _jaccard(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    union = a.union(b)
    return float(len(inter)) / float(len(union))


def _overlap(a: set, b: set) -> float:
    if not a or not b:
        return 0.0
    inter = a.intersection(b)
    denom = min(len(a), len(b))
    if denom == 0:
        return 0.0
    return float(len(inter)) / float(denom)


def _clusters_from_edges(nodes: List[str], edges: Dict[str, List[str]]) -> List[List[str]]:
    seen = set()
    clusters: List[List[str]] = []
    for n in nodes:
        if n in seen:
            continue
        stack = [n]
        comp = []
        seen.add(n)
        while stack:
            cur = stack.pop()
            comp.append(cur)
            for nxt in edges.get(cur, []):
                if nxt not in seen:
                    seen.add(nxt)
                    stack.append(nxt)
        clusters.append(sorted(comp))
    return clusters


def _dedupe_clusters(
    paths: List[Path],
    body_chars: int,
    jaccard_hi: float,
    overlap_hi: float,
    overlap_lo: float,
    delta_min: float,
) -> Tuple[List[List[str]], Dict[str, str], Dict[str, Dict[str, object]], Dict[Tuple[str, str], Dict[str, float]]]:
    tokens: Dict[str, set] = {}
    sizes: Dict[str, int] = {}
    titles: Dict[str, str] = {}
    for p in paths:
        key = str(p)
        tokens[key] = _signature_tokens(p, body_chars)
        sizes[key] = len(_read_text(p))
        titles[key] = _normalize_title(_read_heading(p))

    nodes = sorted(tokens.keys())
    edges: Dict[str, List[str]] = {n: [] for n in nodes}
    pair_scores: Dict[Tuple[str, str], Dict[str, float]] = {}

    for i, a in enumerate(nodes):
        for b in nodes[i + 1 :]:
            jac = _jaccard(tokens[a], tokens[b])
            ov = _overlap(tokens[a], tokens[b])
            pair_scores[(a, b)] = {"jaccard": jac, "overlap": ov}
            if ov >= overlap_hi or (jac >= jaccard_hi and ov >= overlap_lo):
                edges[a].append(b)
                edges[b].append(a)

    clusters = _clusters_from_edges(nodes, edges)
    cluster_meta: Dict[str, Dict[str, object]] = {}
    dup_map: Dict[str, str] = {}

    for idx, members in enumerate(clusters, start=1):
        if len(members) == 1:
            continue
        # canonical: largest content length, then lexicographic
        canon = sorted(
            members,
            key=lambda m: (-sizes.get(m, 0), m),
        )[0]
        deltas: List[str] = []
        redundant: List[str] = []
        for m in members:
            if m == canon:
                continue
            unique = tokens[m] - tokens[canon]
            unique_ratio = float(len(unique)) / float(max(1, len(tokens[m])))
            heading_diff = titles.get(m, "") != titles.get(canon, "")
            if unique_ratio >= delta_min or heading_diff:
                deltas.append(m)
            else:
                redundant.append(m)
            dup_map[m] = canon

        cluster_meta[str(idx)] = {
            "canonical": canon,
            "members": members,
            "deltas": sorted(deltas),
            "redundant": sorted(redundant),
        }

    return clusters, dup_map, cluster_meta, pair_scores


def _infer_groups(
    groups: Dict[str, List[Path]],
    loose: List[Path],
    min_score: float,
) -> Dict[str, List[Path]]:
    if not groups:
        return {"root": list(loose)}

    group_tokens = {k: _group_tokens(k, v) for k, v in groups.items()}
    for p in loose:
        tokens = _ticket_tokens(p)
        best = None
        best_score = -1.0
        for name in sorted(group_tokens.keys()):
            score = _jaccard(tokens, group_tokens[name])
            if score > best_score:
                best_score = score
                best = name
        if best is not None and best_score >= min_score:
            groups.setdefault(best, []).append(p)
        else:
            groups.setdefault("misc", []).append(p)
    return groups


def _chunk(paths: List[Path], size: int) -> List[List[Path]]:
    if size <= 0:
        return [paths]
    return [paths[i : i + size] for i in range(0, len(paths), size)]


def _chunk_by_limits(
    paths: List[Path],
    max_files: int,
    max_chars: int,
) -> List[List[Path]]:
    if max_files <= 0 and max_chars <= 0:
        return [paths]
    chunks: List[List[Path]] = []
    cur: List[Path] = []
    cur_chars = 0
    for p in paths:
        size = len(_read_text(p))
        if cur:
            if (max_files > 0 and len(cur) >= max_files) or (
                max_chars > 0 and cur_chars + size > max_chars
            ):
                chunks.append(cur)
                cur = []
                cur_chars = 0
        cur.append(p)
        cur_chars += size
    if cur:
        chunks.append(cur)
    return chunks


def _render_template(template: str, mapping: Dict[str, str]) -> str:
    out = template
    for key, val in mapping.items():
        out = out.replace("{{" + key + "}}", val)
    unresolved = sorted(set(re.findall(r"\{\{([^}]+)\}\}", out)))
    if unresolved:
        raise ValueError(f"Unresolved template placeholders: {unresolved}")
    return out


def _write_merge_file(
    out_dir: Path,
    cluster_id: str,
    canonical: str,
    deltas: List[str],
    redundant: List[str],
    body_chars: int,
) -> Path:
    merge_dir = out_dir / "_ticket_merges"
    merge_dir.mkdir(parents=True, exist_ok=True)
    path = merge_dir / f"cluster-{int(cluster_id):04d}.md"

    def _cap(text: str) -> str:
        if len(text) <= body_chars:
            return text
        return text[:body_chars] + "\n[... truncated ...]\n"

    lines: List[str] = []
    lines.append(f"# Ticket Merge Cluster {cluster_id}")
    lines.append("")
    lines.append("## Canonical")
    lines.append(f"- path: {canonical}")
    lines.append("")
    lines.append(_cap(_read_text(Path(canonical))))
    lines.append("")

    members = deltas + redundant
    if members:
        lines.append("## Also reported in")
        for m in members:
            lines.append(f"- {m}")
        lines.append("")

    if deltas:
        lines.append("## Unique details from related tickets")
        for m in deltas:
            text = _read_text(Path(m))
            toks = _signature_tokens(Path(m), body_chars)
            canon_toks = _signature_tokens(Path(canonical), body_chars)
            unique = toks - canon_toks
            sel: List[str] = []
            for ln in text.splitlines():
                lnt = _tokenize(ln)
                if any(t in unique for t in lnt):
                    sel.append(ln)
                if len(sel) >= 60:
                    break
            lines.append(f"### {m}")
            if sel:
                lines.extend(sel)
            else:
                lines.append("(no unique lines detected within cap)")
            lines.append("")

    path.write_text("\n".join(lines), encoding="utf-8")
    return path


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: generate_grouped_packs.py key=value [key=value ...]")
            return 0

    args = _parse_kv_args(sys.argv[1:])
    codebase_name = args.get("codebase_name", "Unknown")
    out_dir = args.get("out_dir", f"docs/oracle-questions-{_now_slug()}")
    oracle_cmd = args.get("oracle_cmd", "oracle")
    oracle_flags = args.get("oracle_flags", "--files-report")
    extra_files = args.get("extra_files", "")
    ticket_root = args.get("ticket_root", ".tickets")
    ticket_glob = args.get("ticket_glob", "**/*.md")
    ticket_paths = args.get("ticket_paths", "")
    ticket_max_files = args.get("ticket_max_files", "25")
    group_mode = args.get("group_mode", "subdir+infer")
    group_min_score = float(args.get("group_min_score", "0.08"))
    group_max_files = int(args.get("group_max_files", "25"))
    group_max_chars = int(args.get("group_max_chars", "200000"))
    allow_overwrite = args.get("allow_overwrite", "false").lower() in ("1", "true", "yes")
    dedupe_mode = args.get("dedupe_mode", "report")
    dedupe_jaccard = float(args.get("dedupe_jaccard", "0.55"))
    dedupe_overlap_hi = float(args.get("dedupe_overlap_hi", "0.80"))
    dedupe_overlap_lo = float(args.get("dedupe_overlap_lo", "0.70"))
    dedupe_delta_min = float(args.get("dedupe_delta_min", "0.15"))
    dedupe_body_chars = int(args.get("dedupe_body_chars", "2000"))
    mode = args.get("mode", "tickets-grouped-direct")

    template_path = Path(__file__).resolve().parent.parent / "references" / "tickets-pack-template.md"
    template = template_path.read_text(encoding="utf-8")

    paths = _collect_ticket_paths(ticket_root, ticket_glob, ticket_paths)
    paths = sorted((str(p) for p in paths))
    paths = [Path(p) for p in paths]

    original_paths = list(paths)
    dup_map: Dict[str, str] = {}
    cluster_meta: Dict[str, Dict[str, object]] = {}
    dup_pairs: Dict[Tuple[str, str], Dict[str, float]] = {}
    if dedupe_mode != "off":
        _clusters, dup_map, cluster_meta, dup_pairs = _dedupe_clusters(
            paths,
            body_chars=dedupe_body_chars,
            jaccard_hi=dedupe_jaccard,
            overlap_hi=dedupe_overlap_hi,
            overlap_lo=dedupe_overlap_lo,
            delta_min=dedupe_delta_min,
        )

    # Build grouping base: canonical tickets + singletons
    canonical_set = {meta["canonical"] for meta in cluster_meta.values()}
    dup_set = set(dup_map.keys())
    base_paths: List[Path] = []
    for p in paths:
        sp = str(p)
        if sp in dup_set:
            continue
        base_paths.append(p)

    groups: Dict[str, List[Path]] = {}
    loose: List[Path] = []
    if "subdir" in group_mode:
        groups, loose = _group_by_subdir(base_paths, ticket_root)
    else:
        loose = list(base_paths)

    if "infer" in group_mode:
        groups = _infer_groups(groups, loose, group_min_score)
    else:
        groups.setdefault("misc", []).extend(loose)

    dedupe_plan: Dict[str, Dict[str, object]] = {}
    merge_files: Dict[str, str] = {}
    if cluster_meta:
        primary_to_group: Dict[str, str] = {}
        for gname in groups:
            for p in groups[gname]:
                primary_to_group[str(p)] = gname

        for cluster_id, meta in sorted(cluster_meta.items(), key=lambda x: int(x[0])):
            canonical = meta["canonical"]
            deltas = list(meta["deltas"])
            redundant = list(meta["redundant"])
            gname = primary_to_group.get(canonical, "misc")

            if dedupe_mode == "merge":
                merge_path = _write_merge_file(
                    Path(out_dir),
                    cluster_id=cluster_id,
                    canonical=canonical,
                    deltas=deltas,
                    redundant=redundant,
                    body_chars=dedupe_body_chars,
                )
                merge_files[canonical] = str(merge_path)
                # Replace canonical in group with merge file
                groups[gname] = [p for p in groups[gname] if str(p) != canonical]
                groups[gname].append(merge_path)
            else:
                # report/prune: append related tickets to canonical group
                keep = deltas if dedupe_mode == "prune" else deltas + redundant
                for p in keep:
                    groups.setdefault(gname, []).append(Path(p))

            dedupe_plan[cluster_id] = {
                "canonical": canonical,
                "group": gname,
                "deltas": sorted(deltas),
                "redundant": sorted(redundant),
                "mode": dedupe_mode,
            }

    # Ensure stable order
    for k in sorted(groups.keys()):
        groups[k] = sorted((str(p) for p in groups[k]))
        groups[k] = [Path(p) for p in groups[k]]

    original_set = {str(p) for p in original_paths}
    assignment: Dict[str, str] = {}
    for gname, gpaths in groups.items():
        for p in gpaths:
            sp = str(p)
            if sp in original_set:
                if sp in assignment:
                    raise SystemExit(f"[ERROR] Ticket assigned to multiple groups: {sp}")
                assignment[sp] = gname

    for meta in dedupe_plan.values():
        gname = meta["group"]
        for sp in [meta["canonical"]] + meta["deltas"] + meta["redundant"]:
            if sp not in assignment:
                assignment[sp] = gname

    missing = sorted(original_set - set(assignment.keys()))
    if missing:
        raise SystemExit(f"[ERROR] Tickets missing group assignment: {missing}")

    out_dir = _ensure_run_dir(out_dir, allow_overwrite)
    base_out = Path(out_dir)
    packs_dir = base_out / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    grouping_report: Dict[str, List[str]] = {}
    manifest_groups: List[Dict[str, object]] = []
    group_originals: Dict[str, List[str]] = {g: [] for g in groups.keys()}
    for ticket, gname in assignment.items():
        group_originals.setdefault(gname, []).append(ticket)
    for group_name in sorted(groups.keys()):
        group_paths = groups[group_name]
        grouping_report[group_name] = [str(p) for p in group_paths]

        parts = _chunk_by_limits(group_paths, group_max_files, group_max_chars)
        for idx, part in enumerate(parts, start=1):
            part_suffix = f"-part-{idx:02d}" if len(parts) > 1 else ""
            group_slug = _slugify(group_name + part_suffix)

            pack_out_dir = str(base_out / group_slug)
            pack_file = packs_dir / f"{group_slug}.md"

            mapping = {
                "codebase_name": codebase_name,
                "out_dir": pack_out_dir,
                "oracle_cmd": oracle_cmd,
                "oracle_flags": oracle_flags,
                "extra_files": extra_files,
                "ticket_root": ticket_root,
                "ticket_glob": ticket_glob,
                "ticket_paths": ",".join(str(p) for p in part),
                "ticket_max_files": str(min(len(part), max(1, group_max_files))),
                "group_name": group_name,
                "group_slug": group_slug,
                "mode": mode,
            }

            content = _render_template(template, mapping)
            pack_file.write_text(content, encoding="utf-8")

            manifest_groups.append(
                {
                    "group": group_name,
                    "slug": group_slug,
                    "part": idx,
                    "pack_path": str(pack_file),
                    "out_dir": pack_out_dir,
                    "attached_paths": [str(p) for p in part],
                    "original_tickets": sorted(group_originals.get(group_name, [])),
                }
            )

    (base_out / "_groups.json").write_text(
        json.dumps(grouping_report, indent=2, sort_keys=True),
        encoding="utf-8",
    )

    if dup_map:
        (base_out / "_duplicates.json").write_text(
            json.dumps(dup_map, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    if dedupe_plan:
        (base_out / "_dedupe_plan.json").write_text(
            json.dumps(dedupe_plan, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    if cluster_meta:
        pairs_out = [
            {"a": a, "b": b, **scores} for (a, b), scores in sorted(dup_pairs.items())
        ]
        (base_out / "_dupes_possible.json").write_text(
            json.dumps({"clusters": cluster_meta, "pairs": pairs_out}, indent=2, sort_keys=True),
            encoding="utf-8",
        )

    (base_out / "manifest.json").write_text(
        json.dumps({"groups": manifest_groups}, indent=2, sort_keys=True),
        encoding="utf-8",
    )

    print(f"[OK] wrote packs to: {packs_dir}")
    print(f"[OK] wrote grouping map: {base_out / '_groups.json'}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
