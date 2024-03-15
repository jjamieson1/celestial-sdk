package cms

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetAllCms(tenantId, token string, baseUrl string) ([]models.Cms, error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + token,
	}

	url := baseUrl + "/api/v1/cms/content"
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, headers, nil)

	var cms []models.Cms
	json.Unmarshal(body, &cms)

	return cms, err
}

func GetCmsByCategoryId(tenantId, token, categoryId, baseUrl string) (content []models.Cms, status int, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + token,
	}
	url := baseUrl + "/api/v1/cms/content/category/" + categoryId
	method := "GET"
	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
	var cms []models.Cms
	json.Unmarshal(body, &cms)
	return cms, status, err
}

func GetCmsByCmsId(cmsId, tenantId, baseUrl string) (models.Cms, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content/" + cmsId
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)

	var cms models.Cms
	json.Unmarshal(body, &cms)

	return cms, status, err
}

func DeleteCmsByCmsId(tenantId string, cmsId string, baseUrl string) (int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content/" + cmsId
	method := "DELETE"

	_, status, err := client.CallRestEndPoint(url, method, headers, nil)

	return status, err
}

func AddUpdateCmsItem(cms models.Cms, tenantId, jwt string, baseUrl string) (cmsResponse models.Cms, status int, err error) {
	url := baseUrl + "/api/v1/cms/content"
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}
	b, err := json.Marshal(cms)
	if err != nil {
		return cmsResponse, status, err
	}

	response, status, err := client.CallRestEndPoint(url, "POST", headers, b)
	if status != 200 {
		err = fmt.Errorf(string(response))
	}

	json.Unmarshal(response, &cmsResponse)
	return cmsResponse, status, err
}
