package storage

import (
	"context"
	"errors"
	"order/internal/model"

	"github.com/google/uuid"
)

type OrderStorage interface {
	CreateOrder(ctx context.Context, order_uid uuid.UUID, order model.Order) error
	GetOrder(ctx context.Context, uuid uuid.UUID) (*model.Order, error)
	DeleteOrder(ctx context.Context, uuid uuid.UUID) error
}

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
)
