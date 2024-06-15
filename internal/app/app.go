package app

import (
	"fmt"
	"log/slog"
	"order/internal/handlers"
	"order/internal/service"
	"order/internal/storage/postgresql"
	"order/internal/storage/redis"
	"order/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	log   *slog.Logger
	fiber *fiber.App
	port  int
}

func NewApp(log *slog.Logger, port int, postgresPath string, redisPath string) *App {

	psStorage, err := postgresql.NewPostgresStorage(postgresPath)
	if err != nil {
		log.Error("Failed to init storage:", logger.Err(err))
	}

	redisStorage, err := redis.NewRedisStorage(redisPath)
	if err != nil {
		log.Error("Failed to init storage:", logger.Err(err))
	}

	orderService := service.NewOrderService(log, *psStorage, *redisStorage)
	orderCtrl := handlers.NewOrderHandlers(log, orderService)

	fiberApp := handlers.SetupFiber()

	handlers.SetupRoutes(fiberApp, orderCtrl)

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
