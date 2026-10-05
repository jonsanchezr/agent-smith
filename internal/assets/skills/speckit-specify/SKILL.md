---
name: speckit-specify
description: "Trigger: speckit-specify, requirements, user stories, specify. Define requirements and user stories using the spec-kit specify workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Define requirements and user stories using the spec-kit specify workflow.

## Commands

 Delegates to: `specify specify`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify specify` command.

## Execution

Run `specify specify` to invoke the spec-kit specify workflow.
