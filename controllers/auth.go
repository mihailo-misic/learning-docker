package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	. "github.com/mihailo-misic/learning-docker/models"
	"github.com/mihailo-misic/learning-docker/res"
)

func FormRegister(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func FormLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{})
}

func Login(c *gin.Context) {
	var loginData FormLoginStruct
	c.Request.ParseForm()
	
	if err := c.Bind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}
	
	var user User
	ok := db.Where("email=? AND password=?", loginData.Email, loginData.Password).Find(&user)
	if ok.RowsAffected == 0 {
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		return
	}
	
	log.Println(user)
	c.SetCookie("token", user.Email,
		60*60*8, "/", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func IsAuth(c *gin.Context) {
	cook, _ := c.Cookie("token")
	log.Println(cook)
	if cook == "" {
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
	}
}

func Logout(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}
