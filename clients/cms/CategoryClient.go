package cms

import (
	"encoding/json"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetCategoryFromProvider(tenantId, lang string, provider models.TenantProvider) ([]models.CmsCategory, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := provider.Adapter.AdapterUrl + "/categories"
	method := "GET"

	body, _, err := clients.CallRestEndPoint(url, method, headers, nil)

	var categories []models.CmsCategory
	json.Unmarshal(body, &categories)

	return categories, err
}
