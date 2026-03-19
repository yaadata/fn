package forgedomain

import . "github.com/yaadata/optionsgo"

type IssueContext struct {
	Assignee Option[string]
	Labels   []string
	Url      Option[string]
}
