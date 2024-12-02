package middleware

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

const maxFileSize = 5 * 1024 * 1024

func FileValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.BadRequestException("No file provider"))
			ctx.Abort()
			return
		}

		defer file.Close()

		allowedExtensions := []string{".jpg", ".png", ".webp"}
		fileExt := strings.ToLower(filepath.Ext(header.Filename))
		isValidExtension := false

		for _, ext := range allowedExtensions {
			if fileExt == ext {
				isValidExtension = true
				break
			}
		}

		if !isValidExtension {
			ctx.JSON(http.StatusBadRequest, util.BadRequestException("Invalid file extension: Only .jpg, .png, .webp are allowed"))
			ctx.Abort()
			return
		}

		if header.Size > maxFileSize {
			ctx.JSON(http.StatusBadRequest, util.BadRequestException("File size exceeds 5MB"))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
