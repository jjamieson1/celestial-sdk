package user

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt"
	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/clients/notification"
	"github.com/jjamieson1/celestial-sdk/models"
)

func AuthenticateWithKeys(tenantId, appKey, apiKey string) (au models.AuthenticatedUser, status int, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
		"appKey":   appKey,
		"apiKey":   apiKey,
	}
	url := "http://localhost:3000/api/v1/auth/authenticate"
	method := "GET"
	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	json.Unmarshal(response, &au)
	return au, status, err
}

func FindTenantIdByEmail(email string) (tenant map[string]string, status int, err error) {

	url := "http://localhost:3000/api/v1/auth/tenant/" + email
	method := "GET"
	response, status, err := clients.CallRestEndPoint(url, method, nil, nil)
	if err != nil {
		return tenant, status, err
	}
	err = json.Unmarshal(response, &tenant)
	return tenant, status, err
}

func Authenticate(identifier, password string, tenantId string) (models.AuthenticatedUser, int, error) {
	headers := map[string]string{
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

func CheckToken(token string) (jwt.Token, int, error) {
	url := "http://localhost:3000/api/v1/auth/authenticate/" + token
	var result jwt.Token
	response, status, err := clients.CallRestEndPoint(url, "GET", nil, nil)
	json.Unmarshal(response, &result)
	return result, status, err
}

func BeginEmailAuth(email, tenantId string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

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

	response, status, err = clients.CallRestEndPoint(url, "POST", headers, nil)

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

func GetRolesByToken(token, tenantId string) (roles []models.Role, status int, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + token,
	}
	url := "http://localhost:3000/api/v1/auth/account/roles/token"

	response, status, err := clients.CallRestEndPoint(url, "GET", headers, nil)
	json.Unmarshal(response, &roles)
	return roles, status, err
}
