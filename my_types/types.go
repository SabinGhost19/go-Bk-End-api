package mytypes

import (
	"gorm.io/gorm"
)

type UserStore interface{
	GetUserByEmail(email string)(*User,error)
	GetUserById(id int)(*User,error)
	CreateUser(*User)error
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


type UserType struct{
	FirstName string `json:"firstname" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}