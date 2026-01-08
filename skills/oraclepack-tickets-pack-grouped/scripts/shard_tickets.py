#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import math
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple

STOPWORDS = {
    "the", "and", "for", "with", "from", "this", "that", "into", "over", "under", "when",
    "then", "than", "else", "only", "must", "should", "could", "would", "will", "shall",
    "ticket", "tickets", "oraclepack", "oracle", "pack", "packs",
}

SECTION_KEYS = {"summary", "acceptance", "criteria", "background", "context"}


@dataclass
class Ticket:
    path: Path
    text: str
    tokens: List[str]
    vector: List[float]


def _tokenize(text: str) -> List[str]:
    text = text.lower()
    text = re.sub(r"[^a-z0-9]+", " ", text)
    return [t for t in text.split() if len(t) >= 3 and t not in STOPWORDS]


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def _extract_repr(text: str, stem: str, max_chars: int) -> str:
    lines = text.splitlines()
    heading = ""
    sections: List[str] = []
    capture = False
    for line in lines:
        s = line.strip()
        if not heading and s.startswith("#"):
            heading = s.lstrip("#").strip()
        if s.startswith("#"):
            key = s.lstrip("#").strip().lower()
            capture = any(k in key for k in SECTION_KEYS)
            continue
        if capture and s:
            sections.append(s)
        if len(" ".join(sections)) >= max_chars:
            break
    body = " ".join(sections)
    base = " ".join([stem, heading, body])
    return base[:max_chars]


def _tfidf_vectors(texts: List[str]) -> Tuple[List[List[float]], List[str]]:
    docs = [
        [tok for tok in _tokenize(t)]
        for t in texts
    ]
    vocab: Dict[str, int] = {}
    df: Dict[str, int] = {}
    for toks in docs:
        seen = set()
        for tok in toks:
            if tok not in vocab:
                vocab[tok] = len(vocab)
            if tok not in seen:
                df[tok] = df.get(tok, 0) + 1
                seen.add(tok)

    n_docs = len(docs)
    idf = [0.0] * len(vocab)
    for tok, idx in vocab.items():
        idf[idx] = math.log((1 + n_docs) / (1 + df.get(tok, 1))) + 1.0

    vectors: List[List[float]] = []
    for toks in docs:
        tf: Dict[int, float] = {}
        for tok in toks:
            tf[vocab[tok]] = tf.get(vocab[tok], 0.0) + 1.0
        vec = [0.0] * len(vocab)
        for idx, count in tf.items():
            vec[idx] = count * idf[idx]
        # L2 normalize
        norm = math.sqrt(sum(v * v for v in vec)) or 1.0
        vec = [v / norm for v in vec]
        vectors.append(vec)

    inv_vocab = [None] * len(vocab)
    for tok, idx in vocab.items():
        inv_vocab[idx] = tok
    return vectors, inv_vocab


def _cosine(a: List[float], b: List[float]) -> float:
    return sum(x * y for x, y in zip(a, b))


def _centroid(vectors: List[List[float]]) -> List[float]:
    if not vectors:
        return []
    dim = len(vectors[0])
    out = [0.0] * dim
    for v in vectors:
        for i, val in enumerate(v):
            out[i] += val
    n = float(len(vectors)) or 1.0
    out = [v / n for v in out]
    norm = math.sqrt(sum(v * v for v in out)) or 1.0
    return [v / norm for v in out]


