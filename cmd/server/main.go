package main

import (
	"log"
	"product_srv/internal/database"
	"product_srv/internal/routers"

	"github.com/joho/godotenv"
)

func main() {
	// Load biến môi trường từ .env
	err := godotenv.Load(".env.dev", ".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// 👉 Gán DB vào biến Global để sử dụng ở mọi nơi
	database.NewDatabase()

	r := routers.InitRouter()

	// Khởi động server
	r.Run(":8080")

	log.Println("✅ Application started successfully!")
}
