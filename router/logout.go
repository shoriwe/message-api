package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/session"
)

func (r *Router) Logout(ctx *gin.Context) {
	s := ctx.MustGet(SessionVariable).(*session.Session)
	lErr := r.Controller.Logout(s)
	if lErr == nil {
		ctx.JSON(http.StatusOK, SucceedResponse)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: lErr.Error()})
}
