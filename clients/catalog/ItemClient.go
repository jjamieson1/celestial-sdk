package catalog

import (
	"encoding/json"
	"fmt"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetItemGroupsByCatalogId(catalogId string, tenantId, baseURL string) ([]models.ItemGroup, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/itemgroups/%v", baseURL, catalogId)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return nil, status, err
	}

	var itemGroups []models.ItemGroup
	if err := json.Unmarshal(body, &itemGroups); err != nil {
		return nil, status, err
	}

	return itemGroups, status, err
}

func GetItemGroupById(itemGroupId string, tenantId, baseURL string) (models.ItemGroup, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/itemgroup/%v", baseURL, itemGroupId)
	method := "GET"

	body, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return models.ItemGroup{}, status, err
	}

	var itemGroup models.ItemGroup
	if err := json.Unmarshal(body, &itemGroup); err != nil {
		return models.ItemGroup{}, status, err
	}

	return itemGroup, status, err
}

func AddUpdateItem(item models.Item, tenantId, baseURL string) (responseItem models.Item, status int, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := fmt.Sprintf("%s/api/v1/catalog/item/addupdate", baseURL)
	method := "POST"

	b, err := json.Marshal(item)
	if err != nil {
		return responseItem, 400, err
	}

	revel.AppLog.Debugf("\n\n %v to URL %v with item \n\n %+v \n\n", method, url, string(b))

	body, status, err := clients.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return responseItem, status, err
	}

	err = json.Unmarshal(body, &responseItem)
	return responseItem, status, err
}
