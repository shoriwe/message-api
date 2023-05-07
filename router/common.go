package router

const (
	RootRoute     = "/"
	APIRoute      = "/api"
	RegisterRoute = "/register"
	LoginRoute    = "/login"
	LogoutRoute   = "/logout"
	UsersRoute    = "/users"
	MessageRoute  = "/message"
)

const (
	SessionVariable = "SESSION"
)

type Response struct {
	Result string `json:"result"`
}

var (
	UnauthorizedResponse = Response{Result: "unauthorized"}
	SucceedResponse      = Response{Result: "succeed"}
)
