package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Load .env error: %v\n", err)
	}

	if os.Getenv("mysql") != "" {
		InitMysql()
	}

	if os.Getenv("postgres") != "" {
		InitPostgres()
	}

	if os.Getenv("redis_addr") != "" {
		InitRedis()
	}
}
