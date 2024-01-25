package models

type OptionalPageElements struct {
	PageId 		string              `json:"pageId"`
	PageName 	string                `json:"pageName"`
	File 		string                `json:"file"`
	Slug 		string                `json:"slug"`
	Forms 		[]Form               `json:"forms"`
	Rules 		[]Rule               `json:"rules"`
	ContentObjects 	[]ContentObject `json:"contentObjects"`
}

type PageElements struct {
	User                 User                 `json:"user"`
	TenantId             string               `json:"tenantId"`
	Lang                 string               `json:"lang"`
	CompanyDetails       Tenant               `json:"companyDetails"`
	CategoryString       string               `json:"categoryString"`
	ProductCategories    []Category           `json:"productCategories"`
	ContentCategories    []CmsCategory        `json:"contentCategories"`
	OptionalPageElements OptionalPageElements `json:"optionalPageElements"`
}

type Form struct {
}

type ContentObject struct {
	ObjectName 		string						`json:"objectName"`
	ObjectArg 		string 						`json:"objectArg"`
	ObjectData		map[string]interface{} 		`json:"objectData"`
}