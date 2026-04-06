// Code generated from Pkl module `fn.pkl`. DO NOT EDIT.
package golang

import "github.com/yaadata/fn/pkl/golang/provider"

type ProviderConfig struct {
	Name provider.Provider `pkl:"name"`

	AccessScheme string `pkl:"accessScheme"`

	Fetch FetchConfig `pkl:"fetch"`
}
