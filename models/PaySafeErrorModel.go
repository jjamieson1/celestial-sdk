package models

type PaySafeError struct {
	Links          []Links `json:"links"`
	Id             string  `json:"id"`
	MerchantRefNum string  `json:"merchantRefNum"`
	Error          Error   `json:"error"`
	SettleWithAuth bool    `json:"settleWithAuth"`
}


type Links struct {
	Rel 	string 	`json:"rel"`
	Href	string	`json:"href"`

}

type Error struct {
	Code 			string   `json:"code"`
	Message 		string    `json:"message"`
	Links			[]Links `json:"links"`
}
