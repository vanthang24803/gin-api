package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		if len(token) > 6 && token[:7] == "Bearer " {
			token = token[7:]
		}

		secretKey := util.GetEnv("JWT_SECRET")

		claims := &jwt.RegisteredClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
