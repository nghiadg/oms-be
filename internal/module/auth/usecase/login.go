package usecase

import (
	"context"
	"oms-be/internal/common"
	commonentity "oms-be/internal/entity"
	"oms-be/internal/module/auth/entity"
	"oms-be/internal/module/auth/message"
	"oms-be/internal/module/auth/utils"
)

type LoginUseCase struct {
	repo ILoginRepo
}

func NewLoginUseCase(repo ILoginRepo) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (uc *LoginUseCase) Login(ctx context.Context, payload entity.LoginPayload) (*entity.LoginRes, *common.AppError) {
	user, err := uc.repo.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		return nil, common.NewAppError(err, string(message.MsgErrUserNotFound))
	}

	if ok := common.CheckHashPassword(payload.LoginPass, user.LoginPass); !ok {
		return nil, common.NewAppError(nil, string(message.MsgErrInvalidPassword))
	}

	accessToken, err := utils.GenerateAccessToken(utils.ClaimsAccessToken{UserId: user.UserId, Email: payload.Email})

	if err != nil {
		return nil, common.NewAppError(err, string(message.MsgErrGenerateAccessToken))
	}

	refreshToken, err := utils.GenerateRefreshToken(utils.ClaimsRefreshToken{UserId: user.UserId})

	if err != nil {
		return nil, common.NewAppError(err, string(message.MsgErrGenerateRefreshToken))
	}

	data := commonentity.UserTokenCreation{
		RefreshToken: refreshToken,
		UserId:       user.UserId,
		CreateId:     user.UserId,
		UpdateId:     user.UserId,
	}

	if ok := uc.repo.CreateUserToken(ctx, &data); !ok {
		return nil, common.NewAppError(err, string(message.MsgErrStoreRefreshToken))
	}

	return &entity.LoginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Email:        user.Email,
			UserId:       user.UserId,
			Phone:        user.Phone,
			Name:         user.Name},
		nil
}
