package clients

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jjamieson1/celestial-sdk/models"
	"gopkg.in/resty.v1"
)

func GetConfigurationForProvider() (models.celestialConfiguration, error) {
	var celestialConfiguration models.celestialConfiguration
	var traceId string
	t, err := uuid.NewUUID()
	if err != nil {
		traceId = err.Error()
	}
	traceId = t.String()

	fmt.Printf("Sending request, to %v  ", "http://configuration-service:3002/api/v1/configuration")
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("traceId", traceId).
		Get("http://configuration-service:3002/api/v1/configuration")
	if err != nil {
		fmt.Errorf("error sending configuration request: %v", err.Error())
	}
	fmt.Println("configuration client response time " + resp.Time().String())

	if resp.StatusCode() != 200 {
		fmt.Errorf("[ERROR]  configuration status Code %v", resp.StatusCode())
	} else {
		err = json.Unmarshal(resp.Body(), &celestialConfiguration)
	}
	return celestialConfiguration, err
}
