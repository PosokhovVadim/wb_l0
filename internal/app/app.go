package app

import (
	"context"
	"fmt"
	"log/slog"
	"order/internal/handlers"
	"order/internal/service"
	"order/internal/storage/postgresql"
	"order/internal/storage/redis"
	nats "order/internal/sub"
	"order/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	log     *slog.Logger
	fiber   *fiber.App
	port    int
	service service.Order
}

func NewApp(log *slog.Logger, port int, postgresPath string, redisPath string, natsURL string) (*App, error) {

	psStorage, err := postgresql.NewPostgresStorage(postgresPath)
	if err != nil {
		log.Error("Failed to init storage:", logger.Err(err))
		return nil, err
	}

	redisStorage, err := redis.NewRedisStorage(redisPath)
	if err != nil {
		log.Error("Failed to init storage:", logger.Err(err))
		return nil, err
	}

	orderService := service.NewOrderService(log, *psStorage, *redisStorage)

	nats := nats.NewNatsSub(log, orderService)

	go func() {
		_ = nats.NatsConnect(natsURL)
	}()

	orderCtrl := handlers.NewOrderHandlers(log, orderService)

	fiberApp := handlers.SetupFiber()

	handlers.SetupRoutes(fiberApp, orderCtrl)

	return &App{
		log:     log,
		fiber:   fiberApp,
		port:    port,
		service: orderService,
	}, nil
}

func (a *App) SyncData(ctx context.Context) error {
	if err := a.service.SyncData(ctx); err != nil {
		a.log.Error("Failed to sync data:", logger.Err(err))
		return err
	}
	return nil
}

func (a *App) Run() error {
	a.log.Info("Starting http server:", slog.Int("port", a.port))

	if err := a.fiber.Listen(fmt.Sprintf(":%d", a.port)); err != nil {
		a.log.Error("Failed to run app:", logger.Err(err))
		return err
	}
	return nil
}
