package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"oms-be/internal/common"
	"oms-be/internal/message"
	"oms-be/internal/module/auth/entity"
	"oms-be/internal/module/auth/repo"
	"oms-be/internal/module/auth/usecase"
	"time"
)

// Login godoc
// @Tags Authentication
// @Produce application/json
// @Param payload body entity.LoginPayload true "LoginPayload"
// @Success 200 {object} entity.LoginRes
// @Router /auth/login [POST]
func login(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var payload entity.LoginPayload
		if err := ctx.ShouldBind(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest,
				common.SendFailedResponse(nil, string(message.MsgErrCommon), err))

			return
		}

		store := repo.NewLoginRepo(db)
		business := usecase.NewLoginUseCase(store)

		data, err := business.Login(ctx.Request.Context(), payload)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.SendFailedResponse(nil, err.Message, err))
			return
		}

		ctx.SetCookie("access_token", data.AccessToken, int(time.Hour), "", "", true, true)

		ctx.JSON(http.StatusOK, common.SendOkResponse(data))

	}

}
