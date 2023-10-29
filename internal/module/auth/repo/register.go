package repo

import (
	"context"
	"gorm.io/gorm"
	commonentity "oms-be/internal/entity"
	"oms-be/internal/module/auth/entity"
)

type RegisterRepo struct {
	db *gorm.DB
}

func NewRegisterRepo(db *gorm.DB) *RegisterRepo {
	return &RegisterRepo{db: db}
}

func (repo *RegisterRepo) CreateUser(ctx context.Context, data *entity.RegisterUserCreation) error {
	if err := repo.db.Table(commonentity.UserMst{}.TableName()).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (repo *RegisterRepo) CheckExistEmail(ctx context.Context, email string) (bool, error) {
	var data commonentity.UserMst

	if err := repo.db.Where("email = ?", email).First(&data).Error; err != nil {
		return true, err
	}

	if data.Email != "" {
		return true, nil
	}

	return false, nil

}
