package service

import (
	"context"
	"pnb/service/db"
	"time"
)

type Health struct {
	Message   string    `json:"message"`
	Service   string    `json:"service"`
	CreatedAt time.Time `json:"created_at"`
}

// HealthCheck check service health
func (s Service) HealthCheck(ctx context.Context) (*Health, error) {
	params := db.InsertHealthParams{
		Message: "OK",
		Service: "pnb",
	}

	dbhealth, err := s.db.InsertHealth(ctx, params)
	if err != nil {
		return nil, err
	}

	return HealthFrom(dbhealth), nil
}

func HealthFrom(dbHealth db.Health) *Health {
	return &Health{
		Message:   dbHealth.Message,
		Service:   dbHealth.Service,
		CreatedAt: dbHealth.CreatedAt,
	}
}
