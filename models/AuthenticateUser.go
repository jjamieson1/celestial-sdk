package models

type AuthenticatedUser struct {
	IsMfa      bool   `json:"isMfa"`
	JWT        string `json:"jwt"`
	UserId     string `json:"userId"`
	User       User   `json:"user"`
	ValidUntil int64  `json:"validUntil"`
}

type ChallengeCredential struct {
	UserId          string `json:"userId"`
	EmailOrUserName string `json:"emailOrUserName"`
	Password        string `json:"password"`
	Jwt             string `json:"jwt"`
}

type LoginCredential struct {
	EmailOrUserName string `json:"emailOrUserName"`
	Password        string `json:"password"`
}
