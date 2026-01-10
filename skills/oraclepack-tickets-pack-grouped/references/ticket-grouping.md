# Ticket grouping (deterministic, inferred)

Objective: split tickets into focused topic/domain groups and generate one pack per group.

## Inputs

- `ticket_root` (default `.tickets`)
- `ticket_glob` (default `**/*.md`, relative to `ticket_root`)
- `ticket_paths` (optional; comma-separated explicit files; if present, ignore `ticket_glob`)
- `group_mode` (default `subdir+infer`)
- `group_min_score` (default `0.08`)
- `group_max_files` (default `25`; max tickets per pack; >0)
- `group_max_chars` (default `200000`; max total chars per pack; >0)
- `dedupe_mode` (default `report`; one of `off`, `report`, `prune`, `merge`)
- `dedupe_jaccard` (default `0.55`)
- `dedupe_overlap_hi` (default `0.80`)
- `dedupe_overlap_lo` (default `0.70`)
- `dedupe_delta_min` (default `0.15`)
- `dedupe_body_chars` (default `2000`)

## Deterministic grouping rules

1) Collect tickets:
- If `ticket_paths` is non-empty: split on commas, trim whitespace, use exactly that list.
- Else: glob `ticket_root/ticket_glob`.
- Always sort lexicographically by path string.
- Apply `.gitignore` from the nearest parent containing one (primary).

2) Detect possible duplicates (if `dedupe_mode != off`):
- Signature: filename stem + first heading + first `dedupe_body_chars` chars.
- Compute `jaccard` + `overlap` between tickets.
- Duplicate edge rule:
  - `overlap >= dedupe_overlap_hi` OR (`jaccard >= dedupe_jaccard` AND `overlap >= dedupe_overlap_lo`)
- Connected components become duplicate clusters.
- Canonical: largest content length; tie-break lexicographic.
- Delta vs redundant:
  - delta if unique token ratio >= `dedupe_delta_min` OR heading differs materially.
  - redundant otherwise.

3) Seed groups by subdir:
- For any path under `ticket_root/<group>/...`, assign to group `<group>`.
- Tickets directly under `ticket_root/` are "loose".

4) Infer loose tickets into groups (if any groups exist):
- Build a token set for each group from:
  - group name tokens
  - ticket filenames (stem tokens)
  - first Markdown heading line (if present)
- For each loose ticket, compute Jaccard overlap score with each group token set.
- If `max_score >= group_min_score`, assign to the best group (stable tie-break by group name).
- Otherwise, assign to `misc`.

5) If no groups exist:
- Put all tickets into a single group named `root`.

6) Merge duplicates into primary group:
- `report`: attach all tickets in the cluster to the canonical’s group.
- `prune`: attach canonical + delta only; drop redundant from attachments.
- `merge`: create `out_dir/_ticket_merges/cluster-XXXX.md` and attach only the merged file.
- Emit `_dupes_possible.json`, `_duplicates.json`, and `_dedupe_plan.json`.

7) Split oversized groups:
- If a group exceeds `group_max_files` or `group_max_chars`, split into parts (1..N)
  in sorted order, chunked deterministically.

Hard rule: do not use mtimes, file sizes, or external ML services.

## Required outputs

- `_groups.json`: mapping of group -> list of ticket paths (lexicographic order)
- Pack file per group (and part), each self-contained and direct-attach
- `manifest.json`: groups with pack path + attached vs original ticket lists
