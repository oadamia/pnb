package api

import (
	"database/sql"
	"errors"
	"net/http"
	"pnb/service/store"

	"github.com/labstack/echo/v4"
)

// ListSource handler
func ListSource(ctx echo.Context) error {

	var req ListSourceReq
	err := ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	listParams := listSourceParamsFrom(req)

	res, err := s.ListSources(ctx.Request().Context(), listParams)
	if err != nil {
		return err
	}

	sources := make([]store.Source, len(res))
	copy(sources, res)

	return ctx.JSON(http.StatusOK, sources)
}

// CreateSource handler
func CreateSource(ctx echo.Context) error {
	req := new(store.CreateSourceParams)
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
	res, err := s.GetSource(ctx.Request().Context(), ctx.Param("id"))
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

	req := new(store.UpdateSourceParams)
	err := ctx.Bind(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := s.UpdateSource(ctx.Request().Context(), ctx.Param("id"), *req)
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
	res, err := s.DeleteSource(ctx.Request().Context(), ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
