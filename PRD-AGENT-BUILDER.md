# PRD: Agent Builder â€” Create Your Own Sub-Agent

> **Your ecosystem, your rules. Build custom AI sub-agents from the TUI â€” no code required.**

**Version**: 0.1.0-draft
**Author**: Gentleman Programming
**Date**: 2026-04-03
**Status**: Draft
**Parent PRD**: [PRD.md](PRD.md) (Gentleman AI Installer)

---

## 1. Problem Statement

The Gentleman AI ecosystem ships with a powerful set of pre-built skills and SDD phases. But every developer works differently. A frontend architect needs a design system reviewer. A security engineer needs a vulnerability scanner agent. A technical writer needs a documentation generator.

**Today, creating a custom sub-agent requires:**

1. Understanding the skill file format (`SKILL.md`) for each AI agent
2. Knowing where each agent stores its configuration (Claude Code: `~/.claude/skills/`, OpenCode: `~/.config/opencode/skills/`, etc.)
3. Writing the system prompt, triggers, and patterns manually
4. Duplicating the skill across every AI agent you use
5. If it's SDD-related, understanding the orchestrator config and how to register a new phase

**This is a barrier that shouldn't exist.** If you can describe what you want your agent to do in natural language, the ecosystem should generate it, validate it, and install it across all your configured AI agents â€” from a single TUI flow.

---

## 2. Vision

**A guided TUI experience where you describe what you want, and the ecosystem uses one of your installed AI agents to generate a production-ready sub-agent skill â€” installed across all your tools instantly.**

Think of it as the "Create Agent" flow from Claude's `/agents` command, but:
- **Cross-agent**: generates once, installs everywhere (Claude Code, OpenCode, Gemini CLI, Cursor, etc.)
- **SDD-aware**: can optionally integrate as a new SDD phase or as support for an existing phase
- **Ecosystem-native**: the generated agent has access to Engram (persistent memory), MCP servers, and follows the Gentleman skill format
- **Preview before install**: you see exactly what will be generated before it touches your filesystem

**Before**: "I want an agent that reviews my CSS for accessibility... I guess I need to write a SKILL.md manually, figure out triggers, and copy it to 4 different directories."

**After**: Select "Create your own Agent" â†’ pick your AI engine â†’ describe what you want â†’ preview the generated skill â†’ it's installed everywhere. Done.

---

## 3. Target Users

### Primary
- **Professional developers** who want specialized AI assistance beyond the pre-built skills
- **Team leads** who want to create standardized agents for their team's workflow (e.g., "our API review agent")
- **SDD users** who need custom phases or phase augmentation for their specific domain

### Secondary
- **Educators** creating custom teaching agents
- **Open source maintainers** who want project-specific review agents

---

## 4. User Experience

### 4.1 Entry Point â€” Welcome Screen

The Agent Builder is a **top-level menu option** on the Welcome screen, alongside the existing options:

```
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚                                                                  â”‚
â”‚   â•”â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•—                      â”‚
â”‚   â•‘   GENTLE AI                          â•‘                      â”‚
â”‚   â•šâ•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•â•                      â”‚
â”‚                                                                  â”‚
â”‚   Supercharge your AI agents. v0.x.x                             â”‚
â”‚                                                                  â”‚
â”‚   Menu                                                           â”‚
â”‚                                                                  â”‚
â”‚     Start installation                                           â”‚
â”‚     Upgrade tools                                                â”‚
â”‚     Sync configs                                                 â”‚
â”‚     Upgrade + Sync                                               â”‚
â”‚     Configure models                                             â”‚
â”‚   â˜… Create your own Agent                                        â”‚
â”‚     Manage backups                                               â”‚
â”‚     Quit                                                         â”‚
â”‚                                                                  â”‚
â”‚   j/k: navigate â€¢ enter: select â€¢ q: quit                       â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
```

### 4.2 Agent Builder Flow

