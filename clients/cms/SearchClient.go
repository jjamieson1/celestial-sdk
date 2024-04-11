package cms

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func SeachCMS(searchText, categoryId, baseUrl, tenantId string) (content []models.CMSSearchResult, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	search := make(map[string]string)
	search["text"] = searchText
	search["categoryId"] = categoryId
	b, err := json.Marshal(search)
	if err != nil {
		return content, err
	}
	url := baseUrl + "/api/v1/cms/content/search/cms"
	method := "POST"

	body, status, err := client.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return content, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return content, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(body, &content)

	return content, err
}
