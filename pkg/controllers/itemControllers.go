package controllers

import (
	"github.com/cris93g/back/pkg/db"
	"github.com/cris93g/back/pkg/models"
	"github.com/gofiber/fiber/v2"
)



func GetItems(c *fiber.Ctx)error{
	db:= db.DBConn
	var items []models.GoItems
	db.Find(&items)
	return c.Status(fiber.StatusOK).JSON(items)
}

//GetItemByID pulls specific item by id
func GetItemByID(c *fiber.Ctx)error{
	id:= c.Params("id")
	db:=db.DBConn
	var item models.GoItems
	db.Find(&item,id)
	if item.Name==""{
		return c.Status(fiber.StatusBadRequest).JSON("could not find item with that id")
	}
	return c.Status(fiber.StatusOK).JSON(item)
}

//GetAllItemsByCategory gets all items specified by query parameter
func GetAllItemsByCategory(c *fiber.Ctx)error{
	category:=c.Params("category")
	db:=db.DBConn
	var items []models.GoItems
	db.Find(&items, "category = ?", category)
	if len(items)==0{
		c.Status(fiber.StatusBadRequest).JSON("found no items")
		
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

//GetAllItemsByBrand gets all items specified by query parameter
func GetAllItemsByBrand(c *fiber.Ctx)error{
	brand:= c.Params("brand")
	db:=db.DBConn
	var items []models.GoItems
	db.Find(&items, "brand = ?", brand)
	if len(items)==0{
	return 	c.Status(fiber.StatusBadRequest).JSON("found no items")
		
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

//AddItem adds an item to db
func AddItem(c *fiber.Ctx)error{
	db:= db.DBConn
	item:= new(models.GoItems)
	if err:= c.BodyParser(item); err!= nil{
		return	c.Status(fiber.StatusBadRequest).JSON(err)
		
	}
	db.Create(&item)
	return c.Status(fiber.StatusOK).JSON(item)
}

//DeleteItem deletes specified item
func DeleteItem(c *fiber.Ctx)error{
	id:=c.Params("id")
	db:=db.DBConn
	var item models.GoItems
	db.First(&item,id)
	if item.Name == "" {
		c.Status(500).JSON("no item found in that id")
	}
	db.Delete(&item)
	return c.Status(fiber.StatusOK).JSON(item)
	
}