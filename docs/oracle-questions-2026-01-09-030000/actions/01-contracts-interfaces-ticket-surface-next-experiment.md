Next smallest concrete experiment

* In the oraclepack repo root, locate the current oracle-only detection + validation callsites to scope the minimal dispatcher/validation expansion: `rg -n --hidden --glob '!.git/' -F -e '^(\\s*)(oracle)\\b' -e 'oracle --dry-run summary' -e 'override validation' -e 'injects flags into commands that begin with' .` 
