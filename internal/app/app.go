package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"oms-be/docs"
	v1 "oms-be/internal/module/auth/controller/http/v1"
	"oms-be/pkg/postgres"
)

func Run() {
	//	Repository
	db := postgres.Connect()

	//	HTTP server
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Swagger
	docs.SwaggerInfo.BasePath = "/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.AuthRouter(router, db)

	err := router.Run()

	if err != nil {
		return
	}
}
