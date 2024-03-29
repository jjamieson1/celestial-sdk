package models

type Tenant struct {
	TenantId         string       `json:"tenantId"`
	ParentTenantId   string       `json:"parentTentantId"`
	Url              string       `json:"url"`
	CommonName       string       `json:"commonName"`
	Contacts         []User       `json:"contacts"`
	LogoUrl          string       `json:"logoUrl"`
	Addresses        []Address    `json:"addresses"`
	Phones           []Phone      `json:"phones"`
	Emails           []Email      `json:"email"`
	IsAvailable      bool         `json:"isAvailable"`
	TenantTypes      []TenantType `json:"tenantType"`
	ServiceProviders []ServiceProviders
	SecretKeys       SecretKeys `json:"keys,omitempty"`
}

type ServiceProviders struct {
	Id              string          `json:"id"`
	ServiceProvider ServiceProvider `json:"providerType"`
	BaseURL         string          `json:"baseURL"`
}

type ServiceProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SecretKeys struct {
	AppKey string `json:"app-key,omitempty"`
	ApiKey string `json:"api-key,omitempty"`
}

type TenantProvider struct {
	Id           string      `json:"id"`
	Adapter      Adapter     `json:"adapter"`
	TenantId     string      `json:"tenantId"`
	CalloutUrl   string      `json:"callOutUrl"`
	UserName     interface{} `json:"username,omitempty"`
	Password     interface{} `json:"password,omitempty"`
	ApiKey       interface{} `json:"apiKey,omitempty"`
	AppKey       interface{} `json:"appKey,omitempty"`
	Token        interface{} `json:"token,omitempty"`
	RefreshToken interface{} `json:"refreshToken,omitempty"`
}

type AuthStrategy struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Parameters string `json:"parameters"`
}

type Adapter struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	PluginName      string          `json:"pluginName"`
	ServiceProvider ServiceProvider `json:"serviceProvider"`
	AuthStrategy    AuthStrategy    `json:"authStrategy"`
	AdapterUrl      string          `json:"adapterUrl"`
	Enabled         bool            `json:"enabled"`
}

type TenantType struct {
	Id   string `json:"tenantTypeId"`
	Name string `json:"name"`
}
