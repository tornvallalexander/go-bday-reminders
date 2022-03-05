package main

import (
	"fmt"
	"go-bday-reminders/sms"
)

func main() {
	resp, err := sms.SendSms("Alexander", "2023-03-02")
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		fmt.Println("Message Sid: " + *resp.Sid)
	}
}
