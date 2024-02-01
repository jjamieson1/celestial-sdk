package notification

import (
	"encoding/json"
	"fmt"
	"github.com/jjamieson1/celestial-sdk/models"
	"gopkg.in/resty.v1"
	"log"
)

func SendSMS(message string, phone string) (response *resty.Response, err error) {

	// body := "{\"MessageBody\": \"" + message + "\"}"

	body := &models.SmsMessage{MessageBody: message}
	jsonBody, err := json.Marshal(body)

	log.Printf("Sending message %v, to %v", message, "http://smsgateway.ca/services/message.svc/lQ183fwPufT0n0Y867Tow1haiOar6SH1/"+phone)
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBody(string(jsonBody)).
		Post("http://smsgateway.ca/services/message.svc/lQ183fwPufT0n0Y867Tow1haiOar6SH1/" + phone)
	if err != nil {
		log.Printf("Error!!!!!!! %v", err.Error())
	}

	fmt.Println("[INFO]  Response time " + resp.Time().String())

	if resp.StatusCode() != 200 {
		fmt.Printf("[ERROR]  SMS Status Code %v", resp.StatusCode())
		fmt.Printf("[ERROR]  Body  %v, error: %v ", string(jsonBody), string(resp.Body()))

	} else {
	}
	return resp, err
}
