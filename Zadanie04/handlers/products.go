package handlers

import (
	"net/http"
	"strconv"
	"go-echo-api/models"
	"github.com/labstack/echo/v4"
)

var products = []models.Product{
	{ID: 1, Name: "Banan", Price: 1.5},
}

func GetAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, p := range products {
		if p.ID == id {
			return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
}

func CreateProduct(c echo.Context) error {
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	p.ID = len(products) + 1
	products = append(products, *p)
	return c.JSON(http.StatusCreated, p)
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	for i, p := range products {
		if p.ID == id {
			updatedProduct := new(models.Product)
			if err := c.Bind(updatedProduct); err != nil {
				return err
			}
			updatedProduct.ID = id
			products[i] = *updatedProduct
			return c.JSON(http.StatusOK, updatedProduct)
		}
	}
	return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Produkt nie istnieje")
}