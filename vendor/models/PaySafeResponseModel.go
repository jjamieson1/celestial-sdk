package models

type PaySafeResponse struct {
	Id                    string             `json:"id"`
	MerchantRefNum        string             `json:"merchantRefNum"`
	UserId                string             `json:"userId"`
	TxnTime               string             `json:"txnTime"`
	Status                string             `json:"status"`
	Amount                int64              `json:"amount"`
	SettleWithAuth        bool               `json:"settleWithAuth"`
	SettleWithOfflineAuth bool               `json:"settleWithOfflineAuth"`
	PreAuth               bool               `json:"preAuth"`
	AvailableToSettle     int64              `json:"availableToSettle"`
	Card                  ResponseCard       `json:"card"`
	AuthCode              string             `json:"authCode"`
	Profile               Profile            `json:"profile"`
	BillingDetails        BillingDetails     `json:"billingDetails"`
	CustomerIp            string             `json:"customerIp"`
	Keywords              []string           `json:"keywords"`
	MerchantDescriptor    MerchantDescriptor `json:"merchantDescriptor"`
	Description           string             `json:"description"`
	CurrencyCode          string             `json:"currencyCode"`
	AvsResponse           string             `json:"avsResponse"`
	CvvVerification       string             `json:"cvvVerification"`
	Filled                bool               `json:"filled"`
}

type ResponseCard struct {
	Type       string     `json:"type"`
	LastDigits string     `json:"lastDigits"`
	CardExpiry CardExpiry `json:"cardExpiry"`
}

type CardExpiry struct {
	Month 					int			`json:"month"`
	Year 					int 		`json:"year"`
}

type MerchantDescriptor struct {
	DynamicDescriptor		string 		`json:"dynamicDescriptor"`
	Phone 					string 		`json:"phone"`
}