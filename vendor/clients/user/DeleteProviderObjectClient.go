package user

import (
	clients2 "github.com/jjamieson1/eden-sdk/clients"

)

func DeleteAddress(tenantId, userId, addressId string) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/address/" + addressId

	return clients2.CallRestEndPoint(url, "DELETE", tenantId, nil)
}

func DeletePhone(tenantId, userId, phoneId string) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients2.CallRestEndPoint(url, "DELETE", tenantId, nil)

}

func DeleteEmail(tenantId, userId, emailId string) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/email/" + emailId

	return clients2.CallRestEndPoint(url, "DELETE", tenantId, nil)

}
