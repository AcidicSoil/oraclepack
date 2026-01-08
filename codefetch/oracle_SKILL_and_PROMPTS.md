<filetree>
Project Structure:
├── .config
│   ├── commands
│   │   └── oracle-pack_v2.toml
│   ├── completion
│   │   └── oraclepack.completion.sh
│   ├── mcp
│   │   ├── oraclepack-gold-pack
│   │   │   ├── references
│   │   │   │   ├── attachment-minimization.md
│   │   │   │   ├── inference-first-discovery.md
│   │   │   │   ├── oracle-pack-template.md
│   │   │   │   └── oracle-scratch-format.md
│   │   │   ├── scripts
│   │   │   │   ├── lint_attachments.py
│   │   │   │   └── validate_pack.py
│   │   │   └── SKILL.md
│   │   └── oraclepack-taskify
│   │       ├── assets
│   │       │   ├── action-pack-template.md
│   │       │   ├── actions-json-schema.md
│   │       │   └── prd-synthesis-prompt.md
│   │       ├── references
│   │       │   ├── determinism-and-safety.md
│   │       │   ├── task-master-cli-cheatsheet.md
│   │       │   └── workflow-overview.md
│   │       ├── scripts
│   │       │   ├── detect-oracle-outputs.sh
│   │       │   └── validate-action-pack.sh
│   │       └── SKILL.md
│   └── skills
│       ├── oracle-pack
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   └── oracle-scratch-format.md
│       │   └── SKILL.md
│       ├── oracle-pack-rpg-infused
│       │   ├── assets
│       │   │   └── oracle-pack-template.md
│       │   ├── references
│       │   │   └── oracle-pack-spec.md
│       │   ├── scripts
│       │   │   └── validate_oracle_pack.py
│       │   └── SKILL.md
│       ├── oracle-strategist-commands
│       │   └── SKILL.md
│       ├── oraclepack-gold-pack
│       │   ├── references
│       │   │   ├── attachment-minimization.md
│       │   │   ├── inference-first-discovery.md
│       │   │   ├── oracle-pack-template.md
│       │   │   └── oracle-scratch-format.md
│       │   ├── scripts
│       │   │   ├── lint_attachments.py
│       │   │   └── validate_pack.py
│       │   └── SKILL.md
│       ├── oraclepack-taskify
│       │   ├── assets
│       │   │   ├── action-pack-template.md
│       │   │   ├── actions-json-schema.md
│       │   │   └── prd-synthesis-prompt.md
│       │   ├── references
│       │   │   ├── determinism-and-safety.md
│       │   │   ├── task-master-cli-cheatsheet.md
│       │   │   └── workflow-overview.md
│       │   ├── scripts
│       │   │   ├── detect-oracle-outputs.sh
│       │   │   └── validate-action-pack.sh
│       │   └── SKILL.md
│       ├── oraclepack-usecase-rpg
│       │   ├── references
│       │   │   └── oracle-pack-template.md
│       │   ├── scripts
│       │   │   └── validate_oraclepack_pack.py
│       │   └── SKILL.md
│       └── strategist-questions-oracle-pack
│           ├── assets
│           │   └── oracle-pack-template.md
│           └── references
│               └── oracle-scratch-format.md
├── .tickets
│   └── mcp
│       ├── oraclepack-MCP.md
│       └── oraclepack_mcp_server.md
├── docs
│   ├── oracle-questions-2026-01-07
│   │   ├── 01-contracts-interfaces-surface.md
│   │   ├── 02-contracts-interfaces-integration.md
│   │   ├── 03-invariants-domain.md
│   │   ├── 04-invariants-validation.md
│   │   ├── 05-caching-state-layers.md
│   │   ├── 06-caching-state-consistency.md
│   │   ├── 07-background-jobs-discovery.md
│   │   ├── 08-background-jobs-reliability.md
│   │   ├── 09-observability-signals.md
│   │   ├── 10-observability-gaps.md
│   │   ├── 11-permissions-model.md
│   │   ├── 12-permissions-enforcement.md
│   │   ├── 13-migrations-schema.md
│   │   ├── 14-migrations-compat.md
│   │   ├── 15-ux-flows-primary.md
│   │   ├── 16-ux-flows-edgecases.md
│   │   ├── 17-failure-modes-taxonomy.md
│   │   ├── 18-failure-modes-resilience.md
│   │   ├── 19-feature-flags-inventory.md
│   │   └── 20-feature-flags-rollout.md
│   ├── oracle-actions-pack-2026-01-07.md
│   └── oracle-pack-2026-01-07.md
├── internal
│   └── exec
│       ├── oracle_scan.go
│       ├── oracle_scan_test.go
│       ├── oracle_validate.go
│       └── oracle_validate_test.go
├── oraclepack-mcp-server
│   ├── .pytest_cache
│   │   ├── v
│   │   │   └── cache
│   │   │       ├── lastfailed
│   │   │       └── nodeids
│   │   └── CACHEDIR.TAG
│   ├── oraclepack_mcp_server
│   │   ├── __init__.py
│   │   ├── __main__.py
│   │   ├── config.py
│   │   ├── oraclepack_cli.py
│   │   ├── security.py
│   │   ├── server.py
│   │   └── taskify.py
│   ├── oraclepack_mcp_server.egg-info
│   │   ├── PKG-INFO
│   │   ├── SOURCES.txt
│   │   ├── dependency_links.txt
│   │   ├── entry_points.txt
│   │   ├── requires.txt
│   │   └── top_level.txt
│   ├── tests
│   │   ├── test_cli.py
│   │   ├── test_config.py
│   │   ├── test_integration.py
│   │   ├── test_security.py
│   │   └── test_taskify.py
│   ├── venv
│   │   ├── bin
│   │   │   ├── Activate.ps1
│   │   │   ├── activate
│   │   │   ├── activate.csh
│   │   │   ├── activate.fish
│   │   │   ├── dotenv
│   │   │   ├── httpx
│   │   │   ├── jsonschema
│   │   │   ├── markdown-it
│   │   │   ├── mcp
│   │   │   ├── oraclepack-mcp
│   │   │   ├── pip
│   │   │   ├── pip3
│   │   │   ├── pip3.12
│   │   │   ├── py.test
│   │   │   ├── pygmentize
│   │   │   ├── pytest
│   │   │   ├── python
│   │   │   ├── python3
│   │   │   ├── python3.12
│   │   │   ├── typer
│   │   │   └── uvicorn
│   │   ├── lib
│   │   │   └── python3.12
│   │   │       └── site-packages
│   │   ├── lib64
│   │   │   └── python3.12
│   │   │       └── site-packages
│   │   └── pyvenv.cfg
│   ├── pyproject.toml
│   └── requirements.txt
├── oracle-pack-2026-01-02.chatgpt-urls.json
└── oracle-pack-2026-01-02.state.json

</filetree>

<source_code>
oracle-pack-2026-01-02.chatgpt-urls.json
```
{
  "default": "",
  "items": [
    {
      "name": "oracle",
[TRUNCATED]
```

oracle-pack-2026-01-02.state.json
```
{
  "schema_version": 1,
  "pack_hash": "",
[TRUNCATED]
```

docs/oracle-actions-pack-2026-01-07.md
```
# Oraclepack Stage 3 — Action Pack (Taskify)

[TRUNCATED]
```

docs/oracle-pack-2026-01-07.md
```
# Oracle Pack — Unknown (Gold Stage 1)

[TRUNCATED]
```

oraclepack-mcp-server/pyproject.toml
```
[build-system]
requires = ["setuptools>=61.0"]
[TRUNCATED]
```

oraclepack-mcp-server/requirements.txt
```
mcp[cli]
pydantic-settings
pydantic>=2.0
```

.config/commands/oracle-pack_v2.toml
```
# path: commands/oraclepack/oracle-pack_v2.toml
[TRUNCATED]
```

.config/completion/oraclepack.completion.sh
```
# bash completion V2 for oraclepack                           -*- shell-script -*-

__oraclepack_debug()
{
[TRUNCATED]
```

.tickets/mcp/oraclepack-MCP.md
```
[TRUNCATED]
```

.tickets/mcp/oraclepack_mcp_server.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/01-contracts-interfaces-surface.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/02-contracts-interfaces-integration.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/03-invariants-domain.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/04-invariants-validation.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/05-caching-state-layers.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/06-caching-state-consistency.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/07-background-jobs-discovery.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/08-background-jobs-reliability.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/09-observability-signals.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/10-observability-gaps.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/11-permissions-model.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/12-permissions-enforcement.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/13-migrations-schema.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/14-migrations-compat.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/15-ux-flows-primary.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/16-ux-flows-edgecases.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/17-failure-modes-taxonomy.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/18-failure-modes-resilience.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/19-feature-flags-inventory.md
```
[TRUNCATED]
```

docs/oracle-questions-2026-01-07/20-feature-flags-rollout.md
```
[TRUNCATED]
```

internal/exec/oracle_scan.go
```
package exec

import (
	"regexp"
	"strings"
)

[TRUNCATED]
```

