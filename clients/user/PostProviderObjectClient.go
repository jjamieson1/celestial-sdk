package user

import (
	clients "github.com/jjamieson1/celestial-sdk/clients"
)

func AddAddress(tenantId string, userId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/address"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}

func AddEmail(tenantId string, userId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/email"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}

func AddPhone(tenantId string, userId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}

func AddUser(tenantId string, body []byte) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile"
	return clients.CallRestEndPoint(url, "POST", headers, body)
}
