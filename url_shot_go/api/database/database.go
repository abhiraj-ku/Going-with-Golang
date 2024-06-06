package database

import (
	"context"
	"os"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) {
	rdb := redis.NewClient(&redis.Options{
		add:      os.Getenv("DB_ADDR"),
		password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	return rdb

}
