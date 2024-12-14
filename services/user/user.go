package user

import (
	mytypes "ecom_test/my_types"
	"ecom_test/services/auth"
	"ecom_test/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)


type Handler struct{
	store mytypes.UserStore
}


func GetUserHandler(store mytypes.UserStore)*Handler{
	return &Handler{store:store};
}

func (h*Handler)RegisterRoutes(router*mux.Router)error{
	
	router.HandleFunc("/login",h.LoginHandler).Methods("POST");	
	router.HandleFunc("/register",h.RegisterHandler).Methods("POST");		
	
	return nil;
}

func (h*Handler)LoginHandler(w http.ResponseWriter,req *http.Request){
	log.Println("Request received for Login in the User service");
	
}

func (h*Handler)RegisterHandler(w http.ResponseWriter,req*http.Request){
	
	var payload mytypes.UserType;
	err:=utils.ParseJson(req,&payload);
	if err!=nil{
		utils.WriteJsonError(w,http.StatusBadRequest,err);
		return;
	}

	//validate de stuct after decoding
	if err=utils.Validator.Struct(payload);err!=nil{
		errors:=err.(validator.ValidationErrors);
		utils.WriteJsonError(w,http.StatusBadRequest,fmt.Errorf("error at formating: %v",errors));
		return ; 
	}
	//check if the user exist	
	_,err=h.store.GetUserByEmail(payload.Email);
	if err==nil{
		utils.WriteJsonError(w,http.StatusBadRequest,fmt.Errorf("user with email: %v already exists",payload.Email));
		return ;		
	}

	hashedPassword,err:=auth.HashPassword(payload.Password);
	if err!=nil{
		utils.WriteJson(w,http.StatusInternalServerError,err);
		return ;
	}

	err=h.store.CreateUser(&mytypes.User{FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	});

	if err!=nil{
		utils.WriteJsonError(w,http.StatusInternalServerError,err);
		return ;
	}
	
	utils.WriteJson(w,http.StatusCreated,nil);

}

