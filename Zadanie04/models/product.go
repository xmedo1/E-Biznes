package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	CategoryID uint `json:"category_id"`
}

func PriceGreaterThan(minPrice float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price > ?", minPrice)
	}
}

func OrderByPriceDesc(db *gorm.DB) *gorm.DB {
	return db.Order("price desc")
}