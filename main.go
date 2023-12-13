package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"pnb/collector"
	"pnb/sqldb/store"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
)

func main() {
	slog.Info("Starting up...")

	var cfg Config
	err := envconfig.Process("pnb", &cfg)
	if err != nil {
		slog.Error("Failed to process config", err)
	}

	slog.Info("Connecting to database...")
	db, err := sql.Open("pgx", cfg.ConnString())
	if err != nil {
		slog.Error("Failed to connect to database", err)
	}
	defer db.Close()

	slog.Info("Pinging database...")
	err = db.Ping()
	if err != nil {
		slog.Error("Failed to ping database", err)
	}

	slog.Info("Running migrations...")
	err = goose.Up(db, cfg.DBMigration)
	if err != nil {
		slog.Error("Failed to run migrations", err)
	}

	slog.Info("Creating store...")
	st := store.New(db)

	slog.Info("Starting collector...")
	collectorService := collector.New(st)
	collectorService.Start()

	e := echo.New()

	Start(e, cfg.Port)
}

type Config struct {
	Port        string `envconfig:"port" required:"true"`
	OpenAIKey   string `envconfig:"openai_api_key" required:"true"`
	DBAddress   string `envconfig:"db_address" required:"true"`
	DBName      string `envconfig:"db_name" required:"true"`
	DBUser      string `envconfig:"db_user" required:"true"`
	DBPassword  string `envconfig:"db_password" required:"true"`
	DBMigration string `envconfig:"db_migration" required:"true"`
}

func (c Config) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.DBUser, c.DBPassword, c.DBAddress, c.DBName)
}
