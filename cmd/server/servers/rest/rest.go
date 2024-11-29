package rest

import (
	"app/internal/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	Port   string
	Router *gin.Engine
}

func New(port string, router *gin.Engine, appRouter *pkg.Router) *Server {
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	appRouter.InitREST(router)
	return &Server{
		Port:   port,
		Router: router,
	}
}

func (r *Server) Start() error {
	log.Print("\033[32m")
	log.Printf("REST server running on port %s", r.Port)
	log.Print("\033[0m")
	return r.Router.Run(":" + r.Port)
}

func (r *Server) Stop() error {
	log.Println("REST server stopping")
	return nil
}
