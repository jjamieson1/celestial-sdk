package messaging

import (
	"encoding/json"
	"fmt"

	client "github.com/jjamieson1/celestial-sdk/clients"
	"github.com/jjamieson1/celestial-sdk/models"
)

func SendSMS(message string, phone string) (response interface{}, err error) {

	url := "https://smsgateway.ca/services/message.svc/lQ183fwPufT0n0Y867Tow1haiOar6SH1/" + phone
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	body := &models.SmsMessage{MessageBody: message}
	b, err := json.Marshal(body)
	if err != nil {
		return response, err
	}

	responseBody, status, err := client.CallRestEndPoint(url, "POST", headers, b)
	if err != nil {
		return response, err
	}
	if status != 200 {
		var message []models.ValidationError
		json.Unmarshal(responseBody, &response)
		return response, fmt.Errorf("%v", message)
	}

	err = json.Unmarshal(responseBody, &response)

	return response, err
}
