package cms

import (
	"encoding/json"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetCategories(tenantId, baseUrl string) ([]models.CmsCategory, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := baseUrl + "/api/v1/cms/categories"
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)

	var categories []models.CmsCategory
	json.Unmarshal(body, &categories)

	return categories, status, err
}
