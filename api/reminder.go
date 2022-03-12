package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	db "go-bday-reminders/db/sqlc"
	"net/http"
)

type createReminderRequest struct {
	FullName       string `json:"full_name" binding:"required"`
	PersonalNumber int64  `json:"personal_number" binding:"required"`
	User           string `json:"user" binding:"required"`
	PhoneNumber    int64  `json:"phone_number" binding:"required"`
}

func (server *Server) createReminder(ctx *gin.Context) {
	var req createReminderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateReminderParams{
		FullName:       req.FullName,
		PersonalNumber: req.PersonalNumber,
		User:           req.User,
		PhoneNumber:    0,
	}

	reminder, err := server.store.CreateReminder(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, reminder)
}

type getReminderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getReminder(ctx *gin.Context) {
	var req getReminderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reminder, err := server.store.GetReminder(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, reminder)
}

type listReminderRequest struct {
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (server *Server) listReminder(ctx *gin.Context) {
	var req listReminderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRemindersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	reminders, err := server.store.ListReminders(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, reminders)
}
