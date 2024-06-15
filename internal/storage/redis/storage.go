package redis

import (
	"context"
	"encoding/json"
	"order/internal/model"
	"order/internal/storage"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(redisPath string) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisPath,
		DB:   0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStorage{
		client: client,
	}, nil
}

func (s *RedisStorage) CreateOrder(ctx context.Context, order_uid string, order model.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = s.client.Set(ctx, order_uid, orderJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisStorage) GetOrder(ctx context.Context, order_uid string) (*model.Order, error) {
	orderJSON, err := s.client.Get(ctx, order_uid).Result()
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

func (s *RedisStorage) DeleteOrder(ctx context.Context, order_uid string) error {
	err := s.client.Del(ctx, order_uid).Err()
	if err != nil {
		if err == redis.Nil {
			return storage.ErrOrderNotFound
		}
		return err
	}
	return nil
}
