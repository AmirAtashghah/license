package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"server/logger"
)

const group = "redis"

type Config struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     int    `yaml:"port" env:"PORT" `
	Password string `yaml:"password" env:"PASSWORD" `
	DB       int    `yaml:"db" env:"DB" `
}

type RedisDB struct {
	RedisClient *redis.Client
	config      Config
}

func RedisConnection(cfg Config) *RedisDB {

	log.Println("-------------------------------------------------", redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	Rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", err.Error())

		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisDB{config: cfg, RedisClient: Rdb}
}
