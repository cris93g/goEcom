package db

import (
	"fmt"

	"github.com/cris93g/back/pkg/models"
	"github.com/jinzhu/gorm"
)

//DBConn pulls function from gorm
var DBConn *gorm.DB

//InitDb creates a connection to db
func InitDb(){

	var err error
	DBConn, err = gorm.Open("postgres","host=ec2-54-235-252-23.compute-1.amazonaws.com port=5432  user=zxfwdderycxcqn dbname=d5f9fcljfn7v2i password=4f697ab683da4c3eb158c85db737abeffe4acbec775cbd05dca56aae6756081bCONNECTION_URL=host=ec2-54-235-252-23.compute-1.amazonaws.com port=5432  user=zxfwdderycxcqn dbname=d5f9fcljfn7v2i password=4f697ab683da4c3eb158c85db737abeffe4acbec775cbd05dca56aae6756081b")
	if err!= nil{
		panic("failed to connect to db")
	}
	fmt.Println("db is connected lets go.........")
	DBConn.AutoMigrate(&models.GoItems{})
	fmt.Println("db has been migrated")
}