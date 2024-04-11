package models

import "time"

type Item struct {
	ItemId            string             `json:"itemID"`
	VendorItemID      string             `json:"vendorItemID"`
	SystemSku         string             `json:"systemSku"`
	DefaultCost       string             `json:"defaultCost"`
	AvgCost           string             `json:"avgCost"`
	Discountable      string             `json:"discountable"`
	Tax               string             `json:"tax"`
	Archived          string             `json:"archived"`
	ItemType          string             `json:"itemType"`
	Serialized        string             `json:"serialized"`
	Description       string             `json:"description"`
	ModelYear         string             `json:"modelYear"`
	Upc               string             `json:"upc"`
	Ean               string             `json:"ean,omitempty"`
	CustomSku         string             `json:"customSku,omitempty"`
	ManufacturerSku   string             `json:"manufacturerSku,omitempty"`
	CreateTime        time.Time          `json:"createTime"`
	TimeStamp         time.Time          `json:"timeStamp"`
	CategoryID        string             `json:"categoryID"`
	TaxClassID        string             `json:"taxClassID"`
	DepartmentID      string             `json:"departmentID"`
	ItemMatrixID      string             `json:"itemMatrixID"`
	ManufacturerID    string             `json:"manufacturerID"`
	SeasonID          string             `json:"seasonID"`
	DefaultVendorID   string             `json:"defaultVendorID"`
	DefaultImage      string             `json:"defaultImage,omitempty"`
	Images            []Image            `json:"images"`
	Category          Category           `json:"category"`
	TaxClass          TaxClass           `json:"taxClass"`
	ItemAttributes    []ItemAttribute    `json:"itemAttributes"`
	Manufacturer      Manufacturer       `json:"manufacturer"`
	Note              Note               `json:"note"`
	ItemShops         []ItemShop         `json:"itemShops"`
	ItemVendorNums    ItemVendorNums     `json:"itemVendorNums"`
	MSRP              string             `json:"msrp"`
	Prices            Prices             `json:"prices"`
	CustomFieldValues []CustomFieldValue `json:"customFieldValues"`
	Tags              Tags               `json:"tags"`
	Score             float64            `json:"score"`
	HasInventory      bool               `json:"hasInventory,omitempty"`
	AddedBy           string             `json:"addedBy"`
	BusinessId        string             `json:"businessId"`
}

type Tags struct {
	TagAttributes TagAttributes `json:"attributes"`
	Tag           []string      `json:"tag"`
}

type TagAttributes struct {
	Count string `json:"count"`
}

type ItemVendorNums struct {
	ItemVendorNum []ItemVendorNum `json:"ItemVendorNum"`
}

type ItemVendorNum struct {
	ItemVendorNumID string    `json:"ItemVendorNumID"`
	Value           string    `json:"value"`
	TimeStamp       time.Time `json:"timeStamp"`
	Cost            string    `json:"cost"`
	ItemID          string    `json:"itemID"`
	VendorID        string    `json:"vendorID"`
}

type Note struct {
	Note      string    `json:"note,omitempty"`
	IsPublic  string    `json:"isPublic,omitempty"`
	TimeStamp time.Time `json:"timeStamp,omitempty"`
}

type Manufacturer struct {
	ManufacturerID string    `json:"manufacturerID,omitempty"`
	Name           string    `json:"name,omitempty"`
	CreateTime     time.Time `json:"createTime,omitempty"`
	TimeStamp      time.Time `json:"timeStamp,omitempty"`
}

type ItemAttributeSet struct {
	ItemAttributeSetID     string                  `json:"itemAttributeSetID"`
	AttributeSetName       string                  `json:"attributeSetName"`
	AttributeNameValueSets []AttributeNameValueSet `json:"attributeNameValueSets"`
}

type AttributeNameValueSet struct {
	AttributeNameValueSetId string           `json:"attributeNameValueSetId"`
	AttributeName           AttributeName    `json:"attributeName"`
	AttributeValues         []AttributeValue `json:"attributeValues"`
}

type AttributeName struct {
	AttributeNameId string `json:"attributeNameId"`
	AttributeName   string `json:"attributeName"`
}

type AttributeValue struct {
	AttributeValueId string `json:"attributeValueId"`
	AttributeValue   string `json:"attributeValue"`
	AttributeImage   Image  `json:"attributeImage"`
}

type ItemAttribute struct {
	ItemAttributeName  AttributeName  `json:"itemAttributeName"`
	ItemAttributeValue AttributeValue `json:"itemAttributeValue"`
}

type ItemShop struct {
	ItemShopID    string    `json:"itemShopID,omitempty"`
	Qoh           string    `json:"qoh"`
	Sellable      string    `json:"sellable"`
	Backorder     string    `json:"backorder"`
	BackOrderEta  time.Time `json:"backOrderEta"`
	TimeStamp     time.Time `json:"timeStamp"`
	Layaways      string    `json:"layaways"`
	SpecialOrders string    `json:"specialorders"`
	WorkOrders    string    `json:"workorders"`
	ItemID        string    `json:"itemID"`
	ShopID        string    `json:"shopID"`
}

type BackOrder struct {
	ItemId           string    `json:"itemId"`
	VendorItemId     string    `json:"vendorItemId"`
	Qty              string    `json:"qty"`
	Eta              time.Time `json:"eta"`
	Description      string    `json:"description"`
	ManufacturerName string    `json:"manufacturerName"`
	ManufacturerId   string    `json:"manufacturerId"`
	Upc              string    `json:"upc"`
	Ean              string    `json:"ean"`
	ItemGroup        ItemGroup `json:"itemGroup"`
	Image            Image     `json:"image"`
}

