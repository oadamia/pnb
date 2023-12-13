package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"pnb/api"
	"pnb/service"
	"pnb/service/db"
	"pnb/worker"
	"time"

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
	pool, err := sql.Open("pgx", cfg.ConnString())
	if err != nil {
		slog.Error("Failed to connect to database", err)
	}
	defer pool.Close()

	slog.Info("Pinging database...")
	err = pool.Ping()
	if err != nil {
		slog.Error("Failed to ping database", err)
	}

	slog.Info("Running migrations...")
	err = goose.Up(pool, cfg.DBMigration)
	if err != nil {
		slog.Error("Failed to run migrations", err)
	}

	slog.Info("Creating store...")
	db := db.New(pool)

	slog.Info("Starting collector...")
	service := service.New(db)
	worker := worker.New(service)
	worker.Start()

	e := echo.New()
	api.RegisterRoutes(e, service)

	Start(e, cfg.Port)
}

func Start(e *echo.Echo, port string) {
	if e != nil {
		go func() {
			if err := e.Start(port); err != nil {
				slog.Error("shutting down the server: ", err)
			}
		}()
	} else {
		panic("Start. echo does not exist")
	}
	slog.Info("Started On ", "Port :", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Shutting down server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
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
