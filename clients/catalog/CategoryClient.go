package catalog

import (
	"encoding/json"
	"fmt"
	"strconv"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetCategories(isDisplayed bool, tenantId, baseURL string) ([]models.Category, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	showAll := strconv.FormatBool(isDisplayed)
	url := fmt.Sprintf("%s/api/v1/catalog/categories?isDisplayed=%s", baseURL, showAll)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return nil, status, err
	}

	var categories []models.Category
	if err := json.Unmarshal(body, &categories); err != nil {
		return nil, status, err
	}

	return categories, status, err
}

func AddUpdateCategory(category models.Category, tenantId, baseURL string) (models.Category, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/categories", baseURL)
	method := "POST"

	body, err := json.Marshal(category)
	if err != nil {
		revel.AppLog.Errorf("AddUpdateCategory: unable to marshal category with error: %v", err.Error())
		return category, 500, err
	}

	body, status, err := clients.CallRestEndPoint(url, method, headers, body)
	if err != nil {
		return category, status, err
	}

	var response models.Category
	if err := json.Unmarshal(body, &response); err != nil {
		return category, status, err
	}

	return response, status, err
}
