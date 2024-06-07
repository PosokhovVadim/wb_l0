package main

import (
	"fmt"
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

	_ = cfg

	log, err := logger.SetupLogger(cfg)
	if err != nil {
		return err
	}

	_ = log

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error running service: %v\n", err)
		os.Exit(1)
	}
}
