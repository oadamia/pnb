package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {
	res, err := s.HealthCheck(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
