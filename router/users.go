package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/controller"
	"github.com/shoriwe/message-api/models"
	"github.com/shoriwe/message-api/session"
)

func (r *Router) Users(ctx *gin.Context) {
	s := ctx.MustGet(SessionVariable).(*session.Session)
	var filter controller.Filter[models.User]
	bErr := ctx.Bind(&filter)
	if bErr != nil {
		return
	}
	results, sErr := r.ListUsers(s, &filter)
	if sErr == nil {
		ctx.JSON(http.StatusOK, results)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: sErr.Error()})
}
