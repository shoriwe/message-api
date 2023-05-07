package router

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/shoriwe/message-api/controller"
)

type Router struct {
	*controller.Controller
	*gin.Engine
}

func New(c *controller.Controller, e *gin.Engine) *Router {
	r := &Router{Controller: c, Engine: e}
	api := r.Group(APIRoute)
	// Public
	public := api.Group(RootRoute)
	public.PUT(RegisterRoute, r.Register)
	public.POST(LoginRoute, r.Login)
	// Auth required
	auth := api.Group(RootRoute, r.Authenticate)
	auth.POST(LogoutRoute, r.Logout)
	auth.POST(UsersRoute, r.Users)
	auth.GET(PictureRouteWithParams, r.DownloadPicture)
	auth.PUT(MessageRoute, r.SendMessage)
	return r
}

func NewMem() *Router {
	return New(controller.NewMem(), gin.Default())
}

func NewTest(t *testing.T) (serverClose func(), r *Router, e *httpexpect.Expect) {
	r = NewMem()
	s := httptest.NewServer(r)
	e = httpexpect.Default(t, s.URL+APIRoute)
	return s.Close, r, e
}
