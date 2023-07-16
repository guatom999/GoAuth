package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	// db *sqlx.DB
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepositoryDB{db}
}

func mockData(db *gorm.DB) error {

	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	fmt.Println("Hello")

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []product{}
	for i := 0; i < 1000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product %v", i+1),
			Price:    random.Intn(9000) + 1000,
			Type:     random.Intn(10),
			Quantity: random.Intn(100),
		})
	}

	return db.Create(&products).Error

}

func (obj productRepositoryDB) GetAll() (product []product, err error) {

	err = obj.db.Limit(30).Find(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (obj productRepositoryDB) GetById(id int) (*product, error) {
	return nil, nil
}