def _kmeans_split(vectors: List[List[float]], k: int, iters: int = 10) -> List[List[int]]:
    if k <= 1:
        return [list(range(len(vectors)))]
    # deterministic init: first k vectors
    centroids = [vectors[i][:] for i in range(k)]
    for _ in range(iters):
        clusters = [[] for _ in range(k)]
        for idx, v in enumerate(vectors):
            best = 0
            best_score = -1.0
            for c_idx, c in enumerate(centroids):
                score = _cosine(v, c)
                if score > best_score:
                    best_score = score
                    best = c_idx
            clusters[best].append(idx)
        new_centroids = []
        for cluster in clusters:
            if cluster:
                new_centroids.append(_centroid([vectors[i] for i in cluster]))
            else:
                new_centroids.append(centroids[len(new_centroids)])
        centroids = new_centroids
    return clusters


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: shard_tickets.py --ticket-root .tickets --out-dir out")
            return 0

    p = argparse.ArgumentParser(description="Shard tickets into topic/domain groups.")
    p.add_argument("--ticket-root", default=".tickets")
    p.add_argument("--ticket-glob", default="**/*.md")
    p.add_argument("--ticket-paths", default="")
    p.add_argument("--out-dir", default="docs/oracle-questions-sharded")
    p.add_argument("--min-sim", type=float, default=0.15)
    p.add_argument("--max-group-size", type=int, default=25)
    p.add_argument("--min-group-size", type=int, default=1)
    p.add_argument("--max-bundle-chars", type=int, default=200000)
    p.add_argument("--repr-chars", type=int, default=2000)
    p.add_argument("--use-llm-for-ambiguous", action="store_true")
    args = p.parse_args()

    ticket_root = Path(args.ticket_root)
    if args.ticket_paths:
        paths = [Path(p.strip()) for p in args.ticket_paths.split(",") if p.strip()]
    else:
        paths = sorted(ticket_root.glob(args.ticket_glob), key=lambda p: str(p)) if ticket_root.exists() else []

    texts: List[str] = []
    tickets: List[Ticket] = []
    for pth in paths:
        txt = _read_text(pth)
        rep = _extract_repr(txt, pth.stem, args.repr_chars)
        texts.append(rep)

    vectors, vocab = _tfidf_vectors(texts)
    for pth, txt, vec in zip(paths, texts, vectors):
        tickets.append(Ticket(path=pth, text=txt, tokens=_tokenize(txt), vector=vec))

    groups: Dict[str, List[int]] = {}
    loose: List[int] = []
    for idx, t in enumerate(tickets):
        try:
            rel = t.path.relative_to(ticket_root)
        except ValueError:
            loose.append(idx)
            continue
        if len(rel.parts) >= 2:
            g = rel.parts[0]
            groups.setdefault(g, []).append(idx)
        else:
            loose.append(idx)

    # Compute centroids for subdir groups
    centroids: Dict[str, List[float]] = {}
    for g, idxs in groups.items():
        centroids[g] = _centroid([tickets[i].vector for i in idxs])

    # Assign loose tickets by similarity
    reasons: Dict[int, Dict[str, object]] = {}
    for idx in loose:
        best_g = None
        best_sim = -1.0
        for g, c in centroids.items():
            sim = _cosine(tickets[idx].vector, c)
            if sim > best_sim:
                best_sim = sim
                best_g = g
        if best_g is not None and best_sim >= args.min_sim:
            groups.setdefault(best_g, []).append(idx)
            reasons[idx] = {"assigned_to": best_g, "sim": best_sim, "reason": "tfidf"}
        else:
            groups.setdefault("misc", []).append(idx)
            reasons[idx] = {
                "assigned_to": "misc",
                "sim": best_sim,
                "reason": "ambiguous" if not args.use_llm_for_ambiguous else "ambiguous_llm_needed",
            }

    # Merge small groups
    if args.min_group_size > 1 and len(groups) > 1:
        for g in sorted(list(groups.keys())):
            if g == "misc":
                continue
            if len(groups[g]) < args.min_group_size:
                # merge into nearest group
                g_centroid = _centroid([tickets[i].vector for i in groups[g]])
                best_g = None
                best_sim = -1.0
                for og, c in centroids.items():
                    if og == g:
                        continue
                    sim = _cosine(g_centroid, c)
                    if sim > best_sim:
                        best_sim = sim
                        best_g = og
                if best_g:
                    groups.setdefault(best_g, []).extend(groups[g])
                    del groups[g]

    # Split large groups using deterministic kmeans
    final_groups: Dict[str, List[int]] = {}
    for g in sorted(groups.keys()):
        idxs = groups[g]
        idxs_sorted = sorted(idxs, key=lambda i: str(tickets[i].path))
        total_chars = sum(len(_read_text(tickets[i].path)) for i in idxs_sorted)
        k = max(
            1,
            math.ceil(len(idxs_sorted) / max(1, args.max_group_size)),
            math.ceil(total_chars / max(1, args.max_bundle_chars)),
        )
        if k <= 1:
            final_groups[g] = idxs_sorted
            continue
        clusters = _kmeans_split([tickets[i].vector for i in idxs_sorted], k)
        part = 1
        for cluster in clusters:
            if not cluster:
                continue
            slug = f"{g}-part-{part:02d}"
            final_groups[slug] = [idxs_sorted[i] for i in cluster]
            part += 1

    # Build manifest
    out_dir = Path(args.out_dir)
    out_dir.mkdir(parents=True, exist_ok=True)
    manifest_groups = []
    for g in sorted(final_groups.keys()):
        idxs = final_groups[g]
        vecs = [tickets[i].vector for i in idxs]
        centroid = _centroid(vecs)
        top_terms = []
        for i, score in sorted(enumerate(centroid), key=lambda x: -x[1])[:8]:
            if score <= 0:
                continue
            top_terms.append(vocab[i])
        sims = []
        for i in idxs:
            sims.append(_cosine(tickets[i].vector, centroid))
        conf = sum(sims) / float(len(sims)) if sims else 0.0

        manifest_groups.append(
            {
                "slug": g,
                "tickets": [str(tickets[i].path) for i in idxs],
                "keywords": top_terms,
                "confidence": conf,
            }
        )

    manifest = {"groups": manifest_groups}
    (out_dir / "manifest.json").write_text(json.dumps(manifest, indent=2, sort_keys=True), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
