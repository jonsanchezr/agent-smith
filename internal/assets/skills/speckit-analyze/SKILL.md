---
name: speckit-analyze
description: "Trigger: speckit-analyze, analyze, consistency, artifact consistency. Check artifact consistency using the spec-kit analyze workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Check artifact consistency using the spec-kit analyze workflow.

## Commands

 Delegates to: `specify analyze`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify analyze` command.

## Execution

Run `specify analyze` to invoke the spec-kit analyze workflow.
