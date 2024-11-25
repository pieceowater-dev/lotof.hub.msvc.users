package user

import (
	"app/internal/core/cfg"
	"app/internal/core/db"
	"app/internal/pkg/user/ctrl"
	"app/internal/pkg/user/svc"
	"log"
)

type Module struct {
	Controller *ctrl.UserController
}

func New() *Module {
	database, err := db.New(
		cfg.Inst().PostgresDatabaseDSN,
		false,
	).Create(db.PostgresDB)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	return &Module{
		Controller: ctrl.NewUserController(
			svc.NewUserService(database),
		),
	}
}