```
"Create your own Agent"
         â”‚
         â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 1: Choose Your AI Engine   â”‚
â”‚                                  â”‚
â”‚  Which installed agent should    â”‚
â”‚  help you build your sub-agent?  â”‚
â”‚                                  â”‚
â”‚  â˜… Claude Code (installed)       â”‚  â† Uses claude --print to generate
â”‚  â—‹ OpenCode (installed)          â”‚  â† Uses opencode run to generate
â”‚  â—‹ Gemini CLI (installed)        â”‚  â† Uses gemini -p to generate
â”‚  â—‹ Codex (installed)             â”‚  â† Uses codex exec to generate
â”‚                                  â”‚
â”‚  Only installed agents shown.    â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 2: Describe Your Agent     â”‚
â”‚                                  â”‚
â”‚  Tell us what you want your      â”‚
â”‚  agent to do. Be as specific     â”‚
â”‚  as you can â€” the more detail,   â”‚
â”‚  the better the result.          â”‚
â”‚                                  â”‚
â”‚  Examples:                       â”‚
â”‚  â€¢ "Review CSS for a11y issues"  â”‚
â”‚  â€¢ "Generate API docs from code" â”‚
â”‚  â€¢ "Validate DB migrations"      â”‚
â”‚                                  â”‚
â”‚  â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”    â”‚
â”‚  â”‚ I want an agent that     â”‚    â”‚
â”‚  â”‚ reviews my React         â”‚    â”‚
â”‚  â”‚ components for           â”‚    â”‚
â”‚  â”‚ accessibility compliance â”‚    â”‚
â”‚  â”‚ following WCAG 2.1 AA    â”‚    â”‚
â”‚  â”‚ standards.               â”‚    â”‚
â”‚  â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜    â”‚
â”‚                                  â”‚
â”‚  [Continue]  [Back]              â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 3: SDD Integration         â”‚
â”‚                                  â”‚
â”‚  Should this agent be part of    â”‚
â”‚  the SDD (Spec-Driven Dev)       â”‚
â”‚  workflow?                       â”‚
â”‚                                  â”‚
â”‚  â˜… Standalone                    â”‚  â† Independent skill, not part of SDD
â”‚  â—‹ New SDD Phase                 â”‚  â† Adds as a new phase in the pipeline
â”‚  â—‹ Support for existing phase    â”‚  â† Augments an existing SDD phase
â”‚                                  â”‚
â”‚  [Continue]  [Back]              â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â”œâ”€â”€ If "New SDD Phase":
           â”‚   â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
           â”‚   â”‚  Where in the SDD pipeline?   â”‚
           â”‚   â”‚                               â”‚
           â”‚   â”‚  explore â†’ propose â†’ spec     â”‚
           â”‚   â”‚  â†’ design â†’ YOUR PHASE        â”‚
           â”‚   â”‚  â†’ tasks â†’ apply â†’ verify     â”‚
           â”‚   â”‚  â†’ archive                    â”‚
           â”‚   â”‚                               â”‚
           â”‚   â”‚  Insert after:                â”‚
           â”‚   â”‚  â—‹ explore                    â”‚
           â”‚   â”‚  â—‹ propose                    â”‚
           â”‚   â”‚  â—‹ spec                       â”‚
           â”‚   â”‚  â˜… design                     â”‚
           â”‚   â”‚  â—‹ tasks                      â”‚
           â”‚   â”‚  â—‹ apply                      â”‚
           â”‚   â”‚  â—‹ verify                     â”‚
           â”‚   â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â”œâ”€â”€ If "Support for existing phase":
           â”‚   â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
           â”‚   â”‚  Which phase to support?      â”‚
           â”‚   â”‚                               â”‚
           â”‚   â”‚  â—‹ explore                    â”‚
           â”‚   â”‚  â—‹ propose                    â”‚
           â”‚   â”‚  â—‹ spec                       â”‚
           â”‚   â”‚  â˜… design                     â”‚
           â”‚   â”‚  â—‹ tasks                      â”‚
           â”‚   â”‚  â—‹ apply                      â”‚
           â”‚   â”‚  â—‹ verify                     â”‚
           â”‚   â”‚  â—‹ archive                    â”‚
           â”‚   â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 4: Generating...           â”‚
â”‚                                  â”‚
â”‚  Using Claude Code to generate   â”‚
â”‚  your sub-agent...               â”‚
â”‚                                  â”‚
â”‚  [â–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–ˆâ–‘â–‘â–‘â–‘] 75%          â”‚
â”‚                                  â”‚
â”‚  â—Œ Analyzing your description    â”‚
â”‚  âœ“ Generating skill definition   â”‚
â”‚  â—Œ Creating trigger patterns     â”‚
â”‚  â—Œ Building agent instructions   â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 5: Preview Your Agent                                  â”‚
â”‚                                                              â”‚
â”‚  Name: a11y-reviewer                                         â”‚
â”‚  Description: Reviews React components for WCAG 2.1 AA       â”‚
â”‚               accessibility compliance                       â”‚
â”‚  Trigger: When reviewing React/JSX files for accessibility   â”‚
â”‚  SDD: Supports "design" phase                                â”‚
â”‚  Engram: âœ“ (reads project patterns, saves a11y decisions)    â”‚
â”‚                                                              â”‚
â”‚  â”€â”€ Generated Skill â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€   â”‚
â”‚  â”‚ # A11y Reviewer                                       â”‚   â”‚
â”‚  â”‚                                                       â”‚   â”‚
â”‚  â”‚ ## Description                                        â”‚   â”‚
â”‚  â”‚ Reviews React components for WCAG 2.1 AA compliance.  â”‚   â”‚
â”‚  â”‚ Checks semantic HTML, ARIA attributes, color contrast, â”‚   â”‚
â”‚  â”‚ keyboard navigation, and focus management.            â”‚   â”‚
â”‚  â”‚                                                       â”‚   â”‚
â”‚  â”‚ ## Trigger                                            â”‚   â”‚
â”‚  â”‚ When reviewing React/JSX/TSX files for accessibility  â”‚   â”‚
â”‚  â”‚ compliance, a11y audits, or WCAG validation.          â”‚   â”‚
â”‚  â”‚                                                       â”‚   â”‚
â”‚  â”‚ ## Instructions                                       â”‚   â”‚
â”‚  â”‚ ...                                                   â”‚   â”‚
â”‚  â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜   â”‚
â”‚                                                              â”‚
â”‚  Will be installed to:                                       â”‚
â”‚    â€¢ ~/.claude/skills/a11y-reviewer/SKILL.md                 â”‚
â”‚    â€¢ ~/.config/opencode/skills/a11y-reviewer/SKILL.md        â”‚
â”‚                                                              â”‚
â”‚  [Install]  [Edit]  [Regenerate]  [Back]                     â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”¬â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
           â”‚
           â”œâ”€â”€ If "Edit":
           â”‚   Opens $EDITOR with the generated SKILL.md
           â”‚   Returns to preview after editor closes
           â”‚
           â”œâ”€â”€ If "Regenerate":
           â”‚   Returns to generating screen with same prompt
           â”‚
           â–¼
â”Œâ”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”
â”‚  Step 6: Installing              â”‚
â”‚                                  â”‚
â”‚  Installing "a11y-reviewer"...   â”‚
â”‚                                  â”‚
â”‚  âœ“ Claude Code â€” skill installed â”‚
â”‚  âœ“ OpenCode â€” skill installed    â”‚
â”‚  âœ“ Skill registered in catalog   â”‚
â”‚  âœ“ SDD integration configured   â”‚
â”‚                                  â”‚
â”‚  Done! Your agent is ready.      â”‚
â”‚                                  â”‚
â”‚  To use it, ask your AI agent    â”‚
â”‚  to review a component for       â”‚
â”‚  accessibility.                  â”‚
â”‚                                  â”‚
â”‚  [Done]                          â”‚
â””â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”€â”˜
```

