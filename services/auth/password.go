package auth

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string)(string, error){
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);
	if err!=nil{
		return "",err;
	}
	return string(hash),nil;
}
func VerifyPassword(password string,hashed_password string)error{
	err:=bcrypt.CompareHashAndPassword([]byte(hashed_password),[]byte(password));
	return err;
}