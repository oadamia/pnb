package service

import (
	"context"
	"pnb/service/db"
	"time"
)

type CreateSourceReq struct {
	Name   string `json:"name" validate:"required"`
	URL    string `json:"url" validate:"required"`
	Driver string `json:"driver" validate:"required"`
}

type Source struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Driver    string    `json:"driver"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InsertSourceParamsFrom(req CreateSourceReq) db.InsertSourceParams {
	return db.InsertSourceParams{
		Name:   req.Name,
		Url:    req.URL,
		Driver: req.Driver,
	}
}

func SourceFrom(dbSource db.Source) *Source {
	return &Source{
		Id:        int(dbSource.ID),
		Name:      dbSource.Name,
		URL:       dbSource.Url,
		Driver:    dbSource.Driver,
		CreatedAt: dbSource.CreatedAt,
		UpdatedAt: dbSource.UpdatedAt,
	}
}

func (s Service) ListSources(ctx context.Context) ([]Source, error) {
	dbsources, err := s.db.ListSources(ctx)
	if err != nil {
		return nil, err
	}

	sources := make([]Source, 0, len(dbsources))
	for _, dbsource := range dbsources {
		sources = append(sources, *SourceFrom(dbsource))
	}

	return sources, nil
}

func (s Service) CreateSource(ctx context.Context, req CreateSourceReq) (*Source, error) {
	params := InsertSourceParamsFrom(req)

	dbsource, err := s.db.InsertSource(ctx, params)
	if err != nil {
		return nil, err
	}

	return SourceFrom(dbsource), nil
}
