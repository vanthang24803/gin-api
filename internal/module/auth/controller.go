package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthang24803/api-ecommerce/internal/models"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Register(ctx *gin.Context) {
	var jsonBody models.Register

	if e := ctx.ShouldBindJSON(&jsonBody); e != nil {
		ctx.JSON(400, util.BadRequestException(e.Error()))
		return
	}

	data, err := AuthRepository().RegisterHandler(&jsonBody)
	if err != nil {
		ctx.JSON(400, util.BadRequestException(err.Error()))
		return
	}

	ctx.JSON(201, util.Created(data))
}

func Login(ctx *gin.Context) {
	var jsonBody models.Login

	if e := ctx.ShouldBindJSON(&jsonBody); e != nil {
		ctx.JSON(400, util.BadRequestException(e.Error()))
		return
	}

	data, err := AuthRepository().LoginHandler(&jsonBody)
	if err != nil {
		ctx.JSON(400, util.BadRequestException(err.Error()))
		return
	}

	ctx.JSON(200, util.OK(data))
}
