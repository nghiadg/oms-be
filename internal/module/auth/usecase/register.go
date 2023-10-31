package usecase

import (
	"context"
	"oms-be/internal/common"
	commonmessage "oms-be/internal/message"
	"oms-be/internal/module/auth/entity"
	"oms-be/internal/module/auth/message"
)

type RegisterUseCase struct {
	repo IRegisterRepo
}

func NewRegisterUseCase(repo IRegisterRepo) *RegisterUseCase {
	return &RegisterUseCase{repo: repo}
}

func (uc *RegisterUseCase) Register(ctx context.Context, payload entity.RegisterPayload) (*entity.RegisterRes, *common.AppError) {
	// Check email is registered
	isExist, err := uc.repo.CheckExistEmail(ctx, payload.Email)
	if err != nil {
		return nil, common.NewAppError(err, string(commonmessage.MsgErrCommon))
	}

	if isExist {
		return nil, common.NewAppError(err, string(message.MsgErrEmailIsExisted))
	}

	hashPassword, err := common.HashPassword(payload.LoginPass)

	if err != nil {
		return nil, common.NewAppError(err, string(commonmessage.MsgErrCommon))
	}

	data := entity.RegisterUserCreation{
		Email:     payload.Email,
		Name:      payload.Name,
		Phone:     payload.Phone,
		LoginPass: hashPassword,
		CreateId:  0,
		UpdateId:  0,
	}

	if err := uc.repo.CreateUser(ctx, &data); err != nil {
		return &entity.RegisterRes{Success: false}, common.NewAppError(err, string(commonmessage.MsgErrCommon))
	}

	return &entity.RegisterRes{Success: true}, nil
}
