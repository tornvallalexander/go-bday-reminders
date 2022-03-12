package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	db "go-bday-reminders/db/sqlc"
	"net/http"
	"time"
)

type createBirthdayRequest struct {
	FullName       string    `json:"full_name" binding:"required"`
	FutureBirthday time.Time `json:"future_birthday" binding:"required"`
}

func (server *Server) createBirthday(ctx *gin.Context) {
	var req createBirthdayRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBirthdayParams{
		FullName:       req.FullName,
		FutureBirthday: req.FutureBirthday,
	}

	birthday, err := server.store.CreateBirthday(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, birthday)
}

type getBirthdayRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBirthday(ctx *gin.Context) {
	var req getBirthdayRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	birthday, err := server.store.GetBirthday(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, birthday)
}

type listBirthdayRequest struct {
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (server *Server) listBirthday(ctx *gin.Context) {
	var req listBirthdayRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBirthdaysParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	birthdays, err := server.store.ListBirthdays(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, birthdays)
}
