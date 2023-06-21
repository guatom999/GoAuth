package service

type SignUpRequest struct {
	Username string
	Password string
}

type SignUpService interface {
	Signup() error
}
