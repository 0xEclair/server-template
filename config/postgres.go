package config

import (
	"fmt"
	"log"
	"os"
	"server-template/cache"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Postgres *gorm.DB

func InitPostgres() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             10 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,             // Don't include params in the SQL log
			Colorful:                  false,            // Disable color
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("postgres"),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: newLogger,
	})

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

	cache.InitCronPostgres(db)
}
