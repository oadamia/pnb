package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"pnb/api"
	"pnb/service"
	"pnb/service/store"
	"pnb/worker"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/pressly/goose/v3"
)

func main() {
	slog.Info("Starting up...")

	var cfg Config
	err := envconfig.Process("pnb", &cfg)
	if err != nil {
		slog.Error("Failed to process config", err)
		panic(err)
	}

	slog.Info("Connecting to database...")
	db, err := sql.Open("pgx", cfg.ConnString())
	if err != nil {
		slog.Error("Failed to connect to database", err)
		panic(err)
	}
	defer db.Close()

	slog.Info("Pinging database...")
	err = db.Ping()
	if err != nil {
		slog.Error("Failed to ping database", err)
	}

	slog.Info("Running migrations...")
	err = goose.Up(db, cfg.MigrationPath())
	if err != nil {
		slog.Error("Failed to run migrations", err)
		panic(err)
	}

	querier := store.New(db)
	service := service.NewService(querier)

	slog.Info("Creating worker...")
	err = worker.Init(service)
	if err != nil {
		slog.Error("Failed to create worker", err)
		panic(err)
	}

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
