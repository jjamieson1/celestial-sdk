package models

type ValidationError struct {
	Message string `json:"Message"`
	Key     string `json:"Key"`
}
