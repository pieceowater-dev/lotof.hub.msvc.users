package ctrl

import (
	pbAuth "app/internal/core/grpc/generated/lotof.hub.msvc.users/auth"
	pbUser "app/internal/core/grpc/generated/lotof.hub.msvc.users/user"
	"app/internal/pkg/auth/svc"
	"app/internal/pkg/user/ent"
	"context"
)

type AuthController struct {
	authService *svc.AuthService
	pbAuth.UnimplementedAuthServiceServer
}

func NewAuthController(service *svc.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (a AuthController) ValidateToken(_ context.Context, req *pbAuth.ValidateTokenRequest) (*pbAuth.ValidateTokenResponse, error) {
	ok, user, err := a.authService.ValidateToken(req.Token)
	if err != nil {
		return &pbAuth.ValidateTokenResponse{
			Valid:   false,
			Message: err.Error(),
		}, nil
	}

	return &pbAuth.ValidateTokenResponse{
		Valid:   ok,
		Message: "",
		User: &pbUser.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (a AuthController) Login(_ context.Context, request *pbAuth.LoginRequest) (*pbAuth.AuthResponse, error) {
	token, user, err := a.authService.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &pbAuth.AuthResponse{
		Token: *token,
		User: &pbUser.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (a AuthController) Register(_ context.Context, request *pbAuth.RegisterRequest) (*pbAuth.AuthResponse, error) {
	token, user, err := a.authService.Register(&ent.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pbAuth.AuthResponse{
		Token: *token,
		User: &pbUser.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}
