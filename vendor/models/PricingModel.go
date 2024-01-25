package models

import "time"

type ProductPrice struct {
  Id                      int         `json:"Id"`
  CompanyId               int         `json:"CompanyId"`
  EntityId                int         `json:"EntityId"`
  CatalogItemId           string      `json:"CatalogItemId"`
  PricingTermId           int         `json:"PricingTermId"`
  RegularPrice            int64       `json:"RegularPrice"`
  IsDiscountable          bool        `json:"IsDiscountable"`
  FloorPrice              int64       `json:"FloorPrice"`
  OriginalPrice           int64       `json:"OriginalPrice"`
  PricingTierId           int         `json:"PricingTierId"`
  PricingGroupId          int         `json:"PricingGroupId"`
  PricingShelfId          int         `json:"PricingShelfId"`
  OverridePriceId         int         `json:"OverridePriceId"`
  OverridePrice           int64       `json:"OverridePrice"`
  OverrideStartDateUtc    time.Time   `json:"OverrideStartDateUtc"`
  OverrideStopDateUtc     time.Time   `json:"OverrideStopDateUtc"`
}

