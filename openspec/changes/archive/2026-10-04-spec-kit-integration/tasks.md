# Tasks: spec-kit-integration

## Review Workload Forecast

| Field | Value |
|-------|-------|
| Estimated changed lines | 650–1000 |
| 400-line budget risk | High |
| Chained PRs recommended | Yes |
| Suggested split | PR 1 → PR 2 → PR 3 → PR 4 |
| Delivery strategy | ask-on-risk |
| Chain strategy | pending |

Decision needed before apply: Yes
Chained PRs recommended: Yes
Chain strategy: stacked-to-main|feature-branch-chain|size-exception|pending
400-line budget risk: High

### Suggested Work Units

| Unit | Goal | Likely PR | Focused test command | Runtime harness | Rollback boundary |
|------|------|-----------|----------------------|-----------------|-------------------|
| 1 | Model + Catalog additions | PR 1 | `go build ./...` | N/A (foundational types) | Drop SkillID constants + catalog entries |
| 2 | 9 SKILL.md asset files | PR 2 | `go test ./internal/assets/...` | N/A (static assets) | Delete speckit-* directories |
| 3 | Orchestrator routing + Auto-install | PR 3 | `go build ./...` | `agent-smith install` (mocked) | Revert orchestrator.md + install code |
| 4 | Tests for all new code | PR 4 | `go test ./...` | `go test ./internal/components/skills/...` | Delete test files |

## Phase 1: Model & Catalog (no dependencies)

- [x] 1.1 Add `SkillSpecKitConstitution`, `SkillSpecKitSpecify`, `SkillSpecKitClarify`, `SkillSpecKitPlan`, `SkillSpecKitChecklist`, `SkillSpecKitTasks`, `SkillSpecKitAnalyze`, `SkillSpecKitImplement`, `SkillSpecKitConverge` constants to `internal/model/types.go` after line 146
- [x] 1.2 Register all 9 speckit skills in `internal/catalog/skills.go` under `// Spec Kit skills` section with `Category: "spec-kit"`, `Priority: "p0"`

## Phase 2: Skill Assets (depends on Phase 1)

- [x] 2.1 Create `internal/assets/skills/speckit-constitution/SKILL.md` — delegates to `specify constitution`
- [x] 2.2 Create `internal/assets/skills/speckit-specify/SKILL.md` — delegates to `specify specify`
- [x] 2.3 Create `internal/assets/skills/speckit-clarify/SKILL.md` — delegates to `specify clarify`
- [x] 2.4 Create `internal/assets/skills/speckit-plan/SKILL.md` — delegates to `specify plan`
- [x] 2.5 Create `internal/assets/skills/speckit-checklist/SKILL.md` — delegates to `specify checklist`
- [x] 2.6 Create `internal/assets/skills/speckit-tasks/SKILL.md` — delegates to `specify tasks`
- [x] 2.7 Create `internal/assets/skills/speckit-analyze/SKILL.md` — delegates to `specify analyze`
- [x] 2.8 Create `internal/assets/skills/speckit-implement/SKILL.md` — delegates to `specify implement`
- [x] 2.9 Create `internal/assets/skills/speckit-converge/SKILL.md` — delegates to `specify converge`
- [x] 2.10 Each SKILL.md handles auto-installation check: call `speckit_available()` on first use; if false, run `speckit_install()` before delegating

## Phase 3: Orchestrator Routing (depends on Phase 2)

- [x] 3.1 Update `internal/assets/skills/_shared/odd-orchestrator-sections.md` — add "Available Workflows" section showing both SDD and Spec Kit paths
- [x] 3.2 Spec Kit path: `speckit-constitution → speckit-specify → speckit-clarify → speckit-plan → speckit-checklist → speckit-tasks → speckit-analyze → speckit-implement → speckit-converge`
- [x] 3.3 Agent Smith SDD path: `sdd-init → sdd-explore → sdd-propose → sdd-spec → sdd-design → sdd-tasks → sdd-apply → sdd-verify → sdd-archive`

## Phase 4: Auto-Install Logic (depends on Phase 1)

- [x] 4.1 Create `internal/components/skills/speckit.go` with `speckit_available()` function
- [x] 4.2 `speckit_available()` checks: Python 3.11+ via `exec.LookPath("python3")`, `uv` via `exec.LookPath("uv")`, `specify` via `exec.LookPath("specify")`
- [x] 4.3 Create `speckit_install()` function that runs `uv tool install specify-cli` then `specify init --integration opencode`
- [x] 4.4 Error handling: clear messages if Python/uv missing, retry after install, offline/air-gapped hint
- [x] 4.5 RED test: `speckit_available()` returns `(false, false, false, nil)` when Python absent
- [x] 4.6 RED test: `speckit_available()` returns `(true, false, false, nil)` when uv absent
- [x] 4.7 RED test: `speckit_available()` returns `(true, true, false, nil)` when specify absent but Python+uv present

## Phase 5: Testing

- [x] 5.1 Add unit tests for `speckit_available()` with mocked exec (Python absent, uv absent, specify absent, all present)
- [x] 5.2 Add unit tests for catalog registration — verify all 9 `SkillSpecKit*` constants present in `MVPSkills()`
- [x] 5.3 Add unit test: `IsSDDSkill("speckit-specify")` returns `false`
- [x] 5.4 Add integration test: `InjectDirectoryWithCapability` with speckit skill IDs writes files to correct directories
- [ ] 5.5 Add E2E install flow test with mocked `uv tool install` — verify SKILL.md files appear in destination **[PENDING FOLLOW-UP: requires injectable exec.Command pattern in speckit.go]**

(End of file - total 68 lines)
