#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import subprocess
from pathlib import Path
from typing import Dict


def _read_text(path: Path) -> str:
    try:
        return path.read_text(encoding="utf-8", errors="replace")
    except FileNotFoundError:
        return ""


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: validate_shards.py --manifest manifest.json")
            return 0

    p = argparse.ArgumentParser(description="Validate sharded packs manifest.")
    p.add_argument("--manifest", default="manifest.json")
    p.add_argument("--max-bundle-chars", type=int, default=200000)
    p.add_argument(
        "--validator",
        default="/home/user/.codex/skills/oraclepack-tickets-pack-common/scripts/validate_pack.py",
    )
    args = p.parse_args()

    manifest_path = Path(args.manifest)
    if not manifest_path.exists():
        raise SystemExit(f"[ERROR] manifest not found: {manifest_path}")

    manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    counts: Dict[str, int] = {}

    for group in manifest.get("groups", []):
        for t in group.get("tickets", []):
            counts[t] = counts.get(t, 0) + 1

    bad = [t for t, c in counts.items() if c != 1]
    if bad:
        raise SystemExit(f"[ERROR] Tickets assigned !=1 times: {bad}")

    for group in manifest.get("groups", []):
        pack_path = Path(group.get("pack_path", ""))
        if not pack_path.exists():
            raise SystemExit(f"[ERROR] pack missing: {pack_path}")

        # validate pack
        subprocess.run(
            [
                "python3",
                args.validator,
                "--mode",
                "bundle",
                str(pack_path),
            ],
            check=True,
        )

        # size check
        total = 0
        for t in group.get("tickets", []):
            total += len(_read_text(Path(t)))
        if total > args.max_bundle_chars:
            raise SystemExit(
                f"[ERROR] group '{group.get('slug')}' exceeds max bundle chars: {total} > {args.max_bundle_chars}"
            )

    print("[OK] Sharded packs manifest validated.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
