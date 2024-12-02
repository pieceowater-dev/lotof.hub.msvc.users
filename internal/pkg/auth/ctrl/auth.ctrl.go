package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/auth/svc"
	"app/internal/pkg/user/ent"
	"context"
)

type AuthController struct {
	authService *svc.AuthService
	pb.UnimplementedAuthServiceServer
}

func NewAuthController(service *svc.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (a AuthController) Login(_ context.Context, request *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, user, err := a.authService.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Token: *token,
		User: &pb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (a AuthController) Register(_ context.Context, request *pb.RegisterRequest) (*pb.AuthResponse, error) {
	token, user, err := a.authService.Register(&ent.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		Token: *token,
		User: &pb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}
