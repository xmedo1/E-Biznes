package main

import (
	"go-echo-api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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