package config

import . "github.com/yaadata/optionsgo"

type AccessToken string

type Credential struct {
	AccessToken      Option[AccessToken]
	UsernamePassword Option[UsernamePassword]
}

type UsernamePassword struct {
	Username string
	Password string
}