internal/exec/oracle_scan_test.go
```
package exec

import (
	"reflect"
	"testing"
)

[TRUNCATED]
```

internal/exec/oracle_validate.go
```
package exec

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"

[TRUNCATED]
```

internal/exec/oracle_validate_test.go
```
package exec

import (
	"context"
	"os"
	"path/filepath"
	"strings"
[TRUNCATED]
```

oraclepack-mcp-server/.pytest_cache/CACHEDIR.TAG
```
Signature: 8a477f597d28d172789f06886806bc55
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/__init__.py
```
```

oraclepack-mcp-server/oraclepack_mcp_server/__main__.py
```
import argparse
import asyncio
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/config.py
```
import os
from pathlib import Path
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/oraclepack_cli.py
```
import asyncio
import time
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/security.py
```
import os
from pathlib import Path
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/server.py
```
import logging
import sys
import os
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server/taskify.py
```
import os
import re
import glob
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/PKG-INFO
```
Metadata-Version: 2.4
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/SOURCES.txt
```
README.md
pyproject.toml
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/dependency_links.txt
```

```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/entry_points.txt
```
[console_scripts]
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/requires.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/oraclepack_mcp_server.egg-info/top_level.txt
```
oraclepack_mcp_server
```

oraclepack-mcp-server/tests/test_cli.py
```
import asyncio
import pytest
[TRUNCATED]
```

oraclepack-mcp-server/tests/test_config.py
```
import os
import pytest
[TRUNCATED]
```

oraclepack-mcp-server/tests/test_integration.py
```
import asyncio
import pytest
import sys
[TRUNCATED]
```

oraclepack-mcp-server/tests/test_security.py
```
import os
import pytest
[TRUNCATED]
```

oraclepack-mcp-server/tests/test_taskify.py
```
import pytest
from pathlib import Path
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/SKILL.md
```
---
name: oraclepack-gold-pack
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/SKILL.md
```
---
name: oraclepack-taskify
[TRUNCATED]
```

.config/skills/oracle-pack/SKILL.md
```
---
name: oracle-pack
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/SKILL.md
```
---
name: oracle-pack-rpg-infused
[TRUNCATED]
```

.config/skills/oracle-strategist-commands/SKILL.md
```
---

name: oracle-strategist-commands
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/SKILL.md
```
---
name: oraclepack-gold-pack
[TRUNCATED]
```

.config/skills/oraclepack-taskify/SKILL.md
```
---
name: oraclepack-taskify
[TRUNCATED]
```

.config/skills/oraclepack-usecase-rpg/SKILL.md
```
---
name: oraclepack-usecase-rpg
[TRUNCATED]
```

oraclepack-mcp-server/venv/pyvenv.cfg
```
home = /usr/bin
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/attachment-minimization.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/inference-first-discovery.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/oracle-pack-template.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/references/oracle-scratch-format.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/scripts/lint_attachments.py
```
[TRUNCATED]
```

.config/mcp/oraclepack-gold-pack/scripts/validate_pack.py
```
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/assets/action-pack-template.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/assets/actions-json-schema.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/assets/prd-synthesis-prompt.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/references/determinism-and-safety.md
```
# Determinism and safety guardrails

[TRUNCATED]
```

.config/mcp/oraclepack-taskify/references/task-master-cli-cheatsheet.md
```
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/references/workflow-overview.md
```
# Stage 3 (oraclepack-taskify) — Workflow overview

[TRUNCATED]
```

.config/mcp/oraclepack-taskify/scripts/detect-oracle-outputs.sh
```
#!/usr/bin/env bash
[TRUNCATED]
```

.config/mcp/oraclepack-taskify/scripts/validate-action-pack.sh
```
#!/usr/bin/env bash
[TRUNCATED]
```

.config/skills/oracle-pack/assets/oracle-pack-template.md
```
[TRUNCATED]
```

.config/skills/oracle-pack/references/attachment-minimization.md
```
[TRUNCATED]
```

.config/skills/oracle-pack/references/inference-first-discovery.md
```
[TRUNCATED]
```

.config/skills/oracle-pack/references/oracle-scratch-format.md
```
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/assets/oracle-pack-template.md
```
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/references/oracle-pack-spec.md
```
[TRUNCATED]
```

.config/skills/oracle-pack-rpg-infused/scripts/validate_oracle_pack.py
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/references/attachment-minimization.md
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/references/inference-first-discovery.md
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/references/oracle-pack-template.md
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/references/oracle-scratch-format.md
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/scripts/lint_attachments.py
```
[TRUNCATED]
```

.config/skills/oraclepack-gold-pack/scripts/validate_pack.py
```
[TRUNCATED]
```

.config/skills/oraclepack-taskify/assets/action-pack-template.md
```
[TRUNCATED]
```

.config/skills/oraclepack-taskify/assets/actions-json-schema.md
```
[TRUNCATED]
```

.config/skills/oraclepack-taskify/assets/prd-synthesis-prompt.md
```
[TRUNCATED]
```

.config/skills/oraclepack-taskify/references/determinism-and-safety.md
```
# Determinism and safety guardrails

[TRUNCATED]
```

.config/skills/oraclepack-taskify/references/task-master-cli-cheatsheet.md
```
[TRUNCATED]
```

.config/skills/oraclepack-taskify/references/workflow-overview.md
```
# Stage 3 (oraclepack-taskify) — Workflow overview

[TRUNCATED]
```

.config/skills/oraclepack-taskify/scripts/detect-oracle-outputs.sh
```
#!/usr/bin/env bash
[TRUNCATED]
```

.config/skills/oraclepack-taskify/scripts/validate-action-pack.sh
```
#!/usr/bin/env bash
[TRUNCATED]
```

.config/skills/oraclepack-usecase-rpg/references/oracle-pack-template.md
```
[TRUNCATED]
```

.config/skills/oraclepack-usecase-rpg/scripts/validate_oraclepack_pack.py
```
[TRUNCATED]
```

.config/skills/strategist-questions-oracle-pack/assets/oracle-pack-template.md
```
[TRUNCATED]
```

.config/skills/strategist-questions-oracle-pack/references/oracle-scratch-format.md
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/Activate.ps1
```
<#
.Synopsis
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/activate
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/activate.csh
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/activate.fish
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/dotenv
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/httpx
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/jsonschema
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/markdown-it
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/mcp
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/oraclepack-mcp
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/pip
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/pip3
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/pip3.12
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/py.test
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/pygmentize
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/pytest
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/python
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/python3
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/python3.12
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/typer
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/bin/uvicorn
```
[TRUNCATED]
```

oraclepack-mcp-server/.pytest_cache/v/cache/lastfailed
```
{}
```

oraclepack-mcp-server/.pytest_cache/v/cache/nodeids
```
[
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/__editable__.oraclepack_mcp_server-0.1.0.pth
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/__editable___oraclepack_mcp_server_0_1_0_finder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/py.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_extensions.py
```
import abc
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/__editable__.oraclepack_mcp_server-0.1.0.pth
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/__editable___oraclepack_mcp_server_0_1_0_finder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/py.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_extensions.py
```
import abc
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/AUTHORS.rst
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/PyJWT-2.10.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_argcomplete.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/cacheprovider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/capture.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/debugging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/deprecated.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/doctest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/faulthandler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/fixtures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/freeze_support.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/helpconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/hookspec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/junitxml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/legacypath.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/monkeypatch.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/nodes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/outcomes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/pastebin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/pathlib.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/pytester.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/pytester_assertions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/python_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/raises.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/recwarn.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/reports.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/runner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/scope.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/setuponly.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/setupplan.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/skipping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/stash.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/stepwise.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/subtests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/terminal.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/terminalprogress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/threadexception.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/timing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/tmpdir.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/tracemalloc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/unittest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/unraisableexception.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/warning_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types/test_cases.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types-0.7.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types-0.7.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types-0.7.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/annotated_types-0.7.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/from_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/functools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/lowlevel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/pytest_plugin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/to_interpreter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/to_process.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/to_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio-4.12.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/__init__.pyi
```
import enum
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_cmp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_cmp.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_funcs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_make.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_next_gen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_typing_compat.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_version_info.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/_version_info.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/converters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/converters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/exceptions.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/filters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/filters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/setters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/setters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attr/validators.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/__init__.pyi
```
import sys

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/converters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/filters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/setters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs-25.4.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs-25.4.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs-25.4.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/attrs-25.4.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi/__main__.py
```
import argparse

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi/core.py
```
"""
certifi.py
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi-2026.1.4.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi-2026.1.4.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi-2026.1.4.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi-2026.1.4.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/certifi-2026.1.4.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/_cffi_errors.h
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/_cffi_include.h
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/_embedding.h
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/_imp_emulation.py
```

