package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "github.com/mihailo-misic/learning-docker/01/controllers"
	"github.com/mihailo-misic/learning-docker/01/database"
	m "github.com/mihailo-misic/learning-docker/01/models"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/home", c.Homers)
	r.GET("/", c.GetProducts)
	r.GET("/products/:id", c.GetProduct)
	r.POST("/products", c.CreateProduct)
	r.PUT("/products/:id", c.UpdateProduct)
	r.DELETE("/products/:id", c.DeleteProduct)

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
