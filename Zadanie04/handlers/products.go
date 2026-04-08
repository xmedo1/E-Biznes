package handlers

import (
	"net/http"
	"gorm.io/gorm"
	"go-echo-api/models"
	"github.com/labstack/echo/v4"
)

var Db *gorm.DB

func GetAllProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var p models.Product
	if err := Db.First(&p, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
	}
	return c.JSON(http.StatusOK, p)
}

func CreateProduct(c echo.Context) error {
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	Db.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var p models.Product

	if err := Db.First(&p, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
	}

	if err := c.Bind(&p); err != nil {
		return err
	}
	Db.Save(&p)
	return c.JSON(http.StatusOK, p)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	Db.Delete(&models.Product{}, id)
	return c.NoContent(http.StatusNoContent)
}