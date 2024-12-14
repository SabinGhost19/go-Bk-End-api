package db

import (
	"ecom_test/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//get the packages:

func Connect_to_POSTGRES_DB()(*gorm.DB,error){
	connection_stirng:=config.Env.GetConnectionString();
	return gorm.Open(postgres.Open(connection_stirng),&gorm.Config{});
}