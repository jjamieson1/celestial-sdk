package user

import (
	clients2 "github.com/jjamieson1/eden-sdk/clients"

)

func UpdateAddress(tenantId string, userId string, addressId string, body []byte) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/address/" + addressId

	return clients2.CallRestEndPoint(url, "PUT", tenantId, body)
}

func UpdatePhone(tenantId string, userId string, phoneId string, body []byte) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/phone/" + phoneId

	return clients2.CallRestEndPoint(url, "PUT", tenantId, body)
}

func UpdateEmail(tenantId string, userId string, emailId string, body []byte) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId + "/email/" + emailId

	return clients2.CallRestEndPoint(url, "PUT", tenantId, body)
}

func UpdateUser(tenantId, userId string, body []byte) (interface{}, int, error) {

	url := "http://localhost:9100/api/profile/" + userId

	return clients2.CallRestEndPoint(url, "PUT", tenantId, body)
}
