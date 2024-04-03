package tenant

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func Test(baseUrl string) (dbStats interface{}, err error) {

	url := baseUrl + "/api/v1/tenant/status/report"
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, nil, nil)
	if err != nil {
		return dbStats, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(body, &message)
		return dbStats, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(body, &dbStats)

	return dbStats, err
}
