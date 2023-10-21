package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"oms-be/internal/module/auth/entity"
	"oms-be/internal/module/auth/repo"
	"oms-be/internal/module/auth/usecase"
)

func newRoutes(handler *gin.RouterGroup, db *gorm.DB) {
	route := handler.Group("/auth")
	{
		route.POST("/login", Login(db))
	}
}

func Login(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var payload entity.LoginPayload
		if err := ctx.ShouldBind(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := repo.NewLoginRepo(db)
		business := usecase.NewLoginUseCase(store)

		if err := business.Login(ctx.Request.Context(), payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": "OK"})

	}

}
