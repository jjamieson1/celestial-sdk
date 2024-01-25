package models

type Orders struct {

	OrderId 			 int 	    `json:"orderId"`
	UserId 				 string 	`json:"userId"`
	PaymentToken		 string 	`json:"paymentToken"`
	Amount 				 int64 		`json:"amount"`
	FirstName			string 		`json:"firstName"`
	LastName			string		`json:"lastName"`
	Email 				string		`json:"email"`
	Street 				string		`json:"street"`
	City 				string		`json:"city"`
	State 				string		`json:"state"`
	Country 			string		`json:"country"`
	Zip 				string		`json:"zip"`
	BillingStreet 				string		`json:"billingStreet"`
	BillingCity 				string		`json:"billingCity"`
	BillingState 				string		`json:"billingState"`
	BillingCountry 			string		`json:"billingCountry"`
	BillingZip 				string		`json:"billingZip"`
	Description			 string		`json:"description"`
	CustomerIp			string		`json:"customerIp"`
	Phone 				string		`json:"phone"`
	Agreed				bool		`json:"agreed"`
	TrackingNumber		string		`json:"trackingNumber"`
}