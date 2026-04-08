package main

import (
	"go-echo-api/handlers"
	"go-echo-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Blad - baza danych")
	}

	db.AutoMigrate(&models.Product{})
	handlers.Db = db

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	g := e.Group("/products")
	g.GET("", handlers.GetAllProducts)
	g.GET("/:id", handlers.GetProduct)
	g.POST("", handlers.CreateProduct)
	g.PUT("/:id", handlers.UpdateProduct)
	g.DELETE("/:id", handlers.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}