package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	db, _ := strconv.ParseUint(os.Getenv("redis_db"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("redis_addr"),
		Password:   os.Getenv("redis_pw"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to redis: %v", err))
	}

	Redis = client
}
