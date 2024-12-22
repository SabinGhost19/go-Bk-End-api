package mytypes

import (
	"gorm.io/gorm"
)

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}
type UserStore interface{
	GetUserByEmail(email string)(*User,error)
	GetUserById(id int)(*User,error)
	CreateUser(*User)error
}

type ProductStore interface{
	GetProducts()([]Product,error);
	GetProductByName(name string) (*Product, error)
	CreateProduct(product Product)error
}

type Order struct {
	gorm.Model
	//ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	//CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	gorm.Model
	//ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	//CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	gorm.Model
	//ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"-"` 
	//CreatedAt time.Time `json:"createdat"`
}

type CartCheckoutItem struct{
	ItemID int `json:"itemID"`
	Quantity int `json:"quantity"`
}

type CartCheckoutPayload struct{
	Items []CartCheckoutItem `json:"item",validate:"required"`
}

type RefreshTypePayload struct{
	RefreshToken string `json:"refreshtoken" validate:"required"`
}

type LoginPayloadType struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type UserType struct{
	FirstName string `json:"firstname" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type ProductPayload struct{
	Name string `json:"name"  validate:"required"`
	Description string `json:"description"  validate:"required"`
	Image string `json:"image"  validate:"required"`
	Price float64 `json:"price"  validate:"required"`
	Quantity int   `json:"quantity"  validate:"required"`
}
type Product struct{
	gorm.Model
	//ID	int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Price float64 `json:"price"`
	Quantity int   `json:"quantity"`
	//CreatedAt time.Time `json:"createdAt"`
}
