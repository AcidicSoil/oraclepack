# path: scripts/build_install_oraclepack.sh
#!/usr/bin/env bash
set -euo pipefail

# Builds oraclepack for:
# - WSL/Linux (oraclepack)
# - Windows amd64 (oraclepack.exe)
# Then installs:
# - WSL binary -> ~/.local/bin/
# - Windows exe -> /mnt/c/Users/<winuser>/.local/bin/
# And writes a Git Bash wrapper:
# - /mnt/c/Users/<winuser>/bin/oraclepack  (calls WSL binary via wsl.exe)

usage() {
  cat <<'USAGE'
Usage:
  scripts/build_install_oraclepack.sh [options]

Options:
  --repo-root <path>      Repo root (default: git toplevel, else current dir)
  --win-user <name>       Windows username (default: auto-detect, else "user")
  --wsl-distro <name>     WSL distro name for wsl.exe -d (default: $WSL_DISTRO_NAME or "Ubuntu-24.04")
  --wsl-user <name>       WSL username for wsl.exe -u (default: current user)
  --no-linux              Skip building/installing Linux binary
  --no-windows            Skip building/installing Windows .exe
  --no-wrapper            Skip writing the Git Bash wrapper script
  --cgo <0|1>             Set CGO_ENABLED (default: leave unchanged)

Examples:
  scripts/build_install_oraclepack.sh
  scripts/build_install_oraclepack.sh --win-user Alice --wsl-distro Ubuntu-24.04
USAGE
}

die() { echo "error: $*" >&2; exit 1; }

have() { command -v "$1" >/dev/null 2>&1; }

detect_repo_root() {
  if have git && git rev-parse --show-toplevel >/dev/null 2>&1; then
    git rev-parse --show-toplevel
  else
    pwd
  fi
}

detect_win_user() {
  # Best-effort: ask Windows for %USERNAME% via cmd.exe (works in most WSL setups).
  if have cmd.exe; then
    local u
    u="$(cmd.exe /c "echo %USERNAME%" 2>/dev/null | tr -d '\r' | tail -n 1 || true)"
    if [[ -n "${u:-}" && "${u:-}" != "%USERNAME%" ]]; then
      echo "$u"
      return
    fi
  fi
  echo "user"
}

REPO_ROOT=""
WIN_USER=""
WSL_DISTRO="${WSL_DISTRO_NAME:-Ubuntu-24.04}"
WSL_USER="$(whoami)"
DO_LINUX=1
DO_WINDOWS=1
DO_WRAPPER=1
CGO=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo-root)  REPO_ROOT="${2:-}"; shift 2;;
    --win-user)   WIN_USER="${2:-}"; shift 2;;
    --wsl-distro) WSL_DISTRO="${2:-}"; shift 2;;
    --wsl-user)   WSL_USER="${2:-}"; shift 2;;
    --no-linux)   DO_LINUX=0; shift;;
    --no-windows) DO_WINDOWS=0; shift;;
    --no-wrapper) DO_WRAPPER=0; shift;;
    --cgo)        CGO="${2:-}"; shift 2;;
    -h|--help)    usage; exit 0;;
    *)            die "unknown option: $1 (use --help)";;
  esac
done

[[ -n "$REPO_ROOT" ]] || REPO_ROOT="$(detect_repo_root)"
[[ -d "$REPO_ROOT" ]] || die "repo root not found: $REPO_ROOT"

[[ -n "$WIN_USER" ]] || WIN_USER="$(detect_win_user)"

have go || die "go not found in PATH"

# Paths
LINUX_OUT="$REPO_ROOT/oraclepack"
WIN_OUT="$REPO_ROOT/oraclepack.exe"

LINUX_INSTALL_DIR="$HOME/.local/bin"
LINUX_INSTALL_PATH="$LINUX_INSTALL_DIR/oraclepack"

WIN_HOME="/mnt/c/Users/$WIN_USER"
WIN_LOCAL_BIN_DIR="$WIN_HOME/.local/bin"
WIN_LOCAL_BIN_PATH="$WIN_LOCAL_BIN_DIR/oraclepack.exe"

WIN_GITBASH_BIN_DIR="$WIN_HOME/bin"
WIN_GITBASH_WRAPPER_PATH="$WIN_GITBASH_BIN_DIR/oraclepack"

# Optional CGO toggle
if [[ -n "$CGO" ]]; then
  [[ "$CGO" == "0" || "$CGO" == "1" ]] || die "--cgo must be 0 or 1"
  export CGO_ENABLED="$CGO"
fi

cd "$REPO_ROOT"

# Build binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Building Linux (WSL) binary: $LINUX_OUT"
  go build -o "$LINUX_OUT" ./cmd/oraclepack
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Building Windows amd64 exe: $WIN_OUT"
  GOOS=windows GOARCH=amd64 go build -o "$WIN_OUT" ./cmd/oraclepack
fi

# Install binaries
if [[ "$DO_LINUX" -eq 1 ]]; then
  echo "==> Installing Linux binary -> $LINUX_INSTALL_PATH"
  mkdir -p "$LINUX_INSTALL_DIR"
  cp -f "$LINUX_OUT" "$LINUX_INSTALL_PATH"
fi

if [[ "$DO_WINDOWS" -eq 1 ]]; then
  echo "==> Installing Windows exe -> $WIN_LOCAL_BIN_PATH"
  mkdir -p "$WIN_LOCAL_BIN_DIR"
  cp -f "$WIN_OUT" "$WIN_LOCAL_BIN_PATH"
fi

# Write Git Bash wrapper (stored on Windows filesystem)
if [[ "$DO_WRAPPER" -eq 1 ]]; then
  echo "==> Writing Git Bash wrapper -> $WIN_GITBASH_WRAPPER_PATH"
  mkdir -p "$WIN_GITBASH_BIN_DIR"

  cat > "$WIN_GITBASH_WRAPPER_PATH" <<EOF
#!/usr/bin/env bash
set -euo pipefail

# Git for Windows (Git Bash) path-conversion off for this exec call.
# Required so /home/... is not rewritten into C:/Program Files/Git/...
MSYS_NO_PATHCONV=1 exec wsl.exe -d ${WSL_DISTRO@Q} -u ${WSL_USER@Q} -- ${LINUX_INSTALL_PATH@Q} "\$@"
EOF

  # Ensure LF line endings (in case editor/tool wrote CRLF) and best-effort executable bit.
  sed -i 's/\r$//' "$WIN_GITBASH_WRAPPER_PATH" || true
  chmod +x "$WIN_GITBASH_WRAPPER_PATH" 2>/dev/null || true

  echo "==> Note: In Git Bash, ensure ~/bin is in PATH (so 'oraclepack' resolves to this wrapper)."
fi

echo "==> Done."
echo "    WSL binary:     $LINUX_INSTALL_PATH"
echo "    Windows exe:    $WIN_LOCAL_BIN_PATH"
echo "    Git Bash wrap:  $WIN_GITBASH_WRAPPER_PATH"
