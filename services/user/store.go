package user

import (
	mytypes "ecom_test/my_types"
	"errors"
	"fmt"

	"gorm.io/gorm"
)


type Store struct{
	db *gorm.DB
}

func NewStore(db*gorm.DB)*Store{
	return &Store{db: db};
}


func (s*Store)GetUserByEmail(email string)(*mytypes.User,error){
	
	//create the user struct to populate it 
	new_user:=new (mytypes.User);

	err := s.db.Unscoped().Where("email = ?", email).First(&new_user).Error
	if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
        return nil, fmt.Errorf("error occurred during query: %v", err)
    } else {

		fmt.Printf("User found: %+v\n", new_user)
	}
	return new_user,nil;
}

func (s*Store) GetUserById(id int)(*mytypes.User,error){

	new_user:=new (mytypes.User);

	err := s.db.Where("id = ?", id).First(&new_user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil,fmt.Errorf("no user found with the given id, %v", err);
		} else {
			return nil,fmt.Errorf("error occurred during query: %v", err);
		}
	} else {

		fmt.Printf("User found: %+v\n", new_user)
	}
	return new_user,nil;
}

func (s *Store) CreateUser(user *mytypes.User) error {

	result := s.db.Exec(
        "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
        user.FirstName, user.LastName, user.Email, user.Password,
    )
    if result.Error != nil {
        return result.Error
    }

    return nil
}

