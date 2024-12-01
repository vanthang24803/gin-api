package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func AuthRoute(app *gin.RouterGroup) {
	app.GET("", func(c *gin.Context) {
		c.JSON(200, util.OK("Hello World"))
	})
}
