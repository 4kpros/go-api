package utils

import (
	"encoding/json"
	"fmt"

	"github.com/4kpros/go-api/config"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

const unavailableMessage = "is not available! Please try again later"

// Sends a SMS to a specified receiver
func SendSMS(message string, receiver string) error {
	return sendSMSWithTwilio(message, receiver)
}

// Send a SMS to multiple recipients
func SendBulkSMS(message string, recipients []string) error {
	return sendBulkSMSWithTwilio(message, recipients)
}

// Sends a SMS to a specific receiver using Twilio API
func sendSMSWithTwilio(message string, receiver string) error {
	if config.TwilioClient == nil {
		return fmt.Errorf("%s", fmt.Sprintf("Twilio %s", unavailableMessage))
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(receiver)
	params.SetFrom(config.Env.TwilioSenderNumber)
	params.SetBody(message)

	resp, err := config.TwilioClient.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("Error sending SMS message: %s", err.Error())
	}

	response, _ := json.Marshal(*resp)
	fmt.Println("Response: " + string(response))
	return nil
}

// Send a SMS to multiple recipients using Twilio API
func sendBulkSMSWithTwilio(message string, receivers []string) error {
	if config.TwilioClient == nil {
		return fmt.Errorf("%s", fmt.Sprintf("Twilio %s", unavailableMessage))
	}
	return nil
}
