---
name: speckit-tasks
description: "Trigger: speckit-tasks, tasks, break down, actionable tasks. Break the plan into actionable tasks using the spec-kit tasks workflow."
license: MIT
metadata:
  author: jonsanchezr
  version: "1.0"
---

## What

Break the plan into actionable tasks using the spec-kit tasks workflow.

## Commands

 Delegates to: `specify tasks`

## Auto-install

Before delegating, check if the specify CLI is available:

```
speckit_available()
```

If it returns `(false, _, _, _)`, run:

```
speckit_install()
```

Then retry the `specify tasks` command.

## Execution

Run `specify tasks` to invoke the spec-kit tasks workflow.
