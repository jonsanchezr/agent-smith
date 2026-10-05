//go:build !windows

package cli

import "github.com/jonsanchezr/agent-smith/v4/internal/model"

func usesAnchoredCompatibilityTransaction() bool {
	return false
}

func newCompatibilityRefreshTransaction(string, []model.ComponentID, model.Selection) (compatibilityRefreshTransaction, error) {
	return nil, nil
}

