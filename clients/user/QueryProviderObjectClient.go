package user

import (
	clients "github.com/jjamieson1/celestial-sdk/clients"
)

func DoesEmailExist(tenantId string, email string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/query/email/" + email

	return clients.CallRestEndPoint(url, "GET", headers, nil)
}

func DoesUserNameExist(tenantId string, userName string) ([]byte, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/query/username" + userName

	return clients.CallRestEndPoint(url, "GET", headers, nil)
}
