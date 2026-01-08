from pathlib import Path
import runpy

COMMON = Path(__file__).resolve().parents[2] / "oraclepack-tickets-pack-common" / "scripts" / "validate_pack.py"
if not COMMON.exists():
    raise SystemExit(f"[ERROR] Shared validator not found: {COMMON}")

runpy.run_path(str(COMMON), run_name="__main__")
