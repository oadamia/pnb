package service

import (
	"context"
	"pnb/service/store"
)

// HealthCheck check service health
func (s Service) HealthCheck(ctx context.Context) (*store.Health, error) {
	params := store.CreateHealthParams{
		Message: "OK",
		Service: "pnb",
	}

	health, err := s.store.CreateHealth(ctx, params)
	if err != nil {
		return nil, err
	}

	return &health, nil
}
