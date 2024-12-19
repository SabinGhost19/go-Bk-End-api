package auth

import (
	"ecom_test/config"
	mytypes "ecom_test/my_types"
	"ecom_test/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
)

func RefreshToken(w http.ResponseWriter,req *http.Request){

	var payload mytypes.RefreshTypePayload;
	err:=utils.ParseJson(req,&payload);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
		return ;
	}
	if err=utils.Validator.Struct(payload);err!=nil{
		errors:=err.(validator.ValidationErrors);
		utils.WriteJsonError(w,http.StatusBadRequest,fmt.Errorf("error at formating: %v",errors));
		return ; 
	}

	token, err := jwt.Parse(payload.RefreshToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Env.JWTSecret), nil
    })

    if err != nil || !token.Valid {
        utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("invalid refresh token"))
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !claims["isRefreshed"].(bool) {
        utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("invalid token type"))
        return
    }

	if int64(claims["expiredAt"].(float64)) < time.Now().Unix() {
        utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("refresh token expired"))
        return
    }

    userId, err := strconv.Atoi(claims["userId"].(string))
    if err != nil {
        utils.WriteJsonError(w, http.StatusInternalServerError, err)
        return
    }

    // Creează un nou Access Token
    accessToken, err := CreateJwt([]byte(config.Env.JWTSecret), userId, false)
    if err != nil {
        utils.WriteJsonError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJson(w, http.StatusOK, map[string]string{
        "accessToken": accessToken,
    })
}
func CreateJwt(secret []byte,user_id int,isRefreshed bool)(string,error){
	
	var expiration_time time.Duration;
	
	if isRefreshed{
		//more time if is refresh token
		
		//expiration_time=time.Second*3;
		
		expiration_time=time.Hour*24*7;
	}else{
		expiration_time=time.Second*time.Duration(config.Env.JWTExpirationTime);
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(user_id),
		"expiredAt": time.Now().Add(expiration_time).Unix(),
		"isRefreshed": isRefreshed,
	});

	tokenString, err := token.SignedString(secret);
	if err!=nil{
		return "",err;
	}

	return tokenString,nil;
}