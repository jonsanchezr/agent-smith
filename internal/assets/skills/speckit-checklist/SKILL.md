---
name: speckit-checklist
description: "Trigger: speckit-checklist, checklist, requirements quality, quality check. Generate requirements-quality checklists using the spec-kit checklist workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Generate requirements-quality checklists using the spec-kit checklist workflow.

## Commands

 Delegates to: `specify checklist`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify checklist` command.

## Execution

Run `specify checklist` to invoke the spec-kit checklist workflow.
