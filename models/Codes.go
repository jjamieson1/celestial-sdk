package models

import "time"

type Code struct {
	CodeId     string    `json:"code_id"`
	Email      string    `json:"email"`
	Code       string    `json:"code"`
	Value      string    `json:"value"` // This code could have a dollar amount set in cents
	Type       string    `json:"type"`  // registration, discount, reset
	Expires    time.Time `json:"expires"`
	BusinessId string    `json:"business_id"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
