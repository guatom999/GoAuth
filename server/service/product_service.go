package service

import "github.com/guatom999/go-auth/repositories"

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductServicee {
	return productService{productRepo}
}

func (obj productService) GetAllProducts() (products []Product, err error) {

	productsFromDB, err := obj.productRepo.GetAll()

	if err != nil {
		return nil, err

	}

	for _, product := range productsFromDB {
		products = append(products, Product{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Type:     product.Type,
			Quantity: product.Quantity,
		})

	}

	return products, nil
}

func (obj productService) GetTotalProductsPrice() (price float64, err error) {

	productsFromDB, err := obj.productRepo.GetAll()

	if err != nil {
		return 0, err

	}

	// var products []Product

	for _, product := range productsFromDB {
		price = price + float64(product.Price)

		// products = append(products, Product{
		// 	ID:       product.ID,
		// 	Name:     product.Name,
		// 	Price:    product.Price,
		// 	Type:     product.Type,
		// 	Quantity: product.Quantity,
		// })

	}

	return price, nil

}
