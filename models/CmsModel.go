package models

import (
	"time"
)

type Cms struct {
	CmsId       string      `json:"cmsId"`
	CmsMetaData CmsMetaData `json:"cmsMetaData"`
	CmsContent  CmsContent  `json:"cmsContent"`
	Created     time.Time   `json:"created"`
	Modified    time.Time   `json:"modified"`
}

type CmsMetaData struct {
	CmsCategoryId string    `json:"cmsCategoryId"`
	CmsType       string    `json:"cmsType"`
	CmsCategory   string    `json:"cmsCategory"`
	Modified      time.Time `json:"modified"`
	Slug          string    `json:"slug"`
	Sort          int       `json:"sort"`
	Status        string    `json:"status"`
	Who           string    `json:"who"`
	IsFeatured    bool      `json:"isFeatured"`
	IsPublic      bool      `json:"isPublic"`
	Created       time.Time `json:"created"`
	BusinessId    string    `json:"businessId"`
}

type CmsContent struct {
	Title     string  `json:"title"`
	ShortText string  `json:"shortText"`
	Body      string  `json:"body"`
	Lang      string  `json:"lang"`
	Images    []Image `json:"images"`
}

type CmsCategory struct {
	CmsCategoryId string  `json:"cmsCategoryId"`
	Name          string  `json:"name"`
	ParentId      string  `json:"parentId"`
	Images        []Image `json:"image,omitempty"`
}
