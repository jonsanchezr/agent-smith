# Skill Registry

> [!NOTE]
> These docs track `main`, which may include unreleased changes. For the latest release, see the [v4.0.0 docs](https://github.com/jonsanchezr/agent-smith/tree/v4.0.0/docs).

â† [Back to README](../README.md)

The skill registry is a project-local index that lets every supported agent find the same skills without rewriting them. It stores skill names, full descriptions, scopes, and exact `SKILL.md` paths.

## When To Use It

Use `agent-smith skill-registry refresh` after you add, remove, rename, or move skills. Normal installs wire this refresh into startup hooks where the agent supports them, including Codex, Claude Code, OpenCode, and Pi through `gentle-pi`.

## Runtime Flow

```text
User task
   â”‚
   â–¼
Orchestrator reads .atl/skill-registry.md
   â”‚
   â–¼
Matches task + file context against full skill descriptions
   â”‚
   â–¼
Passes exact SKILL.md paths to subagent
   â”‚
   â–¼
Subagent reads full skills before work
   â”‚
   â–¼
Subagent executes with original skill intent preserved
```

## Refresh Flow

```text
agent-smith skill-registry refresh
   â”‚
   â”œâ”€ Scan project skill roots first
   â”‚     skills/, .opencode/skills/, .claude/skills/, ...
   â”‚
   â”œâ”€ Scan global agent skill roots second
   â”‚     ~/.config/opencode/skills/, ~/.claude/skills/, ...
   â”‚
   â”œâ”€ Deduplicate by skill name
   â”‚     project skill wins over global skill
   â”‚
   â”œâ”€ Parse frontmatter
   â”‚     name + full description + path + scope
   â”‚
   â””â”€ Write .atl/skill-registry.md + cache
```

## Registry Contract

The registry is an **index**, not a generated summary.

| Field | Meaning |
| --- | --- |
| `Skill` | Skill `name` from frontmatter, or directory name fallback |
| `Trigger / description` | Full `description`, including YAML folded multiline descriptions |
| `Scope` | `project` or `user` |
| `Path` | Exact `SKILL.md` file to load |

## Skill Loading Contract

Delegators pass paths, not digested rules:

```markdown
## Skills to load before work

Read these exact files before reading, writing, reviewing, testing, or creating artifacts:

- /path/to/skills/go-testing/SKILL.md
- /path/to/skills/docs-writer/SKILL.md
```

The subagent then reads those files. This keeps the original `SKILL.md` as the source of truth and avoids breaking author intent through automatic summarization.

## Skill Authoring Flow

```text
New reusable pattern
   â”‚
   â–¼
skill-creator creates SKILL.md
   â”‚
   â–¼
skill-registry indexes SKILL.md path and full description
   â”‚
   â–¼
orchestrator passes matching paths to agents
```

## Skill Improvement Flow

```text
Existing skills
   â”‚
   â–¼
skill-improver reads .atl/skill-registry.md
   â”‚
   â–¼
Audits each indexed SKILL.md against docs/skill-style-guide.md
   â”‚
   â”œâ”€ Audit mode: report issues only
   â”‚
   â””â”€ Apply mode: safely refactor skills and preserve intent
   â”‚
   â–¼
Run agent-smith skill-registry refresh again
```

## Why Not Compact Rules?

Compact rules were cheaper per delegation but could distort skills. The index-first design spends tokens only when a subagent actually needs a skill, and it preserves the complete runtime contract.

| Design | Benefit | Tradeoff |
| --- | --- | --- |
| Compact summaries | Small prompt injection | Can lose nuance and break custom skills |
| Index + paths | Preserves full skill intent | Subagents read selected full skills |

## Excluded Skills

The registry never indexes `_shared`, `skill-registry`, or any `sdd-*` skill.
The first two are internal plumbing; the `sdd-*` prefix belonged to the retired
SDD workflow and stays reserved. This exclusion is intentional and
silent, so a user skill whose name collides with these prefixes is dropped
without a warning.

## Inspecting Without Writing

`skill-registry list` resolves the same deduplicated skill set as `refresh`, but
prints it instead of writing `.atl/skill-registry.md`, the cache, or
`.gitignore`. Handy for debugging what a delegator would see.

```bash
agent-smith skill-registry list          # name<TAB>scope<TAB>path
agent-smith skill-registry list --json   # machine-readable, includes descriptions
```

## Quick Check

```bash
agent-smith skill-registry refresh --force
```

Open `.atl/skill-registry.md` and verify each row has a useful description and a real `SKILL.md` path.

â† [Back to README](../README.md)

