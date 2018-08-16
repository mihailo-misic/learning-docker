package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "github.com/mihailo-misic/learning-docker/controllers"
	"github.com/mihailo-misic/learning-docker/database"
	m "github.com/mihailo-misic/learning-docker/models"
	method "github.com/bu/gin-method-override"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(method.ProcessMethodOverride(r))

	r.LoadHTMLGlob("views/*")

	r.GET("/", c.GetProducts)

	r.GET("/products", c.GetProducts)
	r.GET("/products/create", c.ProductForm)
	r.GET("/product/:id/edit", c.ProductForm)
	r.GET("/product/:id", c.GetProduct)
	r.POST("/products", c.CreateProduct)
	r.PUT("/product/:id", c.UpdateProduct)
	r.GET("/product/:id/delete", c.DeleteProduct)

	return r
}

func main() {
	// Setup database
	db := database.Init()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	db.AutoMigrate(&m.Product{})

	r := setupRouter()

	log.Fatal(r.Run(":8000"))
}
