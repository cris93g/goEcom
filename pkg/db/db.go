package db

import (
	"fmt"
	"os"

	"github.com/cris93g/back/pkg/models"
	"github.com/jinzhu/gorm"
)

//DBConn pulls function from gorm
var DBConn *gorm.DB

//InitDb creates a connection to db
func InitDb(){
	connectionURL:=os.Getenv("CONNECTION_URL")
	var err error
	DBConn, err = gorm.Open("postgres",connectionURL)
	if err!= nil{
		panic("failed to connect to db")
	}
	fmt.Println("db is connected lets go.........")
	DBConn.AutoMigrate(&models.GoItems{})
	fmt.Println("db has been migrated")
}