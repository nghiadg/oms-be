package app

import (
	"github.com/gin-gonic/gin"
	v1 "oms-be/internal/module/auth/controller/http/v1"
	"oms-be/pkg/postgres"
)

func Run() {
	//	Repository
	db := postgres.Connect()

	//	HTTP server
	handler := gin.Default()

	v1.NewRouter(handler, db)

	err := handler.Run()

	if err != nil {
		return
	}
}
