package models

import "time"

type PreRegistration struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	TenantType         string    `json:"tenantType"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Email              string    `json:"email"`
	Phone              string    `json:"phone"`
	URL                string    `json:"url"`
	Mission            string    `json:"mission"`
	Street             string    `json:"street"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	Postal             string    `json:"postal"`
	TenantId           string    `json:"tenantId"`
	ParentOrganization string    `json:"parentOrganizationId"`
	Date               time.Time `json:"date"`
}
