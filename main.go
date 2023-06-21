package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {

	var err error

	db, err = sqlx.Open("mysql", "root:Bossza555@tcp(127.0.0.1:3306)/user")

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	// app.Post("/signup", service.Signup)
	// app.Get("/signin", service.Signup)

	app.Post("/signup", Signup)
	// app.Get("/signin", service.Signup)

	// app.Get("/private", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"success": true, "path": "private"})
	// })

	// app.Get("/public", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"success": true, "path": "public"})
	// })

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

// type User struct {
// 	Id       int    `db:"id"`
// 	Username string `db:"username"`
// 	Password string `db:"password"`
// }

// type SignUpRequest struct {
// 	Username string
// 	Password string
// }

// func Signup(c *fiber.Ctx) error {
// 	request := SignUpRequest{}

// 	// db, err := sqlx.Open("mysql", "root:Bossza555@tcp(127.0.0.1:3306)/user")

// 	// if err != nil {
// 	// 	return fiber.ErrBadGateway
// 	// }

// 	err := c.BodyParser(&request)
// 	if err != nil {
// 		return err
// 	}

// 	if request.Username == "" || request.Password == "" {
// 		return fiber.ErrUnprocessableEntity
// 	}

// 	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	query := "insert into user (username, password) values (? , ?)"

// 	result, err := db.Exec(query, request.Username, string(password))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	user := User{
// 		Id:       int(id),
// 		Username: request.Username,
// 		Password: string(password),
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(user)
// }
