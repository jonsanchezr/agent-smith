package skills

import (
	"fmt"
	"os/exec"
)

// lookPathFunc mimics exec.LookPath for testability.
var lookPathFunc func(name string) error = func(name string) error {
	_, err := exec.LookPath(name)
	return err
}

// Commander abstracts exec.Cmd for testability. Exported so external test
// packages (e.g. internal/cli/install_speckit_test.go) can build fake
// commanders without re-declaring the interface.
type Commander interface {
	Run() error
}

// commandFunc wraps exec.Command for testability.
var commandFunc func(name string, args ...string) Commander = func(name string, args ...string) Commander {
	return exec.Command(name, args...)
}

// speckitAvailable returns (pythonAvailable, uvAvailable, specifyAvailable, error).
func speckitAvailable() (bool, bool, bool, error) {
	pythonErr := lookPathFunc("python3")
	uvErr := lookPathFunc("uv")
	specifyErr := lookPathFunc("specify")
	return pythonErr == nil, uvErr == nil, specifyErr == nil, nil
}

// SpeckitInstallAvailable is the package-public wrapper around speckitAvailable.
// The lowercase name stays unexported (kept for in-package callers and tests);
// external callers (e.g. internal/cli) use this PascalCase wrapper so the
// subsystem is reachable without leaking the internal helper.
func SpeckitInstallAvailable() (bool, bool, bool, error) {
	return speckitAvailable()
}

// Install runs: uv tool install specify-cli && specify init --integration opencode
//
// It is the exported name for what used to be speckitInstall (kept exported
// here so the cli package can drive the spec-kit runtime auto-install).
func Install() error {
	if err := commandFunc("uv", "tool", "install", "specify-cli").Run(); err != nil {
		return fmt.Errorf("uv tool install specify-cli: %w", err)
	}
	if err := commandFunc("specify", "init", "--integration", "opencode", "--non-interactive").Run(); err != nil {
		return fmt.Errorf("specify init: %w", err)
	}
	return nil
}

// RunSpeckitInstall is the package-public PascalCase wrapper around Install,
// mirroring SpeckitInstallAvailable for symmetry. External callers should
// prefer this name when both an availability probe and the installer are
// driven from the same package — it makes the speckit subsystem's public
// surface (SpeckitInstallAvailable + RunSpeckitInstall) self-describing.
func RunSpeckitInstall() error {
	return Install()
}

// SetLookPathFuncForTest swaps the package-level lookPathFunc for the
// duration of a test, auto-registering the restore on t.Cleanup. It accepts
// any *testing.T-shaped value (the looser interface matches engram's
// SetLookPathForTest convention so external tests do not have to import
// testing.T just to wire a seam).
func SetLookPathFuncForTest(t interface {
	Helper()
	Cleanup(func())
}, fn func(name string) error) {
	t.Helper()
	orig := lookPathFunc
	if fn != nil {
		lookPathFunc = fn
	}
	t.Cleanup(func() { lookPathFunc = orig })
}

// SetCommandFuncForTest swaps the package-level commandFunc for the duration
// of a test, auto-registering the restore on t.Cleanup.
func SetCommandFuncForTest(t interface {
	Helper()
	Cleanup(func())
}, fn func(name string, args ...string) Commander) {
	t.Helper()
	orig := commandFunc
	if fn != nil {
		commandFunc = fn
	}
	t.Cleanup(func() { commandFunc = orig })
}
