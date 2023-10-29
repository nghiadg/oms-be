package usecase

import (
	"context"
	commonentity "oms-be/internal/entity"
	"oms-be/internal/module/auth/entity"
)

type ILoginRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*commonentity.UserMst, error)
	CreateUserToken(ctx context.Context, data *commonentity.UserTokenCreation) bool
}

type IRegisterRepo interface {
	CheckExistEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, data *entity.RegisterUserCreation) error
}
