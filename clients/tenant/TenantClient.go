package tenant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	clients "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetTenantIdByUrl(url, baseUrl string) (tenant models.Tenant, err error) {

	hostUrl := baseUrl + "/api/v1/configuration/url/" + url

	method := "GET"

	response, status, err := clients.CallRestEndPoint(hostUrl, method, nil, nil)
	if err != nil {
		revel.AppLog.Errorf("error calling out for a tenantId with error: %v", err.Error())
		return tenant, err
	}
	if status != 200 {
		revel.AppLog.Errorf("non-success status recieved: %v", status)
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return tenant, fmt.Errorf("error %v", message)
	}

	err = json.Unmarshal(response, &tenant)
	return tenant, err
}

// /api/v1/configuration/tenant/name/:tenantId
func GetTenantNameByTenantId(tenantId, baseUrl string) (tenant models.Tenant, err error) {
	hostUrl := baseUrl + "/api/v1/configuration/tenant/name/" + tenantId

	method := "GET"

	response, status, err := clients.CallRestEndPoint(hostUrl, method, nil, nil)
	if err != nil {
		return tenant, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return tenant, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &tenant)

	return tenant, err
}

func GetTenantDetails(appKey, apiKey, baseUrl string) (tenant models.Tenant, err error) {

	hostUrl := baseUrl + "/api/v1/configuration/tenant"
	headers := map[string]string{
		"api-key": apiKey,
		"app-key": appKey,
		"Accept":  "application/json",
	}

	method := "GET"

	response, status, err := clients.CallRestEndPoint(hostUrl, method, headers, nil)
	if err != nil {
		return tenant, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return tenant, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &tenant)

	return tenant, err
}

func GetTenantDetailsByJWT(jwt, tenantId string) (tenant models.Tenant, err error) {

	url := "http://127.0.0.1:3002/api/v1/configuration/tenant/" + tenantId
	method := "GET"

	headers := map[string]string{
		"Authorization": "Bearer " + jwt,
	}

	revel.AppLog.Debugf("headers: %+v", headers)

	response, status, err := clients.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return tenant, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return tenant, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &tenant)
	return tenant, err
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
