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

	businessUrl, err := c.Get("businessUrl")
	if err != nil {
		return s, err
	}

	s.BusinessId = fmt.Sprintf("%v", businessId)
	s.UserId = fmt.Sprintf("%v", userId)
	s.BusinessUrl = fmt.Sprintf("%v", businessUrl)
	return s, err
}
