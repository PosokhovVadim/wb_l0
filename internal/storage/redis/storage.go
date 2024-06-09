package redis

import (
	"context"
	"encoding/json"
	"order/internal/model"
	"order/internal/storage"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(redisPath string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisPath,
		DB:   0,
	})
	return &RedisStorage{
		client: client,
	}, nil
}

func (s *RedisStorage) CreateOrder(ctx context.Context, order_uid uuid.UUID, order model.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	key := order_uid.String()
	err = s.client.Set(ctx, key, orderJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) GetOrder(ctx context.Context, uuid uuid.UUID) (*model.Order, error) {
	key := uuid.String()
	orderJSON, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, storage.ErrOrderNotFound
		}
		return nil, err
	}

	var order model.Order
	err = json.Unmarshal([]byte(orderJSON), &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *RedisStorage) DeleteOrder(ctx context.Context, uuid uuid.UUID) error {
	key := uuid.String()
	err := s.client.Del(ctx, key).Err()
	if err != nil {
		if err == redis.Nil {
			return storage.ErrOrderNotFound
		}
		return err
	}
	return nil
}
