package database

import (
	"fmt"
	"log"
	"os"
	"product_srv/internal/domain/users/models"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var (
	dbInstance *Database
	once       sync.Once // Đảm bảo chỉ khởi tạo một lần
)

// ConnectDB - Tạo kết nối database singleton
func ConnectDB() *gorm.DB {
	once.Do(func() { // Chỉ thực hiện một lần, tránh race condition
		dbname := os.Getenv("DB_DATABASE")
		password := os.Getenv("DB_PASSWORD")
		username := os.Getenv("DB_USERNAME")
		port := os.Getenv("DB_PORT")
		host := os.Getenv("DB_HOST")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, dbname)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to database:", err)
			return
		}

		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.SetConnMaxLifetime(0)
			sqlDB.SetMaxIdleConns(50)
			sqlDB.SetMaxOpenConns(50)
		}

		dbInstance = &Database{DB: db}

		// Chạy migration nếu cần
		AutoMigrate(dbInstance.DB)

		fmt.Println("✅ Database connected successfully")
	})

	return dbInstance.DB
}

// AutoMigrate - Chạy migration
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
