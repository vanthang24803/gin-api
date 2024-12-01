package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/vanthang24803/api-ecommerce/internal/config"
	_ "github.com/vanthang24803/api-ecommerce/internal/database"
	"github.com/vanthang24803/api-ecommerce/internal/middleware"
	"github.com/vanthang24803/api-ecommerce/internal/util"
)

func Application() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	PORT := util.GetEnv("PORT")

	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.ErrorHandlingMiddleware())

	RegisterRouter(app)

	log.Printf("Application listing on port %s ✔️", PORT)

	if err := app.Run(fmt.Sprintf(":%s", PORT)); err != nil {
		log.Panicf("Server error: %v", err)
	}

}
