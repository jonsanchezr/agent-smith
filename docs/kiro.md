# Kiro IDE

> [!NOTE]
> These docs track `main`, which may include unreleased changes. For the latest release, see the [v4.0.0 docs](https://github.com/jonsanchezr/agent-smith/tree/v4.0.0/docs).

â† [Back to README](../README.md)

---

This document explains how agent-smith integrates with **Kiro IDE** and what is installed in your local Kiro configuration.

## Overview

agent-smith supports Kiro as a **native-subagent** platform (`kiro-ide`).

When configured, agent-smith installs:

| Artifact | Path |
|----------|------|
| Steering file | `~/.kiro/steering/agent-smith.md` |
| Native Judgment Day agents | `~/.kiro/agents/jd-fix-agent.md`, `jd-judge-a.md`, `jd-judge-b.md` *(3 files)* |
| Skills directory | `~/.kiro/skills/` |
| MCP config | `~/.kiro/settings/mcp.json` *(separate root â€” see note below)* |

> **Auto-install not supported.** Kiro must be installed manually before running agent-smith.
> Download from: [kiro.dev/downloads](https://kiro.dev/downloads)

---

## Detection

agent-smith uses **two signals** to detect Kiro:

1. **`~/.kiro` directory presence** â€” used by `system.ScanConfigs` for the install/TUI auto-detection flow. If `~/.kiro` exists on disk, Kiro is shown as detected in the installer, regardless of whether the binary is on `PATH`.
2. **`kiro` binary on `PATH`** â€” used by `adapter.Detect()` for the sync/upgrade flow and to confirm the IDE is actually runnable.

In practice: **the installer detects Kiro from `~/.kiro`**, not from `PATH`. If you have Kiro installed but `~/.kiro` hasn't been created yet (e.g., before first launch), run Kiro once to initialize its config dir, then re-run `agent-smith install`.

---

## ODD Execution Model

> **Since v4.0.0:** SDD (Spec-Driven Development) is retired in favor of [ODD](usage.md#organic-driven-development-odd). agent-smith no longer installs `sdd-*` Kiro agents and no longer routes work through `.kiro/specs/`. Existing `sdd-*` agent files from earlier installs are left in place.

Kiro runs with **native sub-agent delegation** via `~/.kiro/agents/`.

The ODD orchestrator stays in the steering file. It keeps understood work inline and delegates bounded delegated-direct work to Kiro's native subagents, with one writer at a time. Engramâ„¢ provides cross-session persistence when available.

The `jd-*` agents run the [Judgment Day](components.md#skills) adversarial review: two blind judges and one fix agent.

---

## Steering Files

**Steering files** at `.kiro/steering/*.md` provide persistent workspace context across sessions â€” treat them like always-on system context for your project conventions, architecture decisions, and team rules.

---

## Steering File Format

The steering file written by agent-smith uses the following frontmatter:

```yaml
---
inclusion: always
---
```

`inclusion: always` ensures Kiro loads this context in every conversation automatically, regardless of workspace or file type.

## Native Agent Frontmatter

Kiro Judgment Day agents are generated with YAML frontmatter including:

- `name`
- `description`
- `tools`
- `model`
- `includeMcpJson: true`

The `model` value is injected during sync from Kiro model assignments, keyed by agent name or `default` (`auto|opus|sonnet|haiku|minimax|glm|deepseek|qwen`) to Kiro-native model IDs.

---

## Config Paths by Platform

### macOS

| Artifact | Path |
|----------|------|
| Global config dir | `~/Library/Application Support/Kiro/User` |
| Steering file | `~/.kiro/steering/agent-smith.md` |
| Skills dir | `~/.kiro/skills/` |
| Settings path | `~/Library/Application Support/Kiro/User/settings.json` |
| MCP config | `~/.kiro/settings/mcp.json` |

### Windows

| Artifact | Path |
|----------|------|
| Global config dir | `%APPDATA%\kiro\User` |
| Steering file | `%USERPROFILE%\.kiro\steering\agent-smith.md` |
| Skills dir | `%USERPROFILE%\.kiro\skills\` |
| Settings path | `%APPDATA%\kiro\User\settings.json` |
| MCP config | `%USERPROFILE%\.kiro\settings\mcp.json` |

### Linux (XDG)

| Artifact | Path |
|----------|------|
| Global config dir | `$XDG_CONFIG_HOME/kiro/user` *(fallback: `~/.config/kiro/user`)* |
| Steering file | `~/.kiro/steering/agent-smith.md` |
| Skills dir | `~/.kiro/skills/` |
| Settings path | `$XDG_CONFIG_HOME/kiro/user/settings.json` |
| MCP config | `~/.kiro/settings/mcp.json` |

---

## âš ï¸ Split-Root Layout

Kiro uses a **split-root layout** â€” agent-smith managed files and IDE settings live in different directories:

- **Steering, skills, and native agents** â†’ `~/.kiro/` (or `%USERPROFILE%\.kiro\` on Windows)
  - `~/.kiro/steering/agent-smith.md` â€” orchestrator persona
  - `~/.kiro/skills/` â€” agent-smith skill files
  - `~/.kiro/agents/` â€” Judgment Day subagents
- **IDE settings** â†’ platform-native Kiro User dir (`settings.json` only)
  - macOS: `~/Library/Application Support/Kiro/User/settings.json`
  - Windows: `%APPDATA%\kiro\User\settings.json`
  - Linux: `$XDG_CONFIG_HOME/kiro/user/settings.json`
- **MCP config** â†’ always `~/.kiro/settings/mcp.json` (or `%USERPROFILE%\.kiro\settings\mcp.json` on Windows)

If MCP tools are not loading, check `~/.kiro/settings/mcp.json`.  
If Kiro app settings are not applying, check the platform-native User dir (`settings.json`).  
If agent-smith skills or steering are missing, check `~/.kiro/skills/` and `~/.kiro/steering/`.

---

## Capability Snapshot

| Capability | Status |
|------------|--------|
| Skills | âœ… Yes |
| System prompt | âœ… Yes |
| MCP | âœ… Yes |
| Output styles | âŒ No |
| Slash commands | âŒ No |
| Delegation model | Full (native subagents) |
| Auto-install | âŒ No â€” manual install required |

