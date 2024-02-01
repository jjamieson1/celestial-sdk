package utilities

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/jjamieson1/celestial-sdk/models"
)

func GetRemoteIPAddress(req *http.Request) string {
	ip := strings.Split(strings.TrimSpace(req.Header.Get("x-forwarded-for")), ",")[0]
	if strings.Compare(ip, "") == 0 {
		i := strings.LastIndex(req.RemoteAddr, ":")
		ip = req.RemoteAddr[:i]
	}
	return ip
}

func GetIdentityIdFromJWT(payload string) string {
	if payload == "" {
		return ""
	}

	b, err := jwt.DecodeSegment(payload)
	if err != nil {
		return ""
	}

	var token models.SessionJwt
	err = json.Unmarshal(b, &token)
	if err != nil {
		return ""
	}
	return token.Subject
}
