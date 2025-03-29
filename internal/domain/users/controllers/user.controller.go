package user_controllers

import (
	"net/http"
	"product_srv/internal/domain/users/models"
	"product_srv/internal/domain/users/services"
	"product_srv/internal/repositories"
	"product_srv/pkg/utils"

	"github.com/gin-gonic/gin"

	lo "github.com/samber/lo"
)

type UserController struct {
	svc *services.UserService
}

func NewUserController(svc *services.UserService) *UserController {
	return &UserController{svc: svc}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc.svc.CreateUser(user)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, &user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	limit, err := utils.GetLimit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offset, err := utils.GetOffset(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := lo.PickByKeys(c.Request.URL.Query(), []string{"id", "name", "email", "age"})

	data, count, err := uc.svc.GetListUser(repositories.GetList{
		Limit:   limit,
		Offset:  offset,
		Filters: params,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
		"data":    data,
		"count":   count,
	})
}
