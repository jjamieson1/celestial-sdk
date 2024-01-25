package models

type CatalogItem struct {
	ItemId            string             `json:"itemId"`
	ItemName          string             `json:"itemName"`
	Sku               string             `json:"sku"`
	IsArchived        bool               `json:"isArchived"`
	Description       string             `json:"description"`
	Category          Category           `json:"category"`
	Manufacturer      string             `json:"manufacturer"`
	IsSaleable        bool               `json:"isSaleable"`
	IsFeatured        bool               `json:"isFeatured"`
	Inventory         int                `json:"inventory"`
	RegularPrice      int64              `json:"regularPrice"`
	SalePrice         int64              `json:"SalePrice"`
	Discount          int                `json:"discount"`
	CatalogItemAssets []CatalogItemAsset `json:"catalogItemAssets"`
	CatalogItemFields []CatalogItemField `json:"catalogItemFields"`
	LocationId        int                `json:"locationId"`
	TenantId          string             `json:"tenantId"`
	CatalogProvider   CatalogProvider    `json:"catalogProvider"`
}

type CatalogItemAsset struct {
	ItemAssetId     string `json:"itemAssetId"`
	ItemId          string `json:"itemId"`
	AssetName       string `json:"assetName"`
	Url             string `json:"url"`
	AssetType       string `json:"assetType"`
	IsVendorManaged bool   `json:"isVendorManaged"`
}

type CatalogItemField struct {
	FieldId         string `json:"fieldId"`
	ItemId          string `json:"itemId"`
	FieldName       string `json:"fieldName"`
	FieldValue      string `json:"fieldValue"`
	FieldType       string `json:"fieldType"`
	FieldUnit       string `json:"fieldUnit"`
	IsVendorManaged bool   `json:"isVendorManaged"`
	IsVisible       bool   `json:"isVisible"`
	IsIgnored       bool   `json:"isIgnored"`
}

type CatalogProvider struct {
	CatalogProviderId   string `json:"catalogProviderId"`
	CatalogProviderName string `json:"catalogProviderName"`
}

type Category struct {
	CategoryId          string `json:"categoryId"`
	CategoryName        string `json:"categoryName"`
	CategoryDescription string `json:"categoryDescription"`
	CategoryImage       string `json:"categoryImage"`
}
