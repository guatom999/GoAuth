package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// const jwtSecret = "12345"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

// func (r customerRepositoryDB) Create(Customer) (*Customer, error) {
// 	return
// }

func (r customerRepositoryDB) Signup(username string, password []byte) (user Customer, err error) {

	query := "insert user (username, password) values (?, ?)"
	// result, err := r.db.Exec(query, request.Username, string(password))
	result, err := r.db.Exec(query, username, string(password))

	if err != nil {
		return user, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user = Customer{
		Id:       int(id),
		Username: username,
		Password: string(password),
	}

	return user, nil

}

// func (r customerRepositoryDB) Signup2(c *fiber.Ctx) error {
// 	request := SignupRequest{}
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

// 	query := "insert user (username, password) values (?, ?)"
// 	result, err := r.db.Exec(query, request.Username, string(password))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
// 	}

// 	user := Customer{
// 		Id:       int(id),
// 		Username: request.Username,
// 		Password: string(password),
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(user)
// }

// *********************|| SignIn ||*********************|| //

// func (r customerRepositoryDB) Signin(c *fiber.Ctx) error {
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
