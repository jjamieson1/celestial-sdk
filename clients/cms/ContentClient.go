package cms

import (
	"encoding/json"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func GetAllCms(tenantId string, provider models.TenantProvider) ([]models.Cms, error) {

	url := provider.Adapter.AdapterUrl + "/content"
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, tenantId, nil)

	var cms []models.Cms
	json.Unmarshal(body, &cms)

	return cms, err
}

func GetCmsByCategoryId(tenantId string, categoryId string, provider models.TenantProvider) ([]models.Cms, error) {

	url := provider.Adapter.AdapterUrl + "/category/" + categoryId
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, tenantId, nil)

	var cms []models.Cms
	json.Unmarshal(body, &cms)
	return cms, err
}

func GetCmsByCmsId(tenantId string, cmsId string, contentType string, provider models.TenantProvider) (models.Cms, error) {

	url := provider.Adapter.AdapterUrl + "/content/" + cmsId
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, tenantId, nil)

	var cmss []models.Cms
	var cms models.Cms
	json.Unmarshal(body, &cmss)

	if len(cmss) == 0 {
		return cms, err
	}

	cms = cmss[0]
	return cms, err
}

func CreateCms(tenantId, provider models.TenantProvider, cms models.Cms) (models.Cms, error) {
	url := provider.Adapter.AdapterUrl + "/content"

	p, err := json.Marshal(cms)

	response, _, err := clients.CallRestEndPoint(url, "POST", tenantId, p)
	var r models.Cms
	json.Unmarshal(response, &r)
	return r, err
}

func UpdateCmsContent(tenantId string, content models.Cms, provider models.TenantProvider) (models.Cms, error) {
	url := provider.Adapter.AdapterUrl + "/content/" + content.CmsId
	p, _ := json.Marshal(content)
	response, _, err := clients.CallRestEndPoint(url, "PUT", tenantId, p)
	var r models.Cms
	json.Unmarshal(response, &r)
	return r, err
}
