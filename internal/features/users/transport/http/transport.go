package users_transport_http

type UsersHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(usersService UserSservice) *UsersHTTPHandler {
	return &UsersHTTPHandler{usersService: usersService}
}