type ItemPrice struct {
	ItemId    string `json:"itemId"`
	Amount    string `json:"amount"`
	UseTypeID string `json:"useTypeID"`
	UseType   string `json:"useType"`
}

type Prices struct {
	ItemPrice []ItemPrice `json:"itemPrice"`
}

type TaxClass struct {
	TaxClassID string    `json:"taxClassID"`
	Name       string    `json:"name"`
	TimeStamp  time.Time `json:"timeStamp"`
}

type ItemGroup struct {
	ItemGroupId      string           `json:"itemGroupId"`
	Category         Category         `json:"category"`
	Manufacturer     Manufacturer     `json:"manufacturer"`
	IsDisplayed      bool             `json:"isDisplayed"`
	Images           []Image          `json:"images"`
	ImageBaseUrl     string           `json:"imageBaseUrl"`
	Name             string           `json:"name"`
	Details          string           `json:"details"`
	Specifications   string           `json:"specifications,omitempty"`
	PartNumbers      []string         `json:"partNumbers,omitempty"`
	Geometry         string           `json:"geometry,omitempty"`
	Sizing           string           `json:"sizing,omitempty"`
	Created          time.Time        `json:"created"`
	TimeStamp        time.Time        `json:"timeStamp"`
	PriceRange       PriceRange       `json:"priceRange"`
	SaleCampaign     SaleCampaign     `json:"saleCampaign,omitempty"`
	Items            []Item           `json:"items"`
	ItemAttributeSet ItemAttributeSet `json:"itemAttributeSet"`
}

type ItemLog struct {
	ItemId   string    `json:"itemId"`
	Action   string    `json:"action"`
	Activity string    `json:"activity"`
	Created  time.Time `json:"created"`
	TenantId string    `json:"tenantId"`
}

type PriceRange struct {
	HasPrice bool   `json:"hasPrice"`
	HasRange bool   `json:"hasRange"`
	High     string `json:"high"`
	Low      string `json:"low"`
}

type CustomFieldValue struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	TimeStamp time.Time `json:"timeStamp"`
}

type AttributeFilterRequest struct {
	AttrNameId  string `json:"attrNameId"`
	AttrValueId string `json:"attrValueId"`
}

type OptionNameList struct {
	AttrNameId       string            `json:"attrNameId"`
	AttrName         string            `json:"attrName"`
	OptionValueLists []OptionValueList `json:"optionValueList"`
}

type OptionValueList struct {
	ValueId     string `json:"valueId"`
	Value       string `json:"value"`
	IsAvailable bool   `json:"isAvailable"`
	IsSelected  bool   `json:"isSelected"`
}

type Review struct {
	ReviewId    string `json:"reviewId"`
	ItemGroupId string `json:"itemGroupId"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Url         string `json:"url"`
	Media       string `json:"media"`
	Score       string `json:"score"`
}

type SaleCampaign struct {
	Id                    string      `json:"id,omitempty"`
	OnSale                bool        `json:"onSale"`
	SaleCampaignName      string      `json:"saleCampaignName,omitempty"`
	SaleCampaignDetails   string      `json:"saleCampaignDetails,omitempty"`
	SaleCampaignDiscount  int         `json:"saleCampaignDiscount,omitempty"`
	SaleCampaignInventory []ItemGroup `json:"saleCampaignInventory,omitempty"`
	StartDate             time.Time   `json:"saleCampaignStartDate,omitempty"`
	EndDate               time.Time   `json:"saleCampaignEndDate,omitempty"`
	CreatedBy             string      `json:"created_by,omitempty"`
	UpdatedBy             string      `json:"updated_by,omitempty"`
	Created               time.Time   `json:"created,omitempty"`
	Updated               time.Time   `json:"updated,omitempty"`
	BusinessId            string      `json:"businessId,omitempty"`
}

type Category struct {
	CategoryId    string     `json:"categoryID"`
	Name          string     `json:"name"`
	NodeDepth     int        `json:"nodeDepth"`
	FullPathName  string     `json:"fullPathName,omitempty"`
	LeftNode      string     `json:"leftNode,omitempty"`
	RightNode     string     `json:"rightNode,omitempty"`
	ParentId      string     `json:"parentId"`
	Children      []Category `json:"children"`
	NumItemGroups int        `json:"numItemGroups"`
	Next          string     `json:"next"`
	Images        []Image    `json:"images"`
	IsDisplayed   bool       `json:"isDisplayed,omitempty"`
	LastSync      time.Time  `json:"lastSync,omitempty"`
	AddedBy       string     `json:"addedBy"`
	CreateTime    time.Time  `json:"createTime,omitempty"`
	TimeStamp     time.Time  `json:"timeStamp,omitempty"`
}

type CategoryList struct {
	MetaData   MetaData   `json:"metaData"`
	Categories []Category `json:"categories"`
}

type MetaData struct {
	NumResults int    `json:"numResults"`
	Next       string `json:"next"`
	Previous   string `json:"previous"`
	Page       int    `json:"page"`
	NumPages   int    `json:"numPages"`
	PageItems  int    `json:"pageItems"`
	Cache      Cache  `json:"cache"`
}

type Cache struct {
	IsCached       bool      `json:"isCached"`
	CacheSetExpiry time.Time `json:"cacheSetExpiry"`
	CacheSet       time.Time `json:"cacheSet"`
}
