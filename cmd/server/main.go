package main

import (
	"fmt"
	"log"
	"os"
	"product_srv/internal/database"
	"product_srv/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load biến môi trường từ .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
		log.Fatal("Error loading .env file", err)
		return
	}
	if os.Getenv("GO_ENV") == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.ReleaseMode)

	// 👉 Gán DB vào biến Global để sử dụng ở mọi nơi
	database.NewDatabase()

	r := routers.InitRouter()

	// Khởi động server
	r.Run(":8080")

	log.Println("✅ Application started successfully!")
}
