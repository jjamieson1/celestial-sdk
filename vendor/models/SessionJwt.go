package models

import (
	"github.com/golang-jwt/jwt"
)

type SessionJwt struct {
	MfaRequired bool `json:"mfaRequired"`
	IsMfa       bool `json:"isMfa"`
	IsDidAuth   bool `json:"isDidAuth"`
	jwt.StandardClaims
}
