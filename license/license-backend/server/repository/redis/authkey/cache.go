package authkey

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/logger"
	rdb "server/repository/redis"
	"time"
)

const group = "authkey"

type DB struct {
	conn *rdb.RedisDB
}

func New(conn *rdb.RedisDB) *DB {
	return &DB{
		conn: conn,
	}
}

func (d *DB) CacheRequestNumber(key, value string) error {

	err := d.conn.RedisClient.Set(context.Background(), key, value, 30*time.Minute).Err()
	if err != nil {
		logger.L().WithGroup(group).Error("error", "error", fmt.Errorf("could not cache number: %v", err).Error())

		return fmt.Errorf("could not cache number: %v", err)
	}

	return nil
}

func (d *DB) GetRequestNumber(key string) (string, error) {

	val, err := d.conn.RedisClient.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		logger.L().WithGroup(group).Error("error", "error", fmt.Errorf("could not get cached number: %v", err).Error())

		return "", fmt.Errorf("could not get cached number: %v", err)
	}

	return val, nil
}