### 4.3 Managing Custom Agents

After creating agents, users need to manage them. A future iteration will add a "Manage custom agents" submenu (out of scope for V1, but architecture must support it):

- List all custom agents
- Edit an existing agent (re-open in preview)
- Delete a custom agent (removes from all configured AI agents)
- Export an agent (for sharing â€” future marketplace feature)

For V1, custom agents can be managed manually by editing/deleting the SKILL.md files.

---

## 5. Agent Generation System

### 5.1 Generation Engine

The Agent Builder uses **one of the user's installed AI agents** as the generation engine. This is a key architectural decision â€” we don't ship our own AI model; we leverage whatever the user already has.

#### Engine Abstraction

Each supported AI agent exposes a different CLI interface for non-interactive use:

| Agent | Command | Mode |
|-------|---------|------|
| Claude Code | `claude --print -p "{prompt}"` | Pipe prompt, get text output |
| OpenCode | `opencode run "{prompt}"` | Run mode, text output |
| Gemini CLI | `gemini -p "{prompt}"` | Pipe mode |
| Codex | `codex exec "{prompt}"` | Exec mode |

The system needs a `GenerationEngine` interface:

```go
type GenerationEngine interface {
    // Agent returns the ID of the AI agent used for generation.
    Agent() model.AgentID

    // Generate sends a prompt and returns the raw text output.
    Generate(ctx context.Context, prompt string) (string, error)

    // Available returns true if the agent binary is installed and accessible.
    Available() bool
}
```

### 5.2 Generation Prompt Strategy

The prompt sent to the AI engine is critical. It must produce a well-structured skill file that follows the ecosystem's conventions.

#### System Prompt (injected before user input)

```
You are an expert AI skill creator for the Gentleman AI ecosystem.
Your task is to generate a SKILL.md file based on the user's description.

The SKILL.md format must follow this structure:

# {Skill Name}

## Description
{What this skill does â€” 1-2 sentences}

## Trigger
{When this skill should be activated â€” specific file types, commands, or contexts}

## Instructions
{Detailed instructions for the AI agent when this skill is active}

## Rules
{Specific constraints, patterns, or conventions the agent must follow}

## Examples
{Concrete examples of input/output or behavior}

CONSTRAINTS:
- The skill MUST be self-contained in a single SKILL.md file
- Instructions MUST be actionable and specific, not vague
- Trigger conditions MUST be precise enough to avoid false activations
- If the user mentions SDD integration, include a section on how this skill
  interacts with the SDD phase it supports
- Include Engram integration instructions (when to save decisions, when to search
  for past context)
- Generate a kebab-case name derived from the description

OUTPUT FORMAT:
Return ONLY the SKILL.md content. No explanations, no wrapping, no code fences.
Start directly with "# {Name}".
```

#### User Prompt Composition

The final prompt combines:
1. The system prompt (above)
2. The user's description (from Step 2)
3. SDD integration context (if selected in Step 3)
4. Additional context about installed agents and ecosystem capabilities

```
{system_prompt}

USER REQUEST:
{user_description}

ADDITIONAL CONTEXT:
- This skill will be installed across these agents: {agent_list}
- SDD Integration: {standalone | new_phase_after_X | supports_phase_X}
- The agent has access to Engram persistent memory (mem_save, mem_search, mem_context)
- The agent can use MCP servers: {installed_mcp_list}
```

### 5.3 Output Parsing

The generation engine returns raw text. The builder must:

1. **Validate structure**: Ensure the output contains required sections (Description, Trigger, Instructions)
2. **Extract metadata**: Parse the skill name, description, and trigger for display in the preview
3. **Clean up**: Remove any code fences, preamble, or trailing text the AI might add
4. **Generate name**: Derive a kebab-case directory name from the skill title

