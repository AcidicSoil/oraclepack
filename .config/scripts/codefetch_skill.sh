# path: scripts/codefetch_skill.sh
#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/codefetch_skill.sh <skill-name> [output-file]

The argument must be the skill directory name under ./skills (no prefix assumptions),
e.g.:
  bash scripts/codefetch_skill.sh oraclepack-tickets-pack
  bash scripts/codefetch_skill.sh tickets-pack
  bash scripts/codefetch_skill.sh my-skill-name ./out.md

Env overrides:
  CODEFETCH_THREADS   (default: 5)
  CODEFETCH_MAXTOKENS (default: 75000)
USAGE
}

if [[ "${1:-}" == "" || "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

skill_name="$1"
threads="${CODEFETCH_THREADS:-5}"
max_tokens="${CODEFETCH_MAXTOKENS:-75000}"

skill_dir="skills/${skill_name%/}"

if [[ ! -d "$skill_dir" ]]; then
  echo "Error: skill directory not found: $skill_dir" >&2
  echo "Expected a folder under ./skills/ named exactly: '$skill_name'" >&2
  exit 1
fi

out="${2:-${skill_name}_skill.md}"

include_files=(
  "${skill_dir}/SKILL.md"
  "${skill_dir}/assets/**/*"
  "${skill_dir}/references/**/*"
  "${skill_dir}/scripts/**/*"
)

include_files_csv=""
for pat in "${include_files[@]}"; do
  if [[ -z "$include_files_csv" ]]; then
    include_files_csv="$pat"
  else
    include_files_csv="${include_files_csv},$pat"
  fi
done

codefetch -t "$threads" \
  --include-dir "$skill_dir" \
  --include-files "$include_files_csv" \
  -o "$out" \
  --max-tokens "$max_tokens"

echo "Wrote: $out"
