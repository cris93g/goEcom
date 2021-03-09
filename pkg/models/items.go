package models

import "github.com/jinzhu/gorm"

type GoItems struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Picture     string  `json:"picture"`
	Category    string  `json:"category"`
}