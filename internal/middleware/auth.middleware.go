package middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vanthang24803/api-ecommerce/internal/dto"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")

		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		if len(tokenStr) > 6 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		secretKey := util.GetEnv("JWT_SECRET")

		token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if payloadData, exists := claims["payload"]; exists {
				payloadBytes, err := json.Marshal(payloadData)
				if err != nil {
					log.Printf("Error marshaling payload: %v", err)
					ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
					ctx.Abort()
					return
				}

				var payload dto.Payload
				if err := json.Unmarshal(payloadBytes, &payload); err != nil {
					log.Printf("Error unmarshaling payload to struct: %v", err)
					ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
					ctx.Abort()
					return
				}

				ctx.Set("user", &payload)

			} else {
				log.Println("Payload not found in token")
				ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
				ctx.Abort()
				return
			}
		} else {
			log.Println("Invalid token")
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AuthorizationMiddleware(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		currentUser, ok := user.(dto.Payload)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, util.UnauthorizedException())
			ctx.Abort()
			return
		}

		hasRole := false
		for _, role := range currentUser.Roles {
			for _, requiredRole := range roles {
				if role == requiredRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			log.Println("User does not have the required role")
			ctx.JSON(http.StatusForbidden, util.ForbiddenException())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func GetCurrentUser(ctx *gin.Context) (*dto.Payload, error) {
	userCtx, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("user not found in context")
	}

	currentUser, ok := userCtx.(*dto.Payload)
	if !ok {
		return nil, errors.New("invalid user data")
	}

	return currentUser, nil
}
