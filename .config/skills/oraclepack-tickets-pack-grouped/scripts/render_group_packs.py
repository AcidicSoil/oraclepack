#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Dict


def _render_template(template: str, mapping: Dict[str, str]) -> str:
    out = template
    for key, val in mapping.items():
        out = out.replace("{{" + key + "}}", val)
    unresolved = sorted(set(re.findall(r"\{\{([^}]+)\}\}", out)))
    if unresolved:
        raise ValueError(f"Unresolved template placeholders: {unresolved}")
    return out


def main() -> int:
    if len(sys.argv) == 1:
        print("Select how to run:")
        print("1) Use defaults (no args)")
        print("2) Provide custom args (show usage)")
        choice = input("Enter choice [1-2]: ").strip() or "1"
        if choice == "2":
            print("Usage: render_group_packs.py --manifest manifest.json --out-dir out")
            return 0

    p = argparse.ArgumentParser(description="Render group-specific bundle packs from manifest.")
    p.add_argument("--manifest", default="manifest.json")
    p.add_argument("--out-dir", default="docs/oracle-questions-sharded")
    p.add_argument("--template", default="/home/user/.codex/skills/oraclepack-tickets-pack-grouped/references/tickets-pack-template-bundle.md")
    p.add_argument("--codebase-name", default="Unknown")
    p.add_argument("--oracle-cmd", default="oracle")
    p.add_argument("--oracle-flags", default="--files-report")
    p.add_argument("--extra-files", default="")
    p.add_argument("--ticket-root", default=".tickets")
    p.add_argument("--ticket-glob", default="**/*.md")
    p.add_argument("--mode", default="tickets-bundle")
    args = p.parse_args()

    manifest_path = Path(args.manifest)
    if not manifest_path.exists():
        raise SystemExit(f"[ERROR] manifest not found: {manifest_path}")

    manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    template = Path(args.template).read_text(encoding="utf-8")

    out_dir = Path(args.out_dir)
    packs_dir = out_dir / "packs"
    packs_dir.mkdir(parents=True, exist_ok=True)

    for group in manifest.get("groups", []):
        slug = group["slug"]
        tickets = group["tickets"]
        pack_dir = packs_dir / slug
        pack_dir.mkdir(parents=True, exist_ok=True)

        pack_path = pack_dir / f"oracle-pack_{slug}.md"
        bundle_path = pack_dir / f"tickets_bundle_{slug}.md"
        out_run_dir = pack_dir / "out"

        mapping = {
            "codebase_name": args.codebase_name,
            "out_dir": str(out_run_dir),
            "oracle_cmd": args.oracle_cmd,
            "oracle_flags": args.oracle_flags,
            "extra_files": args.extra_files,
            "ticket_root": args.ticket_root,
            "ticket_glob": args.ticket_glob,
            "ticket_paths": ",".join(tickets),
            "ticket_bundle_path": str(bundle_path),
            "mode": args.mode,
        }

        content = _render_template(template, mapping)
        pack_path.write_text(content, encoding="utf-8")
        group["pack_path"] = str(pack_path)

    manifest_path.write_text(json.dumps(manifest, indent=2, sort_keys=True), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
