package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
)

func (r *Router) SendMessage(ctx *gin.Context) {
	s := ctx.MustGet(SessionVariable).(*session.Session)
	var msg models.Message
	bErr := ctx.Bind(&msg)
	if bErr != nil {
		return
	}
	sErr := r.Controller.SendMessage(s, &msg)
	if sErr == nil {
		ctx.JSON(http.StatusCreated, SucceedResponse)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: sErr.Error()})
}
