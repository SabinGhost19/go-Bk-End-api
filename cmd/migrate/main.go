package main

import (
	"ecom_test/db"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

//this implmt is a lil messy
//due to the gorm implementation into the main connct to the db
func main() {


	// connect to the data base
	new_database, err := db.Connect_to_POSTGRES_DB()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return
	}

	//get the sql instance
	sqlDB, err := new_database.DB()
	if err != nil {
		log.Fatal("Error getting *sql.DB from GORM: ", err)
		return
	}

	// create the driver
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating postgres driver: ", err)
		return
	}

	// init the migrate
	//using the above driver
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", 
		"postgres",                     
		driver,
	)
	if err != nil {
		log.Fatal("Error initializing migration: ", err)
		return
	}
	cmd:=os.Args[(len(os.Args)-1)];
	if cmd=="up"{
		if err:=m.Up();err!=nil&&err!=migrate.ErrNoChange{
			log.Fatal("Error at migrate UP: ",err);
		}
	}
	if cmd=="down"{
		if err:=m.Down();err!=nil&&err!=migrate.ErrNoChange{
			log.Fatal("Error at migrate DOWN: ",err);
		}
	}
	
	fmt.Println("Database migration completed successfully!")
}