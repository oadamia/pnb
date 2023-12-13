package collector

import (
	"context"
	"pnb/sqldb/store"
)

type Store interface {
	SelectSources(ctx context.Context) ([]store.Source, error)
}

type Collector struct {
	store Store
}

func New(s Store) *Collector {
	return &Collector{
		store: s,
	}
}

func (c *Collector) Start() {
}
