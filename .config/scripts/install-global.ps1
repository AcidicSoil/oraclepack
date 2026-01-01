# path: scripts/install-global.ps1
[CmdletBinding()]
param(
  [string]$InstallDir = "C:\Tools",
  [switch]$AddToPath,
  [switch]$SkipTests
)

$ErrorActionPreference = "Stop"

function Get-RepoRoot {
  try {
    return (git rev-parse --show-toplevel).Trim()
  } catch {
    throw "Not inside a git repo (expected oraclepack repo)."
  }
}

$RepoRoot = Get-RepoRoot
Set-Location $RepoRoot

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
  throw "Go not found in PATH."
}

Write-Host "==> Repo: $RepoRoot"
Write-Host "==> InstallDir: $InstallDir"

if (-not $SkipTests) {
  Write-Host "==> go test ./..."
  go test ./...
}

# Match README build target: ./cmd/oraclepack
Write-Host "==> go build -o oraclepack.exe ./cmd/oraclepack"
go build -o oraclepack.exe ./cmd/oraclepack

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
Copy-Item -Force -Path (Join-Path $RepoRoot "oraclepack.exe") -Destination (Join-Path $InstallDir "oraclepack.exe")

if ($AddToPath) {
  $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
  if ($userPath -notmatch [Regex]::Escape($InstallDir)) {
    $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) { $InstallDir } else { "$userPath;$InstallDir" }
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "==> Added to User PATH: $InstallDir (restart terminal to pick up PATH changes)"
  } else {
    Write-Host "==> InstallDir already in User PATH"
  }
}

Write-Host "==> Installed: $InstallDir\oraclepack.exe"
& (Join-Path $InstallDir "oraclepack.exe") --version 2>$null
