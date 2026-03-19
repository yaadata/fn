package config

import (
	"github.com/yaadata/fn/internal/forge/client"
	forgedomain "github.com/yaadata/fn/internal/forge/domain"
	. "github.com/yaadata/optionsgo"
)

type ProviderConfig struct {
	Name   forgedomain.Provider
	Access forgedomain.ForgeAccess
	Client Option[client.Forge]
	Fetch  FetchConfig
}
