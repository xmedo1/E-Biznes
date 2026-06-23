package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Products []Product `json:"products" gorm:"many2many:cart_products;"`
}