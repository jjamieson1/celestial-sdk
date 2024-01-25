package models

type Phone struct {
	PhoneId   string    `json:"phoneId"`
	UserId    string    `json:"userId"`
	Number    string `json:"number"`
	Verified  bool   `json:"verified"`
	PhoneType string `json:"numberType"`
}
