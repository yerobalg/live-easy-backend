package sql

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func InitDB() (*DB, error) {
	db, err := initMySQL()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func initMySQL() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
	)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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

	fmt.Println("Successfully connected to database!")

	return db, nil
}

func (db *DB) GetDB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func (db *DB) GetWhereClause(params interface{}) map[string]interface{} {
	whereClause := make(map[string]interface{})
	value := reflect.ValueOf(params)
	for i := 0; i < value.NumField(); i++ {
		paramTag := value.Type().Field(i).Tag.Get("param")
		valueField := value.Field(i).Interface()
		isNullableTag := value.Type().Field(i).Tag.Get("is_nullable") == "true"
		isValueNotNull := db.checkIfNotNull(valueField)
		if paramTag == "" || (!isNullableTag && !isValueNotNull) {
			continue
		}

		whereClause[paramTag] = valueField
	}

	return whereClause
}

func (db *DB) checkIfNotNull(value any) bool {
	switch value.(type) {
	case int64:
		if value.(int64) == int64(0) {
			return false
		}
	case string:
		if value.(string) == "" {
			return false
		}
	case bool:
		if value.(bool) == false {
			return false
		}
	case float64:
		if value.(float64) == float64(0) {
			return false
		}
	}

	return true
}