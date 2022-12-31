package sql

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"live-easy-backend/sdk/log"
)

type DB struct {
	*gorm.DB
}

func Init(serverLogger *log.Logger) (*DB, error) {
	db, err := initMySQL(serverLogger)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func initMySQL(serverLogger *log.Logger) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
	)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: log.New(serverLogger),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db, nil
}
