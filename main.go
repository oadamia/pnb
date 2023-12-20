package main

import (
	"fmt"
	"log/slog"
	"pnb/api"
	"pnb/service"
	"pnb/worker"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	slog.Info("Starting up...")

	var cfg Config
	err := envconfig.Process("pnb", &cfg)
	if err != nil {
		slog.Error("Failed to process config", err)
		panic(err)
	}

	slog.Info("Creating service...")
	service, db, err := service.NewService(cfg)
	if err != nil {
		slog.Error("Failed to create service", err)
		panic(err)
	}
	defer db.Close()

	slog.Info("Creating worker...")
	worker.Init(service)

	slog.Info("Init API")
	api.Init(service, cfg.Port)
}

type Config struct {
	Port            string `envconfig:"port" required:"true"`
	OpenAIKey       string `envconfig:"openai_api_key" required:"true"`
	DBAddress       string `envconfig:"db_address" required:"true"`
	DBName          string `envconfig:"db_name" required:"true"`
	DBUser          string `envconfig:"db_user" required:"true"`
	DBPassword      string `envconfig:"db_password" required:"true"`
	DBMigrationPath string `envconfig:"db_migration_path" required:"true"`
}

func (c Config) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.DBUser, c.DBPassword, c.DBAddress, c.DBName)
}

func (c Config) MigrationPath() string {
	return c.DBMigrationPath
}
