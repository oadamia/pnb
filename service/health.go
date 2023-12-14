package service

import (
	"context"
	"pnb/service/db"
)

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

	return healthFrom(dbhealth), nil
}

func healthFrom(dbHealth db.Health) *Health {
	return &Health{
		Message:   dbHealth.Message,
		Service:   dbHealth.Service,
		CreatedAt: dbHealth.CreatedAt,
	}
}
