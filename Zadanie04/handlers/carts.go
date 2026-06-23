package handlers

import (
	"net/http"
	"go-echo-api/models"
	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {
	id := c.Param("id")
	var cart models.Cart
	if err := Db.Preload("Products").First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Koszyk nie istnieje")
	}
	return c.JSON(http.StatusOK, cart)
}

func CreateCart(c echo.Context) error {
	cart := new(models.Cart)
	Db.Create(&cart)
	return c.JSON(http.StatusCreated, cart)
}

func AddProductToCart(c echo.Context) error {
	cartID := c.Param("cartId")
	productID := c.Param("productId")

	var cart models.Cart
	if err := Db.First(&cart, cartID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Koszyk nie istnieje")
	}

	var product models.Product
	if err := Db.First(&product, productID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
	}
	
	Db.Model(&cart).Association("Products").Append(&product)
	return c.JSON(http.StatusOK, map[string]string{"message": "Dodano produkt do koszyka"})
}