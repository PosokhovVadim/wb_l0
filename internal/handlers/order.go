package handlers

import (
	"log/slog"
	"order/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandlers struct {
	order service.Order
	log   *slog.Logger
}

func NewOrderHandlers(log *slog.Logger, order service.Order) *OrderHandlers {
	return &OrderHandlers{
		order: order,
		log:   log,
	}
}

func (h *OrderHandlers) CreateOrder(c *fiber.Ctx) error {
	return nil
}

func (h *OrderHandlers) GetOrder(c *fiber.Ctx) error {
	return nil
}
