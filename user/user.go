package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/warlock/webframerest/database"
)

// User struct
// swagger:model
type User struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Rating  int    `json:"rating"`
}

// GetUsers godoc
// @Summary List users
// @Description get users
// @Tags users
// @Produce  json
// @Router /users [get]
// @Success 200 {array} User
func GetUsers(c echo.Context) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, &users)
}

// GetUser godoc
// @Summary List user
// @Description get user
// @Tags users
// @Produce  json
// @ID get-string-by-int
// @Param id path int true "Account ID"
// @Router /user/{id} [get]
// @Success 200 {object} User
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

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "New User"
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
