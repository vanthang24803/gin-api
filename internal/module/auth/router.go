package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Router(app *gin.RouterGroup) {
	auth := app.Group("/auth")

	auth.GET("", func(c *gin.Context) {
		c.JSON(200, util.OK("Hello World"))
	})

	auth.POST("/register", Register)
	auth.POST("/login", Login)
}
