package models

type ReleaseInformation struct {
	VersionNumber string `json:"versionNumber"`
	AppName       string `json:"appName"`
	TenantName    string `json:"tenantName"`
	StartedOn     string `json:"startedOn"`
}
