package service

import (
	"context"
	"pnb/service/store"
)

func (s Service) GetSource(ctx context.Context, id string) (*store.Source, error) {
	source, err := s.store.GetSource(ctx, id)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) ListSources(ctx context.Context, params store.ListSourceParams) ([]store.Source, error) {
	sources, err := s.store.ListSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func (s Service) CreateSource(ctx context.Context, params store.CreateSourceParams) (*store.Source, error) {
	source, err := s.store.CreateSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) UpdateSource(ctx context.Context, id string, params store.UpdateSourceParams) (*store.Source, error) {
	source, err := s.store.UpdateSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (s Service) DeleteSource(ctx context.Context, id string) (*store.Source, error) {
	source, err := s.store.DeleteSource(ctx, id)
	if err != nil {
		return nil, err
	}

	return &source, nil
}
