package config

import (
	"fmt"
	"log"
	"os"
	"server-template/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Mysql *gorm.DB

func InitMysql() {
	mysqlLogConfig := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), mysqlLogConfig)
	db, err := gorm.Open(mysql.Open(os.Getenv("mysql")), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to mysql: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to mysql: %v", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	Mysql = db

	migrate()
}

func migrate() {
	_ = Mysql.AutoMigrate(
		&model.Record{},
	)
}
