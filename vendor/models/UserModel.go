package models

type User struct {
	UserId             string    `json:"userId"`
	SessionId          string    `json:"sessionId,omitempty"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Emails             []Email   `json:"emails,omitempty"`
	Phones             []Phone   `json:"phones,omitempty"`
	Addresses          []Address `json:"addresses,omitempty"`
	Password           string    `json:"password,omitempty"`
	HashedPassword     []byte    `json:"-"`
	Roles              []Role    `json:"roles,omitempty"`
	Active             bool      `json:"active"`
	RequiresPassChange bool      `json:"requiresPassChange"`
}

type Role struct {
	RoleId          string `json:"roleId"`
	RoleName        string `json:"roleName"`
	RoleDescription string `json:"roleDescription"`
	TenantId        string `json:"tenantId"`
}