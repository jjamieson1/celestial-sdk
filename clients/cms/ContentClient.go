package cms

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetAllCms(tenantId, token string, baseUrl string) (content []models.Cms, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content"
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
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

func GetCmsByCategoryId(tenantId, categoryId, baseUrl string) (content []models.Cms, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content/category/" + categoryId
	method := "GET"
	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return content, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return content, fmt.Errorf("%v", message)
	}

	var cms []models.Cms
	err = json.Unmarshal(body, &cms)

	return cms, err
}

func GetCmsByCmsId(cmsId, tenantId, baseUrl string) (cms models.Cms, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content/" + cmsId
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return cms, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return cms, fmt.Errorf("%v", message)
	}

	json.Unmarshal(body, &cms)

	return cms, err
}

func DeleteCmsByCmsId(tenantId, jwt, cmsId, baseUrl string) error {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}

	url := baseUrl + "/api/v1/cms/content/" + cmsId
	method := "DELETE"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
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

func AddUpdateCmsItem(cms models.Cms, tenantId, jwt string, baseUrl string) (cmsResponse models.Cms, err error) {
	url := baseUrl + "/api/v1/cms/content"
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}

	b, err := json.Marshal(cms)
	if err != nil {
		return cmsResponse, err
	}

	response, status, err := client.CallRestEndPoint(url, "POST", headers, b)
	if err != nil {
		return cmsResponse, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return cmsResponse, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &cmsResponse)

	return cmsResponse, err
}

func SearchCms(searchTerm, tenantId, baseUrl string) (cms []models.Cms, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/cms/content/search/" + searchTerm
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return cms, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return cms, fmt.Errorf("%v", message)
	}

	json.Unmarshal(body, &cms)

	return cms, err
}
