package user

import (
	clients2 "github.com/jjamieson1/celestial-sdk/clients"
)

func AddAddress(tenantId string, userId string, body []byte) ([]byte, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/address"
	return clients2.CallRestEndPoint(url, "POST", tenantId, body)
}

func AddEmail(tenantId string, userId string, body []byte) ([]byte, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/email"
	return clients2.CallRestEndPoint(url, "POST", tenantId, body)
}

func AddPhone(tenantId string, userId string, body []byte) ([]byte, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/phone"
	return clients2.CallRestEndPoint(url, "POST", tenantId, body)
}

func AddUser(tenantId string, body []byte) ([]byte, int, error) {

	url := "http://localhost:9100/api/profile"
	return clients2.CallRestEndPoint(url, "POST", tenantId, body)
}
