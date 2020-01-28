package controllers

import (
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/darwin1224/go-echo-auth-crud/models"
	"github.com/darwin1224/go-echo-auth-crud/services"
	"github.com/labstack/echo"
)

type (
	UserController struct {
		Service *services.UserService
	}
)

func (u *UserController) Index(c echo.Context) error {
	users, err := u.Service.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": users})
}

func (u *UserController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user, err := u.Service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (u *UserController) Store(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user.Password = string(hash)
	if err := u.Service.InsertUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"data": user})
}

func (u *UserController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user, err := u.Service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	c.Bind(user)
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user.Password = string(hash)
	if err := u.Service.UpdateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (u *UserController) Destroy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user, err := u.Service.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	if err := u.Service.DeleteUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
