package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"oms-be/internal/common"
	commonmessage "oms-be/internal/message"
	"oms-be/internal/module/auth/entity"
	"oms-be/internal/module/auth/message"
	"oms-be/internal/module/auth/repo"
	"oms-be/internal/module/auth/usecase"
)

// Register godoc
// @Tags Authentication
// @Produce application/json
// @Param payload body entity.RegisterPayload true "RegisterPayload"
// @Success 200 {object} entity.RegisterRes
// @Router /auth/register [POST]
func register(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var payload entity.RegisterPayload
		if err := ctx.ShouldBind(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, common.SendFailedResponse(nil, string(commonmessage.MsgErrCommon), err))
			return
		}

		// Check confirm password match login password
		if payload.LoginPass != payload.ConfirmLoginPass {
			ctx.JSON(http.StatusBadRequest, common.SendFailedResponse(nil, string(message.MsgErrConfirmLoginPass), nil))
		}

		store := repo.NewRegisterRepo(db)
		business := usecase.NewRegisterUseCase(store)

		data, err := business.Register(ctx, payload)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.SendFailedResponse(nil, err.Message, err))
			return
		}

		ctx.JSON(http.StatusOK, common.SendOkResponse(data))
	}
}
