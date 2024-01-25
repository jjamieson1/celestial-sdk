package cms

import (
	"encoding/json"

	client "github.com/jjamieson1/eden-sdk/clients"
	clients2 "github.com/jjamieson1/eden-sdk/clients"
	"github.com/jjamieson1/eden-sdk/models"
)

func GetCmsFromProvider(tenantId string, lang, contentType string, provider models.TenantProvider) ([]models.Cms, error) {

	url := provider.EdenAdapter.AdapterUrl + "/cms/" + lang + "/type/" + contentType
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, tenantId, nil)

	var cms []models.Cms
	json.Unmarshal(body, &cms)

	return cms, err
}

func GetCmsFromProviderByCategoryId(tenantId string, lang, categoryId string, provider models.TenantProvider) ([]models.Cms, error) {

	url := provider.EdenAdapter.AdapterUrl + "/categories/" + categoryId
	method := "GET"

	body, _, err := client.CallRestEndPoint(url, method, tenantId, nil)

	var cms []models.Cms
	json.Unmarshal(body, &cms)
	return cms, err
}

func GetCmsFromProviderByCmsId(tenantId string, lang, cmsId string, contentType string, provider models.TenantProvider) (models.Cms, error) {

	url := provider.EdenAdapter.AdapterUrl + "/cms/" + lang + "/" + cmsId
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

func PostCmsToContentProvider(tenantId, lang string, provider models.TenantProvider, post models.Cms) (models.Cms, error) {
	url := provider.EdenAdapter.AdapterUrl + "/cms/" + lang

	p, _ := json.Marshal(post)

	response, _, err := clients2.CallRestEndPoint(url, "POST", tenantId, p)
	var r models.Cms
	json.Unmarshal(response, &r)
	return r, err
}

func UpdateCmsContent(tenantId string, content models.Cms, provider models.TenantProvider) (models.Cms, error) {
	url := provider.EdenAdapter.AdapterUrl + "/cms/" + content.CmsContent.Lang + "/" + content.CmsId
	p, _ := json.Marshal(content)
	response, _, err := clients2.CallRestEndPoint(url, "PUT", tenantId, p)
	var r models.Cms
	json.Unmarshal(response, &r)
	return r, err
}
