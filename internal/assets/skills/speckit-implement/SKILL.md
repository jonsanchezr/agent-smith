---
name: speckit-implement
description: "Trigger: speckit-implement, implement, execute, run tasks. Execute the tasks using the spec-kit implement workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Execute the tasks using the spec-kit implement workflow.

## Commands

 Delegates to: `specify implement`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify implement` command.

## Execution

Run `specify implement` to invoke the spec-kit implement workflow.
