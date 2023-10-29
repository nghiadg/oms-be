package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(handler *gin.Engine, db *gorm.DB) {
	//	setup common middleware here
	h := handler.Group("/v1")
	{
		r := h.Group("/auth")
		{
			r.POST("/login", login(db))
		}
	}
}
