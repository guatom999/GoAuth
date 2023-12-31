package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/guatom999/go-auth/handler"
	"github.com/guatom999/go-auth/repositories"
	"github.com/guatom999/go-auth/service"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	var err error

	// db, err = sqlx.Open("mysql", "root:Bossza555@tcp(127.0.0.1:3306)/user")

	initConfig()

	db := initDB()
	db2 := initData()

	accountRepositoryDB := repositories.NewAccountRepository(db)
	productRepositoryDB := repositories.NewProductRepository(db2)
	// _ = productRepositoryDB
	accountService := service.NewAccountService(accountRepositoryDB)

	productService := service.NewProductService(productRepositoryDB)
	catalogHandler := handler.NewCatalogHandler(productService)
	if err != nil {
		panic(err)
	}

	// router := mux.NewRouter()

	// router.HandleFunc("/singup", accountHandler.NewAccount).Methods(http.MethodPost)

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/signin", accountService.Signin)
	app.Post("/signup", accountService.Signup)
	app.Get("/products", catalogHandler.GetProducts)
	app.Get("/totalproductsprice", catalogHandler.GetTotalProductsPrice)
	// app.Post("/signin", accountRepositoryDB.Signin)
	app.Static("/", "./index", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	err = app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.table"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)

	return db
}

func initData() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db2.username"),
		viper.GetString("db2.password"),
		viper.GetString("db2.host"),
		viper.GetString("db2.port"),
		viper.GetString("db2.database"),
	)

	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return db

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
