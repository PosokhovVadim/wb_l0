package handlers

import (
	"context"
	"errors"
	"log/slog"
	"order/internal/service"
	"order/internal/storage"

	_ "order/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the given details
// @Tags order
// @Accept json
// @Produce json
// @Param order body interface{} true "Order Request"
// @Success 201 {string} string "Created"
// @Failure 400 {object} map[string]interface{} "Invalid request body or invalid order data"
// @Failure 409 {object} map[string]interface{} "Order already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /order [post]
func (h *OrderHandlers) CreateOrder(c *fiber.Ctx) error {
	body := c.Body()

	if len(body) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.order.CreateOrder(context.Background(), body)

	if err != nil {

		if errors.Is(err, service.ErrOrderValidation) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid order data",
			})
		}

		if errors.Is(err, storage.ErrOrderAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "order already exists",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.SendStatus(fiber.StatusCreated)

}

// GetOrder godoc
// @Summary Get an order by UUID
// @Description Get details of an order by its UUID
// @Tags order
// @Accept json
// @Produce json
// @Param uuid path string true "Order UUID"
// @Success 200 {object} model.Order
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/{uuid} [get]
func (h *OrderHandlers) GetOrder(c *fiber.Ctx) error {
	orderUID, err := uuid.Parse(c.Params("uuid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid UUID",
		})
	}

	order, err := h.order.GetOrder(context.Background(), orderUID)
	if err != nil {
		if errors.Is(err, storage.ErrOrderNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "order not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(order)
}