```go
type GeneratedAgent struct {
    Name        string // kebab-case, e.g. "a11y-reviewer"
    Title       string // Human-readable, e.g. "A11y Reviewer"
    Description string // From ## Description section
    Trigger     string // From ## Trigger section
    Content     string // Full SKILL.md content
    SDDConfig   *SDDIntegration // nil if standalone
}

type SDDIntegration struct {
    Mode          SDDIntegrationMode // standalone | new_phase | phase_support
    TargetPhase   string             // e.g. "design" â€” which phase to support or insert after
    PhaseName     string             // e.g. "a11y-review" â€” name of new SDD phase (if new_phase)
}

type SDDIntegrationMode string

const (
    SDDStandalone    SDDIntegrationMode = "standalone"
    SDDNewPhase      SDDIntegrationMode = "new-phase"
    SDDPhaseSupport  SDDIntegrationMode = "phase-support"
)
```

---

## 6. Skill Installation

### 6.1 Where Skills Get Installed

Custom agent skills follow the SAME installation pattern as built-in skills. The existing `Adapter.SkillsDir()` method from the agent interface provides the correct path for each agent:

| Agent | Skill Directory | File Path |
|-------|----------------|-----------|
| Claude Code | `~/.claude/skills/` | `~/.claude/skills/{name}/SKILL.md` |
| OpenCode | `~/.config/opencode/skills/` | `~/.config/opencode/skills/{name}/SKILL.md` |
| Gemini CLI | `~/.gemini/skills/` | `~/.gemini/skills/{name}/SKILL.md` |
| Cursor | `~/.cursor/skills/` | `~/.cursor/skills/{name}/SKILL.md` |
| VSCode | `~/.vscode/skills/` | `~/.vscode/skills/{name}/SKILL.md` |
| Codex | `~/.codex/skills/` | `~/.codex/skills/{name}/SKILL.md` |
| Windsurf | `~/.windsurf/skills/` | `~/.windsurf/skills/{name}/SKILL.md` |
| Antigravity | `~/.antigravity/skills/` | `~/.antigravity/skills/{name}/SKILL.md` |

The installer writes the SAME `SKILL.md` to ALL agents that were configured during the initial setup (detected via the ecosystem's state file or by scanning which agent skill directories exist).

### 6.2 Custom Agent Registry

To track which custom agents the user has created (for future management), a local registry file is maintained:

```
~/.config/agent-smith/custom-agents.json
```

```json
{
  "version": 1,
  "agents": [
    {
      "name": "a11y-reviewer",
      "title": "A11y Reviewer",
      "description": "Reviews React components for WCAG 2.1 AA accessibility compliance",
      "created_at": "2026-04-03T14:30:00Z",
      "generation_engine": "claude-code",
      "sdd_integration": {
        "mode": "phase-support",
        "target_phase": "design"
      },
      "installed_agents": ["claude-code", "opencode"]
    }
  ]
}
```

### 6.3 SDD Integration Installation

When the custom agent has SDD integration, additional configuration is needed:

#### Phase Support Mode

The generated skill includes SDD-aware instructions. The orchestrator's system prompt is updated to include a reference to the custom skill:

```markdown
<!-- agent-smith:custom-agent:{name} -->
## Custom Agent: {Title}
When executing the "{target_phase}" phase, also load and apply the "{name}" skill
for additional validation/support.
<!-- /agent-smith:custom-agent:{name} -->
```

This is injected into the agent's system prompt using the existing `StrategyMarkdownSections` approach (marker-based injection that doesn't clobber user content).

#### New Phase Mode

A new SDD phase skill is created with the standard SDD skill structure. The orchestrator's dependency graph and phase list are updated to include the new phase at the specified position.

**Important**: This is a more complex integration. For V1, we inject the phase reference into the orchestrator's system prompt. The orchestrator (being an AI) will interpret the dependency graph and execute accordingly. No code changes to the SDD engine are needed â€” it's all prompt-driven.

---

## 7. Technical Architecture

### 7.1 Package Structure

```
internal/
â”œâ”€â”€ agentbuilder/          # Core agent builder logic
â”‚   â”œâ”€â”€ builder.go         # Main builder orchestrator
â”‚   â”œâ”€â”€ engine.go          # GenerationEngine interface + implementations
â”‚   â”œâ”€â”€ parser.go          # Output parsing and validation
â”‚   â”œâ”€â”€ installer.go       # Skill file installation across agents
â”‚   â”œâ”€â”€ registry.go        # Custom agent registry (JSON file management)
â”‚   â”œâ”€â”€ prompt.go          # Prompt composition logic
â”‚   â””â”€â”€ sdd.go             # SDD integration logic
â”œâ”€â”€ tui/
â”‚   â”œâ”€â”€ screens/
â”‚   â”‚   â”œâ”€â”€ agent_builder_engine.go      # Step 1: Engine selection
â”‚   â”‚   â”œâ”€â”€ agent_builder_prompt.go      # Step 2: Description input
â”‚   â”‚   â”œâ”€â”€ agent_builder_sdd.go         # Step 3: SDD integration
â”‚   â”‚   â”œâ”€â”€ agent_builder_generating.go  # Step 4: Generation progress
â”‚   â”‚   â”œâ”€â”€ agent_builder_preview.go     # Step 5: Preview + edit
â”‚   â”‚   â””â”€â”€ agent_builder_complete.go    # Step 6: Installation complete
â”‚   â”œâ”€â”€ model.go           # Add new Screen constants + AgentBuilder state
â”‚   â””â”€â”€ router.go          # Add agent builder routes
```

