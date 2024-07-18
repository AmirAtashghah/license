package repository

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func CacheNumber(key string, value int) error {
	err := rdb.Set(ctx, key, value, 30*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("could not cache number: %v", err)
	}
	return nil
}

func GetCachedNumber(key string) (int, error) {
	val, err := rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return 0, fmt.Errorf("number not found")
	} else if err != nil {
		return 0, fmt.Errorf("could not get cached number: %v", err)
	}

	number, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("could not convert cached value to int: %v", err)
	}

	return number, nil
}
