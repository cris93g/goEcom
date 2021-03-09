package db

import (
	"fmt"
	"log"
	"os"

	"github.com/cris93g/back/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

//DBConn pulls function from gorm
var DBConn *gorm.DB

//InitDb creates a connection to db
func InitDb(){
	e:=godotenv.Load()
	if e !=nil{
		log.Fatal("Cannot get connection string to DB")
	}
	connectionURL := os.Getenv("CONNECTION_URL")
	var err error
	DBConn, err = gorm.Open("postgres",connectionURL)
	if err!= nil{
		panic("failed to connect to db")
	}
	fmt.Println("db is connected lets go.........")
	DBConn.AutoMigrate(&models.GoItems{})
	fmt.Println("db has been migrated")
}