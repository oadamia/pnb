package service

import (
	"context"
)

// HealthCheck check service health
func (s Service) HealthCheck(ctx context.Context) (*Health, error) {
	params := CreateHealthParams{
		Message: "OK",
		Service: "pnb",
	}

	health, err := s.store.CreateHealth(ctx, params)
	if err != nil {
		return nil, err
	}

	return &health, nil
}
