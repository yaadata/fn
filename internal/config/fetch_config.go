package config

import (
	"time"

	. "github.com/yaadata/optionsgo"
)

type FetchConfig struct {
	Strategy FetchStrategy
	Interval Option[time.Duration]
}
