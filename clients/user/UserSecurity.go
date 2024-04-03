package user

import (
	"encoding/json"
	"errors"
	"fmt"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	notificationClient "github.com/jjamieson1/celestial-sdk/clients/messaging"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
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

func FindTenantIdByEmail(email string) (tenant map[string]string, err error) {

	url := "http://localhost:3000/api/v1/auth/tenant/" + email
	method := "GET"
	response, status, err := clients.CallRestEndPoint(url, method, nil, nil)
	if err != nil {
		return tenant, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return tenant, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &tenant)
	return tenant, err
}

func Authenticate(identifier, password string, tenantId string) (au models.AuthenticatedUser, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	var body models.LoginCredential
	body.Password = password
	body.EmailOrUserName = identifier

	b, _ := json.Marshal(body)

	url := "http://localhost:3000/api/v1/auth/authenticate"
	method := "POST"

	response, status, err := clients.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return au, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return au, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &au)
	return au, err
}

func CheckToken(token string) (result interface{}, err error) {
	url := "http://localhost:3000/api/v1/auth/authenticate/" + token

	response, status, err := clients.CallRestEndPoint(url, "GET", nil, nil)
	if status != 200 {
		return result, fmt.Errorf("error authenticating token with http status: %v", status)
	}
	if err != nil {
		return result, fmt.Errorf("error authenticating token with error: %v", err.Error())
	}

	err = json.Unmarshal(response, &result)
	return result, err
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
		err := notificationClient.SendViaSMTP(email, email, "Your login details", token, 5)
		if err != nil {
			return nil, 500, err
		}
		return nil, 200, err
	}
	return nil, 400, err
}

func GetRolesByToken(token, tenantId string) (roles []models.Role, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + token,
	}
	url := "http://localhost:3000/api/v1/auth/account/roles/token"

	response, status, err := clients.CallRestEndPoint(url, "GET", headers, nil)
	if err != nil {
		return roles, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return roles, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &roles)

	revel.AppLog.Debugf("GetRolesByToken: found roles: %+v", roles)
	return roles, err
}
