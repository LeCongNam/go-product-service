package main

import (
	"fmt"
	"log"
	"os"
	"product_srv/internal/database"
	"product_srv/internal/routers"
	"sync"

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
	database.ConnectDB()

	var wg sync.WaitGroup
	wg.Add(1) // Đánh dấu rằng ta cần chờ 1 goroutine

	r := routers.InitRouter()

	go func() {
		defer wg.Done() // Nếu có lỗi, server dừng lại thì sẽ giải phóng WaitGroup
		if err := r.Run(":8080"); err != nil {
			log.Fatal("❌ Server failed to start:", err)
		}
	}()

	fmt.Println("✅ Server is running on port 8080")

	wg.Wait() // Chờ server chạy

}
