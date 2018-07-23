package main

import (
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	c "github.com/mihailo-misic/learning-docker/01/controllers"
	m "github.com/mihailo-misic/learning-docker/01/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", c.GetProducts)
	// Get one
	// Create
	// Update
	// Delete

	return r
}

func main() {
	// Setup database
	// db := database.Init()
	db, err := gorm.Open("postgres", `
		host=postgres
		port=5432
		user=mygo
		dbname=crm
		password=secret
		sslmode=disable
	`)
	if err != nil {
		log.Fatal(err)
	}
	c.ConnectControllers(db)
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	db.AutoMigrate(&m.Product{})

	r := setupRouter()

	log.Fatal(r.Run(":8000"))
}
