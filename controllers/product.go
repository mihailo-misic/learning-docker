package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/mihailo-misic/learning-docker/models"
	"github.com/mihailo-misic/learning-docker/res"
	"log"
	"net/http"
)

// [GET] all
func GetProducts(c *gin.Context) {
	c.HTML(http.StatusOK, "products.gohtml", gin.H{
		"PageTitle": "Products",
		"Products":  db.Find(&[]Product{}).Value,
	})
}

// [GET] one
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, res.Data(db.First(&Product{}, id)))
}

// [POST] create
func CreateProduct(c *gin.Context) {
	var product FormProductStruct
	c.Request.ParseForm()

	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}

	ok := db.Save(&product)
	if ok.RowsAffected > 0 {
		c.Redirect(http.StatusMovedPermanently, "/products")
		return
	}
	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not create the product", "An error occurred while creating the new product", nil}))
}

// [PUT] update
func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Product
	var newProd ResProduct

	c.Request.ParseForm()

	// Find the product
	db.First(&prod, id)
	// Parse the request data
	if err := c.Bind(&newProd); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not parse the data", "Error occurred while parsing the data", err}))
		return
	}
	// Update the product in the database
	db.Model(&prod).Updates(newProd)

	c.Redirect(http.StatusMovedPermanently, "/products")
}

// [DELETE] delete
func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Product

	ok := db.Delete(&prod, id)

	log.Println(ok.RowsAffected)

	if ok.RowsAffected > 0 {
		c.Redirect(http.StatusMovedPermanently, "/products")
		return
	}

	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Unable to delete", "Error occurred while deleting the product", nil}))
}

// [GET] Form
func FormProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product interface{}
	if id != "" {
		product = db.First(&Product{}, id).Value
	}

	c.HTML(http.StatusOK, "products_form.gohtml", gin.H{
		"PageTitle": "Products",
		"Product":   product,
	})
}
