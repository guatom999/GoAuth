package service

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Id       int    `db:"id"`
// 	Username string `db:"username"`
// 	Password string `db:"password"`
// }

const jwtSecret = "BadzBoss"

// type LoginRequest struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

func Signin(c *fiber.Ctx) error {

	db, err := sqlx.Open("mysql", "root:Bossza555@tcp(127.0.0.1:3306)/user")

	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	request := LoginRequest{}
	err = c.BodyParser(&request)

	if err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	user := User{}
	query := "select id, username, password from user where username=?"
	err = db.Get(&user, query, request.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	cliams := jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
	})
}
