package user

import (
	"app/internal/core/cfg"
	"app/internal/pkg/user/ctrl"
	"app/internal/pkg/user/svc"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type Module struct {
	name    string
	version string
	API     *ctrl.UserController
}

// New creates a new instance of the module.
func New() *Module {
	// Create database instance
	database, err := gossiper.NewDB(
		gossiper.PostgresDB,
		cfg.Inst().PostgresDatabaseDSN,
		false,
		[]any{},
	)
	if err != nil {
		log.Fatalf("Failed to create database instance: %v", err)
	}

	// Create service and controller
	service := svc.NewUserService(database)
	controller := ctrl.NewUserController(service)

	// Initialize and return the module
	return &Module{
		name:    "User",
		version: "v1",
		API:     controller,
	}
}

// Initialize initializes the module. Currently not implemented.
func (m Module) Initialize() error {
	panic("Not implemented")
}

// Version returns the version of the module.
func (m Module) Version() string {
	return m.version
}

// Name returns the name of the module.
func (m Module) Name() string {
	return m.name
}
