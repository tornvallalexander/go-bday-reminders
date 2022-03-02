package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"go-bday-reminders/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: config.AccountSid,
		Password: config.AuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(config.Receiver)
	params.SetFrom(config.Sender)
	params.SetBody("Hello from Go!")

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println("Message Sid: " + *resp.Sid)
	}
}
