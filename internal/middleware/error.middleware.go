package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			if err.Error() == "not found" {
				c.JSON(http.StatusNotFound, util.NotFoundException("Route not found!"))
				return
			}

			c.JSON(http.StatusInternalServerError, util.InternalServerErrorException())
		}
	}
}
