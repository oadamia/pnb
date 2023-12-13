package main

import (
	"database/sql"
	"fmt"
	"log/slog"

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
		fmt.Println(err)
	}

	slog.Info("Connecting to database...")
	db, err := sql.Open("pgx", cfg.ConnString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	slog.Info("Pinging database...")
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	slog.Info("Running migrations...")
	err = goose.Up(db, cfg.DBMigration)
	if err != nil {
		fmt.Println(err)
	}

	slog.Info("Started up!")
}

type Config struct {
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