### 7.2 New Screen Constants

```go
const (
    // ... existing screens ...
    ScreenAgentBuilderEngine    Screen = iota + 100 // Offset to avoid conflicts
    ScreenAgentBuilderPrompt
    ScreenAgentBuilderSDD
    ScreenAgentBuilderSDDPhase
    ScreenAgentBuilderGenerating
    ScreenAgentBuilderPreview
    ScreenAgentBuilderInstalling
    ScreenAgentBuilderComplete
)
```

### 7.3 Model Extensions

```go
type Model struct {
    // ... existing fields ...

    // Agent Builder state
    AgentBuilder AgentBuilderState
}

type AgentBuilderState struct {
    // Step 1: Selected generation engine
    Engine         model.AgentID
    AvailableEngines []model.AgentID

    // Step 2: User prompt
    PromptText     string
    PromptCursorPos int

    // Step 3: SDD integration
    SDDMode        SDDIntegrationMode
    SDDTargetPhase string

    // Step 4: Generation
    Generating     bool
    GenerationErr  error

    // Step 5: Preview
    Generated      *GeneratedAgent
    PreviewScroll  int

    // Step 6: Installation
    Installing     bool
    InstallResult  *InstallResult
}
```

### 7.4 Router Integration

The agent builder has its own sub-flow that's independent of the main installation flow:

```go
var agentBuilderRoutes = map[Screen]Route{
    ScreenAgentBuilderEngine:     {Backward: ScreenWelcome},
    ScreenAgentBuilderPrompt:     {Backward: ScreenAgentBuilderEngine},
    ScreenAgentBuilderSDD:        {Backward: ScreenAgentBuilderPrompt},
    ScreenAgentBuilderSDDPhase:   {Backward: ScreenAgentBuilderSDD},
    ScreenAgentBuilderGenerating: {Backward: ScreenAgentBuilderPrompt},
    ScreenAgentBuilderPreview:    {Forward: ScreenAgentBuilderInstalling, Backward: ScreenAgentBuilderPrompt},
    ScreenAgentBuilderInstalling: {Forward: ScreenAgentBuilderComplete},
    ScreenAgentBuilderComplete:   {Backward: ScreenWelcome},
}
```

### 7.5 Generation Engine Implementations

Each engine implementation wraps the agent's CLI:

```go
// ClaudeEngine generates skills using Claude Code's --print mode.
type ClaudeEngine struct {
    binaryPath string
}

func (e *ClaudeEngine) Generate(ctx context.Context, prompt string) (string, error) {
    cmd := exec.CommandContext(ctx, e.binaryPath, "--print", "-p", prompt)
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("claude generation failed: %w", err)
    }
    return string(output), nil
}

// OpenCodeEngine generates skills using OpenCode's run mode.
type OpenCodeEngine struct {
    binaryPath string
}

func (e *OpenCodeEngine) Generate(ctx context.Context, prompt string) (string, error) {
    cmd := exec.CommandContext(ctx, e.binaryPath, "run", prompt)
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("opencode generation failed: %w", err)
    }
    return string(output), nil
}
```

### 7.6 Sequence Diagram â€” Generation Flow

```mermaid
sequenceDiagram
    participant User
    participant TUI as Agent Smith TUI
    participant Builder as Agent Builder
    participant Engine as Generation Engine<br/>(Claude/OpenCode/etc.)
    participant Parser as Output Parser
    participant Installer as Skill Installer
    participant FS as Filesystem

    User->>TUI: Select "Create your own Agent"
    TUI->>TUI: Show engine selection
    User->>TUI: Select Claude Code
    TUI->>TUI: Show prompt input
    User->>TUI: Enter description
    TUI->>TUI: Show SDD integration options
    User->>TUI: Select "Supports design phase"

    TUI->>Builder: Build(engine, prompt, sddConfig)
    Builder->>Builder: ComposePrompt(systemPrompt, userInput, sddContext)
    Builder->>Engine: Generate(ctx, composedPrompt)
    Engine->>Engine: claude --print -p "{prompt}"
    Engine-->>Builder: Raw SKILL.md text

    Builder->>Parser: Parse(rawText)
    Parser->>Parser: Validate sections
    Parser->>Parser: Extract metadata
    Parser-->>Builder: GeneratedAgent{name, content, ...}

    Builder-->>TUI: GeneratedAgent
    TUI->>TUI: Show preview

    User->>TUI: Click "Install"
    TUI->>Installer: Install(generatedAgent, targetAgents)
    
    loop For each configured agent
        Installer->>FS: Write SKILL.md to agent's skills dir
    end
    
    Installer->>FS: Update custom-agents.json registry
    
    alt SDD Integration
        Installer->>FS: Update system prompt with SDD reference
    end

    Installer-->>TUI: InstallResult
    TUI->>TUI: Show completion
```

