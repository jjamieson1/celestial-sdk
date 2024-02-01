package models

type celestialConfiguration struct {
	OauthClients           []OauthClient           `json:"oauthClients"`
	DatabaseConfigurations []DatabaseConfiguration `json:"databaseConfiguration"`
	AuditService           ServiceConfig           `json:"auditService"`
}

type OauthClient struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthUrl      string `json:"authUrl"`
	TokenUrl     string `json:"tokenUrl"`
	RedirectUrl  string `json:"redirectUrl"`
}

type DatabaseConfiguration struct {
	DataBaseDsn         string `json:"dataBaseDsn"`
	DataBaseSchema      string `json:"dataBaseSchema"`
	DataBaseUser        string `json:"dataBaseUser"`
	DataBasePass        string `json:"dataBasePass"`
	DataBaseHost        string `json:"dataBaseHost"`
	DataBaseFlags       string `json:"dataBaseFlags"`
	IsTls               bool   `json:"isTls"`
	DataBaseCertificate string `json:"dataBaseCertificate"`
}

type ServiceConfig struct {
	URLBase string `json:"urlBase"`
}
