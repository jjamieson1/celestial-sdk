package cms

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func DeleteImageByImageId(imageId, baseUrl, token, tenantId string) (status int, err error) {

	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + token,
	}

	url := baseUrl + "/api/v1/cms/image/" + imageId
	method := "DELETE"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return status, err
	}

	if status != 200 {
		err = fmt.Errorf("error deleting image with error: %+v", body)
	}

	var cms models.Cms
	json.Unmarshal(body, &cms)

	return status, err
}
