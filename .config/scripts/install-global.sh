# path: scripts/install-global.sh
#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'USAGE'
install-global.sh — rebuild oraclepack and install it to a PATH directory.

Usage:
  ./scripts/install-global.sh [--prefix DIR] [--sudo] [--skip-tests]

Options:
  --prefix DIR     Install directory (default: ~/.local/bin; macOS prefers /usr/local/bin if writable)
  --sudo           Use sudo when copying into --prefix (for system dirs like /usr/local/bin)
  --skip-tests     Skip `go test ./...`

Examples:
  ./scripts/install-global.sh
  ./scripts/install-global.sh --prefix "$HOME/.local/bin"
  ./scripts/install-global.sh --prefix /usr/local/bin --sudo
USAGE
}

PREFIX=""
USE_SUDO=0
SKIP_TESTS=0

while [[ $# -gt 0 ]]; do
  case "$1" in
    --prefix) PREFIX="${2:-}"; shift 2 ;;
    --sudo) USE_SUDO=1; shift ;;
    --skip-tests) SKIP_TESTS=1; shift ;;
    -h|--help) usage; exit 0 ;;
    *) echo "Unknown arg: $1" >&2; usage; exit 2 ;;
  esac
done

# Repo root (so script works from anywhere)
REPO_ROOT="$(git rev-parse --show-toplevel 2>/dev/null || true)"
if [[ -z "${REPO_ROOT}" ]]; then
  echo "Error: not inside a git repo (expected oraclepack repo)." >&2
  exit 1
fi
cd "$REPO_ROOT"

# Build output
BUILD_DIR="$REPO_ROOT/.build"
mkdir -p "$BUILD_DIR"
OUT_BIN="$BUILD_DIR/oraclepack"

# Default install prefix
if [[ -z "${PREFIX}" ]]; then
  case "$(uname -s)" in
    Darwin)
      if [[ -w "/usr/local/bin" ]]; then
        PREFIX="/usr/local/bin"
      else
        PREFIX="$HOME/.local/bin"
      fi
      ;;
    *)
      PREFIX="$HOME/.local/bin"
      ;;
  esac
fi

echo "==> Repo: $REPO_ROOT"
echo "==> Building: $OUT_BIN"
echo "==> Installing to: $PREFIX"

command -v go >/dev/null 2>&1 || { echo "Error: go not found in PATH." >&2; exit 1; }

if [[ "$SKIP_TESTS" -eq 0 ]]; then
  echo "==> go test ./..."
  go test ./...
fi

# Match README build target: ./cmd/oraclepack :contentReference[oaicite:0]{index=0}
go build -o "$OUT_BIN" ./cmd/oraclepack

mkdir -p "$PREFIX"

# Install (overwrite existing)
if [[ "$USE_SUDO" -eq 1 ]]; then
  sudo install -m 0755 "$OUT_BIN" "$PREFIX/oraclepack"
else
  install -m 0755 "$OUT_BIN" "$PREFIX/oraclepack"
fi

echo "==> Installed: $PREFIX/oraclepack"
echo "==> Version (if supported):"
"$PREFIX/oraclepack" --version 2>/dev/null || true

# Best-effort: refresh shell hash table (won't error if unsupported)
hash -r 2>/dev/null || true
