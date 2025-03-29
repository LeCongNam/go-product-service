package user_router

import (
	diContainer "product_srv/internal/di"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	userController, _ := diContainer.InitUserDomain()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/", userController.CreateUser)

		userRouter.GET("/", userController.GetUser)
	}

	return userRouter
}
