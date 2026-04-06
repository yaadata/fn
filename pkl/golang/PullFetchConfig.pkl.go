// Code generated from Pkl module `fn.pkl`. DO NOT EDIT.
package golang

import "github.com/apple/pkl-go/pkl"

type PullFetchConfig interface {
	FetchConfig

	GetInterval() pkl.Duration
}

var _ PullFetchConfig = PullFetchConfigImpl{}

type PullFetchConfigImpl struct {
	Strategy string `pkl:"strategy"`

	Interval pkl.Duration `pkl:"interval"`
}

func (rcv PullFetchConfigImpl) GetStrategy() string {
	return rcv.Strategy
}

func (rcv PullFetchConfigImpl) GetInterval() pkl.Duration {
	return rcv.Interval
}
