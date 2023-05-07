package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/models"
)

func (r *Router) Register(ctx *gin.Context) {
	var user models.User
	bErr := ctx.Bind(&user)
	if bErr != nil {
		return
	}
	rErr := r.Controller.Register(&user)
	if rErr == nil {
		ctx.JSON(http.StatusCreated, SucceedResponse)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: rErr.Error()})
}
