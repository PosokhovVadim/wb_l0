package main

import (
	"fmt"
	"order/internal/app"
	"order/internal/config"
	"order/pkg/logger"
	"os"
	_ "order/docs"
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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := run(); err != nil {
		fmt.Printf("Error running service: %v\n", err)
		os.Exit(1)
	}
}
