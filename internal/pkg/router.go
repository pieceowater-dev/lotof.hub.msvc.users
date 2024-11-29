package pkg

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Router struct {
	//userModule       *user.Module
	//authModule       *auth.Module
	//friendshipModule *friendship.Module
}

func NewRouter() *Router {
	return &Router{
		//userModule:       user.New(),
		//authModule:       auth.New(),
		//friendshipModule: friendship.New(),
	}
}

// InitGRPC initializes gRPC routes
func (r *Router) InitGRPC(grpcServer *grpc.Server) {
	// Register gRPC services
	// pb.RegisterUserServiceServer(grpcServer, r.userModule.Controller)
}

// InitREST initializes REST routes using Gin
func (r *Router) InitREST(router *gin.Engine) {
	//api := router.Group("/api")
	{
		// Register GIN routes
		//api.GET("/users", r.userModule.GetUsersHandler)
		//api.POST("/auth/login", r.authModule.LoginHandler)
		//api.POST("/friendships", r.friendshipModule.CreateFriendshipHandler)
	}
}
