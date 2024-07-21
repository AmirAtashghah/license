package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/db"
	"time"
)

func CacheNumber(key, value string) error {

	err := db.Rdb.Set(context.Background(), key, value, 30*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("could not cache number: %v", err)
	}

	return nil
}

func GetCachedNumber(key string) (string, error) {

	val, err := db.Rdb.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("could not get cached number: %v", err)
	}

	return val, nil
}
