package user

import (
	"encoding/json"
	"errors"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/clients/notification"
	"github.com/jjamieson1/celestial-sdk/clients/tenant"
	"github.com/jjamieson1/celestial-sdk/models"
)

func Authenticate(identifier, password string, tenantId string) (models.AuthenticatedUser, int, error) {
	headers := map[string]string {
		"tenantId": tenantId,
	}
	var au models.AuthenticatedUser

	var body models.LoginCredential
	body.Password = password
	body.EmailOrUserName = identifier

	b, _ := json.Marshal(body)

	url := "http://localhost:3000/api/v1/auth/authenticate"
	method := "POST"

	response, status, err := clients.CallRestEndPoint(url, method, headers, b)
	json.Unmarshal(response, &au)
	return au, status, err
}

func CheckToken(email, token, tenantId string) (models.AuthenticatedUser, int, error) {
	headers := map[string]interface{} {
		"tenantId": tenantId,
		"Authorization": "bearer " + token 
	}
	url := "http://localhost:3000/api/v1/auth/authenticate/" + token
	var result interface{}
	response, status, err := clients.CallRestEndPoint(url, "GET", headers, b)
	json.Unmarshal(response, &result)
	return result, status, err
}

func BeginEmailAuth(email, tenantId string) ([]byte, int, error) {

	response, status, err := DoesEmailExist(tenantId, email)
	if err != nil {
		err = errors.New(err.Error())
		return nil, 500, err
	}
	if status != 200 {
		err = errors.New("email does not exist")
		return nil, 400, err
	}

	url := "http://localhost:9100/api/security/authenticate/token/create/" + email

	response, status, err = clients.CallRestEndPoint(url, "POST", tenantId, nil)

	if status == 200 {
		token := make(map[string]string)
		json.Unmarshal(response, &token)
		err := notification.SendViaSMTP(email, email, "Your login details", token, 5)
		if err != nil {
			return nil, 500, err
		}
		return nil, 200, err
	}
	return nil, 400, err
}
