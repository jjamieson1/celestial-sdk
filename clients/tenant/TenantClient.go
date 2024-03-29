package tenant

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetProvidersForTenantByType(tenantId, providerType string) ([]models.TenantProvider, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://127.0.0.1:9001/api/tenant/provider/" + providerType + "/" + tenantId
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, headers, nil)

	var tenantProvider []models.TenantProvider
	json.Unmarshal(body, &tenantProvider)

	if len(tenantProvider) == 0 {
		err = errors.New(fmt.Sprintf("System Error:  No %v provider selected for this tenant", providerType))
	}

	return tenantProvider, err
}

func GetTenantIdByUrl(url string) (string, int, error) {

	var tenant map[string]string
	hostUrl := "http://127.0.0.1:3002/api/v1/configuration/url/" + url

	method := "GET"

	response, status, err := clients.CallRestEndPoint(hostUrl, method, nil, nil)
	json.Unmarshal(response, &tenant)
	return tenant["tenantId"], status, err
}

func GetTenantDetails(appKey, apiKey string) (*models.Tenant, int, error) {
	var tenant models.Tenant
	tenantServiceUrl := "http://127.0.0.1:3002/api/v1/configuration/tenant"

	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, tenantServiceUrl, nil)
	if err != nil {
		return &tenant, 0, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("api-key", apiKey)
	req.Header.Add("app-key", appKey)
	res, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("error calling the url: %v, with error: %v", tenantServiceUrl, err.Error())
		return &tenant, 500, errors.New(message)
	}

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		message := fmt.Sprintf("expected 200, but recieved %v calling the url: %v, with error: %v", res.StatusCode, tenantServiceUrl, string(body))
		return &tenant, res.StatusCode, errors.New(message)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	json.Unmarshal(body, &tenant)

	return &tenant, 200, err
}

func GetTenantDetailsByJWT(jwt, tenantId string) (tenant models.Tenant, status int, err error) {

	url := "http://127.0.0.1:3002/api/v1/configuration/tenant/" + tenantId
	method := "GET"

	headers := map[string]string{
		"Authorization": "Bearer " + jwt,
	}

	revel.AppLog.Debugf("headers: %+v", headers)

	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	json.Unmarshal(response, &tenant)
	return tenant, status, err
}

func GetUserServiceProvider(tenantId string) ([]models.TenantProvider, error) {
	url := fmt.Sprintf("http://127.0.0.1:9001/api/tenant/provider/user/%v", tenantId)
	method := "GET"
	var provider []models.TenantProvider
	var client = &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return provider, err
	}

	req.Header.Add("Accept", "application/json")
	//req.Header.Add("Authorization", authorizationString)

	res, err := client.Do(req)

	if err != nil || res.StatusCode == 0 {
		return provider, err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return provider, err
	}

	json.Unmarshal(body, &provider)
	return provider, err
}

func GetTenantTypes(tenantId string) ([]models.TenantType, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := "http://127.0.0.1:9001/api/tenant/type/"
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, headers, nil)

	var tenantTypes []models.TenantType
	json.Unmarshal(body, &tenantTypes)

	return tenantTypes, err
}

func GetTenantChildren(tenantId string) ([]models.Tenant, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://127.0.0.1:9001/api/tenant/children/" + tenantId
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, headers, nil)

	var tenant []models.Tenant
	json.Unmarshal(body, &tenant)

	return tenant, err
}

func GetTenants(tenantId string) ([]models.Tenant, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}
	url := "http://127.0.0.1:9001/api/tenants"
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, headers, nil)

	var tenant []models.Tenant
	json.Unmarshal(body, &tenant)

	return tenant, err
}
