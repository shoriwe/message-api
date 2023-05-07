package router

import (
	"encoding/base64"
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/controller"
)

var bearerRegexp = regexp.MustCompile(`(?m)Bearer\s+`)

func (r *Router) Authenticate(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	b64Token := bearerRegexp.ReplaceAllString(auth, "")
	token, pErr := base64.StdEncoding.DecodeString(b64Token)
	if pErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Result: pErr.Error()})
		return
	}
	session, tok, pErr := r.Controller.JWT.Parse(string(token))
	if pErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Result: pErr.Error()})
		return
	}
	if !tok.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse)
		return
	}
	_, authErr := r.Controller.Authenticate(session)
	if authErr != nil {
		if errors.Is(authErr, controller.ErrUnauthorized) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, UnauthorizedResponse)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{Result: authErr.Error()})
		return
	}
	ctx.Set(SessionVariable, session)
	ctx.Next()
}
