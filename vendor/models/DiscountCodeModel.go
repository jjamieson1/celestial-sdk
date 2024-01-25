package models

import "time"

type DiscountCode struct {
	DisCountCodeId	int			`json:"discountCodeId"`
	Code 			string		`json:"code"`
	Amount 			int64		`json:"amount"`
	CartId			string		`json:"cartId"`
	Active 			bool		`json:"active"`
	Created			time.Time	`json:"created"`
	IsEmailEd		bool		`json:"isEmailed"`
}
