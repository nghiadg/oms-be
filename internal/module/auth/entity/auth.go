package entity

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
