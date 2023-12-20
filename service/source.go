package service

import (
	"context"
)

func (s Service) GetSource(ctx context.Context, id int) (*Source, error) {
	source, err := s.store.GetSource(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) ListSources(ctx context.Context) ([]Source, error) {
	sources, err := s.store.ListSources(ctx)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func (s Service) CreateSource(ctx context.Context, params CreateSourceParams) (*Source, error) {
	source, err := s.store.CreateSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) UpdateSource(ctx context.Context, id int, params UpdateSourceParams) (*Source, error) {
	source, err := s.store.UpdateSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) DeleteSource(ctx context.Context, id int) (*Source, error) {
	source, err := s.store.DeleteSource(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &source, nil
}
