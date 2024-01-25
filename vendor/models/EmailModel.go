package models

type Email struct {
	EmailId   string    `json:"emailId"`
	UserId    string    `json:"userId"`
	Email    string `json:"email"`
	Verified  bool   `json:"verified"`
}
