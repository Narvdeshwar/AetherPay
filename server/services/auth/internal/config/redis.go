package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *Config) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	ctx := context.Background()
	if _, err := rds.Ping(ctx).Result(); err != nil {
		log.Fatalf("Redis connection failed:%v", err.Error())
	}
	log.Println("Redis connected successfully")
	return rds
}
