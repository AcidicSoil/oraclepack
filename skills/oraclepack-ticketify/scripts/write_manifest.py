#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
from pathlib import Path


def main() -> int:
    p = argparse.ArgumentParser(description="Write a minimal manifest for ticketify outputs.")
    p.add_argument("--out-dir", required=True)
    p.add_argument("--pack-path", required=True)
    p.add_argument("--tickets-index", required=True)
    p.add_argument("--actions-json", required=True)
    p.add_argument("--actions-md", required=True)
    p.add_argument("--prd-path", required=True)
    p.add_argument("--pipelines", default="")
    p.add_argument("--autopilot", default="")
    p.add_argument("--manifest-path", default="")
    args = p.parse_args()

    out_dir = Path(args.out_dir)
    manifest_path = Path(args.manifest_path) if args.manifest_path else out_dir / "manifest.json"

    manifest = {
        "out_dir": str(out_dir),
        "pack_path": args.pack_path,
        "tickets_index": args.tickets_index,
        "actions_json": args.actions_json,
        "actions_md": args.actions_md,
        "prd_path": args.prd_path,
    }
    if args.pipelines:
        manifest["pipelines"] = args.pipelines
    if args.autopilot:
        manifest["autopilot"] = args.autopilot

    manifest_path.parent.mkdir(parents=True, exist_ok=True)
    manifest_path.write_text(json.dumps(manifest, indent=2, sort_keys=True), encoding="utf-8")
    print(f"[OK] wrote manifest: {manifest_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
