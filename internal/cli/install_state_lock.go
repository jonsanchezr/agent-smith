package cli

import "github.com/jonsanchezr/agent-smith/v4/internal/statecoord"

func withInstallStateLock(homeDir string, operation func() error) error {
	return statecoord.WithLock(homeDir, operation)
}

