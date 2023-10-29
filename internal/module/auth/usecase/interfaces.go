package usecase

import (
	"context"
	commonentity "oms-be/internal/entity"
)

type ILoginRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*commonentity.UserMst, error)
	CreateUserToken(ctx context.Context, data *commonentity.UserTokenCreation) bool
}
