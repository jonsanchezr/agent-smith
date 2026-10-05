---
name: speckit-clarify
description: "Trigger: speckit-clarify, clarify, ambiguity, requirements clarity. Resolve ambiguity before planning using the spec-kit clarify workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Resolve ambiguity before planning using the spec-kit clarify workflow.

## Commands

 Delegates to: `specify clarify`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify clarify` command.

## Execution

Run `specify clarify` to invoke the spec-kit clarify workflow.
