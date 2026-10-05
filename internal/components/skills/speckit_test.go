package skills

import (
	"errors"
	"strings"
	"testing"
)

// fakeCmd mimics exec.Cmd for testability.
type fakeCmd struct {
	err error
}

func (c *fakeCmd) Run() error { return c.err }

// TestSpeckitAvailablePythonAbsent verifies speckit_available returns
// (false, _, _, nil) when Python is absent.
func TestSpeckitAvailablePythonAbsent(t *testing.T) {
	// Override lookPathFunc to simulate python3 not found.
	original := lookPathFunc
	lookPathFunc = func(name string) error {
		if name == "python3" {
			return assertAnErr{}
		}
		return nil // uv and specify found
	}
	defer func() { lookPathFunc = original }()

	python, uv, specify, err := speckitAvailable()
	if err != nil {
		t.Fatalf("speckitAvailable() error = %v", err)
	}
	if python {
		t.Errorf("speckitAvailable() python = true, want false")
	}
	// uv and specify may be true or false depending on other calls
	_ = uv
	_ = specify
}

// TestSpeckitAvailableUVAbsent verifies speckit_available returns
// (true, false, _, nil) when Python is present but uv is absent.
func TestSpeckitAvailableUVAbsent(t *testing.T) {
	original := lookPathFunc
	lookPathFunc = func(name string) error {
		if name == "uv" {
			return assertAnErr{}
		}
		return nil // python3 and specify found
	}
	defer func() { lookPathFunc = original }()

	python, uv, specify, err := speckitAvailable()
	if err != nil {
		t.Fatalf("speckitAvailable() error = %v", err)
	}
	if !python {
		t.Errorf("speckitAvailable() python = false, want true")
	}
	if uv {
		t.Errorf("speckitAvailable() uv = true, want false")
	}
	_ = specify
}

// TestSpeckitAvailableSpecifyAbsent verifies speckit_available returns
// (true, true, false, nil) when Python and uv are present but specify is absent.
func TestSpeckitAvailableSpecifyAbsent(t *testing.T) {
	original := lookPathFunc
	lookPathFunc = func(name string) error {
		if name == "specify" {
			return assertAnErr{}
		}
		return nil // python3 and uv found
	}
	defer func() { lookPathFunc = original }()

	python, uv, specify, err := speckitAvailable()
	if err != nil {
		t.Fatalf("speckitAvailable() error = %v", err)
	}
	if !python {
		t.Errorf("speckitAvailable() python = false, want true")
	}
	if !uv {
		t.Errorf("speckitAvailable() uv = false, want true")
	}
	if specify {
		t.Errorf("speckitAvailable() specify = true, want false")
	}
}

// TestSpeckitInstallSucceeds verifies speckitInstall succeeds when both commands succeed.
func TestSpeckitInstallSucceeds(t *testing.T) {
	original := commandFunc
	commandFunc = func(name string, args ...string) commander {
		return &fakeCmd{err: nil}
	}
	defer func() { commandFunc = original }()

	err := speckitInstall()
	if err != nil {
		t.Errorf("speckitInstall() error = %v, want nil", err)
	}
}

// TestSpeckitInstallFailsOnUvToolInstall verifies speckitInstall returns an error
// when uv tool install fails.
func TestSpeckitInstallFailsOnUvToolInstall(t *testing.T) {
	original := commandFunc
	callCount := 0
	commandFunc = func(name string, args ...string) commander {
		callCount++
		if callCount == 1 {
			return &fakeCmd{err: errors.New("uv tool install failed")}
		}
		return &fakeCmd{err: nil}
	}
	defer func() { commandFunc = original }()

	err := speckitInstall()
	if err == nil {
		t.Fatalf("speckitInstall() error = nil, want error containing 'uv tool install'")
	}
	if !errors.Is(err, errors.New("uv tool install failed")) {
		// Check error message contains "uv tool install"
		if !strings.Contains(err.Error(), "uv tool install") {
			t.Errorf("speckitInstall() error = %v, want error containing 'uv tool install'", err)
		}
	}
}

// TestSpeckitInstallFailsOnSpecifyInit verifies speckitInstall returns an error
// when specify init fails.
func TestSpeckitInstallFailsOnSpecifyInit(t *testing.T) {
	original := commandFunc
	callCount := 0
	commandFunc = func(name string, args ...string) commander {
		callCount++
		if callCount == 2 {
			return &fakeCmd{err: errors.New("specify init failed")}
		}
		return &fakeCmd{err: nil}
	}
	defer func() { commandFunc = original }()

	err := speckitInstall()
	if err == nil {
		t.Fatalf("speckitInstall() error = nil, want error containing 'specify init'")
	}
	if !errors.Is(err, errors.New("specify init failed")) {
		// Check error message contains "specify init"
		if !strings.Contains(err.Error(), "specify init") {
			t.Errorf("speckitInstall() error = %v, want error containing 'specify init'", err)
		}
	}
}

// assertAnErr is a sentinel error implementing error interface.
type assertAnErr struct{}

func (e assertAnErr) Error() string { return "assert error" }
