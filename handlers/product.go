package controllers

import (
	"net/http"

	"errors"
	"strconv"

	"github.com/TakasBU/TakasBU/databases"
	"github.com/TakasBU/TakasBU/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}

func NewProduct() *ProductRepo {
	db := databases.InitDb()
	db.AutoMigrate(&models.Product{})
	return &ProductRepo{Db: db}
}

func (repository *ProductRepo) GetProductById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var Product models.Product
	err := models.GetProductById(repository.Db, &Product, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(
				http.StatusNotFound,
				ErrorResponse{Code: http.StatusNotFound, Message: ""},
			)
			return nil
		}
		c.JSON(
			http.StatusInternalServerError,
			ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		)
		return err
	}
	return c.JSON(http.StatusOK, Product)
}

func (repository *ProductRepo) DeleteProduct(c echo.Context) error {
	var Product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteProduct(repository.Db, &Product, id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		)
		return err
	}
	return c.JSON(http.StatusOK, ErrorResponse{Code: http.StatusOK, Message: "Product deleted successfully"})
}
