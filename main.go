package main

import (
	"ecom_test/cmd/api"
	"ecom_test/db"
	"fmt"
	"log"
)


func main(){
	//start of the api server

	new_database,err:=db.Connect_to_POSTGRES_DB();

	if err!=nil{
		log.Fatal("Error at connecting to the database: ",err);
		return ;
	}
	fmt.Printf("Succesfully contected to the DB : %v\n",new_database);

	new_server:=api.GetServer(":8080",new_database);
	if err:=new_server.Run();err!=nil{
		log.Fatalf("Error at running server: %v",err);
	}

}