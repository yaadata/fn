package forgedomain

import (
	. "github.com/yaadata/optionsgo"
)

type Repository struct {
	Group   string
	Project string
	Url     Option[string]
}
