package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guatom999/go-auth/repositories"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	accountRepo repositories.CustomerRepository
}

func NewAccountService(accountRepo repositories.CustomerRepository) AccountService {
	return accountService{accountRepo: accountRepo}
}

func (a accountService) Signup(c *fiber.Ctx) (err error) {
	request := SignUpRequest{}

	err = c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user, err := a.accountRepo.Signup(request.Username, password)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
