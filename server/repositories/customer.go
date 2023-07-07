package repositories

type Customer struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerRepository interface {
	// Create(Customer) (*Customer, error)
	Signup(request SignupRequest) (Customer, error)
	// Signup(c *fiber.Ctx) error
	// Signin(c *fiber.Ctx) error
	//
}
