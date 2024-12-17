package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


type ConfigPostgres struct{
	Host string
	Port string
	DB_Name string
	User string
	Password string
	SSL_Mode string
}


var Env=initConfig();

func initConfig() ConfigPostgres{

	err:=godotenv.Load();
	if err!=nil{
		log.Fatalln("Error at loading the DOTENV...",err);
	}

	return ConfigPostgres{
		Host: getEnv("HOST","localhost"),
		Port: getEnv("PORT","5431"),
		DB_Name: getEnv("DB_NAME","go_bknd_api"),
		User: getEnv("DB_USER","postgres"),
		Password: getEnv("PASSWORD","155015"),
		SSL_Mode: getEnv("SSL_MODE","disable"),
	};
}

func (c*ConfigPostgres)GetConnectionString()string{
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.DB_Name, c.User, c.Password, c.SSL_Mode,
	)
}

func getEnv(key string,fallback string)string{
	if value,ok:=os.LookupEnv(key);ok{
		return value;
	};
	return fallback;
}