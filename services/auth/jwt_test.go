package auth

import (
	"testing"
)

func TestCreateJWT(t *testing.T) {
	//creating a secret
	secret:=[]byte ("secret");

	//generate the token
	//acces token actually
	token,err:=CreateJwt(secret,1,false);
	if err!=nil{
		t.Errorf("error creating JWT: %v",err);
	}
	if token==""{
		t.Error("expected token to be not empty....");
	}
}