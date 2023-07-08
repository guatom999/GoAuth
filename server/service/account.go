package service

import "github.com/gofiber/fiber/v2"

type UserAccount struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AccountService interface {
	Signup(c *fiber.Ctx) error
	// Signin(c *fiber.Ctx) error
}
