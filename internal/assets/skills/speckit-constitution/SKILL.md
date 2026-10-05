---
name: speckit-constitution
description: "Trigger: speckit-constitution, project constitution, establish principles. Establish or update project principles using the spec-kit constitution workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Establish or update project principles using the spec-kit constitution workflow.

## Commands

 Delegates to: `specify constitution`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify constitution` command.

## Execution

Run `specify constitution` to invoke the spec-kit constitution workflow.
