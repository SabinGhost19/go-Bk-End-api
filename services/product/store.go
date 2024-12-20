package product

import (
	mytypes "ecom_test/my_types"

	"gorm.io/gorm"
)

type Store struct{
	db*gorm.DB
}

func NewStore(db*gorm.DB)*Store{
	return &Store{db:db};
}

func (s*Store)GetProducts()([]mytypes.Product,error){
	var products []mytypes.Product

	result := s.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