### 7.7 Architecture Diagram â€” Component Relationships

```mermaid
graph TB
    subgraph TUI_LAYER["TUI Layer (Bubbletea)"]
        WELCOME[Welcome Screen<br/>+ "Create your own Agent"]
        AB_ENGINE[Engine Selection]
        AB_PROMPT[Prompt Input]
        AB_SDD[SDD Integration]
        AB_GENERATING[Generating Screen]
        AB_PREVIEW[Preview Screen]
        AB_INSTALLING[Installing Screen]
        AB_COMPLETE[Complete Screen]

        WELCOME --> AB_ENGINE
        AB_ENGINE --> AB_PROMPT
        AB_PROMPT --> AB_SDD
        AB_SDD --> AB_GENERATING
        AB_GENERATING --> AB_PREVIEW
        AB_PREVIEW --> AB_INSTALLING
        AB_INSTALLING --> AB_COMPLETE
        AB_COMPLETE --> WELCOME
    end

    subgraph BUILDER_LAYER["Agent Builder Core"]
        BUILDER[Builder Orchestrator]
        PROMPT_COMP[Prompt Composer]
        PARSER[Output Parser]
        REGISTRY[Custom Agent Registry]
    end

    subgraph ENGINE_LAYER["Generation Engines"]
        ENGINE_IF{GenerationEngine<br/>Interface}
        CLAUDE_ENG[Claude Engine<br/>claude --print]
        OC_ENG[OpenCode Engine<br/>opencode run]
        GEM_ENG[Gemini Engine<br/>gemini -p]
        CDX_ENG[Codex Engine<br/>codex exec]

        ENGINE_IF --> CLAUDE_ENG
        ENGINE_IF --> OC_ENG
        ENGINE_IF --> GEM_ENG
        ENGINE_IF --> CDX_ENG
    end

    subgraph INSTALLER_LAYER["Skill Installer"]
        SKILL_INST[Skill Installer]
        SDD_INT[SDD Integrator]
        AGENT_ADAPTER[Agent Adapters<br/>SkillsDir / SystemPromptStrategy]
    end

    subgraph STORAGE["Storage"]
        SKILLS_DIR["~/.{agent}/skills/{name}/SKILL.md"]
        REGISTRY_FILE["~/.config/agent-smith/custom-agents.json"]
        SYS_PROMPT["Agent system prompts<br/>(CLAUDE.md / AGENTS.md / etc.)"]
    end

    AB_GENERATING --> BUILDER
    BUILDER --> PROMPT_COMP
    BUILDER --> ENGINE_IF
    BUILDER --> PARSER
    AB_INSTALLING --> SKILL_INST
    SKILL_INST --> AGENT_ADAPTER
    SKILL_INST --> SDD_INT
    SKILL_INST --> REGISTRY
    SKILL_INST --> SKILLS_DIR
    REGISTRY --> REGISTRY_FILE
    SDD_INT --> SYS_PROMPT

    style TUI_LAYER fill:#1a1b26,stroke:#E0C15A,color:#E0C15A
    style BUILDER_LAYER fill:#1a1b26,stroke:#7FB4CA,color:#7FB4CA
    style ENGINE_LAYER fill:#1a1b26,stroke:#B7CC85,color:#B7CC85
    style INSTALLER_LAYER fill:#1a1b26,stroke:#957FB8,color:#957FB8
    style STORAGE fill:#1a1b26,stroke:#CB7C94,color:#CB7C94
```

---

## 8. SDD Integration â€” Deep Dive

### 8.1 Standalone Mode

The simplest mode. The generated skill is a regular skill file â€” no SDD awareness. It triggers based on file context or user invocation, just like `react-19` or `typescript`.

### 8.2 Phase Support Mode

The custom agent **augments** an existing SDD phase. Example: an "a11y-reviewer" that supports the "design" phase.

**How it works at runtime:**

1. The SDD orchestrator reaches the `design` phase
2. The orchestrator's system prompt includes a reference to the custom agent:
   ```
   When executing the "design" phase, also consider loading the
   "a11y-reviewer" skill for accessibility validation of the design.
   ```
3. The AI agent (being intelligent) reads the custom skill and incorporates its guidance into the design phase output
4. The custom skill's Engram instructions tell the agent to save accessibility decisions to memory

**What gets modified:**
- The agent's system prompt (CLAUDE.md, AGENTS.md, etc.) gets a new marker section referencing the custom skill
- The custom SKILL.md includes instructions specific to the supported phase

### 8.3 New Phase Mode

The custom agent becomes a **new phase** in the SDD pipeline. Example: an "a11y-audit" phase that runs between "design" and "tasks".

**How it works at runtime:**

1. The SDD orchestrator's dependency graph is updated:
   ```
   proposal -> specs -> design -> a11y-audit -> tasks -> apply -> verify -> archive
   ```
2. The custom skill follows the SDD phase contract:
   - Reads artifacts from the previous phase
   - Writes its own artifact to Engram (topic key: `sdd/{change}/a11y-audit`)
   - Returns the standard phase result: `status`, `executive_summary`, `artifacts`, `next_recommended`

**What gets modified:**
- A new SDD skill file is created following the phase skill pattern
- The orchestrator's system prompt is updated with the modified dependency graph
- The Engram topic key format is documented in the skill

