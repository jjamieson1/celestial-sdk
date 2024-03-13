package user

import (
	"encoding/json"
	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func AddUpdateEmail(email models.Email, tokenString, tenantId string) (models.Email, int, error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/email"
	method := "POST"
	body, err := json.Marshal(email)
	if err != nil {
		revel.AppLog.Errorf("unable to marshal email with error: %v", err.Error())
	}
	response, status, err := clients.CallRestEndPoint(url, method, headers, body)
	if err != nil {
		revel.AppLog.Errorf("add/update email client responded with  error: %v", err.Error())
	}
	var e models.Email
	err = json.Unmarshal(response, &e)
	return e, status, err
}

func DeleteEmail(tenantId, tokenString, emailId string) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/email/" + emailId

	return clients.CallRestEndPoint(url, "DELETE", headers, nil)

}

func DoesEmailExist(tenantId string, email string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/query/email/" + email

	return clients.CallRestEndPoint(url, "GET", headers, nil)
}

func GetEmailById(emailId, tokenString, tenantId string) (email models.Email, status int, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/email/" + emailId
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)

	err = json.Unmarshal(response, &email)
	return email, status, err
}
