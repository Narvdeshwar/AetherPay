package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *Config) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	ctx := context.Background()
	if _, err := rds.Ping(ctx).Result(); err != nil {
		log.Fatalf("Redis connection failed:%v", err.Error())
	}
	log.Println("Redis connected successfully")
	return rds
}
