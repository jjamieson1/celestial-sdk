package user

import (
	clients "github.com/jjamieson1/celestial-sdk/clients"
)

func DeleteAddress(tenantId, userId, addressId string) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/address/" + addressId

	return clients.CallRestEndPoint(url, "DELETE", headers, nil)
}

func DeletePhone(tenantId, userId, phoneId string) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients.CallRestEndPoint(url, "DELETE", headers, nil)

}

func DeleteEmail(tenantId, userId, emailId string) (interface{}, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://localhost:9100/api/profile/" + userId + "/email/" + emailId

	return clients.CallRestEndPoint(url, "DELETE", headers, nil)

}
