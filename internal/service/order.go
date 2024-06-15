package service

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"order/internal/model"
	"order/internal/storage/postgresql"
	"order/internal/storage/redis"

	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

var (
	ErrOrderValidation = errors.New("failed to validate order")
)

type Order interface {
	CreateOrder(ctx context.Context, data []byte) error
	GetOrder(ctx context.Context, uuid uuid.UUID) (*model.Order, error)
	DeleteOrder(ctx context.Context, uuid uuid.UUID) error
}

type OrderService struct {
	log       *slog.Logger
	storage   postgresql.PostgresStorage
	cache     redis.RedisStorage
	validator *validator.Validate
}

func NewOrderService(log *slog.Logger, storage postgresql.PostgresStorage, cache redis.RedisStorage) *OrderService {
	var validate = validator.New()

	return &OrderService{
		log:       log,
		storage:   storage,
		cache:     cache,
		validator: validate,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, data []byte) error {
	var order model.Order

	if err := json.Unmarshal(data, &order); err != nil {
		s.log.Error("Failed to unmarshal order data:", err)
		return err
	}

	if err := s.validator.Struct(order); err != nil {
		s.log.Error("Failed to validate order data:", err)
		return ErrOrderValidation
	}

	if err := s.storage.CreateOrder(ctx, order.OrderUID, order); err != nil {
		s.log.Error("Failed to create order in storage:", err)
		return err
	}

	if err := s.cache.CreateOrder(ctx, order.OrderUID, order); err != nil {
		s.log.Error("Failed to create order in cache:", err)
		return err
	}

	return nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderUID uuid.UUID) (*model.Order, error) {

	order, err := s.cache.GetOrder(ctx, orderUID)
	if err == nil {
		return order, nil
	} else {
		s.log.Error("Failed to get order from cache:", err)
	}

	order, err = s.storage.GetOrder(ctx, orderUID)
	if err != nil {
		s.log.Error("Failed to get order from storage:", err)
		return nil, err
	}

	return order, nil
}

func (s *OrderService) DeleteOrder(ctx context.Context, uuid uuid.UUID) error {
	return nil
}
