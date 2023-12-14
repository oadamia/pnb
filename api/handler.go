package api

import (
	"net/http"
	"pnb/service"

	"github.com/labstack/echo/v4"
)

// HealthCheck handler
func HealthCheck(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := s.HealthCheck(ctx.Request().Context())
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, res)
	}
}

// ListSources handler
func ListSources(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		res, err := s.ListSources(ctx.Request().Context())
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, res)
	}
}

// CreateSource handler
func CreateSource(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := new(service.CreateSourceReq)
		err := ctx.Bind(req)
		if err != nil {
			return err
		}

		res, err := s.CreateSource(ctx.Request().Context(), *req)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
