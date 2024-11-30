package auth

import (
	"app/internal/core/cfg"
	"app/internal/pkg/auth/ctrl"
	"app/internal/pkg/auth/svc"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.AuthController
}

func New() *Module {
	// Create database instance
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	// Initialize and return the module
	return &Module{
		Controller: ctrl.NewAuthController(
			svc.NewAuthService(database),
		),
	}
}
