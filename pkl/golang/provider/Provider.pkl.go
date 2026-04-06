// Code generated from Pkl module `fn.pkl`. DO NOT EDIT.
package provider

import (
	"encoding"
	"fmt"
)

type Provider string

const (
	Github   Provider = "github"
	Codeberg Provider = "codeberg"
)

// String returns the string representation of Provider
func (rcv Provider) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Provider)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Provider.
func (rcv *Provider) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "github":
		*rcv = Github
	case "codeberg":
		*rcv = Codeberg
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Provider`, str)
	}
	return nil
}
