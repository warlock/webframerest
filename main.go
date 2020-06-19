package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/warlock/webframerest/database"
	"github.com/warlock/webframerest/user"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/users", user.GetUsers)
	e.GET("/user/:id", user.GetUser)
	e.POST("/user", user.CreateUser)
	e.DELETE("/user/:id", user.DeleteUser)
	e.PUT("/user/:id", user.UpdateUser)
}

func initDatabse() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database migrated")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	initDatabse()
	defer database.DBConn.Close()

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
