package friendship

import (
	"app/internal/core/cfg"
	"app/internal/core/db"
	"app/internal/pkg/friendship/ctrl"
	"app/internal/pkg/friendship/svc"
	"log"
)

type Module struct {
	Controller *ctrl.FriendshipController
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

	return &Module{
		Controller: ctrl.NewFriendshipController(
			svc.NewFriendshipService(database),
		),
	}

}
