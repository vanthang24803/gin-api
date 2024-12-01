package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Router(app *gin.RouterGroup) {
	auth := app.Group("/auth")

	auth.GET("", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware([]string{"Customer"}), func(c *gin.Context) {
		user, _ := c.Get("user")

		c.JSON(200, util.OK(user))
	})

	auth.POST("/register", Register)
	auth.POST("/login", Login)
}
