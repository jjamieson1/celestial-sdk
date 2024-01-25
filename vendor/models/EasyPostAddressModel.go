package models

import "time"

type EasyPostAddress struct {
	Id					string	`json:"id"`// Unique identifier, begins with "adr_"
	Object				string	`json:"object"` // "Address"
	Mode				string	`json:"mode"`	// Set based on which api-key you used, either "test" or "production"
	Street1				string	`json:"street1"`	// First line of the address
	Street2				string	`json:"street2"`	// Second line of the address
	City				string	`json:"city"`	// City the address is located in
	State				string	`json:"state"`	// State or province the address is located in
	Zip					string	`json:"zip"`	// ZIP or postal code the address is located in
	Country				string	`json:"country"`	// ISO 3166 country code for the country the address is located in
	Residential			bool	`json:"residential"`	// Whether or not this address would be considered residential
	CarrierFacility		string	`json:"carrier_facility"`	// The specific designation for the address (only relevant if the address is a carrier facility)
	Name				string	`json:"name"`	// Name of the person. Both name and company can be included
	Company				string	`json:"company"`	// Name of the organization. Both name and company can be included
	Phone				string	`json:"phone"`	// Phone number to reach the person or organization
	Email				string	`json:"email"`	// Email to reach the person or organization
	FederalTaxId		string	`json:"federal_tax_id"`	// Federal tax identifier of the person or organization
	StateTaxId			string	`json:"state_tax_id"`	// State tax identifier of the person or organization
}


type EasyPostParcel struct {
	Id 					string    	`json:"id"` // Unique, begins with "prcl_"
	Object    			string    	`json:"object"` // "Parcel"
	Mode    			string    	`json:"mode"` // "test" or "production"
	Length    			float32   	`json:"length"`// (inches)    Required if width and/or height are present
	Width    			float32   	`json:"width"`// (inches)    Required if length and/or height are present
	Height    			float32    	`json:"height"`//(inches)    Required if length and/or width are present
	PredefinedPackage  string    	`json:"predefined_package"`//Optional, one of our predefined_packages
	Weight    			float32     `json:"weight"`// (oz)    Always required
	CreatedAt    		time.Time  	`json:"created_at"`//
	UpdatedAt    		time.Time  	`json:"updated_at"`//
}

