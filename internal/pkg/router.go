package pkg

import (
	"app/internal/pkg/auth"
	"app/internal/pkg/friendship"
	"app/internal/pkg/user"
	"google.golang.org/grpc"
)

type Router struct {
	userModule       *user.Module
	authModule       *auth.Module
	friendshipModule *friendship.Module
}

func NewRouter() *Router {
	return &Router{
		userModule:       user.New(),
		authModule:       auth.New(),
		friendshipModule: friendship.New(),
	}
}

func (r *Router) Init(grpcServer *grpc.Server) {
	// Register gRPC services
	//pb.RegisterTodoServiceServer(grpcServer, r.todoModule.Controller)
}
