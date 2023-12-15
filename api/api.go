package api

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"pnb/service"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func InitAPI(s service.Service, port string) {
	e = echo.New()
	addLogger()
	registerRoutes(s)
	start(port)
}

func addLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogMethod:   true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
}

func registerRoutes(s service.Service) {
	group := e.Group("/pnb/v1")
	group.GET("/healthcheck", HealthCheck(s))

	group.GET("/sources", ListSources(s))
	group.POST("/sources", CreateSource(s))
	group.GET("/sources/:id", GetSource(s))
	group.PUT("/sources/:id", UpdateSource(s))
	group.DELETE("/sources/:id", DeleteSource(s))

}

func start(port string) {
	if e != nil {
		go func() {
			if err := e.Start(port); err != nil {
				slog.Error("shutting down the server: ", err)
			}
		}()
	} else {
		panic("Start. echo does not exist")
	}
	slog.Info("Started On ", "Port :", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Shutting down server...")
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
