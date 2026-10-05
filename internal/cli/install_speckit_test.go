package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jonsanchezr/agent-smith/v4/internal/components/skills"
	"github.com/jonsanchezr/agent-smith/v4/internal/model"
	"github.com/jonsanchezr/agent-smith/v4/internal/system"
)

// TestInstallWritesSpecKitSkillsAndTriggersRuntimeInstall exercises the
// speckit-runtime auto-install path end-to-end. It installs with the default
// preset (full-gentleman) so every speckit-* skill is in the resolved
// selection, then asserts:
//
//  1. The 9 speckit-* SKILL.md files were written to the agent's skills dir.
//  2. The spec-kit runtime install step fired uv tool install specify-cli and
//     specify init --integration opencode (mocked via skills.SetCommandFuncForTest).
//  3. RunInstall returns no error.
//
// The test never touches a real user config dir: it pins osUserHomeDir to a
// t.TempDir() and overrides cmdLookPath so engram/gga are treated as missing
// (matching the production sandbox for tests).
func TestInstallWritesSpecKitSkillsAndTriggersRuntimeInstall(t *testing.T) {
	home := t.TempDir()

	// Sandbox osUserHomeDir so the cli package resolves ~ inside t.TempDir().
	origHome := osUserHomeDir
	osUserHomeDir = func() (string, error) { return home, nil }
	t.Cleanup(func() { osUserHomeDir = origHome })

	// Treat engram/gga as missing so the rest of the pipeline does not shell
	// out for them. cmdLookPath does NOT need to know about python3/uv/specify
	// — those probes go through the per-test skills lookPathFunc.
	origCmdLookPath := cmdLookPath
	cmdLookPath = missingBinaryLookPath
	t.Cleanup(func() { cmdLookPath = origCmdLookPath })

	// Stage the spec-kit runtime: python3 and uv are present, specify is
	// not, so the auto-install path runs.
	skills.SetLookPathFuncForTest(t, func(name string) error {
		switch name {
		case "python3", "uv":
			return nil
		default:
			return errorNotFound()
		}
	})

	// Capture the subprocess invocations instead of letting them touch the
	// real network / filesystem.
	var uvToolCalls, specifyInitCalls int
	skills.SetCommandFuncForTest(t, func(name string, args ...string) skills.Commander {
		switch name {
		case "uv":
			uvToolCalls++
			return &fakeSpeckitCmd{err: nil}
		case "specify":
			specifyInitCalls++
			return &fakeSpeckitCmd{err: nil}
		default:
			return &fakeSpeckitCmd{err: errors.New("unexpected command: " + name)}
		}
	})

	// Mock the runCommand seam (used by ComponentEngram / ComponentGGA) so
	// those components short-circuit without shelling out.
	origRunCommand := runCommand
	runCommand = func(string, ...string) error { return nil }
	t.Cleanup(func() { runCommand = origRunCommand })

	if _, err := RunInstall([]string{"--agent", "opencode", "--preset", "full-gentleman"}, system.DetectionResult{}); err != nil {
		t.Fatalf("RunInstall() error = %v, want nil", err)
	}

	if uvToolCalls != 1 {
		t.Errorf("uv tool install invocations = %d, want 1", uvToolCalls)
	}
	if specifyInitCalls != 1 {
		t.Errorf("specify init invocations = %d, want 1", specifyInitCalls)
	}

	// Assert every speckit-* SKILL.md landed at the opencode adapter's skills
	// dir under our sandbox home. The adapter writes to <home>/.config/opencode/skills.
	wantIDs := []model.SkillID{
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
	for _, id := range wantIDs {
		path := filepath.Join(home, ".config", "opencode", "skills", string(id), "SKILL.md")
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected speckit SKILL.md at %s, stat error = %v", path, err)
		}
	}
}

// TestInstallSpecKitRuntimeInstallFailureIsSurfaced asserts that a real
// failure inside the spec-kit runtime install bubbles up as a wrapped error
// from RunInstall, rather than being swallowed silently. This keeps the
// "fail-loud" contract the rest of the install pipeline already honours.
func TestInstallSpecKitRuntimeInstallFailureIsSurfaced(t *testing.T) {
	home := t.TempDir()

	origHome := osUserHomeDir
	osUserHomeDir = func() (string, error) { return home, nil }
	t.Cleanup(func() { osUserHomeDir = origHome })

	origCmdLookPath := cmdLookPath
	cmdLookPath = missingBinaryLookPath
	t.Cleanup(func() { cmdLookPath = origCmdLookPath })

	skills.SetLookPathFuncForTest(t, func(name string) error {
		switch name {
		case "python3", "uv":
			return nil
		default:
			return errorNotFound()
		}
	})

	var uvToolCalls, specifyInitCalls int
	skills.SetCommandFuncForTest(t, func(name string, args ...string) skills.Commander {
		switch name {
		case "uv":
			uvToolCalls++
			return &fakeSpeckitCmd{err: nil}
		case "specify":
			specifyInitCalls++
			return &fakeSpeckitCmd{err: errors.New("specify init exploded in test")}
		default:
			return &fakeSpeckitCmd{err: errors.New("unexpected command: " + name)}
		}
	})

	origRunCommand := runCommand
	runCommand = func(string, ...string) error { return nil }
	t.Cleanup(func() { runCommand = origRunCommand })

	_, err := RunInstall([]string{"--agent", "opencode", "--preset", "full-gentleman"}, system.DetectionResult{})
	if err == nil {
		t.Fatal("RunInstall() error = nil, want wrapped specify-init error")
	}
	if uvToolCalls != 1 {
		t.Errorf("uv tool calls = %d, want 1 before specify init failure", uvToolCalls)
	}
	if specifyInitCalls != 1 {
		t.Errorf("specify init calls = %d, want 1 (the failing call)", specifyInitCalls)
	}
	if !strings.Contains(err.Error(), "specify init") && !strings.Contains(err.Error(), "uv tool install") {
		t.Errorf("RunInstall() error = %v, want an error mentioning the spec-kit subprocess failure", err)
	}
}

// fakeSpeckitCmd mirrors skills.fakeCmd but is defined in the cli test
// package because the production skills.fakeCmd is unexported.
type fakeSpeckitCmd struct {
	err error
}

func (c *fakeSpeckitCmd) Run() error { return c.err }

// errorNotFound stands in for exec.ErrNotFound so the speckit availability
// probe sees a clean "not found" instead of a generic error string. The
// production code only checks `err == nil`, so the concrete sentinel value
// does not matter — what matters is that the production helper treats a
// non-nil return as "absent". Using a plain errors.New sentinel would also
// work; this named helper just makes the intent self-documenting.
func errorNotFound() error { return errors.New("executable file not found in $PATH") }
