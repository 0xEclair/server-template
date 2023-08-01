package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func InitPostgres() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("postgres"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to postgres: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to mysql: %v", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	Postgres = db
}
