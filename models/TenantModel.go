package models

type Tenant struct {
	TenantId       string     `json:"tenantId"`
	ParentTenantId string     `json:"parentTentantId"`
	Url            string     `json:"url"`
	CommonName     string     `json:"commonName"`
	Contacts       []User     `json:"contacts"`
	LogoUrl        string     `json:"logoUrl"`
	Addresses      []Address  `json:"addresses"`
	Phones         []Phone    `json:"phones"`
	Emails         []Email    `json:"email"`
	IsAvailable    bool       `json:"isAvailable"`
	TenantType     TenantType `json:"tenantType"`
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

type ProviderType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Adapter struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	PluginName   string       `json:"pluginName"`
	ProviderType ProviderType `json:"providerType"`
	AuthStrategy AuthStrategy `json:"authStrategy"`
	AdapterUrl   string       `json:"adapterUrl"`
	Enabled      bool         `json:"enabled"`
}

type TenantType struct {
	Id          string `json:"tenantTypeId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
