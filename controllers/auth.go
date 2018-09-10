package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	. "github.com/mihailo-misic/learning-docker/models"
	"github.com/mihailo-misic/learning-docker/res"
	"log"
	"net/http"
)

func FormRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.gohtml", gin.H{})
}

func Register(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60 * 8,
		Secure:   false,
		HttpOnly: true,
	})

	var newUser FormUserStruct
	c.Request.ParseForm()
	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}

	ok := db.Create(newUser)
	if ok.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"User already exists", "", nil}))
		return
	}

	var user User
	ok = db.Where("email=? AND password=?", newUser.Email, newUser.Password).Find(&user)
	if ok.RowsAffected == 0 {
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		return
	}

	session.Set("user", user.Email)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func FormLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   60 * 60 * 8,
		Secure:   false,
		HttpOnly: true,
	})

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

	session.Set("user", user.Email)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func IsAuth(c *gin.Context) {
	cook := sessions.Default(c).Get("user")
	log.Println("User:", cook)
	if cook == nil {
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}
