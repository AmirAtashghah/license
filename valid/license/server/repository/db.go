package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var rdb *redis.Client
var ctx = context.Background()

func PostgreConnection() {
	conn, err := pgx.Connect(context.Background(), "postgres://license:123@localhost:5432/license")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

}

func RedisConnection() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
