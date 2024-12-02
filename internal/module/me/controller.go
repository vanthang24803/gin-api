package me

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vanthang24803/api-ecommerce/internal/dto"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func GetProfile(ctx *gin.Context) {

	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	result, err := MeRepository().GetProfileHandler(currentUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusOK, util.OK(result))

}

func UpdateProfile(ctx *gin.Context) {

	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	var jsonBody dto.UpdateProfile

	if err := ctx.ShouldBindJSON(&jsonBody); err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	result, err := MeRepository().UpdateProfileHandler(currentUser, &jsonBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusOK, util.OK(result))

}
func UploadAvatar(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException("Upload failed!"))
		return
	}

	uploadDir := "uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.BadRequestException("Failed to create upload directory"))
		return
	}

	originalFilename := file.Filename
	extension := filepath.Ext(originalFilename)

	filePath := fmt.Sprintf("%s/%s%s", uploadDir, uuid.NewString(), extension)

	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.BadRequestException("Failed to save file"))
		return
	}

	result, err := MeRepository().UploadAvatarHandler(currentUser, filePath)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusOK, util.OK(result))
}
