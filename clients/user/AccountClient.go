package user

import (
	"encoding/json"
	"errors"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/clients/notification"
	"github.com/jjamieson1/celestial-sdk/clients/tenant"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetAccount(tokenString, tenantId string) (models.User, int, error) {

	// Add the headers
	headers := map[string]string {
		"tenantId": tenantId,
		"Authorization": "bearer " + tokenString 
	}

	var user models.User

	url := "http://localhost:3000//api/v1/auth/account"
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, headers, b)
	json.Unmarshal(response, &user)
	return user, status, err
}