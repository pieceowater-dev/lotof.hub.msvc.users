package friendship

import (
	"app/internal/core/cfg"
	"app/internal/pkg/friendship/ctrl"
	"app/internal/pkg/friendship/svc"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	Controller *ctrl.FriendshipController
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

	return &Module{
		Controller: ctrl.NewFriendshipController(
			svc.NewFriendshipService(database),
		),
	}

}
