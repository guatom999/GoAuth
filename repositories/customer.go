package repositories

type Customer struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type CustomerRepository interface {
}
