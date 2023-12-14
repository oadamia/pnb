package service

import (
	"context"
	"pnb/service/db"
)

// Service struct
type Service struct {
	db Database
}

// Database store interface
type Database interface {
	ListSources(ctx context.Context) ([]db.Source, error)
	InsertSource(ctx context.Context, arg db.InsertSourceParams) (db.Source, error)
	InsertHealth(ctx context.Context, arg db.InsertHealthParams) (db.Health, error)
}

// New returns new Service
func New(d Database) Service {
	return Service{
		db: d,
	}
}