### 8.4 SDD Phase Positions

The user selects where in the pipeline to insert the new phase:

| Insert After | Resulting Pipeline |
|--------------|-------------------|
| explore | explore â†’ **custom** â†’ propose â†’ spec â†’ design â†’ tasks â†’ apply â†’ verify â†’ archive |
| propose | explore â†’ propose â†’ **custom** â†’ spec â†’ design â†’ tasks â†’ apply â†’ verify â†’ archive |
| spec | explore â†’ propose â†’ spec â†’ **custom** â†’ design â†’ tasks â†’ apply â†’ verify â†’ archive |
| design | explore â†’ propose â†’ spec â†’ design â†’ **custom** â†’ tasks â†’ apply â†’ verify â†’ archive |
| tasks | explore â†’ propose â†’ spec â†’ design â†’ tasks â†’ **custom** â†’ apply â†’ verify â†’ archive |
| apply | explore â†’ propose â†’ spec â†’ design â†’ tasks â†’ apply â†’ **custom** â†’ verify â†’ archive |
| verify | explore â†’ propose â†’ spec â†’ design â†’ tasks â†’ apply â†’ verify â†’ **custom** â†’ archive |

---

## 9. Requirements

### Functional Requirements

| ID | Requirement | Priority |
|----|------------|----------|
| R-AB-01 | The Agent Builder MUST be accessible from the Welcome screen as a top-level menu option | P0 |
| R-AB-02 | The Agent Builder MUST only show installed AI agents as generation engine options | P0 |
| R-AB-03 | The Agent Builder MUST accept free-form natural language description as input | P0 |
| R-AB-04 | The Agent Builder MUST generate a valid SKILL.md file using the selected AI engine | P0 |
| R-AB-05 | The Agent Builder MUST show a preview of the generated skill before installation | P0 |
| R-AB-06 | The Agent Builder MUST install the generated skill to ALL configured AI agents | P0 |
| R-AB-07 | The Agent Builder MUST support SDD integration in three modes: standalone, phase support, new phase | P0 |
| R-AB-08 | The Agent Builder MUST maintain a local registry of custom agents at `~/.config/agent-smith/custom-agents.json` | P0 |
| R-AB-09 | The Agent Builder MUST allow the user to edit the generated skill before installation (open in $EDITOR) | P1 |
| R-AB-10 | The Agent Builder MUST allow the user to regenerate the skill with the same prompt | P0 |
| R-AB-11 | The Agent Builder MUST include Engram integration instructions in every generated skill | P0 |
| R-AB-12 | The Agent Builder MUST handle generation engine errors gracefully with clear error messages | P0 |
| R-AB-13 | The Agent Builder MUST support generation timeouts (configurable, default 120s) | P1 |
| R-AB-14 | The generated skill MUST be a standalone SKILL.md file â€” no external dependencies | P0 |
| R-AB-15 | For SDD phase support mode, the agent's system prompt MUST be updated with a marker-based reference to the custom skill | P0 |
| R-AB-16 | For SDD new phase mode, the orchestrator's dependency graph in the system prompt MUST be updated | P0 |
| R-AB-17 | The Agent Builder MUST detect which agents were configured by the installer (via existing config scan or state file) | P0 |
| R-AB-18 | The text input for the agent description MUST support multi-line input with scrolling | P0 |
| R-AB-19 | The preview screen MUST support scrolling for long skill definitions | P0 |
| R-AB-20 | The Agent Builder flow MUST support Esc to go back at every step | P0 |

### Non-Functional Requirements

| ID | Requirement | Priority |
|----|------------|----------|
| R-AB-NF-01 | Generation MUST complete within 120 seconds or show a timeout error | P0 |
| R-AB-NF-02 | The TUI MUST remain responsive during generation (spinner animation, ability to cancel) | P0 |
| R-AB-NF-03 | The Agent Builder MUST follow the same Bubbletea + Lipgloss styling as the rest of the TUI | P0 |
| R-AB-NF-04 | The Agent Builder architecture MUST allow adding new generation engines by implementing the GenerationEngine interface | P0 |
| R-AB-NF-05 | The custom-agents.json registry format MUST be versioned for forward compatibility | P1 |
| R-AB-NF-06 | Skill installation MUST be atomic â€” if installation to any agent fails, already-installed copies are cleaned up | P1 |

---

## 10. Screens

| Screen | Purpose | Key Actions |
|--------|---------|-------------|
| Agent Builder: Engine Selection | Choose which installed AI agent generates the skill | Select engine, Back |
| Agent Builder: Prompt Input | Multi-line text input describing the desired agent | Type description, Continue, Back |
| Agent Builder: SDD Integration | Choose standalone, phase support, or new phase | Select mode, Continue, Back |
| Agent Builder: SDD Phase Picker | Select which SDD phase to support or insert after | Select phase, Continue, Back |
| Agent Builder: Generating | Show progress while AI generates the skill | Spinner, cancel |
| Agent Builder: Preview | Show generated skill with metadata summary | Install, Edit, Regenerate, Back |
| Agent Builder: Installing | Show installation progress across agents | Progress animation |
| Agent Builder: Complete | Success message with usage instructions | Done (returns to Welcome) |

