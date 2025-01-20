package pkg

import (
	"app/internal/core/generic/interfaces"
	pbAuth "app/internal/core/grpc/generated/lotof.hub.msvc.users/auth"
	pbFriendship "app/internal/core/grpc/generated/lotof.hub.msvc.users/friendship"
	pbUser "app/internal/core/grpc/generated/lotof.hub.msvc.users/user"
	"app/internal/pkg/auth"
	"app/internal/pkg/friendship"
	"app/internal/pkg/user"
	"google.golang.org/grpc"
)

type Router struct {
	modules map[string]interfaces.IModule // Map of module names to their instances.

	server *grpc.Server
}

// NewRouter creates a new Router instance and initializes the module.
func NewRouter(server *grpc.Server) *Router {
	authModule := auth.New()
	userModule := user.New()
	friendModule := friendship.New()

	return &Router{
		server: server,
		modules: map[string]interfaces.IModule{
			authModule.Name():   authModule,
			userModule.Name():   userModule,
			friendModule.Name(): friendModule,
		},
	}
}

// InitializeRouter initializes the router and its gRPC routes.
func (r *Router) InitializeRouter() (any, error) {
	r.InitializeGRPCRoutes(r.server)
	return nil, nil
}

// InitializeGRPCRoutes registers the gRPC routes for the modules.
func (r *Router) InitializeGRPCRoutes(grpcServer *grpc.Server) {
	pbAuth.RegisterAuthServiceServer(grpcServer, r.modules["Auth"].(*auth.Module).API)
	pbUser.RegisterUserServiceServer(grpcServer, r.modules["User"].(*user.Module).API)
	pbFriendship.RegisterFriendshipServiceServer(grpcServer, r.modules["Friendship"].(*friendship.Module).API)
}

// GetModules returns the map of modules.
func (r *Router) GetModules() map[string]interfaces.IModule {
	return r.modules
}
