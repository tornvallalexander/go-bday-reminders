package api

import (
	"github.com/gin-gonic/gin"
	db "go-bday-reminders/db/sqlc"
	"go-bday-reminders/sms"
	"go-bday-reminders/token"
	"go-bday-reminders/utils"
	"net/http"
	"strconv"
	"time"
)

func (server *Server) getSmsReminders(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	reminders, err := server.store.GetSmsReminders(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var monthReminders []db.Reminder
	m := time.Now().Month()
	currentMonth := int(m)

	for i := range reminders {
		reminder := reminders[i]

		pnr := strconv.Itoa(int(reminder.PersonalNumber))
		pnrMonth := utils.Substr(pnr, 4, 2)
		pnrMonthInt, _ := strconv.Atoi(pnrMonth)

		if currentMonth == pnrMonthInt {
			monthReminders = append(monthReminders, reminder)
		}
	}

	_, err = sms.SendSmsReminder(user.PhoneNumber, monthReminders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, monthReminders)
}
