package api

import (
	"net/http"
	"pnb/service"
	"strconv"

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
		res, err := s.SelectSources(ctx.Request().Context())
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

// GetSource handler
func GetSource(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		res, err := s.GetSource(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, res)
	}
}

// UpdateSource handler
func UpdateSource(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}

		req := new(service.UpdateSourceReq)
		err = ctx.Bind(req)
		if err != nil {
			return err
		}

		res, err := s.UpdateSource(ctx.Request().Context(), id, *req)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, res)
	}
}

// DeleteSource handler
func DeleteSource(s service.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return err
		}
		res, err := s.DeleteSource(ctx.Request().Context(), id)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, res)
	}
}
