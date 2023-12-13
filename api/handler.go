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
			return ctx.JSON(http.StatusInternalServerError, res)
		}
		return ctx.JSON(http.StatusOK, res)
	}
}
