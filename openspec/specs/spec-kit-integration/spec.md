# Spec Kit Integration Specification

## Purpose

Embed Spec Kit (github.com/jonsanchezr/spec-kit) as an alternative SDD workflow in Agent Smith. Users invoke `/speckit-*` slash commands to delegate to the `specify` CLI, while existing `/sdd-*` commands remain unchanged. Both workflows coexist in the orchestrator as explicit routing options.

## ADDED Requirements

### Requirement: Spec Kit Skill Bundle

Agent Smith SHALL bundle SKILL.md files for the following Spec Kit phases:

- `speckit-constitution` — establishes project constitution
- `speckit-specify` — captures requirements and specifications
- `speckit-plan` — creates implementation plan
- `speckit-checklist` — review checklist generation
- `speckit-tasks` — task breakdown generation
- `speckit-analyze` — architectural analysis
- `speckit-implement` — implementation guidance
- `speckit-converge` — convergence and validation

Each SKILL.md file SHALL delegate to `specify {phase}` CLI command with OpenCode-compatible frontmatter.

### Requirement: Spec Kit Installation Detection

When a user invokes any `/speckit-*` command, Agent Smith SHALL check whether `specify-cli` is installed via `speckit_available()`.

The system SHALL attempt auto-installation using `uv tool install specify-cli` if the CLI is absent and `uv` is on PATH.

If Python 3.11+ is absent, the system SHALL emit a clear error with installation instructions.

### Requirement: Orchestrator Workflow Routing

The orchestrator SHALL present both SDD workflows as explicit routing options:

- **Agent Smith ODD**: `/sdd-init`, `/sdd-propose`, `/sdd-spec`, `/sdd-design`, `/sdd-tasks`, `/sdd-apply`, `/sdd-verify`
- **Spec Kit SDD**: `/speckit-constitution`, `/speckit-specify`, `/speckit-plan`, `/speckit-tasks`, `/speckit-implement`, `/speckit-converge`

User SHALL explicitly choose per-task which workflow to invoke.

### Requirement: Skill Injection for Non-SDD Prefixed Skills

The `inject.go` `directoryAssets()` function SHALL scan `internal/assets/skills/speckit-*` directories alongside existing skill paths.

The `IsSDDSkill()` guard (prefix `sdd-`) SHALL NOT block `speckit-*` skill IDs.

### Requirement: Spec Kit Skill Catalog Registration

The system SHALL register `speckit-*` skills in `internal/catalog/skills.go` under a `spec-kit` category with distinct SkillID constants:

- `SkillSpecKitConstitution`, `SkillSpecKitSpecify`, `SkillSpecKitPlan`, `SkillSpecKitTasks`, `SkillSpecKitImplement`, `SkillSpecKitConverge`, `SkillSpecKitInstall`

## ADDED Scenarios

### Scenario: User invokes speckit-specify with no Python/uv

- GIVEN user types `/speckit-specify` in a project
- WHEN neither Python 3.11+ nor `uv` is on PATH
- THEN Agent Smith SHALL emit error: "Python 3.11+ and uv are required for Spec Kit. Install Python, then run: uv tool install specify-cli"
- AND SHALL NOT attempt CLI invocation

### Scenario: User invokes speckit-specify with spec-kit already installed

- GIVEN user types `/speckit-specify` and `specify-cli` is already installed
- WHEN `speckit_available()` returns true
- THEN Agent Smith SHALL load `internal/assets/skills/speckit-specify/SKILL.md`
- AND SHALL delegate to `specify specify` CLI command

### Scenario: User chooses Spec Kit workflow in orchestrator

- GIVEN user is in the orchestrator and has both workflows available
- WHEN user selects Spec Kit SDD path
- THEN orchestrator SHALL route to `/speckit-specify` skill
- AND Agent Smith SDD skills remain accessible via `/sdd-*` prefix

### Scenario: User chooses Agent Smith SDD workflow

- GIVEN user is in the orchestrator and has both workflows available
- WHEN user selects Agent Smith ODD path
- THEN orchestrator SHALL route to `/sdd-*` skills as before
- AND Spec Kit skills remain accessible via `/speckit-*` prefix

### Scenario: spec-kit CLI not in PATH after installation

- GIVEN `specify-cli` was installed via `uv tool install` but PATH was not updated
- WHEN user invokes any `/speckit-*` command
- THEN `speckit_available()` SHALL return false
- AND Agent Smith SHALL emit error with: "specify-cli not found in PATH. Ensure uv tool install specify-cli completed and PATH includes uv bin directory."

### Scenario: spec-kit upgrade available

- GIVEN `specify-cli` is installed but an older version is detected
- WHEN `speckit_available()` returns true with outdated version
- THEN Agent Smith SHOULD log: "A newer spec-kit version may be available. Run: uv tool upgrade specify-cli"
- AND SHALL proceed with current installation

### Scenario: Air-gapped offline install attempt

- GIVEN machine is air-gapped with no internet access
- WHEN user invokes `/speckit-*` and `specify-cli` is not installed
- THEN `speckit_available()` SHALL return false
- AND Agent Smith SHALL emit: "specify-cli not found and network unavailable. Install manually: download specify-cli wheel and run uv tool install --offline"

### Scenario: Multiple projects with different spec-kit states

- GIVEN Project A has spec-kit installed and Project B does not
- WHEN user works in Project B and invokes `/speckit-*`
- THEN Agent Smith SHALL attempt installation in Project B context
- AND Project A settings SHALL remain unaffected

## Out of Scope (Documented for Clarity)

The following Spec Kit extensions are NOT integrated and require manual invocation:

- `spec-kit bug` — issue/bug workflow (out of scope for SDD)
- `spec-kit assess` — idea assessment workflow (out of scope for SDD)

These commands MAY be added separately but are not part of this change.

## Data/Artifact Contract

| Artifact | Location | Purpose |
|----------|----------|---------|
| Spec Kit SKILL.md files | `internal/assets/skills/speckit-{phase}/SKILL.md` | OpenCode skill wrappers for each spec-kit phase |
| Skill catalog | `internal/catalog/skills.go` | `speckit-*` skill registration under `spec-kit` category |
| Skill type constants | `internal/model/types.go` | `SkillSpecKit*` constants |
| Skill injection | `internal/components/skills/inject.go` | `directoryAssets()` includes `speckit-*` paths |
| Orchestrator prompt | `internal/assets/opencode/orchestrator.md` | Workflow selection block with both SDD options |

## Installation Artifacts Created

When `agent-smith install` triggers spec-kit setup, the following are created/updated:

- Python virtual environment: `~/.local/share/uv/tools/specify-cli/` (via `uv tool install`)
- Spec Kit CLI: `specify` command on PATH (via `uv tool install specify-cli`)
- Integration setup: SKILL.md files copied to `internal/assets/skills/speckit-*/`
