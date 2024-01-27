package tenant

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	clients "github.com/jjamieson1/eden-sdk/clients"
	"github.com/jjamieson1/eden-sdk/models"
)

func GetProvidersForTenantByType(tenantId, providerType string) ([]models.TenantProvider, error) {
	url := "http://127.0.0.1:9001/api/tenant/provider/" + providerType + "/" + tenantId
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, tenantId, nil)

	var tenantProvider []models.TenantProvider
	json.Unmarshal(body, &tenantProvider)

	if len(tenantProvider) == 0 {
		err = errors.New(fmt.Sprintf("System Error:  No %v provider selected for this tenant", providerType))
	}

	return tenantProvider, err
}

func GetProvidersType(providerType string) ([]models.EdenAdapter, error) {
	url := "http://127.0.0.1:9001/api/tenant/provider/type/" + providerType
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, providerType, nil)

	var tenantProvider []models.EdenAdapter
	json.Unmarshal(body, &tenantProvider)

	if len(tenantProvider) == 0 {
		err = errors.New("System Error:  No user provider selected for this tenant")
	}

	return tenantProvider, err
}

func GetTenantByUrl(url string) (models.Tenant, error) {
	var tenant models.Tenant
	tenantServiceUrl := "http://127.0.0.1:9001/api/tenant/details/url/" + url

	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, tenantServiceUrl, nil)
	if err != nil {
		return tenant, err
	}

	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("error calling the url: %v, with error: %v", url, err.Error())
		return tenant, errors.New(message)
	}

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		message := fmt.Sprintf("expected 200, but recieved %v calling the url: %v, with error: %v", res.StatusCode, tenantServiceUrl, string(body))
		return tenant, errors.New(message)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	json.Unmarshal(body, &tenant)

	return tenant, err
}

func GetTenantDetails(tenantId string) (models.Tenant, error) {
	var tenant models.Tenant
	tenantServiceUrl := "http://127.0.0.1:9001/api/v1/tenant/tenants/" + tenantId

	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, tenantServiceUrl, nil)
	if err != nil {
		return tenant, err
	}

	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		message := fmt.Sprintf("error calling the url: %v, with error: %v", tenantServiceUrl, err.Error())
		return tenant, errors.New(message)
	}

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		message := fmt.Sprintf("expected 200, but recieved %v calling the url: %v, with error: %v", res.StatusCode, tenantServiceUrl, string(body))
		return tenant, errors.New(message)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	json.Unmarshal(body, &tenant)

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
	url := "http://127.0.0.1:9001/api/tenant/type/"
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, tenantId, nil)

	var tenantTypes []models.TenantType
	json.Unmarshal(body, &tenantTypes)

	return tenantTypes, err
}

func GetTenantChildren(tenantId string) ([]models.Tenant, error) {
	url := "http://127.0.0.1:9001/api/tenant/children/" + tenantId
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, tenantId, nil)

	var tenant []models.Tenant
	json.Unmarshal(body, &tenant)

	return tenant, err
}

func GetTenants(tenantId string) ([]models.Tenant, error) {
	url := "http://127.0.0.1:9001/api/tenants"
	method := "GET"
	body, _, err := clients.CallRestEndPoint(url, method, tenantId, nil)

	var tenant []models.Tenant
	json.Unmarshal(body, &tenant)

	return tenant, err
}

func AddOrganizationType(orgType models.TenantType) (models.TenantType, error) {
	url := "http://127.0.0.1:9001/api/tenant/type/"
	method := "POST"
	p, _ := json.Marshal(orgType)
	body, _, err := clients.CallRestEndPoint(url, method, orgType.TenantId, p)

	var r models.TenantType
	json.Unmarshal(body, &r)
	return r, err
}

func AddOrganization(org models.Tenant) (models.Tenant, error) {
	url := "http://127.0.0.1:9001/api/tenant/details"
	method := "POST"
	p, _ := json.Marshal(org)
	body, _, err := clients.CallRestEndPoint(url, method, org.TenantId, p)

	var r models.Tenant
	json.Unmarshal(body, &r)
	return r, err
}

func UpdateOrganization(org models.Tenant) (models.Tenant, error) {
	url := "http://127.0.0.1:9001/api/tenant/details/" + org.TenantId
	method := "PUT"
	p, _ := json.Marshal(org)
	body, _, err := clients.CallRestEndPoint(url, method, org.TenantId, p)
	var r models.Tenant
	json.Unmarshal(body, &r)
	return r, err
}
