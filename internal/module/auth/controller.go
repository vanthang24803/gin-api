package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/dto"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Register(ctx *gin.Context) {
	var jsonBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&jsonBody); err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	data, err := AuthRepository().RegisterHandler(&jsonBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusCreated, util.Created(data))
}

func Login(ctx *gin.Context) {
	var jsonBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&jsonBody); err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	data, err := AuthRepository().LoginHandler(&jsonBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusOK, util.OK(data))
}

func Logout(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	data, err := AuthRepository().LogoutHandler(currentUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.BadRequestException(err))
		return
	}

	ctx.JSON(http.StatusOK, util.OK(data))
}
