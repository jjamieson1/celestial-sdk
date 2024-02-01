package models

type CelestialConfiguration struct {
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

type ConfigModel struct {
	ConfigurationId                  string `json:"configuration_id"`
	RegistrationCode                 bool   `json:"registration_code"`  // Requires a invitation code to register
	CloudFlareApi                    string `json:"cloud_flare_api"`    // Required to use cloud flare Tunrstile
	CloudFlareSecret                 string `json:"cloud_flare_secret"` // Required to use cloud flare Tunrstile
	UseCaptcha                       bool   `json:"use_captcha"`        // Whether to use captcha on registration and login
	MailJetApiKey                    string `json:"mail_jet_api_key"`
	MailJetSecretKey                 string `json:"mail_jet_secret_key"`
	MailJetInvitationTemplateId      int    `json:"mail_jet_invitation_template_id"`
	MailJetRegistrationTemplateId    int    `json:"mail_jet_registration_template_id"`
	MailJetResetPasswordTemplateId   int    `json:"mail_jet_reset_password_template_id"`
	MailJetActivateAccountTemplateId int    `json:"mail_jet_activate_account_template_id"`
	MailJetUnsubscribeTemplateId     int    `json:"mail_jet_unsubscribe_template_id"`
	MailJetReplyToEmailAddress       string `json:"mail_jet_reply_to_email_address"`
	MailJetReplyToName               string `json:"mail_jet_reply_to_name"`
	BaseUrl                          string `json:"base_url"`
	BusinessId                       string `json:"business_id"`
}
