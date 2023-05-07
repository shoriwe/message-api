package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/controller"
)

func (r *Router) DownloadPicture(ctx *gin.Context) {
	userUUID := ctx.Param(UUIDParam)
	contents, dErr := r.Controller.DownloadPicture(userUUID)
	if dErr == nil {
		cType := http.DetectContentType(contents)
		ctx.Data(http.StatusOK, cType, contents)
		return
	}
	if errors.Is(dErr, controller.ErrUnauthorized) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{Result: dErr.Error()})
}
