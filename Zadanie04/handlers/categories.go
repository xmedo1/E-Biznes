package handlers

import (
	"net/http"
	"go-echo-api/models"
	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	Db.Create(&category)
	return c.JSON(http.StatusCreated, category)
}

func GetAllCategories(c echo.Context) error {
	var categories []models.Category
	Db.Preload("Products").Find(&categories)
	return c.JSON(http.StatusOK, categories)
}