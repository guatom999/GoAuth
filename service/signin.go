package service

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninService interface {
	Signin() error
}
