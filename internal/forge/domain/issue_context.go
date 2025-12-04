package forgedomain

import (
	. "github.com/yaadata/optionsgo/core"
)

type IssueContext struct {
	Assignee Option[string]
	Labels   []string
	Url      Option[string]
}
