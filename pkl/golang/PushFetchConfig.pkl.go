// Code generated from Pkl module `fn.pkl`. DO NOT EDIT.
package golang

type PushFetchConfig interface {
	FetchConfig
}

var _ PushFetchConfig = PushFetchConfigImpl{}

type PushFetchConfigImpl struct {
	Strategy string `pkl:"strategy"`
}

func (rcv PushFetchConfigImpl) GetStrategy() string {
	return rcv.Strategy
}
