package service

import "github.com/gofiber/fiber/v2"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninService interface {
	Signin(c *fiber.Ctx) error
}
