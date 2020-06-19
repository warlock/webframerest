package user

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/warlock/webframerest/database"
)

// User ...
type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Rating  int    `json:"rating"`
}

// GetUsers ...
func GetUsers(c echo.Context) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, &users)
}

// GetUser ...
func GetUser(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn
	var user User
	db.First(&user, id)
	if user.Name == "" {
		return c.String(http.StatusNotFound, "Not found")
	}
	return c.JSON(http.StatusOK, &user)
}

// CreateUser ...
func CreateUser(c echo.Context) error {
	db := database.DBConn
	var user User
	user.Name = c.FormValue("name")
	user.Surname = c.FormValue("surname")
	i, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	user.Rating = i
	db.Create(&user)
	return c.JSON(http.StatusOK, &user)
}

// UpdateUser ...
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn
	var user User
	db.First(&user, id)
	if user.Name == "" {
		return c.String(http.StatusNotFound, "Not found")
	}
	user.Name = c.FormValue("name")
	user.Surname = c.FormValue("surname")
	i, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	user.Rating = i
	db.Save(&user)
	return c.String(http.StatusOK, "User successfully updated")
}

// DeleteUser ...
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn
	var user User
	db.First(&user, id)
	if user.Name == "" {
		return c.String(http.StatusNotFound, "Not found")
	}
	db.Delete(&user)
	return c.String(http.StatusOK, "User successfully deleted")
}
