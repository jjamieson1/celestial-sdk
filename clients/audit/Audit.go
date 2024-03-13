package audit

import (
	"encoding/json"
	"fmt"

	"github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func LogItemChange(log models.ItemLog, baseUrl string) (err error) {
	headers := map[string]string{
		"tenantId": log.TenantId,
	}

	url := fmt.Sprintf("%s/api/v1/audit/log-item", baseUrl)
	method := "POST"

	b, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("error marshalling the log to json with error: %v", err.Error())
	}

	_, status, err := clients.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return err
	}
	if status != 200 {
		return fmt.Errorf("non-success status returned:  http status: %v", status)
	}

	return err
}
