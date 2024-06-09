package app

import (
	"fmt"
	"log/slog"
	"order/internal/handlers"
	"order/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	log   *slog.Logger
	fiber *fiber.App
	port  int
}

func NewApp(log *slog.Logger, port int) *App {
	// init storage

	fiberApp := handlers.SetupFiber()

	handlers.SetupRoutes(fiberApp, nil)

	return &App{
		log:   log,
		fiber: fiberApp,
		port:  port,
	}
}

func (a *App) Run() error {
	a.log.Info("Starting http server:", slog.Int("port", a.port))

	if err := a.fiber.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		a.log.Error("Failed to run app:", logger.Err(err))
		return err
	}
	return nil
}
