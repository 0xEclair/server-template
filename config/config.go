package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf("Load .env error: %v", err))
	}

	if os.Getenv("mysql") != "" {
		InitMysql()
	}

	if os.Getenv("redis_addr") != "" {
		InitRedis()
	}
}
