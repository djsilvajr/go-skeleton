package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Service wraps go-redis with convenience methods.
// Mirrors the RedisService used in the Laravel skeleton.
type Service struct {
	client *goredis.Client
}

func NewService(client *goredis.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return s.client.Set(ctx, key, value, ttl).Err()
}

func (s *Service) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *Service) Del(ctx context.Context, keys ...string) error {
	return s.client.Del(ctx, keys...).Err()
}

func (s *Service) Exists(ctx context.Context, key string) (bool, error) {
	n, err := s.client.Exists(ctx, key).Result()
	return n > 0, err
}

func (s *Service) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return s.client.Expire(ctx, key, ttl).Err()
}
