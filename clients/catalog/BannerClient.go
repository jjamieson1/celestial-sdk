package catalog

import (
	"encoding/json"
	"fmt"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetBanners(catagoryId string, tenantId, baseURL string) ([]models.Banner, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/marketing/banners/%v", baseURL, catagoryId)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return nil, status, err
	}

	var banners []models.Banner
	if err := json.Unmarshal(body, &banners); err != nil {
		return nil, status, err
	}

	return banners, status, err
}
