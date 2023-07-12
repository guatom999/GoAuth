package service

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/guatom999/go-auth/repositories"
	"golang.org/x/crypto/bcrypt"
)

// const jwtSecret = "12345"
const jwtSecret = "BadzBoss"

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

// func (a accountService) Signin(c *fiber.Ctx) error {

// 	request := LoginRequest{}
// 	err := c.BodyParser(&request)
// 	if err != nil {
// 		return err
// 	}

// 	if request.Username == "" || request.Password == "" {
// 		return fiber.ErrUnprocessableEntity
// 	}

// 	user := Customer{}
// 	query := "select id, username, password from user where username=?"
// 	err = r.db.Get(&user, query, request.Username)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
// 	}

// 	cliams := jwt.StandardClaims{
// 		Issuer:    strconv.Itoa(user.Id),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
// 	token, err := jwtToken.SignedString([]byte(jwtSecret))
// 	if err != nil {
// 		return fiber.ErrInternalServerError
// 	}

// 	return c.JSON(fiber.Map{
// 		"jwtToken": token,
// 	})
// }

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (a accountService) Signin(c *fiber.Ctx) error {

	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	user, err := a.accountRepo.Signin(request.Username)

	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	claims := &authCustomClaims{
		user.Username,
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(user.Id),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// cliams := jwt.StandardClaims{
	// 	Issuer:    strconv.Itoa(user.Id),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// }

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
	})

	// return c.JSON(user)

	// return nil
}
