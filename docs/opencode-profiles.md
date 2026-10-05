# Native OpenCode background subagents

> [!NOTE]
> These docs track `main`, which may include unreleased changes. For the latest release, see the [v4.0.0 docs](https://github.com/jonsanchezr/agent-smith/tree/v4.0.0/docs).

â† [Back to usage](usage.md)

Agent Smith uses OpenCode's native subagents through its `task` permission. It does not install the legacy `background-agents.ts` plugin by default. Background execution is optional and independent of ODD model configuration.

## Choose an execution preference

Install and sync accept the same preference:

```bash
agent-smith install --agent opencode --opencode-background-subagents=on
agent-smith sync --agent opencode --opencode-background-subagents=off
```

Use `auto`, `on`, or `off`. You can also set `GENTLE_AI_OPENCODE_BACKGROUND_SUBAGENTS=auto|on|off`. The CLI flag takes precedence, then a non-empty environment variable, then the previous managed choice in Agent Smith state; the default is `auto`. An interactive installer may ask for a choice when no prior or explicit setting exists. Cancelling installation does not save a new preference.

When activation is enabled, Agent Smith manages launchers under `~/.agent-smith/bin/` (`opencode` on POSIX; `opencode.cmd` and `opencode.ps1` on Windows). The launcher sets `OPENCODE_EXPERIMENTAL_BACKGROUND_SUBAGENTS=true` only if the variable is unset: an explicit `false` keeps execution in the foreground. Restart OpenCode after activation, and restart the shell if PATH has not refreshed.

Sessions launched through `opencode serve`, `opencode attach`, or OpenCode Desktop may not inherit the managed launcher environment; they fall back to foreground execution. Agent Smith does not rewrite their configuration.

Background jobs are process-local and non-durable: restarting OpenCode loses them. They provide no filesystem isolation, so use them only for independent read-only work, not dependent tasks or parallel writers in one worktree.

For model selection and review roles, use the TUI **Configure Models** screen; see [Usage](usage.md#model-assignment).

