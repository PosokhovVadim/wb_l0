package service

import "log/slog"

type Order interface {
	CreateOrder() error
	GetOrder(id int64) error
}

type OrderService struct {
	log *slog.Logger
	// storage
}

func NewOrderService(log *slog.Logger) *OrderService {
	return &OrderService{
		log: log,
	}
}

func (s *OrderService) CreateOrder() error {
	return nil
}

func (s *OrderService) GetOrder(id int64) error {
	return nil
}
