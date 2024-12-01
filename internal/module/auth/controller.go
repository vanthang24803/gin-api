package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Register(ctx *gin.Context) {
	var jsonBody RegisterRequest

	if e := ctx.ShouldBindJSON(&jsonBody); e != nil {
		ctx.JSON(400, util.BadRequestException(e.Error()))
		return
	}

	data, err := AuthRepository().RegisterHandler(&jsonBody)
	if err != nil {
		ctx.JSON(400, util.BadRequestException(err))
		return
	}

	ctx.JSON(201, util.Created(data))
}

func Login(ctx *gin.Context) {
	var jsonBody LoginRequest

	if e := ctx.ShouldBindJSON(&jsonBody); e != nil {
		ctx.JSON(400, util.BadRequestException(e.Error()))
		return
	}

	data, err := AuthRepository().LoginHandler(&jsonBody)
	if err != nil {
		ctx.JSON(400, util.BadRequestException(err))
		return
	}

	ctx.JSON(200, util.OK(data))
}
