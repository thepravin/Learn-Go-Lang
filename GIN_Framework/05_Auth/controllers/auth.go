package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) InitAuthControllerRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login", a.Login())
	routes.POST("/register", a.Register())
}

func (a *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login....",
		})
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Register....",
		})
	}
}
