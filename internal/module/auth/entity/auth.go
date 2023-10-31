package entity

// Login
type LoginRes struct {
	UserId       int    `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginPayload struct {
	Email     string `json:"email" binding:"required"`
	LoginPass string `json:"login_pass" binding:"required"`
}

// Register
type RegisterPayload struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	LoginPass        string `json:"login_pass" binding:"required"`
	ConfirmLoginPass string `json:"confirm_login_pass" binding:"required"`
}

type RegisterRes struct {
	Success bool `json:"success"`
}

type RegisterUserCreation struct {
	Name      string `json:"name" binding:"required" gorm:"column:name"`
	Email     string `json:"email" binding:"required" gorm:"column:email"`
	Phone     string `json:"phone" binding:"required" gorm:"column:phone"`
	LoginPass string `json:"login_pass" binding:"required" gorm:"column:login_pass"`
	CreateId  int    `json:"create_id" binding:"required" gorm:"column:create_id"`
	UpdateId  int    `json:"update_id" binding:"required" gorm:"column:update_id"`
}
