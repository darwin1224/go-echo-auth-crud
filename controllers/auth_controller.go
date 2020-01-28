package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/darwin1224/go-echo-auth-crud/models"
	"github.com/darwin1224/go-echo-auth-crud/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthController struct {
		Service *services.UserService
	}
)

func (a *AuthController) Login(c echo.Context) error {
	userDto := new(models.User)
	if err := c.Bind(userDto); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Cannot parse json raw input"})
	}
	user, err := a.Service.GetUserByUsername(userDto)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Username or password is wrong"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Username or password is wrong"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["role_user"] = user.RoleUser
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Cannot generate token"})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func (a *AuthController) Credentials(c echo.Context) error {
	credentials := c.Get("user").(*jwt.Token)
	claims := credentials.Claims.(jwt.MapClaims)
	return c.JSON(http.StatusOK, echo.Map{"data": claims})
}
