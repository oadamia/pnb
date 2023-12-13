package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

// Start starts service
func Start(e *echo.Echo, port string) {
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
