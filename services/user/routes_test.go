package user

import (
	"bytes"
	mytypes "ecom_test/my_types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)


func TestUserServiceHandler(t *testing.T){
	
	user_store:=&mockUserStore{};
	handler:=GetUserHandler(user_store);

	//register user payload
	payload:=mytypes.UserType{FirstName: "firstname",
			LastName:"lastname",
			Email: "sabin@gmail.com",
			Password: "as2",
	};
	marshaled,_:=json.Marshal(payload);
	t.Run("show fail if the user payload is invalid ",func(t *testing.T){
		req,err:=http.NewRequest(http.MethodPost,"/register",bytes.NewBuffer(marshaled));
		
		if err!=nil{
			log.Fatalf("Test Failed...%v\n",err.Error());
		}

		rr:=httptest.NewRecorder();
		router:=mux.NewRouter();

		router.HandleFunc("/register",handler.RegisterHandler);
		router.ServeHTTP(rr,req);

		if rr.Code==http.StatusCreated{
			t.Errorf("expected Status code: %d, got %d",http.StatusCreated,rr.Code);
		};
		if rr.Code==http.StatusInternalServerError{
			t.Errorf("expected Status code: %d, got %d",http.StatusInternalServerError,rr.Code);
		};
		if rr.Code==http.StatusBadRequest{
			t.Errorf("expected Status code: %d, got %d",http.StatusBadRequest,rr.Code);
		};

		});

}

type mockUserStore struct{

}

func (m*mockUserStore)GetUserByEmail(email string)(*mytypes.User,error){
	return nil,fmt.Errorf("user not found in the db : %v",email);
	//return nil,nil;

}
func (m*mockUserStore)GetUserById(id int)(*mytypes.User,error){
	return nil,nil;
}
func (m*mockUserStore)CreateUser(*mytypes.User)error{
	return nil;
}