package models

type PaySafeRequestModel struct {
	MerchantReferenceNum string         `json:"merchantRefNum"`
	Amount               int64          `json:"amount"`
	SettleWithAuth       bool           `json:"settleWithAuth"`
	DupCheck             bool           `json:"dupCheck"`
	Card                 Card           `json:"card"`
	Profile              Profile        `json:"profile"`
	BillingDetails       BillingDetails `json:"billingDetails"`
	KeyWords             []string       `json:"keywords"`
	CustomerIp           string         `json:"customerIp"`
	Description          string         `json:"description"`
}

type Card struct {
	PaymentToken		string	`json:"paymentToken"`
}

type Profile struct {
	FirstName			string 	`json:"firstName"`
	LastName			string	`json:"lastName"`
	Email 				string	`json:"email"`
}

type BillingDetails struct {
	Street 				string	`json:"street"`
	City 				string	`json:"city"`
	State 				string	`json:"state"`
	Country 			string	`json:"country"`
	Zip 				string	`json:"zip"`
}


