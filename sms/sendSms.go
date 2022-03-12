package sms

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"go-bday-reminders/utils"
)

// SendSms send a message to the receiver with name and birthday
func SendSms(name string, birthday string) (resp *openapi.ApiV2010Message, err error) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.TwilioAccountSid,
		Password: config.TwilioAuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(config.TwilioReceiver)
	params.SetFrom(config.TwilioSender)
	params.SetBody(fmt.Sprintf("Födelsedag alert: %s, %s", name, birthday))

	resp, err = client.ApiV2010.CreateMessage(params)
	return
}