try:
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/_shimmed_dist_utils.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/backend_ctypes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/cffi_opcode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/commontypes.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/cparser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/error.py
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/ffiplatform.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/lock.py
```
import sys

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/model.py
```
import types
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/parse_c_type.h
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/pkgconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/recompiler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/setuptools_ext.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/vengine_cpy.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/vengine_gen.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi/verifier.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cffi-2.0.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/_termui_impl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/_textwrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/_winconsole.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/decorators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/formatting.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/globals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/parser.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/shell_completion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/termui.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click-8.3.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click-8.3.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click-8.3.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/click-8.3.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/__about__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/fernet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography-46.0.3.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography-46.0.3.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography-46.0.3.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography-46.0.3.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/cli.py
```
import json
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/ipython.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/main.py
```
import io
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/parser.py
```
import codecs
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/variables.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/dotenv/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_abnf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_events.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_headers.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_readers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_receivebuffer.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_state.py
```
################################################################
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/_writers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11/py.typed
```
Marker
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11-0.16.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11-0.16.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11-0.16.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11-0.16.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/h11-0.16.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_exceptions.py
```
import contextlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_ssl.py
```
import ssl

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_synchronization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_trace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore-1.0.9.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore-1.0.9.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore-1.0.9.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore-1.0.9.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/__version__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_client.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_content.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_decoders.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_urlparse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_urls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx-0.28.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx-0.28.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx-0.28.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx-0.28.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx-0.28.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/_decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse-0.4.3.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse-0.4.3.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse-0.4.3.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse-0.4.3.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx_sse-0.4.3.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/codec.py
```
import codecs
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/core.py
```
import bisect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/idnadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/intranges.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/package_data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna/uts46data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna-3.11.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna-3.11.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna-3.11.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/idna-3.11.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig/_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig-2.3.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig-2.3.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig-2.3.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig-2.3.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/iniconfig-2.3.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/__main__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_keywords.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_legacy_keywords.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_typing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/cli.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/protocols.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/validators.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema-4.26.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema-4.26.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema-4.26.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema-4.26.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema-4.26.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/_core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/api_jwk.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/api_jws.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/api_jwt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/help.py
```
import json
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/jwk_set_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/jwks_client.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/utils.py
```
import base64
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jwt/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/_punycode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/parser_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/parser_core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/parser_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/port.yaml
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/renderer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/ruler.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp-1.25.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp-1.25.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp-1.25.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp-1.25.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp-1.25.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/__init__.py
```
__all__ = (
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/_decode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/_encode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl-0.1.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl-0.1.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl-0.1.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mdurl-0.1.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/multipart/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/multipart/decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/multipart/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/multipart/multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/direct_url.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/_elffile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/_manylinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/_musllinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/_structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/specifiers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging-25.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging-25.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging-25.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/packaging-25.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/__main__.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/__pip-runner__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/AUTHORS.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip-24.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/__init__.py
```
__all__ = [
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_callers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_hooks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_result.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_tracing.py
```
"""
Tracing utils
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/_warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy-1.6.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy-1.6.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy-1.6.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy-1.6.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pluggy-1.6.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/__init__.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/_ast_gen.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/_build_tables.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/_c_ast.cfg
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ast_transforms.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/c_ast.py
```
#-----------------------------------------------------------------
# ** ATTENTION **
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/c_generator.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/c_lexer.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/c_parser.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/lextab.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/plyparser.py
```
#-----------------------------------------------------------------
# plyparser.py
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/yacctab.py
```

# yacctab.py
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser-2.23.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser-2.23.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser-2.23.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser-2.23.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser-2.23.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_migration.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/alias_generators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/aliases.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/annotated_handlers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/color.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/dataclasses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/datetime_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/env_settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/error_wrappers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/functional_serializers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/functional_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/json_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/mypy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/networks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/root_model.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/type_adapter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/validate_call_decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic-2.12.5.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic-2.12.5.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic-2.12.5.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic-2.12.5.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core/_pydantic_core.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core/core_schema.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core-2.41.5.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core-2.41.5.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core-2.41.5.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_core-2.41.5.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/__init__.py
```
"""
    Pygments
    ~~~~~~~~

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/__main__.py
```
"""
    pygments.__main__
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/cmdline.py
```
"""
    pygments.cmdline
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/console.py
```
"""
    pygments.console
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/filter.py
```
"""
    pygments.filter
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatter.py
```
"""
    pygments.formatter
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexer.py
```
"""
    pygments.lexer
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/modeline.py
```
"""
    pygments.modeline
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/plugin.py
```
"""
    pygments.plugin
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/regexopt.py
```
"""
    pygments.regexopt
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/scanner.py
```
"""
    pygments.scanner
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/sphinxext.py
```
"""
    pygments.sphinxext
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/style.py
```
"""
    pygments.style
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/unistring.py
```
"""
    pygments.unistring
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/util.py
```
"""
    pygments.util
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments-2.19.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments-2.19.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments-2.19.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments-2.19.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments-2.19.2.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest-9.0.2.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio/plugin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_dotenv-1.2.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart/decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart/multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart-0.0.21.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart-0.0.21.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart-0.0.21.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/python_multipart-0.0.21.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/_attrs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/_attrs.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/jsonschema.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/retrieval.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/typing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing-0.37.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing-0.37.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing-0.37.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing-0.37.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/__main__.py
```
import colorsys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_cell_widths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_emoji_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_emoji_replace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_export_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_extension.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_fileno.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_inspect.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_log_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_loop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_null_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_palettes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_pick.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_ratio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_spinners.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_stack.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_timer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_win32_console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_windows.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_windows_renderer.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/_wrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/align.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/ansi.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/box.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/cells.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/color.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/color_triplet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/columns.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/console.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/constrain.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/containers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/control.py
```
import time
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/default_styles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/diagnose.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/emoji.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/file_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/filesize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/highlighter.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/jupyter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/layout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/live.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/live_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/logging.py
```
import logging
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/markdown.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/markup.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/measure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/pager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/palette.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/panel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/pretty.py
```
import builtins
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/progress_bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/prompt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/region.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/repr.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/rule.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/screen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/segment.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/spinner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/status.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/style.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/styled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/syntax.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/terminal_theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/text.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/theme.py
```
import configparser
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/themes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/traceback.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich-14.2.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich-14.2.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich-14.2.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rich-14.2.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds_py-0.30.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds_py-0.30.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds_py-0.30.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/rpds_py-0.30.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/__init__.py
```
import importlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/nt.py
```
import contextlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham-1.5.4.dist-info/zip-safe
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette/event.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette/sse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette-3.1.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette-3.1.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette-3.1.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette-3.1.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/sse_starlette-3.1.2.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/_exception_handler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/applications.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/authentication.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/background.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/concurrency.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/convertors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/datastructures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/endpoints.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/formparsers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/requests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/responses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/routing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/schemas.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/staticfiles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/status.py
```
"""
HTTP codes
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/templating.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/testclient.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/websockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette-0.50.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette-0.50.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette-0.50.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette-0.50.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/_completion_classes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/_completion_shared.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/_typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/colors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/completion.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/core.py
```
import errno
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/main.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/models.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/params.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/rich_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer/utils.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer-0.21.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer-0.21.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer-0.21.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer-0.21.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typer-0.21.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_extensions-4.15.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_extensions-4.15.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_extensions-4.15.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_extensions-4.15.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection/introspection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection/typing_objects.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection/typing_objects.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection-0.4.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection-0.4.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection-0.4.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/typing_inspection-0.4.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/__main__.py
```
import uvicorn

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/_subprocess.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/importer.py
```
import importlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/py.typed
```

```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/server.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/workers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn-0.40.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn-0.40.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn-0.40.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn-0.40.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn-0.40.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/AUTHORS.rst
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/PyJWT-2.10.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_argcomplete.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/cacheprovider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/capture.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/debugging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/deprecated.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/doctest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/faulthandler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/fixtures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/freeze_support.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/helpconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/hookspec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/junitxml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/legacypath.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/monkeypatch.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/nodes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/outcomes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/pastebin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/pathlib.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/pytester.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/pytester_assertions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/python_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/raises.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/recwarn.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/reports.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/runner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/scope.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/setuponly.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/setupplan.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/skipping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/stash.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/stepwise.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/subtests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/terminal.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/terminalprogress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/threadexception.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/timing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/tmpdir.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/tracemalloc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/unittest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/unraisableexception.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/warning_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types/test_cases.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types-0.7.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types-0.7.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types-0.7.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/annotated_types-0.7.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/from_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/functools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/lowlevel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/pytest_plugin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/to_interpreter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/to_process.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/to_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio-4.12.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/__init__.pyi
```
import enum
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_cmp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_cmp.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_funcs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_make.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_next_gen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_typing_compat.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_version_info.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/_version_info.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/converters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/converters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/exceptions.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/filters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/filters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/setters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/setters.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attr/validators.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/__init__.pyi
```
import sys

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/converters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/filters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/setters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs-25.4.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs-25.4.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs-25.4.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/attrs-25.4.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi/__main__.py
```
import argparse

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi/core.py
```
"""
certifi.py
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi-2026.1.4.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi-2026.1.4.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi-2026.1.4.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi-2026.1.4.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/certifi-2026.1.4.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/_cffi_errors.h
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/_cffi_include.h
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/_embedding.h
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/_imp_emulation.py
```

try:
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/_shimmed_dist_utils.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/backend_ctypes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/cffi_opcode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/commontypes.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/cparser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/error.py
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/ffiplatform.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/lock.py
```
import sys

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/model.py
```
import types
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/parse_c_type.h
```

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/pkgconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/recompiler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/setuptools_ext.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/vengine_cpy.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/vengine_gen.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi/verifier.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cffi-2.0.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/_termui_impl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/_textwrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/_winconsole.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/decorators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/formatting.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/globals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/parser.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/shell_completion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/termui.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click-8.3.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click-8.3.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click-8.3.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/click-8.3.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/__about__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/fernet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography-46.0.3.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography-46.0.3.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography-46.0.3.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography-46.0.3.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/cli.py
```
import json
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/ipython.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/main.py
```
import io
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/parser.py
```
import codecs
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/variables.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/dotenv/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_abnf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_events.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_headers.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_readers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_receivebuffer.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_state.py
```
################################################################
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/_writers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11/py.typed
```
Marker
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11-0.16.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11-0.16.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11-0.16.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11-0.16.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/h11-0.16.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_exceptions.py
```
import contextlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_ssl.py
```
import ssl

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_synchronization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_trace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore-1.0.9.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore-1.0.9.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore-1.0.9.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore-1.0.9.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/__version__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_client.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_content.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_decoders.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_urlparse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_urls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx-0.28.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx-0.28.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx-0.28.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx-0.28.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx-0.28.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/_decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/_models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse-0.4.3.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse-0.4.3.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse-0.4.3.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse-0.4.3.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx_sse-0.4.3.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/codec.py
```
import codecs
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/core.py
```
import bisect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/idnadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/intranges.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/package_data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna/uts46data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna-3.11.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna-3.11.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna-3.11.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/idna-3.11.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig/_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig-2.3.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig-2.3.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig-2.3.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig-2.3.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/iniconfig-2.3.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/__main__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_keywords.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_legacy_keywords.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_typing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/cli.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/protocols.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/validators.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema-4.26.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema-4.26.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema-4.26.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema-4.26.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema-4.26.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/_core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications-2025.9.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/api_jwk.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/api_jws.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/api_jwt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/help.py
```
import json
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/jwk_set_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/jwks_client.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/utils.py
```
import base64
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jwt/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/_punycode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/parser_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/parser_core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/parser_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/port.yaml
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/renderer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/ruler.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it_py-4.0.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp-1.25.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp-1.25.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp-1.25.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp-1.25.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp-1.25.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/__init__.py
```
__all__ = (
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/_decode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/_encode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl-0.1.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl-0.1.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl-0.1.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mdurl-0.1.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/multipart/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/multipart/decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/multipart/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/multipart/multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/direct_url.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/oraclepack_mcp_server-0.1.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/_elffile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/_manylinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/_musllinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/_structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/specifiers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging-25.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging-25.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging-25.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/packaging-25.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/__main__.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/__pip-runner__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/py.typed
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/AUTHORS.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip-24.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/__init__.py
```
__all__ = [
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_callers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_hooks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_result.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_tracing.py
```
"""
Tracing utils
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/_warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy-1.6.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy-1.6.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy-1.6.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy-1.6.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pluggy-1.6.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/__init__.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/_ast_gen.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/_build_tables.py
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/_c_ast.cfg
```
#-----------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ast_transforms.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/c_ast.py
```
#-----------------------------------------------------------------
# ** ATTENTION **
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/c_generator.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/c_lexer.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/c_parser.py
```
#------------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/lextab.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/plyparser.py
```
#-----------------------------------------------------------------
# plyparser.py
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/yacctab.py
```

# yacctab.py
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser-2.23.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser-2.23.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser-2.23.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser-2.23.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser-2.23.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_migration.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/alias_generators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/aliases.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/annotated_handlers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/color.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/dataclasses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/datetime_parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/env_settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/error_wrappers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/functional_serializers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/functional_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/json_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/mypy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/networks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/root_model.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/type_adapter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/validate_call_decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/warnings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic-2.12.5.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic-2.12.5.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic-2.12.5.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic-2.12.5.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core/_pydantic_core.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core/core_schema.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core-2.41.5.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core-2.41.5.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core-2.41.5.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_core-2.41.5.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings-2.12.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/__init__.py
```
"""
    Pygments
    ~~~~~~~~

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/__main__.py
```
"""
    pygments.__main__
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/cmdline.py
```
"""
    pygments.cmdline
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/console.py
```
"""
    pygments.console
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/filter.py
```
"""
    pygments.filter
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatter.py
```
"""
    pygments.formatter
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexer.py
```
"""
    pygments.lexer
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/modeline.py
```
"""
    pygments.modeline
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/plugin.py
```
"""
    pygments.plugin
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/regexopt.py
```
"""
    pygments.regexopt
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/scanner.py
```
"""
    pygments.scanner
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/sphinxext.py
```
"""
    pygments.sphinxext
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/style.py
```
"""
    pygments.style
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/unistring.py
```
"""
    pygments.unistring
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/util.py
```
"""
    pygments.util
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments-2.19.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments-2.19.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments-2.19.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments-2.19.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments-2.19.2.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest-9.0.2.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio/plugin.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/REQUESTED
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pytest_asyncio-1.3.0.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_dotenv-1.2.1.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart/decoders.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart/multipart.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart-0.0.21.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart-0.0.21.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart-0.0.21.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/python_multipart-0.0.21.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/_attrs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/_attrs.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/jsonschema.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/retrieval.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/typing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing-0.37.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing-0.37.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing-0.37.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing-0.37.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/__main__.py
```
import colorsys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_cell_widths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_emoji_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_emoji_replace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_export_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_extension.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_fileno.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_inspect.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_log_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_loop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_null_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_palettes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_pick.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_ratio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_spinners.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_stack.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_timer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_win32_console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_windows.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_windows_renderer.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/_wrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/align.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/ansi.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/box.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/cells.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/color.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/color_triplet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/columns.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/console.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/constrain.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/containers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/control.py
```
import time
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/default_styles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/diagnose.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/emoji.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/file_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/filesize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/highlighter.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/jupyter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/layout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/live.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/live_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/logging.py
```
import logging
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/markdown.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/markup.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/measure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/pager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/palette.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/panel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/pretty.py
```
import builtins
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/progress_bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/prompt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/region.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/repr.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/rule.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/screen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/segment.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/spinner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/status.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/style.py
```
import sys
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/styled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/syntax.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/terminal_theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/text.py
```
import re
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/theme.py
```
import configparser
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/themes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/traceback.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich-14.2.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich-14.2.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich-14.2.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rich-14.2.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds_py-0.30.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds_py-0.30.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds_py-0.30.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/rpds_py-0.30.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/__init__.py
```
import importlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/nt.py
```
import contextlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham-1.5.4.dist-info/zip-safe
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette/event.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette/sse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette-3.1.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette-3.1.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette-3.1.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette-3.1.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/sse_starlette-3.1.2.dist-info/top_level.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/_exception_handler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/applications.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/authentication.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/background.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/concurrency.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/convertors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/datastructures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/endpoints.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/formparsers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/requests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/responses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/routing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/schemas.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/staticfiles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/status.py
```
"""
HTTP codes
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/templating.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/testclient.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/websockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette-0.50.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette-0.50.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette-0.50.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette-0.50.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/_completion_classes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/_completion_shared.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/_typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/colors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/completion.py
```
import os
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/core.py
```
import errno
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/main.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/models.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/params.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/rich_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer/utils.py
```
import inspect
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer-0.21.1.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer-0.21.1.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer-0.21.1.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer-0.21.1.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typer-0.21.1.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_extensions-4.15.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_extensions-4.15.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_extensions-4.15.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_extensions-4.15.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection/introspection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection/typing_objects.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection/typing_objects.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection-0.4.2.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection-0.4.2.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection-0.4.2.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/typing_inspection-0.4.2.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/__main__.py
```
import uvicorn

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/_subprocess.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/importer.py
```
import importlib
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/py.typed
```

```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/server.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/workers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn-0.40.0.dist-info/INSTALLER
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn-0.40.0.dist-info/METADATA
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn-0.40.0.dist-info/RECORD
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn-0.40.0.dist-info/WHEEL
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn-0.40.0.dist-info/entry_points.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_code/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_code/code.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_code/source.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_io/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_io/pprint.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_io/saferepr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_io/terminalwriter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_io/wcwidth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_py/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_py/error.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/_py/path.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/assertion/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/assertion/rewrite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/assertion/truncate.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/assertion/util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/config/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/config/argparsing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/config/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/config/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/config/findpaths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/mark/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/mark/expression.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/_pytest/mark/structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_backends/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_backends/_asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_backends/_trio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_asyncio_selector_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_contextmanagers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_eventloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_fileio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_signals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_sockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_streams.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_subprocesses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_synchronization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_tasks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_tempfile.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/_core/_typedattr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_eventloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_sockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_streams.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_subprocesses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_tasks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/abc/_testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/buffered.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/memory.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/stapled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/anyio/streams/tls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/_oid.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/certificate_transparency.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/extensions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/general_name.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/name.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/ocsp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/oid.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/x509/verification.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/connection_pool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/http11.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/http2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/http_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/interfaces.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_async/socks_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/anyio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/mock.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/sync.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_backends/trio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/connection_pool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/http11.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/http2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/http_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/interfaces.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpcore/_sync/socks_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/asgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/default.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/mock.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/httpx/_transports/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/const_vs_enum.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/contains.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/import_benchmark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/issue232.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/json_schema_test_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/nested_schemas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/subcomponents.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/unused_registry.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/useless_applicator_schemas.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/useless_keywords.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/validator_creation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/_suite.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/fuzz_validate.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_deprecations.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_format.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_jsonschema_test_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/test_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/tests/test_jsonschema_specifications.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/cli/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/cli/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/entities.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/html_blocks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/html_re.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/normalize_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/common/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/helpers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/helpers/parse_link_destination.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/helpers/parse_link_label.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/helpers/parse_link_title.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/presets/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/presets/commonmark.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/presets/default.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/presets/zero.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/blockquote.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/code.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/fence.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/heading.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/hr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/html_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/lheading.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/list.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/paragraph.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/reference.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/state_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_block/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/linkify.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/normalize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/replacements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/smartquotes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/state_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_core/text_join.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/autolink.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/backticks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/balance_pairs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/emphasis.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/entity.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/escape.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/fragments_join.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/html_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/image.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/linkify.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/newline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/state_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/strikethrough.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/markdown_it/rules_inline/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/cli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/cli/claude.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/cli/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/session_group.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/sse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/streamable_http.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/websocket.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/os/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/elicitation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/models.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/session.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/sse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/stdio.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/streamable_http.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/streamable_http_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/transport_security.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/validation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/websocket.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/_httpx_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/auth_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/context.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/memory.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/message.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/metadata_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/response_router.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/tool_name_validation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/build_env.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/configuration.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/pyproject.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/self_outdated_check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/wheel_builder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/six.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/typing_extensions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/vendor.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ply/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ply/cpp.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ply/lex.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ply/yacc.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pycparser/ply/ygen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_core_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_core_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_dataclasses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_decorators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_decorators_v1.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_discriminated_union.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_docs_extraction.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_forward_ref.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_generate_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_git.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_import_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_internal_dataclass.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_known_annotated_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_mock_val_ser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_model_construction.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_namespace_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_repr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_schema_gather.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_schema_generation_shared.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_serializers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_signature.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_typing_extra.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_validate_call.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/_internal/_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/copy_internals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/deprecated/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/experimental/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/experimental/arguments_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/experimental/missing_sentinel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/experimental/pipeline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/plugin/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/plugin/_loader.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/plugin/_schema_validator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/_hypothesis_plugin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/annotated_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/color.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/dataclasses.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/datetime_parse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/env_settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/error_wrappers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/mypy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/networks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/py.typed
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic/v1/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/filters/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/bbcode.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/groff.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/html.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/img.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/irc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/latex.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/other.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/pangomarkup.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/rtf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/svg.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/terminal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/formatters/terminal256.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/abap.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/algol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/algol_nu.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/arduino.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/autumn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/borland.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/bw.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/coffee.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/colorful.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/default.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/dracula.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/emacs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/friendly.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/friendly_grayscale.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/fruity.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/gh_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/gruvbox.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/igor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/inkpot.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/lightbulb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/lilypond.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/lovelace.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/manni.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/material.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/monokai.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/murphy.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/native.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/nord.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/onedark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/paraiso_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/paraiso_light.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/pastie.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/perldoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/rainbow_dash.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/rrt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/sas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/solarized.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/staroffice.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/stata_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/stata_light.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/tango.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/trac.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/vim.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/vs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/xcode.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/styles/zenburn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_ada_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_asy_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_cl_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_cocoa_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_csound_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_css_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_googlesql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_julia_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_lasso_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_lilypond_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_lua_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_luau_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_mql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_mysql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_openedge_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_php_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_postgres_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_qlik_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_scheme_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_scilab_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_sourcemod_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_sql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_stan_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_stata_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_tsql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_usd_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_vbscript_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/_vim_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/actionscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ada.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/agile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/algebra.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ambient.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/amdgpu.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ampl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/apdlexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/apl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/archetype.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/arrow.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/arturo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/asc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/asm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/asn1.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/automation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/bare.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/basic.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/bdd.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/berry.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/bibtex.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/blueprint.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/boa.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/bqn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/business.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/c_cpp.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/c_like.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/capnproto.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/carbon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/cddl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/chapel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/clean.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/codeql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/comal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/compiled.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/configs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/console.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/cplint.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/crystal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/csound.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/css.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/d.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dalvik.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/data.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dax.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/devicetree.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/diff.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dns.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dotnet.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dsls.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/dylan.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ecl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/eiffel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/elm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/elpi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/email.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/erlang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/esoteric.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ezhil.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/factor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/fantom.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/felix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/fift.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/floscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/forth.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/fortran.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/foxpro.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/freefem.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/func.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/functional.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/futhark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/gcodelexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/gdscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/gleam.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/go.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/grammar_notation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/graph.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/graphics.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/graphql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/graphviz.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/gsql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/hare.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/haskell.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/haxe.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/hdl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/hexdump.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/html.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/idl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/igor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/inferno.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/installers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/int_fiction.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/iolang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/j.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/javascript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/jmespath.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/jslt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/json5.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/jsonnet.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/jsx.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/julia.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/jvm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/kuin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/kusto.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ldap.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/lean.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/lilypond.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/lisp.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/macaulay2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/make.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/maple.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/markup.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/math.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/matlab.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/maxima.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/meson.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/mime.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/minecraft.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/mips.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ml.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/modeling.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/modula2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/mojo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/monte.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/mosel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ncl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/nimrod.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/nit.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/nix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/numbair.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/oberon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/objective.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ooc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/openscad.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/other.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/parasail.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/parsers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/pascal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/pawn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/pddl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/perl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/phix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/php.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/pointless.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/pony.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/praat.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/procfile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/prolog.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/promql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/prql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ptx.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/python.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/q.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/qlik.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/qvt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/r.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rdf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rebol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rego.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/resource.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ride.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rita.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rnc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/roboconf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/robotframework.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ruby.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/rust.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/sas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/savi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/scdoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/scripting.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/sgf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/shell.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/sieve.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/slash.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/smalltalk.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/smithy.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/smv.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/snobol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/solidity.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/soong.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/sophia.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/special.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/spice.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/sql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/srcinfo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/stata.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/supercollider.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tablegen.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tact.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tcl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/teal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/templates.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/teraterm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/testing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/text.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/textedit.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/textfmts.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/theorem.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/thingsdb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tlb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tls.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/tnt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/trafficscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/typoscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/typst.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/ul4.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/unicon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/urbi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/usd.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/varnish.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/verification.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/verifpal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/vip.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/vyper.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/web.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/webassembly.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/webidl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/webmisc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/wgsl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/whiley.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/wowtoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/wren.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/x10.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/xorg.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/yang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/yara.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pygments/lexers/zig.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/test_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/test_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/test_jsonschema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/test_referencing_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/referencing/tests/test_retrieval.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/posix/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/posix/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/posix/proc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/shellingham/posix/ps.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/authentication.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/cors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/gzip.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/httpsredirect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/sessions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/trustedhost.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/starlette/middleware/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/lifespan/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/lifespan/off.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/lifespan/on.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/loops/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/loops/asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/loops/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/loops/uvloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/middleware/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/middleware/asgi2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/middleware/message_logger.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/middleware/proxy_headers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/middleware/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/supervisors/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/supervisors/basereload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/supervisors/multiprocess.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/supervisors/statreload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/supervisors/watchfilesreload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_code/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_code/code.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_code/source.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_io/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_io/pprint.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_io/saferepr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_io/terminalwriter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_io/wcwidth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_py/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_py/error.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/_py/path.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/assertion/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/assertion/rewrite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/assertion/truncate.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/assertion/util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/config/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/config/argparsing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/config/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/config/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/config/findpaths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/mark/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/mark/expression.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/_pytest/mark/structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_backends/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_backends/_asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_backends/_trio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_asyncio_selector_thread.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_contextmanagers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_eventloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_fileio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_signals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_sockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_streams.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_subprocesses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_synchronization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_tasks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_tempfile.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/_core/_typedattr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_eventloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_sockets.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_streams.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_subprocesses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_tasks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/abc/_testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/buffered.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/memory.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/stapled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/anyio/streams/tls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/_oid.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/certificate_transparency.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/extensions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/general_name.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/name.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/ocsp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/oid.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/x509/verification.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/connection_pool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/http11.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/http2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/http_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/interfaces.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_async/socks_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/anyio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/mock.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/sync.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_backends/trio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/connection_pool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/http11.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/http2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/http_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/interfaces.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpcore/_sync/socks_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/asgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/default.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/mock.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/httpx/_transports/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/const_vs_enum.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/contains.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/import_benchmark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/issue232.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/json_schema_test_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/nested_schemas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/subcomponents.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/unused_registry.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/useless_applicator_schemas.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/useless_keywords.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/validator_creation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/_suite.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/fuzz_validate.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_deprecations.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_format.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_jsonschema_test_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_types.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/test_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/tests/test_jsonschema_specifications.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/cli/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/cli/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/entities.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/html_blocks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/html_re.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/normalize_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/common/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/helpers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/helpers/parse_link_destination.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/helpers/parse_link_label.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/helpers/parse_link_title.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/presets/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/presets/commonmark.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/presets/default.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/presets/zero.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/blockquote.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/code.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/fence.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/heading.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/hr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/html_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/lheading.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/list.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/paragraph.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/reference.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/state_block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_block/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/block.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/linkify.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/normalize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/replacements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/smartquotes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/state_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_core/text_join.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/autolink.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/backticks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/balance_pairs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/emphasis.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/entity.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/escape.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/fragments_join.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/html_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/image.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/linkify.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/newline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/state_inline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/strikethrough.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/markdown_it/rules_inline/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/cli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/cli/claude.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/cli/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/session_group.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/sse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/streamable_http.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/websocket.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/os/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/elicitation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/models.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/session.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/sse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/stdio.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/streamable_http.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/streamable_http_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/transport_security.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/validation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/websocket.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/_httpx_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/auth_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/context.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/memory.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/message.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/metadata_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/response_router.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/tool_name_validation.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/build_env.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/configuration.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/pyproject.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/self_outdated_check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/wheel_builder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/six.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/typing_extensions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/vendor.txt
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ply/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ply/cpp.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ply/lex.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ply/yacc.py
```
# -----------------------------------------------------------------------------
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pycparser/ply/ygen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_core_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_core_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_dataclasses.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_decorators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_decorators_v1.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_discriminated_union.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_docs_extraction.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_forward_ref.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_generate_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_git.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_import_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_internal_dataclass.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_known_annotated_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_mock_val_ser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_model_construction.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_namespace_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_repr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_schema_gather.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_schema_generation_shared.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_serializers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_signature.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_typing_extra.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_validate_call.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/_internal/_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/copy_internals.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/deprecated/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/experimental/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/experimental/arguments_schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/experimental/missing_sentinel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/experimental/pipeline.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/plugin/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/plugin/_loader.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/plugin/_schema_validator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/_hypothesis_plugin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/annotated_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/class_validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/color.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/config.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/dataclasses.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/datetime_parse.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/decorator.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/env_settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/error_wrappers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/generics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/mypy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/networks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/parse.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/py.typed
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/schema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/tools.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/typing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/validators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic/v1/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/filters/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/bbcode.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/groff.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/html.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/img.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/irc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/latex.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/other.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/pangomarkup.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/rtf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/svg.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/terminal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/formatters/terminal256.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/abap.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/algol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/algol_nu.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/arduino.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/autumn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/borland.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/bw.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/coffee.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/colorful.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/default.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/dracula.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/emacs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/friendly.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/friendly_grayscale.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/fruity.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/gh_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/gruvbox.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/igor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/inkpot.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/lightbulb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/lilypond.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/lovelace.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/manni.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/material.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/monokai.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/murphy.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/native.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/nord.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/onedark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/paraiso_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/paraiso_light.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/pastie.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/perldoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/rainbow_dash.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/rrt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/sas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/solarized.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/staroffice.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/stata_dark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/stata_light.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/tango.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/trac.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/vim.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/vs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/xcode.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/styles/zenburn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_ada_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_asy_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_cl_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_cocoa_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_csound_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_css_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_googlesql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_julia_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_lasso_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_lilypond_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_lua_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_luau_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_mql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_mysql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_openedge_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_php_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_postgres_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_qlik_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_scheme_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_scilab_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_sourcemod_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_sql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_stan_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_stata_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_tsql_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_usd_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_vbscript_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/_vim_builtins.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/actionscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ada.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/agile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/algebra.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ambient.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/amdgpu.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ampl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/apdlexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/apl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/archetype.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/arrow.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/arturo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/asc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/asm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/asn1.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/automation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/bare.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/basic.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/bdd.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/berry.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/bibtex.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/blueprint.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/boa.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/bqn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/business.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/c_cpp.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/c_like.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/capnproto.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/carbon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/cddl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/chapel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/clean.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/codeql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/comal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/compiled.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/configs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/console.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/cplint.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/crystal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/csound.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/css.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/d.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dalvik.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/data.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dax.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/devicetree.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/diff.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dns.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dotnet.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dsls.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/dylan.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ecl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/eiffel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/elm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/elpi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/email.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/erlang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/esoteric.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ezhil.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/factor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/fantom.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/felix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/fift.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/floscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/forth.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/fortran.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/foxpro.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/freefem.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/func.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/functional.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/futhark.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/gcodelexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/gdscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/gleam.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/go.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/grammar_notation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/graph.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/graphics.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/graphql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/graphviz.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/gsql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/hare.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/haskell.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/haxe.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/hdl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/hexdump.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/html.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/idl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/igor.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/inferno.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/installers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/int_fiction.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/iolang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/j.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/javascript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/jmespath.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/jslt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/json5.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/jsonnet.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/jsx.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/julia.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/jvm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/kuin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/kusto.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ldap.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/lean.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/lilypond.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/lisp.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/macaulay2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/make.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/maple.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/markup.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/math.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/matlab.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/maxima.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/meson.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/mime.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/minecraft.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/mips.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ml.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/modeling.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/modula2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/mojo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/monte.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/mosel.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ncl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/nimrod.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/nit.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/nix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/numbair.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/oberon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/objective.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ooc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/openscad.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/other.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/parasail.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/parsers.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/pascal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/pawn.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/pddl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/perl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/phix.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/php.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/pointless.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/pony.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/praat.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/procfile.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/prolog.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/promql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/prql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ptx.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/python.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/q.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/qlik.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/qvt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/r.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rdf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rebol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rego.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/resource.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ride.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rita.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rnc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/roboconf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/robotframework.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ruby.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/rust.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/sas.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/savi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/scdoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/scripting.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/sgf.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/shell.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/sieve.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/slash.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/smalltalk.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/smithy.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/smv.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/snobol.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/solidity.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/soong.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/sophia.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/special.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/spice.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/sql.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/srcinfo.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/stata.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/supercollider.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tablegen.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tact.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tcl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/teal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/templates.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/teraterm.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/testing.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/text.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/textedit.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/textfmts.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/theorem.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/thingsdb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tlb.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tls.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/tnt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/trafficscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/typoscript.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/typst.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/ul4.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/unicon.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/urbi.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/usd.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/varnish.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/verification.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/verifpal.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/vip.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/vyper.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/web.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/webassembly.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/webidl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/webmisc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/wgsl.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/whiley.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/wowtoc.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/wren.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/x10.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/xorg.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/yang.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/yara.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pygments/lexers/zig.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/test_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/test_exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/test_jsonschema.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/test_referencing_suite.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/referencing/tests/test_retrieval.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/posix/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/posix/_core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/posix/proc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/shellingham/posix/ps.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/authentication.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/cors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/gzip.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/httpsredirect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/sessions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/trustedhost.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/starlette/middleware/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/lifespan/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/lifespan/off.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/lifespan/on.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/loops/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/loops/asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/loops/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/loops/uvloop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/middleware/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/middleware/asgi2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/middleware/message_logger.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/middleware/proxy_headers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/middleware/wsgi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/supervisors/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/supervisors/basereload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/supervisors/multiprocess.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/supervisors/statreload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/supervisors/watchfilesreload.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/asn1/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/asn1/asn1.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/backends/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/decrepit/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/_asymmetric.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/_cipheralgorithm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/_serialization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/cmac.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/constant_time.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/hashes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/hmac.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/keywrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/poly1305.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/benchmarks/issue232/issue.json
```
[
    {
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/typing/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema/tests/typing/test_all_concrete_validators_match_protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft3/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft4/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft6/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft7/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/auth/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/auth/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/auth/oauth2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/auth/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/experimental/task_handlers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/experimental/tasks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/stdio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/os/posix/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/os/posix/utilities.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/os/win32/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/os/win32/utilities.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/json_response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/provider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/routes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/request_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/session_features.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/task_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/task_result_handler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/experimental/task_support.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/server.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/lowlevel/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/lowlevel/experimental.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/lowlevel/func_inspection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/lowlevel/helper_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/lowlevel/server.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/autocompletion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/base_command.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/cmdoptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/command_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/main_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/progress_bars.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/req_command.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/spinners.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/cli/status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/completion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/configuration.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/debug.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/download.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/freeze.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/hash.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/help.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/inspect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/install.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/list.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/search.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/show.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/uninstall.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/commands/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/distributions/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/distributions/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/distributions/installed.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/distributions/sdist.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/distributions/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/index/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/index/collector.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/index/package_finder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/index/sources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/locations/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/locations/_distutils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/locations/_sysconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/locations/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/_json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/pkg_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/candidate.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/direct_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/format_control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/installation_report.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/scheme.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/search_scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/selection_prefs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/target_python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/models/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/download.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/lazy_wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/network/xmlrpc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/freeze.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/prepare.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/constructors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/req_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/req_install.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/req_set.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/req/req_uninstall.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/_jaraco_text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/_log.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/appdirs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/compatibility_tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/datetime.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/deprecation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/direct_url_helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/egg_link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/encoding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/entrypoints.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/filesystem.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/filetypes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/glibc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/hashes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/misc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/packaging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/setuptools_build.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/subprocess.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/temp_dir.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/unpacking.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/urls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/virtualenv.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/utils/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/bazaar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/git.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/mercurial.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/subversion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/vcs/versioncontrol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/_cmd.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/adapter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/controller.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/filewrapper.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/heuristics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/serialize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/wrapper.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/certifi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/certifi/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/certifi/core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/big5freq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/big5prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/chardistribution.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/charsetgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/charsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/codingstatemachine.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/codingstatemachinedict.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/cp949prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/enums.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/escprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/escsm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/eucjpprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/euckrfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/euckrprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/euctwfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/euctwprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/gb2312freq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/gb2312prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/hebrewprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/jisfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/johabfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/johabprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/jpcntx.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langbulgarianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langgreekmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langhebrewmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langhungarianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langrussianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langthaimodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/langturkishmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/latin1prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/macromanprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/mbcharsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/mbcsgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/mbcssm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/resultdict.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/sbcharsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/sbcsgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/sjisprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/universaldetector.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/utf1632prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/utf8prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/version.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/ansi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/ansitowin32.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/initialise.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/win32.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/winterm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/database.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/locators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/manifest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/scripts.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/util.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distlib/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distro/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distro/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/distro/distro.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/codec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/idnadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/intranges.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/package_data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/idna/uts46data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pkg_resources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/msgpack/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/msgpack/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/msgpack/ext.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/msgpack/fallback.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/__about__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/_manylinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/_musllinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/_structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/specifiers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/packaging/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/android.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/macos.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/unix.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/platformdirs/windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/__main__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/cmdline.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/console.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/filter.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatter.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/lexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/modeline.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/plugin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/regexopt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/scanner.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/sphinxext.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/style.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/unistring.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/util.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/actions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/common.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/core.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/results.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/unicode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyproject_hooks/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyproject_hooks/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyproject_hooks/_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/__init__.py
```
#   __
#  /__)  _  _     _   _ _/   _
# / (   (- (/ (/ (- _)  /  _)
#          /

"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/__version__.py
```
# .-. .-. .-. . . .-. .-. .-. .-.
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/_internal_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/adapters.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/api.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/auth.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/certs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/compat.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/cookies.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/help.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/hooks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/models.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/packages.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/sessions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/structures.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/requests/utils.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/providers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/reporters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/resolvers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/structs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_cell_widths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_emoji_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_emoji_replace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_export_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_extension.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_fileno.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_inspect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_log_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_loop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_null_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_palettes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_pick.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_ratio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_spinners.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_stack.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_timer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_win32_console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_windows_renderer.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/_wrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/align.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/ansi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/box.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/cells.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/color.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/color_triplet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/columns.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/constrain.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/containers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/default_styles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/diagnose.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/emoji.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/file_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/filesize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/highlighter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/jupyter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/layout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/live.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/live_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/markup.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/measure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/pager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/palette.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/panel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/pretty.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/progress_bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/prompt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/region.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/repr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/rule.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/screen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/segment.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/spinner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/status.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/style.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/styled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/syntax.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/terminal_theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/themes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/traceback.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/rich/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/_asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/after.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/before.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/before_sleep.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/nap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/retry.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/stop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/tornadoweb.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tenacity/wait.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tomli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tomli/_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tomli/_re.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/tomli/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/_macos.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/_openssl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/_ssl_constants.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/truststore/_windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/_collections.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/connectionpool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/filepost.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/poolmanager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/request.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/webencodings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/webencodings/labels.py
```
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/webencodings/mklabels.py
```
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/webencodings/tests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/webencodings/x_user_defined.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/aws.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/azure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/dotenv.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/gcp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/pyproject.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/toml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pydantic_settings/sources/providers/yaml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/http/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/http/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/http/flow_control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/http/h11_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/http/httptools_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/websockets/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/websockets/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/websockets/websockets_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/websockets/websockets_sansio_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/uvicorn/protocols/websockets/wsproto_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/asn1/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/asn1/asn1.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/backends/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/decrepit/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/_asymmetric.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/_cipheralgorithm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/_serialization.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/cmac.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/constant_time.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/hashes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/hmac.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/keywrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/poly1305.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/benchmarks/issue232/issue.json
```
[
    {
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/typing/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema/tests/typing/test_all_concrete_validators_match_protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft3/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft4/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft6/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft7/metaschema.json
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/auth/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/auth/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/auth/oauth2.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/auth/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/experimental/task_handlers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/experimental/tasks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/stdio/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/os/posix/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/os/posix/utilities.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/os/win32/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/os/win32/utilities.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/json_response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/provider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/routes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/settings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/request_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/session_features.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/task_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/task_result_handler.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/experimental/task_support.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/server.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/lowlevel/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/lowlevel/experimental.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/lowlevel/func_inspection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/lowlevel/helper_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/lowlevel/server.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/autocompletion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/base_command.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/cmdoptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/command_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/main.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/main_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/progress_bars.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/req_command.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/spinners.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/cli/status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/completion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/configuration.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/debug.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/download.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/freeze.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/hash.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/help.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/inspect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/install.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/list.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/search.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/show.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/uninstall.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/commands/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/distributions/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/distributions/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/distributions/installed.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/distributions/sdist.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/distributions/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/index/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/index/collector.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/index/package_finder.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/index/sources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/locations/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/locations/_distutils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/locations/_sysconfig.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/locations/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/_json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/pkg_resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/candidate.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/direct_url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/format_control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/installation_report.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/scheme.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/search_scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/selection_prefs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/target_python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/models/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/download.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/lazy_wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/session.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/network/xmlrpc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/check.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/freeze.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/prepare.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/constructors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/req_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/req_install.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/req_set.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/req/req_uninstall.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/_jaraco_text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/_log.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/appdirs.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/compatibility_tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/datetime.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/deprecation.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/direct_url_helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/egg_link.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/encoding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/entrypoints.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/filesystem.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/filetypes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/glibc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/hashes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/misc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/models.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/packaging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/setuptools_build.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/subprocess.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/temp_dir.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/unpacking.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/urls.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/virtualenv.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/utils/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/bazaar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/git.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/mercurial.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/subversion.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/vcs/versioncontrol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/_cmd.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/adapter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/controller.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/filewrapper.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/heuristics.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/serialize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/wrapper.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/certifi/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/certifi/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/certifi/core.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/big5freq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/big5prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/chardistribution.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/charsetgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/charsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/codingstatemachine.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/codingstatemachinedict.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/cp949prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/enums.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/escprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/escsm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/eucjpprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/euckrfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/euckrprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/euctwfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/euctwprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/gb2312freq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/gb2312prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/hebrewprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/jisfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/johabfreq.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/johabprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/jpcntx.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langbulgarianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langgreekmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langhebrewmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langhungarianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langrussianmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langthaimodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/langturkishmodel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/latin1prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/macromanprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/mbcharsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/mbcsgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/mbcssm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/resultdict.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/sbcharsetprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/sbcsgroupprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/sjisprober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/universaldetector.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/utf1632prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/utf8prober.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/version.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/ansi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/ansitowin32.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/initialise.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/win32.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/winterm.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/database.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/index.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/locators.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/manifest.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/resources.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/scripts.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/util.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distlib/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distro/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distro/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/distro/distro.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/codec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/core.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/idnadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/intranges.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/package_data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/idna/uts46data.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/msgpack/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/msgpack/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/msgpack/ext.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/msgpack/fallback.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/__about__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/_manylinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/_musllinux.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/_structures.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/markers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/specifiers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/tags.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/packaging/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pkg_resources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/android.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/macos.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/unix.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/platformdirs/windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/__main__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/cmdline.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/console.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/filter.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatter.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/lexer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/modeline.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/plugin.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/regexopt.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/scanner.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/sphinxext.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/style.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/unistring.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/util.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/actions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/common.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/core.py
```
#
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/results.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/testing.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/unicode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/util.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyproject_hooks/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyproject_hooks/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyproject_hooks/_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/__init__.py
```
#   __
#  /__)  _  _     _   _ _/   _
# / (   (- (/ (/ (- _)  /  _)
#          /

"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/__version__.py
```
# .-. .-. .-. . . .-. .-. .-. .-.
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/_internal_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/adapters.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/api.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/auth.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/certs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/compat.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/cookies.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/exceptions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/help.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/hooks.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/models.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/packages.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/sessions.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/status_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/structures.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/requests/utils.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/providers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/reporters.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/resolvers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/structs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/__main__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_cell_widths.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_emoji_codes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_emoji_replace.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_export_format.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_extension.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_fileno.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_inspect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_log_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_loop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_null_file.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_palettes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_pick.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_ratio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_spinners.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_stack.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_timer.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_win32_console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_windows_renderer.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/_wrap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/align.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/ansi.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/box.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/cells.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/color.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/color_triplet.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/columns.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/console.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/constrain.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/containers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/default_styles.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/diagnose.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/emoji.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/errors.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/file_proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/filesize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/highlighter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/jupyter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/layout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/live.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/live_render.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/markup.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/measure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/pager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/palette.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/panel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/pretty.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/progress.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/progress_bar.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/prompt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/protocol.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/region.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/repr.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/rule.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/scope.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/screen.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/segment.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/spinner.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/status.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/style.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/styled.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/syntax.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/table.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/terminal_theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/text.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/theme.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/themes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/traceback.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/rich/tree.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/_asyncio.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/_utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/after.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/before.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/before_sleep.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/nap.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/retry.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/stop.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/tornadoweb.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tenacity/wait.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tomli/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tomli/_parser.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tomli/_re.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/tomli/_types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/_api.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/_macos.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/_openssl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/_ssl_constants.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/truststore/_windows.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/__init__.py
```
"""
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/_collections.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/_version.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/connectionpool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/exceptions.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/fields.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/filepost.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/poolmanager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/request.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/webencodings/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/webencodings/labels.py
```
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/webencodings/mklabels.py
```
"""

[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/webencodings/tests.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/webencodings/x_user_defined.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/aws.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/azure.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/cli.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/dotenv.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/gcp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/json.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/pyproject.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/toml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pydantic_settings/sources/providers/yaml.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/http/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/http/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/http/flow_control.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/http/h11_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/http/httptools_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/websockets/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/websockets/auto.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/websockets/websockets_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/websockets/websockets_sansio_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/uvicorn/protocols/websockets/wsproto_impl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/backends/openssl/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/backends/openssl/backend.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/_openssl.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/asn1.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/declarative_asn1.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/exceptions.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/ocsp.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/pkcs12.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/pkcs7.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/test_support.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/x509.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/openssl/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/openssl/_conditional.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/openssl/binding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/decrepit/ciphers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/decrepit/ciphers/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/dh.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/dsa.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ed25519.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ed448.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/rsa.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/x25519.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/x448.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/aead.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/modes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/argon2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/concatkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/hkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/kbkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/pbkdf2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/scrypt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/kdf/x963kdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/serialization/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/serialization/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/serialization/pkcs12.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/serialization/pkcs7.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/serialization/ssh.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/hotp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/totp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/applicator
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/content
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/core
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/format
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/meta-data
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/validation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/applicator
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/content
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/core
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/format-annotation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/format-assertion
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/meta-data
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/unevaluated
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/validation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/client/auth/extensions/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/handlers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/handlers/authorize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/handlers/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/handlers/register.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/handlers/revoke.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/middleware/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/middleware/auth_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/middleware/bearer_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/auth/middleware/client_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/prompts/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/prompts/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/prompts/manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/resources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/resources/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/resources/resource_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/resources/templates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/resources/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/tools/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/tools/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/tools/tool_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/utilities/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/utilities/context_injection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/utilities/func_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/utilities/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/server/fastmcp/utilities/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/capabilities.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/in_memory_task_store.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/message_queue.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/polling.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/mcp/shared/experimental/tasks/store.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/importlib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/importlib/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/importlib/_dists.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/metadata/importlib/_envs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/install/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/install/editable_legacy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/operations/install/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/legacy/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/legacy/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/candidates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/factory.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/found_candidates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/provider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/reporter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_internal/resolution/resolvelib/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/caches/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/caches/file_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/cachecontrol/caches/redis_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/cli/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/cli/chardetect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/metadata/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/chardet/metadata/languages.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/ansi_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/ansitowin32_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/initialise_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/isatty_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/colorama/tests/winterm_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/filters/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/bbcode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/groff.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/html.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/img.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/irc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/latex.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/other.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/pangomarkup.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/rtf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/svg.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/terminal.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/formatters/terminal256.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/lexers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/lexers/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/lexers/python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pygments/styles/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyparsing/diagram/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyproject_hooks/_in_process/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/pyproject_hooks/_in_process/_in_process.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/compat/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/resolvelib/compat/collections_abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/_appengine_environ.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/appengine.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/ntlmpool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/pyopenssl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/securetransport.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/socks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/packages/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/packages/six.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/queue.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/request.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/retry.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/ssl_.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/ssl_match_hostname.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/ssltransport.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/timeout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/util/wait.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/backends/openssl/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/backends/openssl/backend.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/_openssl.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/asn1.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/declarative_asn1.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/exceptions.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/ocsp.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/pkcs12.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/pkcs7.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/test_support.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/x509.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/openssl/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/openssl/_conditional.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/openssl/binding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/decrepit/ciphers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/decrepit/ciphers/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/dh.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/dsa.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ec.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ed25519.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/ed448.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/padding.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/rsa.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/x25519.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/asymmetric/x448.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/aead.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/algorithms.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/ciphers/modes.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/argon2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/concatkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/hkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/kbkdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/pbkdf2.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/scrypt.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/kdf/x963kdf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/serialization/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/serialization/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/serialization/pkcs12.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/serialization/pkcs7.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/serialization/ssh.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/hotp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/primitives/twofactor/totp.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/applicator
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/content
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/core
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/format
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/meta-data
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft201909/vocabularies/validation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/applicator
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/content
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/core
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/format-annotation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/format-assertion
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/meta-data
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/unevaluated
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/jsonschema_specifications/schemas/draft202012/vocabularies/validation
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/client/auth/extensions/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/handlers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/handlers/authorize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/handlers/metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/handlers/register.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/handlers/revoke.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/middleware/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/middleware/auth_context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/middleware/bearer_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/auth/middleware/client_auth.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/prompts/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/prompts/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/prompts/manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/resources/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/resources/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/resources/resource_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/resources/templates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/resources/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/tools/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/tools/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/tools/tool_manager.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/utilities/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/utilities/context_injection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/utilities/func_metadata.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/utilities/logging.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/server/fastmcp/utilities/types.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/capabilities.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/context.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/helpers.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/in_memory_task_store.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/message_queue.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/polling.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/mcp/shared/experimental/tasks/store.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/importlib/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/importlib/_compat.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/importlib/_dists.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/metadata/importlib/_envs.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/install/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/install/editable_legacy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/operations/install/wheel.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/legacy/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/legacy/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/base.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/candidates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/factory.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/found_candidates.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/provider.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/reporter.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/requirements.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_internal/resolution/resolvelib/resolver.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/caches/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/caches/file_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/cachecontrol/caches/redis_cache.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/cli/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/cli/chardetect.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/metadata/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/chardet/metadata/languages.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/ansi_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/ansitowin32_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/initialise_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/isatty_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/utils.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/colorama/tests/winterm_test.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/filters/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/bbcode.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/groff.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/html.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/img.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/irc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/latex.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/other.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/pangomarkup.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/rtf.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/svg.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/terminal.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/formatters/terminal256.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/lexers/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/lexers/_mapping.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/lexers/python.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pygments/styles/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyparsing/diagram/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyproject_hooks/_in_process/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/pyproject_hooks/_in_process/_in_process.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/compat/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/resolvelib/compat/collections_abc.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/_appengine_environ.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/appengine.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/ntlmpool.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/pyopenssl.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/securetransport.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/socks.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/packages/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/packages/six.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/__init__.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/connection.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/proxy.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/queue.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/request.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/response.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/retry.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/ssl_.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/ssl_match_hostname.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/ssltransport.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/timeout.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/url.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/util/wait.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/aead.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ciphers.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/cmac.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/dh.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/dsa.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ec.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ed25519.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ed448.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/hashes.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/hmac.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/kdf.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/keys.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/poly1305.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/rsa.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/x25519.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/x448.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/bindings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/low_level.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/__init__.py
```
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/makefile.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/weakref_finalize.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/__init__.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/aead.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ciphers.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/cmac.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/dh.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/dsa.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ec.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ed25519.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/ed448.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/hashes.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/hmac.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/kdf.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/keys.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/poly1305.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/rsa.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/x25519.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/cryptography/hazmat/bindings/_rust/openssl/x448.pyi
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/bindings.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/contrib/_securetransport/low_level.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/__init__.py
```
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/makefile.py
```
[TRUNCATED]
```

oraclepack-mcp-server/venv/lib64/python3.12/site-packages/pip/_vendor/urllib3/packages/backports/weakref_finalize.py
```
[TRUNCATED]
```

</source_code>