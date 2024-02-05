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

type SessionModel struct {
	UserId      string `json:"userId"`
	Jwt         string `json:"jwt"`
	TenantId    string `json:"tenantId"`
	BusinessUrl string `json:"businessUrl"`
}
