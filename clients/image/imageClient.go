package image

import (
	"encoding/json"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

// /api/v1/catalog/image/:imageId

func GetImageById(imageId, tenantId, baseUrl string) (models.Image, int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image/" + imageId
	method := "GET"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)

	var image models.Image
	json.Unmarshal(body, &image)

	return image, status, err
}

func GetImages(imageIds []string, tenantId, baseUrl string) (images []models.Image, status int, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image"
	method := "POST"

	b, err := json.Marshal(imageIds)
	if err != nil {
		revel.AppLog.Errorf("client.GetImages: error marshalling object with error: %v", err.Error())
	}
	body, status, err := client.CallRestEndPoint(url, method, headers, b)

	json.Unmarshal(body, &images)

	return images, status, err
}

func DeleteImageById(imageId, tenantId, baseUrl string) (int, error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image/" + imageId
	method := "DELETE"

	body, status, err := client.CallRestEndPoint(url, method, headers, nil)

	var image models.Image
	json.Unmarshal(body, &image)

	return status, err
}
