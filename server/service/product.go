package service

type Product struct {
	ID       int
	Name     string
	Price    int
	Type     int
	Quantity int
}

type ProductServicee interface {
	GetAllProducts() ([]Product, error)
	// Get
}
