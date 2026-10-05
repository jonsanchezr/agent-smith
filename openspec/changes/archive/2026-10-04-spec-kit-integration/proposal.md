# Proposal: Integrate Spec Kit as Alternative SDD Workflow

## Intent

Agent Smith currently ships a homegrown ODD/Sdd workflow. This change embeds github/spec-kit (a mature Python SDD toolkit with 140k stars) as a second, formally-specified workflow option accessible via `/speckit-*` slash commands. Users choose per-task which workflow to invoke; both coexist in the orchestrator routing table.

## Scope

### In Scope
- Bundle Spec Kit's `SkillsIntegration` output (SKILL.md per phase) as embedded assets in `internal/assets/skills/speckit-*/`
- Add `speckit-install` prerequisite detection and auto-install to the Agent Smith setup/install flow
- Register `speckit-*` skills in `internal/catalog/skills.go` and `internal/model/types.go` (SkillID constants)
- Add injection path in `internal/components/skills/inject.go` for non-`sdd-` prefixed skill IDs
- Update orchestrator prompt to surface both workflows as routing options

### Out of Scope
- Modifying Spec Kit's Python source or CLI (`specify-cli`)
- Bundling Python runtime or `uv` — assumed pre-installed on host
- Porting Spec Kit's bug-fix or idea-assessment extensions (SDD workflow only)
- Changes to the existing ODD/Sdd skill implementations

## Capabilities

### New Capabilities
- `spec-kit-workflow`: Ability to invoke `/speckit-constitution`, `/speckit-specify`, `/speckit-plan`, `/speckit-tasks`, `/speckit-implement`, `/speckit-converge` as Agent Smith skills, each delegating to `specify` CLI with OpenCode-compatible SKILL.md wrappers
- `spec-kit-prerequisite`: Auto-detects and installs `specify-cli` via `uv tool install specify-cli` when first `/speckit-*` command is invoked and the tool is absent

### Modified Capabilities
- None (existing SDD skills unchanged)

## Approach

1. **Spec Kit CLI detection**: Add a `speckit_available()` check in the install/setup path. If absent, prompt user or auto-run `uv tool install specify-cli` (Python 3.11+ required; `uv` must be on PATH).

2. **Skill embedding**: Clone Spec Kit's `SkillsIntegration` output — `skills/speckit-*/SKILL.md` files — into `internal/assets/skills/`. Each SKILL.md wraps the native Spec Kit command format into OpenCode skill frontmatter and execution protocol.

3. **Catalog registration**: Add `SkillSpecKit*` constants to `internal/model/types.go` and register them in `internal/catalog/skills.go` under a new `spec-kit` category.

4. **Injection update**: In `inject.go`, extend `directoryAssets()` to include `skills/speckit-*` directories alongside existing skill paths. The `IsSDDSkill()` guard (prefix `sdd-`) does not block these.

5. **Orchestrator routing**: Update `orchestrator.md` to include a "Workflow Selection" block presenting both Agent Smith ODD and Spec Kit SDD as explicit routing options.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/model/types.go` | Modified | Add `SkillSpecKitConstitution`, `SkillSpecKitSpecify`, `SkillSpecKitPlan`, `SkillSpecKitTasks`, `SkillSpecKitImplement`, `SkillSpecKitConverge`, `SkillSpecKitInstall` |
| `internal/catalog/skills.go` | Modified | Register all `speckit-*` skills under `spec-kit` category |
| `internal/components/skills/inject.go` | Modified | Add `skills/speckit-*` to `directoryAssets()` scan |
| `internal/assets/skills/speckit-*/` | New | ~6–8 SKILL.md files cloned from Spec Kit `SkillsIntegration` output |
| `internal/assets/opencode/orchestrator.md` | Modified | Add workflow selection block surfacing both SDD approaches |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Spec Kit Python dependency missing on host | Medium | Detect in setup flow; emit clear error with `uv tool install specify-cli` instruction |
| Spec Kit output format diverges from OpenCode skill contract | Low | Wrap native SKILL.md output with agent-smith skill frontmatter |
| Skill name collision with future SDD skills | Low | `speckit-` prefix is distinct from `sdd-` prefix |

## Rollback Plan

Remove `speckit-*` entries from `internal/catalog/skills.go`, `internal/model/types.go`, and `internal/assets/skills/speckit-*` directories. Revert `inject.go` to prior `directoryAssets()` implementation. Revert orchestrator prompt. Publish as a breaking change only if users have adopted `/speckit-*` commands.

## Dependencies

- Python 3.11+ on host (required by `specify-cli`)
- `uv` tool on PATH (package manager for `specify-cli`)
- `specify-cli` installed via `uv tool install specify-cli`

## Success Criteria

- [ ] `agent-smith install` installs both ODD and Spec Kit skill sets without error
- [ ] `/speckit-specify` skill loads and delegates to `specify specify` CLI
- [ ] Orchestrator prompt lists both workflows with distinct routing
- [ ] `go test ./...` passes after changes
- [ ] Skill registry refresh correctly indexes all `speckit-*` skills
