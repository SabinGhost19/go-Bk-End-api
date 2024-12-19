package auth

import (
	"ecom_test/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwt(secret []byte,user_id int)(string,error){
	
	expiration_time:=time.Second*time.Duration(config.Env.JWTExpirationTime);

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(user_id),
		"expiredAt": time.Now().Add(expiration_time).Unix(),
	});
		
	tokenString, err := token.SignedString(secret);
	if err!=nil{
		return "",err;
	}

	return tokenString,nil;
}