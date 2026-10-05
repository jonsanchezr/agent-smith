# Design: spec-kit-integration

## Technical Approach

Embed Spec Kit (github.com/jonsanchezr/spec-kit) as an alternative SDD workflow. `speckit-*` SKILL.md files are embedded Go assets under `internal/assets/skills/speckit-*/`, registered alongside existing skills, with installation detection for Python + uv + specify-cli. The orchestrator surfaces both SDD paths as explicit routing options.

## Architecture Decisions

### Decision: Skill ID prefix `speckit-` rather than `spec-kit-`

**Choice**: `speckit-` prefix on all skill IDs
**Alternatives considered**: `spec-kit-`, `specify-`
**Rationale**: `speckit-` mirrors the project name "spec-kit" in a compact form. `IsSDDSkill()` guard (prefix `sdd-`) does not block these — no code change needed to permit non-`sdd-` skills.

### Decision: Embedded SKILL.md files, not dynamic fetch

**Choice**: Bundle SKILL.md files as Go embed assets
**Alternatives considered**: Fetch from spec-kit repo at install time
**Rationale**: `assets.FS` already embeds `all:skills`. Adding `speckit-*/` directories requires no change to `assets.go` or `go:embed`. Offline/air-gapped install works without network.

### Decision: `uv tool install specify-cli` for installation

**Choice**: Delegate installation to `uv` exclusively
**Alternatives considered**: `pip install`, direct wheel download, bundling Python
**Rationale**: Proposal/spec requires `uv`. Consistent with spec-kit's own tooling. `uv tool` installs to user-local dir without sudo.

### Decision: Spec Kit skill catalog under `spec-kit` category

**Choice**: New `spec-kit` category in catalog, parallel to `sdd` and `workflow`
**Alternatives considered**: Merge into `sdd` category
**Rationale**: Two distinct workflows should have distinct categories. Users see both in skill listings without confusion.

## Data Flow

```
User invokes /speckit-specify
        │
        ▼
orchestrator routes to speckit-specify skill
        │
        ▼
speckit_available() checks Python + uv + specify-cli
        │
   ┌────┴────┐
   │ Python  │ uv    │ specify-cli │
   │ absent  │ absent│ absent      │
   └────┬────┴────┬──┴─────────────┘
        │         │         │
        ▼         ▼         ▼
   error msg  error msg  auto-install via
   with hint  with hint  uv tool install
                              │
                              ▼
                    load SKILL.md from
                    internal/assets/skills/speckit-specify/
                              │
                              ▼
                    delegate to `specify specify` CLI
```

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `internal/model/types.go` | Modify | Add `SkillSpecKitConstitution`, `SkillSpecKitSpecify`, `SkillSpecKitPlan`, `SkillSpecKitChecklist`, `SkillSpecKitTasks`, `SkillSpecKitAnalyze`, `SkillSpecKitImplement`, `SkillSpecKitConverge`, `SkillSpecKitInstall` |
| `internal/catalog/skills.go` | Modify | Add `speckit-*` entries to `mvpSkills` under category `"spec-kit"` |
| `internal/assets/skills/speckit-constitution/SKILL.md` | Create | Wrapper for `specify constitution` |
| `internal/assets/skills/speckit-specify/SKILL.md` | Create | Wrapper for `specify specify` |
| `internal/assets/skills/speckit-plan/SKILL.md` | Create | Wrapper for `specify plan` |
| `internal/assets/skills/speckit-checklist/SKILL.md` | Create | Wrapper for `specify checklist` |
| `internal/assets/skills/speckit-tasks/SKILL.md` | Create | Wrapper for `specify tasks` |
| `internal/assets/skills/speckit-analyze/SKILL.md` | Create | Wrapper for `specify analyze` |
| `internal/assets/skills/speckit-implement/SKILL.md` | Create | Wrapper for `specify implement` |
| `internal/assets/skills/speckit-converge/SKILL.md` | Create | Wrapper for `specify converge` |
| `internal/assets/skills/speckit-install/SKILL.md` | Create | Installation prerequisite check and auto-install |
| `internal/assets/opencode/orchestrator.md` | Modify | Add "Available Workflows" section with both SDD paths |

