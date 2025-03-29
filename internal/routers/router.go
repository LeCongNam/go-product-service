package routers

import (
	user_router "product_srv/internal/routers/user_router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(middlewares.Cors())
	// r.Use(middlewares.Logger())

	// Router Here
	user_router.UserRouter(r)

	return r
}
