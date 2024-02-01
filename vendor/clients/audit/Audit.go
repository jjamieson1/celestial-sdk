package audit

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jjamieson1/celestial-sdk/models"
	"gopkg.in/resty.v1"
)

func Error(message models.AuditMessage) {

	//  This is the parent function that will perform the appropriate logging or audit action.
	//  There should be a service endpoint for both audit and logging that may have different strategies.
	//  The type will define either audit or log and take actions based on the type.
	t, err := uuid.NewUUID()
	if err != nil {
		message.TraceId = err.Error()
	}
	message.TraceId = t.String()

	jsonBody, err := json.Marshal(message)

	fmt.Printf("Sending object %+v, to %v  ", message, "http://audit-service:3000/api/v1/audit")
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBody(string(jsonBody)).
		Post("http://audit-service:3000/api/v1/audit")
	if err != nil {
		fmt.Errorf("error sending audit event: %v", err.Error())
	}
	fmt.Println("audit client response time " + resp.Time().String())

	if resp.StatusCode() != 200 {
		fmt.Errorf("[ERROR]  audit status Code %v", resp.StatusCode())
	}
}
