package me

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	ctx.JSON(201, util.OK(result))

}
