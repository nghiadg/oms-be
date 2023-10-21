package entity

type LoginRes struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type LoginPayload struct {
	Email     string `json:"email"`
	LoginPass string `json:"login_pass"`
}
