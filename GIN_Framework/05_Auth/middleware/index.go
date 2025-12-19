package middleware

import (
	"ginLearning/05_Auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")

	if headers == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	token := strings.Split(headers, "")

	if len(token) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	_, err := utils.ParseToken(token[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.Next()
}
