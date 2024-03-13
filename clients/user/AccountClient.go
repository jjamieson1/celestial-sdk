package user

import (
	"encoding/json"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func AddUser(tenantId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}

func UpdateUser(tenantId, tokenString string, user models.User) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}

	body, err := json.Marshal(user)
	if err != nil {
		revel.AppLog.Errorf("error encoding user with error: %v", err.Error())
	}

	url := "http://localhost:3000/api/v1/auth/account/" + user.UserId
	return clients.CallRestEndPoint(url, "PUT", headers, body)
}

func GetAccount(tokenString, tenantId string) (models.User, int, error) {

	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}

	var user models.User

	url := "http://localhost:3000/api/v1/auth/account"
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	json.Unmarshal(response, &user)
	return user, status, err
}

func GetUserProfile(tenantId string, userId string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/profile/" + userId
	return clients.CallRestEndPoint(url, "GET", headers, nil)
}

func GetUserProfileByUserName(tenantId string, userName string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/profile/username" + userName
	return clients.CallRestEndPoint(url, "GET", headers, nil)
}

func GetAllUsers(tenantId string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/profile"

	return clients.CallRestEndPoint(url, "GET", headers, nil)
}

func DoesUserNameExist(tenantId string, userName string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/query/username" + userName

	return clients.CallRestEndPoint(url, "GET", headers, nil)
}
