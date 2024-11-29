package main

import (
	"app/cmd/server/servers"
	grpcServ "app/cmd/server/servers/grpc"
	restServ "app/cmd/server/servers/rest"
	"app/internal/core/cfg"
	"app/internal/pkg"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	appCfg := cfg.Inst()
	appRouter := pkg.NewRouter()

	serverManager := servers.NewServerManager()
	serverManager.AddServer(grpcServ.New(appCfg.GrpcPort, grpc.NewServer(), appRouter))
	serverManager.AddServer(restServ.New(appCfg.RestPort, gin.Default(), appRouter))

	serverManager.StartAll()
	defer serverManager.StopAll()
}
