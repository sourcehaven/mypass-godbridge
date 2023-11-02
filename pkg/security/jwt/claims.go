package jwt

import "github.com/golang-jwt/jwt/v5"

type Fresh bool
type Refresh bool
type Payload map[string]string

type Claim struct {
	Fresh   Fresh   `json:"fre"` // signals whether this is a fresh or non-fresh token
	Refresh Refresh `json:"ref"` // signals whether this jwt is used for token refreshing or not
	Payload Payload `json:"pyl"` // any non-standard stuff goes can go here
	jwt.RegisteredClaims
}
