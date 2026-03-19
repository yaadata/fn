package config

import (
	forgedomain "github.com/yaadata/fn/internal/forge/domain"
)

type Forge struct {
	Name   forgedomain.Provider
	Access forgedomain.ForgeAccess
}
