package routes

import (
	"os"

	"github.com/darwin1224/go-echo-auth-crud/controllers"
	"github.com/darwin1224/go-echo-auth-crud/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAuthRoutes(c *echo.Echo, db *gorm.DB) {
	authController := &controllers.AuthController{Service: &services.UserService{Repo: db}}
	c.POST("/auth/login", authController.Login)
	c.GET("/credentials", authController.Credentials, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}
