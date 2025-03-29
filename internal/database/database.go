package database

import (
	"fmt"
	"log"
	"os"
	"product_srv/internal/domain/users/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db       *gorm.DB
	host     string
	port     string
	username string
	password string
	dbname   string
}

var DB *gorm.DB

// NewDatabase - Khởi tạo newDB
func NewDatabase() (*gorm.DB, error) {
	newDB := &Database{}

	newDB.dbname = os.Getenv("DB_DATABASE")
	newDB.password = os.Getenv("DB_PASSWORD")
	newDB.username = os.Getenv("DB_USERNAME")
	newDB.port = os.Getenv("DB_PORT")
	newDB.host = os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		newDB.username, newDB.password, newDB.host, newDB.port, newDB.dbname)

	// DON'T DO THIS IN PRODUCTION
	// fmt.Println("DSN:", dsn)

	var err error
	newDB.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
		return nil, err
	}

	DB = newDB.db

	// Chạy migration nếu cần
	AutoMigrate(DB)

	fmt.Println("✅ Database connected successfully")
	return DB, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
