package entity

import "oms-be/internal/entity/common"

type UserToken struct {
	TokenId      int    `json:"token_id" gorm:"column:token_id"`
	UserId       int    `json:"user_id" gorm:"column:user_id"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	common.CUEntity
}

func (UserToken) TableName() string {
	return "user_token"
}

type UserTokenCreation struct {
	UserId       int    `json:"user_id" gorm:"column:user_id"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	CreateId     int    `json:"create_id" gorm:"column:create_id"`
	UpdateId     int    `json:"update_id" gorm:"column:update_id"`
}
