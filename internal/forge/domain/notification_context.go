package forgedomain

import (
	. "github.com/yaadata/optionsgo"
)

type NotificationContext struct {
	Url         Option[string]
	Description Option[string]
}
