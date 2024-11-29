package grpc

import (
	"app/internal/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	Port   string
	Server *grpc.Server
	Router *pkg.Router
}

func New(port string, server *grpc.Server, router *pkg.Router) *Server {
	router.InitGRPC(server)
	return &Server{
		Port:   port,
		Server: server,
		Router: router,
	}
}

func (g *Server) Start() error {
	listener, err := net.Listen("tcp", ":"+g.Port)
	if err != nil {
		return err
	}
	log.Print("\033[32m")
	log.Printf("gRPC server running on port %s", g.Port)
	log.Print("\033[0m")
	return g.Server.Serve(listener)
}

func (g *Server) Stop() error {
	g.Server.GracefulStop()
	log.Println("gRPC server stopped")
	return nil
}
