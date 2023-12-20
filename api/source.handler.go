package api

import (
	"database/sql"
	"errors"
	"net/http"
	"pnb/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ListSources handler
func ListSources(ctx echo.Context) error {
	res, err := s.ListSources(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

// CreateSource handler
func CreateSource(ctx echo.Context) error {
	req := new(service.CreateSourceParams)
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := s.CreateSource(ctx.Request().Context(), *req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

// GetSource handler
func GetSource(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res, err := s.GetSource(ctx.Request().Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// UpdateSource handler
func UpdateSource(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := new(service.UpdateSourceParams)
	err = ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := s.UpdateSource(ctx.Request().Context(), id, *req)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

// DeleteSource handler
func DeleteSource(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res, err := s.DeleteSource(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
