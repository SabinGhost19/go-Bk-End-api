package api

import (
	"ecom_test/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


type Server struct{
	_addr string
	database *gorm.DB
}

func GetServer(addr string,database *gorm.DB)*Server{
	return &Server{_addr: addr,database: database};
}

func (s*Server)Run()error{

	router:=mux.NewRouter();
	subrouter:=router.PathPrefix("/api/v1").Subrouter();

	//create new user store
	//pass the database to the store 
	//to know where to operate
	new_user_store:= user.NewStore(s.database);
	//set the new user store to the user package
	//and get the userHandler to register new routes 
	//for user comm>>
	userHandler:=user.GetUserHandler(new_user_store);
	userHandler.RegisterRoutes(subrouter);

	log.Printf("Listen on port %v ...",s._addr);
	return http.ListenAndServe(s._addr,router);
}