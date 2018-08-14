package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/mihailo-misic/learning-docker/models"
	"github.com/mihailo-misic/learning-docker/res"
	"net/http"
)

// [POST] create
func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}

	db.Save(&user)

	if user.ID != 0 {
		c.JSON(http.StatusOK, res.Data(user))
		return
	}
	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not create the user", "An error occurred while creating the new user", nil}))
}

// [GET] all
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, res.Data(db.Find(&[]User{})))
}

// [GET] one
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, res.Data(db.First(&User{}, id)))
}

// [PUT] update
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	var newUser ResUser

	// Find the user
	db.First(&user, id)
	// Parse the request data
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}
	// Update the user in the database
	db.Model(&user).Updates(newUser)

	c.JSON(http.StatusOK, res.Data(user))
}

// [DELETE] delete
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	db.Delete(&user, id)

	if user.ID != 0 {
		c.JSON(http.StatusOK, res.Data(user))
	}

	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Unable to delete", "Error occurred while deleting the user", nil}))
}
