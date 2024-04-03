package image

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

func GetImageById(imageId, tenantId, baseUrl string) (image models.Image, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image/" + imageId
	method := "GET"

	response, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return image, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return image, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &image)

	return image, err
}

func GetImages(imageIds []string, tenantId, baseUrl string) (images []models.Image, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image"
	method := "POST"

	b, err := json.Marshal(imageIds)
	if err != nil {
		revel.AppLog.Errorf("client.GetImages: error marshalling object with error: %v", err.Error())
	}

	response, status, err := client.CallRestEndPoint(url, method, headers, b)
	if err != nil {
		return images, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return images, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &images)

	return images, err
}

func GetImagesByItemId(itemId string, tenantId, baseUrl string) (images []models.Image, err error) {
	headers := map[string]string{
		"tenantId": tenantId,
	}

	url := baseUrl + "/api/v1/image/item/" + itemId
	method := "GET"

	response, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return images, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return images, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &images)

	return images, err
}

func DeleteImageByItemId(itemId, imageId, jwt, tenantId, baseUrl string) (err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}

	url := baseUrl + "/api/v1/image/itemId/" + itemId + "/image/" + imageId
	method := "DELETE"

	response, status, err := client.CallRestEndPoint(url, method, headers, nil)
	if err != nil {
		return err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return fmt.Errorf("%v", message)
	}

	return err
}

func AddImageByURL(newImage models.Image, itemId, jwt, tenantId, baseUrl string) (image models.Image, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}

	url := baseUrl + "/api/v1/image/url/" + itemId
	method := "POST"

	body, err := json.Marshal(newImage)
	if err != nil {
		return image, err
	}

	response, status, err := client.CallRestEndPoint(url, method, headers, body)
	if err != nil {
		return image, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return image, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &image)

	return image, err
}

func AddImageByFile(newImage models.Image, itemId, jwt, tenantId, baseUrl string) (image models.Image, err error) {
	headers := map[string]string{
		"tenantId":      tenantId,
		"Authorization": "Bearer " + jwt,
	}

	url := baseUrl + "/api/v1/image/file/" + itemId
	method := "POST"

	body, err := json.Marshal(newImage)
	if err != nil {
		return image, err
	}

	response, status, err := client.CallRestEndPoint(url, method, headers, body)
	if err != nil {
		return image, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(response, &message)
		return image, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(response, &image)

	return image, err
}
