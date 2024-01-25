package models

import "time"

type Event struct {
	EventId      string           `json:"eventId"`
	Name         EventName        `json:"eventName"`
	Dates        EventDates       `json:"eventDates"`
	Level        Level            `json:"level"`
	Owners       string           `json:"owners"`
	URL          string           `json:"url"`
	Description  EventDescription `json:"description"`
	Location     EventLocation    `json:"location"`
	EventType    EventType        `json:"eventType"`
	VenueName    string           `json:"venueName"`
	Rules        []Rule           `json:"rules"`
	Note         string           `json:"notes"`
	Fee          string           `json:"fee"`
	PaymentTypes []string         `json:"paymentTypes"`
	Approved     bool             `json:"approved"`
	Policy       Policy           `json:"policy"`
	//Files 			[]File 				`json:"files"`
	Files 			string       `json:"files"`
	Links 			[]Link       `json:"links"`
	CustomFields 	[]CustomField `json:"customField"`
	TenantId		string        `json:"tenantId"`
}

type EventName struct {
	En 		string 		`json:"en"`
	Fr 		string 		`json:"fr"`
}

type EventDates struct {
	Publish 				time.Time 	`json:"publish"`
	Retract 				time.Time	`json:"retract"`
	RegistrationOpen		time.Time 	`json:"registrationOpen"`
	EarlyBird 				time.Time	`json:"earlyBird"`
	RegistrationClosed 		time.Time 	`json:"registrationClosed"`
	DateOfEvent 			time.Time	`json:"dateOfEvent"`
	LastDayOfEvent 			time.Time	`json:"lastDayOfEvent"`
}

type EventDescription struct {
	En 			string 		`json:"en"`
	Fr 			string 		`json:"fr"`
}

type EventLocation struct {
	Street 		string 			`json:"street"`
	City 		string 			`json:"city"`
	Province 	string 			`json:"province"`
}

type Rule struct {
	RuleId				string 		`json:"ruleId"`
	RuleName 				string 		`json:"ruleName"`
	TypeName 				string 		`json:"typeName"`
	TypeId 				string 		`json:"typeId"`
}
//
//type Note struct {
//	Id 				string 			`json:"id"`
//	Name 			string 			`json:"name"`
//	Value 			string			`json:value`
//	Type 			string 			`json:"type"`
//	TypeId 			string 			`json:"typeId"`
//}

type Policy struct  {
	Id 				string 			`json:"id"`
	Name  			string 			`json:"name"`
	Details 		string 			`json:"details"`
	Version  		string 			`json:"version"`
}

type File struct {
	Id 				string 			`json:"id"`
	Name 			string 			`json:"name"`
	Location 		string 			`json:"location"`
}

type Link 	struct {
	Id  		string 			`json:"id"`
	Description string 			`json:"description"`
	URL 		string 			`json:"url"`
}

type CustomField 	struct {
	Id 			string 			`json:"id"`
	Name 		string 			`json:"name"`
	Type 		string 			`json:"type"`
}

type Level struct {
	Id 			string 			`json:"id"`
	NameEn 		string 			`json:"nameEn"`
	NameFr 		string 			`json:"nameFr"`
}
type EventType struct {
	EventTypeId 	string 		`json:"eventTypeId"`
	EventTypeName 	string 		`json:"eventTypeName"`
}