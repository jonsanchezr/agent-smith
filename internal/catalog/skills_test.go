package catalog

import (
	"strings"
	"testing"

	"github.com/jonsanchezr/agent-smith/v4/internal/components/skills"
	"github.com/jonsanchezr/agent-smith/v4/internal/model"
)

// TestMVPSkillsCoverAllPresetSkills ensures every skill that presets.go would
// install is also registered in the catalog's mvpSkills allowlist. This
// prevents a future addition to sddSkills or foundationSkills from being
// silently rejected by normalizeSkills in cli/validate.go.
func TestMVPSkillsCoverAllPresetSkills(t *testing.T) {
	catalogSet := make(map[model.SkillID]bool)
	for _, s := range MVPSkills() {
		catalogSet[s.ID] = true
	}

	presetSkills := skills.AllSkillIDs()
	for _, id := range presetSkills {
		if !catalogSet[id] {
			t.Errorf("skill %q is in presets but missing from catalog mvpSkills", id)
		}
	}
}

// TestMVPSkillsNoDuplicates ensures no skill is listed twice in mvpSkills.
func TestMVPSkillsNoDuplicates(t *testing.T) {
	seen := make(map[model.SkillID]bool)
	for _, s := range MVPSkills() {
		if seen[s.ID] {
			t.Errorf("duplicate skill %q in mvpSkills", s.ID)
		}
		seen[s.ID] = true
	}
}

func TestMVPSkillsIncludeRequestedBundledSkillsWithCanonicalNames(t *testing.T) {
	required := map[model.SkillID]string{
		model.SkillCreator:             "skill-creator",
		model.SkillSkillRegistry:       "skill-registry",
		model.SkillCognitiveDoc:        "cognitive-doc-design",
		model.SkillCommentWriter:       "comment-writer",
		model.SkillJudgmentDay:         "judgment-day",
		model.SkillSDDInit:             "sdd-init",
		model.SkillSDDResearch:         "sdd-research",
		model.SkillImprover:            "skill-improver",
		model.SkillRDDDefectWorkflow:   "rdd-defect-workflow",
		model.SkillSystemicIssueTriage: "systemic-issue-triage",
		model.SkillAgentSmithBench:       "agent-smith-bench",
	}

	found := make(map[model.SkillID]string)
	for _, skill := range MVPSkills() {
		found[skill.ID] = skill.Name
		if skill.Name == "judgement-day" {
			t.Fatalf("catalog uses non-canonical spelling %q; want judgment-day", skill.Name)
		}
	}

	for id, wantName := range required {
		name, ok := found[id]
		if !ok {
			t.Fatalf("MVPSkills() missing requested bundled skill %q", id)
		}
		if name != wantName {
			t.Fatalf("MVPSkills() name for %q = %q, want %q", id, name, wantName)
		}
	}
}

// TestMVPSkillsIncludeSpecKitSkills verifies all 9 spec-kit skills are registered.
func TestMVPSkillsIncludeSpecKitSkills(t *testing.T) {
	specKitSkills := []model.SkillID{
		model.SkillSpecKitConstitution,
		model.SkillSpecKitSpecify,
		model.SkillSpecKitClarify,
		model.SkillSpecKitPlan,
		model.SkillSpecKitChecklist,
		model.SkillSpecKitTasks,
		model.SkillSpecKitAnalyze,
		model.SkillSpecKitImplement,
		model.SkillSpecKitConverge,
	}

	found := make(map[model.SkillID]bool)
	for _, skill := range MVPSkills() {
		found[skill.ID] = true
	}

	for _, id := range specKitSkills {
		if !found[id] {
			t.Errorf("MVPSkills() missing spec-kit skill %q", id)
		}
	}
}

// TestSpecKitSkillsHaveCorrectCategory verifies all spec-kit skills have spec-kit category.
func TestSpecKitSkillsHaveCorrectCategory(t *testing.T) {
	for _, skill := range MVPSkills() {
		if strings.HasPrefix(string(skill.ID), "speckit-") {
			if skill.Category != "spec-kit" {
				t.Errorf("spec-kit skill %q has category %q, want spec-kit", skill.ID, skill.Category)
			}
			if skill.Priority != "p0" {
				t.Errorf("spec-kit skill %q has priority %q, want p0", skill.ID, skill.Priority)
			}
		}
	}
}

