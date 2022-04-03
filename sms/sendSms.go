package sms

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	db "go-bday-reminders/db/sqlc"
	"go-bday-reminders/utils"
	"time"
)

// SendSmsReminder sends a message with all reminders for a certain month to current user
func SendSmsReminder(receiver string, birthdays []db.Reminder) (resp *openapi.ApiV2010Message, err error) {
	config, err := utils.LoadConfig("..")
	if err != nil {
		return
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.TwilioAccountSid,
		Password: config.TwilioAuthToken,
	})

	currMonth := time.Now().Month()

	params := &openapi.CreateMessageParams{}
	params.SetTo(receiver)
	params.SetFrom(config.TwilioSender)

	if len(birthdays) == 0 {
		params.SetBody(fmt.Sprintf("No active birthday reminders for month of %s", currMonth))
	} else {
		var message string
		for i := range birthdays {
			b := birthdays[i]
			message = message + fmt.Sprintf("%s %d %s\n", b.FullName, b.PersonalNumber, b.PhoneNumber)
		}
		params.SetBody(message)
	}

	resp, err = client.ApiV2010.CreateMessage(params)
	return
}
