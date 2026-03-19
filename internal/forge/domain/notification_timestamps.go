package forgedomain

import (
	"time"

	. "github.com/yaadata/optionsgo"
)

type NotificationTimestamps struct {
	CreatedAt time.Time
	UpdatedAt Option[time.Time]
}
