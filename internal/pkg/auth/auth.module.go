package auth

import (
	"app/internal/core/cfg"
	"app/internal/core/db"
	"app/internal/pkg/auth/ctrl"
	"app/internal/pkg/auth/svc"
	"log"
)

type Module struct {
	Controller *ctrl.AuthController
}

func New() *Module {
	// Create database instance
	database, err := db.New(
		cfg.Inst().PostgresDatabaseDSN,
		false,
	).Create(db.PostgresDB)
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
