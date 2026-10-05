package engram

import (
	"github.com/jonsanchezr/agent-smith/v4/internal/installcmd"
	"github.com/jonsanchezr/agent-smith/v4/internal/model"
	"github.com/jonsanchezr/agent-smith/v4/internal/system"
)

func InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, model.ComponentEngram)
}