## Interfaces / Contracts

### SkillID constants (types.go)

```go
const (
    SkillSpecKitConstitution SkillID = "speckit-constitution"
    SkillSpecKitSpecify     SkillID = "speckit-specify"
    SkillSpecKitPlan        SkillID = "speckit-plan"
    SkillSpecKitChecklist   SkillID = "speckit-checklist"
    SkillSpecKitTasks       SkillID = "speckit-tasks"
    SkillSpecKitAnalyze     SkillID = "speckit-analyze"
    SkillSpecKitImplement   SkillID = "speckit-implement"
    SkillSpecKitConverge    SkillID = "speckit-converge"
    SkillSpecKitInstall     SkillID = "speckit-install"
)
```

### Catalog entry shape (skills.go)

```go
{ID: model.SkillSpecKitSpecify, Name: "speckit-specify", Category: "spec-kit", Priority: "p0"},
```

### `speckit_available()` signature

```go
// speckit_available returns (python3 present, uv present, specify-cli present, error)
func speckit_available() (bool, bool, bool, error)

// speckit_install attempts uv tool install specify-cli
func speckit_install() error
```

### Orchestrator workflow block (orchestrator.md)

```markdown
## Available Workflows

Two parallel SDD workflows are available:

**Agent Smith ODD** — homegrown ODD workflow:
sdd-init → sdd-explore → sdd-propose → sdd-spec → sdd-design → sdd-tasks → sdd-apply → sdd-verify → sdd-archive

**Spec Kit SDD** — spec-kit integration:
speckit-constitution → speckit-specify → speckit-clarify → speckit-plan → speckit-checklist → speckit-tasks → speckit-analyze → speckit-implement → speckit-converge

Route to the chosen workflow based on user intent. Both coexist; user selects per-task.
```

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | `speckit_available()` with mocked exec | Test Python absent, uv absent, specify-cli absent, all present |
| Unit | Catalog registration | Verify all 9 SkillSpecKit* constants present in MVPSkills() |
| Unit | `IsSDDSkill()` not blocking speckit IDs | Assert `IsSDDSkill("speckit-specify") == false` |
| Integration | `InjectDirectoryWithCapability` writes speckit skills | Call with speckit skill IDs, verify files created |
| E2E | Full install flow | `agent-smith install` with mocked uv; verify SKILL.md files written |

## Threat Matrix

| Boundary | Applicability | Design response | Planned RED tests |
|---|---|---|---|
| Shell command execution | **Applicable**: `uv tool install specify-cli` runs a subprocess | Detect Python + uv first; run installation only if both present | `speckit_available()` returns false when python/uv absent; error message emitted |
| Executable on PATH | **Applicable**: `specify` must be on PATH after install | `speckit_available()` checks `exec.LookPath("specify")` after install | `speckit_available()` returns false when specify not in PATH; hint to uv users |
| Offline/air-gapped install | **Applicable**: spec-kit install attempted without network | Detect network unavailability; emit offline hint with manual install steps | `speckit_available()` returns (true, true, false) offline; specific error message |

N/A — no routing, VCS/PR automation, or executable-file classification boundary applies beyond the shell/process integration above.

## Migration / Rollout

No migration required. New skills are additive. Existing `sdd-*` skills unchanged. If spec-kit install fails, the `speckit-*` skills simply won't be available — other Agent Smith functionality is unaffected.

## Open Questions

- [ ] Should `speckit-install` skill run automatically on first `/speckit-*` invocation, or require explicit user call to `/speckit-install` first?
- [ ] Do we version-pin the spec-kit CLI (e.g., `uv tool install specify-cli==0.3.0`) or always install latest?
- [ ] Should `specify init --integration opencode` be run as part of auto-install, or only on user request?
