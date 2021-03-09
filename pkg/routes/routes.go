package routes

import (
	"github.com/cris93g/back/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

//Routes to our backend app
func Routes(r *fiber.App) {
	r.Get("/api/items", controllers.GetItems)
	r.Get("/api/v1/item/:id",controllers.GetItemByID)
	r.Get("/api/v1/item/category/:category",controllers.GetAllItemsByCategory)
	r.Get("/api/v1/item/brand/:brand",controllers.GetAllItemsByBrand)
	r.Post("/api/v1/items",controllers.AddItem)
	r.Delete("/api/v1/item/:id",controllers.DeleteItem)
}