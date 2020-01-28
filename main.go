package main

import (
	"github.com/darwin1224/go-echo-auth-crud/config"
	"github.com/darwin1224/go-echo-auth-crud/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time: ${time_rfc3339_nano}, method: ${method}, status: ${status}, uri: ${uri}\n",
	}))
	e.Use(middleware.Recover())
	routes.SetAuthRoutes(e, db)
	routes.SetUserRoutes(e, db)
	e.Logger.Fatal(e.Start(":8000"))
}
