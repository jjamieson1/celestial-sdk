package models

type Tenant struct {
	TenantId         string `json:"tenantId"`
	ParentTenantId   string `json:"parentTentantId"`
	Url              string `json:"url"`
	CommonName       string `json:"commonName"`
	Contacts         []User `json:"contacts"`
	LogoUrl          string `json:"logoUrl"`
	LogoSecondaryUrl string `json:"logoSecondaryUrl"`
	Mission          string `json:"mission"`
	Street           string `json:"street"`
	City             string `json:"city"`
	State            string `json:"state"`
	Postal           string `json:"postal"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Hours            Hours  `json:"hours"`
	Promotional      string `json:"promotional"`
	Theme            string `json:"theme"`
	IsAvailable      bool   `json:"isAvailable"`
	TenantType       string `json:"tenantType"`
}

type Hours struct {
	Monday			string 		`json:"monday"`
	Tuesday			string 		`json:"tuesday"`
	Wednesday		string 		`json:"wednesday"`
	Thursday		string 		`json:"thursday"`
	Friday			string 		`json:"friday"`
	Saturday		string 		`json:"saturday"`
	Sunday			string 		`json:"sunday"`
}


type TenantProvider struct {
	Id           string      `json:"id"`
	EdenAdapter  EdenAdapter `json:"edenAdapter"`
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
	Id 				string 		`json:"id"`
	Name 			string 		`json:"name"`
	Parameters		string 		`json:"parameters"`
}

type ProviderType struct {
	Id 				string 		`json:"id"`
	Name 			string 		`json:"name"`
}

type EdenAdapter struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	PluginName   string       `json:"pluginName"`
	ProviderType ProviderType `json:"providerType"`
	AuthStrategy AuthStrategy `json:"authStrategy"`
	AdapterUrl   string       `json:"adapterUrl"`
	Enabled      bool         `json:"enabled"`
}

type TenantType struct {
	Id 				string 			`json:"tenantTypeId"`
	Name 			string 			`json:"name"`
	Description 	string 			`json:"description"`
	TenantId		string 			`json:"tenantId"`
}