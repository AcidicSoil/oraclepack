#!/usr/bin/env bash
set -euo pipefail

out_dir="${1:-oracle-out}"

if [[ ! -d "${out_dir}" ]]; then
  echo "ERROR: out_dir does not exist: ${out_dir}" >&2
  exit 2
fi

shopt -s nullglob

for n in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20; do
  matches=( "${out_dir}/${n}-"*.md )
  if [[ "${#matches[@]}" -eq 0 ]]; then
    echo "ERROR: missing output for prefix ${n}: expected ${out_dir}/${n}-*.md" >&2
    exit 3
  fi
  if [[ "${#matches[@]}" -gt 1 ]]; then
    echo "ERROR: multiple outputs for prefix ${n}; expected exactly one file." >&2
    printf '%s\n' "${matches[@]}" >&2
    exit 4
  fi
  printf '%s\n' "${matches[0]}"
done
