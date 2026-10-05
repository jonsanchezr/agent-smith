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

// commander abstracts exec.Cmd for testability.
type commander interface {
	Run() error
}

// commandFunc wraps exec.Command for testability.
var commandFunc func(name string, args ...string) commander = func(name string, args ...string) commander {
	return exec.Command(name, args...)
}

// speckitAvailable returns (pythonAvailable, uvAvailable, specifyAvailable, error).
func speckitAvailable() (bool, bool, bool, error) {
	pythonErr := lookPathFunc("python3")
	uvErr := lookPathFunc("uv")
	specifyErr := lookPathFunc("specify")
	return pythonErr == nil, uvErr == nil, specifyErr == nil, nil
}

// speckitInstall runs: uv tool install specify-cli && specify init --integration opencode
func speckitInstall() error {
	if err := commandFunc("uv", "tool", "install", "specify-cli").Run(); err != nil {
		return fmt.Errorf("uv tool install specify-cli: %w", err)
	}
	if err := commandFunc("specify", "init", "--integration", "opencode", "--non-interactive").Run(); err != nil {
		return fmt.Errorf("specify init: %w", err)
	}
	return nil
}
