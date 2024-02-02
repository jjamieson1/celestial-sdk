package user

import (
	clients "github.com/jjamieson1/celestial-sdk/clients"
)

func UpdateAddress(tenantId string, userId string, addressId string, body []byte) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/address/" + addressId

	return clients.CallRestEndPoint(url, "PUT", headers, body)
}

func UpdatePhone(tenantId string, userId string, phoneId string, body []byte) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients.CallRestEndPoint(url, "PUT", headers, body)
}

func UpdateEmail(tenantId string, userId string, emailId string, body []byte) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/email/" + emailId

	return clients.CallRestEndPoint(url, "PUT", headers, body)
}

func UpdateUser(tenantId, userId string, body []byte) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId

	return clients.CallRestEndPoint(url, "PUT", headers, body)
}
