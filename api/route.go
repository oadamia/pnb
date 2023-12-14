package api

import (
	"pnb/service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, s service.Service) {
	group := e.Group("/pnb/v1")
	group.GET("/healthcheck", HealthCheck(s))
	group.GET("/sources", ListSources(s))
	group.POST("/sources", CreateSource(s))
}
