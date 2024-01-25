package user

import (
	"encoding/json"
	"errors"

	clients "github.com/jjamieson1/eden-sdk/clients"
	"github.com/jjamieson1/eden-sdk/clients/notification"
	"github.com/jjamieson1/eden-sdk/clients/tenant"
	"github.com/jjamieson1/eden-sdk/models"
)

func Authenticate(username, password string, tenantId string) (models.AuthenticatedUser, int, error) {

	var au models.AuthenticatedUser

	var body models.ChallengeCredential
	body.Password = password
	body.UserName = username

	b, _ := json.Marshal(body)

	url := "http://localhost:9100/api/security/authenticate"
	method := "POST"

	response, status, err := clients.CallRestEndPoint(url, method, tenantId, b)
	json.Unmarshal(response, &au)
	return au, status, err
}

func AuthenticateWithToken(email, token, tenantId string) (models.AuthenticatedUser, int, error) {
	url := "http://localhost:9100/api/security/authenticate/token/verify"
	var au models.AuthenticatedUser

	body := make(map[string]string)
	body["email"] = email
	body["token"] = token
	b, _ := json.Marshal(body)

	response, status, err := clients.CallRestEndPoint(url, "POST", tenantId, b)
	json.Unmarshal(response, &au)
	return au, status, err
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

func UpdateCredentials(userId, password, jwt, tenantId string) error {

	selectedUserProvider, err := tenant.GetProvidersForTenantByType(tenantId, "user")
	if err != nil {
		return err
	}

	var body models.ChallengeCredential
	body.Password = password
	body.UserId = userId
	body.Jwt = jwt

	b, _ := json.Marshal(body)

	url := selectedUserProvider[0].EdenAdapter.AdapterUrl + "/security/credentials"
	method := "POST"

	_, _, err = clients.CallRestEndPoint(url, method, tenantId, b)
	return err
}
