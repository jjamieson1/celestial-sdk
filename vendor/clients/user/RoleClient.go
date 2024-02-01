package user

import (
	"encoding/json"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetRoles(tenantId string) ([]models.Role, int, error) {

	var roles []models.Role

	url := "http://localhost:9100/api/" + "/roles"
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, tenantId, nil)
	json.Unmarshal(response, &roles)
	return roles, status, err
}
