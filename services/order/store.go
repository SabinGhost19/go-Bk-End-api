package order

import (
	mytypes "ecom_test/my_types"

	"gorm.io/gorm"
)


type Store struct{
	db *gorm.DB
}

func NewStore(db *gorm.DB)*Store{
	return &Store{db: db};
}

func (s*Store)CreateOrder(order mytypes.Order) error{
	result:=s.db.Exec("INSERT INTO orders (userId,total,status,address) VALUES (?,?,?,?)",order.UserID,order.Total,order.Status,order.Address)
	if result.Error!=nil{
		return result.Error;
	}
	return nil;
}

func (s*Store)CreateOrderItem(orderItem mytypes.OrderItem) error{
	result := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if result.Error!=nil{
		return result.Error;
	}
	return nil;
}