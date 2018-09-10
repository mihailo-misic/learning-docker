package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/mihailo-misic/learning-docker/models"
	"github.com/mihailo-misic/learning-docker/res"
	"log"
	"net/http"
)

// [GET] all
func GetUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "users.gohtml", gin.H{
		"PageTitle": "Users",
		"Users":     db.Find(&[]User{}).Value,
	})
}

// [GET] one
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, res.Data(db.First(&User{}, id)))
}

// [POST] create
func CreateUser(c *gin.Context) {
	var user FormUserStruct
	c.Request.ParseForm()

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}

	ok := db.Save(&user)

	if ok.RowsAffected > 0 {
		c.Redirect(http.StatusMovedPermanently, "/users")
		return
	}
	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not create the user", "An error occurred while creating the new user", nil}))
}

// [PUT] update
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	var newUser ResUser

	// TODO Implement file uploading
	c.Request.ParseForm()

	// Find the user
	db.First(&user, id)
	// Parse the request data
	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}
	// Update the user in the database
	db.Model(&user).Updates(newUser)

	c.Redirect(http.StatusMovedPermanently, "/users")
}

// [DELETE] delete
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	ok := db.Delete(&user, id)

	log.Println(ok.RowsAffected)

	if ok.RowsAffected > 0 {
		c.Redirect(http.StatusMovedPermanently, "/users")
		return
	}

	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Unable to delete", "Error occurred while deleting the user", nil}))
}

// [GET] Form
func FormUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user interface{}
	if id != "" {
		user = db.First(&User{}, id).Value
	}

	c.HTML(http.StatusOK, "users_form.gohtml", gin.H{
		"PageTitle": "Users",
		"User":      user,
	})
}
