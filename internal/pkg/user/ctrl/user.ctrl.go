package ctrl

import "app/internal/pkg/user/svc"

type UserController struct {
	userService *svc.UserService
}

func NewUserController(service *svc.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// todo: implement methods
