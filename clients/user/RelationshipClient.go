package user

import (
	"encoding/json"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/clients/tenant"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetRelatcelestialtitiesByType(domainType, tenantId, userId string) ([]models.RelationShip, int, error) {

	var relationships []models.RelationShip

	url := "http://localhost:9100/api/relationship/connections/" + userId + "/" + domainType
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, tenantId, nil)
	json.Unmarshal(response, &relationships)
	return relationships, status, err
}

func GetEntityTypes(tenantId string) ([]models.RelationShipType, int, error) {
	var relationshipType []models.RelationShipType

	url := "http://localhost:9100/api/relationship/types/" + tenantId
	method := "GET"

	response, status, err := clients.CallRestEndPoint(url, method, tenantId, nil)
	json.Unmarshal(response, &relationshipType)
	return relationshipType, status, err
}

func AddEntityRelationShip(relationship models.RelationShip, tenantId string) (models.RelationShip, int, error) {

	selectedUserProvider, err := tenant.GetProvidersForTenantByType(tenantId, "user")
	if err != nil {
		return relationship, 500, err
	}
	url := selectedUserProvider[0].celestialAdapter.AdapterUrl + "/relationship/" + tenantId
	method := "POST"

	r, _ := json.Marshal(relationship)

	response, status, err := clients.CallRestEndPoint(url, method, tenantId, r)
	json.Unmarshal(response, &relationship)
	return relationship, status, err
}
