package sms

import (
	"api/config"
	"encoding/json"
	"fmt"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Sends a SMS to a specific receiver using Twilio API
func sendSMSWithTwilio(message string, receiver string) error {
	if config.TwilioClient == nil {
		return fmt.Errorf("%s %s",
			"Twilio",
			unavailableErrMsg,
		)
	}
	if len(receiver) < 1 {
		return fmt.Errorf("%s", emptyReceiverErrMsg)
	}
	if len(message) < 1 {
		return fmt.Errorf("%s", emptyMessageErrMsg)
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(receiver)
	params.SetFrom(config.Env.TwilioSenderNumber)
	params.SetBody(message)

	resp, err := config.TwilioClient.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("%s%s",
			"Error sending Twilio SMS message: ",
			err.Error(),
		)
	}

	response, _ := json.Marshal(*resp)
	fmt.Println("Twilio response: " + string(response))
	return nil
}

// Send a SMS to multiple recipients using Twilio API
func sendBulkSMSWithTwilio(message string, receivers []string) error {
	if config.TwilioClient == nil {
		return fmt.Errorf("%s %s",
			"Twilio",
			unavailableErrMsg,
		)
	}
	if len(receivers) < 1 {
		return fmt.Errorf("%s", emptyRecipientsErrMsg)
	}
	if len(message) < 1 {
		return fmt.Errorf("%s", emptyMessageErrMsg)
	}
	return nil
}
