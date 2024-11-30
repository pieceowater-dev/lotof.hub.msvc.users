package cfg

import (
	frinedship "app/internal/pkg/friendship/ent"
	user "app/internal/pkg/user/ent"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Config struct {
	GrpcPort            string
	RestPort            string
	PostgresDatabaseDSN string
	PostgresModels      []any
}

var (
	once     sync.Once
	instance *Config
)

func Inst() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found, loading from OS environment variables.")
		}

		instance = &Config{
			GrpcPort:            getEnv("GRPC_PORT", "50051"),
			RestPort:            getEnv("REST_PORT", "3000"),
			PostgresDatabaseDSN: getEnv("POSTGRES_DB_DSN", "postgres://pieceouser:pieceopassword@localhost:5432/users?sslmode=disable"),
			PostgresModels: []any{
				// models to migration here:
				// &ent.MyModel{},
				&user.User{},
				&frinedship.Friendship{},
			},
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
