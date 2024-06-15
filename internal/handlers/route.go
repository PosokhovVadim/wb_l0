package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	swagger "github.com/gofiber/swagger"
)

func SetupFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	return app
}

func SetupRoutes(app *fiber.App, orderCtrl *OrderHandlers) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("api/v1/order:uuid", orderCtrl.GetOrder)
	app.Post("api/v1/order", orderCtrl.CreateOrder)
}
