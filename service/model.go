package service

import "time"

type Health struct {
	Message   string    `json:"message"`
	Service   string    `json:"service"`
	CreatedAt time.Time `json:"created_at"`
}

type Source struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Driver    string    `json:"driver"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSourceReq struct {
	Name   string `json:"name" validate:"required"`
	URL    string `json:"url" validate:"required"`
	Driver string `json:"driver" validate:"required"`
}

type UpdateSourceReq struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Driver    string    `json:"driver"`
	CreatedAt time.Time `json:"created_at"`
}

type DeleteSourceReq struct {
	Id int `json:"id" validate:"required"`
}
