package service

import (
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
