package redis

import (
	"context"
	"fmt"
	"log"

	goredis "github.com/redis/go-redis/v9"
	"github.com/djsilvajr/go-skeleton/internal/config"
)

func Connect(cfg *config.Config) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Printf("⚠️  Redis connection warning: %v", err)
	} else {
		log.Println("✅ Redis connected")
	}

	return rdb
}
