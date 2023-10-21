package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(handler *gin.Engine, db *gorm.DB) {
	//	setup common middleware here
	h := handler.Group("/v1")
	{
		newRoutes(h, db)
	}
}
