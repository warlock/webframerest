package main

// @title Gorm Api play
// @version 1.0
// @description Investigant...

// @contact.url https://js.gl
// @contact.email js@js.gl

// @host localhost:1323
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/warlock/webframerest/database"
	_ "github.com/warlock/webframerest/docs"
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
	database.DBConn, err = gorm.Open("sqlite3", "/tmp/db.db")
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
