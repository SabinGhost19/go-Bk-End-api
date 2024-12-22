package product

import (
	mytypes "ecom_test/my_types"
	"fmt"

	"gorm.io/gorm"
)

type Store struct{
	db*gorm.DB
}

func NewStore(db*gorm.DB)*Store{
	return &Store{db:db};
}

func (s*Store)CreateProduct(product mytypes.Product)error{
	result := s.db.Exec(
        "INSERT INTO products (name, description, image, price,quantity) VALUES (?, ?, ?, ?, ?)",
        product.Name, product.Description, product.Image, product.Price,product.Quantity,
    )
    if result.Error != nil {
        return result.Error
    }

    return nil
}
func (s*Store)GetProductsByIds(productIDs []int)([]mytypes.Product,error){
	var products []mytypes.Product

	result := s.db.Where("id IN ?", productIDs).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func (s *Store) GetProductByName(name string) (*mytypes.Product, error) {
	var product mytypes.Product

	result := s.db.Where("name = ?", name).First(&product)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product with name %s not found", name)
		}
		return nil, result.Error
	}
	return &product, nil
}
func (s*Store)GetProducts()([]mytypes.Product,error){
	var products []mytypes.Product

	result := s.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
