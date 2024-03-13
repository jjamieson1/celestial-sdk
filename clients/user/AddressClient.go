package user

import (
	"encoding/json"
	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func AddUpdateAddress(address models.Address, tokenString, tenantId string) (models.Address, int, error) {
	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/address"
	method := "POST"

	body, err := json.Marshal(address)
	response, status, err := clients.CallRestEndPoint(url, method, headers, body)
	var a models.Address
	err = json.Unmarshal(response, &a)
	return a, status, err
}

func FindAddressByAddressId(addressId, tokenString, tenantId string) (models.Address, int, error) {
	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/address/" + addressId
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	var a models.Address
	err = json.Unmarshal(response, &a)
	return a, status, err
}

func DeleteAddress(addressId, tokenString, tenantId string) (status int, err error) {
	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/address/" + addressId
	method := "DELETE"

	_, status, err = clients.CallRestEndPoint(url, method, headers, nil)

	return status, err
}
