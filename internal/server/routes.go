package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/module/auth"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func RegisterRouter(app *gin.Engine) {
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, util.NotFoundException("Not found route!"))
	})

	api := app.Group("api")

	auth.Router(api)
}
