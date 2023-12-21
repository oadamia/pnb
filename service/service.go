package service

import (
	"pnb/service/store"
)

type Configurer interface {
	ConnString() string
	MigrationPath() string
}

// Service struct
type Service struct {
	store store.Querier
}

// New returns new Service

func NewService(db store.Querier) *Service {
	return &Service{
		store: db,
	}
}
