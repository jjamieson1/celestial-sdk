package cms

import (
	"encoding/json"
	"fmt"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetCategories(tenantId, baseUrl string) (categories []models.CmsCategory, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := baseUrl + "/api/v1/cms/categories"
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return categories, err
	}
	if status != 200 {
		revel.AppLog.Errorf("error getting categories with status error: %v", status)
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return categories, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(body, &categories)

	return categories, err
}

func AddCategory(category models.CmsCategory, jwt, tenantId, baseUrl string) (response models.CmsCategory, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}
	url := baseUrl + "/api/v1/cms/categories"
	method := "POST"

	newCategory := category

	b, err := json.Marshal(newCategory)
	if err != nil {
		return category, err
	}

	body, status, err := clients.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return category, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return category, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

func DeleteCategory(categoryId, jwt, tenantId, baseUrl string) error {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}
	url := baseUrl + "/api/v1/cms/categories/" + categoryId
	method := "DELETE"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return fmt.Errorf("%v", message)
	}

	return nil
}
