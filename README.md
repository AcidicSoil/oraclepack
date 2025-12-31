# oraclepack

`oraclepack` is a polished, TUI-driven runner for **Oracle Packs**—interactive bash scripts embedded within Markdown files. It's designed for developers who want to share, validate, and execute complex multi-step workflows with real-time feedback and state persistence.

## 🚀 Features

- **Interactive TUI:** A beautiful Terminal UI built with Bubble Tea for navigating and executing steps.
- **State Persistence:** Automatically saves progress to `.state.json`, allowing you to `--resume` after failure or interruption.
- **Real-time Streaming:** Watch command output live in a styled viewport.
- **Flag Injection:** Safely injects extra flags into `oracle` command calls within your scripts.
- **Machine Readable Reports:** Generates a detailed `.report.json` summary after every run.
- **Markdown Native:** Keep your documentation and execution logic in a single `.md` file.

## 📦 Installation

### From Source

Ensure you have [Go](https://go.dev/) 1.24+ installed:

```bash
git clone https://github.com/user/oraclepack.git
cd oraclepack
go build -o oraclepack ./cmd/oraclepack
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

#### WSL (Windows Subsystem for Linux)
Follow the **Linux** instructions above.

## 🛠 Usage

### 1. Run a Pack (Interactive TUI)

This is the default mode. It opens an interactive list where you can select steps and run them.

```bash
oraclepack run examples/setup-project.md
```

### 2. Resume a Previous Run

If a step failed, fix the issue and resume exactly where you left off:

```bash
oraclepack run examples/setup-project.md --resume
```

### 3. Plain Mode (Non-Interactive)

Ideal for CI/CD or users who prefer standard terminal output:

```bash
oraclepack run examples/setup-project.md --no-tui
```

### 4. Inject Oracle Flags

Pass extra flags down to every `oracle` command invocation inside the pack:

```bash
oraclepack run my-pack.md --oracle-bin="/path/to/oracle" -- --verbose --debug
```

### 5. Validate a Pack

Check if your Markdown file follows the Oracle Pack specification:

```bash
oraclepack validate my-pack.md
```

### 6. List Steps

Quickly view all steps defined in a pack without executing them:

```bash
oraclepack list my-pack.md
```

## 📝 Oracle Pack Format

An Oracle Pack is a standard Markdown file containing exactly one `bash` code block. Steps are identified by a specific header pattern: `# NN)`.

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

# 03) Run oracle query
oracle query "check-integrity"
```
````

### Rules:
1. Steps must start with `# NN)` (e.g., `# 01)`, `# 02)`).
2. Step numbering must be sequential starting from `01`.
3. The first bash code block in the file is the one executed.
4. Everything before the first `# 01)` is the **prelude**, which runs once.

## 📊 Reports and State

- **State File:** `[pack-name].state.json` tracks which steps have succeeded.
- **Report File:** `[pack-name].report.json` contains metadata, duration, and exit codes for every step.

## 🤝 Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.