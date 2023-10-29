package entity

import "oms-be/internal/entity/common"

type UserMst struct {
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	LoginPass string `json:"login_pass"`
	IsDeleted int8   `json:"is_deleted"`
	common.CUEntity
}

func (UserMst) TableName() string {
	return "user_mst"
}
