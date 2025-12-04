package forgedomain

import (
	. "github.com/yaadata/optionsgo"
)

type MergeProposalContext struct {
	Author string
	Body   string
	Labels []string
	Url    Option[string]
}
