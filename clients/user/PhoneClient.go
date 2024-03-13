package user

import (
	clients "github.com/jjamieson1/celestial-sdk/clients"
)

func AddPhone(tenantId string, userId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}

func UpdatePhone(tenantId string, userId string, phoneId string, body []byte) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients.CallRestEndPoint(url, "PUT", headers, body)
}

func DeletePhone(tenantId, userId, phoneId string) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients.CallRestEndPoint(url, "DELETE", headers, nil)

}
