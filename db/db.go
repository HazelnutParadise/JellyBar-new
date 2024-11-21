package db

import (
	"fmt"
	"jellybar/obj"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DB_HOST     = "192.168.1.101"
	DB_PORT     = "5432"
	DB_NAME     = "database"
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_SSLMODE  = "disable"
	DB_TIMEZONE = "Asia/Taipei"
)

var database *gorm.DB

// ConnectDB 連接 PostgreSQL 資料庫
func ConnectDB() (*gorm.DB, error) {
	// 資料庫連線字串
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT, DB_SSLMODE, DB_TIMEZONE)

	// 配置 GORM 日誌模式
	newLogger := logger.Default.LogMode(logger.Info)

	// 建立資料庫連線
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// 驗證資料庫連線
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 設定資料庫連線池
	sqlDB.SetMaxOpenConns(10)               // 最大連線數
	sqlDB.SetMaxIdleConns(5)                // 最大閒置連線數
	sqlDB.SetConnMaxLifetime(1 * time.Hour) // 最大連線存活時間

	return db, nil
}

func InitDB() {
	// 連接資料庫
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("無法連接資料庫: %v", err)
	}
	log.Println("資料庫連接成功！")

	// 自動建立資料表
	err = db.AutoMigrate(&obj.Article{}, &obj.Category{}, &obj.Author{})
	if err != nil {
		log.Fatalf("資料表初始化失敗: %v", err)
	}
	log.Println("資料表初始化成功！")

	database = db
}
