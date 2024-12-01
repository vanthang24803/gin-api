package me

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
)

func Router(app *gin.RouterGroup) {
	router := app.Group("/me", middleware.AuthenticationMiddleware())

	router.GET("", GetProfile)
	router.PUT("", UpdateProfile)

}
