package config

import (
	"github.com/yaadata/fn/internal/forge/client"
	forgedomain "github.com/yaadata/fn/internal/forge/domain"
	. "github.com/yaadata/optionsgo"
)

type Forge struct {
	Name   forgedomain.Provider
	Client Option[client.Forge]
}
