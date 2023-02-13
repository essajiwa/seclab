package api

import (
	"database/sql"
	"net/http"
	"seclab/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IService interface {
	Login(username, password string) (string, error)
	FindProductByCategory(category string) (*[]model.Product, error)
	FindUserByID(id int) (model.User, error)
}

type Handler struct {
	service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) FindUserByID(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
		return err
	}

	user, err := h.service.FindUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "User not found")
		} else {
			c.Error(err)
		}
		return err
	}
	c.JSON(http.StatusOK, user)
	return err
}

// Login handles login request
func (h *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := h.service.Login(username, password)
	if err != nil {

		if err == sql.ErrNoRows {
			c.String(http.StatusUnauthorized, "Invalid username or password")
		} else {
			c.Error(err)
		}

		return err
	}

	c.String(http.StatusOK, "Token: "+token)
	return err
}

// FindProductByCategory handles find product by category request
func (h *Handler) FindProductByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	product, err := h.service.FindProductByCategory(category)
	if err != nil {
		if err == sql.ErrNoRows {
			c.String(http.StatusNotFound, "Product not found")
		} else {
			c.Error(err)
		}

		return err
	}
	c.JSON(http.StatusOK, product)
	return nil
}
