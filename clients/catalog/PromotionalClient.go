package catalog

import (
	"encoding/json"
	"fmt"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

// localhost:3004/api/v1/catalog/campaigns
func GetPromotionalCampaigns(tenantId, baseURL string) (saleCampaigns []models.SaleCampaign, status int, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/campaigns", baseURL)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return saleCampaigns, status, err
	}

	if err := json.Unmarshal(body, &saleCampaigns); err != nil {
		return saleCampaigns, status, err
	}

	return saleCampaigns, status, err
}

func GetPromotionalCampaignById(campaignId, tenantId, baseUrl string) (saleCampaign models.SaleCampaign, status int, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/campaigns/%s", baseUrl, campaignId)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return saleCampaign, status, err
	}

	if err := json.Unmarshal(body, &saleCampaign); err != nil {
		return saleCampaign, status, err
	}

	return saleCampaign, status, err
}
