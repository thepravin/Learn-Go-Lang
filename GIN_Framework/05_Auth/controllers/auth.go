package controllers

import (
	"ginLearning/05_Auth/models"
	"ginLearning/05_Auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
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
		var registerBody models.User
		if err := c.BindJSON(&registerBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.authService.RegisterService(&registerBody.Email, &registerBody.Password, &registerBody.Name) // sequence of parameter matters
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
			"user":    user,
		})

	}
}
