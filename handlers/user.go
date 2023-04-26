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

type UserRepo struct {
	Db *gorm.DB
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewUser() *UserRepo {
	db := databases.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUserById(repository.Db, &user, id)
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
	return c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) DeleteUser(c echo.Context) error {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		)
		return err
	}
	return c.JSON(http.StatusOK, ErrorResponse{Code: http.StatusOK, Message: "User deleted successfully"})
}
