package routes

import (
	"os"

	"github.com/darwin1224/go-echo-auth-crud/middlewares"

	"github.com/darwin1224/go-echo-auth-crud/controllers"
	"github.com/darwin1224/go-echo-auth-crud/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetUserRoutes(c *echo.Echo, db *gorm.DB) {
	userController := &controllers.UserController{Service: &services.UserService{Repo: db}}
	c.GET("/users", userController.Index, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.Authorize("ADMIN"))
	c.GET("/users/:id", userController.Show, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.Authorize("ADMIN"))
	c.POST("/users", userController.Store, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.Authorize("ADMIN"))
	c.PUT("/users/:id", userController.Update, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.Authorize("ADMIN"))
	c.DELETE("/users/:id", userController.Destroy, middleware.JWT([]byte(os.Getenv("JWT_SECRET"))), middlewares.Authorize("ADMIN"))
}
