package catalog

import (
	"encoding/json"
	"fmt"
	"strconv"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetCategories(displayAll bool, page, pageSize, tenantId, baseURL string) (categories []models.Category, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	showAll := strconv.FormatBool(displayAll)
	// + "?page=" + page + "&pageSize=" + pageSize
	url := fmt.Sprintf("%s/api/v1/catalog/categories?isDisplayed=%s&page=%s&pageSize=%s", baseURL, showAll, page, pageSize)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return categories, fmt.Errorf("%v", message)
	}

	if err != nil {
		return categories, err
	}

	err = json.Unmarshal(body, &categories)

	return categories, err
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
