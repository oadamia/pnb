package service

import (
	"context"
	"pnb/service/db"
)

func (s Service) GetSource(ctx context.Context, id int) (*Source, error) {
	dbsource, err := s.db.GetSource(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return sourceFrom(dbsource), nil
}

func (s Service) SelectSources(ctx context.Context) ([]Source, error) {
	dbsources, err := s.db.SelectSources(ctx)
	if err != nil {
		return nil, err
	}

	sources := make([]Source, 0, len(dbsources))
	for _, dbsource := range dbsources {
		sources = append(sources, *sourceFrom(dbsource))
	}

	return sources, nil
}

func (s Service) CreateSource(ctx context.Context, req CreateSourceReq) (*Source, error) {
	params := insertSourceParamsFrom(req)

	dbsource, err := s.db.InsertSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return sourceFrom(dbsource), nil
}

func (s Service) UpdateSource(ctx context.Context, id int, req UpdateSourceReq) (*Source, error) {
	params := updateSourceParamsFrom(id, req)

	dbsource, err := s.db.UpdateSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return sourceFrom(dbsource), nil
}

func (s Service) DeleteSource(ctx context.Context, id int) (*Source, error) {
	dbsource, err := s.db.DeleteSource(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return sourceFrom(dbsource), nil
}

func insertSourceParamsFrom(req CreateSourceReq) db.InsertSourceParams {
	return db.InsertSourceParams{
		Name:   req.Name,
		Url:    req.URL,
		Driver: req.Driver,
	}
}

func sourceFrom(dbSource db.Source) *Source {
	return &Source{
		Id:        int(dbSource.ID),
		Name:      dbSource.Name,
		URL:       dbSource.Url,
		Driver:    dbSource.Driver,
		CreatedAt: dbSource.CreatedAt,
		UpdatedAt: dbSource.UpdatedAt,
	}
}

func updateSourceParamsFrom(id int, req UpdateSourceReq) db.UpdateSourceParams {
	return db.UpdateSourceParams{
		ID:     int32(id),
		Name:   req.Name,
		Url:    req.URL,
		Driver: req.Driver,
	}
}
