package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client
var Conn *pgx.Conn

func PostgreConnection() {

	var err error

	Conn, err = pgx.Connect(context.Background(), "postgres://root:123@localhost:5432/license")
	if err != nil {
		log.Panic(err)
	}
	//defer Conn.Close(context.Background())

}

func RedisConnection() {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
