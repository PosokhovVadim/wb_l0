package main

import (
	"fmt"
	"order/internal/app"
	"order/internal/config"
	"order/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

func run() error {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("Error loading environment: %v\n", err)
	}

	configPath := os.Getenv("CONFIG_PATH")

	cfg, err := config.New(configPath)

	if err != nil {
		return err
	}

	log, err := logger.SetupLogger(cfg)
	if err != nil {
		return err
	}

	orderApp := app.NewApp(log, cfg.HTTPServer.Port, cfg.StoragePath, cfg.RedisPath)

	if err := orderApp.Run(); err != nil {
		log.Error("Failed to run app:", logger.Err(err))
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error running service: %v\n", err)
		os.Exit(1)
	}
}
