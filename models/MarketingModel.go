package models

import (
	"time"
)

type Marketing struct {
	MarketingId   string
	UserId        string
	BusinessId    string
	Transactional bool
	Promotional   bool
	Created       string
	Updated       string
}

type Banner struct {
	BannerId          string            `json:"bannerId"`
	MarketingCategory MarketingCategory `json:"marketingCategory"`
	Author            User              `json:"author"`
	SortOrder         int               `json:"sortOrder"`
	IsDisplayed       bool              `json:"isDisplayed"`
	Heading           string            `json:"heading"`
	Body              string            `json:"body"`
	Theme             string            `json:"theme"`
	Position          string            `json:"position"`
	TagLine           string            `json:"tagLine"`
	CtaLabel          string            `json:"ctaLabel"`
	CtaLink           string            `json:"ctaLink"`
	Image             Image             `json:"image"`
	Created           time.Time         `json:"created"`
	Updated           time.Time         `json:"updated"`
}

type MarketingCategory struct {
	MarketingCategoryId   string `json:"marketingCategoryId"`
	MarketingCategoryName string `json:"marketingCategoryName"`
}
