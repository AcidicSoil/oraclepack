#!/usr/bin/env bash
set -euo pipefail

pack_path="${1:-}"
if [[ -z "${pack_path}" ]]; then
  echo "Usage: validate-action-pack.sh <path/to/oracle-actions-pack.md>" >&2
  exit 2
fi

if [[ ! -f "${pack_path}" ]]; then
  echo "ERROR: file not found: ${pack_path}" >&2
  exit 3
fi

# Rule: exactly one bash fence, and no other fences.
bash_fence_count="$(grep -cE '^[[:space:]]*```bash[[:space:]]*$' "${pack_path}" || true)"
any_fence_count="$(grep -cE '^[[:space:]]*```' "${pack_path}" || true)"

if [[ "${bash_fence_count}" -ne 1 ]]; then
  echo "ERROR: expected exactly one '```bash' fence; found ${bash_fence_count}" >&2
  exit 10
fi

if [[ "${any_fence_count}" -ne 2 ]]; then
  echo "ERROR: expected exactly 2 total fences (start+end); found ${any_fence_count}" >&2
  echo "Fences found:" >&2
  grep -nE '^[[:space:]]*```' "${pack_path}" >&2 || true
  exit 11
fi

# Extract lines within the bash fence and validate step headers.
in_bash=0
headers=()
while IFS= read -r line; do
  if [[ "${line}" =~ ^[[:space:]]*```bash[[:space:]]*$ ]]; then
    in_bash=1
    continue
  fi
  if [[ "${line}" =~ ^[[:space:]]*```[[:space:]]*$ ]]; then
    in_bash=0
    continue
  fi
  if [[ "${in_bash}" -eq 1 ]]; then
    if [[ "${line}" =~ ^#\ ([0-9]{2})\) ]]; then
      headers+=( "${BASH_REMATCH[1]}" )
    fi
  fi
done < "${pack_path}"

if [[ "${#headers[@]}" -lt 1 ]]; then
  echo "ERROR: no step headers found inside bash fence (expected '# NN)')" >&2
  exit 20
fi

# Validate strict sequential: 01..N with no gaps and no duplicates.
seen=""
expected=1
for h in "${headers[@]}"; do
  # Check duplicates
  if [[ " ${seen} " == *" ${h} "* ]]; then
    echo "ERROR: duplicate step header: ${h}" >&2
    exit 21
  fi
  seen="${seen} ${h}"

  exp="$(printf '%02d' "${expected}")"
  if [[ "${h}" != "${exp}" ]]; then
    echo "ERROR: non-sequential step header. Expected ${exp}, got ${h}" >&2
    echo "All headers: ${headers[*]}" >&2
    exit 22
  fi
  expected=$((expected + 1))
done

echo "OK: Action Pack validation passed."
