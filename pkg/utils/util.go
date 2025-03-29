package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLimit(c *gin.Context) (int, error) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // default limit
	}

	return limit, nil
}

func GetOffset(c *gin.Context) (int, error) {
	offsetStr := c.Query("offset")
	limit, err := strconv.Atoi(offsetStr)
	if err != nil {
		limit = 0 // default limit
	}

	return limit, nil
}

func GetInt(c *gin.Context, key string) (int, error) {
	limitStr := c.Query(key)
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatal(err)
	}

	return limit, nil
}
