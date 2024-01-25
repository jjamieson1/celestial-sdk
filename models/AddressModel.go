package models

type Address struct {
	AddressId    string      `json:"addressId"`
	UserId       string      `json:"userId"`
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 string      `json:"addressLine2"`
	City         string      `json:"city"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	PostalCode   string      `json:"postalCode"`
	AddressType  AddressType `json:"addressType"`
	Verified     bool        `json:"verified"`
}

type AddressType struct {
	AddressTypeId   string `json:"addressTypeId"`
	AddressTypeName string `json:"addressTypeName"`
}
