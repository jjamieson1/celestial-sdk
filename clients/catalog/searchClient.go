package catalog

import (
	"encoding/json"
	"fmt"
	"strconv"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func Search(searchText, baseUrl string, catalogId int, tenantId string) ([]models.SearchResult, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	query := make(map[string]string)
	query["catId"] = strconv.Itoa(catalogId)
	query["text"] = searchText

	b, err := json.Marshal(query)
	if err != nil {
		return nil, 500, err
	}

	url := fmt.Sprintf("%s/api/v1/catalog/search", baseUrl)
	method := "POST"

	body, status, err := clients.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return nil, status, err
	}

	var searchResult []models.SearchResult
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return nil, status, err
	}

	return searchResult, status, err
}