---

## 11. Edge Cases & Error Handling

| Scenario | Behavior |
|----------|----------|
| No AI agents installed | "Create your own Agent" menu option is **disabled** with "(no agents)" suffix |
| Selected engine fails to generate | Show error message with the engine's stderr. Offer "Retry" or "Try different engine" |
| Generated output doesn't contain required sections | Show warning: "The generated skill is missing sections: {list}. Edit manually or regenerate." |
| Skill name conflicts with built-in skill | Append `-custom` suffix. Warn user: "Name '{name}' conflicts with built-in skill. Using '{name}-custom'." |
| Skill name conflicts with existing custom agent | Ask user: "Agent '{name}' already exists. Replace it?" |
| $EDITOR not set (Edit action) | Fall back to `vi`. If `vi` not available, show the raw content in a scrollable pane with copy-paste instructions |
| Agent skills directory doesn't exist | Create it (same behavior as the main installer) |
| Generation exceeds timeout | Show timeout error. Offer "Retry with longer timeout" (2x) or "Try different engine" |
| User prompt is empty | "Continue" button is disabled. Show helper text: "Describe what you want your agent to do" |
| Network error during generation | Show clear error. Note: all engines run locally â€” network errors are unlikely but possible with API-based agents |

---

## 12. Future Considerations (Out of Scope for V1)

| Feature | Description | Why Later |
|---------|-------------|-----------|
| **Marketplace** | Share and discover community-created agents | Needs backend infrastructure, auth, trust model |
| **Templates** | Pre-built starting points (Code Reviewer, Doc Writer, Test Generator) | Can be added once the core builder is solid |
| **Agent Management Screen** | List, edit, delete, export custom agents from TUI | Registry is there; UI can come later |
| **Team Sync** | Share custom agents across team members via git | Needs a convention for team-shared skills |
| **Multi-model Generation** | Use multiple AI engines in sequence (e.g., Claude generates, Gemini refines) | Complex orchestration, diminishing returns |
| **Knowledge Files** | Attach reference documents to the custom agent | File management UX is complex |
| **Agent Testing** | "Try your agent" sandbox before installing | Would need a sandboxed agent execution environment |
| **Version Control** | Track versions of custom agents, rollback | Registry versioning is the foundation |

---

## 13. Success Metrics

| Metric | Target | How to Measure |
|--------|--------|---------------|
| Completion rate | >80% of users who start the builder finish creating an agent | Registry entries vs. builder starts (future telemetry) |
| Generation quality | >70% of generated skills used without manual editing | Track "Install" vs "Edit" actions (future telemetry) |
| Cross-agent installation | 100% of configured agents receive the skill | Verified by installation step; logged in registry |
| SDD integration usage | >30% of custom agents use SDD integration | Registry `sdd_integration.mode` distribution |

---

## 14. Implementation Notes

### 14.1 Reusing Existing Infrastructure

The Agent Builder deliberately reuses existing infrastructure:

- **Agent detection**: `agents.Adapter.Detect()` â€” same mechanism as the main installer
- **Skill paths**: `agents.Adapter.SkillsDir()` â€” same paths as built-in skill installation
- **System prompt injection**: `model.StrategyMarkdownSections` â€” same marker-based injection for SDD integration
- **TUI patterns**: Same Bubbletea + Lipgloss styling, same keyboard navigation (j/k, Enter, Esc)
- **Agent registry**: `agents.Registry` â€” used to enumerate available engines

### 14.2 Text Input Considerations

The prompt input (Step 2) is the most complex TUI element. It needs:

- Multi-line text input (not just single-line like backup rename)
- Scrolling for long descriptions
- Word wrap
- Basic cursor navigation (arrows, Home/End)
- Paste support

Consider using [charmbracelet/textarea](https://github.com/charmbracelet/textarea) â€” a Bubbletea component designed for multi-line text input. This avoids building custom text editing logic.

### 14.3 Generation as a Goroutine

The generation step (Step 4) runs the AI engine CLI as a subprocess. This MUST run in a goroutine to keep the TUI responsive. The pattern follows the existing `startInstalling()` / `PipelineDoneMsg` approach:

```go
type AgentBuilderDoneMsg struct {
    Agent *GeneratedAgent
    Err   error
}

func (m Model) startGeneration() tea.Cmd {
    engine := m.AgentBuilder.selectedEngine
    prompt := m.AgentBuilder.composedPrompt
    return func() tea.Msg {
        ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
        defer cancel()
        result, err := engine.Generate(ctx, prompt)
        if err != nil {
            return AgentBuilderDoneMsg{Err: err}
        }
        agent, err := ParseGeneratedAgent(result)
        return AgentBuilderDoneMsg{Agent: agent, Err: err}
    }
}
```

### 14.4 Detecting Configured Agents

The builder needs to know which agents to install the skill to. Options:

1. **Scan filesystem**: Check if each agent's skill directory exists (created by the main installer)
2. **Read state file**: If the installer persists its selections somewhere
3. **Re-run detection**: Use `agents.Adapter.Detect()` for each agent

For V1, option 1 (scan filesystem) is the simplest and most reliable â€” if the skills directory exists, the agent was configured.

