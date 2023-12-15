package service

import (
	"database/sql"
	"log/slog"

	"github.com/pressly/goose/v3"
)

type Configurer interface {
	ConnString() string
	MigrationPath() string
}

// Service struct
type Service struct {
	store Querier
}

// New returns new Service
func NewService(c Configurer) (*Service, *sql.DB, error) {
	slog.Info("Connecting to database...")
	db, err := sql.Open("pgx", c.ConnString())
	if err != nil {
		slog.Error("Failed to connect to database", err)
		return nil, nil, err
	}

	slog.Info("Pinging database...")
	err = db.Ping()
	if err != nil {
		slog.Error("Failed to ping database", err)
		return nil, nil, err
	}

	slog.Info("Running migrations...")
	err = goose.Up(db, c.MigrationPath())
	if err != nil {
		slog.Error("Failed to run migrations", err)
		return nil, nil, err
	}

	slog.Info("Creating store...")
	store := New(db)
	return &Service{
		store: store,
	}, db, nil
}
