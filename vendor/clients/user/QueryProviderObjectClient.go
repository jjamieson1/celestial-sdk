package user

import (
	clients2 "github.com/jjamieson1/eden-sdk/clients"

)

func DoesEmailExist(tenantId string, email string) ([]byte, int, error) {

	url := "http://localhost:9100/api/query/email/" + email

	return clients2.CallRestEndPoint(url, "GET", tenantId, nil)
}

func DoesUserNameExist(tenantId string, userName string) ([]byte, int, error) {

	url := "http://localhost:9100/api/query/username" + userName

	return clients2.CallRestEndPoint(url, "GET", tenantId, nil)
}
