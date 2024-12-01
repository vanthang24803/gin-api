package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
)

func Router(app *gin.RouterGroup) {
	router := app.Group("/auth")

	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/logout", middleware.AuthenticationMiddleware(), Logout)
}
