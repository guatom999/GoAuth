package repositories

type product struct {
	ID       int
	Name     string
	Price    int
	Type     int
	Quantity int
}

type ProductRepository interface {
	GetAll() ([]product, error)
	GetById(id int) (*product, error)
}
