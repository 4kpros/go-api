package config

import (
	"github.com/twilio/twilio-go"
)

var TwilioClient *twilio.RestClient

func SetupTwilioSMS() {
	TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   Env.TwilioApiKey,
		Password:   Env.TwilioApiSecret,
		AccountSid: Env.TwilioAccountSid,
	})
}
