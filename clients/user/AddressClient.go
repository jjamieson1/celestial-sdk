package user

func AddUpdateAddress(address models.Address, tokenString, tenantId string) (models.AuthenticatedUser, int, error) {
	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/address"
	method := "POST"

	response, status, err := clients.CallRestEndPoint(url, method, headers, address)
	var a models.Address
	err := json.Unmarshal(response, &a)
	return address, status, err
}

func DeleteAddress(addressId, tokenString, tenantId string) (status int, err error) {
	// Add the headers
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "bearer " + tokenString,
	}
	url := "http://localhost:3000/api/v1/auth/account/address/" + addressId
	method := "DELETE"

	_, status, err := clients.CallRestEndPoint(url, method, headers, address)

	return status, err
}
