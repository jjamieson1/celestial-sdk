package models

type AuthenticatedUser struct {
	IsMfa      bool   `json:"isMfa"`
	JWT        string `json:"jwt"`
	UserId     string `json:"userId"`
	ValidUntil int64  `json:"validUntil"`
}

type ChallengeCredential struct {
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Jwt      string `json:"jwt"`
}
