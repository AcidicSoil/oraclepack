# path: scripts/tag-release.sh
#!/usr/bin/env bash
set -euo pipefail

read -r -p "Input tag version (e.g., 0.2.0): " version
version="${version//[[:space:]]/}"
[[ -n "${version}" ]] || { echo "Error: version cannot be empty." >&2; exit 1; }

tag="v${version}"

git rev-parse -q --verify "refs/tags/${tag}" >/dev/null && {
  echo "Error: tag ${tag} already exists locally." >&2
  exit 1
}

git ls-remote --tags origin "refs/tags/${tag}" | grep -q . && {
  echo "Error: tag ${tag} already exists on origin." >&2
  exit 1
}

git tag "${tag}"
git push origin "${tag}"
