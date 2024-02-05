package utilities

import (
	"fmt"

	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel/session"
)

func GetSessionVars(c *session.Session) (s models.SessionModel, err error) {
	userId, err := c.Get("userId")
	if err != nil && err.Error() != "Session value not found" {
		return s, err
	}

	tenantId, err := c.Get("tenantId")
	if err != nil {
		return s, err
	}

	jwt, err := c.Get("jwt")
	if err != nil {
		return s, err
	}

	s.Jwt = fmt.Sprintf("%v", jwt)
	s.TenantId = fmt.Sprintf("%v", tenantId)
	s.UserId = fmt.Sprintf("%v", userId)
	return s, err
}
