# oraclepack

<p align="center">
  <a href="https://github.com/acidicsoil/oraclepack/actions/workflows/ci.yml"><img alt="CI" src="https://github.com/acidicsoil/oraclepack/actions/workflows/ci.yml/badge.svg" /></a>
  <a href="https://github.com/acidicsoil/oraclepack/actions/workflows/release.yml"><img alt="Release" src="https://github.com/acidicsoil/oraclepack/actions/workflows/release.yml/badge.svg" /></a>
  <a href="https://github.com/acidicsoil/oraclepack/releases/latest"><img alt="Release Version" src="https://img.shields.io/github/v/release/acidicsoil/oraclepack?sort=semver" /></a>
  <a href="https://github.com/acidicsoil/oraclepack/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/acidicsoil/oraclepack" /></a>
  <a href="https://github.com/acidicsoil/oraclepack/blob/main/go.mod"><img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/acidicsoil/oraclepack" /></a>
</p>

`oraclepack` is a polished, TUI-driven wrapper/runner for **Oracle Packs**—interactive bash workflows embedded in Markdown utilizing [oracle](https://github.com/steipete/oracle). It lets teams ship runbooks, audits, migrations, and LLM evaluation scripts as self-documenting `.md` files that can be validated, resumed, and executed with real-time feedback.

## 🎯 Project Scope & Purpose

Oraclepack is built to make **multi-step operational workflows** reproducible and safe:

- **Runbooks you can execute**: keep instructions and commands in one Markdown file.
- **LLM evaluation flows**: wrap `oracle` CLI calls in steps and validate them with dry runs.
- **Team-friendly automation**: share a pack as documentation, then execute it as a guided TUI.
- **Repeatable ops**: state + report files make it easy to resume or audit past runs.

## 🚀 Features

- **Interactive TUI**: browse steps, view live output, and manage execution with keyboard shortcuts.
- **Run All / Resume**: execute sequentially or continue from the last successful step.
- **Overrides Wizard**: select which steps receive extra official `oracle` flags and validate via dry-run.
- **Step Preview**: view a full step (no truncation), toggle wrap, and copy contents.
- **ROI Filtering**: include/exclude steps by ROI with over/under mode.
- **Project URL Management**: save ChatGPT project URLs globally or per project and pick quickly.
- **State + Report Files**: persistent `.state.json` and `.report.json` outputs for traceability.
- **Plain Mode**: run without TUI for CI or automated environments.
- **Markdown Native**: packs live in a single `.md` file with a bash block.

## 📦 Installation

### Install oracle

Install oracle and setup/configure. Get it [here!](https://github.com/steipete/oracle)

### Setting up for creating oracle-packs

### From Source

Ensure you have [Go](https://go.dev/) 1.24+ installed:

```bash
git clone https://github.com/user/oraclepack.git
cd oraclepack
go build -o oraclepack ./cmd/oraclepack
```

### Building for Windows (.exe)

#### On Windows (PowerShell/CMD)

```powershell
go build -o oraclepack.exe ./cmd/oraclepack
```

#### Cross-Compiling for Windows (from Linux/macOS)

```bash
GOOS=windows GOARCH=amd64 go build -o oraclepack.exe ./cmd/oraclepack
```

### Global Installation (Run from Anywhere)

To run `oraclepack` from any directory, move the binary to a location in your system's `PATH`.

#### Linux & macOS

```bash
# Move the binary to /usr/local/bin (requires sudo)
sudo mv oraclepack /usr/local/bin/

# OR install to your local bin (no sudo required)
mkdir -p ~/.local/bin
mv oraclepack ~/.local/bin/
# Note: Ensure ~/.local/bin is in your shell's PATH
```

#### Windows (PowerShell)

1. Create a directory for your tools (e.g., `C:\Tools`).
2. Move `oraclepack.exe` into that directory.
3. Add that directory to your PATH:

```powershell
$env:Path += ";C:\Tools"
[Environment]::SetEnvironmentVariable("Path", $env:Path, [EnvironmentVariableTarget]::User)
```

```bash
oraclepack completion bash >& oraclepack.completion.sh
mkdir -p ~/.local/share/bash-completion/completions
install -m 0644 oraclepack.completion.sh \
  ~/.local/share/bash-completion/completions/oraclepack

# reload shell
exec bash

```

## Fix: Git Bash wrapper with path-conversion disabled

### Run in Git Bash on Windows

```bash
cat > "$HOME/bin/oraclepack" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail

# Git for Windows (Git Bash) path-conversion off for this exec call.
# Required so /home/... is not rewritten into C:/Program Files/Git/...
MSYS_NO_PATHCONV=1 exec wsl.exe -d Ubuntu-24.04 -u user -- /home/user/.local/bin/oraclepack "$@"
EOF

# ensure LF line endings + executable
sed -i 's/\r$//' "$HOME/bin/oraclepack"
chmod +x "$HOME/bin/oraclepack"
hash -r

```

#### WSL (Windows Subsystem for Linux)

Follow the **Linux** instructions above.

## 🛠 Usage

### 1. Run a Pack (Interactive TUI)

```bash
oraclepack run examples/setup-project.md
```

### 2. Run All Steps Sequentially

```bash
oraclepack run examples/setup-project.md --run-all
```

### 3. Resume a Previous Run

```bash
oraclepack run examples/setup-project.md --resume
```

### 4. Plain Mode (Non-Interactive)

```bash
oraclepack run examples/setup-project.md --no-tui
```

### 5. List Steps

```bash
oraclepack list examples/setup-project.md
```

### 6. Validate a Pack

```bash
oraclepack validate examples/setup-project.md
```

### CLI Flags (run)

```bash
oraclepack run <pack.md> \
  --roi-threshold 2.0 \
  --roi-mode over \
  --run-all \
  --resume \
  --stop-on-fail=true \
  --no-tui \
  --out-dir ./out
```

`oraclepack` expects the `oracle` CLI to be available on your PATH. Overrides let you append official `oracle` flags at runtime.

## ⌨️ TUI Controls (Core)

- `enter`: run selected step
- `a`: run all visible steps sequentially
- `f`: set ROI threshold
- `m`: toggle ROI mode (over/under)
- `v`: step preview (full view)
- `o`: overrides wizard (oracle flags + step targeting)
- `u`: edit ChatGPT project URL
- `U`: open saved URL picker (project/global)
- `q`: quit

### Step Preview Controls

- `b` / `esc`: back to list
- `t`: wrap/un-wrap preview
- `c`: copy step contents (falls back to temp file if clipboard fails)

## 🧭 Overrides Wizard

The overrides flow lets you:

- Select official `oracle` flags (e.g., `--files-report`, `--render`, `--render-plain`, `--copy`, `--wait`).
- Target which steps receive those flags.
- Validate the overridden commands with `oracle --dry-run summary` before running.

## 🔗 Project URL Management

Oraclepack can store ChatGPT project URLs in two scopes:

- **Project scope**: `your-pack.chatgpt-urls.json`
- **Global scope**: `~/.oraclepack/chatgpt-urls.json`

Use `U` to pick, add, edit, delete, or set a default URL.

## 📝 Oracle Pack Format

An Oracle Pack is a Markdown file containing exactly one `bash` code block. Steps are identified by a specific header pattern: `# NN)`.

````markdown
# Project Setup Pack

This pack sets up the development environment.

```bash
# Prelude: Variables defined here are available to all steps
out_dir="dist"

# 01) Initialize dependencies
npm install

# 02) Build the project
npm run build

# 03) Run oracle query (ROI=4.5)
oracle query "check-integrity"
```
````

### Rules

1. Steps must start with `# NN)` (e.g., `# 01)`, `# 02)`).
2. Step numbering must be sequential starting from `01`.
3. The first bash code block in the file is the one executed.
4. Everything before the first `# 01)` is the **prelude**, which runs once.
5. Optional ROI tags (`ROI=2.5`) can be embedded in the step header for filtering.

## 📊 Reports and State

- **State File:** `[pack-name].state.json` tracks step statuses and filter settings.
- **Report File:** `[pack-name].report.json` includes timing, exit codes, and metadata for each step.

## 🤝 Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.
