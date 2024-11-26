package db

import (
	"fmt"
	"jellybar/obj"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

var devConfig = DBConfig{
	DB_HOST:     "192.168.1.164",
	DB_PORT:     "5432",
	DB_NAME:     "jellydev",
	DB_USER:     "user",
	DB_PASSWORD: "password",
	DB_SSLMODE:  "disable",
	DB_TIMEZONE: "Asia/Taipei",
}

var prodConfig = DBConfig{
	DB_HOST:     os.Getenv("DB_HOST"),
	DB_PORT:     os.Getenv("DB_PORT"),
	DB_NAME:     os.Getenv("DB_NAME"),
	DB_USER:     os.Getenv("DB_USER"),
	DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
	DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
}

var database *gorm.DB

// ConnectDB 連接 PostgreSQL 資料庫
func ConnectDB(mode int) (*gorm.DB, error) {
	var config DBConfig
	switch mode {
	case DEV:
		config = devConfig
	case PROD:
		config = prodConfig
		// 檢查環境變數是否設定
		if config.DB_HOST == "" || config.DB_PORT == "" || config.DB_NAME == "" || config.DB_USER == "" || config.DB_PASSWORD == "" || config.DB_SSLMODE == "" || config.DB_TIMEZONE == "" {
			return nil, fmt.Errorf("未設定環境變數")
		}
	default:
		return nil, fmt.Errorf("無效的資料庫模式")
	}
	// 資料庫連線字串
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s", config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_HOST, config.DB_PORT, config.DB_SSLMODE, config.DB_TIMEZONE)

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

const (
	DEV = iota
	PROD
)

func InitDB(mode int) {
	// 連接資料庫
	var db *gorm.DB
	var err error
	for {
		db, err = ConnectDB(mode)
		if err != nil {
			log.Printf("無法連接資料庫: %v，將繼續重試...", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("資料庫連接成功！")
		break
	}

	// 自動建立資料表
	for {
		err = db.AutoMigrate(&obj.Article{}, &obj.Category{}, &obj.Author{}, &obj.User{})
		if err != nil {
			log.Printf("資料表初始化失敗: %v，將繼續重試...", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("資料表初始化成功！")
		break
	}

	database = db
}

func IsDBConnected() bool {
	if database == nil {
		return false
	}

	// 額外檢查數據庫連接是否有效
	sqlDB, err := database.DB()
	if err != nil {
		return false
	}

	err = sqlDB.Ping()
	return err == nil
}
