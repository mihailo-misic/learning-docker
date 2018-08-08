package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/mihailo-misic/learning-docker/01/models"
	"net/http"
	"github.com/mihailo-misic/learning-docker/01/res"
)

func Homers(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"PageTitle": "Yes"})
	c.HTML(http.StatusOK, "home.gohtml", gin.H{"PageTitle": "Yes"})
}

// [GET] all
func GetProducts(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", gin.H{"PageTitle": "Yes"})
	// c.JSON(http.StatusOK, res.Data(db.Find(&[]Product{})))
}

// [GET] one
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, res.Data(db.First(&Product{}, id)))
}

// [POST] create
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}

	db.Save(&product)

	if product.ID != 0 {
		c.JSON(http.StatusOK, res.Data(product))
		return
	}
	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not create the product", "An error occurred while creating the new product", nil}))
}

// [PUT] update
func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Product
	var newProd ResProduct

	// Find the product
	db.First(&prod, id)
	// Parse the request data
	if err := c.BindJSON(&newProd); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}
	// Update the product in the database
	db.Model(&prod).Updates(newProd)

	c.JSON(http.StatusOK, res.Data(prod))
}

// [DELETE] delete
func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Product

	db.Delete(&prod, id)

	if prod.ID != 0 {
		c.JSON(http.StatusOK, res.Data(prod))
	}

	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Unable to delete", "Error occurred while deleting the product", nil}))
}
