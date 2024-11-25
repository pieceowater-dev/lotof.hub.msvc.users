package ctrl

import "app/internal/pkg/auth/svc"

type AuthController struct {
	authService *svc.AuthService
}

func NewAuthController(service *svc.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

// todo: implement methods
