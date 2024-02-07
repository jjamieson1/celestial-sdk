package models

type User struct {
	UserId             string    `json:"userId"`
	SessionId          string    `json:"sessionId,omitempty"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Emails             []Email   `json:"emails,omitempty"`
	Phones             []Phone   `json:"phones,omitempty"`
	Addresses          []Address `json:"addresses,omitempty"`
	Password           []byte    `json:"-"`
	HashedPassword     []byte    `json:"-"`
	NewPassword        string    `json:"newPassword,omitempty"`
	Roles              []Role    `json:"roles,omitempty"`
	Active             bool      `json:"active"`
	TenantId           string    `json:"tenantId"`
	RequiresPassChange bool      `json:"requiresPassChange"`
}

type Role struct {
	RoleId          string `json:"roleId"`
	RoleName        string `json:"roleName"`
	RoleDescription string `json:"roleDescription"`
	TenantId        string `json:"tenantId"`
}
