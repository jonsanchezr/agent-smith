---
name: speckit-plan
description: "Trigger: speckit-plan, planning, implementation plan, technical plan. Create the technical implementation plan using the spec-kit plan workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Create the technical implementation plan using the spec-kit plan workflow.

## Commands

 Delegates to: `specify plan`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify plan` command.

## Execution

Run `specify plan` to invoke the spec-kit plan workflow.
