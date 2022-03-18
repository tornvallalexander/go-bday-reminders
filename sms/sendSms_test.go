package sms

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"go-bday-reminders/utils"
	"testing"
)

func TestSendSms(t *testing.T) {
	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.TwilioAccountSid,
		Password: config.TwilioAuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(config.TwilioReceiver)
	params.SetFrom(config.TwilioSender)
	params.SetBody(fmt.Sprintf("Test SMS: %s, %s", "Test", "0001-01-01"))

	resp, err := client.ApiV2010.CreateMessage(params)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}
