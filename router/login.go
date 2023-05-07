package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/controller"
)

func (r *Router) Login(ctx *gin.Context) {
	var login controller.Login
	bErr := ctx.Bind(&login)
	if bErr != nil {
		return
	}
	_, jwt, lErr := r.Controller.Login(&login)
	if lErr == nil {
		ctx.JSON(http.StatusOK, Response{Result: jwt})
		return
	}
	if errors.Is(lErr, controller.ErrUnauthorized) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: lErr.Error()})
}
