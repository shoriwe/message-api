package router

const (
	UUIDParam = "uuid"
)

const (
	RootRoute              = "/"
	APIRoute               = "/api"
	RegisterRoute          = "/register"
	LoginRoute             = "/login"
	LogoutRoute            = "/logout"
	UsersRoute             = "/users"
	PictureRoute           = "/picture/"
	PictureRouteWithParams = PictureRoute + ":" + UUIDParam
	MessageRoute           = "/message"
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
