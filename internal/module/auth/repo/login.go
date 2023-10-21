package repo

import (
	"context"
	"fmt"
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
	fmt.Println("handle logic login DB")
	return nil, nil
}
