package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/module/auth"
	"github.com/vanthang24803/api-ecommerce/internal/module/me"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func RegisterRouter(app *gin.Engine) {
	app.MaxMultipartMemory = 5 << 20

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, util.NotFoundException("Not found route!"))
	})

	api := app.Group("api")

	auth.Router(api)
	me.Router(api)
}
