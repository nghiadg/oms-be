package repo

import (
	"context"
	"gorm.io/gorm"
	commonentity "oms-be/internal/entity"
)

type LoginRepo struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) *LoginRepo {
	return &LoginRepo{db: db}
}

// GetUserByEmail handle logic communicate with DB here
func (repo *LoginRepo) GetUserByEmail(ctx context.Context, email string) (*commonentity.UserMst, error) {
	var user commonentity.UserMst
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *LoginRepo) CreateUserToken(ctx context.Context, data *commonentity.UserTokenCreation) bool {
	err := repo.db.Table(commonentity.UserToken{}.TableName()).Create(&data).Error

	return err == nil
}
