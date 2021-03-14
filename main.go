package main

import (
	"os"

	"github.com/cris93g/back/pkg/db"
	"github.com/cris93g/back/pkg/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	app:= fiber.New()
	port:=os.Getenv("PORT")
	db.InitDb()
	defer db.DBConn.Close()
	app.Listen("db closed")
	routes.Routes(app)
	app.Listen(":"+port)
}