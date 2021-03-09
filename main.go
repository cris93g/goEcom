package main

import (
	"github.com/cris93g/back/pkg/db"
	"github.com/cris93g/back/pkg/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	app:= fiber.New()
	db.InitDb()
	defer db.DBConn.Close()
	routes.Routes(app)
	app.Listen(":3001")
}