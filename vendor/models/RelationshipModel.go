package models

type RelationShip struct {
	RelationShipId    string           `json:"relationShipId"`
	PrimaryEntityId   string           `json:"primaryEntityId"`
	SecondaryEntityId string           `json:"secondaryEntityId"`
	SecondaryEntity   interface{}      `json:"secondaryEntity"`
	RelationShipType  RelationShipType `json:"relationShipType"`
}

type RelationShipDomain struct {
	RelationShipDomainId string `json:"RelationShipDomainId"`
	DomainLabel          string `json:"domainLabel"`
	DomainDescription    string `json:"domainDescription"`
}

type RelationShipType struct {
	RelationShipTypeId    string             `json:"relationShipTypeId"`
	RelationshipTypeName  string             `json:"relationshipTypeName"`
	RelationshipTypeLabel string             `json:"relationshipTypeLabel"`
	Description           string             `json:"description"`
	RelationshipDomain    RelationShipDomain `json:"relationshipTypeDomain"`
}
