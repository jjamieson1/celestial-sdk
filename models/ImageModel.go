package models

import "time"

type Image struct {
	ImageId          string    `json:"imageId"`
	Type             string    `json:"type,omitempty"`
	Description      string    `json:"description,omitempty"`
	Filename         string    `json:"filename,omitempty"`
	Ordering         string    `json:"ordering,omitempty"`
	PublicID         string    `json:"publicID,omitempty"`
	BaseImageURL     string    `json:"baseImageURL,omitempty"`
	Size             string    `json:"size,omitempty"`
	CreateTime       time.Time `json:"createTime,omitempty"`
	TimeStamp        time.Time `json:"timeStamp,omitempty"`
	ItemID           string    `json:"itemID,omitempty"`
	OriginalFilename string    `json:"originalFilename,omitempty"`
	BusinessId       string    `json:"businessId,omitempty"`
}
