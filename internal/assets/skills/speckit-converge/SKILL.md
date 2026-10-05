---
name: speckit-converge
description: "Trigger: speckit-converge, converge, assess, implementation vs spec. Assess implementation against spec using the spec-kit converge workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Assess implementation against spec using the spec-kit converge workflow.

## Commands

 Delegates to: `specify converge`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify converge` command.

## Execution

Run `specify converge` to invoke the spec-kit converge workflow.
